// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCheckResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCheckResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCheckResultResponseBody) *DescribeCheckResultResponse
	GetBody() *DescribeCheckResultResponseBody
}

type DescribeCheckResultResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCheckResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeCheckResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCheckResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCheckResultResponse) GetBody() *DescribeCheckResultResponseBody {
	return s.Body
}

func (s *DescribeCheckResultResponse) SetHeaders(v map[string]*string) *DescribeCheckResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeCheckResultResponse) SetStatusCode(v int32) *DescribeCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCheckResultResponse) SetBody(v *DescribeCheckResultResponseBody) *DescribeCheckResultResponse {
	s.Body = v
	return s
}

func (s *DescribeCheckResultResponse) Validate() error {
	return dara.Validate(s)
}
