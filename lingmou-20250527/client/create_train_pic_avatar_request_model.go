// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrainPicAvatarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGender(v string) *CreateTrainPicAvatarRequest
	GetGender() *string
	SetGenerateAssets(v bool) *CreateTrainPicAvatarRequest
	GetGenerateAssets() *bool
	SetImageOssPath(v string) *CreateTrainPicAvatarRequest
	GetImageOssPath() *string
	SetName(v string) *CreateTrainPicAvatarRequest
	GetName() *string
	SetTemplateId(v string) *CreateTrainPicAvatarRequest
	GetTemplateId() *string
	SetTransparent(v bool) *CreateTrainPicAvatarRequest
	GetTransparent() *bool
}

type CreateTrainPicAvatarRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// FEMALE
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// true
	GenerateAssets *bool `json:"generateAssets,omitempty" xml:"generateAssets,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// material/INPUT_TRAIN_PIC/Mt.CQEJ2DQ6BBYU2/2.jpg
	ImageOssPath *string `json:"imageOssPath,omitempty" xml:"imageOssPath,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// M16vSG46Pby9HWOrFSZ7QaQA
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// example:
	//
	// false
	Transparent *bool `json:"transparent,omitempty" xml:"transparent,omitempty"`
}

func (s CreateTrainPicAvatarRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainPicAvatarRequest) GoString() string {
	return s.String()
}

func (s *CreateTrainPicAvatarRequest) GetGender() *string {
	return s.Gender
}

func (s *CreateTrainPicAvatarRequest) GetGenerateAssets() *bool {
	return s.GenerateAssets
}

func (s *CreateTrainPicAvatarRequest) GetImageOssPath() *string {
	return s.ImageOssPath
}

func (s *CreateTrainPicAvatarRequest) GetName() *string {
	return s.Name
}

func (s *CreateTrainPicAvatarRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateTrainPicAvatarRequest) GetTransparent() *bool {
	return s.Transparent
}

func (s *CreateTrainPicAvatarRequest) SetGender(v string) *CreateTrainPicAvatarRequest {
	s.Gender = &v
	return s
}

func (s *CreateTrainPicAvatarRequest) SetGenerateAssets(v bool) *CreateTrainPicAvatarRequest {
	s.GenerateAssets = &v
	return s
}

func (s *CreateTrainPicAvatarRequest) SetImageOssPath(v string) *CreateTrainPicAvatarRequest {
	s.ImageOssPath = &v
	return s
}

func (s *CreateTrainPicAvatarRequest) SetName(v string) *CreateTrainPicAvatarRequest {
	s.Name = &v
	return s
}

func (s *CreateTrainPicAvatarRequest) SetTemplateId(v string) *CreateTrainPicAvatarRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateTrainPicAvatarRequest) SetTransparent(v bool) *CreateTrainPicAvatarRequest {
	s.Transparent = &v
	return s
}

func (s *CreateTrainPicAvatarRequest) Validate() error {
	return dara.Validate(s)
}
