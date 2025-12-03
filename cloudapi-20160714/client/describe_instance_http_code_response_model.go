// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceHttpCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceHttpCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceHttpCodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceHttpCodeResponseBody) *DescribeInstanceHttpCodeResponse
	GetBody() *DescribeInstanceHttpCodeResponseBody
}

type DescribeInstanceHttpCodeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceHttpCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceHttpCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHttpCodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHttpCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceHttpCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceHttpCodeResponse) GetBody() *DescribeInstanceHttpCodeResponseBody {
	return s.Body
}

func (s *DescribeInstanceHttpCodeResponse) SetHeaders(v map[string]*string) *DescribeInstanceHttpCodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceHttpCodeResponse) SetStatusCode(v int32) *DescribeInstanceHttpCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceHttpCodeResponse) SetBody(v *DescribeInstanceHttpCodeResponseBody) *DescribeInstanceHttpCodeResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceHttpCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
