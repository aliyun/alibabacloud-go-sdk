// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHealthCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateHealthCheckRequest
	GetDescription() *string
	SetDstIpAddr(v string) *CreateHealthCheckRequest
	GetDstIpAddr() *string
	SetDstPort(v int32) *CreateHealthCheckRequest
	GetDstPort() *int32
	SetFailCountThreshold(v int32) *CreateHealthCheckRequest
	GetFailCountThreshold() *int32
	SetName(v string) *CreateHealthCheckRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateHealthCheckRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateHealthCheckRequest
	GetOwnerId() *int64
	SetProbeCount(v int32) *CreateHealthCheckRequest
	GetProbeCount() *int32
	SetProbeInterval(v int32) *CreateHealthCheckRequest
	GetProbeInterval() *int32
	SetProbeTimeout(v int32) *CreateHealthCheckRequest
	GetProbeTimeout() *int32
	SetRegionId(v string) *CreateHealthCheckRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateHealthCheckRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateHealthCheckRequest
	GetResourceOwnerId() *int64
	SetRttFailThreshold(v int32) *CreateHealthCheckRequest
	GetRttFailThreshold() *int32
	SetRttThreshold(v int32) *CreateHealthCheckRequest
	GetRttThreshold() *int32
	SetSmartAGId(v string) *CreateHealthCheckRequest
	GetSmartAGId() *string
	SetSrcIpAddr(v string) *CreateHealthCheckRequest
	GetSrcIpAddr() *string
	SetSrcPort(v int32) *CreateHealthCheckRequest
	GetSrcPort() *int32
	SetType(v string) *CreateHealthCheckRequest
	GetType() *string
}

type CreateHealthCheckRequest struct {
	// The description of the health check.
	//
	// The description must be 2 to 256 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// hc_123
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination IP address of the health check.
	//
	// This parameter is required.
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
	// 1333
	DstPort *int32 `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// The maximum number of failed probes before a health check is declared failed.
	//
	// Valid values: **1 to 15**.
	//
	// Default value: **3**.
	//
	// example:
	//
	// 3
	FailCountThreshold *int32 `json:"FailCountThreshold,omitempty" xml:"FailCountThreshold,omitempty"`
	// The name of the health check.
	//
	// The name must be 2 to 100 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// hc-123
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of probes performed per health check.
	//
	// Valid values: **1*	- to **20**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	ProbeCount *int32 `json:"ProbeCount,omitempty" xml:"ProbeCount,omitempty"`
	// The time interval at which probes are performed. The next probe does not start before the current one is completed.
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
	// The timeout period of a probe.
	//
	// Valid values: **10*	- to **30000**.
	//
	// Default value: **1000**.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1000
	ProbeTimeout *int32 `json:"ProbeTimeout,omitempty" xml:"ProbeTimeout,omitempty"`
	// The region ID of the SAG instance.
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
	// Default value: **-1**. This value indicates that the maximum RTT is not specified.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1000
	RttThreshold *int32 `json:"RttThreshold,omitempty" xml:"RttThreshold,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-1um5x5nwhilymw****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The source IP address of the health check.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.XX.XX.1
	SrcIpAddr *string `json:"SrcIpAddr,omitempty" xml:"SrcIpAddr,omitempty"`
	// The source port of the health check.
	//
	// >  This parameter is not supported.
	//
	// example:
	//
	// 1344
	SrcPort *int32 `json:"SrcPort,omitempty" xml:"SrcPort,omitempty"`
	// The type of packets used in the health check.
	//
	// >  Only **ICMP_ECHO*	- is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// ICMP_ECHO
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateHealthCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHealthCheckRequest) GoString() string {
	return s.String()
}

func (s *CreateHealthCheckRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateHealthCheckRequest) GetDstIpAddr() *string {
	return s.DstIpAddr
}

func (s *CreateHealthCheckRequest) GetDstPort() *int32 {
	return s.DstPort
}

func (s *CreateHealthCheckRequest) GetFailCountThreshold() *int32 {
	return s.FailCountThreshold
}

func (s *CreateHealthCheckRequest) GetName() *string {
	return s.Name
}

func (s *CreateHealthCheckRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateHealthCheckRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateHealthCheckRequest) GetProbeCount() *int32 {
	return s.ProbeCount
}

func (s *CreateHealthCheckRequest) GetProbeInterval() *int32 {
	return s.ProbeInterval
}

func (s *CreateHealthCheckRequest) GetProbeTimeout() *int32 {
	return s.ProbeTimeout
}

func (s *CreateHealthCheckRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHealthCheckRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateHealthCheckRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateHealthCheckRequest) GetRttFailThreshold() *int32 {
	return s.RttFailThreshold
}

func (s *CreateHealthCheckRequest) GetRttThreshold() *int32 {
	return s.RttThreshold
}

func (s *CreateHealthCheckRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *CreateHealthCheckRequest) GetSrcIpAddr() *string {
	return s.SrcIpAddr
}

func (s *CreateHealthCheckRequest) GetSrcPort() *int32 {
	return s.SrcPort
}

func (s *CreateHealthCheckRequest) GetType() *string {
	return s.Type
}

func (s *CreateHealthCheckRequest) SetDescription(v string) *CreateHealthCheckRequest {
	s.Description = &v
	return s
}

func (s *CreateHealthCheckRequest) SetDstIpAddr(v string) *CreateHealthCheckRequest {
	s.DstIpAddr = &v
	return s
}

func (s *CreateHealthCheckRequest) SetDstPort(v int32) *CreateHealthCheckRequest {
	s.DstPort = &v
	return s
}

func (s *CreateHealthCheckRequest) SetFailCountThreshold(v int32) *CreateHealthCheckRequest {
	s.FailCountThreshold = &v
	return s
}

func (s *CreateHealthCheckRequest) SetName(v string) *CreateHealthCheckRequest {
	s.Name = &v
	return s
}

func (s *CreateHealthCheckRequest) SetOwnerAccount(v string) *CreateHealthCheckRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateHealthCheckRequest) SetOwnerId(v int64) *CreateHealthCheckRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateHealthCheckRequest) SetProbeCount(v int32) *CreateHealthCheckRequest {
	s.ProbeCount = &v
	return s
}

func (s *CreateHealthCheckRequest) SetProbeInterval(v int32) *CreateHealthCheckRequest {
	s.ProbeInterval = &v
	return s
}

func (s *CreateHealthCheckRequest) SetProbeTimeout(v int32) *CreateHealthCheckRequest {
	s.ProbeTimeout = &v
	return s
}

func (s *CreateHealthCheckRequest) SetRegionId(v string) *CreateHealthCheckRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHealthCheckRequest) SetResourceOwnerAccount(v string) *CreateHealthCheckRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateHealthCheckRequest) SetResourceOwnerId(v int64) *CreateHealthCheckRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateHealthCheckRequest) SetRttFailThreshold(v int32) *CreateHealthCheckRequest {
	s.RttFailThreshold = &v
	return s
}

func (s *CreateHealthCheckRequest) SetRttThreshold(v int32) *CreateHealthCheckRequest {
	s.RttThreshold = &v
	return s
}

func (s *CreateHealthCheckRequest) SetSmartAGId(v string) *CreateHealthCheckRequest {
	s.SmartAGId = &v
	return s
}

func (s *CreateHealthCheckRequest) SetSrcIpAddr(v string) *CreateHealthCheckRequest {
	s.SrcIpAddr = &v
	return s
}

func (s *CreateHealthCheckRequest) SetSrcPort(v int32) *CreateHealthCheckRequest {
	s.SrcPort = &v
	return s
}

func (s *CreateHealthCheckRequest) SetType(v string) *CreateHealthCheckRequest {
	s.Type = &v
	return s
}

func (s *CreateHealthCheckRequest) Validate() error {
	return dara.Validate(s)
}
