// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertSourceWithEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncidentUuid(v string) *DescribeAlertSourceWithEventRequest
	GetIncidentUuid() *string
	SetRegionId(v string) *DescribeAlertSourceWithEventRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeAlertSourceWithEventRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeAlertSourceWithEventRequest
	GetRoleType() *int32
}

type DescribeAlertSourceWithEventRequest struct {
	// The globally unique ID of the incident.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The region where the data management center of Threat Analysis is located. Select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in China.
	//
	// - ap-southeast-1: Assets are located in regions outside China.
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

func (s DescribeAlertSourceWithEventRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSourceWithEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceWithEventRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeAlertSourceWithEventRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlertSourceWithEventRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeAlertSourceWithEventRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeAlertSourceWithEventRequest) SetIncidentUuid(v string) *DescribeAlertSourceWithEventRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAlertSourceWithEventRequest) SetRegionId(v string) *DescribeAlertSourceWithEventRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertSourceWithEventRequest) SetRoleFor(v int64) *DescribeAlertSourceWithEventRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertSourceWithEventRequest) SetRoleType(v int32) *DescribeAlertSourceWithEventRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAlertSourceWithEventRequest) Validate() error {
	return dara.Validate(s)
}
