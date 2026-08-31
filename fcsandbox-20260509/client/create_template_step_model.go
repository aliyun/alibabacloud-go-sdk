// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateStep interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v []*string) *CreateTemplateStep
	GetArgs() []*string
	SetType(v string) *CreateTemplateStep
	GetType() *string
}

type CreateTemplateStep struct {
	// The list of step parameters.
	Args []*string `json:"args,omitempty" xml:"args,omitempty" type:"Repeated"`
	// The step type.
	//
	// example:
	//
	// RUN
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateTemplateStep) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateStep) GoString() string {
	return s.String()
}

func (s *CreateTemplateStep) GetArgs() []*string {
	return s.Args
}

func (s *CreateTemplateStep) GetType() *string {
	return s.Type
}

func (s *CreateTemplateStep) SetArgs(v []*string) *CreateTemplateStep {
	s.Args = v
	return s
}

func (s *CreateTemplateStep) SetType(v string) *CreateTemplateStep {
	s.Type = &v
	return s
}

func (s *CreateTemplateStep) Validate() error {
	return dara.Validate(s)
}
