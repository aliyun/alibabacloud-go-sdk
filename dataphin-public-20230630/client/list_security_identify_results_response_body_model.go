// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityIdentifyResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSecurityIdentifyResultsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListSecurityIdentifyResultsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListSecurityIdentifyResultsResponseBody
	GetMessage() *string
	SetPageResult(v *ListSecurityIdentifyResultsResponseBodyPageResult) *ListSecurityIdentifyResultsResponseBody
	GetPageResult() *ListSecurityIdentifyResultsResponseBodyPageResult
	SetRequestId(v string) *ListSecurityIdentifyResultsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSecurityIdentifyResultsResponseBody
	GetSuccess() *bool
}

type ListSecurityIdentifyResultsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message    *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListSecurityIdentifyResultsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSecurityIdentifyResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityIdentifyResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityIdentifyResultsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSecurityIdentifyResultsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSecurityIdentifyResultsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSecurityIdentifyResultsResponseBody) GetPageResult() *ListSecurityIdentifyResultsResponseBodyPageResult {
	return s.PageResult
}

func (s *ListSecurityIdentifyResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecurityIdentifyResultsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSecurityIdentifyResultsResponseBody) SetCode(v string) *ListSecurityIdentifyResultsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBody) SetHttpStatusCode(v int32) *ListSecurityIdentifyResultsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBody) SetMessage(v string) *ListSecurityIdentifyResultsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBody) SetPageResult(v *ListSecurityIdentifyResultsResponseBodyPageResult) *ListSecurityIdentifyResultsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBody) SetRequestId(v string) *ListSecurityIdentifyResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBody) SetSuccess(v bool) *ListSecurityIdentifyResultsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSecurityIdentifyResultsResponseBodyPageResult struct {
	SecurityIdentifyResultList []*ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList `json:"SecurityIdentifyResultList,omitempty" xml:"SecurityIdentifyResultList,omitempty" type:"Repeated"`
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSecurityIdentifyResultsResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityIdentifyResultsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResult) GetSecurityIdentifyResultList() []*ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	return s.SecurityIdentifyResultList
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResult) SetSecurityIdentifyResultList(v []*ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) *ListSecurityIdentifyResultsResponseBodyPageResult {
	s.SecurityIdentifyResultList = v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResult) SetTotalCount(v int32) *ListSecurityIdentifyResultsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResult) Validate() error {
	if s.SecurityIdentifyResultList != nil {
		for _, item := range s.SecurityIdentifyResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList struct {
	// example:
	//
	// 2025-06-30
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// example:
	//
	// LD_test
	BizUnitDisplayName *string `json:"BizUnitDisplayName,omitempty" xml:"BizUnitDisplayName,omitempty"`
	// example:
	//
	// 1
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// example:
	//
	// LD_test
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// example:
	//
	// test
	ClassifyAbbreviation *string `json:"ClassifyAbbreviation,omitempty" xml:"ClassifyAbbreviation,omitempty"`
	// example:
	//
	// 33
	ClassifyId *int64 `json:"ClassifyId,omitempty" xml:"ClassifyId,omitempty"`
	// example:
	//
	// test
	ClassifyName *string `json:"ClassifyName,omitempty" xml:"ClassifyName,omitempty"`
	// example:
	//
	// 2025-06-30 10:30:30
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 30012011
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 101121
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// ds1
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// example:
	//
	// test
	FieldDescription *string `json:"FieldDescription,omitempty" xml:"FieldDescription,omitempty"`
	// example:
	//
	// t_test.col1
	FieldId *string `json:"FieldId,omitempty" xml:"FieldId,omitempty"`
	// example:
	//
	// col1
	FieldName     *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	HasBetterRule *bool   `json:"HasBetterRule,omitempty" xml:"HasBetterRule,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 99
	IdentifyRecordId *int64 `json:"IdentifyRecordId,omitempty" xml:"IdentifyRecordId,omitempty"`
	IsCustomIdentify *bool  `json:"IsCustomIdentify,omitempty" xml:"IsCustomIdentify,omitempty"`
	IsLocked         *bool  `json:"IsLocked,omitempty" xml:"IsLocked,omitempty"`
	// example:
	//
	// test
	LevelAbbreviation *string `json:"LevelAbbreviation,omitempty" xml:"LevelAbbreviation,omitempty"`
	// example:
	//
	// 1
	LevelColor *int64 `json:"LevelColor,omitempty" xml:"LevelColor,omitempty"`
	// example:
	//
	// 1
	LevelIndex *int64 `json:"LevelIndex,omitempty" xml:"LevelIndex,omitempty"`
	// example:
	//
	// test
	LevelName *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	// example:
	//
	// 30012011
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// example:
	//
	// 2025-06-30 10:30:30
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test
	ProjectDisplayName *string `json:"ProjectDisplayName,omitempty" xml:"ProjectDisplayName,omitempty"`
	// example:
	//
	// 101121
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 11
	ScanTaskId *int64 `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// testdb
	TableCatalog *string `json:"TableCatalog,omitempty" xml:"TableCatalog,omitempty"`
	// example:
	//
	// test
	TableDescription *string `json:"TableDescription,omitempty" xml:"TableDescription,omitempty"`
	// example:
	//
	// DEV
	TableEnv *string `json:"TableEnv,omitempty" xml:"TableEnv,omitempty"`
	// example:
	//
	// 1
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// example:
	//
	// t_test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 22
	TableTaskId *int64 `json:"TableTaskId,omitempty" xml:"TableTaskId,omitempty"`
	// example:
	//
	// LOGIC_TABLE
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GoString() string {
	return s.String()
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetBizDate() *string {
	return s.BizDate
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetBizUnitDisplayName() *string {
	return s.BizUnitDisplayName
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetClassifyAbbreviation() *string {
	return s.ClassifyAbbreviation
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetClassifyId() *int64 {
	return s.ClassifyId
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetClassifyName() *string {
	return s.ClassifyName
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetCreator() *string {
	return s.Creator
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetFieldDescription() *string {
	return s.FieldDescription
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetFieldId() *string {
	return s.FieldId
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetFieldName() *string {
	return s.FieldName
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetHasBetterRule() *bool {
	return s.HasBetterRule
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetId() *int64 {
	return s.Id
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetIdentifyRecordId() *int64 {
	return s.IdentifyRecordId
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetIsCustomIdentify() *bool {
	return s.IsCustomIdentify
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetIsLocked() *bool {
	return s.IsLocked
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetLevelAbbreviation() *string {
	return s.LevelAbbreviation
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetLevelColor() *int64 {
	return s.LevelColor
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetLevelIndex() *int64 {
	return s.LevelIndex
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetLevelName() *string {
	return s.LevelName
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetModifier() *string {
	return s.Modifier
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetProjectDisplayName() *string {
	return s.ProjectDisplayName
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetScanTaskId() *int64 {
	return s.ScanTaskId
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetStatus() *string {
	return s.Status
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetTableCatalog() *string {
	return s.TableCatalog
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetTableDescription() *string {
	return s.TableDescription
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetTableEnv() *string {
	return s.TableEnv
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetTableId() *string {
	return s.TableId
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetTableName() *string {
	return s.TableName
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetTableTaskId() *int64 {
	return s.TableTaskId
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) GetTableType() *string {
	return s.TableType
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetBizDate(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.BizDate = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetBizUnitDisplayName(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.BizUnitDisplayName = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetBizUnitId(v int64) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.BizUnitId = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetBizUnitName(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.BizUnitName = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetClassifyAbbreviation(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.ClassifyAbbreviation = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetClassifyId(v int64) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.ClassifyId = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetClassifyName(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.ClassifyName = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetCreateTime(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.CreateTime = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetCreator(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.Creator = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetDatasourceId(v int64) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.DatasourceId = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetDatasourceName(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.DatasourceName = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetFieldDescription(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.FieldDescription = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetFieldId(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.FieldId = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetFieldName(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.FieldName = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetHasBetterRule(v bool) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.HasBetterRule = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetId(v int64) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.Id = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetIdentifyRecordId(v int64) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.IdentifyRecordId = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetIsCustomIdentify(v bool) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.IsCustomIdentify = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetIsLocked(v bool) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.IsLocked = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetLevelAbbreviation(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.LevelAbbreviation = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetLevelColor(v int64) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.LevelColor = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetLevelIndex(v int64) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.LevelIndex = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetLevelName(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.LevelName = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetModifier(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.Modifier = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetModifyTime(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.ModifyTime = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetProjectDisplayName(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.ProjectDisplayName = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetProjectId(v int64) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.ProjectId = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetProjectName(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.ProjectName = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetScanTaskId(v int64) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.ScanTaskId = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetStatus(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.Status = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetTableCatalog(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.TableCatalog = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetTableDescription(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.TableDescription = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetTableEnv(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.TableEnv = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetTableId(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.TableId = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetTableName(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.TableName = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetTableTaskId(v int64) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.TableTaskId = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) SetTableType(v string) *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList {
	s.TableType = &v
	return s
}

func (s *ListSecurityIdentifyResultsResponseBodyPageResultSecurityIdentifyResultList) Validate() error {
	return dara.Validate(s)
}
