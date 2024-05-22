// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ActivateApDeviceRequest struct {
	// This parameter is required.
	ApMac *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s ActivateApDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateApDeviceRequest) GoString() string {
	return s.String()
}

func (s *ActivateApDeviceRequest) SetApMac(v string) *ActivateApDeviceRequest {
	s.ApMac = &v
	return s
}

func (s *ActivateApDeviceRequest) SetStoreId(v string) *ActivateApDeviceRequest {
	s.StoreId = &v
	return s
}

type ActivateApDeviceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ActivateApDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateApDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateApDeviceResponseBody) SetCode(v string) *ActivateApDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *ActivateApDeviceResponseBody) SetDynamicCode(v string) *ActivateApDeviceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ActivateApDeviceResponseBody) SetDynamicMessage(v string) *ActivateApDeviceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ActivateApDeviceResponseBody) SetErrorCode(v string) *ActivateApDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ActivateApDeviceResponseBody) SetErrorMessage(v string) *ActivateApDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ActivateApDeviceResponseBody) SetMessage(v string) *ActivateApDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *ActivateApDeviceResponseBody) SetRequestId(v string) *ActivateApDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateApDeviceResponseBody) SetSuccess(v bool) *ActivateApDeviceResponseBody {
	s.Success = &v
	return s
}

type ActivateApDeviceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateApDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateApDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateApDeviceResponse) GoString() string {
	return s.String()
}

func (s *ActivateApDeviceResponse) SetHeaders(v map[string]*string) *ActivateApDeviceResponse {
	s.Headers = v
	return s
}

func (s *ActivateApDeviceResponse) SetStatusCode(v int32) *ActivateApDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateApDeviceResponse) SetBody(v *ActivateApDeviceResponseBody) *ActivateApDeviceResponse {
	s.Body = v
	return s
}

type AddApDeviceRequest struct {
	// This parameter is required.
	ApMac *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s AddApDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddApDeviceRequest) GoString() string {
	return s.String()
}

func (s *AddApDeviceRequest) SetApMac(v string) *AddApDeviceRequest {
	s.ApMac = &v
	return s
}

func (s *AddApDeviceRequest) SetStoreId(v string) *AddApDeviceRequest {
	s.StoreId = &v
	return s
}

type AddApDeviceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddApDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddApDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *AddApDeviceResponseBody) SetCode(v string) *AddApDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *AddApDeviceResponseBody) SetDynamicCode(v string) *AddApDeviceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AddApDeviceResponseBody) SetDynamicMessage(v string) *AddApDeviceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AddApDeviceResponseBody) SetErrorCode(v string) *AddApDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddApDeviceResponseBody) SetErrorMessage(v string) *AddApDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddApDeviceResponseBody) SetMessage(v string) *AddApDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *AddApDeviceResponseBody) SetRequestId(v string) *AddApDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddApDeviceResponseBody) SetSuccess(v bool) *AddApDeviceResponseBody {
	s.Success = &v
	return s
}

type AddApDeviceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddApDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddApDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddApDeviceResponse) GoString() string {
	return s.String()
}

func (s *AddApDeviceResponse) SetHeaders(v map[string]*string) *AddApDeviceResponse {
	s.Headers = v
	return s
}

func (s *AddApDeviceResponse) SetStatusCode(v int32) *AddApDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddApDeviceResponse) SetBody(v *AddApDeviceResponseBody) *AddApDeviceResponse {
	s.Body = v
	return s
}

type AddEslDeviceRequest struct {
	// This parameter is required.
	EslBarCode *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s AddEslDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEslDeviceRequest) GoString() string {
	return s.String()
}

func (s *AddEslDeviceRequest) SetEslBarCode(v string) *AddEslDeviceRequest {
	s.EslBarCode = &v
	return s
}

func (s *AddEslDeviceRequest) SetStoreId(v string) *AddEslDeviceRequest {
	s.StoreId = &v
	return s
}

type AddEslDeviceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddEslDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddEslDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *AddEslDeviceResponseBody) SetCode(v string) *AddEslDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *AddEslDeviceResponseBody) SetDynamicCode(v string) *AddEslDeviceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AddEslDeviceResponseBody) SetDynamicMessage(v string) *AddEslDeviceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AddEslDeviceResponseBody) SetErrorCode(v string) *AddEslDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddEslDeviceResponseBody) SetErrorMessage(v string) *AddEslDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddEslDeviceResponseBody) SetMessage(v string) *AddEslDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *AddEslDeviceResponseBody) SetRequestId(v string) *AddEslDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEslDeviceResponseBody) SetSuccess(v bool) *AddEslDeviceResponseBody {
	s.Success = &v
	return s
}

type AddEslDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddEslDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddEslDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddEslDeviceResponse) GoString() string {
	return s.String()
}

func (s *AddEslDeviceResponse) SetHeaders(v map[string]*string) *AddEslDeviceResponse {
	s.Headers = v
	return s
}

func (s *AddEslDeviceResponse) SetStatusCode(v int32) *AddEslDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddEslDeviceResponse) SetBody(v *AddEslDeviceResponseBody) *AddEslDeviceResponse {
	s.Body = v
	return s
}

type AddUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// user1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserRequest) GoString() string {
	return s.String()
}

func (s *AddUserRequest) SetUserId(v string) *AddUserRequest {
	s.UserId = &v
	return s
}

type AddUserResponseBody struct {
	// example:
	//
	// -1001
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// PlatformResponseError.%s
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// The specified store %s does not exist.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// MandatoryParameters
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// The specified resource type is invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E69C8998-1787-4999-8C75-D663FF1173CF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserResponseBody) SetCode(v string) *AddUserResponseBody {
	s.Code = &v
	return s
}

func (s *AddUserResponseBody) SetDynamicCode(v string) *AddUserResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AddUserResponseBody) SetDynamicMessage(v string) *AddUserResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AddUserResponseBody) SetErrorCode(v string) *AddUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddUserResponseBody) SetErrorMessage(v string) *AddUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddUserResponseBody) SetMessage(v string) *AddUserResponseBody {
	s.Message = &v
	return s
}

func (s *AddUserResponseBody) SetRequestId(v string) *AddUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserResponseBody) SetSuccess(v bool) *AddUserResponseBody {
	s.Success = &v
	return s
}

type AddUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserResponse) GoString() string {
	return s.String()
}

func (s *AddUserResponse) SetHeaders(v map[string]*string) *AddUserResponse {
	s.Headers = v
	return s
}

func (s *AddUserResponse) SetStatusCode(v int32) *AddUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserResponse) SetBody(v *AddUserResponseBody) *AddUserResponse {
	s.Body = v
	return s
}

type AssignUserRequest struct {
	Stores *string `json:"Stores,omitempty" xml:"Stores,omitempty"`
	// This parameter is required.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// This parameter is required.
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s AssignUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AssignUserRequest) GoString() string {
	return s.String()
}

func (s *AssignUserRequest) SetStores(v string) *AssignUserRequest {
	s.Stores = &v
	return s
}

func (s *AssignUserRequest) SetUserId(v string) *AssignUserRequest {
	s.UserId = &v
	return s
}

func (s *AssignUserRequest) SetUserType(v string) *AssignUserRequest {
	s.UserType = &v
	return s
}

type AssignUserResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AssignUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssignUserResponseBody) GoString() string {
	return s.String()
}

func (s *AssignUserResponseBody) SetCode(v string) *AssignUserResponseBody {
	s.Code = &v
	return s
}

func (s *AssignUserResponseBody) SetDynamicCode(v string) *AssignUserResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AssignUserResponseBody) SetDynamicMessage(v string) *AssignUserResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AssignUserResponseBody) SetErrorCode(v string) *AssignUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AssignUserResponseBody) SetErrorMessage(v string) *AssignUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AssignUserResponseBody) SetMessage(v string) *AssignUserResponseBody {
	s.Message = &v
	return s
}

func (s *AssignUserResponseBody) SetRequestId(v string) *AssignUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignUserResponseBody) SetSuccess(v bool) *AssignUserResponseBody {
	s.Success = &v
	return s
}

type AssignUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignUserResponse) String() string {
	return tea.Prettify(s)
}

func (s AssignUserResponse) GoString() string {
	return s.String()
}

func (s *AssignUserResponse) SetHeaders(v map[string]*string) *AssignUserResponse {
	s.Headers = v
	return s
}

func (s *AssignUserResponse) SetStatusCode(v int32) *AssignUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignUserResponse) SetBody(v *AssignUserResponseBody) *AssignUserResponse {
	s.Body = v
	return s
}

type BatchInsertItemsRequest struct {
	// This parameter is required.
	ItemInfo []*BatchInsertItemsRequestItemInfo `json:"ItemInfo,omitempty" xml:"ItemInfo,omitempty" type:"Repeated"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s BatchInsertItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertItemsRequest) GoString() string {
	return s.String()
}

func (s *BatchInsertItemsRequest) SetItemInfo(v []*BatchInsertItemsRequestItemInfo) *BatchInsertItemsRequest {
	s.ItemInfo = v
	return s
}

func (s *BatchInsertItemsRequest) SetStoreId(v string) *BatchInsertItemsRequest {
	s.StoreId = &v
	return s
}

type BatchInsertItemsRequestItemInfo struct {
	ActionPrice       *int32  `json:"ActionPrice,omitempty" xml:"ActionPrice,omitempty"`
	BePromotion       *bool   `json:"BePromotion,omitempty" xml:"BePromotion,omitempty"`
	BeSourceCode      *bool   `json:"BeSourceCode,omitempty" xml:"BeSourceCode,omitempty"`
	BrandName         *string `json:"BrandName,omitempty" xml:"BrandName,omitempty"`
	CategoryName      *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	CompanyId         *string `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	CustomizeFeatureA *string `json:"CustomizeFeatureA,omitempty" xml:"CustomizeFeatureA,omitempty"`
	CustomizeFeatureB *string `json:"CustomizeFeatureB,omitempty" xml:"CustomizeFeatureB,omitempty"`
	CustomizeFeatureC *string `json:"CustomizeFeatureC,omitempty" xml:"CustomizeFeatureC,omitempty"`
	CustomizeFeatureD *string `json:"CustomizeFeatureD,omitempty" xml:"CustomizeFeatureD,omitempty"`
	CustomizeFeatureE *string `json:"CustomizeFeatureE,omitempty" xml:"CustomizeFeatureE,omitempty"`
	CustomizeFeatureF *string `json:"CustomizeFeatureF,omitempty" xml:"CustomizeFeatureF,omitempty"`
	CustomizeFeatureG *string `json:"CustomizeFeatureG,omitempty" xml:"CustomizeFeatureG,omitempty"`
	CustomizeFeatureH *string `json:"CustomizeFeatureH,omitempty" xml:"CustomizeFeatureH,omitempty"`
	CustomizeFeatureI *string `json:"CustomizeFeatureI,omitempty" xml:"CustomizeFeatureI,omitempty"`
	CustomizeFeatureJ *string `json:"CustomizeFeatureJ,omitempty" xml:"CustomizeFeatureJ,omitempty"`
	EnergyEfficiency  *string `json:"EnergyEfficiency,omitempty" xml:"EnergyEfficiency,omitempty"`
	ExtraAttribute    *string `json:"ExtraAttribute,omitempty" xml:"ExtraAttribute,omitempty"`
	ForestFirstId     *string `json:"ForestFirstId,omitempty" xml:"ForestFirstId,omitempty"`
	ForestSecondId    *string `json:"ForestSecondId,omitempty" xml:"ForestSecondId,omitempty"`
	ItemBarCode       *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	ItemId            *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemInfoIndex     *int32  `json:"ItemInfoIndex,omitempty" xml:"ItemInfoIndex,omitempty"`
	ItemQrCode        *string `json:"ItemQrCode,omitempty" xml:"ItemQrCode,omitempty"`
	ItemShortTitle    *string `json:"ItemShortTitle,omitempty" xml:"ItemShortTitle,omitempty"`
	ItemTitle         *string `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	MemberPrice       *int32  `json:"MemberPrice,omitempty" xml:"MemberPrice,omitempty"`
	ModelNumber       *string `json:"ModelNumber,omitempty" xml:"ModelNumber,omitempty"`
	OptionGroups      *string `json:"OptionGroups,omitempty" xml:"OptionGroups,omitempty"`
	OriginalPrice     *int32  `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	PriceUnit         *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	ProductionPlace   *string `json:"ProductionPlace,omitempty" xml:"ProductionPlace,omitempty"`
	PromotionEnd      *string `json:"PromotionEnd,omitempty" xml:"PromotionEnd,omitempty"`
	PromotionReason   *string `json:"PromotionReason,omitempty" xml:"PromotionReason,omitempty"`
	PromotionStart    *string `json:"PromotionStart,omitempty" xml:"PromotionStart,omitempty"`
	PromotionText     *string `json:"PromotionText,omitempty" xml:"PromotionText,omitempty"`
	Rank              *string `json:"Rank,omitempty" xml:"Rank,omitempty"`
	SaleSpec          *string `json:"SaleSpec,omitempty" xml:"SaleSpec,omitempty"`
	SkuId             *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SourceCode        *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	StoreId           *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	SuggestPrice      *int32  `json:"SuggestPrice,omitempty" xml:"SuggestPrice,omitempty"`
}

func (s BatchInsertItemsRequestItemInfo) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertItemsRequestItemInfo) GoString() string {
	return s.String()
}

func (s *BatchInsertItemsRequestItemInfo) SetActionPrice(v int32) *BatchInsertItemsRequestItemInfo {
	s.ActionPrice = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetBePromotion(v bool) *BatchInsertItemsRequestItemInfo {
	s.BePromotion = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetBeSourceCode(v bool) *BatchInsertItemsRequestItemInfo {
	s.BeSourceCode = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetBrandName(v string) *BatchInsertItemsRequestItemInfo {
	s.BrandName = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCategoryName(v string) *BatchInsertItemsRequestItemInfo {
	s.CategoryName = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCompanyId(v string) *BatchInsertItemsRequestItemInfo {
	s.CompanyId = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureA(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureA = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureB(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureB = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureC(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureC = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureD(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureD = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureE(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureE = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureF(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureF = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureG(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureG = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureH(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureH = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureI(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureI = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetCustomizeFeatureJ(v string) *BatchInsertItemsRequestItemInfo {
	s.CustomizeFeatureJ = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetEnergyEfficiency(v string) *BatchInsertItemsRequestItemInfo {
	s.EnergyEfficiency = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetExtraAttribute(v string) *BatchInsertItemsRequestItemInfo {
	s.ExtraAttribute = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetForestFirstId(v string) *BatchInsertItemsRequestItemInfo {
	s.ForestFirstId = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetForestSecondId(v string) *BatchInsertItemsRequestItemInfo {
	s.ForestSecondId = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetItemBarCode(v string) *BatchInsertItemsRequestItemInfo {
	s.ItemBarCode = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetItemId(v int64) *BatchInsertItemsRequestItemInfo {
	s.ItemId = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetItemInfoIndex(v int32) *BatchInsertItemsRequestItemInfo {
	s.ItemInfoIndex = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetItemQrCode(v string) *BatchInsertItemsRequestItemInfo {
	s.ItemQrCode = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetItemShortTitle(v string) *BatchInsertItemsRequestItemInfo {
	s.ItemShortTitle = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetItemTitle(v string) *BatchInsertItemsRequestItemInfo {
	s.ItemTitle = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetMemberPrice(v int32) *BatchInsertItemsRequestItemInfo {
	s.MemberPrice = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetModelNumber(v string) *BatchInsertItemsRequestItemInfo {
	s.ModelNumber = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetOptionGroups(v string) *BatchInsertItemsRequestItemInfo {
	s.OptionGroups = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetOriginalPrice(v int32) *BatchInsertItemsRequestItemInfo {
	s.OriginalPrice = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetPriceUnit(v string) *BatchInsertItemsRequestItemInfo {
	s.PriceUnit = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetProductionPlace(v string) *BatchInsertItemsRequestItemInfo {
	s.ProductionPlace = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetPromotionEnd(v string) *BatchInsertItemsRequestItemInfo {
	s.PromotionEnd = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetPromotionReason(v string) *BatchInsertItemsRequestItemInfo {
	s.PromotionReason = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetPromotionStart(v string) *BatchInsertItemsRequestItemInfo {
	s.PromotionStart = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetPromotionText(v string) *BatchInsertItemsRequestItemInfo {
	s.PromotionText = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetRank(v string) *BatchInsertItemsRequestItemInfo {
	s.Rank = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetSaleSpec(v string) *BatchInsertItemsRequestItemInfo {
	s.SaleSpec = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetSkuId(v string) *BatchInsertItemsRequestItemInfo {
	s.SkuId = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetSourceCode(v string) *BatchInsertItemsRequestItemInfo {
	s.SourceCode = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetStoreId(v string) *BatchInsertItemsRequestItemInfo {
	s.StoreId = &v
	return s
}

func (s *BatchInsertItemsRequestItemInfo) SetSuggestPrice(v int32) *BatchInsertItemsRequestItemInfo {
	s.SuggestPrice = &v
	return s
}

type BatchInsertItemsResponseBody struct {
	BatchResults   *BatchInsertItemsResponseBodyBatchResults `json:"BatchResults,omitempty" xml:"BatchResults,omitempty" type:"Struct"`
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                                   `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                   `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchInsertItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertItemsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchInsertItemsResponseBody) SetBatchResults(v *BatchInsertItemsResponseBodyBatchResults) *BatchInsertItemsResponseBody {
	s.BatchResults = v
	return s
}

func (s *BatchInsertItemsResponseBody) SetCode(v string) *BatchInsertItemsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchInsertItemsResponseBody) SetDynamicCode(v string) *BatchInsertItemsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *BatchInsertItemsResponseBody) SetDynamicMessage(v string) *BatchInsertItemsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *BatchInsertItemsResponseBody) SetErrorCode(v string) *BatchInsertItemsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchInsertItemsResponseBody) SetErrorMessage(v string) *BatchInsertItemsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BatchInsertItemsResponseBody) SetMessage(v string) *BatchInsertItemsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchInsertItemsResponseBody) SetRequestId(v string) *BatchInsertItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchInsertItemsResponseBody) SetSuccess(v bool) *BatchInsertItemsResponseBody {
	s.Success = &v
	return s
}

type BatchInsertItemsResponseBodyBatchResults struct {
	BatchResult []*BatchInsertItemsResponseBodyBatchResultsBatchResult `json:"BatchResult,omitempty" xml:"BatchResult,omitempty" type:"Repeated"`
}

func (s BatchInsertItemsResponseBodyBatchResults) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertItemsResponseBodyBatchResults) GoString() string {
	return s.String()
}

func (s *BatchInsertItemsResponseBodyBatchResults) SetBatchResult(v []*BatchInsertItemsResponseBodyBatchResultsBatchResult) *BatchInsertItemsResponseBodyBatchResults {
	s.BatchResult = v
	return s
}

type BatchInsertItemsResponseBodyBatchResultsBatchResult struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Index     *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchInsertItemsResponseBodyBatchResultsBatchResult) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertItemsResponseBodyBatchResultsBatchResult) GoString() string {
	return s.String()
}

func (s *BatchInsertItemsResponseBodyBatchResultsBatchResult) SetErrorCode(v string) *BatchInsertItemsResponseBodyBatchResultsBatchResult {
	s.ErrorCode = &v
	return s
}

func (s *BatchInsertItemsResponseBodyBatchResultsBatchResult) SetIndex(v int32) *BatchInsertItemsResponseBodyBatchResultsBatchResult {
	s.Index = &v
	return s
}

func (s *BatchInsertItemsResponseBodyBatchResultsBatchResult) SetMessage(v string) *BatchInsertItemsResponseBodyBatchResultsBatchResult {
	s.Message = &v
	return s
}

func (s *BatchInsertItemsResponseBodyBatchResultsBatchResult) SetSuccess(v bool) *BatchInsertItemsResponseBodyBatchResultsBatchResult {
	s.Success = &v
	return s
}

type BatchInsertItemsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchInsertItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchInsertItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertItemsResponse) GoString() string {
	return s.String()
}

func (s *BatchInsertItemsResponse) SetHeaders(v map[string]*string) *BatchInsertItemsResponse {
	s.Headers = v
	return s
}

func (s *BatchInsertItemsResponse) SetStatusCode(v int32) *BatchInsertItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchInsertItemsResponse) SetBody(v *BatchInsertItemsResponseBody) *BatchInsertItemsResponse {
	s.Body = v
	return s
}

type BindEslDeviceRequest struct {
	// This parameter is required.
	EslBarCode *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	// This parameter is required.
	ItemBarCode *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s BindEslDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s BindEslDeviceRequest) GoString() string {
	return s.String()
}

func (s *BindEslDeviceRequest) SetEslBarCode(v string) *BindEslDeviceRequest {
	s.EslBarCode = &v
	return s
}

func (s *BindEslDeviceRequest) SetItemBarCode(v string) *BindEslDeviceRequest {
	s.ItemBarCode = &v
	return s
}

func (s *BindEslDeviceRequest) SetStoreId(v string) *BindEslDeviceRequest {
	s.StoreId = &v
	return s
}

type BindEslDeviceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindEslDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindEslDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BindEslDeviceResponseBody) SetCode(v string) *BindEslDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *BindEslDeviceResponseBody) SetDynamicCode(v string) *BindEslDeviceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *BindEslDeviceResponseBody) SetDynamicMessage(v string) *BindEslDeviceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *BindEslDeviceResponseBody) SetErrorCode(v string) *BindEslDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BindEslDeviceResponseBody) SetErrorMessage(v string) *BindEslDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BindEslDeviceResponseBody) SetMessage(v string) *BindEslDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *BindEslDeviceResponseBody) SetRequestId(v string) *BindEslDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindEslDeviceResponseBody) SetSuccess(v bool) *BindEslDeviceResponseBody {
	s.Success = &v
	return s
}

type BindEslDeviceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindEslDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindEslDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s BindEslDeviceResponse) GoString() string {
	return s.String()
}

func (s *BindEslDeviceResponse) SetHeaders(v map[string]*string) *BindEslDeviceResponse {
	s.Headers = v
	return s
}

func (s *BindEslDeviceResponse) SetStatusCode(v int32) *BindEslDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *BindEslDeviceResponse) SetBody(v *BindEslDeviceResponseBody) *BindEslDeviceResponse {
	s.Body = v
	return s
}

type BindEslDeviceShelfRequest struct {
	// This parameter is required.
	EslBarCode *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	// This parameter is required.
	ShelfCode *string `json:"ShelfCode,omitempty" xml:"ShelfCode,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s BindEslDeviceShelfRequest) String() string {
	return tea.Prettify(s)
}

func (s BindEslDeviceShelfRequest) GoString() string {
	return s.String()
}

func (s *BindEslDeviceShelfRequest) SetEslBarCode(v string) *BindEslDeviceShelfRequest {
	s.EslBarCode = &v
	return s
}

func (s *BindEslDeviceShelfRequest) SetShelfCode(v string) *BindEslDeviceShelfRequest {
	s.ShelfCode = &v
	return s
}

func (s *BindEslDeviceShelfRequest) SetStoreId(v string) *BindEslDeviceShelfRequest {
	s.StoreId = &v
	return s
}

type BindEslDeviceShelfResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindEslDeviceShelfResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindEslDeviceShelfResponseBody) GoString() string {
	return s.String()
}

func (s *BindEslDeviceShelfResponseBody) SetCode(v string) *BindEslDeviceShelfResponseBody {
	s.Code = &v
	return s
}

func (s *BindEslDeviceShelfResponseBody) SetDynamicCode(v string) *BindEslDeviceShelfResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *BindEslDeviceShelfResponseBody) SetDynamicMessage(v string) *BindEslDeviceShelfResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *BindEslDeviceShelfResponseBody) SetErrorCode(v string) *BindEslDeviceShelfResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BindEslDeviceShelfResponseBody) SetErrorMessage(v string) *BindEslDeviceShelfResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BindEslDeviceShelfResponseBody) SetMessage(v string) *BindEslDeviceShelfResponseBody {
	s.Message = &v
	return s
}

func (s *BindEslDeviceShelfResponseBody) SetRequestId(v string) *BindEslDeviceShelfResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindEslDeviceShelfResponseBody) SetSuccess(v bool) *BindEslDeviceShelfResponseBody {
	s.Success = &v
	return s
}

type BindEslDeviceShelfResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindEslDeviceShelfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindEslDeviceShelfResponse) String() string {
	return tea.Prettify(s)
}

func (s BindEslDeviceShelfResponse) GoString() string {
	return s.String()
}

func (s *BindEslDeviceShelfResponse) SetHeaders(v map[string]*string) *BindEslDeviceShelfResponse {
	s.Headers = v
	return s
}

func (s *BindEslDeviceShelfResponse) SetStatusCode(v int32) *BindEslDeviceShelfResponse {
	s.StatusCode = &v
	return s
}

func (s *BindEslDeviceShelfResponse) SetBody(v *BindEslDeviceShelfResponseBody) *BindEslDeviceShelfResponse {
	s.Body = v
	return s
}

type ConfirmLogisticsRequest struct {
	LogisticsDocuments *string `json:"LogisticsDocuments,omitempty" xml:"LogisticsDocuments,omitempty"`
	// This parameter is required.
	PoNumber *string `json:"PoNumber,omitempty" xml:"PoNumber,omitempty"`
	// This parameter is required.
	PrNumber *string `json:"PrNumber,omitempty" xml:"PrNumber,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ConfirmLogisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmLogisticsRequest) GoString() string {
	return s.String()
}

func (s *ConfirmLogisticsRequest) SetLogisticsDocuments(v string) *ConfirmLogisticsRequest {
	s.LogisticsDocuments = &v
	return s
}

func (s *ConfirmLogisticsRequest) SetPoNumber(v string) *ConfirmLogisticsRequest {
	s.PoNumber = &v
	return s
}

func (s *ConfirmLogisticsRequest) SetPrNumber(v string) *ConfirmLogisticsRequest {
	s.PrNumber = &v
	return s
}

func (s *ConfirmLogisticsRequest) SetStatus(v string) *ConfirmLogisticsRequest {
	s.Status = &v
	return s
}

type ConfirmLogisticsResponseBody struct {
	Acceptance     *string `json:"Acceptance,omitempty" xml:"Acceptance,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfirmLogisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmLogisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmLogisticsResponseBody) SetAcceptance(v string) *ConfirmLogisticsResponseBody {
	s.Acceptance = &v
	return s
}

func (s *ConfirmLogisticsResponseBody) SetCode(v string) *ConfirmLogisticsResponseBody {
	s.Code = &v
	return s
}

func (s *ConfirmLogisticsResponseBody) SetDynamicCode(v string) *ConfirmLogisticsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ConfirmLogisticsResponseBody) SetDynamicMessage(v string) *ConfirmLogisticsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ConfirmLogisticsResponseBody) SetErrorCode(v string) *ConfirmLogisticsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ConfirmLogisticsResponseBody) SetErrorMessage(v string) *ConfirmLogisticsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ConfirmLogisticsResponseBody) SetMessage(v string) *ConfirmLogisticsResponseBody {
	s.Message = &v
	return s
}

func (s *ConfirmLogisticsResponseBody) SetRequestId(v string) *ConfirmLogisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmLogisticsResponseBody) SetSuccess(v bool) *ConfirmLogisticsResponseBody {
	s.Success = &v
	return s
}

type ConfirmLogisticsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmLogisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmLogisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmLogisticsResponse) GoString() string {
	return s.String()
}

func (s *ConfirmLogisticsResponse) SetHeaders(v map[string]*string) *ConfirmLogisticsResponse {
	s.Headers = v
	return s
}

func (s *ConfirmLogisticsResponse) SetStatusCode(v int32) *ConfirmLogisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmLogisticsResponse) SetBody(v *ConfirmLogisticsResponseBody) *ConfirmLogisticsResponse {
	s.Body = v
	return s
}

type CreateStoreRequest struct {
	Brand    *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// This parameter is required.
	CompanyId *string `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	Groups    *string `json:"Groups,omitempty" xml:"Groups,omitempty"`
	OutId     *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	ParentId  *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// This parameter is required.
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// This parameter is required.
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
}

func (s CreateStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStoreRequest) GoString() string {
	return s.String()
}

func (s *CreateStoreRequest) SetBrand(v string) *CreateStoreRequest {
	s.Brand = &v
	return s
}

func (s *CreateStoreRequest) SetComments(v string) *CreateStoreRequest {
	s.Comments = &v
	return s
}

func (s *CreateStoreRequest) SetCompanyId(v string) *CreateStoreRequest {
	s.CompanyId = &v
	return s
}

func (s *CreateStoreRequest) SetGroups(v string) *CreateStoreRequest {
	s.Groups = &v
	return s
}

func (s *CreateStoreRequest) SetOutId(v string) *CreateStoreRequest {
	s.OutId = &v
	return s
}

func (s *CreateStoreRequest) SetParentId(v string) *CreateStoreRequest {
	s.ParentId = &v
	return s
}

func (s *CreateStoreRequest) SetPhone(v string) *CreateStoreRequest {
	s.Phone = &v
	return s
}

func (s *CreateStoreRequest) SetStoreName(v string) *CreateStoreRequest {
	s.StoreName = &v
	return s
}

type CreateStoreResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StoreId        *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStoreResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStoreResponseBody) SetCode(v string) *CreateStoreResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStoreResponseBody) SetDynamicCode(v string) *CreateStoreResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateStoreResponseBody) SetDynamicMessage(v string) *CreateStoreResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateStoreResponseBody) SetErrorCode(v string) *CreateStoreResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateStoreResponseBody) SetErrorMessage(v string) *CreateStoreResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateStoreResponseBody) SetMessage(v string) *CreateStoreResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStoreResponseBody) SetRequestId(v string) *CreateStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStoreResponseBody) SetStoreId(v string) *CreateStoreResponseBody {
	s.StoreId = &v
	return s
}

func (s *CreateStoreResponseBody) SetSuccess(v bool) *CreateStoreResponseBody {
	s.Success = &v
	return s
}

type CreateStoreResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStoreResponse) GoString() string {
	return s.String()
}

func (s *CreateStoreResponse) SetHeaders(v map[string]*string) *CreateStoreResponse {
	s.Headers = v
	return s
}

func (s *CreateStoreResponse) SetStatusCode(v int32) *CreateStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStoreResponse) SetBody(v *CreateStoreResponseBody) *CreateStoreResponse {
	s.Body = v
	return s
}

type DeleteApDeviceRequest struct {
	// This parameter is required.
	ApMac *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DeleteApDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteApDeviceRequest) SetApMac(v string) *DeleteApDeviceRequest {
	s.ApMac = &v
	return s
}

func (s *DeleteApDeviceRequest) SetStoreId(v string) *DeleteApDeviceRequest {
	s.StoreId = &v
	return s
}

type DeleteApDeviceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteApDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApDeviceResponseBody) SetCode(v string) *DeleteApDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApDeviceResponseBody) SetDynamicCode(v string) *DeleteApDeviceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteApDeviceResponseBody) SetDynamicMessage(v string) *DeleteApDeviceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteApDeviceResponseBody) SetErrorCode(v string) *DeleteApDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteApDeviceResponseBody) SetErrorMessage(v string) *DeleteApDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteApDeviceResponseBody) SetMessage(v string) *DeleteApDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApDeviceResponseBody) SetRequestId(v string) *DeleteApDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApDeviceResponseBody) SetSuccess(v bool) *DeleteApDeviceResponseBody {
	s.Success = &v
	return s
}

type DeleteApDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteApDeviceResponse) SetHeaders(v map[string]*string) *DeleteApDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteApDeviceResponse) SetStatusCode(v int32) *DeleteApDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApDeviceResponse) SetBody(v *DeleteApDeviceResponseBody) *DeleteApDeviceResponse {
	s.Body = v
	return s
}

type DeleteEslDeviceRequest struct {
	// This parameter is required.
	EslBarCode *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DeleteEslDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEslDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteEslDeviceRequest) SetEslBarCode(v string) *DeleteEslDeviceRequest {
	s.EslBarCode = &v
	return s
}

func (s *DeleteEslDeviceRequest) SetStoreId(v string) *DeleteEslDeviceRequest {
	s.StoreId = &v
	return s
}

type DeleteEslDeviceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEslDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEslDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEslDeviceResponseBody) SetCode(v string) *DeleteEslDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEslDeviceResponseBody) SetDynamicCode(v string) *DeleteEslDeviceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteEslDeviceResponseBody) SetDynamicMessage(v string) *DeleteEslDeviceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteEslDeviceResponseBody) SetErrorCode(v string) *DeleteEslDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteEslDeviceResponseBody) SetErrorMessage(v string) *DeleteEslDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteEslDeviceResponseBody) SetMessage(v string) *DeleteEslDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEslDeviceResponseBody) SetRequestId(v string) *DeleteEslDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEslDeviceResponseBody) SetSuccess(v bool) *DeleteEslDeviceResponseBody {
	s.Success = &v
	return s
}

type DeleteEslDeviceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEslDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEslDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEslDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteEslDeviceResponse) SetHeaders(v map[string]*string) *DeleteEslDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteEslDeviceResponse) SetStatusCode(v int32) *DeleteEslDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEslDeviceResponse) SetBody(v *DeleteEslDeviceResponseBody) *DeleteEslDeviceResponse {
	s.Body = v
	return s
}

type DeleteItemRequest struct {
	// This parameter is required.
	ItemBarCode *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DeleteItemRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteItemRequest) GoString() string {
	return s.String()
}

func (s *DeleteItemRequest) SetItemBarCode(v string) *DeleteItemRequest {
	s.ItemBarCode = &v
	return s
}

func (s *DeleteItemRequest) SetStoreId(v string) *DeleteItemRequest {
	s.StoreId = &v
	return s
}

type DeleteItemResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteItemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteItemResponseBody) SetCode(v string) *DeleteItemResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteItemResponseBody) SetDynamicCode(v string) *DeleteItemResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteItemResponseBody) SetDynamicMessage(v string) *DeleteItemResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteItemResponseBody) SetErrorCode(v string) *DeleteItemResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteItemResponseBody) SetErrorMessage(v string) *DeleteItemResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteItemResponseBody) SetMessage(v string) *DeleteItemResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteItemResponseBody) SetRequestId(v string) *DeleteItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteItemResponseBody) SetSuccess(v bool) *DeleteItemResponseBody {
	s.Success = &v
	return s
}

type DeleteItemResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteItemResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteItemResponse) SetHeaders(v map[string]*string) *DeleteItemResponse {
	s.Headers = v
	return s
}

func (s *DeleteItemResponse) SetStatusCode(v int32) *DeleteItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteItemResponse) SetBody(v *DeleteItemResponseBody) *DeleteItemResponse {
	s.Body = v
	return s
}

type DeleteItemBySkuIdRequest struct {
	// This parameter is required.
	SkuId *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DeleteItemBySkuIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteItemBySkuIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteItemBySkuIdRequest) SetSkuId(v string) *DeleteItemBySkuIdRequest {
	s.SkuId = &v
	return s
}

func (s *DeleteItemBySkuIdRequest) SetStoreId(v string) *DeleteItemBySkuIdRequest {
	s.StoreId = &v
	return s
}

type DeleteItemBySkuIdResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteItemBySkuIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteItemBySkuIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteItemBySkuIdResponseBody) SetCode(v string) *DeleteItemBySkuIdResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteItemBySkuIdResponseBody) SetDynamicCode(v string) *DeleteItemBySkuIdResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteItemBySkuIdResponseBody) SetDynamicMessage(v string) *DeleteItemBySkuIdResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteItemBySkuIdResponseBody) SetErrorCode(v string) *DeleteItemBySkuIdResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteItemBySkuIdResponseBody) SetErrorMessage(v string) *DeleteItemBySkuIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteItemBySkuIdResponseBody) SetMessage(v string) *DeleteItemBySkuIdResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteItemBySkuIdResponseBody) SetRequestId(v string) *DeleteItemBySkuIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteItemBySkuIdResponseBody) SetSuccess(v bool) *DeleteItemBySkuIdResponseBody {
	s.Success = &v
	return s
}

type DeleteItemBySkuIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteItemBySkuIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteItemBySkuIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteItemBySkuIdResponse) GoString() string {
	return s.String()
}

func (s *DeleteItemBySkuIdResponse) SetHeaders(v map[string]*string) *DeleteItemBySkuIdResponse {
	s.Headers = v
	return s
}

func (s *DeleteItemBySkuIdResponse) SetStatusCode(v int32) *DeleteItemBySkuIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteItemBySkuIdResponse) SetBody(v *DeleteItemBySkuIdResponseBody) *DeleteItemBySkuIdResponse {
	s.Body = v
	return s
}

type DeleteStoreRequest struct {
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DeleteStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStoreRequest) GoString() string {
	return s.String()
}

func (s *DeleteStoreRequest) SetStoreId(v string) *DeleteStoreRequest {
	s.StoreId = &v
	return s
}

type DeleteStoreResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStoreResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStoreResponseBody) SetCode(v string) *DeleteStoreResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteStoreResponseBody) SetDynamicCode(v string) *DeleteStoreResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteStoreResponseBody) SetDynamicMessage(v string) *DeleteStoreResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteStoreResponseBody) SetErrorCode(v string) *DeleteStoreResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteStoreResponseBody) SetErrorMessage(v string) *DeleteStoreResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteStoreResponseBody) SetMessage(v string) *DeleteStoreResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteStoreResponseBody) SetRequestId(v string) *DeleteStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStoreResponseBody) SetSuccess(v bool) *DeleteStoreResponseBody {
	s.Success = &v
	return s
}

type DeleteStoreResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStoreResponse) GoString() string {
	return s.String()
}

func (s *DeleteStoreResponse) SetHeaders(v map[string]*string) *DeleteStoreResponse {
	s.Headers = v
	return s
}

func (s *DeleteStoreResponse) SetStatusCode(v int32) *DeleteStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStoreResponse) SetBody(v *DeleteStoreResponseBody) *DeleteStoreResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	// This parameter is required.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetUserId(v string) *DeleteUserRequest {
	s.UserId = &v
	return s
}

type DeleteUserResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) SetCode(v string) *DeleteUserResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUserResponseBody) SetDynamicCode(v string) *DeleteUserResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteUserResponseBody) SetDynamicMessage(v string) *DeleteUserResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteUserResponseBody) SetErrorCode(v string) *DeleteUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteUserResponseBody) SetErrorMessage(v string) *DeleteUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteUserResponseBody) SetMessage(v string) *DeleteUserResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserResponseBody) SetSuccess(v bool) *DeleteUserResponseBody {
	s.Success = &v
	return s
}

type DeleteUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetStatusCode(v int32) *DeleteUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type DescribeAlarmsRequest struct {
	AlarmId       *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	AlarmStatus   *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	AlarmType     *string `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	ErrorType     *string `json:"ErrorType,omitempty" xml:"ErrorType,omitempty"`
	FromAlarmTime *string `json:"FromAlarmTime,omitempty" xml:"FromAlarmTime,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	StoreId     *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	ToAlarmTime *string `json:"ToAlarmTime,omitempty" xml:"ToAlarmTime,omitempty"`
}

func (s DescribeAlarmsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsRequest) SetAlarmId(v string) *DescribeAlarmsRequest {
	s.AlarmId = &v
	return s
}

func (s *DescribeAlarmsRequest) SetAlarmStatus(v string) *DescribeAlarmsRequest {
	s.AlarmStatus = &v
	return s
}

func (s *DescribeAlarmsRequest) SetAlarmType(v string) *DescribeAlarmsRequest {
	s.AlarmType = &v
	return s
}

func (s *DescribeAlarmsRequest) SetErrorType(v string) *DescribeAlarmsRequest {
	s.ErrorType = &v
	return s
}

func (s *DescribeAlarmsRequest) SetFromAlarmTime(v string) *DescribeAlarmsRequest {
	s.FromAlarmTime = &v
	return s
}

func (s *DescribeAlarmsRequest) SetPageNumber(v int32) *DescribeAlarmsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAlarmsRequest) SetPageSize(v int32) *DescribeAlarmsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlarmsRequest) SetStoreId(v string) *DescribeAlarmsRequest {
	s.StoreId = &v
	return s
}

func (s *DescribeAlarmsRequest) SetToAlarmTime(v string) *DescribeAlarmsRequest {
	s.ToAlarmTime = &v
	return s
}

type DescribeAlarmsResponseBody struct {
	Alarms         *DescribeAlarmsResponseBodyAlarms `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Struct"`
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                           `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                           `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAlarmsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBody) SetAlarms(v *DescribeAlarmsResponseBodyAlarms) *DescribeAlarmsResponseBody {
	s.Alarms = v
	return s
}

func (s *DescribeAlarmsResponseBody) SetCode(v string) *DescribeAlarmsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetDynamicCode(v string) *DescribeAlarmsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetDynamicMessage(v string) *DescribeAlarmsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetErrorCode(v string) *DescribeAlarmsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetErrorMessage(v string) *DescribeAlarmsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetMessage(v string) *DescribeAlarmsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetPageNumber(v int32) *DescribeAlarmsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetPageSize(v int32) *DescribeAlarmsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetRequestId(v string) *DescribeAlarmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetSuccess(v bool) *DescribeAlarmsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetTotalCount(v int32) *DescribeAlarmsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAlarmsResponseBodyAlarms struct {
	AlarmInfo []*DescribeAlarmsResponseBodyAlarmsAlarmInfo `json:"AlarmInfo,omitempty" xml:"AlarmInfo,omitempty" type:"Repeated"`
}

func (s DescribeAlarmsResponseBodyAlarms) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBodyAlarms) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyAlarms) SetAlarmInfo(v []*DescribeAlarmsResponseBodyAlarmsAlarmInfo) *DescribeAlarmsResponseBodyAlarms {
	s.AlarmInfo = v
	return s
}

type DescribeAlarmsResponseBodyAlarmsAlarmInfo struct {
	AlarmId       *string `json:"AlarmId,omitempty" xml:"AlarmId,omitempty"`
	AlarmStatus   *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	AlarmTime     *string `json:"AlarmTime,omitempty" xml:"AlarmTime,omitempty"`
	AlarmType     *string `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	CompanyId     *string `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	DealTime      *string `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	DealUserId    *int64  `json:"DealUserId,omitempty" xml:"DealUserId,omitempty"`
	DeviceBarCode *string `json:"DeviceBarCode,omitempty" xml:"DeviceBarCode,omitempty"`
	DeviceMac     *string `json:"DeviceMac,omitempty" xml:"DeviceMac,omitempty"`
	DeviceType    *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	ErrorType     *string `json:"ErrorType,omitempty" xml:"ErrorType,omitempty"`
	ItemBarCode   *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	ItemTitle     *string `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	Model         *string `json:"Model,omitempty" xml:"Model,omitempty"`
	StoreId       *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	Vendor        *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeAlarmsResponseBodyAlarmsAlarmInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponseBodyAlarmsAlarmInfo) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyAlarmsAlarmInfo) SetAlarmId(v string) *DescribeAlarmsResponseBodyAlarmsAlarmInfo {
	s.AlarmId = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmsAlarmInfo) SetAlarmStatus(v string) *DescribeAlarmsResponseBodyAlarmsAlarmInfo {
	s.AlarmStatus = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmsAlarmInfo) SetAlarmTime(v string) *DescribeAlarmsResponseBodyAlarmsAlarmInfo {
	s.AlarmTime = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmsAlarmInfo) SetAlarmType(v string) *DescribeAlarmsResponseBodyAlarmsAlarmInfo {
	s.AlarmType = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmsAlarmInfo) SetCompanyId(v string) *DescribeAlarmsResponseBodyAlarmsAlarmInfo {
	s.CompanyId = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmsAlarmInfo) SetDealTime(v string) *DescribeAlarmsResponseBodyAlarmsAlarmInfo {
	s.DealTime = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmsAlarmInfo) SetDealUserId(v int64) *DescribeAlarmsResponseBodyAlarmsAlarmInfo {
	s.DealUserId = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmsAlarmInfo) SetDeviceBarCode(v string) *DescribeAlarmsResponseBodyAlarmsAlarmInfo {
	s.DeviceBarCode = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmsAlarmInfo) SetDeviceMac(v string) *DescribeAlarmsResponseBodyAlarmsAlarmInfo {
	s.DeviceMac = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmsAlarmInfo) SetDeviceType(v string) *DescribeAlarmsResponseBodyAlarmsAlarmInfo {
	s.DeviceType = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmsAlarmInfo) SetErrorType(v string) *DescribeAlarmsResponseBodyAlarmsAlarmInfo {
	s.ErrorType = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmsAlarmInfo) SetItemBarCode(v string) *DescribeAlarmsResponseBodyAlarmsAlarmInfo {
	s.ItemBarCode = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmsAlarmInfo) SetItemTitle(v string) *DescribeAlarmsResponseBodyAlarmsAlarmInfo {
	s.ItemTitle = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmsAlarmInfo) SetModel(v string) *DescribeAlarmsResponseBodyAlarmsAlarmInfo {
	s.Model = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmsAlarmInfo) SetStoreId(v string) *DescribeAlarmsResponseBodyAlarmsAlarmInfo {
	s.StoreId = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmsAlarmInfo) SetVendor(v string) *DescribeAlarmsResponseBodyAlarmsAlarmInfo {
	s.Vendor = &v
	return s
}

type DescribeAlarmsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlarmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlarmsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponse) SetHeaders(v map[string]*string) *DescribeAlarmsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlarmsResponse) SetStatusCode(v int32) *DescribeAlarmsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlarmsResponse) SetBody(v *DescribeAlarmsResponseBody) *DescribeAlarmsResponse {
	s.Body = v
	return s
}

type DescribeApDevicesRequest struct {
	Activated  *bool   `json:"Activated,omitempty" xml:"Activated,omitempty"`
	ApMac      *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DescribeApDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApDevicesRequest) SetActivated(v bool) *DescribeApDevicesRequest {
	s.Activated = &v
	return s
}

func (s *DescribeApDevicesRequest) SetApMac(v string) *DescribeApDevicesRequest {
	s.ApMac = &v
	return s
}

func (s *DescribeApDevicesRequest) SetPageNumber(v int32) *DescribeApDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApDevicesRequest) SetPageSize(v int32) *DescribeApDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApDevicesRequest) SetStoreId(v string) *DescribeApDevicesRequest {
	s.StoreId = &v
	return s
}

type DescribeApDevicesResponseBody struct {
	ApDevices      *DescribeApDevicesResponseBodyApDevices `json:"ApDevices,omitempty" xml:"ApDevices,omitempty" type:"Struct"`
	Code           *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                                 `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                 `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApDevicesResponseBody) SetApDevices(v *DescribeApDevicesResponseBodyApDevices) *DescribeApDevicesResponseBody {
	s.ApDevices = v
	return s
}

func (s *DescribeApDevicesResponseBody) SetCode(v string) *DescribeApDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetDynamicCode(v string) *DescribeApDevicesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetDynamicMessage(v string) *DescribeApDevicesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetErrorCode(v string) *DescribeApDevicesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetErrorMessage(v string) *DescribeApDevicesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetMessage(v string) *DescribeApDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetPageNumber(v int32) *DescribeApDevicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetPageSize(v int32) *DescribeApDevicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetRequestId(v string) *DescribeApDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetSuccess(v bool) *DescribeApDevicesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApDevicesResponseBody) SetTotalCount(v int32) *DescribeApDevicesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApDevicesResponseBodyApDevices struct {
	ApInfo []*DescribeApDevicesResponseBodyApDevicesApInfo `json:"ApInfo,omitempty" xml:"ApInfo,omitempty" type:"Repeated"`
}

func (s DescribeApDevicesResponseBodyApDevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeApDevicesResponseBodyApDevices) GoString() string {
	return s.String()
}

func (s *DescribeApDevicesResponseBodyApDevices) SetApInfo(v []*DescribeApDevicesResponseBodyApDevicesApInfo) *DescribeApDevicesResponseBodyApDevices {
	s.ApInfo = v
	return s
}

type DescribeApDevicesResponseBodyApDevicesApInfo struct {
	IsActivate *bool   `json:"IsActivate,omitempty" xml:"IsActivate,omitempty"`
	Mac        *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Model      *string `json:"Model,omitempty" xml:"Model,omitempty"`
	Status     *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeApDevicesResponseBodyApDevicesApInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeApDevicesResponseBodyApDevicesApInfo) GoString() string {
	return s.String()
}

func (s *DescribeApDevicesResponseBodyApDevicesApInfo) SetIsActivate(v bool) *DescribeApDevicesResponseBodyApDevicesApInfo {
	s.IsActivate = &v
	return s
}

func (s *DescribeApDevicesResponseBodyApDevicesApInfo) SetMac(v string) *DescribeApDevicesResponseBodyApDevicesApInfo {
	s.Mac = &v
	return s
}

func (s *DescribeApDevicesResponseBodyApDevicesApInfo) SetModel(v string) *DescribeApDevicesResponseBodyApDevicesApInfo {
	s.Model = &v
	return s
}

func (s *DescribeApDevicesResponseBodyApDevicesApInfo) SetStatus(v bool) *DescribeApDevicesResponseBodyApDevicesApInfo {
	s.Status = &v
	return s
}

type DescribeApDevicesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApDevicesResponse) SetHeaders(v map[string]*string) *DescribeApDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApDevicesResponse) SetStatusCode(v int32) *DescribeApDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApDevicesResponse) SetBody(v *DescribeApDevicesResponseBody) *DescribeApDevicesResponse {
	s.Body = v
	return s
}

type DescribeEslDevicesRequest struct {
	BeBind           *bool   `json:"BeBind,omitempty" xml:"BeBind,omitempty"`
	EslBarCode       *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	EslStatus        *string `json:"EslStatus,omitempty" xml:"EslStatus,omitempty"`
	FromBatteryLevel *int32  `json:"FromBatteryLevel,omitempty" xml:"FromBatteryLevel,omitempty"`
	ItemBarCode      *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	Mac              *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ShelfCode        *string `json:"ShelfCode,omitempty" xml:"ShelfCode,omitempty"`
	// This parameter is required.
	StoreId        *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	ToBatteryLevel *int32  `json:"ToBatteryLevel,omitempty" xml:"ToBatteryLevel,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Vendor         *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeEslDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEslDevicesRequest) SetBeBind(v bool) *DescribeEslDevicesRequest {
	s.BeBind = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetEslBarCode(v string) *DescribeEslDevicesRequest {
	s.EslBarCode = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetEslStatus(v string) *DescribeEslDevicesRequest {
	s.EslStatus = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetFromBatteryLevel(v int32) *DescribeEslDevicesRequest {
	s.FromBatteryLevel = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetItemBarCode(v string) *DescribeEslDevicesRequest {
	s.ItemBarCode = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetMac(v string) *DescribeEslDevicesRequest {
	s.Mac = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetPageNumber(v int32) *DescribeEslDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetPageSize(v int32) *DescribeEslDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetShelfCode(v string) *DescribeEslDevicesRequest {
	s.ShelfCode = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetStoreId(v string) *DescribeEslDevicesRequest {
	s.StoreId = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetToBatteryLevel(v int32) *DescribeEslDevicesRequest {
	s.ToBatteryLevel = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetType(v string) *DescribeEslDevicesRequest {
	s.Type = &v
	return s
}

func (s *DescribeEslDevicesRequest) SetVendor(v string) *DescribeEslDevicesRequest {
	s.Vendor = &v
	return s
}

type DescribeEslDevicesResponseBody struct {
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                                   `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                   `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	EslDevices     *DescribeEslDevicesResponseBodyEslDevices `json:"EslDevices,omitempty" xml:"EslDevices,omitempty" type:"Struct"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEslDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEslDevicesResponseBody) SetCode(v string) *DescribeEslDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetDynamicCode(v string) *DescribeEslDevicesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetDynamicMessage(v string) *DescribeEslDevicesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetErrorCode(v string) *DescribeEslDevicesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetErrorMessage(v string) *DescribeEslDevicesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetEslDevices(v *DescribeEslDevicesResponseBodyEslDevices) *DescribeEslDevicesResponseBody {
	s.EslDevices = v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetMessage(v string) *DescribeEslDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetPageNumber(v int32) *DescribeEslDevicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetPageSize(v int32) *DescribeEslDevicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetRequestId(v string) *DescribeEslDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetSuccess(v bool) *DescribeEslDevicesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEslDevicesResponseBody) SetTotalCount(v int32) *DescribeEslDevicesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeEslDevicesResponseBodyEslDevices struct {
	EslDeviceInfo []*DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo `json:"EslDeviceInfo,omitempty" xml:"EslDeviceInfo,omitempty" type:"Repeated"`
}

func (s DescribeEslDevicesResponseBodyEslDevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslDevicesResponseBodyEslDevices) GoString() string {
	return s.String()
}

func (s *DescribeEslDevicesResponseBodyEslDevices) SetEslDeviceInfo(v []*DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) *DescribeEslDevicesResponseBodyEslDevices {
	s.EslDeviceInfo = v
	return s
}

type DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo struct {
	BatteryLevel        *int32  `json:"BatteryLevel,omitempty" xml:"BatteryLevel,omitempty"`
	BeBind              *bool   `json:"BeBind,omitempty" xml:"BeBind,omitempty"`
	CompanyId           *string `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	ConnectAp           *string `json:"ConnectAp,omitempty" xml:"ConnectAp,omitempty"`
	EslBarCode          *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	EslStatus           *string `json:"EslStatus,omitempty" xml:"EslStatus,omitempty"`
	ItemActionPrice     *int32  `json:"ItemActionPrice,omitempty" xml:"ItemActionPrice,omitempty"`
	ItemBarCode         *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	ItemId              *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemPriceUnit       *string `json:"ItemPriceUnit,omitempty" xml:"ItemPriceUnit,omitempty"`
	ItemTitle           *string `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	LastCommunicateTime *string `json:"LastCommunicateTime,omitempty" xml:"LastCommunicateTime,omitempty"`
	Mac                 *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Model               *string `json:"Model,omitempty" xml:"Model,omitempty"`
	PositionCode        *string `json:"PositionCode,omitempty" xml:"PositionCode,omitempty"`
	ScreenHeight        *int32  `json:"ScreenHeight,omitempty" xml:"ScreenHeight,omitempty"`
	ScreenWidth         *int32  `json:"ScreenWidth,omitempty" xml:"ScreenWidth,omitempty"`
	ShelfCode           *string `json:"ShelfCode,omitempty" xml:"ShelfCode,omitempty"`
	StoreId             *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Vendor              *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) GoString() string {
	return s.String()
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetBatteryLevel(v int32) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.BatteryLevel = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetBeBind(v bool) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.BeBind = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetCompanyId(v string) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.CompanyId = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetConnectAp(v string) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.ConnectAp = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetEslBarCode(v string) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.EslBarCode = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetEslStatus(v string) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.EslStatus = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetItemActionPrice(v int32) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.ItemActionPrice = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetItemBarCode(v string) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.ItemBarCode = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetItemId(v int64) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.ItemId = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetItemPriceUnit(v string) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.ItemPriceUnit = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetItemTitle(v string) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.ItemTitle = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetLastCommunicateTime(v string) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.LastCommunicateTime = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetMac(v string) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.Mac = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetModel(v string) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.Model = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetPositionCode(v string) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.PositionCode = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetScreenHeight(v int32) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.ScreenHeight = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetScreenWidth(v int32) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.ScreenWidth = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetShelfCode(v string) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.ShelfCode = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetStoreId(v string) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.StoreId = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetType(v string) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.Type = &v
	return s
}

func (s *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo) SetVendor(v string) *DescribeEslDevicesResponseBodyEslDevicesEslDeviceInfo {
	s.Vendor = &v
	return s
}

type DescribeEslDevicesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEslDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEslDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEslDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEslDevicesResponse) SetHeaders(v map[string]*string) *DescribeEslDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEslDevicesResponse) SetStatusCode(v int32) *DescribeEslDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEslDevicesResponse) SetBody(v *DescribeEslDevicesResponseBody) *DescribeEslDevicesResponse {
	s.Body = v
	return s
}

type DescribeItemsRequest struct {
	BePromotion *bool   `json:"BePromotion,omitempty" xml:"BePromotion,omitempty"`
	ItemBarCode *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	ItemId      *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemTitle   *string `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SkuId       *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DescribeItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeItemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeItemsRequest) SetBePromotion(v bool) *DescribeItemsRequest {
	s.BePromotion = &v
	return s
}

func (s *DescribeItemsRequest) SetItemBarCode(v string) *DescribeItemsRequest {
	s.ItemBarCode = &v
	return s
}

func (s *DescribeItemsRequest) SetItemId(v int64) *DescribeItemsRequest {
	s.ItemId = &v
	return s
}

func (s *DescribeItemsRequest) SetItemTitle(v string) *DescribeItemsRequest {
	s.ItemTitle = &v
	return s
}

func (s *DescribeItemsRequest) SetPageNumber(v int32) *DescribeItemsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeItemsRequest) SetPageSize(v int32) *DescribeItemsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeItemsRequest) SetSkuId(v string) *DescribeItemsRequest {
	s.SkuId = &v
	return s
}

func (s *DescribeItemsRequest) SetStoreId(v string) *DescribeItemsRequest {
	s.StoreId = &v
	return s
}

type DescribeItemsResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                         `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                         `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Items          *DescribeItemsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeItemsResponseBody) SetCode(v string) *DescribeItemsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeItemsResponseBody) SetDynamicCode(v string) *DescribeItemsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeItemsResponseBody) SetDynamicMessage(v string) *DescribeItemsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeItemsResponseBody) SetErrorCode(v string) *DescribeItemsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeItemsResponseBody) SetErrorMessage(v string) *DescribeItemsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeItemsResponseBody) SetItems(v *DescribeItemsResponseBodyItems) *DescribeItemsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeItemsResponseBody) SetMessage(v string) *DescribeItemsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeItemsResponseBody) SetPageNumber(v int32) *DescribeItemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeItemsResponseBody) SetPageSize(v int32) *DescribeItemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeItemsResponseBody) SetRequestId(v string) *DescribeItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeItemsResponseBody) SetSuccess(v bool) *DescribeItemsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeItemsResponseBody) SetTotalCount(v int32) *DescribeItemsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeItemsResponseBodyItems struct {
	ItemInfo []*DescribeItemsResponseBodyItemsItemInfo `json:"ItemInfo,omitempty" xml:"ItemInfo,omitempty" type:"Repeated"`
}

func (s DescribeItemsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeItemsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeItemsResponseBodyItems) SetItemInfo(v []*DescribeItemsResponseBodyItemsItemInfo) *DescribeItemsResponseBodyItems {
	s.ItemInfo = v
	return s
}

type DescribeItemsResponseBodyItemsItemInfo struct {
	ActionPrice       *int32  `json:"ActionPrice,omitempty" xml:"ActionPrice,omitempty"`
	BePromotion       *bool   `json:"BePromotion,omitempty" xml:"BePromotion,omitempty"`
	BeSourceCode      *bool   `json:"BeSourceCode,omitempty" xml:"BeSourceCode,omitempty"`
	BrandName         *string `json:"BrandName,omitempty" xml:"BrandName,omitempty"`
	CategoryName      *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	CompanyId         *string `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	CustomizeFeatureA *string `json:"CustomizeFeatureA,omitempty" xml:"CustomizeFeatureA,omitempty"`
	CustomizeFeatureB *string `json:"CustomizeFeatureB,omitempty" xml:"CustomizeFeatureB,omitempty"`
	CustomizeFeatureC *string `json:"CustomizeFeatureC,omitempty" xml:"CustomizeFeatureC,omitempty"`
	CustomizeFeatureD *string `json:"CustomizeFeatureD,omitempty" xml:"CustomizeFeatureD,omitempty"`
	CustomizeFeatureE *string `json:"CustomizeFeatureE,omitempty" xml:"CustomizeFeatureE,omitempty"`
	CustomizeFeatureF *string `json:"CustomizeFeatureF,omitempty" xml:"CustomizeFeatureF,omitempty"`
	CustomizeFeatureG *string `json:"CustomizeFeatureG,omitempty" xml:"CustomizeFeatureG,omitempty"`
	CustomizeFeatureH *string `json:"CustomizeFeatureH,omitempty" xml:"CustomizeFeatureH,omitempty"`
	CustomizeFeatureI *string `json:"CustomizeFeatureI,omitempty" xml:"CustomizeFeatureI,omitempty"`
	CustomizeFeatureJ *string `json:"CustomizeFeatureJ,omitempty" xml:"CustomizeFeatureJ,omitempty"`
	EnergyEfficiency  *string `json:"EnergyEfficiency,omitempty" xml:"EnergyEfficiency,omitempty"`
	ExtraAttribute    *string `json:"ExtraAttribute,omitempty" xml:"ExtraAttribute,omitempty"`
	ForestFirstId     *string `json:"ForestFirstId,omitempty" xml:"ForestFirstId,omitempty"`
	ForestSecondId    *string `json:"ForestSecondId,omitempty" xml:"ForestSecondId,omitempty"`
	ItemBarCode       *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	ItemId            *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemInfoIndex     *int32  `json:"ItemInfoIndex,omitempty" xml:"ItemInfoIndex,omitempty"`
	ItemQrCode        *string `json:"ItemQrCode,omitempty" xml:"ItemQrCode,omitempty"`
	ItemShortTitle    *string `json:"ItemShortTitle,omitempty" xml:"ItemShortTitle,omitempty"`
	ItemTitle         *string `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	MemberPrice       *int32  `json:"MemberPrice,omitempty" xml:"MemberPrice,omitempty"`
	ModelNumber       *string `json:"ModelNumber,omitempty" xml:"ModelNumber,omitempty"`
	OptionGroups      *string `json:"OptionGroups,omitempty" xml:"OptionGroups,omitempty"`
	OriginalPrice     *int32  `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	PriceUnit         *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	ProductionPlace   *string `json:"ProductionPlace,omitempty" xml:"ProductionPlace,omitempty"`
	PromotionEnd      *string `json:"PromotionEnd,omitempty" xml:"PromotionEnd,omitempty"`
	PromotionReason   *string `json:"PromotionReason,omitempty" xml:"PromotionReason,omitempty"`
	PromotionStart    *string `json:"PromotionStart,omitempty" xml:"PromotionStart,omitempty"`
	PromotionText     *string `json:"PromotionText,omitempty" xml:"PromotionText,omitempty"`
	Rank              *string `json:"Rank,omitempty" xml:"Rank,omitempty"`
	SaleSpec          *string `json:"SaleSpec,omitempty" xml:"SaleSpec,omitempty"`
	SkuId             *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SourceCode        *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	StoreId           *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	SuggestPrice      *int32  `json:"SuggestPrice,omitempty" xml:"SuggestPrice,omitempty"`
}

func (s DescribeItemsResponseBodyItemsItemInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeItemsResponseBodyItemsItemInfo) GoString() string {
	return s.String()
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetActionPrice(v int32) *DescribeItemsResponseBodyItemsItemInfo {
	s.ActionPrice = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetBePromotion(v bool) *DescribeItemsResponseBodyItemsItemInfo {
	s.BePromotion = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetBeSourceCode(v bool) *DescribeItemsResponseBodyItemsItemInfo {
	s.BeSourceCode = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetBrandName(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.BrandName = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetCategoryName(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.CategoryName = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetCompanyId(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.CompanyId = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetCustomizeFeatureA(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.CustomizeFeatureA = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetCustomizeFeatureB(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.CustomizeFeatureB = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetCustomizeFeatureC(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.CustomizeFeatureC = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetCustomizeFeatureD(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.CustomizeFeatureD = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetCustomizeFeatureE(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.CustomizeFeatureE = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetCustomizeFeatureF(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.CustomizeFeatureF = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetCustomizeFeatureG(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.CustomizeFeatureG = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetCustomizeFeatureH(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.CustomizeFeatureH = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetCustomizeFeatureI(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.CustomizeFeatureI = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetCustomizeFeatureJ(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.CustomizeFeatureJ = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetEnergyEfficiency(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.EnergyEfficiency = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetExtraAttribute(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.ExtraAttribute = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetForestFirstId(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.ForestFirstId = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetForestSecondId(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.ForestSecondId = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetItemBarCode(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.ItemBarCode = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetItemId(v int64) *DescribeItemsResponseBodyItemsItemInfo {
	s.ItemId = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetItemInfoIndex(v int32) *DescribeItemsResponseBodyItemsItemInfo {
	s.ItemInfoIndex = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetItemQrCode(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.ItemQrCode = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetItemShortTitle(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.ItemShortTitle = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetItemTitle(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.ItemTitle = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetMemberPrice(v int32) *DescribeItemsResponseBodyItemsItemInfo {
	s.MemberPrice = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetModelNumber(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.ModelNumber = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetOptionGroups(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.OptionGroups = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetOriginalPrice(v int32) *DescribeItemsResponseBodyItemsItemInfo {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetPriceUnit(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.PriceUnit = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetProductionPlace(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.ProductionPlace = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetPromotionEnd(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.PromotionEnd = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetPromotionReason(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.PromotionReason = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetPromotionStart(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.PromotionStart = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetPromotionText(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.PromotionText = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetRank(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.Rank = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetSaleSpec(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.SaleSpec = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetSkuId(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.SkuId = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetSourceCode(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.SourceCode = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetStoreId(v string) *DescribeItemsResponseBodyItemsItemInfo {
	s.StoreId = &v
	return s
}

func (s *DescribeItemsResponseBodyItemsItemInfo) SetSuggestPrice(v int32) *DescribeItemsResponseBodyItemsItemInfo {
	s.SuggestPrice = &v
	return s
}

type DescribeItemsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeItemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeItemsResponse) SetHeaders(v map[string]*string) *DescribeItemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeItemsResponse) SetStatusCode(v int32) *DescribeItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeItemsResponse) SetBody(v *DescribeItemsResponseBody) *DescribeItemsResponse {
	s.Body = v
	return s
}

type DescribeLogisticsRequest struct {
	// This parameter is required.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s DescribeLogisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogisticsRequest) SetOrderId(v string) *DescribeLogisticsRequest {
	s.OrderId = &v
	return s
}

type DescribeLogisticsResponseBody struct {
	Code           *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                                 `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                 `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Logistics      *DescribeLogisticsResponseBodyLogistics `json:"Logistics,omitempty" xml:"Logistics,omitempty" type:"Struct"`
	Message        *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeLogisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogisticsResponseBody) SetCode(v string) *DescribeLogisticsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogisticsResponseBody) SetDynamicCode(v string) *DescribeLogisticsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeLogisticsResponseBody) SetDynamicMessage(v string) *DescribeLogisticsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeLogisticsResponseBody) SetErrorCode(v string) *DescribeLogisticsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeLogisticsResponseBody) SetErrorMessage(v string) *DescribeLogisticsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeLogisticsResponseBody) SetLogistics(v *DescribeLogisticsResponseBodyLogistics) *DescribeLogisticsResponseBody {
	s.Logistics = v
	return s
}

func (s *DescribeLogisticsResponseBody) SetMessage(v string) *DescribeLogisticsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLogisticsResponseBody) SetRequestId(v string) *DescribeLogisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogisticsResponseBody) SetSuccess(v bool) *DescribeLogisticsResponseBody {
	s.Success = &v
	return s
}

type DescribeLogisticsResponseBodyLogistics struct {
	LogisticsInfo []*DescribeLogisticsResponseBodyLogisticsLogisticsInfo `json:"LogisticsInfo,omitempty" xml:"LogisticsInfo,omitempty" type:"Repeated"`
}

func (s DescribeLogisticsResponseBodyLogistics) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogisticsResponseBodyLogistics) GoString() string {
	return s.String()
}

func (s *DescribeLogisticsResponseBodyLogistics) SetLogisticsInfo(v []*DescribeLogisticsResponseBodyLogisticsLogisticsInfo) *DescribeLogisticsResponseBodyLogistics {
	s.LogisticsInfo = v
	return s
}

type DescribeLogisticsResponseBodyLogisticsLogisticsInfo struct {
	AcceptStatus       *bool   `json:"AcceptStatus,omitempty" xml:"AcceptStatus,omitempty"`
	ApMacList          *string `json:"ApMacList,omitempty" xml:"ApMacList,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EslMacList         *string `json:"EslMacList,omitempty" xml:"EslMacList,omitempty"`
	HasSend            *string `json:"HasSend,omitempty" xml:"HasSend,omitempty"`
	LogisticsDocuments *string `json:"LogisticsDocuments,omitempty" xml:"LogisticsDocuments,omitempty"`
	OrderId            *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PoNumber           *string `json:"PoNumber,omitempty" xml:"PoNumber,omitempty"`
	PrNumber           *string `json:"PrNumber,omitempty" xml:"PrNumber,omitempty"`
}

func (s DescribeLogisticsResponseBodyLogisticsLogisticsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogisticsResponseBodyLogisticsLogisticsInfo) GoString() string {
	return s.String()
}

func (s *DescribeLogisticsResponseBodyLogisticsLogisticsInfo) SetAcceptStatus(v bool) *DescribeLogisticsResponseBodyLogisticsLogisticsInfo {
	s.AcceptStatus = &v
	return s
}

func (s *DescribeLogisticsResponseBodyLogisticsLogisticsInfo) SetApMacList(v string) *DescribeLogisticsResponseBodyLogisticsLogisticsInfo {
	s.ApMacList = &v
	return s
}

func (s *DescribeLogisticsResponseBodyLogisticsLogisticsInfo) SetDescription(v string) *DescribeLogisticsResponseBodyLogisticsLogisticsInfo {
	s.Description = &v
	return s
}

func (s *DescribeLogisticsResponseBodyLogisticsLogisticsInfo) SetEslMacList(v string) *DescribeLogisticsResponseBodyLogisticsLogisticsInfo {
	s.EslMacList = &v
	return s
}

func (s *DescribeLogisticsResponseBodyLogisticsLogisticsInfo) SetHasSend(v string) *DescribeLogisticsResponseBodyLogisticsLogisticsInfo {
	s.HasSend = &v
	return s
}

func (s *DescribeLogisticsResponseBodyLogisticsLogisticsInfo) SetLogisticsDocuments(v string) *DescribeLogisticsResponseBodyLogisticsLogisticsInfo {
	s.LogisticsDocuments = &v
	return s
}

func (s *DescribeLogisticsResponseBodyLogisticsLogisticsInfo) SetOrderId(v string) *DescribeLogisticsResponseBodyLogisticsLogisticsInfo {
	s.OrderId = &v
	return s
}

func (s *DescribeLogisticsResponseBodyLogisticsLogisticsInfo) SetPoNumber(v string) *DescribeLogisticsResponseBodyLogisticsLogisticsInfo {
	s.PoNumber = &v
	return s
}

func (s *DescribeLogisticsResponseBodyLogisticsLogisticsInfo) SetPrNumber(v string) *DescribeLogisticsResponseBodyLogisticsLogisticsInfo {
	s.PrNumber = &v
	return s
}

type DescribeLogisticsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogisticsResponse) SetHeaders(v map[string]*string) *DescribeLogisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogisticsResponse) SetStatusCode(v int32) *DescribeLogisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogisticsResponse) SetBody(v *DescribeLogisticsResponseBody) *DescribeLogisticsResponse {
	s.Body = v
	return s
}

type DescribePayOrdersRequest struct {
	FromDate   *string `json:"FromDate,omitempty" xml:"FromDate,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ToDate     *string `json:"ToDate,omitempty" xml:"ToDate,omitempty"`
}

func (s DescribePayOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePayOrdersRequest) GoString() string {
	return s.String()
}

func (s *DescribePayOrdersRequest) SetFromDate(v string) *DescribePayOrdersRequest {
	s.FromDate = &v
	return s
}

func (s *DescribePayOrdersRequest) SetOrderId(v string) *DescribePayOrdersRequest {
	s.OrderId = &v
	return s
}

func (s *DescribePayOrdersRequest) SetPageNumber(v int32) *DescribePayOrdersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePayOrdersRequest) SetPageSize(v int32) *DescribePayOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePayOrdersRequest) SetToDate(v string) *DescribePayOrdersRequest {
	s.ToDate = &v
	return s
}

type DescribePayOrdersResponseBody struct {
	Code           *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                                 `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                                 `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PayOrders      *DescribePayOrdersResponseBodyPayOrders `json:"PayOrders,omitempty" xml:"PayOrders,omitempty" type:"Struct"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePayOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePayOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePayOrdersResponseBody) SetCode(v string) *DescribePayOrdersResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePayOrdersResponseBody) SetDynamicCode(v string) *DescribePayOrdersResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribePayOrdersResponseBody) SetDynamicMessage(v string) *DescribePayOrdersResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribePayOrdersResponseBody) SetErrorCode(v string) *DescribePayOrdersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribePayOrdersResponseBody) SetErrorMessage(v string) *DescribePayOrdersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribePayOrdersResponseBody) SetMessage(v string) *DescribePayOrdersResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePayOrdersResponseBody) SetPageNumber(v int32) *DescribePayOrdersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePayOrdersResponseBody) SetPageSize(v int32) *DescribePayOrdersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePayOrdersResponseBody) SetPayOrders(v *DescribePayOrdersResponseBodyPayOrders) *DescribePayOrdersResponseBody {
	s.PayOrders = v
	return s
}

func (s *DescribePayOrdersResponseBody) SetRequestId(v string) *DescribePayOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePayOrdersResponseBody) SetSuccess(v bool) *DescribePayOrdersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePayOrdersResponseBody) SetTotalCount(v int32) *DescribePayOrdersResponseBody {
	s.TotalCount = &v
	return s
}

type DescribePayOrdersResponseBodyPayOrders struct {
	PayOrderInfo []*DescribePayOrdersResponseBodyPayOrdersPayOrderInfo `json:"PayOrderInfo,omitempty" xml:"PayOrderInfo,omitempty" type:"Repeated"`
}

func (s DescribePayOrdersResponseBodyPayOrders) String() string {
	return tea.Prettify(s)
}

func (s DescribePayOrdersResponseBodyPayOrders) GoString() string {
	return s.String()
}

func (s *DescribePayOrdersResponseBodyPayOrders) SetPayOrderInfo(v []*DescribePayOrdersResponseBodyPayOrdersPayOrderInfo) *DescribePayOrdersResponseBodyPayOrders {
	s.PayOrderInfo = v
	return s
}

type DescribePayOrdersResponseBodyPayOrdersPayOrderInfo struct {
	CommodityCode  *string  `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CommodityName  *string  `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	DetailName     *string  `json:"DetailName,omitempty" xml:"DetailName,omitempty"`
	GmtCreate      *string  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtPay         *string  `json:"GmtPay,omitempty" xml:"GmtPay,omitempty"`
	IsAccepted     *bool    `json:"IsAccepted,omitempty" xml:"IsAccepted,omitempty"`
	OrderId        *string  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderStatus    *string  `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	OrderType      *string  `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	PayAmount      *float32 `json:"PayAmount,omitempty" xml:"PayAmount,omitempty"`
	Quantity       *int32   `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s DescribePayOrdersResponseBodyPayOrdersPayOrderInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePayOrdersResponseBodyPayOrdersPayOrderInfo) GoString() string {
	return s.String()
}

func (s *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo) SetCommodityCode(v string) *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo {
	s.CommodityCode = &v
	return s
}

func (s *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo) SetCommodityName(v string) *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo {
	s.CommodityName = &v
	return s
}

func (s *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo) SetDetailName(v string) *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo {
	s.DetailName = &v
	return s
}

func (s *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo) SetGmtCreate(v string) *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo {
	s.GmtCreate = &v
	return s
}

func (s *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo) SetGmtPay(v string) *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo {
	s.GmtPay = &v
	return s
}

func (s *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo) SetIsAccepted(v bool) *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo {
	s.IsAccepted = &v
	return s
}

func (s *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo) SetOrderId(v string) *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo {
	s.OrderId = &v
	return s
}

func (s *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo) SetOrderStatus(v string) *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo {
	s.OrderStatus = &v
	return s
}

func (s *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo) SetOrderType(v string) *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo {
	s.OrderType = &v
	return s
}

func (s *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo) SetOriginalAmount(v float32) *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo {
	s.OriginalAmount = &v
	return s
}

func (s *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo) SetPayAmount(v float32) *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo {
	s.PayAmount = &v
	return s
}

func (s *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo) SetQuantity(v int32) *DescribePayOrdersResponseBodyPayOrdersPayOrderInfo {
	s.Quantity = &v
	return s
}

type DescribePayOrdersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePayOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePayOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePayOrdersResponse) GoString() string {
	return s.String()
}

func (s *DescribePayOrdersResponse) SetHeaders(v map[string]*string) *DescribePayOrdersResponse {
	s.Headers = v
	return s
}

func (s *DescribePayOrdersResponse) SetStatusCode(v int32) *DescribePayOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePayOrdersResponse) SetBody(v *DescribePayOrdersResponseBody) *DescribePayOrdersResponse {
	s.Body = v
	return s
}

type DescribePlanogramRailsRequest struct {
	Layer      *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RailCode   *string `json:"RailCode,omitempty" xml:"RailCode,omitempty"`
	Shelf      *string `json:"Shelf,omitempty" xml:"Shelf,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DescribePlanogramRailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlanogramRailsRequest) GoString() string {
	return s.String()
}

func (s *DescribePlanogramRailsRequest) SetLayer(v string) *DescribePlanogramRailsRequest {
	s.Layer = &v
	return s
}

func (s *DescribePlanogramRailsRequest) SetPageNumber(v int32) *DescribePlanogramRailsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePlanogramRailsRequest) SetPageSize(v int32) *DescribePlanogramRailsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePlanogramRailsRequest) SetRailCode(v string) *DescribePlanogramRailsRequest {
	s.RailCode = &v
	return s
}

func (s *DescribePlanogramRailsRequest) SetShelf(v string) *DescribePlanogramRailsRequest {
	s.Shelf = &v
	return s
}

func (s *DescribePlanogramRailsRequest) SetStoreId(v string) *DescribePlanogramRailsRequest {
	s.StoreId = &v
	return s
}

type DescribePlanogramRailsResponseBody struct {
	Code               *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string                                                 `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                                                 `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode          *string                                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage       *string                                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message            *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber         *int32                                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlanogramRailInfos []*DescribePlanogramRailsResponseBodyPlanogramRailInfos `json:"PlanogramRailInfos,omitempty" xml:"PlanogramRailInfos,omitempty" type:"Repeated"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StoreId            *string                                                 `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	Success            *bool                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount         *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePlanogramRailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlanogramRailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlanogramRailsResponseBody) SetCode(v string) *DescribePlanogramRailsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePlanogramRailsResponseBody) SetDynamicCode(v string) *DescribePlanogramRailsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribePlanogramRailsResponseBody) SetDynamicMessage(v string) *DescribePlanogramRailsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribePlanogramRailsResponseBody) SetErrorCode(v string) *DescribePlanogramRailsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribePlanogramRailsResponseBody) SetErrorMessage(v string) *DescribePlanogramRailsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribePlanogramRailsResponseBody) SetMessage(v string) *DescribePlanogramRailsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePlanogramRailsResponseBody) SetPageNumber(v int32) *DescribePlanogramRailsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePlanogramRailsResponseBody) SetPageSize(v int32) *DescribePlanogramRailsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePlanogramRailsResponseBody) SetPlanogramRailInfos(v []*DescribePlanogramRailsResponseBodyPlanogramRailInfos) *DescribePlanogramRailsResponseBody {
	s.PlanogramRailInfos = v
	return s
}

func (s *DescribePlanogramRailsResponseBody) SetRequestId(v string) *DescribePlanogramRailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlanogramRailsResponseBody) SetStoreId(v string) *DescribePlanogramRailsResponseBody {
	s.StoreId = &v
	return s
}

func (s *DescribePlanogramRailsResponseBody) SetSuccess(v bool) *DescribePlanogramRailsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePlanogramRailsResponseBody) SetTotalCount(v int32) *DescribePlanogramRailsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribePlanogramRailsResponseBodyPlanogramRailInfos struct {
	GapUnit  *int32  `json:"GapUnit,omitempty" xml:"GapUnit,omitempty"`
	Layer    *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	RailCode *string `json:"RailCode,omitempty" xml:"RailCode,omitempty"`
	Shelf    *string `json:"Shelf,omitempty" xml:"Shelf,omitempty"`
}

func (s DescribePlanogramRailsResponseBodyPlanogramRailInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribePlanogramRailsResponseBodyPlanogramRailInfos) GoString() string {
	return s.String()
}

func (s *DescribePlanogramRailsResponseBodyPlanogramRailInfos) SetGapUnit(v int32) *DescribePlanogramRailsResponseBodyPlanogramRailInfos {
	s.GapUnit = &v
	return s
}

func (s *DescribePlanogramRailsResponseBodyPlanogramRailInfos) SetLayer(v string) *DescribePlanogramRailsResponseBodyPlanogramRailInfos {
	s.Layer = &v
	return s
}

func (s *DescribePlanogramRailsResponseBodyPlanogramRailInfos) SetRailCode(v string) *DescribePlanogramRailsResponseBodyPlanogramRailInfos {
	s.RailCode = &v
	return s
}

func (s *DescribePlanogramRailsResponseBodyPlanogramRailInfos) SetShelf(v string) *DescribePlanogramRailsResponseBodyPlanogramRailInfos {
	s.Shelf = &v
	return s
}

type DescribePlanogramRailsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlanogramRailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlanogramRailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlanogramRailsResponse) GoString() string {
	return s.String()
}

func (s *DescribePlanogramRailsResponse) SetHeaders(v map[string]*string) *DescribePlanogramRailsResponse {
	s.Headers = v
	return s
}

func (s *DescribePlanogramRailsResponse) SetStatusCode(v int32) *DescribePlanogramRailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlanogramRailsResponse) SetBody(v *DescribePlanogramRailsResponseBody) *DescribePlanogramRailsResponse {
	s.Body = v
	return s
}

type DescribeStoresRequest struct {
	Brand      *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	CompanyId  *string `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	FromDate   *string `json:"FromDate,omitempty" xml:"FromDate,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StoreId    *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	StoreName  *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	ToDate     *string `json:"ToDate,omitempty" xml:"ToDate,omitempty"`
}

func (s DescribeStoresRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoresRequest) GoString() string {
	return s.String()
}

func (s *DescribeStoresRequest) SetBrand(v string) *DescribeStoresRequest {
	s.Brand = &v
	return s
}

func (s *DescribeStoresRequest) SetCompanyId(v string) *DescribeStoresRequest {
	s.CompanyId = &v
	return s
}

func (s *DescribeStoresRequest) SetFromDate(v string) *DescribeStoresRequest {
	s.FromDate = &v
	return s
}

func (s *DescribeStoresRequest) SetPageNumber(v int32) *DescribeStoresRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeStoresRequest) SetPageSize(v int32) *DescribeStoresRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStoresRequest) SetStoreId(v string) *DescribeStoresRequest {
	s.StoreId = &v
	return s
}

func (s *DescribeStoresRequest) SetStoreName(v string) *DescribeStoresRequest {
	s.StoreName = &v
	return s
}

func (s *DescribeStoresRequest) SetToDate(v string) *DescribeStoresRequest {
	s.ToDate = &v
	return s
}

type DescribeStoresResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                           `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                           `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Stores         *DescribeStoresResponseBodyStores `json:"Stores,omitempty" xml:"Stores,omitempty" type:"Struct"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeStoresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoresResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStoresResponseBody) SetCode(v string) *DescribeStoresResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeStoresResponseBody) SetDynamicCode(v string) *DescribeStoresResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeStoresResponseBody) SetDynamicMessage(v string) *DescribeStoresResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeStoresResponseBody) SetErrorCode(v string) *DescribeStoresResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeStoresResponseBody) SetErrorMessage(v string) *DescribeStoresResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeStoresResponseBody) SetMessage(v string) *DescribeStoresResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeStoresResponseBody) SetPageNumber(v int32) *DescribeStoresResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeStoresResponseBody) SetPageSize(v int32) *DescribeStoresResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStoresResponseBody) SetRequestId(v string) *DescribeStoresResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStoresResponseBody) SetStores(v *DescribeStoresResponseBodyStores) *DescribeStoresResponseBody {
	s.Stores = v
	return s
}

func (s *DescribeStoresResponseBody) SetSuccess(v bool) *DescribeStoresResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeStoresResponseBody) SetTotalCount(v int32) *DescribeStoresResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeStoresResponseBodyStores struct {
	StoreInfo []*DescribeStoresResponseBodyStoresStoreInfo `json:"StoreInfo,omitempty" xml:"StoreInfo,omitempty" type:"Repeated"`
}

func (s DescribeStoresResponseBodyStores) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoresResponseBodyStores) GoString() string {
	return s.String()
}

func (s *DescribeStoresResponseBodyStores) SetStoreInfo(v []*DescribeStoresResponseBodyStoresStoreInfo) *DescribeStoresResponseBodyStores {
	s.StoreInfo = v
	return s
}

type DescribeStoresResponseBodyStoresStoreInfo struct {
	Brand       *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	Comments    *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CompanyId   *string `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Groups      *string `json:"Groups,omitempty" xml:"Groups,omitempty"`
	Level       *string `json:"Level,omitempty" xml:"Level,omitempty"`
	OutId       *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	ParentId    *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	StoreId     *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	StoreName   *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
}

func (s DescribeStoresResponseBodyStoresStoreInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoresResponseBodyStoresStoreInfo) GoString() string {
	return s.String()
}

func (s *DescribeStoresResponseBodyStoresStoreInfo) SetBrand(v string) *DescribeStoresResponseBodyStoresStoreInfo {
	s.Brand = &v
	return s
}

func (s *DescribeStoresResponseBodyStoresStoreInfo) SetComments(v string) *DescribeStoresResponseBodyStoresStoreInfo {
	s.Comments = &v
	return s
}

func (s *DescribeStoresResponseBodyStoresStoreInfo) SetCompanyId(v string) *DescribeStoresResponseBodyStoresStoreInfo {
	s.CompanyId = &v
	return s
}

func (s *DescribeStoresResponseBodyStoresStoreInfo) SetGmtCreate(v string) *DescribeStoresResponseBodyStoresStoreInfo {
	s.GmtCreate = &v
	return s
}

func (s *DescribeStoresResponseBodyStoresStoreInfo) SetGmtModified(v string) *DescribeStoresResponseBodyStoresStoreInfo {
	s.GmtModified = &v
	return s
}

func (s *DescribeStoresResponseBodyStoresStoreInfo) SetGroups(v string) *DescribeStoresResponseBodyStoresStoreInfo {
	s.Groups = &v
	return s
}

func (s *DescribeStoresResponseBodyStoresStoreInfo) SetLevel(v string) *DescribeStoresResponseBodyStoresStoreInfo {
	s.Level = &v
	return s
}

func (s *DescribeStoresResponseBodyStoresStoreInfo) SetOutId(v string) *DescribeStoresResponseBodyStoresStoreInfo {
	s.OutId = &v
	return s
}

func (s *DescribeStoresResponseBodyStoresStoreInfo) SetParentId(v string) *DescribeStoresResponseBodyStoresStoreInfo {
	s.ParentId = &v
	return s
}

func (s *DescribeStoresResponseBodyStoresStoreInfo) SetPhone(v string) *DescribeStoresResponseBodyStoresStoreInfo {
	s.Phone = &v
	return s
}

func (s *DescribeStoresResponseBodyStoresStoreInfo) SetStoreId(v string) *DescribeStoresResponseBodyStoresStoreInfo {
	s.StoreId = &v
	return s
}

func (s *DescribeStoresResponseBodyStoresStoreInfo) SetStoreName(v string) *DescribeStoresResponseBodyStoresStoreInfo {
	s.StoreName = &v
	return s
}

type DescribeStoresResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStoresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStoresResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStoresResponse) GoString() string {
	return s.String()
}

func (s *DescribeStoresResponse) SetHeaders(v map[string]*string) *DescribeStoresResponse {
	s.Headers = v
	return s
}

func (s *DescribeStoresResponse) SetStatusCode(v int32) *DescribeStoresResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStoresResponse) SetBody(v *DescribeStoresResponseBody) *DescribeStoresResponse {
	s.Body = v
	return s
}

type DescribeUserLogRequest struct {
	EslBarCode    *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	FromDate      *string `json:"FromDate,omitempty" xml:"FromDate,omitempty"`
	ItemBarCode   *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	ItemId        *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemTitle     *string `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	OperateStatus *string `json:"OperateStatus,omitempty" xml:"OperateStatus,omitempty"`
	OperateType   *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	OperateUserId *int64  `json:"OperateUserId,omitempty" xml:"OperateUserId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Reverse       *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	ToDate  *string `json:"ToDate,omitempty" xml:"ToDate,omitempty"`
}

func (s DescribeUserLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserLogRequest) SetEslBarCode(v string) *DescribeUserLogRequest {
	s.EslBarCode = &v
	return s
}

func (s *DescribeUserLogRequest) SetFromDate(v string) *DescribeUserLogRequest {
	s.FromDate = &v
	return s
}

func (s *DescribeUserLogRequest) SetItemBarCode(v string) *DescribeUserLogRequest {
	s.ItemBarCode = &v
	return s
}

func (s *DescribeUserLogRequest) SetItemId(v int64) *DescribeUserLogRequest {
	s.ItemId = &v
	return s
}

func (s *DescribeUserLogRequest) SetItemTitle(v string) *DescribeUserLogRequest {
	s.ItemTitle = &v
	return s
}

func (s *DescribeUserLogRequest) SetOperateStatus(v string) *DescribeUserLogRequest {
	s.OperateStatus = &v
	return s
}

func (s *DescribeUserLogRequest) SetOperateType(v string) *DescribeUserLogRequest {
	s.OperateType = &v
	return s
}

func (s *DescribeUserLogRequest) SetOperateUserId(v int64) *DescribeUserLogRequest {
	s.OperateUserId = &v
	return s
}

func (s *DescribeUserLogRequest) SetPageNumber(v int32) *DescribeUserLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserLogRequest) SetPageSize(v int32) *DescribeUserLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserLogRequest) SetReverse(v bool) *DescribeUserLogRequest {
	s.Reverse = &v
	return s
}

func (s *DescribeUserLogRequest) SetStoreId(v string) *DescribeUserLogRequest {
	s.StoreId = &v
	return s
}

func (s *DescribeUserLogRequest) SetToDate(v string) *DescribeUserLogRequest {
	s.ToDate = &v
	return s
}

type DescribeUserLogResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                              `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                              `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserLogs       *DescribeUserLogResponseBodyUserLogs `json:"UserLogs,omitempty" xml:"UserLogs,omitempty" type:"Struct"`
}

func (s DescribeUserLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserLogResponseBody) SetCode(v string) *DescribeUserLogResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetDynamicCode(v string) *DescribeUserLogResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetDynamicMessage(v string) *DescribeUserLogResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetErrorCode(v string) *DescribeUserLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetErrorMessage(v string) *DescribeUserLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetMessage(v string) *DescribeUserLogResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetPageNumber(v int32) *DescribeUserLogResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetPageSize(v int32) *DescribeUserLogResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetRequestId(v string) *DescribeUserLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetSuccess(v bool) *DescribeUserLogResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetTotalCount(v int32) *DescribeUserLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUserLogResponseBody) SetUserLogs(v *DescribeUserLogResponseBodyUserLogs) *DescribeUserLogResponseBody {
	s.UserLogs = v
	return s
}

type DescribeUserLogResponseBodyUserLogs struct {
	UserLogInfo []*DescribeUserLogResponseBodyUserLogsUserLogInfo `json:"UserLogInfo,omitempty" xml:"UserLogInfo,omitempty" type:"Repeated"`
}

func (s DescribeUserLogResponseBodyUserLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserLogResponseBodyUserLogs) GoString() string {
	return s.String()
}

func (s *DescribeUserLogResponseBodyUserLogs) SetUserLogInfo(v []*DescribeUserLogResponseBodyUserLogsUserLogInfo) *DescribeUserLogResponseBodyUserLogs {
	s.UserLogInfo = v
	return s
}

type DescribeUserLogResponseBodyUserLogsUserLogInfo struct {
	EslBarCode      *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	ItemActionPrice *int32  `json:"ItemActionPrice,omitempty" xml:"ItemActionPrice,omitempty"`
	ItemBarCode     *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	ItemId          *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemTitle       *string `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	Mac             *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	OperateStatus   *string `json:"OperateStatus,omitempty" xml:"OperateStatus,omitempty"`
	OperateTime     *string `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	OperateType     *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	OperateUserId   *int64  `json:"OperateUserId,omitempty" xml:"OperateUserId,omitempty"`
	ShelfCode       *string `json:"ShelfCode,omitempty" xml:"ShelfCode,omitempty"`
	StoreId         *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s DescribeUserLogResponseBodyUserLogsUserLogInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserLogResponseBodyUserLogsUserLogInfo) GoString() string {
	return s.String()
}

func (s *DescribeUserLogResponseBodyUserLogsUserLogInfo) SetEslBarCode(v string) *DescribeUserLogResponseBodyUserLogsUserLogInfo {
	s.EslBarCode = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogsUserLogInfo) SetItemActionPrice(v int32) *DescribeUserLogResponseBodyUserLogsUserLogInfo {
	s.ItemActionPrice = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogsUserLogInfo) SetItemBarCode(v string) *DescribeUserLogResponseBodyUserLogsUserLogInfo {
	s.ItemBarCode = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogsUserLogInfo) SetItemId(v int64) *DescribeUserLogResponseBodyUserLogsUserLogInfo {
	s.ItemId = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogsUserLogInfo) SetItemTitle(v string) *DescribeUserLogResponseBodyUserLogsUserLogInfo {
	s.ItemTitle = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogsUserLogInfo) SetMac(v string) *DescribeUserLogResponseBodyUserLogsUserLogInfo {
	s.Mac = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogsUserLogInfo) SetOperateStatus(v string) *DescribeUserLogResponseBodyUserLogsUserLogInfo {
	s.OperateStatus = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogsUserLogInfo) SetOperateTime(v string) *DescribeUserLogResponseBodyUserLogsUserLogInfo {
	s.OperateTime = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogsUserLogInfo) SetOperateType(v string) *DescribeUserLogResponseBodyUserLogsUserLogInfo {
	s.OperateType = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogsUserLogInfo) SetOperateUserId(v int64) *DescribeUserLogResponseBodyUserLogsUserLogInfo {
	s.OperateUserId = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogsUserLogInfo) SetShelfCode(v string) *DescribeUserLogResponseBodyUserLogsUserLogInfo {
	s.ShelfCode = &v
	return s
}

func (s *DescribeUserLogResponseBodyUserLogsUserLogInfo) SetStoreId(v string) *DescribeUserLogResponseBodyUserLogsUserLogInfo {
	s.StoreId = &v
	return s
}

type DescribeUserLogResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserLogResponse) SetHeaders(v map[string]*string) *DescribeUserLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserLogResponse) SetStatusCode(v int32) *DescribeUserLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserLogResponse) SetBody(v *DescribeUserLogResponseBody) *DescribeUserLogResponse {
	s.Body = v
	return s
}

type DescribeUsersRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserType   *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsersRequest) SetPageNumber(v int32) *DescribeUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUsersRequest) SetPageSize(v int32) *DescribeUsersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUsersRequest) SetUserId(v string) *DescribeUsersRequest {
	s.UserId = &v
	return s
}

func (s *DescribeUsersRequest) SetUserName(v string) *DescribeUsersRequest {
	s.UserName = &v
	return s
}

func (s *DescribeUsersRequest) SetUserType(v string) *DescribeUsersRequest {
	s.UserType = &v
	return s
}

type DescribeUsersResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                         `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                         `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Users          *DescribeUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s DescribeUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBody) SetCode(v string) *DescribeUsersResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUsersResponseBody) SetDynamicCode(v string) *DescribeUsersResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeUsersResponseBody) SetDynamicMessage(v string) *DescribeUsersResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeUsersResponseBody) SetErrorCode(v string) *DescribeUsersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeUsersResponseBody) SetErrorMessage(v string) *DescribeUsersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeUsersResponseBody) SetMessage(v string) *DescribeUsersResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeUsersResponseBody) SetPageNumber(v int32) *DescribeUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUsersResponseBody) SetPageSize(v int32) *DescribeUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUsersResponseBody) SetRequestId(v string) *DescribeUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsersResponseBody) SetSuccess(v bool) *DescribeUsersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeUsersResponseBody) SetTotalCount(v int32) *DescribeUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUsersResponseBody) SetUsers(v *DescribeUsersResponseBodyUsers) *DescribeUsersResponseBody {
	s.Users = v
	return s
}

type DescribeUsersResponseBodyUsers struct {
	UserInfo []*DescribeUsersResponseBodyUsersUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Repeated"`
}

func (s DescribeUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBodyUsers) SetUserInfo(v []*DescribeUsersResponseBodyUsersUserInfo) *DescribeUsersResponseBodyUsers {
	s.UserInfo = v
	return s
}

type DescribeUsersResponseBodyUsersUserInfo struct {
	Stores   *string `json:"Stores,omitempty" xml:"Stores,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeUsersResponseBodyUsersUserInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponseBodyUsersUserInfo) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBodyUsersUserInfo) SetStores(v string) *DescribeUsersResponseBodyUsersUserInfo {
	s.Stores = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersUserInfo) SetUserId(v string) *DescribeUsersResponseBodyUsersUserInfo {
	s.UserId = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersUserInfo) SetUserName(v string) *DescribeUsersResponseBodyUsersUserInfo {
	s.UserName = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersUserInfo) SetUserType(v string) *DescribeUsersResponseBodyUsersUserInfo {
	s.UserType = &v
	return s
}

type DescribeUsersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponse) SetHeaders(v map[string]*string) *DescribeUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsersResponse) SetStatusCode(v int32) *DescribeUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsersResponse) SetBody(v *DescribeUsersResponseBody) *DescribeUsersResponse {
	s.Body = v
	return s
}

type GetCompanyResponseBody struct {
	Bid            *int64  `json:"Bid,omitempty" xml:"Bid,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CompanyId      *string `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	CompanyType    *string `json:"CompanyType,omitempty" xml:"CompanyType,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCompanyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCompanyResponseBody) GoString() string {
	return s.String()
}

func (s *GetCompanyResponseBody) SetBid(v int64) *GetCompanyResponseBody {
	s.Bid = &v
	return s
}

func (s *GetCompanyResponseBody) SetCode(v string) *GetCompanyResponseBody {
	s.Code = &v
	return s
}

func (s *GetCompanyResponseBody) SetCompanyId(v string) *GetCompanyResponseBody {
	s.CompanyId = &v
	return s
}

func (s *GetCompanyResponseBody) SetCompanyType(v string) *GetCompanyResponseBody {
	s.CompanyType = &v
	return s
}

func (s *GetCompanyResponseBody) SetDynamicCode(v string) *GetCompanyResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetCompanyResponseBody) SetDynamicMessage(v string) *GetCompanyResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetCompanyResponseBody) SetErrorCode(v string) *GetCompanyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCompanyResponseBody) SetErrorMessage(v string) *GetCompanyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetCompanyResponseBody) SetMessage(v string) *GetCompanyResponseBody {
	s.Message = &v
	return s
}

func (s *GetCompanyResponseBody) SetOwnerId(v int64) *GetCompanyResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetCompanyResponseBody) SetRequestId(v string) *GetCompanyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCompanyResponseBody) SetStatus(v string) *GetCompanyResponseBody {
	s.Status = &v
	return s
}

func (s *GetCompanyResponseBody) SetSuccess(v bool) *GetCompanyResponseBody {
	s.Success = &v
	return s
}

type GetCompanyResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCompanyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCompanyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCompanyResponse) GoString() string {
	return s.String()
}

func (s *GetCompanyResponse) SetHeaders(v map[string]*string) *GetCompanyResponse {
	s.Headers = v
	return s
}

func (s *GetCompanyResponse) SetStatusCode(v int32) *GetCompanyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCompanyResponse) SetBody(v *GetCompanyResponseBody) *GetCompanyResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetUserId(v string) *GetUserRequest {
	s.UserId = &v
	return s
}

type GetUserResponseBody struct {
	Code           *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string                  `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string                  `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
	User           *GetUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetCode(v string) *GetUserResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserResponseBody) SetDynamicCode(v string) *GetUserResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetUserResponseBody) SetDynamicMessage(v string) *GetUserResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetUserResponseBody) SetErrorCode(v string) *GetUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserResponseBody) SetErrorMessage(v string) *GetUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserResponseBody) SetMessage(v string) *GetUserResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetSuccess(v bool) *GetUserResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserResponseBody) SetUser(v *GetUserResponseBodyUser) *GetUserResponseBody {
	s.User = v
	return s
}

type GetUserResponseBodyUser struct {
	Stores   *string `json:"Stores,omitempty" xml:"Stores,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetUserResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUser) SetStores(v string) *GetUserResponseBodyUser {
	s.Stores = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserId(v string) *GetUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserName(v string) *GetUserResponseBodyUser {
	s.UserName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserType(v string) *GetUserResponseBodyUser {
	s.UserType = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type MapPlanogramRailRequest struct {
	// This parameter is required.
	Layer *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// This parameter is required.
	RailCode *string `json:"RailCode,omitempty" xml:"RailCode,omitempty"`
	// This parameter is required.
	Shelf *string `json:"Shelf,omitempty" xml:"Shelf,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s MapPlanogramRailRequest) String() string {
	return tea.Prettify(s)
}

func (s MapPlanogramRailRequest) GoString() string {
	return s.String()
}

func (s *MapPlanogramRailRequest) SetLayer(v string) *MapPlanogramRailRequest {
	s.Layer = &v
	return s
}

func (s *MapPlanogramRailRequest) SetRailCode(v string) *MapPlanogramRailRequest {
	s.RailCode = &v
	return s
}

func (s *MapPlanogramRailRequest) SetShelf(v string) *MapPlanogramRailRequest {
	s.Shelf = &v
	return s
}

func (s *MapPlanogramRailRequest) SetStoreId(v string) *MapPlanogramRailRequest {
	s.StoreId = &v
	return s
}

type MapPlanogramRailResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MapPlanogramRailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MapPlanogramRailResponseBody) GoString() string {
	return s.String()
}

func (s *MapPlanogramRailResponseBody) SetCode(v string) *MapPlanogramRailResponseBody {
	s.Code = &v
	return s
}

func (s *MapPlanogramRailResponseBody) SetDynamicCode(v string) *MapPlanogramRailResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *MapPlanogramRailResponseBody) SetDynamicMessage(v string) *MapPlanogramRailResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *MapPlanogramRailResponseBody) SetErrorCode(v string) *MapPlanogramRailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *MapPlanogramRailResponseBody) SetErrorMessage(v string) *MapPlanogramRailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *MapPlanogramRailResponseBody) SetMessage(v string) *MapPlanogramRailResponseBody {
	s.Message = &v
	return s
}

func (s *MapPlanogramRailResponseBody) SetRequestId(v string) *MapPlanogramRailResponseBody {
	s.RequestId = &v
	return s
}

func (s *MapPlanogramRailResponseBody) SetSuccess(v bool) *MapPlanogramRailResponseBody {
	s.Success = &v
	return s
}

type MapPlanogramRailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MapPlanogramRailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MapPlanogramRailResponse) String() string {
	return tea.Prettify(s)
}

func (s MapPlanogramRailResponse) GoString() string {
	return s.String()
}

func (s *MapPlanogramRailResponse) SetHeaders(v map[string]*string) *MapPlanogramRailResponse {
	s.Headers = v
	return s
}

func (s *MapPlanogramRailResponse) SetStatusCode(v int32) *MapPlanogramRailResponse {
	s.StatusCode = &v
	return s
}

func (s *MapPlanogramRailResponse) SetBody(v *MapPlanogramRailResponseBody) *MapPlanogramRailResponse {
	s.Body = v
	return s
}

type RefreshTaobaoItemRequest struct {
	OuterId *string `json:"OuterId,omitempty" xml:"OuterId,omitempty"`
	SkuId   *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	// This parameter is required.
	StoreId      *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	TaobaoItemId *string `json:"TaobaoItemId,omitempty" xml:"TaobaoItemId,omitempty"`
}

func (s RefreshTaobaoItemRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshTaobaoItemRequest) GoString() string {
	return s.String()
}

func (s *RefreshTaobaoItemRequest) SetOuterId(v string) *RefreshTaobaoItemRequest {
	s.OuterId = &v
	return s
}

func (s *RefreshTaobaoItemRequest) SetSkuId(v string) *RefreshTaobaoItemRequest {
	s.SkuId = &v
	return s
}

func (s *RefreshTaobaoItemRequest) SetStoreId(v string) *RefreshTaobaoItemRequest {
	s.StoreId = &v
	return s
}

func (s *RefreshTaobaoItemRequest) SetTaobaoItemId(v string) *RefreshTaobaoItemRequest {
	s.TaobaoItemId = &v
	return s
}

type RefreshTaobaoItemResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RefreshTaobaoItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshTaobaoItemResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshTaobaoItemResponseBody) SetCode(v string) *RefreshTaobaoItemResponseBody {
	s.Code = &v
	return s
}

func (s *RefreshTaobaoItemResponseBody) SetDynamicCode(v string) *RefreshTaobaoItemResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RefreshTaobaoItemResponseBody) SetDynamicMessage(v string) *RefreshTaobaoItemResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RefreshTaobaoItemResponseBody) SetErrorCode(v string) *RefreshTaobaoItemResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RefreshTaobaoItemResponseBody) SetErrorMessage(v string) *RefreshTaobaoItemResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RefreshTaobaoItemResponseBody) SetMessage(v string) *RefreshTaobaoItemResponseBody {
	s.Message = &v
	return s
}

func (s *RefreshTaobaoItemResponseBody) SetRequestId(v string) *RefreshTaobaoItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshTaobaoItemResponseBody) SetSuccess(v bool) *RefreshTaobaoItemResponseBody {
	s.Success = &v
	return s
}

type RefreshTaobaoItemResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshTaobaoItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshTaobaoItemResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshTaobaoItemResponse) GoString() string {
	return s.String()
}

func (s *RefreshTaobaoItemResponse) SetHeaders(v map[string]*string) *RefreshTaobaoItemResponse {
	s.Headers = v
	return s
}

func (s *RefreshTaobaoItemResponse) SetStatusCode(v int32) *RefreshTaobaoItemResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshTaobaoItemResponse) SetBody(v *RefreshTaobaoItemResponseBody) *RefreshTaobaoItemResponse {
	s.Body = v
	return s
}

type UnassignUserRequest struct {
	// This parameter is required.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UnassignUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UnassignUserRequest) GoString() string {
	return s.String()
}

func (s *UnassignUserRequest) SetUserId(v string) *UnassignUserRequest {
	s.UserId = &v
	return s
}

type UnassignUserResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnassignUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnassignUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnassignUserResponseBody) SetCode(v string) *UnassignUserResponseBody {
	s.Code = &v
	return s
}

func (s *UnassignUserResponseBody) SetDynamicCode(v string) *UnassignUserResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UnassignUserResponseBody) SetDynamicMessage(v string) *UnassignUserResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UnassignUserResponseBody) SetErrorCode(v string) *UnassignUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UnassignUserResponseBody) SetErrorMessage(v string) *UnassignUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UnassignUserResponseBody) SetMessage(v string) *UnassignUserResponseBody {
	s.Message = &v
	return s
}

func (s *UnassignUserResponseBody) SetRequestId(v string) *UnassignUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassignUserResponseBody) SetSuccess(v bool) *UnassignUserResponseBody {
	s.Success = &v
	return s
}

type UnassignUserResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassignUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassignUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UnassignUserResponse) GoString() string {
	return s.String()
}

func (s *UnassignUserResponse) SetHeaders(v map[string]*string) *UnassignUserResponse {
	s.Headers = v
	return s
}

func (s *UnassignUserResponse) SetStatusCode(v int32) *UnassignUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassignUserResponse) SetBody(v *UnassignUserResponseBody) *UnassignUserResponse {
	s.Body = v
	return s
}

type UnbindEslDeviceRequest struct {
	// This parameter is required.
	EslBarCode  *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	ItemBarCode *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s UnbindEslDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindEslDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnbindEslDeviceRequest) SetEslBarCode(v string) *UnbindEslDeviceRequest {
	s.EslBarCode = &v
	return s
}

func (s *UnbindEslDeviceRequest) SetItemBarCode(v string) *UnbindEslDeviceRequest {
	s.ItemBarCode = &v
	return s
}

func (s *UnbindEslDeviceRequest) SetStoreId(v string) *UnbindEslDeviceRequest {
	s.StoreId = &v
	return s
}

type UnbindEslDeviceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindEslDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindEslDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindEslDeviceResponseBody) SetCode(v string) *UnbindEslDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindEslDeviceResponseBody) SetDynamicCode(v string) *UnbindEslDeviceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UnbindEslDeviceResponseBody) SetDynamicMessage(v string) *UnbindEslDeviceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UnbindEslDeviceResponseBody) SetErrorCode(v string) *UnbindEslDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UnbindEslDeviceResponseBody) SetErrorMessage(v string) *UnbindEslDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UnbindEslDeviceResponseBody) SetMessage(v string) *UnbindEslDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindEslDeviceResponseBody) SetRequestId(v string) *UnbindEslDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindEslDeviceResponseBody) SetSuccess(v bool) *UnbindEslDeviceResponseBody {
	s.Success = &v
	return s
}

type UnbindEslDeviceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindEslDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindEslDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindEslDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnbindEslDeviceResponse) SetHeaders(v map[string]*string) *UnbindEslDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnbindEslDeviceResponse) SetStatusCode(v int32) *UnbindEslDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindEslDeviceResponse) SetBody(v *UnbindEslDeviceResponseBody) *UnbindEslDeviceResponse {
	s.Body = v
	return s
}

type UnbindEslDeviceShelfRequest struct {
	EslBarCode *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s UnbindEslDeviceShelfRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindEslDeviceShelfRequest) GoString() string {
	return s.String()
}

func (s *UnbindEslDeviceShelfRequest) SetEslBarCode(v string) *UnbindEslDeviceShelfRequest {
	s.EslBarCode = &v
	return s
}

func (s *UnbindEslDeviceShelfRequest) SetStoreId(v string) *UnbindEslDeviceShelfRequest {
	s.StoreId = &v
	return s
}

type UnbindEslDeviceShelfResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindEslDeviceShelfResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindEslDeviceShelfResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindEslDeviceShelfResponseBody) SetCode(v string) *UnbindEslDeviceShelfResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindEslDeviceShelfResponseBody) SetDynamicCode(v string) *UnbindEslDeviceShelfResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UnbindEslDeviceShelfResponseBody) SetDynamicMessage(v string) *UnbindEslDeviceShelfResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UnbindEslDeviceShelfResponseBody) SetErrorCode(v string) *UnbindEslDeviceShelfResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UnbindEslDeviceShelfResponseBody) SetErrorMessage(v string) *UnbindEslDeviceShelfResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UnbindEslDeviceShelfResponseBody) SetMessage(v string) *UnbindEslDeviceShelfResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindEslDeviceShelfResponseBody) SetRequestId(v string) *UnbindEslDeviceShelfResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindEslDeviceShelfResponseBody) SetSuccess(v bool) *UnbindEslDeviceShelfResponseBody {
	s.Success = &v
	return s
}

type UnbindEslDeviceShelfResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindEslDeviceShelfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindEslDeviceShelfResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindEslDeviceShelfResponse) GoString() string {
	return s.String()
}

func (s *UnbindEslDeviceShelfResponse) SetHeaders(v map[string]*string) *UnbindEslDeviceShelfResponse {
	s.Headers = v
	return s
}

func (s *UnbindEslDeviceShelfResponse) SetStatusCode(v int32) *UnbindEslDeviceShelfResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindEslDeviceShelfResponse) SetBody(v *UnbindEslDeviceShelfResponseBody) *UnbindEslDeviceShelfResponse {
	s.Body = v
	return s
}

type UnmapPlanogramRailRequest struct {
	// This parameter is required.
	RailCode *string `json:"RailCode,omitempty" xml:"RailCode,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s UnmapPlanogramRailRequest) String() string {
	return tea.Prettify(s)
}

func (s UnmapPlanogramRailRequest) GoString() string {
	return s.String()
}

func (s *UnmapPlanogramRailRequest) SetRailCode(v string) *UnmapPlanogramRailRequest {
	s.RailCode = &v
	return s
}

func (s *UnmapPlanogramRailRequest) SetStoreId(v string) *UnmapPlanogramRailRequest {
	s.StoreId = &v
	return s
}

type UnmapPlanogramRailResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnmapPlanogramRailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnmapPlanogramRailResponseBody) GoString() string {
	return s.String()
}

func (s *UnmapPlanogramRailResponseBody) SetCode(v string) *UnmapPlanogramRailResponseBody {
	s.Code = &v
	return s
}

func (s *UnmapPlanogramRailResponseBody) SetDynamicCode(v string) *UnmapPlanogramRailResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UnmapPlanogramRailResponseBody) SetDynamicMessage(v string) *UnmapPlanogramRailResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UnmapPlanogramRailResponseBody) SetErrorCode(v string) *UnmapPlanogramRailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UnmapPlanogramRailResponseBody) SetErrorMessage(v string) *UnmapPlanogramRailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UnmapPlanogramRailResponseBody) SetMessage(v string) *UnmapPlanogramRailResponseBody {
	s.Message = &v
	return s
}

func (s *UnmapPlanogramRailResponseBody) SetRequestId(v string) *UnmapPlanogramRailResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnmapPlanogramRailResponseBody) SetSuccess(v bool) *UnmapPlanogramRailResponseBody {
	s.Success = &v
	return s
}

type UnmapPlanogramRailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnmapPlanogramRailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnmapPlanogramRailResponse) String() string {
	return tea.Prettify(s)
}

func (s UnmapPlanogramRailResponse) GoString() string {
	return s.String()
}

func (s *UnmapPlanogramRailResponse) SetHeaders(v map[string]*string) *UnmapPlanogramRailResponse {
	s.Headers = v
	return s
}

func (s *UnmapPlanogramRailResponse) SetStatusCode(v int32) *UnmapPlanogramRailResponse {
	s.StatusCode = &v
	return s
}

func (s *UnmapPlanogramRailResponse) SetBody(v *UnmapPlanogramRailResponseBody) *UnmapPlanogramRailResponse {
	s.Body = v
	return s
}

type UpdateEslDeviceLightRequest struct {
	EslBarCode *string `json:"EslBarCode,omitempty" xml:"EslBarCode,omitempty"`
	// This parameter is required.
	Frequency   *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	ItemBarCode *string `json:"ItemBarCode,omitempty" xml:"ItemBarCode,omitempty"`
	// This parameter is required.
	LedColor *string `json:"LedColor,omitempty" xml:"LedColor,omitempty"`
	// This parameter is required.
	LightUpTime *int32 `json:"LightUpTime,omitempty" xml:"LightUpTime,omitempty"`
	// This parameter is required.
	StoreId *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
}

func (s UpdateEslDeviceLightRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEslDeviceLightRequest) GoString() string {
	return s.String()
}

func (s *UpdateEslDeviceLightRequest) SetEslBarCode(v string) *UpdateEslDeviceLightRequest {
	s.EslBarCode = &v
	return s
}

func (s *UpdateEslDeviceLightRequest) SetFrequency(v string) *UpdateEslDeviceLightRequest {
	s.Frequency = &v
	return s
}

func (s *UpdateEslDeviceLightRequest) SetItemBarCode(v string) *UpdateEslDeviceLightRequest {
	s.ItemBarCode = &v
	return s
}

func (s *UpdateEslDeviceLightRequest) SetLedColor(v string) *UpdateEslDeviceLightRequest {
	s.LedColor = &v
	return s
}

func (s *UpdateEslDeviceLightRequest) SetLightUpTime(v int32) *UpdateEslDeviceLightRequest {
	s.LightUpTime = &v
	return s
}

func (s *UpdateEslDeviceLightRequest) SetStoreId(v string) *UpdateEslDeviceLightRequest {
	s.StoreId = &v
	return s
}

type UpdateEslDeviceLightResponseBody struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode     *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage  *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode       *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage    *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FailCount       *int64  `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	FailEslBarCodes *string `json:"FailEslBarCodes,omitempty" xml:"FailEslBarCodes,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	SuccessCount    *int64  `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s UpdateEslDeviceLightResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEslDeviceLightResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEslDeviceLightResponseBody) SetCode(v string) *UpdateEslDeviceLightResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetDynamicCode(v string) *UpdateEslDeviceLightResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetDynamicMessage(v string) *UpdateEslDeviceLightResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetErrorCode(v string) *UpdateEslDeviceLightResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetErrorMessage(v string) *UpdateEslDeviceLightResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetFailCount(v int64) *UpdateEslDeviceLightResponseBody {
	s.FailCount = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetFailEslBarCodes(v string) *UpdateEslDeviceLightResponseBody {
	s.FailEslBarCodes = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetMessage(v string) *UpdateEslDeviceLightResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetRequestId(v string) *UpdateEslDeviceLightResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetSuccess(v bool) *UpdateEslDeviceLightResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateEslDeviceLightResponseBody) SetSuccessCount(v int64) *UpdateEslDeviceLightResponseBody {
	s.SuccessCount = &v
	return s
}

type UpdateEslDeviceLightResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEslDeviceLightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEslDeviceLightResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEslDeviceLightResponse) GoString() string {
	return s.String()
}

func (s *UpdateEslDeviceLightResponse) SetHeaders(v map[string]*string) *UpdateEslDeviceLightResponse {
	s.Headers = v
	return s
}

func (s *UpdateEslDeviceLightResponse) SetStatusCode(v int32) *UpdateEslDeviceLightResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEslDeviceLightResponse) SetBody(v *UpdateEslDeviceLightResponseBody) *UpdateEslDeviceLightResponse {
	s.Body = v
	return s
}

type UpdateStoreRequest struct {
	Brand    *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	Groups   *string `json:"Groups,omitempty" xml:"Groups,omitempty"`
	OutId    *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	Phone    *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// This parameter is required.
	StoreId   *string `json:"StoreId,omitempty" xml:"StoreId,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
}

func (s UpdateStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoreRequest) GoString() string {
	return s.String()
}

func (s *UpdateStoreRequest) SetBrand(v string) *UpdateStoreRequest {
	s.Brand = &v
	return s
}

func (s *UpdateStoreRequest) SetComments(v string) *UpdateStoreRequest {
	s.Comments = &v
	return s
}

func (s *UpdateStoreRequest) SetGroups(v string) *UpdateStoreRequest {
	s.Groups = &v
	return s
}

func (s *UpdateStoreRequest) SetOutId(v string) *UpdateStoreRequest {
	s.OutId = &v
	return s
}

func (s *UpdateStoreRequest) SetPhone(v string) *UpdateStoreRequest {
	s.Phone = &v
	return s
}

func (s *UpdateStoreRequest) SetStoreId(v string) *UpdateStoreRequest {
	s.StoreId = &v
	return s
}

func (s *UpdateStoreRequest) SetStoreName(v string) *UpdateStoreRequest {
	s.StoreName = &v
	return s
}

type UpdateStoreResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoreResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStoreResponseBody) SetCode(v string) *UpdateStoreResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateStoreResponseBody) SetDynamicCode(v string) *UpdateStoreResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateStoreResponseBody) SetDynamicMessage(v string) *UpdateStoreResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateStoreResponseBody) SetErrorCode(v string) *UpdateStoreResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateStoreResponseBody) SetErrorMessage(v string) *UpdateStoreResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateStoreResponseBody) SetMessage(v string) *UpdateStoreResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateStoreResponseBody) SetRequestId(v string) *UpdateStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStoreResponseBody) SetSuccess(v bool) *UpdateStoreResponseBody {
	s.Success = &v
	return s
}

type UpdateStoreResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoreResponse) GoString() string {
	return s.String()
}

func (s *UpdateStoreResponse) SetHeaders(v map[string]*string) *UpdateStoreResponse {
	s.Headers = v
	return s
}

func (s *UpdateStoreResponse) SetStatusCode(v int32) *UpdateStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStoreResponse) SetBody(v *UpdateStoreResponseBody) *UpdateStoreResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("cloudesl.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("cloudesl.aliyuncs.com"),
		"ap-south-1":                  tea.String("cloudesl.aliyuncs.com"),
		"ap-southeast-1":              tea.String("cloudesl.aliyuncs.com"),
		"ap-southeast-2":              tea.String("cloudesl.aliyuncs.com"),
		"ap-southeast-3":              tea.String("cloudesl.aliyuncs.com"),
		"ap-southeast-5":              tea.String("cloudesl.aliyuncs.com"),
		"cn-beijing":                  tea.String("cloudesl.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("cloudesl.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("cloudesl.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("cloudesl.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("cloudesl.aliyuncs.com"),
		"cn-chengdu":                  tea.String("cloudesl.aliyuncs.com"),
		"cn-edge-1":                   tea.String("cloudesl.aliyuncs.com"),
		"cn-fujian":                   tea.String("cloudesl.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("cloudesl.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("cloudesl.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("cloudesl.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("cloudesl.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("cloudesl.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("cloudesl.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("cloudesl.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("cloudesl.aliyuncs.com"),
		"cn-hongkong":                 tea.String("cloudesl.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("cloudesl.aliyuncs.com"),
		"cn-huhehaote":                tea.String("cloudesl.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("cloudesl.aliyuncs.com"),
		"cn-qingdao":                  tea.String("cloudesl.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("cloudesl.aliyuncs.com"),
		"cn-shanghai":                 tea.String("cloudesl.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("cloudesl.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("cloudesl.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("cloudesl.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("cloudesl.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("cloudesl.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("cloudesl.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("cloudesl.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("cloudesl.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("cloudesl.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("cloudesl.aliyuncs.com"),
		"cn-wuhan":                    tea.String("cloudesl.aliyuncs.com"),
		"cn-yushanfang":               tea.String("cloudesl.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("cloudesl.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("cloudesl.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("cloudesl.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("cloudesl.aliyuncs.com"),
		"eu-central-1":                tea.String("cloudesl.aliyuncs.com"),
		"eu-west-1":                   tea.String("cloudesl.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("cloudesl.aliyuncs.com"),
		"me-east-1":                   tea.String("cloudesl.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("cloudesl.aliyuncs.com"),
		"us-east-1":                   tea.String("cloudesl.aliyuncs.com"),
		"us-west-1":                   tea.String("cloudesl.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudesl"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - ActivateApDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateApDeviceResponse
func (client *Client) ActivateApDeviceWithOptions(request *ActivateApDeviceRequest, runtime *util.RuntimeOptions) (_result *ActivateApDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApMac)) {
		body["ApMac"] = request.ApMac
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ActivateApDevice"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActivateApDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ActivateApDeviceRequest
//
// @return ActivateApDeviceResponse
func (client *Client) ActivateApDevice(request *ActivateApDeviceRequest) (_result *ActivateApDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActivateApDeviceResponse{}
	_body, _err := client.ActivateApDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddApDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddApDeviceResponse
func (client *Client) AddApDeviceWithOptions(request *AddApDeviceRequest, runtime *util.RuntimeOptions) (_result *AddApDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApMac)) {
		body["ApMac"] = request.ApMac
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddApDevice"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddApDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddApDeviceRequest
//
// @return AddApDeviceResponse
func (client *Client) AddApDevice(request *AddApDeviceRequest) (_result *AddApDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddApDeviceResponse{}
	_body, _err := client.AddApDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddEslDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddEslDeviceResponse
func (client *Client) AddEslDeviceWithOptions(request *AddEslDeviceRequest, runtime *util.RuntimeOptions) (_result *AddEslDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EslBarCode)) {
		body["EslBarCode"] = request.EslBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddEslDevice"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddEslDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddEslDeviceRequest
//
// @return AddEslDeviceResponse
func (client *Client) AddEslDevice(request *AddEslDeviceRequest) (_result *AddEslDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddEslDeviceResponse{}
	_body, _err := client.AddEslDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserResponse
func (client *Client) AddUserWithOptions(request *AddUserRequest, runtime *util.RuntimeOptions) (_result *AddUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUser"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddUserRequest
//
// @return AddUserResponse
func (client *Client) AddUser(request *AddUserRequest) (_result *AddUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserResponse{}
	_body, _err := client.AddUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AssignUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignUserResponse
func (client *Client) AssignUserWithOptions(request *AssignUserRequest, runtime *util.RuntimeOptions) (_result *AssignUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Stores)) {
		body["Stores"] = request.Stores
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		body["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssignUser"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssignUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AssignUserRequest
//
// @return AssignUserResponse
func (client *Client) AssignUser(request *AssignUserRequest) (_result *AssignUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssignUserResponse{}
	_body, _err := client.AssignUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BatchInsertItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchInsertItemsResponse
func (client *Client) BatchInsertItemsWithOptions(request *BatchInsertItemsRequest, runtime *util.RuntimeOptions) (_result *BatchInsertItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ItemInfo)) {
		body["ItemInfo"] = request.ItemInfo
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchInsertItems"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchInsertItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - BatchInsertItemsRequest
//
// @return BatchInsertItemsResponse
func (client *Client) BatchInsertItems(request *BatchInsertItemsRequest) (_result *BatchInsertItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchInsertItemsResponse{}
	_body, _err := client.BatchInsertItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BindEslDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindEslDeviceResponse
func (client *Client) BindEslDeviceWithOptions(request *BindEslDeviceRequest, runtime *util.RuntimeOptions) (_result *BindEslDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EslBarCode)) {
		body["EslBarCode"] = request.EslBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.ItemBarCode)) {
		body["ItemBarCode"] = request.ItemBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BindEslDevice"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindEslDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - BindEslDeviceRequest
//
// @return BindEslDeviceResponse
func (client *Client) BindEslDevice(request *BindEslDeviceRequest) (_result *BindEslDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindEslDeviceResponse{}
	_body, _err := client.BindEslDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BindEslDeviceShelfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindEslDeviceShelfResponse
func (client *Client) BindEslDeviceShelfWithOptions(request *BindEslDeviceShelfRequest, runtime *util.RuntimeOptions) (_result *BindEslDeviceShelfResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EslBarCode)) {
		body["EslBarCode"] = request.EslBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.ShelfCode)) {
		body["ShelfCode"] = request.ShelfCode
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BindEslDeviceShelf"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindEslDeviceShelfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - BindEslDeviceShelfRequest
//
// @return BindEslDeviceShelfResponse
func (client *Client) BindEslDeviceShelf(request *BindEslDeviceShelfRequest) (_result *BindEslDeviceShelfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindEslDeviceShelfResponse{}
	_body, _err := client.BindEslDeviceShelfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ConfirmLogisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmLogisticsResponse
func (client *Client) ConfirmLogisticsWithOptions(request *ConfirmLogisticsRequest, runtime *util.RuntimeOptions) (_result *ConfirmLogisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogisticsDocuments)) {
		body["LogisticsDocuments"] = request.LogisticsDocuments
	}

	if !tea.BoolValue(util.IsUnset(request.PoNumber)) {
		body["PoNumber"] = request.PoNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PrNumber)) {
		body["PrNumber"] = request.PrNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfirmLogistics"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfirmLogisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ConfirmLogisticsRequest
//
// @return ConfirmLogisticsResponse
func (client *Client) ConfirmLogistics(request *ConfirmLogisticsRequest) (_result *ConfirmLogisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfirmLogisticsResponse{}
	_body, _err := client.ConfirmLogisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateStoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStoreResponse
func (client *Client) CreateStoreWithOptions(request *CreateStoreRequest, runtime *util.RuntimeOptions) (_result *CreateStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Brand)) {
		body["Brand"] = request.Brand
	}

	if !tea.BoolValue(util.IsUnset(request.Comments)) {
		body["Comments"] = request.Comments
	}

	if !tea.BoolValue(util.IsUnset(request.CompanyId)) {
		body["CompanyId"] = request.CompanyId
	}

	if !tea.BoolValue(util.IsUnset(request.Groups)) {
		body["Groups"] = request.Groups
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		body["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		body["ParentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.StoreName)) {
		body["StoreName"] = request.StoreName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateStore"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateStoreRequest
//
// @return CreateStoreResponse
func (client *Client) CreateStore(request *CreateStoreRequest) (_result *CreateStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStoreResponse{}
	_body, _err := client.CreateStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteApDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApDeviceResponse
func (client *Client) DeleteApDeviceWithOptions(request *DeleteApDeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteApDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApMac)) {
		body["ApMac"] = request.ApMac
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApDevice"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteApDeviceRequest
//
// @return DeleteApDeviceResponse
func (client *Client) DeleteApDevice(request *DeleteApDeviceRequest) (_result *DeleteApDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApDeviceResponse{}
	_body, _err := client.DeleteApDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteEslDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEslDeviceResponse
func (client *Client) DeleteEslDeviceWithOptions(request *DeleteEslDeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteEslDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EslBarCode)) {
		body["EslBarCode"] = request.EslBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEslDevice"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEslDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteEslDeviceRequest
//
// @return DeleteEslDeviceResponse
func (client *Client) DeleteEslDevice(request *DeleteEslDeviceRequest) (_result *DeleteEslDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEslDeviceResponse{}
	_body, _err := client.DeleteEslDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteItemResponse
func (client *Client) DeleteItemWithOptions(request *DeleteItemRequest, runtime *util.RuntimeOptions) (_result *DeleteItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ItemBarCode)) {
		body["ItemBarCode"] = request.ItemBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteItem"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteItemRequest
//
// @return DeleteItemResponse
func (client *Client) DeleteItem(request *DeleteItemRequest) (_result *DeleteItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteItemResponse{}
	_body, _err := client.DeleteItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteItemBySkuIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteItemBySkuIdResponse
func (client *Client) DeleteItemBySkuIdWithOptions(request *DeleteItemBySkuIdRequest, runtime *util.RuntimeOptions) (_result *DeleteItemBySkuIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SkuId)) {
		body["SkuId"] = request.SkuId
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteItemBySkuId"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteItemBySkuIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteItemBySkuIdRequest
//
// @return DeleteItemBySkuIdResponse
func (client *Client) DeleteItemBySkuId(request *DeleteItemBySkuIdRequest) (_result *DeleteItemBySkuIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteItemBySkuIdResponse{}
	_body, _err := client.DeleteItemBySkuIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteStoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStoreResponse
func (client *Client) DeleteStoreWithOptions(request *DeleteStoreRequest, runtime *util.RuntimeOptions) (_result *DeleteStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteStore"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteStoreRequest
//
// @return DeleteStoreResponse
func (client *Client) DeleteStore(request *DeleteStoreRequest) (_result *DeleteStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteStoreResponse{}
	_body, _err := client.DeleteStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteUserRequest
//
// @return DeleteUserResponse
func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAlarmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlarmsResponse
func (client *Client) DescribeAlarmsWithOptions(request *DescribeAlarmsRequest, runtime *util.RuntimeOptions) (_result *DescribeAlarmsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmId)) {
		body["AlarmId"] = request.AlarmId
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmStatus)) {
		body["AlarmStatus"] = request.AlarmStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmType)) {
		body["AlarmType"] = request.AlarmType
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorType)) {
		body["ErrorType"] = request.ErrorType
	}

	if !tea.BoolValue(util.IsUnset(request.FromAlarmTime)) {
		body["FromAlarmTime"] = request.FromAlarmTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.ToAlarmTime)) {
		body["ToAlarmTime"] = request.ToAlarmTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAlarms"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAlarmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAlarmsRequest
//
// @return DescribeAlarmsResponse
func (client *Client) DescribeAlarms(request *DescribeAlarmsRequest) (_result *DescribeAlarmsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlarmsResponse{}
	_body, _err := client.DescribeAlarmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeApDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApDevicesResponse
func (client *Client) DescribeApDevicesWithOptions(request *DescribeApDevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeApDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Activated)) {
		body["Activated"] = request.Activated
	}

	if !tea.BoolValue(util.IsUnset(request.ApMac)) {
		body["ApMac"] = request.ApMac
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApDevices"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeApDevicesRequest
//
// @return DescribeApDevicesResponse
func (client *Client) DescribeApDevices(request *DescribeApDevicesRequest) (_result *DescribeApDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApDevicesResponse{}
	_body, _err := client.DescribeApDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeEslDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEslDevicesResponse
func (client *Client) DescribeEslDevicesWithOptions(request *DescribeEslDevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeEslDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeBind)) {
		body["BeBind"] = request.BeBind
	}

	if !tea.BoolValue(util.IsUnset(request.EslBarCode)) {
		body["EslBarCode"] = request.EslBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.EslStatus)) {
		body["EslStatus"] = request.EslStatus
	}

	if !tea.BoolValue(util.IsUnset(request.FromBatteryLevel)) {
		body["FromBatteryLevel"] = request.FromBatteryLevel
	}

	if !tea.BoolValue(util.IsUnset(request.ItemBarCode)) {
		body["ItemBarCode"] = request.ItemBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.Mac)) {
		body["Mac"] = request.Mac
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ShelfCode)) {
		body["ShelfCode"] = request.ShelfCode
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.ToBatteryLevel)) {
		body["ToBatteryLevel"] = request.ToBatteryLevel
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Vendor)) {
		body["Vendor"] = request.Vendor
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEslDevices"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEslDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeEslDevicesRequest
//
// @return DescribeEslDevicesResponse
func (client *Client) DescribeEslDevices(request *DescribeEslDevicesRequest) (_result *DescribeEslDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEslDevicesResponse{}
	_body, _err := client.DescribeEslDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeItemsResponse
func (client *Client) DescribeItemsWithOptions(request *DescribeItemsRequest, runtime *util.RuntimeOptions) (_result *DescribeItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BePromotion)) {
		body["BePromotion"] = request.BePromotion
	}

	if !tea.BoolValue(util.IsUnset(request.ItemBarCode)) {
		body["ItemBarCode"] = request.ItemBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.ItemId)) {
		body["ItemId"] = request.ItemId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemTitle)) {
		body["ItemTitle"] = request.ItemTitle
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SkuId)) {
		body["SkuId"] = request.SkuId
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeItems"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeItemsRequest
//
// @return DescribeItemsResponse
func (client *Client) DescribeItems(request *DescribeItemsRequest) (_result *DescribeItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeItemsResponse{}
	_body, _err := client.DescribeItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeLogisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogisticsResponse
func (client *Client) DescribeLogisticsWithOptions(request *DescribeLogisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeLogisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogistics"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeLogisticsRequest
//
// @return DescribeLogisticsResponse
func (client *Client) DescribeLogistics(request *DescribeLogisticsRequest) (_result *DescribeLogisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogisticsResponse{}
	_body, _err := client.DescribeLogisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribePayOrdersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePayOrdersResponse
func (client *Client) DescribePayOrdersWithOptions(request *DescribePayOrdersRequest, runtime *util.RuntimeOptions) (_result *DescribePayOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FromDate)) {
		body["FromDate"] = request.FromDate
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ToDate)) {
		body["ToDate"] = request.ToDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePayOrders"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePayOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribePayOrdersRequest
//
// @return DescribePayOrdersResponse
func (client *Client) DescribePayOrders(request *DescribePayOrdersRequest) (_result *DescribePayOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePayOrdersResponse{}
	_body, _err := client.DescribePayOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribePlanogramRailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlanogramRailsResponse
func (client *Client) DescribePlanogramRailsWithOptions(request *DescribePlanogramRailsRequest, runtime *util.RuntimeOptions) (_result *DescribePlanogramRailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Layer)) {
		body["Layer"] = request.Layer
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RailCode)) {
		body["RailCode"] = request.RailCode
	}

	if !tea.BoolValue(util.IsUnset(request.Shelf)) {
		body["Shelf"] = request.Shelf
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePlanogramRails"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePlanogramRailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribePlanogramRailsRequest
//
// @return DescribePlanogramRailsResponse
func (client *Client) DescribePlanogramRails(request *DescribePlanogramRailsRequest) (_result *DescribePlanogramRailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlanogramRailsResponse{}
	_body, _err := client.DescribePlanogramRailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeStoresRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStoresResponse
func (client *Client) DescribeStoresWithOptions(request *DescribeStoresRequest, runtime *util.RuntimeOptions) (_result *DescribeStoresResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Brand)) {
		body["Brand"] = request.Brand
	}

	if !tea.BoolValue(util.IsUnset(request.CompanyId)) {
		body["CompanyId"] = request.CompanyId
	}

	if !tea.BoolValue(util.IsUnset(request.FromDate)) {
		body["FromDate"] = request.FromDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.StoreName)) {
		body["StoreName"] = request.StoreName
	}

	if !tea.BoolValue(util.IsUnset(request.ToDate)) {
		body["ToDate"] = request.ToDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStores"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStoresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeStoresRequest
//
// @return DescribeStoresResponse
func (client *Client) DescribeStores(request *DescribeStoresRequest) (_result *DescribeStoresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStoresResponse{}
	_body, _err := client.DescribeStoresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeUserLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserLogResponse
func (client *Client) DescribeUserLogWithOptions(request *DescribeUserLogRequest, runtime *util.RuntimeOptions) (_result *DescribeUserLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EslBarCode)) {
		body["EslBarCode"] = request.EslBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.FromDate)) {
		body["FromDate"] = request.FromDate
	}

	if !tea.BoolValue(util.IsUnset(request.ItemBarCode)) {
		body["ItemBarCode"] = request.ItemBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.ItemId)) {
		body["ItemId"] = request.ItemId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemTitle)) {
		body["ItemTitle"] = request.ItemTitle
	}

	if !tea.BoolValue(util.IsUnset(request.OperateStatus)) {
		body["OperateStatus"] = request.OperateStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		body["OperateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		body["OperateUserId"] = request.OperateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		body["Reverse"] = request.Reverse
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.ToDate)) {
		body["ToDate"] = request.ToDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserLog"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeUserLogRequest
//
// @return DescribeUserLogResponse
func (client *Client) DescribeUserLog(request *DescribeUserLogRequest) (_result *DescribeUserLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserLogResponse{}
	_body, _err := client.DescribeUserLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsersResponse
func (client *Client) DescribeUsersWithOptions(request *DescribeUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		body["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUsers"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeUsersRequest
//
// @return DescribeUsersResponse
func (client *Client) DescribeUsers(request *DescribeUsersRequest) (_result *DescribeUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUsersResponse{}
	_body, _err := client.DescribeUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetCompanyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCompanyResponse
func (client *Client) GetCompanyWithOptions(runtime *util.RuntimeOptions) (_result *GetCompanyResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetCompany"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCompanyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return GetCompanyResponse
func (client *Client) GetCompany() (_result *GetCompanyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCompanyResponse{}
	_body, _err := client.GetCompanyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetUserRequest
//
// @return GetUserResponse
func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - MapPlanogramRailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MapPlanogramRailResponse
func (client *Client) MapPlanogramRailWithOptions(request *MapPlanogramRailRequest, runtime *util.RuntimeOptions) (_result *MapPlanogramRailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Layer)) {
		body["Layer"] = request.Layer
	}

	if !tea.BoolValue(util.IsUnset(request.RailCode)) {
		body["RailCode"] = request.RailCode
	}

	if !tea.BoolValue(util.IsUnset(request.Shelf)) {
		body["Shelf"] = request.Shelf
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MapPlanogramRail"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MapPlanogramRailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - MapPlanogramRailRequest
//
// @return MapPlanogramRailResponse
func (client *Client) MapPlanogramRail(request *MapPlanogramRailRequest) (_result *MapPlanogramRailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MapPlanogramRailResponse{}
	_body, _err := client.MapPlanogramRailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RefreshTaobaoItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshTaobaoItemResponse
func (client *Client) RefreshTaobaoItemWithOptions(request *RefreshTaobaoItemRequest, runtime *util.RuntimeOptions) (_result *RefreshTaobaoItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OuterId)) {
		body["OuterId"] = request.OuterId
	}

	if !tea.BoolValue(util.IsUnset(request.SkuId)) {
		body["SkuId"] = request.SkuId
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.TaobaoItemId)) {
		body["TaobaoItemId"] = request.TaobaoItemId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshTaobaoItem"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshTaobaoItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - RefreshTaobaoItemRequest
//
// @return RefreshTaobaoItemResponse
func (client *Client) RefreshTaobaoItem(request *RefreshTaobaoItemRequest) (_result *RefreshTaobaoItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshTaobaoItemResponse{}
	_body, _err := client.RefreshTaobaoItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UnassignUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnassignUserResponse
func (client *Client) UnassignUserWithOptions(request *UnassignUserRequest, runtime *util.RuntimeOptions) (_result *UnassignUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnassignUser"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnassignUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UnassignUserRequest
//
// @return UnassignUserResponse
func (client *Client) UnassignUser(request *UnassignUserRequest) (_result *UnassignUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnassignUserResponse{}
	_body, _err := client.UnassignUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UnbindEslDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindEslDeviceResponse
func (client *Client) UnbindEslDeviceWithOptions(request *UnbindEslDeviceRequest, runtime *util.RuntimeOptions) (_result *UnbindEslDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EslBarCode)) {
		body["EslBarCode"] = request.EslBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.ItemBarCode)) {
		body["ItemBarCode"] = request.ItemBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindEslDevice"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindEslDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UnbindEslDeviceRequest
//
// @return UnbindEslDeviceResponse
func (client *Client) UnbindEslDevice(request *UnbindEslDeviceRequest) (_result *UnbindEslDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindEslDeviceResponse{}
	_body, _err := client.UnbindEslDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UnbindEslDeviceShelfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindEslDeviceShelfResponse
func (client *Client) UnbindEslDeviceShelfWithOptions(request *UnbindEslDeviceShelfRequest, runtime *util.RuntimeOptions) (_result *UnbindEslDeviceShelfResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EslBarCode)) {
		body["EslBarCode"] = request.EslBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindEslDeviceShelf"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindEslDeviceShelfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UnbindEslDeviceShelfRequest
//
// @return UnbindEslDeviceShelfResponse
func (client *Client) UnbindEslDeviceShelf(request *UnbindEslDeviceShelfRequest) (_result *UnbindEslDeviceShelfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindEslDeviceShelfResponse{}
	_body, _err := client.UnbindEslDeviceShelfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UnmapPlanogramRailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnmapPlanogramRailResponse
func (client *Client) UnmapPlanogramRailWithOptions(request *UnmapPlanogramRailRequest, runtime *util.RuntimeOptions) (_result *UnmapPlanogramRailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RailCode)) {
		body["RailCode"] = request.RailCode
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnmapPlanogramRail"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnmapPlanogramRailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UnmapPlanogramRailRequest
//
// @return UnmapPlanogramRailResponse
func (client *Client) UnmapPlanogramRail(request *UnmapPlanogramRailRequest) (_result *UnmapPlanogramRailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnmapPlanogramRailResponse{}
	_body, _err := client.UnmapPlanogramRailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateEslDeviceLightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEslDeviceLightResponse
func (client *Client) UpdateEslDeviceLightWithOptions(request *UpdateEslDeviceLightRequest, runtime *util.RuntimeOptions) (_result *UpdateEslDeviceLightResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EslBarCode)) {
		body["EslBarCode"] = request.EslBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.Frequency)) {
		body["Frequency"] = request.Frequency
	}

	if !tea.BoolValue(util.IsUnset(request.ItemBarCode)) {
		body["ItemBarCode"] = request.ItemBarCode
	}

	if !tea.BoolValue(util.IsUnset(request.LedColor)) {
		body["LedColor"] = request.LedColor
	}

	if !tea.BoolValue(util.IsUnset(request.LightUpTime)) {
		body["LightUpTime"] = request.LightUpTime
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEslDeviceLight"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEslDeviceLightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateEslDeviceLightRequest
//
// @return UpdateEslDeviceLightResponse
func (client *Client) UpdateEslDeviceLight(request *UpdateEslDeviceLightRequest) (_result *UpdateEslDeviceLightResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEslDeviceLightResponse{}
	_body, _err := client.UpdateEslDeviceLightWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateStoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStoreResponse
func (client *Client) UpdateStoreWithOptions(request *UpdateStoreRequest, runtime *util.RuntimeOptions) (_result *UpdateStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Brand)) {
		body["Brand"] = request.Brand
	}

	if !tea.BoolValue(util.IsUnset(request.Comments)) {
		body["Comments"] = request.Comments
	}

	if !tea.BoolValue(util.IsUnset(request.Groups)) {
		body["Groups"] = request.Groups
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		body["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		body["StoreId"] = request.StoreId
	}

	if !tea.BoolValue(util.IsUnset(request.StoreName)) {
		body["StoreName"] = request.StoreName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateStore"),
		Version:     tea.String("2019-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateStoreRequest
//
// @return UpdateStoreResponse
func (client *Client) UpdateStore(request *UpdateStoreRequest) (_result *UpdateStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateStoreResponse{}
	_body, _err := client.UpdateStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
