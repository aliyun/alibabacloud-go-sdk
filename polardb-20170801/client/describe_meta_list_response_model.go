// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetaListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMetaListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMetaListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMetaListResponseBody) *DescribeMetaListResponse
	GetBody() *DescribeMetaListResponseBody
}

type DescribeMetaListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetaListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetaListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetaListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMetaListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMetaListResponse) GetBody() *DescribeMetaListResponseBody {
	return s.Body
}

func (s *DescribeMetaListResponse) SetHeaders(v map[string]*string) *DescribeMetaListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetaListResponse) SetStatusCode(v int32) *DescribeMetaListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetaListResponse) SetBody(v *DescribeMetaListResponseBody) *DescribeMetaListResponse {
	s.Body = v
	return s
}

func (s *DescribeMetaListResponse) Validate() error {
	return dara.Validate(s)
}
