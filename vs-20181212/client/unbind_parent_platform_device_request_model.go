// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindParentPlatformDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *UnbindParentPlatformDeviceRequest
	GetDeviceId() *string
	SetOwnerId(v int64) *UnbindParentPlatformDeviceRequest
	GetOwnerId() *int64
	SetParentPlatformId(v string) *UnbindParentPlatformDeviceRequest
	GetParentPlatformId() *string
}

type UnbindParentPlatformDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 361*****212-cn-qingdao
	ParentPlatformId *string `json:"ParentPlatformId,omitempty" xml:"ParentPlatformId,omitempty"`
}

func (s UnbindParentPlatformDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindParentPlatformDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnbindParentPlatformDeviceRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *UnbindParentPlatformDeviceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnbindParentPlatformDeviceRequest) GetParentPlatformId() *string {
	return s.ParentPlatformId
}

func (s *UnbindParentPlatformDeviceRequest) SetDeviceId(v string) *UnbindParentPlatformDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *UnbindParentPlatformDeviceRequest) SetOwnerId(v int64) *UnbindParentPlatformDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindParentPlatformDeviceRequest) SetParentPlatformId(v string) *UnbindParentPlatformDeviceRequest {
	s.ParentPlatformId = &v
	return s
}

func (s *UnbindParentPlatformDeviceRequest) Validate() error {
	return dara.Validate(s)
}
