// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKernelVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKernelVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKernelVersionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKernelVersionResponseBody) *DescribeKernelVersionResponse
	GetBody() *DescribeKernelVersionResponseBody
}

type DescribeKernelVersionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKernelVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKernelVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKernelVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeKernelVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKernelVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKernelVersionResponse) GetBody() *DescribeKernelVersionResponseBody {
	return s.Body
}

func (s *DescribeKernelVersionResponse) SetHeaders(v map[string]*string) *DescribeKernelVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeKernelVersionResponse) SetStatusCode(v int32) *DescribeKernelVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKernelVersionResponse) SetBody(v *DescribeKernelVersionResponseBody) *DescribeKernelVersionResponse {
	s.Body = v
	return s
}

func (s *DescribeKernelVersionResponse) Validate() error {
	return dara.Validate(s)
}
