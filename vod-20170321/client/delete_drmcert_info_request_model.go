// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDRMCertInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *DeleteDRMCertInfoRequest
	GetCertId() *string
	SetOwnerId(v int64) *DeleteDRMCertInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteDRMCertInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDRMCertInfoRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *DeleteDRMCertInfoRequest
	GetResourceRealOwnerId() *int64
}

type DeleteDRMCertInfoRequest struct {
	// This parameter is required.
	CertId               *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
}

func (s DeleteDRMCertInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDRMCertInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteDRMCertInfoRequest) GetCertId() *string {
	return s.CertId
}

func (s *DeleteDRMCertInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDRMCertInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDRMCertInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDRMCertInfoRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *DeleteDRMCertInfoRequest) SetCertId(v string) *DeleteDRMCertInfoRequest {
	s.CertId = &v
	return s
}

func (s *DeleteDRMCertInfoRequest) SetOwnerId(v int64) *DeleteDRMCertInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDRMCertInfoRequest) SetResourceOwnerAccount(v string) *DeleteDRMCertInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDRMCertInfoRequest) SetResourceOwnerId(v int64) *DeleteDRMCertInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDRMCertInfoRequest) SetResourceRealOwnerId(v int64) *DeleteDRMCertInfoRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *DeleteDRMCertInfoRequest) Validate() error {
	return dara.Validate(s)
}
