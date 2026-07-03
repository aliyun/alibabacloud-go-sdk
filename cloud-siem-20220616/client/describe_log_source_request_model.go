// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogType(v string) *DescribeLogSourceRequest
	GetLogType() *string
	SetRegionId(v string) *DescribeLogSourceRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeLogSourceRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeLogSourceRequest
	GetRoleType() *int32
}

type DescribeLogSourceRequest struct {
	// The log type of the rule.
	//
	// example:
	//
	// HTTP_ACTIVITY
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The region where the Data Management center of Threat Analysis is deployed. You must select the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. An administrator can use this parameter to switch to the perspective of a member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts within the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeLogSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogSourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogSourceRequest) GetLogType() *string {
	return s.LogType
}

func (s *DescribeLogSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLogSourceRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeLogSourceRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeLogSourceRequest) SetLogType(v string) *DescribeLogSourceRequest {
	s.LogType = &v
	return s
}

func (s *DescribeLogSourceRequest) SetRegionId(v string) *DescribeLogSourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogSourceRequest) SetRoleFor(v int64) *DescribeLogSourceRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeLogSourceRequest) SetRoleType(v int32) *DescribeLogSourceRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeLogSourceRequest) Validate() error {
	return dara.Validate(s)
}
