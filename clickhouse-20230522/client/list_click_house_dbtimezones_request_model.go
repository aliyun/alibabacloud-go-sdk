// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClickHouseDBTimezonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListClickHouseDBTimezonesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListClickHouseDBTimezonesRequest
	GetPageSize() *int32
}

type ListClickHouseDBTimezonesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClickHouseDBTimezonesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClickHouseDBTimezonesRequest) GoString() string {
	return s.String()
}

func (s *ListClickHouseDBTimezonesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListClickHouseDBTimezonesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClickHouseDBTimezonesRequest) SetPageNumber(v int32) *ListClickHouseDBTimezonesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClickHouseDBTimezonesRequest) SetPageSize(v int32) *ListClickHouseDBTimezonesRequest {
	s.PageSize = &v
	return s
}

func (s *ListClickHouseDBTimezonesRequest) Validate() error {
	return dara.Validate(s)
}
