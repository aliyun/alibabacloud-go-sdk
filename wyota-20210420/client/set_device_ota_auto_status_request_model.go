// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeviceOtaAutoStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUpdate(v int32) *SetDeviceOtaAutoStatusRequest
	GetAutoUpdate() *int32
	SetAutoUpdateTimeSchedule(v string) *SetDeviceOtaAutoStatusRequest
	GetAutoUpdateTimeSchedule() *string
	SetClientType(v int32) *SetDeviceOtaAutoStatusRequest
	GetClientType() *int32
	SetForceUpgrade(v int32) *SetDeviceOtaAutoStatusRequest
	GetForceUpgrade() *int32
	SetStatus(v string) *SetDeviceOtaAutoStatusRequest
	GetStatus() *string
}

type SetDeviceOtaAutoStatusRequest struct {
	AutoUpdate             *int32  `json:"AutoUpdate,omitempty" xml:"AutoUpdate,omitempty"`
	AutoUpdateTimeSchedule *string `json:"AutoUpdateTimeSchedule,omitempty" xml:"AutoUpdateTimeSchedule,omitempty"`
	ClientType             *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ForceUpgrade           *int32  `json:"ForceUpgrade,omitempty" xml:"ForceUpgrade,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetDeviceOtaAutoStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDeviceOtaAutoStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDeviceOtaAutoStatusRequest) GetAutoUpdate() *int32 {
	return s.AutoUpdate
}

func (s *SetDeviceOtaAutoStatusRequest) GetAutoUpdateTimeSchedule() *string {
	return s.AutoUpdateTimeSchedule
}

func (s *SetDeviceOtaAutoStatusRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *SetDeviceOtaAutoStatusRequest) GetForceUpgrade() *int32 {
	return s.ForceUpgrade
}

func (s *SetDeviceOtaAutoStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *SetDeviceOtaAutoStatusRequest) SetAutoUpdate(v int32) *SetDeviceOtaAutoStatusRequest {
	s.AutoUpdate = &v
	return s
}

func (s *SetDeviceOtaAutoStatusRequest) SetAutoUpdateTimeSchedule(v string) *SetDeviceOtaAutoStatusRequest {
	s.AutoUpdateTimeSchedule = &v
	return s
}

func (s *SetDeviceOtaAutoStatusRequest) SetClientType(v int32) *SetDeviceOtaAutoStatusRequest {
	s.ClientType = &v
	return s
}

func (s *SetDeviceOtaAutoStatusRequest) SetForceUpgrade(v int32) *SetDeviceOtaAutoStatusRequest {
	s.ForceUpgrade = &v
	return s
}

func (s *SetDeviceOtaAutoStatusRequest) SetStatus(v string) *SetDeviceOtaAutoStatusRequest {
	s.Status = &v
	return s
}

func (s *SetDeviceOtaAutoStatusRequest) Validate() error {
	return dara.Validate(s)
}
