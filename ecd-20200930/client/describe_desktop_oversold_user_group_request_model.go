// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopOversoldUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeDesktopOversoldUserGroupRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDesktopOversoldUserGroupRequest
	GetNextToken() *string
	SetOversoldGroupId(v string) *DescribeDesktopOversoldUserGroupRequest
	GetOversoldGroupId() *string
	SetUserGroupIds(v []*string) *DescribeDesktopOversoldUserGroupRequest
	GetUserGroupIds() []*string
}

type DescribeDesktopOversoldUserGroupRequest struct {
	MaxResults      *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OversoldGroupId *string   `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
	UserGroupIds    []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
}

func (s DescribeDesktopOversoldUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopOversoldUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopOversoldUserGroupRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDesktopOversoldUserGroupRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopOversoldUserGroupRequest) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *DescribeDesktopOversoldUserGroupRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *DescribeDesktopOversoldUserGroupRequest) SetMaxResults(v int32) *DescribeDesktopOversoldUserGroupRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopOversoldUserGroupRequest) SetNextToken(v string) *DescribeDesktopOversoldUserGroupRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopOversoldUserGroupRequest) SetOversoldGroupId(v string) *DescribeDesktopOversoldUserGroupRequest {
	s.OversoldGroupId = &v
	return s
}

func (s *DescribeDesktopOversoldUserGroupRequest) SetUserGroupIds(v []*string) *DescribeDesktopOversoldUserGroupRequest {
	s.UserGroupIds = v
	return s
}

func (s *DescribeDesktopOversoldUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
