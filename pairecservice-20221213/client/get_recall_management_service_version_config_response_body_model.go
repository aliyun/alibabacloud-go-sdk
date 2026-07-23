// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecallManagementServiceVersionConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigType(v string) *GetRecallManagementServiceVersionConfigResponseBody
	GetConfigType() *string
	SetGmtCreateTime(v string) *GetRecallManagementServiceVersionConfigResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetRecallManagementServiceVersionConfigResponseBody
	GetGmtModifiedTime() *string
	SetMergeConfig(v *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) *GetRecallManagementServiceVersionConfigResponseBody
	GetMergeConfig() *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig
	SetRecallConfig(v *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) *GetRecallManagementServiceVersionConfigResponseBody
	GetRecallConfig() *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig
	SetRecallManagementServiceId(v string) *GetRecallManagementServiceVersionConfigResponseBody
	GetRecallManagementServiceId() *string
	SetRecallManagementServiceVersionConfigId(v string) *GetRecallManagementServiceVersionConfigResponseBody
	GetRecallManagementServiceVersionConfigId() *string
	SetRecallManagementServiceVersionId(v string) *GetRecallManagementServiceVersionConfigResponseBody
	GetRecallManagementServiceVersionId() *string
	SetRequestId(v string) *GetRecallManagementServiceVersionConfigResponseBody
	GetRequestId() *string
}

type GetRecallManagementServiceVersionConfigResponseBody struct {
	// The configuration type. Valid values: Recall (recall configuration) and Merge (merge configuration).
	//
	// example:
	//
	// Recall
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The creation time, in ISO 8601 format.
	//
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The last modification time, in ISO 8601 format.
	//
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The merge configuration.
	MergeConfig *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig `json:"MergeConfig,omitempty" xml:"MergeConfig,omitempty" type:"Struct"`
	// The recall configuration.
	RecallConfig *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig `json:"RecallConfig,omitempty" xml:"RecallConfig,omitempty" type:"Struct"`
	// The recall management service ID.
	//
	// example:
	//
	// 3
	RecallManagementServiceId *string `json:"RecallManagementServiceId,omitempty" xml:"RecallManagementServiceId,omitempty"`
	// The recall management service version configuration ID.
	//
	// example:
	//
	// 2
	RecallManagementServiceVersionConfigId *string `json:"RecallManagementServiceVersionConfigId,omitempty" xml:"RecallManagementServiceVersionConfigId,omitempty"`
	// The recall management service version ID.
	//
	// example:
	//
	// 2
	RecallManagementServiceVersionId *string `json:"RecallManagementServiceVersionId,omitempty" xml:"RecallManagementServiceVersionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRecallManagementServiceVersionConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) GetMergeConfig() *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig {
	return s.MergeConfig
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) GetRecallConfig() *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig {
	return s.RecallConfig
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) GetRecallManagementServiceId() *string {
	return s.RecallManagementServiceId
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) GetRecallManagementServiceVersionConfigId() *string {
	return s.RecallManagementServiceVersionConfigId
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) GetRecallManagementServiceVersionId() *string {
	return s.RecallManagementServiceVersionId
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) SetConfigType(v string) *GetRecallManagementServiceVersionConfigResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) SetGmtCreateTime(v string) *GetRecallManagementServiceVersionConfigResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) SetGmtModifiedTime(v string) *GetRecallManagementServiceVersionConfigResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) SetMergeConfig(v *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) *GetRecallManagementServiceVersionConfigResponseBody {
	s.MergeConfig = v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) SetRecallConfig(v *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) *GetRecallManagementServiceVersionConfigResponseBody {
	s.RecallConfig = v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) SetRecallManagementServiceId(v string) *GetRecallManagementServiceVersionConfigResponseBody {
	s.RecallManagementServiceId = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) SetRecallManagementServiceVersionConfigId(v string) *GetRecallManagementServiceVersionConfigResponseBody {
	s.RecallManagementServiceVersionConfigId = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) SetRecallManagementServiceVersionId(v string) *GetRecallManagementServiceVersionConfigResponseBody {
	s.RecallManagementServiceVersionId = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) SetRequestId(v string) *GetRecallManagementServiceVersionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBody) Validate() error {
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

type GetRecallManagementServiceVersionConfigResponseBodyMergeConfig struct {
	// Additional configuration for the merge operation, reserved for future enhancements.
	//
	// example:
	//
	// {"":""}
	ExtendedConfig *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// The filter expression.
	//
	// example:
	//
	// age>20
	FilterExpression *string `json:"FilterExpression,omitempty" xml:"FilterExpression,omitempty"`
	// A list of filter table IDs.
	FilterRecallManagementTableIds []*string `json:"FilterRecallManagementTableIds,omitempty" xml:"FilterRecallManagementTableIds,omitempty" type:"Repeated"`
	// The item table ID.
	//
	// example:
	//
	// 2
	ItemRecallManagementTableId *string `json:"ItemRecallManagementTableId,omitempty" xml:"ItemRecallManagementTableId,omitempty"`
	// A list of output fields from the item table.
	ItemTableFields []*string `json:"ItemTableFields,omitempty" xml:"ItemTableFields,omitempty" type:"Repeated"`
	// The merge type. Valid values:
	//
	// - Weight: Merges items based on assigned weights.
	//
	// - Alternate: Merges items in an alternating sequence.
	//
	// example:
	//
	// Weight
	MergeType *string `json:"MergeType,omitempty" xml:"MergeType,omitempty"`
	// The recall management service version configuration ID.
	//
	// example:
	//
	// 1
	RecallManagementServiceVersionConfigId *string `json:"RecallManagementServiceVersionConfigId,omitempty" xml:"RecallManagementServiceVersionConfigId,omitempty"`
}

func (s GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) GetExtendedConfig() *string {
	return s.ExtendedConfig
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) GetFilterExpression() *string {
	return s.FilterExpression
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) GetFilterRecallManagementTableIds() []*string {
	return s.FilterRecallManagementTableIds
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) GetItemRecallManagementTableId() *string {
	return s.ItemRecallManagementTableId
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) GetItemTableFields() []*string {
	return s.ItemTableFields
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) GetMergeType() *string {
	return s.MergeType
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) GetRecallManagementServiceVersionConfigId() *string {
	return s.RecallManagementServiceVersionConfigId
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) SetExtendedConfig(v string) *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig {
	s.ExtendedConfig = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) SetFilterExpression(v string) *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig {
	s.FilterExpression = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) SetFilterRecallManagementTableIds(v []*string) *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig {
	s.FilterRecallManagementTableIds = v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) SetItemRecallManagementTableId(v string) *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig {
	s.ItemRecallManagementTableId = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) SetItemTableFields(v []*string) *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig {
	s.ItemTableFields = v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) SetMergeType(v string) *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig {
	s.MergeType = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) SetRecallManagementServiceVersionConfigId(v string) *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig {
	s.RecallManagementServiceVersionConfigId = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyMergeConfig) Validate() error {
	return dara.Validate(s)
}

type GetRecallManagementServiceVersionConfigResponseBodyRecallConfig struct {
	// The recall description.
	//
	// example:
	//
	// this is etrec recall
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Extended configuration, reserved for future use.
	//
	// example:
	//
	// {"":""}
	ExtendedConfig *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// The item conditions, as a JSON array string.
	//
	// example:
	//
	// [{"option":"<","field":"category","type":"STRING","value":"10"}]
	ItemConditionArray *string `json:"ItemConditionArray,omitempty" xml:"ItemConditionArray,omitempty"`
	// The item condition expression.
	//
	// example:
	//
	// age>20
	ItemConditionExpression *string `json:"ItemConditionExpression,omitempty" xml:"ItemConditionExpression,omitempty"`
	// The item vector field.
	//
	// example:
	//
	// item_embedding
	ItemVectorField *string `json:"ItemVectorField,omitempty" xml:"ItemVectorField,omitempty"`
	// The item vector recall table ID.
	//
	// example:
	//
	// 2
	ItemVectorRecallManagementTableId *string `json:"ItemVectorRecallManagementTableId,omitempty" xml:"ItemVectorRecallManagementTableId,omitempty"`
	// The recall name.
	//
	// example:
	//
	// etrec
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A list of operators.
	Operators []*GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Repeated"`
	// The priority. A lower value indicates a higher priority.
	//
	// example:
	//
	// 10
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The recall management service version configuration ID.
	//
	// example:
	//
	// 4
	RecallManagementServiceVersionConfigId *string `json:"RecallManagementServiceVersionConfigId,omitempty" xml:"RecallManagementServiceVersionConfigId,omitempty"`
	// The recall management table ID.
	//
	// example:
	//
	// 3
	RecallManagementTableId *string `json:"RecallManagementTableId,omitempty" xml:"RecallManagementTableId,omitempty"`
	// The recall type.
	//
	// example:
	//
	// X2I
	RecallType *string `json:"RecallType,omitempty" xml:"RecallType,omitempty"`
	// The sort fields.
	//
	// example:
	//
	// name
	SortFields *string `json:"SortFields,omitempty" xml:"SortFields,omitempty"`
	// The user vector field.
	//
	// example:
	//
	// user_embedding
	UserVectorField *string `json:"UserVectorField,omitempty" xml:"UserVectorField,omitempty"`
	// The user vector recall table ID.
	//
	// example:
	//
	// 2
	UserVectorRecallManagementTableId *string `json:"UserVectorRecallManagementTableId,omitempty" xml:"UserVectorRecallManagementTableId,omitempty"`
}

func (s GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) GetDescription() *string {
	return s.Description
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) GetExtendedConfig() *string {
	return s.ExtendedConfig
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) GetItemConditionArray() *string {
	return s.ItemConditionArray
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) GetItemConditionExpression() *string {
	return s.ItemConditionExpression
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) GetItemVectorField() *string {
	return s.ItemVectorField
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) GetItemVectorRecallManagementTableId() *string {
	return s.ItemVectorRecallManagementTableId
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) GetName() *string {
	return s.Name
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) GetOperators() []*GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators {
	return s.Operators
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) GetPriority() *int64 {
	return s.Priority
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) GetRecallManagementServiceVersionConfigId() *string {
	return s.RecallManagementServiceVersionConfigId
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) GetRecallManagementTableId() *string {
	return s.RecallManagementTableId
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) GetRecallType() *string {
	return s.RecallType
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) GetSortFields() *string {
	return s.SortFields
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) GetUserVectorField() *string {
	return s.UserVectorField
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) GetUserVectorRecallManagementTableId() *string {
	return s.UserVectorRecallManagementTableId
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) SetDescription(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig {
	s.Description = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) SetExtendedConfig(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig {
	s.ExtendedConfig = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) SetItemConditionArray(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig {
	s.ItemConditionArray = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) SetItemConditionExpression(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig {
	s.ItemConditionExpression = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) SetItemVectorField(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig {
	s.ItemVectorField = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) SetItemVectorRecallManagementTableId(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig {
	s.ItemVectorRecallManagementTableId = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) SetName(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig {
	s.Name = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) SetOperators(v []*GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig {
	s.Operators = v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) SetPriority(v int64) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig {
	s.Priority = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) SetRecallManagementServiceVersionConfigId(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig {
	s.RecallManagementServiceVersionConfigId = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) SetRecallManagementTableId(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig {
	s.RecallManagementTableId = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) SetRecallType(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig {
	s.RecallType = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) SetSortFields(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig {
	s.SortFields = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) SetUserVectorField(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig {
	s.UserVectorField = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) SetUserVectorRecallManagementTableId(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig {
	s.UserVectorRecallManagementTableId = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfig) Validate() error {
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

type GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators struct {
	// The configuration for the feature operator.
	FeatureConfig *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFeatureConfig `json:"FeatureConfig,omitempty" xml:"FeatureConfig,omitempty" type:"Struct"`
	// The configuration for the filter operator.
	FilterConfig *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFilterConfig `json:"FilterConfig,omitempty" xml:"FilterConfig,omitempty" type:"Struct"`
	// The configuration for the Join operator.
	JoinConfig *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsJoinConfig `json:"JoinConfig,omitempty" xml:"JoinConfig,omitempty" type:"Struct"`
	// The operator type.
	//
	// example:
	//
	// Filter
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// The configuration for the Trigger operator.
	TriggerConfig *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig `json:"TriggerConfig,omitempty" xml:"TriggerConfig,omitempty" type:"Struct"`
}

func (s GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators) GetFeatureConfig() *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFeatureConfig {
	return s.FeatureConfig
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators) GetFilterConfig() *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFilterConfig {
	return s.FilterConfig
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators) GetJoinConfig() *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsJoinConfig {
	return s.JoinConfig
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators) GetOperatorType() *string {
	return s.OperatorType
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators) GetTriggerConfig() *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig {
	return s.TriggerConfig
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators) SetFeatureConfig(v *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFeatureConfig) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators {
	s.FeatureConfig = v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators) SetFilterConfig(v *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFilterConfig) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators {
	s.FilterConfig = v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators) SetJoinConfig(v *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsJoinConfig) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators {
	s.JoinConfig = v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators) SetOperatorType(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators {
	s.OperatorType = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators) SetTriggerConfig(v *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators {
	s.TriggerConfig = v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperators) Validate() error {
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

type GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFeatureConfig struct {
	// The feature expression.
	//
	// example:
	//
	// category=3
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The feature name.
	//
	// example:
	//
	// age
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The feature type.
	//
	// example:
	//
	// STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFeatureConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFeatureConfig) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFeatureConfig) GetExpression() *string {
	return s.Expression
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFeatureConfig) GetName() *string {
	return s.Name
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFeatureConfig) GetType() *string {
	return s.Type
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFeatureConfig) SetExpression(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFeatureConfig {
	s.Expression = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFeatureConfig) SetName(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFeatureConfig {
	s.Name = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFeatureConfig) SetType(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFeatureConfig {
	s.Type = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFeatureConfig) Validate() error {
	return dara.Validate(s)
}

type GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFilterConfig struct {
	// The filter expression.
	//
	// example:
	//
	// age>20
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
}

func (s GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFilterConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFilterConfig) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFilterConfig) GetExpression() *string {
	return s.Expression
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFilterConfig) SetExpression(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFilterConfig {
	s.Expression = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsFilterConfig) Validate() error {
	return dara.Validate(s)
}

type GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsJoinConfig struct {
	// The join field.
	//
	// example:
	//
	// item_id
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The output fields.
	OutputFields []*string `json:"OutputFields,omitempty" xml:"OutputFields,omitempty" type:"Repeated"`
	// The join table ID.
	//
	// example:
	//
	// 3
	RecallManagementTableId *string `json:"RecallManagementTableId,omitempty" xml:"RecallManagementTableId,omitempty"`
}

func (s GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsJoinConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsJoinConfig) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsJoinConfig) GetField() *string {
	return s.Field
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsJoinConfig) GetOutputFields() []*string {
	return s.OutputFields
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsJoinConfig) GetRecallManagementTableId() *string {
	return s.RecallManagementTableId
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsJoinConfig) SetField(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsJoinConfig {
	s.Field = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsJoinConfig) SetOutputFields(v []*string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsJoinConfig {
	s.OutputFields = v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsJoinConfig) SetRecallManagementTableId(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsJoinConfig {
	s.RecallManagementTableId = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsJoinConfig) Validate() error {
	return dara.Validate(s)
}

type GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig struct {
	// The field name.
	//
	// example:
	//
	// user_id
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The limit on the number of fields.
	//
	// example:
	//
	// 20
	FieldQuantityLimit *int32 `json:"FieldQuantityLimit,omitempty" xml:"FieldQuantityLimit,omitempty"`
	// Indicates whether to randomly sort the items.
	//
	// example:
	//
	// false
	IsRandSort *bool `json:"IsRandSort,omitempty" xml:"IsRandSort,omitempty"`
	// The sort field.
	//
	// example:
	//
	// create_time
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
}

func (s GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig) GoString() string {
	return s.String()
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig) GetField() *string {
	return s.Field
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig) GetFieldQuantityLimit() *int32 {
	return s.FieldQuantityLimit
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig) GetIsRandSort() *bool {
	return s.IsRandSort
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig) GetSortField() *string {
	return s.SortField
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig) SetField(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig {
	s.Field = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig) SetFieldQuantityLimit(v int32) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig {
	s.FieldQuantityLimit = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig) SetIsRandSort(v bool) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig {
	s.IsRandSort = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig) SetSortField(v string) *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig {
	s.SortField = &v
	return s
}

func (s *GetRecallManagementServiceVersionConfigResponseBodyRecallConfigOperatorsTriggerConfig) Validate() error {
	return dara.Validate(s)
}
