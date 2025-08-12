// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContactGroupListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContactGroupListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContactGroupListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContactGroupListResponseBody) *DescribeContactGroupListResponse
	GetBody() *DescribeContactGroupListResponseBody
}

type DescribeContactGroupListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContactGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContactGroupListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactGroupListResponse) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContactGroupListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContactGroupListResponse) GetBody() *DescribeContactGroupListResponseBody {
	return s.Body
}

func (s *DescribeContactGroupListResponse) SetHeaders(v map[string]*string) *DescribeContactGroupListResponse {
	s.Headers = v
	return s
}

func (s *DescribeContactGroupListResponse) SetStatusCode(v int32) *DescribeContactGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContactGroupListResponse) SetBody(v *DescribeContactGroupListResponseBody) *DescribeContactGroupListResponse {
	s.Body = v
	return s
}

func (s *DescribeContactGroupListResponse) Validate() error {
	return dara.Validate(s)
}
