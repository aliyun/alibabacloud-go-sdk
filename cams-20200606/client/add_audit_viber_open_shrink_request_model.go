// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuditViberOpenShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditRecordShrink(v string) *AddAuditViberOpenShrinkRequest
	GetAuditRecordShrink() *string
	SetAuditResult(v string) *AddAuditViberOpenShrinkRequest
	GetAuditResult() *string
	SetCustSpaceId(v string) *AddAuditViberOpenShrinkRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *AddAuditViberOpenShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddAuditViberOpenShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddAuditViberOpenShrinkRequest
	GetResourceOwnerId() *int64
}

type AddAuditViberOpenShrinkRequest struct {
	AuditRecordShrink *string `json:"AuditRecord,omitempty" xml:"AuditRecord,omitempty"`
	// example:
	//
	// 示例值
	AuditResult *string `json:"AuditResult,omitempty" xml:"AuditResult,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddAuditViberOpenShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAuditViberOpenShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddAuditViberOpenShrinkRequest) GetAuditRecordShrink() *string {
	return s.AuditRecordShrink
}

func (s *AddAuditViberOpenShrinkRequest) GetAuditResult() *string {
	return s.AuditResult
}

func (s *AddAuditViberOpenShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *AddAuditViberOpenShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddAuditViberOpenShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddAuditViberOpenShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddAuditViberOpenShrinkRequest) SetAuditRecordShrink(v string) *AddAuditViberOpenShrinkRequest {
	s.AuditRecordShrink = &v
	return s
}

func (s *AddAuditViberOpenShrinkRequest) SetAuditResult(v string) *AddAuditViberOpenShrinkRequest {
	s.AuditResult = &v
	return s
}

func (s *AddAuditViberOpenShrinkRequest) SetCustSpaceId(v string) *AddAuditViberOpenShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *AddAuditViberOpenShrinkRequest) SetOwnerId(v int64) *AddAuditViberOpenShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AddAuditViberOpenShrinkRequest) SetResourceOwnerAccount(v string) *AddAuditViberOpenShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddAuditViberOpenShrinkRequest) SetResourceOwnerId(v int64) *AddAuditViberOpenShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddAuditViberOpenShrinkRequest) Validate() error {
	return dara.Validate(s)
}
