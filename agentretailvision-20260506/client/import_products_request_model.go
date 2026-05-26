// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *ImportProductsRequest
	GetDeviceId() *string
	SetExtraImages(v []*string) *ImportProductsRequest
	GetExtraImages() []*string
	SetImageTitle(v string) *ImportProductsRequest
	GetImageTitle() *string
	SetItemUniqueId(v string) *ImportProductsRequest
	GetItemUniqueId() *string
	SetMainImage(v []*string) *ImportProductsRequest
	GetMainImage() []*string
	SetMultiViewImages(v []*ImportProductsRequestMultiViewImages) *ImportProductsRequest
	GetMultiViewImages() []*ImportProductsRequestMultiViewImages
}

type ImportProductsRequest struct {
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
	MainImage       []*string                               `json:"MainImage,omitempty" xml:"MainImage,omitempty" type:"Repeated"`
	MultiViewImages []*ImportProductsRequestMultiViewImages `json:"MultiViewImages,omitempty" xml:"MultiViewImages,omitempty" type:"Repeated"`
}

func (s ImportProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportProductsRequest) GoString() string {
	return s.String()
}

func (s *ImportProductsRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ImportProductsRequest) GetExtraImages() []*string {
	return s.ExtraImages
}

func (s *ImportProductsRequest) GetImageTitle() *string {
	return s.ImageTitle
}

func (s *ImportProductsRequest) GetItemUniqueId() *string {
	return s.ItemUniqueId
}

func (s *ImportProductsRequest) GetMainImage() []*string {
	return s.MainImage
}

func (s *ImportProductsRequest) GetMultiViewImages() []*ImportProductsRequestMultiViewImages {
	return s.MultiViewImages
}

func (s *ImportProductsRequest) SetDeviceId(v string) *ImportProductsRequest {
	s.DeviceId = &v
	return s
}

func (s *ImportProductsRequest) SetExtraImages(v []*string) *ImportProductsRequest {
	s.ExtraImages = v
	return s
}

func (s *ImportProductsRequest) SetImageTitle(v string) *ImportProductsRequest {
	s.ImageTitle = &v
	return s
}

func (s *ImportProductsRequest) SetItemUniqueId(v string) *ImportProductsRequest {
	s.ItemUniqueId = &v
	return s
}

func (s *ImportProductsRequest) SetMainImage(v []*string) *ImportProductsRequest {
	s.MainImage = v
	return s
}

func (s *ImportProductsRequest) SetMultiViewImages(v []*ImportProductsRequestMultiViewImages) *ImportProductsRequest {
	s.MultiViewImages = v
	return s
}

func (s *ImportProductsRequest) Validate() error {
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

type ImportProductsRequestMultiViewImages struct {
	// example:
	//
	// 0
	Angle *string `json:"Angle,omitempty" xml:"Angle,omitempty"`
	// example:
	//
	// https://img5-parcel.oss-cn-hangzhou.aliyuncs.com/2026/01/12/78568805914464s.jpeg?07
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ImportProductsRequestMultiViewImages) String() string {
	return dara.Prettify(s)
}

func (s ImportProductsRequestMultiViewImages) GoString() string {
	return s.String()
}

func (s *ImportProductsRequestMultiViewImages) GetAngle() *string {
	return s.Angle
}

func (s *ImportProductsRequestMultiViewImages) GetUrl() *string {
	return s.Url
}

func (s *ImportProductsRequestMultiViewImages) SetAngle(v string) *ImportProductsRequestMultiViewImages {
	s.Angle = &v
	return s
}

func (s *ImportProductsRequestMultiViewImages) SetUrl(v string) *ImportProductsRequestMultiViewImages {
	s.Url = &v
	return s
}

func (s *ImportProductsRequestMultiViewImages) Validate() error {
	return dara.Validate(s)
}
