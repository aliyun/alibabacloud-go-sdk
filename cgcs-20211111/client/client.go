// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AdaptCreateServiceRequest struct {
	AdaptTarget    *AdaptCreateServiceRequestAdaptTarget `json:"AdaptTarget,omitempty" xml:"AdaptTarget,omitempty" type:"Struct"`
	AppVersionId   *string                               `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	AppVersionName *string                               `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	RequestApp     *string                               `json:"RequestApp,omitempty" xml:"RequestApp,omitempty"`
}

func (s AdaptCreateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AdaptCreateServiceRequest) GoString() string {
	return s.String()
}

func (s *AdaptCreateServiceRequest) SetAdaptTarget(v *AdaptCreateServiceRequestAdaptTarget) *AdaptCreateServiceRequest {
	s.AdaptTarget = v
	return s
}

func (s *AdaptCreateServiceRequest) SetAppVersionId(v string) *AdaptCreateServiceRequest {
	s.AppVersionId = &v
	return s
}

func (s *AdaptCreateServiceRequest) SetAppVersionName(v string) *AdaptCreateServiceRequest {
	s.AppVersionName = &v
	return s
}

func (s *AdaptCreateServiceRequest) SetRequestApp(v string) *AdaptCreateServiceRequest {
	s.RequestApp = &v
	return s
}

type AdaptCreateServiceRequestAdaptTarget struct {
	BitRate      *int32  `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	FrameRate    *int32  `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	Resolution   *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	StartProgram *string `json:"StartProgram,omitempty" xml:"StartProgram,omitempty"`
}

func (s AdaptCreateServiceRequestAdaptTarget) String() string {
	return tea.Prettify(s)
}

func (s AdaptCreateServiceRequestAdaptTarget) GoString() string {
	return s.String()
}

func (s *AdaptCreateServiceRequestAdaptTarget) SetBitRate(v int32) *AdaptCreateServiceRequestAdaptTarget {
	s.BitRate = &v
	return s
}

func (s *AdaptCreateServiceRequestAdaptTarget) SetFrameRate(v int32) *AdaptCreateServiceRequestAdaptTarget {
	s.FrameRate = &v
	return s
}

func (s *AdaptCreateServiceRequestAdaptTarget) SetResolution(v string) *AdaptCreateServiceRequestAdaptTarget {
	s.Resolution = &v
	return s
}

func (s *AdaptCreateServiceRequestAdaptTarget) SetStartProgram(v string) *AdaptCreateServiceRequestAdaptTarget {
	s.StartProgram = &v
	return s
}

type AdaptCreateServiceShrinkRequest struct {
	AdaptTargetShrink *string `json:"AdaptTarget,omitempty" xml:"AdaptTarget,omitempty"`
	AppVersionId      *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	AppVersionName    *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	RequestApp        *string `json:"RequestApp,omitempty" xml:"RequestApp,omitempty"`
}

func (s AdaptCreateServiceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AdaptCreateServiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *AdaptCreateServiceShrinkRequest) SetAdaptTargetShrink(v string) *AdaptCreateServiceShrinkRequest {
	s.AdaptTargetShrink = &v
	return s
}

func (s *AdaptCreateServiceShrinkRequest) SetAppVersionId(v string) *AdaptCreateServiceShrinkRequest {
	s.AppVersionId = &v
	return s
}

func (s *AdaptCreateServiceShrinkRequest) SetAppVersionName(v string) *AdaptCreateServiceShrinkRequest {
	s.AppVersionName = &v
	return s
}

func (s *AdaptCreateServiceShrinkRequest) SetRequestApp(v string) *AdaptCreateServiceShrinkRequest {
	s.RequestApp = &v
	return s
}

type AdaptCreateServiceResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AdaptCreateServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AdaptCreateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AdaptCreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AdaptCreateServiceResponseBody) SetCode(v string) *AdaptCreateServiceResponseBody {
	s.Code = &v
	return s
}

func (s *AdaptCreateServiceResponseBody) SetData(v *AdaptCreateServiceResponseBodyData) *AdaptCreateServiceResponseBody {
	s.Data = v
	return s
}

func (s *AdaptCreateServiceResponseBody) SetMessage(v string) *AdaptCreateServiceResponseBody {
	s.Message = &v
	return s
}

func (s *AdaptCreateServiceResponseBody) SetRequestId(v string) *AdaptCreateServiceResponseBody {
	s.RequestId = &v
	return s
}

type AdaptCreateServiceResponseBodyData struct {
	AdaptApplyId *int64 `json:"AdaptApplyId,omitempty" xml:"AdaptApplyId,omitempty"`
}

func (s AdaptCreateServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AdaptCreateServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AdaptCreateServiceResponseBodyData) SetAdaptApplyId(v int64) *AdaptCreateServiceResponseBodyData {
	s.AdaptApplyId = &v
	return s
}

type AdaptCreateServiceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AdaptCreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AdaptCreateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AdaptCreateServiceResponse) GoString() string {
	return s.String()
}

func (s *AdaptCreateServiceResponse) SetHeaders(v map[string]*string) *AdaptCreateServiceResponse {
	s.Headers = v
	return s
}

func (s *AdaptCreateServiceResponse) SetBody(v *AdaptCreateServiceResponseBody) *AdaptCreateServiceResponse {
	s.Body = v
	return s
}

type AdaptGetServiceRequest struct {
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	RequestApp   *string `json:"RequestApp,omitempty" xml:"RequestApp,omitempty"`
}

func (s AdaptGetServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AdaptGetServiceRequest) GoString() string {
	return s.String()
}

func (s *AdaptGetServiceRequest) SetAppVersionId(v string) *AdaptGetServiceRequest {
	s.AppVersionId = &v
	return s
}

func (s *AdaptGetServiceRequest) SetRequestApp(v string) *AdaptGetServiceRequest {
	s.RequestApp = &v
	return s
}

type AdaptGetServiceResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AdaptGetServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AdaptGetServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AdaptGetServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AdaptGetServiceResponseBody) SetCode(v string) *AdaptGetServiceResponseBody {
	s.Code = &v
	return s
}

func (s *AdaptGetServiceResponseBody) SetData(v *AdaptGetServiceResponseBodyData) *AdaptGetServiceResponseBody {
	s.Data = v
	return s
}

func (s *AdaptGetServiceResponseBody) SetMessage(v string) *AdaptGetServiceResponseBody {
	s.Message = &v
	return s
}

func (s *AdaptGetServiceResponseBody) SetRequestId(v string) *AdaptGetServiceResponseBody {
	s.RequestId = &v
	return s
}

type AdaptGetServiceResponseBodyData struct {
	AdaptStatus  *string                                     `json:"AdaptStatus,omitempty" xml:"AdaptStatus,omitempty"`
	AdaptTarget  *AdaptGetServiceResponseBodyDataAdaptTarget `json:"AdaptTarget,omitempty" xml:"AdaptTarget,omitempty" type:"Struct"`
	AppId        *string                                     `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersionId *string                                     `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	GmtCreate    *string                                     `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *string                                     `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id           *int64                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	TenantId     *int64                                      `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s AdaptGetServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AdaptGetServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AdaptGetServiceResponseBodyData) SetAdaptStatus(v string) *AdaptGetServiceResponseBodyData {
	s.AdaptStatus = &v
	return s
}

func (s *AdaptGetServiceResponseBodyData) SetAdaptTarget(v *AdaptGetServiceResponseBodyDataAdaptTarget) *AdaptGetServiceResponseBodyData {
	s.AdaptTarget = v
	return s
}

func (s *AdaptGetServiceResponseBodyData) SetAppId(v string) *AdaptGetServiceResponseBodyData {
	s.AppId = &v
	return s
}

func (s *AdaptGetServiceResponseBodyData) SetAppVersionId(v string) *AdaptGetServiceResponseBodyData {
	s.AppVersionId = &v
	return s
}

func (s *AdaptGetServiceResponseBodyData) SetGmtCreate(v string) *AdaptGetServiceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *AdaptGetServiceResponseBodyData) SetGmtModified(v string) *AdaptGetServiceResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *AdaptGetServiceResponseBodyData) SetId(v int64) *AdaptGetServiceResponseBodyData {
	s.Id = &v
	return s
}

func (s *AdaptGetServiceResponseBodyData) SetTenantId(v int64) *AdaptGetServiceResponseBodyData {
	s.TenantId = &v
	return s
}

type AdaptGetServiceResponseBodyDataAdaptTarget struct {
	BitRate      *int32  `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	FrameRate    *int32  `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	Resolution   *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	StartProgram *string `json:"StartProgram,omitempty" xml:"StartProgram,omitempty"`
}

func (s AdaptGetServiceResponseBodyDataAdaptTarget) String() string {
	return tea.Prettify(s)
}

func (s AdaptGetServiceResponseBodyDataAdaptTarget) GoString() string {
	return s.String()
}

func (s *AdaptGetServiceResponseBodyDataAdaptTarget) SetBitRate(v int32) *AdaptGetServiceResponseBodyDataAdaptTarget {
	s.BitRate = &v
	return s
}

func (s *AdaptGetServiceResponseBodyDataAdaptTarget) SetFrameRate(v int32) *AdaptGetServiceResponseBodyDataAdaptTarget {
	s.FrameRate = &v
	return s
}

func (s *AdaptGetServiceResponseBodyDataAdaptTarget) SetResolution(v string) *AdaptGetServiceResponseBodyDataAdaptTarget {
	s.Resolution = &v
	return s
}

func (s *AdaptGetServiceResponseBodyDataAdaptTarget) SetStartProgram(v string) *AdaptGetServiceResponseBodyDataAdaptTarget {
	s.StartProgram = &v
	return s
}

type AdaptGetServiceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AdaptGetServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AdaptGetServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AdaptGetServiceResponse) GoString() string {
	return s.String()
}

func (s *AdaptGetServiceResponse) SetHeaders(v map[string]*string) *AdaptGetServiceResponse {
	s.Headers = v
	return s
}

func (s *AdaptGetServiceResponse) SetBody(v *AdaptGetServiceResponseBody) *AdaptGetServiceResponse {
	s.Body = v
	return s
}

type AppCreateServiceRequest struct {
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType    *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	RequestApp *string `json:"RequestApp,omitempty" xml:"RequestApp,omitempty"`
}

func (s AppCreateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AppCreateServiceRequest) GoString() string {
	return s.String()
}

func (s *AppCreateServiceRequest) SetAppName(v string) *AppCreateServiceRequest {
	s.AppName = &v
	return s
}

func (s *AppCreateServiceRequest) SetAppType(v string) *AppCreateServiceRequest {
	s.AppType = &v
	return s
}

func (s *AppCreateServiceRequest) SetRequestApp(v string) *AppCreateServiceRequest {
	s.RequestApp = &v
	return s
}

type AppCreateServiceResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AppCreateServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppCreateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppCreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AppCreateServiceResponseBody) SetCode(v string) *AppCreateServiceResponseBody {
	s.Code = &v
	return s
}

func (s *AppCreateServiceResponseBody) SetData(v *AppCreateServiceResponseBodyData) *AppCreateServiceResponseBody {
	s.Data = v
	return s
}

func (s *AppCreateServiceResponseBody) SetMessage(v string) *AppCreateServiceResponseBody {
	s.Message = &v
	return s
}

func (s *AppCreateServiceResponseBody) SetRequestId(v string) *AppCreateServiceResponseBody {
	s.RequestId = &v
	return s
}

type AppCreateServiceResponseBodyData struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s AppCreateServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AppCreateServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppCreateServiceResponseBodyData) SetAppId(v string) *AppCreateServiceResponseBodyData {
	s.AppId = &v
	return s
}

type AppCreateServiceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppCreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppCreateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AppCreateServiceResponse) GoString() string {
	return s.String()
}

func (s *AppCreateServiceResponse) SetHeaders(v map[string]*string) *AppCreateServiceResponse {
	s.Headers = v
	return s
}

func (s *AppCreateServiceResponse) SetBody(v *AppCreateServiceResponseBody) *AppCreateServiceResponse {
	s.Body = v
	return s
}

type AppDeleteServiceRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestApp *string `json:"RequestApp,omitempty" xml:"RequestApp,omitempty"`
}

func (s AppDeleteServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AppDeleteServiceRequest) GoString() string {
	return s.String()
}

func (s *AppDeleteServiceRequest) SetAppId(v string) *AppDeleteServiceRequest {
	s.AppId = &v
	return s
}

func (s *AppDeleteServiceRequest) SetRequestApp(v string) *AppDeleteServiceRequest {
	s.RequestApp = &v
	return s
}

type AppDeleteServiceResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AppDeleteServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppDeleteServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppDeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AppDeleteServiceResponseBody) SetCode(v string) *AppDeleteServiceResponseBody {
	s.Code = &v
	return s
}

func (s *AppDeleteServiceResponseBody) SetData(v *AppDeleteServiceResponseBodyData) *AppDeleteServiceResponseBody {
	s.Data = v
	return s
}

func (s *AppDeleteServiceResponseBody) SetMessage(v string) *AppDeleteServiceResponseBody {
	s.Message = &v
	return s
}

func (s *AppDeleteServiceResponseBody) SetRequestId(v string) *AppDeleteServiceResponseBody {
	s.RequestId = &v
	return s
}

type AppDeleteServiceResponseBodyData struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s AppDeleteServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AppDeleteServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppDeleteServiceResponseBodyData) SetAppId(v string) *AppDeleteServiceResponseBodyData {
	s.AppId = &v
	return s
}

type AppDeleteServiceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppDeleteServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppDeleteServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AppDeleteServiceResponse) GoString() string {
	return s.String()
}

func (s *AppDeleteServiceResponse) SetHeaders(v map[string]*string) *AppDeleteServiceResponse {
	s.Headers = v
	return s
}

func (s *AppDeleteServiceResponse) SetBody(v *AppDeleteServiceResponseBody) *AppDeleteServiceResponse {
	s.Body = v
	return s
}

type AppGetServiceRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestApp *string `json:"RequestApp,omitempty" xml:"RequestApp,omitempty"`
}

func (s AppGetServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AppGetServiceRequest) GoString() string {
	return s.String()
}

func (s *AppGetServiceRequest) SetAppId(v string) *AppGetServiceRequest {
	s.AppId = &v
	return s
}

func (s *AppGetServiceRequest) SetRequestApp(v string) *AppGetServiceRequest {
	s.RequestApp = &v
	return s
}

type AppGetServiceResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AppGetServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppGetServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppGetServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AppGetServiceResponseBody) SetCode(v string) *AppGetServiceResponseBody {
	s.Code = &v
	return s
}

func (s *AppGetServiceResponseBody) SetData(v *AppGetServiceResponseBodyData) *AppGetServiceResponseBody {
	s.Data = v
	return s
}

func (s *AppGetServiceResponseBody) SetMessage(v string) *AppGetServiceResponseBody {
	s.Message = &v
	return s
}

func (s *AppGetServiceResponseBody) SetRequestId(v string) *AppGetServiceResponseBody {
	s.RequestId = &v
	return s
}

type AppGetServiceResponseBodyData struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType     *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	TenantId    *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s AppGetServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AppGetServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppGetServiceResponseBodyData) SetAppId(v string) *AppGetServiceResponseBodyData {
	s.AppId = &v
	return s
}

func (s *AppGetServiceResponseBodyData) SetAppName(v string) *AppGetServiceResponseBodyData {
	s.AppName = &v
	return s
}

func (s *AppGetServiceResponseBodyData) SetAppType(v string) *AppGetServiceResponseBodyData {
	s.AppType = &v
	return s
}

func (s *AppGetServiceResponseBodyData) SetBizType(v string) *AppGetServiceResponseBodyData {
	s.BizType = &v
	return s
}

func (s *AppGetServiceResponseBodyData) SetGmtCreate(v string) *AppGetServiceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *AppGetServiceResponseBodyData) SetGmtModified(v string) *AppGetServiceResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *AppGetServiceResponseBodyData) SetTenantId(v int64) *AppGetServiceResponseBodyData {
	s.TenantId = &v
	return s
}

type AppGetServiceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppGetServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppGetServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AppGetServiceResponse) GoString() string {
	return s.String()
}

func (s *AppGetServiceResponse) SetHeaders(v map[string]*string) *AppGetServiceResponse {
	s.Headers = v
	return s
}

func (s *AppGetServiceResponse) SetBody(v *AppGetServiceResponseBody) *AppGetServiceResponse {
	s.Body = v
	return s
}

type AppListServiceRequest struct {
	KeySearch  *string `json:"KeySearch,omitempty" xml:"KeySearch,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestApp *string `json:"RequestApp,omitempty" xml:"RequestApp,omitempty"`
}

func (s AppListServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AppListServiceRequest) GoString() string {
	return s.String()
}

func (s *AppListServiceRequest) SetKeySearch(v string) *AppListServiceRequest {
	s.KeySearch = &v
	return s
}

func (s *AppListServiceRequest) SetPageNumber(v int32) *AppListServiceRequest {
	s.PageNumber = &v
	return s
}

func (s *AppListServiceRequest) SetPageSize(v int32) *AppListServiceRequest {
	s.PageSize = &v
	return s
}

func (s *AppListServiceRequest) SetRequestApp(v string) *AppListServiceRequest {
	s.RequestApp = &v
	return s
}

type AppListServiceResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AppListServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppListServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppListServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AppListServiceResponseBody) SetCode(v string) *AppListServiceResponseBody {
	s.Code = &v
	return s
}

func (s *AppListServiceResponseBody) SetData(v *AppListServiceResponseBodyData) *AppListServiceResponseBody {
	s.Data = v
	return s
}

func (s *AppListServiceResponseBody) SetMessage(v string) *AppListServiceResponseBody {
	s.Message = &v
	return s
}

func (s *AppListServiceResponseBody) SetRequestId(v string) *AppListServiceResponseBody {
	s.RequestId = &v
	return s
}

type AppListServiceResponseBodyData struct {
	Apps  []*AppListServiceResponseBodyDataApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	Total *int64                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s AppListServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AppListServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppListServiceResponseBodyData) SetApps(v []*AppListServiceResponseBodyDataApps) *AppListServiceResponseBodyData {
	s.Apps = v
	return s
}

func (s *AppListServiceResponseBodyData) SetTotal(v int64) *AppListServiceResponseBodyData {
	s.Total = &v
	return s
}

type AppListServiceResponseBodyDataApps struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType         *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	GmtCreate       *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	TenantId        *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	VersionAdaptNum *int64  `json:"VersionAdaptNum,omitempty" xml:"VersionAdaptNum,omitempty"`
	VersionTotalNum *int64  `json:"VersionTotalNum,omitempty" xml:"VersionTotalNum,omitempty"`
}

func (s AppListServiceResponseBodyDataApps) String() string {
	return tea.Prettify(s)
}

func (s AppListServiceResponseBodyDataApps) GoString() string {
	return s.String()
}

func (s *AppListServiceResponseBodyDataApps) SetAppId(v string) *AppListServiceResponseBodyDataApps {
	s.AppId = &v
	return s
}

func (s *AppListServiceResponseBodyDataApps) SetAppName(v string) *AppListServiceResponseBodyDataApps {
	s.AppName = &v
	return s
}

func (s *AppListServiceResponseBodyDataApps) SetAppType(v string) *AppListServiceResponseBodyDataApps {
	s.AppType = &v
	return s
}

func (s *AppListServiceResponseBodyDataApps) SetGmtCreate(v string) *AppListServiceResponseBodyDataApps {
	s.GmtCreate = &v
	return s
}

func (s *AppListServiceResponseBodyDataApps) SetGmtModified(v string) *AppListServiceResponseBodyDataApps {
	s.GmtModified = &v
	return s
}

func (s *AppListServiceResponseBodyDataApps) SetTenantId(v int64) *AppListServiceResponseBodyDataApps {
	s.TenantId = &v
	return s
}

func (s *AppListServiceResponseBodyDataApps) SetVersionAdaptNum(v int64) *AppListServiceResponseBodyDataApps {
	s.VersionAdaptNum = &v
	return s
}

func (s *AppListServiceResponseBodyDataApps) SetVersionTotalNum(v int64) *AppListServiceResponseBodyDataApps {
	s.VersionTotalNum = &v
	return s
}

type AppListServiceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppListServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppListServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AppListServiceResponse) GoString() string {
	return s.String()
}

func (s *AppListServiceResponse) SetHeaders(v map[string]*string) *AppListServiceResponse {
	s.Headers = v
	return s
}

func (s *AppListServiceResponse) SetBody(v *AppListServiceResponseBody) *AppListServiceResponse {
	s.Body = v
	return s
}

type AppModifyServiceRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	RequestApp *string `json:"RequestApp,omitempty" xml:"RequestApp,omitempty"`
}

func (s AppModifyServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AppModifyServiceRequest) GoString() string {
	return s.String()
}

func (s *AppModifyServiceRequest) SetAppId(v string) *AppModifyServiceRequest {
	s.AppId = &v
	return s
}

func (s *AppModifyServiceRequest) SetAppName(v string) *AppModifyServiceRequest {
	s.AppName = &v
	return s
}

func (s *AppModifyServiceRequest) SetRequestApp(v string) *AppModifyServiceRequest {
	s.RequestApp = &v
	return s
}

type AppModifyServiceResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AppModifyServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppModifyServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppModifyServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AppModifyServiceResponseBody) SetCode(v string) *AppModifyServiceResponseBody {
	s.Code = &v
	return s
}

func (s *AppModifyServiceResponseBody) SetData(v *AppModifyServiceResponseBodyData) *AppModifyServiceResponseBody {
	s.Data = v
	return s
}

func (s *AppModifyServiceResponseBody) SetMessage(v string) *AppModifyServiceResponseBody {
	s.Message = &v
	return s
}

func (s *AppModifyServiceResponseBody) SetRequestId(v string) *AppModifyServiceResponseBody {
	s.RequestId = &v
	return s
}

type AppModifyServiceResponseBodyData struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s AppModifyServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AppModifyServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppModifyServiceResponseBodyData) SetAppId(v string) *AppModifyServiceResponseBodyData {
	s.AppId = &v
	return s
}

type AppModifyServiceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppModifyServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppModifyServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AppModifyServiceResponse) GoString() string {
	return s.String()
}

func (s *AppModifyServiceResponse) SetHeaders(v map[string]*string) *AppModifyServiceResponse {
	s.Headers = v
	return s
}

func (s *AppModifyServiceResponse) SetBody(v *AppModifyServiceResponseBody) *AppModifyServiceResponse {
	s.Body = v
	return s
}

type AppVersionCreateServiceRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	FileAddress    *string `json:"FileAddress,omitempty" xml:"FileAddress,omitempty"`
	FileSize       *int64  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileUploadType *string `json:"FileUploadType,omitempty" xml:"FileUploadType,omitempty"`
	RequestApp     *string `json:"RequestApp,omitempty" xml:"RequestApp,omitempty"`
}

func (s AppVersionCreateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AppVersionCreateServiceRequest) GoString() string {
	return s.String()
}

func (s *AppVersionCreateServiceRequest) SetAppId(v string) *AppVersionCreateServiceRequest {
	s.AppId = &v
	return s
}

func (s *AppVersionCreateServiceRequest) SetAppVersionName(v string) *AppVersionCreateServiceRequest {
	s.AppVersionName = &v
	return s
}

func (s *AppVersionCreateServiceRequest) SetFileAddress(v string) *AppVersionCreateServiceRequest {
	s.FileAddress = &v
	return s
}

func (s *AppVersionCreateServiceRequest) SetFileSize(v int64) *AppVersionCreateServiceRequest {
	s.FileSize = &v
	return s
}

func (s *AppVersionCreateServiceRequest) SetFileUploadType(v string) *AppVersionCreateServiceRequest {
	s.FileUploadType = &v
	return s
}

func (s *AppVersionCreateServiceRequest) SetRequestApp(v string) *AppVersionCreateServiceRequest {
	s.RequestApp = &v
	return s
}

type AppVersionCreateServiceResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AppVersionCreateServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppVersionCreateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppVersionCreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AppVersionCreateServiceResponseBody) SetCode(v string) *AppVersionCreateServiceResponseBody {
	s.Code = &v
	return s
}

func (s *AppVersionCreateServiceResponseBody) SetData(v *AppVersionCreateServiceResponseBodyData) *AppVersionCreateServiceResponseBody {
	s.Data = v
	return s
}

func (s *AppVersionCreateServiceResponseBody) SetMessage(v string) *AppVersionCreateServiceResponseBody {
	s.Message = &v
	return s
}

func (s *AppVersionCreateServiceResponseBody) SetRequestId(v string) *AppVersionCreateServiceResponseBody {
	s.RequestId = &v
	return s
}

type AppVersionCreateServiceResponseBodyData struct {
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
}

func (s AppVersionCreateServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AppVersionCreateServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppVersionCreateServiceResponseBodyData) SetAppVersionId(v string) *AppVersionCreateServiceResponseBodyData {
	s.AppVersionId = &v
	return s
}

type AppVersionCreateServiceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppVersionCreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppVersionCreateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AppVersionCreateServiceResponse) GoString() string {
	return s.String()
}

func (s *AppVersionCreateServiceResponse) SetHeaders(v map[string]*string) *AppVersionCreateServiceResponse {
	s.Headers = v
	return s
}

func (s *AppVersionCreateServiceResponse) SetBody(v *AppVersionCreateServiceResponseBody) *AppVersionCreateServiceResponse {
	s.Body = v
	return s
}

type AppVersionDeleteServiceRequest struct {
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	RequestApp   *string `json:"RequestApp,omitempty" xml:"RequestApp,omitempty"`
}

func (s AppVersionDeleteServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AppVersionDeleteServiceRequest) GoString() string {
	return s.String()
}

func (s *AppVersionDeleteServiceRequest) SetAppVersionId(v string) *AppVersionDeleteServiceRequest {
	s.AppVersionId = &v
	return s
}

func (s *AppVersionDeleteServiceRequest) SetRequestApp(v string) *AppVersionDeleteServiceRequest {
	s.RequestApp = &v
	return s
}

type AppVersionDeleteServiceResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AppVersionDeleteServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppVersionDeleteServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppVersionDeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AppVersionDeleteServiceResponseBody) SetCode(v string) *AppVersionDeleteServiceResponseBody {
	s.Code = &v
	return s
}

func (s *AppVersionDeleteServiceResponseBody) SetData(v *AppVersionDeleteServiceResponseBodyData) *AppVersionDeleteServiceResponseBody {
	s.Data = v
	return s
}

func (s *AppVersionDeleteServiceResponseBody) SetMessage(v string) *AppVersionDeleteServiceResponseBody {
	s.Message = &v
	return s
}

func (s *AppVersionDeleteServiceResponseBody) SetRequestId(v string) *AppVersionDeleteServiceResponseBody {
	s.RequestId = &v
	return s
}

type AppVersionDeleteServiceResponseBodyData struct {
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
}

func (s AppVersionDeleteServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AppVersionDeleteServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppVersionDeleteServiceResponseBodyData) SetAppVersionId(v string) *AppVersionDeleteServiceResponseBodyData {
	s.AppVersionId = &v
	return s
}

type AppVersionDeleteServiceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppVersionDeleteServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppVersionDeleteServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AppVersionDeleteServiceResponse) GoString() string {
	return s.String()
}

func (s *AppVersionDeleteServiceResponse) SetHeaders(v map[string]*string) *AppVersionDeleteServiceResponse {
	s.Headers = v
	return s
}

func (s *AppVersionDeleteServiceResponse) SetBody(v *AppVersionDeleteServiceResponseBody) *AppVersionDeleteServiceResponse {
	s.Body = v
	return s
}

type AppVersionGetServiceRequest struct {
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	RequestApp   *string `json:"RequestApp,omitempty" xml:"RequestApp,omitempty"`
}

func (s AppVersionGetServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AppVersionGetServiceRequest) GoString() string {
	return s.String()
}

func (s *AppVersionGetServiceRequest) SetAppVersionId(v string) *AppVersionGetServiceRequest {
	s.AppVersionId = &v
	return s
}

func (s *AppVersionGetServiceRequest) SetRequestApp(v string) *AppVersionGetServiceRequest {
	s.RequestApp = &v
	return s
}

type AppVersionGetServiceResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AppVersionGetServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppVersionGetServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppVersionGetServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AppVersionGetServiceResponseBody) SetCode(v string) *AppVersionGetServiceResponseBody {
	s.Code = &v
	return s
}

func (s *AppVersionGetServiceResponseBody) SetData(v *AppVersionGetServiceResponseBodyData) *AppVersionGetServiceResponseBody {
	s.Data = v
	return s
}

func (s *AppVersionGetServiceResponseBody) SetMessage(v string) *AppVersionGetServiceResponseBody {
	s.Message = &v
	return s
}

func (s *AppVersionGetServiceResponseBody) SetRequestId(v string) *AppVersionGetServiceResponseBody {
	s.RequestId = &v
	return s
}

type AppVersionGetServiceResponseBodyData struct {
	AppId                *string  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersionId         *string  `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	AppVersionName       *string  `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	AppVersionStatus     *string  `json:"AppVersionStatus,omitempty" xml:"AppVersionStatus,omitempty"`
	AppVersionStatusMemo *string  `json:"AppVersionStatusMemo,omitempty" xml:"AppVersionStatusMemo,omitempty"`
	ConsumeCu            *float64 `json:"ConsumeCu,omitempty" xml:"ConsumeCu,omitempty"`
	FileAddress          *string  `json:"FileAddress,omitempty" xml:"FileAddress,omitempty"`
	FileSize             *int64   `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileUploadFinishTime *string  `json:"FileUploadFinishTime,omitempty" xml:"FileUploadFinishTime,omitempty"`
	FileUploadType       *string  `json:"FileUploadType,omitempty" xml:"FileUploadType,omitempty"`
	GmtCreate            *string  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified          *string  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	SourceVersionId      *string  `json:"SourceVersionId,omitempty" xml:"SourceVersionId,omitempty"`
	TenantId             *int64   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s AppVersionGetServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AppVersionGetServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppVersionGetServiceResponseBodyData) SetAppId(v string) *AppVersionGetServiceResponseBodyData {
	s.AppId = &v
	return s
}

func (s *AppVersionGetServiceResponseBodyData) SetAppVersionId(v string) *AppVersionGetServiceResponseBodyData {
	s.AppVersionId = &v
	return s
}

func (s *AppVersionGetServiceResponseBodyData) SetAppVersionName(v string) *AppVersionGetServiceResponseBodyData {
	s.AppVersionName = &v
	return s
}

func (s *AppVersionGetServiceResponseBodyData) SetAppVersionStatus(v string) *AppVersionGetServiceResponseBodyData {
	s.AppVersionStatus = &v
	return s
}

func (s *AppVersionGetServiceResponseBodyData) SetAppVersionStatusMemo(v string) *AppVersionGetServiceResponseBodyData {
	s.AppVersionStatusMemo = &v
	return s
}

func (s *AppVersionGetServiceResponseBodyData) SetConsumeCu(v float64) *AppVersionGetServiceResponseBodyData {
	s.ConsumeCu = &v
	return s
}

func (s *AppVersionGetServiceResponseBodyData) SetFileAddress(v string) *AppVersionGetServiceResponseBodyData {
	s.FileAddress = &v
	return s
}

func (s *AppVersionGetServiceResponseBodyData) SetFileSize(v int64) *AppVersionGetServiceResponseBodyData {
	s.FileSize = &v
	return s
}

func (s *AppVersionGetServiceResponseBodyData) SetFileUploadFinishTime(v string) *AppVersionGetServiceResponseBodyData {
	s.FileUploadFinishTime = &v
	return s
}

func (s *AppVersionGetServiceResponseBodyData) SetFileUploadType(v string) *AppVersionGetServiceResponseBodyData {
	s.FileUploadType = &v
	return s
}

func (s *AppVersionGetServiceResponseBodyData) SetGmtCreate(v string) *AppVersionGetServiceResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *AppVersionGetServiceResponseBodyData) SetGmtModified(v string) *AppVersionGetServiceResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *AppVersionGetServiceResponseBodyData) SetSourceVersionId(v string) *AppVersionGetServiceResponseBodyData {
	s.SourceVersionId = &v
	return s
}

func (s *AppVersionGetServiceResponseBodyData) SetTenantId(v int64) *AppVersionGetServiceResponseBodyData {
	s.TenantId = &v
	return s
}

type AppVersionGetServiceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppVersionGetServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppVersionGetServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AppVersionGetServiceResponse) GoString() string {
	return s.String()
}

func (s *AppVersionGetServiceResponse) SetHeaders(v map[string]*string) *AppVersionGetServiceResponse {
	s.Headers = v
	return s
}

func (s *AppVersionGetServiceResponse) SetBody(v *AppVersionGetServiceResponseBody) *AppVersionGetServiceResponse {
	s.Body = v
	return s
}

type AppVersionListServiceRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestApp *string `json:"RequestApp,omitempty" xml:"RequestApp,omitempty"`
}

func (s AppVersionListServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AppVersionListServiceRequest) GoString() string {
	return s.String()
}

func (s *AppVersionListServiceRequest) SetAppId(v string) *AppVersionListServiceRequest {
	s.AppId = &v
	return s
}

func (s *AppVersionListServiceRequest) SetPageNumber(v int32) *AppVersionListServiceRequest {
	s.PageNumber = &v
	return s
}

func (s *AppVersionListServiceRequest) SetPageSize(v int32) *AppVersionListServiceRequest {
	s.PageSize = &v
	return s
}

func (s *AppVersionListServiceRequest) SetRequestApp(v string) *AppVersionListServiceRequest {
	s.RequestApp = &v
	return s
}

type AppVersionListServiceResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AppVersionListServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppVersionListServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppVersionListServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AppVersionListServiceResponseBody) SetCode(v string) *AppVersionListServiceResponseBody {
	s.Code = &v
	return s
}

func (s *AppVersionListServiceResponseBody) SetData(v *AppVersionListServiceResponseBodyData) *AppVersionListServiceResponseBody {
	s.Data = v
	return s
}

func (s *AppVersionListServiceResponseBody) SetMessage(v string) *AppVersionListServiceResponseBody {
	s.Message = &v
	return s
}

func (s *AppVersionListServiceResponseBody) SetRequestId(v string) *AppVersionListServiceResponseBody {
	s.RequestId = &v
	return s
}

type AppVersionListServiceResponseBodyData struct {
	Total    *int64                                           `json:"Total,omitempty" xml:"Total,omitempty"`
	Versions []*AppVersionListServiceResponseBodyDataVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s AppVersionListServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AppVersionListServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppVersionListServiceResponseBodyData) SetTotal(v int64) *AppVersionListServiceResponseBodyData {
	s.Total = &v
	return s
}

func (s *AppVersionListServiceResponseBodyData) SetVersions(v []*AppVersionListServiceResponseBodyDataVersions) *AppVersionListServiceResponseBodyData {
	s.Versions = v
	return s
}

type AppVersionListServiceResponseBodyDataVersions struct {
	AppId                *string  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersionId         *string  `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	AppVersionName       *string  `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	AppVersionStatus     *string  `json:"AppVersionStatus,omitempty" xml:"AppVersionStatus,omitempty"`
	AppVersionStatusMemo *string  `json:"AppVersionStatusMemo,omitempty" xml:"AppVersionStatusMemo,omitempty"`
	ConsumeCu            *float64 `json:"ConsumeCu,omitempty" xml:"ConsumeCu,omitempty"`
	FileAddress          *string  `json:"FileAddress,omitempty" xml:"FileAddress,omitempty"`
	FileSize             *int64   `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileUploadFinishTime *string  `json:"FileUploadFinishTime,omitempty" xml:"FileUploadFinishTime,omitempty"`
	FileUploadType       *string  `json:"FileUploadType,omitempty" xml:"FileUploadType,omitempty"`
	GmtCreate            *string  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified          *string  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	TenantId             *int64   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s AppVersionListServiceResponseBodyDataVersions) String() string {
	return tea.Prettify(s)
}

func (s AppVersionListServiceResponseBodyDataVersions) GoString() string {
	return s.String()
}

func (s *AppVersionListServiceResponseBodyDataVersions) SetAppId(v string) *AppVersionListServiceResponseBodyDataVersions {
	s.AppId = &v
	return s
}

func (s *AppVersionListServiceResponseBodyDataVersions) SetAppVersionId(v string) *AppVersionListServiceResponseBodyDataVersions {
	s.AppVersionId = &v
	return s
}

func (s *AppVersionListServiceResponseBodyDataVersions) SetAppVersionName(v string) *AppVersionListServiceResponseBodyDataVersions {
	s.AppVersionName = &v
	return s
}

func (s *AppVersionListServiceResponseBodyDataVersions) SetAppVersionStatus(v string) *AppVersionListServiceResponseBodyDataVersions {
	s.AppVersionStatus = &v
	return s
}

func (s *AppVersionListServiceResponseBodyDataVersions) SetAppVersionStatusMemo(v string) *AppVersionListServiceResponseBodyDataVersions {
	s.AppVersionStatusMemo = &v
	return s
}

func (s *AppVersionListServiceResponseBodyDataVersions) SetConsumeCu(v float64) *AppVersionListServiceResponseBodyDataVersions {
	s.ConsumeCu = &v
	return s
}

func (s *AppVersionListServiceResponseBodyDataVersions) SetFileAddress(v string) *AppVersionListServiceResponseBodyDataVersions {
	s.FileAddress = &v
	return s
}

func (s *AppVersionListServiceResponseBodyDataVersions) SetFileSize(v int64) *AppVersionListServiceResponseBodyDataVersions {
	s.FileSize = &v
	return s
}

func (s *AppVersionListServiceResponseBodyDataVersions) SetFileUploadFinishTime(v string) *AppVersionListServiceResponseBodyDataVersions {
	s.FileUploadFinishTime = &v
	return s
}

func (s *AppVersionListServiceResponseBodyDataVersions) SetFileUploadType(v string) *AppVersionListServiceResponseBodyDataVersions {
	s.FileUploadType = &v
	return s
}

func (s *AppVersionListServiceResponseBodyDataVersions) SetGmtCreate(v string) *AppVersionListServiceResponseBodyDataVersions {
	s.GmtCreate = &v
	return s
}

func (s *AppVersionListServiceResponseBodyDataVersions) SetGmtModified(v string) *AppVersionListServiceResponseBodyDataVersions {
	s.GmtModified = &v
	return s
}

func (s *AppVersionListServiceResponseBodyDataVersions) SetTenantId(v int64) *AppVersionListServiceResponseBodyDataVersions {
	s.TenantId = &v
	return s
}

type AppVersionListServiceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppVersionListServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppVersionListServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AppVersionListServiceResponse) GoString() string {
	return s.String()
}

func (s *AppVersionListServiceResponse) SetHeaders(v map[string]*string) *AppVersionListServiceResponse {
	s.Headers = v
	return s
}

func (s *AppVersionListServiceResponse) SetBody(v *AppVersionListServiceResponseBody) *AppVersionListServiceResponse {
	s.Body = v
	return s
}

type AppVersionModifyServiceRequest struct {
	AppVersionId   *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	RequestApp     *string `json:"RequestApp,omitempty" xml:"RequestApp,omitempty"`
}

func (s AppVersionModifyServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AppVersionModifyServiceRequest) GoString() string {
	return s.String()
}

func (s *AppVersionModifyServiceRequest) SetAppVersionId(v string) *AppVersionModifyServiceRequest {
	s.AppVersionId = &v
	return s
}

func (s *AppVersionModifyServiceRequest) SetAppVersionName(v string) *AppVersionModifyServiceRequest {
	s.AppVersionName = &v
	return s
}

func (s *AppVersionModifyServiceRequest) SetRequestApp(v string) *AppVersionModifyServiceRequest {
	s.RequestApp = &v
	return s
}

type AppVersionModifyServiceResponseBody struct {
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AppVersionModifyServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppVersionModifyServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppVersionModifyServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AppVersionModifyServiceResponseBody) SetCode(v string) *AppVersionModifyServiceResponseBody {
	s.Code = &v
	return s
}

func (s *AppVersionModifyServiceResponseBody) SetData(v *AppVersionModifyServiceResponseBodyData) *AppVersionModifyServiceResponseBody {
	s.Data = v
	return s
}

func (s *AppVersionModifyServiceResponseBody) SetMessage(v string) *AppVersionModifyServiceResponseBody {
	s.Message = &v
	return s
}

func (s *AppVersionModifyServiceResponseBody) SetRequestId(v string) *AppVersionModifyServiceResponseBody {
	s.RequestId = &v
	return s
}

type AppVersionModifyServiceResponseBodyData struct {
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
}

func (s AppVersionModifyServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AppVersionModifyServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppVersionModifyServiceResponseBodyData) SetAppVersionId(v string) *AppVersionModifyServiceResponseBodyData {
	s.AppVersionId = &v
	return s
}

type AppVersionModifyServiceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppVersionModifyServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppVersionModifyServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AppVersionModifyServiceResponse) GoString() string {
	return s.String()
}

func (s *AppVersionModifyServiceResponse) SetHeaders(v map[string]*string) *AppVersionModifyServiceResponse {
	s.Headers = v
	return s
}

func (s *AppVersionModifyServiceResponse) SetBody(v *AppVersionModifyServiceResponseBody) *AppVersionModifyServiceResponse {
	s.Body = v
	return s
}

type AppVersionQueryServiceRequest struct {
	KeySearch  *string `json:"KeySearch,omitempty" xml:"KeySearch,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestApp *string `json:"RequestApp,omitempty" xml:"RequestApp,omitempty"`
}

func (s AppVersionQueryServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AppVersionQueryServiceRequest) GoString() string {
	return s.String()
}

func (s *AppVersionQueryServiceRequest) SetKeySearch(v string) *AppVersionQueryServiceRequest {
	s.KeySearch = &v
	return s
}

func (s *AppVersionQueryServiceRequest) SetPageNumber(v int32) *AppVersionQueryServiceRequest {
	s.PageNumber = &v
	return s
}

func (s *AppVersionQueryServiceRequest) SetPageSize(v int32) *AppVersionQueryServiceRequest {
	s.PageSize = &v
	return s
}

func (s *AppVersionQueryServiceRequest) SetRequestApp(v string) *AppVersionQueryServiceRequest {
	s.RequestApp = &v
	return s
}

type AppVersionQueryServiceResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *AppVersionQueryServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppVersionQueryServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppVersionQueryServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AppVersionQueryServiceResponseBody) SetCode(v string) *AppVersionQueryServiceResponseBody {
	s.Code = &v
	return s
}

func (s *AppVersionQueryServiceResponseBody) SetData(v *AppVersionQueryServiceResponseBodyData) *AppVersionQueryServiceResponseBody {
	s.Data = v
	return s
}

func (s *AppVersionQueryServiceResponseBody) SetMessage(v string) *AppVersionQueryServiceResponseBody {
	s.Message = &v
	return s
}

func (s *AppVersionQueryServiceResponseBody) SetRequestId(v string) *AppVersionQueryServiceResponseBody {
	s.RequestId = &v
	return s
}

type AppVersionQueryServiceResponseBodyData struct {
	Total    *int64                                            `json:"Total,omitempty" xml:"Total,omitempty"`
	Versions []*AppVersionQueryServiceResponseBodyDataVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s AppVersionQueryServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AppVersionQueryServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppVersionQueryServiceResponseBodyData) SetTotal(v int64) *AppVersionQueryServiceResponseBodyData {
	s.Total = &v
	return s
}

func (s *AppVersionQueryServiceResponseBodyData) SetVersions(v []*AppVersionQueryServiceResponseBodyDataVersions) *AppVersionQueryServiceResponseBodyData {
	s.Versions = v
	return s
}

type AppVersionQueryServiceResponseBodyDataVersions struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppVersionId   *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	TenantId       *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s AppVersionQueryServiceResponseBodyDataVersions) String() string {
	return tea.Prettify(s)
}

func (s AppVersionQueryServiceResponseBodyDataVersions) GoString() string {
	return s.String()
}

func (s *AppVersionQueryServiceResponseBodyDataVersions) SetAppId(v string) *AppVersionQueryServiceResponseBodyDataVersions {
	s.AppId = &v
	return s
}

func (s *AppVersionQueryServiceResponseBodyDataVersions) SetAppName(v string) *AppVersionQueryServiceResponseBodyDataVersions {
	s.AppName = &v
	return s
}

func (s *AppVersionQueryServiceResponseBodyDataVersions) SetAppVersionId(v string) *AppVersionQueryServiceResponseBodyDataVersions {
	s.AppVersionId = &v
	return s
}

func (s *AppVersionQueryServiceResponseBodyDataVersions) SetAppVersionName(v string) *AppVersionQueryServiceResponseBodyDataVersions {
	s.AppVersionName = &v
	return s
}

func (s *AppVersionQueryServiceResponseBodyDataVersions) SetTenantId(v int64) *AppVersionQueryServiceResponseBodyDataVersions {
	s.TenantId = &v
	return s
}

type AppVersionQueryServiceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppVersionQueryServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppVersionQueryServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AppVersionQueryServiceResponse) GoString() string {
	return s.String()
}

func (s *AppVersionQueryServiceResponse) SetHeaders(v map[string]*string) *AppVersionQueryServiceResponse {
	s.Headers = v
	return s
}

func (s *AppVersionQueryServiceResponse) SetBody(v *AppVersionQueryServiceResponseBody) *AppVersionQueryServiceResponse {
	s.Body = v
	return s
}

type AppliedConsumStatRequest struct {
	AppliedId []*string `json:"AppliedId,omitempty" xml:"AppliedId,omitempty" type:"Repeated"`
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// 资源类型,PackageType[CU(cu),code,cssResourceType,desc]
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// 查询结束时间
	QueryEndDate *string `json:"QueryEndDate,omitempty" xml:"QueryEndDate,omitempty"`
	// 查询开始时间
	QueryStartDate *string `json:"QueryStartDate,omitempty" xml:"QueryStartDate,omitempty"`
}

func (s AppliedConsumStatRequest) String() string {
	return tea.Prettify(s)
}

func (s AppliedConsumStatRequest) GoString() string {
	return s.String()
}

func (s *AppliedConsumStatRequest) SetAppliedId(v []*string) *AppliedConsumStatRequest {
	s.AppliedId = v
	return s
}

func (s *AppliedConsumStatRequest) SetOperatorId(v string) *AppliedConsumStatRequest {
	s.OperatorId = &v
	return s
}

func (s *AppliedConsumStatRequest) SetOperatorType(v string) *AppliedConsumStatRequest {
	s.OperatorType = &v
	return s
}

func (s *AppliedConsumStatRequest) SetPackageType(v string) *AppliedConsumStatRequest {
	s.PackageType = &v
	return s
}

func (s *AppliedConsumStatRequest) SetQueryEndDate(v string) *AppliedConsumStatRequest {
	s.QueryEndDate = &v
	return s
}

func (s *AppliedConsumStatRequest) SetQueryStartDate(v string) *AppliedConsumStatRequest {
	s.QueryStartDate = &v
	return s
}

type AppliedConsumStatShrinkRequest struct {
	AppliedIdShrink *string `json:"AppliedId,omitempty" xml:"AppliedId,omitempty"`
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// 资源类型,PackageType[CU(cu),code,cssResourceType,desc]
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// 查询结束时间
	QueryEndDate *string `json:"QueryEndDate,omitempty" xml:"QueryEndDate,omitempty"`
	// 查询开始时间
	QueryStartDate *string `json:"QueryStartDate,omitempty" xml:"QueryStartDate,omitempty"`
}

func (s AppliedConsumStatShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AppliedConsumStatShrinkRequest) GoString() string {
	return s.String()
}

func (s *AppliedConsumStatShrinkRequest) SetAppliedIdShrink(v string) *AppliedConsumStatShrinkRequest {
	s.AppliedIdShrink = &v
	return s
}

func (s *AppliedConsumStatShrinkRequest) SetOperatorId(v string) *AppliedConsumStatShrinkRequest {
	s.OperatorId = &v
	return s
}

func (s *AppliedConsumStatShrinkRequest) SetOperatorType(v string) *AppliedConsumStatShrinkRequest {
	s.OperatorType = &v
	return s
}

func (s *AppliedConsumStatShrinkRequest) SetPackageType(v string) *AppliedConsumStatShrinkRequest {
	s.PackageType = &v
	return s
}

func (s *AppliedConsumStatShrinkRequest) SetQueryEndDate(v string) *AppliedConsumStatShrinkRequest {
	s.QueryEndDate = &v
	return s
}

func (s *AppliedConsumStatShrinkRequest) SetQueryStartDate(v string) *AppliedConsumStatShrinkRequest {
	s.QueryStartDate = &v
	return s
}

type AppliedConsumStatResponseBody struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data *AppliedConsumStatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 业务处理消息摘要
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AppliedConsumStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppliedConsumStatResponseBody) GoString() string {
	return s.String()
}

func (s *AppliedConsumStatResponseBody) SetCode(v string) *AppliedConsumStatResponseBody {
	s.Code = &v
	return s
}

func (s *AppliedConsumStatResponseBody) SetData(v *AppliedConsumStatResponseBodyData) *AppliedConsumStatResponseBody {
	s.Data = v
	return s
}

func (s *AppliedConsumStatResponseBody) SetMessage(v string) *AppliedConsumStatResponseBody {
	s.Message = &v
	return s
}

func (s *AppliedConsumStatResponseBody) SetRequestId(v string) *AppliedConsumStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *AppliedConsumStatResponseBody) SetSuccess(v bool) *AppliedConsumStatResponseBody {
	s.Success = &v
	return s
}

type AppliedConsumStatResponseBodyData struct {
	// 应用消耗Cu统计
	AppliedConsumptionMap map[string][]*DataAppliedConsumptionMapValue `json:"AppliedConsumptionMap,omitempty" xml:"AppliedConsumptionMap,omitempty"`
}

func (s AppliedConsumStatResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AppliedConsumStatResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppliedConsumStatResponseBodyData) SetAppliedConsumptionMap(v map[string][]*DataAppliedConsumptionMapValue) *AppliedConsumStatResponseBodyData {
	s.AppliedConsumptionMap = v
	return s
}

type AppliedConsumStatResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppliedConsumStatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppliedConsumStatResponse) String() string {
	return tea.Prettify(s)
}

func (s AppliedConsumStatResponse) GoString() string {
	return s.String()
}

func (s *AppliedConsumStatResponse) SetHeaders(v map[string]*string) *AppliedConsumStatResponse {
	s.Headers = v
	return s
}

func (s *AppliedConsumStatResponse) SetBody(v *AppliedConsumStatResponseBody) *AppliedConsumStatResponse {
	s.Body = v
	return s
}

type AppliedNearRealStatRequest struct {
	AppliedVersionId []*string `json:"AppliedVersionId,omitempty" xml:"AppliedVersionId,omitempty" type:"Repeated"`
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// 排序类型。默认：AppliedConcurrency_Desc,AppliedNearRealOrderConditionType[AppliedConcurrency_Desc(AppliedConcurrency_Desc,根据实时并发路数降序排列),AppliedConcurrency_Asc(AppliedConcurrency_Asc,根据实时并发路数升序排列),AppliedConsumptionCu_Desc(AppliedConsumptionCu_Desc,根据实时CU消耗降序排列),AppliedConsumptionCu_Asc(AppliedConsumptionCu_Asc,根据实时CU消耗升序排列),orderByType,desc]
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// 资源类型,PackageType[CU(cu),code,cssResourceType,desc]
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// 当前页码，默认1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页项数，默认20,最大100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s AppliedNearRealStatRequest) String() string {
	return tea.Prettify(s)
}

func (s AppliedNearRealStatRequest) GoString() string {
	return s.String()
}

func (s *AppliedNearRealStatRequest) SetAppliedVersionId(v []*string) *AppliedNearRealStatRequest {
	s.AppliedVersionId = v
	return s
}

func (s *AppliedNearRealStatRequest) SetOperatorId(v string) *AppliedNearRealStatRequest {
	s.OperatorId = &v
	return s
}

func (s *AppliedNearRealStatRequest) SetOperatorType(v string) *AppliedNearRealStatRequest {
	s.OperatorType = &v
	return s
}

func (s *AppliedNearRealStatRequest) SetOrderBy(v string) *AppliedNearRealStatRequest {
	s.OrderBy = &v
	return s
}

func (s *AppliedNearRealStatRequest) SetPackageType(v string) *AppliedNearRealStatRequest {
	s.PackageType = &v
	return s
}

func (s *AppliedNearRealStatRequest) SetPageNumber(v int32) *AppliedNearRealStatRequest {
	s.PageNumber = &v
	return s
}

func (s *AppliedNearRealStatRequest) SetPageSize(v int32) *AppliedNearRealStatRequest {
	s.PageSize = &v
	return s
}

type AppliedNearRealStatShrinkRequest struct {
	AppliedVersionIdShrink *string `json:"AppliedVersionId,omitempty" xml:"AppliedVersionId,omitempty"`
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// 排序类型。默认：AppliedConcurrency_Desc,AppliedNearRealOrderConditionType[AppliedConcurrency_Desc(AppliedConcurrency_Desc,根据实时并发路数降序排列),AppliedConcurrency_Asc(AppliedConcurrency_Asc,根据实时并发路数升序排列),AppliedConsumptionCu_Desc(AppliedConsumptionCu_Desc,根据实时CU消耗降序排列),AppliedConsumptionCu_Asc(AppliedConsumptionCu_Asc,根据实时CU消耗升序排列),orderByType,desc]
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// 资源类型,PackageType[CU(cu),code,cssResourceType,desc]
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// 当前页码，默认1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页项数，默认20,最大100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s AppliedNearRealStatShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AppliedNearRealStatShrinkRequest) GoString() string {
	return s.String()
}

func (s *AppliedNearRealStatShrinkRequest) SetAppliedVersionIdShrink(v string) *AppliedNearRealStatShrinkRequest {
	s.AppliedVersionIdShrink = &v
	return s
}

func (s *AppliedNearRealStatShrinkRequest) SetOperatorId(v string) *AppliedNearRealStatShrinkRequest {
	s.OperatorId = &v
	return s
}

func (s *AppliedNearRealStatShrinkRequest) SetOperatorType(v string) *AppliedNearRealStatShrinkRequest {
	s.OperatorType = &v
	return s
}

func (s *AppliedNearRealStatShrinkRequest) SetOrderBy(v string) *AppliedNearRealStatShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *AppliedNearRealStatShrinkRequest) SetPackageType(v string) *AppliedNearRealStatShrinkRequest {
	s.PackageType = &v
	return s
}

func (s *AppliedNearRealStatShrinkRequest) SetPageNumber(v int32) *AppliedNearRealStatShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *AppliedNearRealStatShrinkRequest) SetPageSize(v int32) *AppliedNearRealStatShrinkRequest {
	s.PageSize = &v
	return s
}

type AppliedNearRealStatResponseBody struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data *AppliedNearRealStatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 业务处理消息摘要
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AppliedNearRealStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppliedNearRealStatResponseBody) GoString() string {
	return s.String()
}

func (s *AppliedNearRealStatResponseBody) SetCode(v string) *AppliedNearRealStatResponseBody {
	s.Code = &v
	return s
}

func (s *AppliedNearRealStatResponseBody) SetData(v *AppliedNearRealStatResponseBodyData) *AppliedNearRealStatResponseBody {
	s.Data = v
	return s
}

func (s *AppliedNearRealStatResponseBody) SetMessage(v string) *AppliedNearRealStatResponseBody {
	s.Message = &v
	return s
}

func (s *AppliedNearRealStatResponseBody) SetRequestId(v string) *AppliedNearRealStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *AppliedNearRealStatResponseBody) SetSuccess(v bool) *AppliedNearRealStatResponseBody {
	s.Success = &v
	return s
}

type AppliedNearRealStatResponseBodyData struct {
	// 当前页码，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页项数，默认20,最大100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总页数
	Pages *int64 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// 结果集
	Records []*AppliedNearRealStatResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// 总共项数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s AppliedNearRealStatResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AppliedNearRealStatResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppliedNearRealStatResponseBodyData) SetPageNumber(v int64) *AppliedNearRealStatResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *AppliedNearRealStatResponseBodyData) SetPageSize(v int64) *AppliedNearRealStatResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *AppliedNearRealStatResponseBodyData) SetPages(v int64) *AppliedNearRealStatResponseBodyData {
	s.Pages = &v
	return s
}

func (s *AppliedNearRealStatResponseBodyData) SetRecords(v []*AppliedNearRealStatResponseBodyDataRecords) *AppliedNearRealStatResponseBodyData {
	s.Records = v
	return s
}

func (s *AppliedNearRealStatResponseBodyData) SetTotalCount(v int64) *AppliedNearRealStatResponseBodyData {
	s.TotalCount = &v
	return s
}

type AppliedNearRealStatResponseBodyDataRecords struct {
	// 应用ID
	AppliedId *string `json:"AppliedId,omitempty" xml:"AppliedId,omitempty"`
	// 应用名称
	AppliedName *string `json:"AppliedName,omitempty" xml:"AppliedName,omitempty"`
	// 应用版本ID
	AppliedVersionId *string `json:"AppliedVersionId,omitempty" xml:"AppliedVersionId,omitempty"`
	// 应用版本名称
	AppliedVersionName *string `json:"AppliedVersionName,omitempty" xml:"AppliedVersionName,omitempty"`
	// 实时消耗并发
	Concurrency *int64 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// 实时消耗CU
	ConsumptionCu *float64 `json:"ConsumptionCu,omitempty" xml:"ConsumptionCu,omitempty"`
}

func (s AppliedNearRealStatResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s AppliedNearRealStatResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *AppliedNearRealStatResponseBodyDataRecords) SetAppliedId(v string) *AppliedNearRealStatResponseBodyDataRecords {
	s.AppliedId = &v
	return s
}

func (s *AppliedNearRealStatResponseBodyDataRecords) SetAppliedName(v string) *AppliedNearRealStatResponseBodyDataRecords {
	s.AppliedName = &v
	return s
}

func (s *AppliedNearRealStatResponseBodyDataRecords) SetAppliedVersionId(v string) *AppliedNearRealStatResponseBodyDataRecords {
	s.AppliedVersionId = &v
	return s
}

func (s *AppliedNearRealStatResponseBodyDataRecords) SetAppliedVersionName(v string) *AppliedNearRealStatResponseBodyDataRecords {
	s.AppliedVersionName = &v
	return s
}

func (s *AppliedNearRealStatResponseBodyDataRecords) SetConcurrency(v int64) *AppliedNearRealStatResponseBodyDataRecords {
	s.Concurrency = &v
	return s
}

func (s *AppliedNearRealStatResponseBodyDataRecords) SetConsumptionCu(v float64) *AppliedNearRealStatResponseBodyDataRecords {
	s.ConsumptionCu = &v
	return s
}

type AppliedNearRealStatResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppliedNearRealStatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppliedNearRealStatResponse) String() string {
	return tea.Prettify(s)
}

func (s AppliedNearRealStatResponse) GoString() string {
	return s.String()
}

func (s *AppliedNearRealStatResponse) SetHeaders(v map[string]*string) *AppliedNearRealStatResponse {
	s.Headers = v
	return s
}

func (s *AppliedNearRealStatResponse) SetBody(v *AppliedNearRealStatResponseBody) *AppliedNearRealStatResponse {
	s.Body = v
	return s
}

type AppliedStatRequest struct {
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// 查询结束时间
	QueryEndDate *string `json:"QueryEndDate,omitempty" xml:"QueryEndDate,omitempty"`
	// 查询开始时间
	QueryStartDate *string `json:"QueryStartDate,omitempty" xml:"QueryStartDate,omitempty"`
}

func (s AppliedStatRequest) String() string {
	return tea.Prettify(s)
}

func (s AppliedStatRequest) GoString() string {
	return s.String()
}

func (s *AppliedStatRequest) SetOperatorId(v string) *AppliedStatRequest {
	s.OperatorId = &v
	return s
}

func (s *AppliedStatRequest) SetOperatorType(v string) *AppliedStatRequest {
	s.OperatorType = &v
	return s
}

func (s *AppliedStatRequest) SetQueryEndDate(v string) *AppliedStatRequest {
	s.QueryEndDate = &v
	return s
}

func (s *AppliedStatRequest) SetQueryStartDate(v string) *AppliedStatRequest {
	s.QueryStartDate = &v
	return s
}

type AppliedStatResponseBody struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data *AppliedStatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 业务处理消息摘要
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AppliedStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppliedStatResponseBody) GoString() string {
	return s.String()
}

func (s *AppliedStatResponseBody) SetCode(v string) *AppliedStatResponseBody {
	s.Code = &v
	return s
}

func (s *AppliedStatResponseBody) SetData(v *AppliedStatResponseBodyData) *AppliedStatResponseBody {
	s.Data = v
	return s
}

func (s *AppliedStatResponseBody) SetMessage(v string) *AppliedStatResponseBody {
	s.Message = &v
	return s
}

func (s *AppliedStatResponseBody) SetRequestId(v string) *AppliedStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *AppliedStatResponseBody) SetSuccess(v bool) *AppliedStatResponseBody {
	s.Success = &v
	return s
}

type AppliedStatResponseBodyData struct {
	// 活跃应用个数
	ActiveApplications *int64 `json:"ActiveApplications,omitempty" xml:"ActiveApplications,omitempty"`
	// 日均应用运行时长
	AverageDailyRuntime *int64 `json:"AverageDailyRuntime,omitempty" xml:"AverageDailyRuntime,omitempty"`
	// 应用并发数量峰值
	PeakConcurrency *int64 `json:"PeakConcurrency,omitempty" xml:"PeakConcurrency,omitempty"`
	// 次均应用时长
	SecondaryAverageTime *int64 `json:"SecondaryAverageTime,omitempty" xml:"SecondaryAverageTime,omitempty"`
}

func (s AppliedStatResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AppliedStatResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppliedStatResponseBodyData) SetActiveApplications(v int64) *AppliedStatResponseBodyData {
	s.ActiveApplications = &v
	return s
}

func (s *AppliedStatResponseBodyData) SetAverageDailyRuntime(v int64) *AppliedStatResponseBodyData {
	s.AverageDailyRuntime = &v
	return s
}

func (s *AppliedStatResponseBodyData) SetPeakConcurrency(v int64) *AppliedStatResponseBodyData {
	s.PeakConcurrency = &v
	return s
}

func (s *AppliedStatResponseBodyData) SetSecondaryAverageTime(v int64) *AppliedStatResponseBodyData {
	s.SecondaryAverageTime = &v
	return s
}

type AppliedStatResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppliedStatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppliedStatResponse) String() string {
	return tea.Prettify(s)
}

func (s AppliedStatResponse) GoString() string {
	return s.String()
}

func (s *AppliedStatResponse) SetHeaders(v map[string]*string) *AppliedStatResponse {
	s.Headers = v
	return s
}

func (s *AppliedStatResponse) SetBody(v *AppliedStatResponseBody) *AppliedStatResponse {
	s.Body = v
	return s
}

type CreateAppSessionRequest struct {
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 客户端ip
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 自定义用户id
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	// 启动参数
	StartParameters []*CreateAppSessionRequestStartParameters `json:"StartParameters,omitempty" xml:"StartParameters,omitempty" type:"Repeated"`
	// 系统信息：如端侧机型等信息
	SystemInfo []*CreateAppSessionRequestSystemInfo `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty" type:"Repeated"`
}

func (s CreateAppSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateAppSessionRequest) SetAppId(v string) *CreateAppSessionRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppSessionRequest) SetAppVersion(v string) *CreateAppSessionRequest {
	s.AppVersion = &v
	return s
}

func (s *CreateAppSessionRequest) SetClientIp(v string) *CreateAppSessionRequest {
	s.ClientIp = &v
	return s
}

func (s *CreateAppSessionRequest) SetCustomSessionId(v string) *CreateAppSessionRequest {
	s.CustomSessionId = &v
	return s
}

func (s *CreateAppSessionRequest) SetCustomUserId(v string) *CreateAppSessionRequest {
	s.CustomUserId = &v
	return s
}

func (s *CreateAppSessionRequest) SetStartParameters(v []*CreateAppSessionRequestStartParameters) *CreateAppSessionRequest {
	s.StartParameters = v
	return s
}

func (s *CreateAppSessionRequest) SetSystemInfo(v []*CreateAppSessionRequestSystemInfo) *CreateAppSessionRequest {
	s.SystemInfo = v
	return s
}

type CreateAppSessionRequestStartParameters struct {
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionRequestStartParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionRequestStartParameters) GoString() string {
	return s.String()
}

func (s *CreateAppSessionRequestStartParameters) SetKey(v string) *CreateAppSessionRequestStartParameters {
	s.Key = &v
	return s
}

func (s *CreateAppSessionRequestStartParameters) SetValue(v string) *CreateAppSessionRequestStartParameters {
	s.Value = &v
	return s
}

type CreateAppSessionRequestSystemInfo struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionRequestSystemInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionRequestSystemInfo) GoString() string {
	return s.String()
}

func (s *CreateAppSessionRequestSystemInfo) SetKey(v string) *CreateAppSessionRequestSystemInfo {
	s.Key = &v
	return s
}

func (s *CreateAppSessionRequestSystemInfo) SetValue(v string) *CreateAppSessionRequestSystemInfo {
	s.Value = &v
	return s
}

type CreateAppSessionResponseBody struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppSessionResponseBody) SetAppId(v string) *CreateAppSessionResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateAppSessionResponseBody) SetAppVersion(v string) *CreateAppSessionResponseBody {
	s.AppVersion = &v
	return s
}

func (s *CreateAppSessionResponseBody) SetCustomSessionId(v string) *CreateAppSessionResponseBody {
	s.CustomSessionId = &v
	return s
}

func (s *CreateAppSessionResponseBody) SetPlatformSessionId(v string) *CreateAppSessionResponseBody {
	s.PlatformSessionId = &v
	return s
}

func (s *CreateAppSessionResponseBody) SetRequestId(v string) *CreateAppSessionResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppSessionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateAppSessionResponse) SetHeaders(v map[string]*string) *CreateAppSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateAppSessionResponse) SetBody(v *CreateAppSessionResponseBody) *CreateAppSessionResponse {
	s.Body = v
	return s
}

type CreateAppSessionBatchRequest struct {
	AppInfos     []*CreateAppSessionBatchRequestAppInfos `json:"AppInfos,omitempty" xml:"AppInfos,omitempty" type:"Repeated"`
	CustomTaskId *string                                 `json:"CustomTaskId,omitempty" xml:"CustomTaskId,omitempty"`
	Timeout      *int32                                  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateAppSessionBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchRequest) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchRequest) SetAppInfos(v []*CreateAppSessionBatchRequestAppInfos) *CreateAppSessionBatchRequest {
	s.AppInfos = v
	return s
}

func (s *CreateAppSessionBatchRequest) SetCustomTaskId(v string) *CreateAppSessionBatchRequest {
	s.CustomTaskId = &v
	return s
}

func (s *CreateAppSessionBatchRequest) SetTimeout(v int32) *CreateAppSessionBatchRequest {
	s.Timeout = &v
	return s
}

type CreateAppSessionBatchRequestAppInfos struct {
	AppId             *string                                                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersion        *string                                                `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	ClientIp          *string                                                `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	CustomUserId      *string                                                `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	CustomerSessionId *string                                                `json:"CustomerSessionId,omitempty" xml:"CustomerSessionId,omitempty"`
	StartParameters   []*CreateAppSessionBatchRequestAppInfosStartParameters `json:"StartParameters,omitempty" xml:"StartParameters,omitempty" type:"Repeated"`
	SystemInfo        []*CreateAppSessionBatchRequestAppInfosSystemInfo      `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty" type:"Repeated"`
}

func (s CreateAppSessionBatchRequestAppInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchRequestAppInfos) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchRequestAppInfos) SetAppId(v string) *CreateAppSessionBatchRequestAppInfos {
	s.AppId = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetAppVersion(v string) *CreateAppSessionBatchRequestAppInfos {
	s.AppVersion = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetClientIp(v string) *CreateAppSessionBatchRequestAppInfos {
	s.ClientIp = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetCustomUserId(v string) *CreateAppSessionBatchRequestAppInfos {
	s.CustomUserId = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetCustomerSessionId(v string) *CreateAppSessionBatchRequestAppInfos {
	s.CustomerSessionId = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetStartParameters(v []*CreateAppSessionBatchRequestAppInfosStartParameters) *CreateAppSessionBatchRequestAppInfos {
	s.StartParameters = v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetSystemInfo(v []*CreateAppSessionBatchRequestAppInfosSystemInfo) *CreateAppSessionBatchRequestAppInfos {
	s.SystemInfo = v
	return s
}

type CreateAppSessionBatchRequestAppInfosStartParameters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionBatchRequestAppInfosStartParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchRequestAppInfosStartParameters) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchRequestAppInfosStartParameters) SetKey(v string) *CreateAppSessionBatchRequestAppInfosStartParameters {
	s.Key = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfosStartParameters) SetValue(v string) *CreateAppSessionBatchRequestAppInfosStartParameters {
	s.Value = &v
	return s
}

type CreateAppSessionBatchRequestAppInfosSystemInfo struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionBatchRequestAppInfosSystemInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchRequestAppInfosSystemInfo) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchRequestAppInfosSystemInfo) SetKey(v string) *CreateAppSessionBatchRequestAppInfosSystemInfo {
	s.Key = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfosSystemInfo) SetValue(v string) *CreateAppSessionBatchRequestAppInfosSystemInfo {
	s.Value = &v
	return s
}

type CreateAppSessionBatchShrinkRequest struct {
	AppInfosShrink *string `json:"AppInfos,omitempty" xml:"AppInfos,omitempty"`
	CustomTaskId   *string `json:"CustomTaskId,omitempty" xml:"CustomTaskId,omitempty"`
	Timeout        *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateAppSessionBatchShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchShrinkRequest) SetAppInfosShrink(v string) *CreateAppSessionBatchShrinkRequest {
	s.AppInfosShrink = &v
	return s
}

func (s *CreateAppSessionBatchShrinkRequest) SetCustomTaskId(v string) *CreateAppSessionBatchShrinkRequest {
	s.CustomTaskId = &v
	return s
}

func (s *CreateAppSessionBatchShrinkRequest) SetTimeout(v int32) *CreateAppSessionBatchShrinkRequest {
	s.Timeout = &v
	return s
}

type CreateAppSessionBatchResponseBody struct {
	// 自定义会话id
	CustomTaskId *string `json:"CustomTaskId,omitempty" xml:"CustomTaskId,omitempty"`
	// 平台会话id
	PlatformTaskId *string `json:"PlatformTaskId,omitempty" xml:"PlatformTaskId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppSessionBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchResponseBody) SetCustomTaskId(v string) *CreateAppSessionBatchResponseBody {
	s.CustomTaskId = &v
	return s
}

func (s *CreateAppSessionBatchResponseBody) SetPlatformTaskId(v string) *CreateAppSessionBatchResponseBody {
	s.PlatformTaskId = &v
	return s
}

func (s *CreateAppSessionBatchResponseBody) SetRequestId(v string) *CreateAppSessionBatchResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppSessionBatchResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppSessionBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppSessionBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchResponse) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchResponse) SetHeaders(v map[string]*string) *CreateAppSessionBatchResponse {
	s.Headers = v
	return s
}

func (s *CreateAppSessionBatchResponse) SetBody(v *CreateAppSessionBatchResponseBody) *CreateAppSessionBatchResponse {
	s.Body = v
	return s
}

type CreateUploadTaskRequest struct {
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用类型
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// 上传的bucket名称
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// 环境
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// 游戏链接
	FileAddress *string `json:"FileAddress,omitempty" xml:"FileAddress,omitempty"`
	// 文件大小
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// 上传文件类型
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// 上传进度
	Progress *float64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// 上传的bucket所在region
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 上传状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 上传工具版本
	UploadToolVersion *string `json:"UploadToolVersion,omitempty" xml:"UploadToolVersion,omitempty"`
	// 上传任务类型
	UploadType *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
	// 版本ID
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateUploadTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadTaskRequest) SetAppId(v string) *CreateUploadTaskRequest {
	s.AppId = &v
	return s
}

func (s *CreateUploadTaskRequest) SetAppType(v string) *CreateUploadTaskRequest {
	s.AppType = &v
	return s
}

func (s *CreateUploadTaskRequest) SetBucketName(v string) *CreateUploadTaskRequest {
	s.BucketName = &v
	return s
}

func (s *CreateUploadTaskRequest) SetEnv(v string) *CreateUploadTaskRequest {
	s.Env = &v
	return s
}

func (s *CreateUploadTaskRequest) SetFileAddress(v string) *CreateUploadTaskRequest {
	s.FileAddress = &v
	return s
}

func (s *CreateUploadTaskRequest) SetFileSize(v int64) *CreateUploadTaskRequest {
	s.FileSize = &v
	return s
}

func (s *CreateUploadTaskRequest) SetFileType(v string) *CreateUploadTaskRequest {
	s.FileType = &v
	return s
}

func (s *CreateUploadTaskRequest) SetProgress(v float64) *CreateUploadTaskRequest {
	s.Progress = &v
	return s
}

func (s *CreateUploadTaskRequest) SetRegion(v string) *CreateUploadTaskRequest {
	s.Region = &v
	return s
}

func (s *CreateUploadTaskRequest) SetStatus(v string) *CreateUploadTaskRequest {
	s.Status = &v
	return s
}

func (s *CreateUploadTaskRequest) SetUploadToolVersion(v string) *CreateUploadTaskRequest {
	s.UploadToolVersion = &v
	return s
}

func (s *CreateUploadTaskRequest) SetUploadType(v string) *CreateUploadTaskRequest {
	s.UploadType = &v
	return s
}

func (s *CreateUploadTaskRequest) SetVersionId(v string) *CreateUploadTaskRequest {
	s.VersionId = &v
	return s
}

type CreateUploadTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUploadTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadTaskResponseBody) SetCode(v string) *CreateUploadTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUploadTaskResponseBody) SetData(v bool) *CreateUploadTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateUploadTaskResponseBody) SetMessage(v string) *CreateUploadTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUploadTaskResponseBody) SetRequestId(v string) *CreateUploadTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateUploadTaskResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUploadTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUploadTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateUploadTaskResponse) SetHeaders(v map[string]*string) *CreateUploadTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateUploadTaskResponse) SetBody(v *CreateUploadTaskResponseBody) *CreateUploadTaskResponse {
	s.Body = v
	return s
}

type GetAppListResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetAppListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppListResponseBody) SetCode(v string) *GetAppListResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppListResponseBody) SetData(v []*GetAppListResponseBodyData) *GetAppListResponseBody {
	s.Data = v
	return s
}

func (s *GetAppListResponseBody) SetMessage(v string) *GetAppListResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppListResponseBody) SetRequestId(v string) *GetAppListResponseBody {
	s.RequestId = &v
	return s
}

type GetAppListResponseBodyData struct {
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s GetAppListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAppListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppListResponseBodyData) SetAppId(v string) *GetAppListResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetAppListResponseBodyData) SetAppName(v string) *GetAppListResponseBodyData {
	s.AppName = &v
	return s
}

type GetAppListResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAppListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppListResponse) GoString() string {
	return s.String()
}

func (s *GetAppListResponse) SetHeaders(v map[string]*string) *GetAppListResponse {
	s.Headers = v
	return s
}

func (s *GetAppListResponse) SetBody(v *GetAppListResponseBody) *GetAppListResponse {
	s.Body = v
	return s
}

type GetAppSessionRequest struct {
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
}

func (s GetAppSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppSessionRequest) GoString() string {
	return s.String()
}

func (s *GetAppSessionRequest) SetCustomSessionId(v string) *GetAppSessionRequest {
	s.CustomSessionId = &v
	return s
}

func (s *GetAppSessionRequest) SetPlatformSessionId(v string) *GetAppSessionRequest {
	s.PlatformSessionId = &v
	return s
}

type GetAppSessionResponseBody struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// 请求id
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScheduleInfo []*GetAppSessionResponseBodyScheduleInfo `json:"ScheduleInfo,omitempty" xml:"ScheduleInfo,omitempty" type:"Repeated"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAppSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppSessionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppSessionResponseBody) SetAppId(v string) *GetAppSessionResponseBody {
	s.AppId = &v
	return s
}

func (s *GetAppSessionResponseBody) SetAppVersion(v string) *GetAppSessionResponseBody {
	s.AppVersion = &v
	return s
}

func (s *GetAppSessionResponseBody) SetCustomSessionId(v string) *GetAppSessionResponseBody {
	s.CustomSessionId = &v
	return s
}

func (s *GetAppSessionResponseBody) SetPlatformSessionId(v string) *GetAppSessionResponseBody {
	s.PlatformSessionId = &v
	return s
}

func (s *GetAppSessionResponseBody) SetRequestId(v string) *GetAppSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppSessionResponseBody) SetScheduleInfo(v []*GetAppSessionResponseBodyScheduleInfo) *GetAppSessionResponseBody {
	s.ScheduleInfo = v
	return s
}

func (s *GetAppSessionResponseBody) SetStatus(v string) *GetAppSessionResponseBody {
	s.Status = &v
	return s
}

type GetAppSessionResponseBodyScheduleInfo struct {
	// key数值，枚举有多个数值，例如： RegionId 大区id ServerIp 服务端 IP ServerPort 端口
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAppSessionResponseBodyScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAppSessionResponseBodyScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetAppSessionResponseBodyScheduleInfo) SetKey(v string) *GetAppSessionResponseBodyScheduleInfo {
	s.Key = &v
	return s
}

func (s *GetAppSessionResponseBodyScheduleInfo) SetValue(v string) *GetAppSessionResponseBodyScheduleInfo {
	s.Value = &v
	return s
}

type GetAppSessionResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAppSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppSessionResponse) GoString() string {
	return s.String()
}

func (s *GetAppSessionResponse) SetHeaders(v map[string]*string) *GetAppSessionResponse {
	s.Headers = v
	return s
}

func (s *GetAppSessionResponse) SetBody(v *GetAppSessionResponseBody) *GetAppSessionResponse {
	s.Body = v
	return s
}

type GetNeedUploadFileListRequest struct {
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 环境
	Env      *string   `json:"Env,omitempty" xml:"Env,omitempty"`
	HashList []*string `json:"HashList,omitempty" xml:"HashList,omitempty" type:"Repeated"`
	// 版本ID
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetNeedUploadFileListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNeedUploadFileListRequest) GoString() string {
	return s.String()
}

func (s *GetNeedUploadFileListRequest) SetAppId(v string) *GetNeedUploadFileListRequest {
	s.AppId = &v
	return s
}

func (s *GetNeedUploadFileListRequest) SetEnv(v string) *GetNeedUploadFileListRequest {
	s.Env = &v
	return s
}

func (s *GetNeedUploadFileListRequest) SetHashList(v []*string) *GetNeedUploadFileListRequest {
	s.HashList = v
	return s
}

func (s *GetNeedUploadFileListRequest) SetVersionId(v string) *GetNeedUploadFileListRequest {
	s.VersionId = &v
	return s
}

type GetNeedUploadFileListResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetNeedUploadFileListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNeedUploadFileListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNeedUploadFileListResponseBody) GoString() string {
	return s.String()
}

func (s *GetNeedUploadFileListResponseBody) SetCode(v string) *GetNeedUploadFileListResponseBody {
	s.Code = &v
	return s
}

func (s *GetNeedUploadFileListResponseBody) SetData(v *GetNeedUploadFileListResponseBodyData) *GetNeedUploadFileListResponseBody {
	s.Data = v
	return s
}

func (s *GetNeedUploadFileListResponseBody) SetMessage(v string) *GetNeedUploadFileListResponseBody {
	s.Message = &v
	return s
}

func (s *GetNeedUploadFileListResponseBody) SetRequestId(v string) *GetNeedUploadFileListResponseBody {
	s.RequestId = &v
	return s
}

type GetNeedUploadFileListResponseBodyData struct {
	// 错误信息
	Err *string `json:"Err,omitempty" xml:"Err,omitempty"`
	// 待上传文件列表
	NeedUploadFileList []*string `json:"NeedUploadFileList,omitempty" xml:"NeedUploadFileList,omitempty" type:"Repeated"`
	// 请求结果
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNeedUploadFileListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetNeedUploadFileListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNeedUploadFileListResponseBodyData) SetErr(v string) *GetNeedUploadFileListResponseBodyData {
	s.Err = &v
	return s
}

func (s *GetNeedUploadFileListResponseBodyData) SetNeedUploadFileList(v []*string) *GetNeedUploadFileListResponseBodyData {
	s.NeedUploadFileList = v
	return s
}

func (s *GetNeedUploadFileListResponseBodyData) SetSuccess(v bool) *GetNeedUploadFileListResponseBodyData {
	s.Success = &v
	return s
}

type GetNeedUploadFileListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetNeedUploadFileListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNeedUploadFileListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNeedUploadFileListResponse) GoString() string {
	return s.String()
}

func (s *GetNeedUploadFileListResponse) SetHeaders(v map[string]*string) *GetNeedUploadFileListResponse {
	s.Headers = v
	return s
}

func (s *GetNeedUploadFileListResponse) SetBody(v *GetNeedUploadFileListResponseBody) *GetNeedUploadFileListResponse {
	s.Body = v
	return s
}

type GetOssInfoResponseBody struct {
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetOssInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOssInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOssInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssInfoResponseBody) SetCode(v string) *GetOssInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetOssInfoResponseBody) SetData(v *GetOssInfoResponseBodyData) *GetOssInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetOssInfoResponseBody) SetMessage(v string) *GetOssInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetOssInfoResponseBody) SetRequestId(v string) *GetOssInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetOssInfoResponseBodyData struct {
	First  *string `json:"First,omitempty" xml:"First,omitempty"`
	Second *string `json:"Second,omitempty" xml:"Second,omitempty"`
}

func (s GetOssInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOssInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOssInfoResponseBodyData) SetFirst(v string) *GetOssInfoResponseBodyData {
	s.First = &v
	return s
}

func (s *GetOssInfoResponseBodyData) SetSecond(v string) *GetOssInfoResponseBodyData {
	s.Second = &v
	return s
}

type GetOssInfoResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOssInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOssInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOssInfoResponse) GoString() string {
	return s.String()
}

func (s *GetOssInfoResponse) SetHeaders(v map[string]*string) *GetOssInfoResponse {
	s.Headers = v
	return s
}

func (s *GetOssInfoResponse) SetBody(v *GetOssInfoResponseBody) *GetOssInfoResponse {
	s.Body = v
	return s
}

type GetTenantIdResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTenantIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTenantIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTenantIdResponseBody) SetCode(v string) *GetTenantIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetTenantIdResponseBody) SetData(v int64) *GetTenantIdResponseBody {
	s.Data = &v
	return s
}

func (s *GetTenantIdResponseBody) SetMessage(v string) *GetTenantIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetTenantIdResponseBody) SetRequestId(v string) *GetTenantIdResponseBody {
	s.RequestId = &v
	return s
}

type GetTenantIdResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTenantIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTenantIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTenantIdResponse) GoString() string {
	return s.String()
}

func (s *GetTenantIdResponse) SetHeaders(v map[string]*string) *GetTenantIdResponse {
	s.Headers = v
	return s
}

func (s *GetTenantIdResponse) SetBody(v *GetTenantIdResponseBody) *GetTenantIdResponse {
	s.Body = v
	return s
}

type GetTokenRequest struct {
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 存储桶
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// 环境
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// 区域ID
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 版本ID
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) SetAppId(v string) *GetTokenRequest {
	s.AppId = &v
	return s
}

func (s *GetTokenRequest) SetBucket(v string) *GetTokenRequest {
	s.Bucket = &v
	return s
}

func (s *GetTokenRequest) SetEnv(v string) *GetTokenRequest {
	s.Env = &v
	return s
}

func (s *GetTokenRequest) SetRegion(v string) *GetTokenRequest {
	s.Region = &v
	return s
}

func (s *GetTokenRequest) SetVersionId(v string) *GetTokenRequest {
	s.VersionId = &v
	return s
}

type GetTokenResponseBody struct {
	Code      *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) SetCode(v string) *GetTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetTokenResponseBody) SetData(v *GetTokenResponseBodyData) *GetTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetTokenResponseBody) SetMessage(v string) *GetTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

type GetTokenResponseBodyData struct {
	AccessKeyId      *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret  *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Endpoint         *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	Expiration       *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	InternalEndpoint *string `json:"InternalEndpoint,omitempty" xml:"InternalEndpoint,omitempty"`
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBodyData) SetAccessKeyId(v string) *GetTokenResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetTokenResponseBodyData) SetAccessKeySecret(v string) *GetTokenResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *GetTokenResponseBodyData) SetEndpoint(v string) *GetTokenResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetTokenResponseBodyData) SetExpiration(v string) *GetTokenResponseBodyData {
	s.Expiration = &v
	return s
}

func (s *GetTokenResponseBodyData) SetInternalEndpoint(v string) *GetTokenResponseBodyData {
	s.InternalEndpoint = &v
	return s
}

func (s *GetTokenResponseBodyData) SetSecurityToken(v string) *GetTokenResponseBodyData {
	s.SecurityToken = &v
	return s
}

type GetTokenResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTokenResponse) SetHeaders(v map[string]*string) *GetTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTokenResponse) SetBody(v *GetTokenResponseBody) *GetTokenResponse {
	s.Body = v
	return s
}

type GetUploadToolUrlResponseBody struct {
	Code      *string            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      map[string]*string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUploadToolUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUploadToolUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadToolUrlResponseBody) SetCode(v string) *GetUploadToolUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetUploadToolUrlResponseBody) SetData(v map[string]*string) *GetUploadToolUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetUploadToolUrlResponseBody) SetMessage(v string) *GetUploadToolUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetUploadToolUrlResponseBody) SetRequestId(v string) *GetUploadToolUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetUploadToolUrlResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUploadToolUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUploadToolUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUploadToolUrlResponse) GoString() string {
	return s.String()
}

func (s *GetUploadToolUrlResponse) SetHeaders(v map[string]*string) *GetUploadToolUrlResponse {
	s.Headers = v
	return s
}

func (s *GetUploadToolUrlResponse) SetBody(v *GetUploadToolUrlResponseBody) *GetUploadToolUrlResponse {
	s.Body = v
	return s
}

type HasActivateResponseBody struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data *HasActivateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 业务处理消息摘要
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HasActivateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HasActivateResponseBody) GoString() string {
	return s.String()
}

func (s *HasActivateResponseBody) SetCode(v string) *HasActivateResponseBody {
	s.Code = &v
	return s
}

func (s *HasActivateResponseBody) SetData(v *HasActivateResponseBodyData) *HasActivateResponseBody {
	s.Data = v
	return s
}

func (s *HasActivateResponseBody) SetMessage(v string) *HasActivateResponseBody {
	s.Message = &v
	return s
}

func (s *HasActivateResponseBody) SetRequestId(v string) *HasActivateResponseBody {
	s.RequestId = &v
	return s
}

func (s *HasActivateResponseBody) SetSuccess(v bool) *HasActivateResponseBody {
	s.Success = &v
	return s
}

type HasActivateResponseBodyData struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HasActivateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s HasActivateResponseBodyData) GoString() string {
	return s.String()
}

func (s *HasActivateResponseBodyData) SetSuccess(v bool) *HasActivateResponseBodyData {
	s.Success = &v
	return s
}

type HasActivateResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HasActivateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HasActivateResponse) String() string {
	return tea.Prettify(s)
}

func (s HasActivateResponse) GoString() string {
	return s.String()
}

func (s *HasActivateResponse) SetHeaders(v map[string]*string) *HasActivateResponse {
	s.Headers = v
	return s
}

func (s *HasActivateResponse) SetBody(v *HasActivateResponseBody) *HasActivateResponse {
	s.Body = v
	return s
}

type ListAppSessionsRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 自定义会话id
	CustomSessionIds []*string `json:"CustomSessionIds,omitempty" xml:"CustomSessionIds,omitempty" type:"Repeated"`
	// 页码
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 自定义用户id
	PlatformSessionIds []*string `json:"PlatformSessionIds,omitempty" xml:"PlatformSessionIds,omitempty" type:"Repeated"`
}

func (s ListAppSessionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListAppSessionsRequest) SetAppId(v string) *ListAppSessionsRequest {
	s.AppId = &v
	return s
}

func (s *ListAppSessionsRequest) SetCustomSessionIds(v []*string) *ListAppSessionsRequest {
	s.CustomSessionIds = v
	return s
}

func (s *ListAppSessionsRequest) SetPageNumber(v int32) *ListAppSessionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppSessionsRequest) SetPageSize(v int32) *ListAppSessionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppSessionsRequest) SetPlatformSessionIds(v []*string) *ListAppSessionsRequest {
	s.PlatformSessionIds = v
	return s
}

type ListAppSessionsResponseBody struct {
	AppSessions []*ListAppSessionsResponseBodyAppSessions `json:"AppSessions,omitempty" xml:"AppSessions,omitempty" type:"Repeated"`
	PageNumber  *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAppSessionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppSessionsResponseBody) SetAppSessions(v []*ListAppSessionsResponseBodyAppSessions) *ListAppSessionsResponseBody {
	s.AppSessions = v
	return s
}

func (s *ListAppSessionsResponseBody) SetPageNumber(v int32) *ListAppSessionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppSessionsResponseBody) SetPageSize(v int32) *ListAppSessionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppSessionsResponseBody) SetRequestId(v string) *ListAppSessionsResponseBody {
	s.RequestId = &v
	return s
}

type ListAppSessionsResponseBodyAppSessions struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAppSessionsResponseBodyAppSessions) String() string {
	return tea.Prettify(s)
}

func (s ListAppSessionsResponseBodyAppSessions) GoString() string {
	return s.String()
}

func (s *ListAppSessionsResponseBodyAppSessions) SetAppId(v string) *ListAppSessionsResponseBodyAppSessions {
	s.AppId = &v
	return s
}

func (s *ListAppSessionsResponseBodyAppSessions) SetAppVersion(v string) *ListAppSessionsResponseBodyAppSessions {
	s.AppVersion = &v
	return s
}

func (s *ListAppSessionsResponseBodyAppSessions) SetCustomSessionId(v string) *ListAppSessionsResponseBodyAppSessions {
	s.CustomSessionId = &v
	return s
}

func (s *ListAppSessionsResponseBodyAppSessions) SetPlatformSessionId(v string) *ListAppSessionsResponseBodyAppSessions {
	s.PlatformSessionId = &v
	return s
}

func (s *ListAppSessionsResponseBodyAppSessions) SetStatus(v string) *ListAppSessionsResponseBodyAppSessions {
	s.Status = &v
	return s
}

type ListAppSessionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAppSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppSessionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppSessionsResponse) GoString() string {
	return s.String()
}

func (s *ListAppSessionsResponse) SetHeaders(v map[string]*string) *ListAppSessionsResponse {
	s.Headers = v
	return s
}

func (s *ListAppSessionsResponse) SetBody(v *ListAppSessionsResponseBody) *ListAppSessionsResponse {
	s.Body = v
	return s
}

type PageQueryResourcePackageListRequest struct {
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// 资源包类型,PackageType[CU(cu),code,cssResourceType,desc]
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// 当前页码，默认1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页项数，默认20,最大100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 查询过期的资源包类型,ResourcePackageValidQueryConditionType[All(查询所有资源包),CurrentlyValid(查询当前有效的资源包(已开始，未结束)),PendingValid(未开始,即将生效的资源包),AllValid(已开始未结束 + 即将开始 的资源包),PendingInvalid5m(5min内即将到期的资源包),HasInvalid(已经过期的资源包),queryType,desc]
	QueryValidType *string `json:"QueryValidType,omitempty" xml:"QueryValidType,omitempty"`
}

func (s PageQueryResourcePackageListRequest) String() string {
	return tea.Prettify(s)
}

func (s PageQueryResourcePackageListRequest) GoString() string {
	return s.String()
}

func (s *PageQueryResourcePackageListRequest) SetOperatorId(v string) *PageQueryResourcePackageListRequest {
	s.OperatorId = &v
	return s
}

func (s *PageQueryResourcePackageListRequest) SetOperatorType(v string) *PageQueryResourcePackageListRequest {
	s.OperatorType = &v
	return s
}

func (s *PageQueryResourcePackageListRequest) SetPackageType(v string) *PageQueryResourcePackageListRequest {
	s.PackageType = &v
	return s
}

func (s *PageQueryResourcePackageListRequest) SetPageNumber(v int32) *PageQueryResourcePackageListRequest {
	s.PageNumber = &v
	return s
}

func (s *PageQueryResourcePackageListRequest) SetPageSize(v int32) *PageQueryResourcePackageListRequest {
	s.PageSize = &v
	return s
}

func (s *PageQueryResourcePackageListRequest) SetQueryValidType(v string) *PageQueryResourcePackageListRequest {
	s.QueryValidType = &v
	return s
}

type PageQueryResourcePackageListResponseBody struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data *PageQueryResourcePackageListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 业务处理消息摘要
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PageQueryResourcePackageListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageQueryResourcePackageListResponseBody) GoString() string {
	return s.String()
}

func (s *PageQueryResourcePackageListResponseBody) SetCode(v string) *PageQueryResourcePackageListResponseBody {
	s.Code = &v
	return s
}

func (s *PageQueryResourcePackageListResponseBody) SetData(v *PageQueryResourcePackageListResponseBodyData) *PageQueryResourcePackageListResponseBody {
	s.Data = v
	return s
}

func (s *PageQueryResourcePackageListResponseBody) SetMessage(v string) *PageQueryResourcePackageListResponseBody {
	s.Message = &v
	return s
}

func (s *PageQueryResourcePackageListResponseBody) SetRequestId(v string) *PageQueryResourcePackageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageQueryResourcePackageListResponseBody) SetSuccess(v bool) *PageQueryResourcePackageListResponseBody {
	s.Success = &v
	return s
}

type PageQueryResourcePackageListResponseBodyData struct {
	// 当前页码，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页项数，默认20,最大100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总页数
	Pages *int64 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// 结果集
	Records []*PageQueryResourcePackageListResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// 总共项数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s PageQueryResourcePackageListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PageQueryResourcePackageListResponseBodyData) GoString() string {
	return s.String()
}

func (s *PageQueryResourcePackageListResponseBodyData) SetPageNumber(v int64) *PageQueryResourcePackageListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *PageQueryResourcePackageListResponseBodyData) SetPageSize(v int64) *PageQueryResourcePackageListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *PageQueryResourcePackageListResponseBodyData) SetPages(v int64) *PageQueryResourcePackageListResponseBodyData {
	s.Pages = &v
	return s
}

func (s *PageQueryResourcePackageListResponseBodyData) SetRecords(v []*PageQueryResourcePackageListResponseBodyDataRecords) *PageQueryResourcePackageListResponseBodyData {
	s.Records = v
	return s
}

func (s *PageQueryResourcePackageListResponseBodyData) SetTotalCount(v int64) *PageQueryResourcePackageListResponseBodyData {
	s.TotalCount = &v
	return s
}

type PageQueryResourcePackageListResponseBodyDataRecords struct {
	// 当前资源包剩余总量
	CurrentAmount *int64 `json:"CurrentAmount,omitempty" xml:"CurrentAmount,omitempty"`
	// 资源包有效开始时间
	GmtValidBegin *string `json:"GmtValidBegin,omitempty" xml:"GmtValidBegin,omitempty"`
	// 资源包有效结束时间
	GmtValidEnd *string `json:"GmtValidEnd,omitempty" xml:"GmtValidEnd,omitempty"`
	// 当前资源包购买总量
	InitAmount *int64 `json:"InitAmount,omitempty" xml:"InitAmount,omitempty"`
	// 资源包实例ID
	PackageInstanceId *string `json:"PackageInstanceId,omitempty" xml:"PackageInstanceId,omitempty"`
	// 资源包类型
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
}

func (s PageQueryResourcePackageListResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s PageQueryResourcePackageListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *PageQueryResourcePackageListResponseBodyDataRecords) SetCurrentAmount(v int64) *PageQueryResourcePackageListResponseBodyDataRecords {
	s.CurrentAmount = &v
	return s
}

func (s *PageQueryResourcePackageListResponseBodyDataRecords) SetGmtValidBegin(v string) *PageQueryResourcePackageListResponseBodyDataRecords {
	s.GmtValidBegin = &v
	return s
}

func (s *PageQueryResourcePackageListResponseBodyDataRecords) SetGmtValidEnd(v string) *PageQueryResourcePackageListResponseBodyDataRecords {
	s.GmtValidEnd = &v
	return s
}

func (s *PageQueryResourcePackageListResponseBodyDataRecords) SetInitAmount(v int64) *PageQueryResourcePackageListResponseBodyDataRecords {
	s.InitAmount = &v
	return s
}

func (s *PageQueryResourcePackageListResponseBodyDataRecords) SetPackageInstanceId(v string) *PageQueryResourcePackageListResponseBodyDataRecords {
	s.PackageInstanceId = &v
	return s
}

func (s *PageQueryResourcePackageListResponseBodyDataRecords) SetPackageType(v string) *PageQueryResourcePackageListResponseBodyDataRecords {
	s.PackageType = &v
	return s
}

type PageQueryResourcePackageListResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PageQueryResourcePackageListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PageQueryResourcePackageListResponse) String() string {
	return tea.Prettify(s)
}

func (s PageQueryResourcePackageListResponse) GoString() string {
	return s.String()
}

func (s *PageQueryResourcePackageListResponse) SetHeaders(v map[string]*string) *PageQueryResourcePackageListResponse {
	s.Headers = v
	return s
}

func (s *PageQueryResourcePackageListResponse) SetBody(v *PageQueryResourcePackageListResponseBody) *PageQueryResourcePackageListResponse {
	s.Body = v
	return s
}

type QueryAdaptRecordsRequest struct {
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	RequestApp   *string `json:"RequestApp,omitempty" xml:"RequestApp,omitempty"`
}

func (s QueryAdaptRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAdaptRecordsRequest) GoString() string {
	return s.String()
}

func (s *QueryAdaptRecordsRequest) SetAppVersionId(v string) *QueryAdaptRecordsRequest {
	s.AppVersionId = &v
	return s
}

func (s *QueryAdaptRecordsRequest) SetRequestApp(v string) *QueryAdaptRecordsRequest {
	s.RequestApp = &v
	return s
}

type QueryAdaptRecordsResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryAdaptRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAdaptRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAdaptRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAdaptRecordsResponseBody) SetCode(v string) *QueryAdaptRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAdaptRecordsResponseBody) SetData(v *QueryAdaptRecordsResponseBodyData) *QueryAdaptRecordsResponseBody {
	s.Data = v
	return s
}

func (s *QueryAdaptRecordsResponseBody) SetMessage(v string) *QueryAdaptRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAdaptRecordsResponseBody) SetRequestId(v string) *QueryAdaptRecordsResponseBody {
	s.RequestId = &v
	return s
}

type QueryAdaptRecordsResponseBodyData struct {
	AdaptApplyId         *int64                                           `json:"AdaptApplyId,omitempty" xml:"AdaptApplyId,omitempty"`
	AdaptRecords         []*QueryAdaptRecordsResponseBodyDataAdaptRecords `json:"AdaptRecords,omitempty" xml:"AdaptRecords,omitempty" type:"Repeated"`
	AppId                *string                                          `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName              *string                                          `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType              *string                                          `json:"AppType,omitempty" xml:"AppType,omitempty"`
	AppVersionId         *string                                          `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	AppVersionName       *string                                          `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	AppVersionServiceype *string                                          `json:"AppVersionServiceype,omitempty" xml:"AppVersionServiceype,omitempty"`
	TenantId             *int64                                           `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TenantName           *string                                          `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s QueryAdaptRecordsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAdaptRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAdaptRecordsResponseBodyData) SetAdaptApplyId(v int64) *QueryAdaptRecordsResponseBodyData {
	s.AdaptApplyId = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyData) SetAdaptRecords(v []*QueryAdaptRecordsResponseBodyDataAdaptRecords) *QueryAdaptRecordsResponseBodyData {
	s.AdaptRecords = v
	return s
}

func (s *QueryAdaptRecordsResponseBodyData) SetAppId(v string) *QueryAdaptRecordsResponseBodyData {
	s.AppId = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyData) SetAppName(v string) *QueryAdaptRecordsResponseBodyData {
	s.AppName = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyData) SetAppType(v string) *QueryAdaptRecordsResponseBodyData {
	s.AppType = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyData) SetAppVersionId(v string) *QueryAdaptRecordsResponseBodyData {
	s.AppVersionId = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyData) SetAppVersionName(v string) *QueryAdaptRecordsResponseBodyData {
	s.AppVersionName = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyData) SetAppVersionServiceype(v string) *QueryAdaptRecordsResponseBodyData {
	s.AppVersionServiceype = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyData) SetTenantId(v int64) *QueryAdaptRecordsResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyData) SetTenantName(v string) *QueryAdaptRecordsResponseBodyData {
	s.TenantName = &v
	return s
}

type QueryAdaptRecordsResponseBodyDataAdaptRecords struct {
	AdaptApplyId              *int64                                                                  `json:"AdaptApplyId,omitempty" xml:"AdaptApplyId,omitempty"`
	AdaptRecordId             *int64                                                                  `json:"AdaptRecordId,omitempty" xml:"AdaptRecordId,omitempty"`
	AdaptStatus               *string                                                                 `json:"AdaptStatus,omitempty" xml:"AdaptStatus,omitempty"`
	AdaptTarget               *QueryAdaptRecordsResponseBodyDataAdaptRecordsAdaptTarget               `json:"AdaptTarget,omitempty" xml:"AdaptTarget,omitempty" type:"Struct"`
	AppId                     *string                                                                 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersionId              *string                                                                 `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	CalculationEvaluationInfo *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfo `json:"CalculationEvaluationInfo,omitempty" xml:"CalculationEvaluationInfo,omitempty" type:"Struct"`
	ConsumeCu                 *float64                                                                `json:"ConsumeCu,omitempty" xml:"ConsumeCu,omitempty"`
	// 蔚领：1 独占虚机，2 支持多开 (EXCLUSIVE: 独占虚机, SHARED: 支持多开)
	ContainerType    *string                                                  `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
	FileDownloadPath *string                                                  `json:"FileDownloadPath,omitempty" xml:"FileDownloadPath,omitempty"`
	GmtCreate        *string                                                  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *string                                                  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ImageType        *string                                                  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	IsMustSelect     *bool                                                    `json:"IsMustSelect,omitempty" xml:"IsMustSelect,omitempty"`
	Isv              *string                                                  `json:"Isv,omitempty" xml:"Isv,omitempty"`
	MaxConcurrency   *int32                                                   `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	Memo             *string                                                  `json:"Memo,omitempty" xml:"Memo,omitempty"`
	Priority         *int32                                                   `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ServerInfo       *QueryAdaptRecordsResponseBodyDataAdaptRecordsServerInfo `json:"ServerInfo,omitempty" xml:"ServerInfo,omitempty" type:"Struct"`
	TenantId         *int64                                                   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	VmType           *string                                                  `json:"VmType,omitempty" xml:"VmType,omitempty"`
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecords) String() string {
	return tea.Prettify(s)
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecords) GoString() string {
	return s.String()
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetAdaptApplyId(v int64) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.AdaptApplyId = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetAdaptRecordId(v int64) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.AdaptRecordId = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetAdaptStatus(v string) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.AdaptStatus = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetAdaptTarget(v *QueryAdaptRecordsResponseBodyDataAdaptRecordsAdaptTarget) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.AdaptTarget = v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetAppId(v string) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.AppId = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetAppVersionId(v string) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.AppVersionId = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetCalculationEvaluationInfo(v *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfo) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.CalculationEvaluationInfo = v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetConsumeCu(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.ConsumeCu = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetContainerType(v string) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.ContainerType = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetFileDownloadPath(v string) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.FileDownloadPath = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetGmtCreate(v string) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.GmtCreate = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetGmtModified(v string) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.GmtModified = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetImageType(v string) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.ImageType = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetIsMustSelect(v bool) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.IsMustSelect = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetIsv(v string) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.Isv = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetMaxConcurrency(v int32) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.MaxConcurrency = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetMemo(v string) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.Memo = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetPriority(v int32) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.Priority = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetServerInfo(v *QueryAdaptRecordsResponseBodyDataAdaptRecordsServerInfo) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.ServerInfo = v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetTenantId(v int64) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.TenantId = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecords) SetVmType(v string) *QueryAdaptRecordsResponseBodyDataAdaptRecords {
	s.VmType = &v
	return s
}

type QueryAdaptRecordsResponseBodyDataAdaptRecordsAdaptTarget struct {
	BitRate      *int32  `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	FrameRate    *int32  `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	Resolution   *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	StartProgram *string `json:"StartProgram,omitempty" xml:"StartProgram,omitempty"`
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecordsAdaptTarget) String() string {
	return tea.Prettify(s)
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecordsAdaptTarget) GoString() string {
	return s.String()
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsAdaptTarget) SetBitRate(v int32) *QueryAdaptRecordsResponseBodyDataAdaptRecordsAdaptTarget {
	s.BitRate = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsAdaptTarget) SetFrameRate(v int32) *QueryAdaptRecordsResponseBodyDataAdaptRecordsAdaptTarget {
	s.FrameRate = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsAdaptTarget) SetResolution(v string) *QueryAdaptRecordsResponseBodyDataAdaptRecordsAdaptTarget {
	s.Resolution = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsAdaptTarget) SetStartProgram(v string) *QueryAdaptRecordsResponseBodyDataAdaptRecordsAdaptTarget {
	s.StartProgram = &v
	return s
}

type QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfo struct {
	Cpu *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoCpu `json:"Cpu,omitempty" xml:"Cpu,omitempty" type:"Struct"`
	Gpu *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpu `json:"Gpu,omitempty" xml:"Gpu,omitempty" type:"Struct"`
	Mem *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoMem `json:"Mem,omitempty" xml:"Mem,omitempty" type:"Struct"`
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfo) GoString() string {
	return s.String()
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfo) SetCpu(v *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoCpu) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfo {
	s.Cpu = v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfo) SetGpu(v *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpu) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfo {
	s.Gpu = v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfo) SetMem(v *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoMem) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfo {
	s.Mem = v
	return s
}

type QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoCpu struct {
	Average       *float64 `json:"Average,omitempty" xml:"Average,omitempty"`
	Maximum       *float64 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	Minimum       *float64 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	NumberOfCores *float64 `json:"NumberOfCores,omitempty" xml:"NumberOfCores,omitempty"`
	Quantile80    *float64 `json:"Quantile80,omitempty" xml:"Quantile80,omitempty"`
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoCpu) String() string {
	return tea.Prettify(s)
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoCpu) GoString() string {
	return s.String()
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoCpu) SetAverage(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoCpu {
	s.Average = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoCpu) SetMaximum(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoCpu {
	s.Maximum = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoCpu) SetMinimum(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoCpu {
	s.Minimum = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoCpu) SetNumberOfCores(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoCpu {
	s.NumberOfCores = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoCpu) SetQuantile80(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoCpu {
	s.Quantile80 = &v
	return s
}

type QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpu struct {
	GpuUsedutilization *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuGpuUsedutilization `json:"GpuUsedutilization,omitempty" xml:"GpuUsedutilization,omitempty" type:"Struct"`
	MemUsedutilization *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuMemUsedutilization `json:"MemUsedutilization,omitempty" xml:"MemUsedutilization,omitempty" type:"Struct"`
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpu) String() string {
	return tea.Prettify(s)
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpu) GoString() string {
	return s.String()
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpu) SetGpuUsedutilization(v *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuGpuUsedutilization) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpu {
	s.GpuUsedutilization = v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpu) SetMemUsedutilization(v *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuMemUsedutilization) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpu {
	s.MemUsedutilization = v
	return s
}

type QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuGpuUsedutilization struct {
	Average       *float64 `json:"Average,omitempty" xml:"Average,omitempty"`
	Maximum       *float64 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	Minimum       *float64 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	NumberOfCores *float64 `json:"NumberOfCores,omitempty" xml:"NumberOfCores,omitempty"`
	Quantile80    *float64 `json:"Quantile80,omitempty" xml:"Quantile80,omitempty"`
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuGpuUsedutilization) String() string {
	return tea.Prettify(s)
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuGpuUsedutilization) GoString() string {
	return s.String()
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuGpuUsedutilization) SetAverage(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuGpuUsedutilization {
	s.Average = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuGpuUsedutilization) SetMaximum(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuGpuUsedutilization {
	s.Maximum = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuGpuUsedutilization) SetMinimum(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuGpuUsedutilization {
	s.Minimum = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuGpuUsedutilization) SetNumberOfCores(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuGpuUsedutilization {
	s.NumberOfCores = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuGpuUsedutilization) SetQuantile80(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuGpuUsedutilization {
	s.Quantile80 = &v
	return s
}

type QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuMemUsedutilization struct {
	Average    *float64 `json:"Average,omitempty" xml:"Average,omitempty"`
	Maximum    *float64 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	Minimum    *float64 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	Quantile80 *float64 `json:"Quantile80,omitempty" xml:"Quantile80,omitempty"`
	Total      *float64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuMemUsedutilization) String() string {
	return tea.Prettify(s)
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuMemUsedutilization) GoString() string {
	return s.String()
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuMemUsedutilization) SetAverage(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuMemUsedutilization {
	s.Average = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuMemUsedutilization) SetMaximum(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuMemUsedutilization {
	s.Maximum = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuMemUsedutilization) SetMinimum(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuMemUsedutilization {
	s.Minimum = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuMemUsedutilization) SetQuantile80(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuMemUsedutilization {
	s.Quantile80 = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuMemUsedutilization) SetTotal(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoGpuMemUsedutilization {
	s.Total = &v
	return s
}

type QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoMem struct {
	Average    *float64 `json:"Average,omitempty" xml:"Average,omitempty"`
	Maximum    *float64 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	Minimum    *float64 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	Quantile80 *float64 `json:"Quantile80,omitempty" xml:"Quantile80,omitempty"`
	Total      *float64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoMem) String() string {
	return tea.Prettify(s)
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoMem) GoString() string {
	return s.String()
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoMem) SetAverage(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoMem {
	s.Average = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoMem) SetMaximum(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoMem {
	s.Maximum = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoMem) SetMinimum(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoMem {
	s.Minimum = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoMem) SetQuantile80(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoMem {
	s.Quantile80 = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoMem) SetTotal(v float64) *QueryAdaptRecordsResponseBodyDataAdaptRecordsCalculationEvaluationInfoMem {
	s.Total = &v
	return s
}

type QueryAdaptRecordsResponseBodyDataAdaptRecordsServerInfo struct {
	CpuType *string `json:"CpuType,omitempty" xml:"CpuType,omitempty"`
	GpuType *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecordsServerInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryAdaptRecordsResponseBodyDataAdaptRecordsServerInfo) GoString() string {
	return s.String()
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsServerInfo) SetCpuType(v string) *QueryAdaptRecordsResponseBodyDataAdaptRecordsServerInfo {
	s.CpuType = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsServerInfo) SetGpuType(v string) *QueryAdaptRecordsResponseBodyDataAdaptRecordsServerInfo {
	s.GpuType = &v
	return s
}

func (s *QueryAdaptRecordsResponseBodyDataAdaptRecordsServerInfo) SetName(v string) *QueryAdaptRecordsResponseBodyDataAdaptRecordsServerInfo {
	s.Name = &v
	return s
}

type QueryAdaptRecordsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAdaptRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAdaptRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAdaptRecordsResponse) GoString() string {
	return s.String()
}

func (s *QueryAdaptRecordsResponse) SetHeaders(v map[string]*string) *QueryAdaptRecordsResponse {
	s.Headers = v
	return s
}

func (s *QueryAdaptRecordsResponse) SetBody(v *QueryAdaptRecordsResponseBody) *QueryAdaptRecordsResponse {
	s.Body = v
	return s
}

type QueryUploadProgressRequest struct {
	QueryUploadProgressRequests *string `json:"QueryUploadProgressRequests,omitempty" xml:"QueryUploadProgressRequests,omitempty"`
}

func (s QueryUploadProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUploadProgressRequest) GoString() string {
	return s.String()
}

func (s *QueryUploadProgressRequest) SetQueryUploadProgressRequests(v string) *QueryUploadProgressRequest {
	s.QueryUploadProgressRequests = &v
	return s
}

type QueryUploadProgressResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryUploadProgressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryUploadProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUploadProgressResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUploadProgressResponseBody) SetCode(v string) *QueryUploadProgressResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUploadProgressResponseBody) SetData(v *QueryUploadProgressResponseBodyData) *QueryUploadProgressResponseBody {
	s.Data = v
	return s
}

func (s *QueryUploadProgressResponseBody) SetMessage(v string) *QueryUploadProgressResponseBody {
	s.Message = &v
	return s
}

func (s *QueryUploadProgressResponseBody) SetRequestId(v string) *QueryUploadProgressResponseBody {
	s.RequestId = &v
	return s
}

type QueryUploadProgressResponseBodyData struct {
	// 查询结果
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 进度信息
	Content *QueryUploadProgressResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// 查询信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s QueryUploadProgressResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryUploadProgressResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryUploadProgressResponseBodyData) SetCode(v string) *QueryUploadProgressResponseBodyData {
	s.Code = &v
	return s
}

func (s *QueryUploadProgressResponseBodyData) SetContent(v *QueryUploadProgressResponseBodyDataContent) *QueryUploadProgressResponseBodyData {
	s.Content = v
	return s
}

func (s *QueryUploadProgressResponseBodyData) SetMessage(v string) *QueryUploadProgressResponseBodyData {
	s.Message = &v
	return s
}

type QueryUploadProgressResponseBodyDataContent struct {
	Versions []*QueryUploadProgressResponseBodyDataContentVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s QueryUploadProgressResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s QueryUploadProgressResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *QueryUploadProgressResponseBodyDataContent) SetVersions(v []*QueryUploadProgressResponseBodyDataContentVersions) *QueryUploadProgressResponseBodyDataContent {
	s.Versions = v
	return s
}

type QueryUploadProgressResponseBodyDataContentVersions struct {
	AppId     *string  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Progress  *float64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	TenantId  *int64   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	VersionId *string  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s QueryUploadProgressResponseBodyDataContentVersions) String() string {
	return tea.Prettify(s)
}

func (s QueryUploadProgressResponseBodyDataContentVersions) GoString() string {
	return s.String()
}

func (s *QueryUploadProgressResponseBodyDataContentVersions) SetAppId(v string) *QueryUploadProgressResponseBodyDataContentVersions {
	s.AppId = &v
	return s
}

func (s *QueryUploadProgressResponseBodyDataContentVersions) SetProgress(v float64) *QueryUploadProgressResponseBodyDataContentVersions {
	s.Progress = &v
	return s
}

func (s *QueryUploadProgressResponseBodyDataContentVersions) SetTenantId(v int64) *QueryUploadProgressResponseBodyDataContentVersions {
	s.TenantId = &v
	return s
}

func (s *QueryUploadProgressResponseBodyDataContentVersions) SetVersionId(v string) *QueryUploadProgressResponseBodyDataContentVersions {
	s.VersionId = &v
	return s
}

type QueryUploadProgressResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUploadProgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUploadProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUploadProgressResponse) GoString() string {
	return s.String()
}

func (s *QueryUploadProgressResponse) SetHeaders(v map[string]*string) *QueryUploadProgressResponse {
	s.Headers = v
	return s
}

func (s *QueryUploadProgressResponse) SetBody(v *QueryUploadProgressResponseBody) *QueryUploadProgressResponse {
	s.Body = v
	return s
}

type RecordFinishedFileRequest struct {
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 上传的bucket名称
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// 环境
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// 用于pop传入的文件指纹信息
	FileFingerprintDTOList []*RecordFinishedFileRequestFileFingerprintDTOList `json:"FileFingerprintDTOList,omitempty" xml:"FileFingerprintDTOList,omitempty" type:"Repeated"`
	// 文件大小
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// 上传文件类型
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// 上传的bucket所在region
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 上传工具版本
	ToolVersion *string `json:"ToolVersion,omitempty" xml:"ToolVersion,omitempty"`
	// 版本ID
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s RecordFinishedFileRequest) String() string {
	return tea.Prettify(s)
}

func (s RecordFinishedFileRequest) GoString() string {
	return s.String()
}

func (s *RecordFinishedFileRequest) SetAppId(v string) *RecordFinishedFileRequest {
	s.AppId = &v
	return s
}

func (s *RecordFinishedFileRequest) SetBucketName(v string) *RecordFinishedFileRequest {
	s.BucketName = &v
	return s
}

func (s *RecordFinishedFileRequest) SetEnv(v string) *RecordFinishedFileRequest {
	s.Env = &v
	return s
}

func (s *RecordFinishedFileRequest) SetFileFingerprintDTOList(v []*RecordFinishedFileRequestFileFingerprintDTOList) *RecordFinishedFileRequest {
	s.FileFingerprintDTOList = v
	return s
}

func (s *RecordFinishedFileRequest) SetFileSize(v int64) *RecordFinishedFileRequest {
	s.FileSize = &v
	return s
}

func (s *RecordFinishedFileRequest) SetFileType(v string) *RecordFinishedFileRequest {
	s.FileType = &v
	return s
}

func (s *RecordFinishedFileRequest) SetRegion(v string) *RecordFinishedFileRequest {
	s.Region = &v
	return s
}

func (s *RecordFinishedFileRequest) SetToolVersion(v string) *RecordFinishedFileRequest {
	s.ToolVersion = &v
	return s
}

func (s *RecordFinishedFileRequest) SetVersionId(v string) *RecordFinishedFileRequest {
	s.VersionId = &v
	return s
}

type RecordFinishedFileRequestFileFingerprintDTOList struct {
	// 文件hash
	FileHash *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	// 文件大小
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
}

func (s RecordFinishedFileRequestFileFingerprintDTOList) String() string {
	return tea.Prettify(s)
}

func (s RecordFinishedFileRequestFileFingerprintDTOList) GoString() string {
	return s.String()
}

func (s *RecordFinishedFileRequestFileFingerprintDTOList) SetFileHash(v string) *RecordFinishedFileRequestFileFingerprintDTOList {
	s.FileHash = &v
	return s
}

func (s *RecordFinishedFileRequestFileFingerprintDTOList) SetFileSize(v int64) *RecordFinishedFileRequestFileFingerprintDTOList {
	s.FileSize = &v
	return s
}

type RecordFinishedFileResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecordFinishedFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecordFinishedFileResponseBody) GoString() string {
	return s.String()
}

func (s *RecordFinishedFileResponseBody) SetCode(v string) *RecordFinishedFileResponseBody {
	s.Code = &v
	return s
}

func (s *RecordFinishedFileResponseBody) SetData(v bool) *RecordFinishedFileResponseBody {
	s.Data = &v
	return s
}

func (s *RecordFinishedFileResponseBody) SetMessage(v string) *RecordFinishedFileResponseBody {
	s.Message = &v
	return s
}

func (s *RecordFinishedFileResponseBody) SetRequestId(v string) *RecordFinishedFileResponseBody {
	s.RequestId = &v
	return s
}

type RecordFinishedFileResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecordFinishedFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecordFinishedFileResponse) String() string {
	return tea.Prettify(s)
}

func (s RecordFinishedFileResponse) GoString() string {
	return s.String()
}

func (s *RecordFinishedFileResponse) SetHeaders(v map[string]*string) *RecordFinishedFileResponse {
	s.Headers = v
	return s
}

func (s *RecordFinishedFileResponse) SetBody(v *RecordFinishedFileResponseBody) *RecordFinishedFileResponse {
	s.Body = v
	return s
}

type ReplicateVersionRequest struct {
	// 应用Id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 源头版本Id
	SourceVersionId *string `json:"SourceVersionId,omitempty" xml:"SourceVersionId,omitempty"`
	// 复制目标版本Id
	TargetVersionId *string `json:"TargetVersionId,omitempty" xml:"TargetVersionId,omitempty"`
	// 租户Id
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ReplicateVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplicateVersionRequest) GoString() string {
	return s.String()
}

func (s *ReplicateVersionRequest) SetAppId(v string) *ReplicateVersionRequest {
	s.AppId = &v
	return s
}

func (s *ReplicateVersionRequest) SetSourceVersionId(v string) *ReplicateVersionRequest {
	s.SourceVersionId = &v
	return s
}

func (s *ReplicateVersionRequest) SetTargetVersionId(v string) *ReplicateVersionRequest {
	s.TargetVersionId = &v
	return s
}

func (s *ReplicateVersionRequest) SetTenantId(v int64) *ReplicateVersionRequest {
	s.TenantId = &v
	return s
}

type ReplicateVersionResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ReplicateVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReplicateVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplicateVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ReplicateVersionResponseBody) SetCode(v string) *ReplicateVersionResponseBody {
	s.Code = &v
	return s
}

func (s *ReplicateVersionResponseBody) SetData(v *ReplicateVersionResponseBodyData) *ReplicateVersionResponseBody {
	s.Data = v
	return s
}

func (s *ReplicateVersionResponseBody) SetMessage(v string) *ReplicateVersionResponseBody {
	s.Message = &v
	return s
}

func (s *ReplicateVersionResponseBody) SetRequestId(v string) *ReplicateVersionResponseBody {
	s.RequestId = &v
	return s
}

type ReplicateVersionResponseBodyData struct {
	// 复制结果
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 复制结果信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ReplicateVersionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ReplicateVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReplicateVersionResponseBodyData) SetCode(v string) *ReplicateVersionResponseBodyData {
	s.Code = &v
	return s
}

func (s *ReplicateVersionResponseBodyData) SetMessage(v string) *ReplicateVersionResponseBodyData {
	s.Message = &v
	return s
}

type ReplicateVersionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReplicateVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReplicateVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplicateVersionResponse) GoString() string {
	return s.String()
}

func (s *ReplicateVersionResponse) SetHeaders(v map[string]*string) *ReplicateVersionResponse {
	s.Headers = v
	return s
}

func (s *ReplicateVersionResponse) SetBody(v *ReplicateVersionResponseBody) *ReplicateVersionResponse {
	s.Body = v
	return s
}

type ReportUploadProgressRequest struct {
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 环境
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// 上传进度
	Progress *float64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// 版本ID
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ReportUploadProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportUploadProgressRequest) GoString() string {
	return s.String()
}

func (s *ReportUploadProgressRequest) SetAppId(v string) *ReportUploadProgressRequest {
	s.AppId = &v
	return s
}

func (s *ReportUploadProgressRequest) SetEnv(v string) *ReportUploadProgressRequest {
	s.Env = &v
	return s
}

func (s *ReportUploadProgressRequest) SetProgress(v float64) *ReportUploadProgressRequest {
	s.Progress = &v
	return s
}

func (s *ReportUploadProgressRequest) SetVersionId(v string) *ReportUploadProgressRequest {
	s.VersionId = &v
	return s
}

type ReportUploadProgressResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportUploadProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportUploadProgressResponseBody) GoString() string {
	return s.String()
}

func (s *ReportUploadProgressResponseBody) SetCode(v string) *ReportUploadProgressResponseBody {
	s.Code = &v
	return s
}

func (s *ReportUploadProgressResponseBody) SetData(v bool) *ReportUploadProgressResponseBody {
	s.Data = &v
	return s
}

func (s *ReportUploadProgressResponseBody) SetMessage(v string) *ReportUploadProgressResponseBody {
	s.Message = &v
	return s
}

func (s *ReportUploadProgressResponseBody) SetRequestId(v string) *ReportUploadProgressResponseBody {
	s.RequestId = &v
	return s
}

type ReportUploadProgressResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReportUploadProgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReportUploadProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportUploadProgressResponse) GoString() string {
	return s.String()
}

func (s *ReportUploadProgressResponse) SetHeaders(v map[string]*string) *ReportUploadProgressResponse {
	s.Headers = v
	return s
}

func (s *ReportUploadProgressResponse) SetBody(v *ReportUploadProgressResponseBody) *ReportUploadProgressResponse {
	s.Body = v
	return s
}

type ReportUploadResultRequest struct {
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 上传的bucket名称
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// 环境
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// 用于pop传入的文件指纹信息
	FileFingerprintDTOList []*ReportUploadResultRequestFileFingerprintDTOList `json:"FileFingerprintDTOList,omitempty" xml:"FileFingerprintDTOList,omitempty" type:"Repeated"`
	// 文件大小
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// 上传文件类型
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// 上传的bucket所在region
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 上传工具版本
	ToolVersion *string `json:"ToolVersion,omitempty" xml:"ToolVersion,omitempty"`
	// 版本ID
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ReportUploadResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportUploadResultRequest) GoString() string {
	return s.String()
}

func (s *ReportUploadResultRequest) SetAppId(v string) *ReportUploadResultRequest {
	s.AppId = &v
	return s
}

func (s *ReportUploadResultRequest) SetBucketName(v string) *ReportUploadResultRequest {
	s.BucketName = &v
	return s
}

func (s *ReportUploadResultRequest) SetEnv(v string) *ReportUploadResultRequest {
	s.Env = &v
	return s
}

func (s *ReportUploadResultRequest) SetFileFingerprintDTOList(v []*ReportUploadResultRequestFileFingerprintDTOList) *ReportUploadResultRequest {
	s.FileFingerprintDTOList = v
	return s
}

func (s *ReportUploadResultRequest) SetFileSize(v int64) *ReportUploadResultRequest {
	s.FileSize = &v
	return s
}

func (s *ReportUploadResultRequest) SetFileType(v string) *ReportUploadResultRequest {
	s.FileType = &v
	return s
}

func (s *ReportUploadResultRequest) SetRegion(v string) *ReportUploadResultRequest {
	s.Region = &v
	return s
}

func (s *ReportUploadResultRequest) SetToolVersion(v string) *ReportUploadResultRequest {
	s.ToolVersion = &v
	return s
}

func (s *ReportUploadResultRequest) SetVersionId(v string) *ReportUploadResultRequest {
	s.VersionId = &v
	return s
}

type ReportUploadResultRequestFileFingerprintDTOList struct {
	// 文件hash
	FileHash *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	// 文件大小
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
}

func (s ReportUploadResultRequestFileFingerprintDTOList) String() string {
	return tea.Prettify(s)
}

func (s ReportUploadResultRequestFileFingerprintDTOList) GoString() string {
	return s.String()
}

func (s *ReportUploadResultRequestFileFingerprintDTOList) SetFileHash(v string) *ReportUploadResultRequestFileFingerprintDTOList {
	s.FileHash = &v
	return s
}

func (s *ReportUploadResultRequestFileFingerprintDTOList) SetFileSize(v int64) *ReportUploadResultRequestFileFingerprintDTOList {
	s.FileSize = &v
	return s
}

type ReportUploadResultResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportUploadResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportUploadResultResponseBody) GoString() string {
	return s.String()
}

func (s *ReportUploadResultResponseBody) SetCode(v string) *ReportUploadResultResponseBody {
	s.Code = &v
	return s
}

func (s *ReportUploadResultResponseBody) SetData(v bool) *ReportUploadResultResponseBody {
	s.Data = &v
	return s
}

func (s *ReportUploadResultResponseBody) SetMessage(v string) *ReportUploadResultResponseBody {
	s.Message = &v
	return s
}

func (s *ReportUploadResultResponseBody) SetRequestId(v string) *ReportUploadResultResponseBody {
	s.RequestId = &v
	return s
}

type ReportUploadResultResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReportUploadResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReportUploadResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportUploadResultResponse) GoString() string {
	return s.String()
}

func (s *ReportUploadResultResponse) SetHeaders(v map[string]*string) *ReportUploadResultResponse {
	s.Headers = v
	return s
}

func (s *ReportUploadResultResponse) SetBody(v *ReportUploadResultResponseBody) *ReportUploadResultResponse {
	s.Body = v
	return s
}

type ReportUploadStatusRequest struct {
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 环境
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// 备注信息
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// 上传状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 版本ID
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ReportUploadStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportUploadStatusRequest) GoString() string {
	return s.String()
}

func (s *ReportUploadStatusRequest) SetAppId(v string) *ReportUploadStatusRequest {
	s.AppId = &v
	return s
}

func (s *ReportUploadStatusRequest) SetEnv(v string) *ReportUploadStatusRequest {
	s.Env = &v
	return s
}

func (s *ReportUploadStatusRequest) SetMemo(v string) *ReportUploadStatusRequest {
	s.Memo = &v
	return s
}

func (s *ReportUploadStatusRequest) SetStatus(v string) *ReportUploadStatusRequest {
	s.Status = &v
	return s
}

func (s *ReportUploadStatusRequest) SetVersionId(v string) *ReportUploadStatusRequest {
	s.VersionId = &v
	return s
}

type ReportUploadStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportUploadStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportUploadStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ReportUploadStatusResponseBody) SetCode(v string) *ReportUploadStatusResponseBody {
	s.Code = &v
	return s
}

func (s *ReportUploadStatusResponseBody) SetData(v bool) *ReportUploadStatusResponseBody {
	s.Data = &v
	return s
}

func (s *ReportUploadStatusResponseBody) SetMessage(v string) *ReportUploadStatusResponseBody {
	s.Message = &v
	return s
}

func (s *ReportUploadStatusResponseBody) SetRequestId(v string) *ReportUploadStatusResponseBody {
	s.RequestId = &v
	return s
}

type ReportUploadStatusResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReportUploadStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReportUploadStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportUploadStatusResponse) GoString() string {
	return s.String()
}

func (s *ReportUploadStatusResponse) SetHeaders(v map[string]*string) *ReportUploadStatusResponse {
	s.Headers = v
	return s
}

func (s *ReportUploadStatusResponse) SetBody(v *ReportUploadStatusResponseBody) *ReportUploadStatusResponse {
	s.Body = v
	return s
}

type StopAppSessionRequest struct {
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 自定义用户id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
}

func (s StopAppSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionRequest) GoString() string {
	return s.String()
}

func (s *StopAppSessionRequest) SetCustomSessionId(v string) *StopAppSessionRequest {
	s.CustomSessionId = &v
	return s
}

func (s *StopAppSessionRequest) SetPlatformSessionId(v string) *StopAppSessionRequest {
	s.PlatformSessionId = &v
	return s
}

type StopAppSessionResponseBody struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAppSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionResponseBody) GoString() string {
	return s.String()
}

func (s *StopAppSessionResponseBody) SetAppId(v string) *StopAppSessionResponseBody {
	s.AppId = &v
	return s
}

func (s *StopAppSessionResponseBody) SetAppVersion(v string) *StopAppSessionResponseBody {
	s.AppVersion = &v
	return s
}

func (s *StopAppSessionResponseBody) SetCustomSessionId(v string) *StopAppSessionResponseBody {
	s.CustomSessionId = &v
	return s
}

func (s *StopAppSessionResponseBody) SetPlatformSessionId(v string) *StopAppSessionResponseBody {
	s.PlatformSessionId = &v
	return s
}

func (s *StopAppSessionResponseBody) SetRequestId(v string) *StopAppSessionResponseBody {
	s.RequestId = &v
	return s
}

type StopAppSessionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopAppSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopAppSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionResponse) GoString() string {
	return s.String()
}

func (s *StopAppSessionResponse) SetHeaders(v map[string]*string) *StopAppSessionResponse {
	s.Headers = v
	return s
}

func (s *StopAppSessionResponse) SetBody(v *StopAppSessionResponseBody) *StopAppSessionResponse {
	s.Body = v
	return s
}

type TotalAppliedConsumStatRequest struct {
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// 资源类型,PackageType[CU(cu),code,cssResourceType,desc]
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// 查询结束时间
	QueryEndDate *string `json:"QueryEndDate,omitempty" xml:"QueryEndDate,omitempty"`
	// 查询开始时间
	QueryStartDate *string `json:"QueryStartDate,omitempty" xml:"QueryStartDate,omitempty"`
}

func (s TotalAppliedConsumStatRequest) String() string {
	return tea.Prettify(s)
}

func (s TotalAppliedConsumStatRequest) GoString() string {
	return s.String()
}

func (s *TotalAppliedConsumStatRequest) SetOperatorId(v string) *TotalAppliedConsumStatRequest {
	s.OperatorId = &v
	return s
}

func (s *TotalAppliedConsumStatRequest) SetOperatorType(v string) *TotalAppliedConsumStatRequest {
	s.OperatorType = &v
	return s
}

func (s *TotalAppliedConsumStatRequest) SetPackageType(v string) *TotalAppliedConsumStatRequest {
	s.PackageType = &v
	return s
}

func (s *TotalAppliedConsumStatRequest) SetQueryEndDate(v string) *TotalAppliedConsumStatRequest {
	s.QueryEndDate = &v
	return s
}

func (s *TotalAppliedConsumStatRequest) SetQueryStartDate(v string) *TotalAppliedConsumStatRequest {
	s.QueryStartDate = &v
	return s
}

type TotalAppliedConsumStatResponseBody struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data []*TotalAppliedConsumStatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// 业务处理消息摘要
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TotalAppliedConsumStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TotalAppliedConsumStatResponseBody) GoString() string {
	return s.String()
}

func (s *TotalAppliedConsumStatResponseBody) SetCode(v string) *TotalAppliedConsumStatResponseBody {
	s.Code = &v
	return s
}

func (s *TotalAppliedConsumStatResponseBody) SetData(v []*TotalAppliedConsumStatResponseBodyData) *TotalAppliedConsumStatResponseBody {
	s.Data = v
	return s
}

func (s *TotalAppliedConsumStatResponseBody) SetMessage(v string) *TotalAppliedConsumStatResponseBody {
	s.Message = &v
	return s
}

func (s *TotalAppliedConsumStatResponseBody) SetRequestId(v string) *TotalAppliedConsumStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *TotalAppliedConsumStatResponseBody) SetSuccess(v bool) *TotalAppliedConsumStatResponseBody {
	s.Success = &v
	return s
}

type TotalAppliedConsumStatResponseBodyData struct {
	// 应用ID
	AppliedId *string `json:"AppliedId,omitempty" xml:"AppliedId,omitempty"`
	// 分钟级消耗CU
	ConsumptionCu *int64 `json:"ConsumptionCu,omitempty" xml:"ConsumptionCu,omitempty"`
	// 统计日期
	StatDate *string `json:"StatDate,omitempty" xml:"StatDate,omitempty"`
}

func (s TotalAppliedConsumStatResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TotalAppliedConsumStatResponseBodyData) GoString() string {
	return s.String()
}

func (s *TotalAppliedConsumStatResponseBodyData) SetAppliedId(v string) *TotalAppliedConsumStatResponseBodyData {
	s.AppliedId = &v
	return s
}

func (s *TotalAppliedConsumStatResponseBodyData) SetConsumptionCu(v int64) *TotalAppliedConsumStatResponseBodyData {
	s.ConsumptionCu = &v
	return s
}

func (s *TotalAppliedConsumStatResponseBodyData) SetStatDate(v string) *TotalAppliedConsumStatResponseBodyData {
	s.StatDate = &v
	return s
}

type TotalAppliedConsumStatResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TotalAppliedConsumStatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TotalAppliedConsumStatResponse) String() string {
	return tea.Prettify(s)
}

func (s TotalAppliedConsumStatResponse) GoString() string {
	return s.String()
}

func (s *TotalAppliedConsumStatResponse) SetHeaders(v map[string]*string) *TotalAppliedConsumStatResponse {
	s.Headers = v
	return s
}

func (s *TotalAppliedConsumStatResponse) SetBody(v *TotalAppliedConsumStatResponseBody) *TotalAppliedConsumStatResponse {
	s.Body = v
	return s
}

type TotalAppliedNearRealStatRequest struct {
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// 排序类型。默认：AppliedConcurrency_Desc,AppliedNearRealOrderConditionType[AppliedConcurrency_Desc(AppliedConcurrency_Desc,根据实时并发路数降序排列),AppliedConcurrency_Asc(AppliedConcurrency_Asc,根据实时并发路数升序排列),AppliedConsumptionCu_Desc(AppliedConsumptionCu_Desc,根据实时CU消耗降序排列),AppliedConsumptionCu_Asc(AppliedConsumptionCu_Asc,根据实时CU消耗升序排列),orderByType,desc]
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// 资源类型,PackageType[CU(cu),code,cssResourceType,desc]
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// 当前页码，默认1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页项数，默认20,最大100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s TotalAppliedNearRealStatRequest) String() string {
	return tea.Prettify(s)
}

func (s TotalAppliedNearRealStatRequest) GoString() string {
	return s.String()
}

func (s *TotalAppliedNearRealStatRequest) SetOperatorId(v string) *TotalAppliedNearRealStatRequest {
	s.OperatorId = &v
	return s
}

func (s *TotalAppliedNearRealStatRequest) SetOperatorType(v string) *TotalAppliedNearRealStatRequest {
	s.OperatorType = &v
	return s
}

func (s *TotalAppliedNearRealStatRequest) SetOrderBy(v string) *TotalAppliedNearRealStatRequest {
	s.OrderBy = &v
	return s
}

func (s *TotalAppliedNearRealStatRequest) SetPackageType(v string) *TotalAppliedNearRealStatRequest {
	s.PackageType = &v
	return s
}

func (s *TotalAppliedNearRealStatRequest) SetPageNumber(v int32) *TotalAppliedNearRealStatRequest {
	s.PageNumber = &v
	return s
}

func (s *TotalAppliedNearRealStatRequest) SetPageSize(v int32) *TotalAppliedNearRealStatRequest {
	s.PageSize = &v
	return s
}

type TotalAppliedNearRealStatResponseBody struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data *TotalAppliedNearRealStatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 业务处理消息摘要
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TotalAppliedNearRealStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TotalAppliedNearRealStatResponseBody) GoString() string {
	return s.String()
}

func (s *TotalAppliedNearRealStatResponseBody) SetCode(v string) *TotalAppliedNearRealStatResponseBody {
	s.Code = &v
	return s
}

func (s *TotalAppliedNearRealStatResponseBody) SetData(v *TotalAppliedNearRealStatResponseBodyData) *TotalAppliedNearRealStatResponseBody {
	s.Data = v
	return s
}

func (s *TotalAppliedNearRealStatResponseBody) SetMessage(v string) *TotalAppliedNearRealStatResponseBody {
	s.Message = &v
	return s
}

func (s *TotalAppliedNearRealStatResponseBody) SetRequestId(v string) *TotalAppliedNearRealStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *TotalAppliedNearRealStatResponseBody) SetSuccess(v bool) *TotalAppliedNearRealStatResponseBody {
	s.Success = &v
	return s
}

type TotalAppliedNearRealStatResponseBodyData struct {
	// 实时消耗并发
	TotalConcurrency *int64 `json:"TotalConcurrency,omitempty" xml:"TotalConcurrency,omitempty"`
	// 实时消耗CU
	TotalConsumptionCu *float64 `json:"TotalConsumptionCu,omitempty" xml:"TotalConsumptionCu,omitempty"`
}

func (s TotalAppliedNearRealStatResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TotalAppliedNearRealStatResponseBodyData) GoString() string {
	return s.String()
}

func (s *TotalAppliedNearRealStatResponseBodyData) SetTotalConcurrency(v int64) *TotalAppliedNearRealStatResponseBodyData {
	s.TotalConcurrency = &v
	return s
}

func (s *TotalAppliedNearRealStatResponseBodyData) SetTotalConsumptionCu(v float64) *TotalAppliedNearRealStatResponseBodyData {
	s.TotalConsumptionCu = &v
	return s
}

type TotalAppliedNearRealStatResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TotalAppliedNearRealStatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TotalAppliedNearRealStatResponse) String() string {
	return tea.Prettify(s)
}

func (s TotalAppliedNearRealStatResponse) GoString() string {
	return s.String()
}

func (s *TotalAppliedNearRealStatResponse) SetHeaders(v map[string]*string) *TotalAppliedNearRealStatResponse {
	s.Headers = v
	return s
}

func (s *TotalAppliedNearRealStatResponse) SetBody(v *TotalAppliedNearRealStatResponseBody) *TotalAppliedNearRealStatResponse {
	s.Body = v
	return s
}

type TotalQueryResourcePackageRequest struct {
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// 资源包类型,PackageType[CU(cu),code,cssResourceType,desc]
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
}

func (s TotalQueryResourcePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s TotalQueryResourcePackageRequest) GoString() string {
	return s.String()
}

func (s *TotalQueryResourcePackageRequest) SetOperatorId(v string) *TotalQueryResourcePackageRequest {
	s.OperatorId = &v
	return s
}

func (s *TotalQueryResourcePackageRequest) SetOperatorType(v string) *TotalQueryResourcePackageRequest {
	s.OperatorType = &v
	return s
}

func (s *TotalQueryResourcePackageRequest) SetPackageType(v string) *TotalQueryResourcePackageRequest {
	s.PackageType = &v
	return s
}

type TotalQueryResourcePackageResponseBody struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data *TotalQueryResourcePackageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 业务处理消息摘要
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TotalQueryResourcePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TotalQueryResourcePackageResponseBody) GoString() string {
	return s.String()
}

func (s *TotalQueryResourcePackageResponseBody) SetCode(v string) *TotalQueryResourcePackageResponseBody {
	s.Code = &v
	return s
}

func (s *TotalQueryResourcePackageResponseBody) SetData(v *TotalQueryResourcePackageResponseBodyData) *TotalQueryResourcePackageResponseBody {
	s.Data = v
	return s
}

func (s *TotalQueryResourcePackageResponseBody) SetMessage(v string) *TotalQueryResourcePackageResponseBody {
	s.Message = &v
	return s
}

func (s *TotalQueryResourcePackageResponseBody) SetRequestId(v string) *TotalQueryResourcePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *TotalQueryResourcePackageResponseBody) SetSuccess(v bool) *TotalQueryResourcePackageResponseBody {
	s.Success = &v
	return s
}

type TotalQueryResourcePackageResponseBodyData struct {
	// 租户UserId
	TenantUid *string `json:"TenantUid,omitempty" xml:"TenantUid,omitempty"`
	// 当前所有有效资源包总量
	TotalAmount *int64 `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	// 计算时间
	TotalDate *string `json:"TotalDate,omitempty" xml:"TotalDate,omitempty"`
}

func (s TotalQueryResourcePackageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TotalQueryResourcePackageResponseBodyData) GoString() string {
	return s.String()
}

func (s *TotalQueryResourcePackageResponseBodyData) SetTenantUid(v string) *TotalQueryResourcePackageResponseBodyData {
	s.TenantUid = &v
	return s
}

func (s *TotalQueryResourcePackageResponseBodyData) SetTotalAmount(v int64) *TotalQueryResourcePackageResponseBodyData {
	s.TotalAmount = &v
	return s
}

func (s *TotalQueryResourcePackageResponseBodyData) SetTotalDate(v string) *TotalQueryResourcePackageResponseBodyData {
	s.TotalDate = &v
	return s
}

type TotalQueryResourcePackageResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TotalQueryResourcePackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TotalQueryResourcePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s TotalQueryResourcePackageResponse) GoString() string {
	return s.String()
}

func (s *TotalQueryResourcePackageResponse) SetHeaders(v map[string]*string) *TotalQueryResourcePackageResponse {
	s.Headers = v
	return s
}

func (s *TotalQueryResourcePackageResponse) SetBody(v *TotalQueryResourcePackageResponseBody) *TotalQueryResourcePackageResponse {
	s.Body = v
	return s
}

type DataAppliedConsumptionMapValue struct {
	// 应用ID
	AppliedId *string `json:"AppliedId,omitempty" xml:"AppliedId,omitempty"`
	// 统计日期
	StatDate *string `json:"StatDate,omitempty" xml:"StatDate,omitempty"`
	// 分钟级消耗CU
	ConsumptionCu *int64 `json:"ConsumptionCu,omitempty" xml:"ConsumptionCu,omitempty"`
}

func (s DataAppliedConsumptionMapValue) String() string {
	return tea.Prettify(s)
}

func (s DataAppliedConsumptionMapValue) GoString() string {
	return s.String()
}

func (s *DataAppliedConsumptionMapValue) SetAppliedId(v string) *DataAppliedConsumptionMapValue {
	s.AppliedId = &v
	return s
}

func (s *DataAppliedConsumptionMapValue) SetStatDate(v string) *DataAppliedConsumptionMapValue {
	s.StatDate = &v
	return s
}

func (s *DataAppliedConsumptionMapValue) SetConsumptionCu(v int64) *DataAppliedConsumptionMapValue {
	s.ConsumptionCu = &v
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cgcs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AdaptCreateServiceWithOptions(tmpReq *AdaptCreateServiceRequest, runtime *util.RuntimeOptions) (_result *AdaptCreateServiceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AdaptCreateServiceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.AdaptTarget))) {
		request.AdaptTargetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.AdaptTarget), tea.String("AdaptTarget"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdaptTargetShrink)) {
		body["AdaptTarget"] = request.AdaptTargetShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersionId)) {
		body["AppVersionId"] = request.AppVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersionName)) {
		body["AppVersionName"] = request.AppVersionName
	}

	if !tea.BoolValue(util.IsUnset(request.RequestApp)) {
		body["RequestApp"] = request.RequestApp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AdaptCreateService"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AdaptCreateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AdaptCreateService(request *AdaptCreateServiceRequest) (_result *AdaptCreateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AdaptCreateServiceResponse{}
	_body, _err := client.AdaptCreateServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AdaptGetServiceWithOptions(request *AdaptGetServiceRequest, runtime *util.RuntimeOptions) (_result *AdaptGetServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersionId)) {
		body["AppVersionId"] = request.AppVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestApp)) {
		body["RequestApp"] = request.RequestApp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AdaptGetService"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AdaptGetServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AdaptGetService(request *AdaptGetServiceRequest) (_result *AdaptGetServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AdaptGetServiceResponse{}
	_body, _err := client.AdaptGetServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppCreateServiceWithOptions(request *AppCreateServiceRequest, runtime *util.RuntimeOptions) (_result *AppCreateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.RequestApp)) {
		body["RequestApp"] = request.RequestApp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppCreateService"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppCreateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppCreateService(request *AppCreateServiceRequest) (_result *AppCreateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppCreateServiceResponse{}
	_body, _err := client.AppCreateServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppDeleteServiceWithOptions(request *AppDeleteServiceRequest, runtime *util.RuntimeOptions) (_result *AppDeleteServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestApp)) {
		body["RequestApp"] = request.RequestApp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppDeleteService"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppDeleteServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppDeleteService(request *AppDeleteServiceRequest) (_result *AppDeleteServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppDeleteServiceResponse{}
	_body, _err := client.AppDeleteServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppGetServiceWithOptions(request *AppGetServiceRequest, runtime *util.RuntimeOptions) (_result *AppGetServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestApp)) {
		body["RequestApp"] = request.RequestApp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppGetService"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppGetServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppGetService(request *AppGetServiceRequest) (_result *AppGetServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppGetServiceResponse{}
	_body, _err := client.AppGetServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppListServiceWithOptions(request *AppListServiceRequest, runtime *util.RuntimeOptions) (_result *AppListServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeySearch)) {
		body["KeySearch"] = request.KeySearch
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RequestApp)) {
		body["RequestApp"] = request.RequestApp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppListService"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppListServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppListService(request *AppListServiceRequest) (_result *AppListServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppListServiceResponse{}
	_body, _err := client.AppListServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppModifyServiceWithOptions(request *AppModifyServiceRequest, runtime *util.RuntimeOptions) (_result *AppModifyServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.RequestApp)) {
		body["RequestApp"] = request.RequestApp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppModifyService"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppModifyServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppModifyService(request *AppModifyServiceRequest) (_result *AppModifyServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppModifyServiceResponse{}
	_body, _err := client.AppModifyServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppVersionCreateServiceWithOptions(request *AppVersionCreateServiceRequest, runtime *util.RuntimeOptions) (_result *AppVersionCreateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersionName)) {
		body["AppVersionName"] = request.AppVersionName
	}

	if !tea.BoolValue(util.IsUnset(request.FileAddress)) {
		body["FileAddress"] = request.FileAddress
	}

	if !tea.BoolValue(util.IsUnset(request.FileSize)) {
		body["FileSize"] = request.FileSize
	}

	if !tea.BoolValue(util.IsUnset(request.FileUploadType)) {
		body["FileUploadType"] = request.FileUploadType
	}

	if !tea.BoolValue(util.IsUnset(request.RequestApp)) {
		body["RequestApp"] = request.RequestApp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppVersionCreateService"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppVersionCreateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppVersionCreateService(request *AppVersionCreateServiceRequest) (_result *AppVersionCreateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppVersionCreateServiceResponse{}
	_body, _err := client.AppVersionCreateServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppVersionDeleteServiceWithOptions(request *AppVersionDeleteServiceRequest, runtime *util.RuntimeOptions) (_result *AppVersionDeleteServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersionId)) {
		body["AppVersionId"] = request.AppVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestApp)) {
		body["RequestApp"] = request.RequestApp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppVersionDeleteService"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppVersionDeleteServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppVersionDeleteService(request *AppVersionDeleteServiceRequest) (_result *AppVersionDeleteServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppVersionDeleteServiceResponse{}
	_body, _err := client.AppVersionDeleteServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppVersionGetServiceWithOptions(request *AppVersionGetServiceRequest, runtime *util.RuntimeOptions) (_result *AppVersionGetServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersionId)) {
		body["AppVersionId"] = request.AppVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestApp)) {
		body["RequestApp"] = request.RequestApp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppVersionGetService"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppVersionGetServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppVersionGetService(request *AppVersionGetServiceRequest) (_result *AppVersionGetServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppVersionGetServiceResponse{}
	_body, _err := client.AppVersionGetServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppVersionListServiceWithOptions(request *AppVersionListServiceRequest, runtime *util.RuntimeOptions) (_result *AppVersionListServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RequestApp)) {
		body["RequestApp"] = request.RequestApp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppVersionListService"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppVersionListServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppVersionListService(request *AppVersionListServiceRequest) (_result *AppVersionListServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppVersionListServiceResponse{}
	_body, _err := client.AppVersionListServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppVersionModifyServiceWithOptions(request *AppVersionModifyServiceRequest, runtime *util.RuntimeOptions) (_result *AppVersionModifyServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersionId)) {
		body["AppVersionId"] = request.AppVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersionName)) {
		body["AppVersionName"] = request.AppVersionName
	}

	if !tea.BoolValue(util.IsUnset(request.RequestApp)) {
		body["RequestApp"] = request.RequestApp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppVersionModifyService"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppVersionModifyServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppVersionModifyService(request *AppVersionModifyServiceRequest) (_result *AppVersionModifyServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppVersionModifyServiceResponse{}
	_body, _err := client.AppVersionModifyServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppVersionQueryServiceWithOptions(request *AppVersionQueryServiceRequest, runtime *util.RuntimeOptions) (_result *AppVersionQueryServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeySearch)) {
		body["KeySearch"] = request.KeySearch
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RequestApp)) {
		body["RequestApp"] = request.RequestApp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppVersionQueryService"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppVersionQueryServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppVersionQueryService(request *AppVersionQueryServiceRequest) (_result *AppVersionQueryServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppVersionQueryServiceResponse{}
	_body, _err := client.AppVersionQueryServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppliedConsumStatWithOptions(tmpReq *AppliedConsumStatRequest, runtime *util.RuntimeOptions) (_result *AppliedConsumStatResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AppliedConsumStatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AppliedId)) {
		request.AppliedIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppliedId, tea.String("AppliedId"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppliedIdShrink)) {
		body["AppliedId"] = request.AppliedIdShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorType)) {
		body["OperatorType"] = request.OperatorType
	}

	if !tea.BoolValue(util.IsUnset(request.PackageType)) {
		body["PackageType"] = request.PackageType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryEndDate)) {
		body["QueryEndDate"] = request.QueryEndDate
	}

	if !tea.BoolValue(util.IsUnset(request.QueryStartDate)) {
		body["QueryStartDate"] = request.QueryStartDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppliedConsumStat"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppliedConsumStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppliedConsumStat(request *AppliedConsumStatRequest) (_result *AppliedConsumStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppliedConsumStatResponse{}
	_body, _err := client.AppliedConsumStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppliedNearRealStatWithOptions(tmpReq *AppliedNearRealStatRequest, runtime *util.RuntimeOptions) (_result *AppliedNearRealStatResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AppliedNearRealStatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AppliedVersionId)) {
		request.AppliedVersionIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppliedVersionId, tea.String("AppliedVersionId"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppliedVersionIdShrink)) {
		body["AppliedVersionId"] = request.AppliedVersionIdShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorType)) {
		body["OperatorType"] = request.OperatorType
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		body["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.PackageType)) {
		body["PackageType"] = request.PackageType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppliedNearRealStat"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppliedNearRealStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppliedNearRealStat(request *AppliedNearRealStatRequest) (_result *AppliedNearRealStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppliedNearRealStatResponse{}
	_body, _err := client.AppliedNearRealStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppliedStatWithOptions(request *AppliedStatRequest, runtime *util.RuntimeOptions) (_result *AppliedStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorType)) {
		body["OperatorType"] = request.OperatorType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryEndDate)) {
		body["QueryEndDate"] = request.QueryEndDate
	}

	if !tea.BoolValue(util.IsUnset(request.QueryStartDate)) {
		body["QueryStartDate"] = request.QueryStartDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppliedStat"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppliedStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppliedStat(request *AppliedStatRequest) (_result *AppliedStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppliedStatResponse{}
	_body, _err := client.AppliedStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppSessionWithOptions(request *CreateAppSessionRequest, runtime *util.RuntimeOptions) (_result *CreateAppSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		query["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.CustomSessionId)) {
		query["CustomSessionId"] = request.CustomSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomUserId)) {
		query["CustomUserId"] = request.CustomUserId
	}

	if !tea.BoolValue(util.IsUnset(request.StartParameters)) {
		query["StartParameters"] = request.StartParameters
	}

	if !tea.BoolValue(util.IsUnset(request.SystemInfo)) {
		query["SystemInfo"] = request.SystemInfo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppSession"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppSession(request *CreateAppSessionRequest) (_result *CreateAppSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppSessionResponse{}
	_body, _err := client.CreateAppSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppSessionBatchWithOptions(tmpReq *CreateAppSessionBatchRequest, runtime *util.RuntimeOptions) (_result *CreateAppSessionBatchResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAppSessionBatchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AppInfos)) {
		request.AppInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppInfos, tea.String("AppInfos"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInfosShrink)) {
		query["AppInfos"] = request.AppInfosShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CustomTaskId)) {
		query["CustomTaskId"] = request.CustomTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["Timeout"] = request.Timeout
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppSessionBatch"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppSessionBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppSessionBatch(request *CreateAppSessionBatchRequest) (_result *CreateAppSessionBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppSessionBatchResponse{}
	_body, _err := client.CreateAppSessionBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUploadTaskWithOptions(request *CreateUploadTaskRequest, runtime *util.RuntimeOptions) (_result *CreateUploadTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		body["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.Env)) {
		body["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.FileAddress)) {
		body["FileAddress"] = request.FileAddress
	}

	if !tea.BoolValue(util.IsUnset(request.FileSize)) {
		body["FileSize"] = request.FileSize
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		body["FileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.Progress)) {
		body["Progress"] = request.Progress
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UploadToolVersion)) {
		body["UploadToolVersion"] = request.UploadToolVersion
	}

	if !tea.BoolValue(util.IsUnset(request.UploadType)) {
		body["UploadType"] = request.UploadType
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		body["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUploadTask"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUploadTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUploadTask(request *CreateUploadTaskRequest) (_result *CreateUploadTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUploadTaskResponse{}
	_body, _err := client.CreateUploadTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppListWithOptions(runtime *util.RuntimeOptions) (_result *GetAppListResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetAppList"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppList() (_result *GetAppListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppListResponse{}
	_body, _err := client.GetAppListWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppSessionWithOptions(request *GetAppSessionRequest, runtime *util.RuntimeOptions) (_result *GetAppSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomSessionId)) {
		query["CustomSessionId"] = request.CustomSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformSessionId)) {
		query["PlatformSessionId"] = request.PlatformSessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppSession"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppSession(request *GetAppSessionRequest) (_result *GetAppSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppSessionResponse{}
	_body, _err := client.GetAppSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNeedUploadFileListWithOptions(request *GetNeedUploadFileListRequest, runtime *util.RuntimeOptions) (_result *GetNeedUploadFileListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Env)) {
		body["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.HashList)) {
		body["HashList"] = request.HashList
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		body["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNeedUploadFileList"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNeedUploadFileListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNeedUploadFileList(request *GetNeedUploadFileListRequest) (_result *GetNeedUploadFileListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNeedUploadFileListResponse{}
	_body, _err := client.GetNeedUploadFileListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOssInfoWithOptions(runtime *util.RuntimeOptions) (_result *GetOssInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetOssInfo"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOssInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOssInfo() (_result *GetOssInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOssInfoResponse{}
	_body, _err := client.GetOssInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTenantIdWithOptions(runtime *util.RuntimeOptions) (_result *GetTenantIdResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetTenantId"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTenantIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTenantId() (_result *GetTenantIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTenantIdResponse{}
	_body, _err := client.GetTenantIdWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTokenWithOptions(request *GetTokenRequest, runtime *util.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Bucket)) {
		body["Bucket"] = request.Bucket
	}

	if !tea.BoolValue(util.IsUnset(request.Env)) {
		body["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		body["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetToken"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUploadToolUrlWithOptions(runtime *util.RuntimeOptions) (_result *GetUploadToolUrlResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetUploadToolUrl"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUploadToolUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUploadToolUrl() (_result *GetUploadToolUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUploadToolUrlResponse{}
	_body, _err := client.GetUploadToolUrlWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HasActivateWithOptions(runtime *util.RuntimeOptions) (_result *HasActivateResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("HasActivate"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HasActivateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HasActivate() (_result *HasActivateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HasActivateResponse{}
	_body, _err := client.HasActivateWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppSessionsWithOptions(request *ListAppSessionsRequest, runtime *util.RuntimeOptions) (_result *ListAppSessionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomSessionIds)) {
		query["CustomSessionIds"] = request.CustomSessionIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformSessionIds)) {
		query["PlatformSessionIds"] = request.PlatformSessionIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppSessions"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppSessionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppSessions(request *ListAppSessionsRequest) (_result *ListAppSessionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppSessionsResponse{}
	_body, _err := client.ListAppSessionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageQueryResourcePackageListWithOptions(request *PageQueryResourcePackageListRequest, runtime *util.RuntimeOptions) (_result *PageQueryResourcePackageListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorType)) {
		body["OperatorType"] = request.OperatorType
	}

	if !tea.BoolValue(util.IsUnset(request.PackageType)) {
		body["PackageType"] = request.PackageType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryValidType)) {
		body["QueryValidType"] = request.QueryValidType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PageQueryResourcePackageList"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PageQueryResourcePackageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageQueryResourcePackageList(request *PageQueryResourcePackageListRequest) (_result *PageQueryResourcePackageListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PageQueryResourcePackageListResponse{}
	_body, _err := client.PageQueryResourcePackageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAdaptRecordsWithOptions(request *QueryAdaptRecordsRequest, runtime *util.RuntimeOptions) (_result *QueryAdaptRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersionId)) {
		body["AppVersionId"] = request.AppVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestApp)) {
		body["RequestApp"] = request.RequestApp
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAdaptRecords"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAdaptRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAdaptRecords(request *QueryAdaptRecordsRequest) (_result *QueryAdaptRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAdaptRecordsResponse{}
	_body, _err := client.QueryAdaptRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUploadProgressWithOptions(request *QueryUploadProgressRequest, runtime *util.RuntimeOptions) (_result *QueryUploadProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueryUploadProgressRequests)) {
		body["QueryUploadProgressRequests"] = request.QueryUploadProgressRequests
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryUploadProgress"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUploadProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUploadProgress(request *QueryUploadProgressRequest) (_result *QueryUploadProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUploadProgressResponse{}
	_body, _err := client.QueryUploadProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecordFinishedFileWithOptions(request *RecordFinishedFileRequest, runtime *util.RuntimeOptions) (_result *RecordFinishedFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		body["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.Env)) {
		body["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.FileFingerprintDTOList)) {
		body["FileFingerprintDTOList"] = request.FileFingerprintDTOList
	}

	if !tea.BoolValue(util.IsUnset(request.FileSize)) {
		body["FileSize"] = request.FileSize
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		body["FileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ToolVersion)) {
		body["ToolVersion"] = request.ToolVersion
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		body["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecordFinishedFile"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecordFinishedFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecordFinishedFile(request *RecordFinishedFileRequest) (_result *RecordFinishedFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecordFinishedFileResponse{}
	_body, _err := client.RecordFinishedFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReplicateVersionWithOptions(request *ReplicateVersionRequest, runtime *util.RuntimeOptions) (_result *ReplicateVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceVersionId)) {
		body["SourceVersionId"] = request.SourceVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetVersionId)) {
		body["TargetVersionId"] = request.TargetVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReplicateVersion"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReplicateVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReplicateVersion(request *ReplicateVersionRequest) (_result *ReplicateVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReplicateVersionResponse{}
	_body, _err := client.ReplicateVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportUploadProgressWithOptions(request *ReportUploadProgressRequest, runtime *util.RuntimeOptions) (_result *ReportUploadProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Env)) {
		body["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.Progress)) {
		body["Progress"] = request.Progress
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		body["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportUploadProgress"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportUploadProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportUploadProgress(request *ReportUploadProgressRequest) (_result *ReportUploadProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportUploadProgressResponse{}
	_body, _err := client.ReportUploadProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportUploadResultWithOptions(request *ReportUploadResultRequest, runtime *util.RuntimeOptions) (_result *ReportUploadResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		body["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.Env)) {
		body["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.FileFingerprintDTOList)) {
		body["FileFingerprintDTOList"] = request.FileFingerprintDTOList
	}

	if !tea.BoolValue(util.IsUnset(request.FileSize)) {
		body["FileSize"] = request.FileSize
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		body["FileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ToolVersion)) {
		body["ToolVersion"] = request.ToolVersion
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		body["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportUploadResult"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportUploadResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportUploadResult(request *ReportUploadResultRequest) (_result *ReportUploadResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportUploadResultResponse{}
	_body, _err := client.ReportUploadResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportUploadStatusWithOptions(request *ReportUploadStatusRequest, runtime *util.RuntimeOptions) (_result *ReportUploadStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Env)) {
		body["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.Memo)) {
		body["Memo"] = request.Memo
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		body["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportUploadStatus"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportUploadStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportUploadStatus(request *ReportUploadStatusRequest) (_result *ReportUploadStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportUploadStatusResponse{}
	_body, _err := client.ReportUploadStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopAppSessionWithOptions(request *StopAppSessionRequest, runtime *util.RuntimeOptions) (_result *StopAppSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomSessionId)) {
		query["CustomSessionId"] = request.CustomSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformSessionId)) {
		query["PlatformSessionId"] = request.PlatformSessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopAppSession"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopAppSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopAppSession(request *StopAppSessionRequest) (_result *StopAppSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopAppSessionResponse{}
	_body, _err := client.StopAppSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TotalAppliedConsumStatWithOptions(request *TotalAppliedConsumStatRequest, runtime *util.RuntimeOptions) (_result *TotalAppliedConsumStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorType)) {
		body["OperatorType"] = request.OperatorType
	}

	if !tea.BoolValue(util.IsUnset(request.PackageType)) {
		body["PackageType"] = request.PackageType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryEndDate)) {
		body["QueryEndDate"] = request.QueryEndDate
	}

	if !tea.BoolValue(util.IsUnset(request.QueryStartDate)) {
		body["QueryStartDate"] = request.QueryStartDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TotalAppliedConsumStat"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TotalAppliedConsumStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TotalAppliedConsumStat(request *TotalAppliedConsumStatRequest) (_result *TotalAppliedConsumStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TotalAppliedConsumStatResponse{}
	_body, _err := client.TotalAppliedConsumStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TotalAppliedNearRealStatWithOptions(request *TotalAppliedNearRealStatRequest, runtime *util.RuntimeOptions) (_result *TotalAppliedNearRealStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorType)) {
		body["OperatorType"] = request.OperatorType
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		body["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.PackageType)) {
		body["PackageType"] = request.PackageType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TotalAppliedNearRealStat"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TotalAppliedNearRealStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TotalAppliedNearRealStat(request *TotalAppliedNearRealStatRequest) (_result *TotalAppliedNearRealStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TotalAppliedNearRealStatResponse{}
	_body, _err := client.TotalAppliedNearRealStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TotalQueryResourcePackageWithOptions(request *TotalQueryResourcePackageRequest, runtime *util.RuntimeOptions) (_result *TotalQueryResourcePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorType)) {
		body["OperatorType"] = request.OperatorType
	}

	if !tea.BoolValue(util.IsUnset(request.PackageType)) {
		body["PackageType"] = request.PackageType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TotalQueryResourcePackage"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TotalQueryResourcePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TotalQueryResourcePackage(request *TotalQueryResourcePackageRequest) (_result *TotalQueryResourcePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TotalQueryResourcePackageResponse{}
	_body, _err := client.TotalQueryResourcePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
