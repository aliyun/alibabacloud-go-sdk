// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLumaTable interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *LumaTable
	GetCatalog() *string
	SetColumns(v []*LumaTableColumns) *LumaTable
	GetColumns() []*LumaTableColumns
	SetComment(v string) *LumaTable
	GetComment() *string
	SetCreateTime(v int64) *LumaTable
	GetCreateTime() *int64
	SetName(v string) *LumaTable
	GetName() *string
	SetNamespace(v string) *LumaTable
	GetNamespace() *string
	SetRetentionPolicy(v *LumaTableRetentionPolicy) *LumaTable
	GetRetentionPolicy() *LumaTableRetentionPolicy
	SetUpdateTime(v int64) *LumaTable
	GetUpdateTime() *int64
}

type LumaTable struct {
	// The name of the data catalog to which the table belongs.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The list of column definitions for the table. Each column contains Name (column name), Type (data type), and Comment (comment).
	//
	// example:
	//
	// [{"Name":"id","Type":"bigint","Comment":"Primary key"}]
	Columns []*LumaTableColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The comment of the table.
	//
	// example:
	//
	// Test event table
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The creation time of the table. This value is a UNIX timestamp in milliseconds.
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
	// The data retention policy, including the retention days for hot data and cold data.
	RetentionPolicy *LumaTableRetentionPolicy `json:"RetentionPolicy,omitempty" xml:"RetentionPolicy,omitempty" type:"Struct"`
	// The last update time of the table. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1717948800000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s LumaTable) String() string {
	return dara.Prettify(s)
}

func (s LumaTable) GoString() string {
	return s.String()
}

func (s *LumaTable) GetCatalog() *string {
	return s.Catalog
}

func (s *LumaTable) GetColumns() []*LumaTableColumns {
	return s.Columns
}

func (s *LumaTable) GetComment() *string {
	return s.Comment
}

func (s *LumaTable) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *LumaTable) GetName() *string {
	return s.Name
}

func (s *LumaTable) GetNamespace() *string {
	return s.Namespace
}

func (s *LumaTable) GetRetentionPolicy() *LumaTableRetentionPolicy {
	return s.RetentionPolicy
}

func (s *LumaTable) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *LumaTable) SetCatalog(v string) *LumaTable {
	s.Catalog = &v
	return s
}

func (s *LumaTable) SetColumns(v []*LumaTableColumns) *LumaTable {
	s.Columns = v
	return s
}

func (s *LumaTable) SetComment(v string) *LumaTable {
	s.Comment = &v
	return s
}

func (s *LumaTable) SetCreateTime(v int64) *LumaTable {
	s.CreateTime = &v
	return s
}

func (s *LumaTable) SetName(v string) *LumaTable {
	s.Name = &v
	return s
}

func (s *LumaTable) SetNamespace(v string) *LumaTable {
	s.Namespace = &v
	return s
}

func (s *LumaTable) SetRetentionPolicy(v *LumaTableRetentionPolicy) *LumaTable {
	s.RetentionPolicy = v
	return s
}

func (s *LumaTable) SetUpdateTime(v int64) *LumaTable {
	s.UpdateTime = &v
	return s
}

func (s *LumaTable) Validate() error {
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

type LumaTableColumns struct {
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

func (s LumaTableColumns) String() string {
	return dara.Prettify(s)
}

func (s LumaTableColumns) GoString() string {
	return s.String()
}

func (s *LumaTableColumns) GetComment() *string {
	return s.Comment
}

func (s *LumaTableColumns) GetName() *string {
	return s.Name
}

func (s *LumaTableColumns) GetType() *string {
	return s.Type
}

func (s *LumaTableColumns) SetComment(v string) *LumaTableColumns {
	s.Comment = &v
	return s
}

func (s *LumaTableColumns) SetName(v string) *LumaTableColumns {
	s.Name = &v
	return s
}

func (s *LumaTableColumns) SetType(v string) *LumaTableColumns {
	s.Type = &v
	return s
}

func (s *LumaTableColumns) Validate() error {
	return dara.Validate(s)
}

type LumaTableRetentionPolicy struct {
	// The number of days to retain cold data in low-cost archival storage.
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

func (s LumaTableRetentionPolicy) String() string {
	return dara.Prettify(s)
}

func (s LumaTableRetentionPolicy) GoString() string {
	return s.String()
}

func (s *LumaTableRetentionPolicy) GetColdTTL() *int32 {
	return s.ColdTTL
}

func (s *LumaTableRetentionPolicy) GetHotTTL() *int32 {
	return s.HotTTL
}

func (s *LumaTableRetentionPolicy) SetColdTTL(v int32) *LumaTableRetentionPolicy {
	s.ColdTTL = &v
	return s
}

func (s *LumaTableRetentionPolicy) SetHotTTL(v int32) *LumaTableRetentionPolicy {
	s.HotTTL = &v
	return s
}

func (s *LumaTableRetentionPolicy) Validate() error {
	return dara.Validate(s)
}
