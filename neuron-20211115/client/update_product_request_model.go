// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ProductUpdateCmd) *UpdateProductRequest
	GetBody() *ProductUpdateCmd
}

type UpdateProductRequest struct {
	Body *ProductUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProductRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductRequest) GetBody() *ProductUpdateCmd {
	return s.Body
}

func (s *UpdateProductRequest) SetBody(v *ProductUpdateCmd) *UpdateProductRequest {
	s.Body = v
	return s
}

func (s *UpdateProductRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
