// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAccessConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAccessConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *CreateAccessConfigurationResponseBody) *CreateAccessConfigurationResponse
	GetBody() *CreateAccessConfigurationResponseBody
}

type CreateAccessConfigurationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccessConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAccessConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAccessConfigurationResponse) GetBody() *CreateAccessConfigurationResponseBody {
	return s.Body
}

func (s *CreateAccessConfigurationResponse) SetHeaders(v map[string]*string) *CreateAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessConfigurationResponse) SetStatusCode(v int32) *CreateAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessConfigurationResponse) SetBody(v *CreateAccessConfigurationResponseBody) *CreateAccessConfigurationResponse {
	s.Body = v
	return s
}

func (s *CreateAccessConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
