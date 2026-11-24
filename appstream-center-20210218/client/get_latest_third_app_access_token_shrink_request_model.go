// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLatestThirdAppAccessTokenShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizerId(v string) *GetLatestThirdAppAccessTokenShrinkRequest
	GetAuthorizerId() *string
	SetBusinessChannel(v string) *GetLatestThirdAppAccessTokenShrinkRequest
	GetBusinessChannel() *string
	SetResourceShrink(v string) *GetLatestThirdAppAccessTokenShrinkRequest
	GetResourceShrink() *string
}

type GetLatestThirdAppAccessTokenShrinkRequest struct {
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
	ResourceShrink *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s GetLatestThirdAppAccessTokenShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLatestThirdAppAccessTokenShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetLatestThirdAppAccessTokenShrinkRequest) GetAuthorizerId() *string {
	return s.AuthorizerId
}

func (s *GetLatestThirdAppAccessTokenShrinkRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *GetLatestThirdAppAccessTokenShrinkRequest) GetResourceShrink() *string {
	return s.ResourceShrink
}

func (s *GetLatestThirdAppAccessTokenShrinkRequest) SetAuthorizerId(v string) *GetLatestThirdAppAccessTokenShrinkRequest {
	s.AuthorizerId = &v
	return s
}

func (s *GetLatestThirdAppAccessTokenShrinkRequest) SetBusinessChannel(v string) *GetLatestThirdAppAccessTokenShrinkRequest {
	s.BusinessChannel = &v
	return s
}

func (s *GetLatestThirdAppAccessTokenShrinkRequest) SetResourceShrink(v string) *GetLatestThirdAppAccessTokenShrinkRequest {
	s.ResourceShrink = &v
	return s
}

func (s *GetLatestThirdAppAccessTokenShrinkRequest) Validate() error {
	return dara.Validate(s)
}
