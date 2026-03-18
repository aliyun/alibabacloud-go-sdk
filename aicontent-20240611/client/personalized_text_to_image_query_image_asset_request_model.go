// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedTextToImageQueryImageAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncodeFormat(v string) *PersonalizedTextToImageQueryImageAssetRequest
	GetEncodeFormat() *string
	SetImageId(v string) *PersonalizedTextToImageQueryImageAssetRequest
	GetImageId() *string
}

type PersonalizedTextToImageQueryImageAssetRequest struct {
	// example:
	//
	// base64
	EncodeFormat *string `json:"encodeFormat,omitempty" xml:"encodeFormat,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0000.png
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
}

func (s PersonalizedTextToImageQueryImageAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s PersonalizedTextToImageQueryImageAssetRequest) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageQueryImageAssetRequest) GetEncodeFormat() *string {
	return s.EncodeFormat
}

func (s *PersonalizedTextToImageQueryImageAssetRequest) GetImageId() *string {
	return s.ImageId
}

func (s *PersonalizedTextToImageQueryImageAssetRequest) SetEncodeFormat(v string) *PersonalizedTextToImageQueryImageAssetRequest {
	s.EncodeFormat = &v
	return s
}

func (s *PersonalizedTextToImageQueryImageAssetRequest) SetImageId(v string) *PersonalizedTextToImageQueryImageAssetRequest {
	s.ImageId = &v
	return s
}

func (s *PersonalizedTextToImageQueryImageAssetRequest) Validate() error {
	return dara.Validate(s)
}
