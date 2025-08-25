// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeVideoModelFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddWatermark(v bool) *MergeVideoModelFaceRequest
	GetAddWatermark() *bool
	SetEnhance(v bool) *MergeVideoModelFaceRequest
	GetEnhance() *bool
	SetFaceImageURL(v string) *MergeVideoModelFaceRequest
	GetFaceImageURL() *string
	SetMergeInfos(v []*MergeVideoModelFaceRequestMergeInfos) *MergeVideoModelFaceRequest
	GetMergeInfos() []*MergeVideoModelFaceRequestMergeInfos
	SetTemplateId(v string) *MergeVideoModelFaceRequest
	GetTemplateId() *string
	SetWatermarkType(v string) *MergeVideoModelFaceRequest
	GetWatermarkType() *string
}

type MergeVideoModelFaceRequest struct {
	AddWatermark *bool `json:"AddWatermark,omitempty" xml:"AddWatermark,omitempty"`
	Enhance      *bool `json:"Enhance,omitempty" xml:"Enhance,omitempty"`
	// example:
	//
	// https://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/ref/ref.jpg
	FaceImageURL *string                                 `json:"FaceImageURL,omitempty" xml:"FaceImageURL,omitempty"`
	MergeInfos   []*MergeVideoModelFaceRequestMergeInfos `json:"MergeInfos,omitempty" xml:"MergeInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 3bf2418c-7adf-4002-a9d6-2f7cf1889c0d
	TemplateId    *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
}

func (s MergeVideoModelFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s MergeVideoModelFaceRequest) GoString() string {
	return s.String()
}

func (s *MergeVideoModelFaceRequest) GetAddWatermark() *bool {
	return s.AddWatermark
}

func (s *MergeVideoModelFaceRequest) GetEnhance() *bool {
	return s.Enhance
}

func (s *MergeVideoModelFaceRequest) GetFaceImageURL() *string {
	return s.FaceImageURL
}

func (s *MergeVideoModelFaceRequest) GetMergeInfos() []*MergeVideoModelFaceRequestMergeInfos {
	return s.MergeInfos
}

func (s *MergeVideoModelFaceRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *MergeVideoModelFaceRequest) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *MergeVideoModelFaceRequest) SetAddWatermark(v bool) *MergeVideoModelFaceRequest {
	s.AddWatermark = &v
	return s
}

func (s *MergeVideoModelFaceRequest) SetEnhance(v bool) *MergeVideoModelFaceRequest {
	s.Enhance = &v
	return s
}

func (s *MergeVideoModelFaceRequest) SetFaceImageURL(v string) *MergeVideoModelFaceRequest {
	s.FaceImageURL = &v
	return s
}

func (s *MergeVideoModelFaceRequest) SetMergeInfos(v []*MergeVideoModelFaceRequestMergeInfos) *MergeVideoModelFaceRequest {
	s.MergeInfos = v
	return s
}

func (s *MergeVideoModelFaceRequest) SetTemplateId(v string) *MergeVideoModelFaceRequest {
	s.TemplateId = &v
	return s
}

func (s *MergeVideoModelFaceRequest) SetWatermarkType(v string) *MergeVideoModelFaceRequest {
	s.WatermarkType = &v
	return s
}

func (s *MergeVideoModelFaceRequest) Validate() error {
	return dara.Validate(s)
}

type MergeVideoModelFaceRequestMergeInfos struct {
	ImageURL        *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	TemplateFaceID  *string `json:"TemplateFaceID,omitempty" xml:"TemplateFaceID,omitempty"`
	TemplateFaceURL *string `json:"TemplateFaceURL,omitempty" xml:"TemplateFaceURL,omitempty"`
}

func (s MergeVideoModelFaceRequestMergeInfos) String() string {
	return dara.Prettify(s)
}

func (s MergeVideoModelFaceRequestMergeInfos) GoString() string {
	return s.String()
}

func (s *MergeVideoModelFaceRequestMergeInfos) GetImageURL() *string {
	return s.ImageURL
}

func (s *MergeVideoModelFaceRequestMergeInfos) GetTemplateFaceID() *string {
	return s.TemplateFaceID
}

func (s *MergeVideoModelFaceRequestMergeInfos) GetTemplateFaceURL() *string {
	return s.TemplateFaceURL
}

func (s *MergeVideoModelFaceRequestMergeInfos) SetImageURL(v string) *MergeVideoModelFaceRequestMergeInfos {
	s.ImageURL = &v
	return s
}

func (s *MergeVideoModelFaceRequestMergeInfos) SetTemplateFaceID(v string) *MergeVideoModelFaceRequestMergeInfos {
	s.TemplateFaceID = &v
	return s
}

func (s *MergeVideoModelFaceRequestMergeInfos) SetTemplateFaceURL(v string) *MergeVideoModelFaceRequestMergeInfos {
	s.TemplateFaceURL = &v
	return s
}

func (s *MergeVideoModelFaceRequestMergeInfos) Validate() error {
	return dara.Validate(s)
}
