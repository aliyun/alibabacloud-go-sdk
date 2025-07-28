// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserApiRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserApiRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserApiRequestResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserApiRequestResponseBody) *DescribeUserApiRequestResponse
	GetBody() *DescribeUserApiRequestResponseBody
}

type DescribeUserApiRequestResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserApiRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserApiRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserApiRequestResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserApiRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserApiRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserApiRequestResponse) GetBody() *DescribeUserApiRequestResponseBody {
	return s.Body
}

func (s *DescribeUserApiRequestResponse) SetHeaders(v map[string]*string) *DescribeUserApiRequestResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserApiRequestResponse) SetStatusCode(v int32) *DescribeUserApiRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserApiRequestResponse) SetBody(v *DescribeUserApiRequestResponseBody) *DescribeUserApiRequestResponse {
	s.Body = v
	return s
}

func (s *DescribeUserApiRequestResponse) Validate() error {
	return dara.Validate(s)
}
