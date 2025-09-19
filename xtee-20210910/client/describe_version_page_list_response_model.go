// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVersionPageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVersionPageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVersionPageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVersionPageListResponseBody) *DescribeVersionPageListResponse
	GetBody() *DescribeVersionPageListResponseBody
}

type DescribeVersionPageListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVersionPageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVersionPageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionPageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVersionPageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVersionPageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVersionPageListResponse) GetBody() *DescribeVersionPageListResponseBody {
	return s.Body
}

func (s *DescribeVersionPageListResponse) SetHeaders(v map[string]*string) *DescribeVersionPageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVersionPageListResponse) SetStatusCode(v int32) *DescribeVersionPageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVersionPageListResponse) SetBody(v *DescribeVersionPageListResponseBody) *DescribeVersionPageListResponse {
	s.Body = v
	return s
}

func (s *DescribeVersionPageListResponse) Validate() error {
	return dara.Validate(s)
}
