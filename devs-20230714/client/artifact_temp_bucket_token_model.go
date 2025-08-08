// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArtifactTempBucketToken interface {
	dara.Model
	String() string
	GoString() string
	SetCredentials(v *ArtifactTempBucketTokenCredentials) *ArtifactTempBucketToken
	GetCredentials() *ArtifactTempBucketTokenCredentials
	SetOssBucketName(v string) *ArtifactTempBucketToken
	GetOssBucketName() *string
	SetOssObjectName(v string) *ArtifactTempBucketToken
	GetOssObjectName() *string
	SetOssRegion(v string) *ArtifactTempBucketToken
	GetOssRegion() *string
}

type ArtifactTempBucketToken struct {
	Credentials   *ArtifactTempBucketTokenCredentials `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Struct"`
	OssBucketName *string                             `json:"ossBucketName,omitempty" xml:"ossBucketName,omitempty"`
	OssObjectName *string                             `json:"ossObjectName,omitempty" xml:"ossObjectName,omitempty"`
	OssRegion     *string                             `json:"ossRegion,omitempty" xml:"ossRegion,omitempty"`
}

func (s ArtifactTempBucketToken) String() string {
	return dara.Prettify(s)
}

func (s ArtifactTempBucketToken) GoString() string {
	return s.String()
}

func (s *ArtifactTempBucketToken) GetCredentials() *ArtifactTempBucketTokenCredentials {
	return s.Credentials
}

func (s *ArtifactTempBucketToken) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *ArtifactTempBucketToken) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *ArtifactTempBucketToken) GetOssRegion() *string {
	return s.OssRegion
}

func (s *ArtifactTempBucketToken) SetCredentials(v *ArtifactTempBucketTokenCredentials) *ArtifactTempBucketToken {
	s.Credentials = v
	return s
}

func (s *ArtifactTempBucketToken) SetOssBucketName(v string) *ArtifactTempBucketToken {
	s.OssBucketName = &v
	return s
}

func (s *ArtifactTempBucketToken) SetOssObjectName(v string) *ArtifactTempBucketToken {
	s.OssObjectName = &v
	return s
}

func (s *ArtifactTempBucketToken) SetOssRegion(v string) *ArtifactTempBucketToken {
	s.OssRegion = &v
	return s
}

func (s *ArtifactTempBucketToken) Validate() error {
	return dara.Validate(s)
}

type ArtifactTempBucketTokenCredentials struct {
	AccessKeyId     *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	SecurityToken   *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
}

func (s ArtifactTempBucketTokenCredentials) String() string {
	return dara.Prettify(s)
}

func (s ArtifactTempBucketTokenCredentials) GoString() string {
	return s.String()
}

func (s *ArtifactTempBucketTokenCredentials) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *ArtifactTempBucketTokenCredentials) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *ArtifactTempBucketTokenCredentials) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ArtifactTempBucketTokenCredentials) SetAccessKeyId(v string) *ArtifactTempBucketTokenCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *ArtifactTempBucketTokenCredentials) SetAccessKeySecret(v string) *ArtifactTempBucketTokenCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *ArtifactTempBucketTokenCredentials) SetSecurityToken(v string) *ArtifactTempBucketTokenCredentials {
	s.SecurityToken = &v
	return s
}

func (s *ArtifactTempBucketTokenCredentials) Validate() error {
	return dara.Validate(s)
}
