// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureMigrationJobAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigureMigrationJobAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigureMigrationJobAlertResponse
	GetStatusCode() *int32
	SetBody(v *ConfigureMigrationJobAlertResponseBody) *ConfigureMigrationJobAlertResponse
	GetBody() *ConfigureMigrationJobAlertResponseBody
}

type ConfigureMigrationJobAlertResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigureMigrationJobAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigureMigrationJobAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigureMigrationJobAlertResponse) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigureMigrationJobAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigureMigrationJobAlertResponse) GetBody() *ConfigureMigrationJobAlertResponseBody {
	return s.Body
}

func (s *ConfigureMigrationJobAlertResponse) SetHeaders(v map[string]*string) *ConfigureMigrationJobAlertResponse {
	s.Headers = v
	return s
}

func (s *ConfigureMigrationJobAlertResponse) SetStatusCode(v int32) *ConfigureMigrationJobAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureMigrationJobAlertResponse) SetBody(v *ConfigureMigrationJobAlertResponseBody) *ConfigureMigrationJobAlertResponse {
	s.Body = v
	return s
}

func (s *ConfigureMigrationJobAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
