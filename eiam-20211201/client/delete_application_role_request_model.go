// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DeleteApplicationRoleRequest
	GetApplicationId() *string
	SetApplicationRoleId(v string) *DeleteApplicationRoleRequest
	GetApplicationRoleId() *string
	SetInstanceId(v string) *DeleteApplicationRoleRequest
	GetInstanceId() *string
}

type DeleteApplicationRoleRequest struct {
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

func (s DeleteApplicationRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRoleRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DeleteApplicationRoleRequest) GetApplicationRoleId() *string {
	return s.ApplicationRoleId
}

func (s *DeleteApplicationRoleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteApplicationRoleRequest) SetApplicationId(v string) *DeleteApplicationRoleRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeleteApplicationRoleRequest) SetApplicationRoleId(v string) *DeleteApplicationRoleRequest {
	s.ApplicationRoleId = &v
	return s
}

func (s *DeleteApplicationRoleRequest) SetInstanceId(v string) *DeleteApplicationRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteApplicationRoleRequest) Validate() error {
	return dara.Validate(s)
}
