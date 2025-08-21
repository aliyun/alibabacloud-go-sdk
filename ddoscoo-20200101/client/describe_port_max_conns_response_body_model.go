// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortMaxConnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPortMaxConns(v []*DescribePortMaxConnsResponseBodyPortMaxConns) *DescribePortMaxConnsResponseBody
	GetPortMaxConns() []*DescribePortMaxConnsResponseBodyPortMaxConns
	SetRequestId(v string) *DescribePortMaxConnsResponseBody
	GetRequestId() *string
}

type DescribePortMaxConnsResponseBody struct {
	// The details of the maximum number of connections that can be established over a port of the instance.
	PortMaxConns []*DescribePortMaxConnsResponseBodyPortMaxConns `json:"PortMaxConns,omitempty" xml:"PortMaxConns,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 08F79110-2AF5-4FA7-998E-7C5E75EACF9C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePortMaxConnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePortMaxConnsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortMaxConnsResponseBody) GetPortMaxConns() []*DescribePortMaxConnsResponseBodyPortMaxConns {
	return s.PortMaxConns
}

func (s *DescribePortMaxConnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePortMaxConnsResponseBody) SetPortMaxConns(v []*DescribePortMaxConnsResponseBodyPortMaxConns) *DescribePortMaxConnsResponseBody {
	s.PortMaxConns = v
	return s
}

func (s *DescribePortMaxConnsResponseBody) SetRequestId(v string) *DescribePortMaxConnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortMaxConnsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePortMaxConnsResponseBodyPortMaxConns struct {
	// The maximum number of connections per second (CPS).
	//
	// example:
	//
	// 100
	Cps *int64 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 203.***.***.117
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port of the instance.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribePortMaxConnsResponseBodyPortMaxConns) String() string {
	return dara.Prettify(s)
}

func (s DescribePortMaxConnsResponseBodyPortMaxConns) GoString() string {
	return s.String()
}

func (s *DescribePortMaxConnsResponseBodyPortMaxConns) GetCps() *int64 {
	return s.Cps
}

func (s *DescribePortMaxConnsResponseBodyPortMaxConns) GetIp() *string {
	return s.Ip
}

func (s *DescribePortMaxConnsResponseBodyPortMaxConns) GetPort() *string {
	return s.Port
}

func (s *DescribePortMaxConnsResponseBodyPortMaxConns) SetCps(v int64) *DescribePortMaxConnsResponseBodyPortMaxConns {
	s.Cps = &v
	return s
}

func (s *DescribePortMaxConnsResponseBodyPortMaxConns) SetIp(v string) *DescribePortMaxConnsResponseBodyPortMaxConns {
	s.Ip = &v
	return s
}

func (s *DescribePortMaxConnsResponseBodyPortMaxConns) SetPort(v string) *DescribePortMaxConnsResponseBodyPortMaxConns {
	s.Port = &v
	return s
}

func (s *DescribePortMaxConnsResponseBodyPortMaxConns) Validate() error {
	return dara.Validate(s)
}
