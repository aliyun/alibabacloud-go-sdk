// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuthStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuthStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuthStatusResponseBody) *DescribeAuthStatusResponse
	GetBody() *DescribeAuthStatusResponseBody
}

type DescribeAuthStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuthStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuthStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuthStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuthStatusResponse) GetBody() *DescribeAuthStatusResponseBody {
	return s.Body
}

func (s *DescribeAuthStatusResponse) SetHeaders(v map[string]*string) *DescribeAuthStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuthStatusResponse) SetStatusCode(v int32) *DescribeAuthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuthStatusResponse) SetBody(v *DescribeAuthStatusResponseBody) *DescribeAuthStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeAuthStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
