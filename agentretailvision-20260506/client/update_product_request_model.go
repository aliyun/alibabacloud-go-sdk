// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *UpdateProductRequest
	GetDeviceId() *string
	SetExtraImages(v []*string) *UpdateProductRequest
	GetExtraImages() []*string
	SetImageTitle(v string) *UpdateProductRequest
	GetImageTitle() *string
	SetItemUniqueId(v string) *UpdateProductRequest
	GetItemUniqueId() *string
	SetMainImage(v []*string) *UpdateProductRequest
	GetMainImage() []*string
	SetMultiViewImages(v []*UpdateProductRequestMultiViewImages) *UpdateProductRequest
	GetMultiViewImages() []*UpdateProductRequestMultiViewImages
	SetPlatformItemId(v string) *UpdateProductRequest
	GetPlatformItemId() *string
}

type UpdateProductRequest struct {
	// example:
	//
	// DEVICE_001
	DeviceId    *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	ExtraImages []*string `json:"ExtraImages,omitempty" xml:"ExtraImages,omitempty" type:"Repeated"`
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
	MainImage       []*string                              `json:"MainImage,omitempty" xml:"MainImage,omitempty" type:"Repeated"`
	MultiViewImages []*UpdateProductRequestMultiViewImages `json:"MultiViewImages,omitempty" xml:"MultiViewImages,omitempty" type:"Repeated"`
	// example:
	//
	// PLAT_001
	PlatformItemId *string `json:"PlatformItemId,omitempty" xml:"PlatformItemId,omitempty"`
}

func (s UpdateProductRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *UpdateProductRequest) GetExtraImages() []*string {
	return s.ExtraImages
}

func (s *UpdateProductRequest) GetImageTitle() *string {
	return s.ImageTitle
}

func (s *UpdateProductRequest) GetItemUniqueId() *string {
	return s.ItemUniqueId
}

func (s *UpdateProductRequest) GetMainImage() []*string {
	return s.MainImage
}

func (s *UpdateProductRequest) GetMultiViewImages() []*UpdateProductRequestMultiViewImages {
	return s.MultiViewImages
}

func (s *UpdateProductRequest) GetPlatformItemId() *string {
	return s.PlatformItemId
}

func (s *UpdateProductRequest) SetDeviceId(v string) *UpdateProductRequest {
	s.DeviceId = &v
	return s
}

func (s *UpdateProductRequest) SetExtraImages(v []*string) *UpdateProductRequest {
	s.ExtraImages = v
	return s
}

func (s *UpdateProductRequest) SetImageTitle(v string) *UpdateProductRequest {
	s.ImageTitle = &v
	return s
}

func (s *UpdateProductRequest) SetItemUniqueId(v string) *UpdateProductRequest {
	s.ItemUniqueId = &v
	return s
}

func (s *UpdateProductRequest) SetMainImage(v []*string) *UpdateProductRequest {
	s.MainImage = v
	return s
}

func (s *UpdateProductRequest) SetMultiViewImages(v []*UpdateProductRequestMultiViewImages) *UpdateProductRequest {
	s.MultiViewImages = v
	return s
}

func (s *UpdateProductRequest) SetPlatformItemId(v string) *UpdateProductRequest {
	s.PlatformItemId = &v
	return s
}

func (s *UpdateProductRequest) Validate() error {
	if s.MultiViewImages != nil {
		for _, item := range s.MultiViewImages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateProductRequestMultiViewImages struct {
	// example:
	//
	// 0
	Angle *string `json:"Angle,omitempty" xml:"Angle,omitempty"`
	// example:
	//
	// https://nova-tems.oss-cn-shanghai.aliyuncs.com/crop/33dfc602-c9a4-11f0-ac99-ee21a901d6ec.png?OSSAccessKeyId=****&Expires=1764058353&Signature=****
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateProductRequestMultiViewImages) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductRequestMultiViewImages) GoString() string {
	return s.String()
}

func (s *UpdateProductRequestMultiViewImages) GetAngle() *string {
	return s.Angle
}

func (s *UpdateProductRequestMultiViewImages) GetUrl() *string {
	return s.Url
}

func (s *UpdateProductRequestMultiViewImages) SetAngle(v string) *UpdateProductRequestMultiViewImages {
	s.Angle = &v
	return s
}

func (s *UpdateProductRequestMultiViewImages) SetUrl(v string) *UpdateProductRequestMultiViewImages {
	s.Url = &v
	return s
}

func (s *UpdateProductRequestMultiViewImages) Validate() error {
	return dara.Validate(s)
}
