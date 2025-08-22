// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEncryptRoutineUidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEncryptRoutineUidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEncryptRoutineUidResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEncryptRoutineUidResponseBody) *DescribeEncryptRoutineUidResponse
	GetBody() *DescribeEncryptRoutineUidResponseBody
}

type DescribeEncryptRoutineUidResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEncryptRoutineUidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEncryptRoutineUidResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptRoutineUidResponse) GoString() string {
	return s.String()
}

func (s *DescribeEncryptRoutineUidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEncryptRoutineUidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEncryptRoutineUidResponse) GetBody() *DescribeEncryptRoutineUidResponseBody {
	return s.Body
}

func (s *DescribeEncryptRoutineUidResponse) SetHeaders(v map[string]*string) *DescribeEncryptRoutineUidResponse {
	s.Headers = v
	return s
}

func (s *DescribeEncryptRoutineUidResponse) SetStatusCode(v int32) *DescribeEncryptRoutineUidResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEncryptRoutineUidResponse) SetBody(v *DescribeEncryptRoutineUidResponseBody) *DescribeEncryptRoutineUidResponse {
	s.Body = v
	return s
}

func (s *DescribeEncryptRoutineUidResponse) Validate() error {
	return dara.Validate(s)
}
