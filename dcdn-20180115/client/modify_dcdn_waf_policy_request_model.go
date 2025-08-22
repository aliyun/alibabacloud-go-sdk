// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDcdnWafPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v int64) *ModifyDcdnWafPolicyRequest
	GetPolicyId() *int64
	SetPolicyName(v string) *ModifyDcdnWafPolicyRequest
	GetPolicyName() *string
	SetPolicyStatus(v string) *ModifyDcdnWafPolicyRequest
	GetPolicyStatus() *string
}

type ModifyDcdnWafPolicyRequest struct {
	// The ID of the protection policy that you want to modify. You can specify only one ID in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The new name of the protection policy.
	//
	// > You must specify PolicyName or PolicyStatus.
	//
	// example:
	//
	// policy_test
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The new status of the protection policy. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// > You must specify PolicyName or PolicyStatus.
	//
	// example:
	//
	// on
	PolicyStatus *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
}

func (s ModifyDcdnWafPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDcdnWafPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyDcdnWafPolicyRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *ModifyDcdnWafPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ModifyDcdnWafPolicyRequest) GetPolicyStatus() *string {
	return s.PolicyStatus
}

func (s *ModifyDcdnWafPolicyRequest) SetPolicyId(v int64) *ModifyDcdnWafPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *ModifyDcdnWafPolicyRequest) SetPolicyName(v string) *ModifyDcdnWafPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *ModifyDcdnWafPolicyRequest) SetPolicyStatus(v string) *ModifyDcdnWafPolicyRequest {
	s.PolicyStatus = &v
	return s
}

func (s *ModifyDcdnWafPolicyRequest) Validate() error {
	return dara.Validate(s)
}
