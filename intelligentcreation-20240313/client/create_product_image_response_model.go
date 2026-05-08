// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProductImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProductImageResponse
	GetStatusCode() *int32
	SetBody(v *CreateProductImageResponseBody) *CreateProductImageResponse
	GetBody() *CreateProductImageResponseBody
}

type CreateProductImageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProductImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProductImageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProductImageResponse) GoString() string {
	return s.String()
}

func (s *CreateProductImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProductImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProductImageResponse) GetBody() *CreateProductImageResponseBody {
	return s.Body
}

func (s *CreateProductImageResponse) SetHeaders(v map[string]*string) *CreateProductImageResponse {
	s.Headers = v
	return s
}

func (s *CreateProductImageResponse) SetStatusCode(v int32) *CreateProductImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProductImageResponse) SetBody(v *CreateProductImageResponseBody) *CreateProductImageResponse {
	s.Body = v
	return s
}

func (s *CreateProductImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
