// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDlfDatabase interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseName(v string) *DlfDatabase
	GetDatabaseName() *string
	SetDescription(v string) *DlfDatabase
	GetDescription() *string
	SetTableCount(v int32) *DlfDatabase
	GetTableCount() *int32
}

type DlfDatabase struct {
	// The name of the database.
	//
	// example:
	//
	// my_database
	DatabaseName *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	// The description of the database.
	//
	// example:
	//
	// Test database
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The number of tables in the database. Read-only.
	//
	// example:
	//
	// 10
	TableCount *int32 `json:"tableCount,omitempty" xml:"tableCount,omitempty"`
}

func (s DlfDatabase) String() string {
	return dara.Prettify(s)
}

func (s DlfDatabase) GoString() string {
	return s.String()
}

func (s *DlfDatabase) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DlfDatabase) GetDescription() *string {
	return s.Description
}

func (s *DlfDatabase) GetTableCount() *int32 {
	return s.TableCount
}

func (s *DlfDatabase) SetDatabaseName(v string) *DlfDatabase {
	s.DatabaseName = &v
	return s
}

func (s *DlfDatabase) SetDescription(v string) *DlfDatabase {
	s.Description = &v
	return s
}

func (s *DlfDatabase) SetTableCount(v int32) *DlfDatabase {
	s.TableCount = &v
	return s
}

func (s *DlfDatabase) Validate() error {
	return dara.Validate(s)
}
