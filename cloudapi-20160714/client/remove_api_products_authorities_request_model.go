// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApiProductsAuthoritiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiProductIds(v []*string) *RemoveApiProductsAuthoritiesRequest
	GetApiProductIds() []*string
	SetAppId(v int64) *RemoveApiProductsAuthoritiesRequest
	GetAppId() *int64
	SetSecurityToken(v string) *RemoveApiProductsAuthoritiesRequest
	GetSecurityToken() *string
}

type RemoveApiProductsAuthoritiesRequest struct {
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
	// 110982490
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RemoveApiProductsAuthoritiesRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveApiProductsAuthoritiesRequest) GoString() string {
	return s.String()
}

func (s *RemoveApiProductsAuthoritiesRequest) GetApiProductIds() []*string {
	return s.ApiProductIds
}

func (s *RemoveApiProductsAuthoritiesRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *RemoveApiProductsAuthoritiesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RemoveApiProductsAuthoritiesRequest) SetApiProductIds(v []*string) *RemoveApiProductsAuthoritiesRequest {
	s.ApiProductIds = v
	return s
}

func (s *RemoveApiProductsAuthoritiesRequest) SetAppId(v int64) *RemoveApiProductsAuthoritiesRequest {
	s.AppId = &v
	return s
}

func (s *RemoveApiProductsAuthoritiesRequest) SetSecurityToken(v string) *RemoveApiProductsAuthoritiesRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveApiProductsAuthoritiesRequest) Validate() error {
	return dara.Validate(s)
}
