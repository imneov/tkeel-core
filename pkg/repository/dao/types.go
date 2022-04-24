package dao

import (
	"context"

	"go.etcd.io/etcd/api/v3/mvccpb"
)

var (
	PUT    EnventType = EnventType(mvccpb.PUT)
	DELETE EnventType = EnventType(mvccpb.DELETE)
)

type EnventType mvccpb.Event_EventType

func (et EnventType) String() string {
	return mvccpb.Event_EventType(et).String()
}

type DecodeFunc func([]byte) (Resource, error)
type RangeResourceFunc func([]*mvccpb.KeyValue)
type WatchResourceFunc func(EnventType, *mvccpb.KeyValue)

type Resource interface {
	Codec() KVCodec
}

type IDao interface {
	Close()
	GetLastRevision(ctx context.Context) int64
	// resource etcd interfaces.
	PutResource(ctx context.Context, res Resource) error
	GetResource(ctx context.Context, res Resource) (Resource, error)
	DelResource(ctx context.Context, res Resource) error
	DelResources(ctx context.Context, prefix string) error
	HasResource(ctx context.Context, res Resource) (has bool, err error)
	ListResource(ctx context.Context, rev int64, prefix string, decodeFunc DecodeFunc) ([]Resource, error)
	RangeResource(ctx context.Context, rev int64, prefix string, handler RangeResourceFunc)
	WatchResource(ctx context.Context, rev int64, prefix string, handler WatchResourceFunc)

	// resource store interfaces.
	StoreResource(ctx context.Context, res Resource) error
	GetStoreResource(ctx context.Context, res Resource) (Resource, error)
	RemoveStoreResource(ctx context.Context, res Resource) error
}
