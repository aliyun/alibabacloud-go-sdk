// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpsBasicConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHttpsBasicConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHttpsBasicConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHttpsBasicConfigurationResponseBody) *UpdateHttpsBasicConfigurationResponse
	GetBody() *UpdateHttpsBasicConfigurationResponseBody
}

type UpdateHttpsBasicConfigurationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHttpsBasicConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHttpsBasicConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpsBasicConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateHttpsBasicConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHttpsBasicConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHttpsBasicConfigurationResponse) GetBody() *UpdateHttpsBasicConfigurationResponseBody {
	return s.Body
}

func (s *UpdateHttpsBasicConfigurationResponse) SetHeaders(v map[string]*string) *UpdateHttpsBasicConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateHttpsBasicConfigurationResponse) SetStatusCode(v int32) *UpdateHttpsBasicConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationResponse) SetBody(v *UpdateHttpsBasicConfigurationResponseBody) *UpdateHttpsBasicConfigurationResponse {
	s.Body = v
	return s
}

func (s *UpdateHttpsBasicConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
