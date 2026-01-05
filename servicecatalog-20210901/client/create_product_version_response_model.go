// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProductVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProductVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateProductVersionResponseBody) *CreateProductVersionResponse
	GetBody() *CreateProductVersionResponseBody
}

type CreateProductVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProductVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProductVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateProductVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProductVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProductVersionResponse) GetBody() *CreateProductVersionResponseBody {
	return s.Body
}

func (s *CreateProductVersionResponse) SetHeaders(v map[string]*string) *CreateProductVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateProductVersionResponse) SetStatusCode(v int32) *CreateProductVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProductVersionResponse) SetBody(v *CreateProductVersionResponseBody) *CreateProductVersionResponse {
	s.Body = v
	return s
}

func (s *CreateProductVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
