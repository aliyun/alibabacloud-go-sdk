// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpsApplicationConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHttpsApplicationConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHttpsApplicationConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHttpsApplicationConfigurationResponseBody) *UpdateHttpsApplicationConfigurationResponse
	GetBody() *UpdateHttpsApplicationConfigurationResponseBody
}

type UpdateHttpsApplicationConfigurationResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHttpsApplicationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHttpsApplicationConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpsApplicationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateHttpsApplicationConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHttpsApplicationConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHttpsApplicationConfigurationResponse) GetBody() *UpdateHttpsApplicationConfigurationResponseBody {
	return s.Body
}

func (s *UpdateHttpsApplicationConfigurationResponse) SetHeaders(v map[string]*string) *UpdateHttpsApplicationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateHttpsApplicationConfigurationResponse) SetStatusCode(v int32) *UpdateHttpsApplicationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHttpsApplicationConfigurationResponse) SetBody(v *UpdateHttpsApplicationConfigurationResponseBody) *UpdateHttpsApplicationConfigurationResponse {
	s.Body = v
	return s
}

func (s *UpdateHttpsApplicationConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
