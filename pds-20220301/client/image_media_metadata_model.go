// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageMediaMetadata interface {
	dara.Model
	String() string
	GoString() string
	SetAddressLine(v string) *ImageMediaMetadata
	GetAddressLine() *string
	SetCity(v string) *ImageMediaMetadata
	GetCity() *string
	SetCountry(v string) *ImageMediaMetadata
	GetCountry() *string
	SetDistrict(v string) *ImageMediaMetadata
	GetDistrict() *string
	SetExif(v string) *ImageMediaMetadata
	GetExif() *string
	SetFacesThumbnail(v []*FaceThumbnail) *ImageMediaMetadata
	GetFacesThumbnail() []*FaceThumbnail
	SetHeight(v int64) *ImageMediaMetadata
	GetHeight() *int64
	SetImageQuality(v *ImageQuality) *ImageMediaMetadata
	GetImageQuality() *ImageQuality
	SetImageTags(v []*SystemTag) *ImageMediaMetadata
	GetImageTags() []*SystemTag
	SetLocation(v string) *ImageMediaMetadata
	GetLocation() *string
	SetProvince(v string) *ImageMediaMetadata
	GetProvince() *string
	SetTime(v string) *ImageMediaMetadata
	GetTime() *string
	SetTownship(v string) *ImageMediaMetadata
	GetTownship() *string
	SetWidth(v int64) *ImageMediaMetadata
	GetWidth() *int64
}

type ImageMediaMetadata struct {
	// The full address.
	//
	// example:
	//
	// Jiangling Road, Xixing Street, Binjiang District, Hangzhou, Zhejiang
	AddressLine *string `json:"address_line,omitempty" xml:"address_line,omitempty"`
	// The city in which the image was taken.
	//
	// example:
	//
	// Hangzhou
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// The country or region in which the image was taken.
	//
	// example:
	//
	// China
	Country *string `json:"country,omitempty" xml:"country,omitempty"`
	// The district in which the image was taken.
	//
	// example:
	//
	// Binjiang District
	District *string `json:"district,omitempty" xml:"district,omitempty"`
	// The EXIF information about the image.
	//
	// example:
	//
	// {"Compression":{"value":"6"},"DateTime":{"value":"2020:08:19 17:11:11"}}
	Exif *string `json:"exif,omitempty" xml:"exif,omitempty"`
	// The thumbnails of the faces.
	FacesThumbnail []*FaceThumbnail `json:"faces_thumbnail,omitempty" xml:"faces_thumbnail,omitempty" type:"Repeated"`
	// The height of the image. Unit: pixel.
	//
	// example:
	//
	// 1024
	Height *int64 `json:"height,omitempty" xml:"height,omitempty"`
	// The rating of the image.
	ImageQuality *ImageQuality `json:"image_quality,omitempty" xml:"image_quality,omitempty"`
	// The details of the image tags.
	ImageTags []*SystemTag `json:"image_tags,omitempty" xml:"image_tags,omitempty" type:"Repeated"`
	// The location of the image.
	//
	// example:
	//
	// 30.185453,120.218522
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// The province in which the image was taken.
	//
	// example:
	//
	// Zhejiang
	Province *string `json:"province,omitempty" xml:"province,omitempty"`
	// The time when the image was taken. The time follows the RFC3339 standard.
	//
	// example:
	//
	// 2006-01-02T15:04:05.000Z07:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// The street in which the image was taken.
	//
	// example:
	//
	// Xixing Street
	Township *string `json:"township,omitempty" xml:"township,omitempty"`
	// The width of the image. Unit: pixel.
	//
	// example:
	//
	// 1024
	Width *int64 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s ImageMediaMetadata) String() string {
	return dara.Prettify(s)
}

func (s ImageMediaMetadata) GoString() string {
	return s.String()
}

func (s *ImageMediaMetadata) GetAddressLine() *string {
	return s.AddressLine
}

func (s *ImageMediaMetadata) GetCity() *string {
	return s.City
}

func (s *ImageMediaMetadata) GetCountry() *string {
	return s.Country
}

func (s *ImageMediaMetadata) GetDistrict() *string {
	return s.District
}

func (s *ImageMediaMetadata) GetExif() *string {
	return s.Exif
}

func (s *ImageMediaMetadata) GetFacesThumbnail() []*FaceThumbnail {
	return s.FacesThumbnail
}

func (s *ImageMediaMetadata) GetHeight() *int64 {
	return s.Height
}

func (s *ImageMediaMetadata) GetImageQuality() *ImageQuality {
	return s.ImageQuality
}

func (s *ImageMediaMetadata) GetImageTags() []*SystemTag {
	return s.ImageTags
}

func (s *ImageMediaMetadata) GetLocation() *string {
	return s.Location
}

func (s *ImageMediaMetadata) GetProvince() *string {
	return s.Province
}

func (s *ImageMediaMetadata) GetTime() *string {
	return s.Time
}

func (s *ImageMediaMetadata) GetTownship() *string {
	return s.Township
}

func (s *ImageMediaMetadata) GetWidth() *int64 {
	return s.Width
}

func (s *ImageMediaMetadata) SetAddressLine(v string) *ImageMediaMetadata {
	s.AddressLine = &v
	return s
}

func (s *ImageMediaMetadata) SetCity(v string) *ImageMediaMetadata {
	s.City = &v
	return s
}

func (s *ImageMediaMetadata) SetCountry(v string) *ImageMediaMetadata {
	s.Country = &v
	return s
}

func (s *ImageMediaMetadata) SetDistrict(v string) *ImageMediaMetadata {
	s.District = &v
	return s
}

func (s *ImageMediaMetadata) SetExif(v string) *ImageMediaMetadata {
	s.Exif = &v
	return s
}

func (s *ImageMediaMetadata) SetFacesThumbnail(v []*FaceThumbnail) *ImageMediaMetadata {
	s.FacesThumbnail = v
	return s
}

func (s *ImageMediaMetadata) SetHeight(v int64) *ImageMediaMetadata {
	s.Height = &v
	return s
}

func (s *ImageMediaMetadata) SetImageQuality(v *ImageQuality) *ImageMediaMetadata {
	s.ImageQuality = v
	return s
}

func (s *ImageMediaMetadata) SetImageTags(v []*SystemTag) *ImageMediaMetadata {
	s.ImageTags = v
	return s
}

func (s *ImageMediaMetadata) SetLocation(v string) *ImageMediaMetadata {
	s.Location = &v
	return s
}

func (s *ImageMediaMetadata) SetProvince(v string) *ImageMediaMetadata {
	s.Province = &v
	return s
}

func (s *ImageMediaMetadata) SetTime(v string) *ImageMediaMetadata {
	s.Time = &v
	return s
}

func (s *ImageMediaMetadata) SetTownship(v string) *ImageMediaMetadata {
	s.Township = &v
	return s
}

func (s *ImageMediaMetadata) SetWidth(v int64) *ImageMediaMetadata {
	s.Width = &v
	return s
}

func (s *ImageMediaMetadata) Validate() error {
	if s.FacesThumbnail != nil {
		for _, item := range s.FacesThumbnail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ImageQuality != nil {
		if err := s.ImageQuality.Validate(); err != nil {
			return err
		}
	}
	if s.ImageTags != nil {
		for _, item := range s.ImageTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
