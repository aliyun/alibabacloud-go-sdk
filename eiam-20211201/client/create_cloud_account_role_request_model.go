// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudAccountRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateCloudAccountRoleRequest
	GetClientToken() *string
	SetCloudAccountId(v string) *CreateCloudAccountRoleRequest
	GetCloudAccountId() *string
	SetCloudAccountRoleName(v string) *CreateCloudAccountRoleRequest
	GetCloudAccountRoleName() *string
	SetCloudAccountRoleType(v string) *CreateCloudAccountRoleRequest
	GetCloudAccountRoleType() *string
	SetDescription(v string) *CreateCloudAccountRoleRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateCloudAccountRoleRequest
	GetInstanceId() *string
}

type CreateCloudAccountRoleRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate a parameter value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see References: [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cloud account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The cloud role name.
	//
	// This parameter is required.
	//
	// example:
	//
	// role-test
	CloudAccountRoleName *string `json:"CloudAccountRoleName,omitempty" xml:"CloudAccountRoleName,omitempty"`
	// The cloud role type. The specific format depends on the cloud account type. Valid values:
	//
	// - role: applicable to Alibaba Cloud accounts.
	//
	// example:
	//
	// role
	CloudAccountRoleType *string `json:"CloudAccountRoleType,omitempty" xml:"CloudAccountRoleType,omitempty"`
	// The cloud role description.
	//
	// example:
	//
	// cloud_account_role_description
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

func (s CreateCloudAccountRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudAccountRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudAccountRoleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCloudAccountRoleRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *CreateCloudAccountRoleRequest) GetCloudAccountRoleName() *string {
	return s.CloudAccountRoleName
}

func (s *CreateCloudAccountRoleRequest) GetCloudAccountRoleType() *string {
	return s.CloudAccountRoleType
}

func (s *CreateCloudAccountRoleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCloudAccountRoleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCloudAccountRoleRequest) SetClientToken(v string) *CreateCloudAccountRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCloudAccountRoleRequest) SetCloudAccountId(v string) *CreateCloudAccountRoleRequest {
	s.CloudAccountId = &v
	return s
}

func (s *CreateCloudAccountRoleRequest) SetCloudAccountRoleName(v string) *CreateCloudAccountRoleRequest {
	s.CloudAccountRoleName = &v
	return s
}

func (s *CreateCloudAccountRoleRequest) SetCloudAccountRoleType(v string) *CreateCloudAccountRoleRequest {
	s.CloudAccountRoleType = &v
	return s
}

func (s *CreateCloudAccountRoleRequest) SetDescription(v string) *CreateCloudAccountRoleRequest {
	s.Description = &v
	return s
}

func (s *CreateCloudAccountRoleRequest) SetInstanceId(v string) *CreateCloudAccountRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCloudAccountRoleRequest) Validate() error {
	return dara.Validate(s)
}
