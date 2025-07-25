// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnsGtmMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddrPoolId(v string) *AddDnsGtmMonitorRequest
	GetAddrPoolId() *string
	SetEvaluationCount(v int32) *AddDnsGtmMonitorRequest
	GetEvaluationCount() *int32
	SetInterval(v int32) *AddDnsGtmMonitorRequest
	GetInterval() *int32
	SetIspCityNode(v []*AddDnsGtmMonitorRequestIspCityNode) *AddDnsGtmMonitorRequest
	GetIspCityNode() []*AddDnsGtmMonitorRequestIspCityNode
	SetLang(v string) *AddDnsGtmMonitorRequest
	GetLang() *string
	SetMonitorExtendInfo(v string) *AddDnsGtmMonitorRequest
	GetMonitorExtendInfo() *string
	SetProtocolType(v string) *AddDnsGtmMonitorRequest
	GetProtocolType() *string
	SetTimeout(v int32) *AddDnsGtmMonitorRequest
	GetTimeout() *int32
}

type AddDnsGtmMonitorRequest struct {
	// The ID of the address pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// pool1
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The maximum number of consecutive exceptions detected. If the number of consecutive exceptions detected reaches the maximum number, the application service is deemed abnormal.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The health check interval. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The monitored nodes.
	//
	// This parameter is required.
	IspCityNode []*AddDnsGtmMonitorRequestIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
	// The language of the values of specific response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The extended information. The required parameters vary based on the value of ProtocolType.
	//
	// 	- HTTP or HTTPS
	//
	//     	- port: the port that you want to check
	//
	//     	- host: the host settings
	//
	//     	- path: the URL path
	//
	//     	- code: the response code. The health check result is deemed abnormal if the returned value is greater than the specified value.
	//
	//     	- failureRate: the failure rate
	//
	//     	- sni: specifies whether to enable server name indication (SNI). This parameter is available only when ProtocolType is set to HTTPS. Valid values:
	//
	//         	- true: enables SNI.
	//
	//         	- false: disables SNI.
	//
	//     	- nodeType: the type of the node for monitoring when Type is set to DOMAIN. Valid values:
	//
	//         	- IPV4
	//
	//         	- IPV6
	//
	// 	- PING
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
	// 	- TCP
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
	// This parameter is required.
	//
	// example:
	//
	// {\\"code\\":200,\\"path\\":\\"\\\\index.htm\\",\\"host\\":\\"aliyun.com\\"}
	MonitorExtendInfo *string `json:"MonitorExtendInfo,omitempty" xml:"MonitorExtendInfo,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// http
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The timeout period. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s AddDnsGtmMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDnsGtmMonitorRequest) GoString() string {
	return s.String()
}

func (s *AddDnsGtmMonitorRequest) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *AddDnsGtmMonitorRequest) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *AddDnsGtmMonitorRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *AddDnsGtmMonitorRequest) GetIspCityNode() []*AddDnsGtmMonitorRequestIspCityNode {
	return s.IspCityNode
}

func (s *AddDnsGtmMonitorRequest) GetLang() *string {
	return s.Lang
}

func (s *AddDnsGtmMonitorRequest) GetMonitorExtendInfo() *string {
	return s.MonitorExtendInfo
}

func (s *AddDnsGtmMonitorRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *AddDnsGtmMonitorRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *AddDnsGtmMonitorRequest) SetAddrPoolId(v string) *AddDnsGtmMonitorRequest {
	s.AddrPoolId = &v
	return s
}

func (s *AddDnsGtmMonitorRequest) SetEvaluationCount(v int32) *AddDnsGtmMonitorRequest {
	s.EvaluationCount = &v
	return s
}

func (s *AddDnsGtmMonitorRequest) SetInterval(v int32) *AddDnsGtmMonitorRequest {
	s.Interval = &v
	return s
}

func (s *AddDnsGtmMonitorRequest) SetIspCityNode(v []*AddDnsGtmMonitorRequestIspCityNode) *AddDnsGtmMonitorRequest {
	s.IspCityNode = v
	return s
}

func (s *AddDnsGtmMonitorRequest) SetLang(v string) *AddDnsGtmMonitorRequest {
	s.Lang = &v
	return s
}

func (s *AddDnsGtmMonitorRequest) SetMonitorExtendInfo(v string) *AddDnsGtmMonitorRequest {
	s.MonitorExtendInfo = &v
	return s
}

func (s *AddDnsGtmMonitorRequest) SetProtocolType(v string) *AddDnsGtmMonitorRequest {
	s.ProtocolType = &v
	return s
}

func (s *AddDnsGtmMonitorRequest) SetTimeout(v int32) *AddDnsGtmMonitorRequest {
	s.Timeout = &v
	return s
}

func (s *AddDnsGtmMonitorRequest) Validate() error {
	return dara.Validate(s)
}

type AddDnsGtmMonitorRequestIspCityNode struct {
	// The code of the city where the monitored node is deployed.
	//
	// example:
	//
	// 123
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The code of the Internet service provider (ISP) to which the monitored node belongs.
	//
	// example:
	//
	// 123
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
}

func (s AddDnsGtmMonitorRequestIspCityNode) String() string {
	return dara.Prettify(s)
}

func (s AddDnsGtmMonitorRequestIspCityNode) GoString() string {
	return s.String()
}

func (s *AddDnsGtmMonitorRequestIspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *AddDnsGtmMonitorRequestIspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *AddDnsGtmMonitorRequestIspCityNode) SetCityCode(v string) *AddDnsGtmMonitorRequestIspCityNode {
	s.CityCode = &v
	return s
}

func (s *AddDnsGtmMonitorRequestIspCityNode) SetIspCode(v string) *AddDnsGtmMonitorRequestIspCityNode {
	s.IspCode = &v
	return s
}

func (s *AddDnsGtmMonitorRequestIspCityNode) Validate() error {
	return dara.Validate(s)
}
