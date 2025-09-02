// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionApplyOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyType(v string) *ListPermissionApplyOrdersRequest
	GetApplyType() *string
	SetCatalogName(v string) *ListPermissionApplyOrdersRequest
	GetCatalogName() *string
	SetEndTime(v int64) *ListPermissionApplyOrdersRequest
	GetEndTime() *int64
	SetEngineType(v string) *ListPermissionApplyOrdersRequest
	GetEngineType() *string
	SetFlowStatus(v int32) *ListPermissionApplyOrdersRequest
	GetFlowStatus() *int32
	SetMaxComputeProjectName(v string) *ListPermissionApplyOrdersRequest
	GetMaxComputeProjectName() *string
	SetOrderType(v int32) *ListPermissionApplyOrdersRequest
	GetOrderType() *int32
	SetPageNum(v int32) *ListPermissionApplyOrdersRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListPermissionApplyOrdersRequest
	GetPageSize() *int32
	SetQueryType(v int32) *ListPermissionApplyOrdersRequest
	GetQueryType() *int32
	SetStartTime(v int64) *ListPermissionApplyOrdersRequest
	GetStartTime() *int64
	SetTableName(v string) *ListPermissionApplyOrdersRequest
	GetTableName() *string
	SetWorkspaceId(v int32) *ListPermissionApplyOrdersRequest
	GetWorkspaceId() *int32
}

type ListPermissionApplyOrdersRequest struct {
	ApplyType   *string `json:"ApplyType,omitempty" xml:"ApplyType,omitempty"`
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The end of the time range to query. You can query all the permissions request orders that have been submitted before the time. The parameter value is a UNIX timestamp. If you do not specify the parameter, all permission request orders that are submitted before the current time are queried.
	//
	// example:
	//
	// 1617200471885
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Deprecated
	//
	// The type of the compute engine with which the permission request order is associated. The parameter value is odps and cannot be changed. This value indicates that you can request permissions only on fields of tables in the MaxCompute compute engine.
	//
	// example:
	//
	// odps
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The status of the permission request order. Valid values:
	//
	// 	- 1: to be processed
	//
	// 	- 2: approved and authorized
	//
	// 	- 3: approved but authorization failed
	//
	// 	- 4: rejected
	//
	// example:
	//
	// 1
	FlowStatus *int32 `json:"FlowStatus,omitempty" xml:"FlowStatus,omitempty"`
	// The name of the MaxCompute project with which the permission request order is associated. If you do not specify the parameter, the permission request orders of all MaxCompute projects are returned.
	//
	// example:
	//
	// aMaxComputeProject
	MaxComputeProjectName *string `json:"MaxComputeProjectName,omitempty" xml:"MaxComputeProjectName,omitempty"`
	// Deprecated
	//
	// The type of the permission request order. The parameter value is 1 and cannot be changed. This value indicates ACL-based authorization.
	//
	// example:
	//
	// 1
	OrderType *int32 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query type of the permission request order. Valid values:
	//
	// 	- 0: The permission request orders you submitted.
	//
	// 	- 1: The permission request orders you approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	QueryType *int32 `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The beginning of the time range to query. You can query all the permissions request orders that have been submitted after the time. The parameter value is a UNIX timestamp. If you do not specify the parameter, all permission request orders are queried.
	//
	// example:
	//
	// 1616200471885
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the table with which the permission request order is associated. If you do not specify the parameter, the permission request orders of all tables are returned.
	//
	// example:
	//
	// aTableName
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the DataWorks workspace that is associated with the permission request order. If you do not specify the parameter, the permission request orders of all workspaces are returned. You can go to the Workspace page in the DataWorks console to obtain the workspace ID.
	//
	// example:
	//
	// 12345
	WorkspaceId *int32 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListPermissionApplyOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionApplyOrdersRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionApplyOrdersRequest) GetApplyType() *string {
	return s.ApplyType
}

func (s *ListPermissionApplyOrdersRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListPermissionApplyOrdersRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListPermissionApplyOrdersRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *ListPermissionApplyOrdersRequest) GetFlowStatus() *int32 {
	return s.FlowStatus
}

func (s *ListPermissionApplyOrdersRequest) GetMaxComputeProjectName() *string {
	return s.MaxComputeProjectName
}

func (s *ListPermissionApplyOrdersRequest) GetOrderType() *int32 {
	return s.OrderType
}

func (s *ListPermissionApplyOrdersRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListPermissionApplyOrdersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPermissionApplyOrdersRequest) GetQueryType() *int32 {
	return s.QueryType
}

func (s *ListPermissionApplyOrdersRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListPermissionApplyOrdersRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListPermissionApplyOrdersRequest) GetWorkspaceId() *int32 {
	return s.WorkspaceId
}

func (s *ListPermissionApplyOrdersRequest) SetApplyType(v string) *ListPermissionApplyOrdersRequest {
	s.ApplyType = &v
	return s
}

func (s *ListPermissionApplyOrdersRequest) SetCatalogName(v string) *ListPermissionApplyOrdersRequest {
	s.CatalogName = &v
	return s
}

func (s *ListPermissionApplyOrdersRequest) SetEndTime(v int64) *ListPermissionApplyOrdersRequest {
	s.EndTime = &v
	return s
}

func (s *ListPermissionApplyOrdersRequest) SetEngineType(v string) *ListPermissionApplyOrdersRequest {
	s.EngineType = &v
	return s
}

func (s *ListPermissionApplyOrdersRequest) SetFlowStatus(v int32) *ListPermissionApplyOrdersRequest {
	s.FlowStatus = &v
	return s
}

func (s *ListPermissionApplyOrdersRequest) SetMaxComputeProjectName(v string) *ListPermissionApplyOrdersRequest {
	s.MaxComputeProjectName = &v
	return s
}

func (s *ListPermissionApplyOrdersRequest) SetOrderType(v int32) *ListPermissionApplyOrdersRequest {
	s.OrderType = &v
	return s
}

func (s *ListPermissionApplyOrdersRequest) SetPageNum(v int32) *ListPermissionApplyOrdersRequest {
	s.PageNum = &v
	return s
}

func (s *ListPermissionApplyOrdersRequest) SetPageSize(v int32) *ListPermissionApplyOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *ListPermissionApplyOrdersRequest) SetQueryType(v int32) *ListPermissionApplyOrdersRequest {
	s.QueryType = &v
	return s
}

func (s *ListPermissionApplyOrdersRequest) SetStartTime(v int64) *ListPermissionApplyOrdersRequest {
	s.StartTime = &v
	return s
}

func (s *ListPermissionApplyOrdersRequest) SetTableName(v string) *ListPermissionApplyOrdersRequest {
	s.TableName = &v
	return s
}

func (s *ListPermissionApplyOrdersRequest) SetWorkspaceId(v int32) *ListPermissionApplyOrdersRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListPermissionApplyOrdersRequest) Validate() error {
	return dara.Validate(s)
}
