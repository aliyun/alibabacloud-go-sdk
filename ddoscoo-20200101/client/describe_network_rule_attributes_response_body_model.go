// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkRuleAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkRuleAttributes(v []*DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) *DescribeNetworkRuleAttributesResponseBody
	GetNetworkRuleAttributes() []*DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes
	SetRequestId(v string) *DescribeNetworkRuleAttributesResponseBody
	GetRequestId() *string
}

type DescribeNetworkRuleAttributesResponseBody struct {
	// An array that consists of the mitigation settings of the port forwarding rule for a non-website service. The mitigation settings include session persistence and DDoS mitigation policies.
	NetworkRuleAttributes []*DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes `json:"NetworkRuleAttributes,omitempty" xml:"NetworkRuleAttributes,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// F9F2F77D-307C-4F15-8D02-AB5957EEBF97
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNetworkRuleAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponseBody) GetNetworkRuleAttributes() []*DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes {
	return s.NetworkRuleAttributes
}

func (s *DescribeNetworkRuleAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkRuleAttributesResponseBody) SetNetworkRuleAttributes(v []*DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) *DescribeNetworkRuleAttributesResponseBody {
	s.NetworkRuleAttributes = v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBody) SetRequestId(v string) *DescribeNetworkRuleAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes struct {
	// The mitigation settings of the port forwarding rule.
	Config *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The forwarding port.
	//
	// example:
	//
	// 8080
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) GetConfig() *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig {
	return s.Config
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) SetConfig(v *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes {
	s.Config = v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) SetFrontendPort(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes {
	s.FrontendPort = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) SetInstanceId(v string) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) SetProtocol(v string) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes {
	s.Protocol = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributes) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig struct {
	// The protection policy applied when the number of connections initiated from a source IP address frequently exceeds the limit.
	Cc *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc `json:"Cc,omitempty" xml:"Cc,omitempty" type:"Struct"`
	// The status of the Empty Connection switch. Valid values:
	//
	// 	- **on**: The switch is turned on.
	//
	// 	- **off**: The switch is turned off.
	//
	// example:
	//
	// off
	NodataConn *string `json:"NodataConn,omitempty" xml:"NodataConn,omitempty"`
	// The settings of the Packet Length Limit policy.
	PayloadLen *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen `json:"PayloadLen,omitempty" xml:"PayloadLen,omitempty" type:"Struct"`
	// The timeout period of session persistence. Valid values: **30*	- to **3600**. Unit: seconds. Default value: **0**, which indicates that session persistence is disabled.
	//
	// example:
	//
	// 0
	PersistenceTimeout *int32 `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	// The settings of the Speed Limit for Destination policy.
	Sla *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla `json:"Sla,omitempty" xml:"Sla,omitempty" type:"Struct"`
	// The settings of the Speed Limit for Source policy.
	Slimit *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit `json:"Slimit,omitempty" xml:"Slimit,omitempty" type:"Struct"`
	// The status of the False Source switch. Valid values:
	//
	// 	- **on**: The switch is turned on.
	//
	// 	- **off**: The switch is turned off.
	//
	// example:
	//
	// off
	Synproxy *string `json:"Synproxy,omitempty" xml:"Synproxy,omitempty"`
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) GetCc() *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc {
	return s.Cc
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) GetNodataConn() *string {
	return s.NodataConn
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) GetPayloadLen() *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen {
	return s.PayloadLen
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) GetPersistenceTimeout() *int32 {
	return s.PersistenceTimeout
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) GetSla() *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla {
	return s.Sla
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) GetSlimit() *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit {
	return s.Slimit
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) GetSynproxy() *string {
	return s.Synproxy
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) SetCc(v *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig {
	s.Cc = v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) SetNodataConn(v string) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig {
	s.NodataConn = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) SetPayloadLen(v *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig {
	s.PayloadLen = v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) SetPersistenceTimeout(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig {
	s.PersistenceTimeout = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) SetSla(v *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig {
	s.Sla = v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) SetSlimit(v *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig {
	s.Slimit = v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) SetSynproxy(v string) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig {
	s.Synproxy = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc struct {
	// The protection policy that a source IP address is added to the blacklist when the number of connections initiated from the IP address frequently exceeds the limit.
	Sblack []*DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack `json:"Sblack,omitempty" xml:"Sblack,omitempty" type:"Repeated"`
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc) GetSblack() []*DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack {
	return s.Sblack
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc) SetSblack(v []*DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc {
	s.Sblack = v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCc) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack struct {
	// The threshold that the number of connections initiated from a source IP address can exceed the limit. Set the value to **5**. If the number of connections initiated from a source IP address exceeds the limit five times during the check, the source IP address is added to the blacklist.
	//
	// example:
	//
	// 5
	Cnt *int32 `json:"Cnt,omitempty" xml:"Cnt,omitempty"`
	// The interval at which checks are performed. Set the value to **60**. Unit: seconds.
	//
	// example:
	//
	// 60
	During *int32 `json:"During,omitempty" xml:"During,omitempty"`
	// The validity period of the IP address in the blacklist. Valid values: **60*	- to **604800**. Unit: seconds.
	//
	// example:
	//
	// 600
	Expires *int32 `json:"Expires,omitempty" xml:"Expires,omitempty"`
	// The type of the limit that causes a source IP address to be added to the blacklist. Valid values:
	//
	// 	- **1**: Source New Connection Rate Limit
	//
	// 	- **2**: Source Concurrent Connection Rate Limit
	//
	// 	- **3**: PPS Limit for Source
	//
	// 	- **4**: Bandwidth Limit for Source
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) GetCnt() *int32 {
	return s.Cnt
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) GetDuring() *int32 {
	return s.During
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) GetExpires() *int32 {
	return s.Expires
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) GetType() *int32 {
	return s.Type
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) SetCnt(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack {
	s.Cnt = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) SetDuring(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack {
	s.During = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) SetExpires(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack {
	s.Expires = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) SetType(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack {
	s.Type = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigCcSblack) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen struct {
	// The maximum length of a packet. Valid values: **0*	- to **6000**. Unit: bytes.
	//
	// example:
	//
	// 6000
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// The minimum length of a packet. Valid values: **0*	- to **6000**. Unit: bytes.
	//
	// example:
	//
	// 0
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen) GetMax() *int32 {
	return s.Max
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen) GetMin() *int32 {
	return s.Min
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen) SetMax(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen {
	s.Max = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen) SetMin(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen {
	s.Min = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigPayloadLen) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla struct {
	// The maximum number of new connections per second that can be established over the port of the destination instance. Valid values: **100*	- to **100000**.
	//
	// example:
	//
	// 100000
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// The status of the Destination New Connection Rate Limit switch. Valid values:
	//
	// 	- **0**: The switch is turned off.
	//
	// 	- **1**: The switch is turned on.
	//
	// example:
	//
	// 1
	CpsEnable *int32 `json:"CpsEnable,omitempty" xml:"CpsEnable,omitempty"`
	// The maximum number of concurrent connections that can be established over the port of the destination instance. Valid values: **1000*	- to **1000000**.
	//
	// example:
	//
	// 1000000
	Maxconn *int32 `json:"Maxconn,omitempty" xml:"Maxconn,omitempty"`
	// The status of the Destination Concurrent Connection Rate Limit switch. Valid values:
	//
	// 	- **0**: The switch is turned off.
	//
	// 	- **1**: The switch is turned on.
	//
	// example:
	//
	// 0
	MaxconnEnable *int32 `json:"MaxconnEnable,omitempty" xml:"MaxconnEnable,omitempty"`
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) GetCps() *int32 {
	return s.Cps
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) GetCpsEnable() *int32 {
	return s.CpsEnable
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) GetMaxconn() *int32 {
	return s.Maxconn
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) GetMaxconnEnable() *int32 {
	return s.MaxconnEnable
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) SetCps(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla {
	s.Cps = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) SetCpsEnable(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla {
	s.CpsEnable = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) SetMaxconn(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla {
	s.Maxconn = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) SetMaxconnEnable(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla {
	s.MaxconnEnable = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSla) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit struct {
	// The bandwidth limit for a source IP address. Valid values: **1024*	- to **268435456**. Unit: bytes/s. Default value: **0**, which indicates that the bandwidth for a source IP address is unlimited.
	//
	// example:
	//
	// 0
	Bps *int64 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The maximum number of new connections per second that can be initiated from a source IP address. Valid values: **1*	- to **500000**.
	//
	// example:
	//
	// 0
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// The status of the Source New Connection Rate Limit switch. Valid values:
	//
	// 	- **0**: The switch is turned off.
	//
	// 	- **1**: The switch is turned on.
	//
	// example:
	//
	// 0
	CpsEnable *int32 `json:"CpsEnable,omitempty" xml:"CpsEnable,omitempty"`
	// The mode of the Source New Connection Rate Limit switch. Valid values:
	//
	// 	- **1**: the manual mode
	//
	// 	- **2**: the automatic mode
	//
	// example:
	//
	// 1
	CpsMode *int32 `json:"CpsMode,omitempty" xml:"CpsMode,omitempty"`
	// The maximum number of concurrent connections initiated from a source IP address. Valid values: **1*	- to **500000**.
	//
	// example:
	//
	// 0
	Maxconn *int32 `json:"Maxconn,omitempty" xml:"Maxconn,omitempty"`
	// The status of the Source Concurrent Connection Rate Limit switch. Valid values:
	//
	// 	- **0**: The switch is turned off.
	//
	// 	- **1**: The switch is turned on.
	//
	// example:
	//
	// 0
	MaxconnEnable *int32 `json:"MaxconnEnable,omitempty" xml:"MaxconnEnable,omitempty"`
	// The packets per second (pps) limit for a source IP address. Valid values: **1*	- to **100000**. Unit: packets/s. Default value: **0**, which indicates that the pps for a source IP address is unlimited.
	//
	// example:
	//
	// 0
	Pps *int64 `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) GetBps() *int64 {
	return s.Bps
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) GetCps() *int32 {
	return s.Cps
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) GetCpsEnable() *int32 {
	return s.CpsEnable
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) GetCpsMode() *int32 {
	return s.CpsMode
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) GetMaxconn() *int32 {
	return s.Maxconn
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) GetMaxconnEnable() *int32 {
	return s.MaxconnEnable
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) GetPps() *int64 {
	return s.Pps
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) SetBps(v int64) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit {
	s.Bps = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) SetCps(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit {
	s.Cps = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) SetCpsEnable(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit {
	s.CpsEnable = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) SetCpsMode(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit {
	s.CpsMode = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) SetMaxconn(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit {
	s.Maxconn = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) SetMaxconnEnable(v int32) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit {
	s.MaxconnEnable = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) SetPps(v int64) *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit {
	s.Pps = &v
	return s
}

func (s *DescribeNetworkRuleAttributesResponseBodyNetworkRuleAttributesConfigSlimit) Validate() error {
	return dara.Validate(s)
}
