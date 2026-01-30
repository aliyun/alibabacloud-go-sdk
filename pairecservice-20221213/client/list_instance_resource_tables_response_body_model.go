// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceResourceTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListInstanceResourceTablesResponseBody
	GetRequestId() *string
	SetTables(v []*ListInstanceResourceTablesResponseBodyTables) *ListInstanceResourceTablesResponseBody
	GetTables() []*ListInstanceResourceTablesResponseBodyTables
	SetTotalCount(v string) *ListInstanceResourceTablesResponseBody
	GetTotalCount() *string
}

type ListInstanceResourceTablesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 01D22D08-BA20-5F35-8302-99115F288220
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tables    []*ListInstanceResourceTablesResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// example:
	//
	// 30
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceResourceTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResourceTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResourceTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceResourceTablesResponseBody) GetTables() []*ListInstanceResourceTablesResponseBodyTables {
	return s.Tables
}

func (s *ListInstanceResourceTablesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListInstanceResourceTablesResponseBody) SetRequestId(v string) *ListInstanceResourceTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceResourceTablesResponseBody) SetTables(v []*ListInstanceResourceTablesResponseBodyTables) *ListInstanceResourceTablesResponseBody {
	s.Tables = v
	return s
}

func (s *ListInstanceResourceTablesResponseBody) SetTotalCount(v string) *ListInstanceResourceTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceResourceTablesResponseBody) Validate() error {
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

type ListInstanceResourceTablesResponseBodyTables struct {
	// example:
	//
	// table-1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListInstanceResourceTablesResponseBodyTables) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResourceTablesResponseBodyTables) GoString() string {
	return s.String()
}

func (s *ListInstanceResourceTablesResponseBodyTables) GetTableName() *string {
	return s.TableName
}

func (s *ListInstanceResourceTablesResponseBodyTables) SetTableName(v string) *ListInstanceResourceTablesResponseBodyTables {
	s.TableName = &v
	return s
}

func (s *ListInstanceResourceTablesResponseBodyTables) Validate() error {
	return dara.Validate(s)
}
