// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpAbroadCountryInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIpAbroadCountryInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIpAbroadCountryInfosResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIpAbroadCountryInfosResponseBody) *DescribeIpAbroadCountryInfosResponse
	GetBody() *DescribeIpAbroadCountryInfosResponseBody
}

type DescribeIpAbroadCountryInfosResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpAbroadCountryInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpAbroadCountryInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpAbroadCountryInfosResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpAbroadCountryInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIpAbroadCountryInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIpAbroadCountryInfosResponse) GetBody() *DescribeIpAbroadCountryInfosResponseBody {
	return s.Body
}

func (s *DescribeIpAbroadCountryInfosResponse) SetHeaders(v map[string]*string) *DescribeIpAbroadCountryInfosResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpAbroadCountryInfosResponse) SetStatusCode(v int32) *DescribeIpAbroadCountryInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosResponse) SetBody(v *DescribeIpAbroadCountryInfosResponseBody) *DescribeIpAbroadCountryInfosResponse {
	s.Body = v
	return s
}

func (s *DescribeIpAbroadCountryInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
