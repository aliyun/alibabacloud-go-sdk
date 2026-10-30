// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateEventSourceRequest
	GetDescription() *string
	SetEventBusName(v string) *UpdateEventSourceRequest
	GetEventBusName() *string
	SetEventSourceName(v string) *UpdateEventSourceRequest
	GetEventSourceName() *string
	SetExternalSourceConfig(v map[string]interface{}) *UpdateEventSourceRequest
	GetExternalSourceConfig() map[string]interface{}
	SetExternalSourceType(v string) *UpdateEventSourceRequest
	GetExternalSourceType() *string
	SetLinkedExternalSource(v bool) *UpdateEventSourceRequest
	GetLinkedExternalSource() *bool
	SetSourceHttpEventParameters(v *UpdateEventSourceRequestSourceHttpEventParameters) *UpdateEventSourceRequest
	GetSourceHttpEventParameters() *UpdateEventSourceRequestSourceHttpEventParameters
	SetSourceKafkaParameters(v *UpdateEventSourceRequestSourceKafkaParameters) *UpdateEventSourceRequest
	GetSourceKafkaParameters() *UpdateEventSourceRequestSourceKafkaParameters
	SetSourceMNSParameters(v *UpdateEventSourceRequestSourceMNSParameters) *UpdateEventSourceRequest
	GetSourceMNSParameters() *UpdateEventSourceRequestSourceMNSParameters
	SetSourceOSSEventParameters(v *UpdateEventSourceRequestSourceOSSEventParameters) *UpdateEventSourceRequest
	GetSourceOSSEventParameters() *UpdateEventSourceRequestSourceOSSEventParameters
	SetSourceRabbitMQParameters(v *UpdateEventSourceRequestSourceRabbitMQParameters) *UpdateEventSourceRequest
	GetSourceRabbitMQParameters() *UpdateEventSourceRequestSourceRabbitMQParameters
	SetSourceRocketMQParameters(v *UpdateEventSourceRequestSourceRocketMQParameters) *UpdateEventSourceRequest
	GetSourceRocketMQParameters() *UpdateEventSourceRequestSourceRocketMQParameters
	SetSourceSLSParameters(v *UpdateEventSourceRequestSourceSLSParameters) *UpdateEventSourceRequest
	GetSourceSLSParameters() *UpdateEventSourceRequestSourceSLSParameters
	SetSourceScheduledEventParameters(v *UpdateEventSourceRequestSourceScheduledEventParameters) *UpdateEventSourceRequest
	GetSourceScheduledEventParameters() *UpdateEventSourceRequestSourceScheduledEventParameters
}

type UpdateEventSourceRequest struct {
	// The description of the event source.
	//
	// example:
	//
	// Event source of the RabbitMQ type
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the associated event bus.
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
	// myrabbitmq.source
	EventSourceName *string `json:"EventSourceName,omitempty" xml:"EventSourceName,omitempty"`
	// The external data source configuration.
	//
	// example:
	//
	// {\\"ConsumePosition\\":\\"end\\",\\"LogStore\\":\\"oss_log\\",\\"Project\\":\\"slsaudit-center-5795350335281001-cn-beijing\\",\\"RoleName\\":\\"sls-beijing-tf\\"}
	ExternalSourceConfig map[string]interface{} `json:"ExternalSourceConfig,omitempty" xml:"ExternalSourceConfig,omitempty"`
	// The type of the external data source.
	//
	// example:
	//
	// SLS
	ExternalSourceType *string `json:"ExternalSourceType,omitempty" xml:"ExternalSourceType,omitempty"`
	// Specifies whether to connect to an external data source.
	//
	// example:
	//
	// true
	LinkedExternalSource *bool `json:"LinkedExternalSource,omitempty" xml:"LinkedExternalSource,omitempty"`
	// The SourceHttpEventParameters request parameter.
	SourceHttpEventParameters *UpdateEventSourceRequestSourceHttpEventParameters `json:"SourceHttpEventParameters,omitempty" xml:"SourceHttpEventParameters,omitempty" type:"Struct"`
	// The SourceKafkaParameters request parameter.
	SourceKafkaParameters *UpdateEventSourceRequestSourceKafkaParameters `json:"SourceKafkaParameters,omitempty" xml:"SourceKafkaParameters,omitempty" type:"Struct"`
	// The parameters for a Simple Message Queue (formerly MNS) event source.
	SourceMNSParameters *UpdateEventSourceRequestSourceMNSParameters `json:"SourceMNSParameters,omitempty" xml:"SourceMNSParameters,omitempty" type:"Struct"`
	// The parameters for an OSS event source.
	SourceOSSEventParameters *UpdateEventSourceRequestSourceOSSEventParameters `json:"SourceOSSEventParameters,omitempty" xml:"SourceOSSEventParameters,omitempty" type:"Struct"`
	// The SourceRabbitMQParameters request parameter.
	SourceRabbitMQParameters *UpdateEventSourceRequestSourceRabbitMQParameters `json:"SourceRabbitMQParameters,omitempty" xml:"SourceRabbitMQParameters,omitempty" type:"Struct"`
	// The SourceRocketMQParameters request parameter.
	SourceRocketMQParameters *UpdateEventSourceRequestSourceRocketMQParameters `json:"SourceRocketMQParameters,omitempty" xml:"SourceRocketMQParameters,omitempty" type:"Struct"`
	// The parameters for a Simple Log Service event source.
	SourceSLSParameters *UpdateEventSourceRequestSourceSLSParameters `json:"SourceSLSParameters,omitempty" xml:"SourceSLSParameters,omitempty" type:"Struct"`
	// The SourceScheduledEventParameters request parameter.
	SourceScheduledEventParameters *UpdateEventSourceRequestSourceScheduledEventParameters `json:"SourceScheduledEventParameters,omitempty" xml:"SourceScheduledEventParameters,omitempty" type:"Struct"`
}

func (s UpdateEventSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEventSourceRequest) GetEventBusName() *string {
	return s.EventBusName
}

func (s *UpdateEventSourceRequest) GetEventSourceName() *string {
	return s.EventSourceName
}

func (s *UpdateEventSourceRequest) GetExternalSourceConfig() map[string]interface{} {
	return s.ExternalSourceConfig
}

func (s *UpdateEventSourceRequest) GetExternalSourceType() *string {
	return s.ExternalSourceType
}

func (s *UpdateEventSourceRequest) GetLinkedExternalSource() *bool {
	return s.LinkedExternalSource
}

func (s *UpdateEventSourceRequest) GetSourceHttpEventParameters() *UpdateEventSourceRequestSourceHttpEventParameters {
	return s.SourceHttpEventParameters
}

func (s *UpdateEventSourceRequest) GetSourceKafkaParameters() *UpdateEventSourceRequestSourceKafkaParameters {
	return s.SourceKafkaParameters
}

func (s *UpdateEventSourceRequest) GetSourceMNSParameters() *UpdateEventSourceRequestSourceMNSParameters {
	return s.SourceMNSParameters
}

func (s *UpdateEventSourceRequest) GetSourceOSSEventParameters() *UpdateEventSourceRequestSourceOSSEventParameters {
	return s.SourceOSSEventParameters
}

func (s *UpdateEventSourceRequest) GetSourceRabbitMQParameters() *UpdateEventSourceRequestSourceRabbitMQParameters {
	return s.SourceRabbitMQParameters
}

func (s *UpdateEventSourceRequest) GetSourceRocketMQParameters() *UpdateEventSourceRequestSourceRocketMQParameters {
	return s.SourceRocketMQParameters
}

func (s *UpdateEventSourceRequest) GetSourceSLSParameters() *UpdateEventSourceRequestSourceSLSParameters {
	return s.SourceSLSParameters
}

func (s *UpdateEventSourceRequest) GetSourceScheduledEventParameters() *UpdateEventSourceRequestSourceScheduledEventParameters {
	return s.SourceScheduledEventParameters
}

func (s *UpdateEventSourceRequest) SetDescription(v string) *UpdateEventSourceRequest {
	s.Description = &v
	return s
}

func (s *UpdateEventSourceRequest) SetEventBusName(v string) *UpdateEventSourceRequest {
	s.EventBusName = &v
	return s
}

func (s *UpdateEventSourceRequest) SetEventSourceName(v string) *UpdateEventSourceRequest {
	s.EventSourceName = &v
	return s
}

func (s *UpdateEventSourceRequest) SetExternalSourceConfig(v map[string]interface{}) *UpdateEventSourceRequest {
	s.ExternalSourceConfig = v
	return s
}

func (s *UpdateEventSourceRequest) SetExternalSourceType(v string) *UpdateEventSourceRequest {
	s.ExternalSourceType = &v
	return s
}

func (s *UpdateEventSourceRequest) SetLinkedExternalSource(v bool) *UpdateEventSourceRequest {
	s.LinkedExternalSource = &v
	return s
}

func (s *UpdateEventSourceRequest) SetSourceHttpEventParameters(v *UpdateEventSourceRequestSourceHttpEventParameters) *UpdateEventSourceRequest {
	s.SourceHttpEventParameters = v
	return s
}

func (s *UpdateEventSourceRequest) SetSourceKafkaParameters(v *UpdateEventSourceRequestSourceKafkaParameters) *UpdateEventSourceRequest {
	s.SourceKafkaParameters = v
	return s
}

func (s *UpdateEventSourceRequest) SetSourceMNSParameters(v *UpdateEventSourceRequestSourceMNSParameters) *UpdateEventSourceRequest {
	s.SourceMNSParameters = v
	return s
}

func (s *UpdateEventSourceRequest) SetSourceOSSEventParameters(v *UpdateEventSourceRequestSourceOSSEventParameters) *UpdateEventSourceRequest {
	s.SourceOSSEventParameters = v
	return s
}

func (s *UpdateEventSourceRequest) SetSourceRabbitMQParameters(v *UpdateEventSourceRequestSourceRabbitMQParameters) *UpdateEventSourceRequest {
	s.SourceRabbitMQParameters = v
	return s
}

func (s *UpdateEventSourceRequest) SetSourceRocketMQParameters(v *UpdateEventSourceRequestSourceRocketMQParameters) *UpdateEventSourceRequest {
	s.SourceRocketMQParameters = v
	return s
}

func (s *UpdateEventSourceRequest) SetSourceSLSParameters(v *UpdateEventSourceRequestSourceSLSParameters) *UpdateEventSourceRequest {
	s.SourceSLSParameters = v
	return s
}

func (s *UpdateEventSourceRequest) SetSourceScheduledEventParameters(v *UpdateEventSourceRequestSourceScheduledEventParameters) *UpdateEventSourceRequest {
	s.SourceScheduledEventParameters = v
	return s
}

func (s *UpdateEventSourceRequest) Validate() error {
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

type UpdateEventSourceRequestSourceHttpEventParameters struct {
	// The CIDR block for security configuration. This parameter is required only when the SecurityConfig parameter is set to ip. You can enter a CIDR block or an IP address.
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
	// The HTTP request methods supported by the generated webhook. You can select multiple values. Valid values:
	//
	// - GET
	//
	// - POST
	//
	// - PUT
	//
	// - PATCH
	//
	// - DELETE
	//
	// - HEAD
	//
	// - OPTIONS
	//
	// - TRACE
	//
	// - CONNECT
	Method []*string `json:"Method,omitempty" xml:"Method,omitempty" type:"Repeated"`
	// The secure domain name for security configuration. This parameter is required only when the SecurityConfig parameter is set to referer. You can enter a domain name.
	Referer []*string `json:"Referer,omitempty" xml:"Referer,omitempty" type:"Repeated"`
	// The type of the security configuration. Valid values:
	//
	// - none: No configuration is required.
	//
	// - ip: CIDR block.
	//
	// - referer: Secure domain name.
	//
	// example:
	//
	// none
	SecurityConfig *string `json:"SecurityConfig,omitempty" xml:"SecurityConfig,omitempty"`
	// The protocol type supported by the generated webhook. Valid values:
	//
	// - HTTP
	//
	// - HTTPS
	//
	// - HTTP&HTTPS
	//
	// example:
	//
	// HTTPS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateEventSourceRequestSourceHttpEventParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventSourceRequestSourceHttpEventParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceHttpEventParameters) GetIp() []*string {
	return s.Ip
}

func (s *UpdateEventSourceRequestSourceHttpEventParameters) GetMethod() []*string {
	return s.Method
}

func (s *UpdateEventSourceRequestSourceHttpEventParameters) GetReferer() []*string {
	return s.Referer
}

func (s *UpdateEventSourceRequestSourceHttpEventParameters) GetSecurityConfig() *string {
	return s.SecurityConfig
}

func (s *UpdateEventSourceRequestSourceHttpEventParameters) GetType() *string {
	return s.Type
}

func (s *UpdateEventSourceRequestSourceHttpEventParameters) SetIp(v []*string) *UpdateEventSourceRequestSourceHttpEventParameters {
	s.Ip = v
	return s
}

func (s *UpdateEventSourceRequestSourceHttpEventParameters) SetMethod(v []*string) *UpdateEventSourceRequestSourceHttpEventParameters {
	s.Method = v
	return s
}

func (s *UpdateEventSourceRequestSourceHttpEventParameters) SetReferer(v []*string) *UpdateEventSourceRequestSourceHttpEventParameters {
	s.Referer = v
	return s
}

func (s *UpdateEventSourceRequestSourceHttpEventParameters) SetSecurityConfig(v string) *UpdateEventSourceRequestSourceHttpEventParameters {
	s.SecurityConfig = &v
	return s
}

func (s *UpdateEventSourceRequestSourceHttpEventParameters) SetType(v string) *UpdateEventSourceRequestSourceHttpEventParameters {
	s.Type = &v
	return s
}

func (s *UpdateEventSourceRequestSourceHttpEventParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventSourceRequestSourceKafkaParameters struct {
	// The group ID of the consumer that subscribes to the topic.
	//
	// example:
	//
	// dsp_online_ml_request
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cbwp-bp1o3m66wcjgbkssm3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The concurrency quota, which is the number of consumers.
	//
	// example:
	//
	// 1
	MaximumTasks *int32 `json:"MaximumTasks,omitempty" xml:"MaximumTasks,omitempty"`
	// The network configuration. Valid values: Default and PublicNetwork.
	//
	// example:
	//
	// Default
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The consumption offset.
	//
	// example:
	//
	// latest
	OffsetReset *string `json:"OffsetReset,omitempty" xml:"OffsetReset,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group. This parameter is required for self-managed networks.
	//
	// example:
	//
	// sg-5wz3mjgo9wpvdnwpwnhkj****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// billing_notify
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The ID of the vSwitch. This parameter is required for self-managed networks.
	//
	// example:
	//
	// vsw-bp1xyntcxiwplhqxj****
	VSwitchIds *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// The ID of the VPC. This parameter is required for self-managed networks.
	//
	// example:
	//
	// vpc-2zefu4vfmx6sioguj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateEventSourceRequestSourceKafkaParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventSourceRequestSourceKafkaParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) GetMaximumTasks() *int32 {
	return s.MaximumTasks
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) GetNetwork() *string {
	return s.Network
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) GetOffsetReset() *string {
	return s.OffsetReset
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) GetTopic() *string {
	return s.Topic
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetConsumerGroup(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.ConsumerGroup = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetInstanceId(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetMaximumTasks(v int32) *UpdateEventSourceRequestSourceKafkaParameters {
	s.MaximumTasks = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetNetwork(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.Network = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetOffsetReset(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.OffsetReset = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetRegionId(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetSecurityGroupId(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetTopic(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.Topic = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetVSwitchIds(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.VSwitchIds = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) SetVpcId(v string) *UpdateEventSourceRequestSourceKafkaParameters {
	s.VpcId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceKafkaParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventSourceRequestSourceMNSParameters struct {
	// Specifies whether to enable Base64 decoding. Default value: true.
	//
	// example:
	//
	// true
	IsBase64Decode *bool `json:"IsBase64Decode,omitempty" xml:"IsBase64Decode,omitempty"`
	// The name of the Simple Message Queue (formerly MNS) queue.
	//
	// example:
	//
	// queue_api_bind_1672194645178
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The region where the Simple Message Queue (formerly MNS) queue resides.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateEventSourceRequestSourceMNSParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventSourceRequestSourceMNSParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceMNSParameters) GetIsBase64Decode() *bool {
	return s.IsBase64Decode
}

func (s *UpdateEventSourceRequestSourceMNSParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *UpdateEventSourceRequestSourceMNSParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEventSourceRequestSourceMNSParameters) SetIsBase64Decode(v bool) *UpdateEventSourceRequestSourceMNSParameters {
	s.IsBase64Decode = &v
	return s
}

func (s *UpdateEventSourceRequestSourceMNSParameters) SetQueueName(v string) *UpdateEventSourceRequestSourceMNSParameters {
	s.QueueName = &v
	return s
}

func (s *UpdateEventSourceRequestSourceMNSParameters) SetRegionId(v string) *UpdateEventSourceRequestSourceMNSParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceMNSParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventSourceRequestSourceOSSEventParameters struct {
	// The list of event types.
	EventTypes []*string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty" type:"Repeated"`
	// The matching rules.
	MatchRules [][]*UpdateEventSourceRequestSourceOSSEventParametersMatchRules `json:"MatchRules,omitempty" xml:"MatchRules,omitempty" type:"Repeated"`
	// The ARN of the STS role.
	//
	// example:
	//
	// acs:eventbridge:cn-hangzhou:**
	StsRoleArn *string `json:"StsRoleArn,omitempty" xml:"StsRoleArn,omitempty"`
}

func (s UpdateEventSourceRequestSourceOSSEventParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventSourceRequestSourceOSSEventParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceOSSEventParameters) GetEventTypes() []*string {
	return s.EventTypes
}

func (s *UpdateEventSourceRequestSourceOSSEventParameters) GetMatchRules() [][]*UpdateEventSourceRequestSourceOSSEventParametersMatchRules {
	return s.MatchRules
}

func (s *UpdateEventSourceRequestSourceOSSEventParameters) GetStsRoleArn() *string {
	return s.StsRoleArn
}

func (s *UpdateEventSourceRequestSourceOSSEventParameters) SetEventTypes(v []*string) *UpdateEventSourceRequestSourceOSSEventParameters {
	s.EventTypes = v
	return s
}

func (s *UpdateEventSourceRequestSourceOSSEventParameters) SetMatchRules(v [][]*UpdateEventSourceRequestSourceOSSEventParametersMatchRules) *UpdateEventSourceRequestSourceOSSEventParameters {
	s.MatchRules = v
	return s
}

func (s *UpdateEventSourceRequestSourceOSSEventParameters) SetStsRoleArn(v string) *UpdateEventSourceRequestSourceOSSEventParameters {
	s.StsRoleArn = &v
	return s
}

func (s *UpdateEventSourceRequestSourceOSSEventParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventSourceRequestSourceOSSEventParametersMatchRules struct {
	// The suffix.
	//
	// example:
	//
	// /document
	Suffix *string `json:"Suffix,omitempty" xml:"Suffix,omitempty"`
	// The match state.
	MatchState *bool `json:"MatchState,omitempty" xml:"MatchState,omitempty"`
	// The prefix.
	//
	// example:
	//
	// fun/document/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The name.
	//
	// example:
	//
	// bucketName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateEventSourceRequestSourceOSSEventParametersMatchRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventSourceRequestSourceOSSEventParametersMatchRules) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceOSSEventParametersMatchRules) GetSuffix() *string {
	return s.Suffix
}

func (s *UpdateEventSourceRequestSourceOSSEventParametersMatchRules) GetMatchState() *bool {
	return s.MatchState
}

func (s *UpdateEventSourceRequestSourceOSSEventParametersMatchRules) GetPrefix() *string {
	return s.Prefix
}

func (s *UpdateEventSourceRequestSourceOSSEventParametersMatchRules) GetName() *string {
	return s.Name
}

func (s *UpdateEventSourceRequestSourceOSSEventParametersMatchRules) SetSuffix(v string) *UpdateEventSourceRequestSourceOSSEventParametersMatchRules {
	s.Suffix = &v
	return s
}

func (s *UpdateEventSourceRequestSourceOSSEventParametersMatchRules) SetMatchState(v bool) *UpdateEventSourceRequestSourceOSSEventParametersMatchRules {
	s.MatchState = &v
	return s
}

func (s *UpdateEventSourceRequestSourceOSSEventParametersMatchRules) SetPrefix(v string) *UpdateEventSourceRequestSourceOSSEventParametersMatchRules {
	s.Prefix = &v
	return s
}

func (s *UpdateEventSourceRequestSourceOSSEventParametersMatchRules) SetName(v string) *UpdateEventSourceRequestSourceOSSEventParametersMatchRules {
	s.Name = &v
	return s
}

func (s *UpdateEventSourceRequestSourceOSSEventParametersMatchRules) Validate() error {
	return dara.Validate(s)
}

type UpdateEventSourceRequestSourceRabbitMQParameters struct {
	// The instance ID of the ApsaraMQ for RabbitMQ instance. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// bastionhost-cn-7mz2zky****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// eb-connect: The name of the queue in the ApsaraMQ for RabbitMQ instance. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// eb-connect
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The region where the ApsaraMQ for RabbitMQ instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vhost name of the ApsaraMQ for RabbitMQ instance. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// amqp-cn-nif22u74****
	VirtualHostName *string `json:"VirtualHostName,omitempty" xml:"VirtualHostName,omitempty"`
}

func (s UpdateEventSourceRequestSourceRabbitMQParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventSourceRequestSourceRabbitMQParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceRabbitMQParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateEventSourceRequestSourceRabbitMQParameters) GetQueueName() *string {
	return s.QueueName
}

func (s *UpdateEventSourceRequestSourceRabbitMQParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEventSourceRequestSourceRabbitMQParameters) GetVirtualHostName() *string {
	return s.VirtualHostName
}

func (s *UpdateEventSourceRequestSourceRabbitMQParameters) SetInstanceId(v string) *UpdateEventSourceRequestSourceRabbitMQParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRabbitMQParameters) SetQueueName(v string) *UpdateEventSourceRequestSourceRabbitMQParameters {
	s.QueueName = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRabbitMQParameters) SetRegionId(v string) *UpdateEventSourceRequestSourceRabbitMQParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRabbitMQParameters) SetVirtualHostName(v string) *UpdateEventSourceRequestSourceRabbitMQParameters {
	s.VirtualHostName = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRabbitMQParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventSourceRequestSourceRocketMQParameters struct {
	// The authentication type. Set the value to ACL or leave it empty.
	//
	// example:
	//
	// ACL
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The group ID of the ApsaraMQ for RocketMQ instance.
	//
	// example:
	//
	// GID-test
	GroupID *string `json:"GroupID,omitempty" xml:"GroupID,omitempty"`
	// The endpoint of the instance.
	//
	// example:
	//
	// registry-vpc****.aliyuncs.com
	InstanceEndpoint *string `json:"InstanceEndpoint,omitempty" xml:"InstanceEndpoint,omitempty"`
	// The instance ID of the ApsaraMQ for RocketMQ instance. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// dbaudit-cn-i7m2nx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network type.
	//
	// example:
	//
	// PublicNetwork
	InstanceNetwork *string `json:"InstanceNetwork,omitempty" xml:"InstanceNetwork,omitempty"`
	// The password of the instance.
	//
	// example:
	//
	// yourPassword
	InstancePassword *string `json:"InstancePassword,omitempty" xml:"InstancePassword,omitempty"`
	// The ID of the security group. This parameter is required for VPCs.
	//
	// example:
	//
	// sg-catalog-eventlistener
	InstanceSecurityGroupId *string `json:"InstanceSecurityGroupId,omitempty" xml:"InstanceSecurityGroupId,omitempty"`
	// The instance type. Valid values:
	//
	// - Cloud_4: 4.0 instance.
	//
	// - Cloud_5: 5.0 instance.
	//
	// example:
	//
	// Cloud_4
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The username of the instance.
	//
	// example:
	//
	// root
	InstanceUsername *string `json:"InstanceUsername,omitempty" xml:"InstanceUsername,omitempty"`
	// The ID of the vSwitch. This parameter is required for VPCs.
	//
	// example:
	//
	// vsw-bp10rbrt6rb6vrd89****
	InstanceVSwitchIds *string `json:"InstanceVSwitchIds,omitempty" xml:"InstanceVSwitchIds,omitempty"`
	// The ID of the VPC. This parameter is required for VPCs.
	//
	// example:
	//
	// vpc-bp1a4gmlk31hyg6pt****
	InstanceVpcId *string `json:"InstanceVpcId,omitempty" xml:"InstanceVpcId,omitempty"`
	// The consumption offset of the message. Valid values:
	//
	// - CONSUME_FROM_LAST_OFFSET: Consumes messages from the latest offset.
	//
	// - CONSUME_FROM_FIRST_OFFSET: Consumes messages from the earliest offset.
	//
	// - CONSUME_FROM_TIMESTAMP: Consumes messages from the offset at a specified point in time.
	//
	// Default value: CONSUME_FROM_LAST_OFFSET.
	//
	// example:
	//
	// CONSUMEFROMLASTOFFSET
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The region where the ApsaraMQ for RocketMQ instance resides.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The filter tag of the message.
	//
	// example:
	//
	// KEY2
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The timestamp. This parameter takes effect only when the Offset parameter is set to CONSUME_FROM_TIMESTAMP.
	//
	// example:
	//
	// 1663555399032
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The topic name of the ApsaraMQ for RocketMQ instance. For more information, see [Limits](https://help.aliyun.com/document_detail/163289.html).
	//
	// example:
	//
	// topic_default_195820716552192
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s UpdateEventSourceRequestSourceRocketMQParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventSourceRequestSourceRocketMQParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) GetAuthType() *string {
	return s.AuthType
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) GetGroupID() *string {
	return s.GroupID
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) GetInstanceEndpoint() *string {
	return s.InstanceEndpoint
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) GetInstanceNetwork() *string {
	return s.InstanceNetwork
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) GetInstancePassword() *string {
	return s.InstancePassword
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) GetInstanceSecurityGroupId() *string {
	return s.InstanceSecurityGroupId
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) GetInstanceUsername() *string {
	return s.InstanceUsername
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) GetInstanceVSwitchIds() *string {
	return s.InstanceVSwitchIds
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) GetInstanceVpcId() *string {
	return s.InstanceVpcId
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) GetOffset() *string {
	return s.Offset
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) GetTag() *string {
	return s.Tag
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) GetTopic() *string {
	return s.Topic
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetAuthType(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.AuthType = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetGroupID(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.GroupID = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceEndpoint(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceEndpoint = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceId(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceNetwork(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceNetwork = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstancePassword(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstancePassword = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceSecurityGroupId(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceSecurityGroupId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceType(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceType = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceUsername(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceUsername = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceVSwitchIds(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceVSwitchIds = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetInstanceVpcId(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.InstanceVpcId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetOffset(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.Offset = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetRegionId(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.RegionId = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetTag(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.Tag = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetTimestamp(v int64) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.Timestamp = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) SetTopic(v string) *UpdateEventSourceRequestSourceRocketMQParameters {
	s.Topic = &v
	return s
}

func (s *UpdateEventSourceRequestSourceRocketMQParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventSourceRequestSourceSLSParameters struct {
	// The start consumption offset. You can set this parameter to begin to consume from the earliest offset, end to consume from the latest offset, or a specific timestamp in seconds to consume from a specified point in time.
	//
	// example:
	//
	// end
	ConsumePosition *string `json:"ConsumePosition,omitempty" xml:"ConsumePosition,omitempty"`
	// The Simple Log Service Logstore.
	//
	// example:
	//
	// waf-logstore
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The Simple Log Service project.
	//
	// example:
	//
	// VideoTestProject
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// When authorizing EventBridge to use this role to read Simple Log Service log content, the following conditions must be met: when creating the role in the Resource Access Management (RAM) console, you must select Alibaba Cloud Service, and select EventBridge as the trusted service. For the permission policy of this role, refer to the custom event source Simple Log Service.
	//
	// example:
	//
	// testRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s UpdateEventSourceRequestSourceSLSParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventSourceRequestSourceSLSParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceSLSParameters) GetConsumePosition() *string {
	return s.ConsumePosition
}

func (s *UpdateEventSourceRequestSourceSLSParameters) GetLogStore() *string {
	return s.LogStore
}

func (s *UpdateEventSourceRequestSourceSLSParameters) GetProject() *string {
	return s.Project
}

func (s *UpdateEventSourceRequestSourceSLSParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *UpdateEventSourceRequestSourceSLSParameters) SetConsumePosition(v string) *UpdateEventSourceRequestSourceSLSParameters {
	s.ConsumePosition = &v
	return s
}

func (s *UpdateEventSourceRequestSourceSLSParameters) SetLogStore(v string) *UpdateEventSourceRequestSourceSLSParameters {
	s.LogStore = &v
	return s
}

func (s *UpdateEventSourceRequestSourceSLSParameters) SetProject(v string) *UpdateEventSourceRequestSourceSLSParameters {
	s.Project = &v
	return s
}

func (s *UpdateEventSourceRequestSourceSLSParameters) SetRoleName(v string) *UpdateEventSourceRequestSourceSLSParameters {
	s.RoleName = &v
	return s
}

func (s *UpdateEventSourceRequestSourceSLSParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateEventSourceRequestSourceScheduledEventParameters struct {
	// The cron expression.
	//
	// example:
	//
	// 10 	- 	- 	- 	- *
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The time zone for the cron expression.
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

func (s UpdateEventSourceRequestSourceScheduledEventParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventSourceRequestSourceScheduledEventParameters) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceRequestSourceScheduledEventParameters) GetSchedule() *string {
	return s.Schedule
}

func (s *UpdateEventSourceRequestSourceScheduledEventParameters) GetTimeZone() *string {
	return s.TimeZone
}

func (s *UpdateEventSourceRequestSourceScheduledEventParameters) GetUserData() *string {
	return s.UserData
}

func (s *UpdateEventSourceRequestSourceScheduledEventParameters) SetSchedule(v string) *UpdateEventSourceRequestSourceScheduledEventParameters {
	s.Schedule = &v
	return s
}

func (s *UpdateEventSourceRequestSourceScheduledEventParameters) SetTimeZone(v string) *UpdateEventSourceRequestSourceScheduledEventParameters {
	s.TimeZone = &v
	return s
}

func (s *UpdateEventSourceRequestSourceScheduledEventParameters) SetUserData(v string) *UpdateEventSourceRequestSourceScheduledEventParameters {
	s.UserData = &v
	return s
}

func (s *UpdateEventSourceRequestSourceScheduledEventParameters) Validate() error {
	return dara.Validate(s)
}
