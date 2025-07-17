// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogicTableRouteConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRouteKey(v string) *DeleteLogicTableRouteConfigRequest
	GetRouteKey() *string
	SetTableId(v int64) *DeleteLogicTableRouteConfigRequest
	GetTableId() *int64
	SetTid(v int64) *DeleteLogicTableRouteConfigRequest
	GetTid() *int64
}

type DeleteLogicTableRouteConfigRequest struct {
	// The unique key of the routing algorithm. You can call the [ListLogicTableRouteConfig](https://www.alibabacloud.com/help/en/data-management-service/latest/listlogictablerouteconfig) operation to query the unique key.
	//
	// This parameter is required.
	//
	// example:
	//
	// id-hash-mod15
	RouteKey *string `json:"RouteKey,omitempty" xml:"RouteKey,omitempty"`
	// The ID of the logical table. You can call the [ListLogicTables](https://www.alibabacloud.com/help/en/data-management-service/latest/listlogictables) operation to query the ID of the logical table.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1****
	TableId *int64 `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://www.alibabacloud.com/help/en/data-management-service/latest/getuseractivetenant) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DeleteLogicTableRouteConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogicTableRouteConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLogicTableRouteConfigRequest) GetRouteKey() *string {
	return s.RouteKey
}

func (s *DeleteLogicTableRouteConfigRequest) GetTableId() *int64 {
	return s.TableId
}

func (s *DeleteLogicTableRouteConfigRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteLogicTableRouteConfigRequest) SetRouteKey(v string) *DeleteLogicTableRouteConfigRequest {
	s.RouteKey = &v
	return s
}

func (s *DeleteLogicTableRouteConfigRequest) SetTableId(v int64) *DeleteLogicTableRouteConfigRequest {
	s.TableId = &v
	return s
}

func (s *DeleteLogicTableRouteConfigRequest) SetTid(v int64) *DeleteLogicTableRouteConfigRequest {
	s.Tid = &v
	return s
}

func (s *DeleteLogicTableRouteConfigRequest) Validate() error {
	return dara.Validate(s)
}
