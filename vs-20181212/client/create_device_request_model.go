// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmMethod(v string) *CreateDeviceRequest
	GetAlarmMethod() *string
	SetAutoDirectory(v bool) *CreateDeviceRequest
	GetAutoDirectory() *bool
	SetAutoPos(v bool) *CreateDeviceRequest
	GetAutoPos() *bool
	SetAutoStart(v bool) *CreateDeviceRequest
	GetAutoStart() *bool
	SetDescription(v string) *CreateDeviceRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateDeviceRequest
	GetDirectoryId() *string
	SetDsn(v string) *CreateDeviceRequest
	GetDsn() *string
	SetGbId(v string) *CreateDeviceRequest
	GetGbId() *string
	SetGroupId(v string) *CreateDeviceRequest
	GetGroupId() *string
	SetIp(v string) *CreateDeviceRequest
	GetIp() *string
	SetLatitude(v string) *CreateDeviceRequest
	GetLatitude() *string
	SetLongitude(v string) *CreateDeviceRequest
	GetLongitude() *string
	SetName(v string) *CreateDeviceRequest
	GetName() *string
	SetOwnerId(v int64) *CreateDeviceRequest
	GetOwnerId() *int64
	SetParams(v string) *CreateDeviceRequest
	GetParams() *string
	SetParentId(v string) *CreateDeviceRequest
	GetParentId() *string
	SetPassword(v string) *CreateDeviceRequest
	GetPassword() *string
	SetPort(v int64) *CreateDeviceRequest
	GetPort() *int64
	SetPosInterval(v int64) *CreateDeviceRequest
	GetPosInterval() *int64
	SetType(v string) *CreateDeviceRequest
	GetType() *string
	SetUrl(v string) *CreateDeviceRequest
	GetUrl() *string
	SetUsername(v string) *CreateDeviceRequest
	GetUsername() *string
	SetVendor(v string) *CreateDeviceRequest
	GetVendor() *string
}

type CreateDeviceRequest struct {
	// GB-compliant alarm method to subscribe to. Valid values:
	//
	// - 0 (all)
	//
	// - 5 (video alarm)
	//
	// - 7 (other alarms)
	//
	// > 	- An empty value means no subscription.
	//
	// >
	//
	// > 	- You can specify multiple values, separated by commas (,).
	//
	// example:
	//
	// 0
	AlarmMethod   *string `json:"AlarmMethod,omitempty" xml:"AlarmMethod,omitempty"`
	AutoDirectory *bool   `json:"AutoDirectory,omitempty" xml:"AutoDirectory,omitempty"`
	// Whether to enable location subscription for the device. Default value: false.
	//
	// example:
	//
	// false
	AutoPos *bool `json:"AutoPos,omitempty" xml:"AutoPos,omitempty"`
	// Whether to automatically start the stream. Default value: false.
	//
	// example:
	//
	// false
	AutoStart *bool `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	// Device description.
	//
	// example:
	//
	// xxx路口摄像头
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ID of the folder that contains the device.
	//
	// example:
	//
	// 399*****488-cn-qingdao
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// Device serial number.
	//
	// example:
	//
	// 7D0*****4C0
	Dsn *string `json:"Dsn,omitempty" xml:"Dsn,omitempty"`
	// GB-compliant device ID.
	//
	// > This parameter applies only to GB-compliant protocols.
	//
	// example:
	//
	// 31000000****00000002
	GbId *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	// ID of the space that contains the device.
	//
	// This parameter is required.
	//
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Device IP address.
	//
	// example:
	//
	// 10.10.10.10
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The dimension of the device.
	//
	// example:
	//
	// 119.20
	Latitude *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// Device longitude.
	//
	// example:
	//
	// 45.00
	Longitude *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	// Device name.
	//
	// example:
	//
	// xxx路口摄像头
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Additional device parameters, formatted as a JSON string.
	//
	// example:
	//
	// {}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// ID of the parent device. For example, the ID of the platform that hosts the camera.
	//
	// example:
	//
	// 399*****774-cn-qingdao
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// Device password.
	//
	// example:
	//
	// admin
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Device port.
	//
	// example:
	//
	// 8080
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// Location subscription interval, in seconds.
	//
	// example:
	//
	// 300
	PosInterval *int64 `json:"PosInterval,omitempty" xml:"PosInterval,omitempty"`
	// Device type. Valid values:
	//
	// - ipc (camera)
	//
	// - platform (platform)
	//
	// - ied (intelligent device)
	//
	// This parameter is required.
	//
	// example:
	//
	// ipc
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Stream URL on the device.
	//
	// example:
	//
	// rtmp://xxx/xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// Device username.
	//
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// Device vendor.
	//
	// example:
	//
	// 公司A
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s CreateDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeviceRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceRequest) GetAlarmMethod() *string {
	return s.AlarmMethod
}

func (s *CreateDeviceRequest) GetAutoDirectory() *bool {
	return s.AutoDirectory
}

func (s *CreateDeviceRequest) GetAutoPos() *bool {
	return s.AutoPos
}

func (s *CreateDeviceRequest) GetAutoStart() *bool {
	return s.AutoStart
}

func (s *CreateDeviceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDeviceRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateDeviceRequest) GetDsn() *string {
	return s.Dsn
}

func (s *CreateDeviceRequest) GetGbId() *string {
	return s.GbId
}

func (s *CreateDeviceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateDeviceRequest) GetIp() *string {
	return s.Ip
}

func (s *CreateDeviceRequest) GetLatitude() *string {
	return s.Latitude
}

func (s *CreateDeviceRequest) GetLongitude() *string {
	return s.Longitude
}

func (s *CreateDeviceRequest) GetName() *string {
	return s.Name
}

func (s *CreateDeviceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDeviceRequest) GetParams() *string {
	return s.Params
}

func (s *CreateDeviceRequest) GetParentId() *string {
	return s.ParentId
}

func (s *CreateDeviceRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateDeviceRequest) GetPort() *int64 {
	return s.Port
}

func (s *CreateDeviceRequest) GetPosInterval() *int64 {
	return s.PosInterval
}

func (s *CreateDeviceRequest) GetType() *string {
	return s.Type
}

func (s *CreateDeviceRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateDeviceRequest) GetUsername() *string {
	return s.Username
}

func (s *CreateDeviceRequest) GetVendor() *string {
	return s.Vendor
}

func (s *CreateDeviceRequest) SetAlarmMethod(v string) *CreateDeviceRequest {
	s.AlarmMethod = &v
	return s
}

func (s *CreateDeviceRequest) SetAutoDirectory(v bool) *CreateDeviceRequest {
	s.AutoDirectory = &v
	return s
}

func (s *CreateDeviceRequest) SetAutoPos(v bool) *CreateDeviceRequest {
	s.AutoPos = &v
	return s
}

func (s *CreateDeviceRequest) SetAutoStart(v bool) *CreateDeviceRequest {
	s.AutoStart = &v
	return s
}

func (s *CreateDeviceRequest) SetDescription(v string) *CreateDeviceRequest {
	s.Description = &v
	return s
}

func (s *CreateDeviceRequest) SetDirectoryId(v string) *CreateDeviceRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateDeviceRequest) SetDsn(v string) *CreateDeviceRequest {
	s.Dsn = &v
	return s
}

func (s *CreateDeviceRequest) SetGbId(v string) *CreateDeviceRequest {
	s.GbId = &v
	return s
}

func (s *CreateDeviceRequest) SetGroupId(v string) *CreateDeviceRequest {
	s.GroupId = &v
	return s
}

func (s *CreateDeviceRequest) SetIp(v string) *CreateDeviceRequest {
	s.Ip = &v
	return s
}

func (s *CreateDeviceRequest) SetLatitude(v string) *CreateDeviceRequest {
	s.Latitude = &v
	return s
}

func (s *CreateDeviceRequest) SetLongitude(v string) *CreateDeviceRequest {
	s.Longitude = &v
	return s
}

func (s *CreateDeviceRequest) SetName(v string) *CreateDeviceRequest {
	s.Name = &v
	return s
}

func (s *CreateDeviceRequest) SetOwnerId(v int64) *CreateDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDeviceRequest) SetParams(v string) *CreateDeviceRequest {
	s.Params = &v
	return s
}

func (s *CreateDeviceRequest) SetParentId(v string) *CreateDeviceRequest {
	s.ParentId = &v
	return s
}

func (s *CreateDeviceRequest) SetPassword(v string) *CreateDeviceRequest {
	s.Password = &v
	return s
}

func (s *CreateDeviceRequest) SetPort(v int64) *CreateDeviceRequest {
	s.Port = &v
	return s
}

func (s *CreateDeviceRequest) SetPosInterval(v int64) *CreateDeviceRequest {
	s.PosInterval = &v
	return s
}

func (s *CreateDeviceRequest) SetType(v string) *CreateDeviceRequest {
	s.Type = &v
	return s
}

func (s *CreateDeviceRequest) SetUrl(v string) *CreateDeviceRequest {
	s.Url = &v
	return s
}

func (s *CreateDeviceRequest) SetUsername(v string) *CreateDeviceRequest {
	s.Username = &v
	return s
}

func (s *CreateDeviceRequest) SetVendor(v string) *CreateDeviceRequest {
	s.Vendor = &v
	return s
}

func (s *CreateDeviceRequest) Validate() error {
	return dara.Validate(s)
}
