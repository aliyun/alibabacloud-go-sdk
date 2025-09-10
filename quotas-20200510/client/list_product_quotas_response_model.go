// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductQuotasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProductQuotasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProductQuotasResponse
	GetStatusCode() *int32
	SetBody(v *ListProductQuotasResponseBody) *ListProductQuotasResponse
	GetBody() *ListProductQuotasResponseBody
}

type ListProductQuotasResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductQuotasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProductQuotasResponse) GoString() string {
	return s.String()
}

func (s *ListProductQuotasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProductQuotasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProductQuotasResponse) GetBody() *ListProductQuotasResponseBody {
	return s.Body
}

func (s *ListProductQuotasResponse) SetHeaders(v map[string]*string) *ListProductQuotasResponse {
	s.Headers = v
	return s
}

func (s *ListProductQuotasResponse) SetStatusCode(v int32) *ListProductQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductQuotasResponse) SetBody(v *ListProductQuotasResponseBody) *ListProductQuotasResponse {
	s.Body = v
	return s
}

func (s *ListProductQuotasResponse) Validate() error {
	return dara.Validate(s)
}
