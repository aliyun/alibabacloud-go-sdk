// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelSetModelProfile interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModelSetModelProfile
	GetDescription() *string
	SetDisplayName(v string) *ModelSetModelProfile
	GetDisplayName() *string
	SetEnabled(v bool) *ModelSetModelProfile
	GetEnabled() *bool
	SetName(v string) *ModelSetModelProfile
	GetName() *string
	SetProps(v *ModelSetModelProfileProps) *ModelSetModelProfile
	GetProps() *ModelSetModelProfileProps
	SetSourceType(v string) *ModelSetModelProfile
	GetSourceType() *string
}

type ModelSetModelProfile struct {
	// example:
	//
	// 通义千问系列速度最快、成本很低的模型，适合简单任务。本模型是动态更新版本，模型更新不会提前通知，模型中英文综合能力显著提升，模型人类偏好显著提升，模型推理能力和复杂指令理解能力显著增强，困难任务上的表现更优，数学、代码能力显著提升。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 通义千问-Turbo-Latest
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-turbo-latest
	Name  *string                    `json:"name,omitempty" xml:"name,omitempty"`
	Props *ModelSetModelProfileProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
	// example:
	//
	// predefined
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ModelSetModelProfile) String() string {
	return dara.Prettify(s)
}

func (s ModelSetModelProfile) GoString() string {
	return s.String()
}

func (s *ModelSetModelProfile) GetDescription() *string {
	return s.Description
}

func (s *ModelSetModelProfile) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ModelSetModelProfile) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModelSetModelProfile) GetName() *string {
	return s.Name
}

func (s *ModelSetModelProfile) GetProps() *ModelSetModelProfileProps {
	return s.Props
}

func (s *ModelSetModelProfile) GetSourceType() *string {
	return s.SourceType
}

func (s *ModelSetModelProfile) SetDescription(v string) *ModelSetModelProfile {
	s.Description = &v
	return s
}

func (s *ModelSetModelProfile) SetDisplayName(v string) *ModelSetModelProfile {
	s.DisplayName = &v
	return s
}

func (s *ModelSetModelProfile) SetEnabled(v bool) *ModelSetModelProfile {
	s.Enabled = &v
	return s
}

func (s *ModelSetModelProfile) SetName(v string) *ModelSetModelProfile {
	s.Name = &v
	return s
}

func (s *ModelSetModelProfile) SetProps(v *ModelSetModelProfileProps) *ModelSetModelProfile {
	s.Props = v
	return s
}

func (s *ModelSetModelProfile) SetSourceType(v string) *ModelSetModelProfile {
	s.SourceType = &v
	return s
}

func (s *ModelSetModelProfile) Validate() error {
	return dara.Validate(s)
}

type ModelSetModelProfileProps struct {
	ContextSize *int64  `json:"contextSize,omitempty" xml:"contextSize,omitempty"`
	LlmMode     *string `json:"llmMode,omitempty" xml:"llmMode,omitempty"`
}

func (s ModelSetModelProfileProps) String() string {
	return dara.Prettify(s)
}

func (s ModelSetModelProfileProps) GoString() string {
	return s.String()
}

func (s *ModelSetModelProfileProps) GetContextSize() *int64 {
	return s.ContextSize
}

func (s *ModelSetModelProfileProps) GetLlmMode() *string {
	return s.LlmMode
}

func (s *ModelSetModelProfileProps) SetContextSize(v int64) *ModelSetModelProfileProps {
	s.ContextSize = &v
	return s
}

func (s *ModelSetModelProfileProps) SetLlmMode(v string) *ModelSetModelProfileProps {
	s.LlmMode = &v
	return s
}

func (s *ModelSetModelProfileProps) Validate() error {
	return dara.Validate(s)
}
