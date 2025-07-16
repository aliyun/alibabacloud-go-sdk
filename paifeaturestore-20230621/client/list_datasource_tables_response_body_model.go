// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasourceTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDatasourceTablesResponseBody
	GetRequestId() *string
	SetTables(v []*string) *ListDatasourceTablesResponseBody
	GetTables() []*string
	SetTotalCount(v int64) *ListDatasourceTablesResponseBody
	GetTotalCount() *int64
}

type ListDatasourceTablesResponseBody struct {
	// example:
	//
	// C03B2680-AC9C-59CD-93C5-8142B92537FA
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tables    []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatasourceTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourceTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasourceTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatasourceTablesResponseBody) GetTables() []*string {
	return s.Tables
}

func (s *ListDatasourceTablesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDatasourceTablesResponseBody) SetRequestId(v string) *ListDatasourceTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasourceTablesResponseBody) SetTables(v []*string) *ListDatasourceTablesResponseBody {
	s.Tables = v
	return s
}

func (s *ListDatasourceTablesResponseBody) SetTotalCount(v int64) *ListDatasourceTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatasourceTablesResponseBody) Validate() error {
	return dara.Validate(s)
}
