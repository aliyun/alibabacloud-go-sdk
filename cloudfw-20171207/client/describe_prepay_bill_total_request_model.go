// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrepayBillTotalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillType(v string) *DescribePrepayBillTotalRequest
	GetBillType() *string
	SetCurrentPage(v int64) *DescribePrepayBillTotalRequest
	GetCurrentPage() *int64
	SetEndTime(v string) *DescribePrepayBillTotalRequest
	GetEndTime() *string
	SetLang(v string) *DescribePrepayBillTotalRequest
	GetLang() *string
	SetPageSize(v int64) *DescribePrepayBillTotalRequest
	GetPageSize() *int64
	SetStartTime(v string) *DescribePrepayBillTotalRequest
	GetStartTime() *string
}

type DescribePrepayBillTotalRequest struct {
	// The bill type of the user. This parameter is required. An error is returned if this parameter is not specified. Valid values:
	//
	// - elastic_traffic: elastic traffic
	//
	// - sdl: sensitive data leak detection traffic
	//
	// example:
	//
	// sdl
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// The page number for a paged query. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1646063922
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language. Enumeration value.
	//
	// Default value: zh.
	//
	// Valid value: en.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time of the query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656750960
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePrepayBillTotalRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePrepayBillTotalRequest) GoString() string {
	return s.String()
}

func (s *DescribePrepayBillTotalRequest) GetBillType() *string {
	return s.BillType
}

func (s *DescribePrepayBillTotalRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribePrepayBillTotalRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePrepayBillTotalRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePrepayBillTotalRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePrepayBillTotalRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePrepayBillTotalRequest) SetBillType(v string) *DescribePrepayBillTotalRequest {
	s.BillType = &v
	return s
}

func (s *DescribePrepayBillTotalRequest) SetCurrentPage(v int64) *DescribePrepayBillTotalRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePrepayBillTotalRequest) SetEndTime(v string) *DescribePrepayBillTotalRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePrepayBillTotalRequest) SetLang(v string) *DescribePrepayBillTotalRequest {
	s.Lang = &v
	return s
}

func (s *DescribePrepayBillTotalRequest) SetPageSize(v int64) *DescribePrepayBillTotalRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePrepayBillTotalRequest) SetStartTime(v string) *DescribePrepayBillTotalRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePrepayBillTotalRequest) Validate() error {
	return dara.Validate(s)
}
