// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditLogicDatabaseRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAlias(v string) *EditLogicDatabaseRequest
  GetAlias() *string 
  SetDatabaseIds(v []*int64) *EditLogicDatabaseRequest
  GetDatabaseIds() []*int64 
  SetLogicDbId(v int64) *EditLogicDatabaseRequest
  GetLogicDbId() *int64 
  SetTid(v int64) *EditLogicDatabaseRequest
  GetTid() *int64 
}

type EditLogicDatabaseRequest struct {
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
  DatabaseIds []*int64 `json:"DatabaseIds,omitempty" xml:"DatabaseIds,omitempty" type:"Repeated"`
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

func (s EditLogicDatabaseRequest) String() string {
  return dara.Prettify(s)
}

func (s EditLogicDatabaseRequest) GoString() string {
  return s.String()
}

func (s *EditLogicDatabaseRequest) GetAlias() *string  {
  return s.Alias
}

func (s *EditLogicDatabaseRequest) GetDatabaseIds() []*int64  {
  return s.DatabaseIds
}

func (s *EditLogicDatabaseRequest) GetLogicDbId() *int64  {
  return s.LogicDbId
}

func (s *EditLogicDatabaseRequest) GetTid() *int64  {
  return s.Tid
}

func (s *EditLogicDatabaseRequest) SetAlias(v string) *EditLogicDatabaseRequest {
  s.Alias = &v
  return s
}

func (s *EditLogicDatabaseRequest) SetDatabaseIds(v []*int64) *EditLogicDatabaseRequest {
  s.DatabaseIds = v
  return s
}

func (s *EditLogicDatabaseRequest) SetLogicDbId(v int64) *EditLogicDatabaseRequest {
  s.LogicDbId = &v
  return s
}

func (s *EditLogicDatabaseRequest) SetTid(v int64) *EditLogicDatabaseRequest {
  s.Tid = &v
  return s
}

func (s *EditLogicDatabaseRequest) Validate() error {
  return dara.Validate(s)
}

