// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDiskReplicaPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RemoveDiskReplicaPairRequest
	GetClientToken() *string
	SetRegionId(v string) *RemoveDiskReplicaPairRequest
	GetRegionId() *string
	SetReplicaGroupId(v string) *RemoveDiskReplicaPairRequest
	GetReplicaGroupId() *string
	SetReplicaPairId(v string) *RemoveDiskReplicaPairRequest
	GetReplicaPairId() *string
}

type RemoveDiskReplicaPairRequest struct {
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
	// The ID of the replication pair-consistent group.
	//
	// You can call the [DescribeDiskReplicaGroups](https://help.aliyun.com/document_detail/426614.html) operation to query the IDs of replication pair-consistent groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-myreplica****
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	// The ID of the replication pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// pair-cn-dsa****
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s RemoveDiskReplicaPairRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *RemoveDiskReplicaPairRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveDiskReplicaPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveDiskReplicaPairRequest) GetReplicaGroupId() *string {
	return s.ReplicaGroupId
}

func (s *RemoveDiskReplicaPairRequest) GetReplicaPairId() *string {
	return s.ReplicaPairId
}

func (s *RemoveDiskReplicaPairRequest) SetClientToken(v string) *RemoveDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveDiskReplicaPairRequest) SetRegionId(v string) *RemoveDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveDiskReplicaPairRequest) SetReplicaGroupId(v string) *RemoveDiskReplicaPairRequest {
	s.ReplicaGroupId = &v
	return s
}

func (s *RemoveDiskReplicaPairRequest) SetReplicaPairId(v string) *RemoveDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

func (s *RemoveDiskReplicaPairRequest) Validate() error {
	return dara.Validate(s)
}
