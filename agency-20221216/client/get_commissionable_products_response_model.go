// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommissionableProductsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCommissionableProductsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCommissionableProductsResponse
	GetStatusCode() *int32
	SetBody(v *GetCommissionableProductsResponseBody) *GetCommissionableProductsResponse
	GetBody() *GetCommissionableProductsResponseBody
}

type GetCommissionableProductsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCommissionableProductsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCommissionableProductsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCommissionableProductsResponse) GoString() string {
	return s.String()
}

func (s *GetCommissionableProductsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCommissionableProductsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCommissionableProductsResponse) GetBody() *GetCommissionableProductsResponseBody {
	return s.Body
}

func (s *GetCommissionableProductsResponse) SetHeaders(v map[string]*string) *GetCommissionableProductsResponse {
	s.Headers = v
	return s
}

func (s *GetCommissionableProductsResponse) SetStatusCode(v int32) *GetCommissionableProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCommissionableProductsResponse) SetBody(v *GetCommissionableProductsResponseBody) *GetCommissionableProductsResponse {
	s.Body = v
	return s
}

func (s *GetCommissionableProductsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
