// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceVncUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoginUrl(v string) *DescribeRCInstanceVncUrlResponseBody
	GetLoginUrl() *string
	SetRequestId(v string) *DescribeRCInstanceVncUrlResponseBody
	GetRequestId() *string
}

type DescribeRCInstanceVncUrlResponseBody struct {
	// The VNC logon address.
	//
	// >  The address returned is valid only for 15 seconds. If you do not use the returned address to establish a connection within 15 seconds, the address expires and you must call the operation again to obtain a new address.
	//
	// example:
	//
	// https://g.alicdn.com/aliyun/ecs-console-vnc2/0.0.8/index.html?vncUrl=****&instanceId=i-2vcb1qjj8z5dl8iw****&isWindows=false
	LoginUrl *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F2911788-25E8-42E5-A3A3-1B38D263F01E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRCInstanceVncUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceVncUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceVncUrlResponseBody) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *DescribeRCInstanceVncUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCInstanceVncUrlResponseBody) SetLoginUrl(v string) *DescribeRCInstanceVncUrlResponseBody {
	s.LoginUrl = &v
	return s
}

func (s *DescribeRCInstanceVncUrlResponseBody) SetRequestId(v string) *DescribeRCInstanceVncUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCInstanceVncUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
