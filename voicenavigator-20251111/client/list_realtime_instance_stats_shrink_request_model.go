// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeInstanceStatsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *ListRealtimeInstanceStatsShrinkRequest
	GetInstanceIdsShrink() *string
	SetPageNumber(v int32) *ListRealtimeInstanceStatsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRealtimeInstanceStatsShrinkRequest
	GetPageSize() *int32
}

type ListRealtimeInstanceStatsShrinkRequest struct {
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRealtimeInstanceStatsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeInstanceStatsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListRealtimeInstanceStatsShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *ListRealtimeInstanceStatsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRealtimeInstanceStatsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRealtimeInstanceStatsShrinkRequest) SetInstanceIdsShrink(v string) *ListRealtimeInstanceStatsShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *ListRealtimeInstanceStatsShrinkRequest) SetPageNumber(v int32) *ListRealtimeInstanceStatsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRealtimeInstanceStatsShrinkRequest) SetPageSize(v int32) *ListRealtimeInstanceStatsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListRealtimeInstanceStatsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
