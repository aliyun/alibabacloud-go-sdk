// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareFacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceImageType(v string) *CompareFacesRequest
	GetSourceImageType() *string
	SetSourceImageValue(v string) *CompareFacesRequest
	GetSourceImageValue() *string
	SetTargetImageType(v string) *CompareFacesRequest
	GetTargetImageType() *string
	SetTargetImageValue(v string) *CompareFacesRequest
	GetTargetImageValue() *string
}

type CompareFacesRequest struct {
	// The type of image 1. Valid values:
	//
	// - **FacePic**: A face photo of the user.
	//
	// - **IDPic**: A headshot photo from the chip of the user\\"s second-generation ID card (typically read and decoded by a second-generation ID card reader device).
	//
	// example:
	//
	// FacePic
	SourceImageType *string `json:"SourceImageType,omitempty" xml:"SourceImageType,omitempty"`
	// The URL of image 1. For more information, see the description of image URL upload.
	//
	// example:
	//
	// http%3A%2F%2Fjiangsu.china.com.cn%2Fuploadfile%2F2015%2F0114%2F1421221304095989.jpg
	SourceImageValue *string `json:"SourceImageValue,omitempty" xml:"SourceImageValue,omitempty"`
	// The type of image 2. Valid values:
	//
	// - **FacePic**: A face photo of the user.
	//
	// - **IDPic**: A headshot photo from the chip of the user\\"s second-generation ID card (typically read and decoded by a second-generation ID card reader device).
	//
	// example:
	//
	// FacePic
	TargetImageType *string `json:"TargetImageType,omitempty" xml:"TargetImageType,omitempty"`
	// The URL of image 2. For more information, see the description of image URL upload.
	//
	// example:
	//
	// http%3A%2F%2Fjiangsu.china.com.cn%2Fuploadfile%2F2015%2F0114%2F1421221304095989.jpg
	TargetImageValue *string `json:"TargetImageValue,omitempty" xml:"TargetImageValue,omitempty"`
}

func (s CompareFacesRequest) String() string {
	return dara.Prettify(s)
}

func (s CompareFacesRequest) GoString() string {
	return s.String()
}

func (s *CompareFacesRequest) GetSourceImageType() *string {
	return s.SourceImageType
}

func (s *CompareFacesRequest) GetSourceImageValue() *string {
	return s.SourceImageValue
}

func (s *CompareFacesRequest) GetTargetImageType() *string {
	return s.TargetImageType
}

func (s *CompareFacesRequest) GetTargetImageValue() *string {
	return s.TargetImageValue
}

func (s *CompareFacesRequest) SetSourceImageType(v string) *CompareFacesRequest {
	s.SourceImageType = &v
	return s
}

func (s *CompareFacesRequest) SetSourceImageValue(v string) *CompareFacesRequest {
	s.SourceImageValue = &v
	return s
}

func (s *CompareFacesRequest) SetTargetImageType(v string) *CompareFacesRequest {
	s.TargetImageType = &v
	return s
}

func (s *CompareFacesRequest) SetTargetImageValue(v string) *CompareFacesRequest {
	s.TargetImageValue = &v
	return s
}

func (s *CompareFacesRequest) Validate() error {
	return dara.Validate(s)
}
