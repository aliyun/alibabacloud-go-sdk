// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateEventSourceRequest
	GetDescription() *string
	SetEventBusName(v string) *CreateEventSourceRequest
	GetEventBusName() *string
	SetEventSourceName(v string) *CreateEventSourceRequest
	GetEventSourceName() *string
	SetExternalSourceConfig(v map[string]interface{}) *CreateEventSourceRequest
	GetExternalSourceConfig() map[string]interface{}
	SetExternalSourceType(v []byte) *CreateEventSourceRequest
	GetExternalSourceType() []byte
	SetLinkedExternalSource(v bool) *CreateEventSourceRequest
	GetLinkedExternalSource() *bool
	SetSourceHttpEventParameters(v *CreateEventSourceRequestSourceHttpEventParameters) *CreateEventSourceRequest
	GetSourceHttpEventParameters() *CreateEventSourceRequestSourceHttpEventParameters
	SetSourceKafkaParameters(v *CreateEventSourceRequestSourceKafkaParameters) *CreateEventSourceRequest
	GetSourceKafkaParameters() *CreateEventSourceRequestSourceKafkaParameters
	SetSourceMNSParameters(v *CreateEventSourceRequestSourceMNSParameters) *CreateEventSourceRequest
	GetSourceMNSParameters() *CreateEventSourceRequestSourceMNSParameters
	SetSourceOSSEventParameters(v *CreateEventSourceRequestSourceOSSEventParameters) *CreateEventSourceRequest
	GetSourceOSSEventParameters() *CreateEventSourceRequestSourceOSSEventParameters
	SetSourceRabbitMQParameters(v *CreateEventSourceRequestSourceRabbitMQParameters) *CreateEventSourceRequest
	GetSourceRabbitMQParameters() *CreateEventSourceRequestSourceRabbitMQParameters
	SetSourceRocketMQParameters(v *CreateEventSourceRequestSourceRocketMQParameters) *CreateEventSourceRequest
	GetSourceRocketMQParameters() *CreateEventSourceRequestSourceRocketMQParameters
	SetSourceSLSParameters(v *CreateEventSourceRequestSourceSLSParameters) *CreateEventSourceRequest
	GetSourceSLSParameters() *CreateEventSourceRequestSourceSLSParameters
	SetSourceScheduledEventParameters(v *CreateEventSourceRequestSourceScheduledEventParameters) *CreateEventSourceRequest
	GetSourceScheduledEventParameters() *CreateEventSourceRequestSourceScheduledEventParameters
}

type CreateEventSourceRequest struct {
	// The description of the event source.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the event bus with which the event source is associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-event-bus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	// The name of the event source.
	//
	// This parameter is required.
	//
	// example:
	//
	// myrabbitmq.sourc
	EventSourceName *string `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
	// The configurations of the external data source.
	ExternalSourceConfig map[string]interface{} `json:"ExternalSourceConfig,omitempty" xml:"ExternalSourceConfig,omitempty"`
	// The type of the external data source.
	//
	// example:
	//
	// RabbitMQ
	ExternalSourceType []byte `json:"ExternalSourceType,omitempty" xml:"ExternalSourceType,omitempty"`
	// Specify whether to connect to an external data source.
	//
	// example:
	//
	// true
	LinkedExternalSource *bool `json:"LinkedExternalSource,omitempty" xml:"LinkedExternalSource,omitempty"`
	// The parameters that are configured if the event source is HTTP events.
	SourceHttpEventParameters *CreateEventSourceRequestSourceHttpEventParameters `json:"SourceHttpEventParameters,omitempty" xml:"SourceHttpEventParameters,omitempty" type:"Struct"`
	// The parameters that are configured if the event source is Message Queue for Apache Kafka.
	SourceKafkaParameters *CreateEventSourceRequestSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify Simple Message Queue (formerly MNS) (SMQ) as the event source. If you specify SMQ as the event source, you must configure RegionId, IsBase64Decode, and QueueName.
	SourceMNSParameters      *CreateEventSourceRequestSourceMNSParameters      `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	SourceOSSEventParameters *CreateEventSourceRequestSourceOSSEventParameters `json:"SourceOSSEventParameters,omitempty" xml:"SourceOSSEventParameters,omitempty" type:"Struct"`
	// The parameters that are configured if the event source is Message Queue for RabbitMQ.
	SourceRabbitMQParameters *CreateEventSourceRequestSourceRabbitMQParameters `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if the event source is Message Queue for Apache RocketMQ.
	SourceRocketMQParameters *CreateEventSourceRequestSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// The parameters that are configured if the event source is Log Service.
	SourceSLSParameters *CreateEventSourceRequestSourceSLSParameters `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
	// The parameters that are configured if you specify scheduled events as the event source.
	SourceScheduledEventParameters *CreateEventSourceRequestSourceScheduledEventParameters `json:"SourceScheduledEventParameters,omitempty" xml:"SourceScheduledEventParameters,omitempty" type:"Struct"`
}

func (s CreateEventSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEventSourceRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *CreateEventSourceRequest) GetEventSourceName() *string {
	return s.EventSourceName
}

func (s *CreateEventSourceRequest) GetExternalSourceConfig() map[string]interface{} {
	return s.ExternalSourceConfig
}

func (s *CreateEventSourceRequest) GetExternalSourceType() []byte {
	return s.ExternalSourceType
}

func (s *CreateEventSourceRequest) GetLinkedExternalSource() *bool {
	return s.LinkedExternalSource
}

func (s *CreateEventSourceRequest) GetSourceHttpEventParameters() *CreateEventSourceRequestSourceHttpEventParameters {
	return s.SourceHttpEventParameters
}

func (s *CreateEventSourceRequest) GetSourceKafkaParameters() *CreateEventSourceRequestSourceKafkaParameters {
	return s.SourceKafkaParameters
}

func (s *CreateEventSourceRequest) GetSourceMNSParameters() *CreateEventSourceRequestSourceMNSParameters {
	return s.SourceMNSParameters
}

func (s *CreateEventSourceRequest) GetSourceOSSEventParameters() *CreateEventSourceRequestSourceOSSEventParameters {
	return s.SourceOSSEventParameters
}

func (s *CreateEventSourceRequest) GetSourceRabbitMQParameters() *CreateEventSourceRequestSourceRabbitMQParameters {
	return s.SourceRabbitMQParameters
}

func (s *CreateEventSourceRequest) GetSourceRocketMQParameters() *CreateEventSourceRequestSourceRocketMQParameters {
	return s.SourceRocketMQParameters
}

func (s *CreateEventSourceRequest) GetSourceSLSParameters() *CreateEventSourceRequestSourceSLSParameters {
	return s.SourceSLSParameters
}

func (s *CreateEventSourceRequest) GetSourceScheduledEventParameters() *CreateEventSourceRequestSourceScheduledEventParameters {
	return s.SourceScheduledEventParameters
}

func (s *CreateEventSourceRequest) SetDescription(v string) *CreateEventSourceRequest {
	s.Description = &v
	return s
}

func (s *CreateEventSourceRequest) SetEventBusName(v string) *CreateEventSourceRequest {
	s.EventBusName = &v
	return s
}

func (s *CreateEventSourceRequest) SetEventSourceName(v string) *CreateEventSourceRequest {
	s.EventSourceName = &v
	return s
}

func (s *CreateEventSourceRequest) SetExternalSourceConfig(v map[string]interface{}) *CreateEventSourceRequest {
	s.ExternalSourceConfig = v
	return s
}

func (s *CreateEventSourceRequest) SetExternalSourceType(v []byte) *CreateEventSourceRequest {
	s.ExternalSourceType = v
	return s
}

func (s *CreateEventSourceRequest) SetLinkedExternalSource(v bool) *CreateEventSourceRequest {
	s.LinkedExternalSource = &v
	return s
}

func (s *CreateEventSourceRequest) SetSourceHttpEventParameters(v *CreateEventSourceRequestSourceHttpEventParameters) *CreateEventSourceRequest {
	s.SourceHttpEventParameters = v
	return s
}

func (s *CreateEventSourceRequest) SetSourceKafkaParameters(v *CreateEventSourceRequestSourceKafkaParameters) *CreateEventSourceRequest {
	s.SourceKafkaParameters = v
	return s
}

func (s *CreateEventSourceRequest) SetSourceMNSParameters(v *CreateEventSourceRequestSourceMNSParameters) *CreateEventSourceRequest {
	s.SourceMNSParameters = v
	return s
}

func (s *CreateEventSourceRequest) SetSourceOSSEventParameters(v *CreateEventSourceRequestSourceOSSEventParameters) *CreateEventSourceRequest {
	s.SourceOSSEventParameters = v
	return s
}

func (s *CreateEventSourceRequest) SetSourceRabbitMQParameters(v *CreateEventSourceRequestSourceRabbitMQParameters) *CreateEventSourceRequest {
	s.SourceRabbitMQParameters = v
	return s
}

func (s *CreateEventSourceRequest) SetSourceRocketMQParameters(v *CreateEventSourceRequestSourceRocketMQParameters) *CreateEventSourceRequest {
	s.SourceRocketMQParameters = v
	return s
}

func (s *CreateEventSourceRequest) SetSourceSLSParameters(v *CreateEventSourceRequestSourceSLSParameters) *CreateEventSourceRequest {
	s.SourceSLSParameters = v
	return s
}

func (s *CreateEventSourceRequest) SetSourceScheduledEventParameters(v *CreateEventSourceRequestSourceScheduledEventParameters) *CreateEventSourceRequest {
	s.SourceScheduledEventParameters = v
	return s
}

func (s *CreateEventSourceRequest) Validate() error {
	if s.SourceHttpEventParameters != nil {
		if err := s.SourceHttpEventParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceKafkaParameters != nil {
		if err := s.SourceKafkaParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceMNSParameters != nil {
		if err := s.SourceMNSParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceOSSEventParameters != nil {
		if err := s.SourceOSSEventParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceRabbitMQParameters != nil {
		if err := s.SourceRabbitMQParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceRocketMQParameters != nil {
		if err := s.SourceRocketMQParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceSLSParameters != nil {
		if err := s.SourceSLSParameters.Validate(); err != nil {
			return err
		}
	}
	if s.SourceScheduledEventParameters != nil {
		if err := s.SourceScheduledEventParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEventSourceRequestSourceHttpEventParameters struct {
	// The CIDR block that is used for security settings. This parameter is required only if you set SecurityConfig to ip. You can enter a CIDR block or an IP address.
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
	// The HTTP request method supported by the generated webhook URL. You can select multiple values. Valid values:
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
	// The security domain name. This parameter is required only if you set SecurityConfig to referer. You can enter a domain name.
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
}

func (s CreateEventSourceRequestSourceHttpEventParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSourceRequestSourceHttpEventParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceHttpEventParameters) GetIp() []*string {
	return s.Ip
}

func (s *CreateEventSourceRequestSourceHttpEventParameters) GetMethod() []*string {
	return s.Method
}

func (s *CreateEventSourceRequestSourceHttpEventParameters) GetReferer() []*string {
	return s.Referer
}

func (s *CreateEventSourceRequestSourceHttpEventParameters) GetSecurityConfig() *string {
	return s.SecurityConfig
}

func (s *CreateEventSourceRequestSourceHttpEventParameters) GetType() *string {
	return s.Type
}

func (s *CreateEventSourceRequestSourceHttpEventParameters) SetIp(v []*string) *CreateEventSourceRequestSourceHttpEventParameters {
	s.Ip = v
	return s
}

func (s *CreateEventSourceRequestSourceHttpEventParameters) SetMethod(v []*string) *CreateEventSourceRequestSourceHttpEventParameters {
	s.Method = v
	return s
}

func (s *CreateEventSourceRequestSourceHttpEventParameters) SetReferer(v []*string) *CreateEventSourceRequestSourceHttpEventParameters {
	s.Referer = v
	return s
}

func (s *CreateEventSourceRequestSourceHttpEventParameters) SetSecurityConfig(v string) *CreateEventSourceRequestSourceHttpEventParameters {
	s.SecurityConfig = &v
	return s
}

func (s *CreateEventSourceRequestSourceHttpEventParameters) SetType(v string) *CreateEventSourceRequestSourceHttpEventParameters {
	s.Type = &v
	return s
}

func (s *CreateEventSourceRequestSourceHttpEventParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventSourceRequestSourceKafkaParameters struct {
	// The ID of the consumer group that subscribes to the topic.
	//
	// example:
	//
	// wechat_peer_prod
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The ID of the Message Queue for Apache Kafka instance.
	//
	// example:
	//
	// pc-2zehmg67txzuyuuwlxv4f
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of consumers.
	//
	// example:
	//
	// 1
	MaximumTasks *int32 `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	// The network. Valid values: Default and PublicNetwork. Default value: Default. The value PublicNetwork indicates a self-managed network.
	//
	// example:
	//
	// Default
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The consumer offset.
	//
	// example:
	//
	// latest
	OffsetReset *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	// The ID of the region where the Message Queue for Apache Kafka instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group to which the Message Queue for Apache Kafka instance belongs. This parameter is required only if you set Network to PublicNetwork.
	//
	// example:
	//
	// sg-8vbf66aoyp0wfzrzxlmy
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the topic on the Message Queue for Apache Kafka instance.
	//
	// example:
	//
	// prod_ma_dispatch_center_call_record
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The ID of the vSwitch with which the Message Queue for Apache Kafka instance is associated. This parameter is required only if you set Network to PublicNetwork.
	//
	// example:
	//
	// vsw-bp127azpeirmwu4q9ttqi
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The ID of the VPC in which the Message Queue for Apache Kafka instance resides. This parameter is required only if you set Network to PublicNetwork.
	//
	// example:
	//
	// vpc-2ze5ejm986a73qq3vshlk
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateEventSourceRequestSourceKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSourceRequestSourceKafkaParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceKafkaParameters) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *CreateEventSourceRequestSourceKafkaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEventSourceRequestSourceKafkaParameters) GetMaximumTasks() *int32 {
	return s.MaximumTasks
}

func (s *CreateEventSourceRequestSourceKafkaParameters) GetNetwork() *string {
	return s.Network
}

func (s *CreateEventSourceRequestSourceKafkaParameters) GetOffsetReset() *string {
	return s.OffsetReset
}

func (s *CreateEventSourceRequestSourceKafkaParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEventSourceRequestSourceKafkaParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEventSourceRequestSourceKafkaParameters) GetTopic() *string {
	return s.Topic
}

func (s *CreateEventSourceRequestSourceKafkaParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *CreateEventSourceRequestSourceKafkaParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetConsumerGroup(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetInstanceId(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetMaximumTasks(v int32) *CreateEventSourceRequestSourceKafkaParameters {
	s.MaximumTasks = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetNetwork(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.Network = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetOffsetReset(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetRegionId(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetSecurityGroupId(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetTopic(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.Topic = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetVSwitchIds(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) SetVpcId(v string) *CreateEventSourceRequestSourceKafkaParameters {
	s.VpcId = &v
	return s
}

func (s *CreateEventSourceRequestSourceKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventSourceRequestSourceMNSParameters struct {
	// Specify whether to enable Base64 decoding. Valid values: true and false. If you set this parameter to true, Base64 decoding is enabled.
	//
	// example:
	//
	// true
	IsBase64Decode *bool `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	// The name of the SMQ queue.
	//
	// example:
	//
	// MyQueue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The ID of the region where the SMQ queue resides. Valid values: cn-qingdao, cn-beijing, cn-zhangjiakou, cn-huhehaote, cn-wulanchabu, cn-hangzhou, cn-shanghai, cn-shenzhen, cn-guangzhou, cn-chengdu, cn-hongkong, ap-southeast-1, ap-southeast-2, ap-southeast-3, ap-southeast-5, ap-northeast-1, eu-central-1, us-west-1, us-east-1, ap-south-1, me-east-1, and cn-north-2-gov-1.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateEventSourceRequestSourceMNSParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSourceRequestSourceMNSParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceMNSParameters) GetIsBase64Decode() *bool {
	return s.IsBase64Decode
}

func (s *CreateEventSourceRequestSourceMNSParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateEventSourceRequestSourceMNSParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEventSourceRequestSourceMNSParameters) SetIsBase64Decode(v bool) *CreateEventSourceRequestSourceMNSParameters {
	s.IsBase64Decode = &v
	return s
}

func (s *CreateEventSourceRequestSourceMNSParameters) SetQueueName(v string) *CreateEventSourceRequestSourceMNSParameters {
	s.QueueName = &v
	return s
}

func (s *CreateEventSourceRequestSourceMNSParameters) SetRegionId(v string) *CreateEventSourceRequestSourceMNSParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventSourceRequestSourceMNSParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventSourceRequestSourceOSSEventParameters struct {
	EventTypes []*string                                                       `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	MatchRules [][]*CreateEventSourceRequestSourceOSSEventParametersMatchRules `json:"MatchRules,omitempty" xml:"MatchRules,omitempty" type:"Repeated"`
	StsRoleArn *string                                                         `json:"StsRoleArn,omitempty" xml:"StsRoleArn,omitempty"`
}

func (s CreateEventSourceRequestSourceOSSEventParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSourceRequestSourceOSSEventParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceOSSEventParameters) GetEventTypes() []*string {
	return s.EventTypes
}

func (s *CreateEventSourceRequestSourceOSSEventParameters) GetMatchRules() [][]*CreateEventSourceRequestSourceOSSEventParametersMatchRules {
	return s.MatchRules
}

func (s *CreateEventSourceRequestSourceOSSEventParameters) GetStsRoleArn() *string {
	return s.StsRoleArn
}

func (s *CreateEventSourceRequestSourceOSSEventParameters) SetEventTypes(v []*string) *CreateEventSourceRequestSourceOSSEventParameters {
	s.EventTypes = v
	return s
}

func (s *CreateEventSourceRequestSourceOSSEventParameters) SetMatchRules(v [][]*CreateEventSourceRequestSourceOSSEventParametersMatchRules) *CreateEventSourceRequestSourceOSSEventParameters {
	s.MatchRules = v
	return s
}

func (s *CreateEventSourceRequestSourceOSSEventParameters) SetStsRoleArn(v string) *CreateEventSourceRequestSourceOSSEventParameters {
	s.StsRoleArn = &v
	return s
}

func (s *CreateEventSourceRequestSourceOSSEventParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventSourceRequestSourceOSSEventParametersMatchRules struct {
	Prefix     *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Suffix     *string `json:"Suffix,omitempty" xml:"Suffix,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	MatchState *bool   `json:"MatchState,omitempty" xml:"MatchState,omitempty"`
}

func (s CreateEventSourceRequestSourceOSSEventParametersMatchRules) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSourceRequestSourceOSSEventParametersMatchRules) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceOSSEventParametersMatchRules) GetPrefix() *string {
	return s.Prefix
}

func (s *CreateEventSourceRequestSourceOSSEventParametersMatchRules) GetSuffix() *string {
	return s.Suffix
}

func (s *CreateEventSourceRequestSourceOSSEventParametersMatchRules) GetName() *string {
	return s.Name
}

func (s *CreateEventSourceRequestSourceOSSEventParametersMatchRules) GetMatchState() *bool {
	return s.MatchState
}

func (s *CreateEventSourceRequestSourceOSSEventParametersMatchRules) SetPrefix(v string) *CreateEventSourceRequestSourceOSSEventParametersMatchRules {
	s.Prefix = &v
	return s
}

func (s *CreateEventSourceRequestSourceOSSEventParametersMatchRules) SetSuffix(v string) *CreateEventSourceRequestSourceOSSEventParametersMatchRules {
	s.Suffix = &v
	return s
}

func (s *CreateEventSourceRequestSourceOSSEventParametersMatchRules) SetName(v string) *CreateEventSourceRequestSourceOSSEventParametersMatchRules {
	s.Name = &v
	return s
}

func (s *CreateEventSourceRequestSourceOSSEventParametersMatchRules) SetMatchState(v bool) *CreateEventSourceRequestSourceOSSEventParametersMatchRules {
	s.MatchState = &v
	return s
}

func (s *CreateEventSourceRequestSourceOSSEventParametersMatchRules) Validate() error {
	return dara.Validate(s)
}

type CreateEventSourceRequestSourceRabbitMQParameters struct {
	// The ID of the Message Queue for RabbitMQ instance. For more information, see Limits.
	//
	// example:
	//
	// amqp-cn-nif22u74****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the queue on the Message Queue for RabbitMQ instance. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// demo
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

func (s CreateEventSourceRequestSourceRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSourceRequestSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceRabbitMQParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEventSourceRequestSourceRabbitMQParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateEventSourceRequestSourceRabbitMQParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEventSourceRequestSourceRabbitMQParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *CreateEventSourceRequestSourceRabbitMQParameters) SetInstanceId(v string) *CreateEventSourceRequestSourceRabbitMQParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventSourceRequestSourceRabbitMQParameters) SetQueueName(v string) *CreateEventSourceRequestSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *CreateEventSourceRequestSourceRabbitMQParameters) SetRegionId(v string) *CreateEventSourceRequestSourceRabbitMQParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventSourceRequestSourceRabbitMQParameters) SetVirtualHostName(v string) *CreateEventSourceRequestSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

func (s *CreateEventSourceRequestSourceRabbitMQParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventSourceRequestSourceRocketMQParameters struct {
	// The authentication type. You can set this parameter to ACL or leave this parameter empty.
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
	GroupID *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	// The endpoint that is used to access the Message Queue for Apache RocketMQ instance.
	//
	// example:
	//
	// registry-vpc****.aliyuncs.com
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The ID of the Message Queue for Apache RocketMQ instance. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// MQ_INST_164901546557****_BAAN****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// None.
	//
	// example:
	//
	// None
	InstanceNetwork *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	// The password that is used to access the Message Queue for Apache RocketMQ instance.
	//
	// example:
	//
	// ******
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// The ID of the security group to which the Message Queue for Apache RocketMQ instance belongs.
	//
	// example:
	//
	// sg-catalog-eventlistener
	InstanceSecurityGroupId *string `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	// The type of the Message Queue for Apache RocketMQ instance. Valid values:
	//
	// 	- Cloud_4: Message Queue for Apache RocketMQ 4.0 instance.
	//
	// 	- Cloud_5: Message Queue for Apache RocketMQ 5.0 instance.
	//
	// example:
	//
	// Cloud_4
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
	// vsw-bp10rbrt6rb6vrd89****
	InstanceVSwitchIds *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the Message Queue for Apache RocketMQ instance resides.
	//
	// example:
	//
	// vpc-bp1a4gmlk31hyg6ptl3ss
	InstanceVpcId *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// The offset from which message consumption starts. Valid values: CONSUME_FROM_LAST_OFFSET: Start message consumption from the latest offset. CONSUME_FROM_FIRST_OFFSET: Start message consumption from the earliest offset. CONSUME_FROM_TIMESTAMP: Start message consumption from the offset at the specified point in time. Default value: CONSUME_FROM_LAST_OFFSET.
	//
	// example:
	//
	// CONSUME_FROM_LAST_OFFSET
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The region where the Message Queue for Apache RocketMQ instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tag that is used to filter messages.
	//
	// example:
	//
	// test
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The timestamp that specifies the time from which messages are consumed. This parameter is valid only if you set Offset to CONSUME_FROM_TIMESTAMP.
	//
	// example:
	//
	// 1636597951964
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The name of the topic on the Message Queue for Apache RocketMQ instance. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// mytopic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s CreateEventSourceRequestSourceRocketMQParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSourceRequestSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) GetAuthType() *string {
	return s.AuthType
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) GetGroupID() *string {
	return s.GroupID
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) GetInstanceEndpoint() *string {
	return s.InstanceEndpoint
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) GetInstanceNetwork() *string {
	return s.InstanceNetwork
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) GetInstancePassword() *string {
	return s.InstancePassword
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) GetInstanceSecurityGroupId() *string {
	return s.InstanceSecurityGroupId
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) GetInstanceUsername() *string {
	return s.InstanceUsername
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) GetInstanceVSwitchIds() *string {
	return s.InstanceVSwitchIds
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) GetInstanceVpcId() *string {
	return s.InstanceVpcId
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) GetOffset() *string {
	return s.Offset
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) GetTag() *string {
	return s.Tag
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) GetTopic() *string {
	return s.Topic
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetAuthType(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.AuthType = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetGroupID(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceEndpoint(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceId(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceId = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceNetwork(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceNetwork = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstancePassword(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstancePassword = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceType(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceType = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceUsername(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceUsername = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceVSwitchIds(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceVSwitchIds = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetInstanceVpcId(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.InstanceVpcId = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetOffset(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.Offset = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetRegionId(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.RegionId = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetTag(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.Tag = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetTimestamp(v int64) *CreateEventSourceRequestSourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) SetTopic(v string) *CreateEventSourceRequestSourceRocketMQParameters {
	s.Topic = &v
	return s
}

func (s *CreateEventSourceRequestSourceRocketMQParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventSourceRequestSourceSLSParameters struct {
	// The starting consumer offset. The value begin specifies the earliest offset, and the value end specifies the latest offset. You can also specify a time in seconds to start consumption.
	//
	// example:
	//
	// end
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	// The Log Service Logstore.
	//
	// example:
	//
	// test-logstore
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The Log Service project.
	//
	// example:
	//
	// test-project
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The role name. If you want to authorize EventBridge to use this role to read logs in Log Service, you must select Alibaba Cloud Service for Selected Trusted Entity and EventBridge for Select Trusted Service when you create the role in the Resource Access Management (RAM) console. For information about the permission policy of this role, see Create a custom event source of the Log Service type.
	//
	// example:
	//
	// testRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateEventSourceRequestSourceSLSParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSourceRequestSourceSLSParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceSLSParameters) GetConsumePosition() *string {
	return s.ConsumePosition
}

func (s *CreateEventSourceRequestSourceSLSParameters) GetLogStore() *string {
	return s.LogStore
}

func (s *CreateEventSourceRequestSourceSLSParameters) GetProject() *string {
	return s.Project
}

func (s *CreateEventSourceRequestSourceSLSParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateEventSourceRequestSourceSLSParameters) SetConsumePosition(v string) *CreateEventSourceRequestSourceSLSParameters {
	s.ConsumePosition = &v
	return s
}

func (s *CreateEventSourceRequestSourceSLSParameters) SetLogStore(v string) *CreateEventSourceRequestSourceSLSParameters {
	s.LogStore = &v
	return s
}

func (s *CreateEventSourceRequestSourceSLSParameters) SetProject(v string) *CreateEventSourceRequestSourceSLSParameters {
	s.Project = &v
	return s
}

func (s *CreateEventSourceRequestSourceSLSParameters) SetRoleName(v string) *CreateEventSourceRequestSourceSLSParameters {
	s.RoleName = &v
	return s
}

func (s *CreateEventSourceRequestSourceSLSParameters) Validate() error {
	return dara.Validate(s)
}

type CreateEventSourceRequestSourceScheduledEventParameters struct {
	// The cron expression.
	//
	// example:
	//
	// 10 	- 	- 	- 	- *
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The time zone in which the cron expression is executed.
	//
	// example:
	//
	// GMT+0:00
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// example:
	//
	// {"a": "b"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateEventSourceRequestSourceScheduledEventParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSourceRequestSourceScheduledEventParameters) GoString() string {
	return s.String()
}

func (s *CreateEventSourceRequestSourceScheduledEventParameters) GetSchedule() *string {
	return s.Schedule
}

func (s *CreateEventSourceRequestSourceScheduledEventParameters) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateEventSourceRequestSourceScheduledEventParameters) GetUserData() *string {
	return s.UserData
}

func (s *CreateEventSourceRequestSourceScheduledEventParameters) SetSchedule(v string) *CreateEventSourceRequestSourceScheduledEventParameters {
	s.Schedule = &v
	return s
}

func (s *CreateEventSourceRequestSourceScheduledEventParameters) SetTimeZone(v string) *CreateEventSourceRequestSourceScheduledEventParameters {
	s.TimeZone = &v
	return s
}

func (s *CreateEventSourceRequestSourceScheduledEventParameters) SetUserData(v string) *CreateEventSourceRequestSourceScheduledEventParameters {
	s.UserData = &v
	return s
}

func (s *CreateEventSourceRequestSourceScheduledEventParameters) Validate() error {
	return dara.Validate(s)
}
