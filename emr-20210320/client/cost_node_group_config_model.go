// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostNodeGroupConfig interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceTypes(v []*CostInstanceType) *CostNodeGroupConfig
	GetInstanceTypes() []*CostInstanceType
	SetMaximalNodeCount(v int32) *CostNodeGroupConfig
	GetMaximalNodeCount() *int32
	SetMinimalNodeCount(v int32) *CostNodeGroupConfig
	GetMinimalNodeCount() *int32
	SetNodeCount(v int32) *CostNodeGroupConfig
	GetNodeCount() *int32
	SetNodeGroupName(v string) *CostNodeGroupConfig
	GetNodeGroupName() *string
	SetNodeGroupType(v string) *CostNodeGroupConfig
	GetNodeGroupType() *string
	SetPaymentType(v string) *CostNodeGroupConfig
	GetPaymentType() *string
}

type CostNodeGroupConfig struct {
	// 实例类型列表。
	InstanceTypes []*CostInstanceType `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// 最大节点数限制。
	//
	// example:
	//
	// 100
	MaximalNodeCount *int32 `json:"MaximalNodeCount,omitempty" xml:"MaximalNodeCount,omitempty"`
	// 最小节点数限制。
	//
	// example:
	//
	// 0
	MinimalNodeCount *int32 `json:"MinimalNodeCount,omitempty" xml:"MinimalNodeCount,omitempty"`
	// 节点数。
	//
	// example:
	//
	// 1
	NodeCount     *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// 节点组类型。取值范围：
	//
	// - MASTER：管理类型节点组。
	//
	// - CORE：存储类型节点组。
	//
	// - TASK：计算类型节点组。
	//
	// example:
	//
	// CORE
	NodeGroupType *string `json:"NodeGroupType,omitempty" xml:"NodeGroupType,omitempty"`
	// 付费类型。
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
}

func (s CostNodeGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s CostNodeGroupConfig) GoString() string {
	return s.String()
}

func (s *CostNodeGroupConfig) GetInstanceTypes() []*CostInstanceType {
	return s.InstanceTypes
}

func (s *CostNodeGroupConfig) GetMaximalNodeCount() *int32 {
	return s.MaximalNodeCount
}

func (s *CostNodeGroupConfig) GetMinimalNodeCount() *int32 {
	return s.MinimalNodeCount
}

func (s *CostNodeGroupConfig) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *CostNodeGroupConfig) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *CostNodeGroupConfig) GetNodeGroupType() *string {
	return s.NodeGroupType
}

func (s *CostNodeGroupConfig) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CostNodeGroupConfig) SetInstanceTypes(v []*CostInstanceType) *CostNodeGroupConfig {
	s.InstanceTypes = v
	return s
}

func (s *CostNodeGroupConfig) SetMaximalNodeCount(v int32) *CostNodeGroupConfig {
	s.MaximalNodeCount = &v
	return s
}

func (s *CostNodeGroupConfig) SetMinimalNodeCount(v int32) *CostNodeGroupConfig {
	s.MinimalNodeCount = &v
	return s
}

func (s *CostNodeGroupConfig) SetNodeCount(v int32) *CostNodeGroupConfig {
	s.NodeCount = &v
	return s
}

func (s *CostNodeGroupConfig) SetNodeGroupName(v string) *CostNodeGroupConfig {
	s.NodeGroupName = &v
	return s
}

func (s *CostNodeGroupConfig) SetNodeGroupType(v string) *CostNodeGroupConfig {
	s.NodeGroupType = &v
	return s
}

func (s *CostNodeGroupConfig) SetPaymentType(v string) *CostNodeGroupConfig {
	s.PaymentType = &v
	return s
}

func (s *CostNodeGroupConfig) Validate() error {
	return dara.Validate(s)
}
