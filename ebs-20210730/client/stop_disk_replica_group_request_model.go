// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDiskReplicaGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StopDiskReplicaGroupRequest
	GetClientToken() *string
	SetRegionId(v string) *StopDiskReplicaGroupRequest
	GetRegionId() *string
	SetReplicaGroupId(v string) *StopDiskReplicaGroupRequest
	GetReplicaGroupId() *string
}

type StopDiskReplicaGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the replication pair-consistent group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the replication pair-consistent group. You can call the [DescribeDiskReplicaGroups](https://help.aliyun.com/document_detail/426614.html) operation to query the IDs of replication pair-consistent groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-myreplica****
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
}

func (s StopDiskReplicaGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StopDiskReplicaGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopDiskReplicaGroupRequest) GetReplicaGroupId() *string {
	return s.ReplicaGroupId
}

func (s *StopDiskReplicaGroupRequest) SetClientToken(v string) *StopDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *StopDiskReplicaGroupRequest) SetRegionId(v string) *StopDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *StopDiskReplicaGroupRequest) SetReplicaGroupId(v string) *StopDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

func (s *StopDiskReplicaGroupRequest) Validate() error {
	return dara.Validate(s)
}
