// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDatabasesResponseBody
	GetRequestId() *string
	SetResult(v *ListDatabasesResponseBodyResult) *ListDatabasesResponseBody
	GetResult() *ListDatabasesResponseBodyResult
}

type ListDatabasesResponseBody struct {
	// id of request
	//
	// example:
	//
	// E45380E8-994A-5402-9806-F114B3295FCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// NodeTreeVO
	Result *ListDatabasesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ListDatabasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatabasesResponseBody) GetResult() *ListDatabasesResponseBodyResult {
	return s.Result
}

func (s *ListDatabasesResponseBody) SetRequestId(v string) *ListDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabasesResponseBody) SetResult(v *ListDatabasesResponseBodyResult) *ListDatabasesResponseBody {
	s.Result = v
	return s
}

func (s *ListDatabasesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDatabasesResponseBodyResult struct {
	Databases []*ListDatabasesResponseBodyResultDatabases `json:"databases,omitempty" xml:"databases,omitempty" type:"Repeated"`
}

func (s ListDatabasesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyResult) GetDatabases() []*ListDatabasesResponseBodyResultDatabases {
	return s.Databases
}

func (s *ListDatabasesResponseBodyResult) SetDatabases(v []*ListDatabasesResponseBodyResultDatabases) *ListDatabasesResponseBodyResult {
	s.Databases = v
	return s
}

func (s *ListDatabasesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListDatabasesResponseBodyResultDatabases struct {
	// example:
	//
	// general
	Database     *string                                                 `json:"database,omitempty" xml:"database,omitempty"`
	Functions    map[string][]*ResultDatabasesFunctionsValue             `json:"functions,omitempty" xml:"functions,omitempty"`
	SqlInstances []*ListDatabasesResponseBodyResultDatabasesSqlInstances `json:"sqlInstances,omitempty" xml:"sqlInstances,omitempty" type:"Repeated"`
	Tables       []*ListDatabasesResponseBodyResultDatabasesTables       `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
	Templates    []*ListDatabasesResponseBodyResultDatabasesTemplates    `json:"templates,omitempty" xml:"templates,omitempty" type:"Repeated"`
}

func (s ListDatabasesResponseBodyResultDatabases) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBodyResultDatabases) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyResultDatabases) GetDatabase() *string {
	return s.Database
}

func (s *ListDatabasesResponseBodyResultDatabases) GetFunctions() map[string][]*ResultDatabasesFunctionsValue {
	return s.Functions
}

func (s *ListDatabasesResponseBodyResultDatabases) GetSqlInstances() []*ListDatabasesResponseBodyResultDatabasesSqlInstances {
	return s.SqlInstances
}

func (s *ListDatabasesResponseBodyResultDatabases) GetTables() []*ListDatabasesResponseBodyResultDatabasesTables {
	return s.Tables
}

func (s *ListDatabasesResponseBodyResultDatabases) GetTemplates() []*ListDatabasesResponseBodyResultDatabasesTemplates {
	return s.Templates
}

func (s *ListDatabasesResponseBodyResultDatabases) SetDatabase(v string) *ListDatabasesResponseBodyResultDatabases {
	s.Database = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabases) SetFunctions(v map[string][]*ResultDatabasesFunctionsValue) *ListDatabasesResponseBodyResultDatabases {
	s.Functions = v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabases) SetSqlInstances(v []*ListDatabasesResponseBodyResultDatabasesSqlInstances) *ListDatabasesResponseBodyResultDatabases {
	s.SqlInstances = v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabases) SetTables(v []*ListDatabasesResponseBodyResultDatabasesTables) *ListDatabasesResponseBodyResultDatabases {
	s.Tables = v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabases) SetTemplates(v []*ListDatabasesResponseBodyResultDatabasesTemplates) *ListDatabasesResponseBodyResultDatabases {
	s.Templates = v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabases) Validate() error {
	return dara.Validate(s)
}

type ListDatabasesResponseBodyResultDatabasesSqlInstances struct {
	Children []interface{} `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// example:
	//
	// 12190
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1
	InstanceId *int64 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// true
	IsDir *int32 `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// example:
	//
	// general
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// -1
	Parent *int64 `json:"parent,omitempty" xml:"parent,omitempty"`
	// example:
	//
	// 1
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// table, instance, template, function
	//
	// example:
	//
	// instance
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDatabasesResponseBodyResultDatabasesSqlInstances) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBodyResultDatabasesSqlInstances) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) GetChildren() []interface{} {
	return s.Children
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) GetId() *int64 {
	return s.Id
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) GetIsDir() *int32 {
	return s.IsDir
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) GetName() *string {
	return s.Name
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) GetParent() *int64 {
	return s.Parent
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) GetType() *string {
	return s.Type
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) SetChildren(v []interface{}) *ListDatabasesResponseBodyResultDatabasesSqlInstances {
	s.Children = v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) SetId(v int64) *ListDatabasesResponseBodyResultDatabasesSqlInstances {
	s.Id = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) SetInstanceId(v int64) *ListDatabasesResponseBodyResultDatabasesSqlInstances {
	s.InstanceId = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) SetIsDir(v int32) *ListDatabasesResponseBodyResultDatabasesSqlInstances {
	s.IsDir = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) SetName(v string) *ListDatabasesResponseBodyResultDatabasesSqlInstances {
	s.Name = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) SetParent(v int64) *ListDatabasesResponseBodyResultDatabasesSqlInstances {
	s.Parent = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) SetTemplateId(v int64) *ListDatabasesResponseBodyResultDatabasesSqlInstances {
	s.TemplateId = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) SetType(v string) *ListDatabasesResponseBodyResultDatabasesSqlInstances {
	s.Type = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesSqlInstances) Validate() error {
	return dara.Validate(s)
}

type ListDatabasesResponseBodyResultDatabasesTables struct {
	Children []interface{} `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// example:
	//
	// 56
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1
	InstanceId *int64 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// true
	IsDir *int32 `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// example:
	//
	// general
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// -1
	Parent *int64 `json:"parent,omitempty" xml:"parent,omitempty"`
	// example:
	//
	// 1
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// table, instance, template, function
	//
	// example:
	//
	// table
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDatabasesResponseBodyResultDatabasesTables) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBodyResultDatabasesTables) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) GetChildren() []interface{} {
	return s.Children
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) GetId() *int64 {
	return s.Id
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) GetIsDir() *int32 {
	return s.IsDir
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) GetName() *string {
	return s.Name
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) GetParent() *int64 {
	return s.Parent
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) GetType() *string {
	return s.Type
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) SetChildren(v []interface{}) *ListDatabasesResponseBodyResultDatabasesTables {
	s.Children = v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) SetId(v int64) *ListDatabasesResponseBodyResultDatabasesTables {
	s.Id = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) SetInstanceId(v int64) *ListDatabasesResponseBodyResultDatabasesTables {
	s.InstanceId = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) SetIsDir(v int32) *ListDatabasesResponseBodyResultDatabasesTables {
	s.IsDir = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) SetName(v string) *ListDatabasesResponseBodyResultDatabasesTables {
	s.Name = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) SetParent(v int64) *ListDatabasesResponseBodyResultDatabasesTables {
	s.Parent = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) SetTemplateId(v int64) *ListDatabasesResponseBodyResultDatabasesTables {
	s.TemplateId = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) SetType(v string) *ListDatabasesResponseBodyResultDatabasesTables {
	s.Type = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesTables) Validate() error {
	return dara.Validate(s)
}

type ListDatabasesResponseBodyResultDatabasesTemplates struct {
	Children []interface{} `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1
	InstanceId *int64 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// true
	IsDir *int32 `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// example:
	//
	// c26_schema
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// -1
	Parent *int64 `json:"parent,omitempty" xml:"parent,omitempty"`
	// example:
	//
	// 1
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// table, instance, template, function
	//
	// example:
	//
	// template
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListDatabasesResponseBodyResultDatabasesTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBodyResultDatabasesTemplates) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) GetChildren() []interface{} {
	return s.Children
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) GetId() *int64 {
	return s.Id
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) GetIsDir() *int32 {
	return s.IsDir
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) GetName() *string {
	return s.Name
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) GetParent() *int64 {
	return s.Parent
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) GetType() *string {
	return s.Type
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) SetChildren(v []interface{}) *ListDatabasesResponseBodyResultDatabasesTemplates {
	s.Children = v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) SetId(v int64) *ListDatabasesResponseBodyResultDatabasesTemplates {
	s.Id = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) SetInstanceId(v int64) *ListDatabasesResponseBodyResultDatabasesTemplates {
	s.InstanceId = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) SetIsDir(v int32) *ListDatabasesResponseBodyResultDatabasesTemplates {
	s.IsDir = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) SetName(v string) *ListDatabasesResponseBodyResultDatabasesTemplates {
	s.Name = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) SetParent(v int64) *ListDatabasesResponseBodyResultDatabasesTemplates {
	s.Parent = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) SetTemplateId(v int64) *ListDatabasesResponseBodyResultDatabasesTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) SetType(v string) *ListDatabasesResponseBodyResultDatabasesTemplates {
	s.Type = &v
	return s
}

func (s *ListDatabasesResponseBodyResultDatabasesTemplates) Validate() error {
	return dara.Validate(s)
}
