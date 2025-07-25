// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmMonitorConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeDnsGtmMonitorConfigResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeDnsGtmMonitorConfigResponseBody
	GetCreateTimestamp() *int64
	SetEvaluationCount(v int32) *DescribeDnsGtmMonitorConfigResponseBody
	GetEvaluationCount() *int32
	SetInterval(v int32) *DescribeDnsGtmMonitorConfigResponseBody
	GetInterval() *int32
	SetIspCityNodes(v *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes) *DescribeDnsGtmMonitorConfigResponseBody
	GetIspCityNodes() *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes
	SetMonitorConfigId(v string) *DescribeDnsGtmMonitorConfigResponseBody
	GetMonitorConfigId() *string
	SetMonitorExtendInfo(v string) *DescribeDnsGtmMonitorConfigResponseBody
	GetMonitorExtendInfo() *string
	SetProtocolType(v string) *DescribeDnsGtmMonitorConfigResponseBody
	GetProtocolType() *string
	SetRequestId(v string) *DescribeDnsGtmMonitorConfigResponseBody
	GetRequestId() *string
	SetTimeout(v int32) *DescribeDnsGtmMonitorConfigResponseBody
	GetTimeout() *int32
	SetUpdateTime(v string) *DescribeDnsGtmMonitorConfigResponseBody
	GetUpdateTime() *string
	SetUpdateTimestamp(v int64) *DescribeDnsGtmMonitorConfigResponseBody
	GetUpdateTimestamp() *int64
}

type DescribeDnsGtmMonitorConfigResponseBody struct {
	// The time when the health check configuration was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the health check configuration was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The number of consecutive failures.
	//
	// example:
	//
	// 1
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The health check interval. Unit: seconds.
	//
	// example:
	//
	// 1
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The health check nodes.
	IspCityNodes *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes `json:"IspCityNodes,omitempty" xml:"IspCityNodes,omitempty" type:"Struct"`
	// The ID of the health check configuration.
	//
	// example:
	//
	// MonitorConfigId1
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
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
	//     	- nodeType: the type of the node for monitoring when the address pool type is domain name. Valid values:
	//
	//         	- IPV4
	//
	//         	- IPV6
	//
	// 	- PING:
	//
	//     	- failureRate: the failure rate
	//
	//     	- packetNum: the number of ping packets
	//
	//     	- packetLossRate: the loss rate of ping packets
	//
	//     	- nodeType: the type of the node for monitoring when the address pool type is domain name. Valid values:
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
	//     	- nodeType: the type of the node for monitoring when the address pool type is domain name. Valid values:
	//
	//         	- IPV4
	//
	//         	- IPV6
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
	// example:
	//
	// http
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timeout period. Unit: milliseconds.
	//
	// example:
	//
	// 3000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The time when the health check configuration was updated. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-01-03T08:57Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the health check configuration was updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s DescribeDnsGtmMonitorConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) GetIspCityNodes() *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes {
	return s.IspCityNodes
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) GetMonitorExtendInfo() *string {
	return s.MonitorExtendInfo
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetCreateTime(v string) *DescribeDnsGtmMonitorConfigResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetCreateTimestamp(v int64) *DescribeDnsGtmMonitorConfigResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetEvaluationCount(v int32) *DescribeDnsGtmMonitorConfigResponseBody {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetInterval(v int32) *DescribeDnsGtmMonitorConfigResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetIspCityNodes(v *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes) *DescribeDnsGtmMonitorConfigResponseBody {
	s.IspCityNodes = v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetMonitorConfigId(v string) *DescribeDnsGtmMonitorConfigResponseBody {
	s.MonitorConfigId = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetMonitorExtendInfo(v string) *DescribeDnsGtmMonitorConfigResponseBody {
	s.MonitorExtendInfo = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetProtocolType(v string) *DescribeDnsGtmMonitorConfigResponseBody {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetRequestId(v string) *DescribeDnsGtmMonitorConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetTimeout(v int32) *DescribeDnsGtmMonitorConfigResponseBody {
	s.Timeout = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetUpdateTime(v string) *DescribeDnsGtmMonitorConfigResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) SetUpdateTimestamp(v int64) *DescribeDnsGtmMonitorConfigResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes struct {
	IspCityNode []*DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes) GetIspCityNode() []*DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode {
	return s.IspCityNode
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes) SetIspCityNode(v []*DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes {
	s.IspCityNode = v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode struct {
	// The city code.
	//
	// example:
	//
	// 572
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The display name of the city.
	//
	// example:
	//
	// Qingdao
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// The code of the country or region.
	//
	// example:
	//
	// 001
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The display name of the country or region.
	//
	// example:
	//
	// China
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	// The Internet service provider (ISP) code.
	//
	// example:
	//
	// 123
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	// The display name of the ISP.
	//
	// example:
	//
	// Alibaba
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) GetCityName() *string {
	return s.CityName
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) GetCountryCode() *string {
	return s.CountryCode
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) GetCountryName() *string {
	return s.CountryName
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) GetIspName() *string {
	return s.IspName
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) SetCityCode(v string) *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode {
	s.CityCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) SetCityName(v string) *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode {
	s.CityName = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) SetCountryCode(v string) *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode {
	s.CountryCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) SetCountryName(v string) *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode {
	s.CountryName = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) SetIspCode(v string) *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode {
	s.IspCode = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) SetIspName(v string) *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode {
	s.IspName = &v
	return s
}

func (s *DescribeDnsGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) Validate() error {
	return dara.Validate(s)
}
