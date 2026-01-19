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
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId                   *string                                                                             `json:"cloudAccountId,omitempty" xml:"cloudAccountId,omitempty"`
	CloudAccountRoleAccessCredential *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential `json:"cloudAccountRoleAccessCredential,omitempty" xml:"cloudAccountRoleAccessCredential,omitempty" type:"Struct"`
	// example:
	//
	// acs:ram::xxx:role/role-test
	CloudAccountRoleExternalId *string `json:"cloudAccountRoleExternalId,omitempty" xml:"cloudAccountRoleExternalId,omitempty"`
	// example:
	//
	// carole_01kmek49aqxxxx
	CloudAccountRoleId *string `json:"cloudAccountRoleId,omitempty" xml:"cloudAccountRoleId,omitempty"`
	// example:
	//
	// role-test
	CloudAccountRoleName *string `json:"cloudAccountRoleName,omitempty" xml:"cloudAccountRoleName,omitempty"`
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
	// example:
	//
	// 1767196800
	AccessCredentialExpiresAt *int64                                                                                                  `json:"accessCredentialExpiresAt,omitempty" xml:"accessCredentialExpiresAt,omitempty"`
	AlibabaCloudStsToken      *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken `json:"alibabaCloudStsToken,omitempty" xml:"alibabaCloudStsToken,omitempty" type:"Struct"`
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

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential) SetAccessCredentialExpiresAt(v int64) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential {
	s.AccessCredentialExpiresAt = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential) SetAlibabaCloudStsToken(v *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken) *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential {
	s.AlibabaCloudStsToken = v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredential) Validate() error {
	if s.AlibabaCloudStsToken != nil {
		if err := s.AlibabaCloudStsToken.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ObtainCloudAccountRoleAccessCredentialResponseBodyCloudAccountRoleAccessCredentialAlibabaCloudStsToken struct {
	// example:
	//
	// STS.NUgYrLnoC37mZZCNnAbez****
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// example:
	//
	// CVwjCkNzTMupZ8NbTCxCBRq3K16jtcWFTJAyBEv2****
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// example:
	//
	// 2021-10-20T04:27:09Z
	Expiration *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
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
