// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAgentDescription(v string) *AgentInfo
	GetAgentDescription() *string
	SetAgentName(v string) *AgentInfo
	GetAgentName() *string
	SetId(v int64) *AgentInfo
	GetId() *int64
	SetInputType(v string) *AgentInfo
	GetInputType() *string
	SetInstructionType(v string) *AgentInfo
	GetInstructionType() *string
	SetInstructionTypeParam(v *AgentInfoInstructionTypeParam) *AgentInfo
	GetInstructionTypeParam() *AgentInfoInstructionTypeParam
	SetModelType(v string) *AgentInfo
	GetModelType() *string
}

type AgentInfo struct {
	AgentDescription     *string                        `json:"AgentDescription,omitempty" xml:"AgentDescription,omitempty"`
	AgentName            *string                        `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	Id                   *int64                         `json:"Id,omitempty" xml:"Id,omitempty"`
	InputType            *string                        `json:"InputType,omitempty" xml:"InputType,omitempty"`
	InstructionType      *string                        `json:"InstructionType,omitempty" xml:"InstructionType,omitempty"`
	InstructionTypeParam *AgentInfoInstructionTypeParam `json:"InstructionTypeParam,omitempty" xml:"InstructionTypeParam,omitempty" type:"Struct"`
	ModelType            *string                        `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
}

func (s AgentInfo) String() string {
	return dara.Prettify(s)
}

func (s AgentInfo) GoString() string {
	return s.String()
}

func (s *AgentInfo) GetAgentDescription() *string {
	return s.AgentDescription
}

func (s *AgentInfo) GetAgentName() *string {
	return s.AgentName
}

func (s *AgentInfo) GetId() *int64 {
	return s.Id
}

func (s *AgentInfo) GetInputType() *string {
	return s.InputType
}

func (s *AgentInfo) GetInstructionType() *string {
	return s.InstructionType
}

func (s *AgentInfo) GetInstructionTypeParam() *AgentInfoInstructionTypeParam {
	return s.InstructionTypeParam
}

func (s *AgentInfo) GetModelType() *string {
	return s.ModelType
}

func (s *AgentInfo) SetAgentDescription(v string) *AgentInfo {
	s.AgentDescription = &v
	return s
}

func (s *AgentInfo) SetAgentName(v string) *AgentInfo {
	s.AgentName = &v
	return s
}

func (s *AgentInfo) SetId(v int64) *AgentInfo {
	s.Id = &v
	return s
}

func (s *AgentInfo) SetInputType(v string) *AgentInfo {
	s.InputType = &v
	return s
}

func (s *AgentInfo) SetInstructionType(v string) *AgentInfo {
	s.InstructionType = &v
	return s
}

func (s *AgentInfo) SetInstructionTypeParam(v *AgentInfoInstructionTypeParam) *AgentInfo {
	s.InstructionTypeParam = v
	return s
}

func (s *AgentInfo) SetModelType(v string) *AgentInfo {
	s.ModelType = &v
	return s
}

func (s *AgentInfo) Validate() error {
	if s.InstructionTypeParam != nil {
		if err := s.InstructionTypeParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AgentInfoInstructionTypeParam struct {
	CustomPromptParam      *AgentInfoInstructionTypeParamCustomPromptParam      `json:"CustomPromptParam,omitempty" xml:"CustomPromptParam,omitempty" type:"Struct"`
	FieldsParam            *AgentInfoInstructionTypeParamFieldsParam            `json:"FieldsParam,omitempty" xml:"FieldsParam,omitempty" type:"Struct"`
	ServiceInspectionParam *AgentInfoInstructionTypeParamServiceInspectionParam `json:"ServiceInspectionParam,omitempty" xml:"ServiceInspectionParam,omitempty" type:"Struct"`
	TagCategoryParam       *AgentInfoInstructionTypeParamTagCategoryParam       `json:"TagCategoryParam,omitempty" xml:"TagCategoryParam,omitempty" type:"Struct"`
}

func (s AgentInfoInstructionTypeParam) String() string {
	return dara.Prettify(s)
}

func (s AgentInfoInstructionTypeParam) GoString() string {
	return s.String()
}

func (s *AgentInfoInstructionTypeParam) GetCustomPromptParam() *AgentInfoInstructionTypeParamCustomPromptParam {
	return s.CustomPromptParam
}

func (s *AgentInfoInstructionTypeParam) GetFieldsParam() *AgentInfoInstructionTypeParamFieldsParam {
	return s.FieldsParam
}

func (s *AgentInfoInstructionTypeParam) GetServiceInspectionParam() *AgentInfoInstructionTypeParamServiceInspectionParam {
	return s.ServiceInspectionParam
}

func (s *AgentInfoInstructionTypeParam) GetTagCategoryParam() *AgentInfoInstructionTypeParamTagCategoryParam {
	return s.TagCategoryParam
}

func (s *AgentInfoInstructionTypeParam) SetCustomPromptParam(v *AgentInfoInstructionTypeParamCustomPromptParam) *AgentInfoInstructionTypeParam {
	s.CustomPromptParam = v
	return s
}

func (s *AgentInfoInstructionTypeParam) SetFieldsParam(v *AgentInfoInstructionTypeParamFieldsParam) *AgentInfoInstructionTypeParam {
	s.FieldsParam = v
	return s
}

func (s *AgentInfoInstructionTypeParam) SetServiceInspectionParam(v *AgentInfoInstructionTypeParamServiceInspectionParam) *AgentInfoInstructionTypeParam {
	s.ServiceInspectionParam = v
	return s
}

func (s *AgentInfoInstructionTypeParam) SetTagCategoryParam(v *AgentInfoInstructionTypeParamTagCategoryParam) *AgentInfoInstructionTypeParam {
	s.TagCategoryParam = v
	return s
}

func (s *AgentInfoInstructionTypeParam) Validate() error {
	if s.CustomPromptParam != nil {
		if err := s.CustomPromptParam.Validate(); err != nil {
			return err
		}
	}
	if s.FieldsParam != nil {
		if err := s.FieldsParam.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceInspectionParam != nil {
		if err := s.ServiceInspectionParam.Validate(); err != nil {
			return err
		}
	}
	if s.TagCategoryParam != nil {
		if err := s.TagCategoryParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AgentInfoInstructionTypeParamCustomPromptParam struct {
	CustomPrompt     *string                                                           `json:"CustomPrompt,omitempty" xml:"CustomPrompt,omitempty"`
	NameDescPairList []*AgentInfoInstructionTypeParamCustomPromptParamNameDescPairList `json:"NameDescPairList,omitempty" xml:"NameDescPairList,omitempty" type:"Repeated"`
}

func (s AgentInfoInstructionTypeParamCustomPromptParam) String() string {
	return dara.Prettify(s)
}

func (s AgentInfoInstructionTypeParamCustomPromptParam) GoString() string {
	return s.String()
}

func (s *AgentInfoInstructionTypeParamCustomPromptParam) GetCustomPrompt() *string {
	return s.CustomPrompt
}

func (s *AgentInfoInstructionTypeParamCustomPromptParam) GetNameDescPairList() []*AgentInfoInstructionTypeParamCustomPromptParamNameDescPairList {
	return s.NameDescPairList
}

func (s *AgentInfoInstructionTypeParamCustomPromptParam) SetCustomPrompt(v string) *AgentInfoInstructionTypeParamCustomPromptParam {
	s.CustomPrompt = &v
	return s
}

func (s *AgentInfoInstructionTypeParamCustomPromptParam) SetNameDescPairList(v []*AgentInfoInstructionTypeParamCustomPromptParamNameDescPairList) *AgentInfoInstructionTypeParamCustomPromptParam {
	s.NameDescPairList = v
	return s
}

func (s *AgentInfoInstructionTypeParamCustomPromptParam) Validate() error {
	if s.NameDescPairList != nil {
		for _, item := range s.NameDescPairList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AgentInfoInstructionTypeParamCustomPromptParamNameDescPairList struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AgentInfoInstructionTypeParamCustomPromptParamNameDescPairList) String() string {
	return dara.Prettify(s)
}

func (s AgentInfoInstructionTypeParamCustomPromptParamNameDescPairList) GoString() string {
	return s.String()
}

func (s *AgentInfoInstructionTypeParamCustomPromptParamNameDescPairList) GetName() *string {
	return s.Name
}

func (s *AgentInfoInstructionTypeParamCustomPromptParamNameDescPairList) GetValue() *string {
	return s.Value
}

func (s *AgentInfoInstructionTypeParamCustomPromptParamNameDescPairList) SetName(v string) *AgentInfoInstructionTypeParamCustomPromptParamNameDescPairList {
	s.Name = &v
	return s
}

func (s *AgentInfoInstructionTypeParamCustomPromptParamNameDescPairList) SetValue(v string) *AgentInfoInstructionTypeParamCustomPromptParamNameDescPairList {
	s.Value = &v
	return s
}

func (s *AgentInfoInstructionTypeParamCustomPromptParamNameDescPairList) Validate() error {
	return dara.Validate(s)
}

type AgentInfoInstructionTypeParamFieldsParam struct {
	NameDescPairList []*AgentInfoInstructionTypeParamFieldsParamNameDescPairList `json:"NameDescPairList,omitempty" xml:"NameDescPairList,omitempty" type:"Repeated"`
}

func (s AgentInfoInstructionTypeParamFieldsParam) String() string {
	return dara.Prettify(s)
}

func (s AgentInfoInstructionTypeParamFieldsParam) GoString() string {
	return s.String()
}

func (s *AgentInfoInstructionTypeParamFieldsParam) GetNameDescPairList() []*AgentInfoInstructionTypeParamFieldsParamNameDescPairList {
	return s.NameDescPairList
}

func (s *AgentInfoInstructionTypeParamFieldsParam) SetNameDescPairList(v []*AgentInfoInstructionTypeParamFieldsParamNameDescPairList) *AgentInfoInstructionTypeParamFieldsParam {
	s.NameDescPairList = v
	return s
}

func (s *AgentInfoInstructionTypeParamFieldsParam) Validate() error {
	if s.NameDescPairList != nil {
		for _, item := range s.NameDescPairList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AgentInfoInstructionTypeParamFieldsParamNameDescPairList struct {
	Desc  *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AgentInfoInstructionTypeParamFieldsParamNameDescPairList) String() string {
	return dara.Prettify(s)
}

func (s AgentInfoInstructionTypeParamFieldsParamNameDescPairList) GoString() string {
	return s.String()
}

func (s *AgentInfoInstructionTypeParamFieldsParamNameDescPairList) GetDesc() *string {
	return s.Desc
}

func (s *AgentInfoInstructionTypeParamFieldsParamNameDescPairList) GetName() *string {
	return s.Name
}

func (s *AgentInfoInstructionTypeParamFieldsParamNameDescPairList) GetValue() *string {
	return s.Value
}

func (s *AgentInfoInstructionTypeParamFieldsParamNameDescPairList) SetDesc(v string) *AgentInfoInstructionTypeParamFieldsParamNameDescPairList {
	s.Desc = &v
	return s
}

func (s *AgentInfoInstructionTypeParamFieldsParamNameDescPairList) SetName(v string) *AgentInfoInstructionTypeParamFieldsParamNameDescPairList {
	s.Name = &v
	return s
}

func (s *AgentInfoInstructionTypeParamFieldsParamNameDescPairList) SetValue(v string) *AgentInfoInstructionTypeParamFieldsParamNameDescPairList {
	s.Value = &v
	return s
}

func (s *AgentInfoInstructionTypeParamFieldsParamNameDescPairList) Validate() error {
	return dara.Validate(s)
}

type AgentInfoInstructionTypeParamServiceInspectionParam struct {
	Dimensions       []*AgentInfoInstructionTypeParamServiceInspectionParamDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	SceneDescription *string                                                          `json:"SceneDescription,omitempty" xml:"SceneDescription,omitempty"`
	SceneName        *string                                                          `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s AgentInfoInstructionTypeParamServiceInspectionParam) String() string {
	return dara.Prettify(s)
}

func (s AgentInfoInstructionTypeParamServiceInspectionParam) GoString() string {
	return s.String()
}

func (s *AgentInfoInstructionTypeParamServiceInspectionParam) GetDimensions() []*AgentInfoInstructionTypeParamServiceInspectionParamDimensions {
	return s.Dimensions
}

func (s *AgentInfoInstructionTypeParamServiceInspectionParam) GetSceneDescription() *string {
	return s.SceneDescription
}

func (s *AgentInfoInstructionTypeParamServiceInspectionParam) GetSceneName() *string {
	return s.SceneName
}

func (s *AgentInfoInstructionTypeParamServiceInspectionParam) SetDimensions(v []*AgentInfoInstructionTypeParamServiceInspectionParamDimensions) *AgentInfoInstructionTypeParamServiceInspectionParam {
	s.Dimensions = v
	return s
}

func (s *AgentInfoInstructionTypeParamServiceInspectionParam) SetSceneDescription(v string) *AgentInfoInstructionTypeParamServiceInspectionParam {
	s.SceneDescription = &v
	return s
}

func (s *AgentInfoInstructionTypeParamServiceInspectionParam) SetSceneName(v string) *AgentInfoInstructionTypeParamServiceInspectionParam {
	s.SceneName = &v
	return s
}

func (s *AgentInfoInstructionTypeParamServiceInspectionParam) Validate() error {
	if s.Dimensions != nil {
		for _, item := range s.Dimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AgentInfoInstructionTypeParamServiceInspectionParamDimensions struct {
	Desc      *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
}

func (s AgentInfoInstructionTypeParamServiceInspectionParamDimensions) String() string {
	return dara.Prettify(s)
}

func (s AgentInfoInstructionTypeParamServiceInspectionParamDimensions) GoString() string {
	return s.String()
}

func (s *AgentInfoInstructionTypeParamServiceInspectionParamDimensions) GetDesc() *string {
	return s.Desc
}

func (s *AgentInfoInstructionTypeParamServiceInspectionParamDimensions) GetDimension() *string {
	return s.Dimension
}

func (s *AgentInfoInstructionTypeParamServiceInspectionParamDimensions) SetDesc(v string) *AgentInfoInstructionTypeParamServiceInspectionParamDimensions {
	s.Desc = &v
	return s
}

func (s *AgentInfoInstructionTypeParamServiceInspectionParamDimensions) SetDimension(v string) *AgentInfoInstructionTypeParamServiceInspectionParamDimensions {
	s.Dimension = &v
	return s
}

func (s *AgentInfoInstructionTypeParamServiceInspectionParamDimensions) Validate() error {
	return dara.Validate(s)
}

type AgentInfoInstructionTypeParamTagCategoryParam struct {
	NameDescPairList []*AgentInfoInstructionTypeParamTagCategoryParamNameDescPairList `json:"NameDescPairList,omitempty" xml:"NameDescPairList,omitempty" type:"Repeated"`
}

func (s AgentInfoInstructionTypeParamTagCategoryParam) String() string {
	return dara.Prettify(s)
}

func (s AgentInfoInstructionTypeParamTagCategoryParam) GoString() string {
	return s.String()
}

func (s *AgentInfoInstructionTypeParamTagCategoryParam) GetNameDescPairList() []*AgentInfoInstructionTypeParamTagCategoryParamNameDescPairList {
	return s.NameDescPairList
}

func (s *AgentInfoInstructionTypeParamTagCategoryParam) SetNameDescPairList(v []*AgentInfoInstructionTypeParamTagCategoryParamNameDescPairList) *AgentInfoInstructionTypeParamTagCategoryParam {
	s.NameDescPairList = v
	return s
}

func (s *AgentInfoInstructionTypeParamTagCategoryParam) Validate() error {
	if s.NameDescPairList != nil {
		for _, item := range s.NameDescPairList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AgentInfoInstructionTypeParamTagCategoryParamNameDescPairList struct {
	Desc      *string   `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Name      *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s AgentInfoInstructionTypeParamTagCategoryParamNameDescPairList) String() string {
	return dara.Prettify(s)
}

func (s AgentInfoInstructionTypeParamTagCategoryParamNameDescPairList) GoString() string {
	return s.String()
}

func (s *AgentInfoInstructionTypeParamTagCategoryParamNameDescPairList) GetDesc() *string {
	return s.Desc
}

func (s *AgentInfoInstructionTypeParamTagCategoryParamNameDescPairList) GetName() *string {
	return s.Name
}

func (s *AgentInfoInstructionTypeParamTagCategoryParamNameDescPairList) GetValueList() []*string {
	return s.ValueList
}

func (s *AgentInfoInstructionTypeParamTagCategoryParamNameDescPairList) SetDesc(v string) *AgentInfoInstructionTypeParamTagCategoryParamNameDescPairList {
	s.Desc = &v
	return s
}

func (s *AgentInfoInstructionTypeParamTagCategoryParamNameDescPairList) SetName(v string) *AgentInfoInstructionTypeParamTagCategoryParamNameDescPairList {
	s.Name = &v
	return s
}

func (s *AgentInfoInstructionTypeParamTagCategoryParamNameDescPairList) SetValueList(v []*string) *AgentInfoInstructionTypeParamTagCategoryParamNameDescPairList {
	s.ValueList = v
	return s
}

func (s *AgentInfoInstructionTypeParamTagCategoryParamNameDescPairList) Validate() error {
	return dara.Validate(s)
}
