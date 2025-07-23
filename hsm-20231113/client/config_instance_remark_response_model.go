// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigInstanceRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigInstanceRemarkResponse
	GetStatusCode() *int32
	SetBody(v *ConfigInstanceRemarkResponseBody) *ConfigInstanceRemarkResponse
	GetBody() *ConfigInstanceRemarkResponseBody
}

type ConfigInstanceRemarkResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigInstanceRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigInstanceRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceRemarkResponse) GoString() string {
	return s.String()
}

func (s *ConfigInstanceRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigInstanceRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigInstanceRemarkResponse) GetBody() *ConfigInstanceRemarkResponseBody {
	return s.Body
}

func (s *ConfigInstanceRemarkResponse) SetHeaders(v map[string]*string) *ConfigInstanceRemarkResponse {
	s.Headers = v
	return s
}

func (s *ConfigInstanceRemarkResponse) SetStatusCode(v int32) *ConfigInstanceRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigInstanceRemarkResponse) SetBody(v *ConfigInstanceRemarkResponseBody) *ConfigInstanceRemarkResponse {
	s.Body = v
	return s
}

func (s *ConfigInstanceRemarkResponse) Validate() error {
	return dara.Validate(s)
}
