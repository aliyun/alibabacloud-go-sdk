// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmarttagTemplateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QuerySmarttagTemplateListResponseBody
	GetRequestId() *string
	SetTemplates(v *QuerySmarttagTemplateListResponseBodyTemplates) *QuerySmarttagTemplateListResponseBody
	GetTemplates() *QuerySmarttagTemplateListResponseBodyTemplates
}

type QuerySmarttagTemplateListResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5210DBB0-E327-4D45-ADBC-0B83C8796E26
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates *QuerySmarttagTemplateListResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Struct"`
}

func (s QuerySmarttagTemplateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySmarttagTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySmarttagTemplateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySmarttagTemplateListResponseBody) GetTemplates() *QuerySmarttagTemplateListResponseBodyTemplates {
	return s.Templates
}

func (s *QuerySmarttagTemplateListResponseBody) SetRequestId(v string) *QuerySmarttagTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySmarttagTemplateListResponseBody) SetTemplates(v *QuerySmarttagTemplateListResponseBodyTemplates) *QuerySmarttagTemplateListResponseBody {
	s.Templates = v
	return s
}

func (s *QuerySmarttagTemplateListResponseBody) Validate() error {
	if s.Templates != nil {
		if err := s.Templates.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySmarttagTemplateListResponseBodyTemplates struct {
	Template []*QuerySmarttagTemplateListResponseBodyTemplatesTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s QuerySmarttagTemplateListResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s QuerySmarttagTemplateListResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *QuerySmarttagTemplateListResponseBodyTemplates) GetTemplate() []*QuerySmarttagTemplateListResponseBodyTemplatesTemplate {
	return s.Template
}

func (s *QuerySmarttagTemplateListResponseBodyTemplates) SetTemplate(v []*QuerySmarttagTemplateListResponseBodyTemplatesTemplate) *QuerySmarttagTemplateListResponseBodyTemplates {
	s.Template = v
	return s
}

func (s *QuerySmarttagTemplateListResponseBodyTemplates) Validate() error {
	if s.Template != nil {
		for _, item := range s.Template {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySmarttagTemplateListResponseBodyTemplatesTemplate struct {
	AnalyseTypes           *string `json:"AnalyseTypes,omitempty" xml:"AnalyseTypes,omitempty"`
	FaceCategoryIds        *string `json:"FaceCategoryIds,omitempty" xml:"FaceCategoryIds,omitempty"`
	FaceCustomParamsConfig *string `json:"FaceCustomParamsConfig,omitempty" xml:"FaceCustomParamsConfig,omitempty"`
	Industry               *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	IsDefault              *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	KeywordConfig          *string `json:"KeywordConfig,omitempty" xml:"KeywordConfig,omitempty"`
	KnowledgeConfig        *string `json:"KnowledgeConfig,omitempty" xml:"KnowledgeConfig,omitempty"`
	LabelType              *string `json:"LabelType,omitempty" xml:"LabelType,omitempty"`
	LabelVersion           *string `json:"LabelVersion,omitempty" xml:"LabelVersion,omitempty"`
	LandmarkGroupIds       *string `json:"LandmarkGroupIds,omitempty" xml:"LandmarkGroupIds,omitempty"`
	ObjectGroupIds         *string `json:"ObjectGroupIds,omitempty" xml:"ObjectGroupIds,omitempty"`
	Scene                  *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	TemplateId             *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName           *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s QuerySmarttagTemplateListResponseBodyTemplatesTemplate) String() string {
	return dara.Prettify(s)
}

func (s QuerySmarttagTemplateListResponseBodyTemplatesTemplate) GoString() string {
	return s.String()
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) GetAnalyseTypes() *string {
	return s.AnalyseTypes
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) GetFaceCategoryIds() *string {
	return s.FaceCategoryIds
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) GetFaceCustomParamsConfig() *string {
	return s.FaceCustomParamsConfig
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) GetIndustry() *string {
	return s.Industry
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) GetKeywordConfig() *string {
	return s.KeywordConfig
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) GetKnowledgeConfig() *string {
	return s.KnowledgeConfig
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) GetLabelType() *string {
	return s.LabelType
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) GetLabelVersion() *string {
	return s.LabelVersion
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) GetLandmarkGroupIds() *string {
	return s.LandmarkGroupIds
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) GetObjectGroupIds() *string {
	return s.ObjectGroupIds
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) GetScene() *string {
	return s.Scene
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) SetAnalyseTypes(v string) *QuerySmarttagTemplateListResponseBodyTemplatesTemplate {
	s.AnalyseTypes = &v
	return s
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) SetFaceCategoryIds(v string) *QuerySmarttagTemplateListResponseBodyTemplatesTemplate {
	s.FaceCategoryIds = &v
	return s
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) SetFaceCustomParamsConfig(v string) *QuerySmarttagTemplateListResponseBodyTemplatesTemplate {
	s.FaceCustomParamsConfig = &v
	return s
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) SetIndustry(v string) *QuerySmarttagTemplateListResponseBodyTemplatesTemplate {
	s.Industry = &v
	return s
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) SetIsDefault(v bool) *QuerySmarttagTemplateListResponseBodyTemplatesTemplate {
	s.IsDefault = &v
	return s
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) SetKeywordConfig(v string) *QuerySmarttagTemplateListResponseBodyTemplatesTemplate {
	s.KeywordConfig = &v
	return s
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) SetKnowledgeConfig(v string) *QuerySmarttagTemplateListResponseBodyTemplatesTemplate {
	s.KnowledgeConfig = &v
	return s
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) SetLabelType(v string) *QuerySmarttagTemplateListResponseBodyTemplatesTemplate {
	s.LabelType = &v
	return s
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) SetLabelVersion(v string) *QuerySmarttagTemplateListResponseBodyTemplatesTemplate {
	s.LabelVersion = &v
	return s
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) SetLandmarkGroupIds(v string) *QuerySmarttagTemplateListResponseBodyTemplatesTemplate {
	s.LandmarkGroupIds = &v
	return s
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) SetObjectGroupIds(v string) *QuerySmarttagTemplateListResponseBodyTemplatesTemplate {
	s.ObjectGroupIds = &v
	return s
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) SetScene(v string) *QuerySmarttagTemplateListResponseBodyTemplatesTemplate {
	s.Scene = &v
	return s
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) SetTemplateId(v string) *QuerySmarttagTemplateListResponseBodyTemplatesTemplate {
	s.TemplateId = &v
	return s
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) SetTemplateName(v string) *QuerySmarttagTemplateListResponseBodyTemplatesTemplate {
	s.TemplateName = &v
	return s
}

func (s *QuerySmarttagTemplateListResponseBodyTemplatesTemplate) Validate() error {
	return dara.Validate(s)
}
