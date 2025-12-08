// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLiquifyFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *LiquifyFaceRequest
	GetImageURL() *string
	SetSlimDegree(v float32) *LiquifyFaceRequest
	GetSlimDegree() *float32
}

type LiquifyFaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/LiquifyFace/LiquifyFace1.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// 1.0
	SlimDegree *float32 `json:"SlimDegree,omitempty" xml:"SlimDegree,omitempty"`
}

func (s LiquifyFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s LiquifyFaceRequest) GoString() string {
	return s.String()
}

func (s *LiquifyFaceRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *LiquifyFaceRequest) GetSlimDegree() *float32 {
	return s.SlimDegree
}

func (s *LiquifyFaceRequest) SetImageURL(v string) *LiquifyFaceRequest {
	s.ImageURL = &v
	return s
}

func (s *LiquifyFaceRequest) SetSlimDegree(v float32) *LiquifyFaceRequest {
	s.SlimDegree = &v
	return s
}

func (s *LiquifyFaceRequest) Validate() error {
	return dara.Validate(s)
}
