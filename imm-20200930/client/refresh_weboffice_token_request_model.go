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
	// Weboffice access token. Obtain it through the [GenerateWebofficeToken](https://help.aliyun.com/document_detail/478226.html) or [RefreshWebofficeToken](https://help.aliyun.com/document_detail/478227.html) interfaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// 99d1b8b478b641c1b3372f5bd6********
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// **If there are no special requirements, leave it blank.**
	//
	// Chained authorization configuration, optional. For more information, see [Access Other Entity Resources Using Chained Authorization](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// Project name. For more information on how to obtain it, see [Create Project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Weboffice refresh token. Obtain it through the [GenerateWebofficeToken](https://help.aliyun.com/document_detail/478226.html) or [RefreshWebofficeToken](https://help.aliyun.com/document_detail/478227.html) interfaces.
	//
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
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
