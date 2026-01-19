// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudAccountRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountId(v string) *DeleteCloudAccountRoleRequest
	GetCloudAccountId() *string
	SetCloudAccountRoleId(v string) *DeleteCloudAccountRoleRequest
	GetCloudAccountRoleId() *string
	SetInstanceId(v string) *DeleteCloudAccountRoleRequest
	GetInstanceId() *string
}

type DeleteCloudAccountRoleRequest struct {
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

func (s DeleteCloudAccountRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudAccountRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudAccountRoleRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *DeleteCloudAccountRoleRequest) GetCloudAccountRoleId() *string {
	return s.CloudAccountRoleId
}

func (s *DeleteCloudAccountRoleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCloudAccountRoleRequest) SetCloudAccountId(v string) *DeleteCloudAccountRoleRequest {
	s.CloudAccountId = &v
	return s
}

func (s *DeleteCloudAccountRoleRequest) SetCloudAccountRoleId(v string) *DeleteCloudAccountRoleRequest {
	s.CloudAccountRoleId = &v
	return s
}

func (s *DeleteCloudAccountRoleRequest) SetInstanceId(v string) *DeleteCloudAccountRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCloudAccountRoleRequest) Validate() error {
	return dara.Validate(s)
}
