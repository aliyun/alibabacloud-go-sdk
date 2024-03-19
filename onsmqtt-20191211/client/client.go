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

type BatchSendMessagesRequest struct {
	InstanceId *string                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Messages   []*BatchSendMessagesRequestMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
}

func (s BatchSendMessagesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSendMessagesRequest) GoString() string {
	return s.String()
}

func (s *BatchSendMessagesRequest) SetInstanceId(v string) *BatchSendMessagesRequest {
	s.InstanceId = &v
	return s
}

func (s *BatchSendMessagesRequest) SetMessages(v []*BatchSendMessagesRequestMessages) *BatchSendMessagesRequest {
	s.Messages = v
	return s
}

type BatchSendMessagesRequestMessages struct {
	Id        *int32    `json:"Id,omitempty" xml:"Id,omitempty"`
	Payload   *string   `json:"Payload,omitempty" xml:"Payload,omitempty"`
	ReceiptId *string   `json:"ReceiptId,omitempty" xml:"ReceiptId,omitempty"`
	Topics    []*string `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
}

func (s BatchSendMessagesRequestMessages) String() string {
	return tea.Prettify(s)
}

func (s BatchSendMessagesRequestMessages) GoString() string {
	return s.String()
}

func (s *BatchSendMessagesRequestMessages) SetId(v int32) *BatchSendMessagesRequestMessages {
	s.Id = &v
	return s
}

func (s *BatchSendMessagesRequestMessages) SetPayload(v string) *BatchSendMessagesRequestMessages {
	s.Payload = &v
	return s
}

func (s *BatchSendMessagesRequestMessages) SetReceiptId(v string) *BatchSendMessagesRequestMessages {
	s.ReceiptId = &v
	return s
}

func (s *BatchSendMessagesRequestMessages) SetTopics(v []*string) *BatchSendMessagesRequestMessages {
	s.Topics = v
	return s
}

type BatchSendMessagesResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Responses []*BatchSendMessagesResponseBodyResponses `json:"Responses,omitempty" xml:"Responses,omitempty" type:"Repeated"`
}

func (s BatchSendMessagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSendMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSendMessagesResponseBody) SetRequestId(v string) *BatchSendMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSendMessagesResponseBody) SetResponses(v []*BatchSendMessagesResponseBodyResponses) *BatchSendMessagesResponseBody {
	s.Responses = v
	return s
}

type BatchSendMessagesResponseBodyResponses struct {
	ErrorCode    *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Id           *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	MsgId        *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
}

func (s BatchSendMessagesResponseBodyResponses) String() string {
	return tea.Prettify(s)
}

func (s BatchSendMessagesResponseBodyResponses) GoString() string {
	return s.String()
}

func (s *BatchSendMessagesResponseBodyResponses) SetErrorCode(v int32) *BatchSendMessagesResponseBodyResponses {
	s.ErrorCode = &v
	return s
}

func (s *BatchSendMessagesResponseBodyResponses) SetErrorMessage(v string) *BatchSendMessagesResponseBodyResponses {
	s.ErrorMessage = &v
	return s
}

func (s *BatchSendMessagesResponseBodyResponses) SetId(v int32) *BatchSendMessagesResponseBodyResponses {
	s.Id = &v
	return s
}

func (s *BatchSendMessagesResponseBodyResponses) SetMsgId(v string) *BatchSendMessagesResponseBodyResponses {
	s.MsgId = &v
	return s
}

type BatchSendMessagesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSendMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSendMessagesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSendMessagesResponse) GoString() string {
	return s.String()
}

func (s *BatchSendMessagesResponse) SetHeaders(v map[string]*string) *BatchSendMessagesResponse {
	s.Headers = v
	return s
}

func (s *BatchSendMessagesResponse) SetStatusCode(v int32) *BatchSendMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSendMessagesResponse) SetBody(v *BatchSendMessagesResponseBody) *BatchSendMessagesResponse {
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MqttTopic  *string `json:"MqttTopic,omitempty" xml:"MqttTopic,omitempty"`
	Payload    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	ReceiptId  *string `json:"ReceiptId,omitempty" xml:"ReceiptId,omitempty"`
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

func (s *SendMessageRequest) SetReceiptId(v string) *SendMessageRequest {
	s.ReceiptId = &v
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
		Version:     tea.String("2019-12-11"),
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
		Version:     tea.String("2019-12-11"),
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

func (client *Client) BatchSendMessagesWithOptions(request *BatchSendMessagesRequest, runtime *util.RuntimeOptions) (_result *BatchSendMessagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Messages)) {
		query["Messages"] = request.Messages
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchSendMessages"),
		Version:     tea.String("2019-12-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchSendMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchSendMessages(request *BatchSendMessagesRequest) (_result *BatchSendMessagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchSendMessagesResponse{}
	_body, _err := client.BatchSendMessagesWithOptions(request, runtime)
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
		Version:     tea.String("2019-12-11"),
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
		Version:     tea.String("2019-12-11"),
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
		Version:     tea.String("2019-12-11"),
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
		Version:     tea.String("2019-12-11"),
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
		Version:     tea.String("2019-12-11"),
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
		Version:     tea.String("2019-12-11"),
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

	if !tea.BoolValue(util.IsUnset(request.ReceiptId)) {
		query["ReceiptId"] = request.ReceiptId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendMessage"),
		Version:     tea.String("2019-12-11"),
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
