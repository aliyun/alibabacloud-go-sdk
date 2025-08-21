// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSecurityGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSecurityGroupsRequest
	GetPageSize() *int32
	SetSecurityGroupId(v string) *DescribeSecurityGroupsRequest
	GetSecurityGroupId() *string
	SetSecurityGroupName(v string) *DescribeSecurityGroupsRequest
	GetSecurityGroupName() *string
}

type DescribeSecurityGroupsRequest struct {
	// The page number.
	//
	// 	- Pages start from page 1.
	//
	// 	- Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// 	- Maximum value: 50.
	//
	// 	- Default value: 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-bp67acfmxazb4ph***
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the security group.
	//
	// example:
	//
	// DocTest
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s DescribeSecurityGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSecurityGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSecurityGroupsRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeSecurityGroupsRequest) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *DescribeSecurityGroupsRequest) SetPageNumber(v int32) *DescribeSecurityGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetPageSize(v int32) *DescribeSecurityGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetSecurityGroupId(v string) *DescribeSecurityGroupsRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetSecurityGroupName(v string) *DescribeSecurityGroupsRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) Validate() error {
	return dara.Validate(s)
}
