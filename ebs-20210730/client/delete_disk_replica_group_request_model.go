// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiskReplicaGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteDiskReplicaGroupRequest
	GetClientToken() *string
	SetRegionId(v string) *DeleteDiskReplicaGroupRequest
	GetRegionId() *string
	SetReplicaGroupId(v string) *DeleteDiskReplicaGroupRequest
	GetReplicaGroupId() *string
}

type DeleteDiskReplicaGroupRequest struct {
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

func (s DeleteDiskReplicaGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteDiskReplicaGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDiskReplicaGroupRequest) GetReplicaGroupId() *string {
	return s.ReplicaGroupId
}

func (s *DeleteDiskReplicaGroupRequest) SetClientToken(v string) *DeleteDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDiskReplicaGroupRequest) SetRegionId(v string) *DeleteDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDiskReplicaGroupRequest) SetReplicaGroupId(v string) *DeleteDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

func (s *DeleteDiskReplicaGroupRequest) Validate() error {
	return dara.Validate(s)
}
