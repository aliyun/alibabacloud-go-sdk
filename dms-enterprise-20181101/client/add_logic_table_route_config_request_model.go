// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLogicTableRouteConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRouteExpr(v string) *AddLogicTableRouteConfigRequest
	GetRouteExpr() *string
	SetRouteKey(v string) *AddLogicTableRouteConfigRequest
	GetRouteKey() *string
	SetTableId(v int64) *AddLogicTableRouteConfigRequest
	GetTableId() *int64
	SetTid(v int64) *AddLogicTableRouteConfigRequest
	GetTid() *int64
}

type AddLogicTableRouteConfigRequest struct {
	// The routing algorithm expression. For more information about how to configure a routing algorithm expression, see [Configure a routing algorithm](https://www.alibabacloud.com/help/en/data-management-service/latest/configure-a-routing-algorithm).
	//
	// This parameter is required.
	//
	// example:
	//
	// #id#%16
	RouteExpr *string `json:"RouteExpr,omitempty" xml:"RouteExpr,omitempty"`
	// The unique key of the routing algorithm.
	//
	// > - You can create a custom unique key for the routing algorithm. No requirements are imposed on custom unique keys.
	//
	// > - The unique key of the routing algorithm in the same logical table must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// id-hash-mod16
	RouteKey *string `json:"RouteKey,omitempty" xml:"RouteKey,omitempty"`
	// The ID of the logical table. You can call the [ListLogicTables](https://www.alibabacloud.com/help/en/data-management-service/latest/listlogictables) operation to query the ID of the logical table.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4****
	TableId *int64 `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://www.alibabacloud.com/help/en/data-management-service/latest/getuseractivetenant) operation to query the tenant ID.
	//
	// example:
	//
	// 4***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s AddLogicTableRouteConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLogicTableRouteConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLogicTableRouteConfigRequest) GetRouteExpr() *string {
	return s.RouteExpr
}

func (s *AddLogicTableRouteConfigRequest) GetRouteKey() *string {
	return s.RouteKey
}

func (s *AddLogicTableRouteConfigRequest) GetTableId() *int64 {
	return s.TableId
}

func (s *AddLogicTableRouteConfigRequest) GetTid() *int64 {
	return s.Tid
}

func (s *AddLogicTableRouteConfigRequest) SetRouteExpr(v string) *AddLogicTableRouteConfigRequest {
	s.RouteExpr = &v
	return s
}

func (s *AddLogicTableRouteConfigRequest) SetRouteKey(v string) *AddLogicTableRouteConfigRequest {
	s.RouteKey = &v
	return s
}

func (s *AddLogicTableRouteConfigRequest) SetTableId(v int64) *AddLogicTableRouteConfigRequest {
	s.TableId = &v
	return s
}

func (s *AddLogicTableRouteConfigRequest) SetTid(v int64) *AddLogicTableRouteConfigRequest {
	s.Tid = &v
	return s
}

func (s *AddLogicTableRouteConfigRequest) Validate() error {
	return dara.Validate(s)
}
