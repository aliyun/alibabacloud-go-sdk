// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListPageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNameListPageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNameListPageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNameListPageListResponseBody) *DescribeNameListPageListResponse
	GetBody() *DescribeNameListPageListResponseBody
}

type DescribeNameListPageListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNameListPageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNameListPageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListPageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeNameListPageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNameListPageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNameListPageListResponse) GetBody() *DescribeNameListPageListResponseBody {
	return s.Body
}

func (s *DescribeNameListPageListResponse) SetHeaders(v map[string]*string) *DescribeNameListPageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeNameListPageListResponse) SetStatusCode(v int32) *DescribeNameListPageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNameListPageListResponse) SetBody(v *DescribeNameListPageListResponseBody) *DescribeNameListPageListResponse {
	s.Body = v
	return s
}

func (s *DescribeNameListPageListResponse) Validate() error {
	return dara.Validate(s)
}
