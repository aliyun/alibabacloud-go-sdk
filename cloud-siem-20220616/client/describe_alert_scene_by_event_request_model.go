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
	// The incident ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The region where the data management center of Threat Analysis is located. You must select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Assets are located outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose perspective the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts in the enterprise.
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
