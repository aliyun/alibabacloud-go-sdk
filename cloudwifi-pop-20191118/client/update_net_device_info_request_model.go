// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *UpdateNetDeviceInfoRequest
	GetAppCode() *string
	SetAppName(v string) *UpdateNetDeviceInfoRequest
	GetAppName() *string
	SetDevices(v []*UpdateNetDeviceInfoRequestDevices) *UpdateNetDeviceInfoRequest
	GetDevices() []*UpdateNetDeviceInfoRequestDevices
	SetRequestId(v string) *UpdateNetDeviceInfoRequest
	GetRequestId() *string
}

type UpdateNetDeviceInfoRequest struct {
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Devices []*UpdateNetDeviceInfoRequestDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNetDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateNetDeviceInfoRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *UpdateNetDeviceInfoRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateNetDeviceInfoRequest) GetDevices() []*UpdateNetDeviceInfoRequestDevices {
	return s.Devices
}

func (s *UpdateNetDeviceInfoRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNetDeviceInfoRequest) SetAppCode(v string) *UpdateNetDeviceInfoRequest {
	s.AppCode = &v
	return s
}

func (s *UpdateNetDeviceInfoRequest) SetAppName(v string) *UpdateNetDeviceInfoRequest {
	s.AppName = &v
	return s
}

func (s *UpdateNetDeviceInfoRequest) SetDevices(v []*UpdateNetDeviceInfoRequestDevices) *UpdateNetDeviceInfoRequest {
	s.Devices = v
	return s
}

func (s *UpdateNetDeviceInfoRequest) SetRequestId(v string) *UpdateNetDeviceInfoRequest {
	s.RequestId = &v
	return s
}

func (s *UpdateNetDeviceInfoRequest) Validate() error {
	if s.Devices != nil {
		for _, item := range s.Devices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateNetDeviceInfoRequestDevices struct {
	// This parameter is required.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// This parameter is required.
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Idc         *string `json:"Idc,omitempty" xml:"Idc,omitempty"`
	LogicNetPod *string `json:"LogicNetPod,omitempty" xml:"LogicNetPod,omitempty"`
	// This parameter is required.
	Manufacturer *string `json:"Manufacturer,omitempty" xml:"Manufacturer,omitempty"`
	MgnIp        *string `json:"MgnIp,omitempty" xml:"MgnIp,omitempty"`
	Model        *string `json:"Model,omitempty" xml:"Model,omitempty"`
	NetPod       *string `json:"NetPod,omitempty" xml:"NetPod,omitempty"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// This parameter is required.
	ServiceTag *string `json:"ServiceTag,omitempty" xml:"ServiceTag,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateNetDeviceInfoRequestDevices) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetDeviceInfoRequestDevices) GoString() string {
	return s.String()
}

func (s *UpdateNetDeviceInfoRequestDevices) GetHostName() *string {
	return s.HostName
}

func (s *UpdateNetDeviceInfoRequestDevices) GetId() *int64 {
	return s.Id
}

func (s *UpdateNetDeviceInfoRequestDevices) GetIdc() *string {
	return s.Idc
}

func (s *UpdateNetDeviceInfoRequestDevices) GetLogicNetPod() *string {
	return s.LogicNetPod
}

func (s *UpdateNetDeviceInfoRequestDevices) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *UpdateNetDeviceInfoRequestDevices) GetMgnIp() *string {
	return s.MgnIp
}

func (s *UpdateNetDeviceInfoRequestDevices) GetModel() *string {
	return s.Model
}

func (s *UpdateNetDeviceInfoRequestDevices) GetNetPod() *string {
	return s.NetPod
}

func (s *UpdateNetDeviceInfoRequestDevices) GetPassword() *string {
	return s.Password
}

func (s *UpdateNetDeviceInfoRequestDevices) GetRole() *string {
	return s.Role
}

func (s *UpdateNetDeviceInfoRequestDevices) GetServiceTag() *string {
	return s.ServiceTag
}

func (s *UpdateNetDeviceInfoRequestDevices) GetUsername() *string {
	return s.Username
}

func (s *UpdateNetDeviceInfoRequestDevices) SetHostName(v string) *UpdateNetDeviceInfoRequestDevices {
	s.HostName = &v
	return s
}

func (s *UpdateNetDeviceInfoRequestDevices) SetId(v int64) *UpdateNetDeviceInfoRequestDevices {
	s.Id = &v
	return s
}

func (s *UpdateNetDeviceInfoRequestDevices) SetIdc(v string) *UpdateNetDeviceInfoRequestDevices {
	s.Idc = &v
	return s
}

func (s *UpdateNetDeviceInfoRequestDevices) SetLogicNetPod(v string) *UpdateNetDeviceInfoRequestDevices {
	s.LogicNetPod = &v
	return s
}

func (s *UpdateNetDeviceInfoRequestDevices) SetManufacturer(v string) *UpdateNetDeviceInfoRequestDevices {
	s.Manufacturer = &v
	return s
}

func (s *UpdateNetDeviceInfoRequestDevices) SetMgnIp(v string) *UpdateNetDeviceInfoRequestDevices {
	s.MgnIp = &v
	return s
}

func (s *UpdateNetDeviceInfoRequestDevices) SetModel(v string) *UpdateNetDeviceInfoRequestDevices {
	s.Model = &v
	return s
}

func (s *UpdateNetDeviceInfoRequestDevices) SetNetPod(v string) *UpdateNetDeviceInfoRequestDevices {
	s.NetPod = &v
	return s
}

func (s *UpdateNetDeviceInfoRequestDevices) SetPassword(v string) *UpdateNetDeviceInfoRequestDevices {
	s.Password = &v
	return s
}

func (s *UpdateNetDeviceInfoRequestDevices) SetRole(v string) *UpdateNetDeviceInfoRequestDevices {
	s.Role = &v
	return s
}

func (s *UpdateNetDeviceInfoRequestDevices) SetServiceTag(v string) *UpdateNetDeviceInfoRequestDevices {
	s.ServiceTag = &v
	return s
}

func (s *UpdateNetDeviceInfoRequestDevices) SetUsername(v string) *UpdateNetDeviceInfoRequestDevices {
	s.Username = &v
	return s
}

func (s *UpdateNetDeviceInfoRequestDevices) Validate() error {
	return dara.Validate(s)
}
