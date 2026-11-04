// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventRecordPlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryEventRecordPlansRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *QueryEventRecordPlansRequest
	GetPageSize() *int32
}

type QueryEventRecordPlansRequest struct {
	// The page number to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Minimum value: 1. Maximum value: 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryEventRecordPlansRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEventRecordPlansRequest) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlansRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryEventRecordPlansRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryEventRecordPlansRequest) SetCurrentPage(v int32) *QueryEventRecordPlansRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryEventRecordPlansRequest) SetPageSize(v int32) *QueryEventRecordPlansRequest {
	s.PageSize = &v
	return s
}

func (s *QueryEventRecordPlansRequest) Validate() error {
	return dara.Validate(s)
}
