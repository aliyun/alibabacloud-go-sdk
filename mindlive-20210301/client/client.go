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

type LoginDeviceRequest struct {
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserSource *string `json:"UserSource,omitempty" xml:"UserSource,omitempty"`
}

func (s LoginDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s LoginDeviceRequest) GoString() string {
	return s.String()
}

func (s *LoginDeviceRequest) SetUserId(v string) *LoginDeviceRequest {
	s.UserId = &v
	return s
}

func (s *LoginDeviceRequest) SetUserSource(v string) *LoginDeviceRequest {
	s.UserSource = &v
	return s
}

type LoginDeviceResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LoginDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoginDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *LoginDeviceResponseBody) SetErrorCode(v string) *LoginDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *LoginDeviceResponseBody) SetErrorMessage(v string) *LoginDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *LoginDeviceResponseBody) SetRequestId(v string) *LoginDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *LoginDeviceResponseBody) SetSuccess(v bool) *LoginDeviceResponseBody {
	s.Success = &v
	return s
}

type LoginDeviceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LoginDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LoginDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s LoginDeviceResponse) GoString() string {
	return s.String()
}

func (s *LoginDeviceResponse) SetHeaders(v map[string]*string) *LoginDeviceResponse {
	s.Headers = v
	return s
}

func (s *LoginDeviceResponse) SetStatusCode(v int32) *LoginDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *LoginDeviceResponse) SetBody(v *LoginDeviceResponseBody) *LoginDeviceResponse {
	s.Body = v
	return s
}

type LogoutDeviceRequest struct {
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserSource *string `json:"UserSource,omitempty" xml:"UserSource,omitempty"`
}

func (s LogoutDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s LogoutDeviceRequest) GoString() string {
	return s.String()
}

func (s *LogoutDeviceRequest) SetUserId(v string) *LogoutDeviceRequest {
	s.UserId = &v
	return s
}

func (s *LogoutDeviceRequest) SetUserSource(v string) *LogoutDeviceRequest {
	s.UserSource = &v
	return s
}

type LogoutDeviceResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LogoutDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LogoutDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *LogoutDeviceResponseBody) SetErrorCode(v string) *LogoutDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *LogoutDeviceResponseBody) SetErrorMessage(v string) *LogoutDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *LogoutDeviceResponseBody) SetRequestId(v string) *LogoutDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *LogoutDeviceResponseBody) SetSuccess(v bool) *LogoutDeviceResponseBody {
	s.Success = &v
	return s
}

type LogoutDeviceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LogoutDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LogoutDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s LogoutDeviceResponse) GoString() string {
	return s.String()
}

func (s *LogoutDeviceResponse) SetHeaders(v map[string]*string) *LogoutDeviceResponse {
	s.Headers = v
	return s
}

func (s *LogoutDeviceResponse) SetStatusCode(v int32) *LogoutDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *LogoutDeviceResponse) SetBody(v *LogoutDeviceResponseBody) *LogoutDeviceResponse {
	s.Body = v
	return s
}

type QueryItemBackgroundsRequest struct {
	ItemIds map[string]interface{} `json:"ItemIds,omitempty" xml:"ItemIds,omitempty"`
}

func (s QueryItemBackgroundsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryItemBackgroundsRequest) GoString() string {
	return s.String()
}

func (s *QueryItemBackgroundsRequest) SetItemIds(v map[string]interface{}) *QueryItemBackgroundsRequest {
	s.ItemIds = v
	return s
}

type QueryItemBackgroundsShrinkRequest struct {
	ItemIdsShrink *string `json:"ItemIds,omitempty" xml:"ItemIds,omitempty"`
}

func (s QueryItemBackgroundsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryItemBackgroundsShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryItemBackgroundsShrinkRequest) SetItemIdsShrink(v string) *QueryItemBackgroundsShrinkRequest {
	s.ItemIdsShrink = &v
	return s
}

type QueryItemBackgroundsResponseBody struct {
	Data         map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *string                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryItemBackgroundsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryItemBackgroundsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryItemBackgroundsResponseBody) SetData(v map[string]interface{}) *QueryItemBackgroundsResponseBody {
	s.Data = v
	return s
}

func (s *QueryItemBackgroundsResponseBody) SetErrorCode(v string) *QueryItemBackgroundsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryItemBackgroundsResponseBody) SetErrorMessage(v string) *QueryItemBackgroundsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryItemBackgroundsResponseBody) SetRequestId(v string) *QueryItemBackgroundsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryItemBackgroundsResponseBody) SetSuccess(v bool) *QueryItemBackgroundsResponseBody {
	s.Success = &v
	return s
}

type QueryItemBackgroundsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryItemBackgroundsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryItemBackgroundsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryItemBackgroundsResponse) GoString() string {
	return s.String()
}

func (s *QueryItemBackgroundsResponse) SetHeaders(v map[string]*string) *QueryItemBackgroundsResponse {
	s.Headers = v
	return s
}

func (s *QueryItemBackgroundsResponse) SetStatusCode(v int32) *QueryItemBackgroundsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryItemBackgroundsResponse) SetBody(v *QueryItemBackgroundsResponseBody) *QueryItemBackgroundsResponse {
	s.Body = v
	return s
}

type ReportCurrentBackgroundRequest struct {
	BgConfig     map[string]interface{} `json:"BgConfig,omitempty" xml:"BgConfig,omitempty"`
	Mode         *string                `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Open         *bool                  `json:"Open,omitempty" xml:"Open,omitempty"`
	ResourceUuid *string                `json:"ResourceUuid,omitempty" xml:"ResourceUuid,omitempty"`
}

func (s ReportCurrentBackgroundRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportCurrentBackgroundRequest) GoString() string {
	return s.String()
}

func (s *ReportCurrentBackgroundRequest) SetBgConfig(v map[string]interface{}) *ReportCurrentBackgroundRequest {
	s.BgConfig = v
	return s
}

func (s *ReportCurrentBackgroundRequest) SetMode(v string) *ReportCurrentBackgroundRequest {
	s.Mode = &v
	return s
}

func (s *ReportCurrentBackgroundRequest) SetOpen(v bool) *ReportCurrentBackgroundRequest {
	s.Open = &v
	return s
}

func (s *ReportCurrentBackgroundRequest) SetResourceUuid(v string) *ReportCurrentBackgroundRequest {
	s.ResourceUuid = &v
	return s
}

type ReportCurrentBackgroundShrinkRequest struct {
	BgConfigShrink *string `json:"BgConfig,omitempty" xml:"BgConfig,omitempty"`
	Mode           *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Open           *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	ResourceUuid   *string `json:"ResourceUuid,omitempty" xml:"ResourceUuid,omitempty"`
}

func (s ReportCurrentBackgroundShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportCurrentBackgroundShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReportCurrentBackgroundShrinkRequest) SetBgConfigShrink(v string) *ReportCurrentBackgroundShrinkRequest {
	s.BgConfigShrink = &v
	return s
}

func (s *ReportCurrentBackgroundShrinkRequest) SetMode(v string) *ReportCurrentBackgroundShrinkRequest {
	s.Mode = &v
	return s
}

func (s *ReportCurrentBackgroundShrinkRequest) SetOpen(v bool) *ReportCurrentBackgroundShrinkRequest {
	s.Open = &v
	return s
}

func (s *ReportCurrentBackgroundShrinkRequest) SetResourceUuid(v string) *ReportCurrentBackgroundShrinkRequest {
	s.ResourceUuid = &v
	return s
}

type ReportCurrentBackgroundResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReportCurrentBackgroundResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportCurrentBackgroundResponseBody) GoString() string {
	return s.String()
}

func (s *ReportCurrentBackgroundResponseBody) SetErrorCode(v string) *ReportCurrentBackgroundResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReportCurrentBackgroundResponseBody) SetErrorMessage(v string) *ReportCurrentBackgroundResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ReportCurrentBackgroundResponseBody) SetRequestId(v string) *ReportCurrentBackgroundResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportCurrentBackgroundResponseBody) SetSuccess(v bool) *ReportCurrentBackgroundResponseBody {
	s.Success = &v
	return s
}

type ReportCurrentBackgroundResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportCurrentBackgroundResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportCurrentBackgroundResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportCurrentBackgroundResponse) GoString() string {
	return s.String()
}

func (s *ReportCurrentBackgroundResponse) SetHeaders(v map[string]*string) *ReportCurrentBackgroundResponse {
	s.Headers = v
	return s
}

func (s *ReportCurrentBackgroundResponse) SetStatusCode(v int32) *ReportCurrentBackgroundResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportCurrentBackgroundResponse) SetBody(v *ReportCurrentBackgroundResponseBody) *ReportCurrentBackgroundResponse {
	s.Body = v
	return s
}

type ReportDevConfigRequest struct {
	Configs *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
}

func (s ReportDevConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportDevConfigRequest) GoString() string {
	return s.String()
}

func (s *ReportDevConfigRequest) SetConfigs(v string) *ReportDevConfigRequest {
	s.Configs = &v
	return s
}

type ReportDevConfigResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReportDevConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportDevConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ReportDevConfigResponseBody) SetErrorCode(v string) *ReportDevConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReportDevConfigResponseBody) SetErrorMessage(v string) *ReportDevConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ReportDevConfigResponseBody) SetRequestId(v string) *ReportDevConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportDevConfigResponseBody) SetSuccess(v bool) *ReportDevConfigResponseBody {
	s.Success = &v
	return s
}

type ReportDevConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportDevConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportDevConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportDevConfigResponse) GoString() string {
	return s.String()
}

func (s *ReportDevConfigResponse) SetHeaders(v map[string]*string) *ReportDevConfigResponse {
	s.Headers = v
	return s
}

func (s *ReportDevConfigResponse) SetStatusCode(v int32) *ReportDevConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportDevConfigResponse) SetBody(v *ReportDevConfigResponseBody) *ReportDevConfigResponse {
	s.Body = v
	return s
}

type ReportLiveStateRequest struct {
	AnchorId  *string `json:"AnchorId,omitempty" xml:"AnchorId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	LiveMode  *string `json:"LiveMode,omitempty" xml:"LiveMode,omitempty"`
	LiveState *string `json:"LiveState,omitempty" xml:"LiveState,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ReportLiveStateRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportLiveStateRequest) GoString() string {
	return s.String()
}

func (s *ReportLiveStateRequest) SetAnchorId(v string) *ReportLiveStateRequest {
	s.AnchorId = &v
	return s
}

func (s *ReportLiveStateRequest) SetId(v string) *ReportLiveStateRequest {
	s.Id = &v
	return s
}

func (s *ReportLiveStateRequest) SetLiveMode(v string) *ReportLiveStateRequest {
	s.LiveMode = &v
	return s
}

func (s *ReportLiveStateRequest) SetLiveState(v string) *ReportLiveStateRequest {
	s.LiveState = &v
	return s
}

func (s *ReportLiveStateRequest) SetStartTime(v int64) *ReportLiveStateRequest {
	s.StartTime = &v
	return s
}

func (s *ReportLiveStateRequest) SetType(v string) *ReportLiveStateRequest {
	s.Type = &v
	return s
}

type ReportLiveStateResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReportLiveStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportLiveStateResponseBody) GoString() string {
	return s.String()
}

func (s *ReportLiveStateResponseBody) SetErrorCode(v string) *ReportLiveStateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReportLiveStateResponseBody) SetErrorMessage(v string) *ReportLiveStateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ReportLiveStateResponseBody) SetRequestId(v string) *ReportLiveStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportLiveStateResponseBody) SetSuccess(v bool) *ReportLiveStateResponseBody {
	s.Success = &v
	return s
}

type ReportLiveStateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportLiveStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportLiveStateResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportLiveStateResponse) GoString() string {
	return s.String()
}

func (s *ReportLiveStateResponse) SetHeaders(v map[string]*string) *ReportLiveStateResponse {
	s.Headers = v
	return s
}

func (s *ReportLiveStateResponse) SetStatusCode(v int32) *ReportLiveStateResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportLiveStateResponse) SetBody(v *ReportLiveStateResponseBody) *ReportLiveStateResponse {
	s.Body = v
	return s
}

type ReportScreenRequest struct {
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	OssObjectKey  *string `json:"OssObjectKey,omitempty" xml:"OssObjectKey,omitempty"`
}

func (s ReportScreenRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportScreenRequest) GoString() string {
	return s.String()
}

func (s *ReportScreenRequest) SetOssBucketName(v string) *ReportScreenRequest {
	s.OssBucketName = &v
	return s
}

func (s *ReportScreenRequest) SetOssObjectKey(v string) *ReportScreenRequest {
	s.OssObjectKey = &v
	return s
}

type ReportScreenResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReportScreenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportScreenResponseBody) GoString() string {
	return s.String()
}

func (s *ReportScreenResponseBody) SetErrorCode(v string) *ReportScreenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReportScreenResponseBody) SetErrorMessage(v string) *ReportScreenResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ReportScreenResponseBody) SetRequestId(v string) *ReportScreenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportScreenResponseBody) SetSuccess(v bool) *ReportScreenResponseBody {
	s.Success = &v
	return s
}

type ReportScreenResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportScreenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportScreenResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportScreenResponse) GoString() string {
	return s.String()
}

func (s *ReportScreenResponse) SetHeaders(v map[string]*string) *ReportScreenResponse {
	s.Headers = v
	return s
}

func (s *ReportScreenResponse) SetStatusCode(v int32) *ReportScreenResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportScreenResponse) SetBody(v *ReportScreenResponseBody) *ReportScreenResponse {
	s.Body = v
	return s
}

type ReportUserConfigRequest struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ReportUserConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportUserConfigRequest) GoString() string {
	return s.String()
}

func (s *ReportUserConfigRequest) SetKey(v string) *ReportUserConfigRequest {
	s.Key = &v
	return s
}

func (s *ReportUserConfigRequest) SetValue(v string) *ReportUserConfigRequest {
	s.Value = &v
	return s
}

type ReportUserConfigResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReportUserConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ReportUserConfigResponseBody) SetErrorCode(v string) *ReportUserConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReportUserConfigResponseBody) SetErrorMessage(v string) *ReportUserConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ReportUserConfigResponseBody) SetRequestId(v string) *ReportUserConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportUserConfigResponseBody) SetSuccess(v bool) *ReportUserConfigResponseBody {
	s.Success = &v
	return s
}

type ReportUserConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportUserConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportUserConfigResponse) GoString() string {
	return s.String()
}

func (s *ReportUserConfigResponse) SetHeaders(v map[string]*string) *ReportUserConfigResponse {
	s.Headers = v
	return s
}

func (s *ReportUserConfigResponse) SetStatusCode(v int32) *ReportUserConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportUserConfigResponse) SetBody(v *ReportUserConfigResponseBody) *ReportUserConfigResponse {
	s.Body = v
	return s
}

type RequestAuthorizationDataResponseBody struct {
	Data         *RequestAuthorizationDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequestAuthorizationDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestAuthorizationDataResponseBody) GoString() string {
	return s.String()
}

func (s *RequestAuthorizationDataResponseBody) SetData(v *RequestAuthorizationDataResponseBodyData) *RequestAuthorizationDataResponseBody {
	s.Data = v
	return s
}

func (s *RequestAuthorizationDataResponseBody) SetErrorCode(v string) *RequestAuthorizationDataResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RequestAuthorizationDataResponseBody) SetErrorMessage(v string) *RequestAuthorizationDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RequestAuthorizationDataResponseBody) SetRequestId(v string) *RequestAuthorizationDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestAuthorizationDataResponseBody) SetSuccess(v bool) *RequestAuthorizationDataResponseBody {
	s.Success = &v
	return s
}

type RequestAuthorizationDataResponseBodyData struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RequestAuthorizationDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RequestAuthorizationDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *RequestAuthorizationDataResponseBodyData) SetUrl(v string) *RequestAuthorizationDataResponseBodyData {
	s.Url = &v
	return s
}

type RequestAuthorizationDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestAuthorizationDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestAuthorizationDataResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestAuthorizationDataResponse) GoString() string {
	return s.String()
}

func (s *RequestAuthorizationDataResponse) SetHeaders(v map[string]*string) *RequestAuthorizationDataResponse {
	s.Headers = v
	return s
}

func (s *RequestAuthorizationDataResponse) SetStatusCode(v int32) *RequestAuthorizationDataResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestAuthorizationDataResponse) SetBody(v *RequestAuthorizationDataResponseBody) *RequestAuthorizationDataResponse {
	s.Body = v
	return s
}

type RequestBackgroundResponseBody struct {
	Data         *RequestBackgroundResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequestBackgroundResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestBackgroundResponseBody) GoString() string {
	return s.String()
}

func (s *RequestBackgroundResponseBody) SetData(v *RequestBackgroundResponseBodyData) *RequestBackgroundResponseBody {
	s.Data = v
	return s
}

func (s *RequestBackgroundResponseBody) SetErrorCode(v string) *RequestBackgroundResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RequestBackgroundResponseBody) SetErrorMessage(v string) *RequestBackgroundResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RequestBackgroundResponseBody) SetRequestId(v string) *RequestBackgroundResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestBackgroundResponseBody) SetSuccess(v bool) *RequestBackgroundResponseBody {
	s.Success = &v
	return s
}

type RequestBackgroundResponseBodyData struct {
	BgConfig     map[string]interface{} `json:"BgConfig,omitempty" xml:"BgConfig,omitempty"`
	DownloadUrl  *string                `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	FileType     *string                `json:"FileType,omitempty" xml:"FileType,omitempty"`
	Mode         *string                `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Open         *bool                  `json:"Open,omitempty" xml:"Open,omitempty"`
	ResourceUuid *string                `json:"ResourceUuid,omitempty" xml:"ResourceUuid,omitempty"`
	Scope        *string                `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s RequestBackgroundResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RequestBackgroundResponseBodyData) GoString() string {
	return s.String()
}

func (s *RequestBackgroundResponseBodyData) SetBgConfig(v map[string]interface{}) *RequestBackgroundResponseBodyData {
	s.BgConfig = v
	return s
}

func (s *RequestBackgroundResponseBodyData) SetDownloadUrl(v string) *RequestBackgroundResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *RequestBackgroundResponseBodyData) SetFileType(v string) *RequestBackgroundResponseBodyData {
	s.FileType = &v
	return s
}

func (s *RequestBackgroundResponseBodyData) SetMode(v string) *RequestBackgroundResponseBodyData {
	s.Mode = &v
	return s
}

func (s *RequestBackgroundResponseBodyData) SetOpen(v bool) *RequestBackgroundResponseBodyData {
	s.Open = &v
	return s
}

func (s *RequestBackgroundResponseBodyData) SetResourceUuid(v string) *RequestBackgroundResponseBodyData {
	s.ResourceUuid = &v
	return s
}

func (s *RequestBackgroundResponseBodyData) SetScope(v string) *RequestBackgroundResponseBodyData {
	s.Scope = &v
	return s
}

type RequestBackgroundResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestBackgroundResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestBackgroundResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestBackgroundResponse) GoString() string {
	return s.String()
}

func (s *RequestBackgroundResponse) SetHeaders(v map[string]*string) *RequestBackgroundResponse {
	s.Headers = v
	return s
}

func (s *RequestBackgroundResponse) SetStatusCode(v int32) *RequestBackgroundResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestBackgroundResponse) SetBody(v *RequestBackgroundResponseBody) *RequestBackgroundResponse {
	s.Body = v
	return s
}

type RequestBindDataRequest struct {
	LiveSource *string `json:"LiveSource,omitempty" xml:"LiveSource,omitempty"`
}

func (s RequestBindDataRequest) String() string {
	return tea.Prettify(s)
}

func (s RequestBindDataRequest) GoString() string {
	return s.String()
}

func (s *RequestBindDataRequest) SetLiveSource(v string) *RequestBindDataRequest {
	s.LiveSource = &v
	return s
}

type RequestBindDataResponseBody struct {
	Data         *RequestBindDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequestBindDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestBindDataResponseBody) GoString() string {
	return s.String()
}

func (s *RequestBindDataResponseBody) SetData(v *RequestBindDataResponseBodyData) *RequestBindDataResponseBody {
	s.Data = v
	return s
}

func (s *RequestBindDataResponseBody) SetErrorCode(v string) *RequestBindDataResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RequestBindDataResponseBody) SetErrorMessage(v string) *RequestBindDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RequestBindDataResponseBody) SetRequestId(v string) *RequestBindDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestBindDataResponseBody) SetSuccess(v bool) *RequestBindDataResponseBody {
	s.Success = &v
	return s
}

type RequestBindDataResponseBodyData struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	MaxKeepSeconds *int32  `json:"MaxKeepSeconds,omitempty" xml:"MaxKeepSeconds,omitempty"`
	ShortUrl       *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RequestBindDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RequestBindDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *RequestBindDataResponseBodyData) SetCode(v string) *RequestBindDataResponseBodyData {
	s.Code = &v
	return s
}

func (s *RequestBindDataResponseBodyData) SetMaxKeepSeconds(v int32) *RequestBindDataResponseBodyData {
	s.MaxKeepSeconds = &v
	return s
}

func (s *RequestBindDataResponseBodyData) SetShortUrl(v string) *RequestBindDataResponseBodyData {
	s.ShortUrl = &v
	return s
}

func (s *RequestBindDataResponseBodyData) SetUrl(v string) *RequestBindDataResponseBodyData {
	s.Url = &v
	return s
}

type RequestBindDataResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestBindDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestBindDataResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestBindDataResponse) GoString() string {
	return s.String()
}

func (s *RequestBindDataResponse) SetHeaders(v map[string]*string) *RequestBindDataResponse {
	s.Headers = v
	return s
}

func (s *RequestBindDataResponse) SetStatusCode(v int32) *RequestBindDataResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestBindDataResponse) SetBody(v *RequestBindDataResponseBody) *RequestBindDataResponse {
	s.Body = v
	return s
}

type RequestBindingStateResponseBody struct {
	Data         *RequestBindingStateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequestBindingStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestBindingStateResponseBody) GoString() string {
	return s.String()
}

func (s *RequestBindingStateResponseBody) SetData(v *RequestBindingStateResponseBodyData) *RequestBindingStateResponseBody {
	s.Data = v
	return s
}

func (s *RequestBindingStateResponseBody) SetErrorCode(v string) *RequestBindingStateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RequestBindingStateResponseBody) SetErrorMessage(v string) *RequestBindingStateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RequestBindingStateResponseBody) SetRequestId(v string) *RequestBindingStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestBindingStateResponseBody) SetSuccess(v bool) *RequestBindingStateResponseBody {
	s.Success = &v
	return s
}

type RequestBindingStateResponseBodyData struct {
	BindAt     *int64  `json:"BindAt,omitempty" xml:"BindAt,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserAvatar *string `json:"UserAvatar,omitempty" xml:"UserAvatar,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserNick   *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
	UserSource *string `json:"UserSource,omitempty" xml:"UserSource,omitempty"`
}

func (s RequestBindingStateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RequestBindingStateResponseBodyData) GoString() string {
	return s.String()
}

func (s *RequestBindingStateResponseBodyData) SetBindAt(v int64) *RequestBindingStateResponseBodyData {
	s.BindAt = &v
	return s
}

func (s *RequestBindingStateResponseBodyData) SetDeviceId(v string) *RequestBindingStateResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *RequestBindingStateResponseBodyData) SetUserAvatar(v string) *RequestBindingStateResponseBodyData {
	s.UserAvatar = &v
	return s
}

func (s *RequestBindingStateResponseBodyData) SetUserId(v string) *RequestBindingStateResponseBodyData {
	s.UserId = &v
	return s
}

func (s *RequestBindingStateResponseBodyData) SetUserNick(v string) *RequestBindingStateResponseBodyData {
	s.UserNick = &v
	return s
}

func (s *RequestBindingStateResponseBodyData) SetUserSource(v string) *RequestBindingStateResponseBodyData {
	s.UserSource = &v
	return s
}

type RequestBindingStateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestBindingStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestBindingStateResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestBindingStateResponse) GoString() string {
	return s.String()
}

func (s *RequestBindingStateResponse) SetHeaders(v map[string]*string) *RequestBindingStateResponse {
	s.Headers = v
	return s
}

func (s *RequestBindingStateResponse) SetStatusCode(v int32) *RequestBindingStateResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestBindingStateResponse) SetBody(v *RequestBindingStateResponseBody) *RequestBindingStateResponse {
	s.Body = v
	return s
}

type RequestDeviceInfoResponseBody struct {
	Data         *RequestDeviceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequestDeviceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *RequestDeviceInfoResponseBody) SetData(v *RequestDeviceInfoResponseBodyData) *RequestDeviceInfoResponseBody {
	s.Data = v
	return s
}

func (s *RequestDeviceInfoResponseBody) SetErrorCode(v string) *RequestDeviceInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RequestDeviceInfoResponseBody) SetErrorMessage(v string) *RequestDeviceInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RequestDeviceInfoResponseBody) SetRequestId(v string) *RequestDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestDeviceInfoResponseBody) SetSuccess(v bool) *RequestDeviceInfoResponseBody {
	s.Success = &v
	return s
}

type RequestDeviceInfoResponseBodyData struct {
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceSn   *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	PublicIp   *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
}

func (s RequestDeviceInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RequestDeviceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *RequestDeviceInfoResponseBodyData) SetDeviceId(v string) *RequestDeviceInfoResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *RequestDeviceInfoResponseBodyData) SetDeviceName(v string) *RequestDeviceInfoResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *RequestDeviceInfoResponseBodyData) SetDeviceSn(v string) *RequestDeviceInfoResponseBodyData {
	s.DeviceSn = &v
	return s
}

func (s *RequestDeviceInfoResponseBodyData) SetPublicIp(v string) *RequestDeviceInfoResponseBodyData {
	s.PublicIp = &v
	return s
}

type RequestDeviceInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestDeviceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *RequestDeviceInfoResponse) SetHeaders(v map[string]*string) *RequestDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *RequestDeviceInfoResponse) SetStatusCode(v int32) *RequestDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestDeviceInfoResponse) SetBody(v *RequestDeviceInfoResponseBody) *RequestDeviceInfoResponse {
	s.Body = v
	return s
}

type RequestIotTriadResponseBody struct {
	Data         *RequestIotTriadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequestIotTriadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestIotTriadResponseBody) GoString() string {
	return s.String()
}

func (s *RequestIotTriadResponseBody) SetData(v *RequestIotTriadResponseBodyData) *RequestIotTriadResponseBody {
	s.Data = v
	return s
}

func (s *RequestIotTriadResponseBody) SetErrorCode(v string) *RequestIotTriadResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RequestIotTriadResponseBody) SetErrorMessage(v string) *RequestIotTriadResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RequestIotTriadResponseBody) SetRequestId(v string) *RequestIotTriadResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestIotTriadResponseBody) SetSuccess(v bool) *RequestIotTriadResponseBody {
	s.Success = &v
	return s
}

type RequestIotTriadResponseBodyData struct {
	DeviceName   *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceSecret *string `json:"DeviceSecret,omitempty" xml:"DeviceSecret,omitempty"`
	ProductKey   *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s RequestIotTriadResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RequestIotTriadResponseBodyData) GoString() string {
	return s.String()
}

func (s *RequestIotTriadResponseBodyData) SetDeviceName(v string) *RequestIotTriadResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *RequestIotTriadResponseBodyData) SetDeviceSecret(v string) *RequestIotTriadResponseBodyData {
	s.DeviceSecret = &v
	return s
}

func (s *RequestIotTriadResponseBodyData) SetProductKey(v string) *RequestIotTriadResponseBodyData {
	s.ProductKey = &v
	return s
}

type RequestIotTriadResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestIotTriadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestIotTriadResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestIotTriadResponse) GoString() string {
	return s.String()
}

func (s *RequestIotTriadResponse) SetHeaders(v map[string]*string) *RequestIotTriadResponse {
	s.Headers = v
	return s
}

func (s *RequestIotTriadResponse) SetStatusCode(v int32) *RequestIotTriadResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestIotTriadResponse) SetBody(v *RequestIotTriadResponseBody) *RequestIotTriadResponse {
	s.Body = v
	return s
}

type RequestLiveSellPointStateResponseBody struct {
	Data         *RequestLiveSellPointStateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequestLiveSellPointStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestLiveSellPointStateResponseBody) GoString() string {
	return s.String()
}

func (s *RequestLiveSellPointStateResponseBody) SetData(v *RequestLiveSellPointStateResponseBodyData) *RequestLiveSellPointStateResponseBody {
	s.Data = v
	return s
}

func (s *RequestLiveSellPointStateResponseBody) SetErrorCode(v string) *RequestLiveSellPointStateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RequestLiveSellPointStateResponseBody) SetErrorMessage(v string) *RequestLiveSellPointStateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RequestLiveSellPointStateResponseBody) SetRequestId(v string) *RequestLiveSellPointStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestLiveSellPointStateResponseBody) SetSuccess(v bool) *RequestLiveSellPointStateResponseBody {
	s.Success = &v
	return s
}

type RequestLiveSellPointStateResponseBodyData struct {
	Display *bool `json:"Display,omitempty" xml:"Display,omitempty"`
}

func (s RequestLiveSellPointStateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RequestLiveSellPointStateResponseBodyData) GoString() string {
	return s.String()
}

func (s *RequestLiveSellPointStateResponseBodyData) SetDisplay(v bool) *RequestLiveSellPointStateResponseBodyData {
	s.Display = &v
	return s
}

type RequestLiveSellPointStateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestLiveSellPointStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestLiveSellPointStateResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestLiveSellPointStateResponse) GoString() string {
	return s.String()
}

func (s *RequestLiveSellPointStateResponse) SetHeaders(v map[string]*string) *RequestLiveSellPointStateResponse {
	s.Headers = v
	return s
}

func (s *RequestLiveSellPointStateResponse) SetStatusCode(v int32) *RequestLiveSellPointStateResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestLiveSellPointStateResponse) SetBody(v *RequestLiveSellPointStateResponseBody) *RequestLiveSellPointStateResponse {
	s.Body = v
	return s
}

type RequestLiveTeleprompterStateResponseBody struct {
	Data         *RequestLiveTeleprompterStateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequestLiveTeleprompterStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestLiveTeleprompterStateResponseBody) GoString() string {
	return s.String()
}

func (s *RequestLiveTeleprompterStateResponseBody) SetData(v *RequestLiveTeleprompterStateResponseBodyData) *RequestLiveTeleprompterStateResponseBody {
	s.Data = v
	return s
}

func (s *RequestLiveTeleprompterStateResponseBody) SetErrorCode(v string) *RequestLiveTeleprompterStateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RequestLiveTeleprompterStateResponseBody) SetErrorMessage(v string) *RequestLiveTeleprompterStateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RequestLiveTeleprompterStateResponseBody) SetRequestId(v string) *RequestLiveTeleprompterStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestLiveTeleprompterStateResponseBody) SetSuccess(v bool) *RequestLiveTeleprompterStateResponseBody {
	s.Success = &v
	return s
}

type RequestLiveTeleprompterStateResponseBodyData struct {
	Display *bool `json:"Display,omitempty" xml:"Display,omitempty"`
}

func (s RequestLiveTeleprompterStateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RequestLiveTeleprompterStateResponseBodyData) GoString() string {
	return s.String()
}

func (s *RequestLiveTeleprompterStateResponseBodyData) SetDisplay(v bool) *RequestLiveTeleprompterStateResponseBodyData {
	s.Display = &v
	return s
}

type RequestLiveTeleprompterStateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestLiveTeleprompterStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestLiveTeleprompterStateResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestLiveTeleprompterStateResponse) GoString() string {
	return s.String()
}

func (s *RequestLiveTeleprompterStateResponse) SetHeaders(v map[string]*string) *RequestLiveTeleprompterStateResponse {
	s.Headers = v
	return s
}

func (s *RequestLiveTeleprompterStateResponse) SetStatusCode(v int32) *RequestLiveTeleprompterStateResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestLiveTeleprompterStateResponse) SetBody(v *RequestLiveTeleprompterStateResponseBody) *RequestLiveTeleprompterStateResponse {
	s.Body = v
	return s
}

type RequestOssStsRequest struct {
	AppCode       *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	ExpireSeconds *int64  `json:"ExpireSeconds,omitempty" xml:"ExpireSeconds,omitempty"`
}

func (s RequestOssStsRequest) String() string {
	return tea.Prettify(s)
}

func (s RequestOssStsRequest) GoString() string {
	return s.String()
}

func (s *RequestOssStsRequest) SetAppCode(v string) *RequestOssStsRequest {
	s.AppCode = &v
	return s
}

func (s *RequestOssStsRequest) SetExpireSeconds(v int64) *RequestOssStsRequest {
	s.ExpireSeconds = &v
	return s
}

type RequestOssStsResponseBody struct {
	Data         *RequestOssStsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequestOssStsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestOssStsResponseBody) GoString() string {
	return s.String()
}

func (s *RequestOssStsResponseBody) SetData(v *RequestOssStsResponseBodyData) *RequestOssStsResponseBody {
	s.Data = v
	return s
}

func (s *RequestOssStsResponseBody) SetErrorCode(v string) *RequestOssStsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RequestOssStsResponseBody) SetErrorMessage(v string) *RequestOssStsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RequestOssStsResponseBody) SetRequestId(v string) *RequestOssStsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestOssStsResponseBody) SetSuccess(v bool) *RequestOssStsResponseBody {
	s.Success = &v
	return s
}

type RequestOssStsResponseBodyData struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Bucket          *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	EndPoint        *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	Expire          *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	ObjectKeyPrefix *string `json:"ObjectKeyPrefix,omitempty" xml:"ObjectKeyPrefix,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RequestOssStsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RequestOssStsResponseBodyData) GoString() string {
	return s.String()
}

func (s *RequestOssStsResponseBodyData) SetAccessKeyId(v string) *RequestOssStsResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *RequestOssStsResponseBodyData) SetAccessKeySecret(v string) *RequestOssStsResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *RequestOssStsResponseBodyData) SetBucket(v string) *RequestOssStsResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *RequestOssStsResponseBodyData) SetEndPoint(v string) *RequestOssStsResponseBodyData {
	s.EndPoint = &v
	return s
}

func (s *RequestOssStsResponseBodyData) SetExpire(v string) *RequestOssStsResponseBodyData {
	s.Expire = &v
	return s
}

func (s *RequestOssStsResponseBodyData) SetObjectKeyPrefix(v string) *RequestOssStsResponseBodyData {
	s.ObjectKeyPrefix = &v
	return s
}

func (s *RequestOssStsResponseBodyData) SetSecurityToken(v string) *RequestOssStsResponseBodyData {
	s.SecurityToken = &v
	return s
}

type RequestOssStsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestOssStsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestOssStsResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestOssStsResponse) GoString() string {
	return s.String()
}

func (s *RequestOssStsResponse) SetHeaders(v map[string]*string) *RequestOssStsResponse {
	s.Headers = v
	return s
}

func (s *RequestOssStsResponse) SetStatusCode(v int32) *RequestOssStsResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestOssStsResponse) SetBody(v *RequestOssStsResponseBody) *RequestOssStsResponse {
	s.Body = v
	return s
}

type RequestPasterResponseBody struct {
	Data         []*RequestPasterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorCode    *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequestPasterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestPasterResponseBody) GoString() string {
	return s.String()
}

func (s *RequestPasterResponseBody) SetData(v []*RequestPasterResponseBodyData) *RequestPasterResponseBody {
	s.Data = v
	return s
}

func (s *RequestPasterResponseBody) SetErrorCode(v string) *RequestPasterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RequestPasterResponseBody) SetErrorMessage(v string) *RequestPasterResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RequestPasterResponseBody) SetRequestId(v string) *RequestPasterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestPasterResponseBody) SetSuccess(v bool) *RequestPasterResponseBody {
	s.Success = &v
	return s
}

type RequestPasterResponseBodyData struct {
	Configs      map[string]interface{} `json:"Configs,omitempty" xml:"Configs,omitempty"`
	DownloadUrl  *string                `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ResourceUuid *string                `json:"ResourceUuid,omitempty" xml:"ResourceUuid,omitempty"`
}

func (s RequestPasterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RequestPasterResponseBodyData) GoString() string {
	return s.String()
}

func (s *RequestPasterResponseBodyData) SetConfigs(v map[string]interface{}) *RequestPasterResponseBodyData {
	s.Configs = v
	return s
}

func (s *RequestPasterResponseBodyData) SetDownloadUrl(v string) *RequestPasterResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *RequestPasterResponseBodyData) SetResourceUuid(v string) *RequestPasterResponseBodyData {
	s.ResourceUuid = &v
	return s
}

type RequestPasterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestPasterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestPasterResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestPasterResponse) GoString() string {
	return s.String()
}

func (s *RequestPasterResponse) SetHeaders(v map[string]*string) *RequestPasterResponse {
	s.Headers = v
	return s
}

func (s *RequestPasterResponse) SetStatusCode(v int32) *RequestPasterResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestPasterResponse) SetBody(v *RequestPasterResponseBody) *RequestPasterResponse {
	s.Body = v
	return s
}

type RequestResetDataRequest struct {
	LiveSource *string `json:"LiveSource,omitempty" xml:"LiveSource,omitempty"`
}

func (s RequestResetDataRequest) String() string {
	return tea.Prettify(s)
}

func (s RequestResetDataRequest) GoString() string {
	return s.String()
}

func (s *RequestResetDataRequest) SetLiveSource(v string) *RequestResetDataRequest {
	s.LiveSource = &v
	return s
}

type RequestResetDataResponseBody struct {
	Data         *RequestResetDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequestResetDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestResetDataResponseBody) GoString() string {
	return s.String()
}

func (s *RequestResetDataResponseBody) SetData(v *RequestResetDataResponseBodyData) *RequestResetDataResponseBody {
	s.Data = v
	return s
}

func (s *RequestResetDataResponseBody) SetErrorCode(v string) *RequestResetDataResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RequestResetDataResponseBody) SetErrorMessage(v string) *RequestResetDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RequestResetDataResponseBody) SetRequestId(v string) *RequestResetDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestResetDataResponseBody) SetSuccess(v bool) *RequestResetDataResponseBody {
	s.Success = &v
	return s
}

type RequestResetDataResponseBodyData struct {
	FullUrl *string `json:"FullUrl,omitempty" xml:"FullUrl,omitempty"`
	Url     *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RequestResetDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RequestResetDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *RequestResetDataResponseBodyData) SetFullUrl(v string) *RequestResetDataResponseBodyData {
	s.FullUrl = &v
	return s
}

func (s *RequestResetDataResponseBodyData) SetUrl(v string) *RequestResetDataResponseBodyData {
	s.Url = &v
	return s
}

type RequestResetDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestResetDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestResetDataResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestResetDataResponse) GoString() string {
	return s.String()
}

func (s *RequestResetDataResponse) SetHeaders(v map[string]*string) *RequestResetDataResponse {
	s.Headers = v
	return s
}

func (s *RequestResetDataResponse) SetStatusCode(v int32) *RequestResetDataResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestResetDataResponse) SetBody(v *RequestResetDataResponseBody) *RequestResetDataResponse {
	s.Body = v
	return s
}

type RequestServiceInfoResponseBody struct {
	Data         *RequestServiceInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequestServiceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestServiceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *RequestServiceInfoResponseBody) SetData(v *RequestServiceInfoResponseBodyData) *RequestServiceInfoResponseBody {
	s.Data = v
	return s
}

func (s *RequestServiceInfoResponseBody) SetErrorCode(v string) *RequestServiceInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RequestServiceInfoResponseBody) SetErrorMessage(v string) *RequestServiceInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RequestServiceInfoResponseBody) SetRequestId(v string) *RequestServiceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestServiceInfoResponseBody) SetSuccess(v bool) *RequestServiceInfoResponseBody {
	s.Success = &v
	return s
}

type RequestServiceInfoResponseBodyData struct {
	ServiceEffectAt *int64  `json:"ServiceEffectAt,omitempty" xml:"ServiceEffectAt,omitempty"`
	ServiceExpireAt *int64  `json:"ServiceExpireAt,omitempty" xml:"ServiceExpireAt,omitempty"`
	ServicePackName *string `json:"ServicePackName,omitempty" xml:"ServicePackName,omitempty"`
}

func (s RequestServiceInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RequestServiceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *RequestServiceInfoResponseBodyData) SetServiceEffectAt(v int64) *RequestServiceInfoResponseBodyData {
	s.ServiceEffectAt = &v
	return s
}

func (s *RequestServiceInfoResponseBodyData) SetServiceExpireAt(v int64) *RequestServiceInfoResponseBodyData {
	s.ServiceExpireAt = &v
	return s
}

func (s *RequestServiceInfoResponseBodyData) SetServicePackName(v string) *RequestServiceInfoResponseBodyData {
	s.ServicePackName = &v
	return s
}

type RequestServiceInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestServiceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestServiceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestServiceInfoResponse) GoString() string {
	return s.String()
}

func (s *RequestServiceInfoResponse) SetHeaders(v map[string]*string) *RequestServiceInfoResponse {
	s.Headers = v
	return s
}

func (s *RequestServiceInfoResponse) SetStatusCode(v int32) *RequestServiceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestServiceInfoResponse) SetBody(v *RequestServiceInfoResponseBody) *RequestServiceInfoResponse {
	s.Body = v
	return s
}

type RequestUserConfigRequest struct {
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s RequestUserConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s RequestUserConfigRequest) GoString() string {
	return s.String()
}

func (s *RequestUserConfigRequest) SetKey(v string) *RequestUserConfigRequest {
	s.Key = &v
	return s
}

type RequestUserConfigResponseBody struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequestUserConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *RequestUserConfigResponseBody) SetData(v string) *RequestUserConfigResponseBody {
	s.Data = &v
	return s
}

func (s *RequestUserConfigResponseBody) SetErrorCode(v string) *RequestUserConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RequestUserConfigResponseBody) SetErrorMessage(v string) *RequestUserConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RequestUserConfigResponseBody) SetRequestId(v string) *RequestUserConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestUserConfigResponseBody) SetSuccess(v bool) *RequestUserConfigResponseBody {
	s.Success = &v
	return s
}

type RequestUserConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestUserConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestUserConfigResponse) GoString() string {
	return s.String()
}

func (s *RequestUserConfigResponse) SetHeaders(v map[string]*string) *RequestUserConfigResponse {
	s.Headers = v
	return s
}

func (s *RequestUserConfigResponse) SetStatusCode(v int32) *RequestUserConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestUserConfigResponse) SetBody(v *RequestUserConfigResponseBody) *RequestUserConfigResponse {
	s.Body = v
	return s
}

type RequestUserSellPointTemplateResponseBody struct {
	Data         *RequestUserSellPointTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RequestUserSellPointTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestUserSellPointTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *RequestUserSellPointTemplateResponseBody) SetData(v *RequestUserSellPointTemplateResponseBodyData) *RequestUserSellPointTemplateResponseBody {
	s.Data = v
	return s
}

func (s *RequestUserSellPointTemplateResponseBody) SetErrorCode(v string) *RequestUserSellPointTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RequestUserSellPointTemplateResponseBody) SetErrorMessage(v string) *RequestUserSellPointTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RequestUserSellPointTemplateResponseBody) SetRequestId(v string) *RequestUserSellPointTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestUserSellPointTemplateResponseBody) SetSuccess(v bool) *RequestUserSellPointTemplateResponseBody {
	s.Success = &v
	return s
}

type RequestUserSellPointTemplateResponseBodyData struct {
	DisplayConfig  map[string]interface{} `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty"`
	TemplateConfig map[string]interface{} `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	TemplateUuid   *string                `json:"TemplateUuid,omitempty" xml:"TemplateUuid,omitempty"`
	Url            *string                `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RequestUserSellPointTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RequestUserSellPointTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *RequestUserSellPointTemplateResponseBodyData) SetDisplayConfig(v map[string]interface{}) *RequestUserSellPointTemplateResponseBodyData {
	s.DisplayConfig = v
	return s
}

func (s *RequestUserSellPointTemplateResponseBodyData) SetTemplateConfig(v map[string]interface{}) *RequestUserSellPointTemplateResponseBodyData {
	s.TemplateConfig = v
	return s
}

func (s *RequestUserSellPointTemplateResponseBodyData) SetTemplateUuid(v string) *RequestUserSellPointTemplateResponseBodyData {
	s.TemplateUuid = &v
	return s
}

func (s *RequestUserSellPointTemplateResponseBodyData) SetUrl(v string) *RequestUserSellPointTemplateResponseBodyData {
	s.Url = &v
	return s
}

type RequestUserSellPointTemplateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestUserSellPointTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestUserSellPointTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestUserSellPointTemplateResponse) GoString() string {
	return s.String()
}

func (s *RequestUserSellPointTemplateResponse) SetHeaders(v map[string]*string) *RequestUserSellPointTemplateResponse {
	s.Headers = v
	return s
}

func (s *RequestUserSellPointTemplateResponse) SetStatusCode(v int32) *RequestUserSellPointTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestUserSellPointTemplateResponse) SetBody(v *RequestUserSellPointTemplateResponseBody) *RequestUserSellPointTemplateResponse {
	s.Body = v
	return s
}

type ResetDeviceRequest struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ResetDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetDeviceRequest) GoString() string {
	return s.String()
}

func (s *ResetDeviceRequest) SetCode(v string) *ResetDeviceRequest {
	s.Code = &v
	return s
}

type ResetDeviceResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResetDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ResetDeviceResponseBody) SetErrorCode(v string) *ResetDeviceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ResetDeviceResponseBody) SetErrorMessage(v string) *ResetDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ResetDeviceResponseBody) SetRequestId(v string) *ResetDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetDeviceResponseBody) SetSuccess(v bool) *ResetDeviceResponseBody {
	s.Success = &v
	return s
}

type ResetDeviceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetDeviceResponse) GoString() string {
	return s.String()
}

func (s *ResetDeviceResponse) SetHeaders(v map[string]*string) *ResetDeviceResponse {
	s.Headers = v
	return s
}

func (s *ResetDeviceResponse) SetStatusCode(v int32) *ResetDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetDeviceResponse) SetBody(v *ResetDeviceResponseBody) *ResetDeviceResponse {
	s.Body = v
	return s
}

type UpdateCurrentItemRequest struct {
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
}

func (s UpdateCurrentItemRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCurrentItemRequest) GoString() string {
	return s.String()
}

func (s *UpdateCurrentItemRequest) SetItemId(v string) *UpdateCurrentItemRequest {
	s.ItemId = &v
	return s
}

type UpdateCurrentItemResponseBody struct {
	Data         *UpdateCurrentItemResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCurrentItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCurrentItemResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCurrentItemResponseBody) SetData(v *UpdateCurrentItemResponseBodyData) *UpdateCurrentItemResponseBody {
	s.Data = v
	return s
}

func (s *UpdateCurrentItemResponseBody) SetErrorCode(v string) *UpdateCurrentItemResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateCurrentItemResponseBody) SetErrorMessage(v string) *UpdateCurrentItemResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateCurrentItemResponseBody) SetRequestId(v string) *UpdateCurrentItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCurrentItemResponseBody) SetSuccess(v bool) *UpdateCurrentItemResponseBody {
	s.Success = &v
	return s
}

type UpdateCurrentItemResponseBodyData struct {
	ItemBackground *UpdateCurrentItemResponseBodyDataItemBackground `json:"ItemBackground,omitempty" xml:"ItemBackground,omitempty" type:"Struct"`
}

func (s UpdateCurrentItemResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateCurrentItemResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateCurrentItemResponseBodyData) SetItemBackground(v *UpdateCurrentItemResponseBodyDataItemBackground) *UpdateCurrentItemResponseBodyData {
	s.ItemBackground = v
	return s
}

type UpdateCurrentItemResponseBodyDataItemBackground struct {
	DownloadUrl  *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	FileType     *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	ItemId       *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ResourceUuid *string `json:"ResourceUuid,omitempty" xml:"ResourceUuid,omitempty"`
	Scope        *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s UpdateCurrentItemResponseBodyDataItemBackground) String() string {
	return tea.Prettify(s)
}

func (s UpdateCurrentItemResponseBodyDataItemBackground) GoString() string {
	return s.String()
}

func (s *UpdateCurrentItemResponseBodyDataItemBackground) SetDownloadUrl(v string) *UpdateCurrentItemResponseBodyDataItemBackground {
	s.DownloadUrl = &v
	return s
}

func (s *UpdateCurrentItemResponseBodyDataItemBackground) SetFileType(v string) *UpdateCurrentItemResponseBodyDataItemBackground {
	s.FileType = &v
	return s
}

func (s *UpdateCurrentItemResponseBodyDataItemBackground) SetItemId(v string) *UpdateCurrentItemResponseBodyDataItemBackground {
	s.ItemId = &v
	return s
}

func (s *UpdateCurrentItemResponseBodyDataItemBackground) SetResourceUuid(v string) *UpdateCurrentItemResponseBodyDataItemBackground {
	s.ResourceUuid = &v
	return s
}

func (s *UpdateCurrentItemResponseBodyDataItemBackground) SetScope(v string) *UpdateCurrentItemResponseBodyDataItemBackground {
	s.Scope = &v
	return s
}

type UpdateCurrentItemResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCurrentItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCurrentItemResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCurrentItemResponse) GoString() string {
	return s.String()
}

func (s *UpdateCurrentItemResponse) SetHeaders(v map[string]*string) *UpdateCurrentItemResponse {
	s.Headers = v
	return s
}

func (s *UpdateCurrentItemResponse) SetStatusCode(v int32) *UpdateCurrentItemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCurrentItemResponse) SetBody(v *UpdateCurrentItemResponseBody) *UpdateCurrentItemResponse {
	s.Body = v
	return s
}

type UpdateLiveSellPointStateRequest struct {
	Display *bool `json:"Display,omitempty" xml:"Display,omitempty"`
}

func (s UpdateLiveSellPointStateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveSellPointStateRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveSellPointStateRequest) SetDisplay(v bool) *UpdateLiveSellPointStateRequest {
	s.Display = &v
	return s
}

type UpdateLiveSellPointStateResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateLiveSellPointStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveSellPointStateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveSellPointStateResponseBody) SetErrorCode(v string) *UpdateLiveSellPointStateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateLiveSellPointStateResponseBody) SetErrorMessage(v string) *UpdateLiveSellPointStateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateLiveSellPointStateResponseBody) SetRequestId(v string) *UpdateLiveSellPointStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveSellPointStateResponseBody) SetSuccess(v bool) *UpdateLiveSellPointStateResponseBody {
	s.Success = &v
	return s
}

type UpdateLiveSellPointStateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveSellPointStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveSellPointStateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveSellPointStateResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveSellPointStateResponse) SetHeaders(v map[string]*string) *UpdateLiveSellPointStateResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveSellPointStateResponse) SetStatusCode(v int32) *UpdateLiveSellPointStateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveSellPointStateResponse) SetBody(v *UpdateLiveSellPointStateResponseBody) *UpdateLiveSellPointStateResponse {
	s.Body = v
	return s
}

type UpdateLiveTeleprompterStateRequest struct {
	Display *bool `json:"Display,omitempty" xml:"Display,omitempty"`
}

func (s UpdateLiveTeleprompterStateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveTeleprompterStateRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveTeleprompterStateRequest) SetDisplay(v bool) *UpdateLiveTeleprompterStateRequest {
	s.Display = &v
	return s
}

type UpdateLiveTeleprompterStateResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateLiveTeleprompterStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveTeleprompterStateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveTeleprompterStateResponseBody) SetErrorCode(v string) *UpdateLiveTeleprompterStateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateLiveTeleprompterStateResponseBody) SetErrorMessage(v string) *UpdateLiveTeleprompterStateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateLiveTeleprompterStateResponseBody) SetRequestId(v string) *UpdateLiveTeleprompterStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveTeleprompterStateResponseBody) SetSuccess(v bool) *UpdateLiveTeleprompterStateResponseBody {
	s.Success = &v
	return s
}

type UpdateLiveTeleprompterStateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveTeleprompterStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveTeleprompterStateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveTeleprompterStateResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveTeleprompterStateResponse) SetHeaders(v map[string]*string) *UpdateLiveTeleprompterStateResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveTeleprompterStateResponse) SetStatusCode(v int32) *UpdateLiveTeleprompterStateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveTeleprompterStateResponse) SetBody(v *UpdateLiveTeleprompterStateResponseBody) *UpdateLiveTeleprompterStateResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("mindlive"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) LoginDeviceWithOptions(request *LoginDeviceRequest, runtime *util.RuntimeOptions) (_result *LoginDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserSource)) {
		query["UserSource"] = request.UserSource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LoginDevice"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LoginDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LoginDevice(request *LoginDeviceRequest) (_result *LoginDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LoginDeviceResponse{}
	_body, _err := client.LoginDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LogoutDeviceWithOptions(request *LogoutDeviceRequest, runtime *util.RuntimeOptions) (_result *LogoutDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserSource)) {
		query["UserSource"] = request.UserSource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LogoutDevice"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LogoutDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LogoutDevice(request *LogoutDeviceRequest) (_result *LogoutDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LogoutDeviceResponse{}
	_body, _err := client.LogoutDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryItemBackgroundsWithOptions(tmpReq *QueryItemBackgroundsRequest, runtime *util.RuntimeOptions) (_result *QueryItemBackgroundsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryItemBackgroundsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ItemIds)) {
		request.ItemIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItemIds, tea.String("ItemIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ItemIdsShrink)) {
		query["ItemIds"] = request.ItemIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryItemBackgrounds"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryItemBackgroundsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryItemBackgrounds(request *QueryItemBackgroundsRequest) (_result *QueryItemBackgroundsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryItemBackgroundsResponse{}
	_body, _err := client.QueryItemBackgroundsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportCurrentBackgroundWithOptions(tmpReq *ReportCurrentBackgroundRequest, runtime *util.RuntimeOptions) (_result *ReportCurrentBackgroundResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ReportCurrentBackgroundShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.BgConfig)) {
		request.BgConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BgConfig, tea.String("BgConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BgConfigShrink)) {
		query["BgConfig"] = request.BgConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.Open)) {
		query["Open"] = request.Open
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceUuid)) {
		query["ResourceUuid"] = request.ResourceUuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportCurrentBackground"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportCurrentBackgroundResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportCurrentBackground(request *ReportCurrentBackgroundRequest) (_result *ReportCurrentBackgroundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportCurrentBackgroundResponse{}
	_body, _err := client.ReportCurrentBackgroundWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportDevConfigWithOptions(request *ReportDevConfigRequest, runtime *util.RuntimeOptions) (_result *ReportDevConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Configs)) {
		query["Configs"] = request.Configs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportDevConfig"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportDevConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportDevConfig(request *ReportDevConfigRequest) (_result *ReportDevConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportDevConfigResponse{}
	_body, _err := client.ReportDevConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportLiveStateWithOptions(request *ReportLiveStateRequest, runtime *util.RuntimeOptions) (_result *ReportLiveStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnchorId)) {
		query["AnchorId"] = request.AnchorId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.LiveMode)) {
		query["LiveMode"] = request.LiveMode
	}

	if !tea.BoolValue(util.IsUnset(request.LiveState)) {
		query["LiveState"] = request.LiveState
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportLiveState"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportLiveStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportLiveState(request *ReportLiveStateRequest) (_result *ReportLiveStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportLiveStateResponse{}
	_body, _err := client.ReportLiveStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportScreenWithOptions(request *ReportScreenRequest, runtime *util.RuntimeOptions) (_result *ReportScreenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OssBucketName)) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.OssObjectKey)) {
		query["OssObjectKey"] = request.OssObjectKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportScreen"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportScreenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportScreen(request *ReportScreenRequest) (_result *ReportScreenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportScreenResponse{}
	_body, _err := client.ReportScreenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportUserConfigWithOptions(request *ReportUserConfigRequest, runtime *util.RuntimeOptions) (_result *ReportUserConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportUserConfig"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportUserConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportUserConfig(request *ReportUserConfigRequest) (_result *ReportUserConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportUserConfigResponse{}
	_body, _err := client.ReportUserConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RequestAuthorizationDataWithOptions(runtime *util.RuntimeOptions) (_result *RequestAuthorizationDataResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("RequestAuthorizationData"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestAuthorizationDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RequestAuthorizationData() (_result *RequestAuthorizationDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestAuthorizationDataResponse{}
	_body, _err := client.RequestAuthorizationDataWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RequestBackgroundWithOptions(runtime *util.RuntimeOptions) (_result *RequestBackgroundResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("RequestBackground"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestBackgroundResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RequestBackground() (_result *RequestBackgroundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestBackgroundResponse{}
	_body, _err := client.RequestBackgroundWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RequestBindDataWithOptions(request *RequestBindDataRequest, runtime *util.RuntimeOptions) (_result *RequestBindDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveSource)) {
		query["LiveSource"] = request.LiveSource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RequestBindData"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestBindDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RequestBindData(request *RequestBindDataRequest) (_result *RequestBindDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestBindDataResponse{}
	_body, _err := client.RequestBindDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RequestBindingStateWithOptions(runtime *util.RuntimeOptions) (_result *RequestBindingStateResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("RequestBindingState"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestBindingStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RequestBindingState() (_result *RequestBindingStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestBindingStateResponse{}
	_body, _err := client.RequestBindingStateWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RequestDeviceInfoWithOptions(runtime *util.RuntimeOptions) (_result *RequestDeviceInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("RequestDeviceInfo"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RequestDeviceInfo() (_result *RequestDeviceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestDeviceInfoResponse{}
	_body, _err := client.RequestDeviceInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RequestIotTriadWithOptions(runtime *util.RuntimeOptions) (_result *RequestIotTriadResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("RequestIotTriad"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestIotTriadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RequestIotTriad() (_result *RequestIotTriadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestIotTriadResponse{}
	_body, _err := client.RequestIotTriadWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RequestLiveSellPointStateWithOptions(runtime *util.RuntimeOptions) (_result *RequestLiveSellPointStateResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("RequestLiveSellPointState"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestLiveSellPointStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RequestLiveSellPointState() (_result *RequestLiveSellPointStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestLiveSellPointStateResponse{}
	_body, _err := client.RequestLiveSellPointStateWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RequestLiveTeleprompterStateWithOptions(runtime *util.RuntimeOptions) (_result *RequestLiveTeleprompterStateResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("RequestLiveTeleprompterState"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestLiveTeleprompterStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RequestLiveTeleprompterState() (_result *RequestLiveTeleprompterStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestLiveTeleprompterStateResponse{}
	_body, _err := client.RequestLiveTeleprompterStateWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RequestOssStsWithOptions(request *RequestOssStsRequest, runtime *util.RuntimeOptions) (_result *RequestOssStsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCode)) {
		query["AppCode"] = request.AppCode
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireSeconds)) {
		query["ExpireSeconds"] = request.ExpireSeconds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RequestOssSts"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestOssStsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RequestOssSts(request *RequestOssStsRequest) (_result *RequestOssStsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestOssStsResponse{}
	_body, _err := client.RequestOssStsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RequestPasterWithOptions(runtime *util.RuntimeOptions) (_result *RequestPasterResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("RequestPaster"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestPasterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RequestPaster() (_result *RequestPasterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestPasterResponse{}
	_body, _err := client.RequestPasterWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RequestResetDataWithOptions(request *RequestResetDataRequest, runtime *util.RuntimeOptions) (_result *RequestResetDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LiveSource)) {
		query["LiveSource"] = request.LiveSource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RequestResetData"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestResetDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RequestResetData(request *RequestResetDataRequest) (_result *RequestResetDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestResetDataResponse{}
	_body, _err := client.RequestResetDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RequestServiceInfoWithOptions(runtime *util.RuntimeOptions) (_result *RequestServiceInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("RequestServiceInfo"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestServiceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RequestServiceInfo() (_result *RequestServiceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestServiceInfoResponse{}
	_body, _err := client.RequestServiceInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RequestUserConfigWithOptions(request *RequestUserConfigRequest, runtime *util.RuntimeOptions) (_result *RequestUserConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RequestUserConfig"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestUserConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RequestUserConfig(request *RequestUserConfigRequest) (_result *RequestUserConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestUserConfigResponse{}
	_body, _err := client.RequestUserConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RequestUserSellPointTemplateWithOptions(runtime *util.RuntimeOptions) (_result *RequestUserSellPointTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("RequestUserSellPointTemplate"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestUserSellPointTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RequestUserSellPointTemplate() (_result *RequestUserSellPointTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestUserSellPointTemplateResponse{}
	_body, _err := client.RequestUserSellPointTemplateWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetDeviceWithOptions(request *ResetDeviceRequest, runtime *util.RuntimeOptions) (_result *ResetDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetDevice"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetDevice(request *ResetDeviceRequest) (_result *ResetDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetDeviceResponse{}
	_body, _err := client.ResetDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCurrentItemWithOptions(request *UpdateCurrentItemRequest, runtime *util.RuntimeOptions) (_result *UpdateCurrentItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ItemId)) {
		query["ItemId"] = request.ItemId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCurrentItem"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCurrentItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCurrentItem(request *UpdateCurrentItemRequest) (_result *UpdateCurrentItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCurrentItemResponse{}
	_body, _err := client.UpdateCurrentItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLiveSellPointStateWithOptions(request *UpdateLiveSellPointStateRequest, runtime *util.RuntimeOptions) (_result *UpdateLiveSellPointStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Display)) {
		query["Display"] = request.Display
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLiveSellPointState"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLiveSellPointStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLiveSellPointState(request *UpdateLiveSellPointStateRequest) (_result *UpdateLiveSellPointStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLiveSellPointStateResponse{}
	_body, _err := client.UpdateLiveSellPointStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLiveTeleprompterStateWithOptions(request *UpdateLiveTeleprompterStateRequest, runtime *util.RuntimeOptions) (_result *UpdateLiveTeleprompterStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Display)) {
		query["Display"] = request.Display
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLiveTeleprompterState"),
		Version:     tea.String("2021-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLiveTeleprompterStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLiveTeleprompterState(request *UpdateLiveTeleprompterStateRequest) (_result *UpdateLiveTeleprompterStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLiveTeleprompterStateResponse{}
	_body, _err := client.UpdateLiveTeleprompterStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
