// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProductResponse
	GetStatusCode() *int32
	SetBody(v *Product) *CreateProductResponse
	GetBody() *Product
}

type CreateProductResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Product           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProductResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProductResponse) GoString() string {
	return s.String()
}

func (s *CreateProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProductResponse) GetBody() *Product {
	return s.Body
}

func (s *CreateProductResponse) SetHeaders(v map[string]*string) *CreateProductResponse {
	s.Headers = v
	return s
}

func (s *CreateProductResponse) SetStatusCode(v int32) *CreateProductResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProductResponse) SetBody(v *Product) *CreateProductResponse {
	s.Body = v
	return s
}

func (s *CreateProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
