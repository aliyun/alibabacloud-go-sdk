// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliasId(v string) *DescribeGroupResponseBody
	GetAliasId() *string
	SetApp(v string) *DescribeGroupResponseBody
	GetApp() *string
	SetCallback(v string) *DescribeGroupResponseBody
	GetCallback() *string
	SetCreatedTime(v string) *DescribeGroupResponseBody
	GetCreatedTime() *string
	SetDescription(v string) *DescribeGroupResponseBody
	GetDescription() *string
	SetEnabled(v bool) *DescribeGroupResponseBody
	GetEnabled() *bool
	SetGbId(v string) *DescribeGroupResponseBody
	GetGbId() *string
	SetGbIp(v string) *DescribeGroupResponseBody
	GetGbIp() *string
	SetGbPort(v int64) *DescribeGroupResponseBody
	GetGbPort() *int64
	SetGbTcpPorts(v []*string) *DescribeGroupResponseBody
	GetGbTcpPorts() []*string
	SetGbUdpPorts(v []*string) *DescribeGroupResponseBody
	GetGbUdpPorts() []*string
	SetId(v string) *DescribeGroupResponseBody
	GetId() *string
	SetInProtocol(v string) *DescribeGroupResponseBody
	GetInProtocol() *string
	SetLazyPull(v bool) *DescribeGroupResponseBody
	GetLazyPull() *bool
	SetName(v string) *DescribeGroupResponseBody
	GetName() *string
	SetOutProtocol(v string) *DescribeGroupResponseBody
	GetOutProtocol() *string
	SetPlayDomain(v string) *DescribeGroupResponseBody
	GetPlayDomain() *string
	SetPushDomain(v string) *DescribeGroupResponseBody
	GetPushDomain() *string
	SetRegion(v string) *DescribeGroupResponseBody
	GetRegion() *string
	SetRequestId(v string) *DescribeGroupResponseBody
	GetRequestId() *string
	SetStats(v *DescribeGroupResponseBodyStats) *DescribeGroupResponseBody
	GetStats() *DescribeGroupResponseBodyStats
	SetStatus(v string) *DescribeGroupResponseBody
	GetStatus() *string
}

type DescribeGroupResponseBody struct {
	// Alias for the space ID.
	//
	// example:
	//
	// 337639*****24964-cn-qingdao
	AliasId *string `json:"AliasId,omitempty" xml:"AliasId,omitempty"`
	// The name of the application used by the group.
	//
	// example:
	//
	// live
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// The callback URL that is used to receive device status updates in the group.
	//
	// example:
	//
	// http://example.com/callback
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// The time when the group was created.
	//
	// example:
	//
	// 2019-02-28T17:00:17Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// A description of the group.
	//
	// example:
	//
	// 上海高速监控
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the group is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The GB/T 28181 ID that is associated with the group.
	//
	// > This parameter is returned only for groups that use the GB/T 28181 protocol for stream ingest.
	//
	// example:
	//
	// 3100000*****0000001
	GbId *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	// The IP address of the GB/T 28181 signaling server that is associated with the group.
	//
	// > This parameter is returned only for groups that use the GB/T 28181 protocol for stream ingest.
	//
	// example:
	//
	// 10.10.10.10
	GbIp *string `json:"GbIp,omitempty" xml:"GbIp,omitempty"`
	// The port of the GB/T 28181 signaling server that is associated with the group.
	//
	// > This parameter is returned only for groups that use the GB/T 28181 protocol for stream ingest.
	//
	// example:
	//
	// 5060
	GbPort *int64 `json:"GbPort,omitempty" xml:"GbPort,omitempty"`
	// The TCP ports of the GB/T 28181 signaling server that are provided by the group.
	//
	// > This parameter is returned only for groups that use the GB/T 28181 protocol for stream ingest.
	GbTcpPorts []*string `json:"GbTcpPorts,omitempty" xml:"GbTcpPorts,omitempty" type:"Repeated"`
	// The UDP ports of the GB/T 28181 signaling server that are provided by the group.
	//
	// > This parameter is returned only for groups that use the GB/T 28181 protocol for stream ingest.
	GbUdpPorts []*string `json:"GbUdpPorts,omitempty" xml:"GbUdpPorts,omitempty" type:"Repeated"`
	// The ID of the space.
	//
	// example:
	//
	// 337639****224964-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ingest protocol used by the group. Valid values:
	//
	// - gb28181
	//
	// - rtmp
	//
	// example:
	//
	// gb28181
	InProtocol *string `json:"InProtocol,omitempty" xml:"InProtocol,omitempty"`
	// Indicates whether on-demand stream pulling is enabled.
	//
	// example:
	//
	// false
	LazyPull *bool `json:"LazyPull,omitempty" xml:"LazyPull,omitempty"`
	// The name of the space.
	//
	// example:
	//
	// 上海高速监控
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The streaming protocol used by the group. Valid values:
	//
	// - flv
	//
	// - hls
	//
	// - rtmp
	//
	// > You can specify multiple protocols. Separate them with commas (,).
	//
	// example:
	//
	// flv,hls,rtmp
	OutProtocol *string `json:"OutProtocol,omitempty" xml:"OutProtocol,omitempty"`
	// The streaming domain used by the group.
	//
	// example:
	//
	// example.aliyundoc.com
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	// The ingest domain used by the group.
	//
	// example:
	//
	// demo.aliyundoc.com
	PushDomain *string `json:"PushDomain,omitempty" xml:"PushDomain,omitempty"`
	// The region where the space is located, which is the service center.
	//
	// example:
	//
	// cn-qingdao
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics of devices in the group.
	Stats *DescribeGroupResponseBodyStats `json:"Stats,omitempty" xml:"Stats,omitempty" type:"Struct"`
	// The status of the group.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupResponseBody) GetAliasId() *string {
	return s.AliasId
}

func (s *DescribeGroupResponseBody) GetApp() *string {
	return s.App
}

func (s *DescribeGroupResponseBody) GetCallback() *string {
	return s.Callback
}

func (s *DescribeGroupResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeGroupResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeGroupResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeGroupResponseBody) GetGbId() *string {
	return s.GbId
}

func (s *DescribeGroupResponseBody) GetGbIp() *string {
	return s.GbIp
}

func (s *DescribeGroupResponseBody) GetGbPort() *int64 {
	return s.GbPort
}

func (s *DescribeGroupResponseBody) GetGbTcpPorts() []*string {
	return s.GbTcpPorts
}

func (s *DescribeGroupResponseBody) GetGbUdpPorts() []*string {
	return s.GbUdpPorts
}

func (s *DescribeGroupResponseBody) GetId() *string {
	return s.Id
}

func (s *DescribeGroupResponseBody) GetInProtocol() *string {
	return s.InProtocol
}

func (s *DescribeGroupResponseBody) GetLazyPull() *bool {
	return s.LazyPull
}

func (s *DescribeGroupResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeGroupResponseBody) GetOutProtocol() *string {
	return s.OutProtocol
}

func (s *DescribeGroupResponseBody) GetPlayDomain() *string {
	return s.PlayDomain
}

func (s *DescribeGroupResponseBody) GetPushDomain() *string {
	return s.PushDomain
}

func (s *DescribeGroupResponseBody) GetRegion() *string {
	return s.Region
}

func (s *DescribeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupResponseBody) GetStats() *DescribeGroupResponseBodyStats {
	return s.Stats
}

func (s *DescribeGroupResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeGroupResponseBody) SetAliasId(v string) *DescribeGroupResponseBody {
	s.AliasId = &v
	return s
}

func (s *DescribeGroupResponseBody) SetApp(v string) *DescribeGroupResponseBody {
	s.App = &v
	return s
}

func (s *DescribeGroupResponseBody) SetCallback(v string) *DescribeGroupResponseBody {
	s.Callback = &v
	return s
}

func (s *DescribeGroupResponseBody) SetCreatedTime(v string) *DescribeGroupResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeGroupResponseBody) SetDescription(v string) *DescribeGroupResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeGroupResponseBody) SetEnabled(v bool) *DescribeGroupResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeGroupResponseBody) SetGbId(v string) *DescribeGroupResponseBody {
	s.GbId = &v
	return s
}

func (s *DescribeGroupResponseBody) SetGbIp(v string) *DescribeGroupResponseBody {
	s.GbIp = &v
	return s
}

func (s *DescribeGroupResponseBody) SetGbPort(v int64) *DescribeGroupResponseBody {
	s.GbPort = &v
	return s
}

func (s *DescribeGroupResponseBody) SetGbTcpPorts(v []*string) *DescribeGroupResponseBody {
	s.GbTcpPorts = v
	return s
}

func (s *DescribeGroupResponseBody) SetGbUdpPorts(v []*string) *DescribeGroupResponseBody {
	s.GbUdpPorts = v
	return s
}

func (s *DescribeGroupResponseBody) SetId(v string) *DescribeGroupResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeGroupResponseBody) SetInProtocol(v string) *DescribeGroupResponseBody {
	s.InProtocol = &v
	return s
}

func (s *DescribeGroupResponseBody) SetLazyPull(v bool) *DescribeGroupResponseBody {
	s.LazyPull = &v
	return s
}

func (s *DescribeGroupResponseBody) SetName(v string) *DescribeGroupResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeGroupResponseBody) SetOutProtocol(v string) *DescribeGroupResponseBody {
	s.OutProtocol = &v
	return s
}

func (s *DescribeGroupResponseBody) SetPlayDomain(v string) *DescribeGroupResponseBody {
	s.PlayDomain = &v
	return s
}

func (s *DescribeGroupResponseBody) SetPushDomain(v string) *DescribeGroupResponseBody {
	s.PushDomain = &v
	return s
}

func (s *DescribeGroupResponseBody) SetRegion(v string) *DescribeGroupResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeGroupResponseBody) SetRequestId(v string) *DescribeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupResponseBody) SetStats(v *DescribeGroupResponseBodyStats) *DescribeGroupResponseBody {
	s.Stats = v
	return s
}

func (s *DescribeGroupResponseBody) SetStatus(v string) *DescribeGroupResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeGroupResponseBody) Validate() error {
	if s.Stats != nil {
		if err := s.Stats.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGroupResponseBodyStats struct {
	// The total number of devices in the group.
	//
	// example:
	//
	// 200
	DeviceNum *int64 `json:"DeviceNum,omitempty" xml:"DeviceNum,omitempty"`
	// The number of intelligent electronic devices (IEDs) in the group.
	//
	// example:
	//
	// 0
	IedNum *int64 `json:"IedNum,omitempty" xml:"IedNum,omitempty"`
	// The number of cameras in the group.
	//
	// example:
	//
	// 100
	IpcNum *int64 `json:"IpcNum,omitempty" xml:"IpcNum,omitempty"`
	// The number of platforms in the group.
	//
	// example:
	//
	// 100
	PlatformNum *int64 `json:"PlatformNum,omitempty" xml:"PlatformNum,omitempty"`
}

func (s DescribeGroupResponseBodyStats) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupResponseBodyStats) GoString() string {
	return s.String()
}

func (s *DescribeGroupResponseBodyStats) GetDeviceNum() *int64 {
	return s.DeviceNum
}

func (s *DescribeGroupResponseBodyStats) GetIedNum() *int64 {
	return s.IedNum
}

func (s *DescribeGroupResponseBodyStats) GetIpcNum() *int64 {
	return s.IpcNum
}

func (s *DescribeGroupResponseBodyStats) GetPlatformNum() *int64 {
	return s.PlatformNum
}

func (s *DescribeGroupResponseBodyStats) SetDeviceNum(v int64) *DescribeGroupResponseBodyStats {
	s.DeviceNum = &v
	return s
}

func (s *DescribeGroupResponseBodyStats) SetIedNum(v int64) *DescribeGroupResponseBodyStats {
	s.IedNum = &v
	return s
}

func (s *DescribeGroupResponseBodyStats) SetIpcNum(v int64) *DescribeGroupResponseBodyStats {
	s.IpcNum = &v
	return s
}

func (s *DescribeGroupResponseBodyStats) SetPlatformNum(v int64) *DescribeGroupResponseBodyStats {
	s.PlatformNum = &v
	return s
}

func (s *DescribeGroupResponseBodyStats) Validate() error {
	return dara.Validate(s)
}
