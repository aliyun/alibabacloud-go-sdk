// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegistDRMCertInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsk(v string) *RegistDRMCertInfoRequest
	GetAsk() *string
	SetCertName(v string) *RegistDRMCertInfoRequest
	GetCertName() *string
	SetDescription(v string) *RegistDRMCertInfoRequest
	GetDescription() *string
	SetOwnerId(v int64) *RegistDRMCertInfoRequest
	GetOwnerId() *int64
	SetPassPhrase(v string) *RegistDRMCertInfoRequest
	GetPassPhrase() *string
	SetPrivateKey(v string) *RegistDRMCertInfoRequest
	GetPrivateKey() *string
	SetResourceOwnerAccount(v string) *RegistDRMCertInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RegistDRMCertInfoRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *RegistDRMCertInfoRequest
	GetResourceRealOwnerId() *int64
	SetServCert(v string) *RegistDRMCertInfoRequest
	GetServCert() *string
}

type RegistDRMCertInfoRequest struct {
	// This parameter is required.
	Ask *string `json:"Ask,omitempty" xml:"Ask,omitempty"`
	// This parameter is required.
	CertName    *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PassPhrase *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	// This parameter is required.
	PrivateKey           *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	ServCert *string `json:"ServCert,omitempty" xml:"ServCert,omitempty"`
}

func (s RegistDRMCertInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s RegistDRMCertInfoRequest) GoString() string {
	return s.String()
}

func (s *RegistDRMCertInfoRequest) GetAsk() *string {
	return s.Ask
}

func (s *RegistDRMCertInfoRequest) GetCertName() *string {
	return s.CertName
}

func (s *RegistDRMCertInfoRequest) GetDescription() *string {
	return s.Description
}

func (s *RegistDRMCertInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RegistDRMCertInfoRequest) GetPassPhrase() *string {
	return s.PassPhrase
}

func (s *RegistDRMCertInfoRequest) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *RegistDRMCertInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RegistDRMCertInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RegistDRMCertInfoRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *RegistDRMCertInfoRequest) GetServCert() *string {
	return s.ServCert
}

func (s *RegistDRMCertInfoRequest) SetAsk(v string) *RegistDRMCertInfoRequest {
	s.Ask = &v
	return s
}

func (s *RegistDRMCertInfoRequest) SetCertName(v string) *RegistDRMCertInfoRequest {
	s.CertName = &v
	return s
}

func (s *RegistDRMCertInfoRequest) SetDescription(v string) *RegistDRMCertInfoRequest {
	s.Description = &v
	return s
}

func (s *RegistDRMCertInfoRequest) SetOwnerId(v int64) *RegistDRMCertInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *RegistDRMCertInfoRequest) SetPassPhrase(v string) *RegistDRMCertInfoRequest {
	s.PassPhrase = &v
	return s
}

func (s *RegistDRMCertInfoRequest) SetPrivateKey(v string) *RegistDRMCertInfoRequest {
	s.PrivateKey = &v
	return s
}

func (s *RegistDRMCertInfoRequest) SetResourceOwnerAccount(v string) *RegistDRMCertInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RegistDRMCertInfoRequest) SetResourceOwnerId(v int64) *RegistDRMCertInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RegistDRMCertInfoRequest) SetResourceRealOwnerId(v int64) *RegistDRMCertInfoRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *RegistDRMCertInfoRequest) SetServCert(v string) *RegistDRMCertInfoRequest {
	s.ServCert = &v
	return s
}

func (s *RegistDRMCertInfoRequest) Validate() error {
	return dara.Validate(s)
}
