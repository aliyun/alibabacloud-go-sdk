// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterRenewApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireAt(v string) *ModelRouterRenewApiKeyRequest
	GetExpireAt() *string
}

type ModelRouterRenewApiKeyRequest struct {
	// The new expiration time in RFC 3339 format. The time must be later than the current time. If this parameter is not specified or is set to null, the API key remains valid indefinitely. This parameter only modifies the validity period and does not change the enabled or disabled status.
	//
	// example:
	//
	// 2027-01-01T00:00:00+08:00
	ExpireAt *string `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
}

func (s ModelRouterRenewApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterRenewApiKeyRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterRenewApiKeyRequest) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *ModelRouterRenewApiKeyRequest) SetExpireAt(v string) *ModelRouterRenewApiKeyRequest {
	s.ExpireAt = &v
	return s
}

func (s *ModelRouterRenewApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
