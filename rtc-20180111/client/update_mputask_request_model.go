// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMPUTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMPUTaskRequest
	GetAppId() *string
	SetBackgroundColor(v int32) *UpdateMPUTaskRequest
	GetBackgroundColor() *int32
	SetBackgrounds(v []*UpdateMPUTaskRequestBackgrounds) *UpdateMPUTaskRequest
	GetBackgrounds() []*UpdateMPUTaskRequestBackgrounds
	SetClockWidgets(v []*UpdateMPUTaskRequestClockWidgets) *UpdateMPUTaskRequest
	GetClockWidgets() []*UpdateMPUTaskRequestClockWidgets
	SetCropMode(v int32) *UpdateMPUTaskRequest
	GetCropMode() *int32
	SetLayoutIds(v []*int64) *UpdateMPUTaskRequest
	GetLayoutIds() []*int64
	SetMediaEncode(v int32) *UpdateMPUTaskRequest
	GetMediaEncode() *int32
	SetMixMode(v int32) *UpdateMPUTaskRequest
	GetMixMode() *int32
	SetOwnerId(v int64) *UpdateMPUTaskRequest
	GetOwnerId() *int64
	SetSourceType(v string) *UpdateMPUTaskRequest
	GetSourceType() *string
	SetStreamType(v int32) *UpdateMPUTaskRequest
	GetStreamType() *int32
	SetSubSpecAudioUsers(v []*string) *UpdateMPUTaskRequest
	GetSubSpecAudioUsers() []*string
	SetSubSpecCameraUsers(v []*string) *UpdateMPUTaskRequest
	GetSubSpecCameraUsers() []*string
	SetSubSpecShareScreenUsers(v []*string) *UpdateMPUTaskRequest
	GetSubSpecShareScreenUsers() []*string
	SetSubSpecUsers(v []*string) *UpdateMPUTaskRequest
	GetSubSpecUsers() []*string
	SetTaskId(v string) *UpdateMPUTaskRequest
	GetTaskId() *string
	SetUnsubSpecAudioUsers(v []*string) *UpdateMPUTaskRequest
	GetUnsubSpecAudioUsers() []*string
	SetUnsubSpecCameraUsers(v []*string) *UpdateMPUTaskRequest
	GetUnsubSpecCameraUsers() []*string
	SetUnsubSpecShareScreenUsers(v []*string) *UpdateMPUTaskRequest
	GetUnsubSpecShareScreenUsers() []*string
	SetUserPanes(v []*UpdateMPUTaskRequestUserPanes) *UpdateMPUTaskRequest
	GetUserPanes() []*UpdateMPUTaskRequestUserPanes
	SetWatermarks(v []*UpdateMPUTaskRequestWatermarks) *UpdateMPUTaskRequest
	GetWatermarks() []*UpdateMPUTaskRequestWatermarks
}

type UpdateMPUTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 0
	BackgroundColor *int32                              `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	Backgrounds     []*UpdateMPUTaskRequestBackgrounds  `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	ClockWidgets    []*UpdateMPUTaskRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CropMode  *int32   `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	LayoutIds []*int64 `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// example:
	//
	// 0
	MixMode *int32 `json:"MixMode,omitempty" xml:"MixMode,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// camera
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 0
	StreamType              *int32    `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	SubSpecAudioUsers       []*string `json:"SubSpecAudioUsers,omitempty" xml:"SubSpecAudioUsers,omitempty" type:"Repeated"`
	SubSpecCameraUsers      []*string `json:"SubSpecCameraUsers,omitempty" xml:"SubSpecCameraUsers,omitempty" type:"Repeated"`
	SubSpecShareScreenUsers []*string `json:"SubSpecShareScreenUsers,omitempty" xml:"SubSpecShareScreenUsers,omitempty" type:"Repeated"`
	SubSpecUsers            []*string `json:"SubSpecUsers,omitempty" xml:"SubSpecUsers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// testId
	TaskId                    *string                           `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	UnsubSpecAudioUsers       []*string                         `json:"UnsubSpecAudioUsers,omitempty" xml:"UnsubSpecAudioUsers,omitempty" type:"Repeated"`
	UnsubSpecCameraUsers      []*string                         `json:"UnsubSpecCameraUsers,omitempty" xml:"UnsubSpecCameraUsers,omitempty" type:"Repeated"`
	UnsubSpecShareScreenUsers []*string                         `json:"UnsubSpecShareScreenUsers,omitempty" xml:"UnsubSpecShareScreenUsers,omitempty" type:"Repeated"`
	UserPanes                 []*UpdateMPUTaskRequestUserPanes  `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
	Watermarks                []*UpdateMPUTaskRequestWatermarks `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
}

func (s UpdateMPUTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMPUTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMPUTaskRequest) GetBackgroundColor() *int32 {
	return s.BackgroundColor
}

func (s *UpdateMPUTaskRequest) GetBackgrounds() []*UpdateMPUTaskRequestBackgrounds {
	return s.Backgrounds
}

func (s *UpdateMPUTaskRequest) GetClockWidgets() []*UpdateMPUTaskRequestClockWidgets {
	return s.ClockWidgets
}

func (s *UpdateMPUTaskRequest) GetCropMode() *int32 {
	return s.CropMode
}

func (s *UpdateMPUTaskRequest) GetLayoutIds() []*int64 {
	return s.LayoutIds
}

func (s *UpdateMPUTaskRequest) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *UpdateMPUTaskRequest) GetMixMode() *int32 {
	return s.MixMode
}

func (s *UpdateMPUTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateMPUTaskRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateMPUTaskRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *UpdateMPUTaskRequest) GetSubSpecAudioUsers() []*string {
	return s.SubSpecAudioUsers
}

func (s *UpdateMPUTaskRequest) GetSubSpecCameraUsers() []*string {
	return s.SubSpecCameraUsers
}

func (s *UpdateMPUTaskRequest) GetSubSpecShareScreenUsers() []*string {
	return s.SubSpecShareScreenUsers
}

func (s *UpdateMPUTaskRequest) GetSubSpecUsers() []*string {
	return s.SubSpecUsers
}

func (s *UpdateMPUTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateMPUTaskRequest) GetUnsubSpecAudioUsers() []*string {
	return s.UnsubSpecAudioUsers
}

func (s *UpdateMPUTaskRequest) GetUnsubSpecCameraUsers() []*string {
	return s.UnsubSpecCameraUsers
}

func (s *UpdateMPUTaskRequest) GetUnsubSpecShareScreenUsers() []*string {
	return s.UnsubSpecShareScreenUsers
}

func (s *UpdateMPUTaskRequest) GetUserPanes() []*UpdateMPUTaskRequestUserPanes {
	return s.UserPanes
}

func (s *UpdateMPUTaskRequest) GetWatermarks() []*UpdateMPUTaskRequestWatermarks {
	return s.Watermarks
}

func (s *UpdateMPUTaskRequest) SetAppId(v string) *UpdateMPUTaskRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetBackgroundColor(v int32) *UpdateMPUTaskRequest {
	s.BackgroundColor = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetBackgrounds(v []*UpdateMPUTaskRequestBackgrounds) *UpdateMPUTaskRequest {
	s.Backgrounds = v
	return s
}

func (s *UpdateMPUTaskRequest) SetClockWidgets(v []*UpdateMPUTaskRequestClockWidgets) *UpdateMPUTaskRequest {
	s.ClockWidgets = v
	return s
}

func (s *UpdateMPUTaskRequest) SetCropMode(v int32) *UpdateMPUTaskRequest {
	s.CropMode = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetLayoutIds(v []*int64) *UpdateMPUTaskRequest {
	s.LayoutIds = v
	return s
}

func (s *UpdateMPUTaskRequest) SetMediaEncode(v int32) *UpdateMPUTaskRequest {
	s.MediaEncode = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetMixMode(v int32) *UpdateMPUTaskRequest {
	s.MixMode = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetOwnerId(v int64) *UpdateMPUTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetSourceType(v string) *UpdateMPUTaskRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetStreamType(v int32) *UpdateMPUTaskRequest {
	s.StreamType = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetSubSpecAudioUsers(v []*string) *UpdateMPUTaskRequest {
	s.SubSpecAudioUsers = v
	return s
}

func (s *UpdateMPUTaskRequest) SetSubSpecCameraUsers(v []*string) *UpdateMPUTaskRequest {
	s.SubSpecCameraUsers = v
	return s
}

func (s *UpdateMPUTaskRequest) SetSubSpecShareScreenUsers(v []*string) *UpdateMPUTaskRequest {
	s.SubSpecShareScreenUsers = v
	return s
}

func (s *UpdateMPUTaskRequest) SetSubSpecUsers(v []*string) *UpdateMPUTaskRequest {
	s.SubSpecUsers = v
	return s
}

func (s *UpdateMPUTaskRequest) SetTaskId(v string) *UpdateMPUTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetUnsubSpecAudioUsers(v []*string) *UpdateMPUTaskRequest {
	s.UnsubSpecAudioUsers = v
	return s
}

func (s *UpdateMPUTaskRequest) SetUnsubSpecCameraUsers(v []*string) *UpdateMPUTaskRequest {
	s.UnsubSpecCameraUsers = v
	return s
}

func (s *UpdateMPUTaskRequest) SetUnsubSpecShareScreenUsers(v []*string) *UpdateMPUTaskRequest {
	s.UnsubSpecShareScreenUsers = v
	return s
}

func (s *UpdateMPUTaskRequest) SetUserPanes(v []*UpdateMPUTaskRequestUserPanes) *UpdateMPUTaskRequest {
	s.UserPanes = v
	return s
}

func (s *UpdateMPUTaskRequest) SetWatermarks(v []*UpdateMPUTaskRequestWatermarks) *UpdateMPUTaskRequest {
	s.Watermarks = v
	return s
}

func (s *UpdateMPUTaskRequest) Validate() error {
	if s.Backgrounds != nil {
		for _, item := range s.Backgrounds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ClockWidgets != nil {
		for _, item := range s.ClockWidgets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	if s.Watermarks != nil {
		for _, item := range s.Watermarks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMPUTaskRequestBackgrounds struct {
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

func (s UpdateMPUTaskRequestBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s UpdateMPUTaskRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskRequestBackgrounds) GetDisplay() *int32 {
	return s.Display
}

func (s *UpdateMPUTaskRequestBackgrounds) GetHeight() *float32 {
	return s.Height
}

func (s *UpdateMPUTaskRequestBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *UpdateMPUTaskRequestBackgrounds) GetWidth() *float32 {
	return s.Width
}

func (s *UpdateMPUTaskRequestBackgrounds) GetX() *float32 {
	return s.X
}

func (s *UpdateMPUTaskRequestBackgrounds) GetY() *float32 {
	return s.Y
}

func (s *UpdateMPUTaskRequestBackgrounds) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *UpdateMPUTaskRequestBackgrounds) SetDisplay(v int32) *UpdateMPUTaskRequestBackgrounds {
	s.Display = &v
	return s
}

func (s *UpdateMPUTaskRequestBackgrounds) SetHeight(v float32) *UpdateMPUTaskRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *UpdateMPUTaskRequestBackgrounds) SetUrl(v string) *UpdateMPUTaskRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *UpdateMPUTaskRequestBackgrounds) SetWidth(v float32) *UpdateMPUTaskRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *UpdateMPUTaskRequestBackgrounds) SetX(v float32) *UpdateMPUTaskRequestBackgrounds {
	s.X = &v
	return s
}

func (s *UpdateMPUTaskRequestBackgrounds) SetY(v float32) *UpdateMPUTaskRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *UpdateMPUTaskRequestBackgrounds) SetZOrder(v int32) *UpdateMPUTaskRequestBackgrounds {
	s.ZOrder = &v
	return s
}

func (s *UpdateMPUTaskRequestBackgrounds) Validate() error {
	return dara.Validate(s)
}

type UpdateMPUTaskRequestClockWidgets struct {
	// example:
	//
	// 0
	Alpha *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// example:
	//
	// 0
	BorderColor *int64 `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// example:
	//
	// 1
	BorderWidth *int32 `json:"BorderWidth,omitempty" xml:"BorderWidth,omitempty"`
	// example:
	//
	// false
	Box *bool `json:"Box,omitempty" xml:"Box,omitempty"`
	// example:
	//
	// 0
	BoxBorderWidth *int32 `json:"BoxBorderWidth,omitempty" xml:"BoxBorderWidth,omitempty"`
	// example:
	//
	// 0
	BoxColor *int64 `json:"BoxColor,omitempty" xml:"BoxColor,omitempty"`
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

func (s UpdateMPUTaskRequestClockWidgets) String() string {
	return dara.Prettify(s)
}

func (s UpdateMPUTaskRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskRequestClockWidgets) GetAlpha() *float32 {
	return s.Alpha
}

func (s *UpdateMPUTaskRequestClockWidgets) GetBorderColor() *int64 {
	return s.BorderColor
}

func (s *UpdateMPUTaskRequestClockWidgets) GetBorderWidth() *int32 {
	return s.BorderWidth
}

func (s *UpdateMPUTaskRequestClockWidgets) GetBox() *bool {
	return s.Box
}

func (s *UpdateMPUTaskRequestClockWidgets) GetBoxBorderWidth() *int32 {
	return s.BoxBorderWidth
}

func (s *UpdateMPUTaskRequestClockWidgets) GetBoxColor() *int64 {
	return s.BoxColor
}

func (s *UpdateMPUTaskRequestClockWidgets) GetFontColor() *int32 {
	return s.FontColor
}

func (s *UpdateMPUTaskRequestClockWidgets) GetFontSize() *int32 {
	return s.FontSize
}

func (s *UpdateMPUTaskRequestClockWidgets) GetFontType() *int32 {
	return s.FontType
}

func (s *UpdateMPUTaskRequestClockWidgets) GetX() *float32 {
	return s.X
}

func (s *UpdateMPUTaskRequestClockWidgets) GetY() *float32 {
	return s.Y
}

func (s *UpdateMPUTaskRequestClockWidgets) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *UpdateMPUTaskRequestClockWidgets) SetAlpha(v float32) *UpdateMPUTaskRequestClockWidgets {
	s.Alpha = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetBorderColor(v int64) *UpdateMPUTaskRequestClockWidgets {
	s.BorderColor = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetBorderWidth(v int32) *UpdateMPUTaskRequestClockWidgets {
	s.BorderWidth = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetBox(v bool) *UpdateMPUTaskRequestClockWidgets {
	s.Box = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetBoxBorderWidth(v int32) *UpdateMPUTaskRequestClockWidgets {
	s.BoxBorderWidth = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetBoxColor(v int64) *UpdateMPUTaskRequestClockWidgets {
	s.BoxColor = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetFontColor(v int32) *UpdateMPUTaskRequestClockWidgets {
	s.FontColor = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetFontSize(v int32) *UpdateMPUTaskRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetFontType(v int32) *UpdateMPUTaskRequestClockWidgets {
	s.FontType = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetX(v float32) *UpdateMPUTaskRequestClockWidgets {
	s.X = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetY(v float32) *UpdateMPUTaskRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetZOrder(v int32) *UpdateMPUTaskRequestClockWidgets {
	s.ZOrder = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) Validate() error {
	return dara.Validate(s)
}

type UpdateMPUTaskRequestUserPanes struct {
	Images []*UpdateMPUTaskRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// example:
	//
	// 0
	SegmentType *int32 `json:"SegmentType,omitempty" xml:"SegmentType,omitempty"`
	// example:
	//
	// camera
	SourceType *string                               `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*UpdateMPUTaskRequestUserPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// example:
	//
	// TestUserID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateMPUTaskRequestUserPanes) String() string {
	return dara.Prettify(s)
}

func (s UpdateMPUTaskRequestUserPanes) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskRequestUserPanes) GetImages() []*UpdateMPUTaskRequestUserPanesImages {
	return s.Images
}

func (s *UpdateMPUTaskRequestUserPanes) GetPaneId() *int32 {
	return s.PaneId
}

func (s *UpdateMPUTaskRequestUserPanes) GetSegmentType() *int32 {
	return s.SegmentType
}

func (s *UpdateMPUTaskRequestUserPanes) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateMPUTaskRequestUserPanes) GetTexts() []*UpdateMPUTaskRequestUserPanesTexts {
	return s.Texts
}

func (s *UpdateMPUTaskRequestUserPanes) GetUserId() *string {
	return s.UserId
}

func (s *UpdateMPUTaskRequestUserPanes) SetImages(v []*UpdateMPUTaskRequestUserPanesImages) *UpdateMPUTaskRequestUserPanes {
	s.Images = v
	return s
}

func (s *UpdateMPUTaskRequestUserPanes) SetPaneId(v int32) *UpdateMPUTaskRequestUserPanes {
	s.PaneId = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanes) SetSegmentType(v int32) *UpdateMPUTaskRequestUserPanes {
	s.SegmentType = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanes) SetSourceType(v string) *UpdateMPUTaskRequestUserPanes {
	s.SourceType = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanes) SetTexts(v []*UpdateMPUTaskRequestUserPanesTexts) *UpdateMPUTaskRequestUserPanes {
	s.Texts = v
	return s
}

func (s *UpdateMPUTaskRequestUserPanes) SetUserId(v string) *UpdateMPUTaskRequestUserPanes {
	s.UserId = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanes) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Texts != nil {
		for _, item := range s.Texts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMPUTaskRequestUserPanesImages struct {
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

func (s UpdateMPUTaskRequestUserPanesImages) String() string {
	return dara.Prettify(s)
}

func (s UpdateMPUTaskRequestUserPanesImages) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskRequestUserPanesImages) GetDisplay() *int32 {
	return s.Display
}

func (s *UpdateMPUTaskRequestUserPanesImages) GetHeight() *float32 {
	return s.Height
}

func (s *UpdateMPUTaskRequestUserPanesImages) GetUrl() *string {
	return s.Url
}

func (s *UpdateMPUTaskRequestUserPanesImages) GetWidth() *float32 {
	return s.Width
}

func (s *UpdateMPUTaskRequestUserPanesImages) GetX() *float32 {
	return s.X
}

func (s *UpdateMPUTaskRequestUserPanesImages) GetY() *float32 {
	return s.Y
}

func (s *UpdateMPUTaskRequestUserPanesImages) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *UpdateMPUTaskRequestUserPanesImages) SetDisplay(v int32) *UpdateMPUTaskRequestUserPanesImages {
	s.Display = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesImages) SetHeight(v float32) *UpdateMPUTaskRequestUserPanesImages {
	s.Height = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesImages) SetUrl(v string) *UpdateMPUTaskRequestUserPanesImages {
	s.Url = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesImages) SetWidth(v float32) *UpdateMPUTaskRequestUserPanesImages {
	s.Width = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesImages) SetX(v float32) *UpdateMPUTaskRequestUserPanesImages {
	s.X = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesImages) SetY(v float32) *UpdateMPUTaskRequestUserPanesImages {
	s.Y = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesImages) SetZOrder(v int32) *UpdateMPUTaskRequestUserPanesImages {
	s.ZOrder = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesImages) Validate() error {
	return dara.Validate(s)
}

type UpdateMPUTaskRequestUserPanesTexts struct {
	// example:
	//
	// 0
	Alpha *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// example:
	//
	// 0
	BorderColor *int64 `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// example:
	//
	// 1
	BorderWidth *int32 `json:"BorderWidth,omitempty" xml:"BorderWidth,omitempty"`
	// example:
	//
	// false
	Box *bool `json:"Box,omitempty" xml:"Box,omitempty"`
	// example:
	//
	// 0
	BoxBorderWidth *int32 `json:"BoxBorderWidth,omitempty" xml:"BoxBorderWidth,omitempty"`
	// example:
	//
	// 0
	BoxColor *int64 `json:"BoxColor,omitempty" xml:"BoxColor,omitempty"`
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

func (s UpdateMPUTaskRequestUserPanesTexts) String() string {
	return dara.Prettify(s)
}

func (s UpdateMPUTaskRequestUserPanesTexts) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskRequestUserPanesTexts) GetAlpha() *float32 {
	return s.Alpha
}

func (s *UpdateMPUTaskRequestUserPanesTexts) GetBorderColor() *int64 {
	return s.BorderColor
}

func (s *UpdateMPUTaskRequestUserPanesTexts) GetBorderWidth() *int32 {
	return s.BorderWidth
}

func (s *UpdateMPUTaskRequestUserPanesTexts) GetBox() *bool {
	return s.Box
}

func (s *UpdateMPUTaskRequestUserPanesTexts) GetBoxBorderWidth() *int32 {
	return s.BoxBorderWidth
}

func (s *UpdateMPUTaskRequestUserPanesTexts) GetBoxColor() *int64 {
	return s.BoxColor
}

func (s *UpdateMPUTaskRequestUserPanesTexts) GetFontColor() *int32 {
	return s.FontColor
}

func (s *UpdateMPUTaskRequestUserPanesTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *UpdateMPUTaskRequestUserPanesTexts) GetFontType() *int32 {
	return s.FontType
}

func (s *UpdateMPUTaskRequestUserPanesTexts) GetText() *string {
	return s.Text
}

func (s *UpdateMPUTaskRequestUserPanesTexts) GetX() *float32 {
	return s.X
}

func (s *UpdateMPUTaskRequestUserPanesTexts) GetY() *float32 {
	return s.Y
}

func (s *UpdateMPUTaskRequestUserPanesTexts) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetAlpha(v float32) *UpdateMPUTaskRequestUserPanesTexts {
	s.Alpha = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetBorderColor(v int64) *UpdateMPUTaskRequestUserPanesTexts {
	s.BorderColor = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetBorderWidth(v int32) *UpdateMPUTaskRequestUserPanesTexts {
	s.BorderWidth = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetBox(v bool) *UpdateMPUTaskRequestUserPanesTexts {
	s.Box = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetBoxBorderWidth(v int32) *UpdateMPUTaskRequestUserPanesTexts {
	s.BoxBorderWidth = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetBoxColor(v int64) *UpdateMPUTaskRequestUserPanesTexts {
	s.BoxColor = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetFontColor(v int32) *UpdateMPUTaskRequestUserPanesTexts {
	s.FontColor = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetFontSize(v int32) *UpdateMPUTaskRequestUserPanesTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetFontType(v int32) *UpdateMPUTaskRequestUserPanesTexts {
	s.FontType = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetText(v string) *UpdateMPUTaskRequestUserPanesTexts {
	s.Text = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetX(v float32) *UpdateMPUTaskRequestUserPanesTexts {
	s.X = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetY(v float32) *UpdateMPUTaskRequestUserPanesTexts {
	s.Y = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetZOrder(v int32) *UpdateMPUTaskRequestUserPanesTexts {
	s.ZOrder = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) Validate() error {
	return dara.Validate(s)
}

type UpdateMPUTaskRequestWatermarks struct {
	// example:
	//
	// 0
	Alpha *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
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

func (s UpdateMPUTaskRequestWatermarks) String() string {
	return dara.Prettify(s)
}

func (s UpdateMPUTaskRequestWatermarks) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskRequestWatermarks) GetAlpha() *float32 {
	return s.Alpha
}

func (s *UpdateMPUTaskRequestWatermarks) GetDisplay() *int32 {
	return s.Display
}

func (s *UpdateMPUTaskRequestWatermarks) GetHeight() *float32 {
	return s.Height
}

func (s *UpdateMPUTaskRequestWatermarks) GetUrl() *string {
	return s.Url
}

func (s *UpdateMPUTaskRequestWatermarks) GetWidth() *float32 {
	return s.Width
}

func (s *UpdateMPUTaskRequestWatermarks) GetX() *float32 {
	return s.X
}

func (s *UpdateMPUTaskRequestWatermarks) GetY() *float32 {
	return s.Y
}

func (s *UpdateMPUTaskRequestWatermarks) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *UpdateMPUTaskRequestWatermarks) SetAlpha(v float32) *UpdateMPUTaskRequestWatermarks {
	s.Alpha = &v
	return s
}

func (s *UpdateMPUTaskRequestWatermarks) SetDisplay(v int32) *UpdateMPUTaskRequestWatermarks {
	s.Display = &v
	return s
}

func (s *UpdateMPUTaskRequestWatermarks) SetHeight(v float32) *UpdateMPUTaskRequestWatermarks {
	s.Height = &v
	return s
}

func (s *UpdateMPUTaskRequestWatermarks) SetUrl(v string) *UpdateMPUTaskRequestWatermarks {
	s.Url = &v
	return s
}

func (s *UpdateMPUTaskRequestWatermarks) SetWidth(v float32) *UpdateMPUTaskRequestWatermarks {
	s.Width = &v
	return s
}

func (s *UpdateMPUTaskRequestWatermarks) SetX(v float32) *UpdateMPUTaskRequestWatermarks {
	s.X = &v
	return s
}

func (s *UpdateMPUTaskRequestWatermarks) SetY(v float32) *UpdateMPUTaskRequestWatermarks {
	s.Y = &v
	return s
}

func (s *UpdateMPUTaskRequestWatermarks) SetZOrder(v int32) *UpdateMPUTaskRequestWatermarks {
	s.ZOrder = &v
	return s
}

func (s *UpdateMPUTaskRequestWatermarks) Validate() error {
	return dara.Validate(s)
}
