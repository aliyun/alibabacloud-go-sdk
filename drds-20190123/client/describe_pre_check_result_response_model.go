// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreCheckResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePreCheckResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePreCheckResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribePreCheckResultResponseBody) *DescribePreCheckResultResponse
	GetBody() *DescribePreCheckResultResponseBody
}

type DescribePreCheckResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePreCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePreCheckResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckResultResponse) GoString() string {
	return s.String()
}

func (s *DescribePreCheckResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePreCheckResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePreCheckResultResponse) GetBody() *DescribePreCheckResultResponseBody {
	return s.Body
}

func (s *DescribePreCheckResultResponse) SetHeaders(v map[string]*string) *DescribePreCheckResultResponse {
	s.Headers = v
	return s
}

func (s *DescribePreCheckResultResponse) SetStatusCode(v int32) *DescribePreCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePreCheckResultResponse) SetBody(v *DescribePreCheckResultResponseBody) *DescribePreCheckResultResponse {
	s.Body = v
	return s
}

func (s *DescribePreCheckResultResponse) Validate() error {
	return dara.Validate(s)
}
