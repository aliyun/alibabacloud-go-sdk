// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCloudAccountRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DisableCloudAccountRoleRequest
	GetClientToken() *string
	SetCloudAccountId(v string) *DisableCloudAccountRoleRequest
	GetCloudAccountId() *string
	SetCloudAccountRoleId(v string) *DisableCloudAccountRoleRequest
	GetCloudAccountRoleId() *string
	SetInstanceId(v string) *DisableCloudAccountRoleRequest
	GetInstanceId() *string
}

type DisableCloudAccountRoleRequest struct {
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
	// 云账号角色ID
	//
	// This parameter is required.
	//
	// example:
	//
	// carole_01kmek49aqxxxx
	CloudAccountRoleId *string `json:"CloudAccountRoleId,omitempty" xml:"CloudAccountRoleId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableCloudAccountRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableCloudAccountRoleRequest) GoString() string {
	return s.String()
}

func (s *DisableCloudAccountRoleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableCloudAccountRoleRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *DisableCloudAccountRoleRequest) GetCloudAccountRoleId() *string {
	return s.CloudAccountRoleId
}

func (s *DisableCloudAccountRoleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableCloudAccountRoleRequest) SetClientToken(v string) *DisableCloudAccountRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableCloudAccountRoleRequest) SetCloudAccountId(v string) *DisableCloudAccountRoleRequest {
	s.CloudAccountId = &v
	return s
}

func (s *DisableCloudAccountRoleRequest) SetCloudAccountRoleId(v string) *DisableCloudAccountRoleRequest {
	s.CloudAccountRoleId = &v
	return s
}

func (s *DisableCloudAccountRoleRequest) SetInstanceId(v string) *DisableCloudAccountRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableCloudAccountRoleRequest) Validate() error {
	return dara.Validate(s)
}
