// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventVariableListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventVariableListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventVariableListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventVariableListResponseBody) *DescribeEventVariableListResponse
	GetBody() *DescribeEventVariableListResponseBody
}

type DescribeEventVariableListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventVariableListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventVariableListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableListResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventVariableListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventVariableListResponse) GetBody() *DescribeEventVariableListResponseBody {
	return s.Body
}

func (s *DescribeEventVariableListResponse) SetHeaders(v map[string]*string) *DescribeEventVariableListResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventVariableListResponse) SetStatusCode(v int32) *DescribeEventVariableListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventVariableListResponse) SetBody(v *DescribeEventVariableListResponseBody) *DescribeEventVariableListResponse {
	s.Body = v
	return s
}

func (s *DescribeEventVariableListResponse) Validate() error {
	return dara.Validate(s)
}
