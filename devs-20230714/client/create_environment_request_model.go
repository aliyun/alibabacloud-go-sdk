// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnvironmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Environment) *CreateEnvironmentRequest
	GetBody() *Environment
}

type CreateEnvironmentRequest struct {
	Body *Environment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnvironmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentRequest) GetBody() *Environment {
	return s.Body
}

func (s *CreateEnvironmentRequest) SetBody(v *Environment) *CreateEnvironmentRequest {
	s.Body = v
	return s
}

func (s *CreateEnvironmentRequest) Validate() error {
	return dara.Validate(s)
}
