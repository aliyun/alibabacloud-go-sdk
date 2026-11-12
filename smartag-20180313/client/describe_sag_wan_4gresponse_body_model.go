// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagWan4GResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIp(v string) *DescribeSagWan4GResponseBody
	GetIp() *string
	SetMac(v string) *DescribeSagWan4GResponseBody
	GetMac() *string
	SetPriority(v int32) *DescribeSagWan4GResponseBody
	GetPriority() *int32
	SetRequestId(v string) *DescribeSagWan4GResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeSagWan4GResponseBody
	GetStatus() *string
	SetStrength(v string) *DescribeSagWan4GResponseBody
	GetStrength() *string
	SetTrafficState(v string) *DescribeSagWan4GResponseBody
	GetTrafficState() *string
}

type DescribeSagWan4GResponseBody struct {
	// The IP address of the 4G NIC.
	//
	// example:
	//
	// 192.XX.XX.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The MAC address of the 4G NIC.
	//
	// example:
	//
	// c4:00:ad:a2:f5:****
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The priority of the 4G NIC. The default value is **99**, which indicates the lowest priority. The priority cannot be modified.
	//
	// example:
	//
	// 99
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CE6642D4-21EB-4168-9BF9-F217953F9892
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the 4G NIC. Valid values:
	//
	// - **Normal**: The status is normal.
	//
	// - **Abnormal**: The status is abnormal.
	//
	// - **Unavailable**: No 4G NIC is available.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The 4G signal strength. Valid values:
	//
	// - **High**: Strong.
	//
	// - **Middle**: Medium.
	//
	// - **Low**: Weak.
	//
	// - **Unavailable**: No signal.
	//
	// example:
	//
	// High
	Strength *string `json:"Strength,omitempty" xml:"Strength,omitempty"`
	// The traffic status of the 4G NIC. Valid values:
	//
	// - **active**: The 4G NIC acts as the active link, and current traffic is transmitted through this 4G NIC.
	//
	// - **standby**: The 4G NIC acts as the standby link and does not forward traffic. If the active link fails, traffic is transmitted through this 4G NIC.
	//
	// example:
	//
	// active
	TrafficState *string `json:"TrafficState,omitempty" xml:"TrafficState,omitempty"`
}

func (s DescribeSagWan4GResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagWan4GResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagWan4GResponseBody) GetIp() *string {
	return s.Ip
}

func (s *DescribeSagWan4GResponseBody) GetMac() *string {
	return s.Mac
}

func (s *DescribeSagWan4GResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeSagWan4GResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagWan4GResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeSagWan4GResponseBody) GetStrength() *string {
	return s.Strength
}

func (s *DescribeSagWan4GResponseBody) GetTrafficState() *string {
	return s.TrafficState
}

func (s *DescribeSagWan4GResponseBody) SetIp(v string) *DescribeSagWan4GResponseBody {
	s.Ip = &v
	return s
}

func (s *DescribeSagWan4GResponseBody) SetMac(v string) *DescribeSagWan4GResponseBody {
	s.Mac = &v
	return s
}

func (s *DescribeSagWan4GResponseBody) SetPriority(v int32) *DescribeSagWan4GResponseBody {
	s.Priority = &v
	return s
}

func (s *DescribeSagWan4GResponseBody) SetRequestId(v string) *DescribeSagWan4GResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagWan4GResponseBody) SetStatus(v string) *DescribeSagWan4GResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSagWan4GResponseBody) SetStrength(v string) *DescribeSagWan4GResponseBody {
	s.Strength = &v
	return s
}

func (s *DescribeSagWan4GResponseBody) SetTrafficState(v string) *DescribeSagWan4GResponseBody {
	s.TrafficState = &v
	return s
}

func (s *DescribeSagWan4GResponseBody) Validate() error {
	return dara.Validate(s)
}
