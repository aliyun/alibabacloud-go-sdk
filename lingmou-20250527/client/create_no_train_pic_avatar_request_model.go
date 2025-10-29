// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNoTrainPicAvatarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpressiveness(v string) *CreateNoTrainPicAvatarRequest
	GetExpressiveness() *string
	SetGender(v string) *CreateNoTrainPicAvatarRequest
	GetGender() *string
	SetGenerateAssets(v bool) *CreateNoTrainPicAvatarRequest
	GetGenerateAssets() *bool
	SetImageOssPath(v string) *CreateNoTrainPicAvatarRequest
	GetImageOssPath() *string
	SetJwtToken(v string) *CreateNoTrainPicAvatarRequest
	GetJwtToken() *string
	SetName(v string) *CreateNoTrainPicAvatarRequest
	GetName() *string
	SetTransparent(v bool) *CreateNoTrainPicAvatarRequest
	GetTransparent() *bool
}

type CreateNoTrainPicAvatarRequest struct {
	// example:
	//
	// NORMAL/ENTHUSIASTIC
	Expressiveness *string `json:"expressiveness,omitempty" xml:"expressiveness,omitempty"`
	// example:
	//
	// FEMALE/MALE
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// true/false
	GenerateAssets *bool `json:"generateAssets,omitempty" xml:"generateAssets,omitempty"`
	// example:
	//
	// material/INPUT_INFER_PIC/Mt.CPQX3T6E25QU2/2e81e20797954440aed4da4264eb7494.webp
	ImageOssPath *string `json:"imageOssPath,omitempty" xml:"imageOssPath,omitempty"`
	// example:
	//
	// Token
	JwtToken *string `json:"jwtToken,omitempty" xml:"jwtToken,omitempty"`
	// example:
	//
	// avatar
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// true/false
	Transparent *bool `json:"transparent,omitempty" xml:"transparent,omitempty"`
}

func (s CreateNoTrainPicAvatarRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNoTrainPicAvatarRequest) GoString() string {
	return s.String()
}

func (s *CreateNoTrainPicAvatarRequest) GetExpressiveness() *string {
	return s.Expressiveness
}

func (s *CreateNoTrainPicAvatarRequest) GetGender() *string {
	return s.Gender
}

func (s *CreateNoTrainPicAvatarRequest) GetGenerateAssets() *bool {
	return s.GenerateAssets
}

func (s *CreateNoTrainPicAvatarRequest) GetImageOssPath() *string {
	return s.ImageOssPath
}

func (s *CreateNoTrainPicAvatarRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *CreateNoTrainPicAvatarRequest) GetName() *string {
	return s.Name
}

func (s *CreateNoTrainPicAvatarRequest) GetTransparent() *bool {
	return s.Transparent
}

func (s *CreateNoTrainPicAvatarRequest) SetExpressiveness(v string) *CreateNoTrainPicAvatarRequest {
	s.Expressiveness = &v
	return s
}

func (s *CreateNoTrainPicAvatarRequest) SetGender(v string) *CreateNoTrainPicAvatarRequest {
	s.Gender = &v
	return s
}

func (s *CreateNoTrainPicAvatarRequest) SetGenerateAssets(v bool) *CreateNoTrainPicAvatarRequest {
	s.GenerateAssets = &v
	return s
}

func (s *CreateNoTrainPicAvatarRequest) SetImageOssPath(v string) *CreateNoTrainPicAvatarRequest {
	s.ImageOssPath = &v
	return s
}

func (s *CreateNoTrainPicAvatarRequest) SetJwtToken(v string) *CreateNoTrainPicAvatarRequest {
	s.JwtToken = &v
	return s
}

func (s *CreateNoTrainPicAvatarRequest) SetName(v string) *CreateNoTrainPicAvatarRequest {
	s.Name = &v
	return s
}

func (s *CreateNoTrainPicAvatarRequest) SetTransparent(v bool) *CreateNoTrainPicAvatarRequest {
	s.Transparent = &v
	return s
}

func (s *CreateNoTrainPicAvatarRequest) Validate() error {
	return dara.Validate(s)
}
