// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterStreamUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DescribeCasterStreamUrlResponseBody
	GetCasterId() *string
	SetCasterStreams(v *DescribeCasterStreamUrlResponseBodyCasterStreams) *DescribeCasterStreamUrlResponseBody
	GetCasterStreams() *DescribeCasterStreamUrlResponseBodyCasterStreams
	SetRequestId(v string) *DescribeCasterStreamUrlResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeCasterStreamUrlResponseBody
	GetTotal() *int32
}

type DescribeCasterStreamUrlResponseBody struct {
	// The ID of the production studio.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The information about the streams of the production studio.
	CasterStreams *DescribeCasterStreamUrlResponseBodyCasterStreams `json:"CasterStreams,omitempty" xml:"CasterStreams,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of streams that were returned.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCasterStreamUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterStreamUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCasterStreamUrlResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeCasterStreamUrlResponseBody) GetCasterStreams() *DescribeCasterStreamUrlResponseBodyCasterStreams {
	return s.CasterStreams
}

func (s *DescribeCasterStreamUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCasterStreamUrlResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeCasterStreamUrlResponseBody) SetCasterId(v string) *DescribeCasterStreamUrlResponseBody {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterStreamUrlResponseBody) SetCasterStreams(v *DescribeCasterStreamUrlResponseBodyCasterStreams) *DescribeCasterStreamUrlResponseBody {
	s.CasterStreams = v
	return s
}

func (s *DescribeCasterStreamUrlResponseBody) SetRequestId(v string) *DescribeCasterStreamUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterStreamUrlResponseBody) SetTotal(v int32) *DescribeCasterStreamUrlResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeCasterStreamUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterStreamUrlResponseBodyCasterStreams struct {
	CasterStream []*DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream `json:"CasterStream,omitempty" xml:"CasterStream,omitempty" type:"Repeated"`
}

func (s DescribeCasterStreamUrlResponseBodyCasterStreams) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterStreamUrlResponseBodyCasterStreams) GoString() string {
	return s.String()
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreams) GetCasterStream() []*DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream {
	return s.CasterStream
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreams) SetCasterStream(v []*DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream) *DescribeCasterStreamUrlResponseBodyCasterStreams {
	s.CasterStream = v
	return s
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreams) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream struct {
	// Indicates whether the output stream is in preview mode or program mode.
	//
	// 	- **0**: indicates that the output videos of the scene are in preview mode.
	//
	// 	- **1**: indicates that the output videos of the scene are in program mode.
	//
	// example:
	//
	// 1
	OutputType *int32 `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	// The Real Time Messaging Protocol (RTMP) URL.
	//
	// example:
	//
	// rtmp://live/caster/example.edu
	RtmpUrl *string `json:"RtmpUrl,omitempty" xml:"RtmpUrl,omitempty"`
	// The ID of the scene.
	//
	// example:
	//
	// 23ca74e0-aca3-4e7a-8561-9d96f525****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The information about the stream.
	StreamInfos *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfos `json:"StreamInfos,omitempty" xml:"StreamInfos,omitempty" type:"Struct"`
	// The streaming URL.
	//
	// example:
	//
	// http://live/caster/example.org
	StreamUrl *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty"`
}

func (s DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream) GoString() string {
	return s.String()
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream) GetOutputType() *int32 {
	return s.OutputType
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream) GetRtmpUrl() *string {
	return s.RtmpUrl
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream) GetSceneId() *string {
	return s.SceneId
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream) GetStreamInfos() *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfos {
	return s.StreamInfos
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream) GetStreamUrl() *string {
	return s.StreamUrl
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream) SetOutputType(v int32) *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream {
	s.OutputType = &v
	return s
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream) SetRtmpUrl(v string) *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream {
	s.RtmpUrl = &v
	return s
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream) SetSceneId(v string) *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream {
	s.SceneId = &v
	return s
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream) SetStreamInfos(v *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfos) *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream {
	s.StreamInfos = v
	return s
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream) SetStreamUrl(v string) *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream {
	s.StreamUrl = &v
	return s
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStream) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfos struct {
	StreamInfo []*DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfosStreamInfo `json:"StreamInfo,omitempty" xml:"StreamInfo,omitempty" type:"Repeated"`
}

func (s DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfos) GoString() string {
	return s.String()
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfos) GetStreamInfo() []*DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfosStreamInfo {
	return s.StreamInfo
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfos) SetStreamInfo(v []*DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfosStreamInfo) *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfos {
	s.StreamInfo = v
	return s
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfosStreamInfo struct {
	// The streaming URL.
	//
	// example:
	//
	// http://out/caster/example.net
	OutputStreamUrl *string `json:"OutputStreamUrl,omitempty" xml:"OutputStreamUrl,omitempty"`
	// The resolution to which the scene transcodes the stream for playback. Valid values:
	//
	// 	- **lsd**: standard definition.
	//
	// 	- **lld**: low definition.
	//
	// 	- **lud**: ultra high definition.
	//
	// 	- **lhd**: high definition.
	//
	// example:
	//
	// lld
	TranscodeConfig *string `json:"TranscodeConfig,omitempty" xml:"TranscodeConfig,omitempty"`
	// The format to which the scene transcodes the stream for playback. Valid values:
	//
	// 	- **flv**.
	//
	// 	- **rtmp**.
	//
	// 	- **m3u8**.
	//
	// example:
	//
	// flv
	VideoFormat *string `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
}

func (s DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfosStreamInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfosStreamInfo) GoString() string {
	return s.String()
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfosStreamInfo) GetOutputStreamUrl() *string {
	return s.OutputStreamUrl
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfosStreamInfo) GetTranscodeConfig() *string {
	return s.TranscodeConfig
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfosStreamInfo) GetVideoFormat() *string {
	return s.VideoFormat
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfosStreamInfo) SetOutputStreamUrl(v string) *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfosStreamInfo {
	s.OutputStreamUrl = &v
	return s
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfosStreamInfo) SetTranscodeConfig(v string) *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfosStreamInfo {
	s.TranscodeConfig = &v
	return s
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfosStreamInfo) SetVideoFormat(v string) *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfosStreamInfo {
	s.VideoFormat = &v
	return s
}

func (s *DescribeCasterStreamUrlResponseBodyCasterStreamsCasterStreamStreamInfosStreamInfo) Validate() error {
	return dara.Validate(s)
}
