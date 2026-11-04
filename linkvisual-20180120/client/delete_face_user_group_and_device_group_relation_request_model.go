// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceUserGroupAndDeviceGroupRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetControlId(v string) *DeleteFaceUserGroupAndDeviceGroupRelationRequest
	GetControlId() *string
	SetIsolationId(v string) *DeleteFaceUserGroupAndDeviceGroupRelationRequest
	GetIsolationId() *string
}

type DeleteFaceUserGroupAndDeviceGroupRelationRequest struct {
	// The control ID of the synchronization relationship between the user group and the device group. Set this parameter to the value of **ControlId*	- that you obtained and saved locally when you called the [AddFaceUserGroupAndDeviceGroupRelation](https://help.aliyun.com/document_detail/288342.html) operation to add the face synchronization relationship between the user group and the device group.
	//
	// This parameter is required.
	//
	// example:
	//
	// qcb2yea4ha4d****
	ControlId *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
	// The business isolation ID, which is used to isolate data between different business applications under the same Alibaba Cloud account. Set this parameter to the same value of **IsolationId*	- that you specified when you called the [AddFaceUserGroupAndDeviceGroupRelation](https://help.aliyun.com/document_detail/288342.html) operation to add the face synchronization relationship between the user group and the device group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZheJiangHZ
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
}

func (s DeleteFaceUserGroupAndDeviceGroupRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceUserGroupAndDeviceGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationRequest) GetControlId() *string {
	return s.ControlId
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationRequest) SetControlId(v string) *DeleteFaceUserGroupAndDeviceGroupRelationRequest {
	s.ControlId = &v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationRequest) SetIsolationId(v string) *DeleteFaceUserGroupAndDeviceGroupRelationRequest {
	s.IsolationId = &v
	return s
}

func (s *DeleteFaceUserGroupAndDeviceGroupRelationRequest) Validate() error {
	return dara.Validate(s)
}
