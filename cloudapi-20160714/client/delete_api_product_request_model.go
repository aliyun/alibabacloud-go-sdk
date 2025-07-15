// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiProductId(v string) *DeleteApiProductRequest
	GetApiProductId() *string
	SetSecurityToken(v string) *DeleteApiProductRequest
	GetSecurityToken() *string
}

type DeleteApiProductRequest struct {
	// The ID of the API product.
	//
	// This parameter is required.
	//
	// example:
	//
	// 117b7a64a8b3f064eaa4a47ac62aac5e
	ApiProductId  *string `json:"ApiProductId,omitempty" xml:"ApiProductId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteApiProductRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiProductRequest) GoString() string {
	return s.String()
}

func (s *DeleteApiProductRequest) GetApiProductId() *string {
	return s.ApiProductId
}

func (s *DeleteApiProductRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteApiProductRequest) SetApiProductId(v string) *DeleteApiProductRequest {
	s.ApiProductId = &v
	return s
}

func (s *DeleteApiProductRequest) SetSecurityToken(v string) *DeleteApiProductRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteApiProductRequest) Validate() error {
	return dara.Validate(s)
}
