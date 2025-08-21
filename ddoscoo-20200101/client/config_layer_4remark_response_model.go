// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigLayer4RemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigLayer4RemarkResponse
	GetStatusCode() *int32
	SetBody(v *ConfigLayer4RemarkResponseBody) *ConfigLayer4RemarkResponse
	GetBody() *ConfigLayer4RemarkResponseBody
}

type ConfigLayer4RemarkResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer4RemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigLayer4RemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RemarkResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigLayer4RemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigLayer4RemarkResponse) GetBody() *ConfigLayer4RemarkResponseBody {
	return s.Body
}

func (s *ConfigLayer4RemarkResponse) SetHeaders(v map[string]*string) *ConfigLayer4RemarkResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer4RemarkResponse) SetStatusCode(v int32) *ConfigLayer4RemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer4RemarkResponse) SetBody(v *ConfigLayer4RemarkResponseBody) *ConfigLayer4RemarkResponse {
	s.Body = v
	return s
}

func (s *ConfigLayer4RemarkResponse) Validate() error {
	return dara.Validate(s)
}
