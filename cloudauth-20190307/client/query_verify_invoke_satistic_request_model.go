// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVerifyInvokeSatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *QueryVerifyInvokeSatisticRequest
	GetCurrentPage() *int64
	SetEndDate(v int64) *QueryVerifyInvokeSatisticRequest
	GetEndDate() *int64
	SetPageSize(v int64) *QueryVerifyInvokeSatisticRequest
	GetPageSize() *int64
	SetProductProgramList(v string) *QueryVerifyInvokeSatisticRequest
	GetProductProgramList() *string
	SetProductType(v string) *QueryVerifyInvokeSatisticRequest
	GetProductType() *string
	SetSceneIdList(v string) *QueryVerifyInvokeSatisticRequest
	GetSceneIdList() *string
	SetStartDate(v int64) *QueryVerifyInvokeSatisticRequest
	GetStartDate() *int64
	SetStatisticsType(v string) *QueryVerifyInvokeSatisticRequest
	GetStatisticsType() *string
}

type QueryVerifyInvokeSatisticRequest struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// End date of the query.
	//
	// example:
	//
	// 1761926399999
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Number of items per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of product codes to query. Please refer to the productCode under the corresponding ProductType.
	//
	// example:
	//
	// ID_PRO
	ProductProgramList *string `json:"ProductProgramList,omitempty" xml:"ProductProgramList,omitempty"`
	// Product type:
	//
	// - **FINANCE_VERIFY**: Financial-grade real-person verification
	//
	// - **SMART_VERIFY**: Enhanced real-person verification (discontinued)
	//
	// - **FACE_VERIFY**: Real-person verification (discontinued)
	//
	// example:
	//
	// FINANCE_VERIFY
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// List of application scenarios.
	//
	// example:
	//
	// []
	SceneIdList *string `json:"SceneIdList,omitempty" xml:"SceneIdList,omitempty"`
	// Start date of the query.
	//
	// example:
	//
	// 1743436800000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Statistics dimension:
	//
	// - **day**: daily
	//
	// - **month**: monthly
	//
	// example:
	//
	// day
	StatisticsType *string `json:"StatisticsType,omitempty" xml:"StatisticsType,omitempty"`
}

func (s QueryVerifyInvokeSatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryVerifyInvokeSatisticRequest) GoString() string {
	return s.String()
}

func (s *QueryVerifyInvokeSatisticRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *QueryVerifyInvokeSatisticRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *QueryVerifyInvokeSatisticRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryVerifyInvokeSatisticRequest) GetProductProgramList() *string {
	return s.ProductProgramList
}

func (s *QueryVerifyInvokeSatisticRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryVerifyInvokeSatisticRequest) GetSceneIdList() *string {
	return s.SceneIdList
}

func (s *QueryVerifyInvokeSatisticRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *QueryVerifyInvokeSatisticRequest) GetStatisticsType() *string {
	return s.StatisticsType
}

func (s *QueryVerifyInvokeSatisticRequest) SetCurrentPage(v int64) *QueryVerifyInvokeSatisticRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryVerifyInvokeSatisticRequest) SetEndDate(v int64) *QueryVerifyInvokeSatisticRequest {
	s.EndDate = &v
	return s
}

func (s *QueryVerifyInvokeSatisticRequest) SetPageSize(v int64) *QueryVerifyInvokeSatisticRequest {
	s.PageSize = &v
	return s
}

func (s *QueryVerifyInvokeSatisticRequest) SetProductProgramList(v string) *QueryVerifyInvokeSatisticRequest {
	s.ProductProgramList = &v
	return s
}

func (s *QueryVerifyInvokeSatisticRequest) SetProductType(v string) *QueryVerifyInvokeSatisticRequest {
	s.ProductType = &v
	return s
}

func (s *QueryVerifyInvokeSatisticRequest) SetSceneIdList(v string) *QueryVerifyInvokeSatisticRequest {
	s.SceneIdList = &v
	return s
}

func (s *QueryVerifyInvokeSatisticRequest) SetStartDate(v int64) *QueryVerifyInvokeSatisticRequest {
	s.StartDate = &v
	return s
}

func (s *QueryVerifyInvokeSatisticRequest) SetStatisticsType(v string) *QueryVerifyInvokeSatisticRequest {
	s.StatisticsType = &v
	return s
}

func (s *QueryVerifyInvokeSatisticRequest) Validate() error {
	return dara.Validate(s)
}
