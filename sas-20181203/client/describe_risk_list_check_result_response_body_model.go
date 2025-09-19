// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskListCheckResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*DescribeRiskListCheckResultResponseBodyList) *DescribeRiskListCheckResultResponseBody
	GetList() []*DescribeRiskListCheckResultResponseBodyList
	SetRequestId(v string) *DescribeRiskListCheckResultResponseBody
	GetRequestId() *string
}

type DescribeRiskListCheckResultResponseBody struct {
	// The number of risk items for each cloud service.
	List []*DescribeRiskListCheckResultResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3BFB4989-A108-46A4-954E-FF7EF02D1078
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRiskListCheckResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskListCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskListCheckResultResponseBody) GetList() []*DescribeRiskListCheckResultResponseBodyList {
	return s.List
}

func (s *DescribeRiskListCheckResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRiskListCheckResultResponseBody) SetList(v []*DescribeRiskListCheckResultResponseBodyList) *DescribeRiskListCheckResultResponseBody {
	s.List = v
	return s
}

func (s *DescribeRiskListCheckResultResponseBody) SetRequestId(v string) *DescribeRiskListCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskListCheckResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRiskListCheckResultResponseBodyList struct {
	// The instance ID of the cloud service.
	//
	// example:
	//
	// rm-bp1e8t4q15sr3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The total number of risk items detected in the current cloud service.
	//
	// example:
	//
	// 3
	RiskCount *int64 `json:"riskCount,omitempty" xml:"riskCount,omitempty"`
}

func (s DescribeRiskListCheckResultResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskListCheckResultResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeRiskListCheckResultResponseBodyList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRiskListCheckResultResponseBodyList) GetRiskCount() *int64 {
	return s.RiskCount
}

func (s *DescribeRiskListCheckResultResponseBodyList) SetInstanceId(v string) *DescribeRiskListCheckResultResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *DescribeRiskListCheckResultResponseBodyList) SetRiskCount(v int64) *DescribeRiskListCheckResultResponseBodyList {
	s.RiskCount = &v
	return s
}

func (s *DescribeRiskListCheckResultResponseBodyList) Validate() error {
	return dara.Validate(s)
}
