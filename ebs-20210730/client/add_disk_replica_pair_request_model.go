// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDiskReplicaPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddDiskReplicaPairRequest
	GetClientToken() *string
	SetRegionId(v string) *AddDiskReplicaPairRequest
	GetRegionId() *string
	SetReplicaGroupId(v string) *AddDiskReplicaPairRequest
	GetReplicaGroupId() *string
	SetReplicaPairId(v string) *AddDiskReplicaPairRequest
	GetReplicaPairId() *string
}

type AddDiskReplicaPairRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// pg-myreplica****
	ReplicaGroupId *string `json:"ReplicaGroupId,omitempty" xml:"ReplicaGroupId,omitempty"`
	// The ID of the replication pair. You can call the [DescribeDiskReplicaPairs](https://help.aliyun.com/document_detail/354206.html) operation to query the IDs of existing replication pairs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pair-cn-dsa****
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s AddDiskReplicaPairRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *AddDiskReplicaPairRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddDiskReplicaPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddDiskReplicaPairRequest) GetReplicaGroupId() *string {
	return s.ReplicaGroupId
}

func (s *AddDiskReplicaPairRequest) GetReplicaPairId() *string {
	return s.ReplicaPairId
}

func (s *AddDiskReplicaPairRequest) SetClientToken(v string) *AddDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *AddDiskReplicaPairRequest) SetRegionId(v string) *AddDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *AddDiskReplicaPairRequest) SetReplicaGroupId(v string) *AddDiskReplicaPairRequest {
	s.ReplicaGroupId = &v
	return s
}

func (s *AddDiskReplicaPairRequest) SetReplicaPairId(v string) *AddDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

func (s *AddDiskReplicaPairRequest) Validate() error {
	return dara.Validate(s)
}
