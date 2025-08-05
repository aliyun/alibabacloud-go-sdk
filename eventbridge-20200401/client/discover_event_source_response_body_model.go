// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiscoverEventSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DiscoverEventSourceResponseBody
	GetCode() *string
	SetData(v *DiscoverEventSourceResponseBodyData) *DiscoverEventSourceResponseBody
	GetData() *DiscoverEventSourceResponseBodyData
	SetMessage(v string) *DiscoverEventSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DiscoverEventSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DiscoverEventSourceResponseBody
	GetSuccess() *bool
}

type DiscoverEventSourceResponseBody struct {
	// example:
	//
	// Success
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DiscoverEventSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C7043799-F4DA-5290-9249-97C35987****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DiscoverEventSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DiscoverEventSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DiscoverEventSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DiscoverEventSourceResponseBody) GetData() *DiscoverEventSourceResponseBodyData {
	return s.Data
}

func (s *DiscoverEventSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DiscoverEventSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DiscoverEventSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DiscoverEventSourceResponseBody) SetCode(v string) *DiscoverEventSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DiscoverEventSourceResponseBody) SetData(v *DiscoverEventSourceResponseBodyData) *DiscoverEventSourceResponseBody {
	s.Data = v
	return s
}

func (s *DiscoverEventSourceResponseBody) SetMessage(v string) *DiscoverEventSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DiscoverEventSourceResponseBody) SetRequestId(v string) *DiscoverEventSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DiscoverEventSourceResponseBody) SetSuccess(v bool) *DiscoverEventSourceResponseBody {
	s.Success = &v
	return s
}

func (s *DiscoverEventSourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DiscoverEventSourceResponseBodyData struct {
	SourceMySQLDiscovery *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery `json:"SourceMySQLDiscovery,omitempty" xml:"SourceMySQLDiscovery,omitempty" type:"Struct"`
}

func (s DiscoverEventSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DiscoverEventSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DiscoverEventSourceResponseBodyData) GetSourceMySQLDiscovery() *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery {
	return s.SourceMySQLDiscovery
}

func (s *DiscoverEventSourceResponseBodyData) SetSourceMySQLDiscovery(v *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) *DiscoverEventSourceResponseBodyData {
	s.SourceMySQLDiscovery = v
	return s
}

func (s *DiscoverEventSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery struct {
	DatabaseNames  []*string `json:"DatabaseNames,omitempty" xml:"DatabaseNames,omitempty" type:"Repeated"`
	EstimatedRows  *int64    `json:"EstimatedRows,omitempty" xml:"EstimatedRows,omitempty"`
	ExpireLogsDays *int32    `json:"ExpireLogsDays,omitempty" xml:"ExpireLogsDays,omitempty"`
	// example:
	//
	// [{\\"is_active\\":\\"1\\",\\"name\\":\\"0c0c5d1a-e844-44a8-902d-4f62cbcb0479\\",\\"id\\":\\"21\\"},{\\"is_active\\":\\"1\\",\\"name\\":\\"18ef033e-70bb-4795-8222-6cfb54f38f26\\",\\"id\\":\\"22\\"}]
	SimpleData  *string                                                             `json:"SimpleData,omitempty" xml:"SimpleData,omitempty"`
	TableNames  []*string                                                           `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Repeated"`
	TableSchema *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchema `json:"TableSchema,omitempty" xml:"TableSchema,omitempty" type:"Struct"`
	WaitTimeout *int32                                                              `json:"WaitTimeout,omitempty" xml:"WaitTimeout,omitempty"`
}

func (s DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) String() string {
	return dara.Prettify(s)
}

func (s DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) GoString() string {
	return s.String()
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) GetDatabaseNames() []*string {
	return s.DatabaseNames
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) GetEstimatedRows() *int64 {
	return s.EstimatedRows
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) GetExpireLogsDays() *int32 {
	return s.ExpireLogsDays
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) GetSimpleData() *string {
	return s.SimpleData
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) GetTableNames() []*string {
	return s.TableNames
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) GetTableSchema() *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchema {
	return s.TableSchema
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) GetWaitTimeout() *int32 {
	return s.WaitTimeout
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) SetDatabaseNames(v []*string) *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery {
	s.DatabaseNames = v
	return s
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) SetEstimatedRows(v int64) *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery {
	s.EstimatedRows = &v
	return s
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) SetExpireLogsDays(v int32) *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery {
	s.ExpireLogsDays = &v
	return s
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) SetSimpleData(v string) *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery {
	s.SimpleData = &v
	return s
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) SetTableNames(v []*string) *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery {
	s.TableNames = v
	return s
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) SetTableSchema(v *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchema) *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery {
	s.TableSchema = v
	return s
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) SetWaitTimeout(v int32) *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery {
	s.WaitTimeout = &v
	return s
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscovery) Validate() error {
	return dara.Validate(s)
}

type DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchema struct {
	Columns []*DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// example:
	//
	// map
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchema) String() string {
	return dara.Prettify(s)
}

func (s DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchema) GoString() string {
	return s.String()
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchema) GetColumns() []*DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns {
	return s.Columns
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchema) GetTableName() *string {
	return s.TableName
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchema) SetColumns(v []*DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns) *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchema {
	s.Columns = v
	return s
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchema) SetTableName(v string) *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchema {
	s.TableName = &v
	return s
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchema) Validate() error {
	return dara.Validate(s)
}

type DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns struct {
	// example:
	//
	// auto_increment
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// id
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// NO
	IsNull *string `json:"IsNull,omitempty" xml:"IsNull,omitempty"`
	// example:
	//
	// PRI
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// int
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns) String() string {
	return dara.Prettify(s)
}

func (s DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns) GoString() string {
	return s.String()
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns) GetExtra() *string {
	return s.Extra
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns) GetField() *string {
	return s.Field
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns) GetIsNull() *string {
	return s.IsNull
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns) GetKey() *string {
	return s.Key
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns) GetType() *string {
	return s.Type
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns) SetExtra(v string) *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns {
	s.Extra = &v
	return s
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns) SetField(v string) *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns {
	s.Field = &v
	return s
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns) SetIsNull(v string) *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns {
	s.IsNull = &v
	return s
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns) SetKey(v string) *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns {
	s.Key = &v
	return s
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns) SetType(v string) *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns {
	s.Type = &v
	return s
}

func (s *DiscoverEventSourceResponseBodyDataSourceMySQLDiscoveryTableSchemaColumns) Validate() error {
	return dara.Validate(s)
}
