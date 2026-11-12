// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceVncUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeInstanceVncUrlResponseBody
	GetRequestId() *string
	SetVncUrl(v string) *DescribeInstanceVncUrlResponseBody
	GetVncUrl() *string
}

type DescribeInstanceVncUrlResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The VNC connection address of the server.
	//
	// example:
	//
	// wss%3A%2F%2Fhz01-vncproxy.aliyun.com%2Fwebsockify%2F%3Fs%3Dwz3L8wEMO6KMt7%252FXInEMtKVubBB%252F7rv055kOm8eUOD%252*****YlmsKjOfz6
	VncUrl *string `json:"VncUrl,omitempty" xml:"VncUrl,omitempty"`
}

func (s DescribeInstanceVncUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceVncUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceVncUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceVncUrlResponseBody) GetVncUrl() *string {
	return s.VncUrl
}

func (s *DescribeInstanceVncUrlResponseBody) SetRequestId(v string) *DescribeInstanceVncUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceVncUrlResponseBody) SetVncUrl(v string) *DescribeInstanceVncUrlResponseBody {
	s.VncUrl = &v
	return s
}

func (s *DescribeInstanceVncUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
