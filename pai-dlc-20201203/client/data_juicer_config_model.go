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
	SetExecutionMode(v string) *DataJuicerConfig
	GetExecutionMode() *string
}

type DataJuicerConfig struct {
	CommandType   *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
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

func (s *DataJuicerConfig) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *DataJuicerConfig) SetCommandType(v string) *DataJuicerConfig {
	s.CommandType = &v
	return s
}

func (s *DataJuicerConfig) SetExecutionMode(v string) *DataJuicerConfig {
	s.ExecutionMode = &v
	return s
}

func (s *DataJuicerConfig) Validate() error {
	return dara.Validate(s)
}
