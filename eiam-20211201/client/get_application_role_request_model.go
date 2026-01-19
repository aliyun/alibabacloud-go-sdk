// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetApplicationRoleRequest
	GetApplicationId() *string
	SetApplicationRoleId(v string) *GetApplicationRoleRequest
	GetApplicationRoleId() *string
	SetInstanceId(v string) *GetApplicationRoleRequest
	GetInstanceId() *string
}

type GetApplicationRoleRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 应用角色ID
	//
	// This parameter is required.
	//
	// example:
	//
	// approle_01kh2vuo8v9splv8maak1d22rxxxx
	ApplicationRoleId *string `json:"ApplicationRoleId,omitempty" xml:"ApplicationRoleId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetApplicationRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationRoleRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationRoleRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationRoleRequest) GetApplicationRoleId() *string {
	return s.ApplicationRoleId
}

func (s *GetApplicationRoleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApplicationRoleRequest) SetApplicationId(v string) *GetApplicationRoleRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationRoleRequest) SetApplicationRoleId(v string) *GetApplicationRoleRequest {
	s.ApplicationRoleId = &v
	return s
}

func (s *GetApplicationRoleRequest) SetInstanceId(v string) *GetApplicationRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationRoleRequest) Validate() error {
	return dara.Validate(s)
}
