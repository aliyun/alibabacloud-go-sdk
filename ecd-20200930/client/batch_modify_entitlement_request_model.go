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
	// The IDs of the cloud computers for which you want to modify end users.
	//
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The IDs of the users.
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	// The number of cloud computers allocated to each user.
	//
	// example:
	//
	// 0
	MaxDesktopPerUser *int32 `json:"MaxDesktopPerUser,omitempty" xml:"MaxDesktopPerUser,omitempty"`
	// The number of users assigned to each cloud computer.
	//
	// example:
	//
	// 1
	MaxUserPerDesktop *int32 `json:"MaxUserPerDesktop,omitempty" xml:"MaxUserPerDesktop,omitempty"`
	// Whether to preview the assign results instead of actually assigning cloud computers.
	//
	// example:
	//
	// true
	Preview *bool `json:"Preview,omitempty" xml:"Preview,omitempty"`
	// The ID of the region. Call the DescribeRegions operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The disproportional assignment policy. Valid values:
	//
	// AVERAGE: The system preferentially guarantees that each user is assigned with at least a cloud computer. If the number of selected cloud computers cannot be proportionally assigned to the selected users, ensure that each user is assigned a cloud computer.
	//
	// CENTRAL: The system preferentially assigns the designated number of cloud computers to each user. If the number of selected cloud computers cannot be proportionally assigned to the selected users, ensure that each user is assigned the specified number of cloud computers.
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
