// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortFlowListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPortFlowList(v []*DescribePortFlowListResponseBodyPortFlowList) *DescribePortFlowListResponseBody
	GetPortFlowList() []*DescribePortFlowListResponseBodyPortFlowList
	SetRequestId(v string) *DescribePortFlowListResponseBody
	GetRequestId() *string
}

type DescribePortFlowListResponseBody struct {
	// The returned traffic data.
	PortFlowList []*DescribePortFlowListResponseBodyPortFlowList `json:"PortFlowList,omitempty" xml:"PortFlowList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// FFC77501-BDF8-4BC8-9BF5-B295FBC3189B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePortFlowListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePortFlowListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortFlowListResponseBody) GetPortFlowList() []*DescribePortFlowListResponseBodyPortFlowList {
	return s.PortFlowList
}

func (s *DescribePortFlowListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePortFlowListResponseBody) SetPortFlowList(v []*DescribePortFlowListResponseBodyPortFlowList) *DescribePortFlowListResponseBody {
	s.PortFlowList = v
	return s
}

func (s *DescribePortFlowListResponseBody) SetRequestId(v string) *DescribePortFlowListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortFlowListResponseBody) Validate() error {
	if s.PortFlowList != nil {
		for _, item := range s.PortFlowList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePortFlowListResponseBodyPortFlowList struct {
	// The bandwidth of attack traffic. Unit: bit/s.
	//
	// example:
	//
	// 0
	AttackBps *int64 `json:"AttackBps,omitempty" xml:"AttackBps,omitempty"`
	// The packet forwarding rate of attack traffic. Unit: pps.
	//
	// example:
	//
	// 0
	AttackPps *int64 `json:"AttackPps,omitempty" xml:"AttackPps,omitempty"`
	// The inbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 2176000
	InBps *int64 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	// The packet forwarding rate of inbound traffic. Unit: packets per second.
	//
	// example:
	//
	// 2934
	InPps *int64 `json:"InPps,omitempty" xml:"InPps,omitempty"`
	// The index number of the returned data.
	//
	// example:
	//
	// 0
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The outbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 4389
	OutBps *int64 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
	// The packet forwarding rate of outbound traffic. Unit: packets per second (pps).
	//
	// example:
	//
	// 5
	OutPps *int64 `json:"OutPps,omitempty" xml:"OutPps,omitempty"`
	// The source region of the traffic. Valid values:
	//
	// 	- **cn**: mainland China
	//
	// 	- **alb-ap-northeast-1-gf-x**: Japan (Tokyo)
	//
	// 	- **alb-ap-southeast-gf-x**: Singapore
	//
	// 	- **alb-cn-hongkong-gf-x**: Hong Kong (China)
	//
	// 	- **alb-eu-central-1-gf-x**: Germany (Frankfurt)
	//
	// 	- **alb-us-west-1-gf-x**: US (Silicon Valley)
	//
	// > The values except **cn*	- are returned only when **RegionId*	- is set to **ap-southeast-1**.
	//
	// example:
	//
	// cn
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SlaBpsDropBps  *int64  `json:"SlaBpsDropBps,omitempty" xml:"SlaBpsDropBps,omitempty"`
	SlaBpsDropPps  *int64  `json:"SlaBpsDropPps,omitempty" xml:"SlaBpsDropPps,omitempty"`
	SlaConnDropBps *int64  `json:"SlaConnDropBps,omitempty" xml:"SlaConnDropBps,omitempty"`
	SlaConnDropPps *int64  `json:"SlaConnDropPps,omitempty" xml:"SlaConnDropPps,omitempty"`
	SlaCpsDropBps  *int64  `json:"SlaCpsDropBps,omitempty" xml:"SlaCpsDropBps,omitempty"`
	SlaCpsDropPps  *int64  `json:"SlaCpsDropPps,omitempty" xml:"SlaCpsDropPps,omitempty"`
	SlaPpsDropBps  *int64  `json:"SlaPpsDropBps,omitempty" xml:"SlaPpsDropBps,omitempty"`
	SlaPpsDropPps  *int64  `json:"SlaPpsDropPps,omitempty" xml:"SlaPpsDropPps,omitempty"`
	// The time when the data was collected. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1582992000
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribePortFlowListResponseBodyPortFlowList) String() string {
	return dara.Prettify(s)
}

func (s DescribePortFlowListResponseBodyPortFlowList) GoString() string {
	return s.String()
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetAttackBps() *int64 {
	return s.AttackBps
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetAttackPps() *int64 {
	return s.AttackPps
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetInBps() *int64 {
	return s.InBps
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetInPps() *int64 {
	return s.InPps
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetIndex() *int64 {
	return s.Index
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetOutBps() *int64 {
	return s.OutBps
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetOutPps() *int64 {
	return s.OutPps
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetRegion() *string {
	return s.Region
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetSlaBpsDropBps() *int64 {
	return s.SlaBpsDropBps
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetSlaBpsDropPps() *int64 {
	return s.SlaBpsDropPps
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetSlaConnDropBps() *int64 {
	return s.SlaConnDropBps
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetSlaConnDropPps() *int64 {
	return s.SlaConnDropPps
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetSlaCpsDropBps() *int64 {
	return s.SlaCpsDropBps
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetSlaCpsDropPps() *int64 {
	return s.SlaCpsDropPps
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetSlaPpsDropBps() *int64 {
	return s.SlaPpsDropBps
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetSlaPpsDropPps() *int64 {
	return s.SlaPpsDropPps
}

func (s *DescribePortFlowListResponseBodyPortFlowList) GetTime() *int64 {
	return s.Time
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetAttackBps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.AttackBps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetAttackPps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.AttackPps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetInBps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.InBps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetInPps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.InPps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetIndex(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.Index = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetOutBps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.OutBps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetOutPps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.OutPps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetRegion(v string) *DescribePortFlowListResponseBodyPortFlowList {
	s.Region = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetSlaBpsDropBps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.SlaBpsDropBps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetSlaBpsDropPps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.SlaBpsDropPps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetSlaConnDropBps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.SlaConnDropBps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetSlaConnDropPps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.SlaConnDropPps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetSlaCpsDropBps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.SlaCpsDropBps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetSlaCpsDropPps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.SlaCpsDropPps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetSlaPpsDropBps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.SlaPpsDropBps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetSlaPpsDropPps(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.SlaPpsDropPps = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) SetTime(v int64) *DescribePortFlowListResponseBodyPortFlowList {
	s.Time = &v
	return s
}

func (s *DescribePortFlowListResponseBodyPortFlowList) Validate() error {
	return dara.Validate(s)
}
