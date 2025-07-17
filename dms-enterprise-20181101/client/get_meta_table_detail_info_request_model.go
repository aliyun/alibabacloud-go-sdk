// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableDetailInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTableGuid(v string) *GetMetaTableDetailInfoRequest
	GetTableGuid() *string
	SetTid(v int64) *GetMetaTableDetailInfoRequest
	GetTid() *int64
}

type GetMetaTableDetailInfoRequest struct {
	// The GUID of the table in Data Management (DMS).
	//
	// >
	//
	// 	- You can call the [ListLogicTables](https://help.aliyun.com/document_detail/141875.html) operation with ReturnGuid set to true to query the GUIDs of logical tables in a specific logical database.
	//
	// 	- You can call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation with ReturnGuid set to true to query the GUIDs of tables in a specific physical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// IDB_L_9032.db-test.yuyang_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the Manage DMS tenants topic.
	//
	// example:
	//
	// 123
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetMetaTableDetailInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableDetailInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMetaTableDetailInfoRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTableDetailInfoRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetMetaTableDetailInfoRequest) SetTableGuid(v string) *GetMetaTableDetailInfoRequest {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTableDetailInfoRequest) SetTid(v int64) *GetMetaTableDetailInfoRequest {
	s.Tid = &v
	return s
}

func (s *GetMetaTableDetailInfoRequest) Validate() error {
	return dara.Validate(s)
}
