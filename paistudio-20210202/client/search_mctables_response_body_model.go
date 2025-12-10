// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMCTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SearchMCTablesResponseBody
	GetRequestId() *string
	SetTables(v []*string) *SearchMCTablesResponseBody
	GetTables() []*string
}

type SearchMCTablesResponseBody struct {
	// example:
	//
	// 18D5A1C6-14B8-545E-8408-0A7DDB4C6B5E
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tables    []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s SearchMCTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMCTablesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMCTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchMCTablesResponseBody) GetTables() []*string {
	return s.Tables
}

func (s *SearchMCTablesResponseBody) SetRequestId(v string) *SearchMCTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMCTablesResponseBody) SetTables(v []*string) *SearchMCTablesResponseBody {
	s.Tables = v
	return s
}

func (s *SearchMCTablesResponseBody) Validate() error {
	return dara.Validate(s)
}
