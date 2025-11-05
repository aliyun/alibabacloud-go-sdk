// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskReplicaGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *ModifyDiskReplicaGroupRequest
	GetBandwidth() *int64
	SetClientToken(v string) *ModifyDiskReplicaGroupRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyDiskReplicaGroupRequest
	GetDescription() *string
	SetEnableRtc(v bool) *ModifyDiskReplicaGroupRequest
	GetEnableRtc() *bool
	SetGroupName(v string) *ModifyDiskReplicaGroupRequest
	GetGroupName() *string
	SetRPO(v int64) *ModifyDiskReplicaGroupRequest
	GetRPO() *int64
	SetRegionId(v string) *ModifyDiskReplicaGroupRequest
	GetRegionId() *string
	SetReplicaGroupId(v string) *ModifyDiskReplicaGroupRequest
	GetReplicaGroupId() *string
}

type ModifyDiskReplicaGroupRequest struct {
	// The bandwidth value. Unit: Kbit/s.
	//
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// -
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the replication pair-consistent group. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether to enable replication time control.
	//
	// example:
	//
	// true
	EnableRtc *bool `json:"EnableRtc,omitempty" xml:"EnableRtc,omitempty"`
	// The name of the replication pair-consistent group. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// myreplicagrouptest
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The RPO of the replication pair-consistent group. Unit: seconds. Valid value: 900.
	//
	// example:
	//
	// 900
	RPO *int64 `json:"RPO,omitempty" xml:"RPO,omitempty"`
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

func (s ModifyDiskReplicaGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaGroupRequest) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *ModifyDiskReplicaGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDiskReplicaGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDiskReplicaGroupRequest) GetEnableRtc() *bool {
	return s.EnableRtc
}

func (s *ModifyDiskReplicaGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyDiskReplicaGroupRequest) GetRPO() *int64 {
	return s.RPO
}

func (s *ModifyDiskReplicaGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDiskReplicaGroupRequest) GetReplicaGroupId() *string {
	return s.ReplicaGroupId
}

func (s *ModifyDiskReplicaGroupRequest) SetBandwidth(v int64) *ModifyDiskReplicaGroupRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetClientToken(v string) *ModifyDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetDescription(v string) *ModifyDiskReplicaGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetEnableRtc(v bool) *ModifyDiskReplicaGroupRequest {
	s.EnableRtc = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetGroupName(v string) *ModifyDiskReplicaGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetRPO(v int64) *ModifyDiskReplicaGroupRequest {
	s.RPO = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetRegionId(v string) *ModifyDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) SetReplicaGroupId(v string) *ModifyDiskReplicaGroupRequest {
	s.ReplicaGroupId = &v
	return s
}

func (s *ModifyDiskReplicaGroupRequest) Validate() error {
	return dara.Validate(s)
}
