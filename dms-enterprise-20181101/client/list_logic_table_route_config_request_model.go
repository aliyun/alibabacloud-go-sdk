// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogicTableRouteConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTableId(v int64) *ListLogicTableRouteConfigRequest
	GetTableId() *int64
	SetTid(v int64) *ListLogicTableRouteConfigRequest
	GetTid() *int64
}

type ListLogicTableRouteConfigRequest struct {
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

func (s ListLogicTableRouteConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLogicTableRouteConfigRequest) GoString() string {
	return s.String()
}

func (s *ListLogicTableRouteConfigRequest) GetTableId() *int64 {
	return s.TableId
}

func (s *ListLogicTableRouteConfigRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListLogicTableRouteConfigRequest) SetTableId(v int64) *ListLogicTableRouteConfigRequest {
	s.TableId = &v
	return s
}

func (s *ListLogicTableRouteConfigRequest) SetTid(v int64) *ListLogicTableRouteConfigRequest {
	s.Tid = &v
	return s
}

func (s *ListLogicTableRouteConfigRequest) Validate() error {
	return dara.Validate(s)
}
