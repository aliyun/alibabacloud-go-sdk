// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvironmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Environment) *UpdateEnvironmentRequest
	GetBody() *Environment
}

type UpdateEnvironmentRequest struct {
	Body *Environment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnvironmentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentRequest) GetBody() *Environment {
	return s.Body
}

func (s *UpdateEnvironmentRequest) SetBody(v *Environment) *UpdateEnvironmentRequest {
	s.Body = v
	return s
}

func (s *UpdateEnvironmentRequest) Validate() error {
	return dara.Validate(s)
}
