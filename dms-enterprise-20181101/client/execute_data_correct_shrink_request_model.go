// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteDataCorrectShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetActionDetailShrink(v string) *ExecuteDataCorrectShrinkRequest
  GetActionDetailShrink() *string 
  SetOrderId(v int64) *ExecuteDataCorrectShrinkRequest
  GetOrderId() *int64 
  SetRealLoginUserUid(v string) *ExecuteDataCorrectShrinkRequest
  GetRealLoginUserUid() *string 
  SetTid(v string) *ExecuteDataCorrectShrinkRequest
  GetTid() *string 
}

type ExecuteDataCorrectShrinkRequest struct {
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
  ActionDetailShrink *string `json:"ActionDetail,omitempty" xml:"ActionDetail,omitempty"`
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

func (s ExecuteDataCorrectShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteDataCorrectShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExecuteDataCorrectShrinkRequest) GetActionDetailShrink() *string  {
  return s.ActionDetailShrink
}

func (s *ExecuteDataCorrectShrinkRequest) GetOrderId() *int64  {
  return s.OrderId
}

func (s *ExecuteDataCorrectShrinkRequest) GetRealLoginUserUid() *string  {
  return s.RealLoginUserUid
}

func (s *ExecuteDataCorrectShrinkRequest) GetTid() *string  {
  return s.Tid
}

func (s *ExecuteDataCorrectShrinkRequest) SetActionDetailShrink(v string) *ExecuteDataCorrectShrinkRequest {
  s.ActionDetailShrink = &v
  return s
}

func (s *ExecuteDataCorrectShrinkRequest) SetOrderId(v int64) *ExecuteDataCorrectShrinkRequest {
  s.OrderId = &v
  return s
}

func (s *ExecuteDataCorrectShrinkRequest) SetRealLoginUserUid(v string) *ExecuteDataCorrectShrinkRequest {
  s.RealLoginUserUid = &v
  return s
}

func (s *ExecuteDataCorrectShrinkRequest) SetTid(v string) *ExecuteDataCorrectShrinkRequest {
  s.Tid = &v
  return s
}

func (s *ExecuteDataCorrectShrinkRequest) Validate() error {
  return dara.Validate(s)
}

