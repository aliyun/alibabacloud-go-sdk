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
	// Type of Image 1, with values:
	//
	// - **FacePic**: User\\"s face photo
	//
	// - **IDPic**: Headshot from the user\\"s second-generation ID card chip (typically obtained and decoded by a second-generation ID card reader)
	//
	// example:
	//
	// FacePic
	SourceImageType *string `json:"SourceImageType,omitempty" xml:"SourceImageType,omitempty"`
	// Address of Image 1. Please refer to the instructions on uploading image addresses.
	//
	// example:
	//
	// http%3A%2F%2Fjiangsu.china.com.cn%2Fuploadfile%2F2015%2F0114%2F1421221304095989.jpg
	SourceImageValue *string `json:"SourceImageValue,omitempty" xml:"SourceImageValue,omitempty"`
	// Type of Image 2, with values:
	//
	// - **FacePic**: User\\"s face photo
	//
	// - **IDPic**: Headshot from the user\\"s second-generation ID card chip (typically obtained and decoded by a second-generation ID card reader)
	//
	// example:
	//
	// FacePic
	TargetImageType *string `json:"TargetImageType,omitempty" xml:"TargetImageType,omitempty"`
	// Address of Image 2. Please refer to the instructions on uploading image addresses.
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
