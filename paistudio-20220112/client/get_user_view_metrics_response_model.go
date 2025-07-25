// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserViewMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserViewMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserViewMetricsResponse
	GetStatusCode() *int32
	SetBody(v *GetUserViewMetricsResponseBody) *GetUserViewMetricsResponse
	GetBody() *GetUserViewMetricsResponseBody
}

type GetUserViewMetricsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserViewMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserViewMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserViewMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetUserViewMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserViewMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserViewMetricsResponse) GetBody() *GetUserViewMetricsResponseBody {
	return s.Body
}

func (s *GetUserViewMetricsResponse) SetHeaders(v map[string]*string) *GetUserViewMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetUserViewMetricsResponse) SetStatusCode(v int32) *GetUserViewMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserViewMetricsResponse) SetBody(v *GetUserViewMetricsResponseBody) *GetUserViewMetricsResponse {
	s.Body = v
	return s
}

func (s *GetUserViewMetricsResponse) Validate() error {
	return dara.Validate(s)
}
