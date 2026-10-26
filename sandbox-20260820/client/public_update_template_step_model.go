// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicUpdateTemplateStep interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v []*string) *PublicUpdateTemplateStep
	GetArgs() []*string
	SetType(v string) *PublicUpdateTemplateStep
	GetType() *string
}

type PublicUpdateTemplateStep struct {
	// The build instruction parameters. The meaning varies depending on type: for USER, it is [username, optional passwordless sudo flag], and the second item configures passwordless sudo when the string is true; for WORKDIR, it is [directory]; for ENV, it is [key, value, ...], and multiple pairs can be set at once; for RUN, it is [command, optional temporary user].
	Args []*string `json:"args" xml:"args" type:"Repeated"`
	// The build instruction type. Valid values: USER, WORKDIR, ENV, and RUN. The instructions are executed sequentially in the order of the array.
	//
	// example:
	//
	// RUN
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PublicUpdateTemplateStep) String() string {
	return dara.Prettify(s)
}

func (s PublicUpdateTemplateStep) GoString() string {
	return s.String()
}

func (s *PublicUpdateTemplateStep) GetArgs() []*string {
	return s.Args
}

func (s *PublicUpdateTemplateStep) GetType() *string {
	return s.Type
}

func (s *PublicUpdateTemplateStep) SetArgs(v []*string) *PublicUpdateTemplateStep {
	s.Args = v
	return s
}

func (s *PublicUpdateTemplateStep) SetType(v string) *PublicUpdateTemplateStep {
	s.Type = &v
	return s
}

func (s *PublicUpdateTemplateStep) Validate() error {
	return dara.Validate(s)
}
