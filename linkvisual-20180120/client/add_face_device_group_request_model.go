// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceDeviceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceGroupName(v string) *AddFaceDeviceGroupRequest
	GetDeviceGroupName() *string
	SetIsolationId(v string) *AddFaceDeviceGroupRequest
	GetIsolationId() *string
}

type AddFaceDeviceGroupRequest struct {
	// The name of the device group. You can customize the value. The name can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestDeviceGroup
	DeviceGroupName *string `json:"DeviceGroupName,omitempty" xml:"DeviceGroupName,omitempty"`
	// The business isolation ID. It is used to isolate data between different business applications under the same Alibaba Cloud account. It supports uppercase and lowercase English letters and digits.
	//
	// For example, if you have face recognition businesses in Region A and Region B, you can set different business isolation IDs to distinguish the face recognition businesses in Region A and Region B.
	//
	// 	Notice:
	//
	// Save the business isolation ID properly for use when you call other API operations in this business isolation zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZheJiangHZ
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
}

func (s AddFaceDeviceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFaceDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *AddFaceDeviceGroupRequest) GetDeviceGroupName() *string {
	return s.DeviceGroupName
}

func (s *AddFaceDeviceGroupRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *AddFaceDeviceGroupRequest) SetDeviceGroupName(v string) *AddFaceDeviceGroupRequest {
	s.DeviceGroupName = &v
	return s
}

func (s *AddFaceDeviceGroupRequest) SetIsolationId(v string) *AddFaceDeviceGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *AddFaceDeviceGroupRequest) Validate() error {
	return dara.Validate(s)
}
