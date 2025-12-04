// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessKeyLastUsedProductsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessKeyLastUsedProductsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessKeyLastUsedProductsResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessKeyLastUsedProductsResponseBody) *GetAccessKeyLastUsedProductsResponse
	GetBody() *GetAccessKeyLastUsedProductsResponseBody
}

type GetAccessKeyLastUsedProductsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessKeyLastUsedProductsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessKeyLastUsedProductsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessKeyLastUsedProductsResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedProductsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessKeyLastUsedProductsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessKeyLastUsedProductsResponse) GetBody() *GetAccessKeyLastUsedProductsResponseBody {
	return s.Body
}

func (s *GetAccessKeyLastUsedProductsResponse) SetHeaders(v map[string]*string) *GetAccessKeyLastUsedProductsResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponse) SetStatusCode(v int32) *GetAccessKeyLastUsedProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponse) SetBody(v *GetAccessKeyLastUsedProductsResponseBody) *GetAccessKeyLastUsedProductsResponse {
	s.Body = v
	return s
}

func (s *GetAccessKeyLastUsedProductsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
