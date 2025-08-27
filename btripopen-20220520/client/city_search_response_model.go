// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCitySearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CitySearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CitySearchResponse
	GetStatusCode() *int32
	SetBody(v *CitySearchResponseBody) *CitySearchResponse
	GetBody() *CitySearchResponseBody
}

type CitySearchResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CitySearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CitySearchResponse) String() string {
	return dara.Prettify(s)
}

func (s CitySearchResponse) GoString() string {
	return s.String()
}

func (s *CitySearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CitySearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CitySearchResponse) GetBody() *CitySearchResponseBody {
	return s.Body
}

func (s *CitySearchResponse) SetHeaders(v map[string]*string) *CitySearchResponse {
	s.Headers = v
	return s
}

func (s *CitySearchResponse) SetStatusCode(v int32) *CitySearchResponse {
	s.StatusCode = &v
	return s
}

func (s *CitySearchResponse) SetBody(v *CitySearchResponseBody) *CitySearchResponse {
	s.Body = v
	return s
}

func (s *CitySearchResponse) Validate() error {
	return dara.Validate(s)
}
