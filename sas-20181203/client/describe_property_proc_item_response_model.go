// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyProcItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertyProcItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertyProcItemResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertyProcItemResponseBody) *DescribePropertyProcItemResponse
	GetBody() *DescribePropertyProcItemResponseBody
}

type DescribePropertyProcItemResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertyProcItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertyProcItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyProcItemResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertyProcItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertyProcItemResponse) GetBody() *DescribePropertyProcItemResponseBody {
	return s.Body
}

func (s *DescribePropertyProcItemResponse) SetHeaders(v map[string]*string) *DescribePropertyProcItemResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyProcItemResponse) SetStatusCode(v int32) *DescribePropertyProcItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertyProcItemResponse) SetBody(v *DescribePropertyProcItemResponseBody) *DescribePropertyProcItemResponse {
	s.Body = v
	return s
}

func (s *DescribePropertyProcItemResponse) Validate() error {
	return dara.Validate(s)
}
