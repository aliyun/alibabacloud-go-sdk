// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityIdentifyRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSecurityIdentifyRecordsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListSecurityIdentifyRecordsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListSecurityIdentifyRecordsResponseBody
	GetMessage() *string
	SetPageResult(v *ListSecurityIdentifyRecordsResponseBodyPageResult) *ListSecurityIdentifyRecordsResponseBody
	GetPageResult() *ListSecurityIdentifyRecordsResponseBodyPageResult
	SetRequestId(v string) *ListSecurityIdentifyRecordsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSecurityIdentifyRecordsResponseBody
	GetSuccess() *bool
}

type ListSecurityIdentifyRecordsResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The paging query result.
	PageResult *ListSecurityIdentifyRecordsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSecurityIdentifyRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityIdentifyRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityIdentifyRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSecurityIdentifyRecordsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSecurityIdentifyRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSecurityIdentifyRecordsResponseBody) GetPageResult() *ListSecurityIdentifyRecordsResponseBodyPageResult {
	return s.PageResult
}

func (s *ListSecurityIdentifyRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecurityIdentifyRecordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSecurityIdentifyRecordsResponseBody) SetCode(v string) *ListSecurityIdentifyRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBody) SetHttpStatusCode(v int32) *ListSecurityIdentifyRecordsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBody) SetMessage(v string) *ListSecurityIdentifyRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBody) SetPageResult(v *ListSecurityIdentifyRecordsResponseBodyPageResult) *ListSecurityIdentifyRecordsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBody) SetRequestId(v string) *ListSecurityIdentifyRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBody) SetSuccess(v bool) *ListSecurityIdentifyRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSecurityIdentifyRecordsResponseBodyPageResult struct {
	// The list of identification records.
	IdentifyRecordList []*ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList `json:"IdentifyRecordList,omitempty" xml:"IdentifyRecordList,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSecurityIdentifyRecordsResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityIdentifyRecordsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResult) GetIdentifyRecordList() []*ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	return s.IdentifyRecordList
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResult) SetIdentifyRecordList(v []*ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) *ListSecurityIdentifyRecordsResponseBodyPageResult {
	s.IdentifyRecordList = v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResult) SetTotalCount(v int32) *ListSecurityIdentifyRecordsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResult) Validate() error {
	if s.IdentifyRecordList != nil {
		for _, item := range s.IdentifyRecordList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList struct {
	// The actual match rate.
	//
	// example:
	//
	// 0.1
	ActualIdentifyRate *float64 `json:"ActualIdentifyRate,omitempty" xml:"ActualIdentifyRate,omitempty"`
	// The business date.
	//
	// example:
	//
	// 2025-06-30
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The display name of the business unit.
	//
	// example:
	//
	// test_abc
	BizUnitDisplayName *string `json:"BizUnitDisplayName,omitempty" xml:"BizUnitDisplayName,omitempty"`
	// The business unit ID.
	//
	// example:
	//
	// 11
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The business unit name.
	//
	// example:
	//
	// test
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// The classification abbreviation.
	//
	// example:
	//
	// test
	ClassifyAbbreviation *string `json:"ClassifyAbbreviation,omitempty" xml:"ClassifyAbbreviation,omitempty"`
	// The classification ID.
	//
	// example:
	//
	// 1
	ClassifyId *int64 `json:"ClassifyId,omitempty" xml:"ClassifyId,omitempty"`
	// The classification name.
	//
	// example:
	//
	// test
	ClassifyName *string `json:"ClassifyName,omitempty" xml:"ClassifyName,omitempty"`
	// The classification effective status. Valid values:
	//
	// - ENABLE: enabled.
	//
	// - DISABLE: disabled.
	//
	// example:
	//
	// ENABLE
	ClassifyStatus *string `json:"ClassifyStatus,omitempty" xml:"ClassifyStatus,omitempty"`
	// The datasource ID.
	//
	// example:
	//
	// 1
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// The datasource name.
	//
	// example:
	//
	// test
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// The field description.
	//
	// example:
	//
	// test
	FieldDescription *string `json:"FieldDescription,omitempty" xml:"FieldDescription,omitempty"`
	// The field ID.
	//
	// example:
	//
	// t_test.col1
	FieldId *string `json:"FieldId,omitempty" xml:"FieldId,omitempty"`
	// The field name.
	//
	// example:
	//
	// col1
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The identification record ID.
	//
	// example:
	//
	// 22
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether this is a better match.
	IsBetterMatch *bool `json:"IsBetterMatch,omitempty" xml:"IsBetterMatch,omitempty"`
	// Indicates whether a custom tagging rule is used.
	IsCustomIdentify *bool `json:"IsCustomIdentify,omitempty" xml:"IsCustomIdentify,omitempty"`
	// Indicates whether the record is locked.
	IsLocked *bool `json:"IsLocked,omitempty" xml:"IsLocked,omitempty"`
	// The classification level abbreviation.
	//
	// example:
	//
	// test
	LevelAbbreviation *string `json:"LevelAbbreviation,omitempty" xml:"LevelAbbreviation,omitempty"`
	// The classification level index.
	//
	// example:
	//
	// 1
	LevelIndex *int64 `json:"LevelIndex,omitempty" xml:"LevelIndex,omitempty"`
	// The classification level name.
	//
	// example:
	//
	// test
	LevelName *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	// The display name of the project.
	//
	// example:
	//
	// test_abc
	ProjectDisplayName *string `json:"ProjectDisplayName,omitempty" xml:"ProjectDisplayName,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 1001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The project name.
	//
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The table description.
	//
	// example:
	//
	// test
	TableDescription *string `json:"TableDescription,omitempty" xml:"TableDescription,omitempty"`
	// The table environment.
	//
	// example:
	//
	// DEV
	TableEnv *string `json:"TableEnv,omitempty" xml:"TableEnv,omitempty"`
	// The table ID.
	//
	// example:
	//
	// 1
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The table name.
	//
	// example:
	//
	// t_test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The table type. Valid values:
	//
	// - LOGIC_TABLE: logical table.
	//
	// - LOGIC_DIM_TABLE: logical dimension table.
	//
	// - LOGIC_FACT_TABLE: logical fact table.
	//
	// - LOGIC_SUM_TABLE: logical aggregate table.
	//
	// - LOGIC_ODM_SOURCE: source table.
	//
	// - LOGIC_LABEL_TABLE: logical label table.
	//
	// - PHYSICAL_TABLE: physical table.
	//
	// - REAL_TIME_LOGIC_TABLE: real-time configured logical table.
	//
	// example:
	//
	// LOGIC_TABLE
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GoString() string {
	return s.String()
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetActualIdentifyRate() *float64 {
	return s.ActualIdentifyRate
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetBizDate() *string {
	return s.BizDate
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetBizUnitDisplayName() *string {
	return s.BizUnitDisplayName
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetClassifyAbbreviation() *string {
	return s.ClassifyAbbreviation
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetClassifyId() *int64 {
	return s.ClassifyId
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetClassifyName() *string {
	return s.ClassifyName
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetClassifyStatus() *string {
	return s.ClassifyStatus
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetFieldDescription() *string {
	return s.FieldDescription
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetFieldId() *string {
	return s.FieldId
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetFieldName() *string {
	return s.FieldName
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetId() *int64 {
	return s.Id
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetIsBetterMatch() *bool {
	return s.IsBetterMatch
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetIsCustomIdentify() *bool {
	return s.IsCustomIdentify
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetIsLocked() *bool {
	return s.IsLocked
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetLevelAbbreviation() *string {
	return s.LevelAbbreviation
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetLevelIndex() *int64 {
	return s.LevelIndex
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetLevelName() *string {
	return s.LevelName
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetProjectDisplayName() *string {
	return s.ProjectDisplayName
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetTableDescription() *string {
	return s.TableDescription
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetTableEnv() *string {
	return s.TableEnv
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetTableId() *string {
	return s.TableId
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetTableName() *string {
	return s.TableName
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) GetTableType() *string {
	return s.TableType
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetActualIdentifyRate(v float64) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.ActualIdentifyRate = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetBizDate(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.BizDate = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetBizUnitDisplayName(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.BizUnitDisplayName = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetBizUnitId(v int64) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.BizUnitId = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetBizUnitName(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.BizUnitName = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetClassifyAbbreviation(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.ClassifyAbbreviation = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetClassifyId(v int64) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.ClassifyId = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetClassifyName(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.ClassifyName = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetClassifyStatus(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.ClassifyStatus = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetDatasourceId(v int64) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.DatasourceId = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetDatasourceName(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.DatasourceName = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetFieldDescription(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.FieldDescription = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetFieldId(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.FieldId = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetFieldName(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.FieldName = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetId(v int64) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.Id = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetIsBetterMatch(v bool) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.IsBetterMatch = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetIsCustomIdentify(v bool) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.IsCustomIdentify = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetIsLocked(v bool) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.IsLocked = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetLevelAbbreviation(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.LevelAbbreviation = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetLevelIndex(v int64) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.LevelIndex = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetLevelName(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.LevelName = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetProjectDisplayName(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.ProjectDisplayName = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetProjectId(v int64) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.ProjectId = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetProjectName(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.ProjectName = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetTableDescription(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.TableDescription = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetTableEnv(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.TableEnv = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetTableId(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.TableId = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetTableName(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.TableName = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) SetTableType(v string) *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList {
	s.TableType = &v
	return s
}

func (s *ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList) Validate() error {
	return dara.Validate(s)
}
