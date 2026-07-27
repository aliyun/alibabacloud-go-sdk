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
	// The list of objects for which permissions are requested.
	//
	// This parameter is required.
	ApplyObject []*CreatePermissionApplyOrderRequestApplyObject `json:"ApplyObject,omitempty" xml:"ApplyObject,omitempty" type:"Repeated"`
	// The reason for the request. This is used by the administrator for evaluation and approval.
	//
	// This parameter is required.
	//
	// example:
	//
	// I need to use this table
	ApplyReason *string `json:"ApplyReason,omitempty" xml:"ApplyReason,omitempty"`
	// The type of the request order. Valid values:
	//
	// - MaxComputeTable: MaxCompute table permission request order.
	//
	// - MaxComputeFunction: MaxCompute function permission request order.
	//
	// - MaxComputeResource: MaxCompute resource permission request order.
	//
	// - DLFSchema: Data Lake Formation (DLF) 1.0 schema permission request order.
	//
	// - DLFTable: DLF 1.0 table permission request order.
	//
	// - DLFColumn: DLF 1.0 column permission request order.
	//
	// - DsApiDeploy: Data service publication permission request order.
	//
	// example:
	//
	// MaxComputeTable
	ApplyType *string `json:"ApplyType,omitempty" xml:"ApplyType,omitempty"`
	// The UIDs of the Alibaba Cloud accounts for which permissions are requested. Separate multiple account UIDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 26784260040899****,26784260040899****
	ApplyUserIds *string `json:"ApplyUserIds,omitempty" xml:"ApplyUserIds,omitempty"`
	// The name of the data catalog to query. Go to the [Data Lake Formation console](https://dlf.console.aliyun.com/ap-southeast-1/metadata/catalog?spm=a2c4g.11186623.0.0.5a225658pT4Dkr) to view the data catalog name.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The expiration time of the requested permissions. Specify a UNIX timestamp. If you do not specify this parameter, the default expiration time is January 1, 2065.
	//
	// If LabelSecurity is not enabled for the MaxCompute project, or the security level of the requested table field is 0 or less than or equal to the security level of the requesting account, you can request only permanent permissions.
	//
	// Go to the management page of the DataWorks workspace and check the advanced configuration page of the MaxCompute engine to verify whether column-level access control is enabled.
	//
	// Go to the DataWorks workspace to view the security level of fields in Data Map and the security level of accounts on the Member Management page.
	//
	// example:
	//
	// 1617115071885
	Deadline *int64 `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// Deprecated
	//
	// This field is deprecated. Set it to empty.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// odps
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The name of the MaxCompute project for which permissions are requested.
	//
	// example:
	//
	// aMaxcomputeProjectName
	MaxComputeProjectName *string `json:"MaxComputeProjectName,omitempty" xml:"MaxComputeProjectName,omitempty"`
	// Deprecated
	//
	// This field is deprecated. Set it to empty.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1
	OrderType *int32 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the DataWorks workspace to which the MaxCompute project belongs. Go to the DataWorks workspace configuration page to obtain the workspace ID.
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
	if s.ApplyObject != nil {
		for _, item := range s.ApplyObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePermissionApplyOrderRequestApplyObject struct {
	// The permission types to request. Separate multiple permission types with commas (,). Only Select, Describe, Drop, Alter, Update, and Download types are supported.
	//
	// example:
	//
	// Select,Describe
	Actions *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	// The list of column objects.
	ColumnMetaList []*CreatePermissionApplyOrderRequestApplyObjectColumnMetaList `json:"ColumnMetaList,omitempty" xml:"ColumnMetaList,omitempty" type:"Repeated"`
	// The object for which permissions are requested. Only MaxCompute table permissions are supported. Enter the name of the target table.
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
	if s.ColumnMetaList != nil {
		for _, item := range s.ColumnMetaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePermissionApplyOrderRequestApplyObjectColumnMetaList struct {
	// The permission types to request. Separate multiple permission types with commas (,). Only Select, Describe, and Download types are supported.
	//
	// example:
	//
	// Select
	Actions *string `json:"Actions,omitempty" xml:"Actions,omitempty"`
	// The name of the column for which permissions are requested. To request permissions on the entire table, enter all column names of the table.
	//
	// You can request permissions on specific columns only if LabelSecurity is enabled for the MaxCompute project. If LabelSecurity is not enabled, you can request permissions only on the entire table.
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
