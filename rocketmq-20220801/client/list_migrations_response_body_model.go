// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMigrationsResponseBody
	GetCode() *string
	SetData(v *ListMigrationsResponseBodyData) *ListMigrationsResponseBody
	GetData() *ListMigrationsResponseBodyData
	SetDynamicCode(v string) *ListMigrationsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListMigrationsResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListMigrationsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListMigrationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMigrationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMigrationsResponseBody
	GetSuccess() *bool
}

type ListMigrationsResponseBody struct {
	// example:
	//
	// MissingInstanceId
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListMigrationsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// E5897B2E-C3AC-56DC-A482-F0E9E53F48D5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListMigrationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMigrationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMigrationsResponseBody) GetData() *ListMigrationsResponseBodyData {
	return s.Data
}

func (s *ListMigrationsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListMigrationsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListMigrationsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMigrationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMigrationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMigrationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMigrationsResponseBody) SetCode(v string) *ListMigrationsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMigrationsResponseBody) SetData(v *ListMigrationsResponseBodyData) *ListMigrationsResponseBody {
	s.Data = v
	return s
}

func (s *ListMigrationsResponseBody) SetDynamicCode(v string) *ListMigrationsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListMigrationsResponseBody) SetDynamicMessage(v string) *ListMigrationsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListMigrationsResponseBody) SetHttpStatusCode(v int32) *ListMigrationsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMigrationsResponseBody) SetMessage(v string) *ListMigrationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMigrationsResponseBody) SetRequestId(v string) *ListMigrationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMigrationsResponseBody) SetSuccess(v bool) *ListMigrationsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMigrationsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMigrationsResponseBodyData struct {
	List []*ListMigrationsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 3
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMigrationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMigrationsResponseBodyData) GetList() []*ListMigrationsResponseBodyDataList {
	return s.List
}

func (s *ListMigrationsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListMigrationsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListMigrationsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMigrationsResponseBodyData) SetList(v []*ListMigrationsResponseBodyDataList) *ListMigrationsResponseBodyData {
	s.List = v
	return s
}

func (s *ListMigrationsResponseBodyData) SetPageNumber(v int64) *ListMigrationsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListMigrationsResponseBodyData) SetPageSize(v int64) *ListMigrationsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMigrationsResponseBodyData) SetTotalCount(v int64) *ListMigrationsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListMigrationsResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMigrationsResponseBodyDataList struct {
	// example:
	//
	// 2022-08-01 00:00:00
	CreateTime   *string                                         `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CurrentStage *ListMigrationsResponseBodyDataListCurrentStage `json:"currentStage,omitempty" xml:"currentStage,omitempty" type:"Struct"`
	// example:
	//
	// 21
	MigrationId *int64 `json:"migrationId,omitempty" xml:"migrationId,omitempty"`
	// example:
	//
	// xxx
	MigrationName   *string                                            `json:"migrationName,omitempty" xml:"migrationName,omitempty"`
	MigrationSource *ListMigrationsResponseBodyDataListMigrationSource `json:"migrationSource,omitempty" xml:"migrationSource,omitempty" type:"Struct"`
	// example:
	//
	// MIGRATING
	MigrationStatus *string                                            `json:"migrationStatus,omitempty" xml:"migrationStatus,omitempty"`
	MigrationTarget *ListMigrationsResponseBodyDataListMigrationTarget `json:"migrationTarget,omitempty" xml:"migrationTarget,omitempty" type:"Struct"`
	// example:
	//
	// INSTANCE_MIGRATION
	MigrationType *string `json:"migrationType,omitempty" xml:"migrationType,omitempty"`
	// example:
	//
	// 2022-08-01 20:05:50
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// 111
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListMigrationsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListMigrationsResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMigrationsResponseBodyDataList) GetCurrentStage() *ListMigrationsResponseBodyDataListCurrentStage {
	return s.CurrentStage
}

func (s *ListMigrationsResponseBodyDataList) GetMigrationId() *int64 {
	return s.MigrationId
}

func (s *ListMigrationsResponseBodyDataList) GetMigrationName() *string {
	return s.MigrationName
}

func (s *ListMigrationsResponseBodyDataList) GetMigrationSource() *ListMigrationsResponseBodyDataListMigrationSource {
	return s.MigrationSource
}

func (s *ListMigrationsResponseBodyDataList) GetMigrationStatus() *string {
	return s.MigrationStatus
}

func (s *ListMigrationsResponseBodyDataList) GetMigrationTarget() *ListMigrationsResponseBodyDataListMigrationTarget {
	return s.MigrationTarget
}

func (s *ListMigrationsResponseBodyDataList) GetMigrationType() *string {
	return s.MigrationType
}

func (s *ListMigrationsResponseBodyDataList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListMigrationsResponseBodyDataList) GetUserId() *string {
	return s.UserId
}

func (s *ListMigrationsResponseBodyDataList) SetCreateTime(v string) *ListMigrationsResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListMigrationsResponseBodyDataList) SetCurrentStage(v *ListMigrationsResponseBodyDataListCurrentStage) *ListMigrationsResponseBodyDataList {
	s.CurrentStage = v
	return s
}

func (s *ListMigrationsResponseBodyDataList) SetMigrationId(v int64) *ListMigrationsResponseBodyDataList {
	s.MigrationId = &v
	return s
}

func (s *ListMigrationsResponseBodyDataList) SetMigrationName(v string) *ListMigrationsResponseBodyDataList {
	s.MigrationName = &v
	return s
}

func (s *ListMigrationsResponseBodyDataList) SetMigrationSource(v *ListMigrationsResponseBodyDataListMigrationSource) *ListMigrationsResponseBodyDataList {
	s.MigrationSource = v
	return s
}

func (s *ListMigrationsResponseBodyDataList) SetMigrationStatus(v string) *ListMigrationsResponseBodyDataList {
	s.MigrationStatus = &v
	return s
}

func (s *ListMigrationsResponseBodyDataList) SetMigrationTarget(v *ListMigrationsResponseBodyDataListMigrationTarget) *ListMigrationsResponseBodyDataList {
	s.MigrationTarget = v
	return s
}

func (s *ListMigrationsResponseBodyDataList) SetMigrationType(v string) *ListMigrationsResponseBodyDataList {
	s.MigrationType = &v
	return s
}

func (s *ListMigrationsResponseBodyDataList) SetUpdateTime(v string) *ListMigrationsResponseBodyDataList {
	s.UpdateTime = &v
	return s
}

func (s *ListMigrationsResponseBodyDataList) SetUserId(v string) *ListMigrationsResponseBodyDataList {
	s.UserId = &v
	return s
}

func (s *ListMigrationsResponseBodyDataList) Validate() error {
	if s.CurrentStage != nil {
		if err := s.CurrentStage.Validate(); err != nil {
			return err
		}
	}
	if s.MigrationSource != nil {
		if err := s.MigrationSource.Validate(); err != nil {
			return err
		}
	}
	if s.MigrationTarget != nil {
		if err := s.MigrationTarget.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMigrationsResponseBodyDataListCurrentStage struct {
	// example:
	//
	// []
	StageData interface{} `json:"stageData,omitempty" xml:"stageData,omitempty"`
	// example:
	//
	// DOING
	StageStatus *string `json:"stageStatus,omitempty" xml:"stageStatus,omitempty"`
	// example:
	//
	// MIGRATE_METADATA
	StageType *string `json:"stageType,omitempty" xml:"stageType,omitempty"`
}

func (s ListMigrationsResponseBodyDataListCurrentStage) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationsResponseBodyDataListCurrentStage) GoString() string {
	return s.String()
}

func (s *ListMigrationsResponseBodyDataListCurrentStage) GetStageData() interface{} {
	return s.StageData
}

func (s *ListMigrationsResponseBodyDataListCurrentStage) GetStageStatus() *string {
	return s.StageStatus
}

func (s *ListMigrationsResponseBodyDataListCurrentStage) GetStageType() *string {
	return s.StageType
}

func (s *ListMigrationsResponseBodyDataListCurrentStage) SetStageData(v interface{}) *ListMigrationsResponseBodyDataListCurrentStage {
	s.StageData = v
	return s
}

func (s *ListMigrationsResponseBodyDataListCurrentStage) SetStageStatus(v string) *ListMigrationsResponseBodyDataListCurrentStage {
	s.StageStatus = &v
	return s
}

func (s *ListMigrationsResponseBodyDataListCurrentStage) SetStageType(v string) *ListMigrationsResponseBodyDataListCurrentStage {
	s.StageType = &v
	return s
}

func (s *ListMigrationsResponseBodyDataListCurrentStage) Validate() error {
	return dara.Validate(s)
}

type ListMigrationsResponseBodyDataListMigrationSource struct {
	// example:
	//
	// []
	SourceData interface{} `json:"sourceData,omitempty" xml:"sourceData,omitempty"`
	// example:
	//
	// EXTERNAL_INSTANCE
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
}

func (s ListMigrationsResponseBodyDataListMigrationSource) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationsResponseBodyDataListMigrationSource) GoString() string {
	return s.String()
}

func (s *ListMigrationsResponseBodyDataListMigrationSource) GetSourceData() interface{} {
	return s.SourceData
}

func (s *ListMigrationsResponseBodyDataListMigrationSource) GetSourceType() *string {
	return s.SourceType
}

func (s *ListMigrationsResponseBodyDataListMigrationSource) SetSourceData(v interface{}) *ListMigrationsResponseBodyDataListMigrationSource {
	s.SourceData = v
	return s
}

func (s *ListMigrationsResponseBodyDataListMigrationSource) SetSourceType(v string) *ListMigrationsResponseBodyDataListMigrationSource {
	s.SourceType = &v
	return s
}

func (s *ListMigrationsResponseBodyDataListMigrationSource) Validate() error {
	return dara.Validate(s)
}

type ListMigrationsResponseBodyDataListMigrationTarget struct {
	// example:
	//
	// []
	TargetData interface{} `json:"targetData,omitempty" xml:"targetData,omitempty"`
	// example:
	//
	// INTERNAL_INSTANCE
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s ListMigrationsResponseBodyDataListMigrationTarget) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationsResponseBodyDataListMigrationTarget) GoString() string {
	return s.String()
}

func (s *ListMigrationsResponseBodyDataListMigrationTarget) GetTargetData() interface{} {
	return s.TargetData
}

func (s *ListMigrationsResponseBodyDataListMigrationTarget) GetTargetType() *string {
	return s.TargetType
}

func (s *ListMigrationsResponseBodyDataListMigrationTarget) SetTargetData(v interface{}) *ListMigrationsResponseBodyDataListMigrationTarget {
	s.TargetData = v
	return s
}

func (s *ListMigrationsResponseBodyDataListMigrationTarget) SetTargetType(v string) *ListMigrationsResponseBodyDataListMigrationTarget {
	s.TargetType = &v
	return s
}

func (s *ListMigrationsResponseBodyDataListMigrationTarget) Validate() error {
	return dara.Validate(s)
}
