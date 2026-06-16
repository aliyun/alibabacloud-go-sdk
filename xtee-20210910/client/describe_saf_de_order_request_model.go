// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafDeOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSafDeOrderRequest
	GetLang() *string
	SetCurrentPage(v int32) *DescribeSafDeOrderRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *DescribeSafDeOrderRequest
	GetEndDate() *string
	SetPageSize(v int32) *DescribeSafDeOrderRequest
	GetPageSize() *int32
	SetRegId(v string) *DescribeSafDeOrderRequest
	GetRegId() *string
	SetStartDate(v string) *DescribeSafDeOrderRequest
	GetStartDate() *string
}

type DescribeSafDeOrderRequest struct {
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// The end time.
	//
	// example:
	//
	// 1728008155799
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// The start date.
	//
	// example:
	//
	// 1728008155799
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s DescribeSafDeOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafDeOrderRequest) GoString() string {
	return s.String()
}

func (s *DescribeSafDeOrderRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSafDeOrderRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSafDeOrderRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeSafDeOrderRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSafDeOrderRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSafDeOrderRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeSafDeOrderRequest) SetLang(v string) *DescribeSafDeOrderRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSafDeOrderRequest) SetCurrentPage(v int32) *DescribeSafDeOrderRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSafDeOrderRequest) SetEndDate(v string) *DescribeSafDeOrderRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeSafDeOrderRequest) SetPageSize(v int32) *DescribeSafDeOrderRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSafDeOrderRequest) SetRegId(v string) *DescribeSafDeOrderRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSafDeOrderRequest) SetStartDate(v string) *DescribeSafDeOrderRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeSafDeOrderRequest) Validate() error {
	return dara.Validate(s)
}
