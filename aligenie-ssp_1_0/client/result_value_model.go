// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResultValue interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceOpenId(v string) *ResultValue
	GetDeviceOpenId() *string
	SetDeviceUnionIds(v []*ResultValueDeviceUnionIds) *ResultValue
	GetDeviceUnionIds() []*ResultValueDeviceUnionIds
	SetName(v string) *ResultValue
	GetName() *string
	SetFirmwareVersion(v string) *ResultValue
	GetFirmwareVersion() *string
	SetMac(v string) *ResultValue
	GetMac() *string
	SetSn(v string) *ResultValue
	GetSn() *string
}

type ResultValue struct {
	// example:
	//
	// A963*0158
	DeviceOpenId   *string                      `json:"DeviceOpenId,omitempty" xml:"DeviceOpenId,omitempty"`
	DeviceUnionIds []*ResultValueDeviceUnionIds `json:"DeviceUnionIds,omitempty" xml:"DeviceUnionIds,omitempty" type:"Repeated"`
	// example:
	//
	// 我的设备
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2.0.3
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" xml:"FirmwareVersion,omitempty"`
	// example:
	//
	// b4:xx:xx:xx:65:2b
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 1200xx048
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s ResultValue) String() string {
	return dara.Prettify(s)
}

func (s ResultValue) GoString() string {
	return s.String()
}

func (s *ResultValue) GetDeviceOpenId() *string {
	return s.DeviceOpenId
}

func (s *ResultValue) GetDeviceUnionIds() []*ResultValueDeviceUnionIds {
	return s.DeviceUnionIds
}

func (s *ResultValue) GetName() *string {
	return s.Name
}

func (s *ResultValue) GetFirmwareVersion() *string {
	return s.FirmwareVersion
}

func (s *ResultValue) GetMac() *string {
	return s.Mac
}

func (s *ResultValue) GetSn() *string {
	return s.Sn
}

func (s *ResultValue) SetDeviceOpenId(v string) *ResultValue {
	s.DeviceOpenId = &v
	return s
}

func (s *ResultValue) SetDeviceUnionIds(v []*ResultValueDeviceUnionIds) *ResultValue {
	s.DeviceUnionIds = v
	return s
}

func (s *ResultValue) SetName(v string) *ResultValue {
	s.Name = &v
	return s
}

func (s *ResultValue) SetFirmwareVersion(v string) *ResultValue {
	s.FirmwareVersion = &v
	return s
}

func (s *ResultValue) SetMac(v string) *ResultValue {
	s.Mac = &v
	return s
}

func (s *ResultValue) SetSn(v string) *ResultValue {
	s.Sn = &v
	return s
}

func (s *ResultValue) Validate() error {
	return dara.Validate(s)
}

type ResultValueDeviceUnionIds struct {
	// example:
	//
	// 1**2
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// 1553*B0C3
	DeviceUnionId *string `json:"DeviceUnionId,omitempty" xml:"DeviceUnionId,omitempty"`
}

func (s ResultValueDeviceUnionIds) String() string {
	return dara.Prettify(s)
}

func (s ResultValueDeviceUnionIds) GoString() string {
	return s.String()
}

func (s *ResultValueDeviceUnionIds) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ResultValueDeviceUnionIds) GetDeviceUnionId() *string {
	return s.DeviceUnionId
}

func (s *ResultValueDeviceUnionIds) SetOrganizationId(v string) *ResultValueDeviceUnionIds {
	s.OrganizationId = &v
	return s
}

func (s *ResultValueDeviceUnionIds) SetDeviceUnionId(v string) *ResultValueDeviceUnionIds {
	s.DeviceUnionId = &v
	return s
}

func (s *ResultValueDeviceUnionIds) Validate() error {
	return dara.Validate(s)
}
