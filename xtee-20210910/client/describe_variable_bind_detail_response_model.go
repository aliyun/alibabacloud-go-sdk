// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableBindDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVariableBindDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVariableBindDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVariableBindDetailResponseBody) *DescribeVariableBindDetailResponse
	GetBody() *DescribeVariableBindDetailResponseBody
}

type DescribeVariableBindDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVariableBindDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVariableBindDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableBindDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVariableBindDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVariableBindDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVariableBindDetailResponse) GetBody() *DescribeVariableBindDetailResponseBody {
	return s.Body
}

func (s *DescribeVariableBindDetailResponse) SetHeaders(v map[string]*string) *DescribeVariableBindDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVariableBindDetailResponse) SetStatusCode(v int32) *DescribeVariableBindDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVariableBindDetailResponse) SetBody(v *DescribeVariableBindDetailResponseBody) *DescribeVariableBindDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeVariableBindDetailResponse) Validate() error {
	return dara.Validate(s)
}
