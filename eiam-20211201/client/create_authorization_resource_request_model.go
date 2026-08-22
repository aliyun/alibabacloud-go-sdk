// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAuthorizationResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationResourceEntityId(v string) *CreateAuthorizationResourceRequest
	GetAuthorizationResourceEntityId() *string
	SetAuthorizationResourceEntityType(v string) *CreateAuthorizationResourceRequest
	GetAuthorizationResourceEntityType() *string
	SetAuthorizationRuleId(v string) *CreateAuthorizationResourceRequest
	GetAuthorizationRuleId() *string
	SetClientToken(v string) *CreateAuthorizationResourceRequest
	GetClientToken() *string
	SetCondition(v *CreateAuthorizationResourceRequestCondition) *CreateAuthorizationResourceRequest
	GetCondition() *CreateAuthorizationResourceRequestCondition
	SetInstanceId(v string) *CreateAuthorizationResourceRequest
	GetInstanceId() *string
}

type CreateAuthorizationResourceRequest struct {
	// The ID of the resource entity associated with the authorization resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// carole_01kmek49aqxxxx
	AuthorizationResourceEntityId *string `json:"AuthorizationResourceEntityId,omitempty" xml:"AuthorizationResourceEntityId,omitempty"`
	// The type of the resource entity associated with the authorization resource. Valid values:
	//
	// - cloud_account_role: cloud role
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_account_role
	AuthorizationResourceEntityType *string `json:"AuthorizationResourceEntityType,omitempty" xml:"AuthorizationResourceEntityType,omitempty"`
	// The authorization rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate a parameter value, but you must make sure that the value is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see References [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The effective condition.
	Condition *CreateAuthorizationResourceRequestCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Struct"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateAuthorizationResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAuthorizationResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationResourceRequest) GetAuthorizationResourceEntityId() *string {
	return s.AuthorizationResourceEntityId
}

func (s *CreateAuthorizationResourceRequest) GetAuthorizationResourceEntityType() *string {
	return s.AuthorizationResourceEntityType
}

func (s *CreateAuthorizationResourceRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *CreateAuthorizationResourceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAuthorizationResourceRequest) GetCondition() *CreateAuthorizationResourceRequestCondition {
	return s.Condition
}

func (s *CreateAuthorizationResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAuthorizationResourceRequest) SetAuthorizationResourceEntityId(v string) *CreateAuthorizationResourceRequest {
	s.AuthorizationResourceEntityId = &v
	return s
}

func (s *CreateAuthorizationResourceRequest) SetAuthorizationResourceEntityType(v string) *CreateAuthorizationResourceRequest {
	s.AuthorizationResourceEntityType = &v
	return s
}

func (s *CreateAuthorizationResourceRequest) SetAuthorizationRuleId(v string) *CreateAuthorizationResourceRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *CreateAuthorizationResourceRequest) SetClientToken(v string) *CreateAuthorizationResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAuthorizationResourceRequest) SetCondition(v *CreateAuthorizationResourceRequestCondition) *CreateAuthorizationResourceRequest {
	s.Condition = v
	return s
}

func (s *CreateAuthorizationResourceRequest) SetInstanceId(v string) *CreateAuthorizationResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAuthorizationResourceRequest) Validate() error {
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAuthorizationResourceRequestCondition struct {
	// The effective condition when used as a credential.
	CredentialCondition *CreateAuthorizationResourceRequestConditionCredentialCondition `json:"CredentialCondition,omitempty" xml:"CredentialCondition,omitempty" type:"Struct"`
}

func (s CreateAuthorizationResourceRequestCondition) String() string {
	return dara.Prettify(s)
}

func (s CreateAuthorizationResourceRequestCondition) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationResourceRequestCondition) GetCredentialCondition() *CreateAuthorizationResourceRequestConditionCredentialCondition {
	return s.CredentialCondition
}

func (s *CreateAuthorizationResourceRequestCondition) SetCredentialCondition(v *CreateAuthorizationResourceRequestConditionCredentialCondition) *CreateAuthorizationResourceRequestCondition {
	s.CredentialCondition = v
	return s
}

func (s *CreateAuthorizationResourceRequestCondition) Validate() error {
	if s.CredentialCondition != nil {
		if err := s.CredentialCondition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAuthorizationResourceRequestConditionCredentialCondition struct {
	// Specifies whether same-name identity accounts are supported.
	AllowSameNameIdentity *bool `json:"AllowSameNameIdentity,omitempty" xml:"AllowSameNameIdentity,omitempty"`
}

func (s CreateAuthorizationResourceRequestConditionCredentialCondition) String() string {
	return dara.Prettify(s)
}

func (s CreateAuthorizationResourceRequestConditionCredentialCondition) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationResourceRequestConditionCredentialCondition) GetAllowSameNameIdentity() *bool {
	return s.AllowSameNameIdentity
}

func (s *CreateAuthorizationResourceRequestConditionCredentialCondition) SetAllowSameNameIdentity(v bool) *CreateAuthorizationResourceRequestConditionCredentialCondition {
	s.AllowSameNameIdentity = &v
	return s
}

func (s *CreateAuthorizationResourceRequestConditionCredentialCondition) Validate() error {
	return dara.Validate(s)
}
