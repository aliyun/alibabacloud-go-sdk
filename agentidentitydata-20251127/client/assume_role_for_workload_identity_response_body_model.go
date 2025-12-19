// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeRoleForWorkloadIdentityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssumedRoleUser(v *AssumeRoleForWorkloadIdentityResponseBodyAssumedRoleUser) *AssumeRoleForWorkloadIdentityResponseBody
	GetAssumedRoleUser() *AssumeRoleForWorkloadIdentityResponseBodyAssumedRoleUser
	SetCredentials(v *AssumeRoleForWorkloadIdentityResponseBodyCredentials) *AssumeRoleForWorkloadIdentityResponseBody
	GetCredentials() *AssumeRoleForWorkloadIdentityResponseBodyCredentials
	SetRequestId(v string) *AssumeRoleForWorkloadIdentityResponseBody
	GetRequestId() *string
	SetWorkloadContextInfo(v *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfo) *AssumeRoleForWorkloadIdentityResponseBody
	GetWorkloadContextInfo() *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfo
}

type AssumeRoleForWorkloadIdentityResponseBody struct {
	AssumedRoleUser *AssumeRoleForWorkloadIdentityResponseBodyAssumedRoleUser `json:"AssumedRoleUser,omitempty" xml:"AssumedRoleUser,omitempty" type:"Struct"`
	Credentials     *AssumeRoleForWorkloadIdentityResponseBodyCredentials     `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
	// example:
	//
	// B7F68680-F3FE-5BF1-A824-4DE0B1AC9B23
	RequestId           *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkloadContextInfo *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfo `json:"WorkloadContextInfo,omitempty" xml:"WorkloadContextInfo,omitempty" type:"Struct"`
}

func (s AssumeRoleForWorkloadIdentityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleForWorkloadIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *AssumeRoleForWorkloadIdentityResponseBody) GetAssumedRoleUser() *AssumeRoleForWorkloadIdentityResponseBodyAssumedRoleUser {
	return s.AssumedRoleUser
}

func (s *AssumeRoleForWorkloadIdentityResponseBody) GetCredentials() *AssumeRoleForWorkloadIdentityResponseBodyCredentials {
	return s.Credentials
}

func (s *AssumeRoleForWorkloadIdentityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssumeRoleForWorkloadIdentityResponseBody) GetWorkloadContextInfo() *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfo {
	return s.WorkloadContextInfo
}

func (s *AssumeRoleForWorkloadIdentityResponseBody) SetAssumedRoleUser(v *AssumeRoleForWorkloadIdentityResponseBodyAssumedRoleUser) *AssumeRoleForWorkloadIdentityResponseBody {
	s.AssumedRoleUser = v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBody) SetCredentials(v *AssumeRoleForWorkloadIdentityResponseBodyCredentials) *AssumeRoleForWorkloadIdentityResponseBody {
	s.Credentials = v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBody) SetRequestId(v string) *AssumeRoleForWorkloadIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBody) SetWorkloadContextInfo(v *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfo) *AssumeRoleForWorkloadIdentityResponseBody {
	s.WorkloadContextInfo = v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBody) Validate() error {
	if s.AssumedRoleUser != nil {
		if err := s.AssumedRoleUser.Validate(); err != nil {
			return err
		}
	}
	if s.Credentials != nil {
		if err := s.Credentials.Validate(); err != nil {
			return err
		}
	}
	if s.WorkloadContextInfo != nil {
		if err := s.WorkloadContextInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssumeRoleForWorkloadIdentityResponseBodyAssumedRoleUser struct {
	// example:
	//
	// acs:ram::113511544585****:role/testoidc/TestOidcAssumedRoleSession
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// example:
	//
	// 33157794895460****
	AssumedRoleId *string `json:"AssumedRoleId,omitempty" xml:"AssumedRoleId,omitempty"`
}

func (s AssumeRoleForWorkloadIdentityResponseBodyAssumedRoleUser) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleForWorkloadIdentityResponseBodyAssumedRoleUser) GoString() string {
	return s.String()
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyAssumedRoleUser) GetArn() *string {
	return s.Arn
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyAssumedRoleUser) GetAssumedRoleId() *string {
	return s.AssumedRoleId
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyAssumedRoleUser) SetArn(v string) *AssumeRoleForWorkloadIdentityResponseBodyAssumedRoleUser {
	s.Arn = &v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyAssumedRoleUser) SetAssumedRoleId(v string) *AssumeRoleForWorkloadIdentityResponseBodyAssumedRoleUser {
	s.AssumedRoleId = &v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyAssumedRoleUser) Validate() error {
	return dara.Validate(s)
}

type AssumeRoleForWorkloadIdentityResponseBodyCredentials struct {
	// example:
	//
	// STS.NTTcrF2UNgeshFRTj9i7hpkah
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// example:
	//
	// 3fNFJk4UA6KLEG2GXiBe9ZE43ejFRJsea1yT1rguTVY6
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// example:
	//
	// 2025-04-28T03:06:19Z
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// example:
	//
	// CAIShwJ1q6Ft5B2yfSjIr5bSEsj4g7BihPWGWHz****
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AssumeRoleForWorkloadIdentityResponseBodyCredentials) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleForWorkloadIdentityResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyCredentials) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyCredentials) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyCredentials) GetExpiration() *string {
	return s.Expiration
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyCredentials) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyCredentials) SetAccessKeyId(v string) *AssumeRoleForWorkloadIdentityResponseBodyCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyCredentials) SetAccessKeySecret(v string) *AssumeRoleForWorkloadIdentityResponseBodyCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyCredentials) SetExpiration(v string) *AssumeRoleForWorkloadIdentityResponseBodyCredentials {
	s.Expiration = &v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyCredentials) SetSecurityToken(v string) *AssumeRoleForWorkloadIdentityResponseBodyCredentials {
	s.SecurityToken = &v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyCredentials) Validate() error {
	return dara.Validate(s)
}

type AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfo struct {
	UserContext *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
	// example:
	//
	// acs:agentidentity:cn-beijing:1778*****8462:workloadidentitydirectory/default/workloadidentity/workload-bc7***
	WorkloadIdentityArn *string `json:"WorkloadIdentityArn,omitempty" xml:"WorkloadIdentityArn,omitempty"`
}

func (s AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfo) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfo) GoString() string {
	return s.String()
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfo) GetUserContext() *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContext {
	return s.UserContext
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfo) GetWorkloadIdentityArn() *string {
	return s.WorkloadIdentityArn
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfo) SetUserContext(v *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContext) *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfo {
	s.UserContext = v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfo) SetWorkloadIdentityArn(v string) *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfo {
	s.WorkloadIdentityArn = &v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfo) Validate() error {
	if s.UserContext != nil {
		if err := s.UserContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContext struct {
	JwtClaims *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContextJwtClaims `json:"JwtClaims,omitempty" xml:"JwtClaims,omitempty" type:"Struct"`
	// example:
	//
	// externalUser
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContext) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContext) GoString() string {
	return s.String()
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContext) GetJwtClaims() *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContextJwtClaims {
	return s.JwtClaims
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContext) GetUserId() *string {
	return s.UserId
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContext) SetJwtClaims(v *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContextJwtClaims) *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContext {
	s.JwtClaims = v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContext) SetUserId(v string) *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContext {
	s.UserId = &v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContext) Validate() error {
	if s.JwtClaims != nil {
		if err := s.JwtClaims.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContextJwtClaims struct {
	// example:
	//
	// ["test-aud","test-aud-2"]
	Audiences *string `json:"Audiences,omitempty" xml:"Audiences,omitempty"`
	// example:
	//
	// https://oauth.aliyun.com
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// example:
	//
	// externalUser
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContextJwtClaims) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContextJwtClaims) GoString() string {
	return s.String()
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContextJwtClaims) GetAudiences() *string {
	return s.Audiences
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContextJwtClaims) GetIssuer() *string {
	return s.Issuer
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContextJwtClaims) GetSubject() *string {
	return s.Subject
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContextJwtClaims) SetAudiences(v string) *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContextJwtClaims {
	s.Audiences = &v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContextJwtClaims) SetIssuer(v string) *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContextJwtClaims {
	s.Issuer = &v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContextJwtClaims) SetSubject(v string) *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContextJwtClaims {
	s.Subject = &v
	return s
}

func (s *AssumeRoleForWorkloadIdentityResponseBodyWorkloadContextInfoUserContextJwtClaims) Validate() error {
	return dara.Validate(s)
}
