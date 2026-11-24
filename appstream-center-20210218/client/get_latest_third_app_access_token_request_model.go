// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLatestThirdAppAccessTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizerId(v string) *GetLatestThirdAppAccessTokenRequest
	GetAuthorizerId() *string
	SetBusinessChannel(v string) *GetLatestThirdAppAccessTokenRequest
	GetBusinessChannel() *string
	SetResource(v *GetLatestThirdAppAccessTokenRequestResource) *GetLatestThirdAppAccessTokenRequest
	GetResource() *GetLatestThirdAppAccessTokenRequestResource
}

type GetLatestThirdAppAccessTokenRequest struct {
	// The WyId of the user used when the token was created. It must correspond to the current account and the target cloud desktop, and cannot be replaced by a username.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1e2d3c4b5a6****
	AuthorizerId *string `json:"AuthorizerId,omitempty" xml:"AuthorizerId,omitempty"`
	// The business channel. Valid values: Enterprise and Business. The value is case-insensitive. If omitted, the default value is Enterprise. If an empty string or other value is explicitly passed, a parameter error is returned. This parameter does not replace user identity verification.
	//
	// example:
	//
	// Enterprise
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The target single cloud desktop resource. You must provide this object along with ResourceType and ResourceId when calling this operation. Pass it as a JSON object to specify only one cloud desktop. Wildcards or authorization for all cloud desktops are not supported.
	Resource *GetLatestThirdAppAccessTokenRequestResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
}

func (s GetLatestThirdAppAccessTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLatestThirdAppAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *GetLatestThirdAppAccessTokenRequest) GetAuthorizerId() *string {
	return s.AuthorizerId
}

func (s *GetLatestThirdAppAccessTokenRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *GetLatestThirdAppAccessTokenRequest) GetResource() *GetLatestThirdAppAccessTokenRequestResource {
	return s.Resource
}

func (s *GetLatestThirdAppAccessTokenRequest) SetAuthorizerId(v string) *GetLatestThirdAppAccessTokenRequest {
	s.AuthorizerId = &v
	return s
}

func (s *GetLatestThirdAppAccessTokenRequest) SetBusinessChannel(v string) *GetLatestThirdAppAccessTokenRequest {
	s.BusinessChannel = &v
	return s
}

func (s *GetLatestThirdAppAccessTokenRequest) SetResource(v *GetLatestThirdAppAccessTokenRequestResource) *GetLatestThirdAppAccessTokenRequest {
	s.Resource = v
	return s
}

func (s *GetLatestThirdAppAccessTokenRequest) Validate() error {
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLatestThirdAppAccessTokenRequestResource struct {
	// The ID of the cloud desktop specified when the token was created. You must provide a non-empty value that is consistent with the value used during creation. Asterisks (*) are not allowed.
	//
	// example:
	//
	// ecd-7a8b9c0d****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. This parameter is required. Valid value: DESKTOP, which indicates a cloud desktop. The value is case-insensitive.
	//
	// example:
	//
	// DESKTOP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetLatestThirdAppAccessTokenRequestResource) String() string {
	return dara.Prettify(s)
}

func (s GetLatestThirdAppAccessTokenRequestResource) GoString() string {
	return s.String()
}

func (s *GetLatestThirdAppAccessTokenRequestResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetLatestThirdAppAccessTokenRequestResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetLatestThirdAppAccessTokenRequestResource) SetResourceId(v string) *GetLatestThirdAppAccessTokenRequestResource {
	s.ResourceId = &v
	return s
}

func (s *GetLatestThirdAppAccessTokenRequestResource) SetResourceType(v string) *GetLatestThirdAppAccessTokenRequestResource {
	s.ResourceType = &v
	return s
}

func (s *GetLatestThirdAppAccessTokenRequestResource) Validate() error {
	return dara.Validate(s)
}
