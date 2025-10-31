// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnycastPopLocationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAnycastPopLocationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAnycastPopLocationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAnycastPopLocationsResponseBody) *DescribeAnycastPopLocationsResponse
	GetBody() *DescribeAnycastPopLocationsResponseBody
}

type DescribeAnycastPopLocationsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAnycastPopLocationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAnycastPopLocationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnycastPopLocationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAnycastPopLocationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAnycastPopLocationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAnycastPopLocationsResponse) GetBody() *DescribeAnycastPopLocationsResponseBody {
	return s.Body
}

func (s *DescribeAnycastPopLocationsResponse) SetHeaders(v map[string]*string) *DescribeAnycastPopLocationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAnycastPopLocationsResponse) SetStatusCode(v int32) *DescribeAnycastPopLocationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAnycastPopLocationsResponse) SetBody(v *DescribeAnycastPopLocationsResponseBody) *DescribeAnycastPopLocationsResponse {
	s.Body = v
	return s
}

func (s *DescribeAnycastPopLocationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
