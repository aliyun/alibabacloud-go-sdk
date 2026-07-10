// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProductShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *UpdateProductShrinkRequest
	GetDeviceId() *string
	SetExtraImagesShrink(v string) *UpdateProductShrinkRequest
	GetExtraImagesShrink() *string
	SetImageTitle(v string) *UpdateProductShrinkRequest
	GetImageTitle() *string
	SetItemUniqueId(v string) *UpdateProductShrinkRequest
	GetItemUniqueId() *string
	SetMainImageShrink(v string) *UpdateProductShrinkRequest
	GetMainImageShrink() *string
	SetMultiViewImagesShrink(v string) *UpdateProductShrinkRequest
	GetMultiViewImagesShrink() *string
	SetPlatformItemId(v string) *UpdateProductShrinkRequest
	GetPlatformItemId() *string
}

type UpdateProductShrinkRequest struct {
	// The device ID, which is used to establish the vector association between the device and the item.
	//
	// example:
	//
	// DEVICE_001
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The list of additional image URLs that can be provided.
	ExtraImagesShrink *string `json:"ExtraImages,omitempty" xml:"ExtraImages,omitempty"`
	// The title of the item.
	//
	// example:
	//
	// 可口可乐330ml
	ImageTitle *string `json:"ImageTitle,omitempty" xml:"ImageTitle,omitempty"`
	// The business-side item ID, which is unique within the same business party.
	//
	// example:
	//
	// ITEM_001
	ItemUniqueId *string `json:"ItemUniqueId,omitempty" xml:"ItemUniqueId,omitempty"`
	// The list of main image URLs for the item. At least one image is required.
	//
	// example:
	//
	// ["https://img.example.com/item1.jpg"]
	MainImageShrink *string `json:"MainImage,omitempty" xml:"MainImage,omitempty"`
	// The list of multi-angle images for the item.
	MultiViewImagesShrink *string `json:"MultiViewImages,omitempty" xml:"MultiViewImages,omitempty"`
	// The platform item ID, which is globally unique.
	//
	// example:
	//
	// PLAT_001
	PlatformItemId *string `json:"PlatformItemId,omitempty" xml:"PlatformItemId,omitempty"`
}

func (s UpdateProductShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductShrinkRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *UpdateProductShrinkRequest) GetExtraImagesShrink() *string {
	return s.ExtraImagesShrink
}

func (s *UpdateProductShrinkRequest) GetImageTitle() *string {
	return s.ImageTitle
}

func (s *UpdateProductShrinkRequest) GetItemUniqueId() *string {
	return s.ItemUniqueId
}

func (s *UpdateProductShrinkRequest) GetMainImageShrink() *string {
	return s.MainImageShrink
}

func (s *UpdateProductShrinkRequest) GetMultiViewImagesShrink() *string {
	return s.MultiViewImagesShrink
}

func (s *UpdateProductShrinkRequest) GetPlatformItemId() *string {
	return s.PlatformItemId
}

func (s *UpdateProductShrinkRequest) SetDeviceId(v string) *UpdateProductShrinkRequest {
	s.DeviceId = &v
	return s
}

func (s *UpdateProductShrinkRequest) SetExtraImagesShrink(v string) *UpdateProductShrinkRequest {
	s.ExtraImagesShrink = &v
	return s
}

func (s *UpdateProductShrinkRequest) SetImageTitle(v string) *UpdateProductShrinkRequest {
	s.ImageTitle = &v
	return s
}

func (s *UpdateProductShrinkRequest) SetItemUniqueId(v string) *UpdateProductShrinkRequest {
	s.ItemUniqueId = &v
	return s
}

func (s *UpdateProductShrinkRequest) SetMainImageShrink(v string) *UpdateProductShrinkRequest {
	s.MainImageShrink = &v
	return s
}

func (s *UpdateProductShrinkRequest) SetMultiViewImagesShrink(v string) *UpdateProductShrinkRequest {
	s.MultiViewImagesShrink = &v
	return s
}

func (s *UpdateProductShrinkRequest) SetPlatformItemId(v string) *UpdateProductShrinkRequest {
	s.PlatformItemId = &v
	return s
}

func (s *UpdateProductShrinkRequest) Validate() error {
	return dara.Validate(s)
}
