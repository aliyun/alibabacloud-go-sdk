// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnsGtmAddressPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddr(v []*AddDnsGtmAddressPoolRequestAddr) *AddDnsGtmAddressPoolRequest
	GetAddr() []*AddDnsGtmAddressPoolRequestAddr
	SetEvaluationCount(v int32) *AddDnsGtmAddressPoolRequest
	GetEvaluationCount() *int32
	SetInstanceId(v string) *AddDnsGtmAddressPoolRequest
	GetInstanceId() *string
	SetInterval(v int32) *AddDnsGtmAddressPoolRequest
	GetInterval() *int32
	SetIspCityNode(v []*AddDnsGtmAddressPoolRequestIspCityNode) *AddDnsGtmAddressPoolRequest
	GetIspCityNode() []*AddDnsGtmAddressPoolRequestIspCityNode
	SetLang(v string) *AddDnsGtmAddressPoolRequest
	GetLang() *string
	SetLbaStrategy(v string) *AddDnsGtmAddressPoolRequest
	GetLbaStrategy() *string
	SetMonitorExtendInfo(v string) *AddDnsGtmAddressPoolRequest
	GetMonitorExtendInfo() *string
	SetMonitorStatus(v string) *AddDnsGtmAddressPoolRequest
	GetMonitorStatus() *string
	SetName(v string) *AddDnsGtmAddressPoolRequest
	GetName() *string
	SetProtocolType(v string) *AddDnsGtmAddressPoolRequest
	GetProtocolType() *string
	SetTimeout(v int32) *AddDnsGtmAddressPoolRequest
	GetTimeout() *int32
	SetType(v string) *AddDnsGtmAddressPoolRequest
	GetType() *string
}

type AddDnsGtmAddressPoolRequest struct {
	// The address pools.
	//
	// This parameter is required.
	Addr []*AddDnsGtmAddressPoolRequestAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Repeated"`
	// The number of consecutive failures.
	//
	// example:
	//
	// 1
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The health check interval. Unit: seconds.
	//
	// example:
	//
	// 1
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The nodes for monitoring.
	IspCityNode []*AddDnsGtmAddressPoolRequestIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
	// The language of the values of specific response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The load balancing policy of the address pool. Valid values:
	//
	// 	- ALL_RR: returns all addresses.
	//
	// 	- RATIO: returns addresses by weight.
	//
	// This parameter is required.
	//
	// example:
	//
	// all_rr
	LbaStrategy *string `json:"LbaStrategy,omitempty" xml:"LbaStrategy,omitempty"`
	// The extended information. The required parameters vary based on the health check protocol.
	//
	// 	- HTTP or HTTPS:
	//
	//     	- port: the port that you want to check
	//
	//     	- host: the host settings
	//
	//     	- path: the URL
	//
	//     	- code: the return code. The health check result is deemed abnormal if the returned value is greater than the specified value. Valid values: 400 and 500.
	//
	//     	- failureRate: the failure rate
	//
	//     	- sni: specifies whether to enable Server Name Indication (SNI). This parameter is available only when ProtocolType is set to HTTPS. Valid values:
	//
	//         	- true: enables SNI.
	//
	//         	- other: disables SNI.
	//
	//     	- nodeType: the type of the node for monitoring when Type is set to DOMAIN. Valid values:
	//
	//         	- IPV4
	//
	//         	- IPV6
	//
	// 	- ping:
	//
	//     	- failureRate: the failure rate
	//
	//     	- packetNum: the number of ping packets
	//
	//     	- packetLossRate: the loss rate of ping packets
	//
	//     	- nodeType: the type of the node for monitoring when Type is set to DOMAIN. Valid values:
	//
	//         	- IPV4
	//
	//         	- IPV6
	//
	// 	- TCP:
	//
	//     	- port: the port that you want to check
	//
	//     	- failureRate: the failure rate
	//
	//     	- nodeType: the type of the node for monitoring when Type is set to DOMAIN. Valid values:
	//
	//         	- IPV4
	//
	//         	- IPV6
	//
	// example:
	//
	// {\\"code\\":200,\\"path\\":\\"\\\\index.htm\\",\\"host\\":\\"aliyun.com\\"}
	MonitorExtendInfo *string `json:"MonitorExtendInfo,omitempty" xml:"MonitorExtendInfo,omitempty"`
	// Specifies whether to enable the health check feature. If you set this parameter to OPEN, the system verifies the health check configurations. If you set this parameter to CLOSE, the system discards the health check configurations. Default value: CLOSE. Valid values:
	//
	// 	- OPEN: enables the health check feature.
	//
	// 	- CLOSE: disables the health check feature.
	//
	// example:
	//
	// open
	MonitorStatus *string `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	// The name of the address pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The health check protocol. Valid values:
	//
	// 	- HTTP
	//
	// 	- HTTPS
	//
	// 	- PING
	//
	// 	- TCP
	//
	// example:
	//
	// http
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The timeout period. Unit: milliseconds.
	//
	// example:
	//
	// 1
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The type of the address pool. Valid values:
	//
	// 	- IPV4: IPv4 address
	//
	// 	- IPV6: IPv6 address
	//
	// 	- DOMAIN: domain name
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv4
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddDnsGtmAddressPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDnsGtmAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAddressPoolRequest) GetAddr() []*AddDnsGtmAddressPoolRequestAddr {
	return s.Addr
}

func (s *AddDnsGtmAddressPoolRequest) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *AddDnsGtmAddressPoolRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddDnsGtmAddressPoolRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *AddDnsGtmAddressPoolRequest) GetIspCityNode() []*AddDnsGtmAddressPoolRequestIspCityNode {
	return s.IspCityNode
}

func (s *AddDnsGtmAddressPoolRequest) GetLang() *string {
	return s.Lang
}

func (s *AddDnsGtmAddressPoolRequest) GetLbaStrategy() *string {
	return s.LbaStrategy
}

func (s *AddDnsGtmAddressPoolRequest) GetMonitorExtendInfo() *string {
	return s.MonitorExtendInfo
}

func (s *AddDnsGtmAddressPoolRequest) GetMonitorStatus() *string {
	return s.MonitorStatus
}

func (s *AddDnsGtmAddressPoolRequest) GetName() *string {
	return s.Name
}

func (s *AddDnsGtmAddressPoolRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *AddDnsGtmAddressPoolRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *AddDnsGtmAddressPoolRequest) GetType() *string {
	return s.Type
}

func (s *AddDnsGtmAddressPoolRequest) SetAddr(v []*AddDnsGtmAddressPoolRequestAddr) *AddDnsGtmAddressPoolRequest {
	s.Addr = v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetEvaluationCount(v int32) *AddDnsGtmAddressPoolRequest {
	s.EvaluationCount = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetInstanceId(v string) *AddDnsGtmAddressPoolRequest {
	s.InstanceId = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetInterval(v int32) *AddDnsGtmAddressPoolRequest {
	s.Interval = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetIspCityNode(v []*AddDnsGtmAddressPoolRequestIspCityNode) *AddDnsGtmAddressPoolRequest {
	s.IspCityNode = v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetLang(v string) *AddDnsGtmAddressPoolRequest {
	s.Lang = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetLbaStrategy(v string) *AddDnsGtmAddressPoolRequest {
	s.LbaStrategy = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetMonitorExtendInfo(v string) *AddDnsGtmAddressPoolRequest {
	s.MonitorExtendInfo = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetMonitorStatus(v string) *AddDnsGtmAddressPoolRequest {
	s.MonitorStatus = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetName(v string) *AddDnsGtmAddressPoolRequest {
	s.Name = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetProtocolType(v string) *AddDnsGtmAddressPoolRequest {
	s.ProtocolType = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetTimeout(v int32) *AddDnsGtmAddressPoolRequest {
	s.Timeout = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) SetType(v string) *AddDnsGtmAddressPoolRequest {
	s.Type = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequest) Validate() error {
	if s.Addr != nil {
		for _, item := range s.Addr {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IspCityNode != nil {
		for _, item := range s.IspCityNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddDnsGtmAddressPoolRequestAddr struct {
	// The address in the address pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.1.1.1
	Addr *string `json:"Addr,omitempty" xml:"Addr,omitempty"`
	// The information about the source region of the address. The value of this parameter is a JSON string. Valid values:
	//
	// 	- lineCode: the line code of the source region for the address
	//
	// 	- lineCodeRectifyType: the rectification type of the line code. Default value: AUTO. Valid values:
	//
	//     	- NO_NEED: no need for rectification
	//
	//     	- RECTIFIED: rectified
	//
	//     	- AUTO: automatic rectification
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	AttributeInfo *string `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	// The weight of the address.
	//
	// example:
	//
	// 1
	LbaWeight *int32 `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	// The return mode of the addresses: Valid values:
	//
	// 	- SMART: smart return
	//
	// 	- ONLINE: always online
	//
	// 	- OFFLINE: always offline
	//
	// This parameter is required.
	//
	// example:
	//
	// online
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The description of the address pool.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s AddDnsGtmAddressPoolRequestAddr) String() string {
	return dara.Prettify(s)
}

func (s AddDnsGtmAddressPoolRequestAddr) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAddressPoolRequestAddr) GetAddr() *string {
	return s.Addr
}

func (s *AddDnsGtmAddressPoolRequestAddr) GetAttributeInfo() *string {
	return s.AttributeInfo
}

func (s *AddDnsGtmAddressPoolRequestAddr) GetLbaWeight() *int32 {
	return s.LbaWeight
}

func (s *AddDnsGtmAddressPoolRequestAddr) GetMode() *string {
	return s.Mode
}

func (s *AddDnsGtmAddressPoolRequestAddr) GetRemark() *string {
	return s.Remark
}

func (s *AddDnsGtmAddressPoolRequestAddr) SetAddr(v string) *AddDnsGtmAddressPoolRequestAddr {
	s.Addr = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequestAddr) SetAttributeInfo(v string) *AddDnsGtmAddressPoolRequestAddr {
	s.AttributeInfo = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequestAddr) SetLbaWeight(v int32) *AddDnsGtmAddressPoolRequestAddr {
	s.LbaWeight = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequestAddr) SetMode(v string) *AddDnsGtmAddressPoolRequestAddr {
	s.Mode = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequestAddr) SetRemark(v string) *AddDnsGtmAddressPoolRequestAddr {
	s.Remark = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequestAddr) Validate() error {
	return dara.Validate(s)
}

type AddDnsGtmAddressPoolRequestIspCityNode struct {
	// The city code.
	//
	// Specify the parameter according to the value of CityCode returned by the DescribeGtmMonitorAvailableConfig operation.
	//
	// example:
	//
	// 503
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// 	- The Internet service provider (ISP) node. Specify the parameter according to the value of IspCode returned by the DescribeGtmMonitorAvailableConfig operation.
	//
	// 	- If the returned value of GroupType for the DescribeGtmMonitorAvailableConfig operation is BGP or Overseas, IspCode is not required and is set to 465 by default.
	//
	// 	- If the returned value of GroupType for the DescribeGtmMonitorAvailableConfig operation is not BGP or Overseas, IspCode is required. When IspCode is specified, CityCode is required.
	//
	// example:
	//
	// 465
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
}

func (s AddDnsGtmAddressPoolRequestIspCityNode) String() string {
	return dara.Prettify(s)
}

func (s AddDnsGtmAddressPoolRequestIspCityNode) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAddressPoolRequestIspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *AddDnsGtmAddressPoolRequestIspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *AddDnsGtmAddressPoolRequestIspCityNode) SetCityCode(v string) *AddDnsGtmAddressPoolRequestIspCityNode {
	s.CityCode = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequestIspCityNode) SetIspCode(v string) *AddDnsGtmAddressPoolRequestIspCityNode {
	s.IspCode = &v
	return s
}

func (s *AddDnsGtmAddressPoolRequestIspCityNode) Validate() error {
	return dara.Validate(s)
}
