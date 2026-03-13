// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrigger interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *Trigger
	GetCreatedTime() *string
	SetDescription(v string) *Trigger
	GetDescription() *string
	SetHttpTrigger(v *HTTPTrigger) *Trigger
	GetHttpTrigger() *HTTPTrigger
	SetInvocationRole(v string) *Trigger
	GetInvocationRole() *string
	SetLastModifiedTime(v string) *Trigger
	GetLastModifiedTime() *string
	SetQualifier(v string) *Trigger
	GetQualifier() *string
	SetSourceArn(v string) *Trigger
	GetSourceArn() *string
	SetStatus(v string) *Trigger
	GetStatus() *string
	SetTargetArn(v string) *Trigger
	GetTargetArn() *string
	SetTriggerConfig(v string) *Trigger
	GetTriggerConfig() *string
	SetTriggerId(v string) *Trigger
	GetTriggerId() *string
	SetTriggerName(v string) *Trigger
	GetTriggerName() *string
	SetTriggerType(v string) *Trigger
	GetTriggerType() *string
}

type Trigger struct {
	// The time when the trigger was created.
	//
	// example:
	//
	// 2020-08-20T02:28:21Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The description of the trigger.
	//
	// example:
	//
	// test_description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The information about the HTTP trigger.
	HttpTrigger *HTTPTrigger `json:"httpTrigger,omitempty" xml:"httpTrigger,omitempty"`
	// The role that is used by the event source such as OSS to invoke the function.
	//
	// example:
	//
	// acs:ram::151641468453****:role/my-role
	InvocationRole *string `json:"invocationRole,omitempty" xml:"invocationRole,omitempty"`
	// The time when the trigger was last modified.
	//
	// example:
	//
	// 2020-04-23T06:32:43Z
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the event source for the trigger.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:151641468453****:my-bucket
	SourceArn *string `json:"sourceArn,omitempty" xml:"sourceArn,omitempty"`
	// The status of the trigger.
	//
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ARN of the function.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:151641468453****:my-bucket
	TargetArn *string `json:"targetArn,omitempty" xml:"targetArn,omitempty"`
	// The configurations of the trigger. The configurations vary based on trigger types.
	//
	// example:
	//
	// {
	//
	//       "events": [
	//
	//             "oss:ObjectCreated:*"
	//
	//       ],
	//
	//       "filter": {
	//
	//             "key": {
	//
	//                   "prefix": "/prefix",
	//
	//                   "suffix": ".zip"
	//
	//             }
	//
	//       }
	//
	// }
	TriggerConfig *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
	// The unique ID of the trigger.
	//
	// example:
	//
	// 546959b5-ce1a-4991-8891-df7a02b25086
	TriggerId *string `json:"triggerId,omitempty" xml:"triggerId,omitempty"`
	// The name of the trigger. The name contains only letters, digits, hyphens (-), and underscores (_). The name must be 1 to 128 characters in length and cannot start with a digit or hyphen (-).
	//
	// example:
	//
	// defaultTrigger
	TriggerName *string `json:"triggerName,omitempty" xml:"triggerName,omitempty"`
	// The type of the trigger.
	//
	// example:
	//
	// http
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s Trigger) String() string {
	return dara.Prettify(s)
}

func (s Trigger) GoString() string {
	return s.String()
}

func (s *Trigger) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Trigger) GetDescription() *string {
	return s.Description
}

func (s *Trigger) GetHttpTrigger() *HTTPTrigger {
	return s.HttpTrigger
}

func (s *Trigger) GetInvocationRole() *string {
	return s.InvocationRole
}

func (s *Trigger) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *Trigger) GetQualifier() *string {
	return s.Qualifier
}

func (s *Trigger) GetSourceArn() *string {
	return s.SourceArn
}

func (s *Trigger) GetStatus() *string {
	return s.Status
}

func (s *Trigger) GetTargetArn() *string {
	return s.TargetArn
}

func (s *Trigger) GetTriggerConfig() *string {
	return s.TriggerConfig
}

func (s *Trigger) GetTriggerId() *string {
	return s.TriggerId
}

func (s *Trigger) GetTriggerName() *string {
	return s.TriggerName
}

func (s *Trigger) GetTriggerType() *string {
	return s.TriggerType
}

func (s *Trigger) SetCreatedTime(v string) *Trigger {
	s.CreatedTime = &v
	return s
}

func (s *Trigger) SetDescription(v string) *Trigger {
	s.Description = &v
	return s
}

func (s *Trigger) SetHttpTrigger(v *HTTPTrigger) *Trigger {
	s.HttpTrigger = v
	return s
}

func (s *Trigger) SetInvocationRole(v string) *Trigger {
	s.InvocationRole = &v
	return s
}

func (s *Trigger) SetLastModifiedTime(v string) *Trigger {
	s.LastModifiedTime = &v
	return s
}

func (s *Trigger) SetQualifier(v string) *Trigger {
	s.Qualifier = &v
	return s
}

func (s *Trigger) SetSourceArn(v string) *Trigger {
	s.SourceArn = &v
	return s
}

func (s *Trigger) SetStatus(v string) *Trigger {
	s.Status = &v
	return s
}

func (s *Trigger) SetTargetArn(v string) *Trigger {
	s.TargetArn = &v
	return s
}

func (s *Trigger) SetTriggerConfig(v string) *Trigger {
	s.TriggerConfig = &v
	return s
}

func (s *Trigger) SetTriggerId(v string) *Trigger {
	s.TriggerId = &v
	return s
}

func (s *Trigger) SetTriggerName(v string) *Trigger {
	s.TriggerName = &v
	return s
}

func (s *Trigger) SetTriggerType(v string) *Trigger {
	s.TriggerType = &v
	return s
}

func (s *Trigger) Validate() error {
	if s.HttpTrigger != nil {
		if err := s.HttpTrigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}
