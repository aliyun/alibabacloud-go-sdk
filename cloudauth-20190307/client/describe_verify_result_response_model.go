// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVerifyResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVerifyResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVerifyResultResponseBody) *DescribeVerifyResultResponse
	GetBody() *DescribeVerifyResultResponseBody
}

type DescribeVerifyResultResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifyResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVerifyResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVerifyResultResponse) GetBody() *DescribeVerifyResultResponseBody {
	return s.Body
}

func (s *DescribeVerifyResultResponse) SetHeaders(v map[string]*string) *DescribeVerifyResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyResultResponse) SetStatusCode(v int32) *DescribeVerifyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifyResultResponse) SetBody(v *DescribeVerifyResultResponseBody) *DescribeVerifyResultResponse {
	s.Body = v
	return s
}

func (s *DescribeVerifyResultResponse) Validate() error {
	return dara.Validate(s)
}
