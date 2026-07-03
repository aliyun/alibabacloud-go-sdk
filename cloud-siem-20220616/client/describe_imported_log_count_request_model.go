// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImportedLogCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeImportedLogCountRequest
	GetRegionId() *string
	SetRoleFor(v string) *DescribeImportedLogCountRequest
	GetRoleFor() *string
	SetRoleType(v string) *DescribeImportedLogCountRequest
	GetRoleType() *string
}

type DescribeImportedLogCountRequest struct {
	// The region where the management hub of Threat Analysis is located. Select the region of the management hub based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. An administrator can specify this parameter to switch to the perspective of the member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *string `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view.
	//
	// - 0: The view of the current Alibaba Cloud account.
	//
	// - 1: The view of all accounts that belong to the enterprise.
	//
	// example:
	//
	// 1
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeImportedLogCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportedLogCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeImportedLogCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImportedLogCountRequest) GetRoleFor() *string {
	return s.RoleFor
}

func (s *DescribeImportedLogCountRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeImportedLogCountRequest) SetRegionId(v string) *DescribeImportedLogCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImportedLogCountRequest) SetRoleFor(v string) *DescribeImportedLogCountRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeImportedLogCountRequest) SetRoleType(v string) *DescribeImportedLogCountRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeImportedLogCountRequest) Validate() error {
	return dara.Validate(s)
}
