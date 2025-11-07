// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVerifyInvokeSatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryVerifyInvokeSatisticResponseBody
	GetCode() *string
	SetCurrentPage(v int64) *QueryVerifyInvokeSatisticResponseBody
	GetCurrentPage() *int64
	SetItems(v []*QueryVerifyInvokeSatisticResponseBodyItems) *QueryVerifyInvokeSatisticResponseBody
	GetItems() []*QueryVerifyInvokeSatisticResponseBodyItems
	SetPageSize(v int64) *QueryVerifyInvokeSatisticResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *QueryVerifyInvokeSatisticResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryVerifyInvokeSatisticResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *QueryVerifyInvokeSatisticResponseBody
	GetTotalCount() *int64
	SetTotalPage(v int64) *QueryVerifyInvokeSatisticResponseBody
	GetTotalPage() *int64
}

type QueryVerifyInvokeSatisticResponseBody struct {
	// Response code, **200*	- indicates a successful response.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of returned data.
	Items []*QueryVerifyInvokeSatisticResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Number of items per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 2FA2C773-47DB-4156-B1EE-5B047321A939
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the response was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total count.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 3
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s QueryVerifyInvokeSatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryVerifyInvokeSatisticResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVerifyInvokeSatisticResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryVerifyInvokeSatisticResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *QueryVerifyInvokeSatisticResponseBody) GetItems() []*QueryVerifyInvokeSatisticResponseBodyItems {
	return s.Items
}

func (s *QueryVerifyInvokeSatisticResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryVerifyInvokeSatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryVerifyInvokeSatisticResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryVerifyInvokeSatisticResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryVerifyInvokeSatisticResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *QueryVerifyInvokeSatisticResponseBody) SetCode(v string) *QueryVerifyInvokeSatisticResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVerifyInvokeSatisticResponseBody) SetCurrentPage(v int64) *QueryVerifyInvokeSatisticResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryVerifyInvokeSatisticResponseBody) SetItems(v []*QueryVerifyInvokeSatisticResponseBodyItems) *QueryVerifyInvokeSatisticResponseBody {
	s.Items = v
	return s
}

func (s *QueryVerifyInvokeSatisticResponseBody) SetPageSize(v int64) *QueryVerifyInvokeSatisticResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryVerifyInvokeSatisticResponseBody) SetRequestId(v string) *QueryVerifyInvokeSatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryVerifyInvokeSatisticResponseBody) SetSuccess(v bool) *QueryVerifyInvokeSatisticResponseBody {
	s.Success = &v
	return s
}

func (s *QueryVerifyInvokeSatisticResponseBody) SetTotalCount(v int64) *QueryVerifyInvokeSatisticResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryVerifyInvokeSatisticResponseBody) SetTotalPage(v int64) *QueryVerifyInvokeSatisticResponseBody {
	s.TotalPage = &v
	return s
}

func (s *QueryVerifyInvokeSatisticResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryVerifyInvokeSatisticResponseBodyItems struct {
	// List of statistical data.
	Data []*QueryVerifyInvokeSatisticResponseBodyItemsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Statistics date.
	//
	// example:
	//
	// 2025-10-16
	StatisticsDate *string `json:"StatisticsDate,omitempty" xml:"StatisticsDate,omitempty"`
}

func (s QueryVerifyInvokeSatisticResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s QueryVerifyInvokeSatisticResponseBodyItems) GoString() string {
	return s.String()
}

func (s *QueryVerifyInvokeSatisticResponseBodyItems) GetData() []*QueryVerifyInvokeSatisticResponseBodyItemsData {
	return s.Data
}

func (s *QueryVerifyInvokeSatisticResponseBodyItems) GetStatisticsDate() *string {
	return s.StatisticsDate
}

func (s *QueryVerifyInvokeSatisticResponseBodyItems) SetData(v []*QueryVerifyInvokeSatisticResponseBodyItemsData) *QueryVerifyInvokeSatisticResponseBodyItems {
	s.Data = v
	return s
}

func (s *QueryVerifyInvokeSatisticResponseBodyItems) SetStatisticsDate(v string) *QueryVerifyInvokeSatisticResponseBodyItems {
	s.StatisticsDate = &v
	return s
}

func (s *QueryVerifyInvokeSatisticResponseBodyItems) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryVerifyInvokeSatisticResponseBodyItemsData struct {
	// Number of occurrences of the statistic.
	//
	// example:
	//
	// 3
	StatisticsCount *string `json:"StatisticsCount,omitempty" xml:"StatisticsCount,omitempty"`
	// ProductCode。
	//
	// example:
	//
	// ID_PRO
	StatisticsType *string `json:"StatisticsType,omitempty" xml:"StatisticsType,omitempty"`
}

func (s QueryVerifyInvokeSatisticResponseBodyItemsData) String() string {
	return dara.Prettify(s)
}

func (s QueryVerifyInvokeSatisticResponseBodyItemsData) GoString() string {
	return s.String()
}

func (s *QueryVerifyInvokeSatisticResponseBodyItemsData) GetStatisticsCount() *string {
	return s.StatisticsCount
}

func (s *QueryVerifyInvokeSatisticResponseBodyItemsData) GetStatisticsType() *string {
	return s.StatisticsType
}

func (s *QueryVerifyInvokeSatisticResponseBodyItemsData) SetStatisticsCount(v string) *QueryVerifyInvokeSatisticResponseBodyItemsData {
	s.StatisticsCount = &v
	return s
}

func (s *QueryVerifyInvokeSatisticResponseBodyItemsData) SetStatisticsType(v string) *QueryVerifyInvokeSatisticResponseBodyItemsData {
	s.StatisticsType = &v
	return s
}

func (s *QueryVerifyInvokeSatisticResponseBodyItemsData) Validate() error {
	return dara.Validate(s)
}
