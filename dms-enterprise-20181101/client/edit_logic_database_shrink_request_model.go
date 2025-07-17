// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditLogicDatabaseShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAlias(v string) *EditLogicDatabaseShrinkRequest
  GetAlias() *string 
  SetDatabaseIdsShrink(v string) *EditLogicDatabaseShrinkRequest
  GetDatabaseIdsShrink() *string 
  SetLogicDbId(v int64) *EditLogicDatabaseShrinkRequest
  GetLogicDbId() *int64 
  SetTid(v int64) *EditLogicDatabaseShrinkRequest
  GetTid() *int64 
}

type EditLogicDatabaseShrinkRequest struct {
  // - The alias of the logical database. If you want to change the alias, specify a new alias.
  // 
  // - If you do not need to change the alias of the logical database, call the [GetLogicDatabase](https://www.alibabacloud.com/help/en/data-management-service/latest/getlogicdatabase) or [GetDBTopology](https://www.alibabacloud.com/help/en/data-management-service/latest/getdbtopology) operation to query the alias of the logical database.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // test_logic_db
  Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
  // - The IDs of the physical databases that compose the logical database. If you want to change the physical databases, you can call the [ListDatabases](https://www.alibabacloud.com/help/en/data-management-service/latest/listdatabases) or [SearchDatabase](https://www.alibabacloud.com/help/en/data-management-service/latest/searchdatabase) operation to query the IDs of the new physical databases that you want to specify.
  // 
  // - If you do not want to change the physical databases, you can call the [GetDBTopology](https://www.alibabacloud.com/help/en/data-management-service/latest/getdbtopology) operation to query the IDs of the physical databases that compose the logical database.
  // 
  // This parameter is required.
  DatabaseIdsShrink *string `json:"DatabaseIds,omitempty" xml:"DatabaseIds,omitempty"`
  // The ID of the logical database. You can call the [ListLogicDatabases](https://www.alibabacloud.com/help/en/data-management-service/latest/listlogicdatabases) operation to query the ID of the logical database.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1***
  LogicDbId *int64 `json:"LogicDbId,omitempty" xml:"LogicDbId,omitempty"`
  // The ID of the tenant. 
  // 
  // >  To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://www.alibabacloud.com/help/en/data-management-service/latest/manage-dms-tenants) topic.
  // 
  // example:
  // 
  // 3***
  Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s EditLogicDatabaseShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EditLogicDatabaseShrinkRequest) GoString() string {
  return s.String()
}

func (s *EditLogicDatabaseShrinkRequest) GetAlias() *string  {
  return s.Alias
}

func (s *EditLogicDatabaseShrinkRequest) GetDatabaseIdsShrink() *string  {
  return s.DatabaseIdsShrink
}

func (s *EditLogicDatabaseShrinkRequest) GetLogicDbId() *int64  {
  return s.LogicDbId
}

func (s *EditLogicDatabaseShrinkRequest) GetTid() *int64  {
  return s.Tid
}

func (s *EditLogicDatabaseShrinkRequest) SetAlias(v string) *EditLogicDatabaseShrinkRequest {
  s.Alias = &v
  return s
}

func (s *EditLogicDatabaseShrinkRequest) SetDatabaseIdsShrink(v string) *EditLogicDatabaseShrinkRequest {
  s.DatabaseIdsShrink = &v
  return s
}

func (s *EditLogicDatabaseShrinkRequest) SetLogicDbId(v int64) *EditLogicDatabaseShrinkRequest {
  s.LogicDbId = &v
  return s
}

func (s *EditLogicDatabaseShrinkRequest) SetTid(v int64) *EditLogicDatabaseShrinkRequest {
  s.Tid = &v
  return s
}

func (s *EditLogicDatabaseShrinkRequest) Validate() error {
  return dara.Validate(s)
}

