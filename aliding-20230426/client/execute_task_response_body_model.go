// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTaskResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecuteTaskResponseBody
  GetRequestId() *string 
  SetVendorRequestId(v string) *ExecuteTaskResponseBody
  GetVendorRequestId() *string 
  SetVendorType(v string) *ExecuteTaskResponseBody
  GetVendorType() *string 
}

type ExecuteTaskResponseBody struct {
  // example:
  // 
  // 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
  VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
  // example:
  // 
  // dingtalk
  VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ExecuteTaskResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTaskResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteTaskResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteTaskResponseBody) GetVendorRequestId() *string  {
  return s.VendorRequestId
}

func (s *ExecuteTaskResponseBody) GetVendorType() *string  {
  return s.VendorType
}

func (s *ExecuteTaskResponseBody) SetRequestId(v string) *ExecuteTaskResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteTaskResponseBody) SetVendorRequestId(v string) *ExecuteTaskResponseBody {
  s.VendorRequestId = &v
  return s
}

func (s *ExecuteTaskResponseBody) SetVendorType(v string) *ExecuteTaskResponseBody {
  s.VendorType = &v
  return s
}

func (s *ExecuteTaskResponseBody) Validate() error {
  return dara.Validate(s)
}

