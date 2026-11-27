// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartViewShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartViewShrinkRequest
	GetAppId() *string
	SetBgColor(v *StartViewShrinkRequestBgColor) *StartViewShrinkRequest
	GetBgColor() *StartViewShrinkRequestBgColor
	SetChannelId(v string) *StartViewShrinkRequest
	GetChannelId() *string
	SetCropMode(v int32) *StartViewShrinkRequest
	GetCropMode() *int32
	SetRegionColor(v *StartViewShrinkRequestRegionColor) *StartViewShrinkRequest
	GetRegionColor() *StartViewShrinkRequestRegionColor
	SetStartWithoutChannel(v bool) *StartViewShrinkRequest
	GetStartWithoutChannel() *bool
	SetStartWithoutChannelWaitTime(v int32) *StartViewShrinkRequest
	GetStartWithoutChannelWaitTime() *int32
	SetTaskId(v string) *StartViewShrinkRequest
	GetTaskId() *string
	SetTemplateId(v string) *StartViewShrinkRequest
	GetTemplateId() *string
	SetViewContent(v string) *StartViewShrinkRequest
	GetViewContent() *string
	SetViewSubscribersShrink(v string) *StartViewShrinkRequest
	GetViewSubscribersShrink() *string
}

type StartViewShrinkRequest struct {
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The global background color.
	BgColor *StartViewShrinkRequestBgColor `json:"BgColor,omitempty" xml:"BgColor,omitempty" type:"Struct"`
	// The ID of an existing channel. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The crop mode. Valid values:
	//
	// - **1**: proportional cropping
	//
	// - **2**: proportional padding (default)
	//
	// - **3**: stretch to fill
	//
	// example:
	//
	// 2
	CropMode *int32 `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	// The global pane background color.
	RegionColor *StartViewShrinkRequestRegionColor `json:"RegionColor,omitempty" xml:"RegionColor,omitempty" type:"Struct"`
	// Specifies whether to allow starting the task when the channel does not exist. Valid values:
	//
	// - true: Yes.
	//
	// - false: No (default).
	StartWithoutChannel *bool `json:"StartWithoutChannel,omitempty" xml:"StartWithoutChannel,omitempty"`
	// The wait time for the task when the channel is not started, in seconds. This parameter is valid only when the StartWithoutChannel parameter is set to true. Valid values: 30 to 120.
	//
	// example:
	//
	// 30
	StartWithoutChannelWaitTime *int32 `json:"StartWithoutChannelWaitTime,omitempty" xml:"StartWithoutChannelWaitTime,omitempty"`
	// The task ID for the current request. Only a single ID is supported. The ID can contain letters, digits, underscores (_), and hyphens (-), and can be up to 55 bytes in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 567
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The stream property. Valid values:
	//
	// - main: main stream
	//
	// - screen-cast: shared stream
	//
	// example:
	//
	// main
	ViewContent *string `json:"ViewContent,omitempty" xml:"ViewContent,omitempty"`
	// The list of user IDs of Session Initiation Protocol (SIP) devices that are allowed to view the main stream or shared stream.
	ViewSubscribersShrink *string `json:"ViewSubscribers,omitempty" xml:"ViewSubscribers,omitempty"`
}

func (s StartViewShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartViewShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartViewShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartViewShrinkRequest) GetBgColor() *StartViewShrinkRequestBgColor {
	return s.BgColor
}

func (s *StartViewShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartViewShrinkRequest) GetCropMode() *int32 {
	return s.CropMode
}

func (s *StartViewShrinkRequest) GetRegionColor() *StartViewShrinkRequestRegionColor {
	return s.RegionColor
}

func (s *StartViewShrinkRequest) GetStartWithoutChannel() *bool {
	return s.StartWithoutChannel
}

func (s *StartViewShrinkRequest) GetStartWithoutChannelWaitTime() *int32 {
	return s.StartWithoutChannelWaitTime
}

func (s *StartViewShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StartViewShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *StartViewShrinkRequest) GetViewContent() *string {
	return s.ViewContent
}

func (s *StartViewShrinkRequest) GetViewSubscribersShrink() *string {
	return s.ViewSubscribersShrink
}

func (s *StartViewShrinkRequest) SetAppId(v string) *StartViewShrinkRequest {
	s.AppId = &v
	return s
}

func (s *StartViewShrinkRequest) SetBgColor(v *StartViewShrinkRequestBgColor) *StartViewShrinkRequest {
	s.BgColor = v
	return s
}

func (s *StartViewShrinkRequest) SetChannelId(v string) *StartViewShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *StartViewShrinkRequest) SetCropMode(v int32) *StartViewShrinkRequest {
	s.CropMode = &v
	return s
}

func (s *StartViewShrinkRequest) SetRegionColor(v *StartViewShrinkRequestRegionColor) *StartViewShrinkRequest {
	s.RegionColor = v
	return s
}

func (s *StartViewShrinkRequest) SetStartWithoutChannel(v bool) *StartViewShrinkRequest {
	s.StartWithoutChannel = &v
	return s
}

func (s *StartViewShrinkRequest) SetStartWithoutChannelWaitTime(v int32) *StartViewShrinkRequest {
	s.StartWithoutChannelWaitTime = &v
	return s
}

func (s *StartViewShrinkRequest) SetTaskId(v string) *StartViewShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *StartViewShrinkRequest) SetTemplateId(v string) *StartViewShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *StartViewShrinkRequest) SetViewContent(v string) *StartViewShrinkRequest {
	s.ViewContent = &v
	return s
}

func (s *StartViewShrinkRequest) SetViewSubscribersShrink(v string) *StartViewShrinkRequest {
	s.ViewSubscribersShrink = &v
	return s
}

func (s *StartViewShrinkRequest) Validate() error {
	if s.BgColor != nil {
		if err := s.BgColor.Validate(); err != nil {
			return err
		}
	}
	if s.RegionColor != nil {
		if err := s.RegionColor.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartViewShrinkRequestBgColor struct {
	// The B value.
	//
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// The G value.
	//
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// The R value.
	//
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartViewShrinkRequestBgColor) String() string {
	return dara.Prettify(s)
}

func (s StartViewShrinkRequestBgColor) GoString() string {
	return s.String()
}

func (s *StartViewShrinkRequestBgColor) GetB() *int32 {
	return s.B
}

func (s *StartViewShrinkRequestBgColor) GetG() *int32 {
	return s.G
}

func (s *StartViewShrinkRequestBgColor) GetR() *int32 {
	return s.R
}

func (s *StartViewShrinkRequestBgColor) SetB(v int32) *StartViewShrinkRequestBgColor {
	s.B = &v
	return s
}

func (s *StartViewShrinkRequestBgColor) SetG(v int32) *StartViewShrinkRequestBgColor {
	s.G = &v
	return s
}

func (s *StartViewShrinkRequestBgColor) SetR(v int32) *StartViewShrinkRequestBgColor {
	s.R = &v
	return s
}

func (s *StartViewShrinkRequestBgColor) Validate() error {
	return dara.Validate(s)
}

type StartViewShrinkRequestRegionColor struct {
	// The B value.
	//
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// The G value.
	//
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// The R value.
	//
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartViewShrinkRequestRegionColor) String() string {
	return dara.Prettify(s)
}

func (s StartViewShrinkRequestRegionColor) GoString() string {
	return s.String()
}

func (s *StartViewShrinkRequestRegionColor) GetB() *int32 {
	return s.B
}

func (s *StartViewShrinkRequestRegionColor) GetG() *int32 {
	return s.G
}

func (s *StartViewShrinkRequestRegionColor) GetR() *int32 {
	return s.R
}

func (s *StartViewShrinkRequestRegionColor) SetB(v int32) *StartViewShrinkRequestRegionColor {
	s.B = &v
	return s
}

func (s *StartViewShrinkRequestRegionColor) SetG(v int32) *StartViewShrinkRequestRegionColor {
	s.G = &v
	return s
}

func (s *StartViewShrinkRequestRegionColor) SetR(v int32) *StartViewShrinkRequestRegionColor {
	s.R = &v
	return s
}

func (s *StartViewShrinkRequestRegionColor) Validate() error {
	return dara.Validate(s)
}
