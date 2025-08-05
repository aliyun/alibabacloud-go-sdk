// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyControlPolicyPriorityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclUuid(v string) *ModifyControlPolicyPriorityRequest
	GetAclUuid() *string
	SetOrder(v string) *ModifyControlPolicyPriorityRequest
	GetOrder() *string
}

type ModifyControlPolicyPriorityRequest struct {
	// The UUID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3770d603-3534-4878-b845-f00095ee5048
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The priority of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s ModifyControlPolicyPriorityRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyControlPolicyPriorityRequest) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyPriorityRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *ModifyControlPolicyPriorityRequest) GetOrder() *string {
	return s.Order
}

func (s *ModifyControlPolicyPriorityRequest) SetAclUuid(v string) *ModifyControlPolicyPriorityRequest {
	s.AclUuid = &v
	return s
}

func (s *ModifyControlPolicyPriorityRequest) SetOrder(v string) *ModifyControlPolicyPriorityRequest {
	s.Order = &v
	return s
}

func (s *ModifyControlPolicyPriorityRequest) Validate() error {
	return dara.Validate(s)
}
