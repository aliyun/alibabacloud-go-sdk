// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePatternTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPatternTypes(v []*DescribePatternTypesResponseBodyPatternTypes) *DescribePatternTypesResponseBody
	GetPatternTypes() []*DescribePatternTypesResponseBodyPatternTypes
	SetRequestId(v string) *DescribePatternTypesResponseBody
	GetRequestId() *string
}

type DescribePatternTypesResponseBody struct {
	// The instance types that meet the specified requirements.
	PatternTypes []*DescribePatternTypesResponseBodyPatternTypes `json:"PatternTypes,omitempty" xml:"PatternTypes,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePatternTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePatternTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePatternTypesResponseBody) GetPatternTypes() []*DescribePatternTypesResponseBodyPatternTypes {
	return s.PatternTypes
}

func (s *DescribePatternTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePatternTypesResponseBody) SetPatternTypes(v []*DescribePatternTypesResponseBodyPatternTypes) *DescribePatternTypesResponseBody {
	s.PatternTypes = v
	return s
}

func (s *DescribePatternTypesResponseBody) SetRequestId(v string) *DescribePatternTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePatternTypesResponseBody) Validate() error {
	if s.PatternTypes != nil {
		for _, item := range s.PatternTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePatternTypesResponseBodyPatternTypes struct {
	// The number of vCPUs that are assigned to the instance type.
	//
	// example:
	//
	// 4
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The level of the instance family.
	//
	// 	- EntryLevel: entry level (shared instance types) Instance types of this level are the most cost-effective but may not provide stable computing performance. Instance types of this level are suitable for business scenarios in which the CPU utilization is low. For more information, see [Shared instance families](https://help.aliyun.com/document_detail/108489.html).
	//
	// 	- EnterpriseLevel: enterprise level. Instance types of this level provide stable performance and dedicated resources and are suitable for scenarios that require high stability. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// 	- CreditEntryLevel: credit-based entry level (burstable instance types). CPU credits are used to ensure computing performance. Instance types of this level are suitable for scenarios in which the CPU utilization is low but may fluctuate in specific cases. For more information, see [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The name of the instance type.
	//
	// example:
	//
	// ecs.c7.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The instance family.
	//
	// example:
	//
	// ecs.c7
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// The memory size that are assigned to the instance type. Unit: GiB.
	//
	// example:
	//
	// 8
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s DescribePatternTypesResponseBodyPatternTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribePatternTypesResponseBodyPatternTypes) GoString() string {
	return s.String()
}

func (s *DescribePatternTypesResponseBodyPatternTypes) GetCores() *int32 {
	return s.Cores
}

func (s *DescribePatternTypesResponseBodyPatternTypes) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *DescribePatternTypesResponseBodyPatternTypes) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribePatternTypesResponseBodyPatternTypes) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribePatternTypesResponseBodyPatternTypes) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribePatternTypesResponseBodyPatternTypes) SetCores(v int32) *DescribePatternTypesResponseBodyPatternTypes {
	s.Cores = &v
	return s
}

func (s *DescribePatternTypesResponseBodyPatternTypes) SetInstanceFamilyLevel(v string) *DescribePatternTypesResponseBodyPatternTypes {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *DescribePatternTypesResponseBodyPatternTypes) SetInstanceType(v string) *DescribePatternTypesResponseBodyPatternTypes {
	s.InstanceType = &v
	return s
}

func (s *DescribePatternTypesResponseBodyPatternTypes) SetInstanceTypeFamily(v string) *DescribePatternTypesResponseBodyPatternTypes {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribePatternTypesResponseBodyPatternTypes) SetMemory(v float32) *DescribePatternTypesResponseBodyPatternTypes {
	s.Memory = &v
	return s
}

func (s *DescribePatternTypesResponseBodyPatternTypes) Validate() error {
	return dara.Validate(s)
}
