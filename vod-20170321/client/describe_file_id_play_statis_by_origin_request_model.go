// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileIdPlayStatisByOriginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeFileIdPlayStatisByOriginRequest
	GetFrom() *string
	SetOwnerId(v string) *DescribeFileIdPlayStatisByOriginRequest
	GetOwnerId() *string
	SetPageSize(v int32) *DescribeFileIdPlayStatisByOriginRequest
	GetPageSize() *int32
	SetScrollToken(v string) *DescribeFileIdPlayStatisByOriginRequest
	GetScrollToken() *string
	SetTo(v string) *DescribeFileIdPlayStatisByOriginRequest
	GetTo() *string
}

type DescribeFileIdPlayStatisByOriginRequest struct {
	// This parameter is required.
	From        *string `json:"From,omitempty" xml:"From,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ScrollToken *string `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
	// This parameter is required.
	To *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s DescribeFileIdPlayStatisByOriginRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileIdPlayStatisByOriginRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileIdPlayStatisByOriginRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeFileIdPlayStatisByOriginRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeFileIdPlayStatisByOriginRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFileIdPlayStatisByOriginRequest) GetScrollToken() *string {
	return s.ScrollToken
}

func (s *DescribeFileIdPlayStatisByOriginRequest) GetTo() *string {
	return s.To
}

func (s *DescribeFileIdPlayStatisByOriginRequest) SetFrom(v string) *DescribeFileIdPlayStatisByOriginRequest {
	s.From = &v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginRequest) SetOwnerId(v string) *DescribeFileIdPlayStatisByOriginRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginRequest) SetPageSize(v int32) *DescribeFileIdPlayStatisByOriginRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginRequest) SetScrollToken(v string) *DescribeFileIdPlayStatisByOriginRequest {
	s.ScrollToken = &v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginRequest) SetTo(v string) *DescribeFileIdPlayStatisByOriginRequest {
	s.To = &v
	return s
}

func (s *DescribeFileIdPlayStatisByOriginRequest) Validate() error {
	return dara.Validate(s)
}
