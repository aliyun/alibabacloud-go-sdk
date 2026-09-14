// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiskReplicaGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int64) *CreateDiskReplicaGroupRequest
	GetBandwidth() *int64
	SetClientToken(v string) *CreateDiskReplicaGroupRequest
	GetClientToken() *string
	SetDescription(v string) *CreateDiskReplicaGroupRequest
	GetDescription() *string
	SetDestinationRegionId(v string) *CreateDiskReplicaGroupRequest
	GetDestinationRegionId() *string
	SetDestinationZoneId(v string) *CreateDiskReplicaGroupRequest
	GetDestinationZoneId() *string
	SetEnableRtc(v bool) *CreateDiskReplicaGroupRequest
	GetEnableRtc() *bool
	SetGroupName(v string) *CreateDiskReplicaGroupRequest
	GetGroupName() *string
	SetRPO(v int64) *CreateDiskReplicaGroupRequest
	GetRPO() *int64
	SetRegionId(v string) *CreateDiskReplicaGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDiskReplicaGroupRequest
	GetResourceGroupId() *string
	SetSourceZoneId(v string) *CreateDiskReplicaGroupRequest
	GetSourceZoneId() *string
	SetTag(v []*CreateDiskReplicaGroupRequestTag) *CreateDiskReplicaGroupRequest
	GetTag() []*CreateDiskReplicaGroupRequestTag
}

type CreateDiskReplicaGroupRequest struct {
	// The bandwidth in Kbps.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// 5
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// A client token to ensure the idempotence of the request. Generate a unique value from your client for this parameter. The \\`ClientToken\\` parameter value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
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
	// The ID of the region where the disaster recovery site is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// The ID of the zone where the disaster recovery site is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai-e
	DestinationZoneId *string `json:"DestinationZoneId,omitempty" xml:"DestinationZoneId,omitempty"`
	// Specifies whether to enable replication time control (RTC). Valid values:
	//
	// - false: Disable RTC.
	//
	// - true: Enable RTC.
	//
	// Default value: false.
	//
	// > If you set this parameter to true, RTC is enabled for the replication pair-consistent group. RTC is also enabled for all asynchronous replication pairs that are added to the group.
	//
	// example:
	//
	// true
	EnableRtc *bool `json:"EnableRtc,omitempty" xml:"EnableRtc,omitempty"`
	// The name of the replication pair-consistent group. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character, and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// myreplicagrouptest
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The recovery point objective (RPO) of the replication pair-consistent group, in seconds. The only supported value is 900.
	//
	// example:
	//
	// 900
	RPO *int64 `json:"RPO,omitempty" xml:"RPO,omitempty"`
	// The ID of the region where the replication pair-consistent group resides. This is the same as the region of the production site.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the replication pair-consistent group belongs.
	//
	// example:
	//
	// rg-acfmvs*******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the zone where the production site is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-f
	SourceZoneId *string `json:"SourceZoneId,omitempty" xml:"SourceZoneId,omitempty"`
	// The tags to add to the resource. You can add up to 20 tags.
	Tag []*CreateDiskReplicaGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateDiskReplicaGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskReplicaGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaGroupRequest) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *CreateDiskReplicaGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDiskReplicaGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDiskReplicaGroupRequest) GetDestinationRegionId() *string {
	return s.DestinationRegionId
}

func (s *CreateDiskReplicaGroupRequest) GetDestinationZoneId() *string {
	return s.DestinationZoneId
}

func (s *CreateDiskReplicaGroupRequest) GetEnableRtc() *bool {
	return s.EnableRtc
}

func (s *CreateDiskReplicaGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateDiskReplicaGroupRequest) GetRPO() *int64 {
	return s.RPO
}

func (s *CreateDiskReplicaGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDiskReplicaGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDiskReplicaGroupRequest) GetSourceZoneId() *string {
	return s.SourceZoneId
}

func (s *CreateDiskReplicaGroupRequest) GetTag() []*CreateDiskReplicaGroupRequestTag {
	return s.Tag
}

func (s *CreateDiskReplicaGroupRequest) SetBandwidth(v int64) *CreateDiskReplicaGroupRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetClientToken(v string) *CreateDiskReplicaGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetDescription(v string) *CreateDiskReplicaGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetDestinationRegionId(v string) *CreateDiskReplicaGroupRequest {
	s.DestinationRegionId = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetDestinationZoneId(v string) *CreateDiskReplicaGroupRequest {
	s.DestinationZoneId = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetEnableRtc(v bool) *CreateDiskReplicaGroupRequest {
	s.EnableRtc = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetGroupName(v string) *CreateDiskReplicaGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetRPO(v int64) *CreateDiskReplicaGroupRequest {
	s.RPO = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetRegionId(v string) *CreateDiskReplicaGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetResourceGroupId(v string) *CreateDiskReplicaGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetSourceZoneId(v string) *CreateDiskReplicaGroupRequest {
	s.SourceZoneId = &v
	return s
}

func (s *CreateDiskReplicaGroupRequest) SetTag(v []*CreateDiskReplicaGroupRequestTag) *CreateDiskReplicaGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateDiskReplicaGroupRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDiskReplicaGroupRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDiskReplicaGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskReplicaGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDiskReplicaGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDiskReplicaGroupRequestTag) SetKey(v string) *CreateDiskReplicaGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDiskReplicaGroupRequestTag) SetValue(v string) *CreateDiskReplicaGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDiskReplicaGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
