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
	// A custom description for the watermark.
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
	// The height of the watermark image, in pixels. This value is relative to `RefHeight` and will be scaled proportionally with the actual video resolution.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The name of the watermark template.
	//
	// This parameter is required.
	//
	// example:
	//
	// livewatermark****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The anchor point for the watermark\\"s position. Valid values:
	//
	// - TopLeft
	//
	// - TopRight
	//
	// - BottomLeft
	//
	// - BottomRight
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
	// The reference height of the video background, in pixels.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1080
	RefHeight *int32 `json:"RefHeight,omitempty" xml:"RefHeight,omitempty"`
	// The reference width of the video background, in pixels.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1920
	RefWidth *int32 `json:"RefWidth,omitempty" xml:"RefWidth,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The opacity of the watermark. Value range: `0` (fully transparent) to `255` (fully opaque).
	//
	// This parameter is required.
	//
	// example:
	//
	// 255
	Transparency *int32 `json:"Transparency,omitempty" xml:"Transparency,omitempty"`
	// The type of the watermark. Valid value:
	//
	// - **0**: image.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The X-axis offset of the watermark, in pixels.
	//
	// > Relative to RefWidth. If OffsetCorner is TopLeft, XOffset is the horizontal distance between the top‑left corner of the watermark and the top‑left corner of the background video. Positive X points to the right.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50.0
	XOffset *float32 `json:"XOffset,omitempty" xml:"XOffset,omitempty"`
	// The Y-axis offset of the watermark, in pixels.
	//
	// > Relative to RefHeight. If OffsetCorner is TopLeft, YOffset is the vertical distance between the top‑left corner of the watermark and the top‑left corner of the background video. Positive Y points downward.
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
