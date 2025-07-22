// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRecordTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartRecordTaskRequest
	GetAppId() *string
	SetChannelId(v string) *StartRecordTaskRequest
	GetChannelId() *string
	SetCropMode(v int64) *StartRecordTaskRequest
	GetCropMode() *int64
	SetLayoutIds(v []*int64) *StartRecordTaskRequest
	GetLayoutIds() []*int64
	SetMediaEncode(v int32) *StartRecordTaskRequest
	GetMediaEncode() *int32
	SetMixMode(v int32) *StartRecordTaskRequest
	GetMixMode() *int32
	SetOwnerId(v int64) *StartRecordTaskRequest
	GetOwnerId() *int64
	SetSourceType(v string) *StartRecordTaskRequest
	GetSourceType() *string
	SetStreamType(v int32) *StartRecordTaskRequest
	GetStreamType() *int32
	SetSubSpecAudioUsers(v []*string) *StartRecordTaskRequest
	GetSubSpecAudioUsers() []*string
	SetSubSpecCameraUsers(v []*string) *StartRecordTaskRequest
	GetSubSpecCameraUsers() []*string
	SetSubSpecShareScreenUsers(v []*string) *StartRecordTaskRequest
	GetSubSpecShareScreenUsers() []*string
	SetSubSpecUsers(v []*string) *StartRecordTaskRequest
	GetSubSpecUsers() []*string
	SetTaskId(v string) *StartRecordTaskRequest
	GetTaskId() *string
	SetTaskProfile(v string) *StartRecordTaskRequest
	GetTaskProfile() *string
	SetTemplateId(v string) *StartRecordTaskRequest
	GetTemplateId() *string
	SetUnsubSpecAudioUsers(v []*string) *StartRecordTaskRequest
	GetUnsubSpecAudioUsers() []*string
	SetUnsubSpecCameraUsers(v []*string) *StartRecordTaskRequest
	GetUnsubSpecCameraUsers() []*string
	SetUnsubSpecShareScreenUsers(v []*string) *StartRecordTaskRequest
	GetUnsubSpecShareScreenUsers() []*string
	SetUserPanes(v []*StartRecordTaskRequestUserPanes) *StartRecordTaskRequest
	GetUserPanes() []*StartRecordTaskRequestUserPanes
}

type StartRecordTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1
	CropMode *int64 `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	// example:
	//
	// 1111
	LayoutIds []*int64 `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// example:
	//
	// 1
	MixMode *int32 `json:"MixMode,omitempty" xml:"MixMode,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// camera
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// example:
	//
	// 1
	SubSpecAudioUsers []*string `json:"SubSpecAudioUsers,omitempty" xml:"SubSpecAudioUsers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	SubSpecCameraUsers []*string `json:"SubSpecCameraUsers,omitempty" xml:"SubSpecCameraUsers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	SubSpecShareScreenUsers []*string `json:"SubSpecShareScreenUsers,omitempty" xml:"SubSpecShareScreenUsers,omitempty" type:"Repeated"`
	// example:
	//
	// userID
	SubSpecUsers []*string `json:"SubSpecUsers,omitempty" xml:"SubSpecUsers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 4IN_1080P
	TaskProfile *string `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 76dasgb****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 1
	UnsubSpecAudioUsers []*string `json:"UnsubSpecAudioUsers,omitempty" xml:"UnsubSpecAudioUsers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	UnsubSpecCameraUsers []*string `json:"UnsubSpecCameraUsers,omitempty" xml:"UnsubSpecCameraUsers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	UnsubSpecShareScreenUsers []*string                          `json:"UnsubSpecShareScreenUsers,omitempty" xml:"UnsubSpecShareScreenUsers,omitempty" type:"Repeated"`
	UserPanes                 []*StartRecordTaskRequestUserPanes `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
}

func (s StartRecordTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRecordTaskRequest) GoString() string {
	return s.String()
}

func (s *StartRecordTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartRecordTaskRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartRecordTaskRequest) GetCropMode() *int64 {
	return s.CropMode
}

func (s *StartRecordTaskRequest) GetLayoutIds() []*int64 {
	return s.LayoutIds
}

func (s *StartRecordTaskRequest) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *StartRecordTaskRequest) GetMixMode() *int32 {
	return s.MixMode
}

func (s *StartRecordTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartRecordTaskRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *StartRecordTaskRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *StartRecordTaskRequest) GetSubSpecAudioUsers() []*string {
	return s.SubSpecAudioUsers
}

func (s *StartRecordTaskRequest) GetSubSpecCameraUsers() []*string {
	return s.SubSpecCameraUsers
}

func (s *StartRecordTaskRequest) GetSubSpecShareScreenUsers() []*string {
	return s.SubSpecShareScreenUsers
}

func (s *StartRecordTaskRequest) GetSubSpecUsers() []*string {
	return s.SubSpecUsers
}

func (s *StartRecordTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StartRecordTaskRequest) GetTaskProfile() *string {
	return s.TaskProfile
}

func (s *StartRecordTaskRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *StartRecordTaskRequest) GetUnsubSpecAudioUsers() []*string {
	return s.UnsubSpecAudioUsers
}

func (s *StartRecordTaskRequest) GetUnsubSpecCameraUsers() []*string {
	return s.UnsubSpecCameraUsers
}

func (s *StartRecordTaskRequest) GetUnsubSpecShareScreenUsers() []*string {
	return s.UnsubSpecShareScreenUsers
}

func (s *StartRecordTaskRequest) GetUserPanes() []*StartRecordTaskRequestUserPanes {
	return s.UserPanes
}

func (s *StartRecordTaskRequest) SetAppId(v string) *StartRecordTaskRequest {
	s.AppId = &v
	return s
}

func (s *StartRecordTaskRequest) SetChannelId(v string) *StartRecordTaskRequest {
	s.ChannelId = &v
	return s
}

func (s *StartRecordTaskRequest) SetCropMode(v int64) *StartRecordTaskRequest {
	s.CropMode = &v
	return s
}

func (s *StartRecordTaskRequest) SetLayoutIds(v []*int64) *StartRecordTaskRequest {
	s.LayoutIds = v
	return s
}

func (s *StartRecordTaskRequest) SetMediaEncode(v int32) *StartRecordTaskRequest {
	s.MediaEncode = &v
	return s
}

func (s *StartRecordTaskRequest) SetMixMode(v int32) *StartRecordTaskRequest {
	s.MixMode = &v
	return s
}

func (s *StartRecordTaskRequest) SetOwnerId(v int64) *StartRecordTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StartRecordTaskRequest) SetSourceType(v string) *StartRecordTaskRequest {
	s.SourceType = &v
	return s
}

func (s *StartRecordTaskRequest) SetStreamType(v int32) *StartRecordTaskRequest {
	s.StreamType = &v
	return s
}

func (s *StartRecordTaskRequest) SetSubSpecAudioUsers(v []*string) *StartRecordTaskRequest {
	s.SubSpecAudioUsers = v
	return s
}

func (s *StartRecordTaskRequest) SetSubSpecCameraUsers(v []*string) *StartRecordTaskRequest {
	s.SubSpecCameraUsers = v
	return s
}

func (s *StartRecordTaskRequest) SetSubSpecShareScreenUsers(v []*string) *StartRecordTaskRequest {
	s.SubSpecShareScreenUsers = v
	return s
}

func (s *StartRecordTaskRequest) SetSubSpecUsers(v []*string) *StartRecordTaskRequest {
	s.SubSpecUsers = v
	return s
}

func (s *StartRecordTaskRequest) SetTaskId(v string) *StartRecordTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StartRecordTaskRequest) SetTaskProfile(v string) *StartRecordTaskRequest {
	s.TaskProfile = &v
	return s
}

func (s *StartRecordTaskRequest) SetTemplateId(v string) *StartRecordTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *StartRecordTaskRequest) SetUnsubSpecAudioUsers(v []*string) *StartRecordTaskRequest {
	s.UnsubSpecAudioUsers = v
	return s
}

func (s *StartRecordTaskRequest) SetUnsubSpecCameraUsers(v []*string) *StartRecordTaskRequest {
	s.UnsubSpecCameraUsers = v
	return s
}

func (s *StartRecordTaskRequest) SetUnsubSpecShareScreenUsers(v []*string) *StartRecordTaskRequest {
	s.UnsubSpecShareScreenUsers = v
	return s
}

func (s *StartRecordTaskRequest) SetUserPanes(v []*StartRecordTaskRequestUserPanes) *StartRecordTaskRequest {
	s.UserPanes = v
	return s
}

func (s *StartRecordTaskRequest) Validate() error {
	return dara.Validate(s)
}

type StartRecordTaskRequestUserPanes struct {
	Images []*StartRecordTaskRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// example:
	//
	// camera
	SourceType *string                                 `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*StartRecordTaskRequestUserPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// example:
	//
	// TestId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StartRecordTaskRequestUserPanes) String() string {
	return dara.Prettify(s)
}

func (s StartRecordTaskRequestUserPanes) GoString() string {
	return s.String()
}

func (s *StartRecordTaskRequestUserPanes) GetImages() []*StartRecordTaskRequestUserPanesImages {
	return s.Images
}

func (s *StartRecordTaskRequestUserPanes) GetPaneId() *int32 {
	return s.PaneId
}

func (s *StartRecordTaskRequestUserPanes) GetSourceType() *string {
	return s.SourceType
}

func (s *StartRecordTaskRequestUserPanes) GetTexts() []*StartRecordTaskRequestUserPanesTexts {
	return s.Texts
}

func (s *StartRecordTaskRequestUserPanes) GetUserId() *string {
	return s.UserId
}

func (s *StartRecordTaskRequestUserPanes) SetImages(v []*StartRecordTaskRequestUserPanesImages) *StartRecordTaskRequestUserPanes {
	s.Images = v
	return s
}

func (s *StartRecordTaskRequestUserPanes) SetPaneId(v int32) *StartRecordTaskRequestUserPanes {
	s.PaneId = &v
	return s
}

func (s *StartRecordTaskRequestUserPanes) SetSourceType(v string) *StartRecordTaskRequestUserPanes {
	s.SourceType = &v
	return s
}

func (s *StartRecordTaskRequestUserPanes) SetTexts(v []*StartRecordTaskRequestUserPanesTexts) *StartRecordTaskRequestUserPanes {
	s.Texts = v
	return s
}

func (s *StartRecordTaskRequestUserPanes) SetUserId(v string) *StartRecordTaskRequestUserPanes {
	s.UserId = &v
	return s
}

func (s *StartRecordTaskRequestUserPanes) Validate() error {
	return dara.Validate(s)
}

type StartRecordTaskRequestUserPanesImages struct {
	// example:
	//
	// 1
	Display *int32 `json:"Display,omitempty" xml:"Display,omitempty"`
	// example:
	//
	// 0.2456
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// https://www.example.com/image.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// 0.2456
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s StartRecordTaskRequestUserPanesImages) String() string {
	return dara.Prettify(s)
}

func (s StartRecordTaskRequestUserPanesImages) GoString() string {
	return s.String()
}

func (s *StartRecordTaskRequestUserPanesImages) GetDisplay() *int32 {
	return s.Display
}

func (s *StartRecordTaskRequestUserPanesImages) GetHeight() *float32 {
	return s.Height
}

func (s *StartRecordTaskRequestUserPanesImages) GetUrl() *string {
	return s.Url
}

func (s *StartRecordTaskRequestUserPanesImages) GetWidth() *float32 {
	return s.Width
}

func (s *StartRecordTaskRequestUserPanesImages) GetX() *float32 {
	return s.X
}

func (s *StartRecordTaskRequestUserPanesImages) GetY() *float32 {
	return s.Y
}

func (s *StartRecordTaskRequestUserPanesImages) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *StartRecordTaskRequestUserPanesImages) SetDisplay(v int32) *StartRecordTaskRequestUserPanesImages {
	s.Display = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetHeight(v float32) *StartRecordTaskRequestUserPanesImages {
	s.Height = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetUrl(v string) *StartRecordTaskRequestUserPanesImages {
	s.Url = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetWidth(v float32) *StartRecordTaskRequestUserPanesImages {
	s.Width = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetX(v float32) *StartRecordTaskRequestUserPanesImages {
	s.X = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetY(v float32) *StartRecordTaskRequestUserPanesImages {
	s.Y = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetZOrder(v int32) *StartRecordTaskRequestUserPanesImages {
	s.ZOrder = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) Validate() error {
	return dara.Validate(s)
}

type StartRecordTaskRequestUserPanesTexts struct {
	// example:
	//
	// 1
	FontColor *int32 `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// example:
	//
	// 1
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// example:
	//
	// 0
	FontType *int32 `json:"FontType,omitempty" xml:"FontType,omitempty"`
	// example:
	//
	// text
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s StartRecordTaskRequestUserPanesTexts) String() string {
	return dara.Prettify(s)
}

func (s StartRecordTaskRequestUserPanesTexts) GoString() string {
	return s.String()
}

func (s *StartRecordTaskRequestUserPanesTexts) GetFontColor() *int32 {
	return s.FontColor
}

func (s *StartRecordTaskRequestUserPanesTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *StartRecordTaskRequestUserPanesTexts) GetFontType() *int32 {
	return s.FontType
}

func (s *StartRecordTaskRequestUserPanesTexts) GetText() *string {
	return s.Text
}

func (s *StartRecordTaskRequestUserPanesTexts) GetX() *float32 {
	return s.X
}

func (s *StartRecordTaskRequestUserPanesTexts) GetY() *float32 {
	return s.Y
}

func (s *StartRecordTaskRequestUserPanesTexts) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *StartRecordTaskRequestUserPanesTexts) SetFontColor(v int32) *StartRecordTaskRequestUserPanesTexts {
	s.FontColor = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetFontSize(v int32) *StartRecordTaskRequestUserPanesTexts {
	s.FontSize = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetFontType(v int32) *StartRecordTaskRequestUserPanesTexts {
	s.FontType = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetText(v string) *StartRecordTaskRequestUserPanesTexts {
	s.Text = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetX(v float32) *StartRecordTaskRequestUserPanesTexts {
	s.X = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetY(v float32) *StartRecordTaskRequestUserPanesTexts {
	s.Y = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetZOrder(v int32) *StartRecordTaskRequestUserPanesTexts {
	s.ZOrder = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) Validate() error {
	return dara.Validate(s)
}
