// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLayer4RulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v []*DescribeLayer4RulesResponseBodyListeners) *DescribeLayer4RulesResponseBody
	GetListeners() []*DescribeLayer4RulesResponseBodyListeners
	SetRequestId(v string) *DescribeLayer4RulesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeLayer4RulesResponseBody
	GetTotal() *int64
}

type DescribeLayer4RulesResponseBody struct {
	// Detailed configuration of port forwarding rules, including the forwarding port, forwarding protocol, and origin server addresses, etc.
	Listeners []*DescribeLayer4RulesResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	// The ID of the current request.
	//
	// example:
	//
	// 949919A2-6636-1444-9213-AB27DD88AAA8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of returned results.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeLayer4RulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesResponseBody) GetListeners() []*DescribeLayer4RulesResponseBodyListeners {
	return s.Listeners
}

func (s *DescribeLayer4RulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLayer4RulesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeLayer4RulesResponseBody) SetListeners(v []*DescribeLayer4RulesResponseBodyListeners) *DescribeLayer4RulesResponseBody {
	s.Listeners = v
	return s
}

func (s *DescribeLayer4RulesResponseBody) SetRequestId(v string) *DescribeLayer4RulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLayer4RulesResponseBody) SetTotal(v int64) *DescribeLayer4RulesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeLayer4RulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLayer4RulesResponseBodyListeners struct {
	// The origin server port.
	//
	// example:
	//
	// 233
	BackendPort *int32 `json:"BackendPort,omitempty" xml:"BackendPort,omitempty"`
	// The origin mode. Values:
	//
	// - **0**: Indicates the default origin mode.
	//
	// - **1**: Indicates the primary/backup origin mode.
	//
	// example:
	//
	// 0
	BakMode *int32 `json:"BakMode,omitempty" xml:"BakMode,omitempty"`
	// The currently effective origin server type. Values:
	//
	// - **1**: Indicates that the primary origin server settings are in effect (DDoS protection forwards business traffic to the primary origin server IP address).
	//
	// - **2**: Indicates that the backup origin server settings are in effect (DDoS protection forwards business traffic to the backup origin server IP address).
	//
	// example:
	//
	// 1
	CurrentIndex *int32 `json:"CurrentIndex,omitempty" xml:"CurrentIndex,omitempty"`
	// The IP address of the DDoS protection instance.
	//
	// example:
	//
	// 203.107.XX.XX
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// The forwarding port.
	//
	// example:
	//
	// 233
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The ID of the DDoS protection instance.
	//
	// example:
	//
	// ddoscoo-cn-zvp2ay9b****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the rule was automatically created. Values:
	//
	// - **true**: Indicates that the rule was automatically created by DDoS protection.
	//
	// - **false**: Indicates that the rule was manually created by you.
	//
	// example:
	//
	// false
	IsAutoCreate *bool `json:"IsAutoCreate,omitempty" xml:"IsAutoCreate,omitempty"`
	// Payload rule module switch. Values:
	//
	// - 1: Enabled
	//
	// - 0: Disabled
	//
	// example:
	//
	// 0
	PayloadRuleEnable *int64 `json:"PayloadRuleEnable,omitempty" xml:"PayloadRuleEnable,omitempty"`
	// The type of forwarding protocol.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Traffic diversion switch. Values:
	//
	// - **0*	- Off.
	//
	// - **1*	- On.
	//
	// example:
	//
	// 0
	ProxyEnable *int32 `json:"ProxyEnable,omitempty" xml:"ProxyEnable,omitempty"`
	// Traffic diversion status. Values:
	//
	// - on: Diversion is effective
	//
	// - off: Diversion is ineffective
	//
	// example:
	//
	// on
	ProxyStatus *string `json:"ProxyStatus,omitempty" xml:"ProxyStatus,omitempty"`
	// The list of origin server IP addresses.
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	// The remarks for the port forwarding rule.
	//
	// example:
	//
	// test-remark
	Remark    *string                                            `json:"Remark,omitempty" xml:"Remark,omitempty"`
	UsTimeout *DescribeLayer4RulesResponseBodyListenersUsTimeout `json:"UsTimeout,omitempty" xml:"UsTimeout,omitempty" type:"Struct"`
}

func (s DescribeLayer4RulesResponseBodyListeners) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RulesResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesResponseBodyListeners) GetBackendPort() *int32 {
	return s.BackendPort
}

func (s *DescribeLayer4RulesResponseBodyListeners) GetBakMode() *int32 {
	return s.BakMode
}

func (s *DescribeLayer4RulesResponseBodyListeners) GetCurrentIndex() *int32 {
	return s.CurrentIndex
}

func (s *DescribeLayer4RulesResponseBodyListeners) GetEip() *string {
	return s.Eip
}

func (s *DescribeLayer4RulesResponseBodyListeners) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *DescribeLayer4RulesResponseBodyListeners) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLayer4RulesResponseBodyListeners) GetIsAutoCreate() *bool {
	return s.IsAutoCreate
}

func (s *DescribeLayer4RulesResponseBodyListeners) GetPayloadRuleEnable() *int64 {
	return s.PayloadRuleEnable
}

func (s *DescribeLayer4RulesResponseBodyListeners) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeLayer4RulesResponseBodyListeners) GetProxyEnable() *int32 {
	return s.ProxyEnable
}

func (s *DescribeLayer4RulesResponseBodyListeners) GetProxyStatus() *string {
	return s.ProxyStatus
}

func (s *DescribeLayer4RulesResponseBodyListeners) GetRealServers() []*string {
	return s.RealServers
}

func (s *DescribeLayer4RulesResponseBodyListeners) GetRemark() *string {
	return s.Remark
}

func (s *DescribeLayer4RulesResponseBodyListeners) GetUsTimeout() *DescribeLayer4RulesResponseBodyListenersUsTimeout {
	return s.UsTimeout
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetBackendPort(v int32) *DescribeLayer4RulesResponseBodyListeners {
	s.BackendPort = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetBakMode(v int32) *DescribeLayer4RulesResponseBodyListeners {
	s.BakMode = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetCurrentIndex(v int32) *DescribeLayer4RulesResponseBodyListeners {
	s.CurrentIndex = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetEip(v string) *DescribeLayer4RulesResponseBodyListeners {
	s.Eip = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetFrontendPort(v int32) *DescribeLayer4RulesResponseBodyListeners {
	s.FrontendPort = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetInstanceId(v string) *DescribeLayer4RulesResponseBodyListeners {
	s.InstanceId = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetIsAutoCreate(v bool) *DescribeLayer4RulesResponseBodyListeners {
	s.IsAutoCreate = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetPayloadRuleEnable(v int64) *DescribeLayer4RulesResponseBodyListeners {
	s.PayloadRuleEnable = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetProtocol(v string) *DescribeLayer4RulesResponseBodyListeners {
	s.Protocol = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetProxyEnable(v int32) *DescribeLayer4RulesResponseBodyListeners {
	s.ProxyEnable = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetProxyStatus(v string) *DescribeLayer4RulesResponseBodyListeners {
	s.ProxyStatus = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetRealServers(v []*string) *DescribeLayer4RulesResponseBodyListeners {
	s.RealServers = v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetRemark(v string) *DescribeLayer4RulesResponseBodyListeners {
	s.Remark = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) SetUsTimeout(v *DescribeLayer4RulesResponseBodyListenersUsTimeout) *DescribeLayer4RulesResponseBodyListeners {
	s.UsTimeout = v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListeners) Validate() error {
	return dara.Validate(s)
}

type DescribeLayer4RulesResponseBodyListenersUsTimeout struct {
	ConnectTimeout *int64 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	RsTimeout      *int64 `json:"RsTimeout,omitempty" xml:"RsTimeout,omitempty"`
}

func (s DescribeLayer4RulesResponseBodyListenersUsTimeout) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RulesResponseBodyListenersUsTimeout) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesResponseBodyListenersUsTimeout) GetConnectTimeout() *int64 {
	return s.ConnectTimeout
}

func (s *DescribeLayer4RulesResponseBodyListenersUsTimeout) GetRsTimeout() *int64 {
	return s.RsTimeout
}

func (s *DescribeLayer4RulesResponseBodyListenersUsTimeout) SetConnectTimeout(v int64) *DescribeLayer4RulesResponseBodyListenersUsTimeout {
	s.ConnectTimeout = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListenersUsTimeout) SetRsTimeout(v int64) *DescribeLayer4RulesResponseBodyListenersUsTimeout {
	s.RsTimeout = &v
	return s
}

func (s *DescribeLayer4RulesResponseBodyListenersUsTimeout) Validate() error {
	return dara.Validate(s)
}
