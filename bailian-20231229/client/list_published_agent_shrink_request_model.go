// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListPublishedAgentShrinkRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListPublishedAgentShrinkRequest
	GetPageSize() *int32
	SetSubTypesShrink(v string) *ListPublishedAgentShrinkRequest
	GetSubTypesShrink() *string
}

type ListPublishedAgentShrinkRequest struct {
	PageNo         *int32  `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	SubTypesShrink *string `json:"subTypes,omitempty" xml:"subTypes,omitempty"`
}

func (s ListPublishedAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentShrinkRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListPublishedAgentShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPublishedAgentShrinkRequest) GetSubTypesShrink() *string {
	return s.SubTypesShrink
}

func (s *ListPublishedAgentShrinkRequest) SetPageNo(v int32) *ListPublishedAgentShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *ListPublishedAgentShrinkRequest) SetPageSize(v int32) *ListPublishedAgentShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListPublishedAgentShrinkRequest) SetSubTypesShrink(v string) *ListPublishedAgentShrinkRequest {
	s.SubTypesShrink = &v
	return s
}

func (s *ListPublishedAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
