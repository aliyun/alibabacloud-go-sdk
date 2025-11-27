// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustinsResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdjustDeadline(v string) *ModifyCustinsResourceRequest
	GetAdjustDeadline() *string
	SetDBInstanceId(v string) *ModifyCustinsResourceRequest
	GetDBInstanceId() *string
	SetIncreaseRatio(v string) *ModifyCustinsResourceRequest
	GetIncreaseRatio() *string
	SetResourceOwnerId(v int64) *ModifyCustinsResourceRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *ModifyCustinsResourceRequest
	GetResourceType() *string
	SetRestoreOriginalSpecification(v string) *ModifyCustinsResourceRequest
	GetRestoreOriginalSpecification() *string
	SetTargetValue(v int32) *ModifyCustinsResourceRequest
	GetTargetValue() *int32
}

type ModifyCustinsResourceRequest struct {
	// The deadline for the modification.
	//
	// example:
	//
	// 2022-12-31 23:59:06
	AdjustDeadline *string `json:"AdjustDeadline,omitempty" xml:"AdjustDeadline,omitempty"`
	// The instance ID. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/26232.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-j5ekvfeengm******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The increase rate in percentage.
	//
	// example:
	//
	// 10
	IncreaseRatio   *string `json:"IncreaseRatio,omitempty" xml:"IncreaseRatio,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// Memory
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The original value. This parameter must be specified when the **ResourceType*	- parameter is set to **instance**.
	//
	// example:
	//
	// 200
	RestoreOriginalSpecification *string `json:"RestoreOriginalSpecification,omitempty" xml:"RestoreOriginalSpecification,omitempty"`
	// The target value. This parameter is available only if you set the ScalingRuleType parameter to TargetTrackingScalingRule or PredictiveScalingRule. The value must be greater than 0 and can contain up to three decimal places.
	//
	// example:
	//
	// 3000
	TargetValue *int32 `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s ModifyCustinsResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustinsResourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyCustinsResourceRequest) GetAdjustDeadline() *string {
	return s.AdjustDeadline
}

func (s *ModifyCustinsResourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyCustinsResourceRequest) GetIncreaseRatio() *string {
	return s.IncreaseRatio
}

func (s *ModifyCustinsResourceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCustinsResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ModifyCustinsResourceRequest) GetRestoreOriginalSpecification() *string {
	return s.RestoreOriginalSpecification
}

func (s *ModifyCustinsResourceRequest) GetTargetValue() *int32 {
	return s.TargetValue
}

func (s *ModifyCustinsResourceRequest) SetAdjustDeadline(v string) *ModifyCustinsResourceRequest {
	s.AdjustDeadline = &v
	return s
}

func (s *ModifyCustinsResourceRequest) SetDBInstanceId(v string) *ModifyCustinsResourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyCustinsResourceRequest) SetIncreaseRatio(v string) *ModifyCustinsResourceRequest {
	s.IncreaseRatio = &v
	return s
}

func (s *ModifyCustinsResourceRequest) SetResourceOwnerId(v int64) *ModifyCustinsResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCustinsResourceRequest) SetResourceType(v string) *ModifyCustinsResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyCustinsResourceRequest) SetRestoreOriginalSpecification(v string) *ModifyCustinsResourceRequest {
	s.RestoreOriginalSpecification = &v
	return s
}

func (s *ModifyCustinsResourceRequest) SetTargetValue(v int32) *ModifyCustinsResourceRequest {
	s.TargetValue = &v
	return s
}

func (s *ModifyCustinsResourceRequest) Validate() error {
	return dara.Validate(s)
}
