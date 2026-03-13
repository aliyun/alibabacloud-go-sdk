// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTriggerInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateTriggerInput
	GetDescription() *string
	SetInvocationRole(v string) *CreateTriggerInput
	GetInvocationRole() *string
	SetQualifier(v string) *CreateTriggerInput
	GetQualifier() *string
	SetSourceArn(v string) *CreateTriggerInput
	GetSourceArn() *string
	SetTriggerConfig(v string) *CreateTriggerInput
	GetTriggerConfig() *string
	SetTriggerName(v string) *CreateTriggerInput
	GetTriggerName() *string
	SetTriggerType(v string) *CreateTriggerInput
	GetTriggerType() *string
}

type CreateTriggerInput struct {
	// The description of the trigger.
	//
	// example:
	//
	// trigger for test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The RAM role that is used by the event source such as Object Storage Service (OSS) to invoke the function.
	//
	// example:
	//
	// acs:ram::1234567890:role/fc-test
	InvocationRole *string `json:"invocationRole,omitempty" xml:"invocationRole,omitempty"`
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the trigger event source.
	//
	// example:
	//
	// acs:oss:cn-shanghai:12345:mybucket
	SourceArn *string `json:"sourceArn,omitempty" xml:"sourceArn,omitempty"`
	// The configurations of the trigger. The configurations vary based on the trigger type. The following items list the data structures for different types of triggers:
	//
	// 	- OSS triggers: [OSSTriggerConfig](https://help.aliyun.com/document_detail/2766465.html).
	//
	// 	- Simple Log Service trigger: [LogTriggerConfig](https://help.aliyun.com/document_detail/2618711.html).
	//
	// 	- Time triggers: [TimerTriggerConfig](https://help.aliyun.com/document_detail/2754638.html).
	//
	// 	- HTTP triggers: [HTTPTriggerConfig](https://help.aliyun.com/document_detail/2766461.html).
	//
	// 	- Tablestore triggers: Specify the **SourceArnm*	- parameter and leave this parameter empty.
	//
	// 	- Alibaba Cloud CDN event triggers: [CDNEventsTriggerConfig](https://help.aliyun.com/document_detail/2766462.html).
	//
	// 	- MNS topic triggers: [MnsTopicTriggerConfig](https://help.aliyun.com/document_detail/2766464.html).
	//
	// 	- EventBridge-based triggers: [EventBridgeTriggerConfig](https://help.aliyun.com/document_detail/2766447.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"events":["oss:ObjectCreated:*"],"filter":{"key":{"prefix":"/prefix","suffix":".zip"}}}
	TriggerConfig *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
	// The name of the trigger. The name can contain only letters, digits, hyphens (-), and underscores (_). The name must be 1 to 128 characters in length and cannot start with a digit or a hyphen (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// oss_create_object_demo
	TriggerName *string `json:"triggerName,omitempty" xml:"triggerName,omitempty"`
	// The type of the trigger. Valid values:
	//
	// 	- **oss**: OSS event triggers. For more information, see [Overview](https://help.aliyun.com/document_detail/2513613.html).
	//
	// 	- **log**: Simple Log Service triggers. For more information, see [Simple Log Service triggers](https://help.aliyun.com/document_detail/2513638.html).
	//
	// 	- **timer**: time triggers. For more information, see [Time triggers](https://help.aliyun.com/document_detail/2513611.html).
	//
	// 	- **http**: HTTP triggers. For more information, see [Overview](https://help.aliyun.com/document_detail/2513634.html).
	//
	// 	- **tablestore**: Tablestore triggers. For more information, see [Tablestore triggers](https://help.aliyun.com/document_detail/2513640.html).
	//
	// 	- **cdn_events**: CDN event triggers. For more information, see [Overview](https://help.aliyun.com/document_detail/2513636.html).
	//
	// 	- **mns_topic**: Message Service (MNS) topic triggers. For more information, see [MNS topic triggers](https://help.aliyun.com/document_detail/2513641.html).
	//
	// 	- **eventbridge**: EventBridge-based triggers.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s CreateTriggerInput) String() string {
	return dara.Prettify(s)
}

func (s CreateTriggerInput) GoString() string {
	return s.String()
}

func (s *CreateTriggerInput) GetDescription() *string {
	return s.Description
}

func (s *CreateTriggerInput) GetInvocationRole() *string {
	return s.InvocationRole
}

func (s *CreateTriggerInput) GetQualifier() *string {
	return s.Qualifier
}

func (s *CreateTriggerInput) GetSourceArn() *string {
	return s.SourceArn
}

func (s *CreateTriggerInput) GetTriggerConfig() *string {
	return s.TriggerConfig
}

func (s *CreateTriggerInput) GetTriggerName() *string {
	return s.TriggerName
}

func (s *CreateTriggerInput) GetTriggerType() *string {
	return s.TriggerType
}

func (s *CreateTriggerInput) SetDescription(v string) *CreateTriggerInput {
	s.Description = &v
	return s
}

func (s *CreateTriggerInput) SetInvocationRole(v string) *CreateTriggerInput {
	s.InvocationRole = &v
	return s
}

func (s *CreateTriggerInput) SetQualifier(v string) *CreateTriggerInput {
	s.Qualifier = &v
	return s
}

func (s *CreateTriggerInput) SetSourceArn(v string) *CreateTriggerInput {
	s.SourceArn = &v
	return s
}

func (s *CreateTriggerInput) SetTriggerConfig(v string) *CreateTriggerInput {
	s.TriggerConfig = &v
	return s
}

func (s *CreateTriggerInput) SetTriggerName(v string) *CreateTriggerInput {
	s.TriggerName = &v
	return s
}

func (s *CreateTriggerInput) SetTriggerType(v string) *CreateTriggerInput {
	s.TriggerType = &v
	return s
}

func (s *CreateTriggerInput) Validate() error {
	return dara.Validate(s)
}
