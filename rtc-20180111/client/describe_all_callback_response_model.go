// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAllCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAllCallbackResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAllCallbackResponseBody) *DescribeAllCallbackResponse
	GetBody() *DescribeAllCallbackResponseBody
}

type DescribeAllCallbackResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAllCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAllCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllCallbackResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAllCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAllCallbackResponse) GetBody() *DescribeAllCallbackResponseBody {
	return s.Body
}

func (s *DescribeAllCallbackResponse) SetHeaders(v map[string]*string) *DescribeAllCallbackResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllCallbackResponse) SetStatusCode(v int32) *DescribeAllCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAllCallbackResponse) SetBody(v *DescribeAllCallbackResponseBody) *DescribeAllCallbackResponse {
	s.Body = v
	return s
}

func (s *DescribeAllCallbackResponse) Validate() error {
	return dara.Validate(s)
}
