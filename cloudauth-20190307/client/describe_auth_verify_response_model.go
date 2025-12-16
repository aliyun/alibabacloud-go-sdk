// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuthVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuthVerifyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuthVerifyResponseBody) *DescribeAuthVerifyResponse
	GetBody() *DescribeAuthVerifyResponseBody
}

type DescribeAuthVerifyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuthVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuthVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthVerifyResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuthVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuthVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuthVerifyResponse) GetBody() *DescribeAuthVerifyResponseBody {
	return s.Body
}

func (s *DescribeAuthVerifyResponse) SetHeaders(v map[string]*string) *DescribeAuthVerifyResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuthVerifyResponse) SetStatusCode(v int32) *DescribeAuthVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuthVerifyResponse) SetBody(v *DescribeAuthVerifyResponseBody) *DescribeAuthVerifyResponse {
	s.Body = v
	return s
}

func (s *DescribeAuthVerifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
