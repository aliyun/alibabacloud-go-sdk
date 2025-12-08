// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iGenerateHumanSketchStyleAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *GenerateHumanSketchStyleAdvanceRequest
	GetImageURLObject() io.Reader
	SetReturnType(v string) *GenerateHumanSketchStyleAdvanceRequest
	GetReturnType() *string
}

type GenerateHumanSketchStyleAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/GenerateHumanSketchStyle/GenerateHumanSketchStyle7.png
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// head
	ReturnType *string `json:"ReturnType,omitempty" xml:"ReturnType,omitempty"`
}

func (s GenerateHumanSketchStyleAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateHumanSketchStyleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateHumanSketchStyleAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *GenerateHumanSketchStyleAdvanceRequest) GetReturnType() *string {
	return s.ReturnType
}

func (s *GenerateHumanSketchStyleAdvanceRequest) SetImageURLObject(v io.Reader) *GenerateHumanSketchStyleAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *GenerateHumanSketchStyleAdvanceRequest) SetReturnType(v string) *GenerateHumanSketchStyleAdvanceRequest {
	s.ReturnType = &v
	return s
}

func (s *GenerateHumanSketchStyleAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
