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
	// The configuration item details in JSON format.
	//
	// example:
	//
	// {\\"description\\":\\"This is a new pytorchjob template\\"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The system-generated unique ID of the debug config.
	//
	// example:
	//
	// dc-vf9lowjt3pso
	DebuggerConfigId *string `json:"DebuggerConfigId,omitempty" xml:"DebuggerConfigId,omitempty"`
	// The description of the configuration item.
	//
	// example:
	//
	// This is a basic Pytorch configuration template
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the template configuration item.
	//
	// example:
	//
	// Pytorch Experiment Config
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The creation time in UTC.
	//
	// example:
	//
	// 2021-01-12T14:35:00Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The modification time in UTC.
	//
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
