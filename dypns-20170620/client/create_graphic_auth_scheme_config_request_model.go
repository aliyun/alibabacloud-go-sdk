// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGraphicAuthSchemeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *CreateGraphicAuthSchemeConfigRequest
	GetOwnerId() *int64
	SetPlatform(v string) *CreateGraphicAuthSchemeConfigRequest
	GetPlatform() *string
	SetResourceOwnerAccount(v string) *CreateGraphicAuthSchemeConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateGraphicAuthSchemeConfigRequest
	GetResourceOwnerId() *int64
	SetSchemeName(v string) *CreateGraphicAuthSchemeConfigRequest
	GetSchemeName() *string
}

type CreateGraphicAuthSchemeConfigRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// Web
	Platform             *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	SchemeName *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
}

func (s CreateGraphicAuthSchemeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGraphicAuthSchemeConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateGraphicAuthSchemeConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateGraphicAuthSchemeConfigRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateGraphicAuthSchemeConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateGraphicAuthSchemeConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateGraphicAuthSchemeConfigRequest) GetSchemeName() *string {
	return s.SchemeName
}

func (s *CreateGraphicAuthSchemeConfigRequest) SetOwnerId(v int64) *CreateGraphicAuthSchemeConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGraphicAuthSchemeConfigRequest) SetPlatform(v string) *CreateGraphicAuthSchemeConfigRequest {
	s.Platform = &v
	return s
}

func (s *CreateGraphicAuthSchemeConfigRequest) SetResourceOwnerAccount(v string) *CreateGraphicAuthSchemeConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateGraphicAuthSchemeConfigRequest) SetResourceOwnerId(v int64) *CreateGraphicAuthSchemeConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateGraphicAuthSchemeConfigRequest) SetSchemeName(v string) *CreateGraphicAuthSchemeConfigRequest {
	s.SchemeName = &v
	return s
}

func (s *CreateGraphicAuthSchemeConfigRequest) Validate() error {
	return dara.Validate(s)
}
