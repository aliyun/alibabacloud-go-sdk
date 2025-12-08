// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRetouchSkinAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *RetouchSkinAdvanceRequest
	GetImageURLObject() io.Reader
	SetRetouchDegree(v float32) *RetouchSkinAdvanceRequest
	GetRetouchDegree() *float32
	SetWhiteningDegree(v float32) *RetouchSkinAdvanceRequest
	GetWhiteningDegree() *float32
}

type RetouchSkinAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/RetouchSkin/RetouchSkin3.png
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// 1.0
	RetouchDegree *float32 `json:"RetouchDegree,omitempty" xml:"RetouchDegree,omitempty"`
	// example:
	//
	// 1.0
	WhiteningDegree *float32 `json:"WhiteningDegree,omitempty" xml:"WhiteningDegree,omitempty"`
}

func (s RetouchSkinAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RetouchSkinAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RetouchSkinAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RetouchSkinAdvanceRequest) GetRetouchDegree() *float32 {
	return s.RetouchDegree
}

func (s *RetouchSkinAdvanceRequest) GetWhiteningDegree() *float32 {
	return s.WhiteningDegree
}

func (s *RetouchSkinAdvanceRequest) SetImageURLObject(v io.Reader) *RetouchSkinAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RetouchSkinAdvanceRequest) SetRetouchDegree(v float32) *RetouchSkinAdvanceRequest {
	s.RetouchDegree = &v
	return s
}

func (s *RetouchSkinAdvanceRequest) SetWhiteningDegree(v float32) *RetouchSkinAdvanceRequest {
	s.WhiteningDegree = &v
	return s
}

func (s *RetouchSkinAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
