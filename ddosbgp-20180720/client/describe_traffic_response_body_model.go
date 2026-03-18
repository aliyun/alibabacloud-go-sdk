// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrafficResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlowList(v []*DescribeTrafficResponseBodyFlowList) *DescribeTrafficResponseBody
	GetFlowList() []*DescribeTrafficResponseBodyFlowList
	SetRequestId(v string) *DescribeTrafficResponseBody
	GetRequestId() *string
}

type DescribeTrafficResponseBody struct {
	// The queried traffic statistics.
	FlowList []*DescribeTrafficResponseBodyFlowList `json:"FlowList,omitempty" xml:"FlowList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 6A507DC8-F657-4C13-84E2-D1D1B9400753
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTrafficResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponseBody) GetFlowList() []*DescribeTrafficResponseBodyFlowList {
	return s.FlowList
}

func (s *DescribeTrafficResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTrafficResponseBody) SetFlowList(v []*DescribeTrafficResponseBodyFlowList) *DescribeTrafficResponseBody {
	s.FlowList = v
	return s
}

func (s *DescribeTrafficResponseBody) SetRequestId(v string) *DescribeTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrafficResponseBody) Validate() error {
	if s.FlowList != nil {
		for _, item := range s.FlowList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTrafficResponseBodyFlowList struct {
	// The bandwidth of attack traffic. Unit: bit/s.
	//
	// >  This parameter is returned only if attack traffic exists.
	//
	// example:
	//
	// 0
	AttackBps *int64 `json:"AttackBps,omitempty" xml:"AttackBps,omitempty"`
	// The packet forwarding rate of attack traffic. Unit: packets per second.
	//
	// >  This parameter is returned only if attack traffic exists.
	//
	// example:
	//
	// 0
	AttackPps *int64 `json:"AttackPps,omitempty" xml:"AttackPps,omitempty"`
	// The type of the traffic statistics. Valid values:
	//
	// 	- **max**: the peak traffic within the specified interval
	//
	// 	- **avg**: the average traffic within the specified interval
	//
	// example:
	//
	// max
	FlowType *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	// The bandwidth of the total traffic. Unit: Kbit/s.
	//
	// example:
	//
	// 417
	Kbps *int32 `json:"Kbps,omitempty" xml:"Kbps,omitempty"`
	// The ID of the traffic statistics.
	//
	// example:
	//
	// 8e33f19e-5644-11eb-b5c1-d89d67182200
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The packet forwarding rate of the total traffic. Unit: packets per second.
	//
	// example:
	//
	// 274
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The time when the traffic statistics are calculated. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1620951900
	Time *int32 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeTrafficResponseBodyFlowList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrafficResponseBodyFlowList) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponseBodyFlowList) GetAttackBps() *int64 {
	return s.AttackBps
}

func (s *DescribeTrafficResponseBodyFlowList) GetAttackPps() *int64 {
	return s.AttackPps
}

func (s *DescribeTrafficResponseBodyFlowList) GetFlowType() *string {
	return s.FlowType
}

func (s *DescribeTrafficResponseBodyFlowList) GetKbps() *int32 {
	return s.Kbps
}

func (s *DescribeTrafficResponseBodyFlowList) GetName() *string {
	return s.Name
}

func (s *DescribeTrafficResponseBodyFlowList) GetPps() *int32 {
	return s.Pps
}

func (s *DescribeTrafficResponseBodyFlowList) GetTime() *int32 {
	return s.Time
}

func (s *DescribeTrafficResponseBodyFlowList) SetAttackBps(v int64) *DescribeTrafficResponseBodyFlowList {
	s.AttackBps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetAttackPps(v int64) *DescribeTrafficResponseBodyFlowList {
	s.AttackPps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetFlowType(v string) *DescribeTrafficResponseBodyFlowList {
	s.FlowType = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetKbps(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Kbps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetName(v string) *DescribeTrafficResponseBodyFlowList {
	s.Name = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetPps(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Pps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetTime(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Time = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) Validate() error {
	return dara.Validate(s)
}
