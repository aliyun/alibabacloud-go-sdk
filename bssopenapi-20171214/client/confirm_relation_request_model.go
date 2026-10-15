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
	// The Alibaba Cloud UID of the linked account.
	//
	// example:
	//
	// 1512996702208737
	ChildUserId *int64 `json:"ChildUserId,omitempty" xml:"ChildUserId,omitempty"`
	// The confirmation action. Valid values:
	//
	// - child_agree: The linked account agrees.
	//
	// - child_disagree: The linked account rejects.
	//
	// - 主账号取消确认: The main account cancels the confirmation.
	//
	// This parameter is required.
	//
	// example:
	//
	// child_agree
	ConfirmCode *string `json:"ConfirmCode,omitempty" xml:"ConfirmCode,omitempty"`
	// The Alibaba Cloud UID of the main account.
	//
	// example:
	//
	// 1738376485192612
	ParentUserId *int64 `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	// The permissions granted to the linked account. Valid values:
	//
	// - SYNCHRONIZE_FINANCE_IDENTITY: Allows synchronizing the credit-based payment identity to the linked account.
	//
	// - SYNCHRONIZE_FINANCE_DISCOUNT_POLICY_TO_TARGET: Allows synchronizing discounts to the linked account.
	//
	// - FORBID_WITHDRAW_CASH: Prohibits the linked account from withdrawing cash from the balance.
	//
	// - FORBID_MANAGE_INVOICE: Prohibits the linked account from managing invoices.
	//
	// - CHECK_FINANCE_INFO: Allows requesting to view the account linking information.
	//
	// - MANAGE_TARGET_INVOICE: Allows managing invoices of the target account.
	//
	// - CHECK_TARGET_CONSUMPTION: Allows viewing the consumption information of the linked account.
	//
	// This parameter is required.
	//
	// example:
	//
	// SYNCHRONIZE_FINANCE_IDENTITY
	PermissionCodes []*string `json:"PermissionCodes,omitempty" xml:"PermissionCodes,omitempty" type:"Repeated"`
	// The relationship ID. You can obtain the value from the relationId field in the response of the QueryRelationList operation.
	//
	// example:
	//
	// 51463
	RelationId *int64 `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	// The relationship type. Valid values: enterprise_group.
	//
	// example:
	//
	// enterprise_group
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The unique identifier of the request.
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
