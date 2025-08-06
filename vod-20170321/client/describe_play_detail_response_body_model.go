// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBaseInfos(v []*DescribePlayDetailResponseBodyBaseInfos) *DescribePlayDetailResponseBody
	GetBaseInfos() []*DescribePlayDetailResponseBodyBaseInfos
	SetRequestId(v string) *DescribePlayDetailResponseBody
	GetRequestId() *string
}

type DescribePlayDetailResponseBody struct {
	BaseInfos []*DescribePlayDetailResponseBodyBaseInfos `json:"BaseInfos,omitempty" xml:"BaseInfos,omitempty" type:"Repeated"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlayDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlayDetailResponseBody) GetBaseInfos() []*DescribePlayDetailResponseBodyBaseInfos {
	return s.BaseInfos
}

func (s *DescribePlayDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlayDetailResponseBody) SetBaseInfos(v []*DescribePlayDetailResponseBodyBaseInfos) *DescribePlayDetailResponseBody {
	s.BaseInfos = v
	return s
}

func (s *DescribePlayDetailResponseBody) SetRequestId(v string) *DescribePlayDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlayDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePlayDetailResponseBodyBaseInfos struct {
	AppName               *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Bps                   *string `json:"Bps,omitempty" xml:"Bps,omitempty"`
	BroadcastPace         *string `json:"BroadcastPace,omitempty" xml:"BroadcastPace,omitempty"`
	ClientIP              *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	Codec                 *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	DecodeStuckTime       *string `json:"DecodeStuckTime,omitempty" xml:"DecodeStuckTime,omitempty"`
	Definition            *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	FirstFrameDuration    *string `json:"FirstFrameDuration,omitempty" xml:"FirstFrameDuration,omitempty"`
	Fps                   *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	IsHardDecode          *string `json:"IsHardDecode,omitempty" xml:"IsHardDecode,omitempty"`
	Mdat                  *string `json:"Mdat,omitempty" xml:"Mdat,omitempty"`
	Moov                  *string `json:"Moov,omitempty" xml:"Moov,omitempty"`
	Network               *string `json:"Network,omitempty" xml:"Network,omitempty"`
	NetworkDuration       *string `json:"NetworkDuration,omitempty" xml:"NetworkDuration,omitempty"`
	NetworkStuckTime      *string `json:"NetworkStuckTime,omitempty" xml:"NetworkStuckTime,omitempty"`
	Os                    *string `json:"Os,omitempty" xml:"Os,omitempty"`
	PlayTs                *string `json:"PlayTs,omitempty" xml:"PlayTs,omitempty"`
	PlayerLoadDuration    *string `json:"PlayerLoadDuration,omitempty" xml:"PlayerLoadDuration,omitempty"`
	PlayerPreDealDuration *string `json:"PlayerPreDealDuration,omitempty" xml:"PlayerPreDealDuration,omitempty"`
	PlayerReadyDuration   *string `json:"PlayerReadyDuration,omitempty" xml:"PlayerReadyDuration,omitempty"`
	SdkVersion            *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	SessionId             *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TerminalType          *string `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
}

func (s DescribePlayDetailResponseBodyBaseInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayDetailResponseBodyBaseInfos) GoString() string {
	return s.String()
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetAppName() *string {
	return s.AppName
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetBps() *string {
	return s.Bps
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetBroadcastPace() *string {
	return s.BroadcastPace
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetClientIP() *string {
	return s.ClientIP
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetCodec() *string {
	return s.Codec
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetDecodeStuckTime() *string {
	return s.DecodeStuckTime
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetDefinition() *string {
	return s.Definition
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetFirstFrameDuration() *string {
	return s.FirstFrameDuration
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetFps() *string {
	return s.Fps
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetIsHardDecode() *string {
	return s.IsHardDecode
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetMdat() *string {
	return s.Mdat
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetMoov() *string {
	return s.Moov
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetNetwork() *string {
	return s.Network
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetNetworkDuration() *string {
	return s.NetworkDuration
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetNetworkStuckTime() *string {
	return s.NetworkStuckTime
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetOs() *string {
	return s.Os
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetPlayTs() *string {
	return s.PlayTs
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetPlayerLoadDuration() *string {
	return s.PlayerLoadDuration
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetPlayerPreDealDuration() *string {
	return s.PlayerPreDealDuration
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetPlayerReadyDuration() *string {
	return s.PlayerReadyDuration
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetSdkVersion() *string {
	return s.SdkVersion
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetStatus() *string {
	return s.Status
}

func (s *DescribePlayDetailResponseBodyBaseInfos) GetTerminalType() *string {
	return s.TerminalType
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetAppName(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.AppName = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetBps(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.Bps = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetBroadcastPace(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.BroadcastPace = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetClientIP(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.ClientIP = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetCodec(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.Codec = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetDecodeStuckTime(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.DecodeStuckTime = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetDefinition(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.Definition = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetFirstFrameDuration(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.FirstFrameDuration = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetFps(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.Fps = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetIsHardDecode(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.IsHardDecode = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetMdat(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.Mdat = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetMoov(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.Moov = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetNetwork(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.Network = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetNetworkDuration(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.NetworkDuration = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetNetworkStuckTime(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.NetworkStuckTime = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetOs(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.Os = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetPlayTs(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.PlayTs = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetPlayerLoadDuration(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.PlayerLoadDuration = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetPlayerPreDealDuration(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.PlayerPreDealDuration = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetPlayerReadyDuration(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.PlayerReadyDuration = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetSdkVersion(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.SdkVersion = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetSessionId(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.SessionId = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetStatus(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.Status = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) SetTerminalType(v string) *DescribePlayDetailResponseBodyBaseInfos {
	s.TerminalType = &v
	return s
}

func (s *DescribePlayDetailResponseBodyBaseInfos) Validate() error {
	return dara.Validate(s)
}
