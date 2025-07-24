// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedScalingPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetManagedScalingPolicyResponseBody
	GetRequestId() *string
	SetScalingPolicy(v *GetManagedScalingPolicyResponseBodyScalingPolicy) *GetManagedScalingPolicyResponseBody
	GetScalingPolicy() *GetManagedScalingPolicyResponseBodyScalingPolicy
}

type GetManagedScalingPolicyResponseBody struct {
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingPolicy *GetManagedScalingPolicyResponseBodyScalingPolicy `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty" type:"Struct"`
}

func (s GetManagedScalingPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetManagedScalingPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetManagedScalingPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetManagedScalingPolicyResponseBody) GetScalingPolicy() *GetManagedScalingPolicyResponseBodyScalingPolicy {
	return s.ScalingPolicy
}

func (s *GetManagedScalingPolicyResponseBody) SetRequestId(v string) *GetManagedScalingPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetManagedScalingPolicyResponseBody) SetScalingPolicy(v *GetManagedScalingPolicyResponseBodyScalingPolicy) *GetManagedScalingPolicyResponseBody {
	s.ScalingPolicy = v
	return s
}

func (s *GetManagedScalingPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetManagedScalingPolicyResponseBodyScalingPolicy struct {
	// 集群ID。
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 最大最小值约束
	Constraints *GetManagedScalingPolicyResponseBodyScalingPolicyConstraints `json:"Constraints,omitempty" xml:"Constraints,omitempty" type:"Struct"`
	// 伸缩策略ID。
	//
	// example:
	//
	// asp-asduwe23znl***
	ScalingPolicyId *string `json:"ScalingPolicyId,omitempty" xml:"ScalingPolicyId,omitempty"`
}

func (s GetManagedScalingPolicyResponseBodyScalingPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetManagedScalingPolicyResponseBodyScalingPolicy) GoString() string {
	return s.String()
}

func (s *GetManagedScalingPolicyResponseBodyScalingPolicy) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetManagedScalingPolicyResponseBodyScalingPolicy) GetConstraints() *GetManagedScalingPolicyResponseBodyScalingPolicyConstraints {
	return s.Constraints
}

func (s *GetManagedScalingPolicyResponseBodyScalingPolicy) GetScalingPolicyId() *string {
	return s.ScalingPolicyId
}

func (s *GetManagedScalingPolicyResponseBodyScalingPolicy) SetClusterId(v string) *GetManagedScalingPolicyResponseBodyScalingPolicy {
	s.ClusterId = &v
	return s
}

func (s *GetManagedScalingPolicyResponseBodyScalingPolicy) SetConstraints(v *GetManagedScalingPolicyResponseBodyScalingPolicyConstraints) *GetManagedScalingPolicyResponseBodyScalingPolicy {
	s.Constraints = v
	return s
}

func (s *GetManagedScalingPolicyResponseBodyScalingPolicy) SetScalingPolicyId(v string) *GetManagedScalingPolicyResponseBodyScalingPolicy {
	s.ScalingPolicyId = &v
	return s
}

func (s *GetManagedScalingPolicyResponseBodyScalingPolicy) Validate() error {
	return dara.Validate(s)
}

type GetManagedScalingPolicyResponseBodyScalingPolicyConstraints struct {
	// 最大值
	//
	// example:
	//
	// 20
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// 最大按量数量
	//
	// example:
	//
	// 10
	MaxOnDemandCapacity *int32 `json:"MaxOnDemandCapacity,omitempty" xml:"MaxOnDemandCapacity,omitempty"`
	// 最小值
	//
	// example:
	//
	// 0
	MinCapacity *int32 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
}

func (s GetManagedScalingPolicyResponseBodyScalingPolicyConstraints) String() string {
	return dara.Prettify(s)
}

func (s GetManagedScalingPolicyResponseBodyScalingPolicyConstraints) GoString() string {
	return s.String()
}

func (s *GetManagedScalingPolicyResponseBodyScalingPolicyConstraints) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *GetManagedScalingPolicyResponseBodyScalingPolicyConstraints) GetMaxOnDemandCapacity() *int32 {
	return s.MaxOnDemandCapacity
}

func (s *GetManagedScalingPolicyResponseBodyScalingPolicyConstraints) GetMinCapacity() *int32 {
	return s.MinCapacity
}

func (s *GetManagedScalingPolicyResponseBodyScalingPolicyConstraints) SetMaxCapacity(v int32) *GetManagedScalingPolicyResponseBodyScalingPolicyConstraints {
	s.MaxCapacity = &v
	return s
}

func (s *GetManagedScalingPolicyResponseBodyScalingPolicyConstraints) SetMaxOnDemandCapacity(v int32) *GetManagedScalingPolicyResponseBodyScalingPolicyConstraints {
	s.MaxOnDemandCapacity = &v
	return s
}

func (s *GetManagedScalingPolicyResponseBodyScalingPolicyConstraints) SetMinCapacity(v int32) *GetManagedScalingPolicyResponseBodyScalingPolicyConstraints {
	s.MinCapacity = &v
	return s
}

func (s *GetManagedScalingPolicyResponseBodyScalingPolicyConstraints) Validate() error {
	return dara.Validate(s)
}
