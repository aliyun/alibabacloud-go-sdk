// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPolicyInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAttachments(v []*Attachment) *PolicyInfo
	GetAttachments() []*Attachment
	SetClassAlias(v string) *PolicyInfo
	GetClassAlias() *string
	SetClassName(v string) *PolicyInfo
	GetClassName() *string
	SetConfig(v string) *PolicyInfo
	GetConfig() *string
	SetDirection(v string) *PolicyInfo
	GetDirection() *string
	SetExecutePriority(v string) *PolicyInfo
	GetExecutePriority() *string
	SetExecuteStage(v string) *PolicyInfo
	GetExecuteStage() *string
	SetName(v string) *PolicyInfo
	GetName() *string
	SetPolicyId(v string) *PolicyInfo
	GetPolicyId() *string
	SetType(v string) *PolicyInfo
	GetType() *string
}

type PolicyInfo struct {
	Attachments     []*Attachment `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	ClassAlias      *string       `json:"classAlias,omitempty" xml:"classAlias,omitempty"`
	ClassName       *string       `json:"className,omitempty" xml:"className,omitempty"`
	Config          *string       `json:"config,omitempty" xml:"config,omitempty"`
	Direction       *string       `json:"direction,omitempty" xml:"direction,omitempty"`
	ExecutePriority *string       `json:"executePriority,omitempty" xml:"executePriority,omitempty"`
	ExecuteStage    *string       `json:"executeStage,omitempty" xml:"executeStage,omitempty"`
	Name            *string       `json:"name,omitempty" xml:"name,omitempty"`
	PolicyId        *string       `json:"policyId,omitempty" xml:"policyId,omitempty"`
	Type            *string       `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PolicyInfo) String() string {
	return dara.Prettify(s)
}

func (s PolicyInfo) GoString() string {
	return s.String()
}

func (s *PolicyInfo) GetAttachments() []*Attachment {
	return s.Attachments
}

func (s *PolicyInfo) GetClassAlias() *string {
	return s.ClassAlias
}

func (s *PolicyInfo) GetClassName() *string {
	return s.ClassName
}

func (s *PolicyInfo) GetConfig() *string {
	return s.Config
}

func (s *PolicyInfo) GetDirection() *string {
	return s.Direction
}

func (s *PolicyInfo) GetExecutePriority() *string {
	return s.ExecutePriority
}

func (s *PolicyInfo) GetExecuteStage() *string {
	return s.ExecuteStage
}

func (s *PolicyInfo) GetName() *string {
	return s.Name
}

func (s *PolicyInfo) GetPolicyId() *string {
	return s.PolicyId
}

func (s *PolicyInfo) GetType() *string {
	return s.Type
}

func (s *PolicyInfo) SetAttachments(v []*Attachment) *PolicyInfo {
	s.Attachments = v
	return s
}

func (s *PolicyInfo) SetClassAlias(v string) *PolicyInfo {
	s.ClassAlias = &v
	return s
}

func (s *PolicyInfo) SetClassName(v string) *PolicyInfo {
	s.ClassName = &v
	return s
}

func (s *PolicyInfo) SetConfig(v string) *PolicyInfo {
	s.Config = &v
	return s
}

func (s *PolicyInfo) SetDirection(v string) *PolicyInfo {
	s.Direction = &v
	return s
}

func (s *PolicyInfo) SetExecutePriority(v string) *PolicyInfo {
	s.ExecutePriority = &v
	return s
}

func (s *PolicyInfo) SetExecuteStage(v string) *PolicyInfo {
	s.ExecuteStage = &v
	return s
}

func (s *PolicyInfo) SetName(v string) *PolicyInfo {
	s.Name = &v
	return s
}

func (s *PolicyInfo) SetPolicyId(v string) *PolicyInfo {
	s.PolicyId = &v
	return s
}

func (s *PolicyInfo) SetType(v string) *PolicyInfo {
	s.Type = &v
	return s
}

func (s *PolicyInfo) Validate() error {
	return dara.Validate(s)
}
