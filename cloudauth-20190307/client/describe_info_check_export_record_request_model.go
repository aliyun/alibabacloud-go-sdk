// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInfoCheckExportRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeInfoCheckExportRecordRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *DescribeInfoCheckExportRecordRequest
	GetEndDate() *string
	SetPageSize(v int32) *DescribeInfoCheckExportRecordRequest
	GetPageSize() *int32
	SetProductType(v string) *DescribeInfoCheckExportRecordRequest
	GetProductType() *string
	SetStartDate(v string) *DescribeInfoCheckExportRecordRequest
	GetStartDate() *string
}

type DescribeInfoCheckExportRecordRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2025-10-11 21:24:48
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ID_CARD_2_META
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 2025-10-11 21:24:48
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeInfoCheckExportRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInfoCheckExportRecordRequest) GoString() string {
	return s.String()
}

func (s *DescribeInfoCheckExportRecordRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInfoCheckExportRecordRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeInfoCheckExportRecordRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInfoCheckExportRecordRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeInfoCheckExportRecordRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeInfoCheckExportRecordRequest) SetCurrentPage(v int32) *DescribeInfoCheckExportRecordRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInfoCheckExportRecordRequest) SetEndDate(v string) *DescribeInfoCheckExportRecordRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeInfoCheckExportRecordRequest) SetPageSize(v int32) *DescribeInfoCheckExportRecordRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInfoCheckExportRecordRequest) SetProductType(v string) *DescribeInfoCheckExportRecordRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeInfoCheckExportRecordRequest) SetStartDate(v string) *DescribeInfoCheckExportRecordRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeInfoCheckExportRecordRequest) Validate() error {
	return dara.Validate(s)
}
