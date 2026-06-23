// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRuleId(v string) *DisableAuthorizationRuleRequest
	GetAuthorizationRuleId() *string
	SetClientToken(v string) *DisableAuthorizationRuleRequest
	GetClientToken() *string
	SetInstanceId(v string) *DisableAuthorizationRuleRequest
	GetInstanceId() *string
}

type DisableAuthorizationRuleRequest struct {
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
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableAuthorizationRuleRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *DisableAuthorizationRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableAuthorizationRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *DisableAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *DisableAuthorizationRuleRequest) SetClientToken(v string) *DisableAuthorizationRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableAuthorizationRuleRequest) SetInstanceId(v string) *DisableAuthorizationRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableAuthorizationRuleRequest) Validate() error {
	return dara.Validate(s)
}
