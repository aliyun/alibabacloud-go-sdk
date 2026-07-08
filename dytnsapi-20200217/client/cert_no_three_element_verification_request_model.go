// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertNoThreeElementVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *CertNoThreeElementVerificationRequest
	GetAuthCode() *string
	SetCertName(v string) *CertNoThreeElementVerificationRequest
	GetCertName() *string
	SetCertNo(v string) *CertNoThreeElementVerificationRequest
	GetCertNo() *string
	SetCertPicture(v string) *CertNoThreeElementVerificationRequest
	GetCertPicture() *string
	SetMask(v string) *CertNoThreeElementVerificationRequest
	GetMask() *string
	SetOwnerId(v int64) *CertNoThreeElementVerificationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CertNoThreeElementVerificationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CertNoThreeElementVerificationRequest
	GetResourceOwnerId() *int64
}

type CertNoThreeElementVerificationRequest struct {
	// The authorization code. Sources:
	//
	// - In the Cell Phone Number Service console, go to the [Tag Square](https://dytns.console.aliyun.com/analysis/square) page, select the **ID Card Three Elements*	- tag, and submit a usage application. After the application is approved, you will obtain the authorization code.
	//
	// - On the [My Applications](https://dytns.console.aliyun.com/analysis/apply) page of the Cell Phone Number Service console, view the approved **ID Card Three Elements*	- authorization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// QASDW@#**
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The name to be verified.
	//
	// This parameter is required.
	//
	// example:
	//
	// 张三
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The ID card number to be verified.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3***************0
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// The BASE64 encoding of the portrait photo to be verified. **Remove the encoded URI information (such as `data:image/png;base64,`) before submission**. The photo size and the BASE64-encoded size must not exceed 50 KB.
	//
	// This parameter is required.
	//
	// example:
	//
	// iVBOFMKODOFNDFP123DFSMOO...
	CertPicture *string `json:"CertPicture,omitempty" xml:"CertPicture,omitempty"`
	// Specifies whether to encrypt. Currently only unencrypted is supported.
	//
	// example:
	//
	// 不加密
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CertNoThreeElementVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s CertNoThreeElementVerificationRequest) GoString() string {
	return s.String()
}

func (s *CertNoThreeElementVerificationRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *CertNoThreeElementVerificationRequest) GetCertName() *string {
	return s.CertName
}

func (s *CertNoThreeElementVerificationRequest) GetCertNo() *string {
	return s.CertNo
}

func (s *CertNoThreeElementVerificationRequest) GetCertPicture() *string {
	return s.CertPicture
}

func (s *CertNoThreeElementVerificationRequest) GetMask() *string {
	return s.Mask
}

func (s *CertNoThreeElementVerificationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CertNoThreeElementVerificationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CertNoThreeElementVerificationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CertNoThreeElementVerificationRequest) SetAuthCode(v string) *CertNoThreeElementVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *CertNoThreeElementVerificationRequest) SetCertName(v string) *CertNoThreeElementVerificationRequest {
	s.CertName = &v
	return s
}

func (s *CertNoThreeElementVerificationRequest) SetCertNo(v string) *CertNoThreeElementVerificationRequest {
	s.CertNo = &v
	return s
}

func (s *CertNoThreeElementVerificationRequest) SetCertPicture(v string) *CertNoThreeElementVerificationRequest {
	s.CertPicture = &v
	return s
}

func (s *CertNoThreeElementVerificationRequest) SetMask(v string) *CertNoThreeElementVerificationRequest {
	s.Mask = &v
	return s
}

func (s *CertNoThreeElementVerificationRequest) SetOwnerId(v int64) *CertNoThreeElementVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *CertNoThreeElementVerificationRequest) SetResourceOwnerAccount(v string) *CertNoThreeElementVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CertNoThreeElementVerificationRequest) SetResourceOwnerId(v int64) *CertNoThreeElementVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CertNoThreeElementVerificationRequest) Validate() error {
	return dara.Validate(s)
}
