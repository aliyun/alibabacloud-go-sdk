// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSwitchNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceSwitchNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceSwitchNetworkResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceSwitchNetworkResponseBody) *DescribeInstanceSwitchNetworkResponse
	GetBody() *DescribeInstanceSwitchNetworkResponseBody
}

type DescribeInstanceSwitchNetworkResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceSwitchNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceSwitchNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSwitchNetworkResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceSwitchNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceSwitchNetworkResponse) GetBody() *DescribeInstanceSwitchNetworkResponseBody {
	return s.Body
}

func (s *DescribeInstanceSwitchNetworkResponse) SetHeaders(v map[string]*string) *DescribeInstanceSwitchNetworkResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponse) SetStatusCode(v int32) *DescribeInstanceSwitchNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponse) SetBody(v *DescribeInstanceSwitchNetworkResponseBody) *DescribeInstanceSwitchNetworkResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceSwitchNetworkResponse) Validate() error {
	return dara.Validate(s)
}
