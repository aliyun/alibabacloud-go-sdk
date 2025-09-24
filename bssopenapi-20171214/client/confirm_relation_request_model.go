// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChildUserId(v int64) *ConfirmRelationRequest
	GetChildUserId() *int64
	SetConfirmCode(v string) *ConfirmRelationRequest
	GetConfirmCode() *string
	SetParentUserId(v int64) *ConfirmRelationRequest
	GetParentUserId() *int64
	SetPermissionCodes(v []*string) *ConfirmRelationRequest
	GetPermissionCodes() []*string
	SetRelationId(v int64) *ConfirmRelationRequest
	GetRelationId() *int64
	SetRelationType(v string) *ConfirmRelationRequest
	GetRelationType() *string
	SetRequestId(v string) *ConfirmRelationRequest
	GetRequestId() *string
}

type ConfirmRelationRequest struct {
	// The ID of the Alibaba Cloud account that is used as the member.
	//
	// example:
	//
	// 1512996702208737
	ChildUserId *int64 `json:"ChildUserId,omitempty" xml:"ChildUserId,omitempty"`
	// The operation to be performed to confirm the invitation. Valid values:
	//
	// 	- child_agree: The member accepts the invitation.
	//
	// 	- child_disagree: The member rejects the invitation.
	//
	// 	- Canceled by the master account: The management account cancels the confirmation.
	//
	// This parameter is required.
	//
	// example:
	//
	// child_agree
	ConfirmCode *string `json:"ConfirmCode,omitempty" xml:"ConfirmCode,omitempty"`
	// The ID of the Alibaba Cloud account that is used as the management account.
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
	// This parameter is required.
	//
	// example:
	//
	// SYNCHRONIZE_FINANCE_IDENTITY
	PermissionCodes []*string `json:"PermissionCodes,omitempty" xml:"PermissionCodes,omitempty" type:"Repeated"`
	// The ID of the financial relationship. Set this parameter to the value of the RelationId response parameter returned by calling the QueryRelationList operation.
	//
	// example:
	//
	// 51463
	RelationId *int64 `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	// The type of the financial relationship. Set the value to enterprise_group.
	//
	// example:
	//
	// Type of the financial relationship
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The unique ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// request_id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfirmRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmRelationRequest) GoString() string {
	return s.String()
}

func (s *ConfirmRelationRequest) GetChildUserId() *int64 {
	return s.ChildUserId
}

func (s *ConfirmRelationRequest) GetConfirmCode() *string {
	return s.ConfirmCode
}

func (s *ConfirmRelationRequest) GetParentUserId() *int64 {
	return s.ParentUserId
}

func (s *ConfirmRelationRequest) GetPermissionCodes() []*string {
	return s.PermissionCodes
}

func (s *ConfirmRelationRequest) GetRelationId() *int64 {
	return s.RelationId
}

func (s *ConfirmRelationRequest) GetRelationType() *string {
	return s.RelationType
}

func (s *ConfirmRelationRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmRelationRequest) SetChildUserId(v int64) *ConfirmRelationRequest {
	s.ChildUserId = &v
	return s
}

func (s *ConfirmRelationRequest) SetConfirmCode(v string) *ConfirmRelationRequest {
	s.ConfirmCode = &v
	return s
}

func (s *ConfirmRelationRequest) SetParentUserId(v int64) *ConfirmRelationRequest {
	s.ParentUserId = &v
	return s
}

func (s *ConfirmRelationRequest) SetPermissionCodes(v []*string) *ConfirmRelationRequest {
	s.PermissionCodes = v
	return s
}

func (s *ConfirmRelationRequest) SetRelationId(v int64) *ConfirmRelationRequest {
	s.RelationId = &v
	return s
}

func (s *ConfirmRelationRequest) SetRelationType(v string) *ConfirmRelationRequest {
	s.RelationType = &v
	return s
}

func (s *ConfirmRelationRequest) SetRequestId(v string) *ConfirmRelationRequest {
	s.RequestId = &v
	return s
}

func (s *ConfirmRelationRequest) Validate() error {
	return dara.Validate(s)
}
