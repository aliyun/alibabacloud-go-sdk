// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigurationRecorderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConfigurationRecorderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConfigurationRecorderResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConfigurationRecorderResponseBody) *UpdateConfigurationRecorderResponse
	GetBody() *UpdateConfigurationRecorderResponseBody
}

type UpdateConfigurationRecorderResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConfigurationRecorderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConfigurationRecorderResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigurationRecorderResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigurationRecorderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConfigurationRecorderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConfigurationRecorderResponse) GetBody() *UpdateConfigurationRecorderResponseBody {
	return s.Body
}

func (s *UpdateConfigurationRecorderResponse) SetHeaders(v map[string]*string) *UpdateConfigurationRecorderResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigurationRecorderResponse) SetStatusCode(v int32) *UpdateConfigurationRecorderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConfigurationRecorderResponse) SetBody(v *UpdateConfigurationRecorderResponseBody) *UpdateConfigurationRecorderResponse {
	s.Body = v
	return s
}

func (s *UpdateConfigurationRecorderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
