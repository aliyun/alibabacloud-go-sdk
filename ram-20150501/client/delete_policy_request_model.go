// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCascadingDelete(v bool) *DeletePolicyRequest
	GetCascadingDelete() *bool
	SetPolicyName(v string) *DeletePolicyRequest
	GetPolicyName() *string
}

type DeletePolicyRequest struct {
	// Specifies whether to delete all versions of the policy. Valid values:
	//
	// 	- true: deletes all versions of the policy.
	//
	// 	- false: does not delete all versions of the policy. If you set the parameter to false, the non-default versions of the policy are not deleted. Before you delete the policy, you must manually delete all non-default versions of the policy.
	//
	// example:
	//
	// true
	CascadingDelete *bool `json:"CascadingDelete,omitempty" xml:"CascadingDelete,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s DeletePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyRequest) GetCascadingDelete() *bool {
	return s.CascadingDelete
}

func (s *DeletePolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DeletePolicyRequest) SetCascadingDelete(v bool) *DeletePolicyRequest {
	s.CascadingDelete = &v
	return s
}

func (s *DeletePolicyRequest) SetPolicyName(v string) *DeletePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *DeletePolicyRequest) Validate() error {
	return dara.Validate(s)
}
