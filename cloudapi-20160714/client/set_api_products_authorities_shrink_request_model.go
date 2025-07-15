// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApiProductsAuthoritiesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiProductIdsShrink(v string) *SetApiProductsAuthoritiesShrinkRequest
	GetApiProductIdsShrink() *string
	SetAppId(v int64) *SetApiProductsAuthoritiesShrinkRequest
	GetAppId() *int64
	SetAuthValidTime(v string) *SetApiProductsAuthoritiesShrinkRequest
	GetAuthValidTime() *string
	SetDescription(v string) *SetApiProductsAuthoritiesShrinkRequest
	GetDescription() *string
	SetSecurityToken(v string) *SetApiProductsAuthoritiesShrinkRequest
	GetSecurityToken() *string
}

type SetApiProductsAuthoritiesShrinkRequest struct {
	// The API products.
	//
	// This parameter is required.
	ApiProductIdsShrink *string `json:"ApiProductIds,omitempty" xml:"ApiProductIds,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 111385984
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 授权有效时间的截止时间，请设置格林尼治标准时间(GMT), 如果为空，即为授权永久有效。
	//
	// example:
	//
	// 2023-06-12T03:07:37Z
	AuthValidTime *string `json:"AuthValidTime,omitempty" xml:"AuthValidTime,omitempty"`
	// The authorization description.
	//
	// example:
	//
	// test
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetApiProductsAuthoritiesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetApiProductsAuthoritiesShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetApiProductsAuthoritiesShrinkRequest) GetApiProductIdsShrink() *string {
	return s.ApiProductIdsShrink
}

func (s *SetApiProductsAuthoritiesShrinkRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *SetApiProductsAuthoritiesShrinkRequest) GetAuthValidTime() *string {
	return s.AuthValidTime
}

func (s *SetApiProductsAuthoritiesShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *SetApiProductsAuthoritiesShrinkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetApiProductsAuthoritiesShrinkRequest) SetApiProductIdsShrink(v string) *SetApiProductsAuthoritiesShrinkRequest {
	s.ApiProductIdsShrink = &v
	return s
}

func (s *SetApiProductsAuthoritiesShrinkRequest) SetAppId(v int64) *SetApiProductsAuthoritiesShrinkRequest {
	s.AppId = &v
	return s
}

func (s *SetApiProductsAuthoritiesShrinkRequest) SetAuthValidTime(v string) *SetApiProductsAuthoritiesShrinkRequest {
	s.AuthValidTime = &v
	return s
}

func (s *SetApiProductsAuthoritiesShrinkRequest) SetDescription(v string) *SetApiProductsAuthoritiesShrinkRequest {
	s.Description = &v
	return s
}

func (s *SetApiProductsAuthoritiesShrinkRequest) SetSecurityToken(v string) *SetApiProductsAuthoritiesShrinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetApiProductsAuthoritiesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
