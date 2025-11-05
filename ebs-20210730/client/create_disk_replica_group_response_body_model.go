// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiskReplicaGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReplicaGroupId(v string) *CreateDiskReplicaGroupResponseBody
	GetReplicaGroupId() *string
	SetRequestId(v string) *CreateDiskReplicaGroupResponseBody
	GetRequestId() *string
}

type CreateDiskReplicaGroupResponseBody struct {
	// The ID of the replication pair-consistent group.
	//
	// example:
	//
	// pg-xxxxxxx
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C123F94F-4E38-19AE-942A-A8D6F44F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDiskReplicaGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskReplicaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaGroupResponseBody) GetReplicaGroupId() *string {
	return s.ReplicaGroupId
}

func (s *CreateDiskReplicaGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDiskReplicaGroupResponseBody) SetReplicaGroupId(v string) *CreateDiskReplicaGroupResponseBody {
	s.ReplicaGroupId = &v
	return s
}

func (s *CreateDiskReplicaGroupResponseBody) SetRequestId(v string) *CreateDiskReplicaGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDiskReplicaGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
