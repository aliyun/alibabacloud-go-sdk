// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserConnectTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeUserConnectTimeRequest
	GetEndTime() *string
	SetEndUserId(v string) *DescribeUserConnectTimeRequest
	GetEndUserId() *string
	SetMaxResults(v int32) *DescribeUserConnectTimeRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeUserConnectTimeRequest
	GetNextToken() *string
	SetOversoldGroupId(v string) *DescribeUserConnectTimeRequest
	GetOversoldGroupId() *string
	SetStartTime(v string) *DescribeUserConnectTimeRequest
	GetStartTime() *string
	SetUserDesktopId(v string) *DescribeUserConnectTimeRequest
	GetUserDesktopId() *string
	SetUserGroupId(v string) *DescribeUserConnectTimeRequest
	GetUserGroupId() *string
}

type DescribeUserConnectTimeRequest struct {
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndUserId       *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OversoldGroupId *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UserDesktopId   *string `json:"UserDesktopId,omitempty" xml:"UserDesktopId,omitempty"`
	UserGroupId     *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DescribeUserConnectTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserConnectTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectTimeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeUserConnectTimeRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeUserConnectTimeRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeUserConnectTimeRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeUserConnectTimeRequest) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *DescribeUserConnectTimeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeUserConnectTimeRequest) GetUserDesktopId() *string {
	return s.UserDesktopId
}

func (s *DescribeUserConnectTimeRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *DescribeUserConnectTimeRequest) SetEndTime(v string) *DescribeUserConnectTimeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUserConnectTimeRequest) SetEndUserId(v string) *DescribeUserConnectTimeRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeUserConnectTimeRequest) SetMaxResults(v int32) *DescribeUserConnectTimeRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUserConnectTimeRequest) SetNextToken(v string) *DescribeUserConnectTimeRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeUserConnectTimeRequest) SetOversoldGroupId(v string) *DescribeUserConnectTimeRequest {
	s.OversoldGroupId = &v
	return s
}

func (s *DescribeUserConnectTimeRequest) SetStartTime(v string) *DescribeUserConnectTimeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUserConnectTimeRequest) SetUserDesktopId(v string) *DescribeUserConnectTimeRequest {
	s.UserDesktopId = &v
	return s
}

func (s *DescribeUserConnectTimeRequest) SetUserGroupId(v string) *DescribeUserConnectTimeRequest {
	s.UserGroupId = &v
	return s
}

func (s *DescribeUserConnectTimeRequest) Validate() error {
	return dara.Validate(s)
}
