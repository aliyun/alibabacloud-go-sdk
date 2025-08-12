// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductResourceTagKeyListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProductResourceTagKeyListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProductResourceTagKeyListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProductResourceTagKeyListResponseBody) *DescribeProductResourceTagKeyListResponse
	GetBody() *DescribeProductResourceTagKeyListResponseBody
}

type DescribeProductResourceTagKeyListResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProductResourceTagKeyListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProductResourceTagKeyListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResourceTagKeyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeProductResourceTagKeyListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProductResourceTagKeyListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProductResourceTagKeyListResponse) GetBody() *DescribeProductResourceTagKeyListResponseBody {
	return s.Body
}

func (s *DescribeProductResourceTagKeyListResponse) SetHeaders(v map[string]*string) *DescribeProductResourceTagKeyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeProductResourceTagKeyListResponse) SetStatusCode(v int32) *DescribeProductResourceTagKeyListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProductResourceTagKeyListResponse) SetBody(v *DescribeProductResourceTagKeyListResponseBody) *DescribeProductResourceTagKeyListResponse {
	s.Body = v
	return s
}

func (s *DescribeProductResourceTagKeyListResponse) Validate() error {
	return dara.Validate(s)
}
