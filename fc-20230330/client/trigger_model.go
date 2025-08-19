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
	CreatedTime      *string      `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Description      *string      `json:"description,omitempty" xml:"description,omitempty"`
	HttpTrigger      *HTTPTrigger `json:"httpTrigger,omitempty" xml:"httpTrigger,omitempty"`
	InvocationRole   *string      `json:"invocationRole,omitempty" xml:"invocationRole,omitempty"`
	LastModifiedTime *string      `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	Qualifier        *string      `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	SourceArn        *string      `json:"sourceArn,omitempty" xml:"sourceArn,omitempty"`
	Status           *string      `json:"status,omitempty" xml:"status,omitempty"`
	TargetArn        *string      `json:"targetArn,omitempty" xml:"targetArn,omitempty"`
	TriggerConfig    *string      `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty"`
	TriggerId        *string      `json:"triggerId,omitempty" xml:"triggerId,omitempty"`
	TriggerName      *string      `json:"triggerName,omitempty" xml:"triggerName,omitempty"`
	TriggerType      *string      `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
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
	return dara.Validate(s)
}
