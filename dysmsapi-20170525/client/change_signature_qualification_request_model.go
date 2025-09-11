// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeSignatureQualificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationLetterId(v int64) *ChangeSignatureQualificationRequest
	GetAuthorizationLetterId() *int64
	SetOwnerId(v int64) *ChangeSignatureQualificationRequest
	GetOwnerId() *int64
	SetQualificationId(v int64) *ChangeSignatureQualificationRequest
	GetQualificationId() *int64
	SetResourceOwnerAccount(v string) *ChangeSignatureQualificationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ChangeSignatureQualificationRequest
	GetResourceOwnerId() *int64
	SetSignatureName(v string) *ChangeSignatureQualificationRequest
	GetSignatureName() *string
}

type ChangeSignatureQualificationRequest struct {
	// 授权委托书id
	//
	// example:
	//
	// 1000********1234
	AuthorizationLetterId *int64 `json:"AuthorizationLetterId,omitempty" xml:"AuthorizationLetterId,omitempty"`
	OwnerId               *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 资质id
	//
	// This parameter is required.
	//
	// example:
	//
	// 1*****2
	QualificationId      *int64  `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 签名
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	SignatureName *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
}

func (s ChangeSignatureQualificationRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeSignatureQualificationRequest) GoString() string {
	return s.String()
}

func (s *ChangeSignatureQualificationRequest) GetAuthorizationLetterId() *int64 {
	return s.AuthorizationLetterId
}

func (s *ChangeSignatureQualificationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChangeSignatureQualificationRequest) GetQualificationId() *int64 {
	return s.QualificationId
}

func (s *ChangeSignatureQualificationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ChangeSignatureQualificationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ChangeSignatureQualificationRequest) GetSignatureName() *string {
	return s.SignatureName
}

func (s *ChangeSignatureQualificationRequest) SetAuthorizationLetterId(v int64) *ChangeSignatureQualificationRequest {
	s.AuthorizationLetterId = &v
	return s
}

func (s *ChangeSignatureQualificationRequest) SetOwnerId(v int64) *ChangeSignatureQualificationRequest {
	s.OwnerId = &v
	return s
}

func (s *ChangeSignatureQualificationRequest) SetQualificationId(v int64) *ChangeSignatureQualificationRequest {
	s.QualificationId = &v
	return s
}

func (s *ChangeSignatureQualificationRequest) SetResourceOwnerAccount(v string) *ChangeSignatureQualificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ChangeSignatureQualificationRequest) SetResourceOwnerId(v int64) *ChangeSignatureQualificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ChangeSignatureQualificationRequest) SetSignatureName(v string) *ChangeSignatureQualificationRequest {
	s.SignatureName = &v
	return s
}

func (s *ChangeSignatureQualificationRequest) Validate() error {
	return dara.Validate(s)
}
