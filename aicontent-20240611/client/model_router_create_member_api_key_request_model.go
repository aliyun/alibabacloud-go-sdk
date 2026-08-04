// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateMemberApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireAt(v string) *ModelRouterCreateMemberApiKeyRequest
	GetExpireAt() *string
	SetName(v string) *ModelRouterCreateMemberApiKeyRequest
	GetName() *string
}

type ModelRouterCreateMemberApiKeyRequest struct {
	// The expiration time in the format of yyyy-MM-dd HH:mm:ss. This parameter is optional. If not specified, the key is permanently valid.
	//
	// example:
	//
	// 2027-07-31 00:00:00
	ExpireAt *string `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	// The name of the API key. This parameter is optional.
	//
	// example:
	//
	// TestKey
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ModelRouterCreateMemberApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateMemberApiKeyRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateMemberApiKeyRequest) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *ModelRouterCreateMemberApiKeyRequest) GetName() *string {
	return s.Name
}

func (s *ModelRouterCreateMemberApiKeyRequest) SetExpireAt(v string) *ModelRouterCreateMemberApiKeyRequest {
	s.ExpireAt = &v
	return s
}

func (s *ModelRouterCreateMemberApiKeyRequest) SetName(v string) *ModelRouterCreateMemberApiKeyRequest {
	s.Name = &v
	return s
}

func (s *ModelRouterCreateMemberApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
