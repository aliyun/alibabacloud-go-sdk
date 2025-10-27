// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStrategyTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeStrategyTargetResponseBody
	GetRequestId() *string
	SetStrategyTargets(v []*DescribeStrategyTargetResponseBodyStrategyTargets) *DescribeStrategyTargetResponseBody
	GetStrategyTargets() []*DescribeStrategyTargetResponseBodyStrategyTargets
}

type DescribeStrategyTargetResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 001BAB34-D70A-54B0-B1D7-91B76DCDD8E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the assets to which the baseline check policy is applied.
	StrategyTargets []*DescribeStrategyTargetResponseBodyStrategyTargets `json:"StrategyTargets,omitempty" xml:"StrategyTargets,omitempty" type:"Repeated"`
}

func (s DescribeStrategyTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStrategyTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStrategyTargetResponseBody) GetStrategyTargets() []*DescribeStrategyTargetResponseBodyStrategyTargets {
	return s.StrategyTargets
}

func (s *DescribeStrategyTargetResponseBody) SetRequestId(v string) *DescribeStrategyTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStrategyTargetResponseBody) SetStrategyTargets(v []*DescribeStrategyTargetResponseBodyStrategyTargets) *DescribeStrategyTargetResponseBody {
	s.StrategyTargets = v
	return s
}

func (s *DescribeStrategyTargetResponseBody) Validate() error {
	if s.StrategyTargets != nil {
		for _, item := range s.StrategyTargets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStrategyTargetResponseBodyStrategyTargets struct {
	// The number of the assets that belong to the asset group.
	//
	// example:
	//
	// 85
	BindUuidCount *int32 `json:"BindUuidCount,omitempty" xml:"BindUuidCount,omitempty"`
	// Indicates whether the baseline check policy is applied to the asset group. Valid values:
	//
	// 	- **add**: The baseline check policy is applied to the asset group.
	//
	// 	- **del**: the baseline check policy is not applied to the asset group.
	//
	// example:
	//
	// add
	Flag *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The ID of the asset group to which the assets belong or the UUID of the asset.
	//
	// example:
	//
	// 9165712
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The method that is used to add the assets to the baseline check policy. Valid values:
	//
	// 	- **groupId**: the ID of the asset group
	//
	// 	- **uuid**: the UUID of the asset
	//
	// example:
	//
	// groupId
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeStrategyTargetResponseBodyStrategyTargets) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyTargetResponseBodyStrategyTargets) GoString() string {
	return s.String()
}

func (s *DescribeStrategyTargetResponseBodyStrategyTargets) GetBindUuidCount() *int32 {
	return s.BindUuidCount
}

func (s *DescribeStrategyTargetResponseBodyStrategyTargets) GetFlag() *string {
	return s.Flag
}

func (s *DescribeStrategyTargetResponseBodyStrategyTargets) GetTarget() *string {
	return s.Target
}

func (s *DescribeStrategyTargetResponseBodyStrategyTargets) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeStrategyTargetResponseBodyStrategyTargets) SetBindUuidCount(v int32) *DescribeStrategyTargetResponseBodyStrategyTargets {
	s.BindUuidCount = &v
	return s
}

func (s *DescribeStrategyTargetResponseBodyStrategyTargets) SetFlag(v string) *DescribeStrategyTargetResponseBodyStrategyTargets {
	s.Flag = &v
	return s
}

func (s *DescribeStrategyTargetResponseBodyStrategyTargets) SetTarget(v string) *DescribeStrategyTargetResponseBodyStrategyTargets {
	s.Target = &v
	return s
}

func (s *DescribeStrategyTargetResponseBodyStrategyTargets) SetTargetType(v string) *DescribeStrategyTargetResponseBodyStrategyTargets {
	s.TargetType = &v
	return s
}

func (s *DescribeStrategyTargetResponseBodyStrategyTargets) Validate() error {
	return dara.Validate(s)
}
