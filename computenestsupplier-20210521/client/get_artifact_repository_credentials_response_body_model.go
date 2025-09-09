// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactRepositoryCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableResources(v []*GetArtifactRepositoryCredentialsResponseBodyAvailableResources) *GetArtifactRepositoryCredentialsResponseBody
	GetAvailableResources() []*GetArtifactRepositoryCredentialsResponseBodyAvailableResources
	SetCredentials(v *GetArtifactRepositoryCredentialsResponseBodyCredentials) *GetArtifactRepositoryCredentialsResponseBody
	GetCredentials() *GetArtifactRepositoryCredentialsResponseBodyCredentials
	SetExpireDate(v string) *GetArtifactRepositoryCredentialsResponseBody
	GetExpireDate() *string
	SetRequestId(v string) *GetArtifactRepositoryCredentialsResponseBody
	GetRequestId() *string
}

type GetArtifactRepositoryCredentialsResponseBody struct {
	// The information about the resources that can be uploaded.
	AvailableResources []*GetArtifactRepositoryCredentialsResponseBodyAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Repeated"`
	// The credentials.
	Credentials *GetArtifactRepositoryCredentialsResponseBodyCredentials `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
	// The time when the credentials expired.
	//
	// example:
	//
	// 1526549792000
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 94E89857-B994-44B6-9C4F-DBD200E9XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetArtifactRepositoryCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsResponseBody) GetAvailableResources() []*GetArtifactRepositoryCredentialsResponseBodyAvailableResources {
	return s.AvailableResources
}

func (s *GetArtifactRepositoryCredentialsResponseBody) GetCredentials() *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	return s.Credentials
}

func (s *GetArtifactRepositoryCredentialsResponseBody) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *GetArtifactRepositoryCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetArtifactRepositoryCredentialsResponseBody) SetAvailableResources(v []*GetArtifactRepositoryCredentialsResponseBodyAvailableResources) *GetArtifactRepositoryCredentialsResponseBody {
	s.AvailableResources = v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBody) SetCredentials(v *GetArtifactRepositoryCredentialsResponseBodyCredentials) *GetArtifactRepositoryCredentialsResponseBody {
	s.Credentials = v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBody) SetExpireDate(v string) *GetArtifactRepositoryCredentialsResponseBody {
	s.ExpireDate = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBody) SetRequestId(v string) *GetArtifactRepositoryCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetArtifactRepositoryCredentialsResponseBodyAvailableResources struct {
	// The path.
	//
	// example:
	//
	// "/xxx/"
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The region ID.
	//
	// example:
	//
	// oss-cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The repository name.
	//
	// example:
	//
	// computenest-artifacts-draft-cn-hangzhou
	RepositoryName *string `json:"RepositoryName,omitempty" xml:"RepositoryName,omitempty"`
}

func (s GetArtifactRepositoryCredentialsResponseBodyAvailableResources) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsResponseBodyAvailableResources) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsResponseBodyAvailableResources) GetPath() *string {
	return s.Path
}

func (s *GetArtifactRepositoryCredentialsResponseBodyAvailableResources) GetRegionId() *string {
	return s.RegionId
}

func (s *GetArtifactRepositoryCredentialsResponseBodyAvailableResources) GetRepositoryName() *string {
	return s.RepositoryName
}

func (s *GetArtifactRepositoryCredentialsResponseBodyAvailableResources) SetPath(v string) *GetArtifactRepositoryCredentialsResponseBodyAvailableResources {
	s.Path = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyAvailableResources) SetRegionId(v string) *GetArtifactRepositoryCredentialsResponseBodyAvailableResources {
	s.RegionId = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyAvailableResources) SetRepositoryName(v string) *GetArtifactRepositoryCredentialsResponseBodyAvailableResources {
	s.RepositoryName = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyAvailableResources) Validate() error {
	return dara.Validate(s)
}

type GetArtifactRepositoryCredentialsResponseBodyCredentials struct {
	// The AccessKey ID.
	//
	// example:
	//
	// STS.xxx
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret.
	//
	// example:
	//
	// xxx
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The password.
	//
	// example:
	//
	// eyJ0aW1lIjoiMTUyNjU0OTc5:0705733****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The Security Token Service (STS) token.
	//
	// example:
	//
	// xxx
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The username.
	//
	// example:
	//
	// xxx
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetArtifactRepositoryCredentialsResponseBodyCredentials) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) GetPassword() *string {
	return s.Password
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) GetUsername() *string {
	return s.Username
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetAccessKeyId(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetAccessKeySecret(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetPassword(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.Password = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetSecurityToken(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.SecurityToken = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetUsername(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.Username = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) Validate() error {
	return dara.Validate(s)
}
