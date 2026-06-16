// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSandbox interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *Sandbox
	GetCreatedAt() *string
	SetEndedAt(v string) *Sandbox
	GetEndedAt() *string
	SetLastUpdatedAt(v string) *Sandbox
	GetLastUpdatedAt() *string
	SetMetadata(v map[string]interface{}) *Sandbox
	GetMetadata() map[string]interface{}
	SetSandboxArn(v string) *Sandbox
	GetSandboxArn() *string
	SetSandboxId(v string) *Sandbox
	GetSandboxId() *string
	SetSandboxIdleTTLInSeconds(v int32) *Sandbox
	GetSandboxIdleTTLInSeconds() *int32
	SetSandboxIdleTimeoutSeconds(v int32) *Sandbox
	GetSandboxIdleTimeoutSeconds() *int32
	SetStatus(v string) *Sandbox
	GetStatus() *string
	SetTemplateId(v string) *Sandbox
	GetTemplateId() *string
	SetTemplateName(v string) *Sandbox
	GetTemplateName() *string
}

type Sandbox struct {
	// Sandbox Creation Time
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-11-26T10:54:17.770719+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Sandbox stop time
	//
	// example:
	//
	// 2025-11-26T10:54:17.770719+08:00
	EndedAt *string `json:"endedAt,omitempty" xml:"endedAt,omitempty"`
	// Last Update Time
	//
	// example:
	//
	// 2025-11-26T10:54:17.770719+08:00
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// Sandbox metadata
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"metadata" xml:"metadata"`
	// Sandbox resource ARN
	//
	// example:
	//
	// acs:ram::1760720386195983:role/aliyunfcdefaultrole
	SandboxArn *string `json:"sandboxArn,omitempty" xml:"sandboxArn,omitempty"`
	// Sandbox ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 01KAWBP6JQD0J3Z34BP4WMX1KG
	SandboxId *string `json:"sandboxId,omitempty" xml:"sandboxId,omitempty"`
	// Sandbox lifecycle duration (deprecated)
	//
	// example:
	//
	// 已弃用
	SandboxIdleTTLInSeconds *int32 `json:"sandboxIdleTTLInSeconds,omitempty" xml:"sandboxIdleTTLInSeconds,omitempty"`
	// Sandbox idle timeout (seconds)
	//
	// example:
	//
	// 1800
	SandboxIdleTimeoutSeconds *int32 `json:"sandboxIdleTimeoutSeconds,omitempty" xml:"sandboxIdleTimeoutSeconds,omitempty"`
	// Status
	//
	// This parameter is required.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Template ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 8d409d30-cac1-445a-95d5-476c47780336.schema
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// Template Name
	//
	// example:
	//
	// my-template
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s Sandbox) String() string {
	return dara.Prettify(s)
}

func (s Sandbox) GoString() string {
	return s.String()
}

func (s *Sandbox) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Sandbox) GetEndedAt() *string {
	return s.EndedAt
}

func (s *Sandbox) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *Sandbox) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *Sandbox) GetSandboxArn() *string {
	return s.SandboxArn
}

func (s *Sandbox) GetSandboxId() *string {
	return s.SandboxId
}

func (s *Sandbox) GetSandboxIdleTTLInSeconds() *int32 {
	return s.SandboxIdleTTLInSeconds
}

func (s *Sandbox) GetSandboxIdleTimeoutSeconds() *int32 {
	return s.SandboxIdleTimeoutSeconds
}

func (s *Sandbox) GetStatus() *string {
	return s.Status
}

func (s *Sandbox) GetTemplateId() *string {
	return s.TemplateId
}

func (s *Sandbox) GetTemplateName() *string {
	return s.TemplateName
}

func (s *Sandbox) SetCreatedAt(v string) *Sandbox {
	s.CreatedAt = &v
	return s
}

func (s *Sandbox) SetEndedAt(v string) *Sandbox {
	s.EndedAt = &v
	return s
}

func (s *Sandbox) SetLastUpdatedAt(v string) *Sandbox {
	s.LastUpdatedAt = &v
	return s
}

func (s *Sandbox) SetMetadata(v map[string]interface{}) *Sandbox {
	s.Metadata = v
	return s
}

func (s *Sandbox) SetSandboxArn(v string) *Sandbox {
	s.SandboxArn = &v
	return s
}

func (s *Sandbox) SetSandboxId(v string) *Sandbox {
	s.SandboxId = &v
	return s
}

func (s *Sandbox) SetSandboxIdleTTLInSeconds(v int32) *Sandbox {
	s.SandboxIdleTTLInSeconds = &v
	return s
}

func (s *Sandbox) SetSandboxIdleTimeoutSeconds(v int32) *Sandbox {
	s.SandboxIdleTimeoutSeconds = &v
	return s
}

func (s *Sandbox) SetStatus(v string) *Sandbox {
	s.Status = &v
	return s
}

func (s *Sandbox) SetTemplateId(v string) *Sandbox {
	s.TemplateId = &v
	return s
}

func (s *Sandbox) SetTemplateName(v string) *Sandbox {
	s.TemplateName = &v
	return s
}

func (s *Sandbox) Validate() error {
	return dara.Validate(s)
}
