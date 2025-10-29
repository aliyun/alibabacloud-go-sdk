// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBrowserSessionInput interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *StartBrowserSessionInput
	GetName() *string
}

type StartBrowserSessionInput struct {
	// if can be null:
	// true
	//
	// example:
	//
	// my-browser-session
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s StartBrowserSessionInput) String() string {
	return dara.Prettify(s)
}

func (s StartBrowserSessionInput) GoString() string {
	return s.String()
}

func (s *StartBrowserSessionInput) GetName() *string {
	return s.Name
}

func (s *StartBrowserSessionInput) SetName(v string) *StartBrowserSessionInput {
	s.Name = &v
	return s
}

func (s *StartBrowserSessionInput) Validate() error {
	return dara.Validate(s)
}
