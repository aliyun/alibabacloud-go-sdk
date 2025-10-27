// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMapRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMapRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMapRunResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMapRunResponseBody) *DescribeMapRunResponse
	GetBody() *DescribeMapRunResponseBody
}

type DescribeMapRunResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMapRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMapRunResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMapRunResponse) GoString() string {
	return s.String()
}

func (s *DescribeMapRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMapRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMapRunResponse) GetBody() *DescribeMapRunResponseBody {
	return s.Body
}

func (s *DescribeMapRunResponse) SetHeaders(v map[string]*string) *DescribeMapRunResponse {
	s.Headers = v
	return s
}

func (s *DescribeMapRunResponse) SetStatusCode(v int32) *DescribeMapRunResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMapRunResponse) SetBody(v *DescribeMapRunResponseBody) *DescribeMapRunResponse {
	s.Body = v
	return s
}

func (s *DescribeMapRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
