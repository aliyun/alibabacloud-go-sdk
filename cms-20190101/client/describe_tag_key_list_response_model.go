// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagKeyListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagKeyListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagKeyListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagKeyListResponseBody) *DescribeTagKeyListResponse
	GetBody() *DescribeTagKeyListResponseBody
}

type DescribeTagKeyListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagKeyListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagKeyListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagKeyListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagKeyListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagKeyListResponse) GetBody() *DescribeTagKeyListResponseBody {
	return s.Body
}

func (s *DescribeTagKeyListResponse) SetHeaders(v map[string]*string) *DescribeTagKeyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagKeyListResponse) SetStatusCode(v int32) *DescribeTagKeyListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagKeyListResponse) SetBody(v *DescribeTagKeyListResponseBody) *DescribeTagKeyListResponse {
	s.Body = v
	return s
}

func (s *DescribeTagKeyListResponse) Validate() error {
	return dara.Validate(s)
}
