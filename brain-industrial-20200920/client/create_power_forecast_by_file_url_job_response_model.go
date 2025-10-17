// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePowerForecastByFileUrlJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePowerForecastByFileUrlJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePowerForecastByFileUrlJobResponse
	GetStatusCode() *int32
	SetBody(v *CreatePowerForecastByFileUrlJobResponseBody) *CreatePowerForecastByFileUrlJobResponse
	GetBody() *CreatePowerForecastByFileUrlJobResponseBody
}

type CreatePowerForecastByFileUrlJobResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePowerForecastByFileUrlJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePowerForecastByFileUrlJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePowerForecastByFileUrlJobResponse) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastByFileUrlJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePowerForecastByFileUrlJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePowerForecastByFileUrlJobResponse) GetBody() *CreatePowerForecastByFileUrlJobResponseBody {
	return s.Body
}

func (s *CreatePowerForecastByFileUrlJobResponse) SetHeaders(v map[string]*string) *CreatePowerForecastByFileUrlJobResponse {
	s.Headers = v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponse) SetStatusCode(v int32) *CreatePowerForecastByFileUrlJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponse) SetBody(v *CreatePowerForecastByFileUrlJobResponseBody) *CreatePowerForecastByFileUrlJobResponse {
	s.Body = v
	return s
}

func (s *CreatePowerForecastByFileUrlJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
