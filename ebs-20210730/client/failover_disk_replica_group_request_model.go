// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailoverDiskReplicaGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *FailoverDiskReplicaGroupRequest
	GetClientToken() *string
	SetRegionId(v string) *FailoverDiskReplicaGroupRequest
	GetRegionId() *string
	SetReplicaGroupId(v string) *FailoverDiskReplicaGroupRequest
	GetReplicaGroupId() *string
}

type FailoverDiskReplicaGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the secondary site of the replication pair-consistent group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair-consistent group.
	//
	// This parameter is required.
	//
	// example:
	//
	// group-myreplica****
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
}

func (s FailoverDiskReplicaGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s FailoverDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *FailoverDiskReplicaGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *FailoverDiskReplicaGroupRequest) GetReplicaGroupId() *string {
	return s.ReplicaGroupId
}

func (s *FailoverDiskReplicaGroupRequest) SetClientToken(v string) *FailoverDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *FailoverDiskReplicaGroupRequest) SetRegionId(v string) *FailoverDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *FailoverDiskReplicaGroupRequest) SetReplicaGroupId(v string) *FailoverDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

func (s *FailoverDiskReplicaGroupRequest) Validate() error {
	return dara.Validate(s)
}
