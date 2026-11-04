// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceUserGroupAndDeviceGroupRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceGroupId(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest
	GetDeviceGroupId() *string
	SetIotInstanceId(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest
	GetIotInstanceId() *string
	SetIsolationId(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest
	GetIsolationId() *string
	SetRelation(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest
	GetRelation() *string
	SetUserGroupId(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest
	GetUserGroupId() *string
}

type AddFaceUserGroupAndDeviceGroupRelationRequest struct {
	// The device group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// kvjyhurc
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// TestInstanceId
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The isolation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestIsolationId
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The face synchronization relationship.
	//
	// This parameter is required.
	//
	// example:
	//
	// SYNC_INFO_PICTURE
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// The user group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// avjyxurd
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s AddFaceUserGroupAndDeviceGroupRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFaceUserGroupAndDeviceGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) GetRelation() *string {
	return s.Relation
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetDeviceGroupId(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetIotInstanceId(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.IotInstanceId = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetIsolationId(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetRelation(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.Relation = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) SetUserGroupId(v string) *AddFaceUserGroupAndDeviceGroupRelationRequest {
	s.UserGroupId = &v
	return s
}

func (s *AddFaceUserGroupAndDeviceGroupRelationRequest) Validate() error {
	return dara.Validate(s)
}
