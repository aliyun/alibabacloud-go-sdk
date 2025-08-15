// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerGroupSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetConsumerGroupSubscriptionResponseBody
	GetCode() *string
	SetData(v []*GetConsumerGroupSubscriptionResponseBodyData) *GetConsumerGroupSubscriptionResponseBody
	GetData() []*GetConsumerGroupSubscriptionResponseBodyData
	SetDynamicCode(v string) *GetConsumerGroupSubscriptionResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetConsumerGroupSubscriptionResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GetConsumerGroupSubscriptionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetConsumerGroupSubscriptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetConsumerGroupSubscriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetConsumerGroupSubscriptionResponseBody
	GetSuccess() *bool
}

type GetConsumerGroupSubscriptionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Instance.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data []*GetConsumerGroupSubscriptionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The response code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 157DF7D4-53FB-58C6-BEBC-A9400E7EF68A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetConsumerGroupSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupSubscriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetConsumerGroupSubscriptionResponseBody) GetData() []*GetConsumerGroupSubscriptionResponseBodyData {
	return s.Data
}

func (s *GetConsumerGroupSubscriptionResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetConsumerGroupSubscriptionResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetConsumerGroupSubscriptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetConsumerGroupSubscriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConsumerGroupSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConsumerGroupSubscriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetConsumerGroupSubscriptionResponseBody) SetCode(v string) *GetConsumerGroupSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBody) SetData(v []*GetConsumerGroupSubscriptionResponseBodyData) *GetConsumerGroupSubscriptionResponseBody {
	s.Data = v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBody) SetDynamicCode(v string) *GetConsumerGroupSubscriptionResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBody) SetDynamicMessage(v string) *GetConsumerGroupSubscriptionResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBody) SetHttpStatusCode(v int32) *GetConsumerGroupSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBody) SetMessage(v string) *GetConsumerGroupSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBody) SetRequestId(v string) *GetConsumerGroupSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBody) SetSuccess(v bool) *GetConsumerGroupSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetConsumerGroupSubscriptionResponseBodyData struct {
	// The connection details.
	ConnectionDTO *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO `json:"connectionDTO,omitempty" xml:"connectionDTO,omitempty" type:"Struct"`
	// The subscription details.
	SubscriptionDTO *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO `json:"subscriptionDTO,omitempty" xml:"subscriptionDTO,omitempty" type:"Struct"`
}

func (s GetConsumerGroupSubscriptionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupSubscriptionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupSubscriptionResponseBodyData) GetConnectionDTO() *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO {
	return s.ConnectionDTO
}

func (s *GetConsumerGroupSubscriptionResponseBodyData) GetSubscriptionDTO() *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO {
	return s.SubscriptionDTO
}

func (s *GetConsumerGroupSubscriptionResponseBodyData) SetConnectionDTO(v *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) *GetConsumerGroupSubscriptionResponseBodyData {
	s.ConnectionDTO = v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyData) SetSubscriptionDTO(v *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) *GetConsumerGroupSubscriptionResponseBodyData {
	s.SubscriptionDTO = v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO struct {
	// The client ID.
	//
	// example:
	//
	// 192.168.50.191@19908#-2093249153#1534215565#40385215750900
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The public IP address of the host.
	//
	// example:
	//
	// xx.xx.xx.xx
	EgressIp *string `json:"egressIp,omitempty" xml:"egressIp,omitempty"`
	// The host name.
	//
	// example:
	//
	// nginx
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// The language used by the client.
	//
	// example:
	//
	// zh
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// The consumption mode of the consumer group. Valid values:
	//
	// 	- BROADCASTING: broadcasting consumption
	//
	// 	- CLUSTERING: clustering consumption
	//
	// example:
	//
	// BROADCASTING
	MessageModel *string `json:"messageModel,omitempty" xml:"messageModel,omitempty"`
	// The client version.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) GetClientId() *string {
	return s.ClientId
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) GetEgressIp() *string {
	return s.EgressIp
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) GetHostname() *string {
	return s.Hostname
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) GetLanguage() *string {
	return s.Language
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) GetMessageModel() *string {
	return s.MessageModel
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) GetVersion() *string {
	return s.Version
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) SetClientId(v string) *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO {
	s.ClientId = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) SetEgressIp(v string) *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO {
	s.EgressIp = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) SetHostname(v string) *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO {
	s.Hostname = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) SetLanguage(v string) *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO {
	s.Language = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) SetMessageModel(v string) *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO {
	s.MessageModel = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) SetVersion(v string) *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO {
	s.Version = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) Validate() error {
	return dara.Validate(s)
}

type GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO struct {
	// The consumer group ID.
	//
	// example:
	//
	// GID_inspector_group
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The filter expression.
	//
	// example:
	//
	// *
	FilterExpression *string `json:"filterExpression,omitempty" xml:"filterExpression,omitempty"`
	// The type of the filter expression. Valid values:
	//
	// 	- SQL: filters messages by using SQL expressions.
	//
	// 	- TAG: filters messages by using tags.
	//
	// example:
	//
	// UNSPECIFIED
	FilterExpressionType *string `json:"filterExpressionType,omitempty" xml:"filterExpressionType,omitempty"`
	// The consumption mode of the consumer group. Valid values:
	//
	// 	- BROADCASTING: broadcasting consumption
	//
	// 	- CLUSTERING: clustering consumption
	//
	// example:
	//
	// BROADCASTING
	MessageModel *string `json:"messageModel,omitempty" xml:"messageModel,omitempty"`
	// The subscription status. Valid values:
	//
	// 	- ONLINE: The consumer group is online. If the consumer group contains multiple consumers, this value is returned if at least one of the consumers is online.
	//
	// 	- OFFLINE: The consumer group is offline. If the consumer group contains multiple consumers, this value is returned only if all consumers are offline.
	//
	// example:
	//
	// ONLINE
	SubscriptionStatus *string `json:"subscriptionStatus,omitempty" xml:"subscriptionStatus,omitempty"`
	// The topic to which the consumer group subscribes.
	//
	// example:
	//
	// Topic_normal_inspector
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) GetFilterExpression() *string {
	return s.FilterExpression
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) GetFilterExpressionType() *string {
	return s.FilterExpressionType
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) GetMessageModel() *string {
	return s.MessageModel
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) GetSubscriptionStatus() *string {
	return s.SubscriptionStatus
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) GetTopicName() *string {
	return s.TopicName
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) SetConsumerGroupId(v string) *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO {
	s.ConsumerGroupId = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) SetFilterExpression(v string) *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO {
	s.FilterExpression = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) SetFilterExpressionType(v string) *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO {
	s.FilterExpressionType = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) SetMessageModel(v string) *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO {
	s.MessageModel = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) SetSubscriptionStatus(v string) *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO {
	s.SubscriptionStatus = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) SetTopicName(v string) *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO {
	s.TopicName = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) Validate() error {
	return dara.Validate(s)
}
