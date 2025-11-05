// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDiskReplicaGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartDiskReplicaGroupRequest
	GetClientToken() *string
	SetOneShot(v bool) *StartDiskReplicaGroupRequest
	GetOneShot() *bool
	SetRegionId(v string) *StartDiskReplicaGroupRequest
	GetRegionId() *string
	SetReplicaGroupId(v string) *StartDiskReplicaGroupRequest
	GetReplicaGroupId() *string
}

type StartDiskReplicaGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to immediately synchronize data once. Valid values:
	//
	// 	- true: immediately synchronizes data once.
	//
	// 	- false: synchronizes data based on the RPO of the replication pair-consistent group.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	OneShot *bool `json:"OneShot,omitempty" xml:"OneShot,omitempty"`
	// The ID of the replication pair-consistent group.
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

func (s StartDiskReplicaGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartDiskReplicaGroupRequest) GetOneShot() *bool {
	return s.OneShot
}

func (s *StartDiskReplicaGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartDiskReplicaGroupRequest) GetReplicaGroupId() *string {
	return s.ReplicaGroupId
}

func (s *StartDiskReplicaGroupRequest) SetClientToken(v string) *StartDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDiskReplicaGroupRequest) SetOneShot(v bool) *StartDiskReplicaGroupRequest {
	s.OneShot = &v
	return s
}

func (s *StartDiskReplicaGroupRequest) SetRegionId(v string) *StartDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *StartDiskReplicaGroupRequest) SetReplicaGroupId(v string) *StartDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

func (s *StartDiskReplicaGroupRequest) Validate() error {
	return dara.Validate(s)
}
