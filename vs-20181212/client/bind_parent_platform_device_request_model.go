// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindParentPlatformDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *BindParentPlatformDeviceRequest
	GetDeviceId() *string
	SetOwnerId(v int64) *BindParentPlatformDeviceRequest
	GetOwnerId() *int64
	SetParentPlatformId(v string) *BindParentPlatformDeviceRequest
	GetParentPlatformId() *string
}

type BindParentPlatformDeviceRequest struct {
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

func (s BindParentPlatformDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s BindParentPlatformDeviceRequest) GoString() string {
	return s.String()
}

func (s *BindParentPlatformDeviceRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BindParentPlatformDeviceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindParentPlatformDeviceRequest) GetParentPlatformId() *string {
	return s.ParentPlatformId
}

func (s *BindParentPlatformDeviceRequest) SetDeviceId(v string) *BindParentPlatformDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *BindParentPlatformDeviceRequest) SetOwnerId(v int64) *BindParentPlatformDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *BindParentPlatformDeviceRequest) SetParentPlatformId(v string) *BindParentPlatformDeviceRequest {
	s.ParentPlatformId = &v
	return s
}

func (s *BindParentPlatformDeviceRequest) Validate() error {
	return dara.Validate(s)
}
