// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpRequestVodTestToolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HttpRequestVodTestToolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HttpRequestVodTestToolResponse
	GetStatusCode() *int32
	SetBody(v *HttpRequestVodTestToolResponseBody) *HttpRequestVodTestToolResponse
	GetBody() *HttpRequestVodTestToolResponseBody
}

type HttpRequestVodTestToolResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HttpRequestVodTestToolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HttpRequestVodTestToolResponse) String() string {
	return dara.Prettify(s)
}

func (s HttpRequestVodTestToolResponse) GoString() string {
	return s.String()
}

func (s *HttpRequestVodTestToolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HttpRequestVodTestToolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HttpRequestVodTestToolResponse) GetBody() *HttpRequestVodTestToolResponseBody {
	return s.Body
}

func (s *HttpRequestVodTestToolResponse) SetHeaders(v map[string]*string) *HttpRequestVodTestToolResponse {
	s.Headers = v
	return s
}

func (s *HttpRequestVodTestToolResponse) SetStatusCode(v int32) *HttpRequestVodTestToolResponse {
	s.StatusCode = &v
	return s
}

func (s *HttpRequestVodTestToolResponse) SetBody(v *HttpRequestVodTestToolResponseBody) *HttpRequestVodTestToolResponse {
	s.Body = v
	return s
}

func (s *HttpRequestVodTestToolResponse) Validate() error {
	return dara.Validate(s)
}
