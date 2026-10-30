// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTable interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *Table
	GetCatalog() *string
	SetColumns(v []*TableColumns) *Table
	GetColumns() []*TableColumns
	SetComment(v string) *Table
	GetComment() *string
	SetCreateTime(v int64) *Table
	GetCreateTime() *int64
	SetName(v string) *Table
	GetName() *string
	SetNamespace(v string) *Table
	GetNamespace() *string
	SetRetentionPolicy(v *TableRetentionPolicy) *Table
	GetRetentionPolicy() *TableRetentionPolicy
	SetUpdateTime(v int64) *Table
	GetUpdateTime() *int64
}

type Table struct {
	// The name of the data catalog to which the table belongs.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The column definitions of the table. Each column includes the name, data type, and comment.
	//
	// example:
	//
	// [{"Name":"id","Type":"bigint","Comment":"Primary key"}]
	Columns []*TableColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The comment of the table.
	//
	// example:
	//
	// Test event table
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The creation time of the table. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1717948800000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The unique name of the event table.
	//
	// example:
	//
	// my_table
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the namespace to which the table belongs.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The data retention policy. This includes the retention periods for hot and cold data in days.
	RetentionPolicy *TableRetentionPolicy `json:"RetentionPolicy,omitempty" xml:"RetentionPolicy,omitempty" type:"Struct"`
	// The last update time of the table. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1717948800000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Table) String() string {
	return dara.Prettify(s)
}

func (s Table) GoString() string {
	return s.String()
}

func (s *Table) GetCatalog() *string {
	return s.Catalog
}

func (s *Table) GetColumns() []*TableColumns {
	return s.Columns
}

func (s *Table) GetComment() *string {
	return s.Comment
}

func (s *Table) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *Table) GetName() *string {
	return s.Name
}

func (s *Table) GetNamespace() *string {
	return s.Namespace
}

func (s *Table) GetRetentionPolicy() *TableRetentionPolicy {
	return s.RetentionPolicy
}

func (s *Table) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *Table) SetCatalog(v string) *Table {
	s.Catalog = &v
	return s
}

func (s *Table) SetColumns(v []*TableColumns) *Table {
	s.Columns = v
	return s
}

func (s *Table) SetComment(v string) *Table {
	s.Comment = &v
	return s
}

func (s *Table) SetCreateTime(v int64) *Table {
	s.CreateTime = &v
	return s
}

func (s *Table) SetName(v string) *Table {
	s.Name = &v
	return s
}

func (s *Table) SetNamespace(v string) *Table {
	s.Namespace = &v
	return s
}

func (s *Table) SetRetentionPolicy(v *TableRetentionPolicy) *Table {
	s.RetentionPolicy = v
	return s
}

func (s *Table) SetUpdateTime(v int64) *Table {
	s.UpdateTime = &v
	return s
}

func (s *Table) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RetentionPolicy != nil {
		if err := s.RetentionPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TableColumns struct {
	// The comment of the column.
	//
	// example:
	//
	// Primary key ID
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the column.
	//
	// example:
	//
	// id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The data type of the column.
	//
	// example:
	//
	// bigint
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s TableColumns) String() string {
	return dara.Prettify(s)
}

func (s TableColumns) GoString() string {
	return s.String()
}

func (s *TableColumns) GetComment() *string {
	return s.Comment
}

func (s *TableColumns) GetName() *string {
	return s.Name
}

func (s *TableColumns) GetType() *string {
	return s.Type
}

func (s *TableColumns) SetComment(v string) *TableColumns {
	s.Comment = &v
	return s
}

func (s *TableColumns) SetName(v string) *TableColumns {
	s.Name = &v
	return s
}

func (s *TableColumns) SetType(v string) *TableColumns {
	s.Type = &v
	return s
}

func (s *TableColumns) Validate() error {
	return dara.Validate(s)
}

type TableRetentionPolicy struct {
	// The number of days to retain cold data in low-cost archive storage.
	//
	// example:
	//
	// 30
	ColdTTL *int32 `json:"ColdTTL,omitempty" xml:"ColdTTL,omitempty"`
	// The number of days to retain hot data in high-performance query storage.
	//
	// example:
	//
	// 7
	HotTTL *int32 `json:"HotTTL,omitempty" xml:"HotTTL,omitempty"`
}

func (s TableRetentionPolicy) String() string {
	return dara.Prettify(s)
}

func (s TableRetentionPolicy) GoString() string {
	return s.String()
}

func (s *TableRetentionPolicy) GetColdTTL() *int32 {
	return s.ColdTTL
}

func (s *TableRetentionPolicy) GetHotTTL() *int32 {
	return s.HotTTL
}

func (s *TableRetentionPolicy) SetColdTTL(v int32) *TableRetentionPolicy {
	s.ColdTTL = &v
	return s
}

func (s *TableRetentionPolicy) SetHotTTL(v int32) *TableRetentionPolicy {
	s.HotTTL = &v
	return s
}

func (s *TableRetentionPolicy) Validate() error {
	return dara.Validate(s)
}
