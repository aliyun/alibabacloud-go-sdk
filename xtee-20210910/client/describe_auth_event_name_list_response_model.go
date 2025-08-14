// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthEventNameListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuthEventNameListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuthEventNameListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuthEventNameListResponseBody) *DescribeAuthEventNameListResponse
	GetBody() *DescribeAuthEventNameListResponseBody
}

type DescribeAuthEventNameListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuthEventNameListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuthEventNameListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthEventNameListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuthEventNameListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuthEventNameListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuthEventNameListResponse) GetBody() *DescribeAuthEventNameListResponseBody {
	return s.Body
}

func (s *DescribeAuthEventNameListResponse) SetHeaders(v map[string]*string) *DescribeAuthEventNameListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuthEventNameListResponse) SetStatusCode(v int32) *DescribeAuthEventNameListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuthEventNameListResponse) SetBody(v *DescribeAuthEventNameListResponseBody) *DescribeAuthEventNameListResponse {
	s.Body = v
	return s
}

func (s *DescribeAuthEventNameListResponse) Validate() error {
	return dara.Validate(s)
}
