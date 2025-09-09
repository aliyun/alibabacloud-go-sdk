// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSqlFlashbackMatchSwitchRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDbName(v string) *EnableSqlFlashbackMatchSwitchRequest
  GetDbName() *string 
  SetDrdsInstanceId(v string) *EnableSqlFlashbackMatchSwitchRequest
  GetDrdsInstanceId() *string 
}

type EnableSqlFlashbackMatchSwitchRequest struct {
  // The name of the database you want to back up.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // test
  DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
  // The ID of the ApsaraDB RDS for PostgreSQL instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // drds***********
  DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s EnableSqlFlashbackMatchSwitchRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSqlFlashbackMatchSwitchRequest) GoString() string {
  return s.String()
}

func (s *EnableSqlFlashbackMatchSwitchRequest) GetDbName() *string  {
  return s.DbName
}

func (s *EnableSqlFlashbackMatchSwitchRequest) GetDrdsInstanceId() *string  {
  return s.DrdsInstanceId
}

func (s *EnableSqlFlashbackMatchSwitchRequest) SetDbName(v string) *EnableSqlFlashbackMatchSwitchRequest {
  s.DbName = &v
  return s
}

func (s *EnableSqlFlashbackMatchSwitchRequest) SetDrdsInstanceId(v string) *EnableSqlFlashbackMatchSwitchRequest {
  s.DrdsInstanceId = &v
  return s
}

func (s *EnableSqlFlashbackMatchSwitchRequest) Validate() error {
  return dara.Validate(s)
}

