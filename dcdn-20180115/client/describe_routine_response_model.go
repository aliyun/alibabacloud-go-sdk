// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoutineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRoutineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRoutineResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRoutineResponseBody) *DescribeRoutineResponse
	GetBody() *DescribeRoutineResponseBody
}

type DescribeRoutineResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRoutineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRoutineResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoutineResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoutineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRoutineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRoutineResponse) GetBody() *DescribeRoutineResponseBody {
	return s.Body
}

func (s *DescribeRoutineResponse) SetHeaders(v map[string]*string) *DescribeRoutineResponse {
	s.Headers = v
	return s
}

func (s *DescribeRoutineResponse) SetStatusCode(v int32) *DescribeRoutineResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRoutineResponse) SetBody(v *DescribeRoutineResponseBody) *DescribeRoutineResponse {
	s.Body = v
	return s
}

func (s *DescribeRoutineResponse) Validate() error {
	return dara.Validate(s)
}
