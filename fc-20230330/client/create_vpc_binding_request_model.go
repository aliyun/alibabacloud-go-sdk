// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcBindingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateVpcBindingInput) *CreateVpcBindingRequest
	GetBody() *CreateVpcBindingInput
}

type CreateVpcBindingRequest struct {
	// The VPC binding configurations.
	//
	// This parameter is required.
	Body *CreateVpcBindingInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcBindingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcBindingRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcBindingRequest) GetBody() *CreateVpcBindingInput {
	return s.Body
}

func (s *CreateVpcBindingRequest) SetBody(v *CreateVpcBindingInput) *CreateVpcBindingRequest {
	s.Body = v
	return s
}

func (s *CreateVpcBindingRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
