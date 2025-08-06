// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWatermarkConsoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v string) *AddWatermarkConsoleRequest
	GetActive() *string
	SetBucket(v string) *AddWatermarkConsoleRequest
	GetBucket() *string
	SetFileName(v string) *AddWatermarkConsoleRequest
	GetFileName() *string
	SetHeight(v string) *AddWatermarkConsoleRequest
	GetHeight() *string
	SetHorizontalOffet(v string) *AddWatermarkConsoleRequest
	GetHorizontalOffet() *string
	SetHorizontalOffset(v string) *AddWatermarkConsoleRequest
	GetHorizontalOffset() *string
	SetName(v string) *AddWatermarkConsoleRequest
	GetName() *string
	SetObject(v string) *AddWatermarkConsoleRequest
	GetObject() *string
	SetOwnerId(v int64) *AddWatermarkConsoleRequest
	GetOwnerId() *int64
	SetPosition(v string) *AddWatermarkConsoleRequest
	GetPosition() *string
	SetResourceOwnerAccount(v string) *AddWatermarkConsoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddWatermarkConsoleRequest
	GetResourceOwnerId() *int64
	SetResourceRealOwnerId(v int64) *AddWatermarkConsoleRequest
	GetResourceRealOwnerId() *int64
	SetScreenMode(v string) *AddWatermarkConsoleRequest
	GetScreenMode() *string
	SetType(v string) *AddWatermarkConsoleRequest
	GetType() *string
	SetVerticalOffset(v string) *AddWatermarkConsoleRequest
	GetVerticalOffset() *string
	SetVideoHeight(v int32) *AddWatermarkConsoleRequest
	GetVideoHeight() *int32
	SetVideoWidth(v int32) *AddWatermarkConsoleRequest
	GetVideoWidth() *int32
	SetWatermarkConfig(v string) *AddWatermarkConsoleRequest
	GetWatermarkConfig() *string
	SetWidth(v string) *AddWatermarkConsoleRequest
	GetWidth() *string
}

type AddWatermarkConsoleRequest struct {
	Active           *string `json:"Active,omitempty" xml:"Active,omitempty"`
	Bucket           *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	FileName         *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Height           *string `json:"Height,omitempty" xml:"Height,omitempty"`
	HorizontalOffet  *string `json:"HorizontalOffet,omitempty" xml:"HorizontalOffet,omitempty"`
	HorizontalOffset *string `json:"HorizontalOffset,omitempty" xml:"HorizontalOffset,omitempty"`
	// This parameter is required.
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
	Width                *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s AddWatermarkConsoleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWatermarkConsoleRequest) GoString() string {
	return s.String()
}

func (s *AddWatermarkConsoleRequest) GetActive() *string {
	return s.Active
}

func (s *AddWatermarkConsoleRequest) GetBucket() *string {
	return s.Bucket
}

func (s *AddWatermarkConsoleRequest) GetFileName() *string {
	return s.FileName
}

func (s *AddWatermarkConsoleRequest) GetHeight() *string {
	return s.Height
}

func (s *AddWatermarkConsoleRequest) GetHorizontalOffet() *string {
	return s.HorizontalOffet
}

func (s *AddWatermarkConsoleRequest) GetHorizontalOffset() *string {
	return s.HorizontalOffset
}

func (s *AddWatermarkConsoleRequest) GetName() *string {
	return s.Name
}

func (s *AddWatermarkConsoleRequest) GetObject() *string {
	return s.Object
}

func (s *AddWatermarkConsoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddWatermarkConsoleRequest) GetPosition() *string {
	return s.Position
}

func (s *AddWatermarkConsoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddWatermarkConsoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddWatermarkConsoleRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *AddWatermarkConsoleRequest) GetScreenMode() *string {
	return s.ScreenMode
}

func (s *AddWatermarkConsoleRequest) GetType() *string {
	return s.Type
}

func (s *AddWatermarkConsoleRequest) GetVerticalOffset() *string {
	return s.VerticalOffset
}

func (s *AddWatermarkConsoleRequest) GetVideoHeight() *int32 {
	return s.VideoHeight
}

func (s *AddWatermarkConsoleRequest) GetVideoWidth() *int32 {
	return s.VideoWidth
}

func (s *AddWatermarkConsoleRequest) GetWatermarkConfig() *string {
	return s.WatermarkConfig
}

func (s *AddWatermarkConsoleRequest) GetWidth() *string {
	return s.Width
}

func (s *AddWatermarkConsoleRequest) SetActive(v string) *AddWatermarkConsoleRequest {
	s.Active = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetBucket(v string) *AddWatermarkConsoleRequest {
	s.Bucket = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetFileName(v string) *AddWatermarkConsoleRequest {
	s.FileName = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetHeight(v string) *AddWatermarkConsoleRequest {
	s.Height = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetHorizontalOffet(v string) *AddWatermarkConsoleRequest {
	s.HorizontalOffet = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetHorizontalOffset(v string) *AddWatermarkConsoleRequest {
	s.HorizontalOffset = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetName(v string) *AddWatermarkConsoleRequest {
	s.Name = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetObject(v string) *AddWatermarkConsoleRequest {
	s.Object = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetOwnerId(v int64) *AddWatermarkConsoleRequest {
	s.OwnerId = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetPosition(v string) *AddWatermarkConsoleRequest {
	s.Position = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetResourceOwnerAccount(v string) *AddWatermarkConsoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetResourceOwnerId(v int64) *AddWatermarkConsoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetResourceRealOwnerId(v int64) *AddWatermarkConsoleRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetScreenMode(v string) *AddWatermarkConsoleRequest {
	s.ScreenMode = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetType(v string) *AddWatermarkConsoleRequest {
	s.Type = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetVerticalOffset(v string) *AddWatermarkConsoleRequest {
	s.VerticalOffset = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetVideoHeight(v int32) *AddWatermarkConsoleRequest {
	s.VideoHeight = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetVideoWidth(v int32) *AddWatermarkConsoleRequest {
	s.VideoWidth = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetWatermarkConfig(v string) *AddWatermarkConsoleRequest {
	s.WatermarkConfig = &v
	return s
}

func (s *AddWatermarkConsoleRequest) SetWidth(v string) *AddWatermarkConsoleRequest {
	s.Width = &v
	return s
}

func (s *AddWatermarkConsoleRequest) Validate() error {
	return dara.Validate(s)
}
