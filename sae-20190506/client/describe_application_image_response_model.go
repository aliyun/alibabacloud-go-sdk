// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationImageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationImageResponseBody) *DescribeApplicationImageResponse
	GetBody() *DescribeApplicationImageResponseBody
}

type DescribeApplicationImageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationImageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationImageResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationImageResponse) GetBody() *DescribeApplicationImageResponseBody {
	return s.Body
}

func (s *DescribeApplicationImageResponse) SetHeaders(v map[string]*string) *DescribeApplicationImageResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationImageResponse) SetStatusCode(v int32) *DescribeApplicationImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationImageResponse) SetBody(v *DescribeApplicationImageResponseBody) *DescribeApplicationImageResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationImageResponse) Validate() error {
	return dara.Validate(s)
}
