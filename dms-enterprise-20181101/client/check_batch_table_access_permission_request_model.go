// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBatchTableAccessPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v int64) *CheckBatchTableAccessPermissionRequest
	GetDbId() *int64
	SetLogic(v bool) *CheckBatchTableAccessPermissionRequest
	GetLogic() *bool
	SetPermissionType(v string) *CheckBatchTableAccessPermissionRequest
	GetPermissionType() *string
	SetTableNameList(v []*string) *CheckBatchTableAccessPermissionRequest
	GetTableNameList() []*string
	SetTid(v int64) *CheckBatchTableAccessPermissionRequest
	GetTid() *int64
}

type CheckBatchTableAccessPermissionRequest struct {
	// The database ID. You can call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) operation to query the ID of a physical database and the [ListLogicDatabases](https://help.aliyun.com/document_detail/141874.html) operation to query the ID of a logical database.
	//
	// >  The value of DatabaseId is that of DbId.
	//
	// This parameter is required.
	//
	// example:
	//
	// 43153
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- true: Logical database.
	//
	// 	- false: Physical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The type of the permission to be verified.
	//
	// Valid values:
	//
	// 	- QUERY
	//
	// 	- EXPORT
	//
	// 	- CORRECT
	//
	// 	- LOGIN
	//
	// 	- PERF
	//
	// This parameter is required.
	//
	// example:
	//
	// QUERY
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
	// The name of the table.
	//
	// This parameter is required.
	TableNameList []*string `json:"TableNameList,omitempty" xml:"TableNameList,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// >  View Tenant ID by hovering over your profile icon in the DMS console. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CheckBatchTableAccessPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckBatchTableAccessPermissionRequest) GoString() string {
	return s.String()
}

func (s *CheckBatchTableAccessPermissionRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *CheckBatchTableAccessPermissionRequest) GetLogic() *bool {
	return s.Logic
}

func (s *CheckBatchTableAccessPermissionRequest) GetPermissionType() *string {
	return s.PermissionType
}

func (s *CheckBatchTableAccessPermissionRequest) GetTableNameList() []*string {
	return s.TableNameList
}

func (s *CheckBatchTableAccessPermissionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CheckBatchTableAccessPermissionRequest) SetDbId(v int64) *CheckBatchTableAccessPermissionRequest {
	s.DbId = &v
	return s
}

func (s *CheckBatchTableAccessPermissionRequest) SetLogic(v bool) *CheckBatchTableAccessPermissionRequest {
	s.Logic = &v
	return s
}

func (s *CheckBatchTableAccessPermissionRequest) SetPermissionType(v string) *CheckBatchTableAccessPermissionRequest {
	s.PermissionType = &v
	return s
}

func (s *CheckBatchTableAccessPermissionRequest) SetTableNameList(v []*string) *CheckBatchTableAccessPermissionRequest {
	s.TableNameList = v
	return s
}

func (s *CheckBatchTableAccessPermissionRequest) SetTid(v int64) *CheckBatchTableAccessPermissionRequest {
	s.Tid = &v
	return s
}

func (s *CheckBatchTableAccessPermissionRequest) Validate() error {
	return dara.Validate(s)
}
