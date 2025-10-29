// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterScenesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCasterScenesResponseBody
	GetRequestId() *string
	SetSceneList(v *DescribeCasterScenesResponseBodySceneList) *DescribeCasterScenesResponseBody
	GetSceneList() *DescribeCasterScenesResponseBodySceneList
	SetTotal(v int32) *DescribeCasterScenesResponseBody
	GetTotal() *int32
}

type DescribeCasterScenesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// The ID of the scene.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scenes.
	SceneList *DescribeCasterScenesResponseBodySceneList `json:"SceneList,omitempty" xml:"SceneList,omitempty" type:"Struct"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCasterScenesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterScenesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCasterScenesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCasterScenesResponseBody) GetSceneList() *DescribeCasterScenesResponseBodySceneList {
	return s.SceneList
}

func (s *DescribeCasterScenesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeCasterScenesResponseBody) SetRequestId(v string) *DescribeCasterScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterScenesResponseBody) SetSceneList(v *DescribeCasterScenesResponseBodySceneList) *DescribeCasterScenesResponseBody {
	s.SceneList = v
	return s
}

func (s *DescribeCasterScenesResponseBody) SetTotal(v int32) *DescribeCasterScenesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeCasterScenesResponseBody) Validate() error {
	if s.SceneList != nil {
		if err := s.SceneList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCasterScenesResponseBodySceneList struct {
	Scene []*DescribeCasterScenesResponseBodySceneListScene `json:"Scene,omitempty" xml:"Scene,omitempty" type:"Repeated"`
}

func (s DescribeCasterScenesResponseBodySceneList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterScenesResponseBodySceneList) GoString() string {
	return s.String()
}

func (s *DescribeCasterScenesResponseBodySceneList) GetScene() []*DescribeCasterScenesResponseBodySceneListScene {
	return s.Scene
}

func (s *DescribeCasterScenesResponseBodySceneList) SetScene(v []*DescribeCasterScenesResponseBodySceneListScene) *DescribeCasterScenesResponseBodySceneList {
	s.Scene = v
	return s
}

func (s *DescribeCasterScenesResponseBodySceneList) Validate() error {
	if s.Scene != nil {
		for _, item := range s.Scene {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCasterScenesResponseBodySceneListScene struct {
	// The components.
	ComponentIds *DescribeCasterScenesResponseBodySceneListSceneComponentIds `json:"ComponentIds,omitempty" xml:"ComponentIds,omitempty" type:"Struct"`
	// The ID of the layout.
	//
	// example:
	//
	// 37cb2f8b-f152-4338-b928-6704f71d****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// Indicates whether the output video is in PVW mode or PGM mode. Valid values:
	//
	// 	- **0**: in PVW mode.
	//
	// 	- **1**: in PGM mode.
	//
	// example:
	//
	// 0
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	// The ID of the scene. You can use the ID as a request parameter in the API operation that is used to modify the audio configurations of the scene, query the audio configurations of the scene, start the scene, or stop the scene.
	//
	// example:
	//
	// b5f8c837-ceeb-424f-b30b-68e94e86****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The name of the scene.
	//
	// example:
	//
	// scene1
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The status of the scene. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about the stream.
	StreamInfos *DescribeCasterScenesResponseBodySceneListSceneStreamInfos `json:"StreamInfos,omitempty" xml:"StreamInfos,omitempty" type:"Struct"`
	// The URL of the output stream.
	//
	// example:
	//
	// rtmp://developer.aliyundoc.com/caster/4a82a3d1b7f0462ea37348366201****?auth_key=1608953344-0-0-ac8c628078541d7055a170ec59a5****
	StreamUrl *string `json:"StreamUrl,omitempty" xml:"StreamUrl,omitempty"`
}

func (s DescribeCasterScenesResponseBodySceneListScene) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterScenesResponseBodySceneListScene) GoString() string {
	return s.String()
}

func (s *DescribeCasterScenesResponseBodySceneListScene) GetComponentIds() *DescribeCasterScenesResponseBodySceneListSceneComponentIds {
	return s.ComponentIds
}

func (s *DescribeCasterScenesResponseBodySceneListScene) GetLayoutId() *string {
	return s.LayoutId
}

func (s *DescribeCasterScenesResponseBodySceneListScene) GetOutputType() *string {
	return s.OutputType
}

func (s *DescribeCasterScenesResponseBodySceneListScene) GetSceneId() *string {
	return s.SceneId
}

func (s *DescribeCasterScenesResponseBodySceneListScene) GetSceneName() *string {
	return s.SceneName
}

func (s *DescribeCasterScenesResponseBodySceneListScene) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCasterScenesResponseBodySceneListScene) GetStreamInfos() *DescribeCasterScenesResponseBodySceneListSceneStreamInfos {
	return s.StreamInfos
}

func (s *DescribeCasterScenesResponseBodySceneListScene) GetStreamUrl() *string {
	return s.StreamUrl
}

func (s *DescribeCasterScenesResponseBodySceneListScene) SetComponentIds(v *DescribeCasterScenesResponseBodySceneListSceneComponentIds) *DescribeCasterScenesResponseBodySceneListScene {
	s.ComponentIds = v
	return s
}

func (s *DescribeCasterScenesResponseBodySceneListScene) SetLayoutId(v string) *DescribeCasterScenesResponseBodySceneListScene {
	s.LayoutId = &v
	return s
}

func (s *DescribeCasterScenesResponseBodySceneListScene) SetOutputType(v string) *DescribeCasterScenesResponseBodySceneListScene {
	s.OutputType = &v
	return s
}

func (s *DescribeCasterScenesResponseBodySceneListScene) SetSceneId(v string) *DescribeCasterScenesResponseBodySceneListScene {
	s.SceneId = &v
	return s
}

func (s *DescribeCasterScenesResponseBodySceneListScene) SetSceneName(v string) *DescribeCasterScenesResponseBodySceneListScene {
	s.SceneName = &v
	return s
}

func (s *DescribeCasterScenesResponseBodySceneListScene) SetStatus(v int32) *DescribeCasterScenesResponseBodySceneListScene {
	s.Status = &v
	return s
}

func (s *DescribeCasterScenesResponseBodySceneListScene) SetStreamInfos(v *DescribeCasterScenesResponseBodySceneListSceneStreamInfos) *DescribeCasterScenesResponseBodySceneListScene {
	s.StreamInfos = v
	return s
}

func (s *DescribeCasterScenesResponseBodySceneListScene) SetStreamUrl(v string) *DescribeCasterScenesResponseBodySceneListScene {
	s.StreamUrl = &v
	return s
}

func (s *DescribeCasterScenesResponseBodySceneListScene) Validate() error {
	if s.ComponentIds != nil {
		if err := s.ComponentIds.Validate(); err != nil {
			return err
		}
	}
	if s.StreamInfos != nil {
		if err := s.StreamInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCasterScenesResponseBodySceneListSceneComponentIds struct {
	ComponentId []*string `json:"componentId,omitempty" xml:"componentId,omitempty" type:"Repeated"`
}

func (s DescribeCasterScenesResponseBodySceneListSceneComponentIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterScenesResponseBodySceneListSceneComponentIds) GoString() string {
	return s.String()
}

func (s *DescribeCasterScenesResponseBodySceneListSceneComponentIds) GetComponentId() []*string {
	return s.ComponentId
}

func (s *DescribeCasterScenesResponseBodySceneListSceneComponentIds) SetComponentId(v []*string) *DescribeCasterScenesResponseBodySceneListSceneComponentIds {
	s.ComponentId = v
	return s
}

func (s *DescribeCasterScenesResponseBodySceneListSceneComponentIds) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterScenesResponseBodySceneListSceneStreamInfos struct {
	StreamInfo []*DescribeCasterScenesResponseBodySceneListSceneStreamInfosStreamInfo `json:"StreamInfo,omitempty" xml:"StreamInfo,omitempty" type:"Repeated"`
}

func (s DescribeCasterScenesResponseBodySceneListSceneStreamInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterScenesResponseBodySceneListSceneStreamInfos) GoString() string {
	return s.String()
}

func (s *DescribeCasterScenesResponseBodySceneListSceneStreamInfos) GetStreamInfo() []*DescribeCasterScenesResponseBodySceneListSceneStreamInfosStreamInfo {
	return s.StreamInfo
}

func (s *DescribeCasterScenesResponseBodySceneListSceneStreamInfos) SetStreamInfo(v []*DescribeCasterScenesResponseBodySceneListSceneStreamInfosStreamInfo) *DescribeCasterScenesResponseBodySceneListSceneStreamInfos {
	s.StreamInfo = v
	return s
}

func (s *DescribeCasterScenesResponseBodySceneListSceneStreamInfos) Validate() error {
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

type DescribeCasterScenesResponseBodySceneListSceneStreamInfosStreamInfo struct {
	// The streaming URL.
	//
	// example:
	//
	// http://live/caster/example.net
	OutputStreamUrl *string `json:"OutputStreamUrl,omitempty" xml:"OutputStreamUrl,omitempty"`
	// The transcoding configuration. Valid values:
	//
	// 	- **sd**: standard definition
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
	// 	- **mp4**
	//
	// 	- **m3u8**
	//
	// example:
	//
	// flv
	VideoFormat *string `json:"VideoFormat,omitempty" xml:"VideoFormat,omitempty"`
}

func (s DescribeCasterScenesResponseBodySceneListSceneStreamInfosStreamInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterScenesResponseBodySceneListSceneStreamInfosStreamInfo) GoString() string {
	return s.String()
}

func (s *DescribeCasterScenesResponseBodySceneListSceneStreamInfosStreamInfo) GetOutputStreamUrl() *string {
	return s.OutputStreamUrl
}

func (s *DescribeCasterScenesResponseBodySceneListSceneStreamInfosStreamInfo) GetTranscodeConfig() *string {
	return s.TranscodeConfig
}

func (s *DescribeCasterScenesResponseBodySceneListSceneStreamInfosStreamInfo) GetVideoFormat() *string {
	return s.VideoFormat
}

func (s *DescribeCasterScenesResponseBodySceneListSceneStreamInfosStreamInfo) SetOutputStreamUrl(v string) *DescribeCasterScenesResponseBodySceneListSceneStreamInfosStreamInfo {
	s.OutputStreamUrl = &v
	return s
}

func (s *DescribeCasterScenesResponseBodySceneListSceneStreamInfosStreamInfo) SetTranscodeConfig(v string) *DescribeCasterScenesResponseBodySceneListSceneStreamInfosStreamInfo {
	s.TranscodeConfig = &v
	return s
}

func (s *DescribeCasterScenesResponseBodySceneListSceneStreamInfosStreamInfo) SetVideoFormat(v string) *DescribeCasterScenesResponseBodySceneListSceneStreamInfosStreamInfo {
	s.VideoFormat = &v
	return s
}

func (s *DescribeCasterScenesResponseBodySceneListSceneStreamInfosStreamInfo) Validate() error {
	return dara.Validate(s)
}
