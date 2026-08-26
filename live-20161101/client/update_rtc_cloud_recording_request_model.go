// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtcCloudRecordingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMixLayoutParams(v *UpdateRtcCloudRecordingRequestMixLayoutParams) *UpdateRtcCloudRecordingRequest
	GetMixLayoutParams() *UpdateRtcCloudRecordingRequestMixLayoutParams
	SetSubscribeParams(v *UpdateRtcCloudRecordingRequestSubscribeParams) *UpdateRtcCloudRecordingRequest
	GetSubscribeParams() *UpdateRtcCloudRecordingRequestSubscribeParams
	SetTaskId(v string) *UpdateRtcCloudRecordingRequest
	GetTaskId() *string
}

type UpdateRtcCloudRecordingRequest struct {
	// The updated layout parameters. Leave this parameter empty in single-stream recording mode. This parameter is required in stream mixing recording mode when the transcoding output is not audio-only.
	MixLayoutParams *UpdateRtcCloudRecordingRequestMixLayoutParams `json:"MixLayoutParams,omitempty" xml:"MixLayoutParams,omitempty" type:"Struct"`
	// The updated subscription parameters.
	//
	// This parameter is required.
	SubscribeParams *UpdateRtcCloudRecordingRequestSubscribeParams `json:"SubscribeParams,omitempty" xml:"SubscribeParams,omitempty" type:"Struct"`
	// The task ID. This ID is returned by StartRtcCloudRecording. Only tasks in the running or abnormal state can be updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// ******73-8501-****-8ac1-72295a******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateRtcCloudRecordingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcCloudRecordingRequest) GoString() string {
	return s.String()
}

func (s *UpdateRtcCloudRecordingRequest) GetMixLayoutParams() *UpdateRtcCloudRecordingRequestMixLayoutParams {
	return s.MixLayoutParams
}

func (s *UpdateRtcCloudRecordingRequest) GetSubscribeParams() *UpdateRtcCloudRecordingRequestSubscribeParams {
	return s.SubscribeParams
}

func (s *UpdateRtcCloudRecordingRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateRtcCloudRecordingRequest) SetMixLayoutParams(v *UpdateRtcCloudRecordingRequestMixLayoutParams) *UpdateRtcCloudRecordingRequest {
	s.MixLayoutParams = v
	return s
}

func (s *UpdateRtcCloudRecordingRequest) SetSubscribeParams(v *UpdateRtcCloudRecordingRequestSubscribeParams) *UpdateRtcCloudRecordingRequest {
	s.SubscribeParams = v
	return s
}

func (s *UpdateRtcCloudRecordingRequest) SetTaskId(v string) *UpdateRtcCloudRecordingRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateRtcCloudRecordingRequest) Validate() error {
	if s.MixLayoutParams != nil {
		if err := s.MixLayoutParams.Validate(); err != nil {
			return err
		}
	}
	if s.SubscribeParams != nil {
		if err := s.SubscribeParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateRtcCloudRecordingRequestMixLayoutParams struct {
	// The global background image for stream mixing.
	MixBackground *UpdateRtcCloudRecordingRequestMixLayoutParamsMixBackground `json:"MixBackground,omitempty" xml:"MixBackground,omitempty" type:"Struct"`
	// The window layout information of the subscribed users. Only UserIds with layout information configured are placed in the output. This parameter is required in stream mixing mode when recording non-audio-only files.
	UserPanes []*UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
}

func (s UpdateRtcCloudRecordingRequestMixLayoutParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcCloudRecordingRequestMixLayoutParams) GoString() string {
	return s.String()
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParams) GetMixBackground() *UpdateRtcCloudRecordingRequestMixLayoutParamsMixBackground {
	return s.MixBackground
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParams) GetUserPanes() []*UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	return s.UserPanes
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParams) SetMixBackground(v *UpdateRtcCloudRecordingRequestMixLayoutParamsMixBackground) *UpdateRtcCloudRecordingRequestMixLayoutParams {
	s.MixBackground = v
	return s
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParams) SetUserPanes(v []*UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) *UpdateRtcCloudRecordingRequestMixLayoutParams {
	s.UserPanes = v
	return s
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParams) Validate() error {
	if s.MixBackground != nil {
		if err := s.MixBackground.Validate(); err != nil {
			return err
		}
	}
	if s.UserPanes != nil {
		for _, item := range s.UserPanes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateRtcCloudRecordingRequestMixLayoutParamsMixBackground struct {
	// The display mode for the output. Valid values:
	//
	// - 0: crop. (Default)
	//
	// - 1: scale and display with black borders.
	//
	// example:
	//
	// 0
	RenderMode *int32 `json:"RenderMode,omitempty" xml:"RenderMode,omitempty"`
	// The URL of the background image. The maximum length is 2048 characters.
	//
	// example:
	//
	// https://xxxx.com/photos/my-test-picture.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateRtcCloudRecordingRequestMixLayoutParamsMixBackground) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcCloudRecordingRequestMixLayoutParamsMixBackground) GoString() string {
	return s.String()
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsMixBackground) GetRenderMode() *int32 {
	return s.RenderMode
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsMixBackground) GetUrl() *string {
	return s.Url
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsMixBackground) SetRenderMode(v int32) *UpdateRtcCloudRecordingRequestMixLayoutParamsMixBackground {
	s.RenderMode = &v
	return s
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsMixBackground) SetUrl(v string) *UpdateRtcCloudRecordingRequestMixLayoutParamsMixBackground {
	s.Url = &v
	return s
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsMixBackground) Validate() error {
	return dara.Validate(s)
}

type UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes struct {
	// The pane height as a normalized percentage. The value must be in the range of [0, 1]. (Default: 0)
	//
	// example:
	//
	// 0.5
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The video input stream type of the UserId. This parameter is invalid if UserId is not specified. Valid values:
	//
	// - 0: camera. (Default)
	//
	// - 1: screen sharing.
	//
	// The combination of UserId and SourceType specified here must be included in SubscribeUserIdList.
	//
	// example:
	//
	// 0
	SourceType *int32 `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The sub-pane background image. When a user turns off the camera, has not started stream ingest after joining, or leaves the channel midway, the corresponding image is displayed at the layout position.
	SubBackground *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground `json:"SubBackground,omitempty" xml:"SubBackground,omitempty" type:"Struct"`
	// The UserId corresponding to this window.
	//
	// - If UserId is not specified, windows are filled in the order in which subscribed users join the channel.
	//
	// - The combination of UserId and SourceType specified here must be included in SubscribeUserIdList.
	//
	// - Audio-only streams cannot be added to the layout.
	//
	// example:
	//
	// userA
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The pane width as a normalized percentage. The value must be in the range of [0, 1]. (Default: 0)
	//
	// example:
	//
	// 0.5
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// The X coordinate as a normalized percentage. The value must be in the range of [0, 1]. (Default: 0)
	//
	// example:
	//
	// 0
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// The Y coordinate as a normalized percentage. The value must be in the range of [0, 1]. (Default: 0)
	//
	// example:
	//
	// 0
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
	// The stacking order. 0 is the bottom layer, layer 1 is above layer 0, and so on. (Default: 0)
	//
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) GoString() string {
	return s.String()
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) GetHeight() *string {
	return s.Height
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) GetSourceType() *int32 {
	return s.SourceType
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) GetSubBackground() *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground {
	return s.SubBackground
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) GetUserId() *string {
	return s.UserId
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) GetWidth() *string {
	return s.Width
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) GetX() *string {
	return s.X
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) GetY() *string {
	return s.Y
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) SetHeight(v string) *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	s.Height = &v
	return s
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) SetSourceType(v int32) *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	s.SourceType = &v
	return s
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) SetSubBackground(v *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground) *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	s.SubBackground = v
	return s
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) SetUserId(v string) *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	s.UserId = &v
	return s
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) SetWidth(v string) *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	s.Width = &v
	return s
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) SetX(v string) *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	s.X = &v
	return s
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) SetY(v string) *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	s.Y = &v
	return s
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) SetZOrder(v int32) *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	s.ZOrder = &v
	return s
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes) Validate() error {
	if s.SubBackground != nil {
		if err := s.SubBackground.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground struct {
	// The display mode for the sub-pane output. Valid values:
	//
	// - 0: crop. (Default)
	//
	// - 1: scale and display with black borders.
	//
	// example:
	//
	// 0
	RenderMode *int32 `json:"RenderMode,omitempty" xml:"RenderMode,omitempty"`
	// The URL of the background image. The maximum length is 2048 characters.
	//
	// example:
	//
	// https://xxxx.com/photos/my-test-pane-picture.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground) GoString() string {
	return s.String()
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground) GetRenderMode() *int32 {
	return s.RenderMode
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground) GetUrl() *string {
	return s.Url
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground) SetRenderMode(v int32) *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground {
	s.RenderMode = &v
	return s
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground) SetUrl(v string) *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground {
	s.Url = &v
	return s
}

func (s *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground) Validate() error {
	return dara.Validate(s)
}

type UpdateRtcCloudRecordingRequestSubscribeParams struct {
	// The list of subscribed UserId entries. In single-stream recording mode, each UserId is recorded separately. In stream mixing recording mode, the audio and video of all UserIds are mixed into a single set of audio and video.
	//
	// >
	//
	// > - The array supports a maximum of 17 elements.
	//
	// This parameter is required.
	SubscribeUserIdList []*UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList `json:"SubscribeUserIdList,omitempty" xml:"SubscribeUserIdList,omitempty" type:"Repeated"`
}

func (s UpdateRtcCloudRecordingRequestSubscribeParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcCloudRecordingRequestSubscribeParams) GoString() string {
	return s.String()
}

func (s *UpdateRtcCloudRecordingRequestSubscribeParams) GetSubscribeUserIdList() []*UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList {
	return s.SubscribeUserIdList
}

func (s *UpdateRtcCloudRecordingRequestSubscribeParams) SetSubscribeUserIdList(v []*UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) *UpdateRtcCloudRecordingRequestSubscribeParams {
	s.SubscribeUserIdList = v
	return s
}

func (s *UpdateRtcCloudRecordingRequestSubscribeParams) Validate() error {
	if s.SubscribeUserIdList != nil {
		for _, item := range s.SubscribeUserIdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList struct {
	// The video input stream type of the UserId. This parameter takes effect only when the video stream is subscribed (StreamType=2). Valid values:
	//
	// - 0: camera. (Default)
	//
	// - 1: screen sharing.
	//
	// example:
	//
	// 0
	SourceType *int32 `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The media type of the subscribed UserId. Valid values:
	//
	// - 0: original stream, which includes both audio and video. (Default)
	//
	// - 1: audio-only stream.
	//
	// - 2: video-only stream.
	//
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The subscribed UserId.
	//
	// This parameter is required.
	//
	// example:
	//
	// userA
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) GoString() string {
	return s.String()
}

func (s *UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) GetSourceType() *int32 {
	return s.SourceType
}

func (s *UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) GetStreamType() *int32 {
	return s.StreamType
}

func (s *UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) GetUserId() *string {
	return s.UserId
}

func (s *UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) SetSourceType(v int32) *UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList {
	s.SourceType = &v
	return s
}

func (s *UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) SetStreamType(v int32) *UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList {
	s.StreamType = &v
	return s
}

func (s *UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) SetUserId(v string) *UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList {
	s.UserId = &v
	return s
}

func (s *UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) Validate() error {
	return dara.Validate(s)
}
