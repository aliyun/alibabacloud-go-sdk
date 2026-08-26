// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotLiveRtcStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *HotLiveRtcStreamRequest
	GetAppName() *string
	SetAudioMsid(v string) *HotLiveRtcStreamRequest
	GetAudioMsid() *string
	SetConnectionTimeout(v string) *HotLiveRtcStreamRequest
	GetConnectionTimeout() *string
	SetDomainName(v string) *HotLiveRtcStreamRequest
	GetDomainName() *string
	SetMediaTimeout(v string) *HotLiveRtcStreamRequest
	GetMediaTimeout() *string
	SetOwnerId(v int64) *HotLiveRtcStreamRequest
	GetOwnerId() *int64
	SetRegionCode(v string) *HotLiveRtcStreamRequest
	GetRegionCode() *string
	SetRegionId(v string) *HotLiveRtcStreamRequest
	GetRegionId() *string
	SetStreamName(v string) *HotLiveRtcStreamRequest
	GetStreamName() *string
	SetVideoMsid(v string) *HotLiveRtcStreamRequest
	GetVideoMsid() *string
}

type HotLiveRtcStreamRequest struct {
	// The application name of the live stream to prefetch.
	//
	// This parameter is required.
	//
	// example:
	//
	// app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The audio Msid.
	//
	// This parameter is required.
	//
	// example:
	//
	// rts audio
	AudioMsid *string `json:"AudioMsid,omitempty" xml:"AudioMsid,omitempty"`
	// The duration to maintain the prefetch connection. Unit: milliseconds. The default value, 0, means the connection is always maintained.
	//
	// example:
	//
	// 0
	ConnectionTimeout *string `json:"ConnectionTimeout,omitempty" xml:"ConnectionTimeout,omitempty"`
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The custom timeout period for a timeout event. Unit: milliseconds.
	//
	// example:
	//
	// 100000
	MediaTimeout *string `json:"MediaTimeout,omitempty" xml:"MediaTimeout,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The prefetch area. For more information, see the RegionCode lookup table.
	//
	// > For regions within China, specify the corresponding code from the "Region codes for China" table. For all other regions, specify the country code.
	//
	// >
	//
	// > - If the CodeRegionHasNoNode error is returned after you specify a RegionCode, the corresponding area is not covered by L1 nodes and cannot be prefetched. In this case, specify a different RegionCode.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZHJ
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the live stream to prefetch.
	//
	// This parameter is required.
	//
	// example:
	//
	// stream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The video Msid.
	//
	// This parameter is required.
	//
	// example:
	//
	// rts video
	VideoMsid *string `json:"VideoMsid,omitempty" xml:"VideoMsid,omitempty"`
}

func (s HotLiveRtcStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s HotLiveRtcStreamRequest) GoString() string {
	return s.String()
}

func (s *HotLiveRtcStreamRequest) GetAppName() *string {
	return s.AppName
}

func (s *HotLiveRtcStreamRequest) GetAudioMsid() *string {
	return s.AudioMsid
}

func (s *HotLiveRtcStreamRequest) GetConnectionTimeout() *string {
	return s.ConnectionTimeout
}

func (s *HotLiveRtcStreamRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *HotLiveRtcStreamRequest) GetMediaTimeout() *string {
	return s.MediaTimeout
}

func (s *HotLiveRtcStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *HotLiveRtcStreamRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *HotLiveRtcStreamRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *HotLiveRtcStreamRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *HotLiveRtcStreamRequest) GetVideoMsid() *string {
	return s.VideoMsid
}

func (s *HotLiveRtcStreamRequest) SetAppName(v string) *HotLiveRtcStreamRequest {
	s.AppName = &v
	return s
}

func (s *HotLiveRtcStreamRequest) SetAudioMsid(v string) *HotLiveRtcStreamRequest {
	s.AudioMsid = &v
	return s
}

func (s *HotLiveRtcStreamRequest) SetConnectionTimeout(v string) *HotLiveRtcStreamRequest {
	s.ConnectionTimeout = &v
	return s
}

func (s *HotLiveRtcStreamRequest) SetDomainName(v string) *HotLiveRtcStreamRequest {
	s.DomainName = &v
	return s
}

func (s *HotLiveRtcStreamRequest) SetMediaTimeout(v string) *HotLiveRtcStreamRequest {
	s.MediaTimeout = &v
	return s
}

func (s *HotLiveRtcStreamRequest) SetOwnerId(v int64) *HotLiveRtcStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *HotLiveRtcStreamRequest) SetRegionCode(v string) *HotLiveRtcStreamRequest {
	s.RegionCode = &v
	return s
}

func (s *HotLiveRtcStreamRequest) SetRegionId(v string) *HotLiveRtcStreamRequest {
	s.RegionId = &v
	return s
}

func (s *HotLiveRtcStreamRequest) SetStreamName(v string) *HotLiveRtcStreamRequest {
	s.StreamName = &v
	return s
}

func (s *HotLiveRtcStreamRequest) SetVideoMsid(v string) *HotLiveRtcStreamRequest {
	s.VideoMsid = &v
	return s
}

func (s *HotLiveRtcStreamRequest) Validate() error {
	return dara.Validate(s)
}
