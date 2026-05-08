// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnchorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnchorCategory(v string) *CreateAnchorRequest
	GetAnchorCategory() *string
	SetAnchorMaterialName(v string) *CreateAnchorRequest
	GetAnchorMaterialName() *string
	SetCoverUrl(v string) *CreateAnchorRequest
	GetCoverUrl() *string
	SetDigitalHumanType(v string) *CreateAnchorRequest
	GetDigitalHumanType() *string
	SetGender(v string) *CreateAnchorRequest
	GetGender() *string
	SetUseScene(v string) *CreateAnchorRequest
	GetUseScene() *string
	SetVideoOssKey(v string) *CreateAnchorRequest
	GetVideoOssKey() *string
}

type CreateAnchorRequest struct {
	// 类型
	//
	// example:
	//
	// VIDEO_ANCHOR
	AnchorCategory     *string `json:"anchorCategory,omitempty" xml:"anchorCategory,omitempty"`
	AnchorMaterialName *string `json:"anchorMaterialName,omitempty" xml:"anchorMaterialName,omitempty"`
	// example:
	//
	// https://xxx
	CoverUrl         *string `json:"coverUrl,omitempty" xml:"coverUrl,omitempty"`
	DigitalHumanType *string `json:"digitalHumanType,omitempty" xml:"digitalHumanType,omitempty"`
	// example:
	//
	// F
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// offlineSynthesis
	UseScene *string `json:"useScene,omitempty" xml:"useScene,omitempty"`
	// 视频文件osskey
	//
	// example:
	//
	// path/key.mp4
	VideoOssKey *string `json:"videoOssKey,omitempty" xml:"videoOssKey,omitempty"`
}

func (s CreateAnchorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAnchorRequest) GoString() string {
	return s.String()
}

func (s *CreateAnchorRequest) GetAnchorCategory() *string {
	return s.AnchorCategory
}

func (s *CreateAnchorRequest) GetAnchorMaterialName() *string {
	return s.AnchorMaterialName
}

func (s *CreateAnchorRequest) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *CreateAnchorRequest) GetDigitalHumanType() *string {
	return s.DigitalHumanType
}

func (s *CreateAnchorRequest) GetGender() *string {
	return s.Gender
}

func (s *CreateAnchorRequest) GetUseScene() *string {
	return s.UseScene
}

func (s *CreateAnchorRequest) GetVideoOssKey() *string {
	return s.VideoOssKey
}

func (s *CreateAnchorRequest) SetAnchorCategory(v string) *CreateAnchorRequest {
	s.AnchorCategory = &v
	return s
}

func (s *CreateAnchorRequest) SetAnchorMaterialName(v string) *CreateAnchorRequest {
	s.AnchorMaterialName = &v
	return s
}

func (s *CreateAnchorRequest) SetCoverUrl(v string) *CreateAnchorRequest {
	s.CoverUrl = &v
	return s
}

func (s *CreateAnchorRequest) SetDigitalHumanType(v string) *CreateAnchorRequest {
	s.DigitalHumanType = &v
	return s
}

func (s *CreateAnchorRequest) SetGender(v string) *CreateAnchorRequest {
	s.Gender = &v
	return s
}

func (s *CreateAnchorRequest) SetUseScene(v string) *CreateAnchorRequest {
	s.UseScene = &v
	return s
}

func (s *CreateAnchorRequest) SetVideoOssKey(v string) *CreateAnchorRequest {
	s.VideoOssKey = &v
	return s
}

func (s *CreateAnchorRequest) Validate() error {
	return dara.Validate(s)
}
