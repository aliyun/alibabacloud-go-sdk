// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyUserItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertyUserItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertyUserItemResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertyUserItemResponseBody) *DescribePropertyUserItemResponse
	GetBody() *DescribePropertyUserItemResponseBody
}

type DescribePropertyUserItemResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertyUserItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertyUserItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyUserItemResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertyUserItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertyUserItemResponse) GetBody() *DescribePropertyUserItemResponseBody {
	return s.Body
}

func (s *DescribePropertyUserItemResponse) SetHeaders(v map[string]*string) *DescribePropertyUserItemResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyUserItemResponse) SetStatusCode(v int32) *DescribePropertyUserItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertyUserItemResponse) SetBody(v *DescribePropertyUserItemResponseBody) *DescribePropertyUserItemResponse {
	s.Body = v
	return s
}

func (s *DescribePropertyUserItemResponse) Validate() error {
	return dara.Validate(s)
}
