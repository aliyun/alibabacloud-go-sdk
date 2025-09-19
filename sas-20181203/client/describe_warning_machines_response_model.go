// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWarningMachinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWarningMachinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWarningMachinesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWarningMachinesResponseBody) *DescribeWarningMachinesResponse
	GetBody() *DescribeWarningMachinesResponseBody
}

type DescribeWarningMachinesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWarningMachinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWarningMachinesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarningMachinesResponse) GoString() string {
	return s.String()
}

func (s *DescribeWarningMachinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWarningMachinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWarningMachinesResponse) GetBody() *DescribeWarningMachinesResponseBody {
	return s.Body
}

func (s *DescribeWarningMachinesResponse) SetHeaders(v map[string]*string) *DescribeWarningMachinesResponse {
	s.Headers = v
	return s
}

func (s *DescribeWarningMachinesResponse) SetStatusCode(v int32) *DescribeWarningMachinesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWarningMachinesResponse) SetBody(v *DescribeWarningMachinesResponseBody) *DescribeWarningMachinesResponse {
	s.Body = v
	return s
}

func (s *DescribeWarningMachinesResponse) Validate() error {
	return dara.Validate(s)
}
