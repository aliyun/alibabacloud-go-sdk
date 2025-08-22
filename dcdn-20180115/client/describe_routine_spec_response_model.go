// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoutineSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRoutineSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRoutineSpecResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRoutineSpecResponseBody) *DescribeRoutineSpecResponse
	GetBody() *DescribeRoutineSpecResponseBody
}

type DescribeRoutineSpecResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRoutineSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRoutineSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoutineSpecResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoutineSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRoutineSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRoutineSpecResponse) GetBody() *DescribeRoutineSpecResponseBody {
	return s.Body
}

func (s *DescribeRoutineSpecResponse) SetHeaders(v map[string]*string) *DescribeRoutineSpecResponse {
	s.Headers = v
	return s
}

func (s *DescribeRoutineSpecResponse) SetStatusCode(v int32) *DescribeRoutineSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRoutineSpecResponse) SetBody(v *DescribeRoutineSpecResponseBody) *DescribeRoutineSpecResponse {
	s.Body = v
	return s
}

func (s *DescribeRoutineSpecResponse) Validate() error {
	return dara.Validate(s)
}
