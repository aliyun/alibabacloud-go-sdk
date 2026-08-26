// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveStreamDelayConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SetLiveStreamDelayConfigRequest
	GetDomainName() *string
	SetFlvDelay(v int32) *SetLiveStreamDelayConfigRequest
	GetFlvDelay() *int32
	SetFlvLevel(v string) *SetLiveStreamDelayConfigRequest
	GetFlvLevel() *string
	SetHlsDelay(v int32) *SetLiveStreamDelayConfigRequest
	GetHlsDelay() *int32
	SetHlsLevel(v string) *SetLiveStreamDelayConfigRequest
	GetHlsLevel() *string
	SetOwnerId(v int64) *SetLiveStreamDelayConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetLiveStreamDelayConfigRequest
	GetRegionId() *string
	SetRtmpDelay(v int32) *SetLiveStreamDelayConfigRequest
	GetRtmpDelay() *int32
	SetRtmpLevel(v string) *SetLiveStreamDelayConfigRequest
	GetRtmpLevel() *string
}

type SetLiveStreamDelayConfigRequest struct {
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The FLV playback latency. Unit: seconds.
	//
	// > If this parameter is left empty, the latency is set based on the value of FlvLevel.
	//
	// example:
	//
	// 8
	FlvDelay *int32 `json:"FlvDelay,omitempty" xml:"FlvDelay,omitempty"`
	// The FLV latency level. This parameter is ignored if FlvDelay is set.
	//
	// Valid values:
	//
	// - **short*	- (default): 2 seconds.
	//
	// - **medium**: 4 seconds.
	//
	// - **long**: more than 4 seconds.
	//
	// > If both FlvDelay and FlvLevel are left empty, the default value **short*	- is used.
	//
	// example:
	//
	// medium
	FlvLevel *string `json:"FlvLevel,omitempty" xml:"FlvLevel,omitempty"`
	// The HLS playback latency. Unit: seconds.
	//
	// > If this parameter is left empty, the latency is set based on the value of HlsLevel.
	//
	// example:
	//
	// 4
	HlsDelay *int32 `json:"HlsDelay,omitempty" xml:"HlsDelay,omitempty"`
	// The HLS latency level. This parameter is ignored if HlsDelay is set.
	//
	// Valid values:
	//
	// - **short*	- (default): 3 seconds.
	//
	// - **medium**: 6 seconds.
	//
	// - **long**: 15 seconds.
	//
	// > If both HlsDelay and HlsLevel are left empty, the default value **short*	- is used.
	//
	// example:
	//
	// short
	HlsLevel *string `json:"HlsLevel,omitempty" xml:"HlsLevel,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The RTMP playback latency. Unit: seconds.
	//
	// > If this parameter is left empty, the latency is set based on the value of RtmpLevel.
	//
	// example:
	//
	// 4
	RtmpDelay *int32 `json:"RtmpDelay,omitempty" xml:"RtmpDelay,omitempty"`
	// The RTMP latency level. This parameter is ignored if RtmpDelay is set.
	//
	// Valid values:
	//
	// - **short*	- (default): 2 seconds.
	//
	// - **medium**: 4 seconds.
	//
	// - **long**: more than 4 seconds.
	//
	// > If both RtmpDelay and RtmpLevel are left empty, the default value **short*	- is used.
	//
	// example:
	//
	// short
	RtmpLevel *string `json:"RtmpLevel,omitempty" xml:"RtmpLevel,omitempty"`
}

func (s SetLiveStreamDelayConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLiveStreamDelayConfigRequest) GoString() string {
	return s.String()
}

func (s *SetLiveStreamDelayConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetLiveStreamDelayConfigRequest) GetFlvDelay() *int32 {
	return s.FlvDelay
}

func (s *SetLiveStreamDelayConfigRequest) GetFlvLevel() *string {
	return s.FlvLevel
}

func (s *SetLiveStreamDelayConfigRequest) GetHlsDelay() *int32 {
	return s.HlsDelay
}

func (s *SetLiveStreamDelayConfigRequest) GetHlsLevel() *string {
	return s.HlsLevel
}

func (s *SetLiveStreamDelayConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLiveStreamDelayConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetLiveStreamDelayConfigRequest) GetRtmpDelay() *int32 {
	return s.RtmpDelay
}

func (s *SetLiveStreamDelayConfigRequest) GetRtmpLevel() *string {
	return s.RtmpLevel
}

func (s *SetLiveStreamDelayConfigRequest) SetDomainName(v string) *SetLiveStreamDelayConfigRequest {
	s.DomainName = &v
	return s
}

func (s *SetLiveStreamDelayConfigRequest) SetFlvDelay(v int32) *SetLiveStreamDelayConfigRequest {
	s.FlvDelay = &v
	return s
}

func (s *SetLiveStreamDelayConfigRequest) SetFlvLevel(v string) *SetLiveStreamDelayConfigRequest {
	s.FlvLevel = &v
	return s
}

func (s *SetLiveStreamDelayConfigRequest) SetHlsDelay(v int32) *SetLiveStreamDelayConfigRequest {
	s.HlsDelay = &v
	return s
}

func (s *SetLiveStreamDelayConfigRequest) SetHlsLevel(v string) *SetLiveStreamDelayConfigRequest {
	s.HlsLevel = &v
	return s
}

func (s *SetLiveStreamDelayConfigRequest) SetOwnerId(v int64) *SetLiveStreamDelayConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLiveStreamDelayConfigRequest) SetRegionId(v string) *SetLiveStreamDelayConfigRequest {
	s.RegionId = &v
	return s
}

func (s *SetLiveStreamDelayConfigRequest) SetRtmpDelay(v int32) *SetLiveStreamDelayConfigRequest {
	s.RtmpDelay = &v
	return s
}

func (s *SetLiveStreamDelayConfigRequest) SetRtmpLevel(v string) *SetLiveStreamDelayConfigRequest {
	s.RtmpLevel = &v
	return s
}

func (s *SetLiveStreamDelayConfigRequest) Validate() error {
	return dara.Validate(s)
}
