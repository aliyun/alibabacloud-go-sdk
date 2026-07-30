// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainCloudAccountRoleAccessCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountId(v string) *ObtainCloudAccountRoleAccessCredentialResponseBody
	GetCloudAccountId() *string
	SetCloudAccountRoleAccessCredential(v *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential) *ObtainCloudAccountRoleAccessCredentialResponseBody
	GetCloudAccountRoleAccessCredential() *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential
	SetCloudAccountRoleExternalId(v string) *ObtainCloudAccountRoleAccessCredentialResponseBody
	GetCloudAccountRoleExternalId() *string
	SetCloudAccountRoleId(v string) *ObtainCloudAccountRoleAccessCredentialResponseBody
	GetCloudAccountRoleId() *string
	SetCloudAccountRoleName(v string) *ObtainCloudAccountRoleAccessCredentialResponseBody
	GetCloudAccountRoleName() *string
	SetCloudAccountVendorType(v string) *ObtainCloudAccountRoleAccessCredentialResponseBody
	GetCloudAccountVendorType() *string
}

type ObtainCloudAccountRoleAccessCredentialResponseBody struct {
	// The cloud account ID.
	//
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"cloudAccountId,omitempty" xml:"cloudAccountId,omitempty"`
	// The temporary access credentials for assuming the cloud account role.
	CloudAccountRoleAccessCredential *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential `json:"cloudAccountRoleAccessCredential,omitempty" xml:"cloudAccountRoleAccessCredential,omitempty" type:"Struct"`
	// The business identifier of the cloud account role.
	//
	// example:
	//
	// acs:ram::xxx:role/role-test
	CloudAccountRoleExternalId *string `json:"cloudAccountRoleExternalId,omitempty" xml:"cloudAccountRoleExternalId,omitempty"`
	// The cloud account role ID.
	//
	// example:
	//
	// carole_01kmek49aqxxxx
	CloudAccountRoleId *string `json:"cloudAccountRoleId,omitempty" xml:"cloudAccountRoleId,omitempty"`
	// The cloud account role name.
	//
	// example:
	//
	// role-test
	CloudAccountRoleName *string `json:"cloudAccountRoleName,omitempty" xml:"cloudAccountRoleName,omitempty"`
	// The cloud account type. Valid values:
	//
	// - alibaba_cloud: Alibaba Cloud.
	//
	// example:
	//
	// alibaba_cloud
	CloudAccountVendorType *string `json:"cloudAccountVendorType,omitempty" xml:"cloudAccountVendorType,omitempty"`
}

func (s ObtainCloudAccountRoleAccessCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ObtainCloudAccountRoleAccessCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBody) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBody) GetCloudAccountRoleAccessCredential() *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential {
	return s.CloudAccountRoleAccessCredential
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBody) GetCloudAccountRoleExternalId() *string {
	return s.CloudAccountRoleExternalId
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBody) GetCloudAccountRoleId() *string {
	return s.CloudAccountRoleId
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBody) GetCloudAccountRoleName() *string {
	return s.CloudAccountRoleName
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBody) GetCloudAccountVendorType() *string {
	return s.CloudAccountVendorType
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBody) SetCloudAccountId(v string) *ObtainCloudAccountRoleAccessCredentialResponseBody {
	s.CloudAccountId = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBody) SetCloudAccountRoleAccessCredential(v *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential) *ObtainCloudAccountRoleAccessCredentialResponseBody {
	s.CloudAccountRoleAccessCredential = v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBody) SetCloudAccountRoleExternalId(v string) *ObtainCloudAccountRoleAccessCredentialResponseBody {
	s.CloudAccountRoleExternalId = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBody) SetCloudAccountRoleId(v string) *ObtainCloudAccountRoleAccessCredentialResponseBody {
	s.CloudAccountRoleId = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBody) SetCloudAccountRoleName(v string) *ObtainCloudAccountRoleAccessCredentialResponseBody {
	s.CloudAccountRoleName = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBody) SetCloudAccountVendorType(v string) *ObtainCloudAccountRoleAccessCredentialResponseBody {
	s.CloudAccountVendorType = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBody) Validate() error {
	if s.CloudAccountRoleAccessCredential != nil {
		if err := s.CloudAccountRoleAccessCredential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential struct {
	// The expiration time of the temporary access credentials for the cloud account role, in UNIX timestamp format and in seconds.
	//
	// example:
	//
	// 1767196800
	AccessCredentialExpiresAt *int64 `json:"accessCredentialExpiresAt,omitempty" xml:"accessCredentialExpiresAt,omitempty"`
	// The temporary identity credentials (STS Token) for assuming an Alibaba Cloud RAM role.
	//
	// > This field is returned only when the cloud account type associated with the cloud account role is Alibaba Cloud (alibaba_cloud).
	AlibabaCloudStsToken *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken `json:"alibabaCloudStsToken,omitempty" xml:"alibabaCloudStsToken,omitempty" type:"Struct"`
	// The STS Token representing an AWS role.
	AwsStsToken          *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken          `json:"awsStsToken,omitempty" xml:"awsStsToken,omitempty" type:"Struct"`
	TencentCloudStsToken *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken `json:"tencentCloudStsToken,omitempty" xml:"tencentCloudStsToken,omitempty" type:"Struct"`
}

func (s ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential) String() string {
	return dara.Prettify(s)
}

func (s ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential) GoString() string {
	return s.String()
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential) GetAccessCredentialExpiresAt() *int64 {
	return s.AccessCredentialExpiresAt
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential) GetAlibabaCloudStsToken() *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken {
	return s.AlibabaCloudStsToken
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential) GetAwsStsToken() *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken {
	return s.AwsStsToken
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential) GetTencentCloudStsToken() *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken {
	return s.TencentCloudStsToken
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential) SetAccessCredentialExpiresAt(v int64) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential {
	s.AccessCredentialExpiresAt = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential) SetAlibabaCloudStsToken(v *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential {
	s.AlibabaCloudStsToken = v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential) SetAwsStsToken(v *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential {
	s.AwsStsToken = v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential) SetTencentCloudStsToken(v *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential {
	s.TencentCloudStsToken = v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential) Validate() error {
	if s.AlibabaCloudStsToken != nil {
		if err := s.AlibabaCloudStsToken.Validate(); err != nil {
			return err
		}
	}
	if s.AwsStsToken != nil {
		if err := s.AwsStsToken.Validate(); err != nil {
			return err
		}
	}
	if s.TencentCloudStsToken != nil {
		if err := s.TencentCloudStsToken.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken struct {
	// The access key ID.
	//
	// example:
	//
	// STS.NUgYrLnoC37mZZCNnAbez****
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// The access key secret.
	//
	// example:
	//
	// CVwjCkNzTMupZ8NbTCxCBRq3K16jtcWFTJAyBEv2****
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// The expiration time of the token (UTC).
	//
	// example:
	//
	// 2021-10-20T04:27:09Z
	Expiration *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// The security token.
	//
	// example:
	//
	// CAIShwJ1q6Ft5B2yfSjIr5bSEsj4g7BihPWGWHz****
	SecurityToken *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
}

func (s ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken) String() string {
	return dara.Prettify(s)
}

func (s ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken) GoString() string {
	return s.String()
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken) GetExpiration() *string {
	return s.Expiration
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken) SetAccessKeyId(v string) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken {
	s.AccessKeyId = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken) SetAccessKeySecret(v string) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken {
	s.AccessKeySecret = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken) SetExpiration(v string) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken {
	s.Expiration = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken) SetSecurityToken(v string) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken {
	s.SecurityToken = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken) Validate() error {
	return dara.Validate(s)
}

type ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken struct {
	// The access key ID.
	//
	// example:
	//
	// ASIAYBGN7XJKRFOM****
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// The expiration time of the STS Token (UTC).
	//
	// example:
	//
	// 2021-10-20T04:27:09Z
	Expiration *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// The secret access key.
	//
	// example:
	//
	// CVwjCkNzTMupZ8NbTCxCBRq3K16jtcWFTJAyBEv2****
	SecretAccessKey *string `json:"secretAccessKey,omitempty" xml:"secretAccessKey,omitempty"`
	// The session token of the temporary credentials.
	//
	// example:
	//
	// FwoDYXdzEJzfSjIr5bSEsj4g7BihPWGWHz****
	SessionToken *string `json:"sessionToken,omitempty" xml:"sessionToken,omitempty"`
}

func (s ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken) String() string {
	return dara.Prettify(s)
}

func (s ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken) GoString() string {
	return s.String()
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken) GetExpiration() *string {
	return s.Expiration
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken) GetSecretAccessKey() *string {
	return s.SecretAccessKey
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken) GetSessionToken() *string {
	return s.SessionToken
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken) SetAccessKeyId(v string) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken {
	s.AccessKeyId = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken) SetExpiration(v string) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken {
	s.Expiration = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken) SetSecretAccessKey(v string) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken {
	s.SecretAccessKey = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken) SetSessionToken(v string) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken {
	s.SessionToken = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAwsStsToken) Validate() error {
	return dara.Validate(s)
}

type ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken struct {
	Expiration   *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	TmpSecretId  *string `json:"tmpSecretId,omitempty" xml:"tmpSecretId,omitempty"`
	TmpSecretKey *string `json:"tmpSecretKey,omitempty" xml:"tmpSecretKey,omitempty"`
	Token        *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken) String() string {
	return dara.Prettify(s)
}

func (s ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken) GoString() string {
	return s.String()
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken) GetExpiration() *string {
	return s.Expiration
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken) GetTmpSecretId() *string {
	return s.TmpSecretId
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken) GetTmpSecretKey() *string {
	return s.TmpSecretKey
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken) GetToken() *string {
	return s.Token
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken) SetExpiration(v string) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken {
	s.Expiration = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken) SetTmpSecretId(v string) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken {
	s.TmpSecretId = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken) SetTmpSecretKey(v string) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken {
	s.TmpSecretKey = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken) SetToken(v string) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken {
	s.Token = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialTencentCloudStsToken) Validate() error {
	return dara.Validate(s)
}
