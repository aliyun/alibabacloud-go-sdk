// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationRoleDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateApplicationRoleDescriptionRequest
	GetApplicationId() *string
	SetApplicationRoleId(v string) *UpdateApplicationRoleDescriptionRequest
	GetApplicationRoleId() *string
	SetDescription(v string) *UpdateApplicationRoleDescriptionRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateApplicationRoleDescriptionRequest
	GetInstanceId() *string
}

type UpdateApplicationRoleDescriptionRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 应用角色的唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// approle_01kh2vuo8v9splv8maak1d22rxxxx
	ApplicationRoleId *string `json:"ApplicationRoleId,omitempty" xml:"ApplicationRoleId,omitempty"`
	// 应用角色的唯一标识
	//
	// example:
	//
	// Admin Role Description
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

func (s UpdateApplicationRoleDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationRoleDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRoleDescriptionRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationRoleDescriptionRequest) GetApplicationRoleId() *string {
	return s.ApplicationRoleId
}

func (s *UpdateApplicationRoleDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationRoleDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateApplicationRoleDescriptionRequest) SetApplicationId(v string) *UpdateApplicationRoleDescriptionRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationRoleDescriptionRequest) SetApplicationRoleId(v string) *UpdateApplicationRoleDescriptionRequest {
	s.ApplicationRoleId = &v
	return s
}

func (s *UpdateApplicationRoleDescriptionRequest) SetDescription(v string) *UpdateApplicationRoleDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateApplicationRoleDescriptionRequest) SetInstanceId(v string) *UpdateApplicationRoleDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateApplicationRoleDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
