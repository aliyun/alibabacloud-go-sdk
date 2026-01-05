// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProvisionedProductsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProvisionedProductsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProvisionedProductsResponse
	GetStatusCode() *int32
	SetBody(v *ListProvisionedProductsResponseBody) *ListProvisionedProductsResponse
	GetBody() *ListProvisionedProductsResponseBody
}

type ListProvisionedProductsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProvisionedProductsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProvisionedProductsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductsResponse) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProvisionedProductsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProvisionedProductsResponse) GetBody() *ListProvisionedProductsResponseBody {
	return s.Body
}

func (s *ListProvisionedProductsResponse) SetHeaders(v map[string]*string) *ListProvisionedProductsResponse {
	s.Headers = v
	return s
}

func (s *ListProvisionedProductsResponse) SetStatusCode(v int32) *ListProvisionedProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProvisionedProductsResponse) SetBody(v *ListProvisionedProductsResponseBody) *ListProvisionedProductsResponse {
	s.Body = v
	return s
}

func (s *ListProvisionedProductsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
