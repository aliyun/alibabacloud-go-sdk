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

type ActiveCaCertificateRequest struct {
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	Sn             *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s ActiveCaCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s ActiveCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *ActiveCaCertificateRequest) SetMqttInstanceId(v string) *ActiveCaCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *ActiveCaCertificateRequest) SetSn(v string) *ActiveCaCertificateRequest {
	s.Sn = &v
	return s
}

type ActiveCaCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sn        *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s ActiveCaCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActiveCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveCaCertificateResponseBody) SetRequestId(v string) *ActiveCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActiveCaCertificateResponseBody) SetSn(v string) *ActiveCaCertificateResponseBody {
	s.Sn = &v
	return s
}

type ActiveCaCertificateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ActiveCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActiveCaCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s ActiveCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *ActiveCaCertificateResponse) SetHeaders(v map[string]*string) *ActiveCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *ActiveCaCertificateResponse) SetStatusCode(v int32) *ActiveCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ActiveCaCertificateResponse) SetBody(v *ActiveCaCertificateResponseBody) *ActiveCaCertificateResponse {
	s.Body = v
	return s
}

type ActiveDeviceCertificateRequest struct {
	CaSn           *string `json:"CaSn,omitempty" xml:"CaSn,omitempty"`
	DeviceSn       *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
}

func (s ActiveDeviceCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s ActiveDeviceCertificateRequest) GoString() string {
	return s.String()
}

func (s *ActiveDeviceCertificateRequest) SetCaSn(v string) *ActiveDeviceCertificateRequest {
	s.CaSn = &v
	return s
}

func (s *ActiveDeviceCertificateRequest) SetDeviceSn(v string) *ActiveDeviceCertificateRequest {
	s.DeviceSn = &v
	return s
}

func (s *ActiveDeviceCertificateRequest) SetMqttInstanceId(v string) *ActiveDeviceCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

type ActiveDeviceCertificateResponseBody struct {
	DeviceSn  *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActiveDeviceCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActiveDeviceCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveDeviceCertificateResponseBody) SetDeviceSn(v string) *ActiveDeviceCertificateResponseBody {
	s.DeviceSn = &v
	return s
}

func (s *ActiveDeviceCertificateResponseBody) SetRequestId(v string) *ActiveDeviceCertificateResponseBody {
	s.RequestId = &v
	return s
}

type ActiveDeviceCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ActiveDeviceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActiveDeviceCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s ActiveDeviceCertificateResponse) GoString() string {
	return s.String()
}

func (s *ActiveDeviceCertificateResponse) SetHeaders(v map[string]*string) *ActiveDeviceCertificateResponse {
	s.Headers = v
	return s
}

func (s *ActiveDeviceCertificateResponse) SetStatusCode(v int32) *ActiveDeviceCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ActiveDeviceCertificateResponse) SetBody(v *ActiveDeviceCertificateResponseBody) *ActiveDeviceCertificateResponse {
	s.Body = v
	return s
}

type ApplyTokenRequest struct {
	Actions    *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Resources  *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s ApplyTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyTokenRequest) GoString() string {
	return s.String()
}

func (s *ApplyTokenRequest) SetActions(v string) *ApplyTokenRequest {
	s.Actions = &v
	return s
}

func (s *ApplyTokenRequest) SetExpireTime(v int64) *ApplyTokenRequest {
	s.ExpireTime = &v
	return s
}

func (s *ApplyTokenRequest) SetInstanceId(v string) *ApplyTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *ApplyTokenRequest) SetResources(v string) *ApplyTokenRequest {
	s.Resources = &v
	return s
}

type ApplyTokenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s ApplyTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyTokenResponseBody) SetRequestId(v string) *ApplyTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyTokenResponseBody) SetToken(v string) *ApplyTokenResponseBody {
	s.Token = &v
	return s
}

type ApplyTokenResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ApplyTokenResponse) SetStatusCode(v int32) *ApplyTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyTokenResponse) SetBody(v *ApplyTokenResponseBody) *ApplyTokenResponse {
	s.Body = v
	return s
}

type BatchQuerySessionByClientIdsRequest struct {
	ClientIdList []*string `json:"ClientIdList,omitempty" xml:"ClientIdList,omitempty" type:"Repeated"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s BatchQuerySessionByClientIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQuerySessionByClientIdsRequest) GoString() string {
	return s.String()
}

func (s *BatchQuerySessionByClientIdsRequest) SetClientIdList(v []*string) *BatchQuerySessionByClientIdsRequest {
	s.ClientIdList = v
	return s
}

func (s *BatchQuerySessionByClientIdsRequest) SetInstanceId(v string) *BatchQuerySessionByClientIdsRequest {
	s.InstanceId = &v
	return s
}

type BatchQuerySessionByClientIdsResponseBody struct {
	OnlineStatusList []*BatchQuerySessionByClientIdsResponseBodyOnlineStatusList `json:"OnlineStatusList,omitempty" xml:"OnlineStatusList,omitempty" type:"Repeated"`
	RequestId        *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchQuerySessionByClientIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQuerySessionByClientIdsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQuerySessionByClientIdsResponseBody) SetOnlineStatusList(v []*BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) *BatchQuerySessionByClientIdsResponseBody {
	s.OnlineStatusList = v
	return s
}

func (s *BatchQuerySessionByClientIdsResponseBody) SetRequestId(v string) *BatchQuerySessionByClientIdsResponseBody {
	s.RequestId = &v
	return s
}

type BatchQuerySessionByClientIdsResponseBodyOnlineStatusList struct {
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	OnlineStatus *bool   `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
}

func (s BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) String() string {
	return tea.Prettify(s)
}

func (s BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) GoString() string {
	return s.String()
}

func (s *BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) SetClientId(v string) *BatchQuerySessionByClientIdsResponseBodyOnlineStatusList {
	s.ClientId = &v
	return s
}

func (s *BatchQuerySessionByClientIdsResponseBodyOnlineStatusList) SetOnlineStatus(v bool) *BatchQuerySessionByClientIdsResponseBodyOnlineStatusList {
	s.OnlineStatus = &v
	return s
}

type BatchQuerySessionByClientIdsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchQuerySessionByClientIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BatchQuerySessionByClientIdsResponse) SetStatusCode(v int32) *BatchQuerySessionByClientIdsResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateGroupIdResponse) SetStatusCode(v int32) *CreateGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupIdResponse) SetBody(v *CreateGroupIdResponseBody) *CreateGroupIdResponse {
	s.Body = v
	return s
}

type DeleteCaCertificateRequest struct {
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	Sn             *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s DeleteCaCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCaCertificateRequest) SetMqttInstanceId(v string) *DeleteCaCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *DeleteCaCertificateRequest) SetSn(v string) *DeleteCaCertificateRequest {
	s.Sn = &v
	return s
}

type DeleteCaCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sn        *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s DeleteCaCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCaCertificateResponseBody) SetRequestId(v string) *DeleteCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCaCertificateResponseBody) SetSn(v string) *DeleteCaCertificateResponseBody {
	s.Sn = &v
	return s
}

type DeleteCaCertificateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCaCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteCaCertificateResponse) SetHeaders(v map[string]*string) *DeleteCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteCaCertificateResponse) SetStatusCode(v int32) *DeleteCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCaCertificateResponse) SetBody(v *DeleteCaCertificateResponseBody) *DeleteCaCertificateResponse {
	s.Body = v
	return s
}

type DeleteDeviceCertificateRequest struct {
	CaSn           *string `json:"CaSn,omitempty" xml:"CaSn,omitempty"`
	DeviceSn       *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
}

func (s DeleteDeviceCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeviceCertificateRequest) SetCaSn(v string) *DeleteDeviceCertificateRequest {
	s.CaSn = &v
	return s
}

func (s *DeleteDeviceCertificateRequest) SetDeviceSn(v string) *DeleteDeviceCertificateRequest {
	s.DeviceSn = &v
	return s
}

func (s *DeleteDeviceCertificateRequest) SetMqttInstanceId(v string) *DeleteDeviceCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

type DeleteDeviceCertificateResponseBody struct {
	DeviceSn  *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDeviceCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeviceCertificateResponseBody) SetDeviceSn(v string) *DeleteDeviceCertificateResponseBody {
	s.DeviceSn = &v
	return s
}

func (s *DeleteDeviceCertificateResponseBody) SetRequestId(v string) *DeleteDeviceCertificateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDeviceCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDeviceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDeviceCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDeviceCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeviceCertificateResponse) SetHeaders(v map[string]*string) *DeleteDeviceCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeviceCertificateResponse) SetStatusCode(v int32) *DeleteDeviceCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeviceCertificateResponse) SetBody(v *DeleteDeviceCertificateResponseBody) *DeleteDeviceCertificateResponse {
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteGroupIdResponse) SetStatusCode(v int32) *DeleteGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupIdResponse) SetBody(v *DeleteGroupIdResponseBody) *DeleteGroupIdResponse {
	s.Body = v
	return s
}

type GetCaCertificateRequest struct {
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	Sn             *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s GetCaCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *GetCaCertificateRequest) SetMqttInstanceId(v string) *GetCaCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *GetCaCertificateRequest) SetSn(v string) *GetCaCertificateRequest {
	s.Sn = &v
	return s
}

type GetCaCertificateResponseBody struct {
	Data      *GetCaCertificateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCaCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetCaCertificateResponseBody) SetData(v *GetCaCertificateResponseBodyData) *GetCaCertificateResponseBody {
	s.Data = v
	return s
}

func (s *GetCaCertificateResponseBody) SetRequestId(v string) *GetCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

type GetCaCertificateResponseBodyData struct {
	CaContent           *string `json:"CaContent,omitempty" xml:"CaContent,omitempty"`
	CaName              *string `json:"CaName,omitempty" xml:"CaName,omitempty"`
	RegistrationCode    *string `json:"RegistrationCode,omitempty" xml:"RegistrationCode,omitempty"`
	Sn                  *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ValidBegin          *string `json:"ValidBegin,omitempty" xml:"ValidBegin,omitempty"`
	ValidEnd            *string `json:"ValidEnd,omitempty" xml:"ValidEnd,omitempty"`
	VerificationContent *string `json:"VerificationContent,omitempty" xml:"VerificationContent,omitempty"`
}

func (s GetCaCertificateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCaCertificateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCaCertificateResponseBodyData) SetCaContent(v string) *GetCaCertificateResponseBodyData {
	s.CaContent = &v
	return s
}

func (s *GetCaCertificateResponseBodyData) SetCaName(v string) *GetCaCertificateResponseBodyData {
	s.CaName = &v
	return s
}

func (s *GetCaCertificateResponseBodyData) SetRegistrationCode(v string) *GetCaCertificateResponseBodyData {
	s.RegistrationCode = &v
	return s
}

func (s *GetCaCertificateResponseBodyData) SetSn(v string) *GetCaCertificateResponseBodyData {
	s.Sn = &v
	return s
}

func (s *GetCaCertificateResponseBodyData) SetStatus(v string) *GetCaCertificateResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCaCertificateResponseBodyData) SetValidBegin(v string) *GetCaCertificateResponseBodyData {
	s.ValidBegin = &v
	return s
}

func (s *GetCaCertificateResponseBodyData) SetValidEnd(v string) *GetCaCertificateResponseBodyData {
	s.ValidEnd = &v
	return s
}

func (s *GetCaCertificateResponseBodyData) SetVerificationContent(v string) *GetCaCertificateResponseBodyData {
	s.VerificationContent = &v
	return s
}

type GetCaCertificateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCaCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *GetCaCertificateResponse) SetHeaders(v map[string]*string) *GetCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *GetCaCertificateResponse) SetStatusCode(v int32) *GetCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCaCertificateResponse) SetBody(v *GetCaCertificateResponseBody) *GetCaCertificateResponse {
	s.Body = v
	return s
}

type GetDeviceCertificateRequest struct {
	CaSn           *string `json:"CaSn,omitempty" xml:"CaSn,omitempty"`
	DeviceSn       *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
}

func (s GetDeviceCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceCertificateRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceCertificateRequest) SetCaSn(v string) *GetDeviceCertificateRequest {
	s.CaSn = &v
	return s
}

func (s *GetDeviceCertificateRequest) SetDeviceSn(v string) *GetDeviceCertificateRequest {
	s.DeviceSn = &v
	return s
}

func (s *GetDeviceCertificateRequest) SetMqttInstanceId(v string) *GetDeviceCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

type GetDeviceCertificateResponseBody struct {
	Data      *GetDeviceCertificateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceCertificateResponseBody) SetData(v *GetDeviceCertificateResponseBodyData) *GetDeviceCertificateResponseBody {
	s.Data = v
	return s
}

func (s *GetDeviceCertificateResponseBody) SetRequestId(v string) *GetDeviceCertificateResponseBody {
	s.RequestId = &v
	return s
}

type GetDeviceCertificateResponseBodyData struct {
	CaSn          *string `json:"CaSn,omitempty" xml:"CaSn,omitempty"`
	DeviceContent *string `json:"DeviceContent,omitempty" xml:"DeviceContent,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceSn      *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ValidBegin    *string `json:"ValidBegin,omitempty" xml:"ValidBegin,omitempty"`
	ValidEnd      *string `json:"ValidEnd,omitempty" xml:"ValidEnd,omitempty"`
}

func (s GetDeviceCertificateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceCertificateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeviceCertificateResponseBodyData) SetCaSn(v string) *GetDeviceCertificateResponseBodyData {
	s.CaSn = &v
	return s
}

func (s *GetDeviceCertificateResponseBodyData) SetDeviceContent(v string) *GetDeviceCertificateResponseBodyData {
	s.DeviceContent = &v
	return s
}

func (s *GetDeviceCertificateResponseBodyData) SetDeviceName(v string) *GetDeviceCertificateResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *GetDeviceCertificateResponseBodyData) SetDeviceSn(v string) *GetDeviceCertificateResponseBodyData {
	s.DeviceSn = &v
	return s
}

func (s *GetDeviceCertificateResponseBodyData) SetStatus(v string) *GetDeviceCertificateResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDeviceCertificateResponseBodyData) SetValidBegin(v string) *GetDeviceCertificateResponseBodyData {
	s.ValidBegin = &v
	return s
}

func (s *GetDeviceCertificateResponseBodyData) SetValidEnd(v string) *GetDeviceCertificateResponseBodyData {
	s.ValidEnd = &v
	return s
}

type GetDeviceCertificateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeviceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeviceCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceCertificateResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceCertificateResponse) SetHeaders(v map[string]*string) *GetDeviceCertificateResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceCertificateResponse) SetStatusCode(v int32) *GetDeviceCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceCertificateResponse) SetBody(v *GetDeviceCertificateResponseBody) *GetDeviceCertificateResponse {
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
	DeviceCredential *GetDeviceCredentialResponseBodyDeviceCredential `json:"DeviceCredential,omitempty" xml:"DeviceCredential,omitempty" type:"Struct"`
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeviceCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeviceCredentialResponseBody) SetDeviceCredential(v *GetDeviceCredentialResponseBodyDeviceCredential) *GetDeviceCredentialResponseBody {
	s.DeviceCredential = v
	return s
}

func (s *GetDeviceCredentialResponseBody) SetRequestId(v string) *GetDeviceCredentialResponseBody {
	s.RequestId = &v
	return s
}

type GetDeviceCredentialResponseBodyDeviceCredential struct {
	ClientId              *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeviceAccessKeyId     *string `json:"DeviceAccessKeyId,omitempty" xml:"DeviceAccessKeyId,omitempty"`
	DeviceAccessKeySecret *string `json:"DeviceAccessKeySecret,omitempty" xml:"DeviceAccessKeySecret,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UpdateTime            *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetDeviceCredentialResponseBodyDeviceCredential) String() string {
	return tea.Prettify(s)
}

func (s GetDeviceCredentialResponseBodyDeviceCredential) GoString() string {
	return s.String()
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetClientId(v string) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.ClientId = &v
	return s
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetCreateTime(v int64) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.CreateTime = &v
	return s
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeyId(v string) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeyId = &v
	return s
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeySecret(v string) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeySecret = &v
	return s
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetInstanceId(v string) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.InstanceId = &v
	return s
}

func (s *GetDeviceCredentialResponseBodyDeviceCredential) SetUpdateTime(v int64) *GetDeviceCredentialResponseBodyDeviceCredential {
	s.UpdateTime = &v
	return s
}

type GetDeviceCredentialResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeviceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDeviceCredentialResponse) SetStatusCode(v int32) *GetDeviceCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceCredentialResponse) SetBody(v *GetDeviceCredentialResponseBody) *GetDeviceCredentialResponse {
	s.Body = v
	return s
}

type GetRegisterCodeRequest struct {
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
}

func (s GetRegisterCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRegisterCodeRequest) GoString() string {
	return s.String()
}

func (s *GetRegisterCodeRequest) SetMqttInstanceId(v string) *GetRegisterCodeRequest {
	s.MqttInstanceId = &v
	return s
}

type GetRegisterCodeResponseBody struct {
	RegisterCode *string `json:"RegisterCode,omitempty" xml:"RegisterCode,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRegisterCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRegisterCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegisterCodeResponseBody) SetRegisterCode(v string) *GetRegisterCodeResponseBody {
	s.RegisterCode = &v
	return s
}

func (s *GetRegisterCodeResponseBody) SetRequestId(v string) *GetRegisterCodeResponseBody {
	s.RequestId = &v
	return s
}

type GetRegisterCodeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRegisterCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRegisterCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRegisterCodeResponse) GoString() string {
	return s.String()
}

func (s *GetRegisterCodeResponse) SetHeaders(v map[string]*string) *GetRegisterCodeResponse {
	s.Headers = v
	return s
}

func (s *GetRegisterCodeResponse) SetStatusCode(v int32) *GetRegisterCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegisterCodeResponse) SetBody(v *GetRegisterCodeResponseBody) *GetRegisterCodeResponse {
	s.Body = v
	return s
}

type InactivateCaCertificateRequest struct {
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	Sn             *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s InactivateCaCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s InactivateCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *InactivateCaCertificateRequest) SetMqttInstanceId(v string) *InactivateCaCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *InactivateCaCertificateRequest) SetSn(v string) *InactivateCaCertificateRequest {
	s.Sn = &v
	return s
}

type InactivateCaCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sn        *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s InactivateCaCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InactivateCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *InactivateCaCertificateResponseBody) SetRequestId(v string) *InactivateCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *InactivateCaCertificateResponseBody) SetSn(v string) *InactivateCaCertificateResponseBody {
	s.Sn = &v
	return s
}

type InactivateCaCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InactivateCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InactivateCaCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s InactivateCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *InactivateCaCertificateResponse) SetHeaders(v map[string]*string) *InactivateCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *InactivateCaCertificateResponse) SetStatusCode(v int32) *InactivateCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *InactivateCaCertificateResponse) SetBody(v *InactivateCaCertificateResponseBody) *InactivateCaCertificateResponse {
	s.Body = v
	return s
}

type InactivateDeviceCertificateRequest struct {
	CaSn           *string `json:"CaSn,omitempty" xml:"CaSn,omitempty"`
	DeviceSn       *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
}

func (s InactivateDeviceCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s InactivateDeviceCertificateRequest) GoString() string {
	return s.String()
}

func (s *InactivateDeviceCertificateRequest) SetCaSn(v string) *InactivateDeviceCertificateRequest {
	s.CaSn = &v
	return s
}

func (s *InactivateDeviceCertificateRequest) SetDeviceSn(v string) *InactivateDeviceCertificateRequest {
	s.DeviceSn = &v
	return s
}

func (s *InactivateDeviceCertificateRequest) SetMqttInstanceId(v string) *InactivateDeviceCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

type InactivateDeviceCertificateResponseBody struct {
	DeviceSn  *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InactivateDeviceCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InactivateDeviceCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *InactivateDeviceCertificateResponseBody) SetDeviceSn(v string) *InactivateDeviceCertificateResponseBody {
	s.DeviceSn = &v
	return s
}

func (s *InactivateDeviceCertificateResponseBody) SetRequestId(v string) *InactivateDeviceCertificateResponseBody {
	s.RequestId = &v
	return s
}

type InactivateDeviceCertificateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InactivateDeviceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InactivateDeviceCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s InactivateDeviceCertificateResponse) GoString() string {
	return s.String()
}

func (s *InactivateDeviceCertificateResponse) SetHeaders(v map[string]*string) *InactivateDeviceCertificateResponse {
	s.Headers = v
	return s
}

func (s *InactivateDeviceCertificateResponse) SetStatusCode(v int32) *InactivateDeviceCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *InactivateDeviceCertificateResponse) SetBody(v *InactivateDeviceCertificateResponseBody) *InactivateDeviceCertificateResponse {
	s.Body = v
	return s
}

type ListCaCertificateRequest struct {
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	PageNo         *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCaCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *ListCaCertificateRequest) SetMqttInstanceId(v string) *ListCaCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *ListCaCertificateRequest) SetPageNo(v string) *ListCaCertificateRequest {
	s.PageNo = &v
	return s
}

func (s *ListCaCertificateRequest) SetPageSize(v string) *ListCaCertificateRequest {
	s.PageSize = &v
	return s
}

type ListCaCertificateResponseBody struct {
	Data      *ListCaCertificateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCaCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ListCaCertificateResponseBody) SetData(v *ListCaCertificateResponseBodyData) *ListCaCertificateResponseBody {
	s.Data = v
	return s
}

func (s *ListCaCertificateResponseBody) SetRequestId(v string) *ListCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

type ListCaCertificateResponseBodyData struct {
	CaCertificateVOS []*ListCaCertificateResponseBodyDataCaCertificateVOS `json:"CaCertificateVOS,omitempty" xml:"CaCertificateVOS,omitempty" type:"Repeated"`
	PageNo           *int32                                               `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize         *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total            *int32                                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListCaCertificateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCaCertificateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCaCertificateResponseBodyData) SetCaCertificateVOS(v []*ListCaCertificateResponseBodyDataCaCertificateVOS) *ListCaCertificateResponseBodyData {
	s.CaCertificateVOS = v
	return s
}

func (s *ListCaCertificateResponseBodyData) SetPageNo(v int32) *ListCaCertificateResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListCaCertificateResponseBodyData) SetPageSize(v int32) *ListCaCertificateResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCaCertificateResponseBodyData) SetTotal(v int32) *ListCaCertificateResponseBodyData {
	s.Total = &v
	return s
}

type ListCaCertificateResponseBodyDataCaCertificateVOS struct {
	CaContent           *string `json:"CaContent,omitempty" xml:"CaContent,omitempty"`
	CaName              *string `json:"CaName,omitempty" xml:"CaName,omitempty"`
	RegistrationCode    *string `json:"RegistrationCode,omitempty" xml:"RegistrationCode,omitempty"`
	Sn                  *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ValidBegin          *string `json:"ValidBegin,omitempty" xml:"ValidBegin,omitempty"`
	ValidEnd            *string `json:"ValidEnd,omitempty" xml:"ValidEnd,omitempty"`
	VerificationContent *string `json:"VerificationContent,omitempty" xml:"VerificationContent,omitempty"`
}

func (s ListCaCertificateResponseBodyDataCaCertificateVOS) String() string {
	return tea.Prettify(s)
}

func (s ListCaCertificateResponseBodyDataCaCertificateVOS) GoString() string {
	return s.String()
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) SetCaContent(v string) *ListCaCertificateResponseBodyDataCaCertificateVOS {
	s.CaContent = &v
	return s
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) SetCaName(v string) *ListCaCertificateResponseBodyDataCaCertificateVOS {
	s.CaName = &v
	return s
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) SetRegistrationCode(v string) *ListCaCertificateResponseBodyDataCaCertificateVOS {
	s.RegistrationCode = &v
	return s
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) SetSn(v string) *ListCaCertificateResponseBodyDataCaCertificateVOS {
	s.Sn = &v
	return s
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) SetStatus(v string) *ListCaCertificateResponseBodyDataCaCertificateVOS {
	s.Status = &v
	return s
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) SetValidBegin(v string) *ListCaCertificateResponseBodyDataCaCertificateVOS {
	s.ValidBegin = &v
	return s
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) SetValidEnd(v string) *ListCaCertificateResponseBodyDataCaCertificateVOS {
	s.ValidEnd = &v
	return s
}

func (s *ListCaCertificateResponseBodyDataCaCertificateVOS) SetVerificationContent(v string) *ListCaCertificateResponseBodyDataCaCertificateVOS {
	s.VerificationContent = &v
	return s
}

type ListCaCertificateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCaCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *ListCaCertificateResponse) SetHeaders(v map[string]*string) *ListCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *ListCaCertificateResponse) SetStatusCode(v int32) *ListCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCaCertificateResponse) SetBody(v *ListCaCertificateResponseBody) *ListCaCertificateResponse {
	s.Body = v
	return s
}

type ListDeviceCertificateRequest struct {
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	PageNo         *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDeviceCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceCertificateRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateRequest) SetMqttInstanceId(v string) *ListDeviceCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *ListDeviceCertificateRequest) SetPageNo(v string) *ListDeviceCertificateRequest {
	s.PageNo = &v
	return s
}

func (s *ListDeviceCertificateRequest) SetPageSize(v string) *ListDeviceCertificateRequest {
	s.PageSize = &v
	return s
}

type ListDeviceCertificateResponseBody struct {
	Data      *ListDeviceCertificateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeviceCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateResponseBody) SetData(v *ListDeviceCertificateResponseBodyData) *ListDeviceCertificateResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceCertificateResponseBody) SetRequestId(v string) *ListDeviceCertificateResponseBody {
	s.RequestId = &v
	return s
}

type ListDeviceCertificateResponseBodyData struct {
	DeviceCertificateVOS []*ListDeviceCertificateResponseBodyDataDeviceCertificateVOS `json:"DeviceCertificateVOS,omitempty" xml:"DeviceCertificateVOS,omitempty" type:"Repeated"`
	PageNo               *int32                                                       `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32                                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total                *int32                                                       `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDeviceCertificateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceCertificateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateResponseBodyData) SetDeviceCertificateVOS(v []*ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) *ListDeviceCertificateResponseBodyData {
	s.DeviceCertificateVOS = v
	return s
}

func (s *ListDeviceCertificateResponseBodyData) SetPageNo(v int32) *ListDeviceCertificateResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyData) SetPageSize(v int32) *ListDeviceCertificateResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyData) SetTotal(v int32) *ListDeviceCertificateResponseBodyData {
	s.Total = &v
	return s
}

type ListDeviceCertificateResponseBodyDataDeviceCertificateVOS struct {
	CaSn          *string `json:"CaSn,omitempty" xml:"CaSn,omitempty"`
	DeviceContent *string `json:"DeviceContent,omitempty" xml:"DeviceContent,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceSn      *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ValidBegin    *string `json:"ValidBegin,omitempty" xml:"ValidBegin,omitempty"`
	ValidEnd      *string `json:"ValidEnd,omitempty" xml:"ValidEnd,omitempty"`
}

func (s ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) SetCaSn(v string) *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS {
	s.CaSn = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) SetDeviceContent(v string) *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS {
	s.DeviceContent = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) SetDeviceName(v string) *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS {
	s.DeviceName = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) SetDeviceSn(v string) *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS {
	s.DeviceSn = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) SetStatus(v string) *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS {
	s.Status = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) SetValidBegin(v string) *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS {
	s.ValidBegin = &v
	return s
}

func (s *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS) SetValidEnd(v string) *ListDeviceCertificateResponseBodyDataDeviceCertificateVOS {
	s.ValidEnd = &v
	return s
}

type ListDeviceCertificateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeviceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceCertificateResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateResponse) SetHeaders(v map[string]*string) *ListDeviceCertificateResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceCertificateResponse) SetStatusCode(v int32) *ListDeviceCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceCertificateResponse) SetBody(v *ListDeviceCertificateResponseBody) *ListDeviceCertificateResponse {
	s.Body = v
	return s
}

type ListDeviceCertificateByCaSnRequest struct {
	CaSn           *string `json:"CaSn,omitempty" xml:"CaSn,omitempty"`
	MqttInstanceId *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	PageNo         *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDeviceCertificateByCaSnRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceCertificateByCaSnRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateByCaSnRequest) SetCaSn(v string) *ListDeviceCertificateByCaSnRequest {
	s.CaSn = &v
	return s
}

func (s *ListDeviceCertificateByCaSnRequest) SetMqttInstanceId(v string) *ListDeviceCertificateByCaSnRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *ListDeviceCertificateByCaSnRequest) SetPageNo(v string) *ListDeviceCertificateByCaSnRequest {
	s.PageNo = &v
	return s
}

func (s *ListDeviceCertificateByCaSnRequest) SetPageSize(v string) *ListDeviceCertificateByCaSnRequest {
	s.PageSize = &v
	return s
}

type ListDeviceCertificateByCaSnResponseBody struct {
	Data      *ListDeviceCertificateByCaSnResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeviceCertificateByCaSnResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceCertificateByCaSnResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateByCaSnResponseBody) SetData(v *ListDeviceCertificateByCaSnResponseBodyData) *ListDeviceCertificateByCaSnResponseBody {
	s.Data = v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBody) SetRequestId(v string) *ListDeviceCertificateByCaSnResponseBody {
	s.RequestId = &v
	return s
}

type ListDeviceCertificateByCaSnResponseBodyData struct {
	DeviceCertificateVOS []*ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS `json:"DeviceCertificateVOS,omitempty" xml:"DeviceCertificateVOS,omitempty" type:"Repeated"`
	PageNo               *int32                                                             `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32                                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total                *int32                                                             `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDeviceCertificateByCaSnResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceCertificateByCaSnResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateByCaSnResponseBodyData) SetDeviceCertificateVOS(v []*ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) *ListDeviceCertificateByCaSnResponseBodyData {
	s.DeviceCertificateVOS = v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyData) SetPageNo(v int32) *ListDeviceCertificateByCaSnResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyData) SetPageSize(v int32) *ListDeviceCertificateByCaSnResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyData) SetTotal(v int32) *ListDeviceCertificateByCaSnResponseBodyData {
	s.Total = &v
	return s
}

type ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS struct {
	CaSn          *string `json:"CaSn,omitempty" xml:"CaSn,omitempty"`
	DeviceContent *string `json:"DeviceContent,omitempty" xml:"DeviceContent,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceSn      *string `json:"DeviceSn,omitempty" xml:"DeviceSn,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ValidBegin    *string `json:"ValidBegin,omitempty" xml:"ValidBegin,omitempty"`
	ValidEnd      *string `json:"ValidEnd,omitempty" xml:"ValidEnd,omitempty"`
}

func (s ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) SetCaSn(v string) *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS {
	s.CaSn = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) SetDeviceContent(v string) *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS {
	s.DeviceContent = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) SetDeviceName(v string) *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS {
	s.DeviceName = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) SetDeviceSn(v string) *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS {
	s.DeviceSn = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) SetStatus(v string) *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS {
	s.Status = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) SetValidBegin(v string) *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS {
	s.ValidBegin = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS) SetValidEnd(v string) *ListDeviceCertificateByCaSnResponseBodyDataDeviceCertificateVOS {
	s.ValidEnd = &v
	return s
}

type ListDeviceCertificateByCaSnResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeviceCertificateByCaSnResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeviceCertificateByCaSnResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceCertificateByCaSnResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateByCaSnResponse) SetHeaders(v map[string]*string) *ListDeviceCertificateByCaSnResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceCertificateByCaSnResponse) SetStatusCode(v int32) *ListDeviceCertificateByCaSnResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponse) SetBody(v *ListDeviceCertificateByCaSnResponseBody) *ListDeviceCertificateByCaSnResponse {
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
	Data      []*ListGroupIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupIdResponseBody) SetData(v []*ListGroupIdResponseBodyData) *ListGroupIdResponseBody {
	s.Data = v
	return s
}

func (s *ListGroupIdResponseBody) SetRequestId(v string) *ListGroupIdResponseBody {
	s.RequestId = &v
	return s
}

type ListGroupIdResponseBodyData struct {
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GroupId           *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	IndependentNaming *bool   `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UpdateTime        *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListGroupIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGroupIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGroupIdResponseBodyData) SetCreateTime(v int64) *ListGroupIdResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListGroupIdResponseBodyData) SetGroupId(v string) *ListGroupIdResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *ListGroupIdResponseBodyData) SetIndependentNaming(v bool) *ListGroupIdResponseBodyData {
	s.IndependentNaming = &v
	return s
}

func (s *ListGroupIdResponseBodyData) SetInstanceId(v string) *ListGroupIdResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListGroupIdResponseBodyData) SetUpdateTime(v int64) *ListGroupIdResponseBodyData {
	s.UpdateTime = &v
	return s
}

type ListGroupIdResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListGroupIdResponse) SetStatusCode(v int32) *ListGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupIdResponse) SetBody(v *ListGroupIdResponseBody) *ListGroupIdResponse {
	s.Body = v
	return s
}

type QueryMqttTraceDeviceRequest struct {
	BeginTime    *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MqttRegionId *string `json:"MqttRegionId,omitempty" xml:"MqttRegionId,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Reverse      *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
}

func (s QueryMqttTraceDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceDeviceRequest) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceDeviceRequest) SetBeginTime(v int64) *QueryMqttTraceDeviceRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetClientId(v string) *QueryMqttTraceDeviceRequest {
	s.ClientId = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetCurrentPage(v int32) *QueryMqttTraceDeviceRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetEndTime(v int64) *QueryMqttTraceDeviceRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetInstanceId(v string) *QueryMqttTraceDeviceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetMqttRegionId(v string) *QueryMqttTraceDeviceRequest {
	s.MqttRegionId = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetPageSize(v int32) *QueryMqttTraceDeviceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMqttTraceDeviceRequest) SetReverse(v bool) *QueryMqttTraceDeviceRequest {
	s.Reverse = &v
	return s
}

type QueryMqttTraceDeviceResponseBody struct {
	CurrentPage    *int32                                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DeviceInfoList []*QueryMqttTraceDeviceResponseBodyDeviceInfoList `json:"DeviceInfoList,omitempty" xml:"DeviceInfoList,omitempty" type:"Repeated"`
	PageSize       *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total          *int64                                            `json:"Total,omitempty" xml:"Total,omitempty"`
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

func (s *QueryMqttTraceDeviceResponseBody) SetDeviceInfoList(v []*QueryMqttTraceDeviceResponseBodyDeviceInfoList) *QueryMqttTraceDeviceResponseBody {
	s.DeviceInfoList = v
	return s
}

func (s *QueryMqttTraceDeviceResponseBody) SetPageSize(v int32) *QueryMqttTraceDeviceResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBody) SetRequestId(v string) *QueryMqttTraceDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBody) SetTotal(v int64) *QueryMqttTraceDeviceResponseBody {
	s.Total = &v
	return s
}

type QueryMqttTraceDeviceResponseBodyDeviceInfoList struct {
	Action     *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	ChannelId  *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryMqttTraceDeviceResponseBodyDeviceInfoList) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceDeviceResponseBodyDeviceInfoList) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) SetAction(v string) *QueryMqttTraceDeviceResponseBodyDeviceInfoList {
	s.Action = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) SetActionCode(v string) *QueryMqttTraceDeviceResponseBodyDeviceInfoList {
	s.ActionCode = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) SetActionInfo(v string) *QueryMqttTraceDeviceResponseBodyDeviceInfoList {
	s.ActionInfo = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) SetChannelId(v string) *QueryMqttTraceDeviceResponseBodyDeviceInfoList {
	s.ChannelId = &v
	return s
}

func (s *QueryMqttTraceDeviceResponseBodyDeviceInfoList) SetTime(v string) *QueryMqttTraceDeviceResponseBodyDeviceInfoList {
	s.Time = &v
	return s
}

type QueryMqttTraceDeviceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMqttTraceDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryMqttTraceDeviceResponse) SetStatusCode(v int32) *QueryMqttTraceDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMqttTraceDeviceResponse) SetBody(v *QueryMqttTraceDeviceResponseBody) *QueryMqttTraceDeviceResponse {
	s.Body = v
	return s
}

type QueryMqttTraceMessageOfClientRequest struct {
	BeginTime    *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MqttRegionId *string `json:"MqttRegionId,omitempty" xml:"MqttRegionId,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Reverse      *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
}

func (s QueryMqttTraceMessageOfClientRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessageOfClientRequest) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageOfClientRequest) SetBeginTime(v int64) *QueryMqttTraceMessageOfClientRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetClientId(v string) *QueryMqttTraceMessageOfClientRequest {
	s.ClientId = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetCurrentPage(v int32) *QueryMqttTraceMessageOfClientRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetEndTime(v int64) *QueryMqttTraceMessageOfClientRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetInstanceId(v string) *QueryMqttTraceMessageOfClientRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientRequest) SetMqttRegionId(v string) *QueryMqttTraceMessageOfClientRequest {
	s.MqttRegionId = &v
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
	MessageOfClientList []*QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList `json:"MessageOfClientList,omitempty" xml:"MessageOfClientList,omitempty" type:"Repeated"`
	PageSize            *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId           *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total               *int64                                                          `json:"Total,omitempty" xml:"Total,omitempty"`
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

func (s *QueryMqttTraceMessageOfClientResponseBody) SetMessageOfClientList(v []*QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) *QueryMqttTraceMessageOfClientResponseBody {
	s.MessageOfClientList = v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBody) SetPageSize(v int32) *QueryMqttTraceMessageOfClientResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBody) SetRequestId(v string) *QueryMqttTraceMessageOfClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBody) SetTotal(v int64) *QueryMqttTraceMessageOfClientResponseBody {
	s.Total = &v
	return s
}

type QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList struct {
	Action     *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) GoString() string {
	return s.String()
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

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) SetClientId(v string) *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList {
	s.ClientId = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) SetMsgId(v string) *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList {
	s.MsgId = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList) SetTime(v string) *QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList {
	s.Time = &v
	return s
}

type QueryMqttTraceMessageOfClientResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMqttTraceMessageOfClientResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryMqttTraceMessageOfClientResponse) SetStatusCode(v int32) *QueryMqttTraceMessageOfClientResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMqttTraceMessageOfClientResponse) SetBody(v *QueryMqttTraceMessageOfClientResponseBody) *QueryMqttTraceMessageOfClientResponse {
	s.Body = v
	return s
}

type QueryMqttTraceMessagePublishRequest struct {
	BeginTime    *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MqttRegionId *string `json:"MqttRegionId,omitempty" xml:"MqttRegionId,omitempty"`
	MsgId        *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s QueryMqttTraceMessagePublishRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessagePublishRequest) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessagePublishRequest) SetBeginTime(v int64) *QueryMqttTraceMessagePublishRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMqttTraceMessagePublishRequest) SetEndTime(v int64) *QueryMqttTraceMessagePublishRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMqttTraceMessagePublishRequest) SetInstanceId(v string) *QueryMqttTraceMessagePublishRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqttTraceMessagePublishRequest) SetMqttRegionId(v string) *QueryMqttTraceMessagePublishRequest {
	s.MqttRegionId = &v
	return s
}

func (s *QueryMqttTraceMessagePublishRequest) SetMsgId(v string) *QueryMqttTraceMessagePublishRequest {
	s.MsgId = &v
	return s
}

type QueryMqttTraceMessagePublishResponseBody struct {
	MessageTraceLists []*QueryMqttTraceMessagePublishResponseBodyMessageTraceLists `json:"MessageTraceLists,omitempty" xml:"MessageTraceLists,omitempty" type:"Repeated"`
	RequestId         *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMqttTraceMessagePublishResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessagePublishResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessagePublishResponseBody) SetMessageTraceLists(v []*QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) *QueryMqttTraceMessagePublishResponseBody {
	s.MessageTraceLists = v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBody) SetRequestId(v string) *QueryMqttTraceMessagePublishResponseBody {
	s.RequestId = &v
	return s
}

type QueryMqttTraceMessagePublishResponseBodyMessageTraceLists struct {
	Action     *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) GoString() string {
	return s.String()
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

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) SetClientId(v string) *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists {
	s.ClientId = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) SetMsgId(v string) *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists {
	s.MsgId = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists) SetTime(v string) *QueryMqttTraceMessagePublishResponseBodyMessageTraceLists {
	s.Time = &v
	return s
}

type QueryMqttTraceMessagePublishResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMqttTraceMessagePublishResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryMqttTraceMessagePublishResponse) SetStatusCode(v int32) *QueryMqttTraceMessagePublishResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMqttTraceMessagePublishResponse) SetBody(v *QueryMqttTraceMessagePublishResponseBody) *QueryMqttTraceMessagePublishResponse {
	s.Body = v
	return s
}

type QueryMqttTraceMessageSubscribeRequest struct {
	BeginTime    *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MqttRegionId *string `json:"MqttRegionId,omitempty" xml:"MqttRegionId,omitempty"`
	MsgId        *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Reverse      *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
}

func (s QueryMqttTraceMessageSubscribeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessageSubscribeRequest) GoString() string {
	return s.String()
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetBeginTime(v int64) *QueryMqttTraceMessageSubscribeRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetClientId(v string) *QueryMqttTraceMessageSubscribeRequest {
	s.ClientId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetCurrentPage(v int32) *QueryMqttTraceMessageSubscribeRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetEndTime(v int64) *QueryMqttTraceMessageSubscribeRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetInstanceId(v string) *QueryMqttTraceMessageSubscribeRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetMqttRegionId(v string) *QueryMqttTraceMessageSubscribeRequest {
	s.MqttRegionId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetMsgId(v string) *QueryMqttTraceMessageSubscribeRequest {
	s.MsgId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetPageSize(v int32) *QueryMqttTraceMessageSubscribeRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeRequest) SetReverse(v bool) *QueryMqttTraceMessageSubscribeRequest {
	s.Reverse = &v
	return s
}

type QueryMqttTraceMessageSubscribeResponseBody struct {
	CurrentPage       *int32                                                         `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	MessageTraceLists []*QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists `json:"MessageTraceLists,omitempty" xml:"MessageTraceLists,omitempty" type:"Repeated"`
	PageSize          *int32                                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total             *int64                                                         `json:"Total,omitempty" xml:"Total,omitempty"`
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

func (s *QueryMqttTraceMessageSubscribeResponseBody) SetMessageTraceLists(v []*QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) *QueryMqttTraceMessageSubscribeResponseBody {
	s.MessageTraceLists = v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) SetPageSize(v int32) *QueryMqttTraceMessageSubscribeResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) SetRequestId(v string) *QueryMqttTraceMessageSubscribeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBody) SetTotal(v int64) *QueryMqttTraceMessageSubscribeResponseBody {
	s.Total = &v
	return s
}

type QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists struct {
	Action     *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	MsgId      *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) String() string {
	return tea.Prettify(s)
}

func (s QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) GoString() string {
	return s.String()
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

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) SetClientId(v string) *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists {
	s.ClientId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) SetMsgId(v string) *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists {
	s.MsgId = &v
	return s
}

func (s *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists) SetTime(v string) *QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists {
	s.Time = &v
	return s
}

type QueryMqttTraceMessageSubscribeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMqttTraceMessageSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryMqttTraceMessageSubscribeResponse) SetStatusCode(v int32) *QueryMqttTraceMessageSubscribeResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySessionByClientIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QuerySessionByClientIdResponse) SetStatusCode(v int32) *QuerySessionByClientIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySessionByClientIdResponse) SetBody(v *QuerySessionByClientIdResponseBody) *QuerySessionByClientIdResponse {
	s.Body = v
	return s
}

type QueryTokenRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s QueryTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTokenRequest) GoString() string {
	return s.String()
}

func (s *QueryTokenRequest) SetInstanceId(v string) *QueryTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTokenRequest) SetToken(v string) *QueryTokenRequest {
	s.Token = &v
	return s
}

type QueryTokenResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TokenStatus *bool   `json:"TokenStatus,omitempty" xml:"TokenStatus,omitempty"`
}

func (s QueryTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTokenResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTokenResponseBody) SetRequestId(v string) *QueryTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTokenResponseBody) SetTokenStatus(v bool) *QueryTokenResponseBody {
	s.TokenStatus = &v
	return s
}

type QueryTokenResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryTokenResponse) SetStatusCode(v int32) *QueryTokenResponse {
	s.StatusCode = &v
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
	DeviceCredential *RefreshDeviceCredentialResponseBodyDeviceCredential `json:"DeviceCredential,omitempty" xml:"DeviceCredential,omitempty" type:"Struct"`
	RequestId        *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshDeviceCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshDeviceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshDeviceCredentialResponseBody) SetDeviceCredential(v *RefreshDeviceCredentialResponseBodyDeviceCredential) *RefreshDeviceCredentialResponseBody {
	s.DeviceCredential = v
	return s
}

func (s *RefreshDeviceCredentialResponseBody) SetRequestId(v string) *RefreshDeviceCredentialResponseBody {
	s.RequestId = &v
	return s
}

type RefreshDeviceCredentialResponseBodyDeviceCredential struct {
	ClientId              *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeviceAccessKeyId     *string `json:"DeviceAccessKeyId,omitempty" xml:"DeviceAccessKeyId,omitempty"`
	DeviceAccessKeySecret *string `json:"DeviceAccessKeySecret,omitempty" xml:"DeviceAccessKeySecret,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UpdateTime            *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s RefreshDeviceCredentialResponseBodyDeviceCredential) String() string {
	return tea.Prettify(s)
}

func (s RefreshDeviceCredentialResponseBodyDeviceCredential) GoString() string {
	return s.String()
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetClientId(v string) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.ClientId = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetCreateTime(v int64) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.CreateTime = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeyId(v string) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeyId = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeySecret(v string) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeySecret = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetInstanceId(v string) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.InstanceId = &v
	return s
}

func (s *RefreshDeviceCredentialResponseBodyDeviceCredential) SetUpdateTime(v int64) *RefreshDeviceCredentialResponseBodyDeviceCredential {
	s.UpdateTime = &v
	return s
}

type RefreshDeviceCredentialResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefreshDeviceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RefreshDeviceCredentialResponse) SetStatusCode(v int32) *RefreshDeviceCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshDeviceCredentialResponse) SetBody(v *RefreshDeviceCredentialResponseBody) *RefreshDeviceCredentialResponse {
	s.Body = v
	return s
}

type RegisterCaCertificateRequest struct {
	CaContent           *string `json:"CaContent,omitempty" xml:"CaContent,omitempty"`
	CaName              *string `json:"CaName,omitempty" xml:"CaName,omitempty"`
	MqttInstanceId      *string `json:"MqttInstanceId,omitempty" xml:"MqttInstanceId,omitempty"`
	VerificationContent *string `json:"VerificationContent,omitempty" xml:"VerificationContent,omitempty"`
}

func (s RegisterCaCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *RegisterCaCertificateRequest) SetCaContent(v string) *RegisterCaCertificateRequest {
	s.CaContent = &v
	return s
}

func (s *RegisterCaCertificateRequest) SetCaName(v string) *RegisterCaCertificateRequest {
	s.CaName = &v
	return s
}

func (s *RegisterCaCertificateRequest) SetMqttInstanceId(v string) *RegisterCaCertificateRequest {
	s.MqttInstanceId = &v
	return s
}

func (s *RegisterCaCertificateRequest) SetVerificationContent(v string) *RegisterCaCertificateRequest {
	s.VerificationContent = &v
	return s
}

type RegisterCaCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sn        *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s RegisterCaCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterCaCertificateResponseBody) SetRequestId(v string) *RegisterCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterCaCertificateResponseBody) SetSn(v string) *RegisterCaCertificateResponseBody {
	s.Sn = &v
	return s
}

type RegisterCaCertificateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RegisterCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterCaCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *RegisterCaCertificateResponse) SetHeaders(v map[string]*string) *RegisterCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *RegisterCaCertificateResponse) SetStatusCode(v int32) *RegisterCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterCaCertificateResponse) SetBody(v *RegisterCaCertificateResponseBody) *RegisterCaCertificateResponse {
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
	DeviceCredential *RegisterDeviceCredentialResponseBodyDeviceCredential `json:"DeviceCredential,omitempty" xml:"DeviceCredential,omitempty" type:"Struct"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterDeviceCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDeviceCredentialResponseBody) SetDeviceCredential(v *RegisterDeviceCredentialResponseBodyDeviceCredential) *RegisterDeviceCredentialResponseBody {
	s.DeviceCredential = v
	return s
}

func (s *RegisterDeviceCredentialResponseBody) SetRequestId(v string) *RegisterDeviceCredentialResponseBody {
	s.RequestId = &v
	return s
}

type RegisterDeviceCredentialResponseBodyDeviceCredential struct {
	ClientId              *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeviceAccessKeyId     *string `json:"DeviceAccessKeyId,omitempty" xml:"DeviceAccessKeyId,omitempty"`
	DeviceAccessKeySecret *string `json:"DeviceAccessKeySecret,omitempty" xml:"DeviceAccessKeySecret,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UpdateTime            *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s RegisterDeviceCredentialResponseBodyDeviceCredential) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceCredentialResponseBodyDeviceCredential) GoString() string {
	return s.String()
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetClientId(v string) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.ClientId = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetCreateTime(v int64) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.CreateTime = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeyId(v string) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeyId = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetDeviceAccessKeySecret(v string) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.DeviceAccessKeySecret = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetInstanceId(v string) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.InstanceId = &v
	return s
}

func (s *RegisterDeviceCredentialResponseBodyDeviceCredential) SetUpdateTime(v int64) *RegisterDeviceCredentialResponseBodyDeviceCredential {
	s.UpdateTime = &v
	return s
}

type RegisterDeviceCredentialResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RegisterDeviceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RegisterDeviceCredentialResponse) SetStatusCode(v int32) *RegisterDeviceCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterDeviceCredentialResponse) SetBody(v *RegisterDeviceCredentialResponseBody) *RegisterDeviceCredentialResponse {
	s.Body = v
	return s
}

type RevokeTokenRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s RevokeTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeTokenRequest) GoString() string {
	return s.String()
}

func (s *RevokeTokenRequest) SetInstanceId(v string) *RevokeTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeTokenRequest) SetToken(v string) *RevokeTokenRequest {
	s.Token = &v
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RevokeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RevokeTokenResponse) SetStatusCode(v int32) *RevokeTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeTokenResponse) SetBody(v *RevokeTokenResponseBody) *RevokeTokenResponse {
	s.Body = v
	return s
}

type SendMessageRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MqttTopic  *string `json:"MqttTopic,omitempty" xml:"MqttTopic,omitempty"`
	Payload    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetInstanceId(v string) *SendMessageRequest {
	s.InstanceId = &v
	return s
}

func (s *SendMessageRequest) SetMqttTopic(v string) *SendMessageRequest {
	s.MqttTopic = &v
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SendMessageResponse) SetStatusCode(v int32) *SendMessageResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnRegisterDeviceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UnRegisterDeviceCredentialResponse) SetStatusCode(v int32) *UnRegisterDeviceCredentialResponse {
	s.StatusCode = &v
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

func (client *Client) ActiveCaCertificateWithOptions(request *ActiveCaCertificateRequest, runtime *util.RuntimeOptions) (_result *ActiveCaCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MqttInstanceId)) {
		query["MqttInstanceId"] = request.MqttInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		query["Sn"] = request.Sn
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ActiveCaCertificate"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActiveCaCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActiveCaCertificate(request *ActiveCaCertificateRequest) (_result *ActiveCaCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActiveCaCertificateResponse{}
	_body, _err := client.ActiveCaCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ActiveDeviceCertificateWithOptions(request *ActiveDeviceCertificateRequest, runtime *util.RuntimeOptions) (_result *ActiveDeviceCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaSn)) {
		query["CaSn"] = request.CaSn
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceSn)) {
		query["DeviceSn"] = request.DeviceSn
	}

	if !tea.BoolValue(util.IsUnset(request.MqttInstanceId)) {
		query["MqttInstanceId"] = request.MqttInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ActiveDeviceCertificate"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActiveDeviceCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActiveDeviceCertificate(request *ActiveDeviceCertificateRequest) (_result *ActiveDeviceCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActiveDeviceCertificateResponse{}
	_body, _err := client.ActiveDeviceCertificateWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Actions)) {
		query["Actions"] = request.Actions
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		query["Resources"] = request.Resources
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyToken"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIdList)) {
		query["ClientIdList"] = request.ClientIdList
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchQuerySessionByClientIds"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchQuerySessionByClientIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroupId"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteCaCertificateWithOptions(request *DeleteCaCertificateRequest, runtime *util.RuntimeOptions) (_result *DeleteCaCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MqttInstanceId)) {
		query["MqttInstanceId"] = request.MqttInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		query["Sn"] = request.Sn
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCaCertificate"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCaCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCaCertificate(request *DeleteCaCertificateRequest) (_result *DeleteCaCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCaCertificateResponse{}
	_body, _err := client.DeleteCaCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDeviceCertificateWithOptions(request *DeleteDeviceCertificateRequest, runtime *util.RuntimeOptions) (_result *DeleteDeviceCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaSn)) {
		query["CaSn"] = request.CaSn
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceSn)) {
		query["DeviceSn"] = request.DeviceSn
	}

	if !tea.BoolValue(util.IsUnset(request.MqttInstanceId)) {
		query["MqttInstanceId"] = request.MqttInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDeviceCertificate"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDeviceCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDeviceCertificate(request *DeleteDeviceCertificateRequest) (_result *DeleteDeviceCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDeviceCertificateResponse{}
	_body, _err := client.DeleteDeviceCertificateWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroupId"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetCaCertificateWithOptions(request *GetCaCertificateRequest, runtime *util.RuntimeOptions) (_result *GetCaCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCaCertificate"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCaCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCaCertificate(request *GetCaCertificateRequest) (_result *GetCaCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCaCertificateResponse{}
	_body, _err := client.GetCaCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeviceCertificateWithOptions(request *GetDeviceCertificateRequest, runtime *util.RuntimeOptions) (_result *GetDeviceCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceCertificate"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeviceCertificate(request *GetDeviceCertificateRequest) (_result *GetDeviceCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeviceCertificateResponse{}
	_body, _err := client.GetDeviceCertificateWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeviceCredential"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeviceCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetRegisterCodeWithOptions(request *GetRegisterCodeRequest, runtime *util.RuntimeOptions) (_result *GetRegisterCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRegisterCode"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRegisterCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRegisterCode(request *GetRegisterCodeRequest) (_result *GetRegisterCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRegisterCodeResponse{}
	_body, _err := client.GetRegisterCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InactivateCaCertificateWithOptions(request *InactivateCaCertificateRequest, runtime *util.RuntimeOptions) (_result *InactivateCaCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MqttInstanceId)) {
		query["MqttInstanceId"] = request.MqttInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		query["Sn"] = request.Sn
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InactivateCaCertificate"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InactivateCaCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InactivateCaCertificate(request *InactivateCaCertificateRequest) (_result *InactivateCaCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InactivateCaCertificateResponse{}
	_body, _err := client.InactivateCaCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InactivateDeviceCertificateWithOptions(request *InactivateDeviceCertificateRequest, runtime *util.RuntimeOptions) (_result *InactivateDeviceCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaSn)) {
		query["CaSn"] = request.CaSn
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceSn)) {
		query["DeviceSn"] = request.DeviceSn
	}

	if !tea.BoolValue(util.IsUnset(request.MqttInstanceId)) {
		query["MqttInstanceId"] = request.MqttInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InactivateDeviceCertificate"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InactivateDeviceCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InactivateDeviceCertificate(request *InactivateDeviceCertificateRequest) (_result *InactivateDeviceCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InactivateDeviceCertificateResponse{}
	_body, _err := client.InactivateDeviceCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCaCertificateWithOptions(request *ListCaCertificateRequest, runtime *util.RuntimeOptions) (_result *ListCaCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCaCertificate"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCaCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCaCertificate(request *ListCaCertificateRequest) (_result *ListCaCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCaCertificateResponse{}
	_body, _err := client.ListCaCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceCertificateWithOptions(request *ListDeviceCertificateRequest, runtime *util.RuntimeOptions) (_result *ListDeviceCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeviceCertificate"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeviceCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceCertificate(request *ListDeviceCertificateRequest) (_result *ListDeviceCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceCertificateResponse{}
	_body, _err := client.ListDeviceCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeviceCertificateByCaSnWithOptions(request *ListDeviceCertificateByCaSnRequest, runtime *util.RuntimeOptions) (_result *ListDeviceCertificateByCaSnResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeviceCertificateByCaSn"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeviceCertificateByCaSnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceCertificateByCaSn(request *ListDeviceCertificateByCaSnRequest) (_result *ListDeviceCertificateByCaSnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceCertificateByCaSnResponse{}
	_body, _err := client.ListDeviceCertificateByCaSnWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroupId"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MqttRegionId)) {
		query["MqttRegionId"] = request.MqttRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		query["Reverse"] = request.Reverse
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMqttTraceDevice"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMqttTraceDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MqttRegionId)) {
		query["MqttRegionId"] = request.MqttRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		query["Reverse"] = request.Reverse
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMqttTraceMessageOfClient"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMqttTraceMessageOfClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MqttRegionId)) {
		query["MqttRegionId"] = request.MqttRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		query["MsgId"] = request.MsgId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMqttTraceMessagePublish"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMqttTraceMessagePublishResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MqttRegionId)) {
		query["MqttRegionId"] = request.MqttRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgId)) {
		query["MsgId"] = request.MsgId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		query["Reverse"] = request.Reverse
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMqttTraceMessageSubscribe"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMqttTraceMessageSubscribeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySessionByClientId"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySessionByClientIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryToken"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshDeviceCredential"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshDeviceCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) RegisterCaCertificateWithOptions(request *RegisterCaCertificateRequest, runtime *util.RuntimeOptions) (_result *RegisterCaCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaContent)) {
		query["CaContent"] = request.CaContent
	}

	if !tea.BoolValue(util.IsUnset(request.CaName)) {
		query["CaName"] = request.CaName
	}

	if !tea.BoolValue(util.IsUnset(request.MqttInstanceId)) {
		query["MqttInstanceId"] = request.MqttInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.VerificationContent)) {
		query["VerificationContent"] = request.VerificationContent
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterCaCertificate"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterCaCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterCaCertificate(request *RegisterCaCertificateRequest) (_result *RegisterCaCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterCaCertificateResponse{}
	_body, _err := client.RegisterCaCertificateWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterDeviceCredential"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterDeviceCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeToken"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MqttTopic)) {
		query["MqttTopic"] = request.MqttTopic
	}

	if !tea.BoolValue(util.IsUnset(request.Payload)) {
		query["Payload"] = request.Payload
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendMessage"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnRegisterDeviceCredential"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnRegisterDeviceCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
