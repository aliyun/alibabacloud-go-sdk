// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApiProductsAuthoritiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiProductIds(v []*string) *SetApiProductsAuthoritiesRequest
	GetApiProductIds() []*string
	SetAppId(v int64) *SetApiProductsAuthoritiesRequest
	GetAppId() *int64
	SetAuthValidTime(v string) *SetApiProductsAuthoritiesRequest
	GetAuthValidTime() *string
	SetDescription(v string) *SetApiProductsAuthoritiesRequest
	GetDescription() *string
	SetSecurityToken(v string) *SetApiProductsAuthoritiesRequest
	GetSecurityToken() *string
}

type SetApiProductsAuthoritiesRequest struct {
	// The API products.
	//
	// This parameter is required.
	ApiProductIds []*string `json:"ApiProductIds,omitempty" xml:"ApiProductIds,omitempty" type:"Repeated"`
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

func (s SetApiProductsAuthoritiesRequest) String() string {
	return dara.Prettify(s)
}

func (s SetApiProductsAuthoritiesRequest) GoString() string {
	return s.String()
}

func (s *SetApiProductsAuthoritiesRequest) GetApiProductIds() []*string {
	return s.ApiProductIds
}

func (s *SetApiProductsAuthoritiesRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *SetApiProductsAuthoritiesRequest) GetAuthValidTime() *string {
	return s.AuthValidTime
}

func (s *SetApiProductsAuthoritiesRequest) GetDescription() *string {
	return s.Description
}

func (s *SetApiProductsAuthoritiesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetApiProductsAuthoritiesRequest) SetApiProductIds(v []*string) *SetApiProductsAuthoritiesRequest {
	s.ApiProductIds = v
	return s
}

func (s *SetApiProductsAuthoritiesRequest) SetAppId(v int64) *SetApiProductsAuthoritiesRequest {
	s.AppId = &v
	return s
}

func (s *SetApiProductsAuthoritiesRequest) SetAuthValidTime(v string) *SetApiProductsAuthoritiesRequest {
	s.AuthValidTime = &v
	return s
}

func (s *SetApiProductsAuthoritiesRequest) SetDescription(v string) *SetApiProductsAuthoritiesRequest {
	s.Description = &v
	return s
}

func (s *SetApiProductsAuthoritiesRequest) SetSecurityToken(v string) *SetApiProductsAuthoritiesRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetApiProductsAuthoritiesRequest) Validate() error {
	return dara.Validate(s)
}
