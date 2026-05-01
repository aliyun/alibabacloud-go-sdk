// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuntimeModelConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetRuntimeModelConfigResponseBodyData) *GetRuntimeModelConfigResponseBody
	GetData() *GetRuntimeModelConfigResponseBodyData
	SetRequestId(v string) *GetRuntimeModelConfigResponseBody
	GetRequestId() *string
}

type GetRuntimeModelConfigResponseBody struct {
	Data *GetRuntimeModelConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRuntimeModelConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRuntimeModelConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuntimeModelConfigResponseBody) GetData() *GetRuntimeModelConfigResponseBodyData {
	return s.Data
}

func (s *GetRuntimeModelConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRuntimeModelConfigResponseBody) SetData(v *GetRuntimeModelConfigResponseBodyData) *GetRuntimeModelConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetRuntimeModelConfigResponseBody) SetRequestId(v string) *GetRuntimeModelConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuntimeModelConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRuntimeModelConfigResponseBodyData struct {
	// example:
	//
	// bailian/qwen3.6-plus
	DefaultModel      *string                                                   `json:"DefaultModel,omitempty" xml:"DefaultModel,omitempty"`
	ModelProviderList []*GetRuntimeModelConfigResponseBodyDataModelProviderList `json:"ModelProviderList,omitempty" xml:"ModelProviderList,omitempty" type:"Repeated"`
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// example:
	//
	// model-template-001
	ModelTemplateName *string `json:"ModelTemplateName,omitempty" xml:"ModelTemplateName,omitempty"`
}

func (s GetRuntimeModelConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRuntimeModelConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRuntimeModelConfigResponseBodyData) GetDefaultModel() *string {
	return s.DefaultModel
}

func (s *GetRuntimeModelConfigResponseBodyData) GetModelProviderList() []*GetRuntimeModelConfigResponseBodyDataModelProviderList {
	return s.ModelProviderList
}

func (s *GetRuntimeModelConfigResponseBodyData) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *GetRuntimeModelConfigResponseBodyData) GetModelTemplateName() *string {
	return s.ModelTemplateName
}

func (s *GetRuntimeModelConfigResponseBodyData) SetDefaultModel(v string) *GetRuntimeModelConfigResponseBodyData {
	s.DefaultModel = &v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyData) SetModelProviderList(v []*GetRuntimeModelConfigResponseBodyDataModelProviderList) *GetRuntimeModelConfigResponseBodyData {
	s.ModelProviderList = v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyData) SetModelTemplateId(v string) *GetRuntimeModelConfigResponseBodyData {
	s.ModelTemplateId = &v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyData) SetModelTemplateName(v string) *GetRuntimeModelConfigResponseBodyData {
	s.ModelTemplateName = &v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyData) Validate() error {
	if s.ModelProviderList != nil {
		for _, item := range s.ModelProviderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRuntimeModelConfigResponseBodyDataModelProviderList struct {
	LlmInfoList []*GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList `json:"LlmInfoList,omitempty" xml:"LlmInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// mpt-xxxx
	ModelProviderTemplateId *string `json:"ModelProviderTemplateId,omitempty" xml:"ModelProviderTemplateId,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// bailian
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s GetRuntimeModelConfigResponseBodyDataModelProviderList) String() string {
	return dara.Prettify(s)
}

func (s GetRuntimeModelConfigResponseBodyDataModelProviderList) GoString() string {
	return s.String()
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderList) GetLlmInfoList() []*GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList {
	return s.LlmInfoList
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderList) GetModelProviderTemplateId() *string {
	return s.ModelProviderTemplateId
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderList) GetName() *string {
	return s.Name
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderList) GetProviderName() *string {
	return s.ProviderName
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderList) SetLlmInfoList(v []*GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) *GetRuntimeModelConfigResponseBodyDataModelProviderList {
	s.LlmInfoList = v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderList) SetModelProviderTemplateId(v string) *GetRuntimeModelConfigResponseBodyDataModelProviderList {
	s.ModelProviderTemplateId = &v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderList) SetName(v string) *GetRuntimeModelConfigResponseBodyDataModelProviderList {
	s.Name = &v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderList) SetProviderName(v string) *GetRuntimeModelConfigResponseBodyDataModelProviderList {
	s.ProviderName = &v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderList) Validate() error {
	if s.LlmInfoList != nil {
		for _, item := range s.LlmInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList struct {
	Description       *string                                                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	Features          []*string                                                                           `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	InferenceMetadata *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoListInferenceMetadata `json:"InferenceMetadata,omitempty" xml:"InferenceMetadata,omitempty" type:"Struct"`
	// example:
	//
	// qwen3.6-plus
	LlmCode *string `json:"LlmCode,omitempty" xml:"LlmCode,omitempty"`
	// example:
	//
	// Qwen3.6-Plus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2026-03-04T06:25:17.000+00:00
	PublishedTime *string `json:"PublishedTime,omitempty" xml:"PublishedTime,omitempty"`
	// example:
	//
	// Normal
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
}

func (s GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) GoString() string {
	return s.String()
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) GetDescription() *string {
	return s.Description
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) GetFeatures() []*string {
	return s.Features
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) GetInferenceMetadata() *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoListInferenceMetadata {
	return s.InferenceMetadata
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) GetLlmCode() *string {
	return s.LlmCode
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) GetName() *string {
	return s.Name
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) GetPublishedTime() *string {
	return s.PublishedTime
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) GetRiskType() *string {
	return s.RiskType
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) SetDescription(v string) *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList {
	s.Description = &v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) SetFeatures(v []*string) *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList {
	s.Features = v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) SetInferenceMetadata(v *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoListInferenceMetadata) *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList {
	s.InferenceMetadata = v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) SetLlmCode(v string) *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList {
	s.LlmCode = &v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) SetName(v string) *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList {
	s.Name = &v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) SetPublishedTime(v string) *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList {
	s.PublishedTime = &v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) SetRiskType(v string) *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList {
	s.RiskType = &v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoList) Validate() error {
	if s.InferenceMetadata != nil {
		if err := s.InferenceMetadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoListInferenceMetadata struct {
	RequestModality  []*string `json:"RequestModality,omitempty" xml:"RequestModality,omitempty" type:"Repeated"`
	ResponseModality []*string `json:"ResponseModality,omitempty" xml:"ResponseModality,omitempty" type:"Repeated"`
}

func (s GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoListInferenceMetadata) String() string {
	return dara.Prettify(s)
}

func (s GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoListInferenceMetadata) GoString() string {
	return s.String()
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoListInferenceMetadata) GetRequestModality() []*string {
	return s.RequestModality
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoListInferenceMetadata) GetResponseModality() []*string {
	return s.ResponseModality
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoListInferenceMetadata) SetRequestModality(v []*string) *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoListInferenceMetadata {
	s.RequestModality = v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoListInferenceMetadata) SetResponseModality(v []*string) *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoListInferenceMetadata {
	s.ResponseModality = v
	return s
}

func (s *GetRuntimeModelConfigResponseBodyDataModelProviderListLlmInfoListInferenceMetadata) Validate() error {
	return dara.Validate(s)
}
