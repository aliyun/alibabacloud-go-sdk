// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnAddrs(v *DescribeEndpointsResponseBodyConnAddrs) *DescribeEndpointsResponseBody
	GetConnAddrs() *DescribeEndpointsResponseBodyConnAddrs
	SetEngine(v string) *DescribeEndpointsResponseBody
	GetEngine() *string
	SetNetType(v string) *DescribeEndpointsResponseBody
	GetNetType() *string
	SetRequestId(v string) *DescribeEndpointsResponseBody
	GetRequestId() *string
	SetVSwitchId(v string) *DescribeEndpointsResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeEndpointsResponseBody
	GetVpcId() *string
}

type DescribeEndpointsResponseBody struct {
	ConnAddrs *DescribeEndpointsResponseBodyConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Struct"`
	// example:
	//
	// hbaseue
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// VPC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// example:
	//
	// F072593C-5234-5B56-9F63-3C7A3AD85D66
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// vsw-bp1foll427ze3d4ps****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp15s22y1a7sff5gj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBody) GetConnAddrs() *DescribeEndpointsResponseBodyConnAddrs {
	return s.ConnAddrs
}

func (s *DescribeEndpointsResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeEndpointsResponseBody) GetNetType() *string {
	return s.NetType
}

func (s *DescribeEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEndpointsResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeEndpointsResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeEndpointsResponseBody) SetConnAddrs(v *DescribeEndpointsResponseBodyConnAddrs) *DescribeEndpointsResponseBody {
	s.ConnAddrs = v
	return s
}

func (s *DescribeEndpointsResponseBody) SetEngine(v string) *DescribeEndpointsResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeEndpointsResponseBody) SetNetType(v string) *DescribeEndpointsResponseBody {
	s.NetType = &v
	return s
}

func (s *DescribeEndpointsResponseBody) SetRequestId(v string) *DescribeEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndpointsResponseBody) SetVSwitchId(v string) *DescribeEndpointsResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeEndpointsResponseBody) SetVpcId(v string) *DescribeEndpointsResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeEndpointsResponseBody) Validate() error {
	if s.ConnAddrs != nil {
		if err := s.ConnAddrs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEndpointsResponseBodyConnAddrs struct {
	ConnAddrInfo []*DescribeEndpointsResponseBodyConnAddrsConnAddrInfo `json:"ConnAddrInfo,omitempty" xml:"ConnAddrInfo,omitempty" type:"Repeated"`
}

func (s DescribeEndpointsResponseBodyConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointsResponseBodyConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBodyConnAddrs) GetConnAddrInfo() []*DescribeEndpointsResponseBodyConnAddrsConnAddrInfo {
	return s.ConnAddrInfo
}

func (s *DescribeEndpointsResponseBodyConnAddrs) SetConnAddrInfo(v []*DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) *DescribeEndpointsResponseBodyConnAddrs {
	s.ConnAddrInfo = v
	return s
}

func (s *DescribeEndpointsResponseBodyConnAddrs) Validate() error {
	if s.ConnAddrInfo != nil {
		for _, item := range s.ConnAddrInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEndpointsResponseBodyConnAddrsConnAddrInfo struct {
	// example:
	//
	// ****
	ConnAddr *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	// example:
	//
	// ****
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	// example:
	//
	// zkConn
	ConnType *string `json:"ConnType,omitempty" xml:"ConnType,omitempty"`
	// example:
	//
	// 2
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) GetConnAddr() *string {
	return s.ConnAddr
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) GetConnAddrPort() *string {
	return s.ConnAddrPort
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) GetConnType() *string {
	return s.ConnType
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) GetNetType() *string {
	return s.NetType
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) SetConnAddr(v string) *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo {
	s.ConnAddr = &v
	return s
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) SetConnAddrPort(v string) *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) SetConnType(v string) *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo {
	s.ConnType = &v
	return s
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) SetNetType(v string) *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo {
	s.NetType = &v
	return s
}

func (s *DescribeEndpointsResponseBodyConnAddrsConnAddrInfo) Validate() error {
	return dara.Validate(s)
}
