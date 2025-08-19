// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartStatisticsPageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeSmartStatisticsPageListResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeSmartStatisticsPageListResponseBodyItems) *DescribeSmartStatisticsPageListResponseBody
	GetItems() []*DescribeSmartStatisticsPageListResponseBodyItems
	SetPageSize(v int32) *DescribeSmartStatisticsPageListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSmartStatisticsPageListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeSmartStatisticsPageListResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *DescribeSmartStatisticsPageListResponseBody
	GetTotalPage() *int32
}

type DescribeSmartStatisticsPageListResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*DescribeSmartStatisticsPageListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 96943***4E39F805
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 29
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 3
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeSmartStatisticsPageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartStatisticsPageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSmartStatisticsPageListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSmartStatisticsPageListResponseBody) GetItems() []*DescribeSmartStatisticsPageListResponseBodyItems {
	return s.Items
}

func (s *DescribeSmartStatisticsPageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSmartStatisticsPageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSmartStatisticsPageListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSmartStatisticsPageListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeSmartStatisticsPageListResponseBody) SetCurrentPage(v int32) *DescribeSmartStatisticsPageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBody) SetItems(v []*DescribeSmartStatisticsPageListResponseBodyItems) *DescribeSmartStatisticsPageListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBody) SetPageSize(v int32) *DescribeSmartStatisticsPageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBody) SetRequestId(v string) *DescribeSmartStatisticsPageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBody) SetTotalCount(v int32) *DescribeSmartStatisticsPageListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBody) SetTotalPage(v int32) *DescribeSmartStatisticsPageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSmartStatisticsPageListResponseBodyItems struct {
	// example:
	//
	// 11/8
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 25
	PassRate *string `json:"PassRate,omitempty" xml:"PassRate,omitempty"`
	// example:
	//
	// SMART_VERIFY
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 20**40
	SceneId   *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// 1
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSmartStatisticsPageListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartStatisticsPageListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) GetDate() *string {
	return s.Date
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) GetPassRate() *string {
	return s.PassRate
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) GetSceneName() *string {
	return s.SceneName
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) SetDate(v string) *DescribeSmartStatisticsPageListResponseBodyItems {
	s.Date = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) SetPassRate(v string) *DescribeSmartStatisticsPageListResponseBodyItems {
	s.PassRate = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) SetProductCode(v string) *DescribeSmartStatisticsPageListResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) SetSceneId(v int64) *DescribeSmartStatisticsPageListResponseBodyItems {
	s.SceneId = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) SetSceneName(v string) *DescribeSmartStatisticsPageListResponseBodyItems {
	s.SceneName = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) SetSuccessCount(v int32) *DescribeSmartStatisticsPageListResponseBodyItems {
	s.SuccessCount = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) SetTotalCount(v int32) *DescribeSmartStatisticsPageListResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeSmartStatisticsPageListResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
