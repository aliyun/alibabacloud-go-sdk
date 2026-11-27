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
	// The cropping mode of the video. Valid values:
	//
	// - **1**: crops the video while maintaining the aspect ratio.
	//
	// - **2*	- (default): adds padding to the video while maintaining the aspect ratio.
	//
	// example:
	//
	// 1
	CropMode *int64 `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	// The list of layout IDs.
	//
	// example:
	//
	// 1111
	LayoutIds []*int64 `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// The encoding option. For more information, see the MediaEncode enum values below.
	//
	// example:
	//
	// 20
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// The recording mode. Valid values:
	//
	// - **0**: single-stream recording. Task parameters cannot be modified during the recording process. You cannot call the UpdateRecordTask operation to modify task parameters.
	//
	// - **1*	- (default): mixed-stream recording. Multiple video streams are mixed and output.
	//
	// > - Valid parameters in single-stream recording mode: StreamType and SourceType.
	//
	// - Valid parameters in mixed-stream recording mode: TaskProfile, MediaEncode, LayoutIds, BackgroundColor, SubSpecUsers, CropMode, UserPanes, Backgrounds, Watermarks, and ClockWidgets.
	//
	// example:
	//
	// 1
	MixMode *int32 `json:"MixMode,omitempty" xml:"MixMode,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The video source in single-stream recording mode. Valid values:
	//
	// - **camera**: the camera.
	//
	// - **shareScreen**: screen sharing.
	//
	// > If the video source is shareScreen and only the camera stream is ingested, the camera stream is recorded. If both the camera stream and the screen sharing stream are ingested, the screen sharing stream is recorded.
	//
	// example:
	//
	// camera
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The single-stream recording mode. Valid values:
	//
	// - **0*	- (default): records the original stream.
	//
	// - **1**: records only the audio stream.
	//
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// Specifies the users whose audio streams are subscribed to in the channel.
	//
	// example:
	//
	// 1
	SubSpecAudioUsers []*string `json:"SubSpecAudioUsers,omitempty" xml:"SubSpecAudioUsers,omitempty" type:"Repeated"`
	// Specifies the users whose camera streams are subscribed to in the channel.
	//
	// example:
	//
	// 1
	SubSpecCameraUsers []*string `json:"SubSpecCameraUsers,omitempty" xml:"SubSpecCameraUsers,omitempty" type:"Repeated"`
	// Specifies the users whose screen sharing streams are subscribed to in the channel.
	//
	// example:
	//
	// 1
	SubSpecShareScreenUsers []*string `json:"SubSpecShareScreenUsers,omitempty" xml:"SubSpecShareScreenUsers,omitempty" type:"Repeated"`
	// Specifies the list of users to which the task subscribes.
	//
	// example:
	//
	// userID
	SubSpecUsers []*string `json:"SubSpecUsers,omitempty" xml:"SubSpecUsers,omitempty" type:"Repeated"`
	// The task ID. Only a single ID is supported and the ID must be unique. The ID must be 1 to 55 characters in length and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The billing configuration of the task. You are charged based on your settings. For more information, see the TaskProfile enum values below.
	//
	// > - The resolution of the MediaEncode encoding option must be less than or equal to the resolution of TaskProfile.
	//
	// - The maximum number of panes in the layout must be less than or equal to the number of input streams specified by TaskProfile.
	//
	// example:
	//
	// 4IN_1080P
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
	//
	// example:
	//
	// 1
	UnsubSpecAudioUsers []*string `json:"UnsubSpecAudioUsers,omitempty" xml:"UnsubSpecAudioUsers,omitempty" type:"Repeated"`
	// Specifies the users whose camera streams are not subscribed to in the channel.
	//
	// example:
	//
	// 1
	UnsubSpecCameraUsers []*string `json:"UnsubSpecCameraUsers,omitempty" xml:"UnsubSpecCameraUsers,omitempty" type:"Repeated"`
	// Specifies the users whose screen sharing streams are not subscribed to in the channel.
	//
	// example:
	//
	// 1
	UnsubSpecShareScreenUsers []*string `json:"UnsubSpecShareScreenUsers,omitempty" xml:"UnsubSpecShareScreenUsers,omitempty" type:"Repeated"`
	// The list of layout parameters.
	UserPanes []*StartRecordTaskRequestUserPanes `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
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

type StartRecordTaskRequestUserPanes struct {
	// The list of image panes in the current layout.
	Images []*StartRecordTaskRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The pane ID. Valid values: **0*	- to **15**.
	//
	// > - If the user specified by PaneId is not found, other streams are displayed by default. The streams are determined based on the order in which users join the channel.
	//
	// - If a parameter related to UserPanes is required, all other related parameters are also required. Otherwise, none of them are required.
	//
	// example:
	//
	// 2
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// The user video input that corresponds to the layout. Valid values:
	//
	// - **camera*	- (default): the video stream.
	//
	// - **shareScreen**: the screen sharing stream.
	//
	// example:
	//
	// camera
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The list of text panes in the current layout.
	Texts []*StartRecordTaskRequestUserPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// The ID of the user that corresponds to the layout pane. Valid values of N: **1 to 16**.
	//
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

type StartRecordTaskRequestUserPanesImages struct {
	// The display mode of the image. Valid values:
	//
	// - **0*	- (None): The image is not displayed.
	//
	// - **1*	- (Always): The image is always displayed.
	//
	// - **2*	- (Backup): The image is displayed when the current user has no video stream.
	//
	// example:
	//
	// 1
	Display *int32 `json:"Display,omitempty" xml:"Display,omitempty"`
	// The height of the pane, represented as a normalized percentage.
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
	// The width of the pane, represented as a normalized percentage.
	//
	// example:
	//
	// 0.2456
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The X coordinate, represented as a normalized percentage.
	//
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// The Y coordinate, represented as a normalized percentage.
	//
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// The stacking order. 0 indicates the bottom layer, 1 indicates the layer above 0, and so on.
	//
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
	// The text color in RGB format. The calculation formula is R + G × 256 + B × 65536. Valid values of R (red), G (green), and B (blue): **0 to 255**.
	//
	// example:
	//
	// 1
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
	// The X coordinate, represented as a normalized percentage.
	//
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// The Y coordinate, represented as a normalized percentage.
	//
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// The stacking order. 0 indicates the bottom layer, 1 indicates the layer above 0, and so on.
	//
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
