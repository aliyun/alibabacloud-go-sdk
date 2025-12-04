// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynDbTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSynDbTablesResponseBody
	GetRequestId() *string
	SetTables(v []*string) *DescribeSynDbTablesResponseBody
	GetTables() []*string
}

type DescribeSynDbTablesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 851D11EA-681C-5B38-A065-C3F90BBD49DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried tables.
	Tables []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s DescribeSynDbTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynDbTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynDbTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSynDbTablesResponseBody) GetTables() []*string {
	return s.Tables
}

func (s *DescribeSynDbTablesResponseBody) SetRequestId(v string) *DescribeSynDbTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynDbTablesResponseBody) SetTables(v []*string) *DescribeSynDbTablesResponseBody {
	s.Tables = v
	return s
}

func (s *DescribeSynDbTablesResponseBody) Validate() error {
	return dara.Validate(s)
}
