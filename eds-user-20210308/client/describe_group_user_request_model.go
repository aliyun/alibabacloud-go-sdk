// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *DescribeGroupUserRequest
	GetBizType() *string
	SetFilter(v string) *DescribeGroupUserRequest
	GetFilter() *string
	SetGroupId(v string) *DescribeGroupUserRequest
	GetGroupId() *string
	SetMaxResults(v int32) *DescribeGroupUserRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeGroupUserRequest
	GetNextToken() *string
	SetSolutionId(v string) *DescribeGroupUserRequest
	GetSolutionId() *string
}

type DescribeGroupUserRequest struct {
	// example:
	//
	// ENTERPRISE
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Filter  *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// example:
	//
	// ug-12341234****
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// co-0esnf80jab***
	SolutionId *string `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s DescribeGroupUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupUserRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupUserRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeGroupUserRequest) GetFilter() *string {
	return s.Filter
}

func (s *DescribeGroupUserRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeGroupUserRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeGroupUserRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeGroupUserRequest) GetSolutionId() *string {
	return s.SolutionId
}

func (s *DescribeGroupUserRequest) SetBizType(v string) *DescribeGroupUserRequest {
	s.BizType = &v
	return s
}

func (s *DescribeGroupUserRequest) SetFilter(v string) *DescribeGroupUserRequest {
	s.Filter = &v
	return s
}

func (s *DescribeGroupUserRequest) SetGroupId(v string) *DescribeGroupUserRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupUserRequest) SetMaxResults(v int32) *DescribeGroupUserRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeGroupUserRequest) SetNextToken(v string) *DescribeGroupUserRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeGroupUserRequest) SetSolutionId(v string) *DescribeGroupUserRequest {
	s.SolutionId = &v
	return s
}

func (s *DescribeGroupUserRequest) Validate() error {
	return dara.Validate(s)
}
