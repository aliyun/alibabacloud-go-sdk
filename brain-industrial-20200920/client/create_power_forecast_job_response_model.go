// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePowerForecastJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePowerForecastJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePowerForecastJobResponse
	GetStatusCode() *int32
	SetBody(v *CreatePowerForecastJobResponseBody) *CreatePowerForecastJobResponse
	GetBody() *CreatePowerForecastJobResponseBody
}

type CreatePowerForecastJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePowerForecastJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePowerForecastJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePowerForecastJobResponse) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePowerForecastJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePowerForecastJobResponse) GetBody() *CreatePowerForecastJobResponseBody {
	return s.Body
}

func (s *CreatePowerForecastJobResponse) SetHeaders(v map[string]*string) *CreatePowerForecastJobResponse {
	s.Headers = v
	return s
}

func (s *CreatePowerForecastJobResponse) SetStatusCode(v int32) *CreatePowerForecastJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePowerForecastJobResponse) SetBody(v *CreatePowerForecastJobResponseBody) *CreatePowerForecastJobResponse {
	s.Body = v
	return s
}

func (s *CreatePowerForecastJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
