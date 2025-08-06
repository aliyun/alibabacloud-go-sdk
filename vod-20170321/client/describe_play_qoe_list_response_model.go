// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayQoeListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlayQoeListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlayQoeListResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlayQoeListResponseBody) *DescribePlayQoeListResponse
	GetBody() *DescribePlayQoeListResponseBody
}

type DescribePlayQoeListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlayQoeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlayQoeListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayQoeListResponse) GoString() string {
	return s.String()
}

func (s *DescribePlayQoeListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlayQoeListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlayQoeListResponse) GetBody() *DescribePlayQoeListResponseBody {
	return s.Body
}

func (s *DescribePlayQoeListResponse) SetHeaders(v map[string]*string) *DescribePlayQoeListResponse {
	s.Headers = v
	return s
}

func (s *DescribePlayQoeListResponse) SetStatusCode(v int32) *DescribePlayQoeListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlayQoeListResponse) SetBody(v *DescribePlayQoeListResponseBody) *DescribePlayQoeListResponse {
	s.Body = v
	return s
}

func (s *DescribePlayQoeListResponse) Validate() error {
	return dara.Validate(s)
}
