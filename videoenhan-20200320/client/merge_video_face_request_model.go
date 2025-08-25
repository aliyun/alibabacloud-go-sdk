// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeVideoFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddWatermark(v bool) *MergeVideoFaceRequest
	GetAddWatermark() *bool
	SetEnhance(v bool) *MergeVideoFaceRequest
	GetEnhance() *bool
	SetReferenceURL(v string) *MergeVideoFaceRequest
	GetReferenceURL() *string
	SetVideoURL(v string) *MergeVideoFaceRequest
	GetVideoURL() *string
	SetWatermarkType(v string) *MergeVideoFaceRequest
	GetWatermarkType() *string
}

type MergeVideoFaceRequest struct {
	AddWatermark *bool `json:"AddWatermark,omitempty" xml:"AddWatermark,omitempty"`
	Enhance      *bool `json:"Enhance,omitempty" xml:"Enhance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/MergeVideoFace/MergeVideoFace-pic1.png
	ReferenceURL *string `json:"ReferenceURL,omitempty" xml:"ReferenceURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/MergeVideoFace/MergeVideoFace1.mp4
	VideoURL      *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
}

func (s MergeVideoFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s MergeVideoFaceRequest) GoString() string {
	return s.String()
}

func (s *MergeVideoFaceRequest) GetAddWatermark() *bool {
	return s.AddWatermark
}

func (s *MergeVideoFaceRequest) GetEnhance() *bool {
	return s.Enhance
}

func (s *MergeVideoFaceRequest) GetReferenceURL() *string {
	return s.ReferenceURL
}

func (s *MergeVideoFaceRequest) GetVideoURL() *string {
	return s.VideoURL
}

func (s *MergeVideoFaceRequest) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *MergeVideoFaceRequest) SetAddWatermark(v bool) *MergeVideoFaceRequest {
	s.AddWatermark = &v
	return s
}

func (s *MergeVideoFaceRequest) SetEnhance(v bool) *MergeVideoFaceRequest {
	s.Enhance = &v
	return s
}

func (s *MergeVideoFaceRequest) SetReferenceURL(v string) *MergeVideoFaceRequest {
	s.ReferenceURL = &v
	return s
}

func (s *MergeVideoFaceRequest) SetVideoURL(v string) *MergeVideoFaceRequest {
	s.VideoURL = &v
	return s
}

func (s *MergeVideoFaceRequest) SetWatermarkType(v string) *MergeVideoFaceRequest {
	s.WatermarkType = &v
	return s
}

func (s *MergeVideoFaceRequest) Validate() error {
	return dara.Validate(s)
}
