// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpsApplicationConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHttpsApplicationConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHttpsApplicationConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *CreateHttpsApplicationConfigurationResponseBody) *CreateHttpsApplicationConfigurationResponse
	GetBody() *CreateHttpsApplicationConfigurationResponseBody
}

type CreateHttpsApplicationConfigurationResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHttpsApplicationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHttpsApplicationConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpsApplicationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateHttpsApplicationConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHttpsApplicationConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHttpsApplicationConfigurationResponse) GetBody() *CreateHttpsApplicationConfigurationResponseBody {
	return s.Body
}

func (s *CreateHttpsApplicationConfigurationResponse) SetHeaders(v map[string]*string) *CreateHttpsApplicationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateHttpsApplicationConfigurationResponse) SetStatusCode(v int32) *CreateHttpsApplicationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHttpsApplicationConfigurationResponse) SetBody(v *CreateHttpsApplicationConfigurationResponseBody) *CreateHttpsApplicationConfigurationResponse {
	s.Body = v
	return s
}

func (s *CreateHttpsApplicationConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
