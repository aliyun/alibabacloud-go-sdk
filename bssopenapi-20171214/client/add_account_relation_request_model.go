// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAccountRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChildNick(v string) *AddAccountRelationRequest
	GetChildNick() *string
	SetChildUserId(v int64) *AddAccountRelationRequest
	GetChildUserId() *int64
	SetParentUserId(v int64) *AddAccountRelationRequest
	GetParentUserId() *int64
	SetPermissionCodes(v []*string) *AddAccountRelationRequest
	GetPermissionCodes() []*string
	SetRelationType(v string) *AddAccountRelationRequest
	GetRelationType() *string
	SetRequestId(v string) *AddAccountRelationRequest
	GetRequestId() *string
	SetRoleCodes(v []*string) *AddAccountRelationRequest
	GetRoleCodes() []*string
}

type AddAccountRelationRequest struct {
	// The display name of the member. This helps clarify the scenario in which the account is used.
	//
	// example:
	//
	// xxx project
	ChildNick *string `json:"ChildNick,omitempty" xml:"ChildNick,omitempty"`
	// The ID of the Alibaba Cloud account that is used as the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1512996702208737
	ChildUserId *int64 `json:"ChildUserId,omitempty" xml:"ChildUserId,omitempty"`
	// The ID of the Alibaba Cloud account that is used as the management account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1738376485192612
	ParentUserId *int64 `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	// The permissions that can be granted to the member. Valid values:
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
	// The params[PermissionCodes, RoleCodes] can not be null at the same time.
	//
	// example:
	//
	// CHECK_TARGET_CONSUMPTION
	PermissionCodes []*string `json:"PermissionCodes,omitempty" xml:"PermissionCodes,omitempty" type:"Repeated"`
	// The type of the financial relationship. Set the value to enterprise_group.
	//
	// This parameter is required.
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
	// 32324242444
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The roles that can be assigned to the member. Set the value to trusteeship.
	//
	// example:
	//
	// trusteeship
	RoleCodes []*string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty" type:"Repeated"`
}

func (s AddAccountRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAccountRelationRequest) GoString() string {
	return s.String()
}

func (s *AddAccountRelationRequest) GetChildNick() *string {
	return s.ChildNick
}

func (s *AddAccountRelationRequest) GetChildUserId() *int64 {
	return s.ChildUserId
}

func (s *AddAccountRelationRequest) GetParentUserId() *int64 {
	return s.ParentUserId
}

func (s *AddAccountRelationRequest) GetPermissionCodes() []*string {
	return s.PermissionCodes
}

func (s *AddAccountRelationRequest) GetRelationType() *string {
	return s.RelationType
}

func (s *AddAccountRelationRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAccountRelationRequest) GetRoleCodes() []*string {
	return s.RoleCodes
}

func (s *AddAccountRelationRequest) SetChildNick(v string) *AddAccountRelationRequest {
	s.ChildNick = &v
	return s
}

func (s *AddAccountRelationRequest) SetChildUserId(v int64) *AddAccountRelationRequest {
	s.ChildUserId = &v
	return s
}

func (s *AddAccountRelationRequest) SetParentUserId(v int64) *AddAccountRelationRequest {
	s.ParentUserId = &v
	return s
}

func (s *AddAccountRelationRequest) SetPermissionCodes(v []*string) *AddAccountRelationRequest {
	s.PermissionCodes = v
	return s
}

func (s *AddAccountRelationRequest) SetRelationType(v string) *AddAccountRelationRequest {
	s.RelationType = &v
	return s
}

func (s *AddAccountRelationRequest) SetRequestId(v string) *AddAccountRelationRequest {
	s.RequestId = &v
	return s
}

func (s *AddAccountRelationRequest) SetRoleCodes(v []*string) *AddAccountRelationRequest {
	s.RoleCodes = v
	return s
}

func (s *AddAccountRelationRequest) Validate() error {
	return dara.Validate(s)
}
