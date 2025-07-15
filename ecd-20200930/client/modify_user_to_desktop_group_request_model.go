// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserToDesktopGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopGroupId(v string) *ModifyUserToDesktopGroupRequest
	GetDesktopGroupId() *string
	SetNewEndUserIds(v []*string) *ModifyUserToDesktopGroupRequest
	GetNewEndUserIds() []*string
	SetOldEndUserIds(v []*string) *ModifyUserToDesktopGroupRequest
	GetOldEndUserIds() []*string
	SetRegionId(v string) *ModifyUserToDesktopGroupRequest
	GetRegionId() *string
}

type ModifyUserToDesktopGroupRequest struct {
	// The ID of the cloud computer share.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The IDs of the end users that you want to add. You can configure 1 to 500 IDs.
	//
	// This parameter is required.
	NewEndUserIds []*string `json:"NewEndUserIds,omitempty" xml:"NewEndUserIds,omitempty" type:"Repeated"`
	// The IDs of the end users that you want to remove. You can configure 1 to 500 IDs.
	//
	// This parameter is required.
	OldEndUserIds []*string `json:"OldEndUserIds,omitempty" xml:"OldEndUserIds,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyUserToDesktopGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserToDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserToDesktopGroupRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *ModifyUserToDesktopGroupRequest) GetNewEndUserIds() []*string {
	return s.NewEndUserIds
}

func (s *ModifyUserToDesktopGroupRequest) GetOldEndUserIds() []*string {
	return s.OldEndUserIds
}

func (s *ModifyUserToDesktopGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyUserToDesktopGroupRequest) SetDesktopGroupId(v string) *ModifyUserToDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *ModifyUserToDesktopGroupRequest) SetNewEndUserIds(v []*string) *ModifyUserToDesktopGroupRequest {
	s.NewEndUserIds = v
	return s
}

func (s *ModifyUserToDesktopGroupRequest) SetOldEndUserIds(v []*string) *ModifyUserToDesktopGroupRequest {
	s.OldEndUserIds = v
	return s
}

func (s *ModifyUserToDesktopGroupRequest) SetRegionId(v string) *ModifyUserToDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserToDesktopGroupRequest) Validate() error {
	return dara.Validate(s)
}
