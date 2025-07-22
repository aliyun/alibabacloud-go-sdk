// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecordTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateRecordTaskRequest
	GetAppId() *string
	SetChannelId(v string) *UpdateRecordTaskRequest
	GetChannelId() *string
	SetCropMode(v int64) *UpdateRecordTaskRequest
	GetCropMode() *int64
	SetLayoutIds(v []*int64) *UpdateRecordTaskRequest
	GetLayoutIds() []*int64
	SetMediaEncode(v int64) *UpdateRecordTaskRequest
	GetMediaEncode() *int64
	SetOwnerId(v int64) *UpdateRecordTaskRequest
	GetOwnerId() *int64
	SetSubSpecAudioUsers(v []*string) *UpdateRecordTaskRequest
	GetSubSpecAudioUsers() []*string
	SetSubSpecCameraUsers(v []*string) *UpdateRecordTaskRequest
	GetSubSpecCameraUsers() []*string
	SetSubSpecShareScreenUsers(v []*string) *UpdateRecordTaskRequest
	GetSubSpecShareScreenUsers() []*string
	SetSubSpecUsers(v []*string) *UpdateRecordTaskRequest
	GetSubSpecUsers() []*string
	SetTaskId(v string) *UpdateRecordTaskRequest
	GetTaskId() *string
	SetTaskProfile(v string) *UpdateRecordTaskRequest
	GetTaskProfile() *string
	SetTemplateId(v string) *UpdateRecordTaskRequest
	GetTemplateId() *string
	SetUnsubSpecAudioUsers(v []*string) *UpdateRecordTaskRequest
	GetUnsubSpecAudioUsers() []*string
	SetUnsubSpecCameraUsers(v []*string) *UpdateRecordTaskRequest
	GetUnsubSpecCameraUsers() []*string
	SetUnsubSpecShareScreenUsers(v []*string) *UpdateRecordTaskRequest
	GetUnsubSpecShareScreenUsers() []*string
	SetUserPanes(v []*UpdateRecordTaskRequestUserPanes) *UpdateRecordTaskRequest
	GetUserPanes() []*UpdateRecordTaskRequestUserPanes
}

type UpdateRecordTaskRequest struct {
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
	ChannelId               *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CropMode                *int64    `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	LayoutIds               []*int64  `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	MediaEncode             *int64    `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	OwnerId                 *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SubSpecAudioUsers       []*string `json:"SubSpecAudioUsers,omitempty" xml:"SubSpecAudioUsers,omitempty" type:"Repeated"`
	SubSpecCameraUsers      []*string `json:"SubSpecCameraUsers,omitempty" xml:"SubSpecCameraUsers,omitempty" type:"Repeated"`
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
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskProfile *string `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 76dasgb****
	TemplateId                *string                             `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	UnsubSpecAudioUsers       []*string                           `json:"UnsubSpecAudioUsers,omitempty" xml:"UnsubSpecAudioUsers,omitempty" type:"Repeated"`
	UnsubSpecCameraUsers      []*string                           `json:"UnsubSpecCameraUsers,omitempty" xml:"UnsubSpecCameraUsers,omitempty" type:"Repeated"`
	UnsubSpecShareScreenUsers []*string                           `json:"UnsubSpecShareScreenUsers,omitempty" xml:"UnsubSpecShareScreenUsers,omitempty" type:"Repeated"`
	UserPanes                 []*UpdateRecordTaskRequestUserPanes `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
}

func (s UpdateRecordTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecordTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateRecordTaskRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpdateRecordTaskRequest) GetCropMode() *int64 {
	return s.CropMode
}

func (s *UpdateRecordTaskRequest) GetLayoutIds() []*int64 {
	return s.LayoutIds
}

func (s *UpdateRecordTaskRequest) GetMediaEncode() *int64 {
	return s.MediaEncode
}

func (s *UpdateRecordTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateRecordTaskRequest) GetSubSpecAudioUsers() []*string {
	return s.SubSpecAudioUsers
}

func (s *UpdateRecordTaskRequest) GetSubSpecCameraUsers() []*string {
	return s.SubSpecCameraUsers
}

func (s *UpdateRecordTaskRequest) GetSubSpecShareScreenUsers() []*string {
	return s.SubSpecShareScreenUsers
}

func (s *UpdateRecordTaskRequest) GetSubSpecUsers() []*string {
	return s.SubSpecUsers
}

func (s *UpdateRecordTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateRecordTaskRequest) GetTaskProfile() *string {
	return s.TaskProfile
}

func (s *UpdateRecordTaskRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateRecordTaskRequest) GetUnsubSpecAudioUsers() []*string {
	return s.UnsubSpecAudioUsers
}

func (s *UpdateRecordTaskRequest) GetUnsubSpecCameraUsers() []*string {
	return s.UnsubSpecCameraUsers
}

func (s *UpdateRecordTaskRequest) GetUnsubSpecShareScreenUsers() []*string {
	return s.UnsubSpecShareScreenUsers
}

func (s *UpdateRecordTaskRequest) GetUserPanes() []*UpdateRecordTaskRequestUserPanes {
	return s.UserPanes
}

func (s *UpdateRecordTaskRequest) SetAppId(v string) *UpdateRecordTaskRequest {
	s.AppId = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetChannelId(v string) *UpdateRecordTaskRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetCropMode(v int64) *UpdateRecordTaskRequest {
	s.CropMode = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetLayoutIds(v []*int64) *UpdateRecordTaskRequest {
	s.LayoutIds = v
	return s
}

func (s *UpdateRecordTaskRequest) SetMediaEncode(v int64) *UpdateRecordTaskRequest {
	s.MediaEncode = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetOwnerId(v int64) *UpdateRecordTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetSubSpecAudioUsers(v []*string) *UpdateRecordTaskRequest {
	s.SubSpecAudioUsers = v
	return s
}

func (s *UpdateRecordTaskRequest) SetSubSpecCameraUsers(v []*string) *UpdateRecordTaskRequest {
	s.SubSpecCameraUsers = v
	return s
}

func (s *UpdateRecordTaskRequest) SetSubSpecShareScreenUsers(v []*string) *UpdateRecordTaskRequest {
	s.SubSpecShareScreenUsers = v
	return s
}

func (s *UpdateRecordTaskRequest) SetSubSpecUsers(v []*string) *UpdateRecordTaskRequest {
	s.SubSpecUsers = v
	return s
}

func (s *UpdateRecordTaskRequest) SetTaskId(v string) *UpdateRecordTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetTaskProfile(v string) *UpdateRecordTaskRequest {
	s.TaskProfile = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetTemplateId(v string) *UpdateRecordTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetUnsubSpecAudioUsers(v []*string) *UpdateRecordTaskRequest {
	s.UnsubSpecAudioUsers = v
	return s
}

func (s *UpdateRecordTaskRequest) SetUnsubSpecCameraUsers(v []*string) *UpdateRecordTaskRequest {
	s.UnsubSpecCameraUsers = v
	return s
}

func (s *UpdateRecordTaskRequest) SetUnsubSpecShareScreenUsers(v []*string) *UpdateRecordTaskRequest {
	s.UnsubSpecShareScreenUsers = v
	return s
}

func (s *UpdateRecordTaskRequest) SetUserPanes(v []*UpdateRecordTaskRequestUserPanes) *UpdateRecordTaskRequest {
	s.UserPanes = v
	return s
}

func (s *UpdateRecordTaskRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateRecordTaskRequestUserPanes struct {
	Images []*UpdateRecordTaskRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// example:
	//
	// camera
	SourceType *string                                  `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*UpdateRecordTaskRequestUserPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// example:
	//
	// TestId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateRecordTaskRequestUserPanes) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordTaskRequestUserPanes) GoString() string {
	return s.String()
}

func (s *UpdateRecordTaskRequestUserPanes) GetImages() []*UpdateRecordTaskRequestUserPanesImages {
	return s.Images
}

func (s *UpdateRecordTaskRequestUserPanes) GetPaneId() *int32 {
	return s.PaneId
}

func (s *UpdateRecordTaskRequestUserPanes) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateRecordTaskRequestUserPanes) GetTexts() []*UpdateRecordTaskRequestUserPanesTexts {
	return s.Texts
}

func (s *UpdateRecordTaskRequestUserPanes) GetUserId() *string {
	return s.UserId
}

func (s *UpdateRecordTaskRequestUserPanes) SetImages(v []*UpdateRecordTaskRequestUserPanesImages) *UpdateRecordTaskRequestUserPanes {
	s.Images = v
	return s
}

func (s *UpdateRecordTaskRequestUserPanes) SetPaneId(v int32) *UpdateRecordTaskRequestUserPanes {
	s.PaneId = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanes) SetSourceType(v string) *UpdateRecordTaskRequestUserPanes {
	s.SourceType = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanes) SetTexts(v []*UpdateRecordTaskRequestUserPanesTexts) *UpdateRecordTaskRequestUserPanes {
	s.Texts = v
	return s
}

func (s *UpdateRecordTaskRequestUserPanes) SetUserId(v string) *UpdateRecordTaskRequestUserPanes {
	s.UserId = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanes) Validate() error {
	return dara.Validate(s)
}

type UpdateRecordTaskRequestUserPanesImages struct {
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

func (s UpdateRecordTaskRequestUserPanesImages) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordTaskRequestUserPanesImages) GoString() string {
	return s.String()
}

func (s *UpdateRecordTaskRequestUserPanesImages) GetDisplay() *int32 {
	return s.Display
}

func (s *UpdateRecordTaskRequestUserPanesImages) GetHeight() *float32 {
	return s.Height
}

func (s *UpdateRecordTaskRequestUserPanesImages) GetUrl() *string {
	return s.Url
}

func (s *UpdateRecordTaskRequestUserPanesImages) GetWidth() *float32 {
	return s.Width
}

func (s *UpdateRecordTaskRequestUserPanesImages) GetX() *float32 {
	return s.X
}

func (s *UpdateRecordTaskRequestUserPanesImages) GetY() *float32 {
	return s.Y
}

func (s *UpdateRecordTaskRequestUserPanesImages) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetDisplay(v int32) *UpdateRecordTaskRequestUserPanesImages {
	s.Display = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetHeight(v float32) *UpdateRecordTaskRequestUserPanesImages {
	s.Height = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetUrl(v string) *UpdateRecordTaskRequestUserPanesImages {
	s.Url = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetWidth(v float32) *UpdateRecordTaskRequestUserPanesImages {
	s.Width = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetX(v float32) *UpdateRecordTaskRequestUserPanesImages {
	s.X = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetY(v float32) *UpdateRecordTaskRequestUserPanesImages {
	s.Y = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetZOrder(v int32) *UpdateRecordTaskRequestUserPanesImages {
	s.ZOrder = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) Validate() error {
	return dara.Validate(s)
}

type UpdateRecordTaskRequestUserPanesTexts struct {
	// example:
	//
	// 0
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

func (s UpdateRecordTaskRequestUserPanesTexts) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordTaskRequestUserPanesTexts) GoString() string {
	return s.String()
}

func (s *UpdateRecordTaskRequestUserPanesTexts) GetFontColor() *int32 {
	return s.FontColor
}

func (s *UpdateRecordTaskRequestUserPanesTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *UpdateRecordTaskRequestUserPanesTexts) GetFontType() *int32 {
	return s.FontType
}

func (s *UpdateRecordTaskRequestUserPanesTexts) GetText() *string {
	return s.Text
}

func (s *UpdateRecordTaskRequestUserPanesTexts) GetX() *float32 {
	return s.X
}

func (s *UpdateRecordTaskRequestUserPanesTexts) GetY() *float32 {
	return s.Y
}

func (s *UpdateRecordTaskRequestUserPanesTexts) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetFontColor(v int32) *UpdateRecordTaskRequestUserPanesTexts {
	s.FontColor = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetFontSize(v int32) *UpdateRecordTaskRequestUserPanesTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetFontType(v int32) *UpdateRecordTaskRequestUserPanesTexts {
	s.FontType = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetText(v string) *UpdateRecordTaskRequestUserPanesTexts {
	s.Text = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetX(v float32) *UpdateRecordTaskRequestUserPanesTexts {
	s.X = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetY(v float32) *UpdateRecordTaskRequestUserPanesTexts {
	s.Y = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetZOrder(v int32) *UpdateRecordTaskRequestUserPanesTexts {
	s.ZOrder = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) Validate() error {
	return dara.Validate(s)
}
