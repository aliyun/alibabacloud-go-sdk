// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCitiesByProvinceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCitiesByProvinceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCitiesByProvinceResponse
	GetStatusCode() *int32
	SetBody(v *ListCitiesByProvinceResponseBody) *ListCitiesByProvinceResponse
	GetBody() *ListCitiesByProvinceResponseBody
}

type ListCitiesByProvinceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCitiesByProvinceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCitiesByProvinceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCitiesByProvinceResponse) GoString() string {
	return s.String()
}

func (s *ListCitiesByProvinceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCitiesByProvinceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCitiesByProvinceResponse) GetBody() *ListCitiesByProvinceResponseBody {
	return s.Body
}

func (s *ListCitiesByProvinceResponse) SetHeaders(v map[string]*string) *ListCitiesByProvinceResponse {
	s.Headers = v
	return s
}

func (s *ListCitiesByProvinceResponse) SetStatusCode(v int32) *ListCitiesByProvinceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCitiesByProvinceResponse) SetBody(v *ListCitiesByProvinceResponseBody) *ListCitiesByProvinceResponse {
	s.Body = v
	return s
}

func (s *ListCitiesByProvinceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
