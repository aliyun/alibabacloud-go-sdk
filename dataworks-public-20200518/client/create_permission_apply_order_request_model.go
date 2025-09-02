// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePermissionApplyOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyObject(v []*CreatePermissionApplyOrderRequestApplyObject) *CreatePermissionApplyOrderRequest
	GetApplyObject() []*CreatePermissionApplyOrderRequestApplyObject
	SetApplyReason(v string) *CreatePermissionApplyOrderRequest
	GetApplyReason() *string
	SetApplyType(v string) *CreatePermissionApplyOrderRequest
	GetApplyType() *string
	SetApplyUserIds(v string) *CreatePermissionApplyOrderRequest
	GetApplyUserIds() *string
	SetCatalogName(v string) *CreatePermissionApplyOrderRequest
	GetCatalogName() *string
	SetDeadline(v int64) *CreatePermissionApplyOrderRequest
	GetDeadline() *int64
	SetEngineType(v string) *CreatePermissionApplyOrderRequest
	GetEngineType() *string
	SetMaxComputeProjectName(v string) *CreatePermissionApplyOrderRequest
	GetMaxComputeProjectName() *string
	SetOrderType(v int32) *CreatePermissionApplyOrderRequest
	GetOrderType() *int32
	SetWorkspaceId(v int32) *CreatePermissionApplyOrderRequest
	GetWorkspaceId() *int32
}

type CreatePermissionApplyOrderRequest struct {
	// The objects on which you want to request permissions.
	//
	// This parameter is required.
	ApplyObject []*CreatePermissionApplyOrderRequestApplyObject `json:"ApplyObject,omitempty" xml:"ApplyObject,omitempty" type:"Repeated"`
	// The reason for your request. The administrator determines whether to approve the request based on the reason.
	//
	// This parameter is required.
	//
	// example:
	//
	// I need to use this table
	ApplyReason *string `json:"ApplyReason,omitempty" xml:"ApplyReason,omitempty"`
	ApplyType   *string `json:"ApplyType,omitempty" xml:"ApplyType,omitempty"`
	// The ID of the Alibaba Cloud account for which you want to request permissions. If you want to request permissions for multiple Alibaba Cloud accounts, separate the IDs of the accounts with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 267842600408993176,267842600408993177
	ApplyUserIds *string `json:"ApplyUserIds,omitempty" xml:"ApplyUserIds,omitempty"`
	CatalogName  *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The expiration time of the permissions that you request. This value is a UNIX timestamp. The default value is January 1, 2065. If LabelSecurity is disabled for the MaxCompute project in which you want to request permissions on the fields of a table, or the security level of the fields is 0 or is lower than or equal to the security level of the Alibaba Cloud account for which you want to request permissions, you can request only permanent permissions. You can go to the Workspace Management page in the DataWorks console, click MaxCompute Management in the left-side navigation pane, and then check whether column-level access control is enabled. You can go to your DataWorks workspace, view the security level of the fields in Data Map, and then view the security level of the Alibaba Cloud account on the User Management page.
	//
	// example:
	//
	// 1617115071885
	Deadline *int64 `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// Deprecated
	//
	// The type of the compute engine in which you want to request permissions on the fields of a table. The parameter value is odps and cannot be changed. This value indicates that you can request permissions only on fields of tables in the MaxCompute compute engine.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// odps
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The name of the MaxCompute project in which you request permissions on the fields of a table.
	//
	// example:
	//
	// aMaxcomputeProjectName
	MaxComputeProjectName *string `json:"MaxComputeProjectName,omitempty" xml:"MaxComputeProjectName,omitempty"`
	// Deprecated
	//
	// The type of the permission request order. The parameter value is 1 and cannot be changed. This value indicates ACL-based authorization.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1
	OrderType *int32 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the DataWorks workspace that is associated with the MaxCompute project in which you want to request permissions on the fields of a table. You can go to the SettingCenter page in the DataWorks console to view the workspace ID.
	//
	// example:
	//
	// 12345
	WorkspaceId *int32 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreatePermissionApplyOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePermissionApplyOrderRequest) GoString() string {
	return s.String()
}

func (s *CreatePermissionApplyOrderRequest) GetApplyObject() []*CreatePermissionApplyOrderRequestApplyObject {
	return s.ApplyObject
}

func (s *CreatePermissionApplyOrderRequest) GetApplyReason() *string {
	return s.ApplyReason
}

func (s *CreatePermissionApplyOrderRequest) GetApplyType() *string {
	return s.ApplyType
}

func (s *CreatePermissionApplyOrderRequest) GetApplyUserIds() *string {
	return s.ApplyUserIds
}

func (s *CreatePermissionApplyOrderRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *CreatePermissionApplyOrderRequest) GetDeadline() *int64 {
	return s.Deadline
}

func (s *CreatePermissionApplyOrderRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *CreatePermissionApplyOrderRequest) GetMaxComputeProjectName() *string {
	return s.MaxComputeProjectName
}

func (s *CreatePermissionApplyOrderRequest) GetOrderType() *int32 {
	return s.OrderType
}

func (s *CreatePermissionApplyOrderRequest) GetWorkspaceId() *int32 {
	return s.WorkspaceId
}

func (s *CreatePermissionApplyOrderRequest) SetApplyObject(v []*CreatePermissionApplyOrderRequestApplyObject) *CreatePermissionApplyOrderRequest {
	s.ApplyObject = v
	return s
}

func (s *CreatePermissionApplyOrderRequest) SetApplyReason(v string) *CreatePermissionApplyOrderRequest {
	s.ApplyReason = &v
	return s
}

func (s *CreatePermissionApplyOrderRequest) SetApplyType(v string) *CreatePermissionApplyOrderRequest {
	s.ApplyType = &v
	return s
}

func (s *CreatePermissionApplyOrderRequest) SetApplyUserIds(v string) *CreatePermissionApplyOrderRequest {
	s.ApplyUserIds = &v
	return s
}

func (s *CreatePermissionApplyOrderRequest) SetCatalogName(v string) *CreatePermissionApplyOrderRequest {
	s.CatalogName = &v
	return s
}

func (s *CreatePermissionApplyOrderRequest) SetDeadline(v int64) *CreatePermissionApplyOrderRequest {
	s.Deadline = &v
	return s
}

func (s *CreatePermissionApplyOrderRequest) SetEngineType(v string) *CreatePermissionApplyOrderRequest {
	s.EngineType = &v
	return s
}

func (s *CreatePermissionApplyOrderRequest) SetMaxComputeProjectName(v string) *CreatePermissionApplyOrderRequest {
	s.MaxComputeProjectName = &v
	return s
}

func (s *CreatePermissionApplyOrderRequest) SetOrderType(v int32) *CreatePermissionApplyOrderRequest {
	s.OrderType = &v
	return s
}

func (s *CreatePermissionApplyOrderRequest) SetWorkspaceId(v int32) *CreatePermissionApplyOrderRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreatePermissionApplyOrderRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePermissionApplyOrderRequestApplyObject struct {
	// The permission that you want to request. If you want to request multiple permissions at the same time, separate them with commas (,). You can request only the following permissions: Select, Describe, Drop, Alter, Update, and Download.
	//
	// example:
	//
	// Select,Describe
	Actions *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	// The fields on which you want to request permissions.
	ColumnMetaList []*CreatePermissionApplyOrderRequestApplyObjectColumnMetaList `json:"ColumnMetaList,omitempty" xml:"ColumnMetaList,omitempty" type:"Repeated"`
	// The name of the object on which you want to request permissions. You can request permissions only on MaxCompute tables. Set this parameter to the name of the table on which you want to request permissions.
	//
	// example:
	//
	// aTableName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreatePermissionApplyOrderRequestApplyObject) String() string {
	return dara.Prettify(s)
}

func (s CreatePermissionApplyOrderRequestApplyObject) GoString() string {
	return s.String()
}

func (s *CreatePermissionApplyOrderRequestApplyObject) GetActions() *string {
	return s.Actions
}

func (s *CreatePermissionApplyOrderRequestApplyObject) GetColumnMetaList() []*CreatePermissionApplyOrderRequestApplyObjectColumnMetaList {
	return s.ColumnMetaList
}

func (s *CreatePermissionApplyOrderRequestApplyObject) GetName() *string {
	return s.Name
}

func (s *CreatePermissionApplyOrderRequestApplyObject) SetActions(v string) *CreatePermissionApplyOrderRequestApplyObject {
	s.Actions = &v
	return s
}

func (s *CreatePermissionApplyOrderRequestApplyObject) SetColumnMetaList(v []*CreatePermissionApplyOrderRequestApplyObjectColumnMetaList) *CreatePermissionApplyOrderRequestApplyObject {
	s.ColumnMetaList = v
	return s
}

func (s *CreatePermissionApplyOrderRequestApplyObject) SetName(v string) *CreatePermissionApplyOrderRequestApplyObject {
	s.Name = &v
	return s
}

func (s *CreatePermissionApplyOrderRequestApplyObject) Validate() error {
	return dara.Validate(s)
}

type CreatePermissionApplyOrderRequestApplyObjectColumnMetaList struct {
	Actions *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	// The field on which you want to request permissions. If you want to request permissions on an entire table, enter all fields in the table. You can request permissions on specific fields of a table in a MaxCompute project only after LabelSecurity is enabled for this project. If LabelSecurity is disabled, you can request permissions only on an entire table.
	//
	// example:
	//
	// aColumnName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreatePermissionApplyOrderRequestApplyObjectColumnMetaList) String() string {
	return dara.Prettify(s)
}

func (s CreatePermissionApplyOrderRequestApplyObjectColumnMetaList) GoString() string {
	return s.String()
}

func (s *CreatePermissionApplyOrderRequestApplyObjectColumnMetaList) GetActions() *string {
	return s.Actions
}

func (s *CreatePermissionApplyOrderRequestApplyObjectColumnMetaList) GetName() *string {
	return s.Name
}

func (s *CreatePermissionApplyOrderRequestApplyObjectColumnMetaList) SetActions(v string) *CreatePermissionApplyOrderRequestApplyObjectColumnMetaList {
	s.Actions = &v
	return s
}

func (s *CreatePermissionApplyOrderRequestApplyObjectColumnMetaList) SetName(v string) *CreatePermissionApplyOrderRequestApplyObjectColumnMetaList {
	s.Name = &v
	return s
}

func (s *CreatePermissionApplyOrderRequestApplyObjectColumnMetaList) Validate() error {
	return dara.Validate(s)
}
