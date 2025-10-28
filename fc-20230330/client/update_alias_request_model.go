// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateAliasInput) *UpdateAliasRequest
	GetBody() *UpdateAliasInput
}

type UpdateAliasRequest struct {
	// The alias information to be updated.
	//
	// This parameter is required.
	Body *UpdateAliasInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAliasRequest) GoString() string {
	return s.String()
}

func (s *UpdateAliasRequest) GetBody() *UpdateAliasInput {
	return s.Body
}

func (s *UpdateAliasRequest) SetBody(v *UpdateAliasInput) *UpdateAliasRequest {
	s.Body = v
	return s
}

func (s *UpdateAliasRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
