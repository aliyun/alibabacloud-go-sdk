// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRebootStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceRebootStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceRebootStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceRebootStatusResponseBody) *DescribeInstanceRebootStatusResponse
	GetBody() *DescribeInstanceRebootStatusResponseBody
}

type DescribeInstanceRebootStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceRebootStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceRebootStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRebootStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRebootStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceRebootStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceRebootStatusResponse) GetBody() *DescribeInstanceRebootStatusResponseBody {
	return s.Body
}

func (s *DescribeInstanceRebootStatusResponse) SetHeaders(v map[string]*string) *DescribeInstanceRebootStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceRebootStatusResponse) SetStatusCode(v int32) *DescribeInstanceRebootStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceRebootStatusResponse) SetBody(v *DescribeInstanceRebootStatusResponseBody) *DescribeInstanceRebootStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceRebootStatusResponse) Validate() error {
	return dara.Validate(s)
}
