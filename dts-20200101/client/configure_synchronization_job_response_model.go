// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSynchronizationJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigureSynchronizationJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigureSynchronizationJobResponse
	GetStatusCode() *int32
	SetBody(v *ConfigureSynchronizationJobResponseBody) *ConfigureSynchronizationJobResponse
	GetBody() *ConfigureSynchronizationJobResponseBody
}

type ConfigureSynchronizationJobResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigureSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigureSynchronizationJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigureSynchronizationJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigureSynchronizationJobResponse) GetBody() *ConfigureSynchronizationJobResponseBody {
	return s.Body
}

func (s *ConfigureSynchronizationJobResponse) SetHeaders(v map[string]*string) *ConfigureSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *ConfigureSynchronizationJobResponse) SetStatusCode(v int32) *ConfigureSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureSynchronizationJobResponse) SetBody(v *ConfigureSynchronizationJobResponseBody) *ConfigureSynchronizationJobResponse {
	s.Body = v
	return s
}

func (s *ConfigureSynchronizationJobResponse) Validate() error {
	return dara.Validate(s)
}
