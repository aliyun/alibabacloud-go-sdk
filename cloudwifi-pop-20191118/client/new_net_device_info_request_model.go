// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNewNetDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *NewNetDeviceInfoRequest
	GetAppCode() *string
	SetAppName(v string) *NewNetDeviceInfoRequest
	GetAppName() *string
	SetDevices(v []*NewNetDeviceInfoRequestDevices) *NewNetDeviceInfoRequest
	GetDevices() []*NewNetDeviceInfoRequestDevices
	SetRequestId(v string) *NewNetDeviceInfoRequest
	GetRequestId() *string
}

type NewNetDeviceInfoRequest struct {
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Devices []*NewNetDeviceInfoRequestDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// This parameter is required.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s NewNetDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s NewNetDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *NewNetDeviceInfoRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *NewNetDeviceInfoRequest) GetAppName() *string {
	return s.AppName
}

func (s *NewNetDeviceInfoRequest) GetDevices() []*NewNetDeviceInfoRequestDevices {
	return s.Devices
}

func (s *NewNetDeviceInfoRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *NewNetDeviceInfoRequest) SetAppCode(v string) *NewNetDeviceInfoRequest {
	s.AppCode = &v
	return s
}

func (s *NewNetDeviceInfoRequest) SetAppName(v string) *NewNetDeviceInfoRequest {
	s.AppName = &v
	return s
}

func (s *NewNetDeviceInfoRequest) SetDevices(v []*NewNetDeviceInfoRequestDevices) *NewNetDeviceInfoRequest {
	s.Devices = v
	return s
}

func (s *NewNetDeviceInfoRequest) SetRequestId(v string) *NewNetDeviceInfoRequest {
	s.RequestId = &v
	return s
}

func (s *NewNetDeviceInfoRequest) Validate() error {
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

type NewNetDeviceInfoRequestDevices struct {
	// This parameter is required.
	HostName     *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Idc          *string `json:"Idc,omitempty" xml:"Idc,omitempty"`
	LogicNetPod  *string `json:"LogicNetPod,omitempty" xml:"LogicNetPod,omitempty"`
	Manufacturer *string `json:"Manufacturer,omitempty" xml:"Manufacturer,omitempty"`
	// This parameter is required.
	MgnIp    *string `json:"MgnIp,omitempty" xml:"MgnIp,omitempty"`
	Model    *string `json:"Model,omitempty" xml:"Model,omitempty"`
	NetPod   *string `json:"NetPod,omitempty" xml:"NetPod,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Role     *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// This parameter is required.
	ServiceTag *string `json:"ServiceTag,omitempty" xml:"ServiceTag,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s NewNetDeviceInfoRequestDevices) String() string {
	return dara.Prettify(s)
}

func (s NewNetDeviceInfoRequestDevices) GoString() string {
	return s.String()
}

func (s *NewNetDeviceInfoRequestDevices) GetHostName() *string {
	return s.HostName
}

func (s *NewNetDeviceInfoRequestDevices) GetIdc() *string {
	return s.Idc
}

func (s *NewNetDeviceInfoRequestDevices) GetLogicNetPod() *string {
	return s.LogicNetPod
}

func (s *NewNetDeviceInfoRequestDevices) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *NewNetDeviceInfoRequestDevices) GetMgnIp() *string {
	return s.MgnIp
}

func (s *NewNetDeviceInfoRequestDevices) GetModel() *string {
	return s.Model
}

func (s *NewNetDeviceInfoRequestDevices) GetNetPod() *string {
	return s.NetPod
}

func (s *NewNetDeviceInfoRequestDevices) GetPassword() *string {
	return s.Password
}

func (s *NewNetDeviceInfoRequestDevices) GetRole() *string {
	return s.Role
}

func (s *NewNetDeviceInfoRequestDevices) GetServiceTag() *string {
	return s.ServiceTag
}

func (s *NewNetDeviceInfoRequestDevices) GetUsername() *string {
	return s.Username
}

func (s *NewNetDeviceInfoRequestDevices) SetHostName(v string) *NewNetDeviceInfoRequestDevices {
	s.HostName = &v
	return s
}

func (s *NewNetDeviceInfoRequestDevices) SetIdc(v string) *NewNetDeviceInfoRequestDevices {
	s.Idc = &v
	return s
}

func (s *NewNetDeviceInfoRequestDevices) SetLogicNetPod(v string) *NewNetDeviceInfoRequestDevices {
	s.LogicNetPod = &v
	return s
}

func (s *NewNetDeviceInfoRequestDevices) SetManufacturer(v string) *NewNetDeviceInfoRequestDevices {
	s.Manufacturer = &v
	return s
}

func (s *NewNetDeviceInfoRequestDevices) SetMgnIp(v string) *NewNetDeviceInfoRequestDevices {
	s.MgnIp = &v
	return s
}

func (s *NewNetDeviceInfoRequestDevices) SetModel(v string) *NewNetDeviceInfoRequestDevices {
	s.Model = &v
	return s
}

func (s *NewNetDeviceInfoRequestDevices) SetNetPod(v string) *NewNetDeviceInfoRequestDevices {
	s.NetPod = &v
	return s
}

func (s *NewNetDeviceInfoRequestDevices) SetPassword(v string) *NewNetDeviceInfoRequestDevices {
	s.Password = &v
	return s
}

func (s *NewNetDeviceInfoRequestDevices) SetRole(v string) *NewNetDeviceInfoRequestDevices {
	s.Role = &v
	return s
}

func (s *NewNetDeviceInfoRequestDevices) SetServiceTag(v string) *NewNetDeviceInfoRequestDevices {
	s.ServiceTag = &v
	return s
}

func (s *NewNetDeviceInfoRequestDevices) SetUsername(v string) *NewNetDeviceInfoRequestDevices {
	s.Username = &v
	return s
}

func (s *NewNetDeviceInfoRequestDevices) Validate() error {
	return dara.Validate(s)
}
