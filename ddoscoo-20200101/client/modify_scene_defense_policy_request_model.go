// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySceneDefensePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ModifySceneDefensePolicyRequest
	GetEndTime() *int64
	SetName(v string) *ModifySceneDefensePolicyRequest
	GetName() *string
	SetPolicyId(v string) *ModifySceneDefensePolicyRequest
	GetPolicyId() *string
	SetStartTime(v int64) *ModifySceneDefensePolicyRequest
	GetStartTime() *int64
	SetTemplate(v string) *ModifySceneDefensePolicyRequest
	GetTemplate() *string
}

type ModifySceneDefensePolicyRequest struct {
	// The end time of the policy. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1586016000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpolicy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the policy that you want to modify.
	//
	// > You can call the [DescribeSceneDefensePolicies](https://help.aliyun.com/document_detail/159382.html) operation to query the IDs of all policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// 321a-fd31-df51-****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The start time of the policy. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1585670400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The template of the policy. Valid values:
	//
	// 	- **promotion**: important activity.
	//
	// 	- **bypass**: all traffic forwarded.
	//
	// This parameter is required.
	//
	// example:
	//
	// promotion
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s ModifySceneDefensePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySceneDefensePolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifySceneDefensePolicyRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModifySceneDefensePolicyRequest) GetName() *string {
	return s.Name
}

func (s *ModifySceneDefensePolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifySceneDefensePolicyRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModifySceneDefensePolicyRequest) GetTemplate() *string {
	return s.Template
}

func (s *ModifySceneDefensePolicyRequest) SetEndTime(v int64) *ModifySceneDefensePolicyRequest {
	s.EndTime = &v
	return s
}

func (s *ModifySceneDefensePolicyRequest) SetName(v string) *ModifySceneDefensePolicyRequest {
	s.Name = &v
	return s
}

func (s *ModifySceneDefensePolicyRequest) SetPolicyId(v string) *ModifySceneDefensePolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *ModifySceneDefensePolicyRequest) SetStartTime(v int64) *ModifySceneDefensePolicyRequest {
	s.StartTime = &v
	return s
}

func (s *ModifySceneDefensePolicyRequest) SetTemplate(v string) *ModifySceneDefensePolicyRequest {
	s.Template = &v
	return s
}

func (s *ModifySceneDefensePolicyRequest) Validate() error {
	return dara.Validate(s)
}
