// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAbacPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbacPolicyId(v int64) *GetAbacPolicyRequest
	GetAbacPolicyId() *int64
	SetAbacPolicyName(v string) *GetAbacPolicyRequest
	GetAbacPolicyName() *string
	SetTid(v int64) *GetAbacPolicyRequest
	GetTid() *int64
}

type GetAbacPolicyRequest struct {
	// The ID of the policy.
	//
	// example:
	//
	// 12****
	AbacPolicyId *int64 `json:"AbacPolicyId,omitempty" xml:"AbacPolicyId,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// policy_test
	AbacPolicyName *string `json:"AbacPolicyName,omitempty" xml:"AbacPolicyName,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetAbacPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAbacPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetAbacPolicyRequest) GetAbacPolicyId() *int64 {
	return s.AbacPolicyId
}

func (s *GetAbacPolicyRequest) GetAbacPolicyName() *string {
	return s.AbacPolicyName
}

func (s *GetAbacPolicyRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetAbacPolicyRequest) SetAbacPolicyId(v int64) *GetAbacPolicyRequest {
	s.AbacPolicyId = &v
	return s
}

func (s *GetAbacPolicyRequest) SetAbacPolicyName(v string) *GetAbacPolicyRequest {
	s.AbacPolicyName = &v
	return s
}

func (s *GetAbacPolicyRequest) SetTid(v int64) *GetAbacPolicyRequest {
	s.Tid = &v
	return s
}

func (s *GetAbacPolicyRequest) Validate() error {
	return dara.Validate(s)
}
