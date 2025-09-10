// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDependentQuotasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDependentQuotasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDependentQuotasResponse
	GetStatusCode() *int32
	SetBody(v *ListDependentQuotasResponseBody) *ListDependentQuotasResponse
	GetBody() *ListDependentQuotasResponseBody
}

type ListDependentQuotasResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDependentQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDependentQuotasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDependentQuotasResponse) GoString() string {
	return s.String()
}

func (s *ListDependentQuotasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDependentQuotasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDependentQuotasResponse) GetBody() *ListDependentQuotasResponseBody {
	return s.Body
}

func (s *ListDependentQuotasResponse) SetHeaders(v map[string]*string) *ListDependentQuotasResponse {
	s.Headers = v
	return s
}

func (s *ListDependentQuotasResponse) SetStatusCode(v int32) *ListDependentQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDependentQuotasResponse) SetBody(v *ListDependentQuotasResponseBody) *ListDependentQuotasResponse {
	s.Body = v
	return s
}

func (s *ListDependentQuotasResponse) Validate() error {
	return dara.Validate(s)
}
