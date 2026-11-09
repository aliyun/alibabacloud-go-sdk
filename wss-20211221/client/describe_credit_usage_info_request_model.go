// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditUsageInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *DescribeCreditUsageInfoRequest
	GetBizType() *string
	SetInstanceIds(v []*string) *DescribeCreditUsageInfoRequest
	GetInstanceIds() []*string
	SetSceneType(v string) *DescribeCreditUsageInfoRequest
	GetSceneType() *string
	SetUsageType(v string) *DescribeCreditUsageInfoRequest
	GetUsageType() *string
}

type DescribeCreditUsageInfoRequest struct {
	// The business type.
	//
	// example:
	//
	// BUSINESS
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The JSON string of the instance ID list. This parameter can be omitted when `UsageType` is set to `User`. Set it to the credit package instance ID when `UsageType` is `CreditPackage`, and the agent ID when `UsageType` is `Agent`.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The operator Scenarios type. Valid values:
	//
	// - Not specified or NaturalMonthlyOwn: Natural monthly exclusive.
	//
	// - JzYearlyShare: Jz annual shared.
	//
	// - JzYearlyOwn: Jz annual exclusive.
	//
	// example:
	//
	// JzYearlyShare
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	// The usage type. Valid values:
	//
	// - User: Returns the used amount, remaining amount, and consumption trend of the currently active credit packages for the user.
	//
	// - CreditPackage: Requires a credit package ID. Returns the total amount and remaining amount of the specified credit package.
	//
	// - Agent: Requires an agent ID. Returns the cumulative credit usage, cumulative allocated quota, and the percentage between the two.
	//
	// - EndUser: Requires a Wuying ID. Returns the shared credit usage, pay-as-you-go credit usage, and total credit usage for the current period of the Wuying user.
	//
	// example:
	//
	// User
	UsageType *string `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
}

func (s DescribeCreditUsageInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditUsageInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCreditUsageInfoRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeCreditUsageInfoRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeCreditUsageInfoRequest) GetSceneType() *string {
	return s.SceneType
}

func (s *DescribeCreditUsageInfoRequest) GetUsageType() *string {
	return s.UsageType
}

func (s *DescribeCreditUsageInfoRequest) SetBizType(v string) *DescribeCreditUsageInfoRequest {
	s.BizType = &v
	return s
}

func (s *DescribeCreditUsageInfoRequest) SetInstanceIds(v []*string) *DescribeCreditUsageInfoRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeCreditUsageInfoRequest) SetSceneType(v string) *DescribeCreditUsageInfoRequest {
	s.SceneType = &v
	return s
}

func (s *DescribeCreditUsageInfoRequest) SetUsageType(v string) *DescribeCreditUsageInfoRequest {
	s.UsageType = &v
	return s
}

func (s *DescribeCreditUsageInfoRequest) Validate() error {
	return dara.Validate(s)
}
