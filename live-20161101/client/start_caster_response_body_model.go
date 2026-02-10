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
	PgmSceneInfos *StartCasterResponseBodyPgmSceneInfos `json:"PgmSceneInfos,omitempty" xml:"PgmSceneInfos,omitempty" type:"Struct"`
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
	if s.PgmSceneInfos != nil {
		if err := s.PgmSceneInfos.Validate(); err != nil {
			return err
		}
	}
	if s.PvwSceneInfos != nil {
		if err := s.PvwSceneInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.SceneInfo != nil {
		for _, item := range s.SceneInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type StartCasterResponseBodyPgmSceneInfosSceneInfo struct {
	SceneId     *string                                                   `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	StreamInfos *StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfos `json:"StreamInfos,omitempty" xml:"StreamInfos,omitempty" type:"Struct"`
	StreamUrl   *string                                                   `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty"`
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
	if s.StreamInfos != nil {
		if err := s.StreamInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.StreamInfo != nil {
		for _, item := range s.StreamInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type StartCasterResponseBodyPgmSceneInfosSceneInfoStreamInfosStreamInfo struct {
	OutputStreamUrl *string `json:"OutputStreamUrl,omitempty" xml:"OutputStreamUrl,omitempty"`
	TranscodeConfig *string `json:"TranscodeConfig,omitempty" xml:"TranscodeConfig,omitempty"`
	VideoFormat     *string `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
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
	if s.SceneInfo != nil {
		for _, item := range s.SceneInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type StartCasterResponseBodyPvwSceneInfosSceneInfo struct {
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
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
