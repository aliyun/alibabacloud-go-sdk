// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmMethod(v string) *DescribeDeviceResponseBody
	GetAlarmMethod() *string
	SetAutoDirectory(v bool) *DescribeDeviceResponseBody
	GetAutoDirectory() *bool
	SetAutoPos(v bool) *DescribeDeviceResponseBody
	GetAutoPos() *bool
	SetAutoStart(v bool) *DescribeDeviceResponseBody
	GetAutoStart() *bool
	SetChannelSyncTime(v string) *DescribeDeviceResponseBody
	GetChannelSyncTime() *string
	SetCreatedTime(v string) *DescribeDeviceResponseBody
	GetCreatedTime() *string
	SetDescription(v string) *DescribeDeviceResponseBody
	GetDescription() *string
	SetDirectory(v *DescribeDeviceResponseBodyDirectory) *DescribeDeviceResponseBody
	GetDirectory() *DescribeDeviceResponseBodyDirectory
	SetDirectoryId(v string) *DescribeDeviceResponseBody
	GetDirectoryId() *string
	SetDsn(v string) *DescribeDeviceResponseBody
	GetDsn() *string
	SetEnabled(v bool) *DescribeDeviceResponseBody
	GetEnabled() *bool
	SetGbId(v string) *DescribeDeviceResponseBody
	GetGbId() *string
	SetGroupId(v string) *DescribeDeviceResponseBody
	GetGroupId() *string
	SetId(v string) *DescribeDeviceResponseBody
	GetId() *string
	SetIp(v string) *DescribeDeviceResponseBody
	GetIp() *string
	SetLatitude(v string) *DescribeDeviceResponseBody
	GetLatitude() *string
	SetLongitude(v string) *DescribeDeviceResponseBody
	GetLongitude() *string
	SetName(v string) *DescribeDeviceResponseBody
	GetName() *string
	SetParams(v string) *DescribeDeviceResponseBody
	GetParams() *string
	SetParentId(v string) *DescribeDeviceResponseBody
	GetParentId() *string
	SetPassword(v string) *DescribeDeviceResponseBody
	GetPassword() *string
	SetPort(v int64) *DescribeDeviceResponseBody
	GetPort() *int64
	SetPosInterval(v int64) *DescribeDeviceResponseBody
	GetPosInterval() *int64
	SetProtocol(v string) *DescribeDeviceResponseBody
	GetProtocol() *string
	SetRegisteredTime(v string) *DescribeDeviceResponseBody
	GetRegisteredTime() *string
	SetRequestId(v string) *DescribeDeviceResponseBody
	GetRequestId() *string
	SetStats(v *DescribeDeviceResponseBodyStats) *DescribeDeviceResponseBody
	GetStats() *DescribeDeviceResponseBodyStats
	SetStatus(v string) *DescribeDeviceResponseBody
	GetStatus() *string
	SetType(v string) *DescribeDeviceResponseBody
	GetType() *string
	SetUrl(v string) *DescribeDeviceResponseBody
	GetUrl() *string
	SetUsername(v string) *DescribeDeviceResponseBody
	GetUsername() *string
	SetVendor(v string) *DescribeDeviceResponseBody
	GetVendor() *string
}

type DescribeDeviceResponseBody struct {
	// example:
	//
	// 5
	AlarmMethod   *string `json:"AlarmMethod,omitempty" xml:"AlarmMethod,omitempty"`
	AutoDirectory *bool   `json:"AutoDirectory,omitempty" xml:"AutoDirectory,omitempty"`
	// example:
	//
	// false
	AutoPos *bool `json:"AutoPos,omitempty" xml:"AutoPos,omitempty"`
	// example:
	//
	// false
	AutoStart *bool `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	// example:
	//
	// 2019-02-28T17:01:17Z
	ChannelSyncTime *string `json:"ChannelSyncTime,omitempty" xml:"ChannelSyncTime,omitempty"`
	// example:
	//
	// 2019-02-28T17:00:17Z
	CreatedTime *string                              `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Directory   *DescribeDeviceResponseBodyDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// example:
	//
	// 3238848****092994-cn-qingdao
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// example:
	//
	// 7D0*****4C0
	Dsn *string `json:"Dsn,omitempty" xml:"Dsn,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 31000000****00000002
	GbId *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	// example:
	//
	// 3238848****092994-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 3238848****092996-cn-qingdao
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
	// example:
	//
	// {}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// example:
	//
	// 3238848****092995-cn-qingdao
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
	// gb28181
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// 2019-02-28T17:00:17Z
	RegisteredTime *string `json:"RegisteredTime,omitempty" xml:"RegisteredTime,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Stats     *DescribeDeviceResponseBodyStats `json:"Stats,omitempty" xml:"Stats,omitempty" type:"Struct"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s DescribeDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceResponseBody) GetAlarmMethod() *string {
	return s.AlarmMethod
}

func (s *DescribeDeviceResponseBody) GetAutoDirectory() *bool {
	return s.AutoDirectory
}

func (s *DescribeDeviceResponseBody) GetAutoPos() *bool {
	return s.AutoPos
}

func (s *DescribeDeviceResponseBody) GetAutoStart() *bool {
	return s.AutoStart
}

func (s *DescribeDeviceResponseBody) GetChannelSyncTime() *string {
	return s.ChannelSyncTime
}

func (s *DescribeDeviceResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeDeviceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeDeviceResponseBody) GetDirectory() *DescribeDeviceResponseBodyDirectory {
	return s.Directory
}

func (s *DescribeDeviceResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeDeviceResponseBody) GetDsn() *string {
	return s.Dsn
}

func (s *DescribeDeviceResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeDeviceResponseBody) GetGbId() *string {
	return s.GbId
}

func (s *DescribeDeviceResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDeviceResponseBody) GetId() *string {
	return s.Id
}

func (s *DescribeDeviceResponseBody) GetIp() *string {
	return s.Ip
}

func (s *DescribeDeviceResponseBody) GetLatitude() *string {
	return s.Latitude
}

func (s *DescribeDeviceResponseBody) GetLongitude() *string {
	return s.Longitude
}

func (s *DescribeDeviceResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeDeviceResponseBody) GetParams() *string {
	return s.Params
}

func (s *DescribeDeviceResponseBody) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeDeviceResponseBody) GetPassword() *string {
	return s.Password
}

func (s *DescribeDeviceResponseBody) GetPort() *int64 {
	return s.Port
}

func (s *DescribeDeviceResponseBody) GetPosInterval() *int64 {
	return s.PosInterval
}

func (s *DescribeDeviceResponseBody) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeDeviceResponseBody) GetRegisteredTime() *string {
	return s.RegisteredTime
}

func (s *DescribeDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeviceResponseBody) GetStats() *DescribeDeviceResponseBodyStats {
	return s.Stats
}

func (s *DescribeDeviceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDeviceResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeDeviceResponseBody) GetUrl() *string {
	return s.Url
}

func (s *DescribeDeviceResponseBody) GetUsername() *string {
	return s.Username
}

func (s *DescribeDeviceResponseBody) GetVendor() *string {
	return s.Vendor
}

func (s *DescribeDeviceResponseBody) SetAlarmMethod(v string) *DescribeDeviceResponseBody {
	s.AlarmMethod = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetAutoDirectory(v bool) *DescribeDeviceResponseBody {
	s.AutoDirectory = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetAutoPos(v bool) *DescribeDeviceResponseBody {
	s.AutoPos = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetAutoStart(v bool) *DescribeDeviceResponseBody {
	s.AutoStart = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetChannelSyncTime(v string) *DescribeDeviceResponseBody {
	s.ChannelSyncTime = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetCreatedTime(v string) *DescribeDeviceResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetDescription(v string) *DescribeDeviceResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetDirectory(v *DescribeDeviceResponseBodyDirectory) *DescribeDeviceResponseBody {
	s.Directory = v
	return s
}

func (s *DescribeDeviceResponseBody) SetDirectoryId(v string) *DescribeDeviceResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetDsn(v string) *DescribeDeviceResponseBody {
	s.Dsn = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetEnabled(v bool) *DescribeDeviceResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetGbId(v string) *DescribeDeviceResponseBody {
	s.GbId = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetGroupId(v string) *DescribeDeviceResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetId(v string) *DescribeDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetIp(v string) *DescribeDeviceResponseBody {
	s.Ip = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetLatitude(v string) *DescribeDeviceResponseBody {
	s.Latitude = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetLongitude(v string) *DescribeDeviceResponseBody {
	s.Longitude = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetName(v string) *DescribeDeviceResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetParams(v string) *DescribeDeviceResponseBody {
	s.Params = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetParentId(v string) *DescribeDeviceResponseBody {
	s.ParentId = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetPassword(v string) *DescribeDeviceResponseBody {
	s.Password = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetPort(v int64) *DescribeDeviceResponseBody {
	s.Port = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetPosInterval(v int64) *DescribeDeviceResponseBody {
	s.PosInterval = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetProtocol(v string) *DescribeDeviceResponseBody {
	s.Protocol = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetRegisteredTime(v string) *DescribeDeviceResponseBody {
	s.RegisteredTime = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetRequestId(v string) *DescribeDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetStats(v *DescribeDeviceResponseBodyStats) *DescribeDeviceResponseBody {
	s.Stats = v
	return s
}

func (s *DescribeDeviceResponseBody) SetStatus(v string) *DescribeDeviceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetType(v string) *DescribeDeviceResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetUrl(v string) *DescribeDeviceResponseBody {
	s.Url = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetUsername(v string) *DescribeDeviceResponseBody {
	s.Username = &v
	return s
}

func (s *DescribeDeviceResponseBody) SetVendor(v string) *DescribeDeviceResponseBody {
	s.Vendor = &v
	return s
}

func (s *DescribeDeviceResponseBody) Validate() error {
	if s.Directory != nil {
		if err := s.Directory.Validate(); err != nil {
			return err
		}
	}
	if s.Stats != nil {
		if err := s.Stats.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDeviceResponseBodyDirectory struct {
	// example:
	//
	// 2019-02-28T17:00:17Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 3238848****092994-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 3238848****092994-cn-qingdao
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 3238848****092995-cn-qingdao
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s DescribeDeviceResponseBodyDirectory) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceResponseBodyDirectory) GoString() string {
	return s.String()
}

func (s *DescribeDeviceResponseBodyDirectory) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeDeviceResponseBodyDirectory) GetDescription() *string {
	return s.Description
}

func (s *DescribeDeviceResponseBodyDirectory) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDeviceResponseBodyDirectory) GetId() *string {
	return s.Id
}

func (s *DescribeDeviceResponseBodyDirectory) GetName() *string {
	return s.Name
}

func (s *DescribeDeviceResponseBodyDirectory) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeDeviceResponseBodyDirectory) SetCreatedTime(v string) *DescribeDeviceResponseBodyDirectory {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDeviceResponseBodyDirectory) SetDescription(v string) *DescribeDeviceResponseBodyDirectory {
	s.Description = &v
	return s
}

func (s *DescribeDeviceResponseBodyDirectory) SetGroupId(v string) *DescribeDeviceResponseBodyDirectory {
	s.GroupId = &v
	return s
}

func (s *DescribeDeviceResponseBodyDirectory) SetId(v string) *DescribeDeviceResponseBodyDirectory {
	s.Id = &v
	return s
}

func (s *DescribeDeviceResponseBodyDirectory) SetName(v string) *DescribeDeviceResponseBodyDirectory {
	s.Name = &v
	return s
}

func (s *DescribeDeviceResponseBodyDirectory) SetParentId(v string) *DescribeDeviceResponseBodyDirectory {
	s.ParentId = &v
	return s
}

func (s *DescribeDeviceResponseBodyDirectory) Validate() error {
	return dara.Validate(s)
}

type DescribeDeviceResponseBodyStats struct {
	// example:
	//
	// 1
	ChannelNum *int64 `json:"ChannelNum,omitempty" xml:"ChannelNum,omitempty"`
	// example:
	//
	// 1
	FailedNum *int64 `json:"FailedNum,omitempty" xml:"FailedNum,omitempty"`
	// example:
	//
	// 1
	OfflineNum *int64 `json:"OfflineNum,omitempty" xml:"OfflineNum,omitempty"`
	// example:
	//
	// 1
	OnlineNum *int64 `json:"OnlineNum,omitempty" xml:"OnlineNum,omitempty"`
	// example:
	//
	// 1
	StreamNum *int64 `json:"StreamNum,omitempty" xml:"StreamNum,omitempty"`
}

func (s DescribeDeviceResponseBodyStats) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceResponseBodyStats) GoString() string {
	return s.String()
}

func (s *DescribeDeviceResponseBodyStats) GetChannelNum() *int64 {
	return s.ChannelNum
}

func (s *DescribeDeviceResponseBodyStats) GetFailedNum() *int64 {
	return s.FailedNum
}

func (s *DescribeDeviceResponseBodyStats) GetOfflineNum() *int64 {
	return s.OfflineNum
}

func (s *DescribeDeviceResponseBodyStats) GetOnlineNum() *int64 {
	return s.OnlineNum
}

func (s *DescribeDeviceResponseBodyStats) GetStreamNum() *int64 {
	return s.StreamNum
}

func (s *DescribeDeviceResponseBodyStats) SetChannelNum(v int64) *DescribeDeviceResponseBodyStats {
	s.ChannelNum = &v
	return s
}

func (s *DescribeDeviceResponseBodyStats) SetFailedNum(v int64) *DescribeDeviceResponseBodyStats {
	s.FailedNum = &v
	return s
}

func (s *DescribeDeviceResponseBodyStats) SetOfflineNum(v int64) *DescribeDeviceResponseBodyStats {
	s.OfflineNum = &v
	return s
}

func (s *DescribeDeviceResponseBodyStats) SetOnlineNum(v int64) *DescribeDeviceResponseBodyStats {
	s.OnlineNum = &v
	return s
}

func (s *DescribeDeviceResponseBodyStats) SetStreamNum(v int64) *DescribeDeviceResponseBodyStats {
	s.StreamNum = &v
	return s
}

func (s *DescribeDeviceResponseBodyStats) Validate() error {
	return dara.Validate(s)
}
