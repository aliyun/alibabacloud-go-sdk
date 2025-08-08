// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeConfig interface {
	dara.Model
	String() string
	GoString() string
	SetSteps(v []interface{}) *InitializeConfig
	GetSteps() []interface{}
}

type InitializeConfig struct {
	Steps []interface{} `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s InitializeConfig) String() string {
	return dara.Prettify(s)
}

func (s InitializeConfig) GoString() string {
	return s.String()
}

func (s *InitializeConfig) GetSteps() []interface{} {
	return s.Steps
}

func (s *InitializeConfig) SetSteps(v []interface{}) *InitializeConfig {
	s.Steps = v
	return s
}

func (s *InitializeConfig) Validate() error {
	return dara.Validate(s)
}
