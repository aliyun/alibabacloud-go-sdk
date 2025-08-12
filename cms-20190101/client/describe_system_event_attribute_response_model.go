// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemEventAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSystemEventAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSystemEventAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSystemEventAttributeResponseBody) *DescribeSystemEventAttributeResponse
	GetBody() *DescribeSystemEventAttributeResponseBody
}

type DescribeSystemEventAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSystemEventAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSystemEventAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemEventAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSystemEventAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSystemEventAttributeResponse) GetBody() *DescribeSystemEventAttributeResponseBody {
	return s.Body
}

func (s *DescribeSystemEventAttributeResponse) SetHeaders(v map[string]*string) *DescribeSystemEventAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeSystemEventAttributeResponse) SetStatusCode(v int32) *DescribeSystemEventAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSystemEventAttributeResponse) SetBody(v *DescribeSystemEventAttributeResponseBody) *DescribeSystemEventAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeSystemEventAttributeResponse) Validate() error {
	return dara.Validate(s)
}
