// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceCommandStep interface {
	dara.Model
	String() string
	GoString() string
	SetPath(v string) *ServiceCommandStep
	GetPath() *string
	SetRun(v string) *ServiceCommandStep
	GetRun() *string
}

type ServiceCommandStep struct {
	// example:
	//
	// ./
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// s invoke
	Run *string `json:"run,omitempty" xml:"run,omitempty"`
}

func (s ServiceCommandStep) String() string {
	return dara.Prettify(s)
}

func (s ServiceCommandStep) GoString() string {
	return s.String()
}

func (s *ServiceCommandStep) GetPath() *string {
	return s.Path
}

func (s *ServiceCommandStep) GetRun() *string {
	return s.Run
}

func (s *ServiceCommandStep) SetPath(v string) *ServiceCommandStep {
	s.Path = &v
	return s
}

func (s *ServiceCommandStep) SetRun(v string) *ServiceCommandStep {
	s.Run = &v
	return s
}

func (s *ServiceCommandStep) Validate() error {
	return dara.Validate(s)
}
