// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *DescribeDesktopsRequest
	GetApiKey() *string
	SetDisplayType(v string) *DescribeDesktopsRequest
	GetDisplayType() *string
	SetMaxResults(v int32) *DescribeDesktopsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDesktopsRequest
	GetNextToken() *string
}

type DescribeDesktopsRequest struct {
	// This parameter is required.
	ApiKey      *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	DisplayType *string `json:"DisplayType,omitempty" xml:"DisplayType,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeDesktopsRequest) GetDisplayType() *string {
	return s.DisplayType
}

func (s *DescribeDesktopsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDesktopsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopsRequest) SetApiKey(v string) *DescribeDesktopsRequest {
	s.ApiKey = &v
	return s
}

func (s *DescribeDesktopsRequest) SetDisplayType(v string) *DescribeDesktopsRequest {
	s.DisplayType = &v
	return s
}

func (s *DescribeDesktopsRequest) SetMaxResults(v int32) *DescribeDesktopsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopsRequest) SetNextToken(v string) *DescribeDesktopsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
