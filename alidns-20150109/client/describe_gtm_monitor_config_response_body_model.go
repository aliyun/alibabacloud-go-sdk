// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmMonitorConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeGtmMonitorConfigResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeGtmMonitorConfigResponseBody
	GetCreateTimestamp() *int64
	SetEvaluationCount(v int32) *DescribeGtmMonitorConfigResponseBody
	GetEvaluationCount() *int32
	SetInterval(v int32) *DescribeGtmMonitorConfigResponseBody
	GetInterval() *int32
	SetIspCityNodes(v *DescribeGtmMonitorConfigResponseBodyIspCityNodes) *DescribeGtmMonitorConfigResponseBody
	GetIspCityNodes() *DescribeGtmMonitorConfigResponseBodyIspCityNodes
	SetMonitorConfigId(v string) *DescribeGtmMonitorConfigResponseBody
	GetMonitorConfigId() *string
	SetMonitorExtendInfo(v string) *DescribeGtmMonitorConfigResponseBody
	GetMonitorExtendInfo() *string
	SetProtocolType(v string) *DescribeGtmMonitorConfigResponseBody
	GetProtocolType() *string
	SetRequestId(v string) *DescribeGtmMonitorConfigResponseBody
	GetRequestId() *string
	SetTimeout(v int32) *DescribeGtmMonitorConfigResponseBody
	GetTimeout() *int32
	SetUpdateTime(v string) *DescribeGtmMonitorConfigResponseBody
	GetUpdateTime() *string
	SetUpdateTimestamp(v int64) *DescribeGtmMonitorConfigResponseBody
	GetUpdateTimestamp() *int64
}

type DescribeGtmMonitorConfigResponseBody struct {
	// The time when the health check configuration was created.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp that indicates the time when the health check configuration was created.
	//
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The maximum number of consecutive exceptions detected. If the number of consecutive exceptions detected reaches the maximum number, the application service is deemed abnormal.
	//
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The health check interval. Unit: seconds. The value is 60.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The monitored nodes.
	IspCityNodes *DescribeGtmMonitorConfigResponseBodyIspCityNodes `json:"IspCityNodes,omitempty" xml:"IspCityNodes,omitempty" type:"Struct"`
	// The ID of the health check configuration.
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
	// example:
	//
	// {\\"code\\":200,\\"path\\":\\"\\\\index.htm\\",\\"host\\":\\"aliyun.com\\"}
	MonitorExtendInfo *string `json:"MonitorExtendInfo,omitempty" xml:"MonitorExtendInfo,omitempty"`
	// The protocol used for the health check.
	//
	// example:
	//
	// HTTP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The health check timeout period. Unit: milliseconds. Valid values: 2000, 3000, 5000, and 10000.
	//
	// example:
	//
	// 3000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The time when the health check configuration was last updated.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The timestamp that indicates the time when the health check configuration was last updated.
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s DescribeGtmMonitorConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmMonitorConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorConfigResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGtmMonitorConfigResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeGtmMonitorConfigResponseBody) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *DescribeGtmMonitorConfigResponseBody) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeGtmMonitorConfigResponseBody) GetIspCityNodes() *DescribeGtmMonitorConfigResponseBodyIspCityNodes {
	return s.IspCityNodes
}

func (s *DescribeGtmMonitorConfigResponseBody) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *DescribeGtmMonitorConfigResponseBody) GetMonitorExtendInfo() *string {
	return s.MonitorExtendInfo
}

func (s *DescribeGtmMonitorConfigResponseBody) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeGtmMonitorConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGtmMonitorConfigResponseBody) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DescribeGtmMonitorConfigResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeGtmMonitorConfigResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeGtmMonitorConfigResponseBody) SetCreateTime(v string) *DescribeGtmMonitorConfigResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetCreateTimestamp(v int64) *DescribeGtmMonitorConfigResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetEvaluationCount(v int32) *DescribeGtmMonitorConfigResponseBody {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetInterval(v int32) *DescribeGtmMonitorConfigResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetIspCityNodes(v *DescribeGtmMonitorConfigResponseBodyIspCityNodes) *DescribeGtmMonitorConfigResponseBody {
	s.IspCityNodes = v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetMonitorConfigId(v string) *DescribeGtmMonitorConfigResponseBody {
	s.MonitorConfigId = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetMonitorExtendInfo(v string) *DescribeGtmMonitorConfigResponseBody {
	s.MonitorExtendInfo = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetProtocolType(v string) *DescribeGtmMonitorConfigResponseBody {
	s.ProtocolType = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetRequestId(v string) *DescribeGtmMonitorConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetTimeout(v int32) *DescribeGtmMonitorConfigResponseBody {
	s.Timeout = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetUpdateTime(v string) *DescribeGtmMonitorConfigResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) SetUpdateTimestamp(v int64) *DescribeGtmMonitorConfigResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBody) Validate() error {
	if s.IspCityNodes != nil {
		if err := s.IspCityNodes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGtmMonitorConfigResponseBodyIspCityNodes struct {
	IspCityNode []*DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode `json:"IspCityNode,omitempty" xml:"IspCityNode,omitempty" type:"Repeated"`
}

func (s DescribeGtmMonitorConfigResponseBodyIspCityNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmMonitorConfigResponseBodyIspCityNodes) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodes) GetIspCityNode() []*DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode {
	return s.IspCityNode
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodes) SetIspCityNode(v []*DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) *DescribeGtmMonitorConfigResponseBodyIspCityNodes {
	s.IspCityNode = v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodes) Validate() error {
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

type DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode struct {
	// The code of the city where the monitored node is deployed.
	//
	// example:
	//
	// 503
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The display name of the city where the monitored node is deployed.
	//
	// example:
	//
	// Zhangjiakou
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// The code of the country where the monitored node is deployed.
	//
	// example:
	//
	// 001
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The display name of the country where the monitored node is deployed.
	//
	// example:
	//
	// China
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	// The code of the Internet service provider (ISP) to which the monitored node belongs.
	//
	// example:
	//
	// 465
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	// The display name of the ISP to which the monitored node belongs.
	//
	// example:
	//
	// Alibaba
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
}

func (s DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) GoString() string {
	return s.String()
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) GetCityCode() *string {
	return s.CityCode
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) GetCityName() *string {
	return s.CityName
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) GetCountryCode() *string {
	return s.CountryCode
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) GetCountryName() *string {
	return s.CountryName
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) GetIspCode() *string {
	return s.IspCode
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) GetIspName() *string {
	return s.IspName
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) SetCityCode(v string) *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode {
	s.CityCode = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) SetCityName(v string) *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode {
	s.CityName = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) SetCountryCode(v string) *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode {
	s.CountryCode = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) SetCountryName(v string) *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode {
	s.CountryName = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) SetIspCode(v string) *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode {
	s.IspCode = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) SetIspName(v string) *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode {
	s.IspName = &v
	return s
}

func (s *DescribeGtmMonitorConfigResponseBodyIspCityNodesIspCityNode) Validate() error {
	return dara.Validate(s)
}
