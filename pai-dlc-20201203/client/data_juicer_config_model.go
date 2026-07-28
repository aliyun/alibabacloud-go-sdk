// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataJuicerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCommandType(v string) *DataJuicerConfig
	GetCommandType() *string
	SetEnableResourceEstimation(v bool) *DataJuicerConfig
	GetEnableResourceEstimation() *bool
	SetExecutionMode(v string) *DataJuicerConfig
	GetExecutionMode() *string
	SetResourceLimit(v *ResourceLimit) *DataJuicerConfig
	GetResourceLimit() *ResourceLimit
}

type DataJuicerConfig struct {
	// The command type. Valid values:
	//
	// - shell: shell command.
	//
	// - config: DataJuicer YAML configuration.
	//
	// example:
	//
	// config
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// Specifies whether to enable resource estimation. When resource estimation is enabled, the execution mode must be distributed, and the command type must be config (DataJuicer YAML configuration).
	//
	// example:
	//
	// true
	EnableResourceEstimation *bool `json:"EnableResourceEstimation,omitempty" xml:"EnableResourceEstimation,omitempty"`
	// The execution mode. Valid values:
	//
	// - standalone: single-node.
	//
	// - distributed: distributed.
	//
	// example:
	//
	// standalone
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The resource estimation limit. This parameter takes effect only when resource estimation is enabled.
	ResourceLimit *ResourceLimit `json:"ResourceLimit,omitempty" xml:"ResourceLimit,omitempty"`
}

func (s DataJuicerConfig) String() string {
	return dara.Prettify(s)
}

func (s DataJuicerConfig) GoString() string {
	return s.String()
}

func (s *DataJuicerConfig) GetCommandType() *string {
	return s.CommandType
}

func (s *DataJuicerConfig) GetEnableResourceEstimation() *bool {
	return s.EnableResourceEstimation
}

func (s *DataJuicerConfig) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *DataJuicerConfig) GetResourceLimit() *ResourceLimit {
	return s.ResourceLimit
}

func (s *DataJuicerConfig) SetCommandType(v string) *DataJuicerConfig {
	s.CommandType = &v
	return s
}

func (s *DataJuicerConfig) SetEnableResourceEstimation(v bool) *DataJuicerConfig {
	s.EnableResourceEstimation = &v
	return s
}

func (s *DataJuicerConfig) SetExecutionMode(v string) *DataJuicerConfig {
	s.ExecutionMode = &v
	return s
}

func (s *DataJuicerConfig) SetResourceLimit(v *ResourceLimit) *DataJuicerConfig {
	s.ResourceLimit = v
	return s
}

func (s *DataJuicerConfig) Validate() error {
	if s.ResourceLimit != nil {
		if err := s.ResourceLimit.Validate(); err != nil {
			return err
		}
	}
	return nil
}
