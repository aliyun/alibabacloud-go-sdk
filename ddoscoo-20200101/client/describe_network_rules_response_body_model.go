// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkRules(v []*DescribeNetworkRulesResponseBodyNetworkRules) *DescribeNetworkRulesResponseBody
	GetNetworkRules() []*DescribeNetworkRulesResponseBodyNetworkRules
	SetRequestId(v string) *DescribeNetworkRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeNetworkRulesResponseBody
	GetTotalCount() *int64
}

type DescribeNetworkRulesResponseBody struct {
	// The details of the port forwarding rules.
	NetworkRules []*DescribeNetworkRulesResponseBodyNetworkRules `json:"NetworkRules,omitempty" xml:"NetworkRules,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8597F235-FA5E-4FC7-BAD9-E4C0B01BC771
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned port forwarding rules.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNetworkRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRulesResponseBody) GetNetworkRules() []*DescribeNetworkRulesResponseBodyNetworkRules {
	return s.NetworkRules
}

func (s *DescribeNetworkRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeNetworkRulesResponseBody) SetNetworkRules(v []*DescribeNetworkRulesResponseBodyNetworkRules) *DescribeNetworkRulesResponseBody {
	s.NetworkRules = v
	return s
}

func (s *DescribeNetworkRulesResponseBody) SetRequestId(v string) *DescribeNetworkRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkRulesResponseBody) SetTotalCount(v int64) *DescribeNetworkRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNetworkRulesResponseBody) Validate() error {
	if s.NetworkRules != nil {
		for _, item := range s.NetworkRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkRulesResponseBodyNetworkRules struct {
	// The port of the origin server.
	//
	// example:
	//
	// 80
	BackendPort *int32 `json:"BackendPort,omitempty" xml:"BackendPort,omitempty"`
	// The forwarding port.
	//
	// example:
	//
	// 80
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the port forwarding rule is automatically created. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsAutoCreate *bool `json:"IsAutoCreate,omitempty" xml:"IsAutoCreate,omitempty"`
	// Indicates whether the payload filtering rule is enabled. Valid values:
	//
	// 	- 1: enabled.
	//
	// 	- 0: disabled.
	//
	// example:
	//
	// 1
	PayloadRuleEnable *int64 `json:"PayloadRuleEnable,omitempty" xml:"PayloadRuleEnable,omitempty"`
	// The forwarding protocol. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Indicates whether the traffic diversion switch is on. Valid values:
	//
	// 	- 0: on.
	//
	// 	- 1: off.
	//
	// example:
	//
	// 0
	ProxyEnable *int64 `json:"ProxyEnable,omitempty" xml:"ProxyEnable,omitempty"`
	// The status of traffic diversion. Valid values:
	//
	// 	- on: Traffic diversion takes effect.
	//
	// 	- off: Traffic diversion does not take effect.
	//
	// example:
	//
	// on
	ProxyStatus *string `json:"ProxyStatus,omitempty" xml:"ProxyStatus,omitempty"`
	// The IP addresses of origin servers.
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	// The remarks of the port forwarding rule.
	//
	// example:
	//
	// Test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s DescribeNetworkRulesResponseBodyNetworkRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRulesResponseBodyNetworkRules) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) GetBackendPort() *int32 {
	return s.BackendPort
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) GetIsAutoCreate() *bool {
	return s.IsAutoCreate
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) GetPayloadRuleEnable() *int64 {
	return s.PayloadRuleEnable
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) GetProxyEnable() *int64 {
	return s.ProxyEnable
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) GetProxyStatus() *string {
	return s.ProxyStatus
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) GetRealServers() []*string {
	return s.RealServers
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) GetRemark() *string {
	return s.Remark
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) SetBackendPort(v int32) *DescribeNetworkRulesResponseBodyNetworkRules {
	s.BackendPort = &v
	return s
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) SetFrontendPort(v int32) *DescribeNetworkRulesResponseBodyNetworkRules {
	s.FrontendPort = &v
	return s
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) SetInstanceId(v string) *DescribeNetworkRulesResponseBodyNetworkRules {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) SetIsAutoCreate(v bool) *DescribeNetworkRulesResponseBodyNetworkRules {
	s.IsAutoCreate = &v
	return s
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) SetPayloadRuleEnable(v int64) *DescribeNetworkRulesResponseBodyNetworkRules {
	s.PayloadRuleEnable = &v
	return s
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) SetProtocol(v string) *DescribeNetworkRulesResponseBodyNetworkRules {
	s.Protocol = &v
	return s
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) SetProxyEnable(v int64) *DescribeNetworkRulesResponseBodyNetworkRules {
	s.ProxyEnable = &v
	return s
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) SetProxyStatus(v string) *DescribeNetworkRulesResponseBodyNetworkRules {
	s.ProxyStatus = &v
	return s
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) SetRealServers(v []*string) *DescribeNetworkRulesResponseBodyNetworkRules {
	s.RealServers = v
	return s
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) SetRemark(v string) *DescribeNetworkRulesResponseBodyNetworkRules {
	s.Remark = &v
	return s
}

func (s *DescribeNetworkRulesResponseBodyNetworkRules) Validate() error {
	return dara.Validate(s)
}
