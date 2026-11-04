// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFaceUserGroupAndDeviceGroupRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetControlId(v string) *UpdateFaceUserGroupAndDeviceGroupRelationRequest
	GetControlId() *string
	SetIsolationId(v string) *UpdateFaceUserGroupAndDeviceGroupRelationRequest
	GetIsolationId() *string
	SetRelation(v string) *UpdateFaceUserGroupAndDeviceGroupRelationRequest
	GetRelation() *string
}

type UpdateFaceUserGroupAndDeviceGroupRelationRequest struct {
	// The control ID of the synchronization relationship between the user group and the device group. Pass the **ControlId*	- value that you obtained and saved locally when you called the [AddFaceUserGroupAndDeviceGroupRelation](https://help.aliyun.com/document_detail/288342.html) operation to add the face synchronization relationship between the user group and the device group.
	//
	// This parameter is required.
	//
	// example:
	//
	// qcb2yea4ha4d****
	ControlId *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
	// The business isolation ID, which is used to isolate data between different business applications under the same Alibaba Cloud account. The value must be the same as the **IsolationId*	- value that you specified when you called the [AddFaceUserGroupAndDeviceGroupRelation](https://help.aliyun.com/document_detail/288342.html) operation to add the face synchronization relationship between the user group and the device group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZheJiangHZ
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The face synchronization relationship. Currently, only **SYNC_INFO_PICTURE*	- is supported, which indicates that user photos and user information are synchronized.
	//
	// This parameter is required.
	//
	// example:
	//
	// SYNC_INFO_PICTURE
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFaceUserGroupAndDeviceGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationRequest) GetControlId() *string {
	return s.ControlId
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationRequest) GetRelation() *string {
	return s.Relation
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationRequest) SetControlId(v string) *UpdateFaceUserGroupAndDeviceGroupRelationRequest {
	s.ControlId = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationRequest) SetIsolationId(v string) *UpdateFaceUserGroupAndDeviceGroupRelationRequest {
	s.IsolationId = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationRequest) SetRelation(v string) *UpdateFaceUserGroupAndDeviceGroupRelationRequest {
	s.Relation = &v
	return s
}

func (s *UpdateFaceUserGroupAndDeviceGroupRelationRequest) Validate() error {
	return dara.Validate(s)
}
