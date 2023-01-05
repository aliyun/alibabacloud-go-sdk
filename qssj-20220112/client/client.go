// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetAgeDistributionRequest struct {
	CateIds    *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	EndingDate *string `json:"EndingDate,omitempty" xml:"EndingDate,omitempty"`
	StartDate  *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetAgeDistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgeDistributionRequest) GoString() string {
	return s.String()
}

func (s *GetAgeDistributionRequest) SetCateIds(v string) *GetAgeDistributionRequest {
	s.CateIds = &v
	return s
}

func (s *GetAgeDistributionRequest) SetEndingDate(v string) *GetAgeDistributionRequest {
	s.EndingDate = &v
	return s
}

func (s *GetAgeDistributionRequest) SetStartDate(v string) *GetAgeDistributionRequest {
	s.StartDate = &v
	return s
}

type GetAgeDistributionResponseBody struct {
	Code            *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            []*GetAgeDistributionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message         *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *string                               `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetAgeDistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgeDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgeDistributionResponseBody) SetCode(v string) *GetAgeDistributionResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgeDistributionResponseBody) SetData(v []*GetAgeDistributionResponseBodyData) *GetAgeDistributionResponseBody {
	s.Data = v
	return s
}

func (s *GetAgeDistributionResponseBody) SetMessage(v string) *GetAgeDistributionResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgeDistributionResponseBody) SetRequestId(v string) *GetAgeDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgeDistributionResponseBody) SetSuccessResponse(v string) *GetAgeDistributionResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetAgeDistributionResponseBodyData struct {
	AgeRange      *string `json:"AgeRange,omitempty" xml:"AgeRange,omitempty"`
	SaleNumbers   *int64  `json:"SaleNumbers,omitempty" xml:"SaleNumbers,omitempty"`
	SearchNumbers *int64  `json:"SearchNumbers,omitempty" xml:"SearchNumbers,omitempty"`
}

func (s GetAgeDistributionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAgeDistributionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgeDistributionResponseBodyData) SetAgeRange(v string) *GetAgeDistributionResponseBodyData {
	s.AgeRange = &v
	return s
}

func (s *GetAgeDistributionResponseBodyData) SetSaleNumbers(v int64) *GetAgeDistributionResponseBodyData {
	s.SaleNumbers = &v
	return s
}

func (s *GetAgeDistributionResponseBodyData) SetSearchNumbers(v int64) *GetAgeDistributionResponseBodyData {
	s.SearchNumbers = &v
	return s
}

type GetAgeDistributionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAgeDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgeDistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgeDistributionResponse) GoString() string {
	return s.String()
}

func (s *GetAgeDistributionResponse) SetHeaders(v map[string]*string) *GetAgeDistributionResponse {
	s.Headers = v
	return s
}

func (s *GetAgeDistributionResponse) SetStatusCode(v int32) *GetAgeDistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgeDistributionResponse) SetBody(v *GetAgeDistributionResponseBody) *GetAgeDistributionResponse {
	s.Body = v
	return s
}

type GetAllTrendCategoryResponseBody struct {
	Code            *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            []*GetAllTrendCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message         *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool                                  `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetAllTrendCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllTrendCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllTrendCategoryResponseBody) SetCode(v string) *GetAllTrendCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetAllTrendCategoryResponseBody) SetData(v []*GetAllTrendCategoryResponseBodyData) *GetAllTrendCategoryResponseBody {
	s.Data = v
	return s
}

func (s *GetAllTrendCategoryResponseBody) SetMessage(v string) *GetAllTrendCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *GetAllTrendCategoryResponseBody) SetRequestId(v string) *GetAllTrendCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllTrendCategoryResponseBody) SetSuccessResponse(v bool) *GetAllTrendCategoryResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetAllTrendCategoryResponseBodyData struct {
	CategoryId        *int64        `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryLevel     *int32        `json:"CategoryLevel,omitempty" xml:"CategoryLevel,omitempty"`
	CategoryName      *string       `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	ChildCategory     []interface{} `json:"ChildCategory,omitempty" xml:"ChildCategory,omitempty" type:"Repeated"`
	SuperCategoryName *string       `json:"SuperCategoryName,omitempty" xml:"SuperCategoryName,omitempty"`
}

func (s GetAllTrendCategoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAllTrendCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAllTrendCategoryResponseBodyData) SetCategoryId(v int64) *GetAllTrendCategoryResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *GetAllTrendCategoryResponseBodyData) SetCategoryLevel(v int32) *GetAllTrendCategoryResponseBodyData {
	s.CategoryLevel = &v
	return s
}

func (s *GetAllTrendCategoryResponseBodyData) SetCategoryName(v string) *GetAllTrendCategoryResponseBodyData {
	s.CategoryName = &v
	return s
}

func (s *GetAllTrendCategoryResponseBodyData) SetChildCategory(v []interface{}) *GetAllTrendCategoryResponseBodyData {
	s.ChildCategory = v
	return s
}

func (s *GetAllTrendCategoryResponseBodyData) SetSuperCategoryName(v string) *GetAllTrendCategoryResponseBodyData {
	s.SuperCategoryName = &v
	return s
}

type GetAllTrendCategoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAllTrendCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllTrendCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllTrendCategoryResponse) GoString() string {
	return s.String()
}

func (s *GetAllTrendCategoryResponse) SetHeaders(v map[string]*string) *GetAllTrendCategoryResponse {
	s.Headers = v
	return s
}

func (s *GetAllTrendCategoryResponse) SetStatusCode(v int32) *GetAllTrendCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllTrendCategoryResponse) SetBody(v *GetAllTrendCategoryResponseBody) *GetAllTrendCategoryResponse {
	s.Body = v
	return s
}

type GetCrowdDataRequest struct {
	CateIds    *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	EndingDate *string `json:"EndingDate,omitempty" xml:"EndingDate,omitempty"`
	PageIndex  *int64  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	StartDate  *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetCrowdDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCrowdDataRequest) GoString() string {
	return s.String()
}

func (s *GetCrowdDataRequest) SetCateIds(v string) *GetCrowdDataRequest {
	s.CateIds = &v
	return s
}

func (s *GetCrowdDataRequest) SetEndingDate(v string) *GetCrowdDataRequest {
	s.EndingDate = &v
	return s
}

func (s *GetCrowdDataRequest) SetPageIndex(v int64) *GetCrowdDataRequest {
	s.PageIndex = &v
	return s
}

func (s *GetCrowdDataRequest) SetStartDate(v string) *GetCrowdDataRequest {
	s.StartDate = &v
	return s
}

type GetCrowdDataResponseBody struct {
	Code            *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            []*GetCrowdDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message         *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *string                         `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetCrowdDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCrowdDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetCrowdDataResponseBody) SetCode(v string) *GetCrowdDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetCrowdDataResponseBody) SetData(v []*GetCrowdDataResponseBodyData) *GetCrowdDataResponseBody {
	s.Data = v
	return s
}

func (s *GetCrowdDataResponseBody) SetMessage(v string) *GetCrowdDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetCrowdDataResponseBody) SetRequestId(v string) *GetCrowdDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCrowdDataResponseBody) SetSuccessResponse(v string) *GetCrowdDataResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetCrowdDataResponseBodyData struct {
	Age               *int32   `json:"Age,omitempty" xml:"Age,omitempty"`
	BuyerId           *int64   `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	CateId            *int64   `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName          *string  `json:"CateName,omitempty" xml:"CateName,omitempty"`
	Classification    *string  `json:"Classification,omitempty" xml:"Classification,omitempty"`
	ClassificationNew *string  `json:"ClassificationNew,omitempty" xml:"ClassificationNew,omitempty"`
	Gender            *string  `json:"Gender,omitempty" xml:"Gender,omitempty"`
	OnlineDate        *string  `json:"OnlineDate,omitempty" xml:"OnlineDate,omitempty"`
	OrderAmount       *float64 `json:"OrderAmount,omitempty" xml:"OrderAmount,omitempty"`
	OrderQuantity     *int64   `json:"OrderQuantity,omitempty" xml:"OrderQuantity,omitempty"`
	Province          *string  `json:"Province,omitempty" xml:"Province,omitempty"`
	SearchVolume      *int64   `json:"SearchVolume,omitempty" xml:"SearchVolume,omitempty"`
	TrendLevel        *string  `json:"TrendLevel,omitempty" xml:"TrendLevel,omitempty"`
}

func (s GetCrowdDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCrowdDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCrowdDataResponseBodyData) SetAge(v int32) *GetCrowdDataResponseBodyData {
	s.Age = &v
	return s
}

func (s *GetCrowdDataResponseBodyData) SetBuyerId(v int64) *GetCrowdDataResponseBodyData {
	s.BuyerId = &v
	return s
}

func (s *GetCrowdDataResponseBodyData) SetCateId(v int64) *GetCrowdDataResponseBodyData {
	s.CateId = &v
	return s
}

func (s *GetCrowdDataResponseBodyData) SetCateName(v string) *GetCrowdDataResponseBodyData {
	s.CateName = &v
	return s
}

func (s *GetCrowdDataResponseBodyData) SetClassification(v string) *GetCrowdDataResponseBodyData {
	s.Classification = &v
	return s
}

func (s *GetCrowdDataResponseBodyData) SetClassificationNew(v string) *GetCrowdDataResponseBodyData {
	s.ClassificationNew = &v
	return s
}

func (s *GetCrowdDataResponseBodyData) SetGender(v string) *GetCrowdDataResponseBodyData {
	s.Gender = &v
	return s
}

func (s *GetCrowdDataResponseBodyData) SetOnlineDate(v string) *GetCrowdDataResponseBodyData {
	s.OnlineDate = &v
	return s
}

func (s *GetCrowdDataResponseBodyData) SetOrderAmount(v float64) *GetCrowdDataResponseBodyData {
	s.OrderAmount = &v
	return s
}

func (s *GetCrowdDataResponseBodyData) SetOrderQuantity(v int64) *GetCrowdDataResponseBodyData {
	s.OrderQuantity = &v
	return s
}

func (s *GetCrowdDataResponseBodyData) SetProvince(v string) *GetCrowdDataResponseBodyData {
	s.Province = &v
	return s
}

func (s *GetCrowdDataResponseBodyData) SetSearchVolume(v int64) *GetCrowdDataResponseBodyData {
	s.SearchVolume = &v
	return s
}

func (s *GetCrowdDataResponseBodyData) SetTrendLevel(v string) *GetCrowdDataResponseBodyData {
	s.TrendLevel = &v
	return s
}

type GetCrowdDataResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCrowdDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCrowdDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCrowdDataResponse) GoString() string {
	return s.String()
}

func (s *GetCrowdDataResponse) SetHeaders(v map[string]*string) *GetCrowdDataResponse {
	s.Headers = v
	return s
}

func (s *GetCrowdDataResponse) SetStatusCode(v int32) *GetCrowdDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCrowdDataResponse) SetBody(v *GetCrowdDataResponseBody) *GetCrowdDataResponse {
	s.Body = v
	return s
}

type GetCrowdLabelRequest struct {
	CateIds    *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	EndingDate *string `json:"EndingDate,omitempty" xml:"EndingDate,omitempty"`
	StartDate  *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetCrowdLabelRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCrowdLabelRequest) GoString() string {
	return s.String()
}

func (s *GetCrowdLabelRequest) SetCateIds(v string) *GetCrowdLabelRequest {
	s.CateIds = &v
	return s
}

func (s *GetCrowdLabelRequest) SetEndingDate(v string) *GetCrowdLabelRequest {
	s.EndingDate = &v
	return s
}

func (s *GetCrowdLabelRequest) SetStartDate(v string) *GetCrowdLabelRequest {
	s.StartDate = &v
	return s
}

type GetCrowdLabelResponseBody struct {
	Code            *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            []*GetCrowdLabelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message         *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *string                          `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetCrowdLabelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCrowdLabelResponseBody) GoString() string {
	return s.String()
}

func (s *GetCrowdLabelResponseBody) SetCode(v string) *GetCrowdLabelResponseBody {
	s.Code = &v
	return s
}

func (s *GetCrowdLabelResponseBody) SetData(v []*GetCrowdLabelResponseBodyData) *GetCrowdLabelResponseBody {
	s.Data = v
	return s
}

func (s *GetCrowdLabelResponseBody) SetMessage(v string) *GetCrowdLabelResponseBody {
	s.Message = &v
	return s
}

func (s *GetCrowdLabelResponseBody) SetRequestId(v string) *GetCrowdLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCrowdLabelResponseBody) SetSuccessResponse(v string) *GetCrowdLabelResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetCrowdLabelResponseBodyData struct {
	ClosingDate    *string  `json:"ClosingDate,omitempty" xml:"ClosingDate,omitempty"`
	LabelName      *string  `json:"LabelName,omitempty" xml:"LabelName,omitempty"`
	OrderAmount    *float64 `json:"OrderAmount,omitempty" xml:"OrderAmount,omitempty"`
	PurchaseAmount *float64 `json:"PurchaseAmount,omitempty" xml:"PurchaseAmount,omitempty"`
	SearchVolume   *float64 `json:"SearchVolume,omitempty" xml:"SearchVolume,omitempty"`
}

func (s GetCrowdLabelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCrowdLabelResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCrowdLabelResponseBodyData) SetClosingDate(v string) *GetCrowdLabelResponseBodyData {
	s.ClosingDate = &v
	return s
}

func (s *GetCrowdLabelResponseBodyData) SetLabelName(v string) *GetCrowdLabelResponseBodyData {
	s.LabelName = &v
	return s
}

func (s *GetCrowdLabelResponseBodyData) SetOrderAmount(v float64) *GetCrowdLabelResponseBodyData {
	s.OrderAmount = &v
	return s
}

func (s *GetCrowdLabelResponseBodyData) SetPurchaseAmount(v float64) *GetCrowdLabelResponseBodyData {
	s.PurchaseAmount = &v
	return s
}

func (s *GetCrowdLabelResponseBodyData) SetSearchVolume(v float64) *GetCrowdLabelResponseBodyData {
	s.SearchVolume = &v
	return s
}

type GetCrowdLabelResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCrowdLabelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCrowdLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCrowdLabelResponse) GoString() string {
	return s.String()
}

func (s *GetCrowdLabelResponse) SetHeaders(v map[string]*string) *GetCrowdLabelResponse {
	s.Headers = v
	return s
}

func (s *GetCrowdLabelResponse) SetStatusCode(v int32) *GetCrowdLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCrowdLabelResponse) SetBody(v *GetCrowdLabelResponseBody) *GetCrowdLabelResponse {
	s.Body = v
	return s
}

type GetCrowdReginRequest struct {
	CateIds    *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	EndingDate *string `json:"EndingDate,omitempty" xml:"EndingDate,omitempty"`
	StartDate  *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetCrowdReginRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCrowdReginRequest) GoString() string {
	return s.String()
}

func (s *GetCrowdReginRequest) SetCateIds(v string) *GetCrowdReginRequest {
	s.CateIds = &v
	return s
}

func (s *GetCrowdReginRequest) SetEndingDate(v string) *GetCrowdReginRequest {
	s.EndingDate = &v
	return s
}

func (s *GetCrowdReginRequest) SetStartDate(v string) *GetCrowdReginRequest {
	s.StartDate = &v
	return s
}

type GetCrowdReginResponseBody struct {
	Code            *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            *GetCrowdReginResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message         *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool                          `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetCrowdReginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCrowdReginResponseBody) GoString() string {
	return s.String()
}

func (s *GetCrowdReginResponseBody) SetCode(v string) *GetCrowdReginResponseBody {
	s.Code = &v
	return s
}

func (s *GetCrowdReginResponseBody) SetData(v *GetCrowdReginResponseBodyData) *GetCrowdReginResponseBody {
	s.Data = v
	return s
}

func (s *GetCrowdReginResponseBody) SetMessage(v string) *GetCrowdReginResponseBody {
	s.Message = &v
	return s
}

func (s *GetCrowdReginResponseBody) SetRequestId(v string) *GetCrowdReginResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCrowdReginResponseBody) SetSuccessResponse(v bool) *GetCrowdReginResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetCrowdReginResponseBodyData struct {
	SaleNumbers   []*GetCrowdReginResponseBodyDataSaleNumbers   `json:"SaleNumbers,omitempty" xml:"SaleNumbers,omitempty" type:"Repeated"`
	SearchNumbers []*GetCrowdReginResponseBodyDataSearchNumbers `json:"SearchNumbers,omitempty" xml:"SearchNumbers,omitempty" type:"Repeated"`
}

func (s GetCrowdReginResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCrowdReginResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCrowdReginResponseBodyData) SetSaleNumbers(v []*GetCrowdReginResponseBodyDataSaleNumbers) *GetCrowdReginResponseBodyData {
	s.SaleNumbers = v
	return s
}

func (s *GetCrowdReginResponseBodyData) SetSearchNumbers(v []*GetCrowdReginResponseBodyDataSearchNumbers) *GetCrowdReginResponseBodyData {
	s.SearchNumbers = v
	return s
}

type GetCrowdReginResponseBodyDataSaleNumbers struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *int64  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCrowdReginResponseBodyDataSaleNumbers) String() string {
	return tea.Prettify(s)
}

func (s GetCrowdReginResponseBodyDataSaleNumbers) GoString() string {
	return s.String()
}

func (s *GetCrowdReginResponseBodyDataSaleNumbers) SetName(v string) *GetCrowdReginResponseBodyDataSaleNumbers {
	s.Name = &v
	return s
}

func (s *GetCrowdReginResponseBodyDataSaleNumbers) SetValue(v int64) *GetCrowdReginResponseBodyDataSaleNumbers {
	s.Value = &v
	return s
}

type GetCrowdReginResponseBodyDataSearchNumbers struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *int64  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCrowdReginResponseBodyDataSearchNumbers) String() string {
	return tea.Prettify(s)
}

func (s GetCrowdReginResponseBodyDataSearchNumbers) GoString() string {
	return s.String()
}

func (s *GetCrowdReginResponseBodyDataSearchNumbers) SetName(v string) *GetCrowdReginResponseBodyDataSearchNumbers {
	s.Name = &v
	return s
}

func (s *GetCrowdReginResponseBodyDataSearchNumbers) SetValue(v int64) *GetCrowdReginResponseBodyDataSearchNumbers {
	s.Value = &v
	return s
}

type GetCrowdReginResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCrowdReginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCrowdReginResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCrowdReginResponse) GoString() string {
	return s.String()
}

func (s *GetCrowdReginResponse) SetHeaders(v map[string]*string) *GetCrowdReginResponse {
	s.Headers = v
	return s
}

func (s *GetCrowdReginResponse) SetStatusCode(v int32) *GetCrowdReginResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCrowdReginResponse) SetBody(v *GetCrowdReginResponseBody) *GetCrowdReginResponse {
	s.Body = v
	return s
}

type GetOpportunityMarketRequest struct {
	CateIds     *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	TimeDisplay *int64  `json:"TimeDisplay,omitempty" xml:"TimeDisplay,omitempty"`
}

func (s GetOpportunityMarketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOpportunityMarketRequest) GoString() string {
	return s.String()
}

func (s *GetOpportunityMarketRequest) SetCateIds(v string) *GetOpportunityMarketRequest {
	s.CateIds = &v
	return s
}

func (s *GetOpportunityMarketRequest) SetTimeDisplay(v int64) *GetOpportunityMarketRequest {
	s.TimeDisplay = &v
	return s
}

type GetOpportunityMarketResponseBody struct {
	Code            *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            []*GetOpportunityMarketResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message         *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool                                   `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetOpportunityMarketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOpportunityMarketResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpportunityMarketResponseBody) SetCode(v string) *GetOpportunityMarketResponseBody {
	s.Code = &v
	return s
}

func (s *GetOpportunityMarketResponseBody) SetData(v []*GetOpportunityMarketResponseBodyData) *GetOpportunityMarketResponseBody {
	s.Data = v
	return s
}

func (s *GetOpportunityMarketResponseBody) SetMessage(v string) *GetOpportunityMarketResponseBody {
	s.Message = &v
	return s
}

func (s *GetOpportunityMarketResponseBody) SetRequestId(v string) *GetOpportunityMarketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpportunityMarketResponseBody) SetSuccessResponse(v bool) *GetOpportunityMarketResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetOpportunityMarketResponseBodyData struct {
	CateName          *string  `json:"CateName,omitempty" xml:"CateName,omitempty"`
	OpportunityIndex  *float64 `json:"OpportunityIndex,omitempty" xml:"OpportunityIndex,omitempty"`
	RelativeCommodity *float64 `json:"RelativeCommodity,omitempty" xml:"RelativeCommodity,omitempty"`
	RelativeDischarge *float64 `json:"RelativeDischarge,omitempty" xml:"RelativeDischarge,omitempty"`
	TimeUnit          *string  `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s GetOpportunityMarketResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOpportunityMarketResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOpportunityMarketResponseBodyData) SetCateName(v string) *GetOpportunityMarketResponseBodyData {
	s.CateName = &v
	return s
}

func (s *GetOpportunityMarketResponseBodyData) SetOpportunityIndex(v float64) *GetOpportunityMarketResponseBodyData {
	s.OpportunityIndex = &v
	return s
}

func (s *GetOpportunityMarketResponseBodyData) SetRelativeCommodity(v float64) *GetOpportunityMarketResponseBodyData {
	s.RelativeCommodity = &v
	return s
}

func (s *GetOpportunityMarketResponseBodyData) SetRelativeDischarge(v float64) *GetOpportunityMarketResponseBodyData {
	s.RelativeDischarge = &v
	return s
}

func (s *GetOpportunityMarketResponseBodyData) SetTimeUnit(v string) *GetOpportunityMarketResponseBodyData {
	s.TimeUnit = &v
	return s
}

type GetOpportunityMarketResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOpportunityMarketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOpportunityMarketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOpportunityMarketResponse) GoString() string {
	return s.String()
}

func (s *GetOpportunityMarketResponse) SetHeaders(v map[string]*string) *GetOpportunityMarketResponse {
	s.Headers = v
	return s
}

func (s *GetOpportunityMarketResponse) SetStatusCode(v int32) *GetOpportunityMarketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpportunityMarketResponse) SetBody(v *GetOpportunityMarketResponseBody) *GetOpportunityMarketResponse {
	s.Body = v
	return s
}

type GetPriceRangeRequest struct {
	BrandNames *string `json:"BrandNames,omitempty" xml:"BrandNames,omitempty"`
	CateIds    *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	EndingDate *string `json:"EndingDate,omitempty" xml:"EndingDate,omitempty"`
	Section    *int64  `json:"Section,omitempty" xml:"Section,omitempty"`
	ShopIds    *string `json:"ShopIds,omitempty" xml:"ShopIds,omitempty"`
	StartDate  *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetPriceRangeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPriceRangeRequest) GoString() string {
	return s.String()
}

func (s *GetPriceRangeRequest) SetBrandNames(v string) *GetPriceRangeRequest {
	s.BrandNames = &v
	return s
}

func (s *GetPriceRangeRequest) SetCateIds(v string) *GetPriceRangeRequest {
	s.CateIds = &v
	return s
}

func (s *GetPriceRangeRequest) SetEndingDate(v string) *GetPriceRangeRequest {
	s.EndingDate = &v
	return s
}

func (s *GetPriceRangeRequest) SetSection(v int64) *GetPriceRangeRequest {
	s.Section = &v
	return s
}

func (s *GetPriceRangeRequest) SetShopIds(v string) *GetPriceRangeRequest {
	s.ShopIds = &v
	return s
}

func (s *GetPriceRangeRequest) SetStartDate(v string) *GetPriceRangeRequest {
	s.StartDate = &v
	return s
}

type GetPriceRangeResponseBody struct {
	Code            *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            []*GetPriceRangeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message         *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *string                          `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetPriceRangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPriceRangeResponseBody) GoString() string {
	return s.String()
}

func (s *GetPriceRangeResponseBody) SetCode(v string) *GetPriceRangeResponseBody {
	s.Code = &v
	return s
}

func (s *GetPriceRangeResponseBody) SetData(v []*GetPriceRangeResponseBodyData) *GetPriceRangeResponseBody {
	s.Data = v
	return s
}

func (s *GetPriceRangeResponseBody) SetMessage(v string) *GetPriceRangeResponseBody {
	s.Message = &v
	return s
}

func (s *GetPriceRangeResponseBody) SetRequestId(v string) *GetPriceRangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPriceRangeResponseBody) SetSuccessResponse(v string) *GetPriceRangeResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetPriceRangeResponseBodyData struct {
	ClosingDate *string  `json:"ClosingDate,omitempty" xml:"ClosingDate,omitempty"`
	GoodsSales  *int64   `json:"GoodsSales,omitempty" xml:"GoodsSales,omitempty"`
	PriceRange  *string  `json:"PriceRange,omitempty" xml:"PriceRange,omitempty"`
	SalesVolume *float64 `json:"SalesVolume,omitempty" xml:"SalesVolume,omitempty"`
}

func (s GetPriceRangeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPriceRangeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPriceRangeResponseBodyData) SetClosingDate(v string) *GetPriceRangeResponseBodyData {
	s.ClosingDate = &v
	return s
}

func (s *GetPriceRangeResponseBodyData) SetGoodsSales(v int64) *GetPriceRangeResponseBodyData {
	s.GoodsSales = &v
	return s
}

func (s *GetPriceRangeResponseBodyData) SetPriceRange(v string) *GetPriceRangeResponseBodyData {
	s.PriceRange = &v
	return s
}

func (s *GetPriceRangeResponseBodyData) SetSalesVolume(v float64) *GetPriceRangeResponseBodyData {
	s.SalesVolume = &v
	return s
}

type GetPriceRangeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPriceRangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPriceRangeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPriceRangeResponse) GoString() string {
	return s.String()
}

func (s *GetPriceRangeResponse) SetHeaders(v map[string]*string) *GetPriceRangeResponse {
	s.Headers = v
	return s
}

func (s *GetPriceRangeResponse) SetStatusCode(v int32) *GetPriceRangeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPriceRangeResponse) SetBody(v *GetPriceRangeResponseBody) *GetPriceRangeResponse {
	s.Body = v
	return s
}

type GetSexRatioRequest struct {
	CateIds    *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	EndingDate *string `json:"EndingDate,omitempty" xml:"EndingDate,omitempty"`
	StartDate  *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetSexRatioRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSexRatioRequest) GoString() string {
	return s.String()
}

func (s *GetSexRatioRequest) SetCateIds(v string) *GetSexRatioRequest {
	s.CateIds = &v
	return s
}

func (s *GetSexRatioRequest) SetEndingDate(v string) *GetSexRatioRequest {
	s.EndingDate = &v
	return s
}

func (s *GetSexRatioRequest) SetStartDate(v string) *GetSexRatioRequest {
	s.StartDate = &v
	return s
}

type GetSexRatioResponseBody struct {
	Code            *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            *GetSexRatioResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message         *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *string                      `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetSexRatioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSexRatioResponseBody) GoString() string {
	return s.String()
}

func (s *GetSexRatioResponseBody) SetCode(v string) *GetSexRatioResponseBody {
	s.Code = &v
	return s
}

func (s *GetSexRatioResponseBody) SetData(v *GetSexRatioResponseBodyData) *GetSexRatioResponseBody {
	s.Data = v
	return s
}

func (s *GetSexRatioResponseBody) SetMessage(v string) *GetSexRatioResponseBody {
	s.Message = &v
	return s
}

func (s *GetSexRatioResponseBody) SetRequestId(v string) *GetSexRatioResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSexRatioResponseBody) SetSuccessResponse(v string) *GetSexRatioResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetSexRatioResponseBodyData struct {
	SaleNumbers   []*GetSexRatioResponseBodyDataSaleNumbers   `json:"SaleNumbers,omitempty" xml:"SaleNumbers,omitempty" type:"Repeated"`
	SearchNumbers []*GetSexRatioResponseBodyDataSearchNumbers `json:"SearchNumbers,omitempty" xml:"SearchNumbers,omitempty" type:"Repeated"`
}

func (s GetSexRatioResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSexRatioResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSexRatioResponseBodyData) SetSaleNumbers(v []*GetSexRatioResponseBodyDataSaleNumbers) *GetSexRatioResponseBodyData {
	s.SaleNumbers = v
	return s
}

func (s *GetSexRatioResponseBodyData) SetSearchNumbers(v []*GetSexRatioResponseBodyDataSearchNumbers) *GetSexRatioResponseBodyData {
	s.SearchNumbers = v
	return s
}

type GetSexRatioResponseBodyDataSaleNumbers struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *int64  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetSexRatioResponseBodyDataSaleNumbers) String() string {
	return tea.Prettify(s)
}

func (s GetSexRatioResponseBodyDataSaleNumbers) GoString() string {
	return s.String()
}

func (s *GetSexRatioResponseBodyDataSaleNumbers) SetName(v string) *GetSexRatioResponseBodyDataSaleNumbers {
	s.Name = &v
	return s
}

func (s *GetSexRatioResponseBodyDataSaleNumbers) SetValue(v int64) *GetSexRatioResponseBodyDataSaleNumbers {
	s.Value = &v
	return s
}

type GetSexRatioResponseBodyDataSearchNumbers struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *int64  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetSexRatioResponseBodyDataSearchNumbers) String() string {
	return tea.Prettify(s)
}

func (s GetSexRatioResponseBodyDataSearchNumbers) GoString() string {
	return s.String()
}

func (s *GetSexRatioResponseBodyDataSearchNumbers) SetName(v string) *GetSexRatioResponseBodyDataSearchNumbers {
	s.Name = &v
	return s
}

func (s *GetSexRatioResponseBodyDataSearchNumbers) SetValue(v int64) *GetSexRatioResponseBodyDataSearchNumbers {
	s.Value = &v
	return s
}

type GetSexRatioResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSexRatioResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSexRatioResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSexRatioResponse) GoString() string {
	return s.String()
}

func (s *GetSexRatioResponse) SetHeaders(v map[string]*string) *GetSexRatioResponse {
	s.Headers = v
	return s
}

func (s *GetSexRatioResponse) SetStatusCode(v int32) *GetSexRatioResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSexRatioResponse) SetBody(v *GetSexRatioResponseBody) *GetSexRatioResponse {
	s.Body = v
	return s
}

type GetStoreSalesVolumeTopRequest struct {
	BrandNames        *string `json:"BrandNames,omitempty" xml:"BrandNames,omitempty"`
	CateIds           *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	EndingDate        *string `json:"EndingDate,omitempty" xml:"EndingDate,omitempty"`
	EndingPrice       *string `json:"EndingPrice,omitempty" xml:"EndingPrice,omitempty"`
	EndingSalesVolume *string `json:"EndingSalesVolume,omitempty" xml:"EndingSalesVolume,omitempty"`
	PageIndex         *int64  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	ShopIds           *string `json:"ShopIds,omitempty" xml:"ShopIds,omitempty"`
	StartDate         *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	StartPrice        *string `json:"StartPrice,omitempty" xml:"StartPrice,omitempty"`
	StartSalesVolume  *string `json:"StartSalesVolume,omitempty" xml:"StartSalesVolume,omitempty"`
}

func (s GetStoreSalesVolumeTopRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStoreSalesVolumeTopRequest) GoString() string {
	return s.String()
}

func (s *GetStoreSalesVolumeTopRequest) SetBrandNames(v string) *GetStoreSalesVolumeTopRequest {
	s.BrandNames = &v
	return s
}

func (s *GetStoreSalesVolumeTopRequest) SetCateIds(v string) *GetStoreSalesVolumeTopRequest {
	s.CateIds = &v
	return s
}

func (s *GetStoreSalesVolumeTopRequest) SetEndingDate(v string) *GetStoreSalesVolumeTopRequest {
	s.EndingDate = &v
	return s
}

func (s *GetStoreSalesVolumeTopRequest) SetEndingPrice(v string) *GetStoreSalesVolumeTopRequest {
	s.EndingPrice = &v
	return s
}

func (s *GetStoreSalesVolumeTopRequest) SetEndingSalesVolume(v string) *GetStoreSalesVolumeTopRequest {
	s.EndingSalesVolume = &v
	return s
}

func (s *GetStoreSalesVolumeTopRequest) SetPageIndex(v int64) *GetStoreSalesVolumeTopRequest {
	s.PageIndex = &v
	return s
}

func (s *GetStoreSalesVolumeTopRequest) SetShopIds(v string) *GetStoreSalesVolumeTopRequest {
	s.ShopIds = &v
	return s
}

func (s *GetStoreSalesVolumeTopRequest) SetStartDate(v string) *GetStoreSalesVolumeTopRequest {
	s.StartDate = &v
	return s
}

func (s *GetStoreSalesVolumeTopRequest) SetStartPrice(v string) *GetStoreSalesVolumeTopRequest {
	s.StartPrice = &v
	return s
}

func (s *GetStoreSalesVolumeTopRequest) SetStartSalesVolume(v string) *GetStoreSalesVolumeTopRequest {
	s.StartSalesVolume = &v
	return s
}

type GetStoreSalesVolumeTopResponseBody struct {
	Code            *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            []*GetStoreSalesVolumeTopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message         *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool                                     `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetStoreSalesVolumeTopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStoreSalesVolumeTopResponseBody) GoString() string {
	return s.String()
}

func (s *GetStoreSalesVolumeTopResponseBody) SetCode(v string) *GetStoreSalesVolumeTopResponseBody {
	s.Code = &v
	return s
}

func (s *GetStoreSalesVolumeTopResponseBody) SetData(v []*GetStoreSalesVolumeTopResponseBodyData) *GetStoreSalesVolumeTopResponseBody {
	s.Data = v
	return s
}

func (s *GetStoreSalesVolumeTopResponseBody) SetMessage(v string) *GetStoreSalesVolumeTopResponseBody {
	s.Message = &v
	return s
}

func (s *GetStoreSalesVolumeTopResponseBody) SetRequestId(v string) *GetStoreSalesVolumeTopResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStoreSalesVolumeTopResponseBody) SetSuccessResponse(v bool) *GetStoreSalesVolumeTopResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetStoreSalesVolumeTopResponseBodyData struct {
	CommodityQuantity *int64   `json:"CommodityQuantity,omitempty" xml:"CommodityQuantity,omitempty"`
	SalesVolume       *int64   `json:"SalesVolume,omitempty" xml:"SalesVolume,omitempty"`
	SearchVolume      *int64   `json:"SearchVolume,omitempty" xml:"SearchVolume,omitempty"`
	ShopId            *int64   `json:"ShopId,omitempty" xml:"ShopId,omitempty"`
	ShopName          *string  `json:"ShopName,omitempty" xml:"ShopName,omitempty"`
	TotalSales        *float64 `json:"TotalSales,omitempty" xml:"TotalSales,omitempty"`
}

func (s GetStoreSalesVolumeTopResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetStoreSalesVolumeTopResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStoreSalesVolumeTopResponseBodyData) SetCommodityQuantity(v int64) *GetStoreSalesVolumeTopResponseBodyData {
	s.CommodityQuantity = &v
	return s
}

func (s *GetStoreSalesVolumeTopResponseBodyData) SetSalesVolume(v int64) *GetStoreSalesVolumeTopResponseBodyData {
	s.SalesVolume = &v
	return s
}

func (s *GetStoreSalesVolumeTopResponseBodyData) SetSearchVolume(v int64) *GetStoreSalesVolumeTopResponseBodyData {
	s.SearchVolume = &v
	return s
}

func (s *GetStoreSalesVolumeTopResponseBodyData) SetShopId(v int64) *GetStoreSalesVolumeTopResponseBodyData {
	s.ShopId = &v
	return s
}

func (s *GetStoreSalesVolumeTopResponseBodyData) SetShopName(v string) *GetStoreSalesVolumeTopResponseBodyData {
	s.ShopName = &v
	return s
}

func (s *GetStoreSalesVolumeTopResponseBodyData) SetTotalSales(v float64) *GetStoreSalesVolumeTopResponseBodyData {
	s.TotalSales = &v
	return s
}

type GetStoreSalesVolumeTopResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStoreSalesVolumeTopResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStoreSalesVolumeTopResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStoreSalesVolumeTopResponse) GoString() string {
	return s.String()
}

func (s *GetStoreSalesVolumeTopResponse) SetHeaders(v map[string]*string) *GetStoreSalesVolumeTopResponse {
	s.Headers = v
	return s
}

func (s *GetStoreSalesVolumeTopResponse) SetStatusCode(v int32) *GetStoreSalesVolumeTopResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStoreSalesVolumeTopResponse) SetBody(v *GetStoreSalesVolumeTopResponseBody) *GetStoreSalesVolumeTopResponse {
	s.Body = v
	return s
}

type GetStoreSearchTopRequest struct {
	BrandNames        *string `json:"BrandNames,omitempty" xml:"BrandNames,omitempty"`
	CateIds           *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	EndingDate        *string `json:"EndingDate,omitempty" xml:"EndingDate,omitempty"`
	EndingPrice       *string `json:"EndingPrice,omitempty" xml:"EndingPrice,omitempty"`
	EndingSalesVolume *string `json:"EndingSalesVolume,omitempty" xml:"EndingSalesVolume,omitempty"`
	PageIndex         *int64  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	ShopIds           *string `json:"ShopIds,omitempty" xml:"ShopIds,omitempty"`
	StartDate         *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	StartPrice        *string `json:"StartPrice,omitempty" xml:"StartPrice,omitempty"`
	StartSalesVolume  *string `json:"StartSalesVolume,omitempty" xml:"StartSalesVolume,omitempty"`
}

func (s GetStoreSearchTopRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStoreSearchTopRequest) GoString() string {
	return s.String()
}

func (s *GetStoreSearchTopRequest) SetBrandNames(v string) *GetStoreSearchTopRequest {
	s.BrandNames = &v
	return s
}

func (s *GetStoreSearchTopRequest) SetCateIds(v string) *GetStoreSearchTopRequest {
	s.CateIds = &v
	return s
}

func (s *GetStoreSearchTopRequest) SetEndingDate(v string) *GetStoreSearchTopRequest {
	s.EndingDate = &v
	return s
}

func (s *GetStoreSearchTopRequest) SetEndingPrice(v string) *GetStoreSearchTopRequest {
	s.EndingPrice = &v
	return s
}

func (s *GetStoreSearchTopRequest) SetEndingSalesVolume(v string) *GetStoreSearchTopRequest {
	s.EndingSalesVolume = &v
	return s
}

func (s *GetStoreSearchTopRequest) SetPageIndex(v int64) *GetStoreSearchTopRequest {
	s.PageIndex = &v
	return s
}

func (s *GetStoreSearchTopRequest) SetShopIds(v string) *GetStoreSearchTopRequest {
	s.ShopIds = &v
	return s
}

func (s *GetStoreSearchTopRequest) SetStartDate(v string) *GetStoreSearchTopRequest {
	s.StartDate = &v
	return s
}

func (s *GetStoreSearchTopRequest) SetStartPrice(v string) *GetStoreSearchTopRequest {
	s.StartPrice = &v
	return s
}

func (s *GetStoreSearchTopRequest) SetStartSalesVolume(v string) *GetStoreSearchTopRequest {
	s.StartSalesVolume = &v
	return s
}

type GetStoreSearchTopResponseBody struct {
	Code            *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            []*GetStoreSearchTopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message         *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool                                `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetStoreSearchTopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStoreSearchTopResponseBody) GoString() string {
	return s.String()
}

func (s *GetStoreSearchTopResponseBody) SetCode(v string) *GetStoreSearchTopResponseBody {
	s.Code = &v
	return s
}

func (s *GetStoreSearchTopResponseBody) SetData(v []*GetStoreSearchTopResponseBodyData) *GetStoreSearchTopResponseBody {
	s.Data = v
	return s
}

func (s *GetStoreSearchTopResponseBody) SetMessage(v string) *GetStoreSearchTopResponseBody {
	s.Message = &v
	return s
}

func (s *GetStoreSearchTopResponseBody) SetRequestId(v string) *GetStoreSearchTopResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStoreSearchTopResponseBody) SetSuccessResponse(v bool) *GetStoreSearchTopResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetStoreSearchTopResponseBodyData struct {
	CommodityQuantity *int64   `json:"CommodityQuantity,omitempty" xml:"CommodityQuantity,omitempty"`
	SalesVolume       *int64   `json:"SalesVolume,omitempty" xml:"SalesVolume,omitempty"`
	SearchVolume      *int64   `json:"SearchVolume,omitempty" xml:"SearchVolume,omitempty"`
	ShopId            *int64   `json:"ShopId,omitempty" xml:"ShopId,omitempty"`
	ShopName          *string  `json:"ShopName,omitempty" xml:"ShopName,omitempty"`
	TotalSales        *float64 `json:"TotalSales,omitempty" xml:"TotalSales,omitempty"`
}

func (s GetStoreSearchTopResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetStoreSearchTopResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStoreSearchTopResponseBodyData) SetCommodityQuantity(v int64) *GetStoreSearchTopResponseBodyData {
	s.CommodityQuantity = &v
	return s
}

func (s *GetStoreSearchTopResponseBodyData) SetSalesVolume(v int64) *GetStoreSearchTopResponseBodyData {
	s.SalesVolume = &v
	return s
}

func (s *GetStoreSearchTopResponseBodyData) SetSearchVolume(v int64) *GetStoreSearchTopResponseBodyData {
	s.SearchVolume = &v
	return s
}

func (s *GetStoreSearchTopResponseBodyData) SetShopId(v int64) *GetStoreSearchTopResponseBodyData {
	s.ShopId = &v
	return s
}

func (s *GetStoreSearchTopResponseBodyData) SetShopName(v string) *GetStoreSearchTopResponseBodyData {
	s.ShopName = &v
	return s
}

func (s *GetStoreSearchTopResponseBodyData) SetTotalSales(v float64) *GetStoreSearchTopResponseBodyData {
	s.TotalSales = &v
	return s
}

type GetStoreSearchTopResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStoreSearchTopResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStoreSearchTopResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStoreSearchTopResponse) GoString() string {
	return s.String()
}

func (s *GetStoreSearchTopResponse) SetHeaders(v map[string]*string) *GetStoreSearchTopResponse {
	s.Headers = v
	return s
}

func (s *GetStoreSearchTopResponse) SetStatusCode(v int32) *GetStoreSearchTopResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStoreSearchTopResponse) SetBody(v *GetStoreSearchTopResponseBody) *GetStoreSearchTopResponse {
	s.Body = v
	return s
}

type GetStyleOnlineRequest struct {
	BrandNames *string `json:"BrandNames,omitempty" xml:"BrandNames,omitempty"`
	CateIds    *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	EndingDate *string `json:"EndingDate,omitempty" xml:"EndingDate,omitempty"`
	PageIndex  *int64  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	ShopIds    *string `json:"ShopIds,omitempty" xml:"ShopIds,omitempty"`
	StartDate  *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetStyleOnlineRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStyleOnlineRequest) GoString() string {
	return s.String()
}

func (s *GetStyleOnlineRequest) SetBrandNames(v string) *GetStyleOnlineRequest {
	s.BrandNames = &v
	return s
}

func (s *GetStyleOnlineRequest) SetCateIds(v string) *GetStyleOnlineRequest {
	s.CateIds = &v
	return s
}

func (s *GetStyleOnlineRequest) SetEndingDate(v string) *GetStyleOnlineRequest {
	s.EndingDate = &v
	return s
}

func (s *GetStyleOnlineRequest) SetPageIndex(v int64) *GetStyleOnlineRequest {
	s.PageIndex = &v
	return s
}

func (s *GetStyleOnlineRequest) SetShopIds(v string) *GetStyleOnlineRequest {
	s.ShopIds = &v
	return s
}

func (s *GetStyleOnlineRequest) SetStartDate(v string) *GetStyleOnlineRequest {
	s.StartDate = &v
	return s
}

type GetStyleOnlineResponseBody struct {
	Code            *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            []*GetStyleOnlineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message         *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *string                           `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetStyleOnlineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStyleOnlineResponseBody) GoString() string {
	return s.String()
}

func (s *GetStyleOnlineResponseBody) SetCode(v string) *GetStyleOnlineResponseBody {
	s.Code = &v
	return s
}

func (s *GetStyleOnlineResponseBody) SetData(v []*GetStyleOnlineResponseBodyData) *GetStyleOnlineResponseBody {
	s.Data = v
	return s
}

func (s *GetStyleOnlineResponseBody) SetMessage(v string) *GetStyleOnlineResponseBody {
	s.Message = &v
	return s
}

func (s *GetStyleOnlineResponseBody) SetRequestId(v string) *GetStyleOnlineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStyleOnlineResponseBody) SetSuccessResponse(v string) *GetStyleOnlineResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetStyleOnlineResponseBodyData struct {
	AttributeContent *string   `json:"AttributeContent,omitempty" xml:"AttributeContent,omitempty"`
	BrandName        *string   `json:"BrandName,omitempty" xml:"BrandName,omitempty"`
	BuyerTags        *string   `json:"BuyerTags,omitempty" xml:"BuyerTags,omitempty"`
	CateName         *string   `json:"CateName,omitempty" xml:"CateName,omitempty"`
	Color            *string   `json:"Color,omitempty" xml:"Color,omitempty"`
	ExposureValue    *int64    `json:"ExposureValue,omitempty" xml:"ExposureValue,omitempty"`
	Images           []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	Material         *string   `json:"Material,omitempty" xml:"Material,omitempty"`
	Price            *float64  `json:"Price,omitempty" xml:"Price,omitempty"`
	ProductLink      *string   `json:"ProductLink,omitempty" xml:"ProductLink,omitempty"`
	Sales            *float64  `json:"Sales,omitempty" xml:"Sales,omitempty"`
	SalesVolume      *int64    `json:"SalesVolume,omitempty" xml:"SalesVolume,omitempty"`
	SearchVolume     *float64  `json:"SearchVolume,omitempty" xml:"SearchVolume,omitempty"`
	ShopId           *int64    `json:"ShopId,omitempty" xml:"ShopId,omitempty"`
	ShopName         *string   `json:"ShopName,omitempty" xml:"ShopName,omitempty"`
	Style            *string   `json:"Style,omitempty" xml:"Style,omitempty"`
	Title            *string   `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetStyleOnlineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetStyleOnlineResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStyleOnlineResponseBodyData) SetAttributeContent(v string) *GetStyleOnlineResponseBodyData {
	s.AttributeContent = &v
	return s
}

func (s *GetStyleOnlineResponseBodyData) SetBrandName(v string) *GetStyleOnlineResponseBodyData {
	s.BrandName = &v
	return s
}

func (s *GetStyleOnlineResponseBodyData) SetBuyerTags(v string) *GetStyleOnlineResponseBodyData {
	s.BuyerTags = &v
	return s
}

func (s *GetStyleOnlineResponseBodyData) SetCateName(v string) *GetStyleOnlineResponseBodyData {
	s.CateName = &v
	return s
}

func (s *GetStyleOnlineResponseBodyData) SetColor(v string) *GetStyleOnlineResponseBodyData {
	s.Color = &v
	return s
}

func (s *GetStyleOnlineResponseBodyData) SetExposureValue(v int64) *GetStyleOnlineResponseBodyData {
	s.ExposureValue = &v
	return s
}

func (s *GetStyleOnlineResponseBodyData) SetImages(v []*string) *GetStyleOnlineResponseBodyData {
	s.Images = v
	return s
}

func (s *GetStyleOnlineResponseBodyData) SetMaterial(v string) *GetStyleOnlineResponseBodyData {
	s.Material = &v
	return s
}

func (s *GetStyleOnlineResponseBodyData) SetPrice(v float64) *GetStyleOnlineResponseBodyData {
	s.Price = &v
	return s
}

func (s *GetStyleOnlineResponseBodyData) SetProductLink(v string) *GetStyleOnlineResponseBodyData {
	s.ProductLink = &v
	return s
}

func (s *GetStyleOnlineResponseBodyData) SetSales(v float64) *GetStyleOnlineResponseBodyData {
	s.Sales = &v
	return s
}

func (s *GetStyleOnlineResponseBodyData) SetSalesVolume(v int64) *GetStyleOnlineResponseBodyData {
	s.SalesVolume = &v
	return s
}

func (s *GetStyleOnlineResponseBodyData) SetSearchVolume(v float64) *GetStyleOnlineResponseBodyData {
	s.SearchVolume = &v
	return s
}

func (s *GetStyleOnlineResponseBodyData) SetShopId(v int64) *GetStyleOnlineResponseBodyData {
	s.ShopId = &v
	return s
}

func (s *GetStyleOnlineResponseBodyData) SetShopName(v string) *GetStyleOnlineResponseBodyData {
	s.ShopName = &v
	return s
}

func (s *GetStyleOnlineResponseBodyData) SetStyle(v string) *GetStyleOnlineResponseBodyData {
	s.Style = &v
	return s
}

func (s *GetStyleOnlineResponseBodyData) SetTitle(v string) *GetStyleOnlineResponseBodyData {
	s.Title = &v
	return s
}

type GetStyleOnlineResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStyleOnlineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStyleOnlineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStyleOnlineResponse) GoString() string {
	return s.String()
}

func (s *GetStyleOnlineResponse) SetHeaders(v map[string]*string) *GetStyleOnlineResponse {
	s.Headers = v
	return s
}

func (s *GetStyleOnlineResponse) SetStatusCode(v int32) *GetStyleOnlineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStyleOnlineResponse) SetBody(v *GetStyleOnlineResponseBody) *GetStyleOnlineResponse {
	s.Body = v
	return s
}

type GetStyleTopRequest struct {
	BrandNames  *string `json:"BrandNames,omitempty" xml:"BrandNames,omitempty"`
	CateIds     *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	EndingDate  *string `json:"EndingDate,omitempty" xml:"EndingDate,omitempty"`
	EndingPrice *string `json:"EndingPrice,omitempty" xml:"EndingPrice,omitempty"`
	PageIndex   *int64  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	ShopIds     *string `json:"ShopIds,omitempty" xml:"ShopIds,omitempty"`
	SortOrder   *int64  `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	StartDate   *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	StartPrice  *string `json:"StartPrice,omitempty" xml:"StartPrice,omitempty"`
	TimeDisplay *int64  `json:"TimeDisplay,omitempty" xml:"TimeDisplay,omitempty"`
}

func (s GetStyleTopRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStyleTopRequest) GoString() string {
	return s.String()
}

func (s *GetStyleTopRequest) SetBrandNames(v string) *GetStyleTopRequest {
	s.BrandNames = &v
	return s
}

func (s *GetStyleTopRequest) SetCateIds(v string) *GetStyleTopRequest {
	s.CateIds = &v
	return s
}

func (s *GetStyleTopRequest) SetEndingDate(v string) *GetStyleTopRequest {
	s.EndingDate = &v
	return s
}

func (s *GetStyleTopRequest) SetEndingPrice(v string) *GetStyleTopRequest {
	s.EndingPrice = &v
	return s
}

func (s *GetStyleTopRequest) SetPageIndex(v int64) *GetStyleTopRequest {
	s.PageIndex = &v
	return s
}

func (s *GetStyleTopRequest) SetShopIds(v string) *GetStyleTopRequest {
	s.ShopIds = &v
	return s
}

func (s *GetStyleTopRequest) SetSortOrder(v int64) *GetStyleTopRequest {
	s.SortOrder = &v
	return s
}

func (s *GetStyleTopRequest) SetStartDate(v string) *GetStyleTopRequest {
	s.StartDate = &v
	return s
}

func (s *GetStyleTopRequest) SetStartPrice(v string) *GetStyleTopRequest {
	s.StartPrice = &v
	return s
}

func (s *GetStyleTopRequest) SetTimeDisplay(v int64) *GetStyleTopRequest {
	s.TimeDisplay = &v
	return s
}

type GetStyleTopResponseBody struct {
	Code            *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            []*GetStyleTopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message         *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *string                        `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetStyleTopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStyleTopResponseBody) GoString() string {
	return s.String()
}

func (s *GetStyleTopResponseBody) SetCode(v string) *GetStyleTopResponseBody {
	s.Code = &v
	return s
}

func (s *GetStyleTopResponseBody) SetData(v []*GetStyleTopResponseBodyData) *GetStyleTopResponseBody {
	s.Data = v
	return s
}

func (s *GetStyleTopResponseBody) SetMessage(v string) *GetStyleTopResponseBody {
	s.Message = &v
	return s
}

func (s *GetStyleTopResponseBody) SetRequestId(v string) *GetStyleTopResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStyleTopResponseBody) SetSuccessResponse(v string) *GetStyleTopResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetStyleTopResponseBodyData struct {
	AttributeContent *string   `json:"AttributeContent,omitempty" xml:"AttributeContent,omitempty"`
	BrandName        *string   `json:"BrandName,omitempty" xml:"BrandName,omitempty"`
	BuyerTags        *string   `json:"BuyerTags,omitempty" xml:"BuyerTags,omitempty"`
	CateName         *string   `json:"CateName,omitempty" xml:"CateName,omitempty"`
	Color            *string   `json:"Color,omitempty" xml:"Color,omitempty"`
	ExposureValue    *int64    `json:"ExposureValue,omitempty" xml:"ExposureValue,omitempty"`
	Images           []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	Material         *string   `json:"Material,omitempty" xml:"Material,omitempty"`
	Price            *float64  `json:"Price,omitempty" xml:"Price,omitempty"`
	ProductLink      *string   `json:"ProductLink,omitempty" xml:"ProductLink,omitempty"`
	Sales            *float64  `json:"Sales,omitempty" xml:"Sales,omitempty"`
	SalesVolume      *int64    `json:"SalesVolume,omitempty" xml:"SalesVolume,omitempty"`
	SearchVolume     *float64  `json:"SearchVolume,omitempty" xml:"SearchVolume,omitempty"`
	ShopId           *int64    `json:"ShopId,omitempty" xml:"ShopId,omitempty"`
	ShopName         *string   `json:"ShopName,omitempty" xml:"ShopName,omitempty"`
	Style            *string   `json:"Style,omitempty" xml:"Style,omitempty"`
	Title            *string   `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetStyleTopResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetStyleTopResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStyleTopResponseBodyData) SetAttributeContent(v string) *GetStyleTopResponseBodyData {
	s.AttributeContent = &v
	return s
}

func (s *GetStyleTopResponseBodyData) SetBrandName(v string) *GetStyleTopResponseBodyData {
	s.BrandName = &v
	return s
}

func (s *GetStyleTopResponseBodyData) SetBuyerTags(v string) *GetStyleTopResponseBodyData {
	s.BuyerTags = &v
	return s
}

func (s *GetStyleTopResponseBodyData) SetCateName(v string) *GetStyleTopResponseBodyData {
	s.CateName = &v
	return s
}

func (s *GetStyleTopResponseBodyData) SetColor(v string) *GetStyleTopResponseBodyData {
	s.Color = &v
	return s
}

func (s *GetStyleTopResponseBodyData) SetExposureValue(v int64) *GetStyleTopResponseBodyData {
	s.ExposureValue = &v
	return s
}

func (s *GetStyleTopResponseBodyData) SetImages(v []*string) *GetStyleTopResponseBodyData {
	s.Images = v
	return s
}

func (s *GetStyleTopResponseBodyData) SetMaterial(v string) *GetStyleTopResponseBodyData {
	s.Material = &v
	return s
}

func (s *GetStyleTopResponseBodyData) SetPrice(v float64) *GetStyleTopResponseBodyData {
	s.Price = &v
	return s
}

func (s *GetStyleTopResponseBodyData) SetProductLink(v string) *GetStyleTopResponseBodyData {
	s.ProductLink = &v
	return s
}

func (s *GetStyleTopResponseBodyData) SetSales(v float64) *GetStyleTopResponseBodyData {
	s.Sales = &v
	return s
}

func (s *GetStyleTopResponseBodyData) SetSalesVolume(v int64) *GetStyleTopResponseBodyData {
	s.SalesVolume = &v
	return s
}

func (s *GetStyleTopResponseBodyData) SetSearchVolume(v float64) *GetStyleTopResponseBodyData {
	s.SearchVolume = &v
	return s
}

func (s *GetStyleTopResponseBodyData) SetShopId(v int64) *GetStyleTopResponseBodyData {
	s.ShopId = &v
	return s
}

func (s *GetStyleTopResponseBodyData) SetShopName(v string) *GetStyleTopResponseBodyData {
	s.ShopName = &v
	return s
}

func (s *GetStyleTopResponseBodyData) SetStyle(v string) *GetStyleTopResponseBodyData {
	s.Style = &v
	return s
}

func (s *GetStyleTopResponseBodyData) SetTitle(v string) *GetStyleTopResponseBodyData {
	s.Title = &v
	return s
}

type GetStyleTopResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStyleTopResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStyleTopResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStyleTopResponse) GoString() string {
	return s.String()
}

func (s *GetStyleTopResponse) SetHeaders(v map[string]*string) *GetStyleTopResponse {
	s.Headers = v
	return s
}

func (s *GetStyleTopResponse) SetStatusCode(v int32) *GetStyleTopResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStyleTopResponse) SetBody(v *GetStyleTopResponseBody) *GetStyleTopResponse {
	s.Body = v
	return s
}

type GetTrendImageDetailRequest struct {
	AiImgId *string `json:"AiImgId,omitempty" xml:"AiImgId,omitempty"`
}

func (s GetTrendImageDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrendImageDetailRequest) GoString() string {
	return s.String()
}

func (s *GetTrendImageDetailRequest) SetAiImgId(v string) *GetTrendImageDetailRequest {
	s.AiImgId = &v
	return s
}

type GetTrendImageDetailResponseBody struct {
	Code            *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            *GetTrendImageDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message         *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool                                `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetTrendImageDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrendImageDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrendImageDetailResponseBody) SetCode(v string) *GetTrendImageDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetTrendImageDetailResponseBody) SetData(v *GetTrendImageDetailResponseBodyData) *GetTrendImageDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetTrendImageDetailResponseBody) SetMessage(v string) *GetTrendImageDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetTrendImageDetailResponseBody) SetRequestId(v string) *GetTrendImageDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrendImageDetailResponseBody) SetSuccessResponse(v bool) *GetTrendImageDetailResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetTrendImageDetailResponseBodyData struct {
	AiImgId      *string                                            `json:"AiImgId,omitempty" xml:"AiImgId,omitempty"`
	AiImgUrl     *string                                            `json:"AiImgUrl,omitempty" xml:"AiImgUrl,omitempty"`
	MultiPictUrl *string                                            `json:"MultiPictUrl,omitempty" xml:"MultiPictUrl,omitempty"`
	Population   *string                                            `json:"Population,omitempty" xml:"Population,omitempty"`
	PriceMax     *int64                                             `json:"PriceMax,omitempty" xml:"PriceMax,omitempty"`
	PriceMin     *int64                                             `json:"PriceMin,omitempty" xml:"PriceMin,omitempty"`
	SimilarGoods []*GetTrendImageDetailResponseBodyDataSimilarGoods `json:"SimilarGoods,omitempty" xml:"SimilarGoods,omitempty" type:"Repeated"`
	Tags         *string                                            `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s GetTrendImageDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTrendImageDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrendImageDetailResponseBodyData) SetAiImgId(v string) *GetTrendImageDetailResponseBodyData {
	s.AiImgId = &v
	return s
}

func (s *GetTrendImageDetailResponseBodyData) SetAiImgUrl(v string) *GetTrendImageDetailResponseBodyData {
	s.AiImgUrl = &v
	return s
}

func (s *GetTrendImageDetailResponseBodyData) SetMultiPictUrl(v string) *GetTrendImageDetailResponseBodyData {
	s.MultiPictUrl = &v
	return s
}

func (s *GetTrendImageDetailResponseBodyData) SetPopulation(v string) *GetTrendImageDetailResponseBodyData {
	s.Population = &v
	return s
}

func (s *GetTrendImageDetailResponseBodyData) SetPriceMax(v int64) *GetTrendImageDetailResponseBodyData {
	s.PriceMax = &v
	return s
}

func (s *GetTrendImageDetailResponseBodyData) SetPriceMin(v int64) *GetTrendImageDetailResponseBodyData {
	s.PriceMin = &v
	return s
}

func (s *GetTrendImageDetailResponseBodyData) SetSimilarGoods(v []*GetTrendImageDetailResponseBodyDataSimilarGoods) *GetTrendImageDetailResponseBodyData {
	s.SimilarGoods = v
	return s
}

func (s *GetTrendImageDetailResponseBodyData) SetTags(v string) *GetTrendImageDetailResponseBodyData {
	s.Tags = &v
	return s
}

type GetTrendImageDetailResponseBodyDataSimilarGoods struct {
	AiImgUrl     *string `json:"AiImgUrl,omitempty" xml:"AiImgUrl,omitempty"`
	GoodsSales   *int64  `json:"GoodsSales,omitempty" xml:"GoodsSales,omitempty"`
	SearchVolume *int64  `json:"SearchVolume,omitempty" xml:"SearchVolume,omitempty"`
}

func (s GetTrendImageDetailResponseBodyDataSimilarGoods) String() string {
	return tea.Prettify(s)
}

func (s GetTrendImageDetailResponseBodyDataSimilarGoods) GoString() string {
	return s.String()
}

func (s *GetTrendImageDetailResponseBodyDataSimilarGoods) SetAiImgUrl(v string) *GetTrendImageDetailResponseBodyDataSimilarGoods {
	s.AiImgUrl = &v
	return s
}

func (s *GetTrendImageDetailResponseBodyDataSimilarGoods) SetGoodsSales(v int64) *GetTrendImageDetailResponseBodyDataSimilarGoods {
	s.GoodsSales = &v
	return s
}

func (s *GetTrendImageDetailResponseBodyDataSimilarGoods) SetSearchVolume(v int64) *GetTrendImageDetailResponseBodyDataSimilarGoods {
	s.SearchVolume = &v
	return s
}

type GetTrendImageDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTrendImageDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrendImageDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrendImageDetailResponse) GoString() string {
	return s.String()
}

func (s *GetTrendImageDetailResponse) SetHeaders(v map[string]*string) *GetTrendImageDetailResponse {
	s.Headers = v
	return s
}

func (s *GetTrendImageDetailResponse) SetStatusCode(v int32) *GetTrendImageDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrendImageDetailResponse) SetBody(v *GetTrendImageDetailResponseBody) *GetTrendImageDetailResponse {
	s.Body = v
	return s
}

type GetTrendImageListRequest struct {
	CateIds *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	Query   *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s GetTrendImageListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrendImageListRequest) GoString() string {
	return s.String()
}

func (s *GetTrendImageListRequest) SetCateIds(v string) *GetTrendImageListRequest {
	s.CateIds = &v
	return s
}

func (s *GetTrendImageListRequest) SetQuery(v string) *GetTrendImageListRequest {
	s.Query = &v
	return s
}

type GetTrendImageListResponseBody struct {
	Code            *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            []*GetTrendImageListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message         *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool                                `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetTrendImageListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrendImageListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrendImageListResponseBody) SetCode(v string) *GetTrendImageListResponseBody {
	s.Code = &v
	return s
}

func (s *GetTrendImageListResponseBody) SetData(v []*GetTrendImageListResponseBodyData) *GetTrendImageListResponseBody {
	s.Data = v
	return s
}

func (s *GetTrendImageListResponseBody) SetMessage(v string) *GetTrendImageListResponseBody {
	s.Message = &v
	return s
}

func (s *GetTrendImageListResponseBody) SetRequestId(v string) *GetTrendImageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrendImageListResponseBody) SetSuccessResponse(v bool) *GetTrendImageListResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetTrendImageListResponseBodyData struct {
	AiImgId    *string `json:"AiImgId,omitempty" xml:"AiImgId,omitempty"`
	AiImgUrl   *string `json:"AiImgUrl,omitempty" xml:"AiImgUrl,omitempty"`
	Population *string `json:"Population,omitempty" xml:"Population,omitempty"`
	PriceMax   *int64  `json:"PriceMax,omitempty" xml:"PriceMax,omitempty"`
	PriceMin   *int64  `json:"PriceMin,omitempty" xml:"PriceMin,omitempty"`
}

func (s GetTrendImageListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTrendImageListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrendImageListResponseBodyData) SetAiImgId(v string) *GetTrendImageListResponseBodyData {
	s.AiImgId = &v
	return s
}

func (s *GetTrendImageListResponseBodyData) SetAiImgUrl(v string) *GetTrendImageListResponseBodyData {
	s.AiImgUrl = &v
	return s
}

func (s *GetTrendImageListResponseBodyData) SetPopulation(v string) *GetTrendImageListResponseBodyData {
	s.Population = &v
	return s
}

func (s *GetTrendImageListResponseBodyData) SetPriceMax(v int64) *GetTrendImageListResponseBodyData {
	s.PriceMax = &v
	return s
}

func (s *GetTrendImageListResponseBodyData) SetPriceMin(v int64) *GetTrendImageListResponseBodyData {
	s.PriceMin = &v
	return s
}

type GetTrendImageListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTrendImageListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrendImageListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrendImageListResponse) GoString() string {
	return s.String()
}

func (s *GetTrendImageListResponse) SetHeaders(v map[string]*string) *GetTrendImageListResponse {
	s.Headers = v
	return s
}

func (s *GetTrendImageListResponse) SetStatusCode(v int32) *GetTrendImageListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrendImageListResponse) SetBody(v *GetTrendImageListResponseBody) *GetTrendImageListResponse {
	s.Body = v
	return s
}

type GetTrendIndexRequest struct {
	CateIds  *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	MonthNum *int32  `json:"MonthNum,omitempty" xml:"MonthNum,omitempty"`
}

func (s GetTrendIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrendIndexRequest) GoString() string {
	return s.String()
}

func (s *GetTrendIndexRequest) SetCateIds(v string) *GetTrendIndexRequest {
	s.CateIds = &v
	return s
}

func (s *GetTrendIndexRequest) SetMonthNum(v int32) *GetTrendIndexRequest {
	s.MonthNum = &v
	return s
}

type GetTrendIndexResponseBody struct {
	Code            *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            []*GetTrendIndexResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message         *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool                            `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetTrendIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrendIndexResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrendIndexResponseBody) SetCode(v string) *GetTrendIndexResponseBody {
	s.Code = &v
	return s
}

func (s *GetTrendIndexResponseBody) SetData(v []*GetTrendIndexResponseBodyData) *GetTrendIndexResponseBody {
	s.Data = v
	return s
}

func (s *GetTrendIndexResponseBody) SetMessage(v string) *GetTrendIndexResponseBody {
	s.Message = &v
	return s
}

func (s *GetTrendIndexResponseBody) SetRequestId(v string) *GetTrendIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrendIndexResponseBody) SetSuccessResponse(v bool) *GetTrendIndexResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetTrendIndexResponseBodyData struct {
	BrandIndex         *float32 `json:"BrandIndex,omitempty" xml:"BrandIndex,omitempty"`
	ECommerceIndex     *float32 `json:"ECommerceIndex,omitempty" xml:"ECommerceIndex,omitempty"`
	InstitutionalIndex *float32 `json:"InstitutionalIndex,omitempty" xml:"InstitutionalIndex,omitempty"`
	MediaIndex         *float32 `json:"MediaIndex,omitempty" xml:"MediaIndex,omitempty"`
	SocialIndex        *float32 `json:"SocialIndex,omitempty" xml:"SocialIndex,omitempty"`
	TrendIndex         *float32 `json:"TrendIndex,omitempty" xml:"TrendIndex,omitempty"`
	YearMonth          *string  `json:"YearMonth,omitempty" xml:"YearMonth,omitempty"`
}

func (s GetTrendIndexResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTrendIndexResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrendIndexResponseBodyData) SetBrandIndex(v float32) *GetTrendIndexResponseBodyData {
	s.BrandIndex = &v
	return s
}

func (s *GetTrendIndexResponseBodyData) SetECommerceIndex(v float32) *GetTrendIndexResponseBodyData {
	s.ECommerceIndex = &v
	return s
}

func (s *GetTrendIndexResponseBodyData) SetInstitutionalIndex(v float32) *GetTrendIndexResponseBodyData {
	s.InstitutionalIndex = &v
	return s
}

func (s *GetTrendIndexResponseBodyData) SetMediaIndex(v float32) *GetTrendIndexResponseBodyData {
	s.MediaIndex = &v
	return s
}

func (s *GetTrendIndexResponseBodyData) SetSocialIndex(v float32) *GetTrendIndexResponseBodyData {
	s.SocialIndex = &v
	return s
}

func (s *GetTrendIndexResponseBodyData) SetTrendIndex(v float32) *GetTrendIndexResponseBodyData {
	s.TrendIndex = &v
	return s
}

func (s *GetTrendIndexResponseBodyData) SetYearMonth(v string) *GetTrendIndexResponseBodyData {
	s.YearMonth = &v
	return s
}

type GetTrendIndexResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTrendIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrendIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrendIndexResponse) GoString() string {
	return s.String()
}

func (s *GetTrendIndexResponse) SetHeaders(v map[string]*string) *GetTrendIndexResponse {
	s.Headers = v
	return s
}

func (s *GetTrendIndexResponse) SetStatusCode(v int32) *GetTrendIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrendIndexResponse) SetBody(v *GetTrendIndexResponseBody) *GetTrendIndexResponse {
	s.Body = v
	return s
}

type GetTrendSearchRecordRequest struct {
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s GetTrendSearchRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrendSearchRecordRequest) GoString() string {
	return s.String()
}

func (s *GetTrendSearchRecordRequest) SetKey(v string) *GetTrendSearchRecordRequest {
	s.Key = &v
	return s
}

type GetTrendSearchRecordResponseBody struct {
	Code            *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            []*GetTrendSearchRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message         *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool                                   `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetTrendSearchRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrendSearchRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrendSearchRecordResponseBody) SetCode(v string) *GetTrendSearchRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetTrendSearchRecordResponseBody) SetData(v []*GetTrendSearchRecordResponseBodyData) *GetTrendSearchRecordResponseBody {
	s.Data = v
	return s
}

func (s *GetTrendSearchRecordResponseBody) SetMessage(v string) *GetTrendSearchRecordResponseBody {
	s.Message = &v
	return s
}

func (s *GetTrendSearchRecordResponseBody) SetRequestId(v string) *GetTrendSearchRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrendSearchRecordResponseBody) SetSuccessResponse(v bool) *GetTrendSearchRecordResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetTrendSearchRecordResponseBodyData struct {
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	QueryKey *string `json:"QueryKey,omitempty" xml:"QueryKey,omitempty"`
}

func (s GetTrendSearchRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTrendSearchRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrendSearchRecordResponseBodyData) SetId(v int64) *GetTrendSearchRecordResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetTrendSearchRecordResponseBodyData) SetQueryKey(v string) *GetTrendSearchRecordResponseBodyData {
	s.QueryKey = &v
	return s
}

type GetTrendSearchRecordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTrendSearchRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrendSearchRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrendSearchRecordResponse) GoString() string {
	return s.String()
}

func (s *GetTrendSearchRecordResponse) SetHeaders(v map[string]*string) *GetTrendSearchRecordResponse {
	s.Headers = v
	return s
}

func (s *GetTrendSearchRecordResponse) SetStatusCode(v int32) *GetTrendSearchRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrendSearchRecordResponse) SetBody(v *GetTrendSearchRecordResponseBody) *GetTrendSearchRecordResponse {
	s.Body = v
	return s
}

type GetTrendStatisticRequest struct {
	BrandNames *string `json:"BrandNames,omitempty" xml:"BrandNames,omitempty"`
	CateIds    *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	EndingDate *string `json:"EndingDate,omitempty" xml:"EndingDate,omitempty"`
	ShopIds    *string `json:"ShopIds,omitempty" xml:"ShopIds,omitempty"`
	StartDate  *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetTrendStatisticRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrendStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetTrendStatisticRequest) SetBrandNames(v string) *GetTrendStatisticRequest {
	s.BrandNames = &v
	return s
}

func (s *GetTrendStatisticRequest) SetCateIds(v string) *GetTrendStatisticRequest {
	s.CateIds = &v
	return s
}

func (s *GetTrendStatisticRequest) SetEndingDate(v string) *GetTrendStatisticRequest {
	s.EndingDate = &v
	return s
}

func (s *GetTrendStatisticRequest) SetShopIds(v string) *GetTrendStatisticRequest {
	s.ShopIds = &v
	return s
}

func (s *GetTrendStatisticRequest) SetStartDate(v string) *GetTrendStatisticRequest {
	s.StartDate = &v
	return s
}

type GetTrendStatisticResponseBody struct {
	Code            *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            *GetTrendStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message         *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *bool                              `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetTrendStatisticResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrendStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrendStatisticResponseBody) SetCode(v string) *GetTrendStatisticResponseBody {
	s.Code = &v
	return s
}

func (s *GetTrendStatisticResponseBody) SetData(v *GetTrendStatisticResponseBodyData) *GetTrendStatisticResponseBody {
	s.Data = v
	return s
}

func (s *GetTrendStatisticResponseBody) SetMessage(v string) *GetTrendStatisticResponseBody {
	s.Message = &v
	return s
}

func (s *GetTrendStatisticResponseBody) SetRequestId(v string) *GetTrendStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrendStatisticResponseBody) SetSuccessResponse(v bool) *GetTrendStatisticResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetTrendStatisticResponseBodyData struct {
	BrandCount     *int64   `json:"BrandCount,omitempty" xml:"BrandCount,omitempty"`
	CommodityCount *int64   `json:"CommodityCount,omitempty" xml:"CommodityCount,omitempty"`
	ExposureValue  *int64   `json:"ExposureValue,omitempty" xml:"ExposureValue,omitempty"`
	Hits           *int64   `json:"Hits,omitempty" xml:"Hits,omitempty"`
	Sales          *float64 `json:"Sales,omitempty" xml:"Sales,omitempty"`
	SalesVolume    *int64   `json:"SalesVolume,omitempty" xml:"SalesVolume,omitempty"`
	ShopCount      *int64   `json:"ShopCount,omitempty" xml:"ShopCount,omitempty"`
}

func (s GetTrendStatisticResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTrendStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrendStatisticResponseBodyData) SetBrandCount(v int64) *GetTrendStatisticResponseBodyData {
	s.BrandCount = &v
	return s
}

func (s *GetTrendStatisticResponseBodyData) SetCommodityCount(v int64) *GetTrendStatisticResponseBodyData {
	s.CommodityCount = &v
	return s
}

func (s *GetTrendStatisticResponseBodyData) SetExposureValue(v int64) *GetTrendStatisticResponseBodyData {
	s.ExposureValue = &v
	return s
}

func (s *GetTrendStatisticResponseBodyData) SetHits(v int64) *GetTrendStatisticResponseBodyData {
	s.Hits = &v
	return s
}

func (s *GetTrendStatisticResponseBodyData) SetSales(v float64) *GetTrendStatisticResponseBodyData {
	s.Sales = &v
	return s
}

func (s *GetTrendStatisticResponseBodyData) SetSalesVolume(v int64) *GetTrendStatisticResponseBodyData {
	s.SalesVolume = &v
	return s
}

func (s *GetTrendStatisticResponseBodyData) SetShopCount(v int64) *GetTrendStatisticResponseBodyData {
	s.ShopCount = &v
	return s
}

type GetTrendStatisticResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTrendStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrendStatisticResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrendStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetTrendStatisticResponse) SetHeaders(v map[string]*string) *GetTrendStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetTrendStatisticResponse) SetStatusCode(v int32) *GetTrendStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrendStatisticResponse) SetBody(v *GetTrendStatisticResponseBody) *GetTrendStatisticResponse {
	s.Body = v
	return s
}

type GetTrendWordAndIndexRequest struct {
	CateIds    *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	EndingDate *string `json:"EndingDate,omitempty" xml:"EndingDate,omitempty"`
	PageIndex  *int64  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	StartDate  *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetTrendWordAndIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrendWordAndIndexRequest) GoString() string {
	return s.String()
}

func (s *GetTrendWordAndIndexRequest) SetCateIds(v string) *GetTrendWordAndIndexRequest {
	s.CateIds = &v
	return s
}

func (s *GetTrendWordAndIndexRequest) SetEndingDate(v string) *GetTrendWordAndIndexRequest {
	s.EndingDate = &v
	return s
}

func (s *GetTrendWordAndIndexRequest) SetPageIndex(v int64) *GetTrendWordAndIndexRequest {
	s.PageIndex = &v
	return s
}

func (s *GetTrendWordAndIndexRequest) SetStartDate(v string) *GetTrendWordAndIndexRequest {
	s.StartDate = &v
	return s
}

type GetTrendWordAndIndexResponseBody struct {
	Code            *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data            []*GetTrendWordAndIndexResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Massage         *string                                 `json:"Massage,omitempty" xml:"Massage,omitempty"`
	RequestId       *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessResponse *string                                 `json:"SuccessResponse,omitempty" xml:"SuccessResponse,omitempty"`
}

func (s GetTrendWordAndIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrendWordAndIndexResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrendWordAndIndexResponseBody) SetCode(v string) *GetTrendWordAndIndexResponseBody {
	s.Code = &v
	return s
}

func (s *GetTrendWordAndIndexResponseBody) SetData(v []*GetTrendWordAndIndexResponseBodyData) *GetTrendWordAndIndexResponseBody {
	s.Data = v
	return s
}

func (s *GetTrendWordAndIndexResponseBody) SetMassage(v string) *GetTrendWordAndIndexResponseBody {
	s.Massage = &v
	return s
}

func (s *GetTrendWordAndIndexResponseBody) SetRequestId(v string) *GetTrendWordAndIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTrendWordAndIndexResponseBody) SetSuccessResponse(v string) *GetTrendWordAndIndexResponseBody {
	s.SuccessResponse = &v
	return s
}

type GetTrendWordAndIndexResponseBodyData struct {
	BrandIndex         *float32 `json:"BrandIndex,omitempty" xml:"BrandIndex,omitempty"`
	CateId             *int64   `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName           *string  `json:"CateName,omitempty" xml:"CateName,omitempty"`
	EcommerceIndex     *float32 `json:"EcommerceIndex,omitempty" xml:"EcommerceIndex,omitempty"`
	InstitutionalIndex *float32 `json:"InstitutionalIndex,omitempty" xml:"InstitutionalIndex,omitempty"`
	MediaIndex         *float32 `json:"MediaIndex,omitempty" xml:"MediaIndex,omitempty"`
	SocialIndex        *float32 `json:"SocialIndex,omitempty" xml:"SocialIndex,omitempty"`
	TrendIndex         *float32 `json:"TrendIndex,omitempty" xml:"TrendIndex,omitempty"`
	TrendWord          *string  `json:"TrendWord,omitempty" xml:"TrendWord,omitempty"`
	YearMonth          *string  `json:"YearMonth,omitempty" xml:"YearMonth,omitempty"`
}

func (s GetTrendWordAndIndexResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTrendWordAndIndexResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrendWordAndIndexResponseBodyData) SetBrandIndex(v float32) *GetTrendWordAndIndexResponseBodyData {
	s.BrandIndex = &v
	return s
}

func (s *GetTrendWordAndIndexResponseBodyData) SetCateId(v int64) *GetTrendWordAndIndexResponseBodyData {
	s.CateId = &v
	return s
}

func (s *GetTrendWordAndIndexResponseBodyData) SetCateName(v string) *GetTrendWordAndIndexResponseBodyData {
	s.CateName = &v
	return s
}

func (s *GetTrendWordAndIndexResponseBodyData) SetEcommerceIndex(v float32) *GetTrendWordAndIndexResponseBodyData {
	s.EcommerceIndex = &v
	return s
}

func (s *GetTrendWordAndIndexResponseBodyData) SetInstitutionalIndex(v float32) *GetTrendWordAndIndexResponseBodyData {
	s.InstitutionalIndex = &v
	return s
}

func (s *GetTrendWordAndIndexResponseBodyData) SetMediaIndex(v float32) *GetTrendWordAndIndexResponseBodyData {
	s.MediaIndex = &v
	return s
}

func (s *GetTrendWordAndIndexResponseBodyData) SetSocialIndex(v float32) *GetTrendWordAndIndexResponseBodyData {
	s.SocialIndex = &v
	return s
}

func (s *GetTrendWordAndIndexResponseBodyData) SetTrendIndex(v float32) *GetTrendWordAndIndexResponseBodyData {
	s.TrendIndex = &v
	return s
}

func (s *GetTrendWordAndIndexResponseBodyData) SetTrendWord(v string) *GetTrendWordAndIndexResponseBodyData {
	s.TrendWord = &v
	return s
}

func (s *GetTrendWordAndIndexResponseBodyData) SetYearMonth(v string) *GetTrendWordAndIndexResponseBodyData {
	s.YearMonth = &v
	return s
}

type GetTrendWordAndIndexResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTrendWordAndIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrendWordAndIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrendWordAndIndexResponse) GoString() string {
	return s.String()
}

func (s *GetTrendWordAndIndexResponse) SetHeaders(v map[string]*string) *GetTrendWordAndIndexResponse {
	s.Headers = v
	return s
}

func (s *GetTrendWordAndIndexResponse) SetStatusCode(v int32) *GetTrendWordAndIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrendWordAndIndexResponse) SetBody(v *GetTrendWordAndIndexResponseBody) *GetTrendWordAndIndexResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("qssj"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgeDistributionWithOptions(request *GetAgeDistributionRequest, runtime *util.RuntimeOptions) (_result *GetAgeDistributionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CateIds)) {
		query["CateIds"] = request.CateIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndingDate)) {
		query["EndingDate"] = request.EndingDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAgeDistribution"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAgeDistributionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgeDistribution(request *GetAgeDistributionRequest) (_result *GetAgeDistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgeDistributionResponse{}
	_body, _err := client.GetAgeDistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllTrendCategoryWithOptions(runtime *util.RuntimeOptions) (_result *GetAllTrendCategoryResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetAllTrendCategory"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAllTrendCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllTrendCategory() (_result *GetAllTrendCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAllTrendCategoryResponse{}
	_body, _err := client.GetAllTrendCategoryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCrowdDataWithOptions(request *GetCrowdDataRequest, runtime *util.RuntimeOptions) (_result *GetCrowdDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CateIds)) {
		body["CateIds"] = request.CateIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndingDate)) {
		body["EndingDate"] = request.EndingDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCrowdData"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCrowdDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCrowdData(request *GetCrowdDataRequest) (_result *GetCrowdDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCrowdDataResponse{}
	_body, _err := client.GetCrowdDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCrowdLabelWithOptions(request *GetCrowdLabelRequest, runtime *util.RuntimeOptions) (_result *GetCrowdLabelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CateIds)) {
		body["CateIds"] = request.CateIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndingDate)) {
		body["EndingDate"] = request.EndingDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCrowdLabel"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCrowdLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCrowdLabel(request *GetCrowdLabelRequest) (_result *GetCrowdLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCrowdLabelResponse{}
	_body, _err := client.GetCrowdLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCrowdReginWithOptions(request *GetCrowdReginRequest, runtime *util.RuntimeOptions) (_result *GetCrowdReginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CateIds)) {
		query["CateIds"] = request.CateIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndingDate)) {
		query["EndingDate"] = request.EndingDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCrowdRegin"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCrowdReginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCrowdRegin(request *GetCrowdReginRequest) (_result *GetCrowdReginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCrowdReginResponse{}
	_body, _err := client.GetCrowdReginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOpportunityMarketWithOptions(request *GetOpportunityMarketRequest, runtime *util.RuntimeOptions) (_result *GetOpportunityMarketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CateIds)) {
		body["CateIds"] = request.CateIds
	}

	if !tea.BoolValue(util.IsUnset(request.TimeDisplay)) {
		body["TimeDisplay"] = request.TimeDisplay
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOpportunityMarket"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOpportunityMarketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOpportunityMarket(request *GetOpportunityMarketRequest) (_result *GetOpportunityMarketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOpportunityMarketResponse{}
	_body, _err := client.GetOpportunityMarketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPriceRangeWithOptions(request *GetPriceRangeRequest, runtime *util.RuntimeOptions) (_result *GetPriceRangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BrandNames)) {
		body["BrandNames"] = request.BrandNames
	}

	if !tea.BoolValue(util.IsUnset(request.CateIds)) {
		body["CateIds"] = request.CateIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndingDate)) {
		body["EndingDate"] = request.EndingDate
	}

	if !tea.BoolValue(util.IsUnset(request.Section)) {
		body["Section"] = request.Section
	}

	if !tea.BoolValue(util.IsUnset(request.ShopIds)) {
		body["ShopIds"] = request.ShopIds
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPriceRange"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPriceRangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPriceRange(request *GetPriceRangeRequest) (_result *GetPriceRangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPriceRangeResponse{}
	_body, _err := client.GetPriceRangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSexRatioWithOptions(request *GetSexRatioRequest, runtime *util.RuntimeOptions) (_result *GetSexRatioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CateIds)) {
		query["CateIds"] = request.CateIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndingDate)) {
		query["EndingDate"] = request.EndingDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSexRatio"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSexRatioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSexRatio(request *GetSexRatioRequest) (_result *GetSexRatioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSexRatioResponse{}
	_body, _err := client.GetSexRatioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStoreSalesVolumeTopWithOptions(request *GetStoreSalesVolumeTopRequest, runtime *util.RuntimeOptions) (_result *GetStoreSalesVolumeTopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BrandNames)) {
		body["BrandNames"] = request.BrandNames
	}

	if !tea.BoolValue(util.IsUnset(request.CateIds)) {
		body["CateIds"] = request.CateIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndingDate)) {
		body["EndingDate"] = request.EndingDate
	}

	if !tea.BoolValue(util.IsUnset(request.EndingPrice)) {
		body["EndingPrice"] = request.EndingPrice
	}

	if !tea.BoolValue(util.IsUnset(request.EndingSalesVolume)) {
		body["EndingSalesVolume"] = request.EndingSalesVolume
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.ShopIds)) {
		body["ShopIds"] = request.ShopIds
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartPrice)) {
		body["StartPrice"] = request.StartPrice
	}

	if !tea.BoolValue(util.IsUnset(request.StartSalesVolume)) {
		body["StartSalesVolume"] = request.StartSalesVolume
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStoreSalesVolumeTop"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStoreSalesVolumeTopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStoreSalesVolumeTop(request *GetStoreSalesVolumeTopRequest) (_result *GetStoreSalesVolumeTopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStoreSalesVolumeTopResponse{}
	_body, _err := client.GetStoreSalesVolumeTopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStoreSearchTopWithOptions(request *GetStoreSearchTopRequest, runtime *util.RuntimeOptions) (_result *GetStoreSearchTopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BrandNames)) {
		body["BrandNames"] = request.BrandNames
	}

	if !tea.BoolValue(util.IsUnset(request.CateIds)) {
		body["CateIds"] = request.CateIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndingDate)) {
		body["EndingDate"] = request.EndingDate
	}

	if !tea.BoolValue(util.IsUnset(request.EndingPrice)) {
		body["EndingPrice"] = request.EndingPrice
	}

	if !tea.BoolValue(util.IsUnset(request.EndingSalesVolume)) {
		body["EndingSalesVolume"] = request.EndingSalesVolume
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.ShopIds)) {
		body["ShopIds"] = request.ShopIds
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartPrice)) {
		body["StartPrice"] = request.StartPrice
	}

	if !tea.BoolValue(util.IsUnset(request.StartSalesVolume)) {
		body["StartSalesVolume"] = request.StartSalesVolume
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStoreSearchTop"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStoreSearchTopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStoreSearchTop(request *GetStoreSearchTopRequest) (_result *GetStoreSearchTopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStoreSearchTopResponse{}
	_body, _err := client.GetStoreSearchTopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStyleOnlineWithOptions(request *GetStyleOnlineRequest, runtime *util.RuntimeOptions) (_result *GetStyleOnlineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BrandNames)) {
		query["BrandNames"] = request.BrandNames
	}

	if !tea.BoolValue(util.IsUnset(request.CateIds)) {
		query["CateIds"] = request.CateIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndingDate)) {
		query["EndingDate"] = request.EndingDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.ShopIds)) {
		query["ShopIds"] = request.ShopIds
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStyleOnline"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStyleOnlineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStyleOnline(request *GetStyleOnlineRequest) (_result *GetStyleOnlineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStyleOnlineResponse{}
	_body, _err := client.GetStyleOnlineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStyleTopWithOptions(request *GetStyleTopRequest, runtime *util.RuntimeOptions) (_result *GetStyleTopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BrandNames)) {
		body["BrandNames"] = request.BrandNames
	}

	if !tea.BoolValue(util.IsUnset(request.CateIds)) {
		body["CateIds"] = request.CateIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndingDate)) {
		body["EndingDate"] = request.EndingDate
	}

	if !tea.BoolValue(util.IsUnset(request.EndingPrice)) {
		body["EndingPrice"] = request.EndingPrice
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.ShopIds)) {
		body["ShopIds"] = request.ShopIds
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		body["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartPrice)) {
		body["StartPrice"] = request.StartPrice
	}

	if !tea.BoolValue(util.IsUnset(request.TimeDisplay)) {
		body["TimeDisplay"] = request.TimeDisplay
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStyleTop"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStyleTopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStyleTop(request *GetStyleTopRequest) (_result *GetStyleTopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStyleTopResponse{}
	_body, _err := client.GetStyleTopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrendImageDetailWithOptions(request *GetTrendImageDetailRequest, runtime *util.RuntimeOptions) (_result *GetTrendImageDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AiImgId)) {
		query["AiImgId"] = request.AiImgId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrendImageDetail"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTrendImageDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrendImageDetail(request *GetTrendImageDetailRequest) (_result *GetTrendImageDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTrendImageDetailResponse{}
	_body, _err := client.GetTrendImageDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrendImageListWithOptions(request *GetTrendImageListRequest, runtime *util.RuntimeOptions) (_result *GetTrendImageListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CateIds)) {
		body["CateIds"] = request.CateIds
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["Query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrendImageList"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTrendImageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrendImageList(request *GetTrendImageListRequest) (_result *GetTrendImageListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTrendImageListResponse{}
	_body, _err := client.GetTrendImageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrendIndexWithOptions(request *GetTrendIndexRequest, runtime *util.RuntimeOptions) (_result *GetTrendIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CateIds)) {
		body["CateIds"] = request.CateIds
	}

	if !tea.BoolValue(util.IsUnset(request.MonthNum)) {
		body["MonthNum"] = request.MonthNum
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrendIndex"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTrendIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrendIndex(request *GetTrendIndexRequest) (_result *GetTrendIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTrendIndexResponse{}
	_body, _err := client.GetTrendIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrendSearchRecordWithOptions(request *GetTrendSearchRecordRequest, runtime *util.RuntimeOptions) (_result *GetTrendSearchRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		body["Key"] = request.Key
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrendSearchRecord"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTrendSearchRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrendSearchRecord(request *GetTrendSearchRecordRequest) (_result *GetTrendSearchRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTrendSearchRecordResponse{}
	_body, _err := client.GetTrendSearchRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrendStatisticWithOptions(request *GetTrendStatisticRequest, runtime *util.RuntimeOptions) (_result *GetTrendStatisticResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BrandNames)) {
		body["BrandNames"] = request.BrandNames
	}

	if !tea.BoolValue(util.IsUnset(request.CateIds)) {
		body["CateIds"] = request.CateIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndingDate)) {
		body["EndingDate"] = request.EndingDate
	}

	if !tea.BoolValue(util.IsUnset(request.ShopIds)) {
		body["ShopIds"] = request.ShopIds
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrendStatistic"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTrendStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrendStatistic(request *GetTrendStatisticRequest) (_result *GetTrendStatisticResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTrendStatisticResponse{}
	_body, _err := client.GetTrendStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrendWordAndIndexWithOptions(request *GetTrendWordAndIndexRequest, runtime *util.RuntimeOptions) (_result *GetTrendWordAndIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CateIds)) {
		body["CateIds"] = request.CateIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndingDate)) {
		body["EndingDate"] = request.EndingDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrendWordAndIndex"),
		Version:     tea.String("2022-01-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTrendWordAndIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrendWordAndIndex(request *GetTrendWordAndIndexRequest) (_result *GetTrendWordAndIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTrendWordAndIndexResponse{}
	_body, _err := client.GetTrendWordAndIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
