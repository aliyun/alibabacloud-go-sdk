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
	// example:
	//
	// sdl
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1646063922
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
