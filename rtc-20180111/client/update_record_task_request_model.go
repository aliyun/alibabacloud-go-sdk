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
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The channel ID. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CropMode  *int64  `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	// The array of layout IDs.
	LayoutIds   []*int64 `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	MediaEncode *int64   `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	OwnerId     *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies the users whose audio streams are subscribed to in the channel.
	SubSpecAudioUsers []*string `json:"SubSpecAudioUsers,omitempty" xml:"SubSpecAudioUsers,omitempty" type:"Repeated"`
	// Specifies the users whose camera streams are subscribed to in the channel.
	SubSpecCameraUsers []*string `json:"SubSpecCameraUsers,omitempty" xml:"SubSpecCameraUsers,omitempty" type:"Repeated"`
	// Specifies the users whose screen sharing streams are subscribed to in the channel.
	SubSpecShareScreenUsers []*string `json:"SubSpecShareScreenUsers,omitempty" xml:"SubSpecShareScreenUsers,omitempty" type:"Repeated"`
	// Specifies the list of users to which the task subscribes.
	//
	// example:
	//
	// userID
	SubSpecUsers []*string `json:"SubSpecUsers,omitempty" xml:"SubSpecUsers,omitempty" type:"Repeated"`
	// The task ID. Only a single ID is supported. The value must be the same as the TaskId request parameter of [StartRecordTask](https://help.aliyun.com/document_detail/193528.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskProfile *string `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	// The configuration template ID. For more information about how to obtain the template ID, see [AddRecordTemplate](https://help.aliyun.com/document_detail/193625.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 76dasgb****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Specifies the users whose audio streams are not subscribed to in the channel.
	UnsubSpecAudioUsers []*string `json:"UnsubSpecAudioUsers,omitempty" xml:"UnsubSpecAudioUsers,omitempty" type:"Repeated"`
	// Specifies the users whose camera streams are not subscribed to in the channel.
	UnsubSpecCameraUsers []*string `json:"UnsubSpecCameraUsers,omitempty" xml:"UnsubSpecCameraUsers,omitempty" type:"Repeated"`
	// Specifies the users whose screen sharing streams are not subscribed to in the channel.
	UnsubSpecShareScreenUsers []*string `json:"UnsubSpecShareScreenUsers,omitempty" xml:"UnsubSpecShareScreenUsers,omitempty" type:"Repeated"`
	// The list of layout parameters.
	UserPanes []*UpdateRecordTaskRequestUserPanes `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
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

type UpdateRecordTaskRequestUserPanes struct {
	// The list of image panes in the current layout.
	Images []*UpdateRecordTaskRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The pane ID. Valid values of N: **1 to 16**. Valid values of PaneId: **0 to 15**.
	//
	// > If the user specified by PaneId is not found, other streams are displayed by default. The streams are determined based on the order in which users join the channel.
	//
	// example:
	//
	// 1
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// The video input source of the user corresponding to the layout. Valid values:
	//
	// - **camera*	- (default): the video stream.
	//
	// - **shareScreen**: the screen sharing stream.
	//
	// Valid values of N: **1 to 16**.
	//
	// example:
	//
	// camera
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The list of text panes in the current layout.
	Texts []*UpdateRecordTaskRequestUserPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// The ID of the user corresponding to the layout pane. Valid values of N: **1 to 16**.
	//
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

type UpdateRecordTaskRequestUserPanesImages struct {
	// Specifies whether to display the image. Valid values:
	//
	// - **0*	- (None): Not displayed.
	//
	// - **1*	- (Always): Always displayed.
	//
	// - **2*	- (Backup): Displayed when the current user has no video stream.
	//
	// example:
	//
	// 1
	Display *int32 `json:"Display,omitempty" xml:"Display,omitempty"`
	// The height of the pane, as a normalized percentage.
	//
	// example:
	//
	// 0.2456
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The HTTP or HTTPS URL of the image.
	//
	// example:
	//
	// https://www.example.com/image.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The width of the pane, as a normalized percentage.
	//
	// example:
	//
	// 0.2456
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The X-coordinate, as a normalized percentage.
	//
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// The Y-coordinate, as a normalized percentage.
	//
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// The Z-order. 0 indicates the bottom layer, 1 indicates the layer above 0, and so on.
	//
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
	// The text color in RGB format. The calculation formula is R + G × 256 + B × 65536. Valid values of R (red), G (green), and B (blue): **0 to 255**.
	//
	// example:
	//
	// 0
	FontColor *int32 `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// The font size. Valid values: **(0, 72]**.
	//
	// example:
	//
	// 1
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// The font type. Valid values:
	//
	// - **0*	- (default): NOTO_SERIF_CJKSC_REGULAR.
	//
	// - **1**: ALIBABA_PUHUITI_REGULAR.
	//
	// - **2**: ALIBABA_PUHUITI_BOLD.
	//
	// - **3**: ALIBABA_PUHUITI_Heavy.
	//
	// - **4**: ALIBABA_PUHUITI_LIGHT.
	//
	// - **5**: ALIBABA_PUHUITI_MEDIUM.
	//
	// example:
	//
	// 0
	FontType *int32 `json:"FontType,omitempty" xml:"FontType,omitempty"`
	// The text content.
	//
	// example:
	//
	// text
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The X-coordinate, as a normalized percentage.
	//
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// The Y-coordinate, as a normalized percentage.
	//
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// The Z-order. 0 indicates the bottom layer, 1 indicates the layer above 0, and so on.
	//
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
