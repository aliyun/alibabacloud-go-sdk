// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWatermarkConsoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *UpdateWatermarkConsoleRequest
	GetBucket() *string
	SetFileName(v string) *UpdateWatermarkConsoleRequest
	GetFileName() *string
	SetHeight(v string) *UpdateWatermarkConsoleRequest
	GetHeight() *string
	SetHorizontalOffet(v string) *UpdateWatermarkConsoleRequest
	GetHorizontalOffet() *string
	SetHorizontalOffset(v string) *UpdateWatermarkConsoleRequest
	GetHorizontalOffset() *string
	SetName(v string) *UpdateWatermarkConsoleRequest
	GetName() *string
	SetObject(v string) *UpdateWatermarkConsoleRequest
	GetObject() *string
	SetOwnerId(v int64) *UpdateWatermarkConsoleRequest
	GetOwnerId() *int64
	SetPosition(v string) *UpdateWatermarkConsoleRequest
	GetPosition() *string
	SetResourceOwnerAccount(v string) *UpdateWatermarkConsoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateWatermarkConsoleRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *UpdateWatermarkConsoleRequest
	GetResourceRealOwnerId() *int64
	SetScreenMode(v string) *UpdateWatermarkConsoleRequest
	GetScreenMode() *string
	SetType(v string) *UpdateWatermarkConsoleRequest
	GetType() *string
	SetVerticalOffset(v string) *UpdateWatermarkConsoleRequest
	GetVerticalOffset() *string
	SetVideoHeight(v int32) *UpdateWatermarkConsoleRequest
	GetVideoHeight() *int32
	SetVideoWidth(v int32) *UpdateWatermarkConsoleRequest
	GetVideoWidth() *int32
	SetWatermarkConfig(v string) *UpdateWatermarkConsoleRequest
	GetWatermarkConfig() *string
	SetWatermarkId(v string) *UpdateWatermarkConsoleRequest
	GetWatermarkId() *string
	SetWidth(v string) *UpdateWatermarkConsoleRequest
	GetWidth() *string
}

type UpdateWatermarkConsoleRequest struct {
	Bucket               *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	FileName             *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Height               *string `json:"Height,omitempty" xml:"Height,omitempty"`
	HorizontalOffet      *string `json:"HorizontalOffet,omitempty" xml:"HorizontalOffet,omitempty"`
	HorizontalOffset     *string `json:"HorizontalOffset,omitempty" xml:"HorizontalOffset,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Object               *string `json:"Object,omitempty" xml:"Object,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Position             *string `json:"Position,omitempty" xml:"Position,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	ScreenMode           *string `json:"ScreenMode,omitempty" xml:"ScreenMode,omitempty"`
	Type                 *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VerticalOffset       *string `json:"VerticalOffset,omitempty" xml:"VerticalOffset,omitempty"`
	VideoHeight          *int32  `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	VideoWidth           *int32  `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
	WatermarkConfig      *string `json:"WatermarkConfig,omitempty" xml:"WatermarkConfig,omitempty"`
	WatermarkId          *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
	Width                *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s UpdateWatermarkConsoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWatermarkConsoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkConsoleRequest) GetBucket() *string {
	return s.Bucket
}

func (s *UpdateWatermarkConsoleRequest) GetFileName() *string {
	return s.FileName
}

func (s *UpdateWatermarkConsoleRequest) GetHeight() *string {
	return s.Height
}

func (s *UpdateWatermarkConsoleRequest) GetHorizontalOffet() *string {
	return s.HorizontalOffet
}

func (s *UpdateWatermarkConsoleRequest) GetHorizontalOffset() *string {
	return s.HorizontalOffset
}

func (s *UpdateWatermarkConsoleRequest) GetName() *string {
	return s.Name
}

func (s *UpdateWatermarkConsoleRequest) GetObject() *string {
	return s.Object
}

func (s *UpdateWatermarkConsoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateWatermarkConsoleRequest) GetPosition() *string {
	return s.Position
}

func (s *UpdateWatermarkConsoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateWatermarkConsoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateWatermarkConsoleRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *UpdateWatermarkConsoleRequest) GetScreenMode() *string {
	return s.ScreenMode
}

func (s *UpdateWatermarkConsoleRequest) GetType() *string {
	return s.Type
}

func (s *UpdateWatermarkConsoleRequest) GetVerticalOffset() *string {
	return s.VerticalOffset
}

func (s *UpdateWatermarkConsoleRequest) GetVideoHeight() *int32 {
	return s.VideoHeight
}

func (s *UpdateWatermarkConsoleRequest) GetVideoWidth() *int32 {
	return s.VideoWidth
}

func (s *UpdateWatermarkConsoleRequest) GetWatermarkConfig() *string {
	return s.WatermarkConfig
}

func (s *UpdateWatermarkConsoleRequest) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *UpdateWatermarkConsoleRequest) GetWidth() *string {
	return s.Width
}

func (s *UpdateWatermarkConsoleRequest) SetBucket(v string) *UpdateWatermarkConsoleRequest {
	s.Bucket = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetFileName(v string) *UpdateWatermarkConsoleRequest {
	s.FileName = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetHeight(v string) *UpdateWatermarkConsoleRequest {
	s.Height = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetHorizontalOffet(v string) *UpdateWatermarkConsoleRequest {
	s.HorizontalOffet = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetHorizontalOffset(v string) *UpdateWatermarkConsoleRequest {
	s.HorizontalOffset = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetName(v string) *UpdateWatermarkConsoleRequest {
	s.Name = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetObject(v string) *UpdateWatermarkConsoleRequest {
	s.Object = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetOwnerId(v int64) *UpdateWatermarkConsoleRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetPosition(v string) *UpdateWatermarkConsoleRequest {
	s.Position = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetResourceOwnerAccount(v string) *UpdateWatermarkConsoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetResourceOwnerId(v int64) *UpdateWatermarkConsoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetResourceRealOwnerId(v int64) *UpdateWatermarkConsoleRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetScreenMode(v string) *UpdateWatermarkConsoleRequest {
	s.ScreenMode = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetType(v string) *UpdateWatermarkConsoleRequest {
	s.Type = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetVerticalOffset(v string) *UpdateWatermarkConsoleRequest {
	s.VerticalOffset = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetVideoHeight(v int32) *UpdateWatermarkConsoleRequest {
	s.VideoHeight = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetVideoWidth(v int32) *UpdateWatermarkConsoleRequest {
	s.VideoWidth = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetWatermarkConfig(v string) *UpdateWatermarkConsoleRequest {
	s.WatermarkConfig = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetWatermarkId(v string) *UpdateWatermarkConsoleRequest {
	s.WatermarkId = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) SetWidth(v string) *UpdateWatermarkConsoleRequest {
	s.Width = &v
	return s
}

func (s *UpdateWatermarkConsoleRequest) Validate() error {
	return dara.Validate(s)
}
