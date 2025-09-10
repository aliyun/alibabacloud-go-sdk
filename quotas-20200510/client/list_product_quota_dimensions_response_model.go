// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductQuotaDimensionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProductQuotaDimensionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProductQuotaDimensionsResponse
	GetStatusCode() *int32
	SetBody(v *ListProductQuotaDimensionsResponseBody) *ListProductQuotaDimensionsResponse
	GetBody() *ListProductQuotaDimensionsResponseBody
}

type ListProductQuotaDimensionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductQuotaDimensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductQuotaDimensionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProductQuotaDimensionsResponse) GoString() string {
	return s.String()
}

func (s *ListProductQuotaDimensionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProductQuotaDimensionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProductQuotaDimensionsResponse) GetBody() *ListProductQuotaDimensionsResponseBody {
	return s.Body
}

func (s *ListProductQuotaDimensionsResponse) SetHeaders(v map[string]*string) *ListProductQuotaDimensionsResponse {
	s.Headers = v
	return s
}

func (s *ListProductQuotaDimensionsResponse) SetStatusCode(v int32) *ListProductQuotaDimensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductQuotaDimensionsResponse) SetBody(v *ListProductQuotaDimensionsResponseBody) *ListProductQuotaDimensionsResponse {
	s.Body = v
	return s
}

func (s *ListProductQuotaDimensionsResponse) Validate() error {
	return dara.Validate(s)
}
