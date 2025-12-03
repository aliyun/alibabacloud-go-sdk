// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterConnectAddrsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbType(v string) *DescribeClusterConnectAddrsResponseBody
	GetDbType() *string
	SetIsMultimod(v string) *DescribeClusterConnectAddrsResponseBody
	GetIsMultimod() *string
	SetNetType(v string) *DescribeClusterConnectAddrsResponseBody
	GetNetType() *string
	SetRequestId(v string) *DescribeClusterConnectAddrsResponseBody
	GetRequestId() *string
	SetServiceConnAddrs(v *DescribeClusterConnectAddrsResponseBodyServiceConnAddrs) *DescribeClusterConnectAddrsResponseBody
	GetServiceConnAddrs() *DescribeClusterConnectAddrsResponseBodyServiceConnAddrs
	SetSlbConnAddrs(v *DescribeClusterConnectAddrsResponseBodySlbConnAddrs) *DescribeClusterConnectAddrsResponseBody
	GetSlbConnAddrs() *DescribeClusterConnectAddrsResponseBodySlbConnAddrs
	SetThriftConn(v *DescribeClusterConnectAddrsResponseBodyThriftConn) *DescribeClusterConnectAddrsResponseBody
	GetThriftConn() *DescribeClusterConnectAddrsResponseBodyThriftConn
	SetUiProxyConnAddrInfo(v *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) *DescribeClusterConnectAddrsResponseBody
	GetUiProxyConnAddrInfo() *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo
	SetVSwitchId(v string) *DescribeClusterConnectAddrsResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeClusterConnectAddrsResponseBody
	GetVpcId() *string
	SetZkConnAddrs(v *DescribeClusterConnectAddrsResponseBodyZkConnAddrs) *DescribeClusterConnectAddrsResponseBody
	GetZkConnAddrs() *DescribeClusterConnectAddrsResponseBodyZkConnAddrs
}

type DescribeClusterConnectAddrsResponseBody struct {
	DbType              *string                                                     `json:"DbType,omitempty" xml:"DbType,omitempty"`
	IsMultimod          *string                                                     `json:"IsMultimod,omitempty" xml:"IsMultimod,omitempty"`
	NetType             *string                                                     `json:"NetType,omitempty" xml:"NetType,omitempty"`
	RequestId           *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceConnAddrs    *DescribeClusterConnectAddrsResponseBodyServiceConnAddrs    `json:"ServiceConnAddrs,omitempty" xml:"ServiceConnAddrs,omitempty" type:"Struct"`
	SlbConnAddrs        *DescribeClusterConnectAddrsResponseBodySlbConnAddrs        `json:"SlbConnAddrs,omitempty" xml:"SlbConnAddrs,omitempty" type:"Struct"`
	ThriftConn          *DescribeClusterConnectAddrsResponseBodyThriftConn          `json:"ThriftConn,omitempty" xml:"ThriftConn,omitempty" type:"Struct"`
	UiProxyConnAddrInfo *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo `json:"UiProxyConnAddrInfo,omitempty" xml:"UiProxyConnAddrInfo,omitempty" type:"Struct"`
	VSwitchId           *string                                                     `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId               *string                                                     `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZkConnAddrs         *DescribeClusterConnectAddrsResponseBodyZkConnAddrs         `json:"ZkConnAddrs,omitempty" xml:"ZkConnAddrs,omitempty" type:"Struct"`
}

func (s DescribeClusterConnectAddrsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBody) GetDbType() *string {
	return s.DbType
}

func (s *DescribeClusterConnectAddrsResponseBody) GetIsMultimod() *string {
	return s.IsMultimod
}

func (s *DescribeClusterConnectAddrsResponseBody) GetNetType() *string {
	return s.NetType
}

func (s *DescribeClusterConnectAddrsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterConnectAddrsResponseBody) GetServiceConnAddrs() *DescribeClusterConnectAddrsResponseBodyServiceConnAddrs {
	return s.ServiceConnAddrs
}

func (s *DescribeClusterConnectAddrsResponseBody) GetSlbConnAddrs() *DescribeClusterConnectAddrsResponseBodySlbConnAddrs {
	return s.SlbConnAddrs
}

func (s *DescribeClusterConnectAddrsResponseBody) GetThriftConn() *DescribeClusterConnectAddrsResponseBodyThriftConn {
	return s.ThriftConn
}

func (s *DescribeClusterConnectAddrsResponseBody) GetUiProxyConnAddrInfo() *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo {
	return s.UiProxyConnAddrInfo
}

func (s *DescribeClusterConnectAddrsResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeClusterConnectAddrsResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeClusterConnectAddrsResponseBody) GetZkConnAddrs() *DescribeClusterConnectAddrsResponseBodyZkConnAddrs {
	return s.ZkConnAddrs
}

func (s *DescribeClusterConnectAddrsResponseBody) SetDbType(v string) *DescribeClusterConnectAddrsResponseBody {
	s.DbType = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetIsMultimod(v string) *DescribeClusterConnectAddrsResponseBody {
	s.IsMultimod = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetNetType(v string) *DescribeClusterConnectAddrsResponseBody {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetRequestId(v string) *DescribeClusterConnectAddrsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetServiceConnAddrs(v *DescribeClusterConnectAddrsResponseBodyServiceConnAddrs) *DescribeClusterConnectAddrsResponseBody {
	s.ServiceConnAddrs = v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetSlbConnAddrs(v *DescribeClusterConnectAddrsResponseBodySlbConnAddrs) *DescribeClusterConnectAddrsResponseBody {
	s.SlbConnAddrs = v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetThriftConn(v *DescribeClusterConnectAddrsResponseBodyThriftConn) *DescribeClusterConnectAddrsResponseBody {
	s.ThriftConn = v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetUiProxyConnAddrInfo(v *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) *DescribeClusterConnectAddrsResponseBody {
	s.UiProxyConnAddrInfo = v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetVSwitchId(v string) *DescribeClusterConnectAddrsResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetVpcId(v string) *DescribeClusterConnectAddrsResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) SetZkConnAddrs(v *DescribeClusterConnectAddrsResponseBodyZkConnAddrs) *DescribeClusterConnectAddrsResponseBody {
	s.ZkConnAddrs = v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBody) Validate() error {
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

type DescribeClusterConnectAddrsResponseBodyServiceConnAddrs struct {
	ServiceConnAddr []*DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr `json:"ServiceConnAddr,omitempty" xml:"ServiceConnAddr,omitempty" type:"Repeated"`
}

func (s DescribeClusterConnectAddrsResponseBodyServiceConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodyServiceConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrs) GetServiceConnAddr() []*DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr {
	return s.ServiceConnAddr
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrs) SetServiceConnAddr(v []*DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr) *DescribeClusterConnectAddrsResponseBodyServiceConnAddrs {
	s.ServiceConnAddr = v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrs) Validate() error {
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

type DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr struct {
	ConnAddrInfo *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo `json:"ConnAddrInfo,omitempty" xml:"ConnAddrInfo,omitempty" type:"Struct"`
	ConnType     *string                                                                             `json:"ConnType,omitempty" xml:"ConnType,omitempty"`
}

func (s DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr) GetConnAddrInfo() *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	return s.ConnAddrInfo
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr) GetConnType() *string {
	return s.ConnType
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr) SetConnAddrInfo(v *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr {
	s.ConnAddrInfo = v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr) SetConnType(v string) *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr {
	s.ConnType = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddr) Validate() error {
	if s.ConnAddrInfo != nil {
		if err := s.ConnAddrInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo struct {
	ConnAddr     *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) GetConnAddr() *string {
	return s.ConnAddr
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) GetConnAddrPort() *string {
	return s.ConnAddrPort
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) GetNetType() *string {
	return s.NetType
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) SetConnAddr(v string) *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) SetConnAddrPort(v string) *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) SetNetType(v string) *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyServiceConnAddrsServiceConnAddrConnAddrInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterConnectAddrsResponseBodySlbConnAddrs struct {
	SlbConnAddr []*DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr `json:"SlbConnAddr,omitempty" xml:"SlbConnAddr,omitempty" type:"Repeated"`
}

func (s DescribeClusterConnectAddrsResponseBodySlbConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodySlbConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrs) GetSlbConnAddr() []*DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr {
	return s.SlbConnAddr
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrs) SetSlbConnAddr(v []*DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr) *DescribeClusterConnectAddrsResponseBodySlbConnAddrs {
	s.SlbConnAddr = v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrs) Validate() error {
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

type DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr struct {
	ConnAddrInfo *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo `json:"ConnAddrInfo,omitempty" xml:"ConnAddrInfo,omitempty" type:"Struct"`
	SlbType      *string                                                                     `json:"SlbType,omitempty" xml:"SlbType,omitempty"`
}

func (s DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr) GetConnAddrInfo() *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	return s.ConnAddrInfo
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr) GetSlbType() *string {
	return s.SlbType
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr) SetConnAddrInfo(v *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr {
	s.ConnAddrInfo = v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr) SetSlbType(v string) *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr {
	s.SlbType = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddr) Validate() error {
	if s.ConnAddrInfo != nil {
		if err := s.ConnAddrInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo struct {
	ConnAddr     *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) GetConnAddr() *string {
	return s.ConnAddr
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) GetConnAddrPort() *string {
	return s.ConnAddrPort
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) GetNetType() *string {
	return s.NetType
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) SetConnAddr(v string) *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) SetConnAddrPort(v string) *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) SetNetType(v string) *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodySlbConnAddrsSlbConnAddrConnAddrInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterConnectAddrsResponseBodyThriftConn struct {
	ConnAddr     *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectAddrsResponseBodyThriftConn) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodyThriftConn) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodyThriftConn) GetConnAddr() *string {
	return s.ConnAddr
}

func (s *DescribeClusterConnectAddrsResponseBodyThriftConn) GetConnAddrPort() *string {
	return s.ConnAddrPort
}

func (s *DescribeClusterConnectAddrsResponseBodyThriftConn) GetNetType() *string {
	return s.NetType
}

func (s *DescribeClusterConnectAddrsResponseBodyThriftConn) SetConnAddr(v string) *DescribeClusterConnectAddrsResponseBodyThriftConn {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyThriftConn) SetConnAddrPort(v string) *DescribeClusterConnectAddrsResponseBodyThriftConn {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyThriftConn) SetNetType(v string) *DescribeClusterConnectAddrsResponseBodyThriftConn {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyThriftConn) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo struct {
	ConnAddr     *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) GetConnAddr() *string {
	return s.ConnAddr
}

func (s *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) GetConnAddrPort() *string {
	return s.ConnAddrPort
}

func (s *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) GetNetType() *string {
	return s.NetType
}

func (s *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) SetConnAddr(v string) *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) SetConnAddrPort(v string) *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) SetNetType(v string) *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyUiProxyConnAddrInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterConnectAddrsResponseBodyZkConnAddrs struct {
	ZkConnAddr []*DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr `json:"ZkConnAddr,omitempty" xml:"ZkConnAddr,omitempty" type:"Repeated"`
}

func (s DescribeClusterConnectAddrsResponseBodyZkConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodyZkConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodyZkConnAddrs) GetZkConnAddr() []*DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr {
	return s.ZkConnAddr
}

func (s *DescribeClusterConnectAddrsResponseBodyZkConnAddrs) SetZkConnAddr(v []*DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr) *DescribeClusterConnectAddrsResponseBodyZkConnAddrs {
	s.ZkConnAddr = v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyZkConnAddrs) Validate() error {
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

type DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr struct {
	ConnAddr     *string `json:"ConnAddr,omitempty" xml:"ConnAddr,omitempty"`
	ConnAddrPort *string `json:"ConnAddrPort,omitempty" xml:"ConnAddrPort,omitempty"`
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr) GoString() string {
	return s.String()
}

func (s *DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr) GetConnAddr() *string {
	return s.ConnAddr
}

func (s *DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr) GetConnAddrPort() *string {
	return s.ConnAddrPort
}

func (s *DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr) GetNetType() *string {
	return s.NetType
}

func (s *DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr) SetConnAddr(v string) *DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr {
	s.ConnAddr = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr) SetConnAddrPort(v string) *DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr {
	s.ConnAddrPort = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr) SetNetType(v string) *DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr {
	s.NetType = &v
	return s
}

func (s *DescribeClusterConnectAddrsResponseBodyZkConnAddrsZkConnAddr) Validate() error {
	return dara.Validate(s)
}
