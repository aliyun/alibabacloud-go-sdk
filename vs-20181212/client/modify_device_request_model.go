// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmMethod(v string) *ModifyDeviceRequest
	GetAlarmMethod() *string
	SetAutoDirectory(v bool) *ModifyDeviceRequest
	GetAutoDirectory() *bool
	SetAutoPos(v bool) *ModifyDeviceRequest
	GetAutoPos() *bool
	SetAutoStart(v bool) *ModifyDeviceRequest
	GetAutoStart() *bool
	SetDescription(v string) *ModifyDeviceRequest
	GetDescription() *string
	SetDirectoryId(v string) *ModifyDeviceRequest
	GetDirectoryId() *string
	SetGbId(v string) *ModifyDeviceRequest
	GetGbId() *string
	SetGroupId(v string) *ModifyDeviceRequest
	GetGroupId() *string
	SetId(v string) *ModifyDeviceRequest
	GetId() *string
	SetIp(v string) *ModifyDeviceRequest
	GetIp() *string
	SetLatitude(v string) *ModifyDeviceRequest
	GetLatitude() *string
	SetLongitude(v string) *ModifyDeviceRequest
	GetLongitude() *string
	SetName(v string) *ModifyDeviceRequest
	GetName() *string
	SetOwnerId(v int64) *ModifyDeviceRequest
	GetOwnerId() *int64
	SetParams(v string) *ModifyDeviceRequest
	GetParams() *string
	SetParentId(v string) *ModifyDeviceRequest
	GetParentId() *string
	SetPassword(v string) *ModifyDeviceRequest
	GetPassword() *string
	SetPort(v int64) *ModifyDeviceRequest
	GetPort() *int64
	SetPosInterval(v int64) *ModifyDeviceRequest
	GetPosInterval() *int64
	SetType(v string) *ModifyDeviceRequest
	GetType() *string
	SetUrl(v string) *ModifyDeviceRequest
	GetUrl() *string
	SetUsername(v string) *ModifyDeviceRequest
	GetUsername() *string
	SetVendor(v string) *ModifyDeviceRequest
	GetVendor() *string
}

type ModifyDeviceRequest struct {
	// example:
	//
	// 0
	AlarmMethod   *string `json:"AlarmMethod,omitempty" xml:"AlarmMethod,omitempty"`
	AutoDirectory *bool   `json:"AutoDirectory,omitempty" xml:"AutoDirectory,omitempty"`
	// example:
	//
	// false
	AutoPos *bool `json:"AutoPos,omitempty" xml:"AutoPos,omitempty"`
	// example:
	//
	// false
	AutoStart   *bool   `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 399*****488-cn-qingdao
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// example:
	//
	// 3100000****000000002
	GbId *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 10.10.10.10
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// 119.20
	Latitude *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// example:
	//
	// 45.00
	Longitude *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// {}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// example:
	//
	// 399*****774-cn-qingdao
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// admin
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// 8080
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 300
	PosInterval *int64 `json:"PosInterval,omitempty" xml:"PosInterval,omitempty"`
	// example:
	//
	// ipc
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// rtmp://xxx/xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	Vendor   *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ModifyDeviceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceRequest) GoString() string {
	return s.String()
}

func (s *ModifyDeviceRequest) GetAlarmMethod() *string {
	return s.AlarmMethod
}

func (s *ModifyDeviceRequest) GetAutoDirectory() *bool {
	return s.AutoDirectory
}

func (s *ModifyDeviceRequest) GetAutoPos() *bool {
	return s.AutoPos
}

func (s *ModifyDeviceRequest) GetAutoStart() *bool {
	return s.AutoStart
}

func (s *ModifyDeviceRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDeviceRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ModifyDeviceRequest) GetGbId() *string {
	return s.GbId
}

func (s *ModifyDeviceRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyDeviceRequest) GetId() *string {
	return s.Id
}

func (s *ModifyDeviceRequest) GetIp() *string {
	return s.Ip
}

func (s *ModifyDeviceRequest) GetLatitude() *string {
	return s.Latitude
}

func (s *ModifyDeviceRequest) GetLongitude() *string {
	return s.Longitude
}

func (s *ModifyDeviceRequest) GetName() *string {
	return s.Name
}

func (s *ModifyDeviceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDeviceRequest) GetParams() *string {
	return s.Params
}

func (s *ModifyDeviceRequest) GetParentId() *string {
	return s.ParentId
}

func (s *ModifyDeviceRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyDeviceRequest) GetPort() *int64 {
	return s.Port
}

func (s *ModifyDeviceRequest) GetPosInterval() *int64 {
	return s.PosInterval
}

func (s *ModifyDeviceRequest) GetType() *string {
	return s.Type
}

func (s *ModifyDeviceRequest) GetUrl() *string {
	return s.Url
}

func (s *ModifyDeviceRequest) GetUsername() *string {
	return s.Username
}

func (s *ModifyDeviceRequest) GetVendor() *string {
	return s.Vendor
}

func (s *ModifyDeviceRequest) SetAlarmMethod(v string) *ModifyDeviceRequest {
	s.AlarmMethod = &v
	return s
}

func (s *ModifyDeviceRequest) SetAutoDirectory(v bool) *ModifyDeviceRequest {
	s.AutoDirectory = &v
	return s
}

func (s *ModifyDeviceRequest) SetAutoPos(v bool) *ModifyDeviceRequest {
	s.AutoPos = &v
	return s
}

func (s *ModifyDeviceRequest) SetAutoStart(v bool) *ModifyDeviceRequest {
	s.AutoStart = &v
	return s
}

func (s *ModifyDeviceRequest) SetDescription(v string) *ModifyDeviceRequest {
	s.Description = &v
	return s
}

func (s *ModifyDeviceRequest) SetDirectoryId(v string) *ModifyDeviceRequest {
	s.DirectoryId = &v
	return s
}

func (s *ModifyDeviceRequest) SetGbId(v string) *ModifyDeviceRequest {
	s.GbId = &v
	return s
}

func (s *ModifyDeviceRequest) SetGroupId(v string) *ModifyDeviceRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyDeviceRequest) SetId(v string) *ModifyDeviceRequest {
	s.Id = &v
	return s
}

func (s *ModifyDeviceRequest) SetIp(v string) *ModifyDeviceRequest {
	s.Ip = &v
	return s
}

func (s *ModifyDeviceRequest) SetLatitude(v string) *ModifyDeviceRequest {
	s.Latitude = &v
	return s
}

func (s *ModifyDeviceRequest) SetLongitude(v string) *ModifyDeviceRequest {
	s.Longitude = &v
	return s
}

func (s *ModifyDeviceRequest) SetName(v string) *ModifyDeviceRequest {
	s.Name = &v
	return s
}

func (s *ModifyDeviceRequest) SetOwnerId(v int64) *ModifyDeviceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDeviceRequest) SetParams(v string) *ModifyDeviceRequest {
	s.Params = &v
	return s
}

func (s *ModifyDeviceRequest) SetParentId(v string) *ModifyDeviceRequest {
	s.ParentId = &v
	return s
}

func (s *ModifyDeviceRequest) SetPassword(v string) *ModifyDeviceRequest {
	s.Password = &v
	return s
}

func (s *ModifyDeviceRequest) SetPort(v int64) *ModifyDeviceRequest {
	s.Port = &v
	return s
}

func (s *ModifyDeviceRequest) SetPosInterval(v int64) *ModifyDeviceRequest {
	s.PosInterval = &v
	return s
}

func (s *ModifyDeviceRequest) SetType(v string) *ModifyDeviceRequest {
	s.Type = &v
	return s
}

func (s *ModifyDeviceRequest) SetUrl(v string) *ModifyDeviceRequest {
	s.Url = &v
	return s
}

func (s *ModifyDeviceRequest) SetUsername(v string) *ModifyDeviceRequest {
	s.Username = &v
	return s
}

func (s *ModifyDeviceRequest) SetVendor(v string) *ModifyDeviceRequest {
	s.Vendor = &v
	return s
}

func (s *ModifyDeviceRequest) Validate() error {
	return dara.Validate(s)
}
