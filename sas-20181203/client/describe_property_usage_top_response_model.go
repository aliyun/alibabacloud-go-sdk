// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyUsageTopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertyUsageTopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertyUsageTopResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertyUsageTopResponseBody) *DescribePropertyUsageTopResponse
	GetBody() *DescribePropertyUsageTopResponseBody
}

type DescribePropertyUsageTopResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertyUsageTopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertyUsageTopResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyUsageTopResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyUsageTopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertyUsageTopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertyUsageTopResponse) GetBody() *DescribePropertyUsageTopResponseBody {
	return s.Body
}

func (s *DescribePropertyUsageTopResponse) SetHeaders(v map[string]*string) *DescribePropertyUsageTopResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyUsageTopResponse) SetStatusCode(v int32) *DescribePropertyUsageTopResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertyUsageTopResponse) SetBody(v *DescribePropertyUsageTopResponseBody) *DescribePropertyUsageTopResponse {
	s.Body = v
	return s
}

func (s *DescribePropertyUsageTopResponse) Validate() error {
	return dara.Validate(s)
}
