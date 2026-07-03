// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertSceneByEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncidentUuid(v string) *DescribeAlertSceneByEventRequest
	GetIncidentUuid() *string
	SetRegionId(v string) *DescribeAlertSceneByEventRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeAlertSceneByEventRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeAlertSceneByEventRequest
	GetRoleType() *int32
}

type DescribeAlertSceneByEventRequest struct {
	// The event ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The region of the Data Management center. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: Assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. This parameter is used when an administrator switches to the perspective of a member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts that belong to the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeAlertSceneByEventRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSceneByEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneByEventRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeAlertSceneByEventRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlertSceneByEventRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeAlertSceneByEventRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeAlertSceneByEventRequest) SetIncidentUuid(v string) *DescribeAlertSceneByEventRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAlertSceneByEventRequest) SetRegionId(v string) *DescribeAlertSceneByEventRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertSceneByEventRequest) SetRoleFor(v int64) *DescribeAlertSceneByEventRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertSceneByEventRequest) SetRoleType(v int32) *DescribeAlertSceneByEventRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAlertSceneByEventRequest) Validate() error {
	return dara.Validate(s)
}
