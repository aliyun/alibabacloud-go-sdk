// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultCatalog(v string) *SubmitQueryRequest
	GetDefaultCatalog() *string
	SetDefaultDatabase(v string) *SubmitQueryRequest
	GetDefaultDatabase() *string
	SetLimit(v int32) *SubmitQueryRequest
	GetLimit() *int32
	SetSql(v string) *SubmitQueryRequest
	GetSql() *string
	SetTier(v string) *SubmitQueryRequest
	GetTier() *string
}

type SubmitQueryRequest struct {
	// The default catalog.
	//
	// example:
	//
	// test_catalog
	DefaultCatalog *string `json:"defaultCatalog,omitempty" xml:"defaultCatalog,omitempty"`
	// The default database.
	//
	// example:
	//
	// test_db
	DefaultDatabase *string `json:"defaultDatabase,omitempty" xml:"defaultDatabase,omitempty"`
	// The limit on the number of returned results.
	//
	// example:
	//
	// 1000
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The SQL text. Multiple statements separated by semicolons are supported and executed sequentially within the same session.
	//
	// example:
	//
	// select 	- from table_name;
	Sql *string `json:"sql,omitempty" xml:"sql,omitempty"`
	// The execution specifications. Default value: standard.
	//
	// example:
	//
	// standard
	Tier *string `json:"tier,omitempty" xml:"tier,omitempty"`
}

func (s SubmitQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitQueryRequest) GoString() string {
	return s.String()
}

func (s *SubmitQueryRequest) GetDefaultCatalog() *string {
	return s.DefaultCatalog
}

func (s *SubmitQueryRequest) GetDefaultDatabase() *string {
	return s.DefaultDatabase
}

func (s *SubmitQueryRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *SubmitQueryRequest) GetSql() *string {
	return s.Sql
}

func (s *SubmitQueryRequest) GetTier() *string {
	return s.Tier
}

func (s *SubmitQueryRequest) SetDefaultCatalog(v string) *SubmitQueryRequest {
	s.DefaultCatalog = &v
	return s
}

func (s *SubmitQueryRequest) SetDefaultDatabase(v string) *SubmitQueryRequest {
	s.DefaultDatabase = &v
	return s
}

func (s *SubmitQueryRequest) SetLimit(v int32) *SubmitQueryRequest {
	s.Limit = &v
	return s
}

func (s *SubmitQueryRequest) SetSql(v string) *SubmitQueryRequest {
	s.Sql = &v
	return s
}

func (s *SubmitQueryRequest) SetTier(v string) *SubmitQueryRequest {
	s.Tier = &v
	return s
}

func (s *SubmitQueryRequest) Validate() error {
	return dara.Validate(s)
}
