// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDasProRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EnableDasProRequest
  GetInstanceId() *string 
  SetSqlRetention(v int32) *EnableDasProRequest
  GetSqlRetention() *int32 
  SetUserId(v string) *EnableDasProRequest
  GetUserId() *string 
}

type EnableDasProRequest struct {
  // The instance ID of the database.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // rm-2ze8g2am97624****
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The storage duration of SQL Explorer data. Unit: days. Default value: **30**. Valid values:
  // 
  // - **30**
  // 
  // - **180**
  // 
  // - **365**
  // 
  // - **1095**
  // 
  // - **1825**
  // 
  // example:
  // 
  // 30
  SqlRetention *int32 `json:"SqlRetention,omitempty" xml:"SqlRetention,omitempty"`
  // The instance ID, which is the ID of the Alibaba Cloud account that created the database instance.
  // 
  // > This parameter is optional. After you specify the ID of the target database instance, the system automatically obtains the instance ID when you call the API operation.
  // 
  // example:
  // 
  // 196278346919****
  UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s EnableDasProRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableDasProRequest) GoString() string {
  return s.String()
}

func (s *EnableDasProRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableDasProRequest) GetSqlRetention() *int32  {
  return s.SqlRetention
}

func (s *EnableDasProRequest) GetUserId() *string  {
  return s.UserId
}

func (s *EnableDasProRequest) SetInstanceId(v string) *EnableDasProRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableDasProRequest) SetSqlRetention(v int32) *EnableDasProRequest {
  s.SqlRetention = &v
  return s
}

func (s *EnableDasProRequest) SetUserId(v string) *EnableDasProRequest {
  s.UserId = &v
  return s
}

func (s *EnableDasProRequest) Validate() error {
  return dara.Validate(s)
}

