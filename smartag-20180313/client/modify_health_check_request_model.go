// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHealthCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyHealthCheckRequest
	GetDescription() *string
	SetDstIpAddr(v string) *ModifyHealthCheckRequest
	GetDstIpAddr() *string
	SetDstPort(v int32) *ModifyHealthCheckRequest
	GetDstPort() *int32
	SetFailCountThreshold(v int32) *ModifyHealthCheckRequest
	GetFailCountThreshold() *int32
	SetHcInstanceId(v string) *ModifyHealthCheckRequest
	GetHcInstanceId() *string
	SetName(v string) *ModifyHealthCheckRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyHealthCheckRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyHealthCheckRequest
	GetOwnerId() *int64
	SetProbeCount(v int32) *ModifyHealthCheckRequest
	GetProbeCount() *int32
	SetProbeInterval(v int32) *ModifyHealthCheckRequest
	GetProbeInterval() *int32
	SetProbeTimeout(v int32) *ModifyHealthCheckRequest
	GetProbeTimeout() *int32
	SetRegionId(v string) *ModifyHealthCheckRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyHealthCheckRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyHealthCheckRequest
	GetResourceOwnerId() *int64
	SetRttFailThreshold(v int32) *ModifyHealthCheckRequest
	GetRttFailThreshold() *int32
	SetRttThreshold(v int32) *ModifyHealthCheckRequest
	GetRttThreshold() *int32
	SetSmartAGId(v string) *ModifyHealthCheckRequest
	GetSmartAGId() *string
	SetSrcIpAddr(v string) *ModifyHealthCheckRequest
	GetSrcIpAddr() *string
	SetSrcPort(v int32) *ModifyHealthCheckRequest
	GetSrcPort() *int32
	SetType(v string) *ModifyHealthCheckRequest
	GetType() *string
}

type ModifyHealthCheckRequest struct {
	// The description of the health check.
	//
	// The description must be 2 to 256 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination IP address of the health check.
	//
	// example:
	//
	// 192.XX.XX.1
	DstIpAddr *string `json:"DstIpAddr,omitempty" xml:"DstIpAddr,omitempty"`
	// The destination port of the health check.
	//
	// >  This parameter is not supported.
	//
	// example:
	//
	// 2233
	DstPort *int32 `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// The number of failed probes before a health check is declared failed.
	//
	// Valid values: **1*	- to **15**.
	//
	// Default value: **3**.
	//
	// example:
	//
	// 3
	FailCountThreshold *int32 `json:"FailCountThreshold,omitempty" xml:"FailCountThreshold,omitempty"`
	// The ID of the health check.
	//
	// This parameter is required.
	//
	// example:
	//
	// hc-k9id4loo3lup57****
	HcInstanceId *string `json:"HcInstanceId,omitempty" xml:"HcInstanceId,omitempty"`
	// The name of the health check.
	//
	// The name must be 2 to 100 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// sss333
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of probes that are performed per health check.
	//
	// Valid values: **1*	- to **20**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	ProbeCount *int32 `json:"ProbeCount,omitempty" xml:"ProbeCount,omitempty"`
	// The time interval at which probes are performed. The next probe does not start before the current one is complete.
	//
	// Valid values: **1000*	- to **60000**.
	//
	// Default value: **2000**.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 2000
	ProbeInterval *int32 `json:"ProbeInterval,omitempty" xml:"ProbeInterval,omitempty"`
	// The timeout period of the probe.
	//
	// Valid values: **10*	- to **30000**.
	//
	// Default value: **3000**.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 3000
	ProbeTimeout *int32 `json:"ProbeTimeout,omitempty" xml:"ProbeTimeout,omitempty"`
	// The region ID of the Smart Access Gateway (SAG) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The number of times that the maximum RTT is exceeded before an alert is triggered.
	//
	// Valid values: **1*	- to **15**.
	//
	// Default value: **3**.
	//
	// example:
	//
	// 3
	RttFailThreshold *int32 `json:"RttFailThreshold,omitempty" xml:"RttFailThreshold,omitempty"`
	// The maximum round-trip time (RTT).
	//
	// Valid values: **-1*	- and **1*	- to **5000**.
	//
	// Default value: **-1**. This value indicates that the RTT threshold is not specified.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 3
	RttThreshold *int32 `json:"RttThreshold,omitempty" xml:"RttThreshold,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-i0e85k06v1mzpo****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The source IP address of the health check.
	//
	// example:
	//
	// 192.XX.XX.1
	SrcIpAddr *string `json:"SrcIpAddr,omitempty" xml:"SrcIpAddr,omitempty"`
	// The source port of the health check.
	//
	// >  This parameter is not supported.
	//
	// example:
	//
	// 3333
	SrcPort *int32 `json:"SrcPort,omitempty" xml:"SrcPort,omitempty"`
	// The type of packet used in the health check. Only **ICMP_ECHO*	- is supported.
	//
	// example:
	//
	// ICMP_ECHO
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyHealthCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHealthCheckRequest) GoString() string {
	return s.String()
}

func (s *ModifyHealthCheckRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyHealthCheckRequest) GetDstIpAddr() *string {
	return s.DstIpAddr
}

func (s *ModifyHealthCheckRequest) GetDstPort() *int32 {
	return s.DstPort
}

func (s *ModifyHealthCheckRequest) GetFailCountThreshold() *int32 {
	return s.FailCountThreshold
}

func (s *ModifyHealthCheckRequest) GetHcInstanceId() *string {
	return s.HcInstanceId
}

func (s *ModifyHealthCheckRequest) GetName() *string {
	return s.Name
}

func (s *ModifyHealthCheckRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyHealthCheckRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyHealthCheckRequest) GetProbeCount() *int32 {
	return s.ProbeCount
}

func (s *ModifyHealthCheckRequest) GetProbeInterval() *int32 {
	return s.ProbeInterval
}

func (s *ModifyHealthCheckRequest) GetProbeTimeout() *int32 {
	return s.ProbeTimeout
}

func (s *ModifyHealthCheckRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHealthCheckRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyHealthCheckRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyHealthCheckRequest) GetRttFailThreshold() *int32 {
	return s.RttFailThreshold
}

func (s *ModifyHealthCheckRequest) GetRttThreshold() *int32 {
	return s.RttThreshold
}

func (s *ModifyHealthCheckRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifyHealthCheckRequest) GetSrcIpAddr() *string {
	return s.SrcIpAddr
}

func (s *ModifyHealthCheckRequest) GetSrcPort() *int32 {
	return s.SrcPort
}

func (s *ModifyHealthCheckRequest) GetType() *string {
	return s.Type
}

func (s *ModifyHealthCheckRequest) SetDescription(v string) *ModifyHealthCheckRequest {
	s.Description = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetDstIpAddr(v string) *ModifyHealthCheckRequest {
	s.DstIpAddr = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetDstPort(v int32) *ModifyHealthCheckRequest {
	s.DstPort = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetFailCountThreshold(v int32) *ModifyHealthCheckRequest {
	s.FailCountThreshold = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetHcInstanceId(v string) *ModifyHealthCheckRequest {
	s.HcInstanceId = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetName(v string) *ModifyHealthCheckRequest {
	s.Name = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetOwnerAccount(v string) *ModifyHealthCheckRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetOwnerId(v int64) *ModifyHealthCheckRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetProbeCount(v int32) *ModifyHealthCheckRequest {
	s.ProbeCount = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetProbeInterval(v int32) *ModifyHealthCheckRequest {
	s.ProbeInterval = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetProbeTimeout(v int32) *ModifyHealthCheckRequest {
	s.ProbeTimeout = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetRegionId(v string) *ModifyHealthCheckRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetResourceOwnerAccount(v string) *ModifyHealthCheckRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetResourceOwnerId(v int64) *ModifyHealthCheckRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetRttFailThreshold(v int32) *ModifyHealthCheckRequest {
	s.RttFailThreshold = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetRttThreshold(v int32) *ModifyHealthCheckRequest {
	s.RttThreshold = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetSmartAGId(v string) *ModifyHealthCheckRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetSrcIpAddr(v string) *ModifyHealthCheckRequest {
	s.SrcIpAddr = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetSrcPort(v int32) *ModifyHealthCheckRequest {
	s.SrcPort = &v
	return s
}

func (s *ModifyHealthCheckRequest) SetType(v string) *ModifyHealthCheckRequest {
	s.Type = &v
	return s
}

func (s *ModifyHealthCheckRequest) Validate() error {
	return dara.Validate(s)
}
