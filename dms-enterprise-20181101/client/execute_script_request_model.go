// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteScriptRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDbId(v int32) *ExecuteScriptRequest
  GetDbId() *int32 
  SetLogic(v bool) *ExecuteScriptRequest
  GetLogic() *bool 
  SetScript(v string) *ExecuteScriptRequest
  GetScript() *string 
  SetTid(v int64) *ExecuteScriptRequest
  GetTid() *int64 
}

type ExecuteScriptRequest struct {
  // The ID of the database.
  // 
  // >  This parameter is equivalent to the DatabaseId parameter in the SearchDatabase, ListDatabases, and GetDatabase operations. You can call one of these operations to obtain the required database ID. For more information, see [SearchDatabase](https://help.aliyun.com/document_detail/141876.html), [ListDatabases](https://help.aliyun.com/document_detail/141873.html), and [GetDatabase](https://help.aliyun.com/document_detail/141869.html).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 123
  DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
  // Specifies whether the database is a logical database.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // false
  Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
  // The SQL statements to be executed. Data query language (DQL) statements, data definition language (DDL) statements, and data manipulation language (DML) statements are supported. The control mode of the instance that you want to query determines whether you can execute DDL and DML statements.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // select dt from report_daily
  Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
  // The ID of the tenant.
  // 
  // >  To obtain the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Tenant information](https://help.aliyun.com/document_detail/181330.html).
  // 
  // example:
  // 
  // 234
  Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ExecuteScriptRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteScriptRequest) GoString() string {
  return s.String()
}

func (s *ExecuteScriptRequest) GetDbId() *int32  {
  return s.DbId
}

func (s *ExecuteScriptRequest) GetLogic() *bool  {
  return s.Logic
}

func (s *ExecuteScriptRequest) GetScript() *string  {
  return s.Script
}

func (s *ExecuteScriptRequest) GetTid() *int64  {
  return s.Tid
}

func (s *ExecuteScriptRequest) SetDbId(v int32) *ExecuteScriptRequest {
  s.DbId = &v
  return s
}

func (s *ExecuteScriptRequest) SetLogic(v bool) *ExecuteScriptRequest {
  s.Logic = &v
  return s
}

func (s *ExecuteScriptRequest) SetScript(v string) *ExecuteScriptRequest {
  s.Script = &v
  return s
}

func (s *ExecuteScriptRequest) SetTid(v int64) *ExecuteScriptRequest {
  s.Tid = &v
  return s
}

func (s *ExecuteScriptRequest) Validate() error {
  return dara.Validate(s)
}

