// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ProductCreateCmd) *CreateProductRequest
	GetBody() *ProductCreateCmd
}

type CreateProductRequest struct {
	Body *ProductCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProductRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProductRequest) GoString() string {
	return s.String()
}

func (s *CreateProductRequest) GetBody() *ProductCreateCmd {
	return s.Body
}

func (s *CreateProductRequest) SetBody(v *ProductCreateCmd) *CreateProductRequest {
	s.Body = v
	return s
}

func (s *CreateProductRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
