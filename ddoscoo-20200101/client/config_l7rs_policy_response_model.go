// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigL7RsPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigL7RsPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigL7RsPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ConfigL7RsPolicyResponseBody) *ConfigL7RsPolicyResponse
	GetBody() *ConfigL7RsPolicyResponseBody
}

type ConfigL7RsPolicyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigL7RsPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigL7RsPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigL7RsPolicyResponse) GoString() string {
	return s.String()
}

func (s *ConfigL7RsPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigL7RsPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigL7RsPolicyResponse) GetBody() *ConfigL7RsPolicyResponseBody {
	return s.Body
}

func (s *ConfigL7RsPolicyResponse) SetHeaders(v map[string]*string) *ConfigL7RsPolicyResponse {
	s.Headers = v
	return s
}

func (s *ConfigL7RsPolicyResponse) SetStatusCode(v int32) *ConfigL7RsPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigL7RsPolicyResponse) SetBody(v *ConfigL7RsPolicyResponseBody) *ConfigL7RsPolicyResponse {
	s.Body = v
	return s
}

func (s *ConfigL7RsPolicyResponse) Validate() error {
	return dara.Validate(s)
}
