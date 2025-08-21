// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWeatherResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWeatherResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWeatherResponse
	GetStatusCode() *int32
	SetBody(v *GetWeatherResponseBody) *GetWeatherResponse
	GetBody() *GetWeatherResponseBody
}

type GetWeatherResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWeatherResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWeatherResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWeatherResponse) GoString() string {
	return s.String()
}

func (s *GetWeatherResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWeatherResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWeatherResponse) GetBody() *GetWeatherResponseBody {
	return s.Body
}

func (s *GetWeatherResponse) SetHeaders(v map[string]*string) *GetWeatherResponse {
	s.Headers = v
	return s
}

func (s *GetWeatherResponse) SetStatusCode(v int32) *GetWeatherResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWeatherResponse) SetBody(v *GetWeatherResponseBody) *GetWeatherResponse {
	s.Body = v
	return s
}

func (s *GetWeatherResponse) Validate() error {
	return dara.Validate(s)
}
