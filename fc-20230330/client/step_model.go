// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStep interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v []*string) *Step
	GetArgs() []*string
	SetType(v string) *Step
	GetType() *string
}

type Step struct {
	// The step parameters.
	Args []*string `json:"args" xml:"args" type:"Repeated"`
	// The step type.
	//
	// example:
	//
	// USER, WORKDIR, ENV, RUN
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Step) String() string {
	return dara.Prettify(s)
}

func (s Step) GoString() string {
	return s.String()
}

func (s *Step) GetArgs() []*string {
	return s.Args
}

func (s *Step) GetType() *string {
	return s.Type
}

func (s *Step) SetArgs(v []*string) *Step {
	s.Args = v
	return s
}

func (s *Step) SetType(v string) *Step {
	s.Type = &v
	return s
}

func (s *Step) Validate() error {
	return dara.Validate(s)
}
