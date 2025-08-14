// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInputFeildCountByEventCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInputFeildCountByEventCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInputFeildCountByEventCodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInputFeildCountByEventCodeResponseBody) *DescribeInputFeildCountByEventCodeResponse
	GetBody() *DescribeInputFeildCountByEventCodeResponseBody
}

type DescribeInputFeildCountByEventCodeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInputFeildCountByEventCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInputFeildCountByEventCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInputFeildCountByEventCodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInputFeildCountByEventCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInputFeildCountByEventCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInputFeildCountByEventCodeResponse) GetBody() *DescribeInputFeildCountByEventCodeResponseBody {
	return s.Body
}

func (s *DescribeInputFeildCountByEventCodeResponse) SetHeaders(v map[string]*string) *DescribeInputFeildCountByEventCodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInputFeildCountByEventCodeResponse) SetStatusCode(v int32) *DescribeInputFeildCountByEventCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInputFeildCountByEventCodeResponse) SetBody(v *DescribeInputFeildCountByEventCodeResponseBody) *DescribeInputFeildCountByEventCodeResponse {
	s.Body = v
	return s
}

func (s *DescribeInputFeildCountByEventCodeResponse) Validate() error {
	return dara.Validate(s)
}
