// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllRootVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAllRootVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAllRootVariableResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAllRootVariableResponseBody) *DescribeAllRootVariableResponse
	GetBody() *DescribeAllRootVariableResponseBody
}

type DescribeAllRootVariableResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAllRootVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAllRootVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllRootVariableResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllRootVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAllRootVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAllRootVariableResponse) GetBody() *DescribeAllRootVariableResponseBody {
	return s.Body
}

func (s *DescribeAllRootVariableResponse) SetHeaders(v map[string]*string) *DescribeAllRootVariableResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllRootVariableResponse) SetStatusCode(v int32) *DescribeAllRootVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAllRootVariableResponse) SetBody(v *DescribeAllRootVariableResponseBody) *DescribeAllRootVariableResponse {
	s.Body = v
	return s
}

func (s *DescribeAllRootVariableResponse) Validate() error {
	return dara.Validate(s)
}
