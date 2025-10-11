// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationAttributeResponseBody) *DescribeApplicationAttributeResponse
	GetBody() *DescribeApplicationAttributeResponseBody
}

type DescribeApplicationAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationAttributeResponse) GetBody() *DescribeApplicationAttributeResponseBody {
	return s.Body
}

func (s *DescribeApplicationAttributeResponse) SetHeaders(v map[string]*string) *DescribeApplicationAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationAttributeResponse) SetStatusCode(v int32) *DescribeApplicationAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationAttributeResponse) SetBody(v *DescribeApplicationAttributeResponseBody) *DescribeApplicationAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationAttributeResponse) Validate() error {
	return dara.Validate(s)
}
