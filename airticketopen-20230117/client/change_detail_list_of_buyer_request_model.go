// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDetailListOfBuyerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageIndex(v int32) *ChangeDetailListOfBuyerRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ChangeDetailListOfBuyerRequest
	GetPageSize() *int32
	SetUtcCreateBegin(v int64) *ChangeDetailListOfBuyerRequest
	GetUtcCreateBegin() *int64
	SetUtcCreateEnd(v int64) *ChangeDetailListOfBuyerRequest
	GetUtcCreateEnd() *int64
}

type ChangeDetailListOfBuyerRequest struct {
	// example:
	//
	// 1
	PageIndex *int32 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 1677415276000
	UtcCreateBegin *int64 `json:"utc_create_begin,omitempty" xml:"utc_create_begin,omitempty"`
	// example:
	//
	// 1677415279000
	UtcCreateEnd *int64 `json:"utc_create_end,omitempty" xml:"utc_create_end,omitempty"`
}

func (s ChangeDetailListOfBuyerRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfBuyerRequest) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfBuyerRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ChangeDetailListOfBuyerRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ChangeDetailListOfBuyerRequest) GetUtcCreateBegin() *int64 {
	return s.UtcCreateBegin
}

func (s *ChangeDetailListOfBuyerRequest) GetUtcCreateEnd() *int64 {
	return s.UtcCreateEnd
}

func (s *ChangeDetailListOfBuyerRequest) SetPageIndex(v int32) *ChangeDetailListOfBuyerRequest {
	s.PageIndex = &v
	return s
}

func (s *ChangeDetailListOfBuyerRequest) SetPageSize(v int32) *ChangeDetailListOfBuyerRequest {
	s.PageSize = &v
	return s
}

func (s *ChangeDetailListOfBuyerRequest) SetUtcCreateBegin(v int64) *ChangeDetailListOfBuyerRequest {
	s.UtcCreateBegin = &v
	return s
}

func (s *ChangeDetailListOfBuyerRequest) SetUtcCreateEnd(v int64) *ChangeDetailListOfBuyerRequest {
	s.UtcCreateEnd = &v
	return s
}

func (s *ChangeDetailListOfBuyerRequest) Validate() error {
	return dara.Validate(s)
}
