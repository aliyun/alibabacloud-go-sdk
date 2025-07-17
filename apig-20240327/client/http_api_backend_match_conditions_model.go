// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiBackendMatchConditions interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v []*HttpApiBackendMatchCondition) *HttpApiBackendMatchConditions
	GetConditions() []*HttpApiBackendMatchCondition
	SetDefault(v bool) *HttpApiBackendMatchConditions
	GetDefault() *bool
}

type HttpApiBackendMatchConditions struct {
	Conditions []*HttpApiBackendMatchCondition `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Default *bool `json:"default,omitempty" xml:"default,omitempty"`
}

func (s HttpApiBackendMatchConditions) String() string {
	return dara.Prettify(s)
}

func (s HttpApiBackendMatchConditions) GoString() string {
	return s.String()
}

func (s *HttpApiBackendMatchConditions) GetConditions() []*HttpApiBackendMatchCondition {
	return s.Conditions
}

func (s *HttpApiBackendMatchConditions) GetDefault() *bool {
	return s.Default
}

func (s *HttpApiBackendMatchConditions) SetConditions(v []*HttpApiBackendMatchCondition) *HttpApiBackendMatchConditions {
	s.Conditions = v
	return s
}

func (s *HttpApiBackendMatchConditions) SetDefault(v bool) *HttpApiBackendMatchConditions {
	s.Default = &v
	return s
}

func (s *HttpApiBackendMatchConditions) Validate() error {
	return dara.Validate(s)
}
