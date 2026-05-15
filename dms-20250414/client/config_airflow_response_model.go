// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigAirflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigAirflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigAirflowResponse
	GetStatusCode() *int32
	SetBody(v *ConfigAirflowResponseBody) *ConfigAirflowResponse
	GetBody() *ConfigAirflowResponseBody
}

type ConfigAirflowResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigAirflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigAirflowResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigAirflowResponse) GoString() string {
	return s.String()
}

func (s *ConfigAirflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigAirflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigAirflowResponse) GetBody() *ConfigAirflowResponseBody {
	return s.Body
}

func (s *ConfigAirflowResponse) SetHeaders(v map[string]*string) *ConfigAirflowResponse {
	s.Headers = v
	return s
}

func (s *ConfigAirflowResponse) SetStatusCode(v int32) *ConfigAirflowResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigAirflowResponse) SetBody(v *ConfigAirflowResponseBody) *ConfigAirflowResponse {
	s.Body = v
	return s
}

func (s *ConfigAirflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
