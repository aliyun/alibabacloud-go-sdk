// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientConfStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeClientConfStrategyResponseBody
	GetRequestId() *string
	SetTargetList(v []*DescribeClientConfStrategyResponseBodyTargetList) *DescribeClientConfStrategyResponseBody
	GetTargetList() []*DescribeClientConfStrategyResponseBodyTargetList
	SetTotalCount(v int32) *DescribeClientConfStrategyResponseBody
	GetTotalCount() *int32
}

type DescribeClientConfStrategyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5BD95679-D63A-4151-97D0-188432F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the configurations.
	TargetList []*DescribeClientConfStrategyResponseBodyTargetList `json:"TargetList,omitempty" xml:"TargetList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeClientConfStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientConfStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientConfStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClientConfStrategyResponseBody) GetTargetList() []*DescribeClientConfStrategyResponseBodyTargetList {
	return s.TargetList
}

func (s *DescribeClientConfStrategyResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeClientConfStrategyResponseBody) SetRequestId(v string) *DescribeClientConfStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClientConfStrategyResponseBody) SetTargetList(v []*DescribeClientConfStrategyResponseBodyTargetList) *DescribeClientConfStrategyResponseBody {
	s.TargetList = v
	return s
}

func (s *DescribeClientConfStrategyResponseBody) SetTotalCount(v int32) *DescribeClientConfStrategyResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeClientConfStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClientConfStrategyResponseBodyTargetList struct {
	// The UUID of the Security Center agent.
	//
	// example:
	//
	// 2b1753a6-04d9-448e-ad17-7abdf19f****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeClientConfStrategyResponseBodyTargetList) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientConfStrategyResponseBodyTargetList) GoString() string {
	return s.String()
}

func (s *DescribeClientConfStrategyResponseBodyTargetList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeClientConfStrategyResponseBodyTargetList) SetUuid(v string) *DescribeClientConfStrategyResponseBodyTargetList {
	s.Uuid = &v
	return s
}

func (s *DescribeClientConfStrategyResponseBodyTargetList) Validate() error {
	return dara.Validate(s)
}
