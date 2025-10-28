// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductOrdersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProductOrdersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProductOrdersResponse
	GetStatusCode() *int32
	SetBody(v *CreateProductOrdersResponseBody) *CreateProductOrdersResponse
	GetBody() *CreateProductOrdersResponseBody
}

type CreateProductOrdersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProductOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProductOrdersResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProductOrdersResponse) GoString() string {
	return s.String()
}

func (s *CreateProductOrdersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProductOrdersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProductOrdersResponse) GetBody() *CreateProductOrdersResponseBody {
	return s.Body
}

func (s *CreateProductOrdersResponse) SetHeaders(v map[string]*string) *CreateProductOrdersResponse {
	s.Headers = v
	return s
}

func (s *CreateProductOrdersResponse) SetStatusCode(v int32) *CreateProductOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProductOrdersResponse) SetBody(v *CreateProductOrdersResponseBody) *CreateProductOrdersResponse {
	s.Body = v
	return s
}

func (s *CreateProductOrdersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
