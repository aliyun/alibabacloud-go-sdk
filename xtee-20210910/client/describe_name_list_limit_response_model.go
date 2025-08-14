// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNameListLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNameListLimitResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNameListLimitResponseBody) *DescribeNameListLimitResponse
	GetBody() *DescribeNameListLimitResponseBody
}

type DescribeNameListLimitResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNameListLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNameListLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListLimitResponse) GoString() string {
	return s.String()
}

func (s *DescribeNameListLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNameListLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNameListLimitResponse) GetBody() *DescribeNameListLimitResponseBody {
	return s.Body
}

func (s *DescribeNameListLimitResponse) SetHeaders(v map[string]*string) *DescribeNameListLimitResponse {
	s.Headers = v
	return s
}

func (s *DescribeNameListLimitResponse) SetStatusCode(v int32) *DescribeNameListLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNameListLimitResponse) SetBody(v *DescribeNameListLimitResponseBody) *DescribeNameListLimitResponse {
	s.Body = v
	return s
}

func (s *DescribeNameListLimitResponse) Validate() error {
	return dara.Validate(s)
}
