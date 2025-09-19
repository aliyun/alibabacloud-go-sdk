// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizedStrategyTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCustomizedStrategyTargetsResponseBody
	GetRequestId() *string
	SetStartegyTargets(v []*DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets) *DescribeCustomizedStrategyTargetsResponseBody
	GetStartegyTargets() []*DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets
}

type DescribeCustomizedStrategyTargetsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1EE7B150-D67E-53FD-A52D-3E8E669A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the servers to which custom policies are applied.
	StartegyTargets []*DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets `json:"StartegyTargets,omitempty" xml:"StartegyTargets,omitempty" type:"Repeated"`
}

func (s DescribeCustomizedStrategyTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizedStrategyTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedStrategyTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomizedStrategyTargetsResponseBody) GetStartegyTargets() []*DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets {
	return s.StartegyTargets
}

func (s *DescribeCustomizedStrategyTargetsResponseBody) SetRequestId(v string) *DescribeCustomizedStrategyTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomizedStrategyTargetsResponseBody) SetStartegyTargets(v []*DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets) *DescribeCustomizedStrategyTargetsResponseBody {
	s.StartegyTargets = v
	return s
}

func (s *DescribeCustomizedStrategyTargetsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets struct {
	// The ID of the server group.
	//
	// >  You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to query the IDs of server groups.
	//
	// example:
	//
	// 14590457
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the baseline check policy.
	//
	// example:
	//
	// 1884
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The name of the baseline check policy.
	//
	// example:
	//
	// win
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The UUID of the server.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// example:
	//
	// 2701ad2e-0e8f-428c-8812-ebb2686e****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets) GetStrategyName() *string {
	return s.StrategyName
}

func (s *DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets) SetGroupId(v int64) *DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets {
	s.GroupId = &v
	return s
}

func (s *DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets) SetStrategyId(v int64) *DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets {
	s.StrategyId = &v
	return s
}

func (s *DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets) SetStrategyName(v string) *DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets {
	s.StrategyName = &v
	return s
}

func (s *DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets) SetUuid(v string) *DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets {
	s.Uuid = &v
	return s
}

func (s *DescribeCustomizedStrategyTargetsResponseBodyStartegyTargets) Validate() error {
	return dara.Validate(s)
}
