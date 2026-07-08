// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v []*DescribeGroupsResponseBodyGroups) *DescribeGroupsResponseBody
	GetGroups() []*DescribeGroupsResponseBodyGroups
	SetPageCount(v int64) *DescribeGroupsResponseBody
	GetPageCount() *int64
	SetPageNum(v int64) *DescribeGroupsResponseBody
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeGroupsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeGroupsResponseBody
	GetTotalCount() *int64
}

type DescribeGroupsResponseBody struct {
	// List of groups.
	Groups []*DescribeGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// Total number of pages.
	//
	// example:
	//
	// 10
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// Page number.
	//
	// example:
	//
	// 5
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of groups.
	//
	// example:
	//
	// 200
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponseBody) GetGroups() []*DescribeGroupsResponseBodyGroups {
	return s.Groups
}

func (s *DescribeGroupsResponseBody) GetPageCount() *int64 {
	return s.PageCount
}

func (s *DescribeGroupsResponseBody) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeGroupsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeGroupsResponseBody) SetGroups(v []*DescribeGroupsResponseBodyGroups) *DescribeGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeGroupsResponseBody) SetPageCount(v int64) *DescribeGroupsResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeGroupsResponseBody) SetPageNum(v int64) *DescribeGroupsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeGroupsResponseBody) SetPageSize(v int64) *DescribeGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupsResponseBody) SetRequestId(v string) *DescribeGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupsResponseBody) SetTotalCount(v int64) *DescribeGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGroupsResponseBody) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGroupsResponseBodyGroups struct {
	// Alias for the group ID.
	//
	// example:
	//
	// 337639*****24964-cn-qingdao
	AliasId *string `json:"AliasId,omitempty" xml:"AliasId,omitempty"`
	// Application name used by the group.
	//
	// example:
	//
	// live
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// Callback URL for device status updates in the group.
	//
	// example:
	//
	// http://example.com/callback
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// Time when the group was created.
	//
	// example:
	//
	// 2019-02-28T17:00:17Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// Group description.
	//
	// example:
	//
	// 测试视图计算
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether the group is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// GB ID provided by the group.
	//
	// > This applies only to groups that use the GB protocol.
	//
	// example:
	//
	// 31000000000000000001
	GbId *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	// GB signaling server IP address provided by the group.
	//
	// > This applies only to groups that use the GB protocol.
	//
	// example:
	//
	// 10.10.10.10
	GbIp *string `json:"GbIp,omitempty" xml:"GbIp,omitempty"`
	// GB signaling server port associated with the group.
	//
	// > This applies only to groups that use the GB protocol.
	//
	// example:
	//
	// 5060
	GbPort     *int64    `json:"GbPort,omitempty" xml:"GbPort,omitempty"`
	GbTcpPorts []*string `json:"GbTcpPorts,omitempty" xml:"GbTcpPorts,omitempty" type:"Repeated"`
	GbUdpPorts []*string `json:"GbUdpPorts,omitempty" xml:"GbUdpPorts,omitempty" type:"Repeated"`
	// Group ID.
	//
	// example:
	//
	// 33763950877224964-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Ingest protocol used by the group.
	//
	// example:
	//
	// gb28181
	InProtocol *string `json:"InProtocol,omitempty" xml:"InProtocol,omitempty"`
	// Whether on-demand stream pulling is enabled.
	//
	// example:
	//
	// false
	LazyPull *bool `json:"LazyPull,omitempty" xml:"LazyPull,omitempty"`
	// Group name.
	//
	// example:
	//
	// 我的空间
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Playback protocol used by the group.
	//
	// example:
	//
	// rtmp,flv,hls
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	// Streaming domain used by the group.
	//
	// example:
	//
	// demo.aliyundoc.com
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	// Ingest domain used by the group.
	//
	// example:
	//
	// example.aliyundoc.com
	PushDomain *string `json:"PushDomain,omitempty" xml:"PushDomain,omitempty"`
	// The region where the space is located. This region serves as the service center.
	//
	// example:
	//
	// cn-qingdao
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Device statistics for the group.
	Stats *DescribeGroupsResponseBodyGroupsStats `json:"Stats,omitempty" xml:"Stats,omitempty" type:"Struct"`
	// Group status. Valid values:
	//
	// - on (enabled)
	//
	// - off (disabled)
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGroupsResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponseBodyGroups) GetAliasId() *string {
	return s.AliasId
}

func (s *DescribeGroupsResponseBodyGroups) GetApp() *string {
	return s.App
}

func (s *DescribeGroupsResponseBodyGroups) GetCallback() *string {
	return s.Callback
}

func (s *DescribeGroupsResponseBodyGroups) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeGroupsResponseBodyGroups) GetDescription() *string {
	return s.Description
}

func (s *DescribeGroupsResponseBodyGroups) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeGroupsResponseBodyGroups) GetGbId() *string {
	return s.GbId
}

func (s *DescribeGroupsResponseBodyGroups) GetGbIp() *string {
	return s.GbIp
}

func (s *DescribeGroupsResponseBodyGroups) GetGbPort() *int64 {
	return s.GbPort
}

func (s *DescribeGroupsResponseBodyGroups) GetGbTcpPorts() []*string {
	return s.GbTcpPorts
}

func (s *DescribeGroupsResponseBodyGroups) GetGbUdpPorts() []*string {
	return s.GbUdpPorts
}

func (s *DescribeGroupsResponseBodyGroups) GetId() *string {
	return s.Id
}

func (s *DescribeGroupsResponseBodyGroups) GetInProtocol() *string {
	return s.InProtocol
}

func (s *DescribeGroupsResponseBodyGroups) GetLazyPull() *bool {
	return s.LazyPull
}

func (s *DescribeGroupsResponseBodyGroups) GetName() *string {
	return s.Name
}

func (s *DescribeGroupsResponseBodyGroups) GetOutProtocol() *string {
	return s.OutProtocol
}

func (s *DescribeGroupsResponseBodyGroups) GetPlayDomain() *string {
	return s.PlayDomain
}

func (s *DescribeGroupsResponseBodyGroups) GetPushDomain() *string {
	return s.PushDomain
}

func (s *DescribeGroupsResponseBodyGroups) GetRegion() *string {
	return s.Region
}

func (s *DescribeGroupsResponseBodyGroups) GetStats() *DescribeGroupsResponseBodyGroupsStats {
	return s.Stats
}

func (s *DescribeGroupsResponseBodyGroups) GetStatus() *string {
	return s.Status
}

func (s *DescribeGroupsResponseBodyGroups) SetAliasId(v string) *DescribeGroupsResponseBodyGroups {
	s.AliasId = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetApp(v string) *DescribeGroupsResponseBodyGroups {
	s.App = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetCallback(v string) *DescribeGroupsResponseBodyGroups {
	s.Callback = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetCreatedTime(v string) *DescribeGroupsResponseBodyGroups {
	s.CreatedTime = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetDescription(v string) *DescribeGroupsResponseBodyGroups {
	s.Description = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetEnabled(v bool) *DescribeGroupsResponseBodyGroups {
	s.Enabled = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetGbId(v string) *DescribeGroupsResponseBodyGroups {
	s.GbId = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetGbIp(v string) *DescribeGroupsResponseBodyGroups {
	s.GbIp = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetGbPort(v int64) *DescribeGroupsResponseBodyGroups {
	s.GbPort = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetGbTcpPorts(v []*string) *DescribeGroupsResponseBodyGroups {
	s.GbTcpPorts = v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetGbUdpPorts(v []*string) *DescribeGroupsResponseBodyGroups {
	s.GbUdpPorts = v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetId(v string) *DescribeGroupsResponseBodyGroups {
	s.Id = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetInProtocol(v string) *DescribeGroupsResponseBodyGroups {
	s.InProtocol = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetLazyPull(v bool) *DescribeGroupsResponseBodyGroups {
	s.LazyPull = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetName(v string) *DescribeGroupsResponseBodyGroups {
	s.Name = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetOutProtocol(v string) *DescribeGroupsResponseBodyGroups {
	s.OutProtocol = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetPlayDomain(v string) *DescribeGroupsResponseBodyGroups {
	s.PlayDomain = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetPushDomain(v string) *DescribeGroupsResponseBodyGroups {
	s.PushDomain = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetRegion(v string) *DescribeGroupsResponseBodyGroups {
	s.Region = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetStats(v *DescribeGroupsResponseBodyGroupsStats) *DescribeGroupsResponseBodyGroups {
	s.Stats = v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetStatus(v string) *DescribeGroupsResponseBodyGroups {
	s.Status = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) Validate() error {
	if s.Stats != nil {
		if err := s.Stats.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGroupsResponseBodyGroupsStats struct {
	// Total number of devices in the group.
	//
	// example:
	//
	// 200
	DeviceNum *int64 `json:"DeviceNum,omitempty" xml:"DeviceNum,omitempty"`
	// Number of intelligent devices in the group.
	//
	// example:
	//
	// 0
	IedNum *int64 `json:"IedNum,omitempty" xml:"IedNum,omitempty"`
	// Number of cameras in the group.
	//
	// example:
	//
	// 200
	IpcNum *int64 `json:"IpcNum,omitempty" xml:"IpcNum,omitempty"`
	// Number of platforms in the group.
	//
	// example:
	//
	// 0
	PlatformNum *int64 `json:"PlatformNum,omitempty" xml:"PlatformNum,omitempty"`
}

func (s DescribeGroupsResponseBodyGroupsStats) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupsResponseBodyGroupsStats) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponseBodyGroupsStats) GetDeviceNum() *int64 {
	return s.DeviceNum
}

func (s *DescribeGroupsResponseBodyGroupsStats) GetIedNum() *int64 {
	return s.IedNum
}

func (s *DescribeGroupsResponseBodyGroupsStats) GetIpcNum() *int64 {
	return s.IpcNum
}

func (s *DescribeGroupsResponseBodyGroupsStats) GetPlatformNum() *int64 {
	return s.PlatformNum
}

func (s *DescribeGroupsResponseBodyGroupsStats) SetDeviceNum(v int64) *DescribeGroupsResponseBodyGroupsStats {
	s.DeviceNum = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsStats) SetIedNum(v int64) *DescribeGroupsResponseBodyGroupsStats {
	s.IedNum = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsStats) SetIpcNum(v int64) *DescribeGroupsResponseBodyGroupsStats {
	s.IpcNum = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsStats) SetPlatformNum(v int64) *DescribeGroupsResponseBodyGroupsStats {
	s.PlatformNum = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroupsStats) Validate() error {
	return dara.Validate(s)
}
