// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventBaseInfoByEventCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventBaseInfoByEventCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventBaseInfoByEventCodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventBaseInfoByEventCodeResponseBody) *DescribeEventBaseInfoByEventCodeResponse
	GetBody() *DescribeEventBaseInfoByEventCodeResponseBody
}

type DescribeEventBaseInfoByEventCodeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventBaseInfoByEventCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventBaseInfoByEventCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventBaseInfoByEventCodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventBaseInfoByEventCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventBaseInfoByEventCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventBaseInfoByEventCodeResponse) GetBody() *DescribeEventBaseInfoByEventCodeResponseBody {
	return s.Body
}

func (s *DescribeEventBaseInfoByEventCodeResponse) SetHeaders(v map[string]*string) *DescribeEventBaseInfoByEventCodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponse) SetStatusCode(v int32) *DescribeEventBaseInfoByEventCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponse) SetBody(v *DescribeEventBaseInfoByEventCodeResponseBody) *DescribeEventBaseInfoByEventCodeResponse {
	s.Body = v
	return s
}

func (s *DescribeEventBaseInfoByEventCodeResponse) Validate() error {
	return dara.Validate(s)
}
