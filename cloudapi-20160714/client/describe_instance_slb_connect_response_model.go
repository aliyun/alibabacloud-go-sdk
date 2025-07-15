// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSlbConnectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceSlbConnectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceSlbConnectResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceSlbConnectResponseBody) *DescribeInstanceSlbConnectResponse
	GetBody() *DescribeInstanceSlbConnectResponseBody
}

type DescribeInstanceSlbConnectResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceSlbConnectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceSlbConnectResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSlbConnectResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSlbConnectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceSlbConnectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceSlbConnectResponse) GetBody() *DescribeInstanceSlbConnectResponseBody {
	return s.Body
}

func (s *DescribeInstanceSlbConnectResponse) SetHeaders(v map[string]*string) *DescribeInstanceSlbConnectResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSlbConnectResponse) SetStatusCode(v int32) *DescribeInstanceSlbConnectResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSlbConnectResponse) SetBody(v *DescribeInstanceSlbConnectResponseBody) *DescribeInstanceSlbConnectResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceSlbConnectResponse) Validate() error {
	return dara.Validate(s)
}
