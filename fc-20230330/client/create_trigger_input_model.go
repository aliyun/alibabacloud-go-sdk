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
	// example:
	//
	// trigger for test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// acs:ram::1234567890:role/fc-test
	InvocationRole *string `json:"invocationRole,omitempty" xml:"invocationRole,omitempty"`
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// example:
	//
	// acs:oss:cn-shanghai:12345:mybucket
	SourceArn *string `json:"sourceArn,omitempty" xml:"sourceArn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"events":["oss:ObjectCreated:*"],"filter":{"key":{"prefix":"/prefix","suffix":".zip"}}}
	TriggerConfig *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss_create_object_demo
	TriggerName *string `json:"triggerName,omitempty" xml:"triggerName,omitempty"`
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
