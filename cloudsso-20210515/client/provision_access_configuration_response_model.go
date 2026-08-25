// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProvisionAccessConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProvisionAccessConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProvisionAccessConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *ProvisionAccessConfigurationResponseBody) *ProvisionAccessConfigurationResponse
	GetBody() *ProvisionAccessConfigurationResponseBody
}

type ProvisionAccessConfigurationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProvisionAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProvisionAccessConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s ProvisionAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ProvisionAccessConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProvisionAccessConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProvisionAccessConfigurationResponse) GetBody() *ProvisionAccessConfigurationResponseBody {
	return s.Body
}

func (s *ProvisionAccessConfigurationResponse) SetHeaders(v map[string]*string) *ProvisionAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ProvisionAccessConfigurationResponse) SetStatusCode(v int32) *ProvisionAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ProvisionAccessConfigurationResponse) SetBody(v *ProvisionAccessConfigurationResponseBody) *ProvisionAccessConfigurationResponse {
	s.Body = v
	return s
}

func (s *ProvisionAccessConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
