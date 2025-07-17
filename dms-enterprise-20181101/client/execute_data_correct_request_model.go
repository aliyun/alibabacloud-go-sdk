// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteDataCorrectRequest interface {
  dara.Model
  String() string
  GoString() string
  SetActionDetail(v map[string]interface{}) *ExecuteDataCorrectRequest
  GetActionDetail() map[string]interface{} 
  SetOrderId(v int64) *ExecuteDataCorrectRequest
  GetOrderId() *int64 
  SetRealLoginUserUid(v string) *ExecuteDataCorrectRequest
  GetRealLoginUserUid() *string 
  SetTid(v string) *ExecuteDataCorrectRequest
  GetTid() *string 
}

type ExecuteDataCorrectRequest struct {
  // The parameters that are required to perform the data change.
  // 
  // ```
  // 
  // json
  // 
  // "actionDetail" : {
  // 
  //     "startTime" :"2021-07-01 00:00:00", // Specify the start time to change data. If you want to immediately change data, you do not need to set this parameter. 
  // 
  //     "endTime" : "2021-07-01 01:00:00", // Specify the end time to change data. If you want to immediately change data, you do not need to set this parameter. 
  // 
  //     "transaction" : false, // Specify whether to change data as a transaction. 
  // 
  //     "backupData" : true // Specify whether to back up data. 
  // 
  //   }
  // 
  // ```
  // 
  // example:
  // 
  // { "startTime" : "2021-07-01 00:00:00", "endTime" : "2021-07-01 01:00:00", "transaction" : false, "backupData" : true }
  ActionDetail map[string]interface{} `json:"ActionDetail,omitempty" xml:"ActionDetail,omitempty"`
  // The ID of the ticket. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to query the ID of the ticket.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 406****
  OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
  // The ID of the Alibaba Cloud account that is used to call the API operation.
  // 
  // example:
  // 
  // 21400447956867****
  RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
  // The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
  // 
  // example:
  // 
  // 3***
  Tid *string `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ExecuteDataCorrectRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteDataCorrectRequest) GoString() string {
  return s.String()
}

func (s *ExecuteDataCorrectRequest) GetActionDetail() map[string]interface{}  {
  return s.ActionDetail
}

func (s *ExecuteDataCorrectRequest) GetOrderId() *int64  {
  return s.OrderId
}

func (s *ExecuteDataCorrectRequest) GetRealLoginUserUid() *string  {
  return s.RealLoginUserUid
}

func (s *ExecuteDataCorrectRequest) GetTid() *string  {
  return s.Tid
}

func (s *ExecuteDataCorrectRequest) SetActionDetail(v map[string]interface{}) *ExecuteDataCorrectRequest {
  s.ActionDetail = v
  return s
}

func (s *ExecuteDataCorrectRequest) SetOrderId(v int64) *ExecuteDataCorrectRequest {
  s.OrderId = &v
  return s
}

func (s *ExecuteDataCorrectRequest) SetRealLoginUserUid(v string) *ExecuteDataCorrectRequest {
  s.RealLoginUserUid = &v
  return s
}

func (s *ExecuteDataCorrectRequest) SetTid(v string) *ExecuteDataCorrectRequest {
  s.Tid = &v
  return s
}

func (s *ExecuteDataCorrectRequest) Validate() error {
  return dara.Validate(s)
}

