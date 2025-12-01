// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserStatusResponseBody) *DescribeUserStatusResponse
	GetBody() *DescribeUserStatusResponseBody
}

type DescribeUserStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserStatusResponse) GetBody() *DescribeUserStatusResponseBody {
	return s.Body
}

func (s *DescribeUserStatusResponse) SetHeaders(v map[string]*string) *DescribeUserStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserStatusResponse) SetStatusCode(v int32) *DescribeUserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserStatusResponse) SetBody(v *DescribeUserStatusResponseBody) *DescribeUserStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeUserStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
