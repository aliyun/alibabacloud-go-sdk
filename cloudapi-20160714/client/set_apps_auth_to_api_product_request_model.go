// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAppsAuthToApiProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiProductId(v string) *SetAppsAuthToApiProductRequest
	GetApiProductId() *string
	SetAppIds(v []*int64) *SetAppsAuthToApiProductRequest
	GetAppIds() []*int64
	SetAuthValidTime(v string) *SetAppsAuthToApiProductRequest
	GetAuthValidTime() *string
	SetDescription(v string) *SetAppsAuthToApiProductRequest
	GetDescription() *string
	SetSecurityToken(v string) *SetAppsAuthToApiProductRequest
	GetSecurityToken() *string
}

type SetAppsAuthToApiProductRequest struct {
	// The ID of the API product.
	//
	// This parameter is required.
	//
	// example:
	//
	// 117b7a64a8b3f064eaa4a47ac62aac5e
	ApiProductId *string `json:"ApiProductId,omitempty" xml:"ApiProductId,omitempty"`
	// The IDs of the applications that you want to authorize.
	//
	// This parameter is required.
	AppIds []*int64 `json:"AppIds,omitempty" xml:"AppIds,omitempty" type:"Repeated"`
	// The time (UTC) when the authorization expires. If this parameter is empty, the authorization does not expire.
	//
	// example:
	//
	// 2023-05-31T08:15:39Z
	AuthValidTime *string `json:"AuthValidTime,omitempty" xml:"AuthValidTime,omitempty"`
	// The description of the authorization.
	//
	// example:
	//
	// Test
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetAppsAuthToApiProductRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAppsAuthToApiProductRequest) GoString() string {
	return s.String()
}

func (s *SetAppsAuthToApiProductRequest) GetApiProductId() *string {
	return s.ApiProductId
}

func (s *SetAppsAuthToApiProductRequest) GetAppIds() []*int64 {
	return s.AppIds
}

func (s *SetAppsAuthToApiProductRequest) GetAuthValidTime() *string {
	return s.AuthValidTime
}

func (s *SetAppsAuthToApiProductRequest) GetDescription() *string {
	return s.Description
}

func (s *SetAppsAuthToApiProductRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetAppsAuthToApiProductRequest) SetApiProductId(v string) *SetAppsAuthToApiProductRequest {
	s.ApiProductId = &v
	return s
}

func (s *SetAppsAuthToApiProductRequest) SetAppIds(v []*int64) *SetAppsAuthToApiProductRequest {
	s.AppIds = v
	return s
}

func (s *SetAppsAuthToApiProductRequest) SetAuthValidTime(v string) *SetAppsAuthToApiProductRequest {
	s.AuthValidTime = &v
	return s
}

func (s *SetAppsAuthToApiProductRequest) SetDescription(v string) *SetAppsAuthToApiProductRequest {
	s.Description = &v
	return s
}

func (s *SetAppsAuthToApiProductRequest) SetSecurityToken(v string) *SetAppsAuthToApiProductRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetAppsAuthToApiProductRequest) Validate() error {
	return dara.Validate(s)
}
