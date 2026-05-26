// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportProductsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *ImportProductsShrinkRequest
	GetDeviceId() *string
	SetExtraImagesShrink(v string) *ImportProductsShrinkRequest
	GetExtraImagesShrink() *string
	SetImageTitle(v string) *ImportProductsShrinkRequest
	GetImageTitle() *string
	SetItemUniqueId(v string) *ImportProductsShrinkRequest
	GetItemUniqueId() *string
	SetMainImageShrink(v string) *ImportProductsShrinkRequest
	GetMainImageShrink() *string
	SetMultiViewImagesShrink(v string) *ImportProductsShrinkRequest
	GetMultiViewImagesShrink() *string
}

type ImportProductsShrinkRequest struct {
	// example:
	//
	// DEVICE_001
	DeviceId          *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	ExtraImagesShrink *string `json:"ExtraImages,omitempty" xml:"ExtraImages,omitempty"`
	// example:
	//
	// 可口可乐330ml
	ImageTitle *string `json:"ImageTitle,omitempty" xml:"ImageTitle,omitempty"`
	// example:
	//
	// ITEM_001
	ItemUniqueId *string `json:"ItemUniqueId,omitempty" xml:"ItemUniqueId,omitempty"`
	// example:
	//
	// ["https://img.example.com/item1.jpg"]
	MainImageShrink       *string `json:"MainImage,omitempty" xml:"MainImage,omitempty"`
	MultiViewImagesShrink *string `json:"MultiViewImages,omitempty" xml:"MultiViewImages,omitempty"`
}

func (s ImportProductsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportProductsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportProductsShrinkRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ImportProductsShrinkRequest) GetExtraImagesShrink() *string {
	return s.ExtraImagesShrink
}

func (s *ImportProductsShrinkRequest) GetImageTitle() *string {
	return s.ImageTitle
}

func (s *ImportProductsShrinkRequest) GetItemUniqueId() *string {
	return s.ItemUniqueId
}

func (s *ImportProductsShrinkRequest) GetMainImageShrink() *string {
	return s.MainImageShrink
}

func (s *ImportProductsShrinkRequest) GetMultiViewImagesShrink() *string {
	return s.MultiViewImagesShrink
}

func (s *ImportProductsShrinkRequest) SetDeviceId(v string) *ImportProductsShrinkRequest {
	s.DeviceId = &v
	return s
}

func (s *ImportProductsShrinkRequest) SetExtraImagesShrink(v string) *ImportProductsShrinkRequest {
	s.ExtraImagesShrink = &v
	return s
}

func (s *ImportProductsShrinkRequest) SetImageTitle(v string) *ImportProductsShrinkRequest {
	s.ImageTitle = &v
	return s
}

func (s *ImportProductsShrinkRequest) SetItemUniqueId(v string) *ImportProductsShrinkRequest {
	s.ItemUniqueId = &v
	return s
}

func (s *ImportProductsShrinkRequest) SetMainImageShrink(v string) *ImportProductsShrinkRequest {
	s.MainImageShrink = &v
	return s
}

func (s *ImportProductsShrinkRequest) SetMultiViewImagesShrink(v string) *ImportProductsShrinkRequest {
	s.MultiViewImagesShrink = &v
	return s
}

func (s *ImportProductsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
