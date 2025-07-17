// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableColumnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTableGuid(v string) *GetMetaTableColumnRequest
	GetTableGuid() *string
	SetTid(v int64) *GetMetaTableColumnRequest
	GetTid() *int64
}

type GetMetaTableColumnRequest struct {
	// The globally unique identifier (GUID) of the table in Data Management (DMS).
	//
	// 	- If the database to which the table belongs is a logical database, you can call the [ListLogicTables](https://help.aliyun.com/document_detail/141875.html) operation to obtain the value of this parameter.
	//
	// 	- If the database to which the table belongs is a physical database, you can call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// IDB_40753****.qntest2.activity_setting
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetMetaTableColumnRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableColumnRequest) GoString() string {
	return s.String()
}

func (s *GetMetaTableColumnRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTableColumnRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetMetaTableColumnRequest) SetTableGuid(v string) *GetMetaTableColumnRequest {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTableColumnRequest) SetTid(v int64) *GetMetaTableColumnRequest {
	s.Tid = &v
	return s
}

func (s *GetMetaTableColumnRequest) Validate() error {
	return dara.Validate(s)
}
