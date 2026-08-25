// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeprovisionAccessConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeprovisionAccessConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeprovisionAccessConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DeprovisionAccessConfigurationResponseBody) *DeprovisionAccessConfigurationResponse
	GetBody() *DeprovisionAccessConfigurationResponseBody
}

type DeprovisionAccessConfigurationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeprovisionAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeprovisionAccessConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeprovisionAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeprovisionAccessConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeprovisionAccessConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeprovisionAccessConfigurationResponse) GetBody() *DeprovisionAccessConfigurationResponseBody {
	return s.Body
}

func (s *DeprovisionAccessConfigurationResponse) SetHeaders(v map[string]*string) *DeprovisionAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeprovisionAccessConfigurationResponse) SetStatusCode(v int32) *DeprovisionAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponse) SetBody(v *DeprovisionAccessConfigurationResponseBody) *DeprovisionAccessConfigurationResponse {
	s.Body = v
	return s
}

func (s *DeprovisionAccessConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
