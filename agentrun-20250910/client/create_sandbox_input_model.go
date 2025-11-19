// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSandboxInput interface {
	dara.Model
	String() string
	GoString() string
	SetSandboxIdleTimeoutSeconds(v int32) *CreateSandboxInput
	GetSandboxIdleTimeoutSeconds() *int32
	SetTemplateName(v string) *CreateSandboxInput
	GetTemplateName() *string
}

type CreateSandboxInput struct {
	// 沙箱空闲超时时间（秒）
	SandboxIdleTimeoutSeconds *int32 `json:"sandboxIdleTimeoutSeconds,omitempty" xml:"sandboxIdleTimeoutSeconds,omitempty"`
	// 模板名称（系统内部通过 templateName 查询 template_id）
	//
	// This parameter is required.
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s CreateSandboxInput) String() string {
	return dara.Prettify(s)
}

func (s CreateSandboxInput) GoString() string {
	return s.String()
}

func (s *CreateSandboxInput) GetSandboxIdleTimeoutSeconds() *int32 {
	return s.SandboxIdleTimeoutSeconds
}

func (s *CreateSandboxInput) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateSandboxInput) SetSandboxIdleTimeoutSeconds(v int32) *CreateSandboxInput {
	s.SandboxIdleTimeoutSeconds = &v
	return s
}

func (s *CreateSandboxInput) SetTemplateName(v string) *CreateSandboxInput {
	s.TemplateName = &v
	return s
}

func (s *CreateSandboxInput) Validate() error {
	return dara.Validate(s)
}
