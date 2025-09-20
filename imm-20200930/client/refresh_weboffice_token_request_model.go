// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshWebofficeTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *RefreshWebofficeTokenRequest
	GetAccessToken() *string
	SetCredentialConfig(v *CredentialConfig) *RefreshWebofficeTokenRequest
	GetCredentialConfig() *CredentialConfig
	SetProjectName(v string) *RefreshWebofficeTokenRequest
	GetProjectName() *string
	SetRefreshToken(v string) *RefreshWebofficeTokenRequest
	GetRefreshToken() *string
}

type RefreshWebofficeTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 99d1b8b478b641c1b3372f5bd6********
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a730ae0d7c6a487d87c661d199********
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
}

func (s RefreshWebofficeTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshWebofficeTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *RefreshWebofficeTokenRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *RefreshWebofficeTokenRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *RefreshWebofficeTokenRequest) GetRefreshToken() *string {
	return s.RefreshToken
}

func (s *RefreshWebofficeTokenRequest) SetAccessToken(v string) *RefreshWebofficeTokenRequest {
	s.AccessToken = &v
	return s
}

func (s *RefreshWebofficeTokenRequest) SetCredentialConfig(v *CredentialConfig) *RefreshWebofficeTokenRequest {
	s.CredentialConfig = v
	return s
}

func (s *RefreshWebofficeTokenRequest) SetProjectName(v string) *RefreshWebofficeTokenRequest {
	s.ProjectName = &v
	return s
}

func (s *RefreshWebofficeTokenRequest) SetRefreshToken(v string) *RefreshWebofficeTokenRequest {
	s.RefreshToken = &v
	return s
}

func (s *RefreshWebofficeTokenRequest) Validate() error {
	return dara.Validate(s)
}
