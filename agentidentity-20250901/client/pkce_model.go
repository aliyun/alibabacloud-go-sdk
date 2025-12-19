// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPKCE interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *PKCE
	GetEnabled() *bool
}

type PKCE struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s PKCE) String() string {
	return dara.Prettify(s)
}

func (s PKCE) GoString() string {
	return s.String()
}

func (s *PKCE) GetEnabled() *bool {
	return s.Enabled
}

func (s *PKCE) SetEnabled(v bool) *PKCE {
	s.Enabled = &v
	return s
}

func (s *PKCE) Validate() error {
	return dara.Validate(s)
}
