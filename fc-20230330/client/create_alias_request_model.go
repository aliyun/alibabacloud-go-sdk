// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateAliasInput) *CreateAliasRequest
	GetBody() *CreateAliasInput
}

type CreateAliasRequest struct {
	// The request parameters for creating an alias.
	//
	// This parameter is required.
	Body *CreateAliasInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAliasRequest) GoString() string {
	return s.String()
}

func (s *CreateAliasRequest) GetBody() *CreateAliasInput {
	return s.Body
}

func (s *CreateAliasRequest) SetBody(v *CreateAliasInput) *CreateAliasRequest {
	s.Body = v
	return s
}

func (s *CreateAliasRequest) Validate() error {
	return dara.Validate(s)
}
