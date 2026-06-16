// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCodeInterpreterSessionInput interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *StartCodeInterpreterSessionInput
	GetName() *string
}

type StartCodeInterpreterSessionInput struct {
	// The name of the code interpreter session, used to identify and distinguish between different session instances
	//
	// if can be null:
	// true
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s StartCodeInterpreterSessionInput) String() string {
	return dara.Prettify(s)
}

func (s StartCodeInterpreterSessionInput) GoString() string {
	return s.String()
}

func (s *StartCodeInterpreterSessionInput) GetName() *string {
	return s.Name
}

func (s *StartCodeInterpreterSessionInput) SetName(v string) *StartCodeInterpreterSessionInput {
	s.Name = &v
	return s
}

func (s *StartCodeInterpreterSessionInput) Validate() error {
	return dara.Validate(s)
}
