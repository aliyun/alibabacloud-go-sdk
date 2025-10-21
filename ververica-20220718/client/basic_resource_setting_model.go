// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBasicResourceSetting interface {
	dara.Model
	String() string
	GoString() string
	SetJobmanagerResourceSettingSpec(v *BasicResourceSettingSpec) *BasicResourceSetting
	GetJobmanagerResourceSettingSpec() *BasicResourceSettingSpec
	SetParallelism(v int64) *BasicResourceSetting
	GetParallelism() *int64
	SetTaskmanagerResourceSettingSpec(v *BasicResourceSettingSpec) *BasicResourceSetting
	GetTaskmanagerResourceSettingSpec() *BasicResourceSettingSpec
}

type BasicResourceSetting struct {
	JobmanagerResourceSettingSpec *BasicResourceSettingSpec `json:"jobmanagerResourceSettingSpec,omitempty" xml:"jobmanagerResourceSettingSpec,omitempty"`
	// example:
	//
	// 4
	Parallelism                    *int64                    `json:"parallelism,omitempty" xml:"parallelism,omitempty"`
	TaskmanagerResourceSettingSpec *BasicResourceSettingSpec `json:"taskmanagerResourceSettingSpec,omitempty" xml:"taskmanagerResourceSettingSpec,omitempty"`
}

func (s BasicResourceSetting) String() string {
	return dara.Prettify(s)
}

func (s BasicResourceSetting) GoString() string {
	return s.String()
}

func (s *BasicResourceSetting) GetJobmanagerResourceSettingSpec() *BasicResourceSettingSpec {
	return s.JobmanagerResourceSettingSpec
}

func (s *BasicResourceSetting) GetParallelism() *int64 {
	return s.Parallelism
}

func (s *BasicResourceSetting) GetTaskmanagerResourceSettingSpec() *BasicResourceSettingSpec {
	return s.TaskmanagerResourceSettingSpec
}

func (s *BasicResourceSetting) SetJobmanagerResourceSettingSpec(v *BasicResourceSettingSpec) *BasicResourceSetting {
	s.JobmanagerResourceSettingSpec = v
	return s
}

func (s *BasicResourceSetting) SetParallelism(v int64) *BasicResourceSetting {
	s.Parallelism = &v
	return s
}

func (s *BasicResourceSetting) SetTaskmanagerResourceSettingSpec(v *BasicResourceSettingSpec) *BasicResourceSetting {
	s.TaskmanagerResourceSettingSpec = v
	return s
}

func (s *BasicResourceSetting) Validate() error {
	if s.JobmanagerResourceSettingSpec != nil {
		if err := s.JobmanagerResourceSettingSpec.Validate(); err != nil {
			return err
		}
	}
	if s.TaskmanagerResourceSettingSpec != nil {
		if err := s.TaskmanagerResourceSettingSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}
