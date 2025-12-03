// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiLatencyDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiLatencyDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiLatencyDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiLatencyDataResponseBody) *DescribeApiLatencyDataResponse
	GetBody() *DescribeApiLatencyDataResponseBody
}

type DescribeApiLatencyDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiLatencyDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiLatencyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiLatencyDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiLatencyDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiLatencyDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiLatencyDataResponse) GetBody() *DescribeApiLatencyDataResponseBody {
	return s.Body
}

func (s *DescribeApiLatencyDataResponse) SetHeaders(v map[string]*string) *DescribeApiLatencyDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiLatencyDataResponse) SetStatusCode(v int32) *DescribeApiLatencyDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiLatencyDataResponse) SetBody(v *DescribeApiLatencyDataResponseBody) *DescribeApiLatencyDataResponse {
	s.Body = v
	return s
}

func (s *DescribeApiLatencyDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
