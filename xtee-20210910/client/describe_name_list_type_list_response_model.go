// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListTypeListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNameListTypeListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNameListTypeListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNameListTypeListResponseBody) *DescribeNameListTypeListResponse
	GetBody() *DescribeNameListTypeListResponseBody
}

type DescribeNameListTypeListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNameListTypeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNameListTypeListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListTypeListResponse) GoString() string {
	return s.String()
}

func (s *DescribeNameListTypeListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNameListTypeListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNameListTypeListResponse) GetBody() *DescribeNameListTypeListResponseBody {
	return s.Body
}

func (s *DescribeNameListTypeListResponse) SetHeaders(v map[string]*string) *DescribeNameListTypeListResponse {
	s.Headers = v
	return s
}

func (s *DescribeNameListTypeListResponse) SetStatusCode(v int32) *DescribeNameListTypeListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNameListTypeListResponse) SetBody(v *DescribeNameListTypeListResponseBody) *DescribeNameListTypeListResponse {
	s.Body = v
	return s
}

func (s *DescribeNameListTypeListResponse) Validate() error {
	return dara.Validate(s)
}
