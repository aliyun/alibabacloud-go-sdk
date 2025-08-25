// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iMergeVideoModelFaceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddWatermark(v bool) *MergeVideoModelFaceAdvanceRequest
	GetAddWatermark() *bool
	SetEnhance(v bool) *MergeVideoModelFaceAdvanceRequest
	GetEnhance() *bool
	SetFaceImageURLObject(v io.Reader) *MergeVideoModelFaceAdvanceRequest
	GetFaceImageURLObject() io.Reader
	SetMergeInfos(v []*MergeVideoModelFaceAdvanceRequestMergeInfos) *MergeVideoModelFaceAdvanceRequest
	GetMergeInfos() []*MergeVideoModelFaceAdvanceRequestMergeInfos
	SetTemplateId(v string) *MergeVideoModelFaceAdvanceRequest
	GetTemplateId() *string
	SetWatermarkType(v string) *MergeVideoModelFaceAdvanceRequest
	GetWatermarkType() *string
}

type MergeVideoModelFaceAdvanceRequest struct {
	AddWatermark *bool `json:"AddWatermark,omitempty" xml:"AddWatermark,omitempty"`
	Enhance      *bool `json:"Enhance,omitempty" xml:"Enhance,omitempty"`
	// example:
	//
	// https://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/ref/ref.jpg
	FaceImageURLObject io.Reader                                      `json:"FaceImageURL,omitempty" xml:"FaceImageURL,omitempty"`
	MergeInfos         []*MergeVideoModelFaceAdvanceRequestMergeInfos `json:"MergeInfos,omitempty" xml:"MergeInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 3bf2418c-7adf-4002-a9d6-2f7cf1889c0d
	TemplateId    *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
}

func (s MergeVideoModelFaceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s MergeVideoModelFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *MergeVideoModelFaceAdvanceRequest) GetAddWatermark() *bool {
	return s.AddWatermark
}

func (s *MergeVideoModelFaceAdvanceRequest) GetEnhance() *bool {
	return s.Enhance
}

func (s *MergeVideoModelFaceAdvanceRequest) GetFaceImageURLObject() io.Reader {
	return s.FaceImageURLObject
}

func (s *MergeVideoModelFaceAdvanceRequest) GetMergeInfos() []*MergeVideoModelFaceAdvanceRequestMergeInfos {
	return s.MergeInfos
}

func (s *MergeVideoModelFaceAdvanceRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *MergeVideoModelFaceAdvanceRequest) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *MergeVideoModelFaceAdvanceRequest) SetAddWatermark(v bool) *MergeVideoModelFaceAdvanceRequest {
	s.AddWatermark = &v
	return s
}

func (s *MergeVideoModelFaceAdvanceRequest) SetEnhance(v bool) *MergeVideoModelFaceAdvanceRequest {
	s.Enhance = &v
	return s
}

func (s *MergeVideoModelFaceAdvanceRequest) SetFaceImageURLObject(v io.Reader) *MergeVideoModelFaceAdvanceRequest {
	s.FaceImageURLObject = v
	return s
}

func (s *MergeVideoModelFaceAdvanceRequest) SetMergeInfos(v []*MergeVideoModelFaceAdvanceRequestMergeInfos) *MergeVideoModelFaceAdvanceRequest {
	s.MergeInfos = v
	return s
}

func (s *MergeVideoModelFaceAdvanceRequest) SetTemplateId(v string) *MergeVideoModelFaceAdvanceRequest {
	s.TemplateId = &v
	return s
}

func (s *MergeVideoModelFaceAdvanceRequest) SetWatermarkType(v string) *MergeVideoModelFaceAdvanceRequest {
	s.WatermarkType = &v
	return s
}

func (s *MergeVideoModelFaceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}

type MergeVideoModelFaceAdvanceRequestMergeInfos struct {
	ImageURL        *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	TemplateFaceID  *string `json:"TemplateFaceID,omitempty" xml:"TemplateFaceID,omitempty"`
	TemplateFaceURL *string `json:"TemplateFaceURL,omitempty" xml:"TemplateFaceURL,omitempty"`
}

func (s MergeVideoModelFaceAdvanceRequestMergeInfos) String() string {
	return dara.Prettify(s)
}

func (s MergeVideoModelFaceAdvanceRequestMergeInfos) GoString() string {
	return s.String()
}

func (s *MergeVideoModelFaceAdvanceRequestMergeInfos) GetImageURL() *string {
	return s.ImageURL
}

func (s *MergeVideoModelFaceAdvanceRequestMergeInfos) GetTemplateFaceID() *string {
	return s.TemplateFaceID
}

func (s *MergeVideoModelFaceAdvanceRequestMergeInfos) GetTemplateFaceURL() *string {
	return s.TemplateFaceURL
}

func (s *MergeVideoModelFaceAdvanceRequestMergeInfos) SetImageURL(v string) *MergeVideoModelFaceAdvanceRequestMergeInfos {
	s.ImageURL = &v
	return s
}

func (s *MergeVideoModelFaceAdvanceRequestMergeInfos) SetTemplateFaceID(v string) *MergeVideoModelFaceAdvanceRequestMergeInfos {
	s.TemplateFaceID = &v
	return s
}

func (s *MergeVideoModelFaceAdvanceRequestMergeInfos) SetTemplateFaceURL(v string) *MergeVideoModelFaceAdvanceRequestMergeInfos {
	s.TemplateFaceURL = &v
	return s
}

func (s *MergeVideoModelFaceAdvanceRequestMergeInfos) Validate() error {
	return dara.Validate(s)
}
