// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInitializeVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInitializeVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInitializeVariableResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInitializeVariableResponseBody) *DescribeDBInitializeVariableResponse
	GetBody() *DescribeDBInitializeVariableResponseBody
}

type DescribeDBInitializeVariableResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInitializeVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInitializeVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInitializeVariableResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInitializeVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInitializeVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInitializeVariableResponse) GetBody() *DescribeDBInitializeVariableResponseBody {
	return s.Body
}

func (s *DescribeDBInitializeVariableResponse) SetHeaders(v map[string]*string) *DescribeDBInitializeVariableResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInitializeVariableResponse) SetStatusCode(v int32) *DescribeDBInitializeVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInitializeVariableResponse) SetBody(v *DescribeDBInitializeVariableResponseBody) *DescribeDBInitializeVariableResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInitializeVariableResponse) Validate() error {
	return dara.Validate(s)
}
