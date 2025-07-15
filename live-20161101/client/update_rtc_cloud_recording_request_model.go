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
	MixLayoutParams *UpdateRtcCloudRecordingRequestMixLayoutParams `json:"MixLayoutParams,omitempty" xml:"MixLayoutParams,omitempty" type:"Struct"`
	// This parameter is required.
	SubscribeParams *UpdateRtcCloudRecordingRequestSubscribeParams `json:"SubscribeParams,omitempty" xml:"SubscribeParams,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type UpdateRtcCloudRecordingRequestMixLayoutParams struct {
	MixBackground *UpdateRtcCloudRecordingRequestMixLayoutParamsMixBackground `json:"MixBackground,omitempty" xml:"MixBackground,omitempty" type:"Struct"`
	UserPanes     []*UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanes   `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type UpdateRtcCloudRecordingRequestMixLayoutParamsMixBackground struct {
	// example:
	//
	// 0
	RenderMode *int32 `json:"RenderMode,omitempty" xml:"RenderMode,omitempty"`
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
	// example:
	//
	// 0.5
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0
	SourceType    *int32                                                               `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SubBackground *UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground `json:"SubBackground,omitempty" xml:"SubBackground,omitempty" type:"Struct"`
	// example:
	//
	// userA
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 0.5
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 0
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
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
	return dara.Validate(s)
}

type UpdateRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground struct {
	// example:
	//
	// 0
	RenderMode *int32 `json:"RenderMode,omitempty" xml:"RenderMode,omitempty"`
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
	return dara.Validate(s)
}

type UpdateRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList struct {
	// example:
	//
	// 0
	SourceType *int32 `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
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
