// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSignalsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrder(v string) *ListSignalsRequest
	GetOrder() *string
	SetPageNumber(v string) *ListSignalsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListSignalsRequest
	GetPageSize() *string
	SetSortBy(v string) *ListSignalsRequest
	GetSortBy() *string
	SetStatus(v string) *ListSignalsRequest
	GetStatus() *string
	SetToken(v string) *ListSignalsRequest
	GetToken() *string
}

type ListSignalsRequest struct {
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// eyXXXX-XXXX.XXXXX
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s ListSignalsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSignalsRequest) GoString() string {
	return s.String()
}

func (s *ListSignalsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListSignalsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListSignalsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListSignalsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListSignalsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListSignalsRequest) GetToken() *string {
	return s.Token
}

func (s *ListSignalsRequest) SetOrder(v string) *ListSignalsRequest {
	s.Order = &v
	return s
}

func (s *ListSignalsRequest) SetPageNumber(v string) *ListSignalsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSignalsRequest) SetPageSize(v string) *ListSignalsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSignalsRequest) SetSortBy(v string) *ListSignalsRequest {
	s.SortBy = &v
	return s
}

func (s *ListSignalsRequest) SetStatus(v string) *ListSignalsRequest {
	s.Status = &v
	return s
}

func (s *ListSignalsRequest) SetToken(v string) *ListSignalsRequest {
	s.Token = &v
	return s
}

func (s *ListSignalsRequest) Validate() error {
	return dara.Validate(s)
}
