// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebuggerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DebuggerConfig
	GetContent() *string
	SetDebuggerConfigId(v string) *DebuggerConfig
	GetDebuggerConfigId() *string
	SetDescription(v string) *DebuggerConfig
	GetDescription() *string
	SetDisplayName(v string) *DebuggerConfig
	GetDisplayName() *string
	SetGmtCreateTime(v string) *DebuggerConfig
	GetGmtCreateTime() *string
	SetGmtModifyTime(v string) *DebuggerConfig
	GetGmtModifyTime() *string
}

type DebuggerConfig struct {
	// example:
	//
	// {\"description\":\"这是一个新的pytorchjob模板\"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// dc-vf9lowjt3pso
	DebuggerConfigId *string `json:"DebuggerConfigId,omitempty" xml:"DebuggerConfigId,omitempty"`
	// example:
	//
	// 这是一个Pytorch的基础配置模板
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Pytorch Experiment Config
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:00Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:00Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
}

func (s DebuggerConfig) String() string {
	return dara.Prettify(s)
}

func (s DebuggerConfig) GoString() string {
	return s.String()
}

func (s *DebuggerConfig) GetContent() *string {
	return s.Content
}

func (s *DebuggerConfig) GetDebuggerConfigId() *string {
	return s.DebuggerConfigId
}

func (s *DebuggerConfig) GetDescription() *string {
	return s.Description
}

func (s *DebuggerConfig) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DebuggerConfig) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *DebuggerConfig) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *DebuggerConfig) SetContent(v string) *DebuggerConfig {
	s.Content = &v
	return s
}

func (s *DebuggerConfig) SetDebuggerConfigId(v string) *DebuggerConfig {
	s.DebuggerConfigId = &v
	return s
}

func (s *DebuggerConfig) SetDescription(v string) *DebuggerConfig {
	s.Description = &v
	return s
}

func (s *DebuggerConfig) SetDisplayName(v string) *DebuggerConfig {
	s.DisplayName = &v
	return s
}

func (s *DebuggerConfig) SetGmtCreateTime(v string) *DebuggerConfig {
	s.GmtCreateTime = &v
	return s
}

func (s *DebuggerConfig) SetGmtModifyTime(v string) *DebuggerConfig {
	s.GmtModifyTime = &v
	return s
}

func (s *DebuggerConfig) Validate() error {
	return dara.Validate(s)
}
