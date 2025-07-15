// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveStreamWatermarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateLiveStreamWatermarkRequest
	GetDescription() *string
	SetHeight(v int32) *UpdateLiveStreamWatermarkRequest
	GetHeight() *int32
	SetName(v string) *UpdateLiveStreamWatermarkRequest
	GetName() *string
	SetOffsetCorner(v string) *UpdateLiveStreamWatermarkRequest
	GetOffsetCorner() *string
	SetOwnerId(v int64) *UpdateLiveStreamWatermarkRequest
	GetOwnerId() *int64
	SetPictureUrl(v string) *UpdateLiveStreamWatermarkRequest
	GetPictureUrl() *string
	SetRefHeight(v int32) *UpdateLiveStreamWatermarkRequest
	GetRefHeight() *int32
	SetRefWidth(v int32) *UpdateLiveStreamWatermarkRequest
	GetRefWidth() *int32
	SetRegionId(v string) *UpdateLiveStreamWatermarkRequest
	GetRegionId() *string
	SetTemplateId(v string) *UpdateLiveStreamWatermarkRequest
	GetTemplateId() *string
	SetTransparency(v int32) *UpdateLiveStreamWatermarkRequest
	GetTransparency() *int32
	SetXOffset(v float32) *UpdateLiveStreamWatermarkRequest
	GetXOffset() *float32
	SetYOffset(v float32) *UpdateLiveStreamWatermarkRequest
	GetYOffset() *float32
}

type UpdateLiveStreamWatermarkRequest struct {
	// The description of the watermark.
	//
	// example:
	//
	// my watermark
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The height of the watermark. Unit: pixel. The height of the watermark is scaled in proportion to the height of the background video.
	//
	// example:
	//
	// 200
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The name of the watermark.
	//
	// example:
	//
	// livewatermark****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The position of the watermark. Valid values:
	//
	// 	- TopLeft: the upper-left corner.
	//
	// 	- TopRight: the upper-right corner.
	//
	// 	- BottomLeft: the lower-left corner.
	//
	// 	- BottomRight: the lower-right corner.
	//
	// example:
	//
	// TopRight
	OffsetCorner *string `json:"OffsetCorner,omitempty" xml:"OffsetCorner,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The URL of the watermark image.
	//
	// example:
	//
	// http://example.com
	PictureUrl *string `json:"PictureUrl,omitempty" xml:"PictureUrl,omitempty"`
	// The height of the background video. Unit: pixel.
	//
	// example:
	//
	// 1080
	RefHeight *int32 `json:"RefHeight,omitempty" xml:"RefHeight,omitempty"`
	// The width of the background video. Unit: pixel.
	//
	// example:
	//
	// 1920
	RefWidth *int32  `json:"RefWidth,omitempty" xml:"RefWidth,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the watermark template.
	//
	// >  You can obtain the template ID by checking the value of the TemplateId parameter that is returned by the [AddLiveStreamWatermark](https://help.aliyun.com/document_detail/410759.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The transparency of the watermark. A smaller value indicates a more transparent watermark. Valid values: 0 to 255.
	//
	// example:
	//
	// 255
	Transparency *int32 `json:"Transparency,omitempty" xml:"Transparency,omitempty"`
	// The offset of the watermark along the x-axis. Unit: pixel.
	//
	// >  In this case, the value of the RefWidth parameter is used as the reference. If the OffsetCorner parameter is set to TopLeft, the value of the XOffset parameter indicates the x-axis offset of the upper-left corner of the watermark relative to that of the background video. The directions from the coordinate axes to the center of the background video are positive. In other words, the x-axis is positive toward the right.
	//
	// example:
	//
	// 50.0
	XOffset *float32 `json:"XOffset,omitempty" xml:"XOffset,omitempty"`
	// The offset of the watermark along the y-axis. Unit: pixel.
	//
	// >  In this case, the value of the RefHeight parameter is used as the reference. If the OffsetCorner parameter is set to TopLeft, the value of the YOffset parameter indicates the y-axis offset of the upper-left corner of the watermark relative to that of the background video. The directions from the coordinate axes to the center of the background video are positive. In other words, the y-axis is positive downward.
	//
	// example:
	//
	// 100.0
	YOffset *float32 `json:"YOffset,omitempty" xml:"YOffset,omitempty"`
}

func (s UpdateLiveStreamWatermarkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveStreamWatermarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveStreamWatermarkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLiveStreamWatermarkRequest) GetHeight() *int32 {
	return s.Height
}

func (s *UpdateLiveStreamWatermarkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateLiveStreamWatermarkRequest) GetOffsetCorner() *string {
	return s.OffsetCorner
}

func (s *UpdateLiveStreamWatermarkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveStreamWatermarkRequest) GetPictureUrl() *string {
	return s.PictureUrl
}

func (s *UpdateLiveStreamWatermarkRequest) GetRefHeight() *int32 {
	return s.RefHeight
}

func (s *UpdateLiveStreamWatermarkRequest) GetRefWidth() *int32 {
	return s.RefWidth
}

func (s *UpdateLiveStreamWatermarkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLiveStreamWatermarkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateLiveStreamWatermarkRequest) GetTransparency() *int32 {
	return s.Transparency
}

func (s *UpdateLiveStreamWatermarkRequest) GetXOffset() *float32 {
	return s.XOffset
}

func (s *UpdateLiveStreamWatermarkRequest) GetYOffset() *float32 {
	return s.YOffset
}

func (s *UpdateLiveStreamWatermarkRequest) SetDescription(v string) *UpdateLiveStreamWatermarkRequest {
	s.Description = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRequest) SetHeight(v int32) *UpdateLiveStreamWatermarkRequest {
	s.Height = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRequest) SetName(v string) *UpdateLiveStreamWatermarkRequest {
	s.Name = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRequest) SetOffsetCorner(v string) *UpdateLiveStreamWatermarkRequest {
	s.OffsetCorner = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRequest) SetOwnerId(v int64) *UpdateLiveStreamWatermarkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRequest) SetPictureUrl(v string) *UpdateLiveStreamWatermarkRequest {
	s.PictureUrl = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRequest) SetRefHeight(v int32) *UpdateLiveStreamWatermarkRequest {
	s.RefHeight = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRequest) SetRefWidth(v int32) *UpdateLiveStreamWatermarkRequest {
	s.RefWidth = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRequest) SetRegionId(v string) *UpdateLiveStreamWatermarkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRequest) SetTemplateId(v string) *UpdateLiveStreamWatermarkRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRequest) SetTransparency(v int32) *UpdateLiveStreamWatermarkRequest {
	s.Transparency = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRequest) SetXOffset(v float32) *UpdateLiveStreamWatermarkRequest {
	s.XOffset = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRequest) SetYOffset(v float32) *UpdateLiveStreamWatermarkRequest {
	s.YOffset = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRequest) Validate() error {
	return dara.Validate(s)
}
