// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSqlConcurrencyControlRequest interface {
  dara.Model
  String() string
  GoString() string
  SetConcurrencyControlTime(v int64) *EnableSqlConcurrencyControlRequest
  GetConcurrencyControlTime() *int64 
  SetConsoleContext(v string) *EnableSqlConcurrencyControlRequest
  GetConsoleContext() *string 
  SetInstanceId(v string) *EnableSqlConcurrencyControlRequest
  GetInstanceId() *string 
  SetMaxConcurrency(v int64) *EnableSqlConcurrencyControlRequest
  GetMaxConcurrency() *int64 
  SetSqlKeywords(v string) *EnableSqlConcurrencyControlRequest
  GetSqlKeywords() *string 
  SetSqlType(v string) *EnableSqlConcurrencyControlRequest
  GetSqlType() *string 
}

type EnableSqlConcurrencyControlRequest struct {
  // The duration within which the SQL throttling rule takes effect. Unit: seconds.
  // 
  // >  The throttling rule takes effect only within this duration.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 300
  ConcurrencyControlTime *int64 `json:"ConcurrencyControlTime,omitempty" xml:"ConcurrencyControlTime,omitempty"`
  // The reserved parameter.
  // 
  // example:
  // 
  // None
  ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
  // The instance ID.
  // 
  // >  You must specify the instance ID only if your database instance is an ApsaraDB RDS for MySQL instance or a PolarDB for MySQL cluster.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // rm-2ze1jdv45i7l6****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The maximum number of concurrent SQL statements. Set this parameter to a positive integer.
  // 
  // >  When the number of concurrent SQL statements that contain the specified keywords reaches this upper limit, the throttling rule is triggered.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 3
  MaxConcurrency *int64 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
  // The keywords that are used to identify the SQL statements that need to be throttled.
  // 
  // >  If you specify multiple SQL keywords, separate them with tildes (~). If the number of concurrent SQL statements that contain all the specified SQL keywords reaches the specified upper limit, the throttling rule is triggered.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // call~open~api~test~4~from~POP
  SqlKeywords *string `json:"SqlKeywords,omitempty" xml:"SqlKeywords,omitempty"`
  // The type of the SQL statements. Valid values:
  // 
  // 	- **SELECT**
  // 
  // 	- **UPDATE**
  // 
  // 	- **DELETE**
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // SELECT
  SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
}

func (s EnableSqlConcurrencyControlRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSqlConcurrencyControlRequest) GoString() string {
  return s.String()
}

func (s *EnableSqlConcurrencyControlRequest) GetConcurrencyControlTime() *int64  {
  return s.ConcurrencyControlTime
}

func (s *EnableSqlConcurrencyControlRequest) GetConsoleContext() *string  {
  return s.ConsoleContext
}

func (s *EnableSqlConcurrencyControlRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableSqlConcurrencyControlRequest) GetMaxConcurrency() *int64  {
  return s.MaxConcurrency
}

func (s *EnableSqlConcurrencyControlRequest) GetSqlKeywords() *string  {
  return s.SqlKeywords
}

func (s *EnableSqlConcurrencyControlRequest) GetSqlType() *string  {
  return s.SqlType
}

func (s *EnableSqlConcurrencyControlRequest) SetConcurrencyControlTime(v int64) *EnableSqlConcurrencyControlRequest {
  s.ConcurrencyControlTime = &v
  return s
}

func (s *EnableSqlConcurrencyControlRequest) SetConsoleContext(v string) *EnableSqlConcurrencyControlRequest {
  s.ConsoleContext = &v
  return s
}

func (s *EnableSqlConcurrencyControlRequest) SetInstanceId(v string) *EnableSqlConcurrencyControlRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableSqlConcurrencyControlRequest) SetMaxConcurrency(v int64) *EnableSqlConcurrencyControlRequest {
  s.MaxConcurrency = &v
  return s
}

func (s *EnableSqlConcurrencyControlRequest) SetSqlKeywords(v string) *EnableSqlConcurrencyControlRequest {
  s.SqlKeywords = &v
  return s
}

func (s *EnableSqlConcurrencyControlRequest) SetSqlType(v string) *EnableSqlConcurrencyControlRequest {
  s.SqlType = &v
  return s
}

func (s *EnableSqlConcurrencyControlRequest) Validate() error {
  return dara.Validate(s)
}

