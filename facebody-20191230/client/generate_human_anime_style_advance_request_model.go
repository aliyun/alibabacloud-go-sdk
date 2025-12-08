// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iGenerateHumanAnimeStyleAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgoType(v string) *GenerateHumanAnimeStyleAdvanceRequest
	GetAlgoType() *string
	SetImageURLObject(v io.Reader) *GenerateHumanAnimeStyleAdvanceRequest
	GetImageURLObject() io.Reader
}

type GenerateHumanAnimeStyleAdvanceRequest struct {
	// example:
	//
	// anime
	AlgoType *string `json:"AlgoType,omitempty" xml:"AlgoType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/GenerateHumanAnimeStyle/GenerateHumanAnimeStyle8.png
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s GenerateHumanAnimeStyleAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateHumanAnimeStyleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleAdvanceRequest) GetAlgoType() *string {
	return s.AlgoType
}

func (s *GenerateHumanAnimeStyleAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *GenerateHumanAnimeStyleAdvanceRequest) SetAlgoType(v string) *GenerateHumanAnimeStyleAdvanceRequest {
	s.AlgoType = &v
	return s
}

func (s *GenerateHumanAnimeStyleAdvanceRequest) SetImageURLObject(v io.Reader) *GenerateHumanAnimeStyleAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *GenerateHumanAnimeStyleAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
