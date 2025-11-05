// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCountriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCountriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCountriesResponse
	GetStatusCode() *int32
	SetBody(v *ListCountriesResponseBody) *ListCountriesResponse
	GetBody() *ListCountriesResponseBody
}

type ListCountriesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCountriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCountriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCountriesResponse) GoString() string {
	return s.String()
}

func (s *ListCountriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCountriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCountriesResponse) GetBody() *ListCountriesResponseBody {
	return s.Body
}

func (s *ListCountriesResponse) SetHeaders(v map[string]*string) *ListCountriesResponse {
	s.Headers = v
	return s
}

func (s *ListCountriesResponse) SetStatusCode(v int32) *ListCountriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCountriesResponse) SetBody(v *ListCountriesResponseBody) *ListCountriesResponse {
	s.Body = v
	return s
}

func (s *ListCountriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
