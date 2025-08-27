// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAirportSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AirportSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AirportSearchResponse
	GetStatusCode() *int32
	SetBody(v *AirportSearchResponseBody) *AirportSearchResponse
	GetBody() *AirportSearchResponseBody
}

type AirportSearchResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AirportSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AirportSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s AirportSearchResponse) GoString() string {
	return s.String()
}

func (s *AirportSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AirportSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AirportSearchResponse) GetBody() *AirportSearchResponseBody {
	return s.Body
}

func (s *AirportSearchResponse) SetHeaders(v map[string]*string) *AirportSearchResponse {
	s.Headers = v
	return s
}

func (s *AirportSearchResponse) SetStatusCode(v int32) *AirportSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *AirportSearchResponse) SetBody(v *AirportSearchResponseBody) *AirportSearchResponse {
	s.Body = v
	return s
}

func (s *AirportSearchResponse) Validate() error {
	return dara.Validate(s)
}
