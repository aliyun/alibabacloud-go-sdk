// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDestinationPortEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPortList(v []*DescribeDestinationPortEventResponseBodyPortList) *DescribeDestinationPortEventResponseBody
	GetPortList() []*DescribeDestinationPortEventResponseBodyPortList
	SetRequestId(v string) *DescribeDestinationPortEventResponseBody
	GetRequestId() *string
}

type DescribeDestinationPortEventResponseBody struct {
	// The ports.
	PortList []*DescribeDestinationPortEventResponseBodyPortList `json:"PortList,omitempty" xml:"PortList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9E7F6B2C-03F2-462F-9076-B782CF0DD502
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDestinationPortEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDestinationPortEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDestinationPortEventResponseBody) GetPortList() []*DescribeDestinationPortEventResponseBodyPortList {
	return s.PortList
}

func (s *DescribeDestinationPortEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDestinationPortEventResponseBody) SetPortList(v []*DescribeDestinationPortEventResponseBodyPortList) *DescribeDestinationPortEventResponseBody {
	s.PortList = v
	return s
}

func (s *DescribeDestinationPortEventResponseBody) SetRequestId(v string) *DescribeDestinationPortEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDestinationPortEventResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDestinationPortEventResponseBodyPortList struct {
	// The destination port.
	//
	// example:
	//
	// 80
	DstPort *string `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// The number of request packets received by the destination port.
	//
	// example:
	//
	// 8760950
	InPkts *int64 `json:"InPkts,omitempty" xml:"InPkts,omitempty"`
}

func (s DescribeDestinationPortEventResponseBodyPortList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDestinationPortEventResponseBodyPortList) GoString() string {
	return s.String()
}

func (s *DescribeDestinationPortEventResponseBodyPortList) GetDstPort() *string {
	return s.DstPort
}

func (s *DescribeDestinationPortEventResponseBodyPortList) GetInPkts() *int64 {
	return s.InPkts
}

func (s *DescribeDestinationPortEventResponseBodyPortList) SetDstPort(v string) *DescribeDestinationPortEventResponseBodyPortList {
	s.DstPort = &v
	return s
}

func (s *DescribeDestinationPortEventResponseBodyPortList) SetInPkts(v int64) *DescribeDestinationPortEventResponseBodyPortList {
	s.InPkts = &v
	return s
}

func (s *DescribeDestinationPortEventResponseBodyPortList) Validate() error {
	return dara.Validate(s)
}
