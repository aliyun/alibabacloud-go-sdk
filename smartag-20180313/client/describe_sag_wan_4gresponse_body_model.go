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
	// The IP address of the 4G SIM card.
	//
	// example:
	//
	// 192.XX.XX.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The MAC address of the 4G SIM card.
	//
	// example:
	//
	// c4:00:ad:a2:f5:****
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The priority of the 4G SIM card. The default value is **99**, which indicates the lowest priority and cannot be modified.
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
	// The status of the 4G SIM card. Valid value:
	//
	// 	- **Normal**: The 4G SIM card is available.
	//
	// 	- **Abnormal**: The 4G SIM card encountered an error.
	//
	// 	- **Unavailable**: No 4G SIM card is inserted.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The signal strength of the 4G network. Valid values:
	//
	// 	- **High**: strong signals.
	//
	// 	- **Middle**: medium-strength signals.
	//
	// 	- **Low**: weak signals.
	//
	// 	- **Unavailable**: no signal.
	//
	// example:
	//
	// High
	Strength *string `json:"Strength,omitempty" xml:"Strength,omitempty"`
	// The network status of the 4G SIM card. Valid values:
	//
	// 	- **active**: The 4G SIM card is used to establish the active connection. Network traffic is transmitted over the 4G SIM card first.
	//
	// 	- **standby**: The 4G SIM card is used to establish a standby connection. Network traffic is not transmitted over the 4G SIM card unless the active connection fails.
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
