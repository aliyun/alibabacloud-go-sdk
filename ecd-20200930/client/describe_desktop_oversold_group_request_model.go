// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopOversoldGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeDesktopOversoldGroupRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDesktopOversoldGroupRequest
	GetNextToken() *string
	SetOversoldGroupIds(v []*string) *DescribeDesktopOversoldGroupRequest
	GetOversoldGroupIds() []*string
}

type DescribeDesktopOversoldGroupRequest struct {
	MaxResults       *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OversoldGroupIds []*string `json:"OversoldGroupIds,omitempty" xml:"OversoldGroupIds,omitempty" type:"Repeated"`
}

func (s DescribeDesktopOversoldGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopOversoldGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopOversoldGroupRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDesktopOversoldGroupRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopOversoldGroupRequest) GetOversoldGroupIds() []*string {
	return s.OversoldGroupIds
}

func (s *DescribeDesktopOversoldGroupRequest) SetMaxResults(v int32) *DescribeDesktopOversoldGroupRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopOversoldGroupRequest) SetNextToken(v string) *DescribeDesktopOversoldGroupRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopOversoldGroupRequest) SetOversoldGroupIds(v []*string) *DescribeDesktopOversoldGroupRequest {
	s.OversoldGroupIds = v
	return s
}

func (s *DescribeDesktopOversoldGroupRequest) Validate() error {
	return dara.Validate(s)
}
