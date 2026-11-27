// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartViewRequest
	GetAppId() *string
	SetBgColor(v *StartViewRequestBgColor) *StartViewRequest
	GetBgColor() *StartViewRequestBgColor
	SetChannelId(v string) *StartViewRequest
	GetChannelId() *string
	SetCropMode(v int32) *StartViewRequest
	GetCropMode() *int32
	SetRegionColor(v *StartViewRequestRegionColor) *StartViewRequest
	GetRegionColor() *StartViewRequestRegionColor
	SetStartWithoutChannel(v bool) *StartViewRequest
	GetStartWithoutChannel() *bool
	SetStartWithoutChannelWaitTime(v int32) *StartViewRequest
	GetStartWithoutChannelWaitTime() *int32
	SetTaskId(v string) *StartViewRequest
	GetTaskId() *string
	SetTemplateId(v string) *StartViewRequest
	GetTemplateId() *string
	SetViewContent(v string) *StartViewRequest
	GetViewContent() *string
	SetViewSubscribers(v []*string) *StartViewRequest
	GetViewSubscribers() []*string
}

type StartViewRequest struct {
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The global background color.
	BgColor *StartViewRequestBgColor `json:"BgColor,omitempty" xml:"BgColor,omitempty" type:"Struct"`
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
	RegionColor *StartViewRequestRegionColor `json:"RegionColor,omitempty" xml:"RegionColor,omitempty" type:"Struct"`
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
	ViewSubscribers []*string `json:"ViewSubscribers,omitempty" xml:"ViewSubscribers,omitempty" type:"Repeated"`
}

func (s StartViewRequest) String() string {
	return dara.Prettify(s)
}

func (s StartViewRequest) GoString() string {
	return s.String()
}

func (s *StartViewRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartViewRequest) GetBgColor() *StartViewRequestBgColor {
	return s.BgColor
}

func (s *StartViewRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartViewRequest) GetCropMode() *int32 {
	return s.CropMode
}

func (s *StartViewRequest) GetRegionColor() *StartViewRequestRegionColor {
	return s.RegionColor
}

func (s *StartViewRequest) GetStartWithoutChannel() *bool {
	return s.StartWithoutChannel
}

func (s *StartViewRequest) GetStartWithoutChannelWaitTime() *int32 {
	return s.StartWithoutChannelWaitTime
}

func (s *StartViewRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StartViewRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *StartViewRequest) GetViewContent() *string {
	return s.ViewContent
}

func (s *StartViewRequest) GetViewSubscribers() []*string {
	return s.ViewSubscribers
}

func (s *StartViewRequest) SetAppId(v string) *StartViewRequest {
	s.AppId = &v
	return s
}

func (s *StartViewRequest) SetBgColor(v *StartViewRequestBgColor) *StartViewRequest {
	s.BgColor = v
	return s
}

func (s *StartViewRequest) SetChannelId(v string) *StartViewRequest {
	s.ChannelId = &v
	return s
}

func (s *StartViewRequest) SetCropMode(v int32) *StartViewRequest {
	s.CropMode = &v
	return s
}

func (s *StartViewRequest) SetRegionColor(v *StartViewRequestRegionColor) *StartViewRequest {
	s.RegionColor = v
	return s
}

func (s *StartViewRequest) SetStartWithoutChannel(v bool) *StartViewRequest {
	s.StartWithoutChannel = &v
	return s
}

func (s *StartViewRequest) SetStartWithoutChannelWaitTime(v int32) *StartViewRequest {
	s.StartWithoutChannelWaitTime = &v
	return s
}

func (s *StartViewRequest) SetTaskId(v string) *StartViewRequest {
	s.TaskId = &v
	return s
}

func (s *StartViewRequest) SetTemplateId(v string) *StartViewRequest {
	s.TemplateId = &v
	return s
}

func (s *StartViewRequest) SetViewContent(v string) *StartViewRequest {
	s.ViewContent = &v
	return s
}

func (s *StartViewRequest) SetViewSubscribers(v []*string) *StartViewRequest {
	s.ViewSubscribers = v
	return s
}

func (s *StartViewRequest) Validate() error {
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

type StartViewRequestBgColor struct {
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

func (s StartViewRequestBgColor) String() string {
	return dara.Prettify(s)
}

func (s StartViewRequestBgColor) GoString() string {
	return s.String()
}

func (s *StartViewRequestBgColor) GetB() *int32 {
	return s.B
}

func (s *StartViewRequestBgColor) GetG() *int32 {
	return s.G
}

func (s *StartViewRequestBgColor) GetR() *int32 {
	return s.R
}

func (s *StartViewRequestBgColor) SetB(v int32) *StartViewRequestBgColor {
	s.B = &v
	return s
}

func (s *StartViewRequestBgColor) SetG(v int32) *StartViewRequestBgColor {
	s.G = &v
	return s
}

func (s *StartViewRequestBgColor) SetR(v int32) *StartViewRequestBgColor {
	s.R = &v
	return s
}

func (s *StartViewRequestBgColor) Validate() error {
	return dara.Validate(s)
}

type StartViewRequestRegionColor struct {
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

func (s StartViewRequestRegionColor) String() string {
	return dara.Prettify(s)
}

func (s StartViewRequestRegionColor) GoString() string {
	return s.String()
}

func (s *StartViewRequestRegionColor) GetB() *int32 {
	return s.B
}

func (s *StartViewRequestRegionColor) GetG() *int32 {
	return s.G
}

func (s *StartViewRequestRegionColor) GetR() *int32 {
	return s.R
}

func (s *StartViewRequestRegionColor) SetB(v int32) *StartViewRequestRegionColor {
	s.B = &v
	return s
}

func (s *StartViewRequestRegionColor) SetG(v int32) *StartViewRequestRegionColor {
	s.G = &v
	return s
}

func (s *StartViewRequestRegionColor) SetR(v int32) *StartViewRequestRegionColor {
	s.R = &v
	return s
}

func (s *StartViewRequestRegionColor) Validate() error {
	return dara.Validate(s)
}
