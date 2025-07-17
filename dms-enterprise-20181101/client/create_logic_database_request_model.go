// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogicDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *CreateLogicDatabaseRequest
	GetAlias() *string
	SetDatabaseIds(v []*int64) *CreateLogicDatabaseRequest
	GetDatabaseIds() []*int64
	SetTid(v int64) *CreateLogicDatabaseRequest
	GetTid() *int64
}

type CreateLogicDatabaseRequest struct {
	// The alias of the logical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_logic_db
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The IDs of the physical databases that compose the logical database. You can specify one or more database IDs. You can call the [ListDatabases](https://www.alibabacloud.com/help/en/data-management-service/latest/listdatabases) or [SearchDatabase](https://www.alibabacloud.com/help/en/data-management-service/latest/searchdatabase) operation to query the IDs of the physical databases.
	//
	// This parameter is required.
	DatabaseIds []*int64 `json:"DatabaseIds,omitempty" xml:"DatabaseIds,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the DMS console. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://www.alibabacloud.com/help/en/data-management-service/latest/manage-dms-tenants) topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateLogicDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLogicDatabaseRequest) GoString() string {
	return s.String()
}

func (s *CreateLogicDatabaseRequest) GetAlias() *string {
	return s.Alias
}

func (s *CreateLogicDatabaseRequest) GetDatabaseIds() []*int64 {
	return s.DatabaseIds
}

func (s *CreateLogicDatabaseRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateLogicDatabaseRequest) SetAlias(v string) *CreateLogicDatabaseRequest {
	s.Alias = &v
	return s
}

func (s *CreateLogicDatabaseRequest) SetDatabaseIds(v []*int64) *CreateLogicDatabaseRequest {
	s.DatabaseIds = v
	return s
}

func (s *CreateLogicDatabaseRequest) SetTid(v int64) *CreateLogicDatabaseRequest {
	s.Tid = &v
	return s
}

func (s *CreateLogicDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
