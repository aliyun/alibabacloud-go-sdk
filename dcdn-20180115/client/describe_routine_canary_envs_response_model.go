// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoutineCanaryEnvsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRoutineCanaryEnvsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRoutineCanaryEnvsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRoutineCanaryEnvsResponseBody) *DescribeRoutineCanaryEnvsResponse
	GetBody() *DescribeRoutineCanaryEnvsResponseBody
}

type DescribeRoutineCanaryEnvsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRoutineCanaryEnvsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRoutineCanaryEnvsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoutineCanaryEnvsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoutineCanaryEnvsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRoutineCanaryEnvsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRoutineCanaryEnvsResponse) GetBody() *DescribeRoutineCanaryEnvsResponseBody {
	return s.Body
}

func (s *DescribeRoutineCanaryEnvsResponse) SetHeaders(v map[string]*string) *DescribeRoutineCanaryEnvsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRoutineCanaryEnvsResponse) SetStatusCode(v int32) *DescribeRoutineCanaryEnvsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRoutineCanaryEnvsResponse) SetBody(v *DescribeRoutineCanaryEnvsResponseBody) *DescribeRoutineCanaryEnvsResponse {
	s.Body = v
	return s
}

func (s *DescribeRoutineCanaryEnvsResponse) Validate() error {
	return dara.Validate(s)
}
