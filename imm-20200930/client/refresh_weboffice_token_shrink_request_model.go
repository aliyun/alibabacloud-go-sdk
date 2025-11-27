// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshWebofficeTokenShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *RefreshWebofficeTokenShrinkRequest
	GetAccessToken() *string
	SetCredentialConfigShrink(v string) *RefreshWebofficeTokenShrinkRequest
	GetCredentialConfigShrink() *string
	SetProjectName(v string) *RefreshWebofficeTokenShrinkRequest
	GetProjectName() *string
	SetRefreshToken(v string) *RefreshWebofficeTokenShrinkRequest
	GetRefreshToken() *string
}

type RefreshWebofficeTokenShrinkRequest struct {
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
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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

func (s RefreshWebofficeTokenShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshWebofficeTokenShrinkRequest) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenShrinkRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *RefreshWebofficeTokenShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *RefreshWebofficeTokenShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *RefreshWebofficeTokenShrinkRequest) GetRefreshToken() *string {
	return s.RefreshToken
}

func (s *RefreshWebofficeTokenShrinkRequest) SetAccessToken(v string) *RefreshWebofficeTokenShrinkRequest {
	s.AccessToken = &v
	return s
}

func (s *RefreshWebofficeTokenShrinkRequest) SetCredentialConfigShrink(v string) *RefreshWebofficeTokenShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *RefreshWebofficeTokenShrinkRequest) SetProjectName(v string) *RefreshWebofficeTokenShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *RefreshWebofficeTokenShrinkRequest) SetRefreshToken(v string) *RefreshWebofficeTokenShrinkRequest {
	s.RefreshToken = &v
	return s
}

func (s *RefreshWebofficeTokenShrinkRequest) Validate() error {
	return dara.Validate(s)
}
