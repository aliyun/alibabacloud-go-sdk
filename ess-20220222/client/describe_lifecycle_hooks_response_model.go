// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLifecycleHooksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLifecycleHooksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLifecycleHooksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLifecycleHooksResponseBody) *DescribeLifecycleHooksResponse
	GetBody() *DescribeLifecycleHooksResponseBody
}

type DescribeLifecycleHooksResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLifecycleHooksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLifecycleHooksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecycleHooksResponse) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleHooksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLifecycleHooksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLifecycleHooksResponse) GetBody() *DescribeLifecycleHooksResponseBody {
	return s.Body
}

func (s *DescribeLifecycleHooksResponse) SetHeaders(v map[string]*string) *DescribeLifecycleHooksResponse {
	s.Headers = v
	return s
}

func (s *DescribeLifecycleHooksResponse) SetStatusCode(v int32) *DescribeLifecycleHooksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLifecycleHooksResponse) SetBody(v *DescribeLifecycleHooksResponseBody) *DescribeLifecycleHooksResponse {
	s.Body = v
	return s
}

func (s *DescribeLifecycleHooksResponse) Validate() error {
	return dara.Validate(s)
}
