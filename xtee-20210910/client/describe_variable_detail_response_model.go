// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVariableDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVariableDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVariableDetailResponseBody) *DescribeVariableDetailResponse
	GetBody() *DescribeVariableDetailResponseBody
}

type DescribeVariableDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVariableDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVariableDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVariableDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVariableDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVariableDetailResponse) GetBody() *DescribeVariableDetailResponseBody {
	return s.Body
}

func (s *DescribeVariableDetailResponse) SetHeaders(v map[string]*string) *DescribeVariableDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVariableDetailResponse) SetStatusCode(v int32) *DescribeVariableDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVariableDetailResponse) SetBody(v *DescribeVariableDetailResponseBody) *DescribeVariableDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeVariableDetailResponse) Validate() error {
	return dara.Validate(s)
}
