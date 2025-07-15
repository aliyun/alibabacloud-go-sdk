// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCasterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPgmSceneInfos(v *StartCasterResponseBodyPgmSceneInfos) *StartCasterResponseBody
	GetPgmSceneInfos() *StartCasterResponseBodyPgmSceneInfos
	SetPvwSceneInfos(v *StartCasterResponseBodyPvwSceneInfos) *StartCasterResponseBody
	GetPvwSceneInfos() *StartCasterResponseBodyPvwSceneInfos
	SetRequestId(v string) *StartCasterResponseBody
	GetRequestId() *string
}

type StartCasterResponseBody struct {
	// The PGM scenes.
	PgmSceneInfos *StartCasterResponseBodyPgmSceneInfos `json:"PgmSceneInfos,omitempty" xml:"PgmSceneInfos,omitempty" type:"Struct"`
	// The PVW scenes.
	PvwSceneInfos *StartCasterResponseBodyPvwSceneInfos `json:"PvwSceneInfos,omitempty" xml:"PvwSceneInfos,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6EBD1AC4-C34D-4AE1-963E-B688A228BE31
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartCasterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartCasterResponseBody) GoString() string {
	return s.String()
}

func (s *StartCasterResponseBody) GetPgmSceneInfos() *StartCasterResponseBodyPgmSceneInfos {
	return s.PgmSceneInfos
}

func (s *StartCasterResponseBody) GetPvwSceneInfos() *StartCasterResponseBodyPvwSceneInfos {
	return s.PvwSceneInfos
}

func (s *StartCasterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartCasterResponseBody) SetPgmSceneInfos(v *StartCasterResponseBodyPgmSceneInfos) *StartCasterResponseBody {
	s.PgmSceneInfos = v
	return s
}

func (s *StartCasterResponseBody) SetPvwSceneInfos(v *StartCasterResponseBodyPvwSceneInfos) *StartCasterResponseBody {
	s.PvwSceneInfos = v
	return s
}

func (s *StartCasterResponseBody) SetRequestId(v string) *StartCasterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCasterResponseBody) Validate() error {
	return dara.Validate(s)
}

type StartCasterResponseBodyPgmSceneInfos struct {
	SceneInfo []*StartCasterResponseBodyPgmSceneInfosSceneInfo `json:"SceneInfo,omitempty" xml:"SceneInfo,omitempty" type:"Repeated"`
}

func (s StartCasterResponseBodyPgmSceneInfos) String() string {
	return dara.Prettify(s)
}

func (s StartCasterResponseBodyPgmSceneInfos) GoString() string {
	return s.String()
}

func (s *StartCasterResponseBodyPgmSceneInfos) GetSceneInfo() []*StartCasterResponseBodyPgmSceneInfosSceneInfo {
	return s.SceneInfo
}

func (s *StartCasterResponseBodyPgmSceneInfos) SetSceneInfo(v []*StartCasterResponseBodyPgmSceneInfosSceneInfo) *StartCasterResponseBodyPgmSceneInfos {
	s.SceneInfo = v
	return s
}

func (s *StartCasterResponseBodyPgmSceneInfos) Validate() error {
	return dara.Validate(s)
}

type StartCasterResponseBodyPgmSceneInfosSceneInfo struct {
	// The ID of the scene.
	//
	// example:
	//
	// b5f8c837-ceeb-424f-b30b-68e94e86****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The stream relay URLs.
	StreamInfos *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfos `json:"StreamInfos,omitempty" xml:"StreamInfos,omitempty" type:"Struct"`
	// The streaming URL of the PGM scene in the production studio. The value is not a stream relay URL.
	//
	// example:
	//
	// rtmp://abclive/caster/example.edu
	StreamUrl *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty"`
}

func (s StartCasterResponseBodyPgmSceneInfosSceneInfo) String() string {
	return dara.Prettify(s)
}

func (s StartCasterResponseBodyPgmSceneInfosSceneInfo) GoString() string {
	return s.String()
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfo) GetSceneId() *string {
	return s.SceneId
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfo) GetStreamInfos() *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfos {
	return s.StreamInfos
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfo) GetStreamUrl() *string {
	return s.StreamUrl
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfo) SetSceneId(v string) *StartCasterResponseBodyPgmSceneInfosSceneInfo {
	s.SceneId = &v
	return s
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfo) SetStreamInfos(v *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfos) *StartCasterResponseBodyPgmSceneInfosSceneInfo {
	s.StreamInfos = v
	return s
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfo) SetStreamUrl(v string) *StartCasterResponseBodyPgmSceneInfosSceneInfo {
	s.StreamUrl = &v
	return s
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfo) Validate() error {
	return dara.Validate(s)
}

type StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfos struct {
	StreamInfo []*StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo `json:"StreamInfo,omitempty" xml:"StreamInfo,omitempty" type:"Repeated"`
}

func (s StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfos) String() string {
	return dara.Prettify(s)
}

func (s StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfos) GoString() string {
	return s.String()
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfos) GetStreamInfo() []*StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo {
	return s.StreamInfo
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfos) SetStreamInfo(v []*StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo) *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfos {
	s.StreamInfo = v
	return s
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfos) Validate() error {
	return dara.Validate(s)
}

type StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo struct {
	// The URL.
	//
	// example:
	//
	// rtmp://abclive/caster/example.net
	OutputStreamUrl *string `json:"OutputStreamUrl,omitempty" xml:"OutputStreamUrl,omitempty"`
	// The transcoding configuration. Valid values:
	//
	// 	- **lsd**: standard definition
	//
	// 	- **lld**: low definition
	//
	// 	- **lud**: ultra-high definition
	//
	// 	- **lhd**: high definition
	//
	// example:
	//
	// lld
	TranscodeConfig *string `json:"TranscodeConfig,omitempty" xml:"TranscodeConfig,omitempty"`
	// The format. Valid values:
	//
	// 	- **flv**
	//
	// 	- **rtmp**
	//
	// 	- **m3u8**
	//
	// example:
	//
	// flv
	VideoFormat *string `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
}

func (s StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo) String() string {
	return dara.Prettify(s)
}

func (s StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo) GoString() string {
	return s.String()
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo) GetOutputStreamUrl() *string {
	return s.OutputStreamUrl
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo) GetTranscodeConfig() *string {
	return s.TranscodeConfig
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo) GetVideoFormat() *string {
	return s.VideoFormat
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo) SetOutputStreamUrl(v string) *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo {
	s.OutputStreamUrl = &v
	return s
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo) SetTranscodeConfig(v string) *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo {
	s.TranscodeConfig = &v
	return s
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo) SetVideoFormat(v string) *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo {
	s.VideoFormat = &v
	return s
}

func (s *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo) Validate() error {
	return dara.Validate(s)
}

type StartCasterResponseBodyPvwSceneInfos struct {
	SceneInfo []*StartCasterResponseBodyPvwSceneInfosSceneInfo `json:"SceneInfo,omitempty" xml:"SceneInfo,omitempty" type:"Repeated"`
}

func (s StartCasterResponseBodyPvwSceneInfos) String() string {
	return dara.Prettify(s)
}

func (s StartCasterResponseBodyPvwSceneInfos) GoString() string {
	return s.String()
}

func (s *StartCasterResponseBodyPvwSceneInfos) GetSceneInfo() []*StartCasterResponseBodyPvwSceneInfosSceneInfo {
	return s.SceneInfo
}

func (s *StartCasterResponseBodyPvwSceneInfos) SetSceneInfo(v []*StartCasterResponseBodyPvwSceneInfosSceneInfo) *StartCasterResponseBodyPvwSceneInfos {
	s.SceneInfo = v
	return s
}

func (s *StartCasterResponseBodyPvwSceneInfos) Validate() error {
	return dara.Validate(s)
}

type StartCasterResponseBodyPvwSceneInfosSceneInfo struct {
	// The ID of the scene.
	//
	// example:
	//
	// b5f8c837-ceeb-424f-b30b-68e94e86****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The streaming URL of the PVW scene in the production studio. The value is not a stream relay URL.
	//
	// example:
	//
	// rtmp://abclive/caster/example.net
	StreamUrl *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty"`
}

func (s StartCasterResponseBodyPvwSceneInfosSceneInfo) String() string {
	return dara.Prettify(s)
}

func (s StartCasterResponseBodyPvwSceneInfosSceneInfo) GoString() string {
	return s.String()
}

func (s *StartCasterResponseBodyPvwSceneInfosSceneInfo) GetSceneId() *string {
	return s.SceneId
}

func (s *StartCasterResponseBodyPvwSceneInfosSceneInfo) GetStreamUrl() *string {
	return s.StreamUrl
}

func (s *StartCasterResponseBodyPvwSceneInfosSceneInfo) SetSceneId(v string) *StartCasterResponseBodyPvwSceneInfosSceneInfo {
	s.SceneId = &v
	return s
}

func (s *StartCasterResponseBodyPvwSceneInfosSceneInfo) SetStreamUrl(v string) *StartCasterResponseBodyPvwSceneInfosSceneInfo {
	s.StreamUrl = &v
	return s
}

func (s *StartCasterResponseBodyPvwSceneInfosSceneInfo) Validate() error {
	return dara.Validate(s)
}
