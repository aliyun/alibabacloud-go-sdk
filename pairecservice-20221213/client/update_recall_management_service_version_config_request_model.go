// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecallManagementServiceVersionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UpdateRecallManagementServiceVersionConfigRequest
	GetRegionId() *string
	SetConfigType(v string) *UpdateRecallManagementServiceVersionConfigRequest
	GetConfigType() *string
	SetInstanceId(v string) *UpdateRecallManagementServiceVersionConfigRequest
	GetInstanceId() *string
	SetMergeConfig(v *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) *UpdateRecallManagementServiceVersionConfigRequest
	GetMergeConfig() *UpdateRecallManagementServiceVersionConfigRequestMergeConfig
	SetRecallConfig(v *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) *UpdateRecallManagementServiceVersionConfigRequest
	GetRecallConfig() *UpdateRecallManagementServiceVersionConfigRequestRecallConfig
}

type UpdateRecallManagementServiceVersionConfigRequest struct {
	RegionId     *string                                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ConfigType   *string                                                        `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	InstanceId   *string                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MergeConfig  *UpdateRecallManagementServiceVersionConfigRequestMergeConfig  `json:"MergeConfig,omitempty" xml:"MergeConfig,omitempty" type:"Struct"`
	RecallConfig *UpdateRecallManagementServiceVersionConfigRequestRecallConfig `json:"RecallConfig,omitempty" xml:"RecallConfig,omitempty" type:"Struct"`
}

func (s UpdateRecallManagementServiceVersionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementServiceVersionConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementServiceVersionConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateRecallManagementServiceVersionConfigRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *UpdateRecallManagementServiceVersionConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateRecallManagementServiceVersionConfigRequest) GetMergeConfig() *UpdateRecallManagementServiceVersionConfigRequestMergeConfig {
	return s.MergeConfig
}

func (s *UpdateRecallManagementServiceVersionConfigRequest) GetRecallConfig() *UpdateRecallManagementServiceVersionConfigRequestRecallConfig {
	return s.RecallConfig
}

func (s *UpdateRecallManagementServiceVersionConfigRequest) SetRegionId(v string) *UpdateRecallManagementServiceVersionConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequest) SetConfigType(v string) *UpdateRecallManagementServiceVersionConfigRequest {
	s.ConfigType = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequest) SetInstanceId(v string) *UpdateRecallManagementServiceVersionConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequest) SetMergeConfig(v *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) *UpdateRecallManagementServiceVersionConfigRequest {
	s.MergeConfig = v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequest) SetRecallConfig(v *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) *UpdateRecallManagementServiceVersionConfigRequest {
	s.RecallConfig = v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequest) Validate() error {
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

type UpdateRecallManagementServiceVersionConfigRequestMergeConfig struct {
	ExtendedConfig                         *string   `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	FilterExpression                       *string   `json:"FilterExpression,omitempty" xml:"FilterExpression,omitempty"`
	FilterRecallManagementTableIds         []*string `json:"FilterRecallManagementTableIds,omitempty" xml:"FilterRecallManagementTableIds,omitempty" type:"Repeated"`
	ItemRecallManagementTableId            *string   `json:"ItemRecallManagementTableId,omitempty" xml:"ItemRecallManagementTableId,omitempty"`
	ItemTableFields                        []*string `json:"ItemTableFields,omitempty" xml:"ItemTableFields,omitempty" type:"Repeated"`
	MergeType                              *string   `json:"MergeType,omitempty" xml:"MergeType,omitempty"`
	RecallManagementServiceVersionConfigId *string   `json:"RecallManagementServiceVersionConfigId,omitempty" xml:"RecallManagementServiceVersionConfigId,omitempty"`
}

func (s UpdateRecallManagementServiceVersionConfigRequestMergeConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementServiceVersionConfigRequestMergeConfig) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) GetExtendedConfig() *string {
	return s.ExtendedConfig
}

func (s *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) GetFilterExpression() *string {
	return s.FilterExpression
}

func (s *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) GetFilterRecallManagementTableIds() []*string {
	return s.FilterRecallManagementTableIds
}

func (s *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) GetItemRecallManagementTableId() *string {
	return s.ItemRecallManagementTableId
}

func (s *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) GetItemTableFields() []*string {
	return s.ItemTableFields
}

func (s *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) GetMergeType() *string {
	return s.MergeType
}

func (s *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) GetRecallManagementServiceVersionConfigId() *string {
	return s.RecallManagementServiceVersionConfigId
}

func (s *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) SetExtendedConfig(v string) *UpdateRecallManagementServiceVersionConfigRequestMergeConfig {
	s.ExtendedConfig = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) SetFilterExpression(v string) *UpdateRecallManagementServiceVersionConfigRequestMergeConfig {
	s.FilterExpression = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) SetFilterRecallManagementTableIds(v []*string) *UpdateRecallManagementServiceVersionConfigRequestMergeConfig {
	s.FilterRecallManagementTableIds = v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) SetItemRecallManagementTableId(v string) *UpdateRecallManagementServiceVersionConfigRequestMergeConfig {
	s.ItemRecallManagementTableId = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) SetItemTableFields(v []*string) *UpdateRecallManagementServiceVersionConfigRequestMergeConfig {
	s.ItemTableFields = v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) SetMergeType(v string) *UpdateRecallManagementServiceVersionConfigRequestMergeConfig {
	s.MergeType = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) SetRecallManagementServiceVersionConfigId(v string) *UpdateRecallManagementServiceVersionConfigRequestMergeConfig {
	s.RecallManagementServiceVersionConfigId = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestMergeConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRecallManagementServiceVersionConfigRequestRecallConfig struct {
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExtendedConfig *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// example:
	//
	// ["type":"equal"]
	ItemConditionArray *string `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	// example:
	//
	// age>20
	ItemConditionExpression           *string                                                                   `json:"ItemConditionExpression,omitempty" xml:"ItemConditionExpression,omitempty"`
	ItemVectorField                   *string                                                                   `json:"ItemVectorField,omitempty" xml:"ItemVectorField,omitempty"`
	ItemVectorRecallManagementTableId *string                                                                   `json:"ItemVectorRecallManagementTableId,omitempty" xml:"ItemVectorRecallManagementTableId,omitempty"`
	Name                              *string                                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Operators                         []*UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Repeated"`
	Priority                          *int64                                                                    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RecallManagementTableId           *string                                                                   `json:"RecallManagementTableId,omitempty" xml:"RecallManagementTableId,omitempty"`
	RecallType                        *string                                                                   `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	UserVectorField                   *string                                                                   `json:"UserVectorField,omitempty" xml:"UserVectorField,omitempty"`
	UserVectorRecallManagementTableId *string                                                                   `json:"UserVectorRecallManagementTableId,omitempty" xml:"UserVectorRecallManagementTableId,omitempty"`
}

func (s UpdateRecallManagementServiceVersionConfigRequestRecallConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementServiceVersionConfigRequestRecallConfig) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) GetDescription() *string {
	return s.Description
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) GetExtendedConfig() *string {
	return s.ExtendedConfig
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) GetItemConditionArray() *string {
	return s.ItemConditionArray
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) GetItemConditionExpression() *string {
	return s.ItemConditionExpression
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) GetItemVectorField() *string {
	return s.ItemVectorField
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) GetItemVectorRecallManagementTableId() *string {
	return s.ItemVectorRecallManagementTableId
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) GetName() *string {
	return s.Name
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) GetOperators() []*UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators {
	return s.Operators
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) GetPriority() *int64 {
	return s.Priority
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) GetRecallManagementTableId() *string {
	return s.RecallManagementTableId
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) GetRecallType() *string {
	return s.RecallType
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) GetUserVectorField() *string {
	return s.UserVectorField
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) GetUserVectorRecallManagementTableId() *string {
	return s.UserVectorRecallManagementTableId
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) SetDescription(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.Description = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) SetExtendedConfig(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.ExtendedConfig = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) SetItemConditionArray(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.ItemConditionArray = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) SetItemConditionExpression(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.ItemConditionExpression = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) SetItemVectorField(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.ItemVectorField = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) SetItemVectorRecallManagementTableId(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.ItemVectorRecallManagementTableId = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) SetName(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.Name = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) SetOperators(v []*UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators) *UpdateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.Operators = v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) SetPriority(v int64) *UpdateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.Priority = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) SetRecallManagementTableId(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.RecallManagementTableId = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) SetRecallType(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.RecallType = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) SetUserVectorField(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.UserVectorField = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) SetUserVectorRecallManagementTableId(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfig {
	s.UserVectorRecallManagementTableId = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfig) Validate() error {
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

type UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators struct {
	FeatureConfig *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig `json:"FeatureConfig,omitempty" xml:"FeatureConfig,omitempty" type:"Struct"`
	FilterConfig  *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig  `json:"FilterConfig,omitempty" xml:"FilterConfig,omitempty" type:"Struct"`
	JoinConfig    *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig    `json:"JoinConfig,omitempty" xml:"JoinConfig,omitempty" type:"Struct"`
	OperatorsType *string                                                                              `json:"OperatorsType,omitempty" xml:"OperatorsType,omitempty"`
	TriggerConfig *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty" type:"Struct"`
}

func (s UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators) GetFeatureConfig() *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig {
	return s.FeatureConfig
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators) GetFilterConfig() *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig {
	return s.FilterConfig
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators) GetJoinConfig() *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig {
	return s.JoinConfig
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators) GetOperatorsType() *string {
	return s.OperatorsType
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators) GetTriggerConfig() *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig {
	return s.TriggerConfig
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators) SetFeatureConfig(v *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators {
	s.FeatureConfig = v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators) SetFilterConfig(v *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig) *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators {
	s.FilterConfig = v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators) SetJoinConfig(v *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators {
	s.JoinConfig = v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators) SetOperatorsType(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators {
	s.OperatorsType = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators) SetTriggerConfig(v *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators {
	s.TriggerConfig = v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperators) Validate() error {
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

type UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig struct {
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) GetExpression() *string {
	return s.Expression
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) GetName() *string {
	return s.Name
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) GetType() *string {
	return s.Type
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) SetExpression(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig {
	s.Expression = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) SetName(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig {
	s.Name = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) SetType(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig {
	s.Type = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFeatureConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig struct {
	Experession *string `json:"Experession,omitempty" xml:"Experession,omitempty"`
}

func (s UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig) GetExperession() *string {
	return s.Experession
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig) SetExperession(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig {
	s.Experession = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsFilterConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig struct {
	Field                   *string `json:"Field,omitempty" xml:"Field,omitempty"`
	OutputFields            *string `json:"OutputFields,omitempty" xml:"OutputFields,omitempty"`
	RecallManagementTableId *string `json:"RecallManagementTableId,omitempty" xml:"RecallManagementTableId,omitempty"`
}

func (s UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) GetField() *string {
	return s.Field
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) GetOutputFields() *string {
	return s.OutputFields
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) GetRecallManagementTableId() *string {
	return s.RecallManagementTableId
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) SetField(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig {
	s.Field = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) SetOutputFields(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig {
	s.OutputFields = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) SetRecallManagementTableId(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig {
	s.RecallManagementTableId = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsJoinConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig struct {
	Field              *string `json:"Field,omitempty" xml:"Field,omitempty"`
	FieldQuantityLimit *string `json:"FieldQuantityLimit,omitempty" xml:"FieldQuantityLimit,omitempty"`
	IsRandSort         *string `json:"IsRandSort,omitempty" xml:"IsRandSort,omitempty"`
	SortField          *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
}

func (s UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) GetField() *string {
	return s.Field
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) GetFieldQuantityLimit() *string {
	return s.FieldQuantityLimit
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) GetIsRandSort() *string {
	return s.IsRandSort
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) GetSortField() *string {
	return s.SortField
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) SetField(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig {
	s.Field = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) SetFieldQuantityLimit(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig {
	s.FieldQuantityLimit = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) SetIsRandSort(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig {
	s.IsRandSort = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) SetSortField(v string) *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig {
	s.SortField = &v
	return s
}

func (s *UpdateRecallManagementServiceVersionConfigRequestRecallConfigOperatorsTriggerConfig) Validate() error {
	return dara.Validate(s)
}
