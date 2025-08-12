// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomEventHistogramResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomEventHistogramResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomEventHistogramResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomEventHistogramResponseBody) *DescribeCustomEventHistogramResponse
	GetBody() *DescribeCustomEventHistogramResponseBody
}

type DescribeCustomEventHistogramResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomEventHistogramResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomEventHistogramResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEventHistogramResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventHistogramResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomEventHistogramResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomEventHistogramResponse) GetBody() *DescribeCustomEventHistogramResponseBody {
	return s.Body
}

func (s *DescribeCustomEventHistogramResponse) SetHeaders(v map[string]*string) *DescribeCustomEventHistogramResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomEventHistogramResponse) SetStatusCode(v int32) *DescribeCustomEventHistogramResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomEventHistogramResponse) SetBody(v *DescribeCustomEventHistogramResponseBody) *DescribeCustomEventHistogramResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomEventHistogramResponse) Validate() error {
	return dara.Validate(s)
}
