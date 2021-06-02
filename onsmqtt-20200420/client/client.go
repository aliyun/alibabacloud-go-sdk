// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ApplyTokenRequest struct {
	Resources  *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Actions    *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
}

func (s ApplyTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyTokenRequest) GoString() string {
	return s.String()
}

func (s *ApplyTokenRequest) SetResources(v string) *ApplyTokenRequest {
	s.Resources = &v
	return s
}

func (s *ApplyTokenRequest) SetInstanceId(v string) *ApplyTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *ApplyTokenRequest) SetExpireTime(v int64) *ApplyTokenRequest {
	s.ExpireTime = &v
	return s
}

func (s *ApplyTokenRequest) SetActions(v string) *ApplyTokenRequest {
	s.Actions = &v
	return s
}

type ApplyTokenResponseBody struct {
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyTokenResponseBody) SetToken(v string) *ApplyTokenResponseBody {
	s.Token = &v
	return s
}

func (s *ApplyTokenResponseBody) SetRequestId(v string) *ApplyTokenResponseBody {
	s.RequestId = &v
	return s
}

type ApplyTokenResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyTokenResponse) GoString() string {
	return s.String()
}

func (s *ApplyTokenResponse) SetHeaders(v map[string]*string) *ApplyTokenResponse {
	s.Headers = v
	return s
}

func (s *ApplyTokenResponse) SetBody(v *ApplyTokenResponseBody) *ApplyTokenResponse {
	s.Body = v
	return s
}

type BatchQuerySessionByClientIdsRequest struct {
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ClientIdList []*string `json:"ClientIdList,omitempty" xml:"ClientIdList,omitempty" type:"Repeated"`
}

func (s BatchQuerySessionByClientIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQuerySessionByClientIdsRequest) GoString() string {
	return s.String()
}

func (s *BatchQuerySessionByClientIdsRequest) SetInstanceId(v string) *BatchQuerySessionByClientIdsRequest {
	s.InstanceId = &v
	return s
}

func (s *BatchQuerySessionByClientIdsRequest) SetClientIdList(v []*string) *BatchQuerySessionByClientIdsRequest {
	s.ClientIdList = v
	return s
}

type BatchQuerySessionByClientIdsResponseBody struct {
	RequestId        *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OnlineStatusList []*BatchQuerySessionByClientIdsResponseBodyOnlineStatusList `json:"OnlineStatusList,omitempty" xml:"OnlineStatusList,omitempty" type:"Repeated"`
}

func (s BatchQuerySessionByClientIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQuerySessionByClientIdsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQuerySessionByClientIdsResponseBody) SetRequestId(v string) *BatchQuerySessionByClientIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchQuerySessionByClientIdsResponseBody) SetOnlineStatusList(v []*BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) *BatchQuerySessionByClientIdsResponseBody {
	s.OnlineStatusList = v
	return s
}

type BatchQuerySessionByClientIdsResponseBodyOnlineStatusList struct {
	OnlineStatus *bool   `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) String() string {
	return tea.Prettify(s)
}

func (s BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) GoString() string {
	return s.String()
}

func (s *BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) SetOnlineStatus(v bool) *BatchQuerySessionByClientIdsResponseBodyOnlineStatusList {
	s.OnlineStatus = &v
	return s
}

func (s *BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) SetClientId(v string) *BatchQuerySessionByClientIdsResponseBodyOnlineStatusList {
	s.ClientId = &v
	return s
}

type BatchQuerySessionByClientIdsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchQuerySessionByClientIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchQuerySessionByClientIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQuerySessionByClientIdsResponse) GoString() string {
	return s.String()
}

func (s *BatchQuerySessionByClientIdsResponse) SetHeaders(v map[string]*string) *BatchQuerySessionByClientIdsResponse {
	s.Headers = v
	return s
}

func (s *BatchQuerySessionByClientIdsResponse) SetBody(v *BatchQuerySessionByClientIdsResponseBody) *BatchQuerySessionByClientIdsResponse {
	s.Body = v
	return s
}

type CreateGroupIdRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupIdRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupIdRequest) SetGroupId(v string) *CreateGroupIdRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupIdRequest) SetInstanceId(v string) *CreateGroupIdRequest {
	s.InstanceId = &v
	return s
}

type CreateGroupIdResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupIdResponseBody) SetRequestId(v string) *CreateGroupIdResponseBody {
	s.RequestId = &v
	return s
}

type CreateGroupIdResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupIdResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupIdResponse) SetHeaders(v map[string]*string) *CreateGroupIdResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupIdResponse) SetBody(v *CreateGroupIdResponseBody) *CreateGroupIdResponse {
	s.Body = v
	return s
}

type DeleteGroupIdRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupIdRequest) SetGroupId(v string) *DeleteGroupIdRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteGroupIdRequest) SetInstanceId(v string) *DeleteGroupIdRequest {
	s.InstanceId = &v
	return s
}

type DeleteGroupIdResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupIdResponseBody) SetRequestId(v string) *DeleteGroupIdResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGroupIdResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupIdResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupIdResponse) SetHeaders(v map[string]*string) *DeleteGroupIdResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupIdResponse) SetBody(v *DeleteGroupIdResponseBody) *DeleteGroupIdResponse {
	s.Body = v
	return s
}

type GetDeviceCredentialRequest struct {
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDeviceCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceCredentialRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceCredentialRequest) SetClientId(v string) *GetDeviceCredentialRequest {
	s.ClientId = &v
	return s
}

func (s *GetDeviceCredentialRequest) SetInstanceId(v string) *GetDeviceCredentialRequest {
	s.InstanceId = &v
	return s
}

type GetDeviceCredentialResponseBody struct {
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceCredential *GetDeviceCredentialResponseBodyDeviceCredential `json:"DeviceCredential,omitempty" xml:"DeviceCredential,omitempty" type:"Struct"`
}

func (s GetDeviceCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceCredentialResponseBody) SetRequestId(v string) *GetDeviceCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeviceCredentialResponseBody) SetDeviceCredential(v *GetDeviceCredentialResponseBodyDeviceCredential) *GetDeviceCredentialResponseBody {
	s.DeviceCredential = v
	return s
}

type GetDeviceCredentialResponseBodyDeviceCredential struct {
	UpdateTime            *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	DeviceAccessKeyId     *string `json:"DeviceAccessKeyId,omitempty" xml:"DeviceAccessKeyId,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DeviceAccessKeySecret *string `json:"DeviceAccessKeySecret,omitempty" xml:"DeviceAccessKeySecret,omitempty"`
	ClientId              *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s GetDeviceCredentialResponseBodyDeviceCredential) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceCredentialResponseBodyDeviceCredential) GoString() string {
	return s.String()
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetUpdateTime(v int64) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.UpdateTime = &v
	return s
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeyId(v string) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeyId = &v
	return s
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetCreateTime(v int64) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.CreateTime = &v
	return s
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetInstanceId(v string) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.InstanceId = &v
	return s
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeySecret(v string) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeySecret = &v
	return s
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetClientId(v string) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.ClientId = &v
	return s
}

type GetDeviceCredentialResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeviceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceCredentialResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceCredentialResponse) SetHeaders(v map[string]*string) *GetDeviceCredentialResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceCredentialResponse) SetBody(v *GetDeviceCredentialResponseBody) *GetDeviceCredentialResponse {
	s.Body = v
	return s
}

type ListGroupIdRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupIdRequest) GoString() string {
	return s.String()
}

func (s *ListGroupIdRequest) SetInstanceId(v string) *ListGroupIdRequest {
	s.InstanceId = &v
	return s
}

type ListGroupIdResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*ListGroupIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupIdResponseBody) SetRequestId(v string) *ListGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupIdResponseBody) SetData(v []*ListGroupIdResponseBodyData) *ListGroupIdResponseBody {
	s.Data = v
	return s
}

type ListGroupIdResponseBodyData struct {
	UpdateTime        *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IndependentNaming *bool   `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	GroupId           *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s ListGroupIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGroupIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGroupIdResponseBodyData) SetUpdateTime(v int64) *ListGroupIdResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListGroupIdResponseBodyData) SetInstanceId(v string) *ListGroupIdResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListGroupIdResponseBodyData) SetIndependentNaming(v bool) *ListGroupIdResponseBodyData {
	s.IndependentNaming = &v
	return s
}

func (s *ListGroupIdResponseBodyData) SetGroupId(v string) *ListGroupIdResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *ListGroupIdResponseBodyData) SetCreateTime(v int64) *ListGroupIdResponseBodyData {
	s.CreateTime = &v
	return s
}

type ListGroupIdResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupIdResponse) GoString() string {
	return s.String()
}

func (s *ListGroupIdResponse) SetHeaders(v map[string]*string) *ListGroupIdResponse {
	s.Headers = v
	return s
}

func (s *ListGroupIdResponse) SetBody(v *ListGroupIdResponseBody) *ListGroupIdResponse {
	s.Body = v
	return s
}

type QueryMqttTraceDeviceRequest struct {
	MqttRegionId *string `json:"MqttRegionId,omitempty" xml:"MqttRegionId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Reverse      *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	BeginTime    *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryMqttTraceDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceDeviceRequest) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceDeviceRequest) SetMqttRegionId(v string) *QueryMqttTraceDeviceRequest {
	s.MqttRegionId = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetInstanceId(v string) *QueryMqttTraceDeviceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetReverse(v bool) *QueryMqttTraceDeviceRequest {
	s.Reverse = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetClientId(v string) *QueryMqttTraceDeviceRequest {
	s.ClientId = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetBeginTime(v int64) *QueryMqttTraceDeviceRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetEndTime(v int64) *QueryMqttTraceDeviceRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetCurrentPage(v int32) *QueryMqttTraceDeviceRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetPageSize(v int32) *QueryMqttTraceDeviceRequest {
	s.PageSize = &v
	return s
}

type QueryMqttTraceDeviceResponseBody struct {
	CurrentPage    *int32                                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize       *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total          *int64                                            `json:"Total,omitempty" xml:"Total,omitempty"`
	DeviceInfoList []*QueryMqttTraceDeviceResponseBodyDeviceInfoList `json:"DeviceInfoList,omitempty" xml:"DeviceInfoList,omitempty" type:"Repeated"`
}

func (s QueryMqttTraceDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceDeviceResponseBody) SetCurrentPage(v int32) *QueryMqttTraceDeviceResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBody) SetRequestId(v string) *QueryMqttTraceDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBody) SetPageSize(v int32) *QueryMqttTraceDeviceResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBody) SetTotal(v int64) *QueryMqttTraceDeviceResponseBody {
	s.Total = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBody) SetDeviceInfoList(v []*QueryMqttTraceDeviceResponseBodyDeviceInfoList) *QueryMqttTraceDeviceResponseBody {
	s.DeviceInfoList = v
	return s
}

type QueryMqttTraceDeviceResponseBodyDeviceInfoList struct {
	ChannelId  *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty"`
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	Action     *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
}

func (s QueryMqttTraceDeviceResponseBodyDeviceInfoList) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceDeviceResponseBodyDeviceInfoList) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) SetChannelId(v string) *QueryMqttTraceDeviceResponseBodyDeviceInfoList {
	s.ChannelId = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) SetTime(v string) *QueryMqttTraceDeviceResponseBodyDeviceInfoList {
	s.Time = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) SetActionCode(v string) *QueryMqttTraceDeviceResponseBodyDeviceInfoList {
	s.ActionCode = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) SetAction(v string) *QueryMqttTraceDeviceResponseBodyDeviceInfoList {
	s.Action = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) SetActionInfo(v string) *QueryMqttTraceDeviceResponseBodyDeviceInfoList {
	s.ActionInfo = &v
	return s
}

type QueryMqttTraceDeviceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMqttTraceDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMqttTraceDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceDeviceResponse) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceDeviceResponse) SetHeaders(v map[string]*string) *QueryMqttTraceDeviceResponse {
	s.Headers = v
	return s
}

func (s *QueryMqttTraceDeviceResponse) SetBody(v *QueryMqttTraceDeviceResponseBody) *QueryMqttTraceDeviceResponse {
	s.Body = v
	return s
}

type QueryMqttTraceMessageOfClientRequest struct {
	MqttRegionId *string `json:"MqttRegionId,omitempty" xml:"MqttRegionId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	BeginTime    *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Reverse      *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
}

func (s QueryMqttTraceMessageOfClientRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessageOfClientRequest) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageOfClientRequest) SetMqttRegionId(v string) *QueryMqttTraceMessageOfClientRequest {
	s.MqttRegionId = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetInstanceId(v string) *QueryMqttTraceMessageOfClientRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetClientId(v string) *QueryMqttTraceMessageOfClientRequest {
	s.ClientId = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetBeginTime(v int64) *QueryMqttTraceMessageOfClientRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetEndTime(v int64) *QueryMqttTraceMessageOfClientRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetCurrentPage(v int32) *QueryMqttTraceMessageOfClientRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetPageSize(v int32) *QueryMqttTraceMessageOfClientRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetReverse(v bool) *QueryMqttTraceMessageOfClientRequest {
	s.Reverse = &v
	return s
}

type QueryMqttTraceMessageOfClientResponseBody struct {
	CurrentPage         *int32                                                          `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId           *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize            *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total               *int64                                                          `json:"Total,omitempty" xml:"Total,omitempty"`
	MessageOfClientList []*QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList `json:"MessageOfClientList,omitempty" xml:"MessageOfClientList,omitempty" type:"Repeated"`
}

func (s QueryMqttTraceMessageOfClientResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessageOfClientResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageOfClientResponseBody) SetCurrentPage(v int32) *QueryMqttTraceMessageOfClientResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBody) SetRequestId(v string) *QueryMqttTraceMessageOfClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBody) SetPageSize(v int32) *QueryMqttTraceMessageOfClientResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBody) SetTotal(v int64) *QueryMqttTraceMessageOfClientResponseBody {
	s.Total = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBody) SetMessageOfClientList(v []*QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) *QueryMqttTraceMessageOfClientResponseBody {
	s.MessageOfClientList = v
	return s
}

type QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList struct {
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Action     *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) SetTime(v string) *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList {
	s.Time = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) SetAction(v string) *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList {
	s.Action = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) SetActionCode(v string) *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList {
	s.ActionCode = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) SetActionInfo(v string) *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList {
	s.ActionInfo = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) SetMsgId(v string) *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList {
	s.MsgId = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) SetClientId(v string) *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList {
	s.ClientId = &v
	return s
}

type QueryMqttTraceMessageOfClientResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMqttTraceMessageOfClientResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMqttTraceMessageOfClientResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessageOfClientResponse) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageOfClientResponse) SetHeaders(v map[string]*string) *QueryMqttTraceMessageOfClientResponse {
	s.Headers = v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponse) SetBody(v *QueryMqttTraceMessageOfClientResponseBody) *QueryMqttTraceMessageOfClientResponse {
	s.Body = v
	return s
}

type QueryMqttTraceMessagePublishRequest struct {
	MqttRegionId *string `json:"MqttRegionId,omitempty" xml:"MqttRegionId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MsgId        *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	BeginTime    *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s QueryMqttTraceMessagePublishRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessagePublishRequest) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessagePublishRequest) SetMqttRegionId(v string) *QueryMqttTraceMessagePublishRequest {
	s.MqttRegionId = &v
	return s
}

func (s *QueryMqttTraceMessagePublishRequest) SetInstanceId(v string) *QueryMqttTraceMessagePublishRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqttTraceMessagePublishRequest) SetMsgId(v string) *QueryMqttTraceMessagePublishRequest {
	s.MsgId = &v
	return s
}

func (s *QueryMqttTraceMessagePublishRequest) SetBeginTime(v int64) *QueryMqttTraceMessagePublishRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMqttTraceMessagePublishRequest) SetEndTime(v int64) *QueryMqttTraceMessagePublishRequest {
	s.EndTime = &v
	return s
}

type QueryMqttTraceMessagePublishResponseBody struct {
	RequestId         *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MessageTraceLists []*QueryMqttTraceMessagePublishResponseBodyMessageTraceLists `json:"MessageTraceLists,omitempty" xml:"MessageTraceLists,omitempty" type:"Repeated"`
}

func (s QueryMqttTraceMessagePublishResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessagePublishResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessagePublishResponseBody) SetRequestId(v string) *QueryMqttTraceMessagePublishResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBody) SetMessageTraceLists(v []*QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) *QueryMqttTraceMessagePublishResponseBody {
	s.MessageTraceLists = v
	return s
}

type QueryMqttTraceMessagePublishResponseBodyMessageTraceLists struct {
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Action     *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) SetTime(v string) *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists {
	s.Time = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) SetAction(v string) *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists {
	s.Action = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) SetActionCode(v string) *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists {
	s.ActionCode = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) SetActionInfo(v string) *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists {
	s.ActionInfo = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) SetMsgId(v string) *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists {
	s.MsgId = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) SetClientId(v string) *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists {
	s.ClientId = &v
	return s
}

type QueryMqttTraceMessagePublishResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMqttTraceMessagePublishResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMqttTraceMessagePublishResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessagePublishResponse) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessagePublishResponse) SetHeaders(v map[string]*string) *QueryMqttTraceMessagePublishResponse {
	s.Headers = v
	return s
}

func (s *QueryMqttTraceMessagePublishResponse) SetBody(v *QueryMqttTraceMessagePublishResponseBody) *QueryMqttTraceMessagePublishResponse {
	s.Body = v
	return s
}

type QueryMqttTraceMessageSubscribeRequest struct {
	MqttRegionId *string `json:"MqttRegionId,omitempty" xml:"MqttRegionId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Reverse      *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	BeginTime    *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	MsgId        *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s QueryMqttTraceMessageSubscribeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessageSubscribeRequest) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetMqttRegionId(v string) *QueryMqttTraceMessageSubscribeRequest {
	s.MqttRegionId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetInstanceId(v string) *QueryMqttTraceMessageSubscribeRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetReverse(v bool) *QueryMqttTraceMessageSubscribeRequest {
	s.Reverse = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetClientId(v string) *QueryMqttTraceMessageSubscribeRequest {
	s.ClientId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetBeginTime(v int64) *QueryMqttTraceMessageSubscribeRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetEndTime(v int64) *QueryMqttTraceMessageSubscribeRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetCurrentPage(v int32) *QueryMqttTraceMessageSubscribeRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetPageSize(v int32) *QueryMqttTraceMessageSubscribeRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetMsgId(v string) *QueryMqttTraceMessageSubscribeRequest {
	s.MsgId = &v
	return s
}

type QueryMqttTraceMessageSubscribeResponseBody struct {
	CurrentPage       *int32                                                         `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RequestId         *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize          *int32                                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total             *int64                                                         `json:"Total,omitempty" xml:"Total,omitempty"`
	MessageTraceLists []*QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists `json:"MessageTraceLists,omitempty" xml:"MessageTraceLists,omitempty" type:"Repeated"`
}

func (s QueryMqttTraceMessageSubscribeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessageSubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) SetCurrentPage(v int32) *QueryMqttTraceMessageSubscribeResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) SetRequestId(v string) *QueryMqttTraceMessageSubscribeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) SetPageSize(v int32) *QueryMqttTraceMessageSubscribeResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) SetTotal(v int64) *QueryMqttTraceMessageSubscribeResponseBody {
	s.Total = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) SetMessageTraceLists(v []*QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) *QueryMqttTraceMessageSubscribeResponseBody {
	s.MessageTraceLists = v
	return s
}

type QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists struct {
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Action     *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) SetTime(v string) *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists {
	s.Time = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) SetAction(v string) *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists {
	s.Action = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) SetActionCode(v string) *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists {
	s.ActionCode = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) SetActionInfo(v string) *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists {
	s.ActionInfo = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) SetMsgId(v string) *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists {
	s.MsgId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) SetClientId(v string) *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists {
	s.ClientId = &v
	return s
}

type QueryMqttTraceMessageSubscribeResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMqttTraceMessageSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMqttTraceMessageSubscribeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessageSubscribeResponse) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageSubscribeResponse) SetHeaders(v map[string]*string) *QueryMqttTraceMessageSubscribeResponse {
	s.Headers = v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponse) SetBody(v *QueryMqttTraceMessageSubscribeResponseBody) *QueryMqttTraceMessageSubscribeResponse {
	s.Body = v
	return s
}

type QuerySessionByClientIdRequest struct {
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QuerySessionByClientIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySessionByClientIdRequest) GoString() string {
	return s.String()
}

func (s *QuerySessionByClientIdRequest) SetClientId(v string) *QuerySessionByClientIdRequest {
	s.ClientId = &v
	return s
}

func (s *QuerySessionByClientIdRequest) SetInstanceId(v string) *QuerySessionByClientIdRequest {
	s.InstanceId = &v
	return s
}

type QuerySessionByClientIdResponseBody struct {
	OnlineStatus *bool   `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuerySessionByClientIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySessionByClientIdResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySessionByClientIdResponseBody) SetOnlineStatus(v bool) *QuerySessionByClientIdResponseBody {
	s.OnlineStatus = &v
	return s
}

func (s *QuerySessionByClientIdResponseBody) SetRequestId(v string) *QuerySessionByClientIdResponseBody {
	s.RequestId = &v
	return s
}

type QuerySessionByClientIdResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySessionByClientIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySessionByClientIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySessionByClientIdResponse) GoString() string {
	return s.String()
}

func (s *QuerySessionByClientIdResponse) SetHeaders(v map[string]*string) *QuerySessionByClientIdResponse {
	s.Headers = v
	return s
}

func (s *QuerySessionByClientIdResponse) SetBody(v *QuerySessionByClientIdResponseBody) *QuerySessionByClientIdResponse {
	s.Body = v
	return s
}

type QueryTokenRequest struct {
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTokenRequest) GoString() string {
	return s.String()
}

func (s *QueryTokenRequest) SetToken(v string) *QueryTokenRequest {
	s.Token = &v
	return s
}

func (s *QueryTokenRequest) SetInstanceId(v string) *QueryTokenRequest {
	s.InstanceId = &v
	return s
}

type QueryTokenResponseBody struct {
	TokenStatus *bool   `json:"TokenStatus,omitempty" xml:"TokenStatus,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTokenResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTokenResponseBody) SetTokenStatus(v bool) *QueryTokenResponseBody {
	s.TokenStatus = &v
	return s
}

func (s *QueryTokenResponseBody) SetRequestId(v string) *QueryTokenResponseBody {
	s.RequestId = &v
	return s
}

type QueryTokenResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTokenResponse) GoString() string {
	return s.String()
}

func (s *QueryTokenResponse) SetHeaders(v map[string]*string) *QueryTokenResponse {
	s.Headers = v
	return s
}

func (s *QueryTokenResponse) SetBody(v *QueryTokenResponseBody) *QueryTokenResponse {
	s.Body = v
	return s
}

type RefreshDeviceCredentialRequest struct {
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RefreshDeviceCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshDeviceCredentialRequest) GoString() string {
	return s.String()
}

func (s *RefreshDeviceCredentialRequest) SetClientId(v string) *RefreshDeviceCredentialRequest {
	s.ClientId = &v
	return s
}

func (s *RefreshDeviceCredentialRequest) SetInstanceId(v string) *RefreshDeviceCredentialRequest {
	s.InstanceId = &v
	return s
}

type RefreshDeviceCredentialResponseBody struct {
	RequestId        *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceCredential *RefreshDeviceCredentialResponseBodyDeviceCredential `json:"DeviceCredential,omitempty" xml:"DeviceCredential,omitempty" type:"Struct"`
}

func (s RefreshDeviceCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshDeviceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshDeviceCredentialResponseBody) SetRequestId(v string) *RefreshDeviceCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBody) SetDeviceCredential(v *RefreshDeviceCredentialResponseBodyDeviceCredential) *RefreshDeviceCredentialResponseBody {
	s.DeviceCredential = v
	return s
}

type RefreshDeviceCredentialResponseBodyDeviceCredential struct {
	UpdateTime            *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	DeviceAccessKeyId     *string `json:"DeviceAccessKeyId,omitempty" xml:"DeviceAccessKeyId,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DeviceAccessKeySecret *string `json:"DeviceAccessKeySecret,omitempty" xml:"DeviceAccessKeySecret,omitempty"`
	ClientId              *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s RefreshDeviceCredentialResponseBodyDeviceCredential) String() string {
	return tea.Prettify(s)
}

func (s RefreshDeviceCredentialResponseBodyDeviceCredential) GoString() string {
	return s.String()
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetUpdateTime(v int64) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.UpdateTime = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeyId(v string) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeyId = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetCreateTime(v int64) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.CreateTime = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetInstanceId(v string) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.InstanceId = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeySecret(v string) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeySecret = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetClientId(v string) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.ClientId = &v
	return s
}

type RefreshDeviceCredentialResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefreshDeviceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshDeviceCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshDeviceCredentialResponse) GoString() string {
	return s.String()
}

func (s *RefreshDeviceCredentialResponse) SetHeaders(v map[string]*string) *RefreshDeviceCredentialResponse {
	s.Headers = v
	return s
}

func (s *RefreshDeviceCredentialResponse) SetBody(v *RefreshDeviceCredentialResponseBody) *RefreshDeviceCredentialResponse {
	s.Body = v
	return s
}

type RegisterDeviceCredentialRequest struct {
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RegisterDeviceCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceCredentialRequest) GoString() string {
	return s.String()
}

func (s *RegisterDeviceCredentialRequest) SetClientId(v string) *RegisterDeviceCredentialRequest {
	s.ClientId = &v
	return s
}

func (s *RegisterDeviceCredentialRequest) SetInstanceId(v string) *RegisterDeviceCredentialRequest {
	s.InstanceId = &v
	return s
}

type RegisterDeviceCredentialResponseBody struct {
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceCredential *RegisterDeviceCredentialResponseBodyDeviceCredential `json:"DeviceCredential,omitempty" xml:"DeviceCredential,omitempty" type:"Struct"`
}

func (s RegisterDeviceCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDeviceCredentialResponseBody) SetRequestId(v string) *RegisterDeviceCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBody) SetDeviceCredential(v *RegisterDeviceCredentialResponseBodyDeviceCredential) *RegisterDeviceCredentialResponseBody {
	s.DeviceCredential = v
	return s
}

type RegisterDeviceCredentialResponseBodyDeviceCredential struct {
	UpdateTime            *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	DeviceAccessKeyId     *string `json:"DeviceAccessKeyId,omitempty" xml:"DeviceAccessKeyId,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DeviceAccessKeySecret *string `json:"DeviceAccessKeySecret,omitempty" xml:"DeviceAccessKeySecret,omitempty"`
	ClientId              *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s RegisterDeviceCredentialResponseBodyDeviceCredential) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceCredentialResponseBodyDeviceCredential) GoString() string {
	return s.String()
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetUpdateTime(v int64) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.UpdateTime = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeyId(v string) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeyId = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetCreateTime(v int64) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.CreateTime = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetInstanceId(v string) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.InstanceId = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeySecret(v string) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeySecret = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetClientId(v string) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.ClientId = &v
	return s
}

type RegisterDeviceCredentialResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterDeviceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterDeviceCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceCredentialResponse) GoString() string {
	return s.String()
}

func (s *RegisterDeviceCredentialResponse) SetHeaders(v map[string]*string) *RegisterDeviceCredentialResponse {
	s.Headers = v
	return s
}

func (s *RegisterDeviceCredentialResponse) SetBody(v *RegisterDeviceCredentialResponseBody) *RegisterDeviceCredentialResponse {
	s.Body = v
	return s
}

type RevokeTokenRequest struct {
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RevokeTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeTokenRequest) GoString() string {
	return s.String()
}

func (s *RevokeTokenRequest) SetToken(v string) *RevokeTokenRequest {
	s.Token = &v
	return s
}

func (s *RevokeTokenRequest) SetInstanceId(v string) *RevokeTokenRequest {
	s.InstanceId = &v
	return s
}

type RevokeTokenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeTokenResponseBody) SetRequestId(v string) *RevokeTokenResponseBody {
	s.RequestId = &v
	return s
}

type RevokeTokenResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RevokeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokeTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeTokenResponse) GoString() string {
	return s.String()
}

func (s *RevokeTokenResponse) SetHeaders(v map[string]*string) *RevokeTokenResponse {
	s.Headers = v
	return s
}

func (s *RevokeTokenResponse) SetBody(v *RevokeTokenResponseBody) *RevokeTokenResponse {
	s.Body = v
	return s
}

type SendMessageRequest struct {
	MqttTopic  *string `json:"MqttTopic,omitempty" xml:"MqttTopic,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Payload    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetMqttTopic(v string) *SendMessageRequest {
	s.MqttTopic = &v
	return s
}

func (s *SendMessageRequest) SetInstanceId(v string) *SendMessageRequest {
	s.InstanceId = &v
	return s
}

func (s *SendMessageRequest) SetPayload(v string) *SendMessageRequest {
	s.Payload = &v
	return s
}

type SendMessageResponseBody struct {
	MsgId     *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) SetMsgId(v string) *SendMessageResponseBody {
	s.MsgId = &v
	return s
}

func (s *SendMessageResponseBody) SetRequestId(v string) *SendMessageResponseBody {
	s.RequestId = &v
	return s
}

type SendMessageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponse) GoString() string {
	return s.String()
}

func (s *SendMessageResponse) SetHeaders(v map[string]*string) *SendMessageResponse {
	s.Headers = v
	return s
}

func (s *SendMessageResponse) SetBody(v *SendMessageResponseBody) *SendMessageResponse {
	s.Body = v
	return s
}

type UnRegisterDeviceCredentialRequest struct {
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UnRegisterDeviceCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s UnRegisterDeviceCredentialRequest) GoString() string {
	return s.String()
}

func (s *UnRegisterDeviceCredentialRequest) SetClientId(v string) *UnRegisterDeviceCredentialRequest {
	s.ClientId = &v
	return s
}

func (s *UnRegisterDeviceCredentialRequest) SetInstanceId(v string) *UnRegisterDeviceCredentialRequest {
	s.InstanceId = &v
	return s
}

type UnRegisterDeviceCredentialResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnRegisterDeviceCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnRegisterDeviceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *UnRegisterDeviceCredentialResponseBody) SetRequestId(v string) *UnRegisterDeviceCredentialResponseBody {
	s.RequestId = &v
	return s
}

type UnRegisterDeviceCredentialResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnRegisterDeviceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnRegisterDeviceCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s UnRegisterDeviceCredentialResponse) GoString() string {
	return s.String()
}

func (s *UnRegisterDeviceCredentialResponse) SetHeaders(v map[string]*string) *UnRegisterDeviceCredentialResponse {
	s.Headers = v
	return s
}

func (s *UnRegisterDeviceCredentialResponse) SetBody(v *UnRegisterDeviceCredentialResponseBody) *UnRegisterDeviceCredentialResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("onsmqtt"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ApplyTokenWithOptions(request *ApplyTokenRequest, runtime *util.RuntimeOptions) (_result *ApplyTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApplyTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApplyToken"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyToken(request *ApplyTokenRequest) (_result *ApplyTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyTokenResponse{}
	_body, _err := client.ApplyTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchQuerySessionByClientIdsWithOptions(request *BatchQuerySessionByClientIdsRequest, runtime *util.RuntimeOptions) (_result *BatchQuerySessionByClientIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchQuerySessionByClientIdsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchQuerySessionByClientIds"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchQuerySessionByClientIds(request *BatchQuerySessionByClientIdsRequest) (_result *BatchQuerySessionByClientIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchQuerySessionByClientIdsResponse{}
	_body, _err := client.BatchQuerySessionByClientIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupIdWithOptions(request *CreateGroupIdRequest, runtime *util.RuntimeOptions) (_result *CreateGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateGroupIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateGroupId"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroupId(request *CreateGroupIdRequest) (_result *CreateGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupIdResponse{}
	_body, _err := client.CreateGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGroupIdWithOptions(request *DeleteGroupIdRequest, runtime *util.RuntimeOptions) (_result *DeleteGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteGroupIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteGroupId"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGroupId(request *DeleteGroupIdRequest) (_result *DeleteGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGroupIdResponse{}
	_body, _err := client.DeleteGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceCredentialWithOptions(request *GetDeviceCredentialRequest, runtime *util.RuntimeOptions) (_result *GetDeviceCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDeviceCredentialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDeviceCredential"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceCredential(request *GetDeviceCredentialRequest) (_result *GetDeviceCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceCredentialResponse{}
	_body, _err := client.GetDeviceCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGroupIdWithOptions(request *ListGroupIdRequest, runtime *util.RuntimeOptions) (_result *ListGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListGroupIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListGroupId"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGroupId(request *ListGroupIdRequest) (_result *ListGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupIdResponse{}
	_body, _err := client.ListGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMqttTraceDeviceWithOptions(request *QueryMqttTraceDeviceRequest, runtime *util.RuntimeOptions) (_result *QueryMqttTraceDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMqttTraceDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMqttTraceDevice"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMqttTraceDevice(request *QueryMqttTraceDeviceRequest) (_result *QueryMqttTraceDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMqttTraceDeviceResponse{}
	_body, _err := client.QueryMqttTraceDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMqttTraceMessageOfClientWithOptions(request *QueryMqttTraceMessageOfClientRequest, runtime *util.RuntimeOptions) (_result *QueryMqttTraceMessageOfClientResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMqttTraceMessageOfClientResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMqttTraceMessageOfClient"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMqttTraceMessageOfClient(request *QueryMqttTraceMessageOfClientRequest) (_result *QueryMqttTraceMessageOfClientResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMqttTraceMessageOfClientResponse{}
	_body, _err := client.QueryMqttTraceMessageOfClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMqttTraceMessagePublishWithOptions(request *QueryMqttTraceMessagePublishRequest, runtime *util.RuntimeOptions) (_result *QueryMqttTraceMessagePublishResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMqttTraceMessagePublishResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMqttTraceMessagePublish"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMqttTraceMessagePublish(request *QueryMqttTraceMessagePublishRequest) (_result *QueryMqttTraceMessagePublishResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMqttTraceMessagePublishResponse{}
	_body, _err := client.QueryMqttTraceMessagePublishWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMqttTraceMessageSubscribeWithOptions(request *QueryMqttTraceMessageSubscribeRequest, runtime *util.RuntimeOptions) (_result *QueryMqttTraceMessageSubscribeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMqttTraceMessageSubscribeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMqttTraceMessageSubscribe"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMqttTraceMessageSubscribe(request *QueryMqttTraceMessageSubscribeRequest) (_result *QueryMqttTraceMessageSubscribeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMqttTraceMessageSubscribeResponse{}
	_body, _err := client.QueryMqttTraceMessageSubscribeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySessionByClientIdWithOptions(request *QuerySessionByClientIdRequest, runtime *util.RuntimeOptions) (_result *QuerySessionByClientIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySessionByClientIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySessionByClientId"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySessionByClientId(request *QuerySessionByClientIdRequest) (_result *QuerySessionByClientIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySessionByClientIdResponse{}
	_body, _err := client.QuerySessionByClientIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTokenWithOptions(request *QueryTokenRequest, runtime *util.RuntimeOptions) (_result *QueryTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryToken"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryToken(request *QueryTokenRequest) (_result *QueryTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTokenResponse{}
	_body, _err := client.QueryTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshDeviceCredentialWithOptions(request *RefreshDeviceCredentialRequest, runtime *util.RuntimeOptions) (_result *RefreshDeviceCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefreshDeviceCredentialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefreshDeviceCredential"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshDeviceCredential(request *RefreshDeviceCredentialRequest) (_result *RefreshDeviceCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshDeviceCredentialResponse{}
	_body, _err := client.RefreshDeviceCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterDeviceCredentialWithOptions(request *RegisterDeviceCredentialRequest, runtime *util.RuntimeOptions) (_result *RegisterDeviceCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegisterDeviceCredentialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterDeviceCredential"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterDeviceCredential(request *RegisterDeviceCredentialRequest) (_result *RegisterDeviceCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterDeviceCredentialResponse{}
	_body, _err := client.RegisterDeviceCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeTokenWithOptions(request *RevokeTokenRequest, runtime *util.RuntimeOptions) (_result *RevokeTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RevokeTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RevokeToken"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeToken(request *RevokeTokenRequest) (_result *RevokeTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeTokenResponse{}
	_body, _err := client.RevokeTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMessageWithOptions(request *SendMessageRequest, runtime *util.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendMessage"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMessage(request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnRegisterDeviceCredentialWithOptions(request *UnRegisterDeviceCredentialRequest, runtime *util.RuntimeOptions) (_result *UnRegisterDeviceCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnRegisterDeviceCredentialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnRegisterDeviceCredential"), tea.String("2020-04-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnRegisterDeviceCredential(request *UnRegisterDeviceCredentialRequest) (_result *UnRegisterDeviceCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnRegisterDeviceCredentialResponse{}
	_body, _err := client.UnRegisterDeviceCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
