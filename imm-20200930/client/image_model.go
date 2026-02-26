// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImage interface {
	dara.Model
	String() string
	GoString() string
	SetCroppingSuggestions(v []*CroppingSuggestion) *Image
	GetCroppingSuggestions() []*CroppingSuggestion
	SetEXIF(v string) *Image
	GetEXIF() *string
	SetImageHeight(v int64) *Image
	GetImageHeight() *int64
	SetImageScore(v *ImageScore) *Image
	GetImageScore() *ImageScore
	SetImageWidth(v int64) *Image
	GetImageWidth() *int64
	SetOCRContents(v []*OCRContents) *Image
	GetOCRContents() []*OCRContents
}

type Image struct {
	// The image cropping suggestions. This parameter is reserved and not available.
	CroppingSuggestions []*CroppingSuggestion `json:"CroppingSuggestions,omitempty" xml:"CroppingSuggestions,omitempty" type:"Repeated"`
	// The original EXIF information about the image. The EXIF information is stored in the serialized JSON format. For more information, see [Query image information](https://help.aliyun.com/document_detail/44975.html).
	//
	// example:
	//
	// {"FileSize":{"value":"29304"},"Format":{"value":"jpg"}}
	EXIF *string `json:"EXIF,omitempty" xml:"EXIF,omitempty"`
	// The height of the image. Unit: pixels.
	//
	// example:
	//
	// 820
	ImageHeight *int64 `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	// The image scoring information.
	ImageScore *ImageScore `json:"ImageScore,omitempty" xml:"ImageScore,omitempty"`
	// The width of the image. Unit: pixels.
	//
	// example:
	//
	// 500
	ImageWidth *int64 `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	// The results of optical character recognition (OCR). This parameter is reserved and not available.
	OCRContents []*OCRContents `json:"OCRContents,omitempty" xml:"OCRContents,omitempty" type:"Repeated"`
}

func (s Image) String() string {
	return dara.Prettify(s)
}

func (s Image) GoString() string {
	return s.String()
}

func (s *Image) GetCroppingSuggestions() []*CroppingSuggestion {
	return s.CroppingSuggestions
}

func (s *Image) GetEXIF() *string {
	return s.EXIF
}

func (s *Image) GetImageHeight() *int64 {
	return s.ImageHeight
}

func (s *Image) GetImageScore() *ImageScore {
	return s.ImageScore
}

func (s *Image) GetImageWidth() *int64 {
	return s.ImageWidth
}

func (s *Image) GetOCRContents() []*OCRContents {
	return s.OCRContents
}

func (s *Image) SetCroppingSuggestions(v []*CroppingSuggestion) *Image {
	s.CroppingSuggestions = v
	return s
}

func (s *Image) SetEXIF(v string) *Image {
	s.EXIF = &v
	return s
}

func (s *Image) SetImageHeight(v int64) *Image {
	s.ImageHeight = &v
	return s
}

func (s *Image) SetImageScore(v *ImageScore) *Image {
	s.ImageScore = v
	return s
}

func (s *Image) SetImageWidth(v int64) *Image {
	s.ImageWidth = &v
	return s
}

func (s *Image) SetOCRContents(v []*OCRContents) *Image {
	s.OCRContents = v
	return s
}

func (s *Image) Validate() error {
	if s.CroppingSuggestions != nil {
		for _, item := range s.CroppingSuggestions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ImageScore != nil {
		if err := s.ImageScore.Validate(); err != nil {
			return err
		}
	}
	if s.OCRContents != nil {
		for _, item := range s.OCRContents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
