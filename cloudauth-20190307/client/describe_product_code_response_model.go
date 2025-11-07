// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProductCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProductCodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProductCodeResponseBody) *DescribeProductCodeResponse
	GetBody() *DescribeProductCodeResponseBody
}

type DescribeProductCodeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProductCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProductCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductCodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeProductCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProductCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProductCodeResponse) GetBody() *DescribeProductCodeResponseBody {
	return s.Body
}

func (s *DescribeProductCodeResponse) SetHeaders(v map[string]*string) *DescribeProductCodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeProductCodeResponse) SetStatusCode(v int32) *DescribeProductCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProductCodeResponse) SetBody(v *DescribeProductCodeResponseBody) *DescribeProductCodeResponse {
	s.Body = v
	return s
}

func (s *DescribeProductCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
