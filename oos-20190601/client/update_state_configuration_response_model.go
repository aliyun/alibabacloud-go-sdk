// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStateConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStateConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStateConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStateConfigurationResponseBody) *UpdateStateConfigurationResponse
	GetBody() *UpdateStateConfigurationResponseBody
}

type UpdateStateConfigurationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStateConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStateConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStateConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateStateConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStateConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStateConfigurationResponse) GetBody() *UpdateStateConfigurationResponseBody {
	return s.Body
}

func (s *UpdateStateConfigurationResponse) SetHeaders(v map[string]*string) *UpdateStateConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateStateConfigurationResponse) SetStatusCode(v int32) *UpdateStateConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStateConfigurationResponse) SetBody(v *UpdateStateConfigurationResponseBody) *UpdateStateConfigurationResponse {
	s.Body = v
	return s
}

func (s *UpdateStateConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
