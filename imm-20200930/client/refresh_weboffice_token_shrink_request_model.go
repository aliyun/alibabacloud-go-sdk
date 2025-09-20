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
	// This parameter is required.
	//
	// example:
	//
	// 99d1b8b478b641c1b3372f5bd6********
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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
