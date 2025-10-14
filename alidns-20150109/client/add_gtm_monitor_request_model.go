// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGtmMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddrPoolId(v string) *AddGtmMonitorRequest
	GetAddrPoolId() *string
	SetEvaluationCount(v int32) *AddGtmMonitorRequest
	GetEvaluationCount() *int32
	SetInterval(v int32) *AddGtmMonitorRequest
	GetInterval() *int32
	SetIspCityNode(v []*AddGtmMonitorRequestIspCityNode) *AddGtmMonitorRequest
	GetIspCityNode() []*AddGtmMonitorRequestIspCityNode
	SetLang(v string) *AddGtmMonitorRequest
	GetLang() *string
	SetMonitorExtendInfo(v string) *AddGtmMonitorRequest
	GetMonitorExtendInfo() *string
	SetProtocolType(v string) *AddGtmMonitorRequest
	GetProtocolType() *string
	SetTimeout(v int32) *AddGtmMonitorRequest
	GetTimeout() *int32
}

type AddGtmMonitorRequest struct {
	// The ID of the address pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The number of consecutive failures.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The health check interval. Unit: seconds. Set the value to 60.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The nodes for monitoring.
	//
	// This parameter is required.
	IspCityNode []*AddGtmMonitorRequestIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The extended information. The required parameters vary based on the health check protocol.
	//
	// HTTP or HTTPS
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
	// PING
	//
	// 	- packetNum: the number of ping packets
	//
	// 	- packetLossRate: the packet loss rate
	//
	// 	- failureRate: the failure rate
	//
	// TCP
	//
	// 	- port: the port that you want to check
	//
	// 	- failureRate: the failure rate
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"code\\":200,\\"path\\":\\"/index.htm\\",\\"host\\":\\"aliyun.com\\"}
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
	// HTTP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The health check timeout period. Unit: milliseconds. Valid values: 2000, 3000, 5000, and 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s AddGtmMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGtmMonitorRequest) GoString() string {
	return s.String()
}

func (s *AddGtmMonitorRequest) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *AddGtmMonitorRequest) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *AddGtmMonitorRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *AddGtmMonitorRequest) GetIspCityNode() []*AddGtmMonitorRequestIspCityNode {
	return s.IspCityNode
}

func (s *AddGtmMonitorRequest) GetLang() *string {
	return s.Lang
}

func (s *AddGtmMonitorRequest) GetMonitorExtendInfo() *string {
	return s.MonitorExtendInfo
}

func (s *AddGtmMonitorRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *AddGtmMonitorRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *AddGtmMonitorRequest) SetAddrPoolId(v string) *AddGtmMonitorRequest {
	s.AddrPoolId = &v
	return s
}

func (s *AddGtmMonitorRequest) SetEvaluationCount(v int32) *AddGtmMonitorRequest {
	s.EvaluationCount = &v
	return s
}

func (s *AddGtmMonitorRequest) SetInterval(v int32) *AddGtmMonitorRequest {
	s.Interval = &v
	return s
}

func (s *AddGtmMonitorRequest) SetIspCityNode(v []*AddGtmMonitorRequestIspCityNode) *AddGtmMonitorRequest {
	s.IspCityNode = v
	return s
}

func (s *AddGtmMonitorRequest) SetLang(v string) *AddGtmMonitorRequest {
	s.Lang = &v
	return s
}

func (s *AddGtmMonitorRequest) SetMonitorExtendInfo(v string) *AddGtmMonitorRequest {
	s.MonitorExtendInfo = &v
	return s
}

func (s *AddGtmMonitorRequest) SetProtocolType(v string) *AddGtmMonitorRequest {
	s.ProtocolType = &v
	return s
}

func (s *AddGtmMonitorRequest) SetTimeout(v int32) *AddGtmMonitorRequest {
	s.Timeout = &v
	return s
}

func (s *AddGtmMonitorRequest) Validate() error {
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

type AddGtmMonitorRequestIspCityNode struct {
	// The city code.
	//
	// Specify the parameter according to the value of CityCode returned by the DescribeGtmMonitorAvailableConfig operation.
	//
	// example:
	//
	// 503
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The Internet service provider (ISP) node. Specify the parameter according to the value of IspCode returned by the DescribeGtmMonitorAvailableConfig operation.
	//
	// 	- If the return value of GroupType for the DescribeGtmMonitorAvailableConfig operation is BGP or Overseas, IspCode is not required and is set to 465 by default.
	//
	// 	- If the return value of GroupType for the DescribeGtmMonitorAvailableConfig operation is not BGP or Overseas, IspCode is required. When IspCode is specified, CityCode is required.
	//
	// example:
	//
	// 465
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
}

func (s AddGtmMonitorRequestIspCityNode) String() string {
	return dara.Prettify(s)
}

func (s AddGtmMonitorRequestIspCityNode) GoString() string {
	return s.String()
}

func (s *AddGtmMonitorRequestIspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *AddGtmMonitorRequestIspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *AddGtmMonitorRequestIspCityNode) SetCityCode(v string) *AddGtmMonitorRequestIspCityNode {
	s.CityCode = &v
	return s
}

func (s *AddGtmMonitorRequestIspCityNode) SetIspCode(v string) *AddGtmMonitorRequestIspCityNode {
	s.IspCode = &v
	return s
}

func (s *AddGtmMonitorRequestIspCityNode) Validate() error {
	return dara.Validate(s)
}
