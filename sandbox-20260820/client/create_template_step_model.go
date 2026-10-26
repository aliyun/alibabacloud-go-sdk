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
	// The build instruction parameters, whose meaning varies depending on the type: for USER, it is [username, optional passwordless sudo flag], where the second item configures passwordless sudo when it is the string true; for WORKDIR, it is [directory]; for ENV, it is [key, value, ...], allowing multiple pairs to be set at once; for RUN, it is [command, optional temporary user].
	Args []*string `json:"args" xml:"args" type:"Repeated"`
	// The build instruction type, which supports four values: USER, WORKDIR, ENV, and RUN, executed sequentially in the order of the array.
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
