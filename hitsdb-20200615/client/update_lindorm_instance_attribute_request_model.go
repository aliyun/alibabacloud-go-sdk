// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLindormInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeletionProtection(v bool) *UpdateLindormInstanceAttributeRequest
	GetDeletionProtection() *bool
	SetDispositionType(v string) *UpdateLindormInstanceAttributeRequest
	GetDispositionType() *string
	SetInstanceAlias(v string) *UpdateLindormInstanceAttributeRequest
	GetInstanceAlias() *string
	SetInstanceId(v string) *UpdateLindormInstanceAttributeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *UpdateLindormInstanceAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateLindormInstanceAttributeRequest
	GetOwnerId() *int64
	SetPlannedCompletionDate(v string) *UpdateLindormInstanceAttributeRequest
	GetPlannedCompletionDate() *string
	SetResourceOwnerAccount(v string) *UpdateLindormInstanceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateLindormInstanceAttributeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *UpdateLindormInstanceAttributeRequest
	GetSecurityToken() *string
}

type UpdateLindormInstanceAttributeRequest struct {
	// Specifies whether to enable deletion protection. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The risk disposition type for a single zone.
	//
	// example:
	//
	// PLAN_MIGRATION
	DispositionType *string `json:"DispositionType,omitempty" xml:"DispositionType,omitempty"`
	// The instance name.
	//
	// example:
	//
	// lindorm-test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1z3506imz2f****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The planned completion date for risk disposition, in the yyyy-MM-dd format.
	//
	// example:
	//
	// 2026-10-01
	PlannedCompletionDate *string `json:"PlannedCompletionDate,omitempty" xml:"PlannedCompletionDate,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken         *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UpdateLindormInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLindormInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateLindormInstanceAttributeRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *UpdateLindormInstanceAttributeRequest) GetDispositionType() *string {
	return s.DispositionType
}

func (s *UpdateLindormInstanceAttributeRequest) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *UpdateLindormInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateLindormInstanceAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateLindormInstanceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLindormInstanceAttributeRequest) GetPlannedCompletionDate() *string {
	return s.PlannedCompletionDate
}

func (s *UpdateLindormInstanceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateLindormInstanceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateLindormInstanceAttributeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdateLindormInstanceAttributeRequest) SetDeletionProtection(v bool) *UpdateLindormInstanceAttributeRequest {
	s.DeletionProtection = &v
	return s
}

func (s *UpdateLindormInstanceAttributeRequest) SetDispositionType(v string) *UpdateLindormInstanceAttributeRequest {
	s.DispositionType = &v
	return s
}

func (s *UpdateLindormInstanceAttributeRequest) SetInstanceAlias(v string) *UpdateLindormInstanceAttributeRequest {
	s.InstanceAlias = &v
	return s
}

func (s *UpdateLindormInstanceAttributeRequest) SetInstanceId(v string) *UpdateLindormInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLindormInstanceAttributeRequest) SetOwnerAccount(v string) *UpdateLindormInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateLindormInstanceAttributeRequest) SetOwnerId(v int64) *UpdateLindormInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLindormInstanceAttributeRequest) SetPlannedCompletionDate(v string) *UpdateLindormInstanceAttributeRequest {
	s.PlannedCompletionDate = &v
	return s
}

func (s *UpdateLindormInstanceAttributeRequest) SetResourceOwnerAccount(v string) *UpdateLindormInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateLindormInstanceAttributeRequest) SetResourceOwnerId(v int64) *UpdateLindormInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateLindormInstanceAttributeRequest) SetSecurityToken(v string) *UpdateLindormInstanceAttributeRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateLindormInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
