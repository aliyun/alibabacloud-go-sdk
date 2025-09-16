// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySearcherReplicaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPartition(v int32) *ModifySearcherReplicaRequest
	GetPartition() *int32
	SetReplica(v int32) *ModifySearcherReplicaRequest
	GetReplica() *int32
}

type ModifySearcherReplicaRequest struct {
	// example:
	//
	// 2
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
	// example:
	//
	// 2
	Replica *int32 `json:"replica,omitempty" xml:"replica,omitempty"`
}

func (s ModifySearcherReplicaRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySearcherReplicaRequest) GoString() string {
	return s.String()
}

func (s *ModifySearcherReplicaRequest) GetPartition() *int32 {
	return s.Partition
}

func (s *ModifySearcherReplicaRequest) GetReplica() *int32 {
	return s.Replica
}

func (s *ModifySearcherReplicaRequest) SetPartition(v int32) *ModifySearcherReplicaRequest {
	s.Partition = &v
	return s
}

func (s *ModifySearcherReplicaRequest) SetReplica(v int32) *ModifySearcherReplicaRequest {
	s.Replica = &v
	return s
}

func (s *ModifySearcherReplicaRequest) Validate() error {
	return dara.Validate(s)
}
