// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataDistResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataDistResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataDistResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataDistResultResponseBody) *DescribeDataDistResultResponse
	GetBody() *DescribeDataDistResultResponseBody
}

type DescribeDataDistResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataDistResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataDistResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataDistResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataDistResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataDistResultResponse) GetBody() *DescribeDataDistResultResponseBody {
	return s.Body
}

func (s *DescribeDataDistResultResponse) SetHeaders(v map[string]*string) *DescribeDataDistResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataDistResultResponse) SetStatusCode(v int32) *DescribeDataDistResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataDistResultResponse) SetBody(v *DescribeDataDistResultResponseBody) *DescribeDataDistResultResponse {
	s.Body = v
	return s
}

func (s *DescribeDataDistResultResponse) Validate() error {
	return dara.Validate(s)
}
