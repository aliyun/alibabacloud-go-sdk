// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadForecastJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLoadForecastJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLoadForecastJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateLoadForecastJobResponseBody) *CreateLoadForecastJobResponse
	GetBody() *CreateLoadForecastJobResponseBody
}

type CreateLoadForecastJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLoadForecastJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLoadForecastJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadForecastJobResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLoadForecastJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLoadForecastJobResponse) GetBody() *CreateLoadForecastJobResponseBody {
	return s.Body
}

func (s *CreateLoadForecastJobResponse) SetHeaders(v map[string]*string) *CreateLoadForecastJobResponse {
	s.Headers = v
	return s
}

func (s *CreateLoadForecastJobResponse) SetStatusCode(v int32) *CreateLoadForecastJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoadForecastJobResponse) SetBody(v *CreateLoadForecastJobResponseBody) *CreateLoadForecastJobResponse {
	s.Body = v
	return s
}

func (s *CreateLoadForecastJobResponse) Validate() error {
	return dara.Validate(s)
}
