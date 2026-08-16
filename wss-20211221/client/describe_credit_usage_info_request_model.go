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
	// The JSON string of instance IDs. This parameter can be omitted when `UsageType=User`. Pass a credit package instance ID when `UsageType=CreditPackage`, or pass an `AgentId` when `UsageType=Agent`.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The usage type (UsageType). Valid values:
	//
	// 	- User: Returns the usage, remaining credits, and consumption trends of the current user\\"s active credit packages.
	//
	// 	- CreditPackage: Requires a CreditPackageId. Returns the total and remaining credits of the specified credit package.
	//
	// 	- Agent: Requires an AgentId. Returns the cumulative credit usage, cumulative allocated quota, and their percentages.
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

func (s *DescribeCreditUsageInfoRequest) SetUsageType(v string) *DescribeCreditUsageInfoRequest {
	s.UsageType = &v
	return s
}

func (s *DescribeCreditUsageInfoRequest) Validate() error {
	return dara.Validate(s)
}
