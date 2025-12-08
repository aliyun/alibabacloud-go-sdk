// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iMergeImageFaceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddWatermark(v bool) *MergeImageFaceAdvanceRequest
	GetAddWatermark() *bool
	SetImageURLObject(v io.Reader) *MergeImageFaceAdvanceRequest
	GetImageURLObject() io.Reader
	SetMergeInfos(v []*MergeImageFaceAdvanceRequestMergeInfos) *MergeImageFaceAdvanceRequest
	GetMergeInfos() []*MergeImageFaceAdvanceRequestMergeInfos
	SetModelVersion(v string) *MergeImageFaceAdvanceRequest
	GetModelVersion() *string
	SetTemplateId(v string) *MergeImageFaceAdvanceRequest
	GetTemplateId() *string
	SetWatermarkType(v string) *MergeImageFaceAdvanceRequest
	GetWatermarkType() *string
}

type MergeImageFaceAdvanceRequest struct {
	// example:
	//
	// False
	AddWatermark *bool `json:"AddWatermark,omitempty" xml:"AddWatermark,omitempty"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/MergeImageFace/MergeImageFace-1.png
	ImageURLObject io.Reader                                 `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	MergeInfos     []*MergeImageFaceAdvanceRequestMergeInfos `json:"MergeInfos,omitempty" xml:"MergeInfos,omitempty" type:"Repeated"`
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

func (s MergeImageFaceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s MergeImageFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *MergeImageFaceAdvanceRequest) GetAddWatermark() *bool {
	return s.AddWatermark
}

func (s *MergeImageFaceAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *MergeImageFaceAdvanceRequest) GetMergeInfos() []*MergeImageFaceAdvanceRequestMergeInfos {
	return s.MergeInfos
}

func (s *MergeImageFaceAdvanceRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *MergeImageFaceAdvanceRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *MergeImageFaceAdvanceRequest) GetWatermarkType() *string {
	return s.WatermarkType
}

func (s *MergeImageFaceAdvanceRequest) SetAddWatermark(v bool) *MergeImageFaceAdvanceRequest {
	s.AddWatermark = &v
	return s
}

func (s *MergeImageFaceAdvanceRequest) SetImageURLObject(v io.Reader) *MergeImageFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *MergeImageFaceAdvanceRequest) SetMergeInfos(v []*MergeImageFaceAdvanceRequestMergeInfos) *MergeImageFaceAdvanceRequest {
	s.MergeInfos = v
	return s
}

func (s *MergeImageFaceAdvanceRequest) SetModelVersion(v string) *MergeImageFaceAdvanceRequest {
	s.ModelVersion = &v
	return s
}

func (s *MergeImageFaceAdvanceRequest) SetTemplateId(v string) *MergeImageFaceAdvanceRequest {
	s.TemplateId = &v
	return s
}

func (s *MergeImageFaceAdvanceRequest) SetWatermarkType(v string) *MergeImageFaceAdvanceRequest {
	s.WatermarkType = &v
	return s
}

func (s *MergeImageFaceAdvanceRequest) Validate() error {
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

type MergeImageFaceAdvanceRequestMergeInfos struct {
	// example:
	//
	// http://invi-label.oss-cn-shanghai.aliyuncs.com/label/temp/faceswap/ref/ref.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// 6cd509ea-54fa-4730-8e9d-c94cadcda048_0
	TemplateFaceID *string `json:"TemplateFaceID,omitempty" xml:"TemplateFaceID,omitempty"`
}

func (s MergeImageFaceAdvanceRequestMergeInfos) String() string {
	return dara.Prettify(s)
}

func (s MergeImageFaceAdvanceRequestMergeInfos) GoString() string {
	return s.String()
}

func (s *MergeImageFaceAdvanceRequestMergeInfos) GetImageURL() *string {
	return s.ImageURL
}

func (s *MergeImageFaceAdvanceRequestMergeInfos) GetTemplateFaceID() *string {
	return s.TemplateFaceID
}

func (s *MergeImageFaceAdvanceRequestMergeInfos) SetImageURL(v string) *MergeImageFaceAdvanceRequestMergeInfos {
	s.ImageURL = &v
	return s
}

func (s *MergeImageFaceAdvanceRequestMergeInfos) SetTemplateFaceID(v string) *MergeImageFaceAdvanceRequestMergeInfos {
	s.TemplateFaceID = &v
	return s
}

func (s *MergeImageFaceAdvanceRequestMergeInfos) Validate() error {
	return dara.Validate(s)
}
