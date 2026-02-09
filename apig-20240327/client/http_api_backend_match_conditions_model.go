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
	// The configurations of the matching conditions.
	Conditions []*HttpApiBackendMatchCondition `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// Specifies whether the matching condition is the default one.
	//
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
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
