// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigHealthCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigHealthCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigHealthCheckResponse
	GetStatusCode() *int32
	SetBody(v *ConfigHealthCheckResponseBody) *ConfigHealthCheckResponse
	GetBody() *ConfigHealthCheckResponseBody
}

type ConfigHealthCheckResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigHealthCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigHealthCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigHealthCheckResponse) GoString() string {
	return s.String()
}

func (s *ConfigHealthCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigHealthCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigHealthCheckResponse) GetBody() *ConfigHealthCheckResponseBody {
	return s.Body
}

func (s *ConfigHealthCheckResponse) SetHeaders(v map[string]*string) *ConfigHealthCheckResponse {
	s.Headers = v
	return s
}

func (s *ConfigHealthCheckResponse) SetStatusCode(v int32) *ConfigHealthCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigHealthCheckResponse) SetBody(v *ConfigHealthCheckResponseBody) *ConfigHealthCheckResponse {
	s.Body = v
	return s
}

func (s *ConfigHealthCheckResponse) Validate() error {
	return dara.Validate(s)
}
