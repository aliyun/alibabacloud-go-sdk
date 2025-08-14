// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableFeeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVariableFeeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVariableFeeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVariableFeeResponseBody) *DescribeVariableFeeResponse
	GetBody() *DescribeVariableFeeResponseBody
}

type DescribeVariableFeeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVariableFeeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVariableFeeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableFeeResponse) GoString() string {
	return s.String()
}

func (s *DescribeVariableFeeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVariableFeeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVariableFeeResponse) GetBody() *DescribeVariableFeeResponseBody {
	return s.Body
}

func (s *DescribeVariableFeeResponse) SetHeaders(v map[string]*string) *DescribeVariableFeeResponse {
	s.Headers = v
	return s
}

func (s *DescribeVariableFeeResponse) SetStatusCode(v int32) *DescribeVariableFeeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVariableFeeResponse) SetBody(v *DescribeVariableFeeResponseBody) *DescribeVariableFeeResponse {
	s.Body = v
	return s
}

func (s *DescribeVariableFeeResponse) Validate() error {
	return dara.Validate(s)
}
