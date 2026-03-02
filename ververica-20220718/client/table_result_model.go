// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTableResult interface {
	dara.Model
	String() string
	GoString() string
	SetCollectSinkOperatorId(v string) *TableResult
	GetCollectSinkOperatorId() *string
	SetRowUpdates(v []*RowUpdate) *TableResult
	GetRowUpdates() []*RowUpdate
	SetTableName(v string) *TableResult
	GetTableName() *string
}

type TableResult struct {
	CollectSinkOperatorId *string      `json:"collectSinkOperatorId,omitempty" xml:"collectSinkOperatorId,omitempty"`
	RowUpdates            []*RowUpdate `json:"rowUpdates,omitempty" xml:"rowUpdates,omitempty" type:"Repeated"`
	TableName             *string      `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s TableResult) String() string {
	return dara.Prettify(s)
}

func (s TableResult) GoString() string {
	return s.String()
}

func (s *TableResult) GetCollectSinkOperatorId() *string {
	return s.CollectSinkOperatorId
}

func (s *TableResult) GetRowUpdates() []*RowUpdate {
	return s.RowUpdates
}

func (s *TableResult) GetTableName() *string {
	return s.TableName
}

func (s *TableResult) SetCollectSinkOperatorId(v string) *TableResult {
	s.CollectSinkOperatorId = &v
	return s
}

func (s *TableResult) SetRowUpdates(v []*RowUpdate) *TableResult {
	s.RowUpdates = v
	return s
}

func (s *TableResult) SetTableName(v string) *TableResult {
	s.TableName = &v
	return s
}

func (s *TableResult) Validate() error {
	if s.RowUpdates != nil {
		for _, item := range s.RowUpdates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
