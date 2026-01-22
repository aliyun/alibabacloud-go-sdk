// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecallManagementServiceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v *CreateRecallManagementServiceVersionRequestConfigs) *CreateRecallManagementServiceVersionRequest
	GetConfigs() *CreateRecallManagementServiceVersionRequestConfigs
	SetSourceRecallManagementServiceVersionId(v string) *CreateRecallManagementServiceVersionRequest
	GetSourceRecallManagementServiceVersionId() *string
}

type CreateRecallManagementServiceVersionRequest struct {
	Configs *CreateRecallManagementServiceVersionRequestConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Struct"`
	// example:
	//
	// 4
	SourceRecallManagementServiceVersionId *string `json:"SourceRecallManagementServiceVersionId,omitempty" xml:"SourceRecallManagementServiceVersionId,omitempty"`
}

func (s CreateRecallManagementServiceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionRequest) GetConfigs() *CreateRecallManagementServiceVersionRequestConfigs {
	return s.Configs
}

func (s *CreateRecallManagementServiceVersionRequest) GetSourceRecallManagementServiceVersionId() *string {
	return s.SourceRecallManagementServiceVersionId
}

func (s *CreateRecallManagementServiceVersionRequest) SetConfigs(v *CreateRecallManagementServiceVersionRequestConfigs) *CreateRecallManagementServiceVersionRequest {
	s.Configs = v
	return s
}

func (s *CreateRecallManagementServiceVersionRequest) SetSourceRecallManagementServiceVersionId(v string) *CreateRecallManagementServiceVersionRequest {
	s.SourceRecallManagementServiceVersionId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequest) Validate() error {
	if s.Configs != nil {
		if err := s.Configs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRecallManagementServiceVersionRequestConfigs struct {
	MergeConfig   *CreateRecallManagementServiceVersionRequestConfigsMergeConfig     `json:"MergeConfig,omitempty" xml:"MergeConfig,omitempty" type:"Struct"`
	RecallConfigs []*CreateRecallManagementServiceVersionRequestConfigsRecallConfigs `json:"RecallConfigs,omitempty" xml:"RecallConfigs,omitempty" type:"Repeated"`
}

func (s CreateRecallManagementServiceVersionRequestConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionRequestConfigs) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionRequestConfigs) GetMergeConfig() *CreateRecallManagementServiceVersionRequestConfigsMergeConfig {
	return s.MergeConfig
}

func (s *CreateRecallManagementServiceVersionRequestConfigs) GetRecallConfigs() []*CreateRecallManagementServiceVersionRequestConfigsRecallConfigs {
	return s.RecallConfigs
}

func (s *CreateRecallManagementServiceVersionRequestConfigs) SetMergeConfig(v *CreateRecallManagementServiceVersionRequestConfigsMergeConfig) *CreateRecallManagementServiceVersionRequestConfigs {
	s.MergeConfig = v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigs) SetRecallConfigs(v []*CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) *CreateRecallManagementServiceVersionRequestConfigs {
	s.RecallConfigs = v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigs) Validate() error {
	if s.MergeConfig != nil {
		if err := s.MergeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RecallConfigs != nil {
		for _, item := range s.RecallConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRecallManagementServiceVersionRequestConfigsMergeConfig struct {
	// example:
	//
	// {"":""}
	ExtendedConfig *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// example:
	//
	// age>10
	FilterExpression               *string   `json:"FilterExpression,omitempty" xml:"FilterExpression,omitempty"`
	FilterRecallManagementTableIds []*string `json:"FilterRecallManagementTableIds,omitempty" xml:"FilterRecallManagementTableIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	ItemRecallManagementTableId *string   `json:"ItemRecallManagementTableId,omitempty" xml:"ItemRecallManagementTableId,omitempty"`
	ItemTableFields             []*string `json:"ItemTableFields,omitempty" xml:"ItemTableFields,omitempty" type:"Repeated"`
	// example:
	//
	// Weight
	MergeType *string `json:"MergeType,omitempty" xml:"MergeType,omitempty"`
	// example:
	//
	// 1
	RecallManagementServiceVersionConfigId *string `json:"RecallManagementServiceVersionConfigId,omitempty" xml:"RecallManagementServiceVersionConfigId,omitempty"`
}

func (s CreateRecallManagementServiceVersionRequestConfigsMergeConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionRequestConfigsMergeConfig) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionRequestConfigsMergeConfig) GetExtendedConfig() *string {
	return s.ExtendedConfig
}

func (s *CreateRecallManagementServiceVersionRequestConfigsMergeConfig) GetFilterExpression() *string {
	return s.FilterExpression
}

func (s *CreateRecallManagementServiceVersionRequestConfigsMergeConfig) GetFilterRecallManagementTableIds() []*string {
	return s.FilterRecallManagementTableIds
}

func (s *CreateRecallManagementServiceVersionRequestConfigsMergeConfig) GetItemRecallManagementTableId() *string {
	return s.ItemRecallManagementTableId
}

func (s *CreateRecallManagementServiceVersionRequestConfigsMergeConfig) GetItemTableFields() []*string {
	return s.ItemTableFields
}

func (s *CreateRecallManagementServiceVersionRequestConfigsMergeConfig) GetMergeType() *string {
	return s.MergeType
}

func (s *CreateRecallManagementServiceVersionRequestConfigsMergeConfig) GetRecallManagementServiceVersionConfigId() *string {
	return s.RecallManagementServiceVersionConfigId
}

func (s *CreateRecallManagementServiceVersionRequestConfigsMergeConfig) SetExtendedConfig(v string) *CreateRecallManagementServiceVersionRequestConfigsMergeConfig {
	s.ExtendedConfig = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsMergeConfig) SetFilterExpression(v string) *CreateRecallManagementServiceVersionRequestConfigsMergeConfig {
	s.FilterExpression = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsMergeConfig) SetFilterRecallManagementTableIds(v []*string) *CreateRecallManagementServiceVersionRequestConfigsMergeConfig {
	s.FilterRecallManagementTableIds = v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsMergeConfig) SetItemRecallManagementTableId(v string) *CreateRecallManagementServiceVersionRequestConfigsMergeConfig {
	s.ItemRecallManagementTableId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsMergeConfig) SetItemTableFields(v []*string) *CreateRecallManagementServiceVersionRequestConfigsMergeConfig {
	s.ItemTableFields = v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsMergeConfig) SetMergeType(v string) *CreateRecallManagementServiceVersionRequestConfigsMergeConfig {
	s.MergeType = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsMergeConfig) SetRecallManagementServiceVersionConfigId(v string) *CreateRecallManagementServiceVersionRequestConfigsMergeConfig {
	s.RecallManagementServiceVersionConfigId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsMergeConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRecallManagementServiceVersionRequestConfigsRecallConfigs struct {
	// example:
	//
	// this is etrec recall
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ""
	ExtendedConfig *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// example:
	//
	// [{"option":"<","field":"category","type":"STRING","value":"10"}]
	ItemConditionArray *string `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	// example:
	//
	// age>10
	ItemConditionExpression *string `json:"ItemConditionExpression,omitempty" xml:"ItemConditionExpression,omitempty"`
	// example:
	//
	// item_embedding
	ItemVectorField *string `json:"ItemVectorField,omitempty" xml:"ItemVectorField,omitempty"`
	// example:
	//
	// 4
	ItemVectorRecallManagementTableId *string `json:"ItemVectorRecallManagementTableId,omitempty" xml:"ItemVectorRecallManagementTableId,omitempty"`
	// example:
	//
	// etrec
	Name      *string                                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	Operators []*CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1
	RecallManagementTableId *string `json:"RecallManagementTableId,omitempty" xml:"RecallManagementTableId,omitempty"`
	// example:
	//
	// X2I
	RecallType *string `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// example:
	//
	// user_embedding
	UserVectorField *string `json:"UserVectorField,omitempty" xml:"UserVectorField,omitempty"`
	// example:
	//
	// 3
	UserVectorRecallManagementTableId *string `json:"UserVectorRecallManagementTableId,omitempty" xml:"UserVectorRecallManagementTableId,omitempty"`
}

func (s CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) GetDescription() *string {
	return s.Description
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) GetExtendedConfig() *string {
	return s.ExtendedConfig
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) GetItemConditionArray() *string {
	return s.ItemConditionArray
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) GetItemConditionExpression() *string {
	return s.ItemConditionExpression
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) GetItemVectorField() *string {
	return s.ItemVectorField
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) GetItemVectorRecallManagementTableId() *string {
	return s.ItemVectorRecallManagementTableId
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) GetName() *string {
	return s.Name
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) GetOperators() []*CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators {
	return s.Operators
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) GetPriority() *int64 {
	return s.Priority
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) GetRecallManagementTableId() *string {
	return s.RecallManagementTableId
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) GetRecallType() *string {
	return s.RecallType
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) GetUserVectorField() *string {
	return s.UserVectorField
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) GetUserVectorRecallManagementTableId() *string {
	return s.UserVectorRecallManagementTableId
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) SetDescription(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs {
	s.Description = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) SetExtendedConfig(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs {
	s.ExtendedConfig = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) SetItemConditionArray(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs {
	s.ItemConditionArray = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) SetItemConditionExpression(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs {
	s.ItemConditionExpression = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) SetItemVectorField(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs {
	s.ItemVectorField = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) SetItemVectorRecallManagementTableId(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs {
	s.ItemVectorRecallManagementTableId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) SetName(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs {
	s.Name = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) SetOperators(v []*CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs {
	s.Operators = v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) SetPriority(v int64) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs {
	s.Priority = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) SetRecallManagementTableId(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs {
	s.RecallManagementTableId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) SetRecallType(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs {
	s.RecallType = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) SetUserVectorField(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs {
	s.UserVectorField = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) SetUserVectorRecallManagementTableId(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs {
	s.UserVectorRecallManagementTableId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigs) Validate() error {
	if s.Operators != nil {
		for _, item := range s.Operators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators struct {
	FeatureConfig *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFeatureConfig `json:"FeatureConfig,omitempty" xml:"FeatureConfig,omitempty" type:"Struct"`
	FilterConfig  *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFilterConfig  `json:"FilterConfig,omitempty" xml:"FilterConfig,omitempty" type:"Struct"`
	JoinConfig    *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsJoinConfig    `json:"JoinConfig,omitempty" xml:"JoinConfig,omitempty" type:"Struct"`
	// example:
	//
	// Filter
	OperatorType  *string                                                                                `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	TriggerConfig *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty" type:"Struct"`
}

func (s CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators) GetFeatureConfig() *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFeatureConfig {
	return s.FeatureConfig
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators) GetFilterConfig() *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFilterConfig {
	return s.FilterConfig
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators) GetJoinConfig() *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsJoinConfig {
	return s.JoinConfig
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators) GetOperatorType() *string {
	return s.OperatorType
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators) GetTriggerConfig() *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig {
	return s.TriggerConfig
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators) SetFeatureConfig(v *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFeatureConfig) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators {
	s.FeatureConfig = v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators) SetFilterConfig(v *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFilterConfig) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators {
	s.FilterConfig = v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators) SetJoinConfig(v *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsJoinConfig) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators {
	s.JoinConfig = v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators) SetOperatorType(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators {
	s.OperatorType = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators) SetTriggerConfig(v *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators {
	s.TriggerConfig = v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperators) Validate() error {
	if s.FeatureConfig != nil {
		if err := s.FeatureConfig.Validate(); err != nil {
			return err
		}
	}
	if s.FilterConfig != nil {
		if err := s.FilterConfig.Validate(); err != nil {
			return err
		}
	}
	if s.JoinConfig != nil {
		if err := s.JoinConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TriggerConfig != nil {
		if err := s.TriggerConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFeatureConfig struct {
	// example:
	//
	// category=3
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// city
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFeatureConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFeatureConfig) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFeatureConfig) GetExpression() *string {
	return s.Expression
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFeatureConfig) GetName() *string {
	return s.Name
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFeatureConfig) GetType() *string {
	return s.Type
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFeatureConfig) SetExpression(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFeatureConfig {
	s.Expression = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFeatureConfig) SetName(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFeatureConfig {
	s.Name = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFeatureConfig) SetType(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFeatureConfig {
	s.Type = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFeatureConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFilterConfig struct {
	// example:
	//
	// age>20
	Experession *string `json:"Experession,omitempty" xml:"Experession,omitempty"`
}

func (s CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFilterConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFilterConfig) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFilterConfig) GetExperession() *string {
	return s.Experession
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFilterConfig) SetExperession(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFilterConfig {
	s.Experession = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsFilterConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsJoinConfig struct {
	// example:
	//
	// item_id
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// item_id
	OutputFields []*string `json:"OutputFields,omitempty" xml:"OutputFields,omitempty" type:"Repeated"`
	// example:
	//
	// 4
	RecallManagementTableId *string `json:"RecallManagementTableId,omitempty" xml:"RecallManagementTableId,omitempty"`
}

func (s CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsJoinConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsJoinConfig) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsJoinConfig) GetField() *string {
	return s.Field
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsJoinConfig) GetOutputFields() []*string {
	return s.OutputFields
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsJoinConfig) GetRecallManagementTableId() *string {
	return s.RecallManagementTableId
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsJoinConfig) SetField(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsJoinConfig {
	s.Field = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsJoinConfig) SetOutputFields(v []*string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsJoinConfig {
	s.OutputFields = v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsJoinConfig) SetRecallManagementTableId(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsJoinConfig {
	s.RecallManagementTableId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsJoinConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig struct {
	// example:
	//
	// user_id
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// 20
	FieldQuantityLimit *string `json:"FieldQuantityLimit,omitempty" xml:"FieldQuantityLimit,omitempty"`
	// example:
	//
	// false
	IsRandSort *string `json:"IsRandSort,omitempty" xml:"IsRandSort,omitempty"`
	// example:
	//
	// create_time
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
}

func (s CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig) GetField() *string {
	return s.Field
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig) GetFieldQuantityLimit() *string {
	return s.FieldQuantityLimit
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig) GetIsRandSort() *string {
	return s.IsRandSort
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig) GetSortField() *string {
	return s.SortField
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig) SetField(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig {
	s.Field = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig) SetFieldQuantityLimit(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig {
	s.FieldQuantityLimit = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig) SetIsRandSort(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig {
	s.IsRandSort = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig) SetSortField(v string) *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig {
	s.SortField = &v
	return s
}

func (s *CreateRecallManagementServiceVersionRequestConfigsRecallConfigsOperatorsTriggerConfig) Validate() error {
	return dara.Validate(s)
}
