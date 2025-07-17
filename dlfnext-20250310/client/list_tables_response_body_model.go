// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListTablesResponseBody
	GetNextPageToken() *string
	SetTables(v []*string) *ListTablesResponseBody
	GetTables() []*string
}

type ListTablesResponseBody struct {
	// example:
	//
	// E8ABEB1C3DB893D16576269017992F57
	NextPageToken *string   `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Tables        []*string `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
}

func (s ListTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListTablesResponseBody) GetTables() []*string {
	return s.Tables
}

func (s *ListTablesResponseBody) SetNextPageToken(v string) *ListTablesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListTablesResponseBody) SetTables(v []*string) *ListTablesResponseBody {
	s.Tables = v
	return s
}

func (s *ListTablesResponseBody) Validate() error {
	return dara.Validate(s)
}
