// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceComponentStep interface {
	dara.Model
	String() string
	GoString() string
	SetComponent(v string) *ServiceComponentStep
	GetComponent() *string
}

type ServiceComponentStep struct {
	// example:
	//
	// fc3 invoke
	Component *string `json:"component,omitempty" xml:"component,omitempty"`
}

func (s ServiceComponentStep) String() string {
	return dara.Prettify(s)
}

func (s ServiceComponentStep) GoString() string {
	return s.String()
}

func (s *ServiceComponentStep) GetComponent() *string {
	return s.Component
}

func (s *ServiceComponentStep) SetComponent(v string) *ServiceComponentStep {
	s.Component = &v
	return s
}

func (s *ServiceComponentStep) Validate() error {
	return dara.Validate(s)
}
