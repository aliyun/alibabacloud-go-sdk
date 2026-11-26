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
	// The description of the agent.
	//
	// example:
	//
	// Agent for customer service session quality inspection
	AgentDescription *string `json:"AgentDescription,omitempty" xml:"AgentDescription,omitempty"`
	// The name of the agent.
	//
	// example:
	//
	// Customer service quality inspection assistant
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The unique identifier of the agent configuration.
	//
	// example:
	//
	// 1001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The object type for analysis. Valid values: audio and text.
	//
	// example:
	//
	// text
	InputType *string `json:"InputType,omitempty" xml:"InputType,omitempty"`
	// The instruction type. Valid values: serviceInspection (service quality inspection), field (key property extraction), tagCategory (label classification), multiLevelTag (multi-level label recognition), and customPrompt (custom instruction).
	//
	// example:
	//
	// field
	InstructionType *string `json:"InstructionType,omitempty" xml:"InstructionType,omitempty"`
	// The parameters corresponding to the instruction type. Different InstructionType values require different sub-parameters. Specify only the sub-parameters that correspond to the current InstructionType and leave the others empty.
	InstructionTypeParam *AgentInfoInstructionTypeParam `json:"InstructionTypeParam,omitempty" xml:"InstructionTypeParam,omitempty" type:"Struct"`
	// The type of the large language model.
	//
	// example:
	//
	// qwen-max
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
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
	// The custom instruction parameters. Used when InstructionType is set to customPrompt.
	CustomPromptParam *AgentInfoInstructionTypeParamCustomPromptParam `json:"CustomPromptParam,omitempty" xml:"CustomPromptParam,omitempty" type:"Struct"`
	// The key attribute extraction parameters. Used when InstructionType is set to field.
	FieldsParam *AgentInfoInstructionTypeParamFieldsParam `json:"FieldsParam,omitempty" xml:"FieldsParam,omitempty" type:"Struct"`
	// The service quality inspection parameters. Used when InstructionType is set to serviceInspection.
	ServiceInspectionParam *AgentInfoInstructionTypeParamServiceInspectionParam `json:"ServiceInspectionParam,omitempty" xml:"ServiceInspectionParam,omitempty" type:"Struct"`
	// The label classification parameters. Used when InstructionType is set to tagCategory.
	TagCategoryParam *AgentInfoInstructionTypeParamTagCategoryParam `json:"TagCategoryParam,omitempty" xml:"TagCategoryParam,omitempty" type:"Struct"`
	// The multi-level tag recognition parameters. Used when InstructionType is set to multiLevelTag.
	TagTreeLevelParam *AgentInfoInstructionTypeParamTagTreeLevelParam `json:"TagTreeLevelParam,omitempty" xml:"TagTreeLevelParam,omitempty" type:"Struct"`
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

func (s *AgentInfoInstructionTypeParam) GetTagTreeLevelParam() *AgentInfoInstructionTypeParamTagTreeLevelParam {
	return s.TagTreeLevelParam
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

func (s *AgentInfoInstructionTypeParam) SetTagTreeLevelParam(v *AgentInfoInstructionTypeParamTagTreeLevelParam) *AgentInfoInstructionTypeParam {
	s.TagTreeLevelParam = v
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
	if s.TagTreeLevelParam != nil {
		if err := s.TagTreeLevelParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AgentInfoInstructionTypeParamCustomPromptParam struct {
	// The content of the user-defined prompt template. This parameter is required.
	//
	// example:
	//
	// Analyze the service attitude of the customer service representative in the following conversation
	CustomPrompt *string `json:"CustomPrompt,omitempty" xml:"CustomPrompt,omitempty"`
	// The list of name-value pairs attached to the custom instruction.
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
	// The name.
	//
	// example:
	//
	// Scenario
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value.
	//
	// example:
	//
	// After-sales inquiry
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
	// The list of attribute dimensions to be extracted.
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
	// The attribute description. This parameter is optional.
	//
	// example:
	//
	// The specific request raised by the customer during the session
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The attribute name. This parameter is required.
	//
	// example:
	//
	// Customer request
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The attribute value. In the key attribute extraction scenario, this is the extracted result value corresponding to the dimension.
	//
	// example:
	//
	// Refund
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
	// The list of quality inspection dimensions. This parameter is required.
	Dimensions []*AgentInfoInstructionTypeParamServiceInspectionParamDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The description of the quality inspection scenario. This parameter is optional.
	//
	// example:
	//
	// Service quality inspection for after-sales customer service sessions
	SceneDescription *string `json:"SceneDescription,omitempty" xml:"SceneDescription,omitempty"`
	// The name of the quality inspection scenario. This parameter is required.
	//
	// example:
	//
	// After-sales service quality inspection
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
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
	// The dimension description. This parameter is optional.
	//
	// example:
	//
	// Whether the customer service representative is polite and patient
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The dimension name. This parameter is required.
	//
	// example:
	//
	// Service attitude
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
	// The list of label classification dimensions. This parameter is required.
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
	// The dimension description.
	//
	// example:
	//
	// The classification of the types of issues consulted by the customer
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The dimension name.
	//
	// example:
	//
	// Issue type
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of candidate tag values.
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

type AgentInfoInstructionTypeParamTagTreeLevelParam struct {
	// The list of starting parent tag IDs at runtime. The first ID is treated as the matched parent tag, and recognition proceeds level by level from its child nodes, skipping the preceding levels. This field also takes effect when AgentId is specified.
	TagIds []*int64 `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The recognition configuration for each level, arranged in hierarchical order. This parameter is required.
	TagTreeLevels []*AgentInfoInstructionTypeParamTagTreeLevelParamTagTreeLevels `json:"TagTreeLevels,omitempty" xml:"TagTreeLevels,omitempty" type:"Repeated"`
}

func (s AgentInfoInstructionTypeParamTagTreeLevelParam) String() string {
	return dara.Prettify(s)
}

func (s AgentInfoInstructionTypeParamTagTreeLevelParam) GoString() string {
	return s.String()
}

func (s *AgentInfoInstructionTypeParamTagTreeLevelParam) GetTagIds() []*int64 {
	return s.TagIds
}

func (s *AgentInfoInstructionTypeParamTagTreeLevelParam) GetTagTreeLevels() []*AgentInfoInstructionTypeParamTagTreeLevelParamTagTreeLevels {
	return s.TagTreeLevels
}

func (s *AgentInfoInstructionTypeParamTagTreeLevelParam) SetTagIds(v []*int64) *AgentInfoInstructionTypeParamTagTreeLevelParam {
	s.TagIds = v
	return s
}

func (s *AgentInfoInstructionTypeParamTagTreeLevelParam) SetTagTreeLevels(v []*AgentInfoInstructionTypeParamTagTreeLevelParamTagTreeLevels) *AgentInfoInstructionTypeParamTagTreeLevelParam {
	s.TagTreeLevels = v
	return s
}

func (s *AgentInfoInstructionTypeParamTagTreeLevelParam) Validate() error {
	if s.TagTreeLevels != nil {
		for _, item := range s.TagTreeLevels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AgentInfoInstructionTypeParamTagTreeLevelParamTagTreeLevels struct {
	// The analysis prompt for this level. It is concatenated with the tag data and sent to the large language model. This parameter is required.
	//
	// example:
	//
	// Determine which tag category this conversation belongs to
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The list of candidate tag IDs for this level. A value of -1 indicates that all tags at this level are selected. This parameter is required.
	TagTreeIds []*int64 `json:"TagTreeIds,omitempty" xml:"TagTreeIds,omitempty" type:"Repeated"`
}

func (s AgentInfoInstructionTypeParamTagTreeLevelParamTagTreeLevels) String() string {
	return dara.Prettify(s)
}

func (s AgentInfoInstructionTypeParamTagTreeLevelParamTagTreeLevels) GoString() string {
	return s.String()
}

func (s *AgentInfoInstructionTypeParamTagTreeLevelParamTagTreeLevels) GetPrompt() *string {
	return s.Prompt
}

func (s *AgentInfoInstructionTypeParamTagTreeLevelParamTagTreeLevels) GetTagTreeIds() []*int64 {
	return s.TagTreeIds
}

func (s *AgentInfoInstructionTypeParamTagTreeLevelParamTagTreeLevels) SetPrompt(v string) *AgentInfoInstructionTypeParamTagTreeLevelParamTagTreeLevels {
	s.Prompt = &v
	return s
}

func (s *AgentInfoInstructionTypeParamTagTreeLevelParamTagTreeLevels) SetTagTreeIds(v []*int64) *AgentInfoInstructionTypeParamTagTreeLevelParamTagTreeLevels {
	s.TagTreeIds = v
	return s
}

func (s *AgentInfoInstructionTypeParamTagTreeLevelParamTagTreeLevels) Validate() error {
	return dara.Validate(s)
}
