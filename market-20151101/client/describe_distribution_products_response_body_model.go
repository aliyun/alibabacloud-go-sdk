// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDistributionProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDistributionProductsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDistributionProductsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDistributionProductsResponseBody
	GetRequestId() *string
	SetResults(v []*DescribeDistributionProductsResponseBodyResults) *DescribeDistributionProductsResponseBody
	GetResults() []*DescribeDistributionProductsResponseBodyResults
	SetTotalCount(v int32) *DescribeDistributionProductsResponseBody
	GetTotalCount() *int32
}

type DescribeDistributionProductsResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 5BD09171-MB74-18D8-890E-C70C067527BE
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*DescribeDistributionProductsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDistributionProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistributionProductsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDistributionProductsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDistributionProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDistributionProductsResponseBody) GetResults() []*DescribeDistributionProductsResponseBodyResults {
	return s.Results
}

func (s *DescribeDistributionProductsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDistributionProductsResponseBody) SetPageNumber(v int32) *DescribeDistributionProductsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDistributionProductsResponseBody) SetPageSize(v int32) *DescribeDistributionProductsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDistributionProductsResponseBody) SetRequestId(v string) *DescribeDistributionProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDistributionProductsResponseBody) SetResults(v []*DescribeDistributionProductsResponseBodyResults) *DescribeDistributionProductsResponseBody {
	s.Results = v
	return s
}

func (s *DescribeDistributionProductsResponseBody) SetTotalCount(v int32) *DescribeDistributionProductsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDistributionProductsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDistributionProductsResponseBodyResults struct {
	// example:
	//
	// cmap*****
	Code              *string `json:"Code,omitempty" xml:"Code,omitempty"`
	FirstCategoryName *string `json:"FirstCategoryName,omitempty" xml:"FirstCategoryName,omitempty"`
	// example:
	//
	// //photogallery.oss-cn-hangzhou.aliyuncs.com/photo/1744526877246715/09605255-87fd-44d1-8143-96ebc8019d46.jpeg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 100
	Price *string `json:"Price,omitempty" xml:"Price,omitempty"`
	// example:
	//
	// 5
	Score              *string `json:"Score,omitempty" xml:"Score,omitempty"`
	SecondCategoryName *string `json:"SecondCategoryName,omitempty" xml:"SecondCategoryName,omitempty"`
	ShortDescription   *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	// example:
	//
	// 30
	SubmissionRadio *string `json:"SubmissionRadio,omitempty" xml:"SubmissionRadio,omitempty"`
	SupplierName    *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// example:
	//
	// 1911534921******
	SupplierUId *string `json:"SupplierUId,omitempty" xml:"SupplierUId,omitempty"`
	// example:
	//
	// 109
	TradeCount *string `json:"TradeCount,omitempty" xml:"TradeCount,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 55
	UserCommentCount *string `json:"UserCommentCount,omitempty" xml:"UserCommentCount,omitempty"`
}

func (s DescribeDistributionProductsResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistributionProductsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *DescribeDistributionProductsResponseBodyResults) GetFirstCategoryName() *string {
	return s.FirstCategoryName
}

func (s *DescribeDistributionProductsResponseBodyResults) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeDistributionProductsResponseBodyResults) GetName() *string {
	return s.Name
}

func (s *DescribeDistributionProductsResponseBodyResults) GetPrice() *string {
	return s.Price
}

func (s *DescribeDistributionProductsResponseBodyResults) GetScore() *string {
	return s.Score
}

func (s *DescribeDistributionProductsResponseBodyResults) GetSecondCategoryName() *string {
	return s.SecondCategoryName
}

func (s *DescribeDistributionProductsResponseBodyResults) GetShortDescription() *string {
	return s.ShortDescription
}

func (s *DescribeDistributionProductsResponseBodyResults) GetSubmissionRadio() *string {
	return s.SubmissionRadio
}

func (s *DescribeDistributionProductsResponseBodyResults) GetSupplierName() *string {
	return s.SupplierName
}

func (s *DescribeDistributionProductsResponseBodyResults) GetSupplierUId() *string {
	return s.SupplierUId
}

func (s *DescribeDistributionProductsResponseBodyResults) GetTradeCount() *string {
	return s.TradeCount
}

func (s *DescribeDistributionProductsResponseBodyResults) GetType() *string {
	return s.Type
}

func (s *DescribeDistributionProductsResponseBodyResults) GetUserCommentCount() *string {
	return s.UserCommentCount
}

func (s *DescribeDistributionProductsResponseBodyResults) SetCode(v string) *DescribeDistributionProductsResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetFirstCategoryName(v string) *DescribeDistributionProductsResponseBodyResults {
	s.FirstCategoryName = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetImageUrl(v string) *DescribeDistributionProductsResponseBodyResults {
	s.ImageUrl = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetName(v string) *DescribeDistributionProductsResponseBodyResults {
	s.Name = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetPrice(v string) *DescribeDistributionProductsResponseBodyResults {
	s.Price = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetScore(v string) *DescribeDistributionProductsResponseBodyResults {
	s.Score = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetSecondCategoryName(v string) *DescribeDistributionProductsResponseBodyResults {
	s.SecondCategoryName = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetShortDescription(v string) *DescribeDistributionProductsResponseBodyResults {
	s.ShortDescription = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetSubmissionRadio(v string) *DescribeDistributionProductsResponseBodyResults {
	s.SubmissionRadio = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetSupplierName(v string) *DescribeDistributionProductsResponseBodyResults {
	s.SupplierName = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetSupplierUId(v string) *DescribeDistributionProductsResponseBodyResults {
	s.SupplierUId = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetTradeCount(v string) *DescribeDistributionProductsResponseBodyResults {
	s.TradeCount = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetType(v string) *DescribeDistributionProductsResponseBodyResults {
	s.Type = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) SetUserCommentCount(v string) *DescribeDistributionProductsResponseBodyResults {
	s.UserCommentCount = &v
	return s
}

func (s *DescribeDistributionProductsResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
