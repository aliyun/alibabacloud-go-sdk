// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQueryVariablePageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQueryVariablePageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQueryVariablePageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQueryVariablePageListResponseBody) *DescribeQueryVariablePageListResponse
	GetBody() *DescribeQueryVariablePageListResponseBody
}

type DescribeQueryVariablePageListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQueryVariablePageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQueryVariablePageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryVariablePageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeQueryVariablePageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQueryVariablePageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQueryVariablePageListResponse) GetBody() *DescribeQueryVariablePageListResponseBody {
	return s.Body
}

func (s *DescribeQueryVariablePageListResponse) SetHeaders(v map[string]*string) *DescribeQueryVariablePageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeQueryVariablePageListResponse) SetStatusCode(v int32) *DescribeQueryVariablePageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQueryVariablePageListResponse) SetBody(v *DescribeQueryVariablePageListResponseBody) *DescribeQueryVariablePageListResponse {
	s.Body = v
	return s
}

func (s *DescribeQueryVariablePageListResponse) Validate() error {
	return dara.Validate(s)
}
