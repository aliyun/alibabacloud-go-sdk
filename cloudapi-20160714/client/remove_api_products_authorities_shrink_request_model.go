// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApiProductsAuthoritiesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiProductIdsShrink(v string) *RemoveApiProductsAuthoritiesShrinkRequest
	GetApiProductIdsShrink() *string
	SetAppId(v int64) *RemoveApiProductsAuthoritiesShrinkRequest
	GetAppId() *int64
	SetSecurityToken(v string) *RemoveApiProductsAuthoritiesShrinkRequest
	GetSecurityToken() *string
}

type RemoveApiProductsAuthoritiesShrinkRequest struct {
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
	// 110982490
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RemoveApiProductsAuthoritiesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveApiProductsAuthoritiesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveApiProductsAuthoritiesShrinkRequest) GetApiProductIdsShrink() *string {
	return s.ApiProductIdsShrink
}

func (s *RemoveApiProductsAuthoritiesShrinkRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *RemoveApiProductsAuthoritiesShrinkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RemoveApiProductsAuthoritiesShrinkRequest) SetApiProductIdsShrink(v string) *RemoveApiProductsAuthoritiesShrinkRequest {
	s.ApiProductIdsShrink = &v
	return s
}

func (s *RemoveApiProductsAuthoritiesShrinkRequest) SetAppId(v int64) *RemoveApiProductsAuthoritiesShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RemoveApiProductsAuthoritiesShrinkRequest) SetSecurityToken(v string) *RemoveApiProductsAuthoritiesShrinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveApiProductsAuthoritiesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
