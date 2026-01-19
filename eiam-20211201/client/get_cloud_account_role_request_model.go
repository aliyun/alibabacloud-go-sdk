// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAccountRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountId(v string) *GetCloudAccountRoleRequest
	GetCloudAccountId() *string
	SetCloudAccountRoleId(v string) *GetCloudAccountRoleRequest
	GetCloudAccountRoleId() *string
	SetInstanceId(v string) *GetCloudAccountRoleRequest
	GetInstanceId() *string
}

type GetCloudAccountRoleRequest struct {
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
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCloudAccountRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountRoleRequest) GoString() string {
	return s.String()
}

func (s *GetCloudAccountRoleRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *GetCloudAccountRoleRequest) GetCloudAccountRoleId() *string {
	return s.CloudAccountRoleId
}

func (s *GetCloudAccountRoleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCloudAccountRoleRequest) SetCloudAccountId(v string) *GetCloudAccountRoleRequest {
	s.CloudAccountId = &v
	return s
}

func (s *GetCloudAccountRoleRequest) SetCloudAccountRoleId(v string) *GetCloudAccountRoleRequest {
	s.CloudAccountRoleId = &v
	return s
}

func (s *GetCloudAccountRoleRequest) SetInstanceId(v string) *GetCloudAccountRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCloudAccountRoleRequest) Validate() error {
	return dara.Validate(s)
}
