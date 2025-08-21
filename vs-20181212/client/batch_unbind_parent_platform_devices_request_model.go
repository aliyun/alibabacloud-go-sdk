// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUnbindParentPlatformDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *BatchUnbindParentPlatformDevicesRequest
	GetDeviceId() *string
	SetOwnerId(v int64) *BatchUnbindParentPlatformDevicesRequest
	GetOwnerId() *int64
	SetParentPlatformId(v string) *BatchUnbindParentPlatformDevicesRequest
	GetParentPlatformId() *string
}

type BatchUnbindParentPlatformDevicesRequest struct {
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

func (s BatchUnbindParentPlatformDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindParentPlatformDevicesRequest) GoString() string {
	return s.String()
}

func (s *BatchUnbindParentPlatformDevicesRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BatchUnbindParentPlatformDevicesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchUnbindParentPlatformDevicesRequest) GetParentPlatformId() *string {
	return s.ParentPlatformId
}

func (s *BatchUnbindParentPlatformDevicesRequest) SetDeviceId(v string) *BatchUnbindParentPlatformDevicesRequest {
	s.DeviceId = &v
	return s
}

func (s *BatchUnbindParentPlatformDevicesRequest) SetOwnerId(v int64) *BatchUnbindParentPlatformDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchUnbindParentPlatformDevicesRequest) SetParentPlatformId(v string) *BatchUnbindParentPlatformDevicesRequest {
	s.ParentPlatformId = &v
	return s
}

func (s *BatchUnbindParentPlatformDevicesRequest) Validate() error {
	return dara.Validate(s)
}
