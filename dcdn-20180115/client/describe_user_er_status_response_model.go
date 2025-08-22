// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserErStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserErStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserErStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserErStatusResponseBody) *DescribeUserErStatusResponse
	GetBody() *DescribeUserErStatusResponseBody
}

type DescribeUserErStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserErStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserErStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserErStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserErStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserErStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserErStatusResponse) GetBody() *DescribeUserErStatusResponseBody {
	return s.Body
}

func (s *DescribeUserErStatusResponse) SetHeaders(v map[string]*string) *DescribeUserErStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserErStatusResponse) SetStatusCode(v int32) *DescribeUserErStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserErStatusResponse) SetBody(v *DescribeUserErStatusResponseBody) *DescribeUserErStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeUserErStatusResponse) Validate() error {
	return dara.Validate(s)
}
