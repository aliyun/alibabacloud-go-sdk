// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceDeviceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceGroupId(v string) *DeleteFaceDeviceGroupRequest
	GetDeviceGroupId() *string
	SetIsolationId(v string) *DeleteFaceDeviceGroupRequest
	GetIsolationId() *string
}

type DeleteFaceDeviceGroupRequest struct {
	// The device group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// qaf2tihq
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	// The isolation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestIsolationId
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
}

func (s DeleteFaceDeviceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceDeviceGroupRequest) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *DeleteFaceDeviceGroupRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *DeleteFaceDeviceGroupRequest) SetDeviceGroupId(v string) *DeleteFaceDeviceGroupRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *DeleteFaceDeviceGroupRequest) SetIsolationId(v string) *DeleteFaceDeviceGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *DeleteFaceDeviceGroupRequest) Validate() error {
	return dara.Validate(s)
}
