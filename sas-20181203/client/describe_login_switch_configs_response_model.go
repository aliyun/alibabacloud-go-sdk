// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoginSwitchConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLoginSwitchConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLoginSwitchConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLoginSwitchConfigsResponseBody) *DescribeLoginSwitchConfigsResponse
	GetBody() *DescribeLoginSwitchConfigsResponseBody
}

type DescribeLoginSwitchConfigsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLoginSwitchConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLoginSwitchConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoginSwitchConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoginSwitchConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLoginSwitchConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLoginSwitchConfigsResponse) GetBody() *DescribeLoginSwitchConfigsResponseBody {
	return s.Body
}

func (s *DescribeLoginSwitchConfigsResponse) SetHeaders(v map[string]*string) *DescribeLoginSwitchConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoginSwitchConfigsResponse) SetStatusCode(v int32) *DescribeLoginSwitchConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoginSwitchConfigsResponse) SetBody(v *DescribeLoginSwitchConfigsResponseBody) *DescribeLoginSwitchConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeLoginSwitchConfigsResponse) Validate() error {
	return dara.Validate(s)
}
