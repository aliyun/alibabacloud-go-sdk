// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RealLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigLayer4RealLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigLayer4RealLimitResponse
	GetStatusCode() *int32
	SetBody(v *ConfigLayer4RealLimitResponseBody) *ConfigLayer4RealLimitResponse
	GetBody() *ConfigLayer4RealLimitResponseBody
}

type ConfigLayer4RealLimitResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer4RealLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigLayer4RealLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RealLimitResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RealLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigLayer4RealLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigLayer4RealLimitResponse) GetBody() *ConfigLayer4RealLimitResponseBody {
	return s.Body
}

func (s *ConfigLayer4RealLimitResponse) SetHeaders(v map[string]*string) *ConfigLayer4RealLimitResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer4RealLimitResponse) SetStatusCode(v int32) *ConfigLayer4RealLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer4RealLimitResponse) SetBody(v *ConfigLayer4RealLimitResponseBody) *ConfigLayer4RealLimitResponse {
	s.Body = v
	return s
}

func (s *ConfigLayer4RealLimitResponse) Validate() error {
	return dara.Validate(s)
}
