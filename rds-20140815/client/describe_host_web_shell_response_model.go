// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHostWebShellResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHostWebShellResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHostWebShellResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHostWebShellResponseBody) *DescribeHostWebShellResponse
	GetBody() *DescribeHostWebShellResponseBody
}

type DescribeHostWebShellResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHostWebShellResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHostWebShellResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostWebShellResponse) GoString() string {
	return s.String()
}

func (s *DescribeHostWebShellResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHostWebShellResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHostWebShellResponse) GetBody() *DescribeHostWebShellResponseBody {
	return s.Body
}

func (s *DescribeHostWebShellResponse) SetHeaders(v map[string]*string) *DescribeHostWebShellResponse {
	s.Headers = v
	return s
}

func (s *DescribeHostWebShellResponse) SetStatusCode(v int32) *DescribeHostWebShellResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHostWebShellResponse) SetBody(v *DescribeHostWebShellResponseBody) *DescribeHostWebShellResponse {
	s.Body = v
	return s
}

func (s *DescribeHostWebShellResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
