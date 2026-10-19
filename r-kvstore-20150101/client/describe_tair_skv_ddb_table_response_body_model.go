// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairSkvDdbTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTairSkvDdbTableResponseBody
	GetRequestId() *string
	SetTables(v *DescribeTairSkvDdbTableResponseBodyTables) *DescribeTairSkvDdbTableResponseBody
	GetTables() *DescribeTairSkvDdbTableResponseBodyTables
}

type DescribeTairSkvDdbTableResponseBody struct {
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 2363CEDF-C697-14B6-AB9E-C57A4AA0EAD4
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tables    *DescribeTairSkvDdbTableResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
}

func (s DescribeTairSkvDdbTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairSkvDdbTableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTairSkvDdbTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTairSkvDdbTableResponseBody) GetTables() *DescribeTairSkvDdbTableResponseBodyTables {
	return s.Tables
}

func (s *DescribeTairSkvDdbTableResponseBody) SetRequestId(v string) *DescribeTairSkvDdbTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTairSkvDdbTableResponseBody) SetTables(v *DescribeTairSkvDdbTableResponseBodyTables) *DescribeTairSkvDdbTableResponseBody {
	s.Tables = v
	return s
}

func (s *DescribeTairSkvDdbTableResponseBody) Validate() error {
	if s.Tables != nil {
		if err := s.Tables.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTairSkvDdbTableResponseBodyTables struct {
	Table []*DescribeTairSkvDdbTableResponseBodyTablesTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeTairSkvDdbTableResponseBodyTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairSkvDdbTableResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeTairSkvDdbTableResponseBodyTables) GetTable() []*DescribeTairSkvDdbTableResponseBodyTablesTable {
	return s.Table
}

func (s *DescribeTairSkvDdbTableResponseBodyTables) SetTable(v []*DescribeTairSkvDdbTableResponseBodyTablesTable) *DescribeTairSkvDdbTableResponseBodyTables {
	s.Table = v
	return s
}

func (s *DescribeTairSkvDdbTableResponseBodyTables) Validate() error {
	if s.Table != nil {
		for _, item := range s.Table {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTairSkvDdbTableResponseBodyTablesTable struct {
	Bandwidth                *int64                                                                  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Connections              *int64                                                                  `json:"Connections,omitempty" xml:"Connections,omitempty"`
	GlobalSecondaryIndexList *DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexList `json:"GlobalSecondaryIndexList,omitempty" xml:"GlobalSecondaryIndexList,omitempty" type:"Struct"`
	TableId                  *string                                                                 `json:"TableId,omitempty" xml:"TableId,omitempty"`
	TableName                *string                                                                 `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTairSkvDdbTableResponseBodyTablesTable) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairSkvDdbTableResponseBodyTablesTable) GoString() string {
	return s.String()
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTable) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTable) GetConnections() *int64 {
	return s.Connections
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTable) GetGlobalSecondaryIndexList() *DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexList {
	return s.GlobalSecondaryIndexList
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTable) GetTableId() *string {
	return s.TableId
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTable) GetTableName() *string {
	return s.TableName
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTable) SetBandwidth(v int64) *DescribeTairSkvDdbTableResponseBodyTablesTable {
	s.Bandwidth = &v
	return s
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTable) SetConnections(v int64) *DescribeTairSkvDdbTableResponseBodyTablesTable {
	s.Connections = &v
	return s
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTable) SetGlobalSecondaryIndexList(v *DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexList) *DescribeTairSkvDdbTableResponseBodyTablesTable {
	s.GlobalSecondaryIndexList = v
	return s
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTable) SetTableId(v string) *DescribeTairSkvDdbTableResponseBodyTablesTable {
	s.TableId = &v
	return s
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTable) SetTableName(v string) *DescribeTairSkvDdbTableResponseBodyTablesTable {
	s.TableName = &v
	return s
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTable) Validate() error {
	if s.GlobalSecondaryIndexList != nil {
		if err := s.GlobalSecondaryIndexList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexList struct {
	GlobalSecondaryIndex []*DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexListGlobalSecondaryIndex `json:"GlobalSecondaryIndex,omitempty" xml:"GlobalSecondaryIndex,omitempty" type:"Repeated"`
}

func (s DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexList) GoString() string {
	return s.String()
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexList) GetGlobalSecondaryIndex() []*DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexListGlobalSecondaryIndex {
	return s.GlobalSecondaryIndex
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexList) SetGlobalSecondaryIndex(v []*DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexListGlobalSecondaryIndex) *DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexList {
	s.GlobalSecondaryIndex = v
	return s
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexList) Validate() error {
	if s.GlobalSecondaryIndex != nil {
		for _, item := range s.GlobalSecondaryIndex {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexListGlobalSecondaryIndex struct {
	// example:
	//
	// tt-gsi-xxxxxxxxx
	GsiId *string `json:"GsiId,omitempty" xml:"GsiId,omitempty"`
	// example:
	//
	// indexKey1
	IndexName *string `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
}

func (s DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexListGlobalSecondaryIndex) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexListGlobalSecondaryIndex) GoString() string {
	return s.String()
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexListGlobalSecondaryIndex) GetGsiId() *string {
	return s.GsiId
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexListGlobalSecondaryIndex) GetIndexName() *string {
	return s.IndexName
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexListGlobalSecondaryIndex) SetGsiId(v string) *DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexListGlobalSecondaryIndex {
	s.GsiId = &v
	return s
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexListGlobalSecondaryIndex) SetIndexName(v string) *DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexListGlobalSecondaryIndex {
	s.IndexName = &v
	return s
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTableGlobalSecondaryIndexListGlobalSecondaryIndex) Validate() error {
	return dara.Validate(s)
}
