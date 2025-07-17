// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogicDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogicDbId(v int64) *DeleteLogicDatabaseRequest
	GetLogicDbId() *int64
	SetTid(v int64) *DeleteLogicDatabaseRequest
	GetTid() *int64
}

type DeleteLogicDatabaseRequest struct {
	// The ID of the logical database. You can call the [ListLogicDatabases](https://www.alibabacloud.com/help/en/data-management-service/latest/listlogicdatabases) or [SearchDatabase](https://www.alibabacloud.com/help/en/data-management-service/latest/searchdatabase) operation to query the ID of the logical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1***
	LogicDbId *int64 `json:"LogicDbId,omitempty" xml:"LogicDbId,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the DMS console. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://www.alibabacloud.com/help/en/data-management-service/latest/manage-dms-tenants) topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DeleteLogicDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogicDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteLogicDatabaseRequest) GetLogicDbId() *int64 {
	return s.LogicDbId
}

func (s *DeleteLogicDatabaseRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteLogicDatabaseRequest) SetLogicDbId(v int64) *DeleteLogicDatabaseRequest {
	s.LogicDbId = &v
	return s
}

func (s *DeleteLogicDatabaseRequest) SetTid(v int64) *DeleteLogicDatabaseRequest {
	s.Tid = &v
	return s
}

func (s *DeleteLogicDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
