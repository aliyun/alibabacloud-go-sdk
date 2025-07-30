// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWhiteIpListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WhiteIpListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WhiteIpListResponse
	GetStatusCode() *int32
	SetBody(v *WhiteIpListResponseBody) *WhiteIpListResponse
	GetBody() *WhiteIpListResponseBody
}

type WhiteIpListResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WhiteIpListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WhiteIpListResponse) String() string {
	return dara.Prettify(s)
}

func (s WhiteIpListResponse) GoString() string {
	return s.String()
}

func (s *WhiteIpListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WhiteIpListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WhiteIpListResponse) GetBody() *WhiteIpListResponseBody {
	return s.Body
}

func (s *WhiteIpListResponse) SetHeaders(v map[string]*string) *WhiteIpListResponse {
	s.Headers = v
	return s
}

func (s *WhiteIpListResponse) SetStatusCode(v int32) *WhiteIpListResponse {
	s.StatusCode = &v
	return s
}

func (s *WhiteIpListResponse) SetBody(v *WhiteIpListResponseBody) *WhiteIpListResponse {
	s.Body = v
	return s
}

func (s *WhiteIpListResponse) Validate() error {
	return dara.Validate(s)
}
