// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSmarttagTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalyseTypes(v string) *AddSmarttagTemplateRequest
	GetAnalyseTypes() *string
	SetFaceCategoryIds(v string) *AddSmarttagTemplateRequest
	GetFaceCategoryIds() *string
	SetFaceCustomParamsConfig(v string) *AddSmarttagTemplateRequest
	GetFaceCustomParamsConfig() *string
	SetIndustry(v string) *AddSmarttagTemplateRequest
	GetIndustry() *string
	SetIsDefault(v bool) *AddSmarttagTemplateRequest
	GetIsDefault() *bool
	SetKeywordConfig(v string) *AddSmarttagTemplateRequest
	GetKeywordConfig() *string
	SetKnowledgeConfig(v string) *AddSmarttagTemplateRequest
	GetKnowledgeConfig() *string
	SetLabelCustomCategoryIds(v string) *AddSmarttagTemplateRequest
	GetLabelCustomCategoryIds() *string
	SetLabelCustomParamsConfig(v string) *AddSmarttagTemplateRequest
	GetLabelCustomParamsConfig() *string
	SetLabelType(v string) *AddSmarttagTemplateRequest
	GetLabelType() *string
	SetLabelVersion(v string) *AddSmarttagTemplateRequest
	GetLabelVersion() *string
	SetLandmarkGroupIds(v string) *AddSmarttagTemplateRequest
	GetLandmarkGroupIds() *string
	SetObjectGroupIds(v string) *AddSmarttagTemplateRequest
	GetObjectGroupIds() *string
	SetOwnerAccount(v string) *AddSmarttagTemplateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddSmarttagTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddSmarttagTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddSmarttagTemplateRequest
	GetResourceOwnerId() *int64
	SetScene(v string) *AddSmarttagTemplateRequest
	GetScene() *string
	SetTemplateConfig(v string) *AddSmarttagTemplateRequest
	GetTemplateConfig() *string
	SetTemplateName(v string) *AddSmarttagTemplateRequest
	GetTemplateName() *string
}

type AddSmarttagTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ocr
	AnalyseTypes *string `json:"AnalyseTypes,omitempty" xml:"AnalyseTypes,omitempty"`
	// example:
	//
	// celebrity
	FaceCategoryIds *string `json:"FaceCategoryIds,omitempty" xml:"FaceCategoryIds,omitempty"`
	// example:
	//
	// { "faceDetThreshold":0.999, "faceRegThreshold":0.9 }
	FaceCustomParamsConfig *string `json:"FaceCustomParamsConfig,omitempty" xml:"FaceCustomParamsConfig,omitempty"`
	// This parameter is required.
	//
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
	// "type": "name,location,organization,other" }
	KeywordConfig *string `json:"KeywordConfig,omitempty" xml:"KeywordConfig,omitempty"`
	// example:
	//
	// { "movie":"name,alias,chnl,genre", "music":"songName,artistName", "person":"name,gender" }
	KnowledgeConfig         *string `json:"KnowledgeConfig,omitempty" xml:"KnowledgeConfig,omitempty"`
	LabelCustomCategoryIds  *string `json:"LabelCustomCategoryIds,omitempty" xml:"LabelCustomCategoryIds,omitempty"`
	LabelCustomParamsConfig *string `json:"LabelCustomParamsConfig,omitempty" xml:"LabelCustomParamsConfig,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// search
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// template-example-****
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s AddSmarttagTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSmarttagTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddSmarttagTemplateRequest) GetAnalyseTypes() *string {
	return s.AnalyseTypes
}

func (s *AddSmarttagTemplateRequest) GetFaceCategoryIds() *string {
	return s.FaceCategoryIds
}

func (s *AddSmarttagTemplateRequest) GetFaceCustomParamsConfig() *string {
	return s.FaceCustomParamsConfig
}

func (s *AddSmarttagTemplateRequest) GetIndustry() *string {
	return s.Industry
}

func (s *AddSmarttagTemplateRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *AddSmarttagTemplateRequest) GetKeywordConfig() *string {
	return s.KeywordConfig
}

func (s *AddSmarttagTemplateRequest) GetKnowledgeConfig() *string {
	return s.KnowledgeConfig
}

func (s *AddSmarttagTemplateRequest) GetLabelCustomCategoryIds() *string {
	return s.LabelCustomCategoryIds
}

func (s *AddSmarttagTemplateRequest) GetLabelCustomParamsConfig() *string {
	return s.LabelCustomParamsConfig
}

func (s *AddSmarttagTemplateRequest) GetLabelType() *string {
	return s.LabelType
}

func (s *AddSmarttagTemplateRequest) GetLabelVersion() *string {
	return s.LabelVersion
}

func (s *AddSmarttagTemplateRequest) GetLandmarkGroupIds() *string {
	return s.LandmarkGroupIds
}

func (s *AddSmarttagTemplateRequest) GetObjectGroupIds() *string {
	return s.ObjectGroupIds
}

func (s *AddSmarttagTemplateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddSmarttagTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddSmarttagTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddSmarttagTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddSmarttagTemplateRequest) GetScene() *string {
	return s.Scene
}

func (s *AddSmarttagTemplateRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *AddSmarttagTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *AddSmarttagTemplateRequest) SetAnalyseTypes(v string) *AddSmarttagTemplateRequest {
	s.AnalyseTypes = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetFaceCategoryIds(v string) *AddSmarttagTemplateRequest {
	s.FaceCategoryIds = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetFaceCustomParamsConfig(v string) *AddSmarttagTemplateRequest {
	s.FaceCustomParamsConfig = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetIndustry(v string) *AddSmarttagTemplateRequest {
	s.Industry = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetIsDefault(v bool) *AddSmarttagTemplateRequest {
	s.IsDefault = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetKeywordConfig(v string) *AddSmarttagTemplateRequest {
	s.KeywordConfig = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetKnowledgeConfig(v string) *AddSmarttagTemplateRequest {
	s.KnowledgeConfig = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetLabelCustomCategoryIds(v string) *AddSmarttagTemplateRequest {
	s.LabelCustomCategoryIds = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetLabelCustomParamsConfig(v string) *AddSmarttagTemplateRequest {
	s.LabelCustomParamsConfig = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetLabelType(v string) *AddSmarttagTemplateRequest {
	s.LabelType = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetLabelVersion(v string) *AddSmarttagTemplateRequest {
	s.LabelVersion = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetLandmarkGroupIds(v string) *AddSmarttagTemplateRequest {
	s.LandmarkGroupIds = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetObjectGroupIds(v string) *AddSmarttagTemplateRequest {
	s.ObjectGroupIds = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetOwnerAccount(v string) *AddSmarttagTemplateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetOwnerId(v int64) *AddSmarttagTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetResourceOwnerAccount(v string) *AddSmarttagTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetResourceOwnerId(v int64) *AddSmarttagTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetScene(v string) *AddSmarttagTemplateRequest {
	s.Scene = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetTemplateConfig(v string) *AddSmarttagTemplateRequest {
	s.TemplateConfig = &v
	return s
}

func (s *AddSmarttagTemplateRequest) SetTemplateName(v string) *AddSmarttagTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *AddSmarttagTemplateRequest) Validate() error {
	return dara.Validate(s)
}
