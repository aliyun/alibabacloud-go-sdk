// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsDbsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListMmsDbsShrinkRequest
	GetName() *string
	SetPageNum(v int32) *ListMmsDbsShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMmsDbsShrinkRequest
	GetPageSize() *int32
	SetSorterShrink(v string) *ListMmsDbsShrinkRequest
	GetSorterShrink() *string
	SetStatus(v string) *ListMmsDbsShrinkRequest
	GetStatus() *string
}

type ListMmsDbsShrinkRequest struct {
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 10
	PageSize     *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	SorterShrink *string `json:"sorter,omitempty" xml:"sorter,omitempty"`
	// example:
	//
	// STARTED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListMmsDbsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDbsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMmsDbsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListMmsDbsShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsDbsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsDbsShrinkRequest) GetSorterShrink() *string {
	return s.SorterShrink
}

func (s *ListMmsDbsShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListMmsDbsShrinkRequest) SetName(v string) *ListMmsDbsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListMmsDbsShrinkRequest) SetPageNum(v int32) *ListMmsDbsShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListMmsDbsShrinkRequest) SetPageSize(v int32) *ListMmsDbsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListMmsDbsShrinkRequest) SetSorterShrink(v string) *ListMmsDbsShrinkRequest {
	s.SorterShrink = &v
	return s
}

func (s *ListMmsDbsShrinkRequest) SetStatus(v string) *ListMmsDbsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListMmsDbsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
