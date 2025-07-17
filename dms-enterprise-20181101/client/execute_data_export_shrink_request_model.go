// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteDataExportShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetActionDetailShrink(v string) *ExecuteDataExportShrinkRequest
  GetActionDetailShrink() *string 
  SetOrderId(v int64) *ExecuteDataExportShrinkRequest
  GetOrderId() *int64 
  SetRealLoginUserUid(v string) *ExecuteDataExportShrinkRequest
  GetRealLoginUserUid() *string 
  SetTid(v int64) *ExecuteDataExportShrinkRequest
  GetTid() *int64 
}

type ExecuteDataExportShrinkRequest struct {
  // The parameters that are required to perform the operation. Sample code:
  // 
  // ```json
  // 
  // {
  // 
  //   "mode" : "FAST",   // The mode in which data is exported. Default value: FAST. A value of NORMAL specifies that the export task can be terminated during the export.  "encoding" : "UTF8",  // The encoding format.  "startTime" : "2022-12-22 00:00:00",  // The point in time at which data export starts.  "transaction" : false,    // Specifies whether to enable transactions.  "fileType" : "SQL"    // The format of the exported file.}
  // 
  // ```
  // 
  // >  You can also set mode, encoding, and fileType to the following values:
  // 
  // 	- mode: NORMAL
  // 
  // 	- encoding: UTF8MB4, GB2312, ISO_8859_1, GBK, LATAIN1, or CP1252
  // 
  // 	- fileType: XLSX, CSV, JSON, or TXT
  // 
  // example:
  // 
  // {    "fileType": "CSV",    "encoding": ""  }
  ActionDetailShrink *string `json:"ActionDetail,omitempty" xml:"ActionDetail,omitempty"`
  // The ID of the ticket.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1234
  OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
  // The ID of the Alibaba Cloud account that is used to call the API operation.
  // 
  // example:
  // 
  // 21400447956867****
  RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
  // The ID of the tenant.
  // 
  // > : To view the ID of the tenant, log on to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
  // 
  // example:
  // 
  // -1
  Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ExecuteDataExportShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteDataExportShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExecuteDataExportShrinkRequest) GetActionDetailShrink() *string  {
  return s.ActionDetailShrink
}

func (s *ExecuteDataExportShrinkRequest) GetOrderId() *int64  {
  return s.OrderId
}

func (s *ExecuteDataExportShrinkRequest) GetRealLoginUserUid() *string  {
  return s.RealLoginUserUid
}

func (s *ExecuteDataExportShrinkRequest) GetTid() *int64  {
  return s.Tid
}

func (s *ExecuteDataExportShrinkRequest) SetActionDetailShrink(v string) *ExecuteDataExportShrinkRequest {
  s.ActionDetailShrink = &v
  return s
}

func (s *ExecuteDataExportShrinkRequest) SetOrderId(v int64) *ExecuteDataExportShrinkRequest {
  s.OrderId = &v
  return s
}

func (s *ExecuteDataExportShrinkRequest) SetRealLoginUserUid(v string) *ExecuteDataExportShrinkRequest {
  s.RealLoginUserUid = &v
  return s
}

func (s *ExecuteDataExportShrinkRequest) SetTid(v int64) *ExecuteDataExportShrinkRequest {
  s.Tid = &v
  return s
}

func (s *ExecuteDataExportShrinkRequest) Validate() error {
  return dara.Validate(s)
}

