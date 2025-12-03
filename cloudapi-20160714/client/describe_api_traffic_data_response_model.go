// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiTrafficDataResponseBody) *DescribeApiTrafficDataResponse
	GetBody() *DescribeApiTrafficDataResponseBody
}

type DescribeApiTrafficDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiTrafficDataResponse) GetBody() *DescribeApiTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeApiTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeApiTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiTrafficDataResponse) SetStatusCode(v int32) *DescribeApiTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiTrafficDataResponse) SetBody(v *DescribeApiTrafficDataResponseBody) *DescribeApiTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeApiTrafficDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
