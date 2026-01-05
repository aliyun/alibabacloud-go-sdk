// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsAsAdminResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProductsAsAdminResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProductsAsAdminResponse
	GetStatusCode() *int32
	SetBody(v *ListProductsAsAdminResponseBody) *ListProductsAsAdminResponse
	GetBody() *ListProductsAsAdminResponseBody
}

type ListProductsAsAdminResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductsAsAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductsAsAdminResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProductsAsAdminResponse) GoString() string {
	return s.String()
}

func (s *ListProductsAsAdminResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProductsAsAdminResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProductsAsAdminResponse) GetBody() *ListProductsAsAdminResponseBody {
	return s.Body
}

func (s *ListProductsAsAdminResponse) SetHeaders(v map[string]*string) *ListProductsAsAdminResponse {
	s.Headers = v
	return s
}

func (s *ListProductsAsAdminResponse) SetStatusCode(v int32) *ListProductsAsAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductsAsAdminResponse) SetBody(v *ListProductsAsAdminResponseBody) *ListProductsAsAdminResponse {
	s.Body = v
	return s
}

func (s *ListProductsAsAdminResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
