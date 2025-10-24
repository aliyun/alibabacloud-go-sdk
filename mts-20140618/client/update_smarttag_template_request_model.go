// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmarttagTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalyseTypes(v string) *UpdateSmarttagTemplateRequest
	GetAnalyseTypes() *string
	SetFaceCategoryIds(v string) *UpdateSmarttagTemplateRequest
	GetFaceCategoryIds() *string
	SetFaceCustomParamsConfig(v string) *UpdateSmarttagTemplateRequest
	GetFaceCustomParamsConfig() *string
	SetIndustry(v string) *UpdateSmarttagTemplateRequest
	GetIndustry() *string
	SetIsDefault(v bool) *UpdateSmarttagTemplateRequest
	GetIsDefault() *bool
	SetKeywordConfig(v string) *UpdateSmarttagTemplateRequest
	GetKeywordConfig() *string
	SetKnowledgeConfig(v string) *UpdateSmarttagTemplateRequest
	GetKnowledgeConfig() *string
	SetLabelType(v string) *UpdateSmarttagTemplateRequest
	GetLabelType() *string
	SetLabelVersion(v string) *UpdateSmarttagTemplateRequest
	GetLabelVersion() *string
	SetLandmarkGroupIds(v string) *UpdateSmarttagTemplateRequest
	GetLandmarkGroupIds() *string
	SetObjectGroupIds(v string) *UpdateSmarttagTemplateRequest
	GetObjectGroupIds() *string
	SetOwnerAccount(v string) *UpdateSmarttagTemplateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateSmarttagTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateSmarttagTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateSmarttagTemplateRequest
	GetResourceOwnerId() *int64
	SetScene(v string) *UpdateSmarttagTemplateRequest
	GetScene() *string
	SetTemplateConfig(v string) *UpdateSmarttagTemplateRequest
	GetTemplateConfig() *string
	SetTemplateId(v string) *UpdateSmarttagTemplateRequest
	GetTemplateId() *string
	SetTemplateName(v string) *UpdateSmarttagTemplateRequest
	GetTemplateName() *string
}

type UpdateSmarttagTemplateRequest struct {
	// example:
	//
	// ocr,asr
	AnalyseTypes *string `json:"AnalyseTypes,omitempty" xml:"AnalyseTypes,omitempty"`
	// example:
	//
	// celebrity
	FaceCategoryIds *string `json:"FaceCategoryIds,omitempty" xml:"FaceCategoryIds,omitempty"`
	// example:
	//
	// { "faceDetThreshold":0.999, "faceRegThreshold":0.9 }
	FaceCustomParamsConfig *string `json:"FaceCustomParamsConfig,omitempty" xml:"FaceCustomParamsConfig,omitempty"`
	// example:
	//
	// common
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// { "type": "name,location,organization,other" }
	KeywordConfig *string `json:"KeywordConfig,omitempty" xml:"KeywordConfig,omitempty"`
	// example:
	//
	// { "movie":"name,alias,chnl,genre", "music":"songName,artistName", "person":"name,gender" }
	KnowledgeConfig *string `json:"KnowledgeConfig,omitempty" xml:"KnowledgeConfig,omitempty"`
	// example:
	//
	// hmi
	LabelType *string `json:"LabelType,omitempty" xml:"LabelType,omitempty"`
	// example:
	//
	// 1.0
	LabelVersion *string `json:"LabelVersion,omitempty" xml:"LabelVersion,omitempty"`
	// example:
	//
	// common
	LandmarkGroupIds *string `json:"LandmarkGroupIds,omitempty" xml:"LandmarkGroupIds,omitempty"`
	// example:
	//
	// general,item,weapon,animal
	ObjectGroupIds       *string `json:"ObjectGroupIds,omitempty" xml:"ObjectGroupIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// search
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 05de22f255284c7a8d2aab535dde****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// template-example-****
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdateSmarttagTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmarttagTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmarttagTemplateRequest) GetAnalyseTypes() *string {
	return s.AnalyseTypes
}

func (s *UpdateSmarttagTemplateRequest) GetFaceCategoryIds() *string {
	return s.FaceCategoryIds
}

func (s *UpdateSmarttagTemplateRequest) GetFaceCustomParamsConfig() *string {
	return s.FaceCustomParamsConfig
}

func (s *UpdateSmarttagTemplateRequest) GetIndustry() *string {
	return s.Industry
}

func (s *UpdateSmarttagTemplateRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *UpdateSmarttagTemplateRequest) GetKeywordConfig() *string {
	return s.KeywordConfig
}

func (s *UpdateSmarttagTemplateRequest) GetKnowledgeConfig() *string {
	return s.KnowledgeConfig
}

func (s *UpdateSmarttagTemplateRequest) GetLabelType() *string {
	return s.LabelType
}

func (s *UpdateSmarttagTemplateRequest) GetLabelVersion() *string {
	return s.LabelVersion
}

func (s *UpdateSmarttagTemplateRequest) GetLandmarkGroupIds() *string {
	return s.LandmarkGroupIds
}

func (s *UpdateSmarttagTemplateRequest) GetObjectGroupIds() *string {
	return s.ObjectGroupIds
}

func (s *UpdateSmarttagTemplateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateSmarttagTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateSmarttagTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateSmarttagTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateSmarttagTemplateRequest) GetScene() *string {
	return s.Scene
}

func (s *UpdateSmarttagTemplateRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *UpdateSmarttagTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateSmarttagTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateSmarttagTemplateRequest) SetAnalyseTypes(v string) *UpdateSmarttagTemplateRequest {
	s.AnalyseTypes = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetFaceCategoryIds(v string) *UpdateSmarttagTemplateRequest {
	s.FaceCategoryIds = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetFaceCustomParamsConfig(v string) *UpdateSmarttagTemplateRequest {
	s.FaceCustomParamsConfig = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetIndustry(v string) *UpdateSmarttagTemplateRequest {
	s.Industry = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetIsDefault(v bool) *UpdateSmarttagTemplateRequest {
	s.IsDefault = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetKeywordConfig(v string) *UpdateSmarttagTemplateRequest {
	s.KeywordConfig = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetKnowledgeConfig(v string) *UpdateSmarttagTemplateRequest {
	s.KnowledgeConfig = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetLabelType(v string) *UpdateSmarttagTemplateRequest {
	s.LabelType = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetLabelVersion(v string) *UpdateSmarttagTemplateRequest {
	s.LabelVersion = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetLandmarkGroupIds(v string) *UpdateSmarttagTemplateRequest {
	s.LandmarkGroupIds = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetObjectGroupIds(v string) *UpdateSmarttagTemplateRequest {
	s.ObjectGroupIds = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetOwnerAccount(v string) *UpdateSmarttagTemplateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetOwnerId(v int64) *UpdateSmarttagTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetResourceOwnerAccount(v string) *UpdateSmarttagTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetResourceOwnerId(v int64) *UpdateSmarttagTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetScene(v string) *UpdateSmarttagTemplateRequest {
	s.Scene = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetTemplateConfig(v string) *UpdateSmarttagTemplateRequest {
	s.TemplateConfig = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetTemplateId(v string) *UpdateSmarttagTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) SetTemplateName(v string) *UpdateSmarttagTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateSmarttagTemplateRequest) Validate() error {
	return dara.Validate(s)
}
