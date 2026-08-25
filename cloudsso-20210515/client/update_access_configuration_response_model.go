// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccessConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAccessConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAccessConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAccessConfigurationResponseBody) *UpdateAccessConfigurationResponse
	GetBody() *UpdateAccessConfigurationResponseBody
}

type UpdateAccessConfigurationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccessConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccessConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAccessConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAccessConfigurationResponse) GetBody() *UpdateAccessConfigurationResponseBody {
	return s.Body
}

func (s *UpdateAccessConfigurationResponse) SetHeaders(v map[string]*string) *UpdateAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccessConfigurationResponse) SetStatusCode(v int32) *UpdateAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccessConfigurationResponse) SetBody(v *UpdateAccessConfigurationResponseBody) *UpdateAccessConfigurationResponse {
	s.Body = v
	return s
}

func (s *UpdateAccessConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
