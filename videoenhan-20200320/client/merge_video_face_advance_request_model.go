// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iMergeVideoFaceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddWatermark(v bool) *MergeVideoFaceAdvanceRequest
	GetAddWatermark() *bool
	SetEnhance(v bool) *MergeVideoFaceAdvanceRequest
	GetEnhance() *bool
	SetReferenceURLObject(v io.Reader) *MergeVideoFaceAdvanceRequest
	GetReferenceURLObject() io.Reader
	SetVideoURLObject(v io.Reader) *MergeVideoFaceAdvanceRequest
	GetVideoURLObject() io.Reader
	SetWatermarkType(v string) *MergeVideoFaceAdvanceRequest
	GetWatermarkType() *string
}

type MergeVideoFaceAdvanceRequest struct {
	AddWatermark *bool `json:"AddWatermark,omitempty" xml:"AddWatermark,omitempty"`
	Enhance      *bool `json:"Enhance,omitempty" xml:"Enhance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/MergeVideoFace/MergeVideoFace-pic1.png
	ReferenceURLObject io.Reader `json:"ReferenceURL,omitempty" xml:"ReferenceURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/MergeVideoFace/MergeVideoFace1.mp4
	VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
	WatermarkType  *string   `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
}

func (s MergeVideoFaceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s MergeVideoFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *MergeVideoFaceAdvanceRequest) GetAddWatermark() *bool {
	return s.AddWatermark
}

func (s *MergeVideoFaceAdvanceRequest) GetEnhance() *bool {
	return s.Enhance
}

func (s *MergeVideoFaceAdvanceRequest) GetReferenceURLObject() io.Reader {
	return s.ReferenceURLObject
}

func (s *MergeVideoFaceAdvanceRequest) GetVideoURLObject() io.Reader {
	return s.VideoURLObject
}

func (s *MergeVideoFaceAdvanceRequest) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *MergeVideoFaceAdvanceRequest) SetAddWatermark(v bool) *MergeVideoFaceAdvanceRequest {
	s.AddWatermark = &v
	return s
}

func (s *MergeVideoFaceAdvanceRequest) SetEnhance(v bool) *MergeVideoFaceAdvanceRequest {
	s.Enhance = &v
	return s
}

func (s *MergeVideoFaceAdvanceRequest) SetReferenceURLObject(v io.Reader) *MergeVideoFaceAdvanceRequest {
	s.ReferenceURLObject = v
	return s
}

func (s *MergeVideoFaceAdvanceRequest) SetVideoURLObject(v io.Reader) *MergeVideoFaceAdvanceRequest {
	s.VideoURLObject = v
	return s
}

func (s *MergeVideoFaceAdvanceRequest) SetWatermarkType(v string) *MergeVideoFaceAdvanceRequest {
	s.WatermarkType = &v
	return s
}

func (s *MergeVideoFaceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
