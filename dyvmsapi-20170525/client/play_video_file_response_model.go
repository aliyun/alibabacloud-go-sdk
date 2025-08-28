// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayVideoFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PlayVideoFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PlayVideoFileResponse
	GetStatusCode() *int32
	SetBody(v *PlayVideoFileResponseBody) *PlayVideoFileResponse
	GetBody() *PlayVideoFileResponseBody
}

type PlayVideoFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PlayVideoFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PlayVideoFileResponse) String() string {
	return dara.Prettify(s)
}

func (s PlayVideoFileResponse) GoString() string {
	return s.String()
}

func (s *PlayVideoFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PlayVideoFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PlayVideoFileResponse) GetBody() *PlayVideoFileResponseBody {
	return s.Body
}

func (s *PlayVideoFileResponse) SetHeaders(v map[string]*string) *PlayVideoFileResponse {
	s.Headers = v
	return s
}

func (s *PlayVideoFileResponse) SetStatusCode(v int32) *PlayVideoFileResponse {
	s.StatusCode = &v
	return s
}

func (s *PlayVideoFileResponse) SetBody(v *PlayVideoFileResponseBody) *PlayVideoFileResponse {
	s.Body = v
	return s
}

func (s *PlayVideoFileResponse) Validate() error {
	return dara.Validate(s)
}
