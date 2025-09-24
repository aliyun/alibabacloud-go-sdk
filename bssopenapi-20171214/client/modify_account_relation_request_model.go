// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChildNick(v string) *ModifyAccountRelationRequest
	GetChildNick() *string
	SetChildUserId(v int64) *ModifyAccountRelationRequest
	GetChildUserId() *int64
	SetParentUserId(v int64) *ModifyAccountRelationRequest
	GetParentUserId() *int64
	SetPermissionCodes(v []*string) *ModifyAccountRelationRequest
	GetPermissionCodes() []*string
	SetRelationId(v int64) *ModifyAccountRelationRequest
	GetRelationId() *int64
	SetRelationOperation(v string) *ModifyAccountRelationRequest
	GetRelationOperation() *string
	SetRelationType(v string) *ModifyAccountRelationRequest
	GetRelationType() *string
	SetRequestId(v string) *ModifyAccountRelationRequest
	GetRequestId() *string
	SetRoleCodes(v []*string) *ModifyAccountRelationRequest
	GetRoleCodes() []*string
}

type ModifyAccountRelationRequest struct {
	// The display name of the member. This helps clarify the scenario in which the account is used.
	//
	// example:
	//
	// Display name of the member
	ChildNick *string `json:"ChildNick,omitempty" xml:"ChildNick,omitempty"`
	// The ID of the Alibaba Cloud account that is used as the member.
	//
	// example:
	//
	// 1512996702208737
	ChildUserId *int64 `json:"ChildUserId,omitempty" xml:"ChildUserId,omitempty"`
	// The ID of the Alibaba Cloud account that is used as the management account.
	//
	// example:
	//
	// 1738376485192612
	ParentUserId *int64 `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	// The permissions that can be modified. Valid values:
	//
	// 	- SYNCHRONIZE_FINANCE_IDENTITY: allows the credit control identity to be shared with the member.
	//
	// 	- SYNCHRONIZE_FINANCE_DISCOUNT_POLICY_TO_TARGET: allows the discount policy to be shared with the member.
	//
	// 	- FORBID_WITHDRAW_CASH: does not allow the member to withdraw the balance.
	//
	// 	- FORBID_MANAGE_INVOICE: does not allow the member to manage invoices.
	//
	// 	- CHECK_FINANCE_INFO: requests to view information about the financial relationship.
	//
	// 	- MANAGE_TARGET_INVOICE: allows the member to manage invoices.
	//
	// 	- CHECK_TARGET_CONSUMPTION: allows the member to view the bills.
	//
	// example:
	//
	// SYNCHRONIZE_FINANCE_IDENTITY
	PermissionCodes []*string `json:"PermissionCodes,omitempty" xml:"PermissionCodes,omitempty" type:"Repeated"`
	// The ID of the financial relationship. Set this parameter to the value of the relationId response parameter returned by calling the QueryRelationList operation.
	//
	// example:
	//
	// 51463
	RelationId *int64 `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	// The operation to be performed. Valid values:
	//
	// 	- ADD
	//
	// 	- DELETE
	//
	// This parameter is required.
	//
	// example:
	//
	// ADD
	RelationOperation *string `json:"RelationOperation,omitempty" xml:"RelationOperation,omitempty"`
	// The type of the financial relationship. Set the value to enterprise_group.
	//
	// example:
	//
	// enterprise_group
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The unique ID of the request. The ID is used to mark a request and troubleshoot a problem.
	//
	// This parameter is required.
	//
	// example:
	//
	// request_id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The roles that can be assigned to the member. You cannot modify the roles.
	//
	// example:
	//
	// trusteeship
	RoleCodes []*string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty" type:"Repeated"`
}

func (s ModifyAccountRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountRelationRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountRelationRequest) GetChildNick() *string {
	return s.ChildNick
}

func (s *ModifyAccountRelationRequest) GetChildUserId() *int64 {
	return s.ChildUserId
}

func (s *ModifyAccountRelationRequest) GetParentUserId() *int64 {
	return s.ParentUserId
}

func (s *ModifyAccountRelationRequest) GetPermissionCodes() []*string {
	return s.PermissionCodes
}

func (s *ModifyAccountRelationRequest) GetRelationId() *int64 {
	return s.RelationId
}

func (s *ModifyAccountRelationRequest) GetRelationOperation() *string {
	return s.RelationOperation
}

func (s *ModifyAccountRelationRequest) GetRelationType() *string {
	return s.RelationType
}

func (s *ModifyAccountRelationRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountRelationRequest) GetRoleCodes() []*string {
	return s.RoleCodes
}

func (s *ModifyAccountRelationRequest) SetChildNick(v string) *ModifyAccountRelationRequest {
	s.ChildNick = &v
	return s
}

func (s *ModifyAccountRelationRequest) SetChildUserId(v int64) *ModifyAccountRelationRequest {
	s.ChildUserId = &v
	return s
}

func (s *ModifyAccountRelationRequest) SetParentUserId(v int64) *ModifyAccountRelationRequest {
	s.ParentUserId = &v
	return s
}

func (s *ModifyAccountRelationRequest) SetPermissionCodes(v []*string) *ModifyAccountRelationRequest {
	s.PermissionCodes = v
	return s
}

func (s *ModifyAccountRelationRequest) SetRelationId(v int64) *ModifyAccountRelationRequest {
	s.RelationId = &v
	return s
}

func (s *ModifyAccountRelationRequest) SetRelationOperation(v string) *ModifyAccountRelationRequest {
	s.RelationOperation = &v
	return s
}

func (s *ModifyAccountRelationRequest) SetRelationType(v string) *ModifyAccountRelationRequest {
	s.RelationType = &v
	return s
}

func (s *ModifyAccountRelationRequest) SetRequestId(v string) *ModifyAccountRelationRequest {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountRelationRequest) SetRoleCodes(v []*string) *ModifyAccountRelationRequest {
	s.RoleCodes = v
	return s
}

func (s *ModifyAccountRelationRequest) Validate() error {
	return dara.Validate(s)
}
