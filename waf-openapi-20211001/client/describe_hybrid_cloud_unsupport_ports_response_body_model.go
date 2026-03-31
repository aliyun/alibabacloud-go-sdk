// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudUnsupportPortsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHttpPorts(v string) *DescribeHybridCloudUnsupportPortsResponseBody
	GetHttpPorts() *string
	SetHttpsPorts(v string) *DescribeHybridCloudUnsupportPortsResponseBody
	GetHttpsPorts() *string
	SetRequestId(v string) *DescribeHybridCloudUnsupportPortsResponseBody
	GetRequestId() *string
}

type DescribeHybridCloudUnsupportPortsResponseBody struct {
	// The HTTP ports. The value is a string. If multiple ports are returned, the value is in the **port1,port2,port3*	- format.
	//
	// example:
	//
	// 80,8080
	HttpPorts *string `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty"`
	// The HTTPS ports. The value is a string. If multiple ports are returned, the value is in the **port1,port2,port3*	- format.
	//
	// example:
	//
	// 443,8443,8021,3443,2443,9011
	HttpsPorts *string `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C2E97B3F-1623-4CDF-A7E2-FD9D****027A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHybridCloudUnsupportPortsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudUnsupportPortsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudUnsupportPortsResponseBody) GetHttpPorts() *string {
	return s.HttpPorts
}

func (s *DescribeHybridCloudUnsupportPortsResponseBody) GetHttpsPorts() *string {
	return s.HttpsPorts
}

func (s *DescribeHybridCloudUnsupportPortsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridCloudUnsupportPortsResponseBody) SetHttpPorts(v string) *DescribeHybridCloudUnsupportPortsResponseBody {
	s.HttpPorts = &v
	return s
}

func (s *DescribeHybridCloudUnsupportPortsResponseBody) SetHttpsPorts(v string) *DescribeHybridCloudUnsupportPortsResponseBody {
	s.HttpsPorts = &v
	return s
}

func (s *DescribeHybridCloudUnsupportPortsResponseBody) SetRequestId(v string) *DescribeHybridCloudUnsupportPortsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudUnsupportPortsResponseBody) Validate() error {
	return dara.Validate(s)
}
