// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureDtsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigureDtsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigureDtsJobResponse
	GetStatusCode() *int32
	SetBody(v *ConfigureDtsJobResponseBody) *ConfigureDtsJobResponse
	GetBody() *ConfigureDtsJobResponseBody
}

type ConfigureDtsJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigureDtsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigureDtsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigureDtsJobResponse) GoString() string {
	return s.String()
}

func (s *ConfigureDtsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigureDtsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigureDtsJobResponse) GetBody() *ConfigureDtsJobResponseBody {
	return s.Body
}

func (s *ConfigureDtsJobResponse) SetHeaders(v map[string]*string) *ConfigureDtsJobResponse {
	s.Headers = v
	return s
}

func (s *ConfigureDtsJobResponse) SetStatusCode(v int32) *ConfigureDtsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureDtsJobResponse) SetBody(v *ConfigureDtsJobResponseBody) *ConfigureDtsJobResponse {
	s.Body = v
	return s
}

func (s *ConfigureDtsJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
