// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStreamingResourceSetting interface {
	dara.Model
	String() string
	GoString() string
	SetBasicResourceSetting(v *BasicResourceSetting) *StreamingResourceSetting
	GetBasicResourceSetting() *BasicResourceSetting
	SetExpertResourceSetting(v *ExpertResourceSetting) *StreamingResourceSetting
	GetExpertResourceSetting() *ExpertResourceSetting
	SetResourceSettingMode(v string) *StreamingResourceSetting
	GetResourceSettingMode() *string
}

type StreamingResourceSetting struct {
	BasicResourceSetting  *BasicResourceSetting  `json:"basicResourceSetting,omitempty" xml:"basicResourceSetting,omitempty"`
	ExpertResourceSetting *ExpertResourceSetting `json:"expertResourceSetting,omitempty" xml:"expertResourceSetting,omitempty"`
	// example:
	//
	// EXPERT
	ResourceSettingMode *string `json:"resourceSettingMode,omitempty" xml:"resourceSettingMode,omitempty"`
}

func (s StreamingResourceSetting) String() string {
	return dara.Prettify(s)
}

func (s StreamingResourceSetting) GoString() string {
	return s.String()
}

func (s *StreamingResourceSetting) GetBasicResourceSetting() *BasicResourceSetting {
	return s.BasicResourceSetting
}

func (s *StreamingResourceSetting) GetExpertResourceSetting() *ExpertResourceSetting {
	return s.ExpertResourceSetting
}

func (s *StreamingResourceSetting) GetResourceSettingMode() *string {
	return s.ResourceSettingMode
}

func (s *StreamingResourceSetting) SetBasicResourceSetting(v *BasicResourceSetting) *StreamingResourceSetting {
	s.BasicResourceSetting = v
	return s
}

func (s *StreamingResourceSetting) SetExpertResourceSetting(v *ExpertResourceSetting) *StreamingResourceSetting {
	s.ExpertResourceSetting = v
	return s
}

func (s *StreamingResourceSetting) SetResourceSettingMode(v string) *StreamingResourceSetting {
	s.ResourceSettingMode = &v
	return s
}

func (s *StreamingResourceSetting) Validate() error {
	if s.BasicResourceSetting != nil {
		if err := s.BasicResourceSetting.Validate(); err != nil {
			return err
		}
	}
	if s.ExpertResourceSetting != nil {
		if err := s.ExpertResourceSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}
