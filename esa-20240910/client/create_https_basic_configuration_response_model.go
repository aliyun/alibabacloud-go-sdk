// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpsBasicConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHttpsBasicConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHttpsBasicConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *CreateHttpsBasicConfigurationResponseBody) *CreateHttpsBasicConfigurationResponse
	GetBody() *CreateHttpsBasicConfigurationResponseBody
}

type CreateHttpsBasicConfigurationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHttpsBasicConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHttpsBasicConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpsBasicConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateHttpsBasicConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHttpsBasicConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHttpsBasicConfigurationResponse) GetBody() *CreateHttpsBasicConfigurationResponseBody {
	return s.Body
}

func (s *CreateHttpsBasicConfigurationResponse) SetHeaders(v map[string]*string) *CreateHttpsBasicConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateHttpsBasicConfigurationResponse) SetStatusCode(v int32) *CreateHttpsBasicConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHttpsBasicConfigurationResponse) SetBody(v *CreateHttpsBasicConfigurationResponseBody) *CreateHttpsBasicConfigurationResponse {
	s.Body = v
	return s
}

func (s *CreateHttpsBasicConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
