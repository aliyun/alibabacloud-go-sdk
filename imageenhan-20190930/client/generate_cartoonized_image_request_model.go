// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCartoonizedImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageType(v string) *GenerateCartoonizedImageRequest
	GetImageType() *string
	SetImageUrl(v string) *GenerateCartoonizedImageRequest
	GetImageUrl() *string
	SetIndex(v string) *GenerateCartoonizedImageRequest
	GetIndex() *string
}

type GenerateCartoonizedImageRequest struct {
	// example:
	//
	// girl
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/zhangchaorun/1025.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
}

func (s GenerateCartoonizedImageRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateCartoonizedImageRequest) GoString() string {
	return s.String()
}

func (s *GenerateCartoonizedImageRequest) GetImageType() *string {
	return s.ImageType
}

func (s *GenerateCartoonizedImageRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *GenerateCartoonizedImageRequest) GetIndex() *string {
	return s.Index
}

func (s *GenerateCartoonizedImageRequest) SetImageType(v string) *GenerateCartoonizedImageRequest {
	s.ImageType = &v
	return s
}

func (s *GenerateCartoonizedImageRequest) SetImageUrl(v string) *GenerateCartoonizedImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *GenerateCartoonizedImageRequest) SetIndex(v string) *GenerateCartoonizedImageRequest {
	s.Index = &v
	return s
}

func (s *GenerateCartoonizedImageRequest) Validate() error {
	return dara.Validate(s)
}
