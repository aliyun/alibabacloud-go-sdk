// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDropPacketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceDropPacketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceDropPacketResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceDropPacketResponseBody) *DescribeInstanceDropPacketResponse
	GetBody() *DescribeInstanceDropPacketResponseBody
}

type DescribeInstanceDropPacketResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceDropPacketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceDropPacketResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDropPacketResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDropPacketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceDropPacketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceDropPacketResponse) GetBody() *DescribeInstanceDropPacketResponseBody {
	return s.Body
}

func (s *DescribeInstanceDropPacketResponse) SetHeaders(v map[string]*string) *DescribeInstanceDropPacketResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceDropPacketResponse) SetStatusCode(v int32) *DescribeInstanceDropPacketResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceDropPacketResponse) SetBody(v *DescribeInstanceDropPacketResponseBody) *DescribeInstanceDropPacketResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceDropPacketResponse) Validate() error {
	return dara.Validate(s)
}
