// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvCustomJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnvCustomJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnvCustomJobResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnvCustomJobResponseBody) *DescribeEnvCustomJobResponse
	GetBody() *DescribeEnvCustomJobResponseBody
}

type DescribeEnvCustomJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnvCustomJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnvCustomJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvCustomJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnvCustomJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnvCustomJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnvCustomJobResponse) GetBody() *DescribeEnvCustomJobResponseBody {
	return s.Body
}

func (s *DescribeEnvCustomJobResponse) SetHeaders(v map[string]*string) *DescribeEnvCustomJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnvCustomJobResponse) SetStatusCode(v int32) *DescribeEnvCustomJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnvCustomJobResponse) SetBody(v *DescribeEnvCustomJobResponseBody) *DescribeEnvCustomJobResponse {
	s.Body = v
	return s
}

func (s *DescribeEnvCustomJobResponse) Validate() error {
	return dara.Validate(s)
}
