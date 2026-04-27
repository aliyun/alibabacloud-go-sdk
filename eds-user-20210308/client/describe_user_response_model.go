// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserResponseBody) *DescribeUserResponse
	GetBody() *DescribeUserResponseBody
}

type DescribeUserResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserResponse) GetBody() *DescribeUserResponseBody {
	return s.Body
}

func (s *DescribeUserResponse) SetHeaders(v map[string]*string) *DescribeUserResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserResponse) SetStatusCode(v int32) *DescribeUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserResponse) SetBody(v *DescribeUserResponseBody) *DescribeUserResponse {
	s.Body = v
	return s
}

func (s *DescribeUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
