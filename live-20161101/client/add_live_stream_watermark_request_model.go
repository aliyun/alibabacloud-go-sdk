// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveStreamWatermarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *AddLiveStreamWatermarkRequest
	GetDescription() *string
	SetDomain(v string) *AddLiveStreamWatermarkRequest
	GetDomain() *string
	SetHeight(v int32) *AddLiveStreamWatermarkRequest
	GetHeight() *int32
	SetName(v string) *AddLiveStreamWatermarkRequest
	GetName() *string
	SetOffsetCorner(v string) *AddLiveStreamWatermarkRequest
	GetOffsetCorner() *string
	SetOwnerId(v int64) *AddLiveStreamWatermarkRequest
	GetOwnerId() *int64
	SetPictureUrl(v string) *AddLiveStreamWatermarkRequest
	GetPictureUrl() *string
	SetRefHeight(v int32) *AddLiveStreamWatermarkRequest
	GetRefHeight() *int32
	SetRefWidth(v int32) *AddLiveStreamWatermarkRequest
	GetRefWidth() *int32
	SetRegionId(v string) *AddLiveStreamWatermarkRequest
	GetRegionId() *string
	SetTransparency(v int32) *AddLiveStreamWatermarkRequest
	GetTransparency() *int32
	SetType(v int32) *AddLiveStreamWatermarkRequest
	GetType() *int32
	SetXOffset(v float32) *AddLiveStreamWatermarkRequest
	GetXOffset() *float32
	SetYOffset(v float32) *AddLiveStreamWatermarkRequest
	GetYOffset() *float32
}

type AddLiveStreamWatermarkRequest struct {
	// The description of the watermark.
	//
	// example:
	//
	// my watermark
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The streaming domain.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The height of the watermark. Unit: pixels. The height of the watermark is scaled in proportion to the height of the background video.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The name of the watermark.
	//
	// This parameter is required.
	//
	// example:
	//
	// livewatermark****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The location of the watermark. Valid values:
	//
	// 	- TopLeft: the upper-left corner.
	//
	// 	- TopRight: the upper-right corner.
	//
	// 	- BottomLeft: the lower-left corner.
	//
	// 	- BottomRight: the lower-right corner.
	//
	// This parameter is required.
	//
	// example:
	//
	// TopRight
	OffsetCorner *string `json:"OffsetCorner,omitempty" xml:"OffsetCorner,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The URL of the watermark image.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example.com
	PictureUrl *string `json:"PictureUrl,omitempty" xml:"PictureUrl,omitempty"`
	// The height of the background video. Unit: pixels.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1080
	RefHeight *int32 `json:"RefHeight,omitempty" xml:"RefHeight,omitempty"`
	// The width of the background video. Unit: pixels.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1920
	RefWidth *int32  `json:"RefWidth,omitempty" xml:"RefWidth,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The transparency of the watermark. A smaller value indicates a more transparent watermark. Valid values: 0 to 255.
	//
	// This parameter is required.
	//
	// example:
	//
	// 255
	Transparency *int32 `json:"Transparency,omitempty" xml:"Transparency,omitempty"`
	// The type of the watermark. Valid values:
	//
	// 	- **0**: image.
	//
	// 	- **1**: text. Only image watermarks are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The offset of the watermark along the x-axis. Unit: pixels.
	//
	// >  In this case, the value of the RefWidth parameter is used as the reference. If the OffsetCorner parameter is set to TopLeft, the value of the XOffset parameter indicates the x-axis offset of the upper-left corner of the watermark relative to that of the background video. The directions from the coordinate axes to the center of the background video are positive. In other words, the x-axis is positive toward the right.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50.0
	XOffset *float32 `json:"XOffset,omitempty" xml:"XOffset,omitempty"`
	// The offset of the watermark along the y-axis. Unit: pixels.
	//
	// >  In this case, the value of the RefHeight parameter is used as the reference. If the OffsetCorner parameter is set to TopLeft, the value of the YOffset parameter indicates the y-axis offset of the upper-left corner of the watermark relative to that of the background video. The directions from the coordinate axes to the center of the background video are positive. In other words, the y-axis is positive downward.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100.0
	YOffset *float32 `json:"YOffset,omitempty" xml:"YOffset,omitempty"`
}

func (s AddLiveStreamWatermarkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveStreamWatermarkRequest) GoString() string {
	return s.String()
}

func (s *AddLiveStreamWatermarkRequest) GetDescription() *string {
	return s.Description
}

func (s *AddLiveStreamWatermarkRequest) GetDomain() *string {
	return s.Domain
}

func (s *AddLiveStreamWatermarkRequest) GetHeight() *int32 {
	return s.Height
}

func (s *AddLiveStreamWatermarkRequest) GetName() *string {
	return s.Name
}

func (s *AddLiveStreamWatermarkRequest) GetOffsetCorner() *string {
	return s.OffsetCorner
}

func (s *AddLiveStreamWatermarkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveStreamWatermarkRequest) GetPictureUrl() *string {
	return s.PictureUrl
}

func (s *AddLiveStreamWatermarkRequest) GetRefHeight() *int32 {
	return s.RefHeight
}

func (s *AddLiveStreamWatermarkRequest) GetRefWidth() *int32 {
	return s.RefWidth
}

func (s *AddLiveStreamWatermarkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddLiveStreamWatermarkRequest) GetTransparency() *int32 {
	return s.Transparency
}

func (s *AddLiveStreamWatermarkRequest) GetType() *int32 {
	return s.Type
}

func (s *AddLiveStreamWatermarkRequest) GetXOffset() *float32 {
	return s.XOffset
}

func (s *AddLiveStreamWatermarkRequest) GetYOffset() *float32 {
	return s.YOffset
}

func (s *AddLiveStreamWatermarkRequest) SetDescription(v string) *AddLiveStreamWatermarkRequest {
	s.Description = &v
	return s
}

func (s *AddLiveStreamWatermarkRequest) SetDomain(v string) *AddLiveStreamWatermarkRequest {
	s.Domain = &v
	return s
}

func (s *AddLiveStreamWatermarkRequest) SetHeight(v int32) *AddLiveStreamWatermarkRequest {
	s.Height = &v
	return s
}

func (s *AddLiveStreamWatermarkRequest) SetName(v string) *AddLiveStreamWatermarkRequest {
	s.Name = &v
	return s
}

func (s *AddLiveStreamWatermarkRequest) SetOffsetCorner(v string) *AddLiveStreamWatermarkRequest {
	s.OffsetCorner = &v
	return s
}

func (s *AddLiveStreamWatermarkRequest) SetOwnerId(v int64) *AddLiveStreamWatermarkRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveStreamWatermarkRequest) SetPictureUrl(v string) *AddLiveStreamWatermarkRequest {
	s.PictureUrl = &v
	return s
}

func (s *AddLiveStreamWatermarkRequest) SetRefHeight(v int32) *AddLiveStreamWatermarkRequest {
	s.RefHeight = &v
	return s
}

func (s *AddLiveStreamWatermarkRequest) SetRefWidth(v int32) *AddLiveStreamWatermarkRequest {
	s.RefWidth = &v
	return s
}

func (s *AddLiveStreamWatermarkRequest) SetRegionId(v string) *AddLiveStreamWatermarkRequest {
	s.RegionId = &v
	return s
}

func (s *AddLiveStreamWatermarkRequest) SetTransparency(v int32) *AddLiveStreamWatermarkRequest {
	s.Transparency = &v
	return s
}

func (s *AddLiveStreamWatermarkRequest) SetType(v int32) *AddLiveStreamWatermarkRequest {
	s.Type = &v
	return s
}

func (s *AddLiveStreamWatermarkRequest) SetXOffset(v float32) *AddLiveStreamWatermarkRequest {
	s.XOffset = &v
	return s
}

func (s *AddLiveStreamWatermarkRequest) SetYOffset(v float32) *AddLiveStreamWatermarkRequest {
	s.YOffset = &v
	return s
}

func (s *AddLiveStreamWatermarkRequest) Validate() error {
	return dara.Validate(s)
}
