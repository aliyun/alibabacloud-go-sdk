// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDevices(v []*DescribeDevicesResponseBodyDevices) *DescribeDevicesResponseBody
	GetDevices() []*DescribeDevicesResponseBodyDevices
	SetPageCount(v int64) *DescribeDevicesResponseBody
	GetPageCount() *int64
	SetPageNum(v int64) *DescribeDevicesResponseBody
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeDevicesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDevicesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDevicesResponseBody
	GetTotalCount() *int64
}

type DescribeDevicesResponseBody struct {
	Devices []*DescribeDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 77
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseBody) GetDevices() []*DescribeDevicesResponseBodyDevices {
	return s.Devices
}

func (s *DescribeDevicesResponseBody) GetPageCount() *int64 {
	return s.PageCount
}

func (s *DescribeDevicesResponseBody) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeDevicesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDevicesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDevicesResponseBody) SetDevices(v []*DescribeDevicesResponseBodyDevices) *DescribeDevicesResponseBody {
	s.Devices = v
	return s
}

func (s *DescribeDevicesResponseBody) SetPageCount(v int64) *DescribeDevicesResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeDevicesResponseBody) SetPageNum(v int64) *DescribeDevicesResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeDevicesResponseBody) SetPageSize(v int64) *DescribeDevicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDevicesResponseBody) SetRequestId(v string) *DescribeDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDevicesResponseBody) SetTotalCount(v int64) *DescribeDevicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDevicesResponseBody) Validate() error {
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

type DescribeDevicesResponseBodyDevices struct {
	// example:
	//
	// 0
	AlarmMethod *string `json:"AlarmMethod,omitempty" xml:"AlarmMethod,omitempty"`
	// example:
	//
	// true
	AutoDirectory *bool `json:"AutoDirectory,omitempty" xml:"AutoDirectory,omitempty"`
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
	CreatedTime *string                                      `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	Directory   *DescribeDevicesResponseBodyDevicesDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// example:
	//
	// 399*****488-cn-qingdao
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
	// 310000000****0000002
	GbId *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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
	// example:
	//
	// {}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// example:
	//
	// 3238848****092995
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
	RegisteredTime *string                                  `json:"RegisteredTime,omitempty" xml:"RegisteredTime,omitempty"`
	Stats          *DescribeDevicesResponseBodyDevicesStats `json:"Stats,omitempty" xml:"Stats,omitempty" type:"Struct"`
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

func (s DescribeDevicesResponseBodyDevices) String() string {
	return dara.Prettify(s)
}

func (s DescribeDevicesResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseBodyDevices) GetAlarmMethod() *string {
	return s.AlarmMethod
}

func (s *DescribeDevicesResponseBodyDevices) GetAutoDirectory() *bool {
	return s.AutoDirectory
}

func (s *DescribeDevicesResponseBodyDevices) GetAutoPos() *bool {
	return s.AutoPos
}

func (s *DescribeDevicesResponseBodyDevices) GetAutoStart() *bool {
	return s.AutoStart
}

func (s *DescribeDevicesResponseBodyDevices) GetChannelSyncTime() *string {
	return s.ChannelSyncTime
}

func (s *DescribeDevicesResponseBodyDevices) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeDevicesResponseBodyDevices) GetDescription() *string {
	return s.Description
}

func (s *DescribeDevicesResponseBodyDevices) GetDirectory() *DescribeDevicesResponseBodyDevicesDirectory {
	return s.Directory
}

func (s *DescribeDevicesResponseBodyDevices) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeDevicesResponseBodyDevices) GetDsn() *string {
	return s.Dsn
}

func (s *DescribeDevicesResponseBodyDevices) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeDevicesResponseBodyDevices) GetGbId() *string {
	return s.GbId
}

func (s *DescribeDevicesResponseBodyDevices) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDevicesResponseBodyDevices) GetId() *string {
	return s.Id
}

func (s *DescribeDevicesResponseBodyDevices) GetIp() *string {
	return s.Ip
}

func (s *DescribeDevicesResponseBodyDevices) GetLatitude() *string {
	return s.Latitude
}

func (s *DescribeDevicesResponseBodyDevices) GetLongitude() *string {
	return s.Longitude
}

func (s *DescribeDevicesResponseBodyDevices) GetName() *string {
	return s.Name
}

func (s *DescribeDevicesResponseBodyDevices) GetParams() *string {
	return s.Params
}

func (s *DescribeDevicesResponseBodyDevices) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeDevicesResponseBodyDevices) GetPassword() *string {
	return s.Password
}

func (s *DescribeDevicesResponseBodyDevices) GetPort() *int64 {
	return s.Port
}

func (s *DescribeDevicesResponseBodyDevices) GetPosInterval() *int64 {
	return s.PosInterval
}

func (s *DescribeDevicesResponseBodyDevices) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeDevicesResponseBodyDevices) GetRegisteredTime() *string {
	return s.RegisteredTime
}

func (s *DescribeDevicesResponseBodyDevices) GetStats() *DescribeDevicesResponseBodyDevicesStats {
	return s.Stats
}

func (s *DescribeDevicesResponseBodyDevices) GetStatus() *string {
	return s.Status
}

func (s *DescribeDevicesResponseBodyDevices) GetType() *string {
	return s.Type
}

func (s *DescribeDevicesResponseBodyDevices) GetUrl() *string {
	return s.Url
}

func (s *DescribeDevicesResponseBodyDevices) GetUsername() *string {
	return s.Username
}

func (s *DescribeDevicesResponseBodyDevices) GetVendor() *string {
	return s.Vendor
}

func (s *DescribeDevicesResponseBodyDevices) SetAlarmMethod(v string) *DescribeDevicesResponseBodyDevices {
	s.AlarmMethod = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetAutoDirectory(v bool) *DescribeDevicesResponseBodyDevices {
	s.AutoDirectory = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetAutoPos(v bool) *DescribeDevicesResponseBodyDevices {
	s.AutoPos = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetAutoStart(v bool) *DescribeDevicesResponseBodyDevices {
	s.AutoStart = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetChannelSyncTime(v string) *DescribeDevicesResponseBodyDevices {
	s.ChannelSyncTime = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetCreatedTime(v string) *DescribeDevicesResponseBodyDevices {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetDescription(v string) *DescribeDevicesResponseBodyDevices {
	s.Description = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetDirectory(v *DescribeDevicesResponseBodyDevicesDirectory) *DescribeDevicesResponseBodyDevices {
	s.Directory = v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetDirectoryId(v string) *DescribeDevicesResponseBodyDevices {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetDsn(v string) *DescribeDevicesResponseBodyDevices {
	s.Dsn = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetEnabled(v bool) *DescribeDevicesResponseBodyDevices {
	s.Enabled = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetGbId(v string) *DescribeDevicesResponseBodyDevices {
	s.GbId = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetGroupId(v string) *DescribeDevicesResponseBodyDevices {
	s.GroupId = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetId(v string) *DescribeDevicesResponseBodyDevices {
	s.Id = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetIp(v string) *DescribeDevicesResponseBodyDevices {
	s.Ip = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetLatitude(v string) *DescribeDevicesResponseBodyDevices {
	s.Latitude = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetLongitude(v string) *DescribeDevicesResponseBodyDevices {
	s.Longitude = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetName(v string) *DescribeDevicesResponseBodyDevices {
	s.Name = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetParams(v string) *DescribeDevicesResponseBodyDevices {
	s.Params = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetParentId(v string) *DescribeDevicesResponseBodyDevices {
	s.ParentId = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetPassword(v string) *DescribeDevicesResponseBodyDevices {
	s.Password = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetPort(v int64) *DescribeDevicesResponseBodyDevices {
	s.Port = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetPosInterval(v int64) *DescribeDevicesResponseBodyDevices {
	s.PosInterval = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetProtocol(v string) *DescribeDevicesResponseBodyDevices {
	s.Protocol = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetRegisteredTime(v string) *DescribeDevicesResponseBodyDevices {
	s.RegisteredTime = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetStats(v *DescribeDevicesResponseBodyDevicesStats) *DescribeDevicesResponseBodyDevices {
	s.Stats = v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetStatus(v string) *DescribeDevicesResponseBodyDevices {
	s.Status = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetType(v string) *DescribeDevicesResponseBodyDevices {
	s.Type = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetUrl(v string) *DescribeDevicesResponseBodyDevices {
	s.Url = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetUsername(v string) *DescribeDevicesResponseBodyDevices {
	s.Username = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetVendor(v string) *DescribeDevicesResponseBodyDevices {
	s.Vendor = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) Validate() error {
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

type DescribeDevicesResponseBodyDevicesDirectory struct {
	// example:
	//
	// 2019-02-28T17:00:17Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 399*****488-cn-qingdao
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 399*****774-cn-qingdao
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s DescribeDevicesResponseBodyDevicesDirectory) String() string {
	return dara.Prettify(s)
}

func (s DescribeDevicesResponseBodyDevicesDirectory) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) GetDescription() *string {
	return s.Description
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) GetId() *string {
	return s.Id
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) GetName() *string {
	return s.Name
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) SetCreatedTime(v string) *DescribeDevicesResponseBodyDevicesDirectory {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) SetDescription(v string) *DescribeDevicesResponseBodyDevicesDirectory {
	s.Description = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) SetGroupId(v string) *DescribeDevicesResponseBodyDevicesDirectory {
	s.GroupId = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) SetId(v string) *DescribeDevicesResponseBodyDevicesDirectory {
	s.Id = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) SetName(v string) *DescribeDevicesResponseBodyDevicesDirectory {
	s.Name = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) SetParentId(v string) *DescribeDevicesResponseBodyDevicesDirectory {
	s.ParentId = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesDirectory) Validate() error {
	return dara.Validate(s)
}

type DescribeDevicesResponseBodyDevicesStats struct {
	// example:
	//
	// 0
	ChannelNum *int64 `json:"ChannelNum,omitempty" xml:"ChannelNum,omitempty"`
	// example:
	//
	// 0
	FailedNum *int64 `json:"FailedNum,omitempty" xml:"FailedNum,omitempty"`
	// example:
	//
	// 0
	OfflineNum *int64 `json:"OfflineNum,omitempty" xml:"OfflineNum,omitempty"`
	// example:
	//
	// 0
	OnlineNum *int64 `json:"OnlineNum,omitempty" xml:"OnlineNum,omitempty"`
	// example:
	//
	// 0
	StreamNum *int64 `json:"StreamNum,omitempty" xml:"StreamNum,omitempty"`
}

func (s DescribeDevicesResponseBodyDevicesStats) String() string {
	return dara.Prettify(s)
}

func (s DescribeDevicesResponseBodyDevicesStats) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseBodyDevicesStats) GetChannelNum() *int64 {
	return s.ChannelNum
}

func (s *DescribeDevicesResponseBodyDevicesStats) GetFailedNum() *int64 {
	return s.FailedNum
}

func (s *DescribeDevicesResponseBodyDevicesStats) GetOfflineNum() *int64 {
	return s.OfflineNum
}

func (s *DescribeDevicesResponseBodyDevicesStats) GetOnlineNum() *int64 {
	return s.OnlineNum
}

func (s *DescribeDevicesResponseBodyDevicesStats) GetStreamNum() *int64 {
	return s.StreamNum
}

func (s *DescribeDevicesResponseBodyDevicesStats) SetChannelNum(v int64) *DescribeDevicesResponseBodyDevicesStats {
	s.ChannelNum = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesStats) SetFailedNum(v int64) *DescribeDevicesResponseBodyDevicesStats {
	s.FailedNum = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesStats) SetOfflineNum(v int64) *DescribeDevicesResponseBodyDevicesStats {
	s.OfflineNum = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesStats) SetOnlineNum(v int64) *DescribeDevicesResponseBodyDevicesStats {
	s.OnlineNum = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesStats) SetStreamNum(v int64) *DescribeDevicesResponseBodyDevicesStats {
	s.StreamNum = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesStats) Validate() error {
	return dara.Validate(s)
}
