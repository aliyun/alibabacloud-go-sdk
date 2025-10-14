// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemEventHistogramResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSystemEventHistogramResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSystemEventHistogramResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSystemEventHistogramResponseBody) *DescribeSystemEventHistogramResponse
	GetBody() *DescribeSystemEventHistogramResponseBody
}

type DescribeSystemEventHistogramResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSystemEventHistogramResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSystemEventHistogramResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventHistogramResponse) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventHistogramResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSystemEventHistogramResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSystemEventHistogramResponse) GetBody() *DescribeSystemEventHistogramResponseBody {
	return s.Body
}

func (s *DescribeSystemEventHistogramResponse) SetHeaders(v map[string]*string) *DescribeSystemEventHistogramResponse {
	s.Headers = v
	return s
}

func (s *DescribeSystemEventHistogramResponse) SetStatusCode(v int32) *DescribeSystemEventHistogramResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSystemEventHistogramResponse) SetBody(v *DescribeSystemEventHistogramResponseBody) *DescribeSystemEventHistogramResponse {
	s.Body = v
	return s
}

func (s *DescribeSystemEventHistogramResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
