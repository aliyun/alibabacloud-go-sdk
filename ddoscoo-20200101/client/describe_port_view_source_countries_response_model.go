// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortViewSourceCountriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePortViewSourceCountriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePortViewSourceCountriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribePortViewSourceCountriesResponseBody) *DescribePortViewSourceCountriesResponse
	GetBody() *DescribePortViewSourceCountriesResponseBody
}

type DescribePortViewSourceCountriesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePortViewSourceCountriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePortViewSourceCountriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePortViewSourceCountriesResponse) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceCountriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePortViewSourceCountriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePortViewSourceCountriesResponse) GetBody() *DescribePortViewSourceCountriesResponseBody {
	return s.Body
}

func (s *DescribePortViewSourceCountriesResponse) SetHeaders(v map[string]*string) *DescribePortViewSourceCountriesResponse {
	s.Headers = v
	return s
}

func (s *DescribePortViewSourceCountriesResponse) SetStatusCode(v int32) *DescribePortViewSourceCountriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePortViewSourceCountriesResponse) SetBody(v *DescribePortViewSourceCountriesResponseBody) *DescribePortViewSourceCountriesResponse {
	s.Body = v
	return s
}

func (s *DescribePortViewSourceCountriesResponse) Validate() error {
	return dara.Validate(s)
}
