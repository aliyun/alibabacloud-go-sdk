// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopOversoldUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeDesktopOversoldUserRequest
	GetClientToken() *string
	SetEndUserId(v string) *DescribeDesktopOversoldUserRequest
	GetEndUserId() *string
	SetMaxResults(v int32) *DescribeDesktopOversoldUserRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDesktopOversoldUserRequest
	GetNextToken() *string
	SetOversoldGroupId(v string) *DescribeDesktopOversoldUserRequest
	GetOversoldGroupId() *string
	SetUserDesktopIds(v []*string) *DescribeDesktopOversoldUserRequest
	GetUserDesktopIds() []*string
	SetUserGroupId(v string) *DescribeDesktopOversoldUserRequest
	GetUserGroupId() *string
}

type DescribeDesktopOversoldUserRequest struct {
	ClientToken     *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndUserId       *string   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	MaxResults      *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OversoldGroupId *string   `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
	UserDesktopIds  []*string `json:"UserDesktopIds,omitempty" xml:"UserDesktopIds,omitempty" type:"Repeated"`
	UserGroupId     *string   `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DescribeDesktopOversoldUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopOversoldUserRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopOversoldUserRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDesktopOversoldUserRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeDesktopOversoldUserRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDesktopOversoldUserRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopOversoldUserRequest) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *DescribeDesktopOversoldUserRequest) GetUserDesktopIds() []*string {
	return s.UserDesktopIds
}

func (s *DescribeDesktopOversoldUserRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *DescribeDesktopOversoldUserRequest) SetClientToken(v string) *DescribeDesktopOversoldUserRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDesktopOversoldUserRequest) SetEndUserId(v string) *DescribeDesktopOversoldUserRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopOversoldUserRequest) SetMaxResults(v int32) *DescribeDesktopOversoldUserRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopOversoldUserRequest) SetNextToken(v string) *DescribeDesktopOversoldUserRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopOversoldUserRequest) SetOversoldGroupId(v string) *DescribeDesktopOversoldUserRequest {
	s.OversoldGroupId = &v
	return s
}

func (s *DescribeDesktopOversoldUserRequest) SetUserDesktopIds(v []*string) *DescribeDesktopOversoldUserRequest {
	s.UserDesktopIds = v
	return s
}

func (s *DescribeDesktopOversoldUserRequest) SetUserGroupId(v string) *DescribeDesktopOversoldUserRequest {
	s.UserGroupId = &v
	return s
}

func (s *DescribeDesktopOversoldUserRequest) Validate() error {
	return dara.Validate(s)
}
