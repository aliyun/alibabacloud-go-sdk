// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTaskInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTaskInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTaskInfoResponseBody) *DescribeTaskInfoResponse
	GetBody() *DescribeTaskInfoResponseBody
}

type DescribeTaskInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTaskInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTaskInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTaskInfoResponse) GetBody() *DescribeTaskInfoResponseBody {
	return s.Body
}

func (s *DescribeTaskInfoResponse) SetHeaders(v map[string]*string) *DescribeTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskInfoResponse) SetStatusCode(v int32) *DescribeTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTaskInfoResponse) SetBody(v *DescribeTaskInfoResponseBody) *DescribeTaskInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeTaskInfoResponse) Validate() error {
	return dara.Validate(s)
}
