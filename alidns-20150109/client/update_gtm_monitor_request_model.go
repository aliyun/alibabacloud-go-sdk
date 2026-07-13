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
	// The number of consecutive health checks.
	//
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The interval between health checks. Unit: seconds. The value must be 60.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The list of monitoring nodes.
	//
	// This parameter is required.
	IspCityNode []*UpdateGtmMonitorRequestIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
	// The language of the response.
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
	// The extended information. The parameters vary based on the protocol.
	//
	// HTTP and HTTPS:
	//
	// - port: The health check port.
	//
	// - failureRate: The failure rate.
	//
	// - code: The return code. A response with a status code greater than the specified value is considered abnormal. Valid values: 400 and 500.
	//
	// - host: The host settings.
	//
	// - path: The URL path.
	//
	// PING:
	//
	// - packetNum: The number of ping packets.
	//
	// - packetLossRate: The packet loss rate.
	//
	// - failureRate: The failure rate.
	//
	// TCP:
	//
	// - port: The health check port.
	//
	// - failureRate: The failure rate.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"code\\":200,\\"path\\":\\"\\\\index.htm\\",\\"host\\":\\"aliyun.com\\"}
	MonitorExtendInfo *string `json:"MonitorExtendInfo,omitempty" xml:"MonitorExtendInfo,omitempty"`
	// The health check protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The timeout period for a health check. Unit: milliseconds. Valid values: 2000, 3000, 5000, and 10000.
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
	// The city code.
	//
	// example:
	//
	// 572
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// - For more information about the valid values, see the response of the DescribeGtmMonitorAvailableConfig operation.
	//
	// - If GroupType is set to Border Gateway Protocol (BGP) or Overseas, IspCityNode.N.IspCode is optional. The default value is 465.
	//
	// - If GroupType is not set to BGP or Overseas, IspCityNode.N.IspCode is required. You must specify a value that matches the value of IspCityNode.N.CityCode.
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
