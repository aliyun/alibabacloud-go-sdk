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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The templates.
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
	// The analysis types that are used in the template. One or more values are returned. Valid values:
	//
	// 	- **ocr**: text recognition
	//
	// 	- **asr**: speech recognition
	//
	// 	- **classification**: video classification
	//
	// 	- **shows**: program recognition
	//
	// 	- **face**: facial recognition
	//
	// 	- **role**: figure recognition
	//
	// 	- **object**: object recognition
	//
	// 	- **tvstation**: logo recognition
	//
	// 	- **action**: action recognition
	//
	// 	- **emotion**: facial expression recognition
	//
	// 	- **landmark**: landmark recognition
	//
	// 	- **scene**: scene recognition
	//
	// 	- **movieip**: movie intellectual property recognition
	//
	// 	- **subtitle**: subtitle extraction
	//
	// example:
	//
	// ocr,asr,classification,shows,face,role,object,tvstation,action,emotion,landmark,scene
	AnalyseTypes *string `json:"AnalyseTypes,omitempty" xml:"AnalyseTypes,omitempty"`
	// The IDs of the system facial image libraries that are used in the template. One or more values are returned. Valid values:
	//
	// 	- celebrity: the facial image library of celebrities
	//
	// 	- politician: the facial image library of politicians
	//
	// 	- sensitive: the facial image library of sensitive figures
	//
	// example:
	//
	// politician,sensitive,celebrity
	FaceCategoryIds *string `json:"FaceCategoryIds,omitempty" xml:"FaceCategoryIds,omitempty"`
	// The configurations of face-related algorithms. The value of this parameter is a JSON string and consists of the thresholds set for face detection and facial recognition. Valid values:
	//
	// 	- **faceDetThreshold**: The default threshold for face detection is 0.999. The threshold takes effect only for the faces that are strange to the system.
	//
	// 	- **faceRegThreshold**: The default threshold for facial recognition is 0.9.
	//
	// example:
	//
	// { "faceDetThreshold":0.999, "faceRegThreshold":0.9 }
	FaceCustomParamsConfig *string `json:"FaceCustomParamsConfig,omitempty" xml:"FaceCustomParamsConfig,omitempty"`
	// The industry to which the template applies. Default value: **common**. Valid values:
	//
	// 	- **microVideo**: short video industry
	//
	// 	- **common**: general industries
	//
	// example:
	//
	// common
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// Indicates whether the template is the default template. Valid values:
	//
	// 	- **true**: The template is the default template.
	//
	// 	- **false**: The template is not the default template.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The configuration of keyword tags. The type field specifies the category of a keyword tag. You can specify one or more values and separate the values with commas (,). Valid values:
	//
	// 	- name
	//
	// 	- location
	//
	// 	- organization
	//
	// 	- other
	//
	// > Keyword tags of all the categories are returned in one of the following scenarios: The KeywordConfig parameter is not specified or the Keyword field is invalid because it is not a JSON string, or the KeywordConfig parameter does not contain the type field or the type field is invalid.
	//
	// example:
	//
	// { "type": "name,location,organization,other" }
	KeywordConfig *string `json:"KeywordConfig,omitempty" xml:"KeywordConfig,omitempty"`
	// The fields to be identified as knowledge graph information when tags are returned in Smart tagging V2.0 and Smart tagging V2.0-custom modes. For more information, see [Knowledge graph fields in smart tagging jobs](https://help.aliyun.com/document_detail/356383.html). If this parameter is not specified or the specified value is NULL or invalid because it is not a JSON string, the following fields are returned:
	//
	// 	- movie-related fields:
	//
	//     	- name: the name of the intellectual property that is featured in the movie
	//
	//     	- alias: the alias of the intellectual property that is featured in the movie
	//
	//     	- chnl: the category of the movie
	//
	//     	- genre: the genre of the movie
	//
	//     	- country: the country or region in which the movie was produced
	//
	//     	- language: the language of the movie
	//
	//     	- releaseYear: the year when the movie was released
	//
	// 	- music-related fields:
	//
	//     	- songName: the name of the song
	//
	//     	- artistName: the name of the singer
	//
	//     	- artistArea: the region to which the singer belongs, such as China, Japan, Korea, Europe, and America, or others.
	//
	//     	- albumName: the name of the album
	//
	// 	- person-related fields:
	//
	//     	- name: the name of the person
	//
	//     	- gender: the gender of the person
	//
	//     	- citizenship: the nationality of the person
	//
	//     	- occupation: the occupation of the person
	//
	//     	- classification: the type into which the person is classified
	//
	//     	- nationality: the ethnic group of the person
	//
	//     	- birthPlace: the place where the person was born
	//
	//     	- birthDate: the date when the person was born
	//
	// 	- landmark-related fields:
	//
	//     	- name: the display name of the landmark
	//
	//     	- nameEn: the English name of the landmark
	//
	//     	- Description: the description of the parameter
	//
	//     	- address: the address of the landmark
	//
	// 	- item-related fields:
	//
	//     	- brandName: the brand of the item
	//
	//     	- finegrainName: the fine-grained description of the item
	//
	// example:
	//
	// { "movie":"name,alias,chnl,genre", "music":"songName,artistName", "person":"name,gender" }
	KnowledgeConfig *string `json:"KnowledgeConfig,omitempty" xml:"KnowledgeConfig,omitempty"`
	// The type of the tagging. Default value: **auto**. Valid values:
	//
	// 	- **auto**: machine tagging
	//
	// 	- **hmi**: tagging by human and machine
	//
	// example:
	//
	// hmi
	LabelType *string `json:"LabelType,omitempty" xml:"LabelType,omitempty"`
	// The version of the smart tagging feature. Default value: 1.0. Valid values:
	//
	// 	- 1.0: Smart tagging V1.0
	//
	// 	- 2.0: Smart tagging V2.0 (CPV tagging)
	//
	// 	- 2.0-custom: Smart tagging V2.0-custom (CPV tagging by using custom models)
	//
	// example:
	//
	// 1.0
	LabelVersion *string `json:"LabelVersion,omitempty" xml:"LabelVersion,omitempty"`
	// The IDs of the landmark libraries that are used in the template.
	//
	// example:
	//
	// common
	LandmarkGroupIds *string `json:"LandmarkGroupIds,omitempty" xml:"LandmarkGroupIds,omitempty"`
	// The IDs of the object libraries that are used in the template.
	//
	// example:
	//
	// general,item,weapon,animal
	ObjectGroupIds *string `json:"ObjectGroupIds,omitempty" xml:"ObjectGroupIds,omitempty"`
	// The scenario in which the template is used. Valid values:
	//
	// 	- **search**: search scenarios
	//
	// 	- **recommend**: recommendation scenarios
	//
	// example:
	//
	// search
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// 05de22f255284c7a8d2aab535dde****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// example-****
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
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
