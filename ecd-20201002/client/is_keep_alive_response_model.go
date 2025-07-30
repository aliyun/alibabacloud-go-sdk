// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsKeepAliveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IsKeepAliveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IsKeepAliveResponse
	GetStatusCode() *int32
	SetBody(v *IsKeepAliveResponseBody) *IsKeepAliveResponse
	GetBody() *IsKeepAliveResponseBody
}

type IsKeepAliveResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IsKeepAliveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IsKeepAliveResponse) String() string {
	return dara.Prettify(s)
}

func (s IsKeepAliveResponse) GoString() string {
	return s.String()
}

func (s *IsKeepAliveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IsKeepAliveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IsKeepAliveResponse) GetBody() *IsKeepAliveResponseBody {
	return s.Body
}

func (s *IsKeepAliveResponse) SetHeaders(v map[string]*string) *IsKeepAliveResponse {
	s.Headers = v
	return s
}

func (s *IsKeepAliveResponse) SetStatusCode(v int32) *IsKeepAliveResponse {
	s.StatusCode = &v
	return s
}

func (s *IsKeepAliveResponse) SetBody(v *IsKeepAliveResponseBody) *IsKeepAliveResponse {
	s.Body = v
	return s
}

func (s *IsKeepAliveResponse) Validate() error {
	return dara.Validate(s)
}
