// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudAccountRoleDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateCloudAccountRoleDescriptionRequest
	GetClientToken() *string
	SetCloudAccountId(v string) *UpdateCloudAccountRoleDescriptionRequest
	GetCloudAccountId() *string
	SetCloudAccountRoleId(v string) *UpdateCloudAccountRoleDescriptionRequest
	GetCloudAccountRoleId() *string
	SetDescription(v string) *UpdateCloudAccountRoleDescriptionRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateCloudAccountRoleDescriptionRequest
	GetInstanceId() *string
}

type UpdateCloudAccountRoleDescriptionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// 云账号角色ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// carole_01kmek49aqxxxx
	CloudAccountRoleId *string `json:"CloudAccountRoleId,omitempty" xml:"CloudAccountRoleId,omitempty"`
	// 描述
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_account_role_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateCloudAccountRoleDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudAccountRoleDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudAccountRoleDescriptionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudAccountRoleDescriptionRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *UpdateCloudAccountRoleDescriptionRequest) GetCloudAccountRoleId() *string {
	return s.CloudAccountRoleId
}

func (s *UpdateCloudAccountRoleDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateCloudAccountRoleDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCloudAccountRoleDescriptionRequest) SetClientToken(v string) *UpdateCloudAccountRoleDescriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudAccountRoleDescriptionRequest) SetCloudAccountId(v string) *UpdateCloudAccountRoleDescriptionRequest {
	s.CloudAccountId = &v
	return s
}

func (s *UpdateCloudAccountRoleDescriptionRequest) SetCloudAccountRoleId(v string) *UpdateCloudAccountRoleDescriptionRequest {
	s.CloudAccountRoleId = &v
	return s
}

func (s *UpdateCloudAccountRoleDescriptionRequest) SetDescription(v string) *UpdateCloudAccountRoleDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateCloudAccountRoleDescriptionRequest) SetInstanceId(v string) *UpdateCloudAccountRoleDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCloudAccountRoleDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
