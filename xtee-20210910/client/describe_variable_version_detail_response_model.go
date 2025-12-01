// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableVersionDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVariableVersionDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVariableVersionDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVariableVersionDetailResponseBody) *DescribeVariableVersionDetailResponse
	GetBody() *DescribeVariableVersionDetailResponseBody
}

type DescribeVariableVersionDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVariableVersionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVariableVersionDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableVersionDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVariableVersionDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVariableVersionDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVariableVersionDetailResponse) GetBody() *DescribeVariableVersionDetailResponseBody {
	return s.Body
}

func (s *DescribeVariableVersionDetailResponse) SetHeaders(v map[string]*string) *DescribeVariableVersionDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVariableVersionDetailResponse) SetStatusCode(v int32) *DescribeVariableVersionDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVariableVersionDetailResponse) SetBody(v *DescribeVariableVersionDetailResponseBody) *DescribeVariableVersionDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeVariableVersionDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
