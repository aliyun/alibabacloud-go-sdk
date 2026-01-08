// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAddressRecoverSuspendShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditRecordShrink(v string) *AddAddressRecoverSuspendShrinkRequest
	GetAuditRecordShrink() *string
	SetCustSpaceId(v string) *AddAddressRecoverSuspendShrinkRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *AddAddressRecoverSuspendShrinkRequest
	GetOwnerId() *int64
	SetRequestType(v string) *AddAddressRecoverSuspendShrinkRequest
	GetRequestType() *string
	SetResourceOwnerAccount(v string) *AddAddressRecoverSuspendShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddAddressRecoverSuspendShrinkRequest
	GetResourceOwnerId() *int64
}

type AddAddressRecoverSuspendShrinkRequest struct {
	AuditRecordShrink *string `json:"AuditRecord,omitempty" xml:"AuditRecord,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 示例值
	RequestType          *string `json:"RequestType,omitempty" xml:"RequestType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddAddressRecoverSuspendShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAddressRecoverSuspendShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddAddressRecoverSuspendShrinkRequest) GetAuditRecordShrink() *string {
	return s.AuditRecordShrink
}

func (s *AddAddressRecoverSuspendShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *AddAddressRecoverSuspendShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddAddressRecoverSuspendShrinkRequest) GetRequestType() *string {
	return s.RequestType
}

func (s *AddAddressRecoverSuspendShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddAddressRecoverSuspendShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddAddressRecoverSuspendShrinkRequest) SetAuditRecordShrink(v string) *AddAddressRecoverSuspendShrinkRequest {
	s.AuditRecordShrink = &v
	return s
}

func (s *AddAddressRecoverSuspendShrinkRequest) SetCustSpaceId(v string) *AddAddressRecoverSuspendShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *AddAddressRecoverSuspendShrinkRequest) SetOwnerId(v int64) *AddAddressRecoverSuspendShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AddAddressRecoverSuspendShrinkRequest) SetRequestType(v string) *AddAddressRecoverSuspendShrinkRequest {
	s.RequestType = &v
	return s
}

func (s *AddAddressRecoverSuspendShrinkRequest) SetResourceOwnerAccount(v string) *AddAddressRecoverSuspendShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddAddressRecoverSuspendShrinkRequest) SetResourceOwnerId(v int64) *AddAddressRecoverSuspendShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddAddressRecoverSuspendShrinkRequest) Validate() error {
	return dara.Validate(s)
}
