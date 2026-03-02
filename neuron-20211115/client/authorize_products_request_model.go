// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *AuthorizeProductsCmd) *AuthorizeProductsRequest
	GetBody() *AuthorizeProductsCmd
}

type AuthorizeProductsRequest struct {
	Body *AuthorizeProductsCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeProductsRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeProductsRequest) GetBody() *AuthorizeProductsCmd {
	return s.Body
}

func (s *AuthorizeProductsRequest) SetBody(v *AuthorizeProductsCmd) *AuthorizeProductsRequest {
	s.Body = v
	return s
}

func (s *AuthorizeProductsRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
