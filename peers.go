package cache

import pb "cache/geecachepb"
type PeerPicker interface{
	PickPeer(key string) (peer PeerGetter,ok bool)
}

type PeerGetter interface {
	Get(in *pb.Request,out *pb.Response) error
	//Get(group string,key string)([]byte,error)
}


