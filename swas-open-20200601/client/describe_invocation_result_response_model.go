// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvocationResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInvocationResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInvocationResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInvocationResultResponseBody) *DescribeInvocationResultResponse
	GetBody() *DescribeInvocationResultResponseBody
}

type DescribeInvocationResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvocationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvocationResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInvocationResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInvocationResultResponse) GetBody() *DescribeInvocationResultResponseBody {
	return s.Body
}

func (s *DescribeInvocationResultResponse) SetHeaders(v map[string]*string) *DescribeInvocationResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvocationResultResponse) SetStatusCode(v int32) *DescribeInvocationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvocationResultResponse) SetBody(v *DescribeInvocationResultResponseBody) *DescribeInvocationResultResponse {
	s.Body = v
	return s
}

func (s *DescribeInvocationResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
