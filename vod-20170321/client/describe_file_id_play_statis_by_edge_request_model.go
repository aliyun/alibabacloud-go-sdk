// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileIdPlayStatisByEdgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeFileIdPlayStatisByEdgeRequest
	GetFrom() *string
	SetOwnerId(v string) *DescribeFileIdPlayStatisByEdgeRequest
	GetOwnerId() *string
	SetPageSize(v int32) *DescribeFileIdPlayStatisByEdgeRequest
	GetPageSize() *int32
	SetScrollToken(v string) *DescribeFileIdPlayStatisByEdgeRequest
	GetScrollToken() *string
	SetTo(v string) *DescribeFileIdPlayStatisByEdgeRequest
	GetTo() *string
}

type DescribeFileIdPlayStatisByEdgeRequest struct {
	// This parameter is required.
	From        *string `json:"From,omitempty" xml:"From,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ScrollToken *string `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
	// This parameter is required.
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s DescribeFileIdPlayStatisByEdgeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileIdPlayStatisByEdgeRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileIdPlayStatisByEdgeRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeFileIdPlayStatisByEdgeRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeFileIdPlayStatisByEdgeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFileIdPlayStatisByEdgeRequest) GetScrollToken() *string {
	return s.ScrollToken
}

func (s *DescribeFileIdPlayStatisByEdgeRequest) GetTo() *string {
	return s.To
}

func (s *DescribeFileIdPlayStatisByEdgeRequest) SetFrom(v string) *DescribeFileIdPlayStatisByEdgeRequest {
	s.From = &v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeRequest) SetOwnerId(v string) *DescribeFileIdPlayStatisByEdgeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeRequest) SetPageSize(v int32) *DescribeFileIdPlayStatisByEdgeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeRequest) SetScrollToken(v string) *DescribeFileIdPlayStatisByEdgeRequest {
	s.ScrollToken = &v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeRequest) SetTo(v string) *DescribeFileIdPlayStatisByEdgeRequest {
	s.To = &v
	return s
}

func (s *DescribeFileIdPlayStatisByEdgeRequest) Validate() error {
	return dara.Validate(s)
}
