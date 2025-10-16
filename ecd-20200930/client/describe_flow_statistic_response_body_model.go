// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopCount(v int32) *DescribeFlowStatisticResponseBody
	GetDesktopCount() *int32
	SetDesktopFlowStatistic(v []*DescribeFlowStatisticResponseBodyDesktopFlowStatistic) *DescribeFlowStatisticResponseBody
	GetDesktopFlowStatistic() []*DescribeFlowStatisticResponseBodyDesktopFlowStatistic
	SetRequestId(v string) *DescribeFlowStatisticResponseBody
	GetRequestId() *string
}

type DescribeFlowStatisticResponseBody struct {
	// The number of available cloud computers in the office network.
	//
	// example:
	//
	// 10
	DesktopCount *int32 `json:"DesktopCount,omitempty" xml:"DesktopCount,omitempty"`
	// The traffic statistics.
	DesktopFlowStatistic []*DescribeFlowStatisticResponseBodyDesktopFlowStatistic `json:"DesktopFlowStatistic,omitempty" xml:"DesktopFlowStatistic,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 269BDB16-2CD8-4865-84BD-11C40BC2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFlowStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowStatisticResponseBody) GetDesktopCount() *int32 {
	return s.DesktopCount
}

func (s *DescribeFlowStatisticResponseBody) GetDesktopFlowStatistic() []*DescribeFlowStatisticResponseBodyDesktopFlowStatistic {
	return s.DesktopFlowStatistic
}

func (s *DescribeFlowStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFlowStatisticResponseBody) SetDesktopCount(v int32) *DescribeFlowStatisticResponseBody {
	s.DesktopCount = &v
	return s
}

func (s *DescribeFlowStatisticResponseBody) SetDesktopFlowStatistic(v []*DescribeFlowStatisticResponseBodyDesktopFlowStatistic) *DescribeFlowStatisticResponseBody {
	s.DesktopFlowStatistic = v
	return s
}

func (s *DescribeFlowStatisticResponseBody) SetRequestId(v string) *DescribeFlowStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowStatisticResponseBody) Validate() error {
	if s.DesktopFlowStatistic != nil {
		for _, item := range s.DesktopFlowStatistic {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFlowStatisticResponseBodyDesktopFlowStatistic struct {
	// The ID of the cloud computer.
	//
	// example:
	//
	// ecd-8bslxqq0csytn****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The name of the cloud computer.
	//
	// example:
	//
	// desktop-1
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The traffic amount. Unit: KB.
	//
	// example:
	//
	// 1000
	FlowIn *string `json:"FlowIn,omitempty" xml:"FlowIn,omitempty"`
	// The traffic ranking.
	//
	// example:
	//
	// 1
	FlowRank *int32 `json:"FlowRank,omitempty" xml:"FlowRank,omitempty"`
}

func (s DescribeFlowStatisticResponseBodyDesktopFlowStatistic) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowStatisticResponseBodyDesktopFlowStatistic) GoString() string {
	return s.String()
}

func (s *DescribeFlowStatisticResponseBodyDesktopFlowStatistic) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeFlowStatisticResponseBodyDesktopFlowStatistic) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeFlowStatisticResponseBodyDesktopFlowStatistic) GetFlowIn() *string {
	return s.FlowIn
}

func (s *DescribeFlowStatisticResponseBodyDesktopFlowStatistic) GetFlowRank() *int32 {
	return s.FlowRank
}

func (s *DescribeFlowStatisticResponseBodyDesktopFlowStatistic) SetDesktopId(v string) *DescribeFlowStatisticResponseBodyDesktopFlowStatistic {
	s.DesktopId = &v
	return s
}

func (s *DescribeFlowStatisticResponseBodyDesktopFlowStatistic) SetDesktopName(v string) *DescribeFlowStatisticResponseBodyDesktopFlowStatistic {
	s.DesktopName = &v
	return s
}

func (s *DescribeFlowStatisticResponseBodyDesktopFlowStatistic) SetFlowIn(v string) *DescribeFlowStatisticResponseBodyDesktopFlowStatistic {
	s.FlowIn = &v
	return s
}

func (s *DescribeFlowStatisticResponseBodyDesktopFlowStatistic) SetFlowRank(v int32) *DescribeFlowStatisticResponseBodyDesktopFlowStatistic {
	s.FlowRank = &v
	return s
}

func (s *DescribeFlowStatisticResponseBodyDesktopFlowStatistic) Validate() error {
	return dara.Validate(s)
}
