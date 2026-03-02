// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNode interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *Node
	GetCatalogName() *string
	SetConnector(v string) *Node
	GetConnector() *string
	SetDatabaseName(v string) *Node
	GetDatabaseName() *string
	SetId(v string) *Node
	GetId() *string
	SetIsTemporary(v bool) *Node
	GetIsTemporary() *bool
	SetTables(v []*LineageTable) *Node
	GetTables() []*LineageTable
}

type Node struct {
	// The name of the catalog.
	//
	// example:
	//
	// paimon
	CatalogName *string `json:"catalogName,omitempty" xml:"catalogName,omitempty"`
	// The name of the connector.
	//
	// example:
	//
	// paimon
	Connector *string `json:"connector,omitempty" xml:"connector,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// paimon-ods
	DatabaseName *string `json:"databaseName,omitempty" xml:"databaseName,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 97383
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Specifies whether the table is a temporary table.
	//
	// example:
	//
	// false
	IsTemporary *bool `json:"isTemporary,omitempty" xml:"isTemporary,omitempty"`
	// The information about the table.
	Tables []*LineageTable `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
}

func (s Node) String() string {
	return dara.Prettify(s)
}

func (s Node) GoString() string {
	return s.String()
}

func (s *Node) GetCatalogName() *string {
	return s.CatalogName
}

func (s *Node) GetConnector() *string {
	return s.Connector
}

func (s *Node) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *Node) GetId() *string {
	return s.Id
}

func (s *Node) GetIsTemporary() *bool {
	return s.IsTemporary
}

func (s *Node) GetTables() []*LineageTable {
	return s.Tables
}

func (s *Node) SetCatalogName(v string) *Node {
	s.CatalogName = &v
	return s
}

func (s *Node) SetConnector(v string) *Node {
	s.Connector = &v
	return s
}

func (s *Node) SetDatabaseName(v string) *Node {
	s.DatabaseName = &v
	return s
}

func (s *Node) SetId(v string) *Node {
	s.Id = &v
	return s
}

func (s *Node) SetIsTemporary(v bool) *Node {
	s.IsTemporary = &v
	return s
}

func (s *Node) SetTables(v []*LineageTable) *Node {
	s.Tables = v
	return s
}

func (s *Node) Validate() error {
	if s.Tables != nil {
		for _, item := range s.Tables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
