// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQueryVariableDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQueryVariableDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQueryVariableDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQueryVariableDetailResponseBody) *DescribeQueryVariableDetailResponse
	GetBody() *DescribeQueryVariableDetailResponseBody
}

type DescribeQueryVariableDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQueryVariableDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQueryVariableDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryVariableDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeQueryVariableDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQueryVariableDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQueryVariableDetailResponse) GetBody() *DescribeQueryVariableDetailResponseBody {
	return s.Body
}

func (s *DescribeQueryVariableDetailResponse) SetHeaders(v map[string]*string) *DescribeQueryVariableDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeQueryVariableDetailResponse) SetStatusCode(v int32) *DescribeQueryVariableDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQueryVariableDetailResponse) SetBody(v *DescribeQueryVariableDetailResponseBody) *DescribeQueryVariableDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeQueryVariableDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
