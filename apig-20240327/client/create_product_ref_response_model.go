// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductRefResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProductRefResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProductRefResponse
	GetStatusCode() *int32
	SetBody(v *CreateProductRefResponseBody) *CreateProductRefResponse
	GetBody() *CreateProductRefResponseBody
}

type CreateProductRefResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProductRefResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProductRefResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProductRefResponse) GoString() string {
	return s.String()
}

func (s *CreateProductRefResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProductRefResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProductRefResponse) GetBody() *CreateProductRefResponseBody {
	return s.Body
}

func (s *CreateProductRefResponse) SetHeaders(v map[string]*string) *CreateProductRefResponse {
	s.Headers = v
	return s
}

func (s *CreateProductRefResponse) SetStatusCode(v int32) *CreateProductRefResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProductRefResponse) SetBody(v *CreateProductRefResponseBody) *CreateProductRefResponse {
	s.Body = v
	return s
}

func (s *CreateProductRefResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
