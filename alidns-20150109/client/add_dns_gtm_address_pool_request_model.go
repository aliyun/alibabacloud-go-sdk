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
	// The list of addresses in the address pool.
	//
	// This parameter is required.
	Addr []*AddDnsGtmAddressPoolRequestAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Repeated"`
	// The number of consecutive failed health checks.
	//
	// example:
	//
	// 2
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The instance ID.
	//
	// <props="intl">Call [DescribeDnsGtmInstances](https://www.alibabacloud.com/help/en/dns/api-alidns-2015-01-09-describednsgtminstances) to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gtm-cn-cs02xyk****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The health check interval. Unit: seconds.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The list of health check nodes. If MonitorStatus is set to OPEN, you must specify at least one valid health check node.
	IspCityNode []*AddDnsGtmAddressPoolRequestIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
	// The language of some returned parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The load balancing policy. Valid values:
	//
	// - ALL_RR: returns all addresses.
	//
	// - RATIO: returns addresses by weight.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL_RR
	LbaStrategy *string `json:"LbaStrategy,omitempty" xml:"LbaStrategy,omitempty"`
	// The extended information. This parameter is a JSON string. The required parameters vary based on the health check protocol:
	//
	// - HTTP and HTTPS:
	//
	//   - port: The health check port.
	//
	//   - host: The host settings.
	//
	//   - path: The URL path.
	//
	//   - code: The return code. A response with a status code greater than this value is considered abnormal. Valid values: 400 and 500.
	//
	//   - failureRate: The failure rate.
	//
	//   - sni: Specifies whether to enable Server Name Indication (SNI). This parameter is available only for the HTTPS protocol.
	//
	//     - true: Enable SNI.
	//
	//     - Other values: Disable SNI.
	//
	//   - nodeType: The type of the health check node when the address pool type is DOMAIN. Valid values:
	//
	//     - IPV4
	//
	//     - IPV6
	//
	// - PING:
	//
	//   - failureRate: The failure rate.
	//
	//   - packetNum: The number of ping packets.
	//
	//   - packetLossRate: The packet loss rate.
	//
	//   - nodeType: The type of the health check node when the address pool type is DOMAIN. Valid values:
	//
	//     - IPV4
	//
	//     - IPV6
	//
	// - TCP:
	//
	//   - port: The health check port.
	//
	//   - failureRate: The failure rate.
	//
	//   - nodeType: The type of the health check node when the address pool type is DOMAIN. Valid values:
	//
	//     - IPV4
	//
	//     - IPV6
	//
	// example:
	//
	// {"failureRate":50,"port":80}
	MonitorExtendInfo *string `json:"MonitorExtendInfo,omitempty" xml:"MonitorExtendInfo,omitempty"`
	// The status of the health check feature. Default value: CLOSE. If you set this parameter to OPEN, the health check configuration is verified. Otherwise, the configuration is ignored.
	//
	// - OPEN: enabled
	//
	// - CLOSE: disabled
	//
	// example:
	//
	// OPEN
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
	// - HTTP
	//
	// - HTTPS
	//
	// - PING
	//
	// - TCP
	//
	// example:
	//
	// TCP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The timeout period. Unit: milliseconds.
	//
	// example:
	//
	// 2000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The type of the address pool. Valid values:
	//
	// - IPV4: IPv4 address
	//
	// - IPV6: IPv6 address
	//
	// - DOMAIN: domain name
	//
	// This parameter is required.
	//
	// example:
	//
	// IPV4
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
	// The address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 223.5.XX.XX
	Addr *string `json:"Addr,omitempty" xml:"Addr,omitempty"`
	// The source region of the address. This parameter is a JSON string.
	//
	// - lineCode: The line code of the source region.
	//
	// - lineCodeRectifyType: The rectification type for the line code. Default value: AUTO.
	//
	//   - NO_NEED: No rectification is performed.
	//
	//   - RECTIFIED: The line code is rectified.
	//
	//   - AUTO: The line code is automatically rectified.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"lineCodeRectifyType":"AUTO", "lineCode":[]}
	AttributeInfo *string `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	// The weight.
	//
	// example:
	//
	// 1
	LbaWeight *int32 `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	// The mode. Valid values:
	//
	// - SMART: smart return
	//
	// - ONLINE: always online
	//
	// - OFFLINE: always offline
	//
	// This parameter is required.
	//
	// example:
	//
	// SMART
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The remarks.
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
	// The city code. If MonitorStatus is set to OPEN, CityCode is required.
	//
	// For information about the valid values, see the response of the DescribeDnsGtmMonitorAvailableConfig operation.
	//
	// example:
	//
	// 503
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// If MonitorStatus is set to OPEN, IspCode is required.
	//
	// For information about the valid values, see the response of the DescribeDnsGtmMonitorAvailableConfig operation.
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
