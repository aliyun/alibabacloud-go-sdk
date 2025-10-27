// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckWarningMachinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCheckWarningMachinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCheckWarningMachinesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCheckWarningMachinesResponseBody) *DescribeCheckWarningMachinesResponse
	GetBody() *DescribeCheckWarningMachinesResponseBody
}

type DescribeCheckWarningMachinesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCheckWarningMachinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCheckWarningMachinesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningMachinesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningMachinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCheckWarningMachinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCheckWarningMachinesResponse) GetBody() *DescribeCheckWarningMachinesResponseBody {
	return s.Body
}

func (s *DescribeCheckWarningMachinesResponse) SetHeaders(v map[string]*string) *DescribeCheckWarningMachinesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCheckWarningMachinesResponse) SetStatusCode(v int32) *DescribeCheckWarningMachinesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCheckWarningMachinesResponse) SetBody(v *DescribeCheckWarningMachinesResponseBody) *DescribeCheckWarningMachinesResponse {
	s.Body = v
	return s
}

func (s *DescribeCheckWarningMachinesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
