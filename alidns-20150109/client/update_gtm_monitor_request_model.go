// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGtmMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEvaluationCount(v int32) *UpdateGtmMonitorRequest
	GetEvaluationCount() *int32
	SetInterval(v int32) *UpdateGtmMonitorRequest
	GetInterval() *int32
	SetIspCityNode(v []*UpdateGtmMonitorRequestIspCityNode) *UpdateGtmMonitorRequest
	GetIspCityNode() []*UpdateGtmMonitorRequestIspCityNode
	SetLang(v string) *UpdateGtmMonitorRequest
	GetLang() *string
	SetMonitorConfigId(v string) *UpdateGtmMonitorRequest
	GetMonitorConfigId() *string
	SetMonitorExtendInfo(v string) *UpdateGtmMonitorRequest
	GetMonitorExtendInfo() *string
	SetProtocolType(v string) *UpdateGtmMonitorRequest
	GetProtocolType() *string
	SetTimeout(v int32) *UpdateGtmMonitorRequest
	GetTimeout() *int32
}

type UpdateGtmMonitorRequest struct {
	// The maximum number of consecutive exceptions detected. If the number of consecutive exceptions detected reaches the maximum number, the application service is deemed abnormal.
	//
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The health check interval. Unit: seconds. Set the value to 60.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The monitored nodes.
	//
	// This parameter is required.
	IspCityNode []*UpdateGtmMonitorRequestIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
	// The language of the values of specific response parameters.
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
	// 1234abc
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	// The extended information, that is, the parameters required for the protocol. Different protocols require different parameters:
	//
	// HTTP or HTTPS:
	//
	// 	- port: the port to check.
	//
	// 	- failureRate: the failure rate.
	//
	// 	- code: the status code threshold. If the returned status code is greater than the specified threshold, the application service is deemed abnormal. Valid values: 400 and 500.
	//
	// 	- host: the host configuration.
	//
	// 	- path: the health check URL.
	//
	// PING:
	//
	// 	- packetNum: the number of ping packets.
	//
	// 	- packetLossRate: the loss rate of ping packets.
	//
	// 	- failureRate: the failure rate.
	//
	// TCP:
	//
	// 	- port: the port to check.
	//
	// 	- failureRate: the failure rate.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"code\\":200,\\"path\\":\\"\\\\index.htm\\",\\"host\\":\\"aliyun.com\\"}
	MonitorExtendInfo *string `json:"MonitorExtendInfo,omitempty" xml:"MonitorExtendInfo,omitempty"`
	// The protocol used for the health check.
	//
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The health check timeout period. Unit: milliseconds. Valid values: 2000, 3000, 5000, and 10000.
	//
	// example:
	//
	// 3000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateGtmMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmMonitorRequest) GoString() string {
	return s.String()
}

func (s *UpdateGtmMonitorRequest) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *UpdateGtmMonitorRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *UpdateGtmMonitorRequest) GetIspCityNode() []*UpdateGtmMonitorRequestIspCityNode {
	return s.IspCityNode
}

func (s *UpdateGtmMonitorRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateGtmMonitorRequest) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *UpdateGtmMonitorRequest) GetMonitorExtendInfo() *string {
	return s.MonitorExtendInfo
}

func (s *UpdateGtmMonitorRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *UpdateGtmMonitorRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateGtmMonitorRequest) SetEvaluationCount(v int32) *UpdateGtmMonitorRequest {
	s.EvaluationCount = &v
	return s
}

func (s *UpdateGtmMonitorRequest) SetInterval(v int32) *UpdateGtmMonitorRequest {
	s.Interval = &v
	return s
}

func (s *UpdateGtmMonitorRequest) SetIspCityNode(v []*UpdateGtmMonitorRequestIspCityNode) *UpdateGtmMonitorRequest {
	s.IspCityNode = v
	return s
}

func (s *UpdateGtmMonitorRequest) SetLang(v string) *UpdateGtmMonitorRequest {
	s.Lang = &v
	return s
}

func (s *UpdateGtmMonitorRequest) SetMonitorConfigId(v string) *UpdateGtmMonitorRequest {
	s.MonitorConfigId = &v
	return s
}

func (s *UpdateGtmMonitorRequest) SetMonitorExtendInfo(v string) *UpdateGtmMonitorRequest {
	s.MonitorExtendInfo = &v
	return s
}

func (s *UpdateGtmMonitorRequest) SetProtocolType(v string) *UpdateGtmMonitorRequest {
	s.ProtocolType = &v
	return s
}

func (s *UpdateGtmMonitorRequest) SetTimeout(v int32) *UpdateGtmMonitorRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateGtmMonitorRequest) Validate() error {
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

type UpdateGtmMonitorRequestIspCityNode struct {
	// The code of the city where the monitored node is deployed.
	//
	// example:
	//
	// 572
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

func (s UpdateGtmMonitorRequestIspCityNode) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmMonitorRequestIspCityNode) GoString() string {
	return s.String()
}

func (s *UpdateGtmMonitorRequestIspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *UpdateGtmMonitorRequestIspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *UpdateGtmMonitorRequestIspCityNode) SetCityCode(v string) *UpdateGtmMonitorRequestIspCityNode {
	s.CityCode = &v
	return s
}

func (s *UpdateGtmMonitorRequestIspCityNode) SetIspCode(v string) *UpdateGtmMonitorRequestIspCityNode {
	s.IspCode = &v
	return s
}

func (s *UpdateGtmMonitorRequestIspCityNode) Validate() error {
	return dara.Validate(s)
}
