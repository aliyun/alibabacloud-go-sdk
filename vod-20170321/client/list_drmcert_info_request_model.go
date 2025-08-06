// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDRMCertInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ListDRMCertInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListDRMCertInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListDRMCertInfoRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *ListDRMCertInfoRequest
	GetResourceRealOwnerId() *int64
}

type ListDRMCertInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
}

func (s ListDRMCertInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDRMCertInfoRequest) GoString() string {
	return s.String()
}

func (s *ListDRMCertInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListDRMCertInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListDRMCertInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListDRMCertInfoRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *ListDRMCertInfoRequest) SetOwnerId(v int64) *ListDRMCertInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDRMCertInfoRequest) SetResourceOwnerAccount(v string) *ListDRMCertInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListDRMCertInfoRequest) SetResourceOwnerId(v int64) *ListDRMCertInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListDRMCertInfoRequest) SetResourceRealOwnerId(v int64) *ListDRMCertInfoRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *ListDRMCertInfoRequest) Validate() error {
	return dara.Validate(s)
}
