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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActiveCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActiveDeviceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The permission type of the token. Valid values:
	//
	// *   **R**: read-only. You can only subscribe to the specified topics.
	// *   **W**: write-only. You can only send messages to the specified topics.
	// *   **R,W**: read and write. You can send messages to and subscribe to the specified topics. Separate **R** and **W** with a comma (,).
	Actions *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	// The timestamp that indicates the point in time when the token expires. Unit: milliseconds. The minimum validity period of a token is 60 seconds, and the maximum validity period of a token is 30 days. If you specify a validity period of more than 30 days for a token, no errors are returned. However, the token is valid only for 30 days.
	//
	// For example, you want to specify a validity period of 60 seconds for a token. If the current system timestamp is 1609434061000, you must set this parameter to **1609434121000**. The value is calculated by using the following formula: 1609434061000 + 60 x 1000 = 1609434121000.
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can obtain the instance ID on the **Instance Details** page that corresponds to the instance in the [ApsaraMQ for MQTT console](https://mqtt.console.aliyun.com/).
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The topics on the ApsaraMQ for MQTT instance. Separate multiple topics with commas (,). Each token can be used to access up to 100 topics. Multiple topics are sorted in alphabetic order. MQTT wildcards, including single-level wildcards represented by plus signs (+) and multi-level wildcards represented by number signs (#), can be used for the Resources parameter that you register to apply for a token.
	//
	// For example, if you set the **Resources** parameter to Topic1/+ when you apply for a token, the ApsaraMQ for MQTT client can manage the topics in Topic1/xxx. If you set the **Resources** parameter to Topic1/# when you apply for a token, the ApsaraMQ for MQTT client can manage topics of any level in Topic1/xxx/xxx/xxx.
	//
	// >  ApsaraMQ for MQTT supports subtopics. You can specify subtopics in the code for messaging instead of configuring them in the ApsaraMQ for MQTT console. Forward slashes (/) are used to separate topics of different levels. For more information, see [Terms](~~42420~~).
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
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
	// The request ID. This parameter is a common parameter.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The token that is returned by the ApsaraMQ for MQTT broker.
	//
	// >  Do not assume the length, format, or rule of the token to return. The actual returned token shall prevail.
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ApsaraMQ for MQTT clients.
	ClientIdList []*string `json:"ClientIdList,omitempty" xml:"ClientIdList,omitempty" type:"Repeated"`
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can obtain the instance ID on the **Instance Details** page that corresponds to the instance in the [ApsaraMQ for MQTT console](https://mqtt.console.aliyun.com).
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// The status list of all queried ApsaraMQ for MQTT clients.
	OnlineStatusList []*BatchQuerySessionByClientIdsResponseBodyOnlineStatusList `json:"OnlineStatusList,omitempty" xml:"OnlineStatusList,omitempty" type:"Repeated"`
	// The request ID. This parameter is a common parameter.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The ID of the ApsaraMQ for MQTT client. For more information about client IDs, see [Terms](~~42420~~).
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// Indicates whether the ApsaraMQ for MQTT client is online. Valid values:
	//
	// *   **true**
	// *   **false**
	OnlineStatus *bool `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQuerySessionByClientIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the group that you want to create. The group ID must meet the following conventions:
	//
	// *   The ID must be 7 to 64 characters in length. It must start with GID\_ or GID- and can contain only letters, digits, hyphens (-), and underscores (\_).
	// *   The ID cannot be changed after the group is created. For more information, see [Terms](~~42420~~).
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance to which the group belongs.
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
	// The request ID. This parameter is a common parameter.
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeviceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the group that you want to delete.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance from which you want to delete a group.
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
	// The request ID. This parameter is a common parameter.
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The client ID of the device whose access credential you want to query.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can obtain the instance ID on the **Instance Details** page that corresponds to the instance in the ApsaraMQ for MQTT console.
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
	// The information about the access credential of the device.
	DeviceCredential *GetDeviceCredentialResponseBodyDeviceCredential `json:"DeviceCredential,omitempty" xml:"DeviceCredential,omitempty" type:"Struct"`
	// The request ID. This parameter is a common parameter.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The client ID of the device.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The timestamp that indicates when the access credential of the device was created. Unit: milliseconds.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The AccessKey ID of the device.
	DeviceAccessKeyId *string `json:"DeviceAccessKeyId,omitempty" xml:"DeviceAccessKeyId,omitempty"`
	// The AccessKey secret of the device.
	DeviceAccessKeySecret *string `json:"DeviceAccessKeySecret,omitempty" xml:"DeviceAccessKeySecret,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The timestamp that indicates when the access credential of the device was last updated. The value of this parameter is a UNIX timestamp in milliseconds.
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRegisterCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InactivateCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InactivateDeviceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceCertificateByCaSnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListDeviceCredentialClientIdRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageNo     *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDeviceCredentialClientIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceCredentialClientIdRequest) GoString() string {
	return s.String()
}

func (s *ListDeviceCredentialClientIdRequest) SetGroupId(v string) *ListDeviceCredentialClientIdRequest {
	s.GroupId = &v
	return s
}

func (s *ListDeviceCredentialClientIdRequest) SetInstanceId(v string) *ListDeviceCredentialClientIdRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDeviceCredentialClientIdRequest) SetNextToken(v string) *ListDeviceCredentialClientIdRequest {
	s.NextToken = &v
	return s
}

func (s *ListDeviceCredentialClientIdRequest) SetPageNo(v string) *ListDeviceCredentialClientIdRequest {
	s.PageNo = &v
	return s
}

func (s *ListDeviceCredentialClientIdRequest) SetPageSize(v string) *ListDeviceCredentialClientIdRequest {
	s.PageSize = &v
	return s
}

type ListDeviceCredentialClientIdResponseBody struct {
	DeviceCredentialClientIdList *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList `json:"DeviceCredentialClientIdList,omitempty" xml:"DeviceCredentialClientIdList,omitempty" type:"Struct"`
	RequestId                    *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeviceCredentialClientIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceCredentialClientIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceCredentialClientIdResponseBody) SetDeviceCredentialClientIdList(v *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) *ListDeviceCredentialClientIdResponseBody {
	s.DeviceCredentialClientIdList = v
	return s
}

func (s *ListDeviceCredentialClientIdResponseBody) SetRequestId(v string) *ListDeviceCredentialClientIdResponseBody {
	s.RequestId = &v
	return s
}

type ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList struct {
	ClientIdList []*string `json:"ClientIdList,omitempty" xml:"ClientIdList,omitempty" type:"Repeated"`
	NextToken    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageNo       *string   `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize     *string   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total        *string   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) GoString() string {
	return s.String()
}

func (s *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) SetClientIdList(v []*string) *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList {
	s.ClientIdList = v
	return s
}

func (s *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) SetNextToken(v string) *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList {
	s.NextToken = &v
	return s
}

func (s *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) SetPageNo(v string) *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList {
	s.PageNo = &v
	return s
}

func (s *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) SetPageSize(v string) *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList {
	s.PageSize = &v
	return s
}

func (s *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) SetTotal(v string) *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList {
	s.Total = &v
	return s
}

type ListDeviceCredentialClientIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceCredentialClientIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceCredentialClientIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeviceCredentialClientIdResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceCredentialClientIdResponse) SetHeaders(v map[string]*string) *ListDeviceCredentialClientIdResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceCredentialClientIdResponse) SetStatusCode(v int32) *ListDeviceCredentialClientIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceCredentialClientIdResponse) SetBody(v *ListDeviceCredentialClientIdResponseBody) *ListDeviceCredentialClientIdResponse {
	s.Body = v
	return s
}

type ListGroupIdRequest struct {
	// The ID of the ApsaraMQ for MQTT instance whose groups you want to query.
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
	// The details of a queried group.
	Data []*ListGroupIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID. This parameter is a common parameter.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The time when the group was created.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The queried group that belongs to the ApsaraMQ for MQTT instance.
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Indicates whether a separate namespace is configured for the ApsaraMQ for MQTT instance. Valid values:
	//
	// *   **true**: A separate namespace is configured for the ApsaraMQ for MQTT instance. Resource names must be unique within an ApsaraMQ for MQTT instance but can be the same across ApsaraMQ for MQTT instances.
	// *   **false**: No separate namespace is configured for the ApsaraMQ for MQTT instance. Resource names must be globally unique within an ApsaraMQ for MQTT instance and across ApsaraMQ for MQTT instances.
	IndependentNaming *bool `json:"IndependentNaming,omitempty" xml:"IndependentNaming,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance to which the group belongs.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the group was last updated.
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The beginning of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The client ID of the device whose trace you want to query.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The number of the page to return. Pages start from page 1. If the input parameter value is greater than the total number of pages, the returned result is empty.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can view the instance ID in the **Basic Information** section on the **Instance Details** page that corresponds to the instance in the ApsaraMQ for MQTT console.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the ApsaraMQ for MQTT instance resides. For more information, see [Endpoints](~~181438~~).
	MqttRegionId *string `json:"MqttRegionId,omitempty" xml:"MqttRegionId,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether the returned results are displayed in reverse chronological order. Valid values:
	//
	// *   **true**: The returned results are displayed in reverse time order of actions on the device. This means that the information about the latest action on the device is displayed as the first entry and the information about the earliest action on the device is displayed as the last entry.
	// *   **false**: The returned results are displayed in time order of actions on the device. This means that the information about the earliest action on the device is displayed as the first entry and the information about the latest action on the device is displayed as the last entry.
	//
	// If you do not specify this parameter, the returned results are displayed in time order of actions on the device by default.
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
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
	// The page number of the returned page.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The details of the action on the device.
	DeviceInfoList []*QueryMqttTraceDeviceResponseBodyDeviceInfoList `json:"DeviceInfoList,omitempty" xml:"DeviceInfoList,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID. You can use the ID to troubleshoot issues. This parameter is a common parameter.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned actions on the device.
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// The action on the device. Valid values:
	//
	// *   **connect**: The ApsaraMQ for MQTT client requests a connection to the ApsaraMQ for MQTT broker.
	// *   **close**: The TCP connection is closed.
	// *   **disconnect**: The ApsaraMQ for MQTT client requests a disconnection from the ApsaraMQ for MQTT broker.
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The returned code for the action on the device. Valid values:
	//
	// *   **mqtt.trace.action.connect**: This value is returned if the value of Action is **connect**.
	// *   **mqtt.trace.action.close**: This value is returned if the value of Action is **close**.
	// *   **mqtt.trace.action.disconnect**: This value is returned if the value of Action is **disconnect**.
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	// The returned information for the action on the device. Valid values:
	//
	// *   **accepted**: The ApsaraMQ for MQTT broker accepts the connection request from the ApsaraMQ for MQTT client.
	// *   **not authorized**: The TCP connection is closed because the permission verification of the client to access the instance fails.
	// *   **clientId conflict**: The TCP connection is closed due to a conflict in the ID of the ApsaraMQ for MQTT client.
	// *   **resource auth failed**: The TCP connection is closed because the permission verification for the ApsaraMQ for MQTT client to access the topic or group fails.
	// *   **no heart**: The TCP connection is closed because no heartbeat is detected on the client.
	// *   **closed by client**: The TCP connection is closed because an exception occurs on the client.
	// *   **disconnected by client**: The client requests a disconnection.
	// *   **invalid param**: The TCP connection is closed due to invalid request parameters.
	// *   **Socket IOException**: The TCP connection is closed due to network jitter or packet loss.
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	// The connection ID.
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The time when the action occurred on the device.
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMqttTraceDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The beginning of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The client ID of the device whose messages you want to query.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The number of the page to return. Pages start from page 1. If the input parameter value is greater than the total number of pages, the returned result is empty.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can view the instance ID in the **Basic Information** section of the **Instance Details** page that corresponds to the instance in the ApsaraMQ for MQTT console.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the ApsaraMQ for MQTT instance resides. For more information, see [Endpoints](~~181438~~).
	MqttRegionId *string `json:"MqttRegionId,omitempty" xml:"MqttRegionId,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether the returned results are displayed in reverse chronological order. Valid values:
	//
	// *   **true**: The returned results are displayed in reverse order of the time when messages are sent or received. This means that the latest sent or received message is displayed as the first entry and the earliest sent or received message is displayed as the last entry.
	// *   **false**: The returned results are displayed in order of the time when messages are sent or received. This means that the earliest sent or received message is displayed as the first entry and the latest sent or received message is displayed as the last entry.
	//
	// If this parameter is not specified, the returned results are displayed in order of the time when messages are sent or received.
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
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
	// The page number of the returned page.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The returned messages.
	MessageOfClientList []*QueryMqttTraceMessageOfClientResponseBodyMessageOfClientList `json:"MessageOfClientList,omitempty" xml:"MessageOfClientList,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID. You can use the ID to troubleshoot issues. This parameter is a common parameter.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of messages returned.
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// The action on the message. Valid values:
	//
	// *   **pub_mqtt**: The ApsaraMQ for MQTT client sends the message.
	// *   **sub**: The ApsaraMQ for MQTT client subscribes to the message.
	// *   **push_offline**: The ApsaraMQ for MQTT broker pushes the offline message to the ApsaraMQ for MQTT client.
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The returned code for the action on the message. Valid values:
	//
	// *   **mqtt.trace.action.msg.pub.mqtt**: This value is returned if the value of Action is **pub_mqtt**.
	// *   **mqtt.trace.action.msg.sub**: This value is returned if the value of Action is **sub**.
	// *   **mqtt.trace.action.msg.push.offline**: This value is returned if the value of Action is **push_offline**.
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	// The information returned for the action on the message. Valid values:
	//
	// *   **Pub From Mqtt Client**: This value is returned if the value of Action is **pub_mqtt**.
	// *   **Push To Mqtt Client**: This value is returned if the value of Action is **sub**.
	// *   **Push Offline Msg To Mqtt Client**: This value is returned if the value of Action is **push_offline**.
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	// The client ID of the device.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The message ID.
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The time when the message was sent or received.
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMqttTraceMessageOfClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The beginning of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can view the instance ID in the **Basic Information** section on the **Instance Details** page that corresponds to the instance in the ApsaraMQ for MQTT console.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the ApsaraMQ for MQTT instance resides. For more information, see [Endpoints](~~181438~~).
	MqttRegionId *string `json:"MqttRegionId,omitempty" xml:"MqttRegionId,omitempty"`
	// The message ID.
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
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
	// The message traces.
	MessageTraceLists []*QueryMqttTraceMessagePublishResponseBodyMessageTraceLists `json:"MessageTraceLists,omitempty" xml:"MessageTraceLists,omitempty" type:"Repeated"`
	// The request ID. You can use the ID to troubleshoot issues. This parameter is a common parameter.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The action on the message. Valid values:
	//
	// *   **pub_mqtt**: indicates that the message was sent by an ApsaraMQ for MQTT client.
	// *   **pub_mq**: indicates that the message was sent by an ApsaraMQ for RocketMQ client.
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The returned code for the action on the message. Valid values:
	//
	// *   **mqtt.trace.action.msg.pub.mqtt**: This value is returned if the value of Action is **pub_mqtt**.
	// *   **mqtt.trace.action.msg.pub.mq**: This value is returned if the value of Action is **pub_mq**.
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	// The returned information for the action on the message. Valid values:
	//
	// *   **Pub From Mqtt Client**: This value is returned if the value of Action is **pub_mqtt**.
	// *   **Pub From MQ**: This value is returned if the value of Action is **pub_mq**.
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	// The ID of the client that sends the message.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The message ID.
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The time when the message was sent.
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMqttTraceMessagePublishResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The beginning of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The ID of the client that subscribes to the message. If you do not specify this parameter, the IDs of all clients that subscribe to the message are returned.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The number of the page to return. Pages start from page 1. If the input parameter value is greater than the total number of pages, the returned result is empty.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can view the instance ID in the **Basic Information** section of the **Instance Details** page that corresponds to the instance in the ApsaraMQ for MQTT console.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the ApsaraMQ for MQTT instance resides. For more information, see [Endpoints](~~181438~~).
	MqttRegionId *string `json:"MqttRegionId,omitempty" xml:"MqttRegionId,omitempty"`
	// The message ID.
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether the returned results are displayed in reverse chronological order. Valid values:
	//
	// *   **true**: The returned results are displayed in reverse order of the time when messages are delivered. This means that the latest consumed message is displayed as the first entry and the earliest consumed message is displayed as the last entry.
	// *   **false**: The returned results are displayed in order of the time when messages are delivered. This means that the earliest consumed message is displayed as the first entry and the latest consumed message is displayed as the last entry.
	//
	// If you do not specify this parameter, the returned results are displayed in order of time when messages are delivered.
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
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
	// The page number of the returned page.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The delivery trace of the queried message.
	MessageTraceLists []*QueryMqttTraceMessageSubscribeResponseBodyMessageTraceLists `json:"MessageTraceLists,omitempty" xml:"MessageTraceLists,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique ID that the system generates for the request. You can use the ID to troubleshoot issues. This parameter is a common parameter.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned delivery traces.
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// The action on the message. Valid values:
	//
	// *   **sub**: The ApsaraMQ for MQTT client subscribes to the message.
	// *   **push_offline**: The ApsaraMQ for MQTT broker pushes the offline message to the ApsaraMQ for MQTT client.
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The code returned for the action on the message. Valid values:
	//
	// *   **mqtt.trace.action.msg.sub**: The value that is returned if the value of Action is **sub**.
	// *   **mqtt.trace.action.msg.push.offline**: The value that is returned if the value of Action is **push_offline**.
	ActionCode *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	// The returned information for the action on the message. Valid values:
	//
	// *   **Push To Mqtt Client**: The value that is returned if the value of Action is **sub**.
	// *   **Push Offline Msg To Mqtt Client**: The value that is returned if the value of Action is **push_offline**.
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	// The ID of the client that subscribes to the message.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The message ID.
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The time when the message was delivered.
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMqttTraceMessageSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the ApsaraMQ for MQTT client that you want to query.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can obtain the instance ID on the **Instance Details** page that corresponds to the instance in the [ApsaraMQ for MQTT console](https://mqtt.console.aliyun.com).
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
	// Indicates whether the ApsaraMQ for MQTT client is connected to the ApsaraMQ for MQTT broker. Valid values:
	//
	// *   **true**
	// *   **false**
	OnlineStatus *bool `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// The request ID. This parameter is a common parameter.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySessionByClientIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can obtain the instance ID on the **Instance Details** page that corresponds to the instance in the [ApsaraMQ for MQTT console](https://mqtt.console.aliyun.com/).
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The token that you want to query.
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
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
	// The unique ID that the system generates for the request. This parameter is a common parameter.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the queried token. Valid values:
	//
	// *   **true**: indicates the token is valid.
	// *   **false**: indicates the token is invalid.
	TokenStatus *bool `json:"TokenStatus,omitempty" xml:"TokenStatus,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The client ID of the device whose access credential you want to update.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can obtain the instance ID on the **Instance Details** page that corresponds to the instance in the ApsaraMQ for MQTT console.
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
	// The access credential of the device.
	DeviceCredential *RefreshDeviceCredentialResponseBodyDeviceCredential `json:"DeviceCredential,omitempty" xml:"DeviceCredential,omitempty" type:"Struct"`
	// The request ID. This parameter is a common parameter.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The client ID of the device.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The timestamp that indicates when the access credential of the device was created. The value of this parameter is a UNIX timestamp in milliseconds.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The AccessKey ID of the device.
	DeviceAccessKeyId *string `json:"DeviceAccessKeyId,omitempty" xml:"DeviceAccessKeyId,omitempty"`
	// The AccessKey secret of the device.
	DeviceAccessKeySecret *string `json:"DeviceAccessKeySecret,omitempty" xml:"DeviceAccessKeySecret,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The timestamp that indicates when the access credential of the device was last updated. The value of this parameter is a UNIX timestamp in milliseconds.
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshDeviceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The client ID of the device for which you want to create an access credential.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can obtain the instance ID on the **Instance Details** page that corresponds to the instance in the ApsaraMQ for MQTT console.
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
	// The access credential of the device.
	DeviceCredential *RegisterDeviceCredentialResponseBodyDeviceCredential `json:"DeviceCredential,omitempty" xml:"DeviceCredential,omitempty" type:"Struct"`
	// The request ID. This parameter is a common parameter.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The client ID of the device.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The timestamp that indicates when the access credential of the device was created. Unit: milliseconds.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The AccessKey ID of the device.
	DeviceAccessKeyId *string `json:"DeviceAccessKeyId,omitempty" xml:"DeviceAccessKeyId,omitempty"`
	// The AccessKey secret of the device.
	DeviceAccessKeySecret *string `json:"DeviceAccessKeySecret,omitempty" xml:"DeviceAccessKeySecret,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The timestamp that indicates when the access credential of the device was last updated. Unit: milliseconds.
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterDeviceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can obtain the instance ID on the **Instance Details** page that corresponds to the instance in the [ApsaraMQ for MQTT console](https://mqtt.console.aliyun.com/).
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The token that you want to revoke.
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
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
	// The request ID. This parameter is a common parameter.
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can view the instance ID in the **Basic Information** section on the **Instance Details** page that corresponds to the instance in the [ApsaraMQ for MQTT console](https://mqtt.console.aliyun.com).
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The topic to which you want to send a message on the ApsaraMQ for MQTT instance.
	MqttTopic *string `json:"MqttTopic,omitempty" xml:"MqttTopic,omitempty"`
	// The message content, which is the payload of the message. We recommend that you encode the content in Base64 to prevent non-printable characters from being transmitted.
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
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
	// The unique message ID that is returned by the ApsaraMQ for MQTT broker after the message is sent.
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// The unique ID that the system generates for the request. This parameter is a common parameter.
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The client ID of the device whose access credential you want to deregister.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance. The ID must be consistent with the ID of the instance that the ApsaraMQ for MQTT client uses. You can obtain the instance ID on the **Instance Details** page that corresponds to the instance in the ApsaraMQ for MQTT console.
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
	// The unique ID that the system generates for the request. This parameter is a common parameter.
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnRegisterDeviceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

/**
 * *   You can call this operation up to 100 times per second per account. If you want to increase the limit, join the DingTalk group 35228338 to contact ApsaraMQ for MQTT technical support.
 * *   Each successful call to the **ApplyToken** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request ApplyTokenRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ApplyTokenResponse
 */
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

/**
 * *   You can call this operation up to 100 times per second per account. If you want to increase the limit, join the DingTalk group 35228338 to contact ApsaraMQ for MQTT technical support.
 * *   Each successful call to the **ApplyToken** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request ApplyTokenRequest
 * @return ApplyTokenResponse
 */
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

/**
 * *   You can call the **BatchQuerySessionByClientIds** operation up to 100 times per second. For more information, see [Limits on QPS](~~163047~~).
 * *   You can call the **BatchQuerySessionByClientIds** operation to query the status of up to 10 ApsaraMQ for MQTT clients in a single query.
 * *   Each successful call to the **BatchQuerySessionByClientIds** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request BatchQuerySessionByClientIdsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return BatchQuerySessionByClientIdsResponse
 */
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

/**
 * *   You can call the **BatchQuerySessionByClientIds** operation up to 100 times per second. For more information, see [Limits on QPS](~~163047~~).
 * *   You can call the **BatchQuerySessionByClientIds** operation to query the status of up to 10 ApsaraMQ for MQTT clients in a single query.
 * *   Each successful call to the **BatchQuerySessionByClientIds** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request BatchQuerySessionByClientIdsRequest
 * @return BatchQuerySessionByClientIdsResponse
 */
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

/**
 * Each successful call to the **CreateGroupId** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request CreateGroupIdRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateGroupIdResponse
 */
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

/**
 * Each successful call to the **CreateGroupId** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request CreateGroupIdRequest
 * @return CreateGroupIdResponse
 */
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

/**
 * Each successful call to the **DeleteGroupId** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request DeleteGroupIdRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteGroupIdResponse
 */
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

/**
 * Each successful call to the **DeleteGroupId** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request DeleteGroupIdRequest
 * @return DeleteGroupIdResponse
 */
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

/**
 * *   You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](~~163047~~).
 * *   Each successful call to the **GetDeviceCredential** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request GetDeviceCredentialRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetDeviceCredentialResponse
 */
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

/**
 * *   You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](~~163047~~).
 * *   Each successful call to the **GetDeviceCredential** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request GetDeviceCredentialRequest
 * @return GetDeviceCredentialResponse
 */
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

func (client *Client) ListDeviceCredentialClientIdWithOptions(request *ListDeviceCredentialClientIdRequest, runtime *util.RuntimeOptions) (_result *ListDeviceCredentialClientIdResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeviceCredentialClientId"),
		Version:     tea.String("2020-04-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeviceCredentialClientIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeviceCredentialClientId(request *ListDeviceCredentialClientIdRequest) (_result *ListDeviceCredentialClientIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeviceCredentialClientIdResponse{}
	_body, _err := client.ListDeviceCredentialClientIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Each successful call to the **ListGroupId** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request ListGroupIdRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListGroupIdResponse
 */
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

/**
 * Each successful call to the **ListGroupId** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request ListGroupIdRequest
 * @return ListGroupIdResponse
 */
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

/**
 * *   Each successful call to the **QueryMqttTraceDevice** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 * *   You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](~~163047~~).
 *
 * @param request QueryMqttTraceDeviceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryMqttTraceDeviceResponse
 */
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

/**
 * *   Each successful call to the **QueryMqttTraceDevice** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 * *   You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](~~163047~~).
 *
 * @param request QueryMqttTraceDeviceRequest
 * @return QueryMqttTraceDeviceResponse
 */
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

/**
 * *   Each successful call to the **QueryMqttTraceMessageOfClient** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 * *   You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](~~163047~~).
 *
 * @param request QueryMqttTraceMessageOfClientRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryMqttTraceMessageOfClientResponse
 */
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

/**
 * *   Each successful call to the **QueryMqttTraceMessageOfClient** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 * *   You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](~~163047~~).
 *
 * @param request QueryMqttTraceMessageOfClientRequest
 * @return QueryMqttTraceMessageOfClientResponse
 */
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

/**
 * *   Each successful call to the **QueryMqttTraceMessagePublish** operation increases the messaging transactions per second (TPS). This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 * *   You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](~~163047~~).
 *
 * @param request QueryMqttTraceMessagePublishRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryMqttTraceMessagePublishResponse
 */
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

/**
 * *   Each successful call to the **QueryMqttTraceMessagePublish** operation increases the messaging transactions per second (TPS). This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 * *   You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](~~163047~~).
 *
 * @param request QueryMqttTraceMessagePublishRequest
 * @return QueryMqttTraceMessagePublishResponse
 */
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

/**
 * *   Each successful call to the **QueryMqttTraceMessageSubscribe** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 * *   You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](~~163047~~).
 *
 * @param request QueryMqttTraceMessageSubscribeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryMqttTraceMessageSubscribeResponse
 */
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

/**
 * *   Each successful call to the **QueryMqttTraceMessageSubscribe** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 * *   You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](~~163047~~).
 *
 * @param request QueryMqttTraceMessageSubscribeRequest
 * @return QueryMqttTraceMessageSubscribeResponse
 */
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

/**
 * *   You can call this operation up to 500 times per second.**** For more information, see [Limits on QPS](~~163047~~).
 * *   Each successful call to the **QuerySessionByClientId** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request QuerySessionByClientIdRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QuerySessionByClientIdResponse
 */
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

/**
 * *   You can call this operation up to 500 times per second.**** For more information, see [Limits on QPS](~~163047~~).
 * *   Each successful call to the **QuerySessionByClientId** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request QuerySessionByClientIdRequest
 * @return QuerySessionByClientIdResponse
 */
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

/**
 * *   You can call this operation up to 100 times per second per account. If you want to increase the limit, join the DingTalk group 35228338 to contact ApsaraMQ for MQTT technical support.
 * *   Each successful call to the **QueryToken** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request QueryTokenRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryTokenResponse
 */
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

/**
 * *   You can call this operation up to 100 times per second per account. If you want to increase the limit, join the DingTalk group 35228338 to contact ApsaraMQ for MQTT technical support.
 * *   Each successful call to the **QueryToken** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request QueryTokenRequest
 * @return QueryTokenResponse
 */
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

/**
 * ## [](#)Limits
 * You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](~~163047~~).
 * >  Each successful call to the **RefreshDeviceCredential** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request RefreshDeviceCredentialRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RefreshDeviceCredentialResponse
 */
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

/**
 * ## [](#)Limits
 * You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](~~163047~~).
 * >  Each successful call to the **RefreshDeviceCredential** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request RefreshDeviceCredentialRequest
 * @return RefreshDeviceCredentialResponse
 */
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

/**
 * *   You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](~~163047~~).
 * *   Each successful call to the **RegisterDeviceCredential** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request RegisterDeviceCredentialRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RegisterDeviceCredentialResponse
 */
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

/**
 * *   You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](~~163047~~).
 * *   Each successful call to the **RegisterDeviceCredential** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request RegisterDeviceCredentialRequest
 * @return RegisterDeviceCredentialResponse
 */
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

/**
 * *   You can call this operation up to 5 times per second per account. If you want to increase the limit, join the DingTalk group 35228338 to contact ApsaraMQ for MQTT technical support.
 * *   Each successful call to the **RevokeToken** operation increases the messaging transactions per second (TPS). This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request RevokeTokenRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RevokeTokenResponse
 */
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

/**
 * *   You can call this operation up to 5 times per second per account. If you want to increase the limit, join the DingTalk group 35228338 to contact ApsaraMQ for MQTT technical support.
 * *   Each successful call to the **RevokeToken** operation increases the messaging transactions per second (TPS). This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request RevokeTokenRequest
 * @return RevokeTokenResponse
 */
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

/**
 * *   The **SendMessage** operation is called by applications on cloud servers. It is complementary to the operation that is called by ApsaraMQ for MQTT clients to send messages. For information about the differences between the scenarios of sending messages from applications on cloud servers and the scenarios of sending messages from ApsaraMQ for MQTT clients, see [Developer guide](~~179160~~).
 * *   Before you call the **SendMessage** operation, make sure that the kernel version of your ApsaraMQ for MQTT instance is 3.3.0 or later. You can obtain the information about the kernel version on the [Instance Details](https://mqtt.console.aliyun.com) page that corresponds to the instance in the **ApsaraMQ for MQTT console**.
 * *   Messages that are sent by calling the **SendMessage** operation cannot be forwarded to ApsaraMQ for RocketMQ. If you want to use an ApsaraMQ for MQTT to forward messages to ApsaraMQ for RocketMQ, send the messages by using an SDK. For more information, see [Export data from ApsaraMQ for MQTT to other Alibaba Cloud services](~~174527~~). You can call the **SendMessage** operation up to 1,000 times per second. For more information, see [Limits on QPS](~~163047~~).
 * *   Each successful call to the **SendMessage** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For information about the billing details, see [Billing rules](~~52819~~).
 *
 * @param request SendMessageRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SendMessageResponse
 */
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

/**
 * *   The **SendMessage** operation is called by applications on cloud servers. It is complementary to the operation that is called by ApsaraMQ for MQTT clients to send messages. For information about the differences between the scenarios of sending messages from applications on cloud servers and the scenarios of sending messages from ApsaraMQ for MQTT clients, see [Developer guide](~~179160~~).
 * *   Before you call the **SendMessage** operation, make sure that the kernel version of your ApsaraMQ for MQTT instance is 3.3.0 or later. You can obtain the information about the kernel version on the [Instance Details](https://mqtt.console.aliyun.com) page that corresponds to the instance in the **ApsaraMQ for MQTT console**.
 * *   Messages that are sent by calling the **SendMessage** operation cannot be forwarded to ApsaraMQ for RocketMQ. If you want to use an ApsaraMQ for MQTT to forward messages to ApsaraMQ for RocketMQ, send the messages by using an SDK. For more information, see [Export data from ApsaraMQ for MQTT to other Alibaba Cloud services](~~174527~~). You can call the **SendMessage** operation up to 1,000 times per second. For more information, see [Limits on QPS](~~163047~~).
 * *   Each successful call to the **SendMessage** operation increases the messaging transactions per second (TPS) by one. This affects the billing of your instance. For information about the billing details, see [Billing rules](~~52819~~).
 *
 * @param request SendMessageRequest
 * @return SendMessageResponse
 */
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

/**
 * *   You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](~~163047~~).
 * *   Each successful call to the **UnRegisterDeviceCredential** operation increases the number of transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request UnRegisterDeviceCredentialRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UnRegisterDeviceCredentialResponse
 */
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

/**
 * *   You can call this operation up to 500 times per second per account. If the limit is exceeded, throttling is triggered. This may affect your business. We recommend that you take note of this limit when you call this operation. For more information, see [Limits on QPS](~~163047~~).
 * *   Each successful call to the **UnRegisterDeviceCredential** operation increases the number of transactions per second (TPS) by one. This affects the billing of your instance. For more information, see [Billing rules](~~52819~~).
 *
 * @param request UnRegisterDeviceCredentialRequest
 * @return UnRegisterDeviceCredentialResponse
 */
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
