// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSystemLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSystemLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSystemLogResponseBody) *DescribeSystemLogResponse
	GetBody() *DescribeSystemLogResponseBody
}

type DescribeSystemLogResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSystemLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSystemLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeSystemLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSystemLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSystemLogResponse) GetBody() *DescribeSystemLogResponseBody {
	return s.Body
}

func (s *DescribeSystemLogResponse) SetHeaders(v map[string]*string) *DescribeSystemLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeSystemLogResponse) SetStatusCode(v int32) *DescribeSystemLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSystemLogResponse) SetBody(v *DescribeSystemLogResponseBody) *DescribeSystemLogResponse {
	s.Body = v
	return s
}

func (s *DescribeSystemLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
