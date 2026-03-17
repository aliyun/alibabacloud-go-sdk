// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthCheckAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *DescribeHealthCheckAttributeResponseBody
	GetCreateTime() *int64
	SetDescription(v string) *DescribeHealthCheckAttributeResponseBody
	GetDescription() *string
	SetDstIpAddr(v string) *DescribeHealthCheckAttributeResponseBody
	GetDstIpAddr() *string
	SetDstPort(v int32) *DescribeHealthCheckAttributeResponseBody
	GetDstPort() *int32
	SetFailCountThreshold(v int32) *DescribeHealthCheckAttributeResponseBody
	GetFailCountThreshold() *int32
	SetHcInstanceId(v string) *DescribeHealthCheckAttributeResponseBody
	GetHcInstanceId() *string
	SetName(v string) *DescribeHealthCheckAttributeResponseBody
	GetName() *string
	SetProbeCount(v int32) *DescribeHealthCheckAttributeResponseBody
	GetProbeCount() *int32
	SetProbeInterval(v int32) *DescribeHealthCheckAttributeResponseBody
	GetProbeInterval() *int32
	SetProbeTimeout(v int32) *DescribeHealthCheckAttributeResponseBody
	GetProbeTimeout() *int32
	SetRequestId(v string) *DescribeHealthCheckAttributeResponseBody
	GetRequestId() *string
	SetRttFailThreshold(v int32) *DescribeHealthCheckAttributeResponseBody
	GetRttFailThreshold() *int32
	SetRttThreshold(v int32) *DescribeHealthCheckAttributeResponseBody
	GetRttThreshold() *int32
	SetSmartAGId(v string) *DescribeHealthCheckAttributeResponseBody
	GetSmartAGId() *string
	SetSrcIpAddr(v string) *DescribeHealthCheckAttributeResponseBody
	GetSrcIpAddr() *string
	SetSrcPort(v int32) *DescribeHealthCheckAttributeResponseBody
	GetSrcPort() *int32
	SetType(v string) *DescribeHealthCheckAttributeResponseBody
	GetType() *string
}

type DescribeHealthCheckAttributeResponseBody struct {
	// The time when the health check instance was created. This value is a UNIX timestamp.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1586759657000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the health check instance.
	//
	// example:
	//
	// hc-123
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination IP address of the health check.
	//
	// example:
	//
	// 192.XX.XX.1
	DstIpAddr *string `json:"DstIpAddr,omitempty" xml:"DstIpAddr,omitempty"`
	// The destination port of the health check instance.
	//
	// >  This feature is not supported.
	//
	// example:
	//
	// 1223
	DstPort *int32 `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// The maximum number of failed probes before the health check is declared failed.
	//
	// Valid values: **1*	- to **15**.
	//
	// Default value: **3**.
	//
	// example:
	//
	// 3
	FailCountThreshold *int32 `json:"FailCountThreshold,omitempty" xml:"FailCountThreshold,omitempty"`
	// The ID of the health check instance.
	//
	// example:
	//
	// hc-1k4ucuq77b56x4****
	HcInstanceId *string `json:"HcInstanceId,omitempty" xml:"HcInstanceId,omitempty"`
	// The name of the health check instance.
	//
	// example:
	//
	// bvt-test-03****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of probes performed per health check.
	//
	// Valid values: **1*	- to **20**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 3
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
	// Default value: **1000**.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// 1000
	ProbeTimeout *int32 `json:"ProbeTimeout,omitempty" xml:"ProbeTimeout,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DDA08B78-5634-4A83-94E4-5C58FD7EBA19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// Value values: **-1*	- and **1*	- to **5000**.
	//
	// Default value: **-1**. This value indicates that the RTT threshold is not specified.
	//
	// example:
	//
	// 300
	RttThreshold *int32 `json:"RttThreshold,omitempty" xml:"RttThreshold,omitempty"`
	// The ID of the SAG instance.
	//
	// example:
	//
	// sag-1um5x5nwhilymw****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The source IP address of the health check.
	//
	// example:
	//
	// 10.XX.XX.1
	SrcIpAddr *string `json:"SrcIpAddr,omitempty" xml:"SrcIpAddr,omitempty"`
	// The source port of the health check instance.
	//
	// >  This feature is not supported.
	//
	// example:
	//
	// 2334
	SrcPort *int32 `json:"SrcPort,omitempty" xml:"SrcPort,omitempty"`
	// The type of packets used in the health check.
	//
	// Only **ICMP_ECHO*	- is supported.
	//
	// example:
	//
	// ICMP_ECHO
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeHealthCheckAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthCheckAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHealthCheckAttributeResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeHealthCheckAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeHealthCheckAttributeResponseBody) GetDstIpAddr() *string {
	return s.DstIpAddr
}

func (s *DescribeHealthCheckAttributeResponseBody) GetDstPort() *int32 {
	return s.DstPort
}

func (s *DescribeHealthCheckAttributeResponseBody) GetFailCountThreshold() *int32 {
	return s.FailCountThreshold
}

func (s *DescribeHealthCheckAttributeResponseBody) GetHcInstanceId() *string {
	return s.HcInstanceId
}

func (s *DescribeHealthCheckAttributeResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeHealthCheckAttributeResponseBody) GetProbeCount() *int32 {
	return s.ProbeCount
}

func (s *DescribeHealthCheckAttributeResponseBody) GetProbeInterval() *int32 {
	return s.ProbeInterval
}

func (s *DescribeHealthCheckAttributeResponseBody) GetProbeTimeout() *int32 {
	return s.ProbeTimeout
}

func (s *DescribeHealthCheckAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHealthCheckAttributeResponseBody) GetRttFailThreshold() *int32 {
	return s.RttFailThreshold
}

func (s *DescribeHealthCheckAttributeResponseBody) GetRttThreshold() *int32 {
	return s.RttThreshold
}

func (s *DescribeHealthCheckAttributeResponseBody) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeHealthCheckAttributeResponseBody) GetSrcIpAddr() *string {
	return s.SrcIpAddr
}

func (s *DescribeHealthCheckAttributeResponseBody) GetSrcPort() *int32 {
	return s.SrcPort
}

func (s *DescribeHealthCheckAttributeResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeHealthCheckAttributeResponseBody) SetCreateTime(v int64) *DescribeHealthCheckAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) SetDescription(v string) *DescribeHealthCheckAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) SetDstIpAddr(v string) *DescribeHealthCheckAttributeResponseBody {
	s.DstIpAddr = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) SetDstPort(v int32) *DescribeHealthCheckAttributeResponseBody {
	s.DstPort = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) SetFailCountThreshold(v int32) *DescribeHealthCheckAttributeResponseBody {
	s.FailCountThreshold = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) SetHcInstanceId(v string) *DescribeHealthCheckAttributeResponseBody {
	s.HcInstanceId = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) SetName(v string) *DescribeHealthCheckAttributeResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) SetProbeCount(v int32) *DescribeHealthCheckAttributeResponseBody {
	s.ProbeCount = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) SetProbeInterval(v int32) *DescribeHealthCheckAttributeResponseBody {
	s.ProbeInterval = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) SetProbeTimeout(v int32) *DescribeHealthCheckAttributeResponseBody {
	s.ProbeTimeout = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) SetRequestId(v string) *DescribeHealthCheckAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) SetRttFailThreshold(v int32) *DescribeHealthCheckAttributeResponseBody {
	s.RttFailThreshold = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) SetRttThreshold(v int32) *DescribeHealthCheckAttributeResponseBody {
	s.RttThreshold = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) SetSmartAGId(v string) *DescribeHealthCheckAttributeResponseBody {
	s.SmartAGId = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) SetSrcIpAddr(v string) *DescribeHealthCheckAttributeResponseBody {
	s.SrcIpAddr = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) SetSrcPort(v int32) *DescribeHealthCheckAttributeResponseBody {
	s.SrcPort = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) SetType(v string) *DescribeHealthCheckAttributeResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeHealthCheckAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
