// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigDataAgentMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *ConfigDataAgentMemoryRequest
	GetDMSUnit() *string
	SetEnabled(v bool) *ConfigDataAgentMemoryRequest
	GetEnabled() *bool
	SetRecallEnabled(v bool) *ConfigDataAgentMemoryRequest
	GetRecallEnabled() *bool
}

type ConfigDataAgentMemoryRequest struct {
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// true
	RecallEnabled *bool `json:"RecallEnabled,omitempty" xml:"RecallEnabled,omitempty"`
}

func (s ConfigDataAgentMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigDataAgentMemoryRequest) GoString() string {
	return s.String()
}

func (s *ConfigDataAgentMemoryRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *ConfigDataAgentMemoryRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ConfigDataAgentMemoryRequest) GetRecallEnabled() *bool {
	return s.RecallEnabled
}

func (s *ConfigDataAgentMemoryRequest) SetDMSUnit(v string) *ConfigDataAgentMemoryRequest {
	s.DMSUnit = &v
	return s
}

func (s *ConfigDataAgentMemoryRequest) SetEnabled(v bool) *ConfigDataAgentMemoryRequest {
	s.Enabled = &v
	return s
}

func (s *ConfigDataAgentMemoryRequest) SetRecallEnabled(v bool) *ConfigDataAgentMemoryRequest {
	s.RecallEnabled = &v
	return s
}

func (s *ConfigDataAgentMemoryRequest) Validate() error {
	return dara.Validate(s)
}
