// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDLPartition interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *DLPartition
	GetCatalogName() *string
	SetCreateTime(v int32) *DLPartition
	GetCreateTime() *int32
	SetDbName(v string) *DLPartition
	GetDbName() *string
	SetLastAccessTime(v int32) *DLPartition
	GetLastAccessTime() *int32
	SetParameters(v map[string]*string) *DLPartition
	GetParameters() map[string]*string
	SetSd(v *DLStorageDescriptor) *DLPartition
	GetSd() *DLStorageDescriptor
	SetTableName(v string) *DLPartition
	GetTableName() *string
	SetValues(v []*string) *DLPartition
	GetValues() []*string
}

type DLPartition struct {
	// The name of the catalog.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The time when the partition was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1735109884
	CreateTime *int32 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The time when the table was last accessed.
	//
	// example:
	//
	// 1608707407
	LastAccessTime *int32 `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	// The key-value pair of the partition.
	Parameters map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The description of the data storage.
	Sd *DLStorageDescriptor `json:"Sd,omitempty" xml:"Sd,omitempty"`
	// The table name.
	//
	// example:
	//
	// table_name
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The values in a partition key column.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DLPartition) String() string {
	return dara.Prettify(s)
}

func (s DLPartition) GoString() string {
	return s.String()
}

func (s *DLPartition) GetCatalogName() *string {
	return s.CatalogName
}

func (s *DLPartition) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *DLPartition) GetDbName() *string {
	return s.DbName
}

func (s *DLPartition) GetLastAccessTime() *int32 {
	return s.LastAccessTime
}

func (s *DLPartition) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *DLPartition) GetSd() *DLStorageDescriptor {
	return s.Sd
}

func (s *DLPartition) GetTableName() *string {
	return s.TableName
}

func (s *DLPartition) GetValues() []*string {
	return s.Values
}

func (s *DLPartition) SetCatalogName(v string) *DLPartition {
	s.CatalogName = &v
	return s
}

func (s *DLPartition) SetCreateTime(v int32) *DLPartition {
	s.CreateTime = &v
	return s
}

func (s *DLPartition) SetDbName(v string) *DLPartition {
	s.DbName = &v
	return s
}

func (s *DLPartition) SetLastAccessTime(v int32) *DLPartition {
	s.LastAccessTime = &v
	return s
}

func (s *DLPartition) SetParameters(v map[string]*string) *DLPartition {
	s.Parameters = v
	return s
}

func (s *DLPartition) SetSd(v *DLStorageDescriptor) *DLPartition {
	s.Sd = v
	return s
}

func (s *DLPartition) SetTableName(v string) *DLPartition {
	s.TableName = &v
	return s
}

func (s *DLPartition) SetValues(v []*string) *DLPartition {
	s.Values = v
	return s
}

func (s *DLPartition) Validate() error {
	if s.Sd != nil {
		if err := s.Sd.Validate(); err != nil {
			return err
		}
	}
	return nil
}
