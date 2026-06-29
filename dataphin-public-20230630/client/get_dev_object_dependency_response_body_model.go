// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDevObjectDependencyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDevObjectDependencyResponseBody
	GetCode() *string
	SetDevObjectDependencyList(v []*GetDevObjectDependencyResponseBodyDevObjectDependencyList) *GetDevObjectDependencyResponseBody
	GetDevObjectDependencyList() []*GetDevObjectDependencyResponseBodyDevObjectDependencyList
	SetHttpStatusCode(v int32) *GetDevObjectDependencyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDevObjectDependencyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDevObjectDependencyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDevObjectDependencyResponseBody
	GetSuccess() *bool
}

type GetDevObjectDependencyResponseBody struct {
	// Error code. OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Dependency list.
	DevObjectDependencyList []*GetDevObjectDependencyResponseBodyDevObjectDependencyList `json:"DevObjectDependencyList,omitempty" xml:"DevObjectDependencyList,omitempty" type:"Repeated"`
	// HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDevObjectDependencyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDevObjectDependencyResponseBody) GoString() string {
	return s.String()
}

func (s *GetDevObjectDependencyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDevObjectDependencyResponseBody) GetDevObjectDependencyList() []*GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	return s.DevObjectDependencyList
}

func (s *GetDevObjectDependencyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDevObjectDependencyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDevObjectDependencyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDevObjectDependencyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDevObjectDependencyResponseBody) SetCode(v string) *GetDevObjectDependencyResponseBody {
	s.Code = &v
	return s
}

func (s *GetDevObjectDependencyResponseBody) SetDevObjectDependencyList(v []*GetDevObjectDependencyResponseBodyDevObjectDependencyList) *GetDevObjectDependencyResponseBody {
	s.DevObjectDependencyList = v
	return s
}

func (s *GetDevObjectDependencyResponseBody) SetHttpStatusCode(v int32) *GetDevObjectDependencyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDevObjectDependencyResponseBody) SetMessage(v string) *GetDevObjectDependencyResponseBody {
	s.Message = &v
	return s
}

func (s *GetDevObjectDependencyResponseBody) SetRequestId(v string) *GetDevObjectDependencyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDevObjectDependencyResponseBody) SetSuccess(v bool) *GetDevObjectDependencyResponseBody {
	s.Success = &v
	return s
}

func (s *GetDevObjectDependencyResponseBody) Validate() error {
	if s.DevObjectDependencyList != nil {
		for _, item := range s.DevObjectDependencyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDevObjectDependencyResponseBodyDevObjectDependencyList struct {
	// Indicates whether the task is automatically parsed.
	//
	// example:
	//
	// true
	AutoParse *bool `json:"AutoParse,omitempty" xml:"AutoParse,omitempty"`
	// Business type.
	//
	// - SCRIPT: Script
	//
	// - LOGICAL_TABLE: Logical table
	//
	// example:
	//
	// SCRIPT
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// Business unit ID.
	//
	// example:
	//
	// 13111
	BizUnitId *string `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// Business unit name.
	//
	// example:
	//
	// xx测试
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// Cron expression for the scheduling node.
	//
	// example:
	//
	// 0 0 0 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// Indicates whether a custom cron expression is used for the scheduling node.
	//
	// example:
	//
	// true
	CustomCronExpression *bool `json:"CustomCronExpression,omitempty" xml:"CustomCronExpression,omitempty"`
	// Dependency fields.
	DependFieldList []*string `json:"DependFieldList,omitempty" xml:"DependFieldList,omitempty" type:"Repeated"`
	// Dependency period configuration.
	DependencyPeriod *GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod `json:"DependencyPeriod,omitempty" xml:"DependencyPeriod,omitempty" type:"Struct"`
	// Dependency strategy.
	//
	// - ALL
	//
	// - FIRST
	//
	// - LAST
	//
	// - NEAR
	//
	// example:
	//
	// ALL
	DependencyStrategy *string `json:"DependencyStrategy,omitempty" xml:"DependencyStrategy,omitempty"`
	// Indicates whether the node is a dimension table intermediate node.
	//
	// example:
	//
	// true
	DimMidNode *bool `json:"DimMidNode,omitempty" xml:"DimMidNode,omitempty"`
	// Effect fields.
	EffectFieldList []*string `json:"EffectFieldList,omitempty" xml:"EffectFieldList,omitempty" type:"Repeated"`
	// Additional business information.
	//
	// example:
	//
	// 所有字段信息/hasProd/hasDev等信息
	ExternalBizInfo *string `json:"ExternalBizInfo,omitempty" xml:"ExternalBizInfo,omitempty"`
	// Indicates whether the dependency is manually added.
	//
	// example:
	//
	// false
	ManuallyAdd *bool `json:"ManuallyAdd,omitempty" xml:"ManuallyAdd,omitempty"`
	// Node ID.
	//
	// example:
	//
	// n_13211
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Node name.
	//
	// example:
	//
	// xx测试
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// Node output name.
	//
	// example:
	//
	// n_xx
	NodeOutputName *string `json:"NodeOutputName,omitempty" xml:"NodeOutputName,omitempty"`
	// Output table information.
	//
	// example:
	//
	// t_xx
	NodeOutputTableName *string `json:"NodeOutputTableName,omitempty" xml:"NodeOutputTableName,omitempty"`
	// Node type.
	//
	// - DATA_PROCESS: Code task
	//
	// - BBOX_LOGIC_TABLE_NODE: Black box logical table node
	//
	// - ONE_ID_LABEL: ID label node
	//
	// - ONE_ID_RULE: ID rule node
	//
	// - PIPELINE_NODE: Pipeline node
	//
	// example:
	//
	// DATA_PROCESS
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// Cross-node output parameters.
	OutputContextParamList []*GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList `json:"OutputContextParamList,omitempty" xml:"OutputContextParamList,omitempty" type:"Repeated"`
	// Node owners.
	OwnerList []*GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	// Dependency period difference.
	//
	// example:
	//
	// 1
	PeriodDiff *int32 `json:"PeriodDiff,omitempty" xml:"PeriodDiff,omitempty"`
	// Project ID.
	//
	// example:
	//
	// 123131
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Project name.
	//
	// example:
	//
	// xx测试
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Schedule type.
	//
	// - MINUTELY: Minute
	//
	// - HOURLY: Hour
	//
	// - DAILY: Day
	//
	// - WEEKLY: Week
	//
	// - MONTHLY: Month
	//
	// - YEARLY: Year
	//
	// example:
	//
	// DAILY
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// Indicates whether the node has a self-dependency.
	//
	// example:
	//
	// true
	SelfDepend *bool `json:"SelfDepend,omitempty" xml:"SelfDepend,omitempty"`
	// Sub-business type.
	//
	// - MAX_COMPUTE_SQL
	//
	// - HIVE_SQL
	//
	// - SHELL
	//
	// - PYTHON
	//
	// - ONE_SERVICE_SQL
	//
	// - DATABASE_SQL, etc.
	//
	// example:
	//
	// SHELL
	SubBizType *string `json:"SubBizType,omitempty" xml:"SubBizType,omitempty"`
	// Indicates whether the configuration is valid.
	//
	// example:
	//
	// true
	Valid *bool `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s GetDevObjectDependencyResponseBodyDevObjectDependencyList) String() string {
	return dara.Prettify(s)
}

func (s GetDevObjectDependencyResponseBodyDevObjectDependencyList) GoString() string {
	return s.String()
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetAutoParse() *bool {
	return s.AutoParse
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetBizType() *string {
	return s.BizType
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetBizUnitId() *string {
	return s.BizUnitId
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetCronExpression() *string {
	return s.CronExpression
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetCustomCronExpression() *bool {
	return s.CustomCronExpression
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetDependFieldList() []*string {
	return s.DependFieldList
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetDependencyPeriod() *GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod {
	return s.DependencyPeriod
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetDependencyStrategy() *string {
	return s.DependencyStrategy
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetDimMidNode() *bool {
	return s.DimMidNode
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetEffectFieldList() []*string {
	return s.EffectFieldList
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetExternalBizInfo() *string {
	return s.ExternalBizInfo
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetManuallyAdd() *bool {
	return s.ManuallyAdd
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetNodeId() *string {
	return s.NodeId
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetNodeName() *string {
	return s.NodeName
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetNodeOutputName() *string {
	return s.NodeOutputName
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetNodeOutputTableName() *string {
	return s.NodeOutputTableName
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetNodeType() *string {
	return s.NodeType
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetOutputContextParamList() []*GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList {
	return s.OutputContextParamList
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetOwnerList() []*GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList {
	return s.OwnerList
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetPeriodDiff() *int32 {
	return s.PeriodDiff
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetSelfDepend() *bool {
	return s.SelfDepend
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetSubBizType() *string {
	return s.SubBizType
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) GetValid() *bool {
	return s.Valid
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetAutoParse(v bool) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.AutoParse = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetBizType(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.BizType = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetBizUnitId(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.BizUnitId = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetBizUnitName(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.BizUnitName = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetCronExpression(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.CronExpression = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetCustomCronExpression(v bool) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.CustomCronExpression = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetDependFieldList(v []*string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.DependFieldList = v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetDependencyPeriod(v *GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.DependencyPeriod = v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetDependencyStrategy(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.DependencyStrategy = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetDimMidNode(v bool) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.DimMidNode = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetEffectFieldList(v []*string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.EffectFieldList = v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetExternalBizInfo(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.ExternalBizInfo = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetManuallyAdd(v bool) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.ManuallyAdd = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetNodeId(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.NodeId = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetNodeName(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.NodeName = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetNodeOutputName(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.NodeOutputName = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetNodeOutputTableName(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.NodeOutputTableName = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetNodeType(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.NodeType = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetOutputContextParamList(v []*GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.OutputContextParamList = v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetOwnerList(v []*GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.OwnerList = v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetPeriodDiff(v int32) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.PeriodDiff = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetProjectId(v int64) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.ProjectId = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetProjectName(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.ProjectName = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetScheduleType(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.ScheduleType = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetSelfDepend(v bool) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.SelfDepend = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetSubBizType(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.SubBizType = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) SetValid(v bool) *GetDevObjectDependencyResponseBodyDevObjectDependencyList {
	s.Valid = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyList) Validate() error {
	if s.DependencyPeriod != nil {
		if err := s.DependencyPeriod.Validate(); err != nil {
			return err
		}
	}
	if s.OutputContextParamList != nil {
		for _, item := range s.OutputContextParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OwnerList != nil {
		for _, item := range s.OwnerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod struct {
	// Period offset. This parameter is required when the dependency period type is LAST_N_PERIOD.
	//
	// example:
	//
	// 1
	PeriodOffset *int32 `json:"PeriodOffset,omitempty" xml:"PeriodOffset,omitempty"`
	// Dependency period type.
	//
	// - CURRENT_PERIOD
	//
	// - LAST_PERIOD
	//
	// - LAST_N_PERIOD
	//
	// - LAST_24_HOUR
	//
	// example:
	//
	// CURRENT_PERIOD
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
}

func (s GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod) String() string {
	return dara.Prettify(s)
}

func (s GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod) GoString() string {
	return s.String()
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod) GetPeriodOffset() *int32 {
	return s.PeriodOffset
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod) GetPeriodType() *string {
	return s.PeriodType
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod) SetPeriodOffset(v int32) *GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod {
	s.PeriodOffset = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod) SetPeriodType(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod {
	s.PeriodType = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListDependencyPeriod) Validate() error {
	return dara.Validate(s)
}

type GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList struct {
	// Default value.
	//
	// example:
	//
	// v1
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// Description.
	//
	// example:
	//
	// xxtest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Parameter key.
	//
	// example:
	//
	// v1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList) String() string {
	return dara.Prettify(s)
}

func (s GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList) GoString() string {
	return s.String()
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList) GetDescription() *string {
	return s.Description
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList) GetKey() *string {
	return s.Key
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList) SetDefaultValue(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList {
	s.DefaultValue = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList) SetDescription(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList {
	s.Description = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList) SetKey(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList {
	s.Key = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOutputContextParamList) Validate() error {
	return dara.Validate(s)
}

type GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList struct {
	// Node ID.
	//
	// example:
	//
	// 11123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Node name.
	//
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList) String() string {
	return dara.Prettify(s)
}

func (s GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList) GoString() string {
	return s.String()
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList) GetId() *string {
	return s.Id
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList) GetName() *string {
	return s.Name
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList) SetId(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList {
	s.Id = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList) SetName(v string) *GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList {
	s.Name = &v
	return s
}

func (s *GetDevObjectDependencyResponseBodyDevObjectDependencyListOwnerList) Validate() error {
	return dara.Validate(s)
}
