// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentList(v []*GetAgentListResponseBodyAgentList) *GetAgentListResponseBody
	GetAgentList() []*GetAgentListResponseBodyAgentList
	SetRequestId(v string) *GetAgentListResponseBody
	GetRequestId() *string
}

type GetAgentListResponseBody struct {
	AgentList []*GetAgentListResponseBodyAgentList `json:"AgentList,omitempty" xml:"AgentList,omitempty" type:"Repeated"`
	// example:
	//
	// 1B217656-2267-4FBF-B26C-CDD201BDC3B8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAgentListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentListResponseBody) GetAgentList() []*GetAgentListResponseBodyAgentList {
	return s.AgentList
}

func (s *GetAgentListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentListResponseBody) SetAgentList(v []*GetAgentListResponseBodyAgentList) *GetAgentListResponseBody {
	s.AgentList = v
	return s
}

func (s *GetAgentListResponseBody) SetRequestId(v string) *GetAgentListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentListResponseBody) Validate() error {
	if s.AgentList != nil {
		for _, item := range s.AgentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAgentListResponseBodyAgentList struct {
	// example:
	//
	// rmagent_2ze68o4fden9o1s****
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// Windows
	AgentOs *string `json:"AgentOs,omitempty" xml:"AgentOs,omitempty"`
	// example:
	//
	// 49046
	AgentPort *string `json:"AgentPort,omitempty" xml:"AgentPort,omitempty"`
	// example:
	//
	// 1
	AgentStatus *int32 `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// example:
	//
	// 3.2.5.3-1
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// example:
	//
	// i-2zel0147166****
	EcsId *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	// example:
	//
	// 2020-09-29 11:05:22
	FirstLoginTime *string `json:"FirstLoginTime,omitempty" xml:"FirstLoginTime,omitempty"`
	// example:
	//
	// 2020-09-29 15:50:22
	LastActiveTime *string `json:"LastActiveTime,omitempty" xml:"LastActiveTime,omitempty"`
	OsCpu          *int32  `json:"OsCpu,omitempty" xml:"OsCpu,omitempty"`
	OsMem          *int32  `json:"OsMem,omitempty" xml:"OsMem,omitempty"`
	// example:
	//
	// 5
	PktLoss *int32 `json:"PktLoss,omitempty" xml:"PktLoss,omitempty"`
	// example:
	//
	// 192.168.XX.XX
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// example:
	//
	// 1
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// example:
	//
	// 5
	RmagentCpu *int32 `json:"RmagentCpu,omitempty" xml:"RmagentCpu,omitempty"`
	// example:
	//
	// 5
	RmagentMem *int32 `json:"RmagentMem,omitempty" xml:"RmagentMem,omitempty"`
	// example:
	//
	// 1000000
	SendByteCount *int64 `json:"SendByteCount,omitempty" xml:"SendByteCount,omitempty"`
	// example:
	//
	// 100000
	SendBytes *int64 `json:"SendBytes,omitempty" xml:"SendBytes,omitempty"`
	// example:
	//
	// 1000000
	SendPacketCount *int64 `json:"SendPacketCount,omitempty" xml:"SendPacketCount,omitempty"`
	// example:
	//
	// 100000
	SendPackets *int64                                   `json:"SendPackets,omitempty" xml:"SendPackets,omitempty"`
	SysConfig   *string                                  `json:"SysConfig,omitempty" xml:"SysConfig,omitempty"`
	Tags        []*GetAgentListResponseBodyAgentListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// v-asd2rrddgs****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetAgentListResponseBodyAgentList) String() string {
	return dara.Prettify(s)
}

func (s GetAgentListResponseBodyAgentList) GoString() string {
	return s.String()
}

func (s *GetAgentListResponseBodyAgentList) GetAgentId() *string {
	return s.AgentId
}

func (s *GetAgentListResponseBodyAgentList) GetAgentOs() *string {
	return s.AgentOs
}

func (s *GetAgentListResponseBodyAgentList) GetAgentPort() *string {
	return s.AgentPort
}

func (s *GetAgentListResponseBodyAgentList) GetAgentStatus() *int32 {
	return s.AgentStatus
}

func (s *GetAgentListResponseBodyAgentList) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *GetAgentListResponseBodyAgentList) GetEcsId() *string {
	return s.EcsId
}

func (s *GetAgentListResponseBodyAgentList) GetFirstLoginTime() *string {
	return s.FirstLoginTime
}

func (s *GetAgentListResponseBodyAgentList) GetLastActiveTime() *string {
	return s.LastActiveTime
}

func (s *GetAgentListResponseBodyAgentList) GetOsCpu() *int32 {
	return s.OsCpu
}

func (s *GetAgentListResponseBodyAgentList) GetOsMem() *int32 {
	return s.OsMem
}

func (s *GetAgentListResponseBodyAgentList) GetPktLoss() *int32 {
	return s.PktLoss
}

func (s *GetAgentListResponseBodyAgentList) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *GetAgentListResponseBodyAgentList) GetPublicIp() *string {
	return s.PublicIp
}

func (s *GetAgentListResponseBodyAgentList) GetRmagentCpu() *int32 {
	return s.RmagentCpu
}

func (s *GetAgentListResponseBodyAgentList) GetRmagentMem() *int32 {
	return s.RmagentMem
}

func (s *GetAgentListResponseBodyAgentList) GetSendByteCount() *int64 {
	return s.SendByteCount
}

func (s *GetAgentListResponseBodyAgentList) GetSendBytes() *int64 {
	return s.SendBytes
}

func (s *GetAgentListResponseBodyAgentList) GetSendPacketCount() *int64 {
	return s.SendPacketCount
}

func (s *GetAgentListResponseBodyAgentList) GetSendPackets() *int64 {
	return s.SendPackets
}

func (s *GetAgentListResponseBodyAgentList) GetSysConfig() *string {
	return s.SysConfig
}

func (s *GetAgentListResponseBodyAgentList) GetTags() []*GetAgentListResponseBodyAgentListTags {
	return s.Tags
}

func (s *GetAgentListResponseBodyAgentList) GetVpcId() *string {
	return s.VpcId
}

func (s *GetAgentListResponseBodyAgentList) SetAgentId(v string) *GetAgentListResponseBodyAgentList {
	s.AgentId = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetAgentOs(v string) *GetAgentListResponseBodyAgentList {
	s.AgentOs = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetAgentPort(v string) *GetAgentListResponseBodyAgentList {
	s.AgentPort = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetAgentStatus(v int32) *GetAgentListResponseBodyAgentList {
	s.AgentStatus = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetAgentVersion(v string) *GetAgentListResponseBodyAgentList {
	s.AgentVersion = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetEcsId(v string) *GetAgentListResponseBodyAgentList {
	s.EcsId = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetFirstLoginTime(v string) *GetAgentListResponseBodyAgentList {
	s.FirstLoginTime = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetLastActiveTime(v string) *GetAgentListResponseBodyAgentList {
	s.LastActiveTime = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetOsCpu(v int32) *GetAgentListResponseBodyAgentList {
	s.OsCpu = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetOsMem(v int32) *GetAgentListResponseBodyAgentList {
	s.OsMem = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetPktLoss(v int32) *GetAgentListResponseBodyAgentList {
	s.PktLoss = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetPrivateIp(v string) *GetAgentListResponseBodyAgentList {
	s.PrivateIp = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetPublicIp(v string) *GetAgentListResponseBodyAgentList {
	s.PublicIp = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetRmagentCpu(v int32) *GetAgentListResponseBodyAgentList {
	s.RmagentCpu = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetRmagentMem(v int32) *GetAgentListResponseBodyAgentList {
	s.RmagentMem = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetSendByteCount(v int64) *GetAgentListResponseBodyAgentList {
	s.SendByteCount = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetSendBytes(v int64) *GetAgentListResponseBodyAgentList {
	s.SendBytes = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetSendPacketCount(v int64) *GetAgentListResponseBodyAgentList {
	s.SendPacketCount = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetSendPackets(v int64) *GetAgentListResponseBodyAgentList {
	s.SendPackets = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetSysConfig(v string) *GetAgentListResponseBodyAgentList {
	s.SysConfig = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetTags(v []*GetAgentListResponseBodyAgentListTags) *GetAgentListResponseBodyAgentList {
	s.Tags = v
	return s
}

func (s *GetAgentListResponseBodyAgentList) SetVpcId(v string) *GetAgentListResponseBodyAgentList {
	s.VpcId = &v
	return s
}

func (s *GetAgentListResponseBodyAgentList) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAgentListResponseBodyAgentListTags struct {
	TagId    *int64  `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetAgentListResponseBodyAgentListTags) String() string {
	return dara.Prettify(s)
}

func (s GetAgentListResponseBodyAgentListTags) GoString() string {
	return s.String()
}

func (s *GetAgentListResponseBodyAgentListTags) GetTagId() *int64 {
	return s.TagId
}

func (s *GetAgentListResponseBodyAgentListTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetAgentListResponseBodyAgentListTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetAgentListResponseBodyAgentListTags) SetTagId(v int64) *GetAgentListResponseBodyAgentListTags {
	s.TagId = &v
	return s
}

func (s *GetAgentListResponseBodyAgentListTags) SetTagKey(v string) *GetAgentListResponseBodyAgentListTags {
	s.TagKey = &v
	return s
}

func (s *GetAgentListResponseBodyAgentListTags) SetTagValue(v string) *GetAgentListResponseBodyAgentListTags {
	s.TagValue = &v
	return s
}

func (s *GetAgentListResponseBodyAgentListTags) Validate() error {
	return dara.Validate(s)
}
