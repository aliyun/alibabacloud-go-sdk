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
	PageResult *ListSecurityIdentifyRecordsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	IdentifyRecordList []*ListSecurityIdentifyRecordsResponseBodyPageResultIdentifyRecordList `json:"IdentifyRecordList,omitempty" xml:"IdentifyRecordList,omitempty" type:"Repeated"`
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
	// example:
	//
	// 0.1
	ActualIdentifyRate *float64 `json:"ActualIdentifyRate,omitempty" xml:"ActualIdentifyRate,omitempty"`
	// example:
	//
	// 2025-06-30
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// example:
	//
	// test_abc
	BizUnitDisplayName *string `json:"BizUnitDisplayName,omitempty" xml:"BizUnitDisplayName,omitempty"`
	// example:
	//
	// 11
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// example:
	//
	// test
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// example:
	//
	// test
	ClassifyAbbreviation *string `json:"ClassifyAbbreviation,omitempty" xml:"ClassifyAbbreviation,omitempty"`
	// example:
	//
	// 1
	ClassifyId *int64 `json:"ClassifyId,omitempty" xml:"ClassifyId,omitempty"`
	// example:
	//
	// test
	ClassifyName *string `json:"ClassifyName,omitempty" xml:"ClassifyName,omitempty"`
	// example:
	//
	// ENABLE
	ClassifyStatus *string `json:"ClassifyStatus,omitempty" xml:"ClassifyStatus,omitempty"`
	// example:
	//
	// 1
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// test
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
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// example:
	//
	// 22
	Id               *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	IsBetterMatch    *bool  `json:"IsBetterMatch,omitempty" xml:"IsBetterMatch,omitempty"`
	IsCustomIdentify *bool  `json:"IsCustomIdentify,omitempty" xml:"IsCustomIdentify,omitempty"`
	IsLocked         *bool  `json:"IsLocked,omitempty" xml:"IsLocked,omitempty"`
	// example:
	//
	// test
	LevelAbbreviation *string `json:"LevelAbbreviation,omitempty" xml:"LevelAbbreviation,omitempty"`
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
	// test_abc
	ProjectDisplayName *string `json:"ProjectDisplayName,omitempty" xml:"ProjectDisplayName,omitempty"`
	// example:
	//
	// 1001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
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
