// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateRCSnapshotRequest
	GetDescription() *string
	SetDiskId(v string) *CreateRCSnapshotRequest
	GetDiskId() *string
	SetInstantAccess(v bool) *CreateRCSnapshotRequest
	GetInstantAccess() *bool
	SetInstantAccessRetentionDays(v int32) *CreateRCSnapshotRequest
	GetInstantAccessRetentionDays() *int32
	SetRegionId(v string) *CreateRCSnapshotRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateRCSnapshotRequest
	GetResourceGroupId() *string
	SetRetentionDays(v int32) *CreateRCSnapshotRequest
	GetRetentionDays() *int32
	SetTag(v []*CreateRCSnapshotRequestTag) *CreateRCSnapshotRequest
	GetTag() []*CreateRCSnapshotRequestTag
	SetZoneId(v string) *CreateRCSnapshotRequest
	GetZoneId() *string
}

type CreateRCSnapshotRequest struct {
	// The snapshot description. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The cloud disk ID.
	//
	// example:
	//
	// rcd-wz9f3peueu5npsl****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// none
	InstantAccess *bool `json:"InstantAccess,omitempty" xml:"InstantAccess,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// none
	InstantAccessRetentionDays *int32 `json:"InstantAccessRetentionDays,omitempty" xml:"InstantAccessRetentionDays,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// None
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The retention period of the snapshot. Valid values: 1 to 65536. Unit: days. The snapshot is automatically released when its retention period expires.
	//
	// By default, this parameter is left empty, which specifies that the snapshot is not automatically released.
	//
	// example:
	//
	// 2
	RetentionDays *int32                        `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	Tag           []*CreateRCSnapshotRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter has been deprecated.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateRCSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRCSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateRCSnapshotRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRCSnapshotRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *CreateRCSnapshotRequest) GetInstantAccess() *bool {
	return s.InstantAccess
}

func (s *CreateRCSnapshotRequest) GetInstantAccessRetentionDays() *int32 {
	return s.InstantAccessRetentionDays
}

func (s *CreateRCSnapshotRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRCSnapshotRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateRCSnapshotRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *CreateRCSnapshotRequest) GetTag() []*CreateRCSnapshotRequestTag {
	return s.Tag
}

func (s *CreateRCSnapshotRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateRCSnapshotRequest) SetDescription(v string) *CreateRCSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateRCSnapshotRequest) SetDiskId(v string) *CreateRCSnapshotRequest {
	s.DiskId = &v
	return s
}

func (s *CreateRCSnapshotRequest) SetInstantAccess(v bool) *CreateRCSnapshotRequest {
	s.InstantAccess = &v
	return s
}

func (s *CreateRCSnapshotRequest) SetInstantAccessRetentionDays(v int32) *CreateRCSnapshotRequest {
	s.InstantAccessRetentionDays = &v
	return s
}

func (s *CreateRCSnapshotRequest) SetRegionId(v string) *CreateRCSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRCSnapshotRequest) SetResourceGroupId(v string) *CreateRCSnapshotRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateRCSnapshotRequest) SetRetentionDays(v int32) *CreateRCSnapshotRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateRCSnapshotRequest) SetTag(v []*CreateRCSnapshotRequestTag) *CreateRCSnapshotRequest {
	s.Tag = v
	return s
}

func (s *CreateRCSnapshotRequest) SetZoneId(v string) *CreateRCSnapshotRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateRCSnapshotRequest) Validate() error {
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

type CreateRCSnapshotRequestTag struct {
	// example:
	//
	// None
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// None
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRCSnapshotRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateRCSnapshotRequestTag) GoString() string {
	return s.String()
}

func (s *CreateRCSnapshotRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateRCSnapshotRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateRCSnapshotRequestTag) SetKey(v string) *CreateRCSnapshotRequestTag {
	s.Key = &v
	return s
}

func (s *CreateRCSnapshotRequestTag) SetValue(v string) *CreateRCSnapshotRequestTag {
	s.Value = &v
	return s
}

func (s *CreateRCSnapshotRequestTag) Validate() error {
	return dara.Validate(s)
}
