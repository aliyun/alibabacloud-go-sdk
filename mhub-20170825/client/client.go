// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateAppRequest struct {
	// example:
	//
	// com.test.ios
	BundleId    *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	EncodedIcon *string `json:"EncodedIcon,omitempty" xml:"EncodedIcon,omitempty"`
	// example:
	//
	// 1
	IndustryId *string `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// com.test.android
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetBundleId(v string) *CreateAppRequest {
	s.BundleId = &v
	return s
}

func (s *CreateAppRequest) SetEncodedIcon(v string) *CreateAppRequest {
	s.EncodedIcon = &v
	return s
}

func (s *CreateAppRequest) SetIndustryId(v string) *CreateAppRequest {
	s.IndustryId = &v
	return s
}

func (s *CreateAppRequest) SetName(v string) *CreateAppRequest {
	s.Name = &v
	return s
}

func (s *CreateAppRequest) SetPackageName(v string) *CreateAppRequest {
	s.PackageName = &v
	return s
}

func (s *CreateAppRequest) SetProductId(v string) *CreateAppRequest {
	s.ProductId = &v
	return s
}

func (s *CreateAppRequest) SetType(v int32) *CreateAppRequest {
	s.Type = &v
	return s
}

type CreateAppResponseBody struct {
	AppInfo *CreateAppResponseBodyAppInfo `json:"AppInfo,omitempty" xml:"AppInfo,omitempty" type:"Struct"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetAppInfo(v *CreateAppResponseBodyAppInfo) *CreateAppResponseBody {
	s.AppInfo = v
	return s
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppResponseBodyAppInfo struct {
	// example:
	//
	// 123456
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// com.test.ios
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// example:
	//
	// 2020-12-16T06:25:52Z
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2020-12-16T06:25:52Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// com.test.android
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// example:
	//
	// 123456
	ProductId *int32 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAppResponseBodyAppInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyAppInfo) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyAppInfo) SetAppKey(v string) *CreateAppResponseBodyAppInfo {
	s.AppKey = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) SetBundleId(v string) *CreateAppResponseBodyAppInfo {
	s.BundleId = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) SetCreateTime(v string) *CreateAppResponseBodyAppInfo {
	s.CreateTime = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) SetDescription(v string) *CreateAppResponseBodyAppInfo {
	s.Description = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) SetModifyTime(v string) *CreateAppResponseBodyAppInfo {
	s.ModifyTime = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) SetName(v string) *CreateAppResponseBodyAppInfo {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) SetPackageName(v string) *CreateAppResponseBodyAppInfo {
	s.PackageName = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) SetProductId(v int32) *CreateAppResponseBodyAppInfo {
	s.ProductId = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) SetType(v int32) *CreateAppResponseBodyAppInfo {
	s.Type = &v
	return s
}

type CreateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetStatusCode(v int32) *CreateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateProductRequest struct {
	// example:
	//
	// AAA
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateProductRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProductRequest) GoString() string {
	return s.String()
}

func (s *CreateProductRequest) SetDescription(v string) *CreateProductRequest {
	s.Description = &v
	return s
}

func (s *CreateProductRequest) SetName(v string) *CreateProductRequest {
	s.Name = &v
	return s
}

type CreateProductResponseBody struct {
	// example:
	//
	// 123456
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProductResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductResponseBody) SetProductId(v int64) *CreateProductResponseBody {
	s.ProductId = &v
	return s
}

func (s *CreateProductResponseBody) SetRequestId(v string) *CreateProductResponseBody {
	s.RequestId = &v
	return s
}

type CreateProductResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProductResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProductResponse) GoString() string {
	return s.String()
}

func (s *CreateProductResponse) SetHeaders(v map[string]*string) *CreateProductResponse {
	s.Headers = v
	return s
}

func (s *CreateProductResponse) SetStatusCode(v int32) *CreateProductResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProductResponse) SetBody(v *CreateProductResponseBody) *CreateProductResponse {
	s.Body = v
	return s
}

type DeleteAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s DeleteAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRequest) SetAppKey(v string) *DeleteAppRequest {
	s.AppKey = &v
	return s
}

type DeleteAppResponseBody struct {
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppResponse) SetHeaders(v map[string]*string) *DeleteAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppResponse) SetStatusCode(v int32) *DeleteAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppResponse) SetBody(v *DeleteAppResponseBody) *DeleteAppResponse {
	s.Body = v
	return s
}

type DeleteProductRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s DeleteProductRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductRequest) GoString() string {
	return s.String()
}

func (s *DeleteProductRequest) SetProductId(v string) *DeleteProductRequest {
	s.ProductId = &v
	return s
}

type DeleteProductResponseBody struct {
	// example:
	//
	// PRODUCT_NOT_ALONE
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductResponseBody) SetMessage(v string) *DeleteProductResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteProductResponseBody) SetRequestId(v string) *DeleteProductResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProductResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProductResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductResponse) SetHeaders(v map[string]*string) *DeleteProductResponse {
	s.Headers = v
	return s
}

func (s *DeleteProductResponse) SetStatusCode(v int32) *DeleteProductResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProductResponse) SetBody(v *DeleteProductResponseBody) *DeleteProductResponse {
	s.Body = v
	return s
}

type DescribeDashboardRequest struct {
	// example:
	//
	// 29201799
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 1
	AppType *int32 `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// 4.8
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 1681985390256
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 4.8
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// queryAppVersions
	ProxyAction *string `json:"ProxyAction,omitempty" xml:"ProxyAction,omitempty"`
	// example:
	//
	// mqc
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// example:
	//
	// 1681369984564
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDashboardRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDashboardRequest) GoString() string {
	return s.String()
}

func (s *DescribeDashboardRequest) SetAppKey(v string) *DescribeDashboardRequest {
	s.AppKey = &v
	return s
}

func (s *DescribeDashboardRequest) SetAppType(v int32) *DescribeDashboardRequest {
	s.AppType = &v
	return s
}

func (s *DescribeDashboardRequest) SetAppVersion(v string) *DescribeDashboardRequest {
	s.AppVersion = &v
	return s
}

func (s *DescribeDashboardRequest) SetEndTime(v int64) *DescribeDashboardRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDashboardRequest) SetKeyword(v string) *DescribeDashboardRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDashboardRequest) SetProxyAction(v string) *DescribeDashboardRequest {
	s.ProxyAction = &v
	return s
}

func (s *DescribeDashboardRequest) SetServiceName(v string) *DescribeDashboardRequest {
	s.ServiceName = &v
	return s
}

func (s *DescribeDashboardRequest) SetStartTime(v int64) *DescribeDashboardRequest {
	s.StartTime = &v
	return s
}

type DescribeDashboardResponseBody struct {
	// example:
	//
	// {
	//
	// 	"success":true,
	//
	// 	"model":{
	//
	// 		"perfMonthCount":0,
	//
	// 		"compatibilityMonthCount":0,
	//
	// 		"perfLastMonthCount":0,
	//
	// 		"compatibilityLastMonthCount":0,
	//
	// 		"automationMonthCount":0,
	//
	// 		"automationLastMonthCount":0
	//
	// 	}
	//
	// }
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 4CC30A8F-787C-55CA-87A6-7D1BED56067E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDashboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDashboardResponseBody) SetModel(v string) *DescribeDashboardResponseBody {
	s.Model = &v
	return s
}

func (s *DescribeDashboardResponseBody) SetRequestId(v string) *DescribeDashboardResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDashboardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDashboardResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDashboardResponse) GoString() string {
	return s.String()
}

func (s *DescribeDashboardResponse) SetHeaders(v map[string]*string) *DescribeDashboardResponse {
	s.Headers = v
	return s
}

func (s *DescribeDashboardResponse) SetStatusCode(v int32) *DescribeDashboardResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDashboardResponse) SetBody(v *DescribeDashboardResponseBody) *DescribeDashboardResponse {
	s.Body = v
	return s
}

type ListAppsRequest struct {
	// example:
	//
	// 1
	OsType *int32 `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// example:
	//
	// 1
	Page *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s ListAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) SetOsType(v int32) *ListAppsRequest {
	s.OsType = &v
	return s
}

func (s *ListAppsRequest) SetPage(v string) *ListAppsRequest {
	s.Page = &v
	return s
}

func (s *ListAppsRequest) SetPageSize(v string) *ListAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppsRequest) SetProductId(v string) *ListAppsRequest {
	s.ProductId = &v
	return s
}

type ListAppsResponseBody struct {
	AppInfos *ListAppsResponseBodyAppInfos `json:"AppInfos,omitempty" xml:"AppInfos,omitempty" type:"Struct"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// enabled
	UbsmsStatus *string `json:"UbsmsStatus,omitempty" xml:"UbsmsStatus,omitempty"`
}

func (s ListAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBody) SetAppInfos(v *ListAppsResponseBodyAppInfos) *ListAppsResponseBody {
	s.AppInfos = v
	return s
}

func (s *ListAppsResponseBody) SetRequestId(v string) *ListAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppsResponseBody) SetTotal(v int32) *ListAppsResponseBody {
	s.Total = &v
	return s
}

func (s *ListAppsResponseBody) SetUbsmsStatus(v string) *ListAppsResponseBody {
	s.UbsmsStatus = &v
	return s
}

type ListAppsResponseBodyAppInfos struct {
	AppInfo []*ListAppsResponseBodyAppInfosAppInfo `json:"AppInfo,omitempty" xml:"AppInfo,omitempty" type:"Repeated"`
}

func (s ListAppsResponseBodyAppInfos) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyAppInfos) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyAppInfos) SetAppInfo(v []*ListAppsResponseBodyAppInfosAppInfo) *ListAppsResponseBodyAppInfos {
	s.AppInfo = v
	return s
}

type ListAppsResponseBodyAppInfosAppInfo struct {
	// example:
	//
	// 123456
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// com.test.ios
	BundleId    *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	EncodedIcon *string `json:"EncodedIcon,omitempty" xml:"EncodedIcon,omitempty"`
	// example:
	//
	// 1
	IndustryId *int32  `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// com.test.android
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// example:
	//
	// false
	Readonly *bool `json:"Readonly,omitempty" xml:"Readonly,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAppsResponseBodyAppInfosAppInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyAppInfosAppInfo) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetAppKey(v string) *ListAppsResponseBodyAppInfosAppInfo {
	s.AppKey = &v
	return s
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetBundleId(v string) *ListAppsResponseBodyAppInfosAppInfo {
	s.BundleId = &v
	return s
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetEncodedIcon(v string) *ListAppsResponseBodyAppInfosAppInfo {
	s.EncodedIcon = &v
	return s
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetIndustryId(v int32) *ListAppsResponseBodyAppInfosAppInfo {
	s.IndustryId = &v
	return s
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetName(v string) *ListAppsResponseBodyAppInfosAppInfo {
	s.Name = &v
	return s
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetPackageName(v string) *ListAppsResponseBodyAppInfosAppInfo {
	s.PackageName = &v
	return s
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetReadonly(v bool) *ListAppsResponseBodyAppInfosAppInfo {
	s.Readonly = &v
	return s
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetType(v int32) *ListAppsResponseBodyAppInfosAppInfo {
	s.Type = &v
	return s
}

type ListAppsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponse) GoString() string {
	return s.String()
}

func (s *ListAppsResponse) SetHeaders(v map[string]*string) *ListAppsResponse {
	s.Headers = v
	return s
}

func (s *ListAppsResponse) SetStatusCode(v int32) *ListAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppsResponse) SetBody(v *ListAppsResponseBody) *ListAppsResponse {
	s.Body = v
	return s
}

type ListProductsRequest struct {
	// example:
	//
	// 1
	Offset      *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// false
	Simple *bool `json:"Simple,omitempty" xml:"Simple,omitempty"`
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListProductsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) SetOffset(v int32) *ListProductsRequest {
	s.Offset = &v
	return s
}

func (s *ListProductsRequest) SetProductName(v string) *ListProductsRequest {
	s.ProductName = &v
	return s
}

func (s *ListProductsRequest) SetSimple(v bool) *ListProductsRequest {
	s.Simple = &v
	return s
}

func (s *ListProductsRequest) SetSize(v int32) *ListProductsRequest {
	s.Size = &v
	return s
}

type ListProductsResponseBody struct {
	ProductInfos *ListProductsResponseBodyProductInfos `json:"ProductInfos,omitempty" xml:"ProductInfos,omitempty" type:"Struct"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// enabled
	UbsmsStatus *string `json:"UbsmsStatus,omitempty" xml:"UbsmsStatus,omitempty"`
}

func (s ListProductsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) SetProductInfos(v *ListProductsResponseBodyProductInfos) *ListProductsResponseBody {
	s.ProductInfos = v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) SetTotal(v int32) *ListProductsResponseBody {
	s.Total = &v
	return s
}

func (s *ListProductsResponseBody) SetUbsmsStatus(v string) *ListProductsResponseBody {
	s.UbsmsStatus = &v
	return s
}

type ListProductsResponseBodyProductInfos struct {
	ProductInfo []*ListProductsResponseBodyProductInfosProductInfo `json:"ProductInfo,omitempty" xml:"ProductInfo,omitempty" type:"Repeated"`
}

func (s ListProductsResponseBodyProductInfos) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyProductInfos) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProductInfos) SetProductInfo(v []*ListProductsResponseBodyProductInfosProductInfo) *ListProductsResponseBodyProductInfos {
	s.ProductInfo = v
	return s
}

type ListProductsResponseBodyProductInfosProductInfo struct {
	// example:
	//
	// 2020-12-16T06:25:52Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// xxxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EncodedIcon *string `json:"EncodedIcon,omitempty" xml:"EncodedIcon,omitempty"`
	// example:
	//
	// 1
	IndustryId *int32  `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// iOS
	Platforms *string `json:"Platforms,omitempty" xml:"Platforms,omitempty"`
	// example:
	//
	// 1234
	ProductId *int32 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// false
	Readonly *bool `json:"Readonly,omitempty" xml:"Readonly,omitempty"`
}

func (s ListProductsResponseBodyProductInfosProductInfo) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyProductInfosProductInfo) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProductInfosProductInfo) SetCreateTime(v string) *ListProductsResponseBodyProductInfosProductInfo {
	s.CreateTime = &v
	return s
}

func (s *ListProductsResponseBodyProductInfosProductInfo) SetDescription(v string) *ListProductsResponseBodyProductInfosProductInfo {
	s.Description = &v
	return s
}

func (s *ListProductsResponseBodyProductInfosProductInfo) SetEncodedIcon(v string) *ListProductsResponseBodyProductInfosProductInfo {
	s.EncodedIcon = &v
	return s
}

func (s *ListProductsResponseBodyProductInfosProductInfo) SetIndustryId(v int32) *ListProductsResponseBodyProductInfosProductInfo {
	s.IndustryId = &v
	return s
}

func (s *ListProductsResponseBodyProductInfosProductInfo) SetName(v string) *ListProductsResponseBodyProductInfosProductInfo {
	s.Name = &v
	return s
}

func (s *ListProductsResponseBodyProductInfosProductInfo) SetPlatforms(v string) *ListProductsResponseBodyProductInfosProductInfo {
	s.Platforms = &v
	return s
}

func (s *ListProductsResponseBodyProductInfosProductInfo) SetProductId(v int32) *ListProductsResponseBodyProductInfosProductInfo {
	s.ProductId = &v
	return s
}

func (s *ListProductsResponseBodyProductInfosProductInfo) SetReadonly(v bool) *ListProductsResponseBodyProductInfosProductInfo {
	s.Readonly = &v
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

type ModifyAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// com.test.ios
	BundleId    *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	EncodedIcon *string `json:"EncodedIcon,omitempty" xml:"EncodedIcon,omitempty"`
	// example:
	//
	// 1
	IndustryId *int32  `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// com.test.android
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
}

func (s ModifyAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppRequest) SetAppKey(v string) *ModifyAppRequest {
	s.AppKey = &v
	return s
}

func (s *ModifyAppRequest) SetBundleId(v string) *ModifyAppRequest {
	s.BundleId = &v
	return s
}

func (s *ModifyAppRequest) SetEncodedIcon(v string) *ModifyAppRequest {
	s.EncodedIcon = &v
	return s
}

func (s *ModifyAppRequest) SetIndustryId(v int32) *ModifyAppRequest {
	s.IndustryId = &v
	return s
}

func (s *ModifyAppRequest) SetName(v string) *ModifyAppRequest {
	s.Name = &v
	return s
}

func (s *ModifyAppRequest) SetPackageName(v string) *ModifyAppRequest {
	s.PackageName = &v
	return s
}

type ModifyAppResponseBody struct {
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppResponseBody) SetRequestId(v string) *ModifyAppResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppResponse) SetHeaders(v map[string]*string) *ModifyAppResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppResponse) SetStatusCode(v int32) *ModifyAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppResponse) SetBody(v *ModifyAppResponseBody) *ModifyAppResponse {
	s.Body = v
	return s
}

type ModifyProductRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s ModifyProductRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyProductRequest) GoString() string {
	return s.String()
}

func (s *ModifyProductRequest) SetDescription(v string) *ModifyProductRequest {
	s.Description = &v
	return s
}

func (s *ModifyProductRequest) SetName(v string) *ModifyProductRequest {
	s.Name = &v
	return s
}

func (s *ModifyProductRequest) SetProductId(v string) *ModifyProductRequest {
	s.ProductId = &v
	return s
}

type ModifyProductResponseBody struct {
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyProductResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyProductResponseBody) SetRequestId(v string) *ModifyProductResponseBody {
	s.RequestId = &v
	return s
}

type ModifyProductResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyProductResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyProductResponse) GoString() string {
	return s.String()
}

func (s *ModifyProductResponse) SetHeaders(v map[string]*string) *ModifyProductResponse {
	s.Headers = v
	return s
}

func (s *ModifyProductResponse) SetStatusCode(v int32) *ModifyProductResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProductResponse) SetBody(v *ModifyProductResponseBody) *ModifyProductResponse {
	s.Body = v
	return s
}

type OpenEmasServiceRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s OpenEmasServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenEmasServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenEmasServiceRequest) SetOwnerId(v int64) *OpenEmasServiceRequest {
	s.OwnerId = &v
	return s
}

type OpenEmasServiceResponseBody struct {
	// example:
	//
	// 20671870151****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenEmasServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenEmasServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenEmasServiceResponseBody) SetOrderId(v string) *OpenEmasServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenEmasServiceResponseBody) SetRequestId(v string) *OpenEmasServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenEmasServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenEmasServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenEmasServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenEmasServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenEmasServiceResponse) SetHeaders(v map[string]*string) *OpenEmasServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenEmasServiceResponse) SetStatusCode(v int32) *OpenEmasServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenEmasServiceResponse) SetBody(v *OpenEmasServiceResponseBody) *OpenEmasServiceResponse {
	s.Body = v
	return s
}

type QueryAppInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s QueryAppInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAppInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryAppInfoRequest) SetAppKey(v string) *QueryAppInfoRequest {
	s.AppKey = &v
	return s
}

type QueryAppInfoResponseBody struct {
	AppInfo *QueryAppInfoResponseBodyAppInfo `json:"AppInfo,omitempty" xml:"AppInfo,omitempty" type:"Struct"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAppInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAppInfoResponseBody) SetAppInfo(v *QueryAppInfoResponseBodyAppInfo) *QueryAppInfoResponseBody {
	s.AppInfo = v
	return s
}

func (s *QueryAppInfoResponseBody) SetRequestId(v string) *QueryAppInfoResponseBody {
	s.RequestId = &v
	return s
}

type QueryAppInfoResponseBodyAppInfo struct {
	// example:
	//
	// 123456
	AppKey  *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// com.test.ios
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// example:
	//
	// false
	CertDevelopAvail *bool `json:"CertDevelopAvail,omitempty" xml:"CertDevelopAvail,omitempty"`
	// example:
	//
	// 2020-12-16T06:25:52Z
	CertDevelopExpiration *string `json:"CertDevelopExpiration,omitempty" xml:"CertDevelopExpiration,omitempty"`
	// example:
	//
	// false
	CertProductAvail *bool `json:"CertProductAvail,omitempty" xml:"CertProductAvail,omitempty"`
	// example:
	//
	// 2020-12-16T06:25:52Z
	CertProductExpiration *string `json:"CertProductExpiration,omitempty" xml:"CertProductExpiration,omitempty"`
	// example:
	//
	// 2020-12-16T06:25:52Z
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EncodedIcon *string `json:"EncodedIcon,omitempty" xml:"EncodedIcon,omitempty"`
	// example:
	//
	// 1
	IndustryId *int32 `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	// example:
	//
	// com.test.android
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// example:
	//
	// 1234
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// false
	Readonly *bool `json:"Readonly,omitempty" xml:"Readonly,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryAppInfoResponseBodyAppInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryAppInfoResponseBodyAppInfo) GoString() string {
	return s.String()
}

func (s *QueryAppInfoResponseBodyAppInfo) SetAppKey(v string) *QueryAppInfoResponseBodyAppInfo {
	s.AppKey = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetAppName(v string) *QueryAppInfoResponseBodyAppInfo {
	s.AppName = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetBundleId(v string) *QueryAppInfoResponseBodyAppInfo {
	s.BundleId = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetCertDevelopAvail(v bool) *QueryAppInfoResponseBodyAppInfo {
	s.CertDevelopAvail = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetCertDevelopExpiration(v string) *QueryAppInfoResponseBodyAppInfo {
	s.CertDevelopExpiration = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetCertProductAvail(v bool) *QueryAppInfoResponseBodyAppInfo {
	s.CertProductAvail = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetCertProductExpiration(v string) *QueryAppInfoResponseBodyAppInfo {
	s.CertProductExpiration = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetCreateTime(v string) *QueryAppInfoResponseBodyAppInfo {
	s.CreateTime = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetEncodedIcon(v string) *QueryAppInfoResponseBodyAppInfo {
	s.EncodedIcon = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetIndustryId(v int32) *QueryAppInfoResponseBodyAppInfo {
	s.IndustryId = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetPackageName(v string) *QueryAppInfoResponseBodyAppInfo {
	s.PackageName = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetProductId(v int64) *QueryAppInfoResponseBodyAppInfo {
	s.ProductId = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetReadonly(v bool) *QueryAppInfoResponseBodyAppInfo {
	s.Readonly = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetStatus(v int32) *QueryAppInfoResponseBodyAppInfo {
	s.Status = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetType(v int32) *QueryAppInfoResponseBodyAppInfo {
	s.Type = &v
	return s
}

type QueryAppInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAppInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAppInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryAppInfoResponse) SetHeaders(v map[string]*string) *QueryAppInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryAppInfoResponse) SetStatusCode(v int32) *QueryAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAppInfoResponse) SetBody(v *QueryAppInfoResponseBody) *QueryAppInfoResponse {
	s.Body = v
	return s
}

type QueryAppSecurityInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s QueryAppSecurityInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAppSecurityInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryAppSecurityInfoRequest) SetAppKey(v string) *QueryAppSecurityInfoRequest {
	s.AppKey = &v
	return s
}

type QueryAppSecurityInfoResponseBody struct {
	AppSecurityInfo *QueryAppSecurityInfoResponseBodyAppSecurityInfo `json:"AppSecurityInfo,omitempty" xml:"AppSecurityInfo,omitempty" type:"Struct"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAppSecurityInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAppSecurityInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAppSecurityInfoResponseBody) SetAppSecurityInfo(v *QueryAppSecurityInfoResponseBodyAppSecurityInfo) *QueryAppSecurityInfoResponseBody {
	s.AppSecurityInfo = v
	return s
}

func (s *QueryAppSecurityInfoResponseBody) SetRequestId(v string) *QueryAppSecurityInfoResponseBody {
	s.RequestId = &v
	return s
}

type QueryAppSecurityInfoResponseBodyAppSecurityInfo struct {
	// example:
	//
	// 123456
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// abc123abc123
	AppSecret *string                                                   `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	ExtConfig *QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig `json:"ExtConfig,omitempty" xml:"ExtConfig,omitempty" type:"Struct"`
}

func (s QueryAppSecurityInfoResponseBodyAppSecurityInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryAppSecurityInfoResponseBodyAppSecurityInfo) GoString() string {
	return s.String()
}

func (s *QueryAppSecurityInfoResponseBodyAppSecurityInfo) SetAppKey(v string) *QueryAppSecurityInfoResponseBodyAppSecurityInfo {
	s.AppKey = &v
	return s
}

func (s *QueryAppSecurityInfoResponseBodyAppSecurityInfo) SetAppSecret(v string) *QueryAppSecurityInfoResponseBodyAppSecurityInfo {
	s.AppSecret = &v
	return s
}

func (s *QueryAppSecurityInfoResponseBodyAppSecurityInfo) SetExtConfig(v *QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig) *QueryAppSecurityInfoResponseBodyAppSecurityInfo {
	s.ExtConfig = v
	return s
}

type QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig struct {
	TlogRsaSecret *string `json:"TlogRsaSecret,omitempty" xml:"TlogRsaSecret,omitempty"`
}

func (s QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig) String() string {
	return tea.Prettify(s)
}

func (s QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig) GoString() string {
	return s.String()
}

func (s *QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig) SetTlogRsaSecret(v string) *QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig {
	s.TlogRsaSecret = &v
	return s
}

type QueryAppSecurityInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAppSecurityInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAppSecurityInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAppSecurityInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryAppSecurityInfoResponse) SetHeaders(v map[string]*string) *QueryAppSecurityInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryAppSecurityInfoResponse) SetStatusCode(v int32) *QueryAppSecurityInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAppSecurityInfoResponse) SetBody(v *QueryAppSecurityInfoResponseBody) *QueryAppSecurityInfoResponse {
	s.Body = v
	return s
}

type QueryProductInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s QueryProductInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProductInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryProductInfoRequest) SetProductId(v string) *QueryProductInfoRequest {
	s.ProductId = &v
	return s
}

type QueryProductInfoResponseBody struct {
	ProductInfo *QueryProductInfoResponseBodyProductInfo `json:"ProductInfo,omitempty" xml:"ProductInfo,omitempty" type:"Struct"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryProductInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProductInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProductInfoResponseBody) SetProductInfo(v *QueryProductInfoResponseBodyProductInfo) *QueryProductInfoResponseBody {
	s.ProductInfo = v
	return s
}

func (s *QueryProductInfoResponseBody) SetRequestId(v string) *QueryProductInfoResponseBody {
	s.RequestId = &v
	return s
}

type QueryProductInfoResponseBodyProductInfo struct {
	EncodedIcon *string `json:"EncodedIcon,omitempty" xml:"EncodedIcon,omitempty"`
	IconName    *string `json:"IconName,omitempty" xml:"IconName,omitempty"`
	// example:
	//
	// 1
	IndustryId *int32  `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	Readonly *bool `json:"Readonly,omitempty" xml:"Readonly,omitempty"`
}

func (s QueryProductInfoResponseBodyProductInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryProductInfoResponseBodyProductInfo) GoString() string {
	return s.String()
}

func (s *QueryProductInfoResponseBodyProductInfo) SetEncodedIcon(v string) *QueryProductInfoResponseBodyProductInfo {
	s.EncodedIcon = &v
	return s
}

func (s *QueryProductInfoResponseBodyProductInfo) SetIconName(v string) *QueryProductInfoResponseBodyProductInfo {
	s.IconName = &v
	return s
}

func (s *QueryProductInfoResponseBodyProductInfo) SetIndustryId(v int32) *QueryProductInfoResponseBodyProductInfo {
	s.IndustryId = &v
	return s
}

func (s *QueryProductInfoResponseBodyProductInfo) SetName(v string) *QueryProductInfoResponseBodyProductInfo {
	s.Name = &v
	return s
}

func (s *QueryProductInfoResponseBodyProductInfo) SetReadonly(v bool) *QueryProductInfoResponseBodyProductInfo {
	s.Readonly = &v
	return s
}

type QueryProductInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProductInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProductInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProductInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryProductInfoResponse) SetHeaders(v map[string]*string) *QueryProductInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryProductInfoResponse) SetStatusCode(v int32) *QueryProductInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProductInfoResponse) SetBody(v *QueryProductInfoResponseBody) *QueryProductInfoResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("mhub"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - CreateAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppResponse
func (client *Client) CreateAppWithOptions(request *CreateAppRequest, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		query["BundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.EncodedIcon)) {
		query["EncodedIcon"] = request.EncodedIcon
	}

	if !tea.BoolValue(util.IsUnset(request.IndustryId)) {
		query["IndustryId"] = request.IndustryId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PackageName)) {
		query["PackageName"] = request.PackageName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApp"),
		Version:     tea.String("2017-08-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAppResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAppResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CreateAppRequest
//
// @return CreateAppResponse
func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProductResponse
func (client *Client) CreateProductWithOptions(request *CreateProductRequest, runtime *util.RuntimeOptions) (_result *CreateProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProduct"),
		Version:     tea.String("2017-08-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateProductResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateProductResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CreateProductRequest
//
// @return CreateProductResponse
func (client *Client) CreateProduct(request *CreateProductRequest) (_result *CreateProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProductResponse{}
	_body, _err := client.CreateProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppResponse
func (client *Client) DeleteAppWithOptions(request *DeleteAppRequest, runtime *util.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApp"),
		Version:     tea.String("2017-08-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAppResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAppResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DeleteAppRequest
//
// @return DeleteAppResponse
func (client *Client) DeleteApp(request *DeleteAppRequest) (_result *DeleteAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProductResponse
func (client *Client) DeleteProductWithOptions(request *DeleteProductRequest, runtime *util.RuntimeOptions) (_result *DeleteProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProduct"),
		Version:     tea.String("2017-08-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteProductResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteProductResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DeleteProductRequest
//
// @return DeleteProductResponse
func (client *Client) DeleteProduct(request *DeleteProductRequest) (_result *DeleteProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProductResponse{}
	_body, _err := client.DeleteProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取emas dashboard
//
// @param request - DescribeDashboardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDashboardResponse
func (client *Client) DescribeDashboardWithOptions(request *DescribeDashboardRequest, runtime *util.RuntimeOptions) (_result *DescribeDashboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyAction)) {
		query["ProxyAction"] = request.ProxyAction
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDashboard"),
		Version:     tea.String("2017-08-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeDashboardResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeDashboardResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取emas dashboard
//
// @param request - DescribeDashboardRequest
//
// @return DescribeDashboardResponse
func (client *Client) DescribeDashboard(request *DescribeDashboardRequest) (_result *DescribeDashboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDashboardResponse{}
	_body, _err := client.DescribeDashboardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 展示用户应用列表
//
// @param request - ListAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppsResponse
func (client *Client) ListAppsWithOptions(request *ListAppsRequest, runtime *util.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApps"),
		Version:     tea.String("2017-08-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAppsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAppsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 展示用户应用列表
//
// @param request - ListAppsRequest
//
// @return ListAppsResponse
func (client *Client) ListApps(request *ListAppsRequest) (_result *ListAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppsResponse{}
	_body, _err := client.ListAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductsResponse
func (client *Client) ListProductsWithOptions(request *ListProductsRequest, runtime *util.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		query["ProductName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.Simple)) {
		query["Simple"] = request.Simple
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProducts"),
		Version:     tea.String("2017-08-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListProductsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListProductsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListProductsRequest
//
// @return ListProductsResponse
func (client *Client) ListProducts(request *ListProductsRequest) (_result *ListProductsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProductsResponse{}
	_body, _err := client.ListProductsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppResponse
func (client *Client) ModifyAppWithOptions(request *ModifyAppRequest, runtime *util.RuntimeOptions) (_result *ModifyAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		query["BundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.EncodedIcon)) {
		query["EncodedIcon"] = request.EncodedIcon
	}

	if !tea.BoolValue(util.IsUnset(request.IndustryId)) {
		query["IndustryId"] = request.IndustryId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PackageName)) {
		query["PackageName"] = request.PackageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApp"),
		Version:     tea.String("2017-08-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyAppResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyAppResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyAppRequest
//
// @return ModifyAppResponse
func (client *Client) ModifyApp(request *ModifyAppRequest) (_result *ModifyAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAppResponse{}
	_body, _err := client.ModifyAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyProductResponse
func (client *Client) ModifyProductWithOptions(request *ModifyProductRequest, runtime *util.RuntimeOptions) (_result *ModifyProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyProduct"),
		Version:     tea.String("2017-08-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyProductResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyProductResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyProductRequest
//
// @return ModifyProductResponse
func (client *Client) ModifyProduct(request *ModifyProductRequest) (_result *ModifyProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyProductResponse{}
	_body, _err := client.ModifyProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - OpenEmasServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenEmasServiceResponse
func (client *Client) OpenEmasServiceWithOptions(request *OpenEmasServiceRequest, runtime *util.RuntimeOptions) (_result *OpenEmasServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenEmasService"),
		Version:     tea.String("2017-08-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OpenEmasServiceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OpenEmasServiceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - OpenEmasServiceRequest
//
// @return OpenEmasServiceResponse
func (client *Client) OpenEmasService(request *OpenEmasServiceRequest) (_result *OpenEmasServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenEmasServiceResponse{}
	_body, _err := client.OpenEmasServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAppInfoResponse
func (client *Client) QueryAppInfoWithOptions(request *QueryAppInfoRequest, runtime *util.RuntimeOptions) (_result *QueryAppInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAppInfo"),
		Version:     tea.String("2017-08-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &QueryAppInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &QueryAppInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - QueryAppInfoRequest
//
// @return QueryAppInfoResponse
func (client *Client) QueryAppInfo(request *QueryAppInfoRequest) (_result *QueryAppInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAppInfoResponse{}
	_body, _err := client.QueryAppInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询应用对应的安全字段
//
// @param request - QueryAppSecurityInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAppSecurityInfoResponse
func (client *Client) QueryAppSecurityInfoWithOptions(request *QueryAppSecurityInfoRequest, runtime *util.RuntimeOptions) (_result *QueryAppSecurityInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAppSecurityInfo"),
		Version:     tea.String("2017-08-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &QueryAppSecurityInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &QueryAppSecurityInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询应用对应的安全字段
//
// @param request - QueryAppSecurityInfoRequest
//
// @return QueryAppSecurityInfoResponse
func (client *Client) QueryAppSecurityInfo(request *QueryAppSecurityInfoRequest) (_result *QueryAppSecurityInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAppSecurityInfoResponse{}
	_body, _err := client.QueryAppSecurityInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryProductInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProductInfoResponse
func (client *Client) QueryProductInfoWithOptions(request *QueryProductInfoRequest, runtime *util.RuntimeOptions) (_result *QueryProductInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryProductInfo"),
		Version:     tea.String("2017-08-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &QueryProductInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &QueryProductInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - QueryProductInfoRequest
//
// @return QueryProductInfoResponse
func (client *Client) QueryProductInfo(request *QueryProductInfoRequest) (_result *QueryProductInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryProductInfoResponse{}
	_body, _err := client.QueryProductInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
