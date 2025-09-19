// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListStrategyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeWhiteListStrategyListResponseBody
	GetRequestId() *string
	SetStrategies(v []*DescribeWhiteListStrategyListResponseBodyStrategies) *DescribeWhiteListStrategyListResponseBody
	GetStrategies() []*DescribeWhiteListStrategyListResponseBodyStrategies
}

type DescribeWhiteListStrategyListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 571B2642-BF51-5BDD-906B-D2340DB9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the policies.
	Strategies []*DescribeWhiteListStrategyListResponseBodyStrategies `json:"Strategies,omitempty" xml:"Strategies,omitempty" type:"Repeated"`
}

func (s DescribeWhiteListStrategyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListStrategyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListStrategyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWhiteListStrategyListResponseBody) GetStrategies() []*DescribeWhiteListStrategyListResponseBodyStrategies {
	return s.Strategies
}

func (s *DescribeWhiteListStrategyListResponseBody) SetRequestId(v string) *DescribeWhiteListStrategyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhiteListStrategyListResponseBody) SetStrategies(v []*DescribeWhiteListStrategyListResponseBodyStrategies) *DescribeWhiteListStrategyListResponseBody {
	s.Strategies = v
	return s
}

func (s *DescribeWhiteListStrategyListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWhiteListStrategyListResponseBodyStrategies struct {
	// The status of the policy. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: learning
	//
	// 	- **2**: paused
	//
	// 	- **3**: learning completed
	//
	// 	- **4**: enabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// 8795555
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// test
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The intelligent learning duration. Unit: hour.
	//
	// example:
	//
	// 5
	StudyTime *int32 `json:"StudyTime,omitempty" xml:"StudyTime,omitempty"`
}

func (s DescribeWhiteListStrategyListResponseBodyStrategies) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListStrategyListResponseBodyStrategies) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListStrategyListResponseBodyStrategies) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeWhiteListStrategyListResponseBodyStrategies) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *DescribeWhiteListStrategyListResponseBodyStrategies) GetStrategyName() *string {
	return s.StrategyName
}

func (s *DescribeWhiteListStrategyListResponseBodyStrategies) GetStudyTime() *int32 {
	return s.StudyTime
}

func (s *DescribeWhiteListStrategyListResponseBodyStrategies) SetStatus(v int32) *DescribeWhiteListStrategyListResponseBodyStrategies {
	s.Status = &v
	return s
}

func (s *DescribeWhiteListStrategyListResponseBodyStrategies) SetStrategyId(v int64) *DescribeWhiteListStrategyListResponseBodyStrategies {
	s.StrategyId = &v
	return s
}

func (s *DescribeWhiteListStrategyListResponseBodyStrategies) SetStrategyName(v string) *DescribeWhiteListStrategyListResponseBodyStrategies {
	s.StrategyName = &v
	return s
}

func (s *DescribeWhiteListStrategyListResponseBodyStrategies) SetStudyTime(v int32) *DescribeWhiteListStrategyListResponseBodyStrategies {
	s.StudyTime = &v
	return s
}

func (s *DescribeWhiteListStrategyListResponseBodyStrategies) Validate() error {
	return dara.Validate(s)
}
