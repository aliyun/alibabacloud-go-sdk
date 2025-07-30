// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureMigrationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigureMigrationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigureMigrationJobResponse
	GetStatusCode() *int32
	SetBody(v *ConfigureMigrationJobResponseBody) *ConfigureMigrationJobResponse
	GetBody() *ConfigureMigrationJobResponseBody
}

type ConfigureMigrationJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigureMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigureMigrationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigureMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigureMigrationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigureMigrationJobResponse) GetBody() *ConfigureMigrationJobResponseBody {
	return s.Body
}

func (s *ConfigureMigrationJobResponse) SetHeaders(v map[string]*string) *ConfigureMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *ConfigureMigrationJobResponse) SetStatusCode(v int32) *ConfigureMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureMigrationJobResponse) SetBody(v *ConfigureMigrationJobResponseBody) *ConfigureMigrationJobResponse {
	s.Body = v
	return s
}

func (s *ConfigureMigrationJobResponse) Validate() error {
	return dara.Validate(s)
}
