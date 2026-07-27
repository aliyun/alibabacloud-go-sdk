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
	// The type of the application order. Valid values:
	//
	// - [MaxComputeTable] MaxCompute table permission application order.
	//
	// - [MaxComputeFunction] MaxCompute function application order.
	//
	// - [MaxComputeResource] MaxCompute resource application order.
	//
	// - [DLFSchema] DLF 1.0 schema permission application order.
	//
	// - [DLFTable] DLF 1.0 table permission application order.
	//
	// - [DLFColumn] DLF 1.0 column permission application order.
	//
	// - [DsApiDeploy] DataService publishing permission application order.
	//
	// example:
	//
	// MaxComputeTable
	ApplyType *string `json:"ApplyType,omitempty" xml:"ApplyType,omitempty"`
	// The name of the data catalog to query.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The end time for querying application orders, specified as a UNIX timestamp. If this parameter is not specified, application orders up to the current time are queried.
	//
	// example:
	//
	// 1617200471885
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is deprecated and does not take effect.
	//
	// example:
	//
	// odps
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The status of the application order. Valid values:
	//
	// - 1: Pending approval.
	//
	// - 2: Approved, authorization succeeded.
	//
	// - 3: Approved, authorization failed.
	//
	// - 4: Rejected.
	//
	// - 5: Withdrawn.
	//
	// example:
	//
	// 1
	FlowStatus *int32 `json:"FlowStatus,omitempty" xml:"FlowStatus,omitempty"`
	// The name of the MaxCompute project to which the application order belongs. If this parameter is not specified, application orders from all MaxCompute projects are returned.
	//
	// example:
	//
	// aMaxComputeProject
	MaxComputeProjectName *string `json:"MaxComputeProjectName,omitempty" xml:"MaxComputeProjectName,omitempty"`
	// This parameter is deprecated and does not take effect.
	//
	// example:
	//
	// 1
	OrderType *int32 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The page number for paginated queries. The value must be a positive integer greater than or equal to 1. Default value: 1.
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
	// The query type of the application order. Valid values:
	//
	// - 0: Application orders submitted by me.
	//
	// - 1: Application orders approved by me.
	//
	// - 2: All application orders.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	QueryType *int32 `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The start time for querying application orders, specified as a UNIX timestamp. If this parameter is not specified, all application orders are queried.
	//
	// example:
	//
	// 1616200471885
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The table name included in the application order. If this parameter is not specified, application orders for all tables are returned.
	//
	// example:
	//
	// aTableName
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the workspace to which the application order belongs. If this parameter is not specified, application orders from all workspaces are returned. You can log on to the DataWorks console and go to the Workspace Settings page to obtain the workspace ID.
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
