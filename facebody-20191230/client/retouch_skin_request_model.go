// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetouchSkinRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *RetouchSkinRequest
	GetImageURL() *string
	SetRetouchDegree(v float32) *RetouchSkinRequest
	GetRetouchDegree() *float32
	SetWhiteningDegree(v float32) *RetouchSkinRequest
	GetWhiteningDegree() *float32
}

type RetouchSkinRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/RetouchSkin/RetouchSkin3.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// 1.0
	RetouchDegree *float32 `json:"RetouchDegree,omitempty" xml:"RetouchDegree,omitempty"`
	// example:
	//
	// 1.0
	WhiteningDegree *float32 `json:"WhiteningDegree,omitempty" xml:"WhiteningDegree,omitempty"`
}

func (s RetouchSkinRequest) String() string {
	return dara.Prettify(s)
}

func (s RetouchSkinRequest) GoString() string {
	return s.String()
}

func (s *RetouchSkinRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RetouchSkinRequest) GetRetouchDegree() *float32 {
	return s.RetouchDegree
}

func (s *RetouchSkinRequest) GetWhiteningDegree() *float32 {
	return s.WhiteningDegree
}

func (s *RetouchSkinRequest) SetImageURL(v string) *RetouchSkinRequest {
	s.ImageURL = &v
	return s
}

func (s *RetouchSkinRequest) SetRetouchDegree(v float32) *RetouchSkinRequest {
	s.RetouchDegree = &v
	return s
}

func (s *RetouchSkinRequest) SetWhiteningDegree(v float32) *RetouchSkinRequest {
	s.WhiteningDegree = &v
	return s
}

func (s *RetouchSkinRequest) Validate() error {
	return dara.Validate(s)
}
