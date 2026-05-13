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
	Bandwidth   *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Connections *int64  `json:"Connections,omitempty" xml:"Connections,omitempty"`
	TableId     *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	TableName   *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
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

func (s *DescribeTairSkvDdbTableResponseBodyTablesTable) SetTableId(v string) *DescribeTairSkvDdbTableResponseBodyTablesTable {
	s.TableId = &v
	return s
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTable) SetTableName(v string) *DescribeTairSkvDdbTableResponseBodyTablesTable {
	s.TableName = &v
	return s
}

func (s *DescribeTairSkvDdbTableResponseBodyTablesTable) Validate() error {
	return dara.Validate(s)
}
