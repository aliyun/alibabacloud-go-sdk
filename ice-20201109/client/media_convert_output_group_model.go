// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertOutputGroup interface {
	dara.Model
	String() string
	GoString() string
	SetGroupConfig(v *MediaConvertOutputGroupConfig) *MediaConvertOutputGroup
	GetGroupConfig() *MediaConvertOutputGroupConfig
	SetName(v string) *MediaConvertOutputGroup
	GetName() *string
	SetOutputs(v []*MediaConvertOutputGroupOutput) *MediaConvertOutputGroup
	GetOutputs() []*MediaConvertOutputGroupOutput
}

type MediaConvertOutputGroup struct {
	GroupConfig *MediaConvertOutputGroupConfig   `json:"GroupConfig,omitempty" xml:"GroupConfig,omitempty"`
	Name        *string                          `json:"Name,omitempty" xml:"Name,omitempty"`
	Outputs     []*MediaConvertOutputGroupOutput `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
}

func (s MediaConvertOutputGroup) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertOutputGroup) GoString() string {
	return s.String()
}

func (s *MediaConvertOutputGroup) GetGroupConfig() *MediaConvertOutputGroupConfig {
	return s.GroupConfig
}

func (s *MediaConvertOutputGroup) GetName() *string {
	return s.Name
}

func (s *MediaConvertOutputGroup) GetOutputs() []*MediaConvertOutputGroupOutput {
	return s.Outputs
}

func (s *MediaConvertOutputGroup) SetGroupConfig(v *MediaConvertOutputGroupConfig) *MediaConvertOutputGroup {
	s.GroupConfig = v
	return s
}

func (s *MediaConvertOutputGroup) SetName(v string) *MediaConvertOutputGroup {
	s.Name = &v
	return s
}

func (s *MediaConvertOutputGroup) SetOutputs(v []*MediaConvertOutputGroupOutput) *MediaConvertOutputGroup {
	s.Outputs = v
	return s
}

func (s *MediaConvertOutputGroup) Validate() error {
	if s.GroupConfig != nil {
		if err := s.GroupConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Outputs != nil {
		for _, item := range s.Outputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
