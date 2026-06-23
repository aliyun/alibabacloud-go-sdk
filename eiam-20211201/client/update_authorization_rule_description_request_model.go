// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationRuleDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRuleId(v string) *UpdateAuthorizationRuleDescriptionRequest
	GetAuthorizationRuleId() *string
	SetClientToken(v string) *UpdateAuthorizationRuleDescriptionRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateAuthorizationRuleDescriptionRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateAuthorizationRuleDescriptionRequest
	GetInstanceId() *string
}

type UpdateAuthorizationRuleDescriptionRequest struct {
	// The authorization rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate a parameter value, but make sure that the value is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see References: [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the authorization rule. The description can be up to 128 characters in length.
	//
	// example:
	//
	// this is a test rule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateAuthorizationRuleDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleDescriptionRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *UpdateAuthorizationRuleDescriptionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAuthorizationRuleDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAuthorizationRuleDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAuthorizationRuleDescriptionRequest) SetAuthorizationRuleId(v string) *UpdateAuthorizationRuleDescriptionRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *UpdateAuthorizationRuleDescriptionRequest) SetClientToken(v string) *UpdateAuthorizationRuleDescriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAuthorizationRuleDescriptionRequest) SetDescription(v string) *UpdateAuthorizationRuleDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateAuthorizationRuleDescriptionRequest) SetInstanceId(v string) *UpdateAuthorizationRuleDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAuthorizationRuleDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
