// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementServiceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v *GetRecallManagementServiceVersionResponseBodyConfigs) *GetRecallManagementServiceVersionResponseBody
	GetConfigs() *GetRecallManagementServiceVersionResponseBodyConfigs
	SetGmtCreateTime(v string) *GetRecallManagementServiceVersionResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetRecallManagementServiceVersionResponseBody
	GetGmtModifiedTime() *string
	SetIsDefault(v string) *GetRecallManagementServiceVersionResponseBody
	GetIsDefault() *string
	SetName(v string) *GetRecallManagementServiceVersionResponseBody
	GetName() *string
	SetRecallManagementServiceVersionId(v string) *GetRecallManagementServiceVersionResponseBody
	GetRecallManagementServiceVersionId() *string
	SetRequestId(v string) *GetRecallManagementServiceVersionResponseBody
	GetRequestId() *string
}

type GetRecallManagementServiceVersionResponseBody struct {
	Configs *GetRecallManagementServiceVersionResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Struct"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// false
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// V1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2
	RecallManagementServiceVersionId *string `json:"RecallManagementServiceVersionId,omitempty" xml:"RecallManagementServiceVersionId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E15A1443-7917-5BE0-AE70-25538ECF398D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRecallManagementServiceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionResponseBody) GetConfigs() *GetRecallManagementServiceVersionResponseBodyConfigs {
	return s.Configs
}

func (s *GetRecallManagementServiceVersionResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetRecallManagementServiceVersionResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetRecallManagementServiceVersionResponseBody) GetIsDefault() *string {
	return s.IsDefault
}

func (s *GetRecallManagementServiceVersionResponseBody) GetName() *string {
	return s.Name
}

func (s *GetRecallManagementServiceVersionResponseBody) GetRecallManagementServiceVersionId() *string {
	return s.RecallManagementServiceVersionId
}

func (s *GetRecallManagementServiceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRecallManagementServiceVersionResponseBody) SetConfigs(v *GetRecallManagementServiceVersionResponseBodyConfigs) *GetRecallManagementServiceVersionResponseBody {
	s.Configs = v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBody) SetGmtCreateTime(v string) *GetRecallManagementServiceVersionResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBody) SetGmtModifiedTime(v string) *GetRecallManagementServiceVersionResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBody) SetIsDefault(v string) *GetRecallManagementServiceVersionResponseBody {
	s.IsDefault = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBody) SetName(v string) *GetRecallManagementServiceVersionResponseBody {
	s.Name = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBody) SetRecallManagementServiceVersionId(v string) *GetRecallManagementServiceVersionResponseBody {
	s.RecallManagementServiceVersionId = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBody) SetRequestId(v string) *GetRecallManagementServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBody) Validate() error {
	if s.Configs != nil {
		if err := s.Configs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRecallManagementServiceVersionResponseBodyConfigs struct {
	MergeConfig   *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig     `json:"MergeConfig,omitempty" xml:"MergeConfig,omitempty" type:"Struct"`
	RecallConfigs []*GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs `json:"RecallConfigs,omitempty" xml:"RecallConfigs,omitempty" type:"Repeated"`
}

func (s GetRecallManagementServiceVersionResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigs) GetMergeConfig() *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig {
	return s.MergeConfig
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigs) GetRecallConfigs() []*GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs {
	return s.RecallConfigs
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigs) SetMergeConfig(v *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) *GetRecallManagementServiceVersionResponseBodyConfigs {
	s.MergeConfig = v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigs) SetRecallConfigs(v []*GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) *GetRecallManagementServiceVersionResponseBodyConfigs {
	s.RecallConfigs = v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigs) Validate() error {
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

type GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig struct {
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

func (s GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) GetExtendedConfig() *string {
	return s.ExtendedConfig
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) GetFilterExpression() *string {
	return s.FilterExpression
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) GetFilterRecallManagementTableIds() []*string {
	return s.FilterRecallManagementTableIds
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) GetItemRecallManagementTableId() *string {
	return s.ItemRecallManagementTableId
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) GetItemTableFields() []*string {
	return s.ItemTableFields
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) GetMergeType() *string {
	return s.MergeType
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) GetRecallManagementServiceVersionConfigId() *string {
	return s.RecallManagementServiceVersionConfigId
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) SetExtendedConfig(v string) *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig {
	s.ExtendedConfig = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) SetFilterExpression(v string) *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig {
	s.FilterExpression = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) SetFilterRecallManagementTableIds(v []*string) *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig {
	s.FilterRecallManagementTableIds = v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) SetItemRecallManagementTableId(v string) *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig {
	s.ItemRecallManagementTableId = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) SetItemTableFields(v []*string) *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig {
	s.ItemTableFields = v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) SetMergeType(v string) *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig {
	s.MergeType = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) SetRecallManagementServiceVersionConfigId(v string) *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig {
	s.RecallManagementServiceVersionConfigId = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsMergeConfig) Validate() error {
	return dara.Validate(s)
}

type GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs struct {
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
	// 4
	ItemVectorRecallManagementTableId *string `json:"ItemVectorRecallManagementTableId,omitempty" xml:"ItemVectorRecallManagementTableId,omitempty"`
	// example:
	//
	// etrec
	Name      *string                                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	Operators []*GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 5
	RecallManagementServiceVersionConfigId *string `json:"RecallManagementServiceVersionConfigId,omitempty" xml:"RecallManagementServiceVersionConfigId,omitempty"`
	// example:
	//
	// 2
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

func (s GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) GetDescription() *string {
	return s.Description
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) GetExtendedConfig() *string {
	return s.ExtendedConfig
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) GetItemConditionArray() *string {
	return s.ItemConditionArray
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) GetItemVectorField() *string {
	return s.ItemVectorField
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) GetItemVectorRecallManagementTableId() *string {
	return s.ItemVectorRecallManagementTableId
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) GetName() *string {
	return s.Name
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) GetOperators() []*GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators {
	return s.Operators
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) GetPriority() *int64 {
	return s.Priority
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) GetRecallManagementServiceVersionConfigId() *string {
	return s.RecallManagementServiceVersionConfigId
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) GetRecallManagementTableId() *string {
	return s.RecallManagementTableId
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) GetRecallType() *string {
	return s.RecallType
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) GetUserVectorField() *string {
	return s.UserVectorField
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) GetUserVectorRecallManagementTableId() *string {
	return s.UserVectorRecallManagementTableId
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) SetDescription(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs {
	s.Description = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) SetExtendedConfig(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs {
	s.ExtendedConfig = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) SetItemConditionArray(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs {
	s.ItemConditionArray = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) SetItemVectorField(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs {
	s.ItemVectorField = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) SetItemVectorRecallManagementTableId(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs {
	s.ItemVectorRecallManagementTableId = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) SetName(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs {
	s.Name = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) SetOperators(v []*GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs {
	s.Operators = v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) SetPriority(v int64) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs {
	s.Priority = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) SetRecallManagementServiceVersionConfigId(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs {
	s.RecallManagementServiceVersionConfigId = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) SetRecallManagementTableId(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs {
	s.RecallManagementTableId = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) SetRecallType(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs {
	s.RecallType = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) SetUserVectorField(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs {
	s.UserVectorField = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) SetUserVectorRecallManagementTableId(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs {
	s.UserVectorRecallManagementTableId = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigs) Validate() error {
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

type GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators struct {
	FeatureConfig *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFeatureConfig `json:"FeatureConfig,omitempty" xml:"FeatureConfig,omitempty" type:"Struct"`
	FilterConfig  *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFilterConfig  `json:"FilterConfig,omitempty" xml:"FilterConfig,omitempty" type:"Struct"`
	JoinConfig    *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsJoinConfig    `json:"JoinConfig,omitempty" xml:"JoinConfig,omitempty" type:"Struct"`
	// example:
	//
	// Filter
	OperatorType  *string                                                                                  `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	TriggerConfig *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty" type:"Struct"`
}

func (s GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators) GetFeatureConfig() *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFeatureConfig {
	return s.FeatureConfig
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators) GetFilterConfig() *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFilterConfig {
	return s.FilterConfig
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators) GetJoinConfig() *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsJoinConfig {
	return s.JoinConfig
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators) GetOperatorType() *string {
	return s.OperatorType
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators) GetTriggerConfig() *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig {
	return s.TriggerConfig
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators) SetFeatureConfig(v *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFeatureConfig) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators {
	s.FeatureConfig = v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators) SetFilterConfig(v *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFilterConfig) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators {
	s.FilterConfig = v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators) SetJoinConfig(v *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsJoinConfig) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators {
	s.JoinConfig = v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators) SetOperatorType(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators {
	s.OperatorType = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators) SetTriggerConfig(v *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators {
	s.TriggerConfig = v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperators) Validate() error {
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

type GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFeatureConfig struct {
	// example:
	//
	// city = \\"hangzhou\\"
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// city
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFeatureConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFeatureConfig) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFeatureConfig) GetExpression() *string {
	return s.Expression
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFeatureConfig) GetName() *string {
	return s.Name
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFeatureConfig) GetType() *string {
	return s.Type
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFeatureConfig) SetExpression(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFeatureConfig {
	s.Expression = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFeatureConfig) SetName(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFeatureConfig {
	s.Name = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFeatureConfig) SetType(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFeatureConfig {
	s.Type = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFeatureConfig) Validate() error {
	return dara.Validate(s)
}

type GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFilterConfig struct {
	// example:
	//
	// age > 10
	Experession *string `json:"Experession,omitempty" xml:"Experession,omitempty"`
}

func (s GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFilterConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFilterConfig) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFilterConfig) GetExperession() *string {
	return s.Experession
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFilterConfig) SetExperession(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFilterConfig {
	s.Experession = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsFilterConfig) Validate() error {
	return dara.Validate(s)
}

type GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsJoinConfig struct {
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

func (s GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsJoinConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsJoinConfig) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsJoinConfig) GetField() *string {
	return s.Field
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsJoinConfig) GetOutputFields() []*string {
	return s.OutputFields
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsJoinConfig) GetRecallManagementTableId() *string {
	return s.RecallManagementTableId
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsJoinConfig) SetField(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsJoinConfig {
	s.Field = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsJoinConfig) SetOutputFields(v []*string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsJoinConfig {
	s.OutputFields = v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsJoinConfig) SetRecallManagementTableId(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsJoinConfig {
	s.RecallManagementTableId = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsJoinConfig) Validate() error {
	return dara.Validate(s)
}

type GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig struct {
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

func (s GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig) GetField() *string {
	return s.Field
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig) GetFieldQuantityLimit() *string {
	return s.FieldQuantityLimit
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig) GetIsRandSort() *string {
	return s.IsRandSort
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig) GetSortField() *string {
	return s.SortField
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig) SetField(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig {
	s.Field = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig) SetFieldQuantityLimit(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig {
	s.FieldQuantityLimit = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig) SetIsRandSort(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig {
	s.IsRandSort = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig) SetSortField(v string) *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig {
	s.SortField = &v
	return s
}

func (s *GetRecallManagementServiceVersionResponseBodyConfigsRecallConfigsOperatorsTriggerConfig) Validate() error {
	return dara.Validate(s)
}
