// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceUserGroupAndDeviceGroupRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetControlId(v string) *QueryFaceUserGroupAndDeviceGroupRelationRequest
	GetControlId() *string
	SetIsolationId(v string) *QueryFaceUserGroupAndDeviceGroupRelationRequest
	GetIsolationId() *string
}

type QueryFaceUserGroupAndDeviceGroupRelationRequest struct {
	// The control ID for the face synchronization relationship between user groups and device groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// ni1iyxqry0j03plv
	ControlId *string `json:"ControlId,omitempty" xml:"ControlId,omitempty"`
	// The isolation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestIsolationId
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
}

func (s QueryFaceUserGroupAndDeviceGroupRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserGroupAndDeviceGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationRequest) GetControlId() *string {
	return s.ControlId
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationRequest) SetControlId(v string) *QueryFaceUserGroupAndDeviceGroupRelationRequest {
	s.ControlId = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationRequest) SetIsolationId(v string) *QueryFaceUserGroupAndDeviceGroupRelationRequest {
	s.IsolationId = &v
	return s
}

func (s *QueryFaceUserGroupAndDeviceGroupRelationRequest) Validate() error {
	return dara.Validate(s)
}
