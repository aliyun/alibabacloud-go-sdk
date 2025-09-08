// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWafScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWafScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWafScopeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWafScopeResponseBody) *DescribeWafScopeResponse
	GetBody() *DescribeWafScopeResponseBody
}

type DescribeWafScopeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWafScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWafScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWafScopeResponse) GoString() string {
	return s.String()
}

func (s *DescribeWafScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWafScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWafScopeResponse) GetBody() *DescribeWafScopeResponseBody {
	return s.Body
}

func (s *DescribeWafScopeResponse) SetHeaders(v map[string]*string) *DescribeWafScopeResponse {
	s.Headers = v
	return s
}

func (s *DescribeWafScopeResponse) SetStatusCode(v int32) *DescribeWafScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWafScopeResponse) SetBody(v *DescribeWafScopeResponseBody) *DescribeWafScopeResponse {
	s.Body = v
	return s
}

func (s *DescribeWafScopeResponse) Validate() error {
	return dara.Validate(s)
}
