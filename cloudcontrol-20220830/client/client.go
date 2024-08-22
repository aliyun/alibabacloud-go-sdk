// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CancelTaskResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CancelTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelTaskResponseBody) SetRequestId(v string) *CancelTaskResponseBody {
	s.RequestId = &v
	return s
}

type CancelTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelTaskResponse) SetHeaders(v map[string]*string) *CancelTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelTaskResponse) SetStatusCode(v int32) *CancelTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelTaskResponse) SetBody(v *CancelTaskResponseBody) *CancelTaskResponse {
	s.Body = v
	return s
}

type CreateResourceRequest struct {
	// example:
	//
	// {
	//
	//      "AccountName": "cctest",
	//
	//      "AccountPassword": "Aa1234****"
	//
	// }
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// 1e810dfe1468721d0664a49b9d9f74f4
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRequest) SetBody(v map[string]interface{}) *CreateResourceRequest {
	s.Body = v
	return s
}

func (s *CreateResourceRequest) SetClientToken(v string) *CreateResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateResourceRequest) SetRegionId(v string) *CreateResourceRequest {
	s.RegionId = &v
	return s
}

type CreateResourceResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// cctest
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// Instance/r-8vbf5abe31c9c4d4/Account/cctest
	ResourcePath *string `json:"resourcePath,omitempty" xml:"resourcePath,omitempty"`
	// example:
	//
	// task-433aead756057fff8189a7ce5****
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceResponseBody) SetRequestId(v string) *CreateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceResponseBody) SetResourceId(v string) *CreateResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *CreateResourceResponseBody) SetResourcePath(v string) *CreateResourceResponseBody {
	s.ResourcePath = &v
	return s
}

func (s *CreateResourceResponseBody) SetTaskId(v string) *CreateResourceResponseBody {
	s.TaskId = &v
	return s
}

type CreateResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceResponse) SetHeaders(v map[string]*string) *CreateResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceResponse) SetStatusCode(v int32) *CreateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceResponse) SetBody(v *CreateResourceResponseBody) *CreateResourceResponse {
	s.Body = v
	return s
}

type DeleteResourceRequest struct {
	// example:
	//
	// 1e810dfe1468721d0664a49b9d9f74f4
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DeleteResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceRequest) SetClientToken(v string) *DeleteResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteResourceRequest) SetRegionId(v string) *DeleteResourceRequest {
	s.RegionId = &v
	return s
}

type DeleteResourceResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// task-433aead756057fff8189a7ce5****
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DeleteResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponseBody) SetRequestId(v string) *DeleteResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceResponseBody) SetTaskId(v string) *DeleteResourceResponseBody {
	s.TaskId = &v
	return s
}

type DeleteResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponse) SetHeaders(v map[string]*string) *DeleteResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceResponse) SetStatusCode(v int32) *DeleteResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceResponse) SetBody(v *DeleteResourceResponseBody) *DeleteResourceResponse {
	s.Body = v
	return s
}

type GetPriceRequest struct {
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// {
	//
	//         "LoadBalancerName": "cc-test",
	//
	//         "LoadBalancerSpec": "slb.s3.small",
	//
	//         "InternetChargeType": "paybybandwidth",
	//
	//         "AddressType": "internet",
	//
	//         "PaymentType": "PayAsYouGo",
	//
	//         "Bandwidth": 6
	//
	//       }
	ResourceAttributes map[string]interface{} `json:"resourceAttributes,omitempty" xml:"resourceAttributes,omitempty"`
}

func (s GetPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPriceRequest) GoString() string {
	return s.String()
}

func (s *GetPriceRequest) SetRegionId(v string) *GetPriceRequest {
	s.RegionId = &v
	return s
}

func (s *GetPriceRequest) SetResourceAttributes(v map[string]interface{}) *GetPriceRequest {
	s.ResourceAttributes = v
	return s
}

type GetPriceShrinkRequest struct {
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// {
	//
	//         "LoadBalancerName": "cc-test",
	//
	//         "LoadBalancerSpec": "slb.s3.small",
	//
	//         "InternetChargeType": "paybybandwidth",
	//
	//         "AddressType": "internet",
	//
	//         "PaymentType": "PayAsYouGo",
	//
	//         "Bandwidth": 6
	//
	//       }
	ResourceAttributesShrink *string `json:"resourceAttributes,omitempty" xml:"resourceAttributes,omitempty"`
}

func (s GetPriceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetPriceShrinkRequest) SetRegionId(v string) *GetPriceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetPriceShrinkRequest) SetResourceAttributesShrink(v string) *GetPriceShrinkRequest {
	s.ResourceAttributesShrink = &v
	return s
}

type GetPriceResponseBody struct {
	Price *GetPriceResponseBodyPrice `json:"price,omitempty" xml:"price,omitempty" type:"Struct"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetPriceResponseBody) SetPrice(v *GetPriceResponseBodyPrice) *GetPriceResponseBody {
	s.Price = v
	return s
}

func (s *GetPriceResponseBody) SetRequestId(v string) *GetPriceResponseBody {
	s.RequestId = &v
	return s
}

type GetPriceResponseBodyPrice struct {
	// example:
	//
	// CNY
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// example:
	//
	// 0.0
	DiscountPrice *float32                                  `json:"discountPrice,omitempty" xml:"discountPrice,omitempty"`
	ModuleDetails []*GetPriceResponseBodyPriceModuleDetails `json:"moduleDetails,omitempty" xml:"moduleDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 760.0
	OriginalPrice    *float32                                     `json:"originalPrice,omitempty" xml:"originalPrice,omitempty"`
	PromotionDetails []*GetPriceResponseBodyPricePromotionDetails `json:"promotionDetails,omitempty" xml:"promotionDetails,omitempty" type:"Repeated"`
	// example:
	//
	// 0.0
	TradePrice *float32 `json:"tradePrice,omitempty" xml:"tradePrice,omitempty"`
}

func (s GetPriceResponseBodyPrice) String() string {
	return tea.Prettify(s)
}

func (s GetPriceResponseBodyPrice) GoString() string {
	return s.String()
}

func (s *GetPriceResponseBodyPrice) SetCurrency(v string) *GetPriceResponseBodyPrice {
	s.Currency = &v
	return s
}

func (s *GetPriceResponseBodyPrice) SetDiscountPrice(v float32) *GetPriceResponseBodyPrice {
	s.DiscountPrice = &v
	return s
}

func (s *GetPriceResponseBodyPrice) SetModuleDetails(v []*GetPriceResponseBodyPriceModuleDetails) *GetPriceResponseBodyPrice {
	s.ModuleDetails = v
	return s
}

func (s *GetPriceResponseBodyPrice) SetOriginalPrice(v float32) *GetPriceResponseBodyPrice {
	s.OriginalPrice = &v
	return s
}

func (s *GetPriceResponseBodyPrice) SetPromotionDetails(v []*GetPriceResponseBodyPricePromotionDetails) *GetPriceResponseBodyPrice {
	s.PromotionDetails = v
	return s
}

func (s *GetPriceResponseBodyPrice) SetTradePrice(v float32) *GetPriceResponseBodyPrice {
	s.TradePrice = &v
	return s
}

type GetPriceResponseBodyPriceModuleDetails struct {
	// example:
	//
	// 0.02
	CostAfterDiscount *float32 `json:"costAfterDiscount,omitempty" xml:"costAfterDiscount,omitempty"`
	// example:
	//
	// 0.0
	InvoiceDiscount *float32 `json:"invoiceDiscount,omitempty" xml:"invoiceDiscount,omitempty"`
	// example:
	//
	// InstanceRent
	ModuleCode *string `json:"moduleCode,omitempty" xml:"moduleCode,omitempty"`
	// example:
	//
	// InstanceRent
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// example:
	//
	// 1000.0
	OriginalCost *float32 `json:"originalCost,omitempty" xml:"originalCost,omitempty"`
	// example:
	//
	// 1.0
	PriceType *string `json:"priceType,omitempty" xml:"priceType,omitempty"`
}

func (s GetPriceResponseBodyPriceModuleDetails) String() string {
	return tea.Prettify(s)
}

func (s GetPriceResponseBodyPriceModuleDetails) GoString() string {
	return s.String()
}

func (s *GetPriceResponseBodyPriceModuleDetails) SetCostAfterDiscount(v float32) *GetPriceResponseBodyPriceModuleDetails {
	s.CostAfterDiscount = &v
	return s
}

func (s *GetPriceResponseBodyPriceModuleDetails) SetInvoiceDiscount(v float32) *GetPriceResponseBodyPriceModuleDetails {
	s.InvoiceDiscount = &v
	return s
}

func (s *GetPriceResponseBodyPriceModuleDetails) SetModuleCode(v string) *GetPriceResponseBodyPriceModuleDetails {
	s.ModuleCode = &v
	return s
}

func (s *GetPriceResponseBodyPriceModuleDetails) SetModuleName(v string) *GetPriceResponseBodyPriceModuleDetails {
	s.ModuleName = &v
	return s
}

func (s *GetPriceResponseBodyPriceModuleDetails) SetOriginalCost(v float32) *GetPriceResponseBodyPriceModuleDetails {
	s.OriginalCost = &v
	return s
}

func (s *GetPriceResponseBodyPriceModuleDetails) SetPriceType(v string) *GetPriceResponseBodyPriceModuleDetails {
	s.PriceType = &v
	return s
}

type GetPriceResponseBodyPricePromotionDetails struct {
	// example:
	//
	// 37284
	PromotionDesc *string `json:"promotionDesc,omitempty" xml:"promotionDesc,omitempty"`
	PromotionId   *int64  `json:"promotionId,omitempty" xml:"promotionId,omitempty"`
	PromotionName *string `json:"promotionName,omitempty" xml:"promotionName,omitempty"`
}

func (s GetPriceResponseBodyPricePromotionDetails) String() string {
	return tea.Prettify(s)
}

func (s GetPriceResponseBodyPricePromotionDetails) GoString() string {
	return s.String()
}

func (s *GetPriceResponseBodyPricePromotionDetails) SetPromotionDesc(v string) *GetPriceResponseBodyPricePromotionDetails {
	s.PromotionDesc = &v
	return s
}

func (s *GetPriceResponseBodyPricePromotionDetails) SetPromotionId(v int64) *GetPriceResponseBodyPricePromotionDetails {
	s.PromotionId = &v
	return s
}

func (s *GetPriceResponseBodyPricePromotionDetails) SetPromotionName(v string) *GetPriceResponseBodyPricePromotionDetails {
	s.PromotionName = &v
	return s
}

type GetPriceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPriceResponse) GoString() string {
	return s.String()
}

func (s *GetPriceResponse) SetHeaders(v map[string]*string) *GetPriceResponse {
	s.Headers = v
	return s
}

func (s *GetPriceResponse) SetStatusCode(v int32) *GetPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPriceResponse) SetBody(v *GetPriceResponseBody) *GetPriceResponse {
	s.Body = v
	return s
}

type GetResourceTypeHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// zh_CH
	XAcsAcceptLanguage *string `json:"x-acs-accept-language,omitempty" xml:"x-acs-accept-language,omitempty"`
}

func (s GetResourceTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeHeaders) GoString() string {
	return s.String()
}

func (s *GetResourceTypeHeaders) SetCommonHeaders(v map[string]*string) *GetResourceTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetResourceTypeHeaders) SetXAcsAcceptLanguage(v string) *GetResourceTypeHeaders {
	s.XAcsAcceptLanguage = &v
	return s
}

type GetResourceTypeResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId    *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResourceType *GetResourceTypeResponseBodyResourceType `json:"resourceType,omitempty" xml:"resourceType,omitempty" type:"Struct"`
}

func (s GetResourceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBody) SetRequestId(v string) *GetResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetResourceType(v *GetResourceTypeResponseBodyResourceType) *GetResourceTypeResponseBody {
	s.ResourceType = v
	return s
}

type GetResourceTypeResponseBodyResourceType struct {
	CreateOnlyProperties   []*string                                        `json:"createOnlyProperties,omitempty" xml:"createOnlyProperties,omitempty" type:"Repeated"`
	DeleteOnlyProperties   []*string                                        `json:"deleteOnlyProperties,omitempty" xml:"deleteOnlyProperties,omitempty" type:"Repeated"`
	FilterProperties       []*string                                        `json:"filterProperties,omitempty" xml:"filterProperties,omitempty" type:"Repeated"`
	GetOnlyProperties      []*string                                        `json:"getOnlyProperties,omitempty" xml:"getOnlyProperties,omitempty" type:"Repeated"`
	GetResponseProperties  []*string                                        `json:"getResponseProperties,omitempty" xml:"getResponseProperties,omitempty" type:"Repeated"`
	Handlers               *GetResourceTypeResponseBodyResourceTypeHandlers `json:"handlers,omitempty" xml:"handlers,omitempty" type:"Struct"`
	Info                   *GetResourceTypeResponseBodyResourceTypeInfo     `json:"info,omitempty" xml:"info,omitempty" type:"Struct"`
	ListOnlyProperties     []*string                                        `json:"listOnlyProperties,omitempty" xml:"listOnlyProperties,omitempty" type:"Repeated"`
	ListResponseProperties []*string                                        `json:"listResponseProperties,omitempty" xml:"listResponseProperties,omitempty" type:"Repeated"`
	// example:
	//
	// /properties/InstanceId
	PrimaryIdentifier *string `json:"primaryIdentifier,omitempty" xml:"primaryIdentifier,omitempty"`
	// example:
	//
	// ECS
	Product            *string                `json:"product,omitempty" xml:"product,omitempty"`
	Properties         map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	PublicProperties   []*string              `json:"publicProperties,omitempty" xml:"publicProperties,omitempty" type:"Repeated"`
	ReadOnlyProperties []*string              `json:"readOnlyProperties,omitempty" xml:"readOnlyProperties,omitempty" type:"Repeated"`
	Required           []*string              `json:"required,omitempty" xml:"required,omitempty" type:"Repeated"`
	// example:
	//
	// Instance
	ResourceType            *string   `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	SensitiveInfoProperties []*string `json:"sensitiveInfoProperties,omitempty" xml:"sensitiveInfoProperties,omitempty" type:"Repeated"`
	UpdateOnlyProperties    []*string `json:"updateOnlyProperties,omitempty" xml:"updateOnlyProperties,omitempty" type:"Repeated"`
	UpdateTypeProperties    []*string `json:"updateTypeProperties,omitempty" xml:"updateTypeProperties,omitempty" type:"Repeated"`
}

func (s GetResourceTypeResponseBodyResourceType) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceType) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceType) SetCreateOnlyProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.CreateOnlyProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetDeleteOnlyProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.DeleteOnlyProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetFilterProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.FilterProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetGetOnlyProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.GetOnlyProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetGetResponseProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.GetResponseProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetHandlers(v *GetResourceTypeResponseBodyResourceTypeHandlers) *GetResourceTypeResponseBodyResourceType {
	s.Handlers = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetInfo(v *GetResourceTypeResponseBodyResourceTypeInfo) *GetResourceTypeResponseBodyResourceType {
	s.Info = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetListOnlyProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.ListOnlyProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetListResponseProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.ListResponseProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetPrimaryIdentifier(v string) *GetResourceTypeResponseBodyResourceType {
	s.PrimaryIdentifier = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetProduct(v string) *GetResourceTypeResponseBodyResourceType {
	s.Product = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetProperties(v map[string]interface{}) *GetResourceTypeResponseBodyResourceType {
	s.Properties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetPublicProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.PublicProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetReadOnlyProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.ReadOnlyProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetRequired(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.Required = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetResourceType(v string) *GetResourceTypeResponseBodyResourceType {
	s.ResourceType = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetSensitiveInfoProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.SensitiveInfoProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetUpdateOnlyProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.UpdateOnlyProperties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetUpdateTypeProperties(v []*string) *GetResourceTypeResponseBodyResourceType {
	s.UpdateTypeProperties = v
	return s
}

type GetResourceTypeResponseBodyResourceTypeHandlers struct {
	Create *GetResourceTypeResponseBodyResourceTypeHandlersCreate `json:"create,omitempty" xml:"create,omitempty" type:"Struct"`
	Delete *GetResourceTypeResponseBodyResourceTypeHandlersDelete `json:"delete,omitempty" xml:"delete,omitempty" type:"Struct"`
	Get    *GetResourceTypeResponseBodyResourceTypeHandlersGet    `json:"get,omitempty" xml:"get,omitempty" type:"Struct"`
	List   *GetResourceTypeResponseBodyResourceTypeHandlersList   `json:"list,omitempty" xml:"list,omitempty" type:"Struct"`
	Update *GetResourceTypeResponseBodyResourceTypeHandlersUpdate `json:"update,omitempty" xml:"update,omitempty" type:"Struct"`
}

func (s GetResourceTypeResponseBodyResourceTypeHandlers) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceTypeHandlers) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlers) SetCreate(v *GetResourceTypeResponseBodyResourceTypeHandlersCreate) *GetResourceTypeResponseBodyResourceTypeHandlers {
	s.Create = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlers) SetDelete(v *GetResourceTypeResponseBodyResourceTypeHandlersDelete) *GetResourceTypeResponseBodyResourceTypeHandlers {
	s.Delete = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlers) SetGet(v *GetResourceTypeResponseBodyResourceTypeHandlersGet) *GetResourceTypeResponseBodyResourceTypeHandlers {
	s.Get = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlers) SetList(v *GetResourceTypeResponseBodyResourceTypeHandlersList) *GetResourceTypeResponseBodyResourceTypeHandlers {
	s.List = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlers) SetUpdate(v *GetResourceTypeResponseBodyResourceTypeHandlersUpdate) *GetResourceTypeResponseBodyResourceTypeHandlers {
	s.Update = v
	return s
}

type GetResourceTypeResponseBodyResourceTypeHandlersCreate struct {
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersCreate) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersCreate) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersCreate) SetPermissions(v []*string) *GetResourceTypeResponseBodyResourceTypeHandlersCreate {
	s.Permissions = v
	return s
}

type GetResourceTypeResponseBodyResourceTypeHandlersDelete struct {
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersDelete) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersDelete) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersDelete) SetPermissions(v []*string) *GetResourceTypeResponseBodyResourceTypeHandlersDelete {
	s.Permissions = v
	return s
}

type GetResourceTypeResponseBodyResourceTypeHandlersGet struct {
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersGet) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersGet) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersGet) SetPermissions(v []*string) *GetResourceTypeResponseBodyResourceTypeHandlersGet {
	s.Permissions = v
	return s
}

type GetResourceTypeResponseBodyResourceTypeHandlersList struct {
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersList) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersList) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersList) SetPermissions(v []*string) *GetResourceTypeResponseBodyResourceTypeHandlersList {
	s.Permissions = v
	return s
}

type GetResourceTypeResponseBodyResourceTypeHandlersUpdate struct {
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersUpdate) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceTypeHandlersUpdate) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceTypeHandlersUpdate) SetPermissions(v []*string) *GetResourceTypeResponseBodyResourceTypeHandlersUpdate {
	s.Permissions = v
	return s
}

type GetResourceTypeResponseBodyResourceTypeInfo struct {
	// example:
	//
	// paid
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// example:
	//
	// region
	DeliveryScope *string `json:"deliveryScope,omitempty" xml:"deliveryScope,omitempty"`
	// example:
	//
	// An ECS instance is equivalent to a virtual machine, including the most basic computing components such as CPU, memory, operating system, network, and disk. You can easily customize and change the configuration of the instance. You have full control over the virtual machine.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// Instance
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetResourceTypeResponseBodyResourceTypeInfo) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceTypeInfo) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceTypeInfo) SetChargeType(v string) *GetResourceTypeResponseBodyResourceTypeInfo {
	s.ChargeType = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeInfo) SetDeliveryScope(v string) *GetResourceTypeResponseBodyResourceTypeInfo {
	s.DeliveryScope = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeInfo) SetDescription(v string) *GetResourceTypeResponseBodyResourceTypeInfo {
	s.Description = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeInfo) SetTitle(v string) *GetResourceTypeResponseBodyResourceTypeInfo {
	s.Title = &v
	return s
}

type GetResourceTypeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponse) SetHeaders(v map[string]*string) *GetResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *GetResourceTypeResponse) SetStatusCode(v int32) *GetResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceTypeResponse) SetBody(v *GetResourceTypeResponseBody) *GetResourceTypeResponse {
	s.Body = v
	return s
}

type GetResourcesRequest struct {
	Filter map[string]interface{} `json:"filter,omitempty" xml:"filter,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourcesRequest) GoString() string {
	return s.String()
}

func (s *GetResourcesRequest) SetFilter(v map[string]interface{}) *GetResourcesRequest {
	s.Filter = v
	return s
}

func (s *GetResourcesRequest) SetMaxResults(v int32) *GetResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *GetResourcesRequest) SetNextToken(v string) *GetResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *GetResourcesRequest) SetRegionId(v string) *GetResourcesRequest {
	s.RegionId = &v
	return s
}

type GetResourcesShrinkRequest struct {
	FilterShrink *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetResourcesShrinkRequest) SetFilterShrink(v string) *GetResourcesShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *GetResourcesShrinkRequest) SetMaxResults(v int32) *GetResourcesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *GetResourcesShrinkRequest) SetNextToken(v string) *GetResourcesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *GetResourcesShrinkRequest) SetRegionId(v string) *GetResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

type GetResourcesResponseBody struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Resource  *GetResourcesResponseBodyResource    `json:"resource,omitempty" xml:"resource,omitempty" type:"Struct"`
	Resources []*GetResourcesResponseBodyResources `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourcesResponseBody) SetMaxResults(v int32) *GetResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetResourcesResponseBody) SetNextToken(v string) *GetResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetResourcesResponseBody) SetRequestId(v string) *GetResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourcesResponseBody) SetResource(v *GetResourcesResponseBodyResource) *GetResourcesResponseBody {
	s.Resource = v
	return s
}

func (s *GetResourcesResponseBody) SetResources(v []*GetResourcesResponseBodyResources) *GetResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *GetResourcesResponseBody) SetTotalCount(v int32) *GetResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type GetResourcesResponseBodyResource struct {
	// example:
	//
	// {"Status":"Available","Description":"","AccountPrivilege":"RoleReadWrite","InstanceId":"r-8vbf5abe31c9c4d4","RegionId":"cn-zhangjiakou","AccountType":"Normal","TypeInfo":{},"AccountName":"cctest"}
	ResourceAttributes map[string]interface{} `json:"resourceAttributes,omitempty" xml:"resourceAttributes,omitempty"`
	// example:
	//
	// cctest
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
}

func (s GetResourcesResponseBodyResource) String() string {
	return tea.Prettify(s)
}

func (s GetResourcesResponseBodyResource) GoString() string {
	return s.String()
}

func (s *GetResourcesResponseBodyResource) SetResourceAttributes(v map[string]interface{}) *GetResourcesResponseBodyResource {
	s.ResourceAttributes = v
	return s
}

func (s *GetResourcesResponseBodyResource) SetResourceId(v string) *GetResourcesResponseBodyResource {
	s.ResourceId = &v
	return s
}

type GetResourcesResponseBodyResources struct {
	// example:
	//
	// {"Status":"Available","Description":"","AccountPrivilege":"RoleReadWrite","InstanceId":"r-8vbf5abe31c9c4d4","RegionId":"cn-zhangjiakou","AccountType":"Normal","TypeInfo":{},"AccountName":"cctest"}
	ResourceAttributes map[string]interface{} `json:"resourceAttributes,omitempty" xml:"resourceAttributes,omitempty"`
	// example:
	//
	// cctest
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
}

func (s GetResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s GetResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *GetResourcesResponseBodyResources) SetResourceAttributes(v map[string]interface{}) *GetResourcesResponseBodyResources {
	s.ResourceAttributes = v
	return s
}

func (s *GetResourcesResponseBodyResources) SetResourceId(v string) *GetResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

type GetResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourcesResponse) GoString() string {
	return s.String()
}

func (s *GetResourcesResponse) SetHeaders(v map[string]*string) *GetResourcesResponse {
	s.Headers = v
	return s
}

func (s *GetResourcesResponse) SetStatusCode(v int32) *GetResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourcesResponse) SetBody(v *GetResourcesResponseBody) *GetResourcesResponse {
	s.Body = v
	return s
}

type GetTaskResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Task      *GetTaskResponseBodyTask `json:"task,omitempty" xml:"task,omitempty" type:"Struct"`
}

func (s GetTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetTask(v *GetTaskResponseBodyTask) *GetTaskResponseBody {
	s.Task = v
	return s
}

type GetTaskResponseBodyTask struct {
	// example:
	//
	// 2022-10-09T00:46:03Z
	CreateTime *string                       `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Error      *GetTaskResponseBodyTaskError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// example:
	//
	// ECS
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// i-8vbascjthm7kzhp3****
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// Instance/i-8vbascjthm7kzhp3****
	//
	// Instance/r-8vbf5abe31c9c4d4/Account/cctest
	ResourcePath *string `json:"resourcePath,omitempty" xml:"resourcePath,omitempty"`
	// example:
	//
	// Instance
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// Create
	TaskAction *string `json:"taskAction,omitempty" xml:"taskAction,omitempty"`
	// example:
	//
	// task-433aead756057fff8189a7ce5****
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetTaskResponseBodyTask) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTask) SetCreateTime(v string) *GetTaskResponseBodyTask {
	s.CreateTime = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetError(v *GetTaskResponseBodyTaskError) *GetTaskResponseBodyTask {
	s.Error = v
	return s
}

func (s *GetTaskResponseBodyTask) SetProduct(v string) *GetTaskResponseBodyTask {
	s.Product = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetRegionId(v string) *GetTaskResponseBodyTask {
	s.RegionId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetResourceId(v string) *GetTaskResponseBodyTask {
	s.ResourceId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetResourcePath(v string) *GetTaskResponseBodyTask {
	s.ResourcePath = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetResourceType(v string) *GetTaskResponseBodyTask {
	s.ResourceType = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetStatus(v string) *GetTaskResponseBodyTask {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTaskAction(v string) *GetTaskResponseBodyTask {
	s.TaskAction = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTaskId(v string) *GetTaskResponseBodyTask {
	s.TaskId = &v
	return s
}

type GetTaskResponseBodyTaskError struct {
	// example:
	//
	// OperationFailure.OperationFailed
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {
	//
	//      "requestId": "123****",
	//
	//      "errorCode": "InvalidRamUser.NoPermission",
	//
	//      "errorMsg": "Ram user is not authorized to perform the operation."
	//
	// }
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetTaskResponseBodyTaskError) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBodyTaskError) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskError) SetCode(v string) *GetTaskResponseBodyTaskError {
	s.Code = &v
	return s
}

func (s *GetTaskResponseBodyTaskError) SetMessage(v string) *GetTaskResponseBodyTaskError {
	s.Message = &v
	return s
}

type GetTaskResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTaskResponse) SetHeaders(v map[string]*string) *GetTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTaskResponse) SetStatusCode(v int32) *GetTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskResponse) SetBody(v *GetTaskResponseBody) *GetTaskResponse {
	s.Body = v
	return s
}

type ListDataSourcesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// RegionId
	AttributeName *string                `json:"attributeName,omitempty" xml:"attributeName,omitempty"`
	Filter        map[string]interface{} `json:"filter,omitempty" xml:"filter,omitempty"`
}

func (s ListDataSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesRequest) SetAttributeName(v string) *ListDataSourcesRequest {
	s.AttributeName = &v
	return s
}

func (s *ListDataSourcesRequest) SetFilter(v map[string]interface{}) *ListDataSourcesRequest {
	s.Filter = v
	return s
}

type ListDataSourcesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// RegionId
	AttributeName *string `json:"attributeName,omitempty" xml:"attributeName,omitempty"`
	FilterShrink  *string `json:"filter,omitempty" xml:"filter,omitempty"`
}

func (s ListDataSourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesShrinkRequest) SetAttributeName(v string) *ListDataSourcesShrinkRequest {
	s.AttributeName = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetFilterShrink(v string) *ListDataSourcesShrinkRequest {
	s.FilterShrink = &v
	return s
}

type ListDataSourcesResponseBody struct {
	DataSources []*ListDataSourcesResponseBodyDataSources `json:"dataSources,omitempty" xml:"dataSources,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListDataSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBody) SetDataSources(v []*ListDataSourcesResponseBodyDataSources) *ListDataSourcesResponseBody {
	s.DataSources = v
	return s
}

func (s *ListDataSourcesResponseBody) SetRequestId(v string) *ListDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListDataSourcesResponseBodyDataSources struct {
	// example:
	//
	// cn-beijing
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s ListDataSourcesResponseBodyDataSources) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBodyDataSources) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyDataSources) SetId(v string) *ListDataSourcesResponseBodyDataSources {
	s.Id = &v
	return s
}

type ListDataSourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponse) SetHeaders(v map[string]*string) *ListDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourcesResponse) SetStatusCode(v int32) *ListDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourcesResponse) SetBody(v *ListDataSourcesResponseBody) *ListDataSourcesResponse {
	s.Body = v
	return s
}

type ListProductsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// zh_CH
	XAcsAcceptLanguage *string `json:"x-acs-accept-language,omitempty" xml:"x-acs-accept-language,omitempty"`
}

func (s ListProductsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListProductsHeaders) GoString() string {
	return s.String()
}

func (s *ListProductsHeaders) SetCommonHeaders(v map[string]*string) *ListProductsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListProductsHeaders) SetXAcsAcceptLanguage(v string) *ListProductsHeaders {
	s.XAcsAcceptLanguage = &v
	return s
}

type ListProductsRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ECS
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListProductsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) SetMaxResults(v int32) *ListProductsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProductsRequest) SetNextToken(v string) *ListProductsRequest {
	s.NextToken = &v
	return s
}

type ListProductsResponseBody struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// Redis
	NextToken *string                             `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Products  []*ListProductsResponseBodyProducts `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) SetMaxResults(v int32) *ListProductsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProductsResponseBody) SetNextToken(v string) *ListProductsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductsResponseBody) SetProducts(v []*ListProductsResponseBodyProducts) *ListProductsResponseBody {
	s.Products = v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) SetTotalCount(v int32) *ListProductsResponseBody {
	s.TotalCount = &v
	return s
}

type ListProductsResponseBodyProducts struct {
	// example:
	//
	// ECS
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	// example:
	//
	// Elastic Compute Service
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
}

func (s ListProductsResponseBodyProducts) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyProducts) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProducts) SetProductCode(v string) *ListProductsResponseBodyProducts {
	s.ProductCode = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetProductName(v string) *ListProductsResponseBodyProducts {
	s.ProductName = &v
	return s
}

type ListProductsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponse) GoString() string {
	return s.String()
}

func (s *ListProductsResponse) SetHeaders(v map[string]*string) *ListProductsResponse {
	s.Headers = v
	return s
}

func (s *ListProductsResponse) SetStatusCode(v int32) *ListProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductsResponse) SetBody(v *ListProductsResponseBody) *ListProductsResponse {
	s.Body = v
	return s
}

type ListResourceTypesHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// zh_CH
	XAcsAcceptLanguage *string `json:"x-acs-accept-language,omitempty" xml:"x-acs-accept-language,omitempty"`
}

func (s ListResourceTypesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesHeaders) GoString() string {
	return s.String()
}

func (s *ListResourceTypesHeaders) SetCommonHeaders(v map[string]*string) *ListResourceTypesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListResourceTypesHeaders) SetXAcsAcceptLanguage(v string) *ListResourceTypesHeaders {
	s.XAcsAcceptLanguage = &v
	return s
}

type ListResourceTypesRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ECS::Disk
	NextToken     *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ResourceTypes []*string `json:"resourceTypes,omitempty" xml:"resourceTypes,omitempty" type:"Repeated"`
}

func (s ListResourceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypesRequest) SetMaxResults(v int32) *ListResourceTypesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceTypesRequest) SetNextToken(v string) *ListResourceTypesRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceTypesRequest) SetResourceTypes(v []*string) *ListResourceTypesRequest {
	s.ResourceTypes = v
	return s
}

type ListResourceTypesShrinkRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ECS::Disk
	NextToken           *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ResourceTypesShrink *string `json:"resourceTypes,omitempty" xml:"resourceTypes,omitempty"`
}

func (s ListResourceTypesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypesShrinkRequest) SetMaxResults(v int32) *ListResourceTypesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) SetNextToken(v string) *ListResourceTypesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) SetResourceTypesShrink(v string) *ListResourceTypesShrinkRequest {
	s.ResourceTypesShrink = &v
	return s
}

type ListResourceTypesResponseBody struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ECS::Disk
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId     *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResourceTypes []*ListResourceTypesResponseBodyResourceTypes `json:"resourceTypes,omitempty" xml:"resourceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListResourceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBody) SetMaxResults(v int32) *ListResourceTypesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetNextToken(v string) *ListResourceTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetRequestId(v string) *ListResourceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetResourceTypes(v []*ListResourceTypesResponseBodyResourceTypes) *ListResourceTypesResponseBody {
	s.ResourceTypes = v
	return s
}

func (s *ListResourceTypesResponseBody) SetTotalCount(v int32) *ListResourceTypesResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourceTypesResponseBodyResourceTypes struct {
	CreateOnlyProperties   []*string                                           `json:"createOnlyProperties,omitempty" xml:"createOnlyProperties,omitempty" type:"Repeated"`
	DeleteOnlyProperties   []*string                                           `json:"deleteOnlyProperties,omitempty" xml:"deleteOnlyProperties,omitempty" type:"Repeated"`
	FilterProperties       []*string                                           `json:"filterProperties,omitempty" xml:"filterProperties,omitempty" type:"Repeated"`
	GetOnlyProperties      []*string                                           `json:"getOnlyProperties,omitempty" xml:"getOnlyProperties,omitempty" type:"Repeated"`
	GetResponseProperties  []*string                                           `json:"getResponseProperties,omitempty" xml:"getResponseProperties,omitempty" type:"Repeated"`
	Handlers               *ListResourceTypesResponseBodyResourceTypesHandlers `json:"handlers,omitempty" xml:"handlers,omitempty" type:"Struct"`
	Info                   *ListResourceTypesResponseBodyResourceTypesInfo     `json:"info,omitempty" xml:"info,omitempty" type:"Struct"`
	ListOnlyProperties     []*string                                           `json:"listOnlyProperties,omitempty" xml:"listOnlyProperties,omitempty" type:"Repeated"`
	ListResponseProperties []*string                                           `json:"listResponseProperties,omitempty" xml:"listResponseProperties,omitempty" type:"Repeated"`
	// example:
	//
	// /properties/InstanceId
	PrimaryIdentifier *string `json:"primaryIdentifier,omitempty" xml:"primaryIdentifier,omitempty"`
	// example:
	//
	// ECS
	Product            *string                `json:"product,omitempty" xml:"product,omitempty"`
	Properties         map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	PublicProperties   []*string              `json:"publicProperties,omitempty" xml:"publicProperties,omitempty" type:"Repeated"`
	ReadOnlyProperties []*string              `json:"readOnlyProperties,omitempty" xml:"readOnlyProperties,omitempty" type:"Repeated"`
	Required           []*string              `json:"required,omitempty" xml:"required,omitempty" type:"Repeated"`
	// example:
	//
	// Instance
	ResourceType            *string   `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	SensitiveInfoProperties []*string `json:"sensitiveInfoProperties,omitempty" xml:"sensitiveInfoProperties,omitempty" type:"Repeated"`
	UpdateOnlyProperties    []*string `json:"updateOnlyProperties,omitempty" xml:"updateOnlyProperties,omitempty" type:"Repeated"`
	UpdateTypeProperties    []*string `json:"updateTypeProperties,omitempty" xml:"updateTypeProperties,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBodyResourceTypes) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypes) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetCreateOnlyProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.CreateOnlyProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetDeleteOnlyProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.DeleteOnlyProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetFilterProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.FilterProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetGetOnlyProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.GetOnlyProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetGetResponseProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.GetResponseProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetHandlers(v *ListResourceTypesResponseBodyResourceTypesHandlers) *ListResourceTypesResponseBodyResourceTypes {
	s.Handlers = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetInfo(v *ListResourceTypesResponseBodyResourceTypesInfo) *ListResourceTypesResponseBodyResourceTypes {
	s.Info = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetListOnlyProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.ListOnlyProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetListResponseProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.ListResponseProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetPrimaryIdentifier(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.PrimaryIdentifier = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetProduct(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.Product = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetProperties(v map[string]interface{}) *ListResourceTypesResponseBodyResourceTypes {
	s.Properties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetPublicProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.PublicProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetReadOnlyProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.ReadOnlyProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetRequired(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.Required = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetResourceType(v string) *ListResourceTypesResponseBodyResourceTypes {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetSensitiveInfoProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.SensitiveInfoProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetUpdateOnlyProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.UpdateOnlyProperties = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypes) SetUpdateTypeProperties(v []*string) *ListResourceTypesResponseBodyResourceTypes {
	s.UpdateTypeProperties = v
	return s
}

type ListResourceTypesResponseBodyResourceTypesHandlers struct {
	Create *ListResourceTypesResponseBodyResourceTypesHandlersCreate `json:"create,omitempty" xml:"create,omitempty" type:"Struct"`
	Delete *ListResourceTypesResponseBodyResourceTypesHandlersDelete `json:"delete,omitempty" xml:"delete,omitempty" type:"Struct"`
	Get    *ListResourceTypesResponseBodyResourceTypesHandlersGet    `json:"get,omitempty" xml:"get,omitempty" type:"Struct"`
	List   *ListResourceTypesResponseBodyResourceTypesHandlersList   `json:"list,omitempty" xml:"list,omitempty" type:"Struct"`
	Update *ListResourceTypesResponseBodyResourceTypesHandlersUpdate `json:"update,omitempty" xml:"update,omitempty" type:"Struct"`
}

func (s ListResourceTypesResponseBodyResourceTypesHandlers) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesHandlers) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlers) SetCreate(v *ListResourceTypesResponseBodyResourceTypesHandlersCreate) *ListResourceTypesResponseBodyResourceTypesHandlers {
	s.Create = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlers) SetDelete(v *ListResourceTypesResponseBodyResourceTypesHandlersDelete) *ListResourceTypesResponseBodyResourceTypesHandlers {
	s.Delete = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlers) SetGet(v *ListResourceTypesResponseBodyResourceTypesHandlersGet) *ListResourceTypesResponseBodyResourceTypesHandlers {
	s.Get = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlers) SetList(v *ListResourceTypesResponseBodyResourceTypesHandlersList) *ListResourceTypesResponseBodyResourceTypesHandlers {
	s.List = v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlers) SetUpdate(v *ListResourceTypesResponseBodyResourceTypesHandlersUpdate) *ListResourceTypesResponseBodyResourceTypesHandlers {
	s.Update = v
	return s
}

type ListResourceTypesResponseBodyResourceTypesHandlersCreate struct {
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersCreate) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersCreate) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersCreate) SetPermissions(v []*string) *ListResourceTypesResponseBodyResourceTypesHandlersCreate {
	s.Permissions = v
	return s
}

type ListResourceTypesResponseBodyResourceTypesHandlersDelete struct {
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersDelete) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersDelete) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersDelete) SetPermissions(v []*string) *ListResourceTypesResponseBodyResourceTypesHandlersDelete {
	s.Permissions = v
	return s
}

type ListResourceTypesResponseBodyResourceTypesHandlersGet struct {
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersGet) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersGet) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersGet) SetPermissions(v []*string) *ListResourceTypesResponseBodyResourceTypesHandlersGet {
	s.Permissions = v
	return s
}

type ListResourceTypesResponseBodyResourceTypesHandlersList struct {
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersList) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersList) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersList) SetPermissions(v []*string) *ListResourceTypesResponseBodyResourceTypesHandlersList {
	s.Permissions = v
	return s
}

type ListResourceTypesResponseBodyResourceTypesHandlersUpdate struct {
	Permissions []*string `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersUpdate) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesHandlersUpdate) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesHandlersUpdate) SetPermissions(v []*string) *ListResourceTypesResponseBodyResourceTypesHandlersUpdate {
	s.Permissions = v
	return s
}

type ListResourceTypesResponseBodyResourceTypesInfo struct {
	// example:
	//
	// paid
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// example:
	//
	// region
	DeliveryScope *string `json:"deliveryScope,omitempty" xml:"deliveryScope,omitempty"`
	// example:
	//
	// An ECS instance is equivalent to a virtual machine, including the most basic computing components such as CPU, memory, operating system, network, and disk. You can easily customize and change the configuration of the instance. You have full control over the virtual machine.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// Instance
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceTypesInfo) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypesInfo) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) SetChargeType(v string) *ListResourceTypesResponseBodyResourceTypesInfo {
	s.ChargeType = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) SetDeliveryScope(v string) *ListResourceTypesResponseBodyResourceTypesInfo {
	s.DeliveryScope = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) SetDescription(v string) *ListResourceTypesResponseBodyResourceTypesInfo {
	s.Description = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypesInfo) SetTitle(v string) *ListResourceTypesResponseBodyResourceTypesInfo {
	s.Title = &v
	return s
}

type ListResourceTypesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponse) SetHeaders(v map[string]*string) *ListResourceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceTypesResponse) SetStatusCode(v int32) *ListResourceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceTypesResponse) SetBody(v *ListResourceTypesResponseBody) *ListResourceTypesResponse {
	s.Body = v
	return s
}

type UpdateResourceRequest struct {
	// example:
	//
	// {
	//
	//      "AccountPassword": "4321****",
	//
	//      "Description": "cctest"
	//
	// }
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// 1e810dfe1468721d0664a49b9d9f74f4
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s UpdateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequest) SetBody(v map[string]interface{}) *UpdateResourceRequest {
	s.Body = v
	return s
}

func (s *UpdateResourceRequest) SetClientToken(v string) *UpdateResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateResourceRequest) SetRegionId(v string) *UpdateResourceRequest {
	s.RegionId = &v
	return s
}

type UpdateResourceResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// task-433aead756057fff8189a7ce5****
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponseBody) SetRequestId(v string) *UpdateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceResponseBody) SetTaskId(v string) *UpdateResourceResponseBody {
	s.TaskId = &v
	return s
}

type UpdateResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponse) SetHeaders(v map[string]*string) *UpdateResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceResponse) SetStatusCode(v int32) *UpdateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceResponse) SetBody(v *UpdateResourceResponseBody) *UpdateResourceResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudcontrol"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 取消任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelTaskResponse
func (client *Client) CancelTaskWithOptions(taskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CancelTask"),
		Version:     tea.String("2022-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId)) + "/operation/cancel"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消任务
//
// @return CancelTaskResponse
func (client *Client) CancelTask(taskId *string) (_result *CancelTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelTaskResponse{}
	_body, _err := client.CancelTaskWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建资源
//
// @param requestPath - the whole path of resource string
//
// @param request - CreateResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceResponse
func (client *Client) CreateResourceWithOptions(requestPath *string, request *CreateResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResource"),
		Version:     tea.String("2022-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String(tea.StringValue(requestPath)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资源
//
// @param requestPath - the whole path of resource string
//
// @param request - CreateResourceRequest
//
// @return CreateResourceResponse
func (client *Client) CreateResource(requestPath *string, request *CreateResourceRequest) (_result *CreateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceResponse{}
	_body, _err := client.CreateResourceWithOptions(requestPath, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除资源
//
// @param requestPath - the whole path of resource string
//
// @param request - DeleteResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceResponse
func (client *Client) DeleteResourceWithOptions(requestPath *string, request *DeleteResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResource"),
		Version:     tea.String("2022-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String(tea.StringValue(requestPath)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除资源
//
// @param requestPath - the whole path of resource string
//
// @param request - DeleteResourceRequest
//
// @return DeleteResourceResponse
func (client *Client) DeleteResource(requestPath *string, request *DeleteResourceRequest) (_result *DeleteResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceResponse{}
	_body, _err := client.DeleteResourceWithOptions(requestPath, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询价格
//
// @param requestPath - the whole path of resource string
//
// @param tmpReq - GetPriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPriceResponse
func (client *Client) GetPriceWithOptions(requestPath *string, tmpReq *GetPriceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPriceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetPriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceAttributes)) {
		request.ResourceAttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceAttributes, tea.String("resourceAttributes"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceAttributesShrink)) {
		query["resourceAttributes"] = request.ResourceAttributesShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPrice"),
		Version:     tea.String("2022-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String(tea.StringValue(requestPath)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询价格
//
// @param requestPath - the whole path of resource string
//
// @param request - GetPriceRequest
//
// @return GetPriceResponse
func (client *Client) GetPrice(requestPath *string, request *GetPriceRequest) (_result *GetPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPriceResponse{}
	_body, _err := client.GetPriceWithOptions(requestPath, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取资源元数据
//
// @param requestPath - the whole path of resource string
//
// @param headers - GetResourceTypeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceTypeResponse
func (client *Client) GetResourceTypeWithOptions(requestPath *string, headers *GetResourceTypeHeaders, runtime *util.RuntimeOptions) (_result *GetResourceTypeResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAcceptLanguage)) {
		realHeaders["x-acs-accept-language"] = util.ToJSONString(headers.XAcsAcceptLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceType"),
		Version:     tea.String("2022-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String(tea.StringValue(requestPath)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源元数据
//
// @param requestPath - the whole path of resource string
//
// @return GetResourceTypeResponse
func (client *Client) GetResourceType(requestPath *string) (_result *GetResourceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetResourceTypeHeaders{}
	_result = &GetResourceTypeResponse{}
	_body, _err := client.GetResourceTypeWithOptions(requestPath, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询资源
//
// @param requestPath - the whole path of resource string
//
// @param tmpReq - GetResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourcesResponse
func (client *Client) GetResourcesWithOptions(requestPath *string, tmpReq *GetResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Filter)) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, tea.String("filter"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterShrink)) {
		query["filter"] = request.FilterShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResources"),
		Version:     tea.String("2022-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String(tea.StringValue(requestPath)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询资源
//
// @param requestPath - the whole path of resource string
//
// @param request - GetResourcesRequest
//
// @return GetResourcesResponse
func (client *Client) GetResources(requestPath *string, request *GetResourcesRequest) (_result *GetResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourcesResponse{}
	_body, _err := client.GetResourcesWithOptions(requestPath, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询任务
//
// Description:
//
// GET /api/v1/tasks/{taskId}。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithOptions(taskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTask"),
		Version:     tea.String("2022-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务
//
// Description:
//
// GET /api/v1/tasks/{taskId}。
//
// @return GetTaskResponse
func (client *Client) GetTask(taskId *string) (_result *GetTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举资源属性可选值
//
// @param requestPath - the whole path of resource string
//
// @param tmpReq - ListDataSourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSourcesWithOptions(requestPath *string, tmpReq *ListDataSourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataSourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListDataSourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Filter)) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, tea.String("filter"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttributeName)) {
		query["attributeName"] = request.AttributeName
	}

	if !tea.BoolValue(util.IsUnset(request.FilterShrink)) {
		query["filter"] = request.FilterShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSources"),
		Version:     tea.String("2022-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String(tea.StringValue(requestPath)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举资源属性可选值
//
// @param requestPath - the whole path of resource string
//
// @param request - ListDataSourcesRequest
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSources(requestPath *string, request *ListDataSourcesRequest) (_result *ListDataSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSourcesResponse{}
	_body, _err := client.ListDataSourcesWithOptions(requestPath, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举资源类型
//
// Description:
//
// GET /api/v1/providers/{provider}/products。
//
// @param request - ListProductsRequest
//
// @param headers - ListProductsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductsResponse
func (client *Client) ListProductsWithOptions(provider *string, request *ListProductsRequest, headers *ListProductsHeaders, runtime *util.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAcceptLanguage)) {
		realHeaders["x-acs-accept-language"] = util.ToJSONString(headers.XAcsAcceptLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProducts"),
		Version:     tea.String("2022-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/providers/" + tea.StringValue(openapiutil.GetEncodeParam(provider)) + "/products"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举资源类型
//
// Description:
//
// GET /api/v1/providers/{provider}/products。
//
// @param request - ListProductsRequest
//
// @return ListProductsResponse
func (client *Client) ListProducts(provider *string, request *ListProductsRequest) (_result *ListProductsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListProductsHeaders{}
	_result = &ListProductsResponse{}
	_body, _err := client.ListProductsWithOptions(provider, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举资源类型
//
// Description:
//
// GET /api/v1/providers/{provider}/products/{product}/resourceTypes。
//
// @param tmpReq - ListResourceTypesRequest
//
// @param headers - ListResourceTypesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceTypesResponse
func (client *Client) ListResourceTypesWithOptions(provider *string, product *string, tmpReq *ListResourceTypesRequest, headers *ListResourceTypesHeaders, runtime *util.RuntimeOptions) (_result *ListResourceTypesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListResourceTypesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceTypes)) {
		request.ResourceTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceTypes, tea.String("resourceTypes"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypesShrink)) {
		query["resourceTypes"] = request.ResourceTypesShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAcceptLanguage)) {
		realHeaders["x-acs-accept-language"] = util.ToJSONString(headers.XAcsAcceptLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceTypes"),
		Version:     tea.String("2022-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/providers/" + tea.StringValue(openapiutil.GetEncodeParam(provider)) + "/products/" + tea.StringValue(openapiutil.GetEncodeParam(product)) + "/resourceTypes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举资源类型
//
// Description:
//
// GET /api/v1/providers/{provider}/products/{product}/resourceTypes。
//
// @param request - ListResourceTypesRequest
//
// @return ListResourceTypesResponse
func (client *Client) ListResourceTypes(provider *string, product *string, request *ListResourceTypesRequest) (_result *ListResourceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListResourceTypesHeaders{}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.ListResourceTypesWithOptions(provider, product, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新资源
//
// @param requestPath - the whole path of resource string
//
// @param request - UpdateResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceResponse
func (client *Client) UpdateResourceWithOptions(requestPath *string, request *UpdateResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResource"),
		Version:     tea.String("2022-08-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String(tea.StringValue(requestPath)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新资源
//
// @param requestPath - the whole path of resource string
//
// @param request - UpdateResourceRequest
//
// @return UpdateResourceResponse
func (client *Client) UpdateResource(requestPath *string, request *UpdateResourceRequest) (_result *UpdateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceResponse{}
	_body, _err := client.UpdateResourceWithOptions(requestPath, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
