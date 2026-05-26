// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProductResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProductResponseBody) *UpdateProductResponse
	GetBody() *UpdateProductResponseBody
}

type UpdateProductResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProductResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductResponse) GoString() string {
	return s.String()
}

func (s *UpdateProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProductResponse) GetBody() *UpdateProductResponseBody {
	return s.Body
}

func (s *UpdateProductResponse) SetHeaders(v map[string]*string) *UpdateProductResponse {
	s.Headers = v
	return s
}

func (s *UpdateProductResponse) SetStatusCode(v int32) *UpdateProductResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProductResponse) SetBody(v *UpdateProductResponseBody) *UpdateProductResponse {
	s.Body = v
	return s
}

func (s *UpdateProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
