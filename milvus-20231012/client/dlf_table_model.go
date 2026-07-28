// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDlfTable interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DlfTable
	GetDescription() *string
	SetLocation(v string) *DlfTable
	GetLocation() *string
	SetTableFormat(v string) *DlfTable
	GetTableFormat() *string
	SetTableName(v string) *DlfTable
	GetTableName() *string
	SetTableType(v string) *DlfTable
	GetTableType() *string
}

type DlfTable struct {
	// A description of the table.
	//
	// example:
	//
	// Test table
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The location of the table data, specified as an Object Storage Service (OSS) URI.
	//
	// example:
	//
	// oss://bucket/path
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// The table format, such as `PAIMON`.
	//
	// example:
	//
	// PAIMON
	TableFormat *string `json:"tableFormat,omitempty" xml:"tableFormat,omitempty"`
	// The table name.
	//
	// example:
	//
	// my_table
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// The table type. For example, `MANAGED` indicates that DLF manages the data and metadata lifecycle.
	//
	// example:
	//
	// MANAGED
	TableType *string `json:"tableType,omitempty" xml:"tableType,omitempty"`
}

func (s DlfTable) String() string {
	return dara.Prettify(s)
}

func (s DlfTable) GoString() string {
	return s.String()
}

func (s *DlfTable) GetDescription() *string {
	return s.Description
}

func (s *DlfTable) GetLocation() *string {
	return s.Location
}

func (s *DlfTable) GetTableFormat() *string {
	return s.TableFormat
}

func (s *DlfTable) GetTableName() *string {
	return s.TableName
}

func (s *DlfTable) GetTableType() *string {
	return s.TableType
}

func (s *DlfTable) SetDescription(v string) *DlfTable {
	s.Description = &v
	return s
}

func (s *DlfTable) SetLocation(v string) *DlfTable {
	s.Location = &v
	return s
}

func (s *DlfTable) SetTableFormat(v string) *DlfTable {
	s.TableFormat = &v
	return s
}

func (s *DlfTable) SetTableName(v string) *DlfTable {
	s.TableName = &v
	return s
}

func (s *DlfTable) SetTableType(v string) *DlfTable {
	s.TableType = &v
	return s
}

func (s *DlfTable) Validate() error {
	return dara.Validate(s)
}
