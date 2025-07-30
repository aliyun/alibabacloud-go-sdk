// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTDEStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceTDEStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceTDEStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceTDEStatusResponseBody) *DescribeInstanceTDEStatusResponse
	GetBody() *DescribeInstanceTDEStatusResponseBody
}

type DescribeInstanceTDEStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceTDEStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceTDEStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTDEStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTDEStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceTDEStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceTDEStatusResponse) GetBody() *DescribeInstanceTDEStatusResponseBody {
	return s.Body
}

func (s *DescribeInstanceTDEStatusResponse) SetHeaders(v map[string]*string) *DescribeInstanceTDEStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceTDEStatusResponse) SetStatusCode(v int32) *DescribeInstanceTDEStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceTDEStatusResponse) SetBody(v *DescribeInstanceTDEStatusResponseBody) *DescribeInstanceTDEStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceTDEStatusResponse) Validate() error {
	return dara.Validate(s)
}
