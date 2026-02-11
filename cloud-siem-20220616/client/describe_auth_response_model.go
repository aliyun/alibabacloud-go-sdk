// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuthResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuthResponseBody) *DescribeAuthResponse
	GetBody() *DescribeAuthResponseBody
}

type DescribeAuthResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuthResponse) GetBody() *DescribeAuthResponseBody {
	return s.Body
}

func (s *DescribeAuthResponse) SetHeaders(v map[string]*string) *DescribeAuthResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuthResponse) SetStatusCode(v int32) *DescribeAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuthResponse) SetBody(v *DescribeAuthResponseBody) *DescribeAuthResponse {
	s.Body = v
	return s
}

func (s *DescribeAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
