// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuditResultDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *GetAuditResultDetailRequest
	GetMediaId() *string
	SetOwnerAccount(v string) *GetAuditResultDetailRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetAuditResultDetailRequest
	GetOwnerId() *string
	SetPageNo(v int32) *GetAuditResultDetailRequest
	GetPageNo() *int32
	SetResourceOwnerAccount(v string) *GetAuditResultDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetAuditResultDetailRequest
	GetResourceOwnerId() *string
}

type GetAuditResultDetailRequest struct {
	// This parameter is required.
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetAuditResultDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAuditResultDetailRequest) GoString() string {
	return s.String()
}

func (s *GetAuditResultDetailRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetAuditResultDetailRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetAuditResultDetailRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetAuditResultDetailRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetAuditResultDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetAuditResultDetailRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetAuditResultDetailRequest) SetMediaId(v string) *GetAuditResultDetailRequest {
	s.MediaId = &v
	return s
}

func (s *GetAuditResultDetailRequest) SetOwnerAccount(v string) *GetAuditResultDetailRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetAuditResultDetailRequest) SetOwnerId(v string) *GetAuditResultDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAuditResultDetailRequest) SetPageNo(v int32) *GetAuditResultDetailRequest {
	s.PageNo = &v
	return s
}

func (s *GetAuditResultDetailRequest) SetResourceOwnerAccount(v string) *GetAuditResultDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAuditResultDetailRequest) SetResourceOwnerId(v string) *GetAuditResultDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAuditResultDetailRequest) Validate() error {
	return dara.Validate(s)
}
