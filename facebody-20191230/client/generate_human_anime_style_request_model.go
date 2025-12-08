// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateHumanAnimeStyleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgoType(v string) *GenerateHumanAnimeStyleRequest
	GetAlgoType() *string
	SetImageURL(v string) *GenerateHumanAnimeStyleRequest
	GetImageURL() *string
}

type GenerateHumanAnimeStyleRequest struct {
	// example:
	//
	// anime
	AlgoType *string `json:"AlgoType,omitempty" xml:"AlgoType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/GenerateHumanAnimeStyle/GenerateHumanAnimeStyle8.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s GenerateHumanAnimeStyleRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateHumanAnimeStyleRequest) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleRequest) GetAlgoType() *string {
	return s.AlgoType
}

func (s *GenerateHumanAnimeStyleRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *GenerateHumanAnimeStyleRequest) SetAlgoType(v string) *GenerateHumanAnimeStyleRequest {
	s.AlgoType = &v
	return s
}

func (s *GenerateHumanAnimeStyleRequest) SetImageURL(v string) *GenerateHumanAnimeStyleRequest {
	s.ImageURL = &v
	return s
}

func (s *GenerateHumanAnimeStyleRequest) Validate() error {
	return dara.Validate(s)
}
