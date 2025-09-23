// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLayer4RuleAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v []*DescribeLayer4RuleAttributesResponseBodyListeners) *DescribeLayer4RuleAttributesResponseBody
	GetListeners() []*DescribeLayer4RuleAttributesResponseBodyListeners
	SetRequestId(v string) *DescribeLayer4RuleAttributesResponseBody
	GetRequestId() *string
}

type DescribeLayer4RuleAttributesResponseBody struct {
	Listeners []*DescribeLayer4RuleAttributesResponseBodyListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBody) GetListeners() []*DescribeLayer4RuleAttributesResponseBodyListeners {
	return s.Listeners
}

func (s *DescribeLayer4RuleAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLayer4RuleAttributesResponseBody) SetListeners(v []*DescribeLayer4RuleAttributesResponseBodyListeners) *DescribeLayer4RuleAttributesResponseBody {
	s.Listeners = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBody) SetRequestId(v string) *DescribeLayer4RuleAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLayer4RuleAttributesResponseBodyListeners struct {
	Config *DescribeLayer4RuleAttributesResponseBodyListenersConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// example:
	//
	// 233
	FrontendPort *int32 `json:"FrontendPort,omitempty" xml:"FrontendPort,omitempty"`
	// example:
	//
	// ddoscoo-cn-XXXXX
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListeners) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListeners) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) GetConfig() *DescribeLayer4RuleAttributesResponseBodyListenersConfig {
	return s.Config
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) GetFrontendPort() *int32 {
	return s.FrontendPort
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) SetConfig(v *DescribeLayer4RuleAttributesResponseBodyListenersConfig) *DescribeLayer4RuleAttributesResponseBodyListeners {
	s.Config = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) SetFrontendPort(v int32) *DescribeLayer4RuleAttributesResponseBodyListeners {
	s.FrontendPort = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) SetInstanceId(v string) *DescribeLayer4RuleAttributesResponseBodyListeners {
	s.InstanceId = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) SetProtocol(v string) *DescribeLayer4RuleAttributesResponseBodyListeners {
	s.Protocol = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListeners) Validate() error {
	return dara.Validate(s)
}

type DescribeLayer4RuleAttributesResponseBodyListenersConfig struct {
	Cc *DescribeLayer4RuleAttributesResponseBodyListenersConfigCc `json:"Cc,omitempty" xml:"Cc,omitempty" type:"Struct"`
	// example:
	//
	// on
	NodataConn *string                                                            `json:"NodataConn,omitempty" xml:"NodataConn,omitempty"`
	PayloadLen *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen `json:"PayloadLen,omitempty" xml:"PayloadLen,omitempty" type:"Struct"`
	// example:
	//
	// 0
	PersistenceTimeout *int32                                                         `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	Sla                *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla    `json:"Sla,omitempty" xml:"Sla,omitempty" type:"Struct"`
	Slimit             *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit `json:"Slimit,omitempty" xml:"Slimit,omitempty" type:"Struct"`
	// example:
	//
	// on
	Synproxy *string `json:"Synproxy,omitempty" xml:"Synproxy,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfig) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) GetCc() *DescribeLayer4RuleAttributesResponseBodyListenersConfigCc {
	return s.Cc
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) GetNodataConn() *string {
	return s.NodataConn
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) GetPayloadLen() *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen {
	return s.PayloadLen
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) GetPersistenceTimeout() *int32 {
	return s.PersistenceTimeout
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) GetSla() *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla {
	return s.Sla
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) GetSlimit() *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	return s.Slimit
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) GetSynproxy() *string {
	return s.Synproxy
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) SetCc(v *DescribeLayer4RuleAttributesResponseBodyListenersConfigCc) *DescribeLayer4RuleAttributesResponseBodyListenersConfig {
	s.Cc = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) SetNodataConn(v string) *DescribeLayer4RuleAttributesResponseBodyListenersConfig {
	s.NodataConn = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) SetPayloadLen(v *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen) *DescribeLayer4RuleAttributesResponseBodyListenersConfig {
	s.PayloadLen = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) SetPersistenceTimeout(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfig {
	s.PersistenceTimeout = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) SetSla(v *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) *DescribeLayer4RuleAttributesResponseBodyListenersConfig {
	s.Sla = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) SetSlimit(v *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) *DescribeLayer4RuleAttributesResponseBodyListenersConfig {
	s.Slimit = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) SetSynproxy(v string) *DescribeLayer4RuleAttributesResponseBodyListenersConfig {
	s.Synproxy = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeLayer4RuleAttributesResponseBodyListenersConfigCc struct {
	Sblack []*DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack `json:"Sblack,omitempty" xml:"Sblack,omitempty" type:"Repeated"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigCc) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigCc) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCc) GetSblack() []*DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack {
	return s.Sblack
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCc) SetSblack(v []*DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) *DescribeLayer4RuleAttributesResponseBodyListenersConfigCc {
	s.Sblack = v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCc) Validate() error {
	return dara.Validate(s)
}

type DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack struct {
	// example:
	//
	// 5
	Cnt *int32 `json:"Cnt,omitempty" xml:"Cnt,omitempty"`
	// example:
	//
	// 60
	During *int32 `json:"During,omitempty" xml:"During,omitempty"`
	// example:
	//
	// 1800
	Expires *int32 `json:"Expires,omitempty" xml:"Expires,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) GetCnt() *int32 {
	return s.Cnt
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) GetDuring() *int32 {
	return s.During
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) GetExpires() *int32 {
	return s.Expires
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) GetType() *int32 {
	return s.Type
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) SetCnt(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack {
	s.Cnt = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) SetDuring(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack {
	s.During = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) SetExpires(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack {
	s.Expires = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) SetType(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack {
	s.Type = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigCcSblack) Validate() error {
	return dara.Validate(s)
}

type DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen struct {
	// example:
	//
	// 2
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// example:
	//
	// 1
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen) GetMax() *int32 {
	return s.Max
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen) GetMin() *int32 {
	return s.Min
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen) SetMax(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen {
	s.Max = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen) SetMin(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen {
	s.Min = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigPayloadLen) Validate() error {
	return dara.Validate(s)
}

type DescribeLayer4RuleAttributesResponseBodyListenersConfigSla struct {
	// example:
	//
	// 100
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// example:
	//
	// 0
	CpsEnable *int32 `json:"CpsEnable,omitempty" xml:"CpsEnable,omitempty"`
	// example:
	//
	// 1000
	Maxconn *int32 `json:"Maxconn,omitempty" xml:"Maxconn,omitempty"`
	// example:
	//
	// 0
	MaxconnEnable *int32 `json:"MaxconnEnable,omitempty" xml:"MaxconnEnable,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) GetCps() *int32 {
	return s.Cps
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) GetCpsEnable() *int32 {
	return s.CpsEnable
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) GetMaxconn() *int32 {
	return s.Maxconn
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) GetMaxconnEnable() *int32 {
	return s.MaxconnEnable
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) SetCps(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla {
	s.Cps = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) SetCpsEnable(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla {
	s.CpsEnable = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) SetMaxconn(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla {
	s.Maxconn = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) SetMaxconnEnable(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla {
	s.MaxconnEnable = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSla) Validate() error {
	return dara.Validate(s)
}

type DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit struct {
	// example:
	//
	// 0
	Bps *int64 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// example:
	//
	// 100
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// example:
	//
	// 0
	CpsEnable *int32 `json:"CpsEnable,omitempty" xml:"CpsEnable,omitempty"`
	// example:
	//
	// 2
	CpsMode *int32 `json:"CpsMode,omitempty" xml:"CpsMode,omitempty"`
	// example:
	//
	// 1000
	Maxconn *int32 `json:"Maxconn,omitempty" xml:"Maxconn,omitempty"`
	// example:
	//
	// 0
	MaxconnEnable *int32 `json:"MaxconnEnable,omitempty" xml:"MaxconnEnable,omitempty"`
	// example:
	//
	// 0
	Pps *int64 `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) GetBps() *int64 {
	return s.Bps
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) GetCps() *int32 {
	return s.Cps
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) GetCpsEnable() *int32 {
	return s.CpsEnable
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) GetCpsMode() *int32 {
	return s.CpsMode
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) GetMaxconn() *int32 {
	return s.Maxconn
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) GetMaxconnEnable() *int32 {
	return s.MaxconnEnable
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) GetPps() *int64 {
	return s.Pps
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetBps(v int64) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.Bps = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetCps(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.Cps = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetCpsEnable(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.CpsEnable = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetCpsMode(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.CpsMode = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetMaxconn(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.Maxconn = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetMaxconnEnable(v int32) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.MaxconnEnable = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) SetPps(v int64) *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit {
	s.Pps = &v
	return s
}

func (s *DescribeLayer4RuleAttributesResponseBodyListenersConfigSlimit) Validate() error {
	return dara.Validate(s)
}
