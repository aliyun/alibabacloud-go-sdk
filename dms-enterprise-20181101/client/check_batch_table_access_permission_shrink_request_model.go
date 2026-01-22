// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBatchTableAccessPermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v int64) *CheckBatchTableAccessPermissionShrinkRequest
	GetDbId() *int64
	SetLogic(v bool) *CheckBatchTableAccessPermissionShrinkRequest
	GetLogic() *bool
	SetPermissionType(v string) *CheckBatchTableAccessPermissionShrinkRequest
	GetPermissionType() *string
	SetTableNameListShrink(v string) *CheckBatchTableAccessPermissionShrinkRequest
	GetTableNameListShrink() *string
	SetTid(v int64) *CheckBatchTableAccessPermissionShrinkRequest
	GetTid() *int64
}

type CheckBatchTableAccessPermissionShrinkRequest struct {
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
	TableNameListShrink *string `json:"TableNameList,omitempty" xml:"TableNameList,omitempty"`
	// The ID of the tenant.
	//
	// >  View Tenant ID by hovering over your profile icon in the DMS console. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CheckBatchTableAccessPermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckBatchTableAccessPermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckBatchTableAccessPermissionShrinkRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *CheckBatchTableAccessPermissionShrinkRequest) GetLogic() *bool {
	return s.Logic
}

func (s *CheckBatchTableAccessPermissionShrinkRequest) GetPermissionType() *string {
	return s.PermissionType
}

func (s *CheckBatchTableAccessPermissionShrinkRequest) GetTableNameListShrink() *string {
	return s.TableNameListShrink
}

func (s *CheckBatchTableAccessPermissionShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CheckBatchTableAccessPermissionShrinkRequest) SetDbId(v int64) *CheckBatchTableAccessPermissionShrinkRequest {
	s.DbId = &v
	return s
}

func (s *CheckBatchTableAccessPermissionShrinkRequest) SetLogic(v bool) *CheckBatchTableAccessPermissionShrinkRequest {
	s.Logic = &v
	return s
}

func (s *CheckBatchTableAccessPermissionShrinkRequest) SetPermissionType(v string) *CheckBatchTableAccessPermissionShrinkRequest {
	s.PermissionType = &v
	return s
}

func (s *CheckBatchTableAccessPermissionShrinkRequest) SetTableNameListShrink(v string) *CheckBatchTableAccessPermissionShrinkRequest {
	s.TableNameListShrink = &v
	return s
}

func (s *CheckBatchTableAccessPermissionShrinkRequest) SetTid(v int64) *CheckBatchTableAccessPermissionShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CheckBatchTableAccessPermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
