// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateThirdAppAccessTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizerId(v string) *CreateThirdAppAccessTokenRequest
	GetAuthorizerId() *string
	SetBusinessChannel(v string) *CreateThirdAppAccessTokenRequest
	GetBusinessChannel() *string
	SetExpiration(v int32) *CreateThirdAppAccessTokenRequest
	GetExpiration() *int32
	SetResource(v *CreateThirdAppAccessTokenRequestResource) *CreateThirdAppAccessTokenRequest
	GetResource() *CreateThirdAppAccessTokenRequestResource
}

type CreateThirdAppAccessTokenRequest struct {
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
	Resource *CreateThirdAppAccessTokenRequestResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
}

func (s CreateThirdAppAccessTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateThirdAppAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateThirdAppAccessTokenRequest) GetAuthorizerId() *string {
	return s.AuthorizerId
}

func (s *CreateThirdAppAccessTokenRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *CreateThirdAppAccessTokenRequest) GetExpiration() *int32 {
	return s.Expiration
}

func (s *CreateThirdAppAccessTokenRequest) GetResource() *CreateThirdAppAccessTokenRequestResource {
	return s.Resource
}

func (s *CreateThirdAppAccessTokenRequest) SetAuthorizerId(v string) *CreateThirdAppAccessTokenRequest {
	s.AuthorizerId = &v
	return s
}

func (s *CreateThirdAppAccessTokenRequest) SetBusinessChannel(v string) *CreateThirdAppAccessTokenRequest {
	s.BusinessChannel = &v
	return s
}

func (s *CreateThirdAppAccessTokenRequest) SetExpiration(v int32) *CreateThirdAppAccessTokenRequest {
	s.Expiration = &v
	return s
}

func (s *CreateThirdAppAccessTokenRequest) SetResource(v *CreateThirdAppAccessTokenRequestResource) *CreateThirdAppAccessTokenRequest {
	s.Resource = v
	return s
}

func (s *CreateThirdAppAccessTokenRequest) Validate() error {
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateThirdAppAccessTokenRequestResource struct {
	// The ID of the target cloud desktop. You must provide a non-empty value when you call this operation. Asterisks (*) are not allowed. The cloud desktop must be authorized to the user specified by AuthorizerId.
	//
	// example:
	//
	// ecd-7a8b9c0d****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. This parameter is required when you call this operation. Set the value to DESKTOP, which indicates a cloud desktop. The value is case-insensitive.
	//
	// example:
	//
	// DESKTOP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateThirdAppAccessTokenRequestResource) String() string {
	return dara.Prettify(s)
}

func (s CreateThirdAppAccessTokenRequestResource) GoString() string {
	return s.String()
}

func (s *CreateThirdAppAccessTokenRequestResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateThirdAppAccessTokenRequestResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateThirdAppAccessTokenRequestResource) SetResourceId(v string) *CreateThirdAppAccessTokenRequestResource {
	s.ResourceId = &v
	return s
}

func (s *CreateThirdAppAccessTokenRequestResource) SetResourceType(v string) *CreateThirdAppAccessTokenRequestResource {
	s.ResourceType = &v
	return s
}

func (s *CreateThirdAppAccessTokenRequestResource) Validate() error {
	return dara.Validate(s)
}
