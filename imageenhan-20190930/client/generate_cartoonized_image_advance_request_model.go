// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iGenerateCartoonizedImageAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageType(v string) *GenerateCartoonizedImageAdvanceRequest
	GetImageType() *string
	SetImageUrlObject(v io.Reader) *GenerateCartoonizedImageAdvanceRequest
	GetImageUrlObject() io.Reader
	SetIndex(v string) *GenerateCartoonizedImageAdvanceRequest
	GetIndex() *string
}

type GenerateCartoonizedImageAdvanceRequest struct {
	// example:
	//
	// girl
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/zhangchaorun/1025.jpg
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
}

func (s GenerateCartoonizedImageAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateCartoonizedImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateCartoonizedImageAdvanceRequest) GetImageType() *string {
	return s.ImageType
}

func (s *GenerateCartoonizedImageAdvanceRequest) GetImageUrlObject() io.Reader {
	return s.ImageUrlObject
}

func (s *GenerateCartoonizedImageAdvanceRequest) GetIndex() *string {
	return s.Index
}

func (s *GenerateCartoonizedImageAdvanceRequest) SetImageType(v string) *GenerateCartoonizedImageAdvanceRequest {
	s.ImageType = &v
	return s
}

func (s *GenerateCartoonizedImageAdvanceRequest) SetImageUrlObject(v io.Reader) *GenerateCartoonizedImageAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *GenerateCartoonizedImageAdvanceRequest) SetIndex(v string) *GenerateCartoonizedImageAdvanceRequest {
	s.Index = &v
	return s
}

func (s *GenerateCartoonizedImageAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
