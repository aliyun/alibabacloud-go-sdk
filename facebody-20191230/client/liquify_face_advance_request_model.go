// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iLiquifyFaceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *LiquifyFaceAdvanceRequest
	GetImageURLObject() io.Reader
	SetSlimDegree(v float32) *LiquifyFaceAdvanceRequest
	GetSlimDegree() *float32
}

type LiquifyFaceAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/LiquifyFace/LiquifyFace1.png
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// 1.0
	SlimDegree *float32 `json:"SlimDegree,omitempty" xml:"SlimDegree,omitempty"`
}

func (s LiquifyFaceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s LiquifyFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *LiquifyFaceAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *LiquifyFaceAdvanceRequest) GetSlimDegree() *float32 {
	return s.SlimDegree
}

func (s *LiquifyFaceAdvanceRequest) SetImageURLObject(v io.Reader) *LiquifyFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *LiquifyFaceAdvanceRequest) SetSlimDegree(v float32) *LiquifyFaceAdvanceRequest {
	s.SlimDegree = &v
	return s
}

func (s *LiquifyFaceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
