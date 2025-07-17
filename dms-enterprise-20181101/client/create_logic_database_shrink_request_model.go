// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogicDatabaseShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *CreateLogicDatabaseShrinkRequest
	GetAlias() *string
	SetDatabaseIdsShrink(v string) *CreateLogicDatabaseShrinkRequest
	GetDatabaseIdsShrink() *string
	SetTid(v int64) *CreateLogicDatabaseShrinkRequest
	GetTid() *int64
}

type CreateLogicDatabaseShrinkRequest struct {
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
	DatabaseIdsShrink *string `json:"DatabaseIds,omitempty" xml:"DatabaseIds,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the DMS console. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://www.alibabacloud.com/help/en/data-management-service/latest/manage-dms-tenants) topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateLogicDatabaseShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLogicDatabaseShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLogicDatabaseShrinkRequest) GetAlias() *string {
	return s.Alias
}

func (s *CreateLogicDatabaseShrinkRequest) GetDatabaseIdsShrink() *string {
	return s.DatabaseIdsShrink
}

func (s *CreateLogicDatabaseShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateLogicDatabaseShrinkRequest) SetAlias(v string) *CreateLogicDatabaseShrinkRequest {
	s.Alias = &v
	return s
}

func (s *CreateLogicDatabaseShrinkRequest) SetDatabaseIdsShrink(v string) *CreateLogicDatabaseShrinkRequest {
	s.DatabaseIdsShrink = &v
	return s
}

func (s *CreateLogicDatabaseShrinkRequest) SetTid(v int64) *CreateLogicDatabaseShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateLogicDatabaseShrinkRequest) Validate() error {
	return dara.Validate(s)
}
