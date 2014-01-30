package cluster

import (
	"fmt"
	"strconv"
	"strings"
	"testing" //import go package for testing related functionality
	"time"
)

func GetMsg(server *Raftserver, mail chan string) {
	//	inter:= <-server.Inbox()
	mail <- (<-server.Inbox()).Msg
}

 /* This is fine :::
func Test_BasicTwoServer(t *testing.T){
	total:=2
	var chans [2]chan string
	var servers [2]*Raftserver
	var output [2]string
        for i:=0;i<total;i++ {
           chans[i] = make(chan string,1)
           servers[i]=New("configBASIC.json",i+1)
        }
        servers[0].Outbox()  <- &Envelope{Pid:-1,MsgId:0, Msg: "Server ONE"}
        servers[1].Outbox()  <- &Envelope{Pid:-1,MsgId:0, Msg: "Server TWO"}
	for i := range chans {
		go GetMsg(servers[i],chans[i])
	}
	for i:=0;i<2;i++{
		output[i]= <-chans[i]
	}

	fmt.Println("In test B")
	if (output[0]!="Server TWO"){
		t.Error("Error 1")
		fmt.Println("Err1:",output[1])
	}else{
		fmt.Println("Test : fine")
	}
	if(output[1]!="Server ONE"){
		t.Error("Error 2")
		fmt.Println("Err2:",output[1])
	}else{
		fmt.Println("Test 2 fine")
	}
//	server1.Wait()
//	server2.Wait()
}*/

func Test_TenServer(t *testing.T) {
	total := 10
	var servers [80]*Raftserver
	for i:= 0; i < total; i++ {
		servers[i] = New("config10.json", i+1)
	}
	totMsg:=0
	for i := 0; i < total; i++ {
		time.Sleep(2 * time.Millisecond)
		smsg := string("Server|" + strconv.Itoa(i+1) + "|says Hi")
		servers[i].Outbox() <- &Envelope{Pid: -1, MsgId: 0, Msg: smsg}
		//fmt.Println("Send msg for server :",i+1)
		for j := 0; j < total; j++ { //you will get total-1 msg ..
		//	fmt.Println("Waiting for channel")
			if(i!=j){
				opt := <-servers[j].Inbox() 
				temp := strings.Split(opt.Msg, "|")
//				fmt.Println("i: ",i,"Sender detected: ", temp[1],"Receiver: ",j+1)
				senderPid, _ := strconv.Atoi(temp[1])
				if(senderPid == i){
					t.Error("Received own message !!")
				}
				totMsg++
			}
		}
		if(totMsg!=(total-1)){
			t.Error("Message dropped in between")
		}
		totMsg=0
	}
}


func Test_TwoHundredServer(t *testing.T) {
        total := 200
        var servers [200]*Raftserver
        for i:= 0; i < total; i++ {
                servers[i] = New("configSevHundred.json", i+1)
        }
        totMsg:=0
        for i := 0; i < total; i++ {
	//	fmt.Println("Good Going")
                time.Sleep(25 * time.Millisecond)
                smsg := string("Server|" + strconv.Itoa(i+1) + "|says Hi")
                servers[i].Outbox() <- &Envelope{Pid: -1, MsgId: 0, Msg: smsg}
                //fmt.Println("Send msg for server :",i+1)
                for j := 0; j < total; j++ { //you will get total-1 msg ..
                //      fmt.Println("Waiting for channel")
                        if(i!=j){
                                opt := <-servers[j].Inbox()
                                temp := strings.Split(opt.Msg, "|")
                              fmt.Println("i: ",i,"Sender detected: ", temp[1],"Receiver: ",j+1)
                                senderPid, _ := strconv.Atoi(temp[1])
                                if(senderPid == i){
                                        t.Error("Received own message !!")
                                }
                                totMsg++
                        }
                }
                if(totMsg!=(total-1)){
                        t.Error("Message dropped in between")
                }
                totMsg=0
        }
}       

