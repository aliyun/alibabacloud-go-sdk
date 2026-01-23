// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecurityIdentifyResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSecurityIdentifyResultResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetSecurityIdentifyResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSecurityIdentifyResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSecurityIdentifyResultResponseBody
	GetRequestId() *string
	SetSecurityIdentifyResultInfo(v *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) *GetSecurityIdentifyResultResponseBody
	GetSecurityIdentifyResultInfo() *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo
	SetSuccess(v bool) *GetSecurityIdentifyResultResponseBody
	GetSuccess() *bool
}

type GetSecurityIdentifyResultResponseBody struct {
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
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId                  *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityIdentifyResultInfo *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo `json:"SecurityIdentifyResultInfo,omitempty" xml:"SecurityIdentifyResultInfo,omitempty" type:"Struct"`
	Success                    *bool                                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSecurityIdentifyResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityIdentifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecurityIdentifyResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSecurityIdentifyResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSecurityIdentifyResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSecurityIdentifyResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecurityIdentifyResultResponseBody) GetSecurityIdentifyResultInfo() *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	return s.SecurityIdentifyResultInfo
}

func (s *GetSecurityIdentifyResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSecurityIdentifyResultResponseBody) SetCode(v string) *GetSecurityIdentifyResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBody) SetHttpStatusCode(v int32) *GetSecurityIdentifyResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBody) SetMessage(v string) *GetSecurityIdentifyResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBody) SetRequestId(v string) *GetSecurityIdentifyResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBody) SetSecurityIdentifyResultInfo(v *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) *GetSecurityIdentifyResultResponseBody {
	s.SecurityIdentifyResultInfo = v
	return s
}

func (s *GetSecurityIdentifyResultResponseBody) SetSuccess(v bool) *GetSecurityIdentifyResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBody) Validate() error {
	if s.SecurityIdentifyResultInfo != nil {
		if err := s.SecurityIdentifyResultInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo struct {
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

func (s GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GoString() string {
	return s.String()
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetBizDate() *string {
	return s.BizDate
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetBizUnitDisplayName() *string {
	return s.BizUnitDisplayName
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetClassifyAbbreviation() *string {
	return s.ClassifyAbbreviation
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetClassifyId() *int64 {
	return s.ClassifyId
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetClassifyName() *string {
	return s.ClassifyName
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetFieldDescription() *string {
	return s.FieldDescription
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetFieldId() *string {
	return s.FieldId
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetFieldName() *string {
	return s.FieldName
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetHasBetterRule() *bool {
	return s.HasBetterRule
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetId() *int64 {
	return s.Id
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetIdentifyRecordId() *int64 {
	return s.IdentifyRecordId
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetIsCustomIdentify() *bool {
	return s.IsCustomIdentify
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetIsLocked() *bool {
	return s.IsLocked
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetLevelAbbreviation() *string {
	return s.LevelAbbreviation
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetLevelColor() *int64 {
	return s.LevelColor
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetLevelIndex() *int64 {
	return s.LevelIndex
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetLevelName() *string {
	return s.LevelName
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetModifier() *string {
	return s.Modifier
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetProjectDisplayName() *string {
	return s.ProjectDisplayName
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetScanTaskId() *int64 {
	return s.ScanTaskId
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetStatus() *string {
	return s.Status
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetTableCatalog() *string {
	return s.TableCatalog
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetTableDescription() *string {
	return s.TableDescription
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetTableEnv() *string {
	return s.TableEnv
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetTableId() *string {
	return s.TableId
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetTableName() *string {
	return s.TableName
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetTableTaskId() *int64 {
	return s.TableTaskId
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) GetTableType() *string {
	return s.TableType
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetBizDate(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.BizDate = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetBizUnitDisplayName(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.BizUnitDisplayName = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetBizUnitId(v int64) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.BizUnitId = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetBizUnitName(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.BizUnitName = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetClassifyAbbreviation(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.ClassifyAbbreviation = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetClassifyId(v int64) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.ClassifyId = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetClassifyName(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.ClassifyName = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetCreateTime(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.CreateTime = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetCreator(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.Creator = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetDatasourceId(v int64) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.DatasourceId = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetDatasourceName(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.DatasourceName = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetFieldDescription(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.FieldDescription = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetFieldId(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.FieldId = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetFieldName(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.FieldName = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetHasBetterRule(v bool) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.HasBetterRule = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetId(v int64) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.Id = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetIdentifyRecordId(v int64) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.IdentifyRecordId = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetIsCustomIdentify(v bool) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.IsCustomIdentify = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetIsLocked(v bool) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.IsLocked = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetLevelAbbreviation(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.LevelAbbreviation = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetLevelColor(v int64) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.LevelColor = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetLevelIndex(v int64) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.LevelIndex = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetLevelName(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.LevelName = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetModifier(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.Modifier = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetModifyTime(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetProjectDisplayName(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.ProjectDisplayName = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetProjectId(v int64) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.ProjectId = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetProjectName(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.ProjectName = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetScanTaskId(v int64) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.ScanTaskId = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetStatus(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.Status = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetTableCatalog(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.TableCatalog = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetTableDescription(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.TableDescription = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetTableEnv(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.TableEnv = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetTableId(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.TableId = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetTableName(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.TableName = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetTableTaskId(v int64) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.TableTaskId = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) SetTableType(v string) *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo {
	s.TableType = &v
	return s
}

func (s *GetSecurityIdentifyResultResponseBodySecurityIdentifyResultInfo) Validate() error {
	return dara.Validate(s)
}
