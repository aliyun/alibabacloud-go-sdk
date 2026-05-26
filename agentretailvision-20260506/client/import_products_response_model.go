// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportProductsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportProductsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportProductsResponse
	GetStatusCode() *int32
	SetBody(v *ImportProductsResponseBody) *ImportProductsResponse
	GetBody() *ImportProductsResponseBody
}

type ImportProductsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportProductsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportProductsResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportProductsResponse) GoString() string {
	return s.String()
}

func (s *ImportProductsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportProductsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportProductsResponse) GetBody() *ImportProductsResponseBody {
	return s.Body
}

func (s *ImportProductsResponse) SetHeaders(v map[string]*string) *ImportProductsResponse {
	s.Headers = v
	return s
}

func (s *ImportProductsResponse) SetStatusCode(v int32) *ImportProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportProductsResponse) SetBody(v *ImportProductsResponseBody) *ImportProductsResponse {
	s.Body = v
	return s
}

func (s *ImportProductsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
