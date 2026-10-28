// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchModifyEntitlementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *BatchModifyEntitlementRequest
	GetDesktopId() []*string
	SetEndUserId(v []*string) *BatchModifyEntitlementRequest
	GetEndUserId() []*string
	SetMaxDesktopPerUser(v int32) *BatchModifyEntitlementRequest
	GetMaxDesktopPerUser() *int32
	SetMaxUserPerDesktop(v int32) *BatchModifyEntitlementRequest
	GetMaxUserPerDesktop() *int32
	SetPreview(v bool) *BatchModifyEntitlementRequest
	GetPreview() *bool
	SetRegionId(v string) *BatchModifyEntitlementRequest
	GetRegionId() *string
	SetStrategy(v string) *BatchModifyEntitlementRequest
	GetStrategy() *string
}

type BatchModifyEntitlementRequest struct {
	// The IDs of the cloud desktops for which you want to modify the authorized users.
	//
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The user IDs, which are the usernames.
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	// The number of cloud desktops assigned to each user.
	//
	// example:
	//
	// 0
	MaxDesktopPerUser *int32 `json:"MaxDesktopPerUser,omitempty" xml:"MaxDesktopPerUser,omitempty"`
	// The number of users assigned to each cloud desktop.
	//
	// example:
	//
	// 1
	MaxUserPerDesktop *int32 `json:"MaxUserPerDesktop,omitempty" xml:"MaxUserPerDesktop,omitempty"`
	// Specifies whether to preview the assignment. A preview does not perform the actual assignment.
	//
	// example:
	//
	// true
	Preview *bool `json:"Preview,omitempty" xml:"Preview,omitempty"`
	// The region ID. You can call DescribeRegions to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The policy used when proportional matching is not possible. Valid values:
	//
	// - AVERAGE: Prioritizes assigning cloud desktops to each user. When the number of selected cloud desktops and users cannot be matched proportionally, cloud desktops are assigned to each user first.
	//
	// - CENTRAL: Prioritizes assigning a specified number of cloud desktops to users. When the number of selected cloud desktops and users cannot be matched proportionally, the specified number of cloud desktops are assigned to users first.
	//
	// example:
	//
	// AVERAGE
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s BatchModifyEntitlementRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchModifyEntitlementRequest) GoString() string {
	return s.String()
}

func (s *BatchModifyEntitlementRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *BatchModifyEntitlementRequest) GetEndUserId() []*string {
	return s.EndUserId
}

func (s *BatchModifyEntitlementRequest) GetMaxDesktopPerUser() *int32 {
	return s.MaxDesktopPerUser
}

func (s *BatchModifyEntitlementRequest) GetMaxUserPerDesktop() *int32 {
	return s.MaxUserPerDesktop
}

func (s *BatchModifyEntitlementRequest) GetPreview() *bool {
	return s.Preview
}

func (s *BatchModifyEntitlementRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BatchModifyEntitlementRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *BatchModifyEntitlementRequest) SetDesktopId(v []*string) *BatchModifyEntitlementRequest {
	s.DesktopId = v
	return s
}

func (s *BatchModifyEntitlementRequest) SetEndUserId(v []*string) *BatchModifyEntitlementRequest {
	s.EndUserId = v
	return s
}

func (s *BatchModifyEntitlementRequest) SetMaxDesktopPerUser(v int32) *BatchModifyEntitlementRequest {
	s.MaxDesktopPerUser = &v
	return s
}

func (s *BatchModifyEntitlementRequest) SetMaxUserPerDesktop(v int32) *BatchModifyEntitlementRequest {
	s.MaxUserPerDesktop = &v
	return s
}

func (s *BatchModifyEntitlementRequest) SetPreview(v bool) *BatchModifyEntitlementRequest {
	s.Preview = &v
	return s
}

func (s *BatchModifyEntitlementRequest) SetRegionId(v string) *BatchModifyEntitlementRequest {
	s.RegionId = &v
	return s
}

func (s *BatchModifyEntitlementRequest) SetStrategy(v string) *BatchModifyEntitlementRequest {
	s.Strategy = &v
	return s
}

func (s *BatchModifyEntitlementRequest) Validate() error {
	return dara.Validate(s)
}
