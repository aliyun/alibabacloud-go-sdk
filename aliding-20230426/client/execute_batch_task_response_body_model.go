// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteBatchTaskResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetFailNumber(v int32) *ExecuteBatchTaskResponseBody
  GetFailNumber() *int32 
  SetRequestId(v string) *ExecuteBatchTaskResponseBody
  GetRequestId() *string 
  SetSuccessNumber(v int32) *ExecuteBatchTaskResponseBody
  GetSuccessNumber() *int32 
  SetTotal(v int32) *ExecuteBatchTaskResponseBody
  GetTotal() *int32 
  SetVendorRequestId(v string) *ExecuteBatchTaskResponseBody
  GetVendorRequestId() *string 
  SetVendorType(v string) *ExecuteBatchTaskResponseBody
  GetVendorType() *string 
}

type ExecuteBatchTaskResponseBody struct {
  // example:
  // 
  // 1
  FailNumber *int32 `json:"failNumber,omitempty" xml:"failNumber,omitempty"`
  // example:
  // 
  // 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // 2
  SuccessNumber *int32 `json:"successNumber,omitempty" xml:"successNumber,omitempty"`
  // example:
  // 
  // 3
  Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
  // example:
  // 
  // 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
  VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
  // example:
  // 
  // dingtalk
  VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ExecuteBatchTaskResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteBatchTaskResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteBatchTaskResponseBody) GetFailNumber() *int32  {
  return s.FailNumber
}

func (s *ExecuteBatchTaskResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteBatchTaskResponseBody) GetSuccessNumber() *int32  {
  return s.SuccessNumber
}

func (s *ExecuteBatchTaskResponseBody) GetTotal() *int32  {
  return s.Total
}

func (s *ExecuteBatchTaskResponseBody) GetVendorRequestId() *string  {
  return s.VendorRequestId
}

func (s *ExecuteBatchTaskResponseBody) GetVendorType() *string  {
  return s.VendorType
}

func (s *ExecuteBatchTaskResponseBody) SetFailNumber(v int32) *ExecuteBatchTaskResponseBody {
  s.FailNumber = &v
  return s
}

func (s *ExecuteBatchTaskResponseBody) SetRequestId(v string) *ExecuteBatchTaskResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteBatchTaskResponseBody) SetSuccessNumber(v int32) *ExecuteBatchTaskResponseBody {
  s.SuccessNumber = &v
  return s
}

func (s *ExecuteBatchTaskResponseBody) SetTotal(v int32) *ExecuteBatchTaskResponseBody {
  s.Total = &v
  return s
}

func (s *ExecuteBatchTaskResponseBody) SetVendorRequestId(v string) *ExecuteBatchTaskResponseBody {
  s.VendorRequestId = &v
  return s
}

func (s *ExecuteBatchTaskResponseBody) SetVendorType(v string) *ExecuteBatchTaskResponseBody {
  s.VendorType = &v
  return s
}

func (s *ExecuteBatchTaskResponseBody) Validate() error {
  return dara.Validate(s)
}

