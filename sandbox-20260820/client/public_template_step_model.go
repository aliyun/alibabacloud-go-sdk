// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicTemplateStep interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v []*string) *PublicTemplateStep
	GetArgs() []*string
	SetType(v string) *PublicTemplateStep
	GetType() *string
}

type PublicTemplateStep struct {
	// The build instruction parameters. The meaning varies depending on the type: for USER, it is [username, optional passwordless sudo flag], where the second item configures passwordless sudo when the string is true; for WORKDIR, it is [directory]; for ENV, it is [key, value, ...], allowing multiple pairs to be set at once; for RUN, it is [command, optional temporary user].
	Args []*string `json:"args" xml:"args" type:"Repeated"`
	// The build instruction type. Valid values: USER, WORKDIR, ENV, and RUN. The instructions are executed sequentially in the order of the array.
	//
	// example:
	//
	// RUN
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PublicTemplateStep) String() string {
	return dara.Prettify(s)
}

func (s PublicTemplateStep) GoString() string {
	return s.String()
}

func (s *PublicTemplateStep) GetArgs() []*string {
	return s.Args
}

func (s *PublicTemplateStep) GetType() *string {
	return s.Type
}

func (s *PublicTemplateStep) SetArgs(v []*string) *PublicTemplateStep {
	s.Args = v
	return s
}

func (s *PublicTemplateStep) SetType(v string) *PublicTemplateStep {
	s.Type = &v
	return s
}

func (s *PublicTemplateStep) Validate() error {
	return dara.Validate(s)
}
