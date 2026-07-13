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
	// The number of consecutive health checks.
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
	// The list of city nodes for health checks.
	//
	// This parameter is required.
	IspCityNode []*UpdateDnsGtmMonitorRequestIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
	// The language of the response. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the health check configuration. You can call the [DescribeDnsGtmInstanceAddressPool](https://www.alibabacloud.com/help/en/dns/api-alidns-2015-01-09-describednsgtminstanceaddresspool) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Monito******
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	// The extended information. The required parameters vary based on the health check protocol.
	//
	// - HTTP(S):
	//
	//   - port: The port for the health check.
	//
	//   - host: The Host header.
	//
	//   - path: The URL path.
	//
	//   - code: The health check is considered abnormal if the returned status code is greater than the specified value. For example, if you set this parameter to 400, a returned status code of 404 is considered abnormal.
	//
	//   - failureRate: The failure rate.
	//
	//   - sni: Specifies whether to enable Server Name Indication (SNI). This parameter is available only for the HTTPS protocol.
	//
	//     - true: Enable SNI.
	//
	//     - false: Disable SNI.
	//
	//   - nodeType: The type of the node for health checks when the address pool type is DOMAIN.
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
	//   - nodeType: The type of the node for health checks when the address pool type is DOMAIN.
	//
	//     - IPV4
	//
	//     - IPV6
	//
	// - TCP:
	//
	//   - port: The port for the health check.
	//
	//   - failureRate: The failure rate.
	//
	//   - nodeType: The type of the node for health checks when the address pool type is DOMAIN.
	//
	//     - IPV4
	//
	//     - IPV6
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"code\\":200,\\"path\\":\\"\\\\index.htm\\",\\"host\\":\\"aliyun.com\\"}
	MonitorExtendInfo *string `json:"MonitorExtendInfo,omitempty" xml:"MonitorExtendInfo,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// http
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The timeout period for a health check. Unit: milliseconds.
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

type UpdateDnsGtmMonitorRequestIspCityNode struct {
	// The city code of the health check node.
	//
	// example:
	//
	// 123
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The carrier code of the health check node.
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
