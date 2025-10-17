// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadForecastByFileUrlJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLoadForecastByFileUrlJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLoadForecastByFileUrlJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateLoadForecastByFileUrlJobResponseBody) *CreateLoadForecastByFileUrlJobResponse
	GetBody() *CreateLoadForecastByFileUrlJobResponseBody
}

type CreateLoadForecastByFileUrlJobResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLoadForecastByFileUrlJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLoadForecastByFileUrlJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadForecastByFileUrlJobResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastByFileUrlJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLoadForecastByFileUrlJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLoadForecastByFileUrlJobResponse) GetBody() *CreateLoadForecastByFileUrlJobResponseBody {
	return s.Body
}

func (s *CreateLoadForecastByFileUrlJobResponse) SetHeaders(v map[string]*string) *CreateLoadForecastByFileUrlJobResponse {
	s.Headers = v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponse) SetStatusCode(v int32) *CreateLoadForecastByFileUrlJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponse) SetBody(v *CreateLoadForecastByFileUrlJobResponseBody) *CreateLoadForecastByFileUrlJobResponse {
	s.Body = v
	return s
}

func (s *CreateLoadForecastByFileUrlJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
