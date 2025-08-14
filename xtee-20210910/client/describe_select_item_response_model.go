// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSelectItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSelectItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSelectItemResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSelectItemResponseBody) *DescribeSelectItemResponse
	GetBody() *DescribeSelectItemResponseBody
}

type DescribeSelectItemResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSelectItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSelectItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSelectItemResponse) GoString() string {
	return s.String()
}

func (s *DescribeSelectItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSelectItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSelectItemResponse) GetBody() *DescribeSelectItemResponseBody {
	return s.Body
}

func (s *DescribeSelectItemResponse) SetHeaders(v map[string]*string) *DescribeSelectItemResponse {
	s.Headers = v
	return s
}

func (s *DescribeSelectItemResponse) SetStatusCode(v int32) *DescribeSelectItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSelectItemResponse) SetBody(v *DescribeSelectItemResponseBody) *DescribeSelectItemResponse {
	s.Body = v
	return s
}

func (s *DescribeSelectItemResponse) Validate() error {
	return dara.Validate(s)
}
