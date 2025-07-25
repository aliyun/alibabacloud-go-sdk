// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsGtmMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEvaluationCount(v int32) *UpdateDnsGtmMonitorRequest
	GetEvaluationCount() *int32
	SetInterval(v int32) *UpdateDnsGtmMonitorRequest
	GetInterval() *int32
	SetIspCityNode(v []*UpdateDnsGtmMonitorRequestIspCityNode) *UpdateDnsGtmMonitorRequest
	GetIspCityNode() []*UpdateDnsGtmMonitorRequestIspCityNode
	SetLang(v string) *UpdateDnsGtmMonitorRequest
	GetLang() *string
	SetMonitorConfigId(v string) *UpdateDnsGtmMonitorRequest
	GetMonitorConfigId() *string
	SetMonitorExtendInfo(v string) *UpdateDnsGtmMonitorRequest
	GetMonitorExtendInfo() *string
	SetProtocolType(v string) *UpdateDnsGtmMonitorRequest
	GetProtocolType() *string
	SetTimeout(v int32) *UpdateDnsGtmMonitorRequest
	GetTimeout() *int32
}

type UpdateDnsGtmMonitorRequest struct {
	// The maximum number of consecutive exceptions detected. If the number of consecutive exceptions detected reaches the maximum number, the application service is deemed abnormal.
	//
	// example:
	//
	// 2
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The health check interval. Unit: seconds.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The monitored nodes.
	//
	// This parameter is required.
	IspCityNode []*UpdateDnsGtmMonitorRequestIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
	// The language of the values of specific response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the health check configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// MonitorConfigId1
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	// The extended information. The required parameters vary based on the health check protocol.
	//
	// 	- HTTP or HTTPS
	//
	//     	- port: the port that you want to check
	//
	//     	- host: the host settings
	//
	//     	- path: the URL path
	//
	//     	- code: the return code. If the return value of code is greater than the specified value, the health check result is deemed abnormal. For example, if code is set to 400 and the code 404 is returned, the health check result is deemed abnormal.
	//
	//     	- failureRate: the failure rate
	//
	//     	- sni: specifies whether to enable server name indication (SNI). This parameter is available only when ProtocolType is set to HTTPS. Valid values:
	//
	//         	- true: enables SNI.
	//
	//         	- false: disables SNI.
	//
	//     	- nodeType: the type of the monitoring node when the address pool type is domain name. Valid values:
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
	//     	- nodeType: the type of the monitoring node when the address pool type is domain name. Valid values:
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
	//     	- nodeType: the type of the monitoring node when the address pool type is domain name. Valid values:
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
	// The protocol used for the health check. Valid values:
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
	// example:
	//
	// 3000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateDnsGtmMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmMonitorRequest) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmMonitorRequest) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *UpdateDnsGtmMonitorRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *UpdateDnsGtmMonitorRequest) GetIspCityNode() []*UpdateDnsGtmMonitorRequestIspCityNode {
	return s.IspCityNode
}

func (s *UpdateDnsGtmMonitorRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDnsGtmMonitorRequest) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *UpdateDnsGtmMonitorRequest) GetMonitorExtendInfo() *string {
	return s.MonitorExtendInfo
}

func (s *UpdateDnsGtmMonitorRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *UpdateDnsGtmMonitorRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateDnsGtmMonitorRequest) SetEvaluationCount(v int32) *UpdateDnsGtmMonitorRequest {
	s.EvaluationCount = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequest) SetInterval(v int32) *UpdateDnsGtmMonitorRequest {
	s.Interval = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequest) SetIspCityNode(v []*UpdateDnsGtmMonitorRequestIspCityNode) *UpdateDnsGtmMonitorRequest {
	s.IspCityNode = v
	return s
}

func (s *UpdateDnsGtmMonitorRequest) SetLang(v string) *UpdateDnsGtmMonitorRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequest) SetMonitorConfigId(v string) *UpdateDnsGtmMonitorRequest {
	s.MonitorConfigId = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequest) SetMonitorExtendInfo(v string) *UpdateDnsGtmMonitorRequest {
	s.MonitorExtendInfo = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequest) SetProtocolType(v string) *UpdateDnsGtmMonitorRequest {
	s.ProtocolType = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequest) SetTimeout(v int32) *UpdateDnsGtmMonitorRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateDnsGtmMonitorRequestIspCityNode struct {
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

func (s UpdateDnsGtmMonitorRequestIspCityNode) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmMonitorRequestIspCityNode) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmMonitorRequestIspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *UpdateDnsGtmMonitorRequestIspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *UpdateDnsGtmMonitorRequestIspCityNode) SetCityCode(v string) *UpdateDnsGtmMonitorRequestIspCityNode {
	s.CityCode = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequestIspCityNode) SetIspCode(v string) *UpdateDnsGtmMonitorRequestIspCityNode {
	s.IspCode = &v
	return s
}

func (s *UpdateDnsGtmMonitorRequestIspCityNode) Validate() error {
	return dara.Validate(s)
}
