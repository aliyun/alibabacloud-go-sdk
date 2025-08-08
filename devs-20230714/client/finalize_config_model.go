// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinalizeConfig interface {
	dara.Model
	String() string
	GoString() string
	SetSteps(v []interface{}) *FinalizeConfig
	GetSteps() []interface{}
}

type FinalizeConfig struct {
	Steps []interface{} `json:"steps,omitempty" xml:"steps,omitempty" type:"Repeated"`
}

func (s FinalizeConfig) String() string {
	return dara.Prettify(s)
}

func (s FinalizeConfig) GoString() string {
	return s.String()
}

func (s *FinalizeConfig) GetSteps() []interface{} {
	return s.Steps
}

func (s *FinalizeConfig) SetSteps(v []interface{}) *FinalizeConfig {
	s.Steps = v
	return s
}

func (s *FinalizeConfig) Validate() error {
	return dara.Validate(s)
}
