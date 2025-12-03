// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbType(v string) *DescribeClusterConnectionResponseBody
	GetDbType() *string
	SetIsMultimod(v string) *DescribeClusterConnectionResponseBody
	GetIsMultimod() *string
	SetNetType(v string) *DescribeClusterConnectionResponseBody
	GetNetType() *string
	SetRequestId(v string) *DescribeClusterConnectionResponseBody
	GetRequestId() *string
	SetServiceConnAddrs(v *DescribeClusterConnectionResponseBodyServiceConnAddrs) *DescribeClusterConnectionResponseBody
	GetServiceConnAddrs() *DescribeClusterConnectionResponseBodyServiceConnAddrs
	SetSlbConnAddrs(v *DescribeClusterConnectionResponseBodySlbConnAddrs) *DescribeClusterConnectionResponseBody
	GetSlbConnAddrs() *DescribeClusterConnectionResponseBodySlbConnAddrs
	SetThriftConn(v *DescribeClusterConnectionResponseBodyThriftConn) *DescribeClusterConnectionResponseBody
	GetThriftConn() *DescribeClusterConnectionResponseBodyThriftConn
	SetUiProxyConnAddrInfo(v *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) *DescribeClusterConnectionResponseBody
	GetUiProxyConnAddrInfo() *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo
	SetVSwitchId(v string) *DescribeClusterConnectionResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeClusterConnectionResponseBody
	GetVpcId() *string
	SetZkConnAddrs(v *DescribeClusterConnectionResponseBodyZkConnAddrs) *DescribeClusterConnectionResponseBody
	GetZkConnAddrs() *DescribeClusterConnectionResponseBodyZkConnAddrs
}

type DescribeClusterConnectionResponseBody struct {
	// example:
	//
	// hbaseue
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// true
	IsMultimod *string `json:"IsMultimod,omitempty" xml:"IsMultimod,omitempty"`
	// example:
	//
	// VPC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// example:
	//
	// 70220050-A465-5DCC-8C0C-C38C6E3DB24D
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceConnAddrs    *DescribeClusterConnectionResponseBodyServiceConnAddrs    `json:"ServiceConnAddrs,omitempty" xml:"ServiceConnAddrs,omitempty" type:"Struct"`
	SlbConnAddrs        *DescribeClusterConnectionResponseBodySlbConnAddrs        `json:"SlbConnAddrs,omitempty" xml:"SlbConnAddrs,omitempty" type:"Struct"`
	ThriftConn          *DescribeClusterConnectionResponseBodyThriftConn          `json:"ThriftConn,omitempty" xml:"ThriftConn,omitempty" type:"Struct"`
	UiProxyConnAddrInfo *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo `json:"UiProxyConnAddrInfo,omitempty" xml:"UiProxyConnAddrInfo,omitempty" type:"Struct"`
	// example:
	//
	// vsw-bp1foll427ze3d4ps****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp15s22y1a7sff5gj****
	VpcId       *string                                           `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZkConnAddrs *DescribeClusterConnectionResponseBodyZkConnAddrs `json:"ZkConnAddrs,omitempty" xml:"ZkConnAddrs,omitempty" type:"Struct"`
}

func (s DescribeClusterConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBody) GetDbType() *string {
	return s.DbType
}

func (s *DescribeClusterConnectionResponseBody) GetIsMultimod() *string {
	return s.IsMultimod
}

func (s *DescribeClusterConnectionResponseBody) GetNetType() *string {
	return s.NetType
}

func (s *DescribeClusterConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterConnectionResponseBody) GetServiceConnAddrs() *DescribeClusterConnectionResponseBodyServiceConnAddrs {
	return s.ServiceConnAddrs
}

func (s *DescribeClusterConnectionResponseBody) GetSlbConnAddrs() *DescribeClusterConnectionResponseBodySlbConnAddrs {
	return s.SlbConnAddrs
}

func (s *DescribeClusterConnectionResponseBody) GetThriftConn() *DescribeClusterConnectionResponseBodyThriftConn {
	return s.ThriftConn
}

func (s *DescribeClusterConnectionResponseBody) GetUiProxyConnAddrInfo() *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo {
	return s.UiProxyConnAddrInfo
}

func (s *DescribeClusterConnectionResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeClusterConnectionResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeClusterConnectionResponseBody) GetZkConnAddrs() *DescribeClusterConnectionResponseBodyZkConnAddrs {
	return s.ZkConnAddrs
}

func (s *DescribeClusterConnectionResponseBody) SetDbType(v string) *DescribeClusterConnectionResponseBody {
	s.DbType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetIsMultimod(v string) *DescribeClusterConnectionResponseBody {
	s.IsMultimod = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetNetType(v string) *DescribeClusterConnectionResponseBody {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetRequestId(v string) *DescribeClusterConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetServiceConnAddrs(v *DescribeClusterConnectionResponseBodyServiceConnAddrs) *DescribeClusterConnectionResponseBody {
	s.ServiceConnAddrs = v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetSlbConnAddrs(v *DescribeClusterConnectionResponseBodySlbConnAddrs) *DescribeClusterConnectionResponseBody {
	s.SlbConnAddrs = v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetThriftConn(v *DescribeClusterConnectionResponseBodyThriftConn) *DescribeClusterConnectionResponseBody {
	s.ThriftConn = v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetUiProxyConnAddrInfo(v *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) *DescribeClusterConnectionResponseBody {
	s.UiProxyConnAddrInfo = v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetVSwitchId(v string) *DescribeClusterConnectionResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetVpcId(v string) *DescribeClusterConnectionResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterConnectionResponseBody) SetZkConnAddrs(v *DescribeClusterConnectionResponseBodyZkConnAddrs) *DescribeClusterConnectionResponseBody {
	s.ZkConnAddrs = v
	return s
}

func (s *DescribeClusterConnectionResponseBody) Validate() error {
	if s.ServiceConnAddrs != nil {
		if err := s.ServiceConnAddrs.Validate(); err != nil {
			return err
		}
	}
	if s.SlbConnAddrs != nil {
		if err := s.SlbConnAddrs.Validate(); err != nil {
			return err
		}
	}
	if s.ThriftConn != nil {
		if err := s.ThriftConn.Validate(); err != nil {
			return err
		}
	}
	if s.UiProxyConnAddrInfo != nil {
		if err := s.UiProxyConnAddrInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ZkConnAddrs != nil {
		if err := s.ZkConnAddrs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterConnectionResponseBodyServiceConnAddrs struct {
	ServiceConnAddr []*DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr `json:"ServiceConnAddr,omitempty" xml:"ServiceConnAddr,omitempty" type:"Repeated"`
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrs) GetServiceConnAddr() []*DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr {
	return s.ServiceConnAddr
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrs) SetServiceConnAddr(v []*DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) *DescribeClusterConnectionResponseBodyServiceConnAddrs {
	s.ServiceConnAddr = v
	return s
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrs) Validate() error {
	if s.ServiceConnAddr != nil {
		for _, item := range s.ServiceConnAddr {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr struct {
	ConnAddrInfo *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo `json:"ConnAddrInfo,omitempty" xml:"ConnAddrInfo,omitempty" type:"Struct"`
	// example:
	//
	// PhoenixConnAddr
	ConnType *string `json:"ConnType,omitempty" xml:"ConnType,omitempty"`
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) GetConnAddrInfo() *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	return s.ConnAddrInfo
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) GetConnType() *string {
	return s.ConnType
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) SetConnAddrInfo(v *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr {
	s.ConnAddrInfo = v
	return s
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) SetConnType(v string) *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr {
	s.ConnType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddr) Validate() error {
	if s.ConnAddrInfo != nil {
		if err := s.ConnAddrInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo struct {
	// example:
	//
	// hb-****-proxy-phoenix.hbase.rds.aliyuncs.com
	ConnAddr *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	// example:
	//
	// 8765
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	// example:
	//
	// 2
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) GetConnAddr() *string {
	return s.ConnAddr
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) GetConnAddrPort() *string {
	return s.ConnAddrPort
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) GetNetType() *string {
	return s.NetType
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) SetConnAddr(v string) *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) SetConnAddrPort(v string) *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) SetNetType(v string) *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterConnectionResponseBodySlbConnAddrs struct {
	SlbConnAddr []*DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr `json:"SlbConnAddr,omitempty" xml:"SlbConnAddr,omitempty" type:"Repeated"`
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrs) GetSlbConnAddr() []*DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr {
	return s.SlbConnAddr
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrs) SetSlbConnAddr(v []*DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) *DescribeClusterConnectionResponseBodySlbConnAddrs {
	s.SlbConnAddr = v
	return s
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrs) Validate() error {
	if s.SlbConnAddr != nil {
		for _, item := range s.SlbConnAddr {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr struct {
	ConnAddrInfo *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo `json:"ConnAddrInfo,omitempty" xml:"ConnAddrInfo,omitempty" type:"Struct"`
	// example:
	//
	// hbaseue
	SlbType *string `json:"SlbType,omitempty" xml:"SlbType,omitempty"`
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) GetConnAddrInfo() *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	return s.ConnAddrInfo
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) GetSlbType() *string {
	return s.SlbType
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) SetConnAddrInfo(v *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr {
	s.ConnAddrInfo = v
	return s
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) SetSlbType(v string) *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr {
	s.SlbType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddr) Validate() error {
	if s.ConnAddrInfo != nil {
		if err := s.ConnAddrInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo struct {
	// example:
	//
	// ld-bp150tns0sjxs****-proxy-hbaseue-pub.hbaseue.rds.aliyuncs.com
	ConnAddr *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	// example:
	//
	// 9190
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	// example:
	//
	// 0
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) GetConnAddr() *string {
	return s.ConnAddr
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) GetConnAddrPort() *string {
	return s.ConnAddrPort
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) GetNetType() *string {
	return s.NetType
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) SetConnAddr(v string) *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) SetConnAddrPort(v string) *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) SetNetType(v string) *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterConnectionResponseBodyThriftConn struct {
	// example:
	//
	// hb-bp1u0639js2h7****-proxy-thrift.hbase.rds.aliyuncs.com
	ConnAddr *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	// example:
	//
	// 9099
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	// example:
	//
	// 2
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectionResponseBodyThriftConn) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyThriftConn) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyThriftConn) GetConnAddr() *string {
	return s.ConnAddr
}

func (s *DescribeClusterConnectionResponseBodyThriftConn) GetConnAddrPort() *string {
	return s.ConnAddrPort
}

func (s *DescribeClusterConnectionResponseBodyThriftConn) GetNetType() *string {
	return s.NetType
}

func (s *DescribeClusterConnectionResponseBodyThriftConn) SetConnAddr(v string) *DescribeClusterConnectionResponseBodyThriftConn {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyThriftConn) SetConnAddrPort(v string) *DescribeClusterConnectionResponseBodyThriftConn {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyThriftConn) SetNetType(v string) *DescribeClusterConnectionResponseBodyThriftConn {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyThriftConn) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo struct {
	// example:
	//
	// ld-bp150tns0sjxs****-master1-001.hbaseue.rds.aliyuncs.com
	ConnAddr *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	// example:
	//
	// 443
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	// example:
	//
	// PUBLIC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) GetConnAddr() *string {
	return s.ConnAddr
}

func (s *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) GetConnAddrPort() *string {
	return s.ConnAddrPort
}

func (s *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) GetNetType() *string {
	return s.NetType
}

func (s *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) SetConnAddr(v string) *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) SetConnAddrPort(v string) *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) SetNetType(v string) *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyUiProxyConnAddrInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterConnectionResponseBodyZkConnAddrs struct {
	ZkConnAddr []*DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr `json:"ZkConnAddr,omitempty" xml:"ZkConnAddr,omitempty" type:"Repeated"`
}

func (s DescribeClusterConnectionResponseBodyZkConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyZkConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrs) GetZkConnAddr() []*DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr {
	return s.ZkConnAddr
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrs) SetZkConnAddr(v []*DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) *DescribeClusterConnectionResponseBodyZkConnAddrs {
	s.ZkConnAddr = v
	return s
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrs) Validate() error {
	if s.ZkConnAddr != nil {
		for _, item := range s.ZkConnAddr {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr struct {
	// example:
	//
	// ld-bp150tns0sjxs****-master1-001.hbaseue.rds.aliyuncs.com
	ConnAddr *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	// example:
	//
	// 2181
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	// example:
	//
	// 2
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) GetConnAddr() *string {
	return s.ConnAddr
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) GetConnAddrPort() *string {
	return s.ConnAddrPort
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) GetNetType() *string {
	return s.NetType
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) SetConnAddr(v string) *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) SetConnAddrPort(v string) *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) SetNetType(v string) *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectionResponseBodyZkConnAddrsZkConnAddr) Validate() error {
	return dara.Validate(s)
}
