// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomLinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomLinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomLinesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomLinesResponseBody) *DescribeCustomLinesResponse
	GetBody() *DescribeCustomLinesResponseBody
}

type DescribeCustomLinesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomLinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomLinesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLinesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomLinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomLinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomLinesResponse) GetBody() *DescribeCustomLinesResponseBody {
	return s.Body
}

func (s *DescribeCustomLinesResponse) SetHeaders(v map[string]*string) *DescribeCustomLinesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomLinesResponse) SetStatusCode(v int32) *DescribeCustomLinesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomLinesResponse) SetBody(v *DescribeCustomLinesResponseBody) *DescribeCustomLinesResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomLinesResponse) Validate() error {
	return dara.Validate(s)
}
