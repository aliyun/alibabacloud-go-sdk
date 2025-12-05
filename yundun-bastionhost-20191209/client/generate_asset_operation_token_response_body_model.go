// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAssetOperationTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetOperationToken(v *GenerateAssetOperationTokenResponseBodyAssetOperationToken) *GenerateAssetOperationTokenResponseBody
	GetAssetOperationToken() *GenerateAssetOperationTokenResponseBodyAssetOperationToken
	SetRequestId(v string) *GenerateAssetOperationTokenResponseBody
	GetRequestId() *string
}

type GenerateAssetOperationTokenResponseBody struct {
	// The data returned.
	AssetOperationToken *GenerateAssetOperationTokenResponseBodyAssetOperationToken `json:"AssetOperationToken,omitempty" xml:"AssetOperationToken,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateAssetOperationTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateAssetOperationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAssetOperationTokenResponseBody) GetAssetOperationToken() *GenerateAssetOperationTokenResponseBodyAssetOperationToken {
	return s.AssetOperationToken
}

func (s *GenerateAssetOperationTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateAssetOperationTokenResponseBody) SetAssetOperationToken(v *GenerateAssetOperationTokenResponseBodyAssetOperationToken) *GenerateAssetOperationTokenResponseBody {
	s.AssetOperationToken = v
	return s
}

func (s *GenerateAssetOperationTokenResponseBody) SetRequestId(v string) *GenerateAssetOperationTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateAssetOperationTokenResponseBody) Validate() error {
	if s.AssetOperationToken != nil {
		if err := s.AssetOperationToken.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateAssetOperationTokenResponseBodyAssetOperationToken struct {
	// The remaining number of times that you can use the O\\&M token.
	//
	// example:
	//
	// 1
	CountLeft *int64 `json:"CountLeft,omitempty" xml:"CountLeft,omitempty"`
	// The time when the O\\&M token expires. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1709110797
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the number of times that you can use the O\\&M token is limited.
	//
	// example:
	//
	// true
	HasCountLimit *bool `json:"HasCountLimit,omitempty" xml:"HasCountLimit,omitempty"`
	// The maximum number of renewals. A value of 0 indicates that renewal is not supported.
	//
	// example:
	//
	// 10
	MaxRenewCount *int64 `json:"MaxRenewCount,omitempty" xml:"MaxRenewCount,omitempty"`
	// The number of times the O\\&M token is renewed.
	//
	// example:
	//
	// 1
	RenewCount *int64 `json:"RenewCount,omitempty" xml:"RenewCount,omitempty"`
	// The single sign-on (SSO) URL.
	//
	// example:
	//
	// sso://eyJOT0RFX0NPTU1PTiI6eyJNb2R******
	SsoUrl *string `json:"SsoUrl,omitempty" xml:"SsoUrl,omitempty"`
	// The O\\&M token that you apply for.
	//
	// example:
	//
	// NmYyMmEzNmMwYzljNGY******
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The ID of the O\\&M token.
	//
	// example:
	//
	// 1
	TokenId *string `json:"TokenId,omitempty" xml:"TokenId,omitempty"`
}

func (s GenerateAssetOperationTokenResponseBodyAssetOperationToken) String() string {
	return dara.Prettify(s)
}

func (s GenerateAssetOperationTokenResponseBodyAssetOperationToken) GoString() string {
	return s.String()
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) GetCountLeft() *int64 {
	return s.CountLeft
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) GetHasCountLimit() *bool {
	return s.HasCountLimit
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) GetMaxRenewCount() *int64 {
	return s.MaxRenewCount
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) GetRenewCount() *int64 {
	return s.RenewCount
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) GetSsoUrl() *string {
	return s.SsoUrl
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) GetToken() *string {
	return s.Token
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) GetTokenId() *string {
	return s.TokenId
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) SetCountLeft(v int64) *GenerateAssetOperationTokenResponseBodyAssetOperationToken {
	s.CountLeft = &v
	return s
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) SetExpireTime(v int64) *GenerateAssetOperationTokenResponseBodyAssetOperationToken {
	s.ExpireTime = &v
	return s
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) SetHasCountLimit(v bool) *GenerateAssetOperationTokenResponseBodyAssetOperationToken {
	s.HasCountLimit = &v
	return s
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) SetMaxRenewCount(v int64) *GenerateAssetOperationTokenResponseBodyAssetOperationToken {
	s.MaxRenewCount = &v
	return s
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) SetRenewCount(v int64) *GenerateAssetOperationTokenResponseBodyAssetOperationToken {
	s.RenewCount = &v
	return s
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) SetSsoUrl(v string) *GenerateAssetOperationTokenResponseBodyAssetOperationToken {
	s.SsoUrl = &v
	return s
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) SetToken(v string) *GenerateAssetOperationTokenResponseBodyAssetOperationToken {
	s.Token = &v
	return s
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) SetTokenId(v string) *GenerateAssetOperationTokenResponseBodyAssetOperationToken {
	s.TokenId = &v
	return s
}

func (s *GenerateAssetOperationTokenResponseBodyAssetOperationToken) Validate() error {
	return dara.Validate(s)
}
