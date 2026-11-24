// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateThirdAppAccessTokenShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizerId(v string) *CreateThirdAppAccessTokenShrinkRequest
	GetAuthorizerId() *string
	SetBusinessChannel(v string) *CreateThirdAppAccessTokenShrinkRequest
	GetBusinessChannel() *string
	SetExpiration(v int32) *CreateThirdAppAccessTokenShrinkRequest
	GetExpiration() *int32
	SetResourceShrink(v string) *CreateThirdAppAccessTokenShrinkRequest
	GetResourceShrink() *string
}

type CreateThirdAppAccessTokenShrinkRequest struct {
	// The WyId of the authorized user. The user must belong to the current calling account. This parameter is not the username.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1e2d3c4b5a6****
	AuthorizerId *string `json:"AuthorizerId,omitempty" xml:"AuthorizerId,omitempty"`
	// The business channel. Valid values: Enterprise and Business. The value is case-insensitive. If omitted, the default value is Enterprise. If an empty string or other values are explicitly passed, a parameter error is returned. This parameter does not replace user identity verification.
	//
	// example:
	//
	// Enterprise
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The validity period of the token, in seconds. Valid values: 1 to 15552000. If omitted, the default value is 15552000, which equals 180 days. Permanent validity is not supported, and 0 cannot be passed.
	//
	// example:
	//
	// 900
	Expiration *int32 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The target single cloud desktop resource. You must provide this object along with ResourceType and ResourceId when you call this operation. Pass this parameter as a JSON object and specify only one cloud desktop. Wildcards or authorization for all cloud desktops are not supported.
	ResourceShrink *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s CreateThirdAppAccessTokenShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateThirdAppAccessTokenShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateThirdAppAccessTokenShrinkRequest) GetAuthorizerId() *string {
	return s.AuthorizerId
}

func (s *CreateThirdAppAccessTokenShrinkRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *CreateThirdAppAccessTokenShrinkRequest) GetExpiration() *int32 {
	return s.Expiration
}

func (s *CreateThirdAppAccessTokenShrinkRequest) GetResourceShrink() *string {
	return s.ResourceShrink
}

func (s *CreateThirdAppAccessTokenShrinkRequest) SetAuthorizerId(v string) *CreateThirdAppAccessTokenShrinkRequest {
	s.AuthorizerId = &v
	return s
}

func (s *CreateThirdAppAccessTokenShrinkRequest) SetBusinessChannel(v string) *CreateThirdAppAccessTokenShrinkRequest {
	s.BusinessChannel = &v
	return s
}

func (s *CreateThirdAppAccessTokenShrinkRequest) SetExpiration(v int32) *CreateThirdAppAccessTokenShrinkRequest {
	s.Expiration = &v
	return s
}

func (s *CreateThirdAppAccessTokenShrinkRequest) SetResourceShrink(v string) *CreateThirdAppAccessTokenShrinkRequest {
	s.ResourceShrink = &v
	return s
}

func (s *CreateThirdAppAccessTokenShrinkRequest) Validate() error {
	return dara.Validate(s)
}
