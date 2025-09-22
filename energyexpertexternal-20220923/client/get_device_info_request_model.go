// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *GetDeviceInfoRequest
	GetDeviceId() *string
	SetDs(v string) *GetDeviceInfoRequest
	GetDs() *string
	SetFactoryId(v string) *GetDeviceInfoRequest
	GetFactoryId() *string
}

type GetDeviceInfoRequest struct {
	// The ID of the device.
	//
	// This parameter is required.
	//
	// example:
	//
	// pn_69873
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// The time string in the YYYY-mm-dd format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-07-26
	Ds *string `json:"ds,omitempty" xml:"ds,omitempty"`
	// The ID of the site.
	//
	// This parameter is required.
	//
	// example:
	//
	// pn_95
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
}

func (s GetDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetDeviceInfoRequest) GetDs() *string {
	return s.Ds
}

func (s *GetDeviceInfoRequest) GetFactoryId() *string {
	return s.FactoryId
}

func (s *GetDeviceInfoRequest) SetDeviceId(v string) *GetDeviceInfoRequest {
	s.DeviceId = &v
	return s
}

func (s *GetDeviceInfoRequest) SetDs(v string) *GetDeviceInfoRequest {
	s.Ds = &v
	return s
}

func (s *GetDeviceInfoRequest) SetFactoryId(v string) *GetDeviceInfoRequest {
	s.FactoryId = &v
	return s
}

func (s *GetDeviceInfoRequest) Validate() error {
	return dara.Validate(s)
}
