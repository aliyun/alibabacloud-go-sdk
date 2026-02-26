// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPurchaserShopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListPurchaserShopsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPurchaserShopsRequest
	GetPageSize() *int32
}

type ListPurchaserShopsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListPurchaserShopsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPurchaserShopsRequest) GoString() string {
	return s.String()
}

func (s *ListPurchaserShopsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPurchaserShopsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPurchaserShopsRequest) SetPageNumber(v int32) *ListPurchaserShopsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPurchaserShopsRequest) SetPageSize(v int32) *ListPurchaserShopsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPurchaserShopsRequest) Validate() error {
	return dara.Validate(s)
}
