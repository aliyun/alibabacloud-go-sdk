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
	SetInstanceId(v string) *CreateAuthorizationResourceRequest
	GetInstanceId() *string
}

type CreateAuthorizationResourceRequest struct {
	// 授权资源关联的资源标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// carole_01kmek49aqxxxx
	AuthorizationResourceEntityId *string `json:"AuthorizationResourceEntityId,omitempty" xml:"AuthorizationResourceEntityId,omitempty"`
	// 授权资源的资源类型。枚举取值:asset(资产)、credential(凭据)、cloudAccountRole(云账号角色)。
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_account_role
	AuthorizationResourceEntityType *string `json:"AuthorizationResourceEntityType,omitempty" xml:"AuthorizationResourceEntityType,omitempty"`
	// 授权规则标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// IDaaS EIAM实例的ID。
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

func (s *CreateAuthorizationResourceRequest) SetInstanceId(v string) *CreateAuthorizationResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAuthorizationResourceRequest) Validate() error {
	return dara.Validate(s)
}
