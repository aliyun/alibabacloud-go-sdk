// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoutineUserInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRoutineUserInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRoutineUserInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRoutineUserInfoResponseBody) *DescribeRoutineUserInfoResponse
	GetBody() *DescribeRoutineUserInfoResponseBody
}

type DescribeRoutineUserInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRoutineUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRoutineUserInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoutineUserInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeRoutineUserInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRoutineUserInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRoutineUserInfoResponse) GetBody() *DescribeRoutineUserInfoResponseBody {
	return s.Body
}

func (s *DescribeRoutineUserInfoResponse) SetHeaders(v map[string]*string) *DescribeRoutineUserInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeRoutineUserInfoResponse) SetStatusCode(v int32) *DescribeRoutineUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRoutineUserInfoResponse) SetBody(v *DescribeRoutineUserInfoResponseBody) *DescribeRoutineUserInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeRoutineUserInfoResponse) Validate() error {
	return dara.Validate(s)
}
