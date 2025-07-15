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
	// The latency of FLV-based playback. Unit: seconds.
	//
	// >  If this parameter is left empty, the latency is set to a value corresponding to the FlvLevel parameter.
	//
	// example:
	//
	// 8
	FlvDelay *int32 `json:"FlvDelay,omitempty" xml:"FlvDelay,omitempty"`
	// The latency level of FLV-based playback. Ignore this parameter if the FlvDelay parameter is configured.
	//
	// Valid values:
	//
	// 	- **short*	- (default): The latency is 4 seconds.
	//
	// 	- **medium**: The latency is 8 seconds.
	//
	// 	- **long**: The latency is 16 seconds.
	//
	// >  If both the FlvDelay and FlvLevel parameters are left empty, FlvLevel is set to **short*	- by default.
	//
	// example:
	//
	// medium
	FlvLevel *string `json:"FlvLevel,omitempty" xml:"FlvLevel,omitempty"`
	// The latency of HLS-based playback. Unit: seconds.
	//
	// >  If this parameter is left empty, the latency is set to a value corresponding to the HlsLevel parameter.
	//
	// example:
	//
	// 4
	HlsDelay *int32 `json:"HlsDelay,omitempty" xml:"HlsDelay,omitempty"`
	// The latency level of HLS-based playback. Ignore this parameter if the HlsDelay parameter is configured.
	//
	// Valid values:
	//
	// 	- **short**: The latency is 3 seconds. This is the default value.
	//
	// 	- **medium**: The latency is 6 seconds.
	//
	// 	- **long**: The latency is 15 seconds.
	//
	// >  If both the HlsDelay and HlsLevel parameters are left empty, HlsLevel is set to **short*	- by default.
	//
	// example:
	//
	// short
	HlsLevel *string `json:"HlsLevel,omitempty" xml:"HlsLevel,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The latency of RTMP-based playback. Unit: seconds.
	//
	// >  If this parameter is left empty, the latency is set to a value corresponding to the RtmpLevel parameter.
	//
	// example:
	//
	// 4
	RtmpDelay *int32 `json:"RtmpDelay,omitempty" xml:"RtmpDelay,omitempty"`
	// The latency level of RTMP-based playback. Ignore this parameter if the RtmpDelay parameter is configured.
	//
	// Valid values:
	//
	// 	- **short*	- (default): The latency is 4 seconds.
	//
	// 	- **medium**: The latency is 8 seconds.
	//
	// 	- **long**: The latency is 16 seconds.
	//
	// >  If both the RtmpDelay and RtmpLevel parameters are left empty, RtmpLevel is set to **short*	- by default.
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
