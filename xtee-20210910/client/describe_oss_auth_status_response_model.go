// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssAuthStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOssAuthStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOssAuthStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOssAuthStatusResponseBody) *DescribeOssAuthStatusResponse
	GetBody() *DescribeOssAuthStatusResponseBody
}

type DescribeOssAuthStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssAuthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssAuthStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssAuthStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssAuthStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOssAuthStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOssAuthStatusResponse) GetBody() *DescribeOssAuthStatusResponseBody {
	return s.Body
}

func (s *DescribeOssAuthStatusResponse) SetHeaders(v map[string]*string) *DescribeOssAuthStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssAuthStatusResponse) SetStatusCode(v int32) *DescribeOssAuthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssAuthStatusResponse) SetBody(v *DescribeOssAuthStatusResponseBody) *DescribeOssAuthStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeOssAuthStatusResponse) Validate() error {
	return dara.Validate(s)
}
