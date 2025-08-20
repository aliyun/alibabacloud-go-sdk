// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdbMySqlTableMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DescribeAdbMySqlTableMetaResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAdbMySqlTableMetaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAdbMySqlTableMetaResponseBody
	GetSuccess() *bool
	SetTableMeta(v *DescribeAdbMySqlTableMetaResponseBodyTableMeta) *DescribeAdbMySqlTableMetaResponseBody
	GetTableMeta() *DescribeAdbMySqlTableMetaResponseBodyTableMeta
}

type DescribeAdbMySqlTableMetaResponseBody struct {
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.****
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The queried table metadata.
	TableMeta *DescribeAdbMySqlTableMetaResponseBodyTableMeta `json:"TableMeta,omitempty" xml:"TableMeta,omitempty" type:"Struct"`
}

func (s DescribeAdbMySqlTableMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlTableMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlTableMetaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAdbMySqlTableMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdbMySqlTableMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAdbMySqlTableMetaResponseBody) GetTableMeta() *DescribeAdbMySqlTableMetaResponseBodyTableMeta {
	return s.TableMeta
}

func (s *DescribeAdbMySqlTableMetaResponseBody) SetMessage(v string) *DescribeAdbMySqlTableMetaResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBody) SetRequestId(v string) *DescribeAdbMySqlTableMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBody) SetSuccess(v bool) *DescribeAdbMySqlTableMetaResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBody) SetTableMeta(v *DescribeAdbMySqlTableMetaResponseBodyTableMeta) *DescribeAdbMySqlTableMetaResponseBody {
	s.TableMeta = v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAdbMySqlTableMetaResponseBodyTableMeta struct {
	// The time when the table was created.
	//
	// example:
	//
	// 2025-03-14 02:18:08.0
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The distribution key column.
	//
	// example:
	//
	// customer_id
	DistributeColumn *string `json:"DistributeColumn,omitempty" xml:"DistributeColumn,omitempty"`
	// The distribution type.
	//
	// example:
	//
	// hash
	DistributeType *string `json:"DistributeType,omitempty" xml:"DistributeType,omitempty"`
	// Indicates whether full-column indexes are used.
	//
	// example:
	//
	// false
	IsAllIndex *bool `json:"IsAllIndex,omitempty" xml:"IsAllIndex,omitempty"`
	// Indicates whether dictionary encoding is used.
	//
	// example:
	//
	// true
	IsDictEncode *bool `json:"IsDictEncode,omitempty" xml:"IsDictEncode,omitempty"`
	// Indicates whether full-text indexes are used.
	//
	// example:
	//
	// true
	IsFullTextDict *bool `json:"IsFullTextDict,omitempty" xml:"IsFullTextDict,omitempty"`
	// Indicates whether pages are hidden.
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	IsHidden *bool `json:"IsHidden,omitempty" xml:"IsHidden,omitempty"`
	// The partition key column.
	//
	// example:
	//
	// DATE_FORMAT(login_time, \\"%Y%m%d\\")
	PartitionColumn *string `json:"PartitionColumn,omitempty" xml:"PartitionColumn,omitempty"`
	// The type of the partition.
	//
	// example:
	//
	// value
	PartitionType *string `json:"PartitionType,omitempty" xml:"PartitionType,omitempty"`
	// The primary key column.
	//
	// example:
	//
	// login_time,customer_id,phone_num
	PrimaryKeyColumn *string `json:"PrimaryKeyColumn,omitempty" xml:"PrimaryKeyColumn,omitempty"`
	// The table engine.
	//
	// example:
	//
	// XUANWU
	TableEngine *string `json:"TableEngine,omitempty" xml:"TableEngine,omitempty"`
	// The name of the table.
	//
	// **
	//
	// ****
	//
	// example:
	//
	// external_supplier
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The database to which the table belongs.
	//
	// example:
	//
	// tpch
	TableSchema *string `json:"TableSchema,omitempty" xml:"TableSchema,omitempty"`
	// The type of the table.
	//
	// example:
	//
	// fact_table
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
	// The time when the table was updated.
	//
	// example:
	//
	// 2024-07-25 02:07:23.0
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeAdbMySqlTableMetaResponseBodyTableMeta) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlTableMetaResponseBodyTableMeta) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) GetDistributeColumn() *string {
	return s.DistributeColumn
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) GetDistributeType() *string {
	return s.DistributeType
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) GetIsAllIndex() *bool {
	return s.IsAllIndex
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) GetIsDictEncode() *bool {
	return s.IsDictEncode
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) GetIsFullTextDict() *bool {
	return s.IsFullTextDict
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) GetIsHidden() *bool {
	return s.IsHidden
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) GetPartitionColumn() *string {
	return s.PartitionColumn
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) GetPartitionType() *string {
	return s.PartitionType
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) GetPrimaryKeyColumn() *string {
	return s.PrimaryKeyColumn
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) GetTableEngine() *string {
	return s.TableEngine
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) GetTableSchema() *string {
	return s.TableSchema
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) GetTableType() *string {
	return s.TableType
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) SetCreateTime(v string) *DescribeAdbMySqlTableMetaResponseBodyTableMeta {
	s.CreateTime = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) SetDistributeColumn(v string) *DescribeAdbMySqlTableMetaResponseBodyTableMeta {
	s.DistributeColumn = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) SetDistributeType(v string) *DescribeAdbMySqlTableMetaResponseBodyTableMeta {
	s.DistributeType = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) SetIsAllIndex(v bool) *DescribeAdbMySqlTableMetaResponseBodyTableMeta {
	s.IsAllIndex = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) SetIsDictEncode(v bool) *DescribeAdbMySqlTableMetaResponseBodyTableMeta {
	s.IsDictEncode = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) SetIsFullTextDict(v bool) *DescribeAdbMySqlTableMetaResponseBodyTableMeta {
	s.IsFullTextDict = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) SetIsHidden(v bool) *DescribeAdbMySqlTableMetaResponseBodyTableMeta {
	s.IsHidden = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) SetPartitionColumn(v string) *DescribeAdbMySqlTableMetaResponseBodyTableMeta {
	s.PartitionColumn = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) SetPartitionType(v string) *DescribeAdbMySqlTableMetaResponseBodyTableMeta {
	s.PartitionType = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) SetPrimaryKeyColumn(v string) *DescribeAdbMySqlTableMetaResponseBodyTableMeta {
	s.PrimaryKeyColumn = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) SetTableEngine(v string) *DescribeAdbMySqlTableMetaResponseBodyTableMeta {
	s.TableEngine = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) SetTableName(v string) *DescribeAdbMySqlTableMetaResponseBodyTableMeta {
	s.TableName = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) SetTableSchema(v string) *DescribeAdbMySqlTableMetaResponseBodyTableMeta {
	s.TableSchema = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) SetTableType(v string) *DescribeAdbMySqlTableMetaResponseBodyTableMeta {
	s.TableType = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) SetUpdateTime(v string) *DescribeAdbMySqlTableMetaResponseBodyTableMeta {
	s.UpdateTime = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaResponseBodyTableMeta) Validate() error {
	return dara.Validate(s)
}
