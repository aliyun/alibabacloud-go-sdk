// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDefinedEventSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUserDefinedEventSourcesResponseBody
	GetCode() *string
	SetData(v *ListUserDefinedEventSourcesResponseBodyData) *ListUserDefinedEventSourcesResponseBody
	GetData() *ListUserDefinedEventSourcesResponseBodyData
	SetMessage(v string) *ListUserDefinedEventSourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListUserDefinedEventSourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUserDefinedEventSourcesResponseBody
	GetSuccess() *bool
}

type ListUserDefinedEventSourcesResponseBody struct {
	// The returned response code. Valid values:
	//
	// 	- Success: The request is successful.
	//
	// 	- Other codes: The request failed. For more information about error codes, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ListUserDefinedEventSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// InvalidArgument
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5169654A-7059-57E3-BFD9-33C7E012EA1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. The value true indicates that the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUserDefinedEventSourcesResponseBody) GetData() *ListUserDefinedEventSourcesResponseBodyData {
	return s.Data
}

func (s *ListUserDefinedEventSourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUserDefinedEventSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserDefinedEventSourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUserDefinedEventSourcesResponseBody) SetCode(v string) *ListUserDefinedEventSourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBody) SetData(v *ListUserDefinedEventSourcesResponseBodyData) *ListUserDefinedEventSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBody) SetMessage(v string) *ListUserDefinedEventSourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBody) SetRequestId(v string) *ListUserDefinedEventSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBody) SetSuccess(v bool) *ListUserDefinedEventSourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUserDefinedEventSourcesResponseBodyData struct {
	// The event sources.
	EventSourceList []*ListUserDefinedEventSourcesResponseBodyDataEventSourceList `json:"EventSourceList,omitempty" xml:"EventSourceList,omitempty" type:"Repeated"`
	// If excess return values exist when you configure Limit, this parameter is returned.
	//
	// example:
	//
	// 100
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 18
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyData) GetEventSourceList() []*ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	return s.EventSourceList
}

func (s *ListUserDefinedEventSourcesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserDefinedEventSourcesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListUserDefinedEventSourcesResponseBodyData) SetEventSourceList(v []*ListUserDefinedEventSourcesResponseBodyDataEventSourceList) *ListUserDefinedEventSourcesResponseBodyData {
	s.EventSourceList = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyData) SetNextToken(v string) *ListUserDefinedEventSourcesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyData) SetTotal(v int32) *ListUserDefinedEventSourcesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceList struct {
	// The Alibaba Cloud Resource Name (ARN) of the queried event source.
	//
	// example:
	//
	// acs:eventbridge:cn-hangzhou:164901546557****:eventbus/my-event-bus/eventsource/myRocketMQ.source
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The timestamp that indicates when the event source was created.
	//
	// example:
	//
	// 1607071602000
	Ctime *float32 `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	// The name of the event bus.
	//
	// example:
	//
	// test-custom-bus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The type of the event source.
	//
	// example:
	//
	// RabbitMQ
	ExternalSourceType *string `json:"ExternalSourceType,omitempty" xml:"ExternalSourceType,omitempty"`
	// The name of the queried event source.
	//
	// example:
	//
	// rocketmq.source
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameters that are returned if HTTP events are specified as the event source.
	SourceHttpEventParameters *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters `json:"SourceHttpEventParameters,omitempty" xml:"SourceHttpEventParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Message Queue for Apache Kafka is specified as the event source.
	SourceKafkaParameters *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Simple Message Queue (formerly MNS) (SMQ) is specified as the event source.
	SourceMNSParameters      *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters      `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	SourceOSSEventParameters *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceOSSEventParameters `json:"SourceOSSEventParameters,omitempty" xml:"SourceOSSEventParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Message Queue for RabbitMQ is specified as the event source.
	SourceRabbitMQParameters *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Message Queue for Apache RocketMQ is specified as the event source.
	SourceRocketMQParameters *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are returned if Simple Log Service is specified as the event source.
	SourceSLSParameters *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
	// The parameters that are returned if scheduled events are specified as the event source.
	SourceScheduledEventParameters *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters `json:"SourceScheduledEventParameters,omitempty" xml:"SourceScheduledEventParameters,omitempty" type:"Struct"`
	// The status of the queried event source. The returned value Activated indicates that the event source is activated.
	//
	// example:
	//
	// Activated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the queried event source. The returned value UserDefined indicates that the event source is a custom event source.
	//
	// example:
	//
	// UserDefined
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceList) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GetArn() *string {
	return s.Arn
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GetCtime() *float32 {
	return s.Ctime
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GetEventBusName() *string {
	return s.EventBusName
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GetExternalSourceType() *string {
	return s.ExternalSourceType
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GetName() *string {
	return s.Name
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GetSourceHttpEventParameters() *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters {
	return s.SourceHttpEventParameters
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GetSourceKafkaParameters() *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	return s.SourceKafkaParameters
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GetSourceMNSParameters() *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters {
	return s.SourceMNSParameters
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GetSourceOSSEventParameters() *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceOSSEventParameters {
	return s.SourceOSSEventParameters
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GetSourceRabbitMQParameters() *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters {
	return s.SourceRabbitMQParameters
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GetSourceRocketMQParameters() *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	return s.SourceRocketMQParameters
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GetSourceSLSParameters() *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters {
	return s.SourceSLSParameters
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GetSourceScheduledEventParameters() *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters {
	return s.SourceScheduledEventParameters
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GetStatus() *string {
	return s.Status
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) GetType() *string {
	return s.Type
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetArn(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.Arn = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetCtime(v float32) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.Ctime = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetEventBusName(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.EventBusName = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetExternalSourceType(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.ExternalSourceType = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetName(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.Name = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetSourceHttpEventParameters(v *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.SourceHttpEventParameters = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetSourceKafkaParameters(v *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.SourceKafkaParameters = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetSourceMNSParameters(v *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.SourceMNSParameters = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetSourceOSSEventParameters(v *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceOSSEventParameters) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.SourceOSSEventParameters = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetSourceRabbitMQParameters(v *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.SourceRabbitMQParameters = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetSourceRocketMQParameters(v *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.SourceRocketMQParameters = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetSourceSLSParameters(v *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.SourceSLSParameters = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetSourceScheduledEventParameters(v *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.SourceScheduledEventParameters = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetStatus(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.Status = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) SetType(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceList {
	s.Type = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceList) Validate() error {
	return dara.Validate(s)
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters struct {
	// The CIDR block that is used for security settings. This parameter is required only if SecurityConfig is set to ip. You can enter a CIDR block or an IP address.
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
	// The HTTP request method that is supported by the generated webhook URL. You can select multiple values. Valid values:
	//
	// 	- GET
	//
	// 	- POST
	//
	// 	- PUT
	//
	// 	- PATCH
	//
	// 	- DELETE
	//
	// 	- HEAD
	//
	// 	- OPTIONS
	//
	// 	- TRACE
	//
	// 	- CONNECT
	Method []*string `json:"Method,omitempty" xml:"Method,omitempty" type:"Repeated"`
	// The Internet request URL.
	PublicWebHookUrl []*string `json:"PublicWebHookUrl,omitempty" xml:"PublicWebHookUrl,omitempty" type:"Repeated"`
	// The security domain name. This parameter is required only if SecurityConfig is set to referer. You can enter a domain name.
	Referer []*string `json:"Referer,omitempty" xml:"Referer,omitempty" type:"Repeated"`
	// The type of security settings. Valid values:
	//
	// 	- none: No configuration is required.
	//
	// 	- ip: CIDR block.
	//
	// 	- referer: security domain name.
	//
	// example:
	//
	// none
	SecurityConfig *string `json:"SecurityConfig,omitempty" xml:"SecurityConfig,omitempty"`
	// The protocol type that is supported by the generated webhook URL. Valid values:
	//
	// 	- HTTP
	//
	// 	- HTTPS
	//
	// 	- HTTP\\&HTTPS
	//
	// example:
	//
	// HTTPS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The internal request URL.
	VpcWebHookUrl []*string `json:"VpcWebHookUrl,omitempty" xml:"VpcWebHookUrl,omitempty" type:"Repeated"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) GetIp() []*string {
	return s.Ip
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) GetMethod() []*string {
	return s.Method
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) GetPublicWebHookUrl() []*string {
	return s.PublicWebHookUrl
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) GetReferer() []*string {
	return s.Referer
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) GetSecurityConfig() *string {
	return s.SecurityConfig
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) GetType() *string {
	return s.Type
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) GetVpcWebHookUrl() []*string {
	return s.VpcWebHookUrl
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) SetIp(v []*string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters {
	s.Ip = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) SetMethod(v []*string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters {
	s.Method = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) SetPublicWebHookUrl(v []*string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters {
	s.PublicWebHookUrl = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) SetReferer(v []*string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters {
	s.Referer = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) SetSecurityConfig(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters {
	s.SecurityConfig = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) SetType(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters {
	s.Type = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) SetVpcWebHookUrl(v []*string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters {
	s.VpcWebHookUrl = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceHttpEventParameters) Validate() error {
	return dara.Validate(s)
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters struct {
	// The ID of the consumer group that subscribes to the topic on the Message Queue for Apache Kafka instance.
	//
	// example:
	//
	// test-gid
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The ID of the Message Queue for Apache Kafka instance.
	//
	// example:
	//
	// i-2ze6kiwzkebf04s5h8ds
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of consumers.
	//
	// example:
	//
	// 2
	MaximumTasks *int32 `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	// The network type. Valid values: Default and PublicNetwork. Default value: Default. The value PublicNetwork indicates a self-managed network.
	//
	// example:
	//
	// Default
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The consumer offset.
	//
	// example:
	//
	// earliest
	OffsetReset *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	// The ID of the region where the Message Queue for Apache Kafka instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group to which the Message Queue for Apache Kafka instance belongs.
	//
	// example:
	//
	// sg-f8zatts5g97x0j***
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The topic name.
	//
	// example:
	//
	// topic_api_1674441611897
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The ID of the vSwitch with which the Message Queue for Apache Kafka instance is associated.
	//
	// example:
	//
	// vsw-bp1hcrxq3mkcik***e
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The ID of the VPC in which the Message Queue for Apache Kafka instance is deployed.
	//
	// example:
	//
	// vpc-bp1kz3ohhzgrau2***
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) GetMaximumTasks() *int32 {
	return s.MaximumTasks
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) GetNetwork() *string {
	return s.Network
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) GetOffsetReset() *string {
	return s.OffsetReset
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) GetTopic() *string {
	return s.Topic
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetConsumerGroup(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetInstanceId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetMaximumTasks(v int32) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.MaximumTasks = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetNetwork(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.Network = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetOffsetReset(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetRegionId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.RegionId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetSecurityGroupId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetTopic(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.Topic = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetVSwitchIds(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) SetVpcId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters {
	s.VpcId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters struct {
	// Indicates whether Base64 decoding is enabled. By default, Base64 decoding is enabled.
	//
	// example:
	//
	// true
	IsBase64Decode *bool `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	// The name of the SMQ queue.
	//
	// example:
	//
	// queue.openapi-sign-callback
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the SMQ queue resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters) GetIsBase64Decode() *bool {
	return s.IsBase64Decode
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters) SetIsBase64Decode(v bool) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters {
	s.IsBase64Decode = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters) SetQueueName(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters {
	s.QueueName = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters) SetRegionId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters {
	s.RegionId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceMNSParameters) Validate() error {
	return dara.Validate(s)
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceOSSEventParameters struct {
	EventTypes []*string   `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	MatchRules interface{} `json:"MatchRules,omitempty" xml:"MatchRules,omitempty"`
	StsRoleArn *string     `json:"StsRoleArn,omitempty" xml:"StsRoleArn,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceOSSEventParameters) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceOSSEventParameters) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceOSSEventParameters) GetEventTypes() []*string {
	return s.EventTypes
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceOSSEventParameters) GetMatchRules() interface{} {
	return s.MatchRules
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceOSSEventParameters) GetStsRoleArn() *string {
	return s.StsRoleArn
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceOSSEventParameters) SetEventTypes(v []*string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceOSSEventParameters {
	s.EventTypes = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceOSSEventParameters) SetMatchRules(v interface{}) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceOSSEventParameters {
	s.MatchRules = v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceOSSEventParameters) SetStsRoleArn(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceOSSEventParameters {
	s.StsRoleArn = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceOSSEventParameters) Validate() error {
	return dara.Validate(s)
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters struct {
	// The ID of the Message Queue for RabbitMQ instance. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// bastionhost-cn-0ju2x28fj07
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the queue on the Message Queue for RabbitMQ instance. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// file-upload-queue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the Message Queue for RabbitMQ instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the vhost of the Message Queue for RabbitMQ instance. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// eb-connect
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) SetInstanceId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters {
	s.InstanceId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) SetQueueName(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) SetRegionId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters {
	s.RegionId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) SetVirtualHostName(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRabbitMQParameters) Validate() error {
	return dara.Validate(s)
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters struct {
	// The authentication type. This parameter can be set to ACL or left empty.
	//
	// example:
	//
	// ACL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The ID of the consumer group on the Message Queue for Apache RocketMQ instance.
	//
	// example:
	//
	// GID-test
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The endpoint that is used to access the Message Queue for Apache RocketMQ instance.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The ID of the Message Queue for Apache RocketMQ instance. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// bastionhost-cn-7mz293s9d1p
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of network over which the Message Queue for Apache RocketMQ instance is accessed.
	//
	// example:
	//
	// PublicNetwork
	InstanceNetwork *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	// The password that is used to access the Message Queue for Apache RocketMQ instance.
	//
	// example:
	//
	// ***
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// The ID of the security group to which the Message Queue for Apache RocketMQ instance belongs.
	//
	// example:
	//
	// eb-167adad548***
	InstanceSecurityGroupId *string `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	// The instance type. Valid values: CLOUD_4, CLOUD_5, and SELF_BUILT. The value CLOUD_4 indicates that the instance is a Message Queue for Apache RocketMQ 4.0 instance. The value CLOUD_5 indicates that the instance is a Message Queue for Apache RocketMQ 5.0 instance. The value SELF_BUILT indicates that the instance is a self-managed RocketMQ instance.
	//
	// example:
	//
	// CLOUD_5
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The username that is used to access the Message Queue for Apache RocketMQ instance.
	//
	// example:
	//
	// root
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// The ID of the vSwitch with which the Message Queue for Apache RocketMQ instance is associated.
	//
	// example:
	//
	// vsw-bp1iu***
	InstanceVSwitchIds *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the Message Queue for Apache RocketMQ instance is deployed.
	//
	// example:
	//
	// vpc-***
	InstanceVpcId *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// The offset from which messages are consumed. Valid values:
	//
	// 	- CONSUME_FROM_LAST_OFFSET: Messages are consumed from the latest offset.
	//
	// 	- CONSUME_FROM_FIRST_OFFSET: Messages are consumed from the earliest offset.
	//
	// 	- CONSUME_FROM_TIMESTAMP: Messages are consumed from the offset at the specified point in time.
	//
	// Default value: CONSUME_FROM_LAST_OFFSET.
	//
	// example:
	//
	// CONSUMEFROMLASTOFFSET
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The ID of the region where the Message Queue for Apache RocketMQ instance resides.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tag that is used to filter messages.
	//
	// example:
	//
	// dataact
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The timestamp that indicates the time from which messages are consumed. This parameter is valid only if Offset is set to CONSUME_FROM_TIMESTAMP.
	//
	// example:
	//
	// 1664591760
	Timestamp *float32 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The name of the topic on the Message Queue for Apache RocketMQ instance. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// migration_instance
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GetAuthType() *string {
	return s.AuthType
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GetGroupId() *string {
	return s.GroupId
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GetInstanceEndpoint() *string {
	return s.InstanceEndpoint
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GetInstanceNetwork() *string {
	return s.InstanceNetwork
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GetInstancePassword() *string {
	return s.InstancePassword
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GetInstanceSecurityGroupId() *string {
	return s.InstanceSecurityGroupId
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GetInstanceUsername() *string {
	return s.InstanceUsername
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GetInstanceVSwitchIds() *string {
	return s.InstanceVSwitchIds
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GetInstanceVpcId() *string {
	return s.InstanceVpcId
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GetOffset() *string {
	return s.Offset
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GetTag() *string {
	return s.Tag
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GetTimestamp() *float32 {
	return s.Timestamp
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) GetTopic() *string {
	return s.Topic
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetAuthType(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.AuthType = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetGroupId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.GroupId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceEndpoint(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceNetwork(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceNetwork = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstancePassword(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstancePassword = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceType(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceType = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceUsername(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceUsername = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceVSwitchIds(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceVSwitchIds = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetInstanceVpcId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.InstanceVpcId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetOffset(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.Offset = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetRegionId(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.RegionId = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetTag(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.Tag = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetTimestamp(v float32) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) SetTopic(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters {
	s.Topic = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceRocketMQParameters) Validate() error {
	return dara.Validate(s)
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters struct {
	// The consumer offset. The value begin indicates the earliest offset, and the value end indicates the latest offset. You can also specify a time in seconds to start consumption.
	//
	// example:
	//
	// end
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	// The Simple Log Service Logstore.
	//
	// example:
	//
	// cloudfirewall-logstore
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The Simple Log Service project.
	//
	// example:
	//
	// VideoTestProject
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Simple Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the Resource Access Management (RAM) console. For information about the permission policy of this role, see Create a custom event source of the Log Service type.
	//
	// example:
	//
	// testRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) GetConsumePosition() *string {
	return s.ConsumePosition
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) GetLogStore() *string {
	return s.LogStore
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) GetProject() *string {
	return s.Project
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) SetConsumePosition(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters {
	s.ConsumePosition = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) SetLogStore(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters {
	s.LogStore = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) SetProject(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters {
	s.Project = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) SetRoleName(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters {
	s.RoleName = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceSLSParameters) Validate() error {
	return dara.Validate(s)
}

type ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters struct {
	// The cron expression.
	//
	// example:
	//
	// 0 1 	- 	- 	- *
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The time zone in which the cron expression is executed.
	//
	// example:
	//
	// GMT+0:00
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// The JSON string.
	//
	// example:
	//
	// {"a": "b"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters) String() string {
	return dara.Prettify(s)
}

func (s ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters) GoString() string {
	return s.String()
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters) GetSchedule() *string {
	return s.Schedule
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters) GetUserData() *string {
	return s.UserData
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters) SetSchedule(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters {
	s.Schedule = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters) SetTimeZone(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters {
	s.TimeZone = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters) SetUserData(v string) *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters {
	s.UserData = &v
	return s
}

func (s *ListUserDefinedEventSourcesResponseBodyDataEventSourceListSourceScheduledEventParameters) Validate() error {
	return dara.Validate(s)
}
