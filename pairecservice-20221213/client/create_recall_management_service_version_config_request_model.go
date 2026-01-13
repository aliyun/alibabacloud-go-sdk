// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecallManagementServiceVersionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigType(v string) *CreateRecallManagementServiceVersionConfigRequest
	GetConfigType() *string
	SetInstanceId(v string) *CreateRecallManagementServiceVersionConfigRequest
	GetInstanceId() *string
	SetMergeConfig(v *CreateRecallManagementServiceVersionConfigRequestMergeConfig) *CreateRecallManagementServiceVersionConfigRequest
	GetMergeConfig() *CreateRecallManagementServiceVersionConfigRequestMergeConfig
	SetRecallConfig(v *CreateRecallManagementServiceVersionConfigRequestRecallConfig) *CreateRecallManagementServiceVersionConfigRequest
	GetRecallConfig() *CreateRecallManagementServiceVersionConfigRequestRecallConfig
}

type CreateRecallManagementServiceVersionConfigRequest struct {
	// example:
	//
	// Recall
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// example:
	//
	// pai-teest-1
	InstanceId   *string                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MergeConfig  *CreateRecallManagementServiceVersionConfigRequestMergeConfig  `json:"MergeConfig,omitempty" xml:"MergeConfig,omitempty" type:"Struct"`
	RecallConfig *CreateRecallManagementServiceVersionConfigRequestRecallConfig `json:"RecallConfig,omitempty" xml:"RecallConfig,omitempty" type:"Struct"`
}

func (s CreateRecallManagementServiceVersionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionConfigRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *CreateRecallManagementServiceVersionConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRecallManagementServiceVersionConfigRequest) GetMergeConfig() *CreateRecallManagementServiceVersionConfigRequestMergeConfig {
	return s.MergeConfig
}

func (s *CreateRecallManagementServiceVersionConfigRequest) GetRecallConfig() *CreateRecallManagementServiceVersionConfigRequestRecallConfig {
	return s.RecallConfig
}

func (s *CreateRecallManagementServiceVersionConfigRequest) SetConfigType(v string) *CreateRecallManagementServiceVersionConfigRequest {
	s.ConfigType = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequest) SetInstanceId(v string) *CreateRecallManagementServiceVersionConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequest) SetMergeConfig(v *CreateRecallManagementServiceVersionConfigRequestMergeConfig) *CreateRecallManagementServiceVersionConfigRequest {
	s.MergeConfig = v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequest) SetRecallConfig(v *CreateRecallManagementServiceVersionConfigRequestRecallConfig) *CreateRecallManagementServiceVersionConfigRequest {
	s.RecallConfig = v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequest) Validate() error {
	if s.MergeConfig != nil {
		if err := s.MergeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RecallConfig != nil {
		if err := s.RecallConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRecallManagementServiceVersionConfigRequestMergeConfig struct {
	// example:
	//
	// ""
	ExtendedConfig *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// example:
	//
	// age>20
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

func (s CreateRecallManagementServiceVersionConfigRequestMergeConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionConfigRequestMergeConfig) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionConfigRequestMergeConfig) GetExtendedConfig() *string {
	return s.ExtendedConfig
}

func (s *CreateRecallManagementServiceVersionConfigRequestMergeConfig) GetFilterExpression() *string {
	return s.FilterExpression
}

func (s *CreateRecallManagementServiceVersionConfigRequestMergeConfig) GetFilterRecallManagementTableIds() []*string {
	return s.FilterRecallManagementTableIds
}

func (s *CreateRecallManagementServiceVersionConfigRequestMergeConfig) GetItemRecallManagementTableId() *string {
	return s.ItemRecallManagementTableId
}

func (s *CreateRecallManagementServiceVersionConfigRequestMergeConfig) GetItemTableFields() []*string {
	return s.ItemTableFields
}

func (s *CreateRecallManagementServiceVersionConfigRequestMergeConfig) GetMergeType() *string {
	return s.MergeType
}

func (s *CreateRecallManagementServiceVersionConfigRequestMergeConfig) GetRecallManagementServiceVersionConfigId() *string {
	return s.RecallManagementServiceVersionConfigId
}

func (s *CreateRecallManagementServiceVersionConfigRequestMergeConfig) SetExtendedConfig(v string) *CreateRecallManagementServiceVersionConfigRequestMergeConfig {
	s.ExtendedConfig = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestMergeConfig) SetFilterExpression(v string) *CreateRecallManagementServiceVersionConfigRequestMergeConfig {
	s.FilterExpression = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestMergeConfig) SetFilterRecallManagementTableIds(v []*string) *CreateRecallManagementServiceVersionConfigRequestMergeConfig {
	s.FilterRecallManagementTableIds = v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestMergeConfig) SetItemRecallManagementTableId(v string) *CreateRecallManagementServiceVersionConfigRequestMergeConfig {
	s.ItemRecallManagementTableId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestMergeConfig) SetItemTableFields(v []*string) *CreateRecallManagementServiceVersionConfigRequestMergeConfig {
	s.ItemTableFields = v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestMergeConfig) SetMergeType(v string) *CreateRecallManagementServiceVersionConfigRequestMergeConfig {
	s.MergeType = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestMergeConfig) SetRecallManagementServiceVersionConfigId(v string) *CreateRecallManagementServiceVersionConfigRequestMergeConfig {
	s.RecallManagementServiceVersionConfigId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestMergeConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRecallManagementServiceVersionConfigRequestRecallConfig struct {
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
	// item_embedding
	ItemVectorField *string `json:"ItemVectorField,omitempty" xml:"ItemVectorField,omitempty"`
	// example:
	//
	// 5
	ItemVectorRecallManagementTableId *string `json:"ItemVectorRecallManagementTableId,omitempty" xml:"ItemVectorRecallManagementTableId,omitempty"`
	// example:
	//
	// etrec
	Name      *string                                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Operators []*CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 3
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
	// 4
	UserVectorRecallManagementTableId *string `json:"UserVectorRecallManagementTableId,omitempty" xml:"UserVectorRecallManagementTableId,omitempty"`
}

func (s CreateRecallManagementServiceVersionConfigRequestRecallConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionConfigRequestRecallConfig) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) GetDescription() *string {
	return s.Description
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) GetExtendedConfig() *string {
	return s.ExtendedConfig
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) GetItemConditionArray() *string {
	return s.ItemConditionArray
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) GetItemVectorField() *string {
	return s.ItemVectorField
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) GetItemVectorRecallManagementTableId() *string {
	return s.ItemVectorRecallManagementTableId
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) GetName() *string {
	return s.Name
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) GetOperators() []*CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators {
	return s.Operators
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) GetPriority() *int64 {
	return s.Priority
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) GetRecallManagementTableId() *string {
	return s.RecallManagementTableId
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) GetRecallType() *string {
	return s.RecallType
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) GetUserVectorField() *string {
	return s.UserVectorField
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) GetUserVectorRecallManagementTableId() *string {
	return s.UserVectorRecallManagementTableId
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) SetDescription(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.Description = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) SetExtendedConfig(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.ExtendedConfig = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) SetItemConditionArray(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.ItemConditionArray = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) SetItemVectorField(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.ItemVectorField = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) SetItemVectorRecallManagementTableId(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.ItemVectorRecallManagementTableId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) SetName(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.Name = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) SetOperators(v []*CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators) *CreateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.Operators = v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) SetPriority(v int64) *CreateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.Priority = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) SetRecallManagementTableId(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.RecallManagementTableId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) SetRecallType(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.RecallType = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) SetUserVectorField(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.UserVectorField = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) SetUserVectorRecallManagementTableId(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.UserVectorRecallManagementTableId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfig) Validate() error {
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

type CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators struct {
	FeatureConfig *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig `json:"FeatureConfig,omitempty" xml:"FeatureConfig,omitempty" type:"Struct"`
	FilterConfig  *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig  `json:"FilterConfig,omitempty" xml:"FilterConfig,omitempty" type:"Struct"`
	JoinConfig    *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig    `json:"JoinConfig,omitempty" xml:"JoinConfig,omitempty" type:"Struct"`
	// example:
	//
	// Filter
	OperatorType  *string                                                                              `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	TriggerConfig *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty" type:"Struct"`
}

func (s CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators) GetFeatureConfig() *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig {
	return s.FeatureConfig
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators) GetFilterConfig() *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig {
	return s.FilterConfig
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators) GetJoinConfig() *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig {
	return s.JoinConfig
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators) GetOperatorType() *string {
	return s.OperatorType
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators) GetTriggerConfig() *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig {
	return s.TriggerConfig
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators) SetFeatureConfig(v *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators {
	s.FeatureConfig = v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators) SetFilterConfig(v *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig) *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators {
	s.FilterConfig = v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators) SetJoinConfig(v *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators {
	s.JoinConfig = v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators) SetOperatorType(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators {
	s.OperatorType = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators) SetTriggerConfig(v *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators {
	s.TriggerConfig = v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperators) Validate() error {
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

type CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig struct {
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

func (s CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) GetExpression() *string {
	return s.Expression
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) GetName() *string {
	return s.Name
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) GetType() *string {
	return s.Type
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) SetExpression(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig {
	s.Expression = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) SetName(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig {
	s.Name = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) SetType(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig {
	s.Type = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig struct {
	// example:
	//
	// age>20
	Experession *string `json:"Experession,omitempty" xml:"Experession,omitempty"`
}

func (s CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig) GetExperession() *string {
	return s.Experession
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig) SetExperession(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig {
	s.Experession = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig struct {
	// example:
	//
	// item_id
	Field        *string   `json:"Field,omitempty" xml:"Field,omitempty"`
	OutputFields []*string `json:"OutputFields,omitempty" xml:"OutputFields,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	RecallManagementTableId *string `json:"RecallManagementTableId,omitempty" xml:"RecallManagementTableId,omitempty"`
}

func (s CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) GetField() *string {
	return s.Field
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) GetOutputFields() []*string {
	return s.OutputFields
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) GetRecallManagementTableId() *string {
	return s.RecallManagementTableId
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) SetField(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig {
	s.Field = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) SetOutputFields(v []*string) *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig {
	s.OutputFields = v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) SetRecallManagementTableId(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig {
	s.RecallManagementTableId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) Validate() error {
	return dara.Validate(s)
}

type CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig struct {
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

func (s CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) GetField() *string {
	return s.Field
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) GetFieldQuantityLimit() *string {
	return s.FieldQuantityLimit
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) GetIsRandSort() *string {
	return s.IsRandSort
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) GetSortField() *string {
	return s.SortField
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) SetField(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig {
	s.Field = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) SetFieldQuantityLimit(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig {
	s.FieldQuantityLimit = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) SetIsRandSort(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig {
	s.IsRandSort = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) SetSortField(v string) *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig {
	s.SortField = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) Validate() error {
	return dara.Validate(s)
}
