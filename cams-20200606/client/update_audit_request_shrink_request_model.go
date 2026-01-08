// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuditRequestShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditRecordShrink(v string) *UpdateAuditRequestShrinkRequest
	GetAuditRecordShrink() *string
	SetAuditResult(v string) *UpdateAuditRequestShrinkRequest
	GetAuditResult() *string
	SetCustSpaceId(v string) *UpdateAuditRequestShrinkRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *UpdateAuditRequestShrinkRequest
	GetOwnerId() *int64
	SetRequestNo(v string) *UpdateAuditRequestShrinkRequest
	GetRequestNo() *string
	SetResourceOwnerAccount(v string) *UpdateAuditRequestShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateAuditRequestShrinkRequest
	GetResourceOwnerId() *int64
}

type UpdateAuditRequestShrinkRequest struct {
	// This parameter is required.
	AuditRecordShrink *string `json:"AuditRecord,omitempty" xml:"AuditRecord,omitempty"`
	// example:
	//
	// unAudit
	AuditResult *string `json:"AuditResult,omitempty" xml:"AuditResult,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cams-8pi**urn5s
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 114624518**5956096
	RequestNo            *string `json:"RequestNo,omitempty" xml:"RequestNo,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateAuditRequestShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuditRequestShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuditRequestShrinkRequest) GetAuditRecordShrink() *string {
	return s.AuditRecordShrink
}

func (s *UpdateAuditRequestShrinkRequest) GetAuditResult() *string {
	return s.AuditResult
}

func (s *UpdateAuditRequestShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *UpdateAuditRequestShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateAuditRequestShrinkRequest) GetRequestNo() *string {
	return s.RequestNo
}

func (s *UpdateAuditRequestShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateAuditRequestShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateAuditRequestShrinkRequest) SetAuditRecordShrink(v string) *UpdateAuditRequestShrinkRequest {
	s.AuditRecordShrink = &v
	return s
}

func (s *UpdateAuditRequestShrinkRequest) SetAuditResult(v string) *UpdateAuditRequestShrinkRequest {
	s.AuditResult = &v
	return s
}

func (s *UpdateAuditRequestShrinkRequest) SetCustSpaceId(v string) *UpdateAuditRequestShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdateAuditRequestShrinkRequest) SetOwnerId(v int64) *UpdateAuditRequestShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAuditRequestShrinkRequest) SetRequestNo(v string) *UpdateAuditRequestShrinkRequest {
	s.RequestNo = &v
	return s
}

func (s *UpdateAuditRequestShrinkRequest) SetResourceOwnerAccount(v string) *UpdateAuditRequestShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateAuditRequestShrinkRequest) SetResourceOwnerId(v int64) *UpdateAuditRequestShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateAuditRequestShrinkRequest) Validate() error {
	return dara.Validate(s)
}
