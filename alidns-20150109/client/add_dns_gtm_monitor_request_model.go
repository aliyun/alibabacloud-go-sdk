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
	// The ID of the address pool. You can call the [DescribeDnsGtmInstanceAddressPools](https://www.alibabacloud.com/help/en/dns/api-alidns-2015-01-09-describednsgtminstanceaddresspools) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// po**
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The number of consecutive health checks.
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
	// The list of monitoring nodes.
	//
	// This parameter is required.
	IspCityNode []*AddDnsGtmMonitorRequestIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
	// The language of the response. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The extended information. The parameters vary based on the protocol type.
	//
	// - HTTP or HTTPS
	//
	//   - port: The health check port.
	//
	//   - host: The Host header.
	//
	//   - path: The URL path.
	//
	//   - code: The health check is considered abnormal if the returned HTTP status code is greater than this value.
	//
	//   - failureRate: The failure rate.
	//
	//   - sni: Specifies whether to enable Server Name Indication (SNI). This parameter is used only when the health check protocol is HTTPS. Valid values:
	//
	//     - true
	//
	//     - false
	//
	//   - nodeType: The type of the monitoring node. This parameter is used when the address pool type is DOMAIN. Valid values:
	//
	//     - IPV4
	//
	//     - IPV6
	//
	// - PING
	//
	//   - failureRate: The failure rate.
	//
	//   - packetNum: The number of ping packets.
	//
	//   - packetLossRate: The packet loss rate.
	//
	//   - nodeType: The type of the monitoring node. This parameter is used when the address pool type is DOMAIN. Valid values:
	//
	//     - IPV4
	//
	//     - IPV6
	//
	// - TCP
	//
	//   - port: The health check port.
	//
	//   - failureRate: The failure rate.
	//
	//   - nodeType: The type of the monitoring node. This parameter is used when the address pool type is DOMAIN. Valid values:
	//
	//     - IPV4
	//
	//     - IPV6
	//
	// > This parameter must be a JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"failureRate":50,"port":80}
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
	// TCP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The timeout period. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3000
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

type AddDnsGtmMonitorRequestIspCityNode struct {
	// The city code of the monitoring node.
	//
	// example:
	//
	// 123
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The carrier code of the monitoring node.
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
