// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeAlertSceneRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeAlertSceneRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeAlertSceneRequest
	GetRoleType() *int32
}

type DescribeAlertSceneRequest struct {
	// The location of the data management center for Threat Analysis. Select the location of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Assets are located in regions outside the Chinese mainland.
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

func (s DescribeAlertSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSceneRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertSceneRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlertSceneRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeAlertSceneRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeAlertSceneRequest) SetRegionId(v string) *DescribeAlertSceneRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertSceneRequest) SetRoleFor(v int64) *DescribeAlertSceneRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertSceneRequest) SetRoleType(v int32) *DescribeAlertSceneRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAlertSceneRequest) Validate() error {
	return dara.Validate(s)
}
