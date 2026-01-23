// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTableColumnsResponseBody
	GetCode() *string
	SetColumnList(v []*GetTableColumnsResponseBodyColumnList) *GetTableColumnsResponseBody
	GetColumnList() []*GetTableColumnsResponseBodyColumnList
	SetHttpStatusCode(v int32) *GetTableColumnsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTableColumnsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTableColumnsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableColumnsResponseBody
	GetSuccess() *bool
}

type GetTableColumnsResponseBody struct {
	// example:
	//
	// OK
	Code       *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	ColumnList []*GetTableColumnsResponseBodyColumnList `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Repeated"`
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTableColumnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableColumnsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTableColumnsResponseBody) GetColumnList() []*GetTableColumnsResponseBodyColumnList {
	return s.ColumnList
}

func (s *GetTableColumnsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTableColumnsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTableColumnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableColumnsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableColumnsResponseBody) SetCode(v string) *GetTableColumnsResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableColumnsResponseBody) SetColumnList(v []*GetTableColumnsResponseBodyColumnList) *GetTableColumnsResponseBody {
	s.ColumnList = v
	return s
}

func (s *GetTableColumnsResponseBody) SetHttpStatusCode(v int32) *GetTableColumnsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTableColumnsResponseBody) SetMessage(v string) *GetTableColumnsResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableColumnsResponseBody) SetRequestId(v string) *GetTableColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableColumnsResponseBody) SetSuccess(v bool) *GetTableColumnsResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableColumnsResponseBody) Validate() error {
	if s.ColumnList != nil {
		for _, item := range s.ColumnList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTableColumnsResponseBodyColumnList struct {
	// example:
	//
	// 12345
	ClassifyId *int64 `json:"ClassifyId,omitempty" xml:"ClassifyId,omitempty"`
	// example:
	//
	// test
	ClassifyName *string `json:"ClassifyName,omitempty" xml:"ClassifyName,omitempty"`
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 30012011
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 3301
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// example:
	//
	// varchar
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// test
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// 年龄
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// dev
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 1121.col1
	Guid              *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	IsForeignKey      *bool   `json:"IsForeignKey,omitempty" xml:"IsForeignKey,omitempty"`
	IsPartitionColumn *bool   `json:"IsPartitionColumn,omitempty" xml:"IsPartitionColumn,omitempty"`
	IsPrimaryKey      *bool   `json:"IsPrimaryKey,omitempty" xml:"IsPrimaryKey,omitempty"`
	// example:
	//
	// 30012011
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// L1
	LevelAbbreviation *string `json:"LevelAbbreviation,omitempty" xml:"LevelAbbreviation,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// age
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Nullable *bool   `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	// example:
	//
	// 1
	SeqNumber *int32 `json:"SeqNumber,omitempty" xml:"SeqNumber,omitempty"`
	// example:
	//
	// 1121
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// example:
	//
	// t_test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 99
	VisitCount30d *int64 `json:"VisitCount30d,omitempty" xml:"VisitCount30d,omitempty"`
}

func (s GetTableColumnsResponseBodyColumnList) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnsResponseBodyColumnList) GoString() string {
	return s.String()
}

func (s *GetTableColumnsResponseBodyColumnList) GetClassifyId() *int64 {
	return s.ClassifyId
}

func (s *GetTableColumnsResponseBodyColumnList) GetClassifyName() *string {
	return s.ClassifyName
}

func (s *GetTableColumnsResponseBodyColumnList) GetComment() *string {
	return s.Comment
}

func (s *GetTableColumnsResponseBodyColumnList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTableColumnsResponseBodyColumnList) GetCreator() *string {
	return s.Creator
}

func (s *GetTableColumnsResponseBodyColumnList) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *GetTableColumnsResponseBodyColumnList) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *GetTableColumnsResponseBodyColumnList) GetDataType() *string {
	return s.DataType
}

func (s *GetTableColumnsResponseBodyColumnList) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetTableColumnsResponseBodyColumnList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetTableColumnsResponseBodyColumnList) GetEnv() *string {
	return s.Env
}

func (s *GetTableColumnsResponseBodyColumnList) GetGuid() *string {
	return s.Guid
}

func (s *GetTableColumnsResponseBodyColumnList) GetIsForeignKey() *bool {
	return s.IsForeignKey
}

func (s *GetTableColumnsResponseBodyColumnList) GetIsPartitionColumn() *bool {
	return s.IsPartitionColumn
}

func (s *GetTableColumnsResponseBodyColumnList) GetIsPrimaryKey() *bool {
	return s.IsPrimaryKey
}

func (s *GetTableColumnsResponseBodyColumnList) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetTableColumnsResponseBodyColumnList) GetLevelAbbreviation() *string {
	return s.LevelAbbreviation
}

func (s *GetTableColumnsResponseBodyColumnList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetTableColumnsResponseBodyColumnList) GetName() *string {
	return s.Name
}

func (s *GetTableColumnsResponseBodyColumnList) GetNullable() *bool {
	return s.Nullable
}

func (s *GetTableColumnsResponseBodyColumnList) GetSeqNumber() *int32 {
	return s.SeqNumber
}

func (s *GetTableColumnsResponseBodyColumnList) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetTableColumnsResponseBodyColumnList) GetTableName() *string {
	return s.TableName
}

func (s *GetTableColumnsResponseBodyColumnList) GetVisitCount30d() *int64 {
	return s.VisitCount30d
}

func (s *GetTableColumnsResponseBodyColumnList) SetClassifyId(v int64) *GetTableColumnsResponseBodyColumnList {
	s.ClassifyId = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetClassifyName(v string) *GetTableColumnsResponseBodyColumnList {
	s.ClassifyName = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetComment(v string) *GetTableColumnsResponseBodyColumnList {
	s.Comment = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetCreateTime(v string) *GetTableColumnsResponseBodyColumnList {
	s.CreateTime = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetCreator(v string) *GetTableColumnsResponseBodyColumnList {
	s.Creator = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetDataSourceId(v int64) *GetTableColumnsResponseBodyColumnList {
	s.DataSourceId = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetDataSourceType(v string) *GetTableColumnsResponseBodyColumnList {
	s.DataSourceType = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetDataType(v string) *GetTableColumnsResponseBodyColumnList {
	s.DataType = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetDefaultValue(v string) *GetTableColumnsResponseBodyColumnList {
	s.DefaultValue = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetDisplayName(v string) *GetTableColumnsResponseBodyColumnList {
	s.DisplayName = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetEnv(v string) *GetTableColumnsResponseBodyColumnList {
	s.Env = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetGuid(v string) *GetTableColumnsResponseBodyColumnList {
	s.Guid = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetIsForeignKey(v bool) *GetTableColumnsResponseBodyColumnList {
	s.IsForeignKey = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetIsPartitionColumn(v bool) *GetTableColumnsResponseBodyColumnList {
	s.IsPartitionColumn = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetIsPrimaryKey(v bool) *GetTableColumnsResponseBodyColumnList {
	s.IsPrimaryKey = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetLastModifier(v string) *GetTableColumnsResponseBodyColumnList {
	s.LastModifier = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetLevelAbbreviation(v string) *GetTableColumnsResponseBodyColumnList {
	s.LevelAbbreviation = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetModifyTime(v string) *GetTableColumnsResponseBodyColumnList {
	s.ModifyTime = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetName(v string) *GetTableColumnsResponseBodyColumnList {
	s.Name = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetNullable(v bool) *GetTableColumnsResponseBodyColumnList {
	s.Nullable = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetSeqNumber(v int32) *GetTableColumnsResponseBodyColumnList {
	s.SeqNumber = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetTableGuid(v string) *GetTableColumnsResponseBodyColumnList {
	s.TableGuid = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetTableName(v string) *GetTableColumnsResponseBodyColumnList {
	s.TableName = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) SetVisitCount30d(v int64) *GetTableColumnsResponseBodyColumnList {
	s.VisitCount30d = &v
	return s
}

func (s *GetTableColumnsResponseBodyColumnList) Validate() error {
	return dara.Validate(s)
}
