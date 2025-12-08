// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateHumanSketchStyleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *GenerateHumanSketchStyleRequest
	GetImageURL() *string
	SetReturnType(v string) *GenerateHumanSketchStyleRequest
	GetReturnType() *string
}

type GenerateHumanSketchStyleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/GenerateHumanSketchStyle/GenerateHumanSketchStyle7.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// head
	ReturnType *string `json:"ReturnType,omitempty" xml:"ReturnType,omitempty"`
}

func (s GenerateHumanSketchStyleRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateHumanSketchStyleRequest) GoString() string {
	return s.String()
}

func (s *GenerateHumanSketchStyleRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *GenerateHumanSketchStyleRequest) GetReturnType() *string {
	return s.ReturnType
}

func (s *GenerateHumanSketchStyleRequest) SetImageURL(v string) *GenerateHumanSketchStyleRequest {
	s.ImageURL = &v
	return s
}

func (s *GenerateHumanSketchStyleRequest) SetReturnType(v string) *GenerateHumanSketchStyleRequest {
	s.ReturnType = &v
	return s
}

func (s *GenerateHumanSketchStyleRequest) Validate() error {
	return dara.Validate(s)
}
