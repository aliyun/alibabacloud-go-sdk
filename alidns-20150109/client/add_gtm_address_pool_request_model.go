// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGtmAddressPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddr(v []*AddGtmAddressPoolRequestAddr) *AddGtmAddressPoolRequest
	GetAddr() []*AddGtmAddressPoolRequestAddr
	SetEvaluationCount(v int32) *AddGtmAddressPoolRequest
	GetEvaluationCount() *int32
	SetInstanceId(v string) *AddGtmAddressPoolRequest
	GetInstanceId() *string
	SetInterval(v int32) *AddGtmAddressPoolRequest
	GetInterval() *int32
	SetIspCityNode(v []*AddGtmAddressPoolRequestIspCityNode) *AddGtmAddressPoolRequest
	GetIspCityNode() []*AddGtmAddressPoolRequestIspCityNode
	SetLang(v string) *AddGtmAddressPoolRequest
	GetLang() *string
	SetMinAvailableAddrNum(v int32) *AddGtmAddressPoolRequest
	GetMinAvailableAddrNum() *int32
	SetMonitorExtendInfo(v string) *AddGtmAddressPoolRequest
	GetMonitorExtendInfo() *string
	SetMonitorStatus(v string) *AddGtmAddressPoolRequest
	GetMonitorStatus() *string
	SetName(v string) *AddGtmAddressPoolRequest
	GetName() *string
	SetProtocolType(v string) *AddGtmAddressPoolRequest
	GetProtocolType() *string
	SetTimeout(v int32) *AddGtmAddressPoolRequest
	GetTimeout() *int32
	SetType(v string) *AddGtmAddressPoolRequest
	GetType() *string
}

type AddGtmAddressPoolRequest struct {
	// The address pools.
	//
	// This parameter is required.
	Addr []*AddGtmAddressPoolRequestAddr `json:"Addr,omitempty" xml:"Addr,omitempty" type:"Repeated"`
	// The number of consecutive failures.
	//
	// example:
	//
	// 2
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The ID of the GTM instance for which you want to create an address pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// gtm-cn-xxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The health check interval. Unit: seconds. Set the value to 60.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The monitored nodes.
	IspCityNode []*AddGtmAddressPoolRequestIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
	// The language of the values of specific response parameters.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The minimum number of available addresses in the address pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	MinAvailableAddrNum *int32 `json:"MinAvailableAddrNum,omitempty" xml:"MinAvailableAddrNum,omitempty"`
	// The extended information. The required parameters vary based on the value of ProtocolType.
	//
	// When ProtocolType is set to HTTP or HTTPS:
	//
	// 	- port: the port that you want to check
	//
	// 	- failureRate: the failure rate
	//
	// 	- code: the return code. The health check result is deemed abnormal if the returned value is greater than the specified value. Valid values: 400 and 500.
	//
	// 	- host: the host settings
	//
	// 	- path: the URL path
	//
	// When ProtocolType is set to PING:
	//
	// 	- packetNum: the number of ping packets
	//
	// 	- packetLossRate: the packet loss rate
	//
	// 	- failureRate: the failure rate
	//
	// When ProtocolType is set to TCP:
	//
	// 	- port: the port that you want to check
	//
	// 	- failureRate: the failure rate
	//
	// example:
	//
	// {"host":"aliyun.com","port":80}
	MonitorExtendInfo *string `json:"MonitorExtendInfo,omitempty" xml:"MonitorExtendInfo,omitempty"`
	// Specifies whether to enable the health check. Valid values:
	//
	// 	- **OPEN**: enables the health check.
	//
	// 	- **CLOSE**: disables the health check. This is the default value.
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
	// Alibaba Cloud cluster
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The health check protocol. Valid values:
	//
	// 	- HTTP
	//
	// 	- HTTPS
	//
	// 	- Ping
	//
	// 	- TCP
	//
	// example:
	//
	// HTTPS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The timeout period. Unit: milliseconds. Valid values: 2000, 3000, 5000, and 10000.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The type of the address pool. Valid values:
	//
	// 	- **IP**: IPv4 address
	//
	// 	- **DOMAIN**: domain name
	//
	// This parameter is required.
	//
	// example:
	//
	// IP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddGtmAddressPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGtmAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *AddGtmAddressPoolRequest) GetAddr() []*AddGtmAddressPoolRequestAddr {
	return s.Addr
}

func (s *AddGtmAddressPoolRequest) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *AddGtmAddressPoolRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddGtmAddressPoolRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *AddGtmAddressPoolRequest) GetIspCityNode() []*AddGtmAddressPoolRequestIspCityNode {
	return s.IspCityNode
}

func (s *AddGtmAddressPoolRequest) GetLang() *string {
	return s.Lang
}

func (s *AddGtmAddressPoolRequest) GetMinAvailableAddrNum() *int32 {
	return s.MinAvailableAddrNum
}

func (s *AddGtmAddressPoolRequest) GetMonitorExtendInfo() *string {
	return s.MonitorExtendInfo
}

func (s *AddGtmAddressPoolRequest) GetMonitorStatus() *string {
	return s.MonitorStatus
}

func (s *AddGtmAddressPoolRequest) GetName() *string {
	return s.Name
}

func (s *AddGtmAddressPoolRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *AddGtmAddressPoolRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *AddGtmAddressPoolRequest) GetType() *string {
	return s.Type
}

func (s *AddGtmAddressPoolRequest) SetAddr(v []*AddGtmAddressPoolRequestAddr) *AddGtmAddressPoolRequest {
	s.Addr = v
	return s
}

func (s *AddGtmAddressPoolRequest) SetEvaluationCount(v int32) *AddGtmAddressPoolRequest {
	s.EvaluationCount = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetInstanceId(v string) *AddGtmAddressPoolRequest {
	s.InstanceId = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetInterval(v int32) *AddGtmAddressPoolRequest {
	s.Interval = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetIspCityNode(v []*AddGtmAddressPoolRequestIspCityNode) *AddGtmAddressPoolRequest {
	s.IspCityNode = v
	return s
}

func (s *AddGtmAddressPoolRequest) SetLang(v string) *AddGtmAddressPoolRequest {
	s.Lang = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetMinAvailableAddrNum(v int32) *AddGtmAddressPoolRequest {
	s.MinAvailableAddrNum = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetMonitorExtendInfo(v string) *AddGtmAddressPoolRequest {
	s.MonitorExtendInfo = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetMonitorStatus(v string) *AddGtmAddressPoolRequest {
	s.MonitorStatus = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetName(v string) *AddGtmAddressPoolRequest {
	s.Name = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetProtocolType(v string) *AddGtmAddressPoolRequest {
	s.ProtocolType = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetTimeout(v int32) *AddGtmAddressPoolRequest {
	s.Timeout = &v
	return s
}

func (s *AddGtmAddressPoolRequest) SetType(v string) *AddGtmAddressPoolRequest {
	s.Type = &v
	return s
}

func (s *AddGtmAddressPoolRequest) Validate() error {
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

type AddGtmAddressPoolRequestAddr struct {
	// The weight of the address pool.
	//
	// example:
	//
	// 1
	LbaWeight *int32 `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	// The mode of the address pool. Valid values:
	//
	// 	- **SMART**: smart return
	//
	// 	- **ONLINE**: always online
	//
	// 	- **OFFLINE**: always offline
	//
	// example:
	//
	// SMART
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The address in the address pool.
	//
	// example:
	//
	// 1.1.1.1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddGtmAddressPoolRequestAddr) String() string {
	return dara.Prettify(s)
}

func (s AddGtmAddressPoolRequestAddr) GoString() string {
	return s.String()
}

func (s *AddGtmAddressPoolRequestAddr) GetLbaWeight() *int32 {
	return s.LbaWeight
}

func (s *AddGtmAddressPoolRequestAddr) GetMode() *string {
	return s.Mode
}

func (s *AddGtmAddressPoolRequestAddr) GetValue() *string {
	return s.Value
}

func (s *AddGtmAddressPoolRequestAddr) SetLbaWeight(v int32) *AddGtmAddressPoolRequestAddr {
	s.LbaWeight = &v
	return s
}

func (s *AddGtmAddressPoolRequestAddr) SetMode(v string) *AddGtmAddressPoolRequestAddr {
	s.Mode = &v
	return s
}

func (s *AddGtmAddressPoolRequestAddr) SetValue(v string) *AddGtmAddressPoolRequestAddr {
	s.Value = &v
	return s
}

func (s *AddGtmAddressPoolRequestAddr) Validate() error {
	return dara.Validate(s)
}

type AddGtmAddressPoolRequestIspCityNode struct {
	// The code of the city where the monitored node is deployed. For more information about specific values, see the response parameters of DescribeGtmMonitorAvailableConfig.
	//
	// example:
	//
	// 546
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// 	- The code of the Internet service provider (ISP) to which the monitored node belongs. For more information about specific values, see the response parameters of DescribeGtmMonitorAvailableConfig.
	//
	// 	- If the value of the GroupType parameter is BGP or OVERSEAS, IspCode is optional. The default value is 465.
	//
	// 	- If the value of the GroupType parameter is not BGP or OVERSEAS, IspCode is required and is used together with CityCode.
	//
	// example:
	//
	// 465
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
}

func (s AddGtmAddressPoolRequestIspCityNode) String() string {
	return dara.Prettify(s)
}

func (s AddGtmAddressPoolRequestIspCityNode) GoString() string {
	return s.String()
}

func (s *AddGtmAddressPoolRequestIspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *AddGtmAddressPoolRequestIspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *AddGtmAddressPoolRequestIspCityNode) SetCityCode(v string) *AddGtmAddressPoolRequestIspCityNode {
	s.CityCode = &v
	return s
}

func (s *AddGtmAddressPoolRequestIspCityNode) SetIspCode(v string) *AddGtmAddressPoolRequestIspCityNode {
	s.IspCode = &v
	return s
}

func (s *AddGtmAddressPoolRequestIspCityNode) Validate() error {
	return dara.Validate(s)
}
