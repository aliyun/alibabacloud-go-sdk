// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHookConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHookConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHookConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHookConfigurationResponseBody) *UpdateHookConfigurationResponse
	GetBody() *UpdateHookConfigurationResponseBody
}

type UpdateHookConfigurationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHookConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHookConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHookConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateHookConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHookConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHookConfigurationResponse) GetBody() *UpdateHookConfigurationResponseBody {
	return s.Body
}

func (s *UpdateHookConfigurationResponse) SetHeaders(v map[string]*string) *UpdateHookConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateHookConfigurationResponse) SetStatusCode(v int32) *UpdateHookConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHookConfigurationResponse) SetBody(v *UpdateHookConfigurationResponseBody) *UpdateHookConfigurationResponse {
	s.Body = v
	return s
}

func (s *UpdateHookConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
