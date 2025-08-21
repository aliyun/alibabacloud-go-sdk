// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebRulesResponseBody) *DescribeWebRulesResponse
	GetBody() *DescribeWebRulesResponseBody
}

type DescribeWebRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebRulesResponse) GetBody() *DescribeWebRulesResponseBody {
	return s.Body
}

func (s *DescribeWebRulesResponse) SetHeaders(v map[string]*string) *DescribeWebRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebRulesResponse) SetStatusCode(v int32) *DescribeWebRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebRulesResponse) SetBody(v *DescribeWebRulesResponseBody) *DescribeWebRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeWebRulesResponse) Validate() error {
	return dara.Validate(s)
}
