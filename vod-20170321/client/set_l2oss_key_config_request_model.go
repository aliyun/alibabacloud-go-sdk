// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetL2OssKeyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SetL2OssKeyConfigRequest
	GetDomainName() *string
	SetOwnerAccount(v string) *SetL2OssKeyConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SetL2OssKeyConfigRequest
	GetOwnerId() *string
	SetPrivateOssAuth(v string) *SetL2OssKeyConfigRequest
	GetPrivateOssAuth() *string
	SetResourceOwnerAccount(v string) *SetL2OssKeyConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SetL2OssKeyConfigRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *SetL2OssKeyConfigRequest
	GetResourceRealOwnerId() *int64
}

type SetL2OssKeyConfigRequest struct {
	// This parameter is required.
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PrivateOssAuth       *string `json:"PrivateOssAuth,omitempty" xml:"PrivateOssAuth,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
}

func (s SetL2OssKeyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetL2OssKeyConfigRequest) GoString() string {
	return s.String()
}

func (s *SetL2OssKeyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetL2OssKeyConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetL2OssKeyConfigRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SetL2OssKeyConfigRequest) GetPrivateOssAuth() *string {
	return s.PrivateOssAuth
}

func (s *SetL2OssKeyConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetL2OssKeyConfigRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SetL2OssKeyConfigRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *SetL2OssKeyConfigRequest) SetDomainName(v string) *SetL2OssKeyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetL2OssKeyConfigRequest) SetOwnerAccount(v string) *SetL2OssKeyConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetL2OssKeyConfigRequest) SetOwnerId(v string) *SetL2OssKeyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetL2OssKeyConfigRequest) SetPrivateOssAuth(v string) *SetL2OssKeyConfigRequest {
	s.PrivateOssAuth = &v
	return s
}

func (s *SetL2OssKeyConfigRequest) SetResourceOwnerAccount(v string) *SetL2OssKeyConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetL2OssKeyConfigRequest) SetResourceOwnerId(v string) *SetL2OssKeyConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetL2OssKeyConfigRequest) SetResourceRealOwnerId(v int64) *SetL2OssKeyConfigRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *SetL2OssKeyConfigRequest) Validate() error {
	return dara.Validate(s)
}
