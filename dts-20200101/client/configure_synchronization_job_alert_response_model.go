// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSynchronizationJobAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigureSynchronizationJobAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigureSynchronizationJobAlertResponse
	GetStatusCode() *int32
	SetBody(v *ConfigureSynchronizationJobAlertResponseBody) *ConfigureSynchronizationJobAlertResponse
	GetBody() *ConfigureSynchronizationJobAlertResponseBody
}

type ConfigureSynchronizationJobAlertResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigureSynchronizationJobAlertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigureSynchronizationJobAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSynchronizationJobAlertResponse) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigureSynchronizationJobAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigureSynchronizationJobAlertResponse) GetBody() *ConfigureSynchronizationJobAlertResponseBody {
	return s.Body
}

func (s *ConfigureSynchronizationJobAlertResponse) SetHeaders(v map[string]*string) *ConfigureSynchronizationJobAlertResponse {
	s.Headers = v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponse) SetStatusCode(v int32) *ConfigureSynchronizationJobAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponse) SetBody(v *ConfigureSynchronizationJobAlertResponseBody) *ConfigureSynchronizationJobAlertResponse {
	s.Body = v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
