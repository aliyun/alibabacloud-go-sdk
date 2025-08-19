// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTriggerInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateTriggerInput
	GetDescription() *string
	SetInvocationRole(v string) *UpdateTriggerInput
	GetInvocationRole() *string
	SetQualifier(v string) *UpdateTriggerInput
	GetQualifier() *string
	SetTriggerConfig(v string) *UpdateTriggerInput
	GetTriggerConfig() *string
}

type UpdateTriggerInput struct {
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
	// {"events":["oss:ObjectCreated:*"],"filter":{"key":{"prefix":"/prefix","suffix":".zip"}}}
	TriggerConfig *string `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
}

func (s UpdateTriggerInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateTriggerInput) GoString() string {
	return s.String()
}

func (s *UpdateTriggerInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateTriggerInput) GetInvocationRole() *string {
	return s.InvocationRole
}

func (s *UpdateTriggerInput) GetQualifier() *string {
	return s.Qualifier
}

func (s *UpdateTriggerInput) GetTriggerConfig() *string {
	return s.TriggerConfig
}

func (s *UpdateTriggerInput) SetDescription(v string) *UpdateTriggerInput {
	s.Description = &v
	return s
}

func (s *UpdateTriggerInput) SetInvocationRole(v string) *UpdateTriggerInput {
	s.InvocationRole = &v
	return s
}

func (s *UpdateTriggerInput) SetQualifier(v string) *UpdateTriggerInput {
	s.Qualifier = &v
	return s
}

func (s *UpdateTriggerInput) SetTriggerConfig(v string) *UpdateTriggerInput {
	s.TriggerConfig = &v
	return s
}

func (s *UpdateTriggerInput) Validate() error {
	return dara.Validate(s)
}
