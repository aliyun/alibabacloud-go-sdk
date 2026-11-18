// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationServerDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationServerId(v string) *UpdateAuthorizationServerDescriptionRequest
	GetAuthorizationServerId() *string
	SetClientToken(v string) *UpdateAuthorizationServerDescriptionRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateAuthorizationServerDescriptionRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateAuthorizationServerDescriptionRequest
	GetInstanceId() *string
}

type UpdateAuthorizationServerDescriptionRequest struct {
	// The ID of the authorization server.
	//
	// This parameter is required.
	//
	// example:
	//
	// iauths_system
	AuthorizationServerId *string `json:"AuthorizationServerId,omitempty" xml:"AuthorizationServerId,omitempty"`
	// The client token that is used to ensure idempotence. You can generate a client token on your client and make sure that the token is unique across different requests. The client token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/en/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// abc123xyz
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the authorization server.
	//
	// example:
	//
	// description of authorization server
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateAuthorizationServerDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationServerDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationServerDescriptionRequest) GetAuthorizationServerId() *string {
	return s.AuthorizationServerId
}

func (s *UpdateAuthorizationServerDescriptionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAuthorizationServerDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAuthorizationServerDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAuthorizationServerDescriptionRequest) SetAuthorizationServerId(v string) *UpdateAuthorizationServerDescriptionRequest {
	s.AuthorizationServerId = &v
	return s
}

func (s *UpdateAuthorizationServerDescriptionRequest) SetClientToken(v string) *UpdateAuthorizationServerDescriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAuthorizationServerDescriptionRequest) SetDescription(v string) *UpdateAuthorizationServerDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateAuthorizationServerDescriptionRequest) SetInstanceId(v string) *UpdateAuthorizationServerDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAuthorizationServerDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
