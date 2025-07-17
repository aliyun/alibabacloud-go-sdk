// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeSQLLineageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v int64) *AnalyzeSQLLineageRequest
	GetDbId() *int64
	SetSqlContent(v string) *AnalyzeSQLLineageRequest
	GetSqlContent() *string
	SetTid(v int64) *AnalyzeSQLLineageRequest
	GetTid() *int64
}

type AnalyzeSQLLineageRequest struct {
	// The database ID.
	//
	// >  You can call one of the [SearchDatabase](https://help.aliyun.com/document_detail/141876.html), [ListDatabases](https://help.aliyun.com/document_detail/141873.html), and [GetDatabase](https://help.aliyun.com/document_detail/141869.html) operations to obtain the database ID provided in the DatabaseId response parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123***
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The SQL statement.
	//
	// This parameter is required.
	//
	// example:
	//
	// insert into a (id) select id from b;
	SqlContent *string `json:"SqlContent,omitempty" xml:"SqlContent,omitempty"`
	// The tenant ID.
	//
	// >  To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s AnalyzeSQLLineageRequest) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeSQLLineageRequest) GoString() string {
	return s.String()
}

func (s *AnalyzeSQLLineageRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *AnalyzeSQLLineageRequest) GetSqlContent() *string {
	return s.SqlContent
}

func (s *AnalyzeSQLLineageRequest) GetTid() *int64 {
	return s.Tid
}

func (s *AnalyzeSQLLineageRequest) SetDbId(v int64) *AnalyzeSQLLineageRequest {
	s.DbId = &v
	return s
}

func (s *AnalyzeSQLLineageRequest) SetSqlContent(v string) *AnalyzeSQLLineageRequest {
	s.SqlContent = &v
	return s
}

func (s *AnalyzeSQLLineageRequest) SetTid(v int64) *AnalyzeSQLLineageRequest {
	s.Tid = &v
	return s
}

func (s *AnalyzeSQLLineageRequest) Validate() error {
	return dara.Validate(s)
}
