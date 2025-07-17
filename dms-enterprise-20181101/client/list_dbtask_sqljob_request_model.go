// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDBTaskSQLJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBTaskGroupId(v int64) *ListDBTaskSQLJobRequest
	GetDBTaskGroupId() *int64
	SetPageNumber(v int64) *ListDBTaskSQLJobRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListDBTaskSQLJobRequest
	GetPageSize() *int64
	SetTid(v int64) *ListDBTaskSQLJobRequest
	GetTid() *int64
}

type ListDBTaskSQLJobRequest struct {
	// The ID of the SQL task group. You can call the [GetStructSyncJobDetail](https://help.aliyun.com/document_detail/206160.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1324235
	DBTaskGroupId *int64 `json:"DBTaskGroupId,omitempty" xml:"DBTaskGroupId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, log on to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListDBTaskSQLJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDBTaskSQLJobRequest) GoString() string {
	return s.String()
}

func (s *ListDBTaskSQLJobRequest) GetDBTaskGroupId() *int64 {
	return s.DBTaskGroupId
}

func (s *ListDBTaskSQLJobRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDBTaskSQLJobRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDBTaskSQLJobRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDBTaskSQLJobRequest) SetDBTaskGroupId(v int64) *ListDBTaskSQLJobRequest {
	s.DBTaskGroupId = &v
	return s
}

func (s *ListDBTaskSQLJobRequest) SetPageNumber(v int64) *ListDBTaskSQLJobRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDBTaskSQLJobRequest) SetPageSize(v int64) *ListDBTaskSQLJobRequest {
	s.PageSize = &v
	return s
}

func (s *ListDBTaskSQLJobRequest) SetTid(v int64) *ListDBTaskSQLJobRequest {
	s.Tid = &v
	return s
}

func (s *ListDBTaskSQLJobRequest) Validate() error {
	return dara.Validate(s)
}
