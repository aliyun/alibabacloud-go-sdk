// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeImageFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddWatermark(v bool) *MergeImageFaceRequest
	GetAddWatermark() *bool
	SetImageURL(v string) *MergeImageFaceRequest
	GetImageURL() *string
	SetMergeInfos(v []*MergeImageFaceRequestMergeInfos) *MergeImageFaceRequest
	GetMergeInfos() []*MergeImageFaceRequestMergeInfos
	SetModelVersion(v string) *MergeImageFaceRequest
	GetModelVersion() *string
	SetTemplateId(v string) *MergeImageFaceRequest
	GetTemplateId() *string
	SetWatermarkType(v string) *MergeImageFaceRequest
	GetWatermarkType() *string
}

type MergeImageFaceRequest struct {
	// example:
	//
	// False
	AddWatermark *bool `json:"AddWatermark,omitempty" xml:"AddWatermark,omitempty"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/MergeImageFace/MergeImageFace-1.png
	ImageURL   *string                            `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	MergeInfos []*MergeImageFaceRequestMergeInfos `json:"MergeInfos,omitempty" xml:"MergeInfos,omitempty" type:"Repeated"`
	// example:
	//
	// v1
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6cd509ea-54fa-4730-8e9d-c94cadcda048
	TemplateId    *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	WatermarkType *string `json:"WatermarkType,omitempty" xml:"WatermarkType,omitempty"`
}

func (s MergeImageFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s MergeImageFaceRequest) GoString() string {
	return s.String()
}

func (s *MergeImageFaceRequest) GetAddWatermark() *bool {
	return s.AddWatermark
}

func (s *MergeImageFaceRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *MergeImageFaceRequest) GetMergeInfos() []*MergeImageFaceRequestMergeInfos {
	return s.MergeInfos
}

func (s *MergeImageFaceRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *MergeImageFaceRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *MergeImageFaceRequest) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *MergeImageFaceRequest) SetAddWatermark(v bool) *MergeImageFaceRequest {
	s.AddWatermark = &v
	return s
}

func (s *MergeImageFaceRequest) SetImageURL(v string) *MergeImageFaceRequest {
	s.ImageURL = &v
	return s
}

func (s *MergeImageFaceRequest) SetMergeInfos(v []*MergeImageFaceRequestMergeInfos) *MergeImageFaceRequest {
	s.MergeInfos = v
	return s
}

func (s *MergeImageFaceRequest) SetModelVersion(v string) *MergeImageFaceRequest {
	s.ModelVersion = &v
	return s
}

func (s *MergeImageFaceRequest) SetTemplateId(v string) *MergeImageFaceRequest {
	s.TemplateId = &v
	return s
}

func (s *MergeImageFaceRequest) SetWatermarkType(v string) *MergeImageFaceRequest {
	s.WatermarkType = &v
	return s
}

func (s *MergeImageFaceRequest) Validate() error {
	if s.MergeInfos != nil {
		for _, item := range s.MergeInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MergeImageFaceRequestMergeInfos struct {
	// example:
	//
	// http://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/ref/ref.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// 6cd509ea-54fa-4730-8e9d-c94cadcda048_0
	TemplateFaceID *string `json:"TemplateFaceID,omitempty" xml:"TemplateFaceID,omitempty"`
}

func (s MergeImageFaceRequestMergeInfos) String() string {
	return dara.Prettify(s)
}

func (s MergeImageFaceRequestMergeInfos) GoString() string {
	return s.String()
}

func (s *MergeImageFaceRequestMergeInfos) GetImageURL() *string {
	return s.ImageURL
}

func (s *MergeImageFaceRequestMergeInfos) GetTemplateFaceID() *string {
	return s.TemplateFaceID
}

func (s *MergeImageFaceRequestMergeInfos) SetImageURL(v string) *MergeImageFaceRequestMergeInfos {
	s.ImageURL = &v
	return s
}

func (s *MergeImageFaceRequestMergeInfos) SetTemplateFaceID(v string) *MergeImageFaceRequestMergeInfos {
	s.TemplateFaceID = &v
	return s
}

func (s *MergeImageFaceRequestMergeInfos) Validate() error {
	return dara.Validate(s)
}
