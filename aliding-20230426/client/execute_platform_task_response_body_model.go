// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecutePlatformTaskResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecutePlatformTaskResponseBody
  GetRequestId() *string 
  SetVendorRequestId(v string) *ExecutePlatformTaskResponseBody
  GetVendorRequestId() *string 
  SetVendorType(v string) *ExecutePlatformTaskResponseBody
  GetVendorType() *string 
}

type ExecutePlatformTaskResponseBody struct {
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

func (s ExecutePlatformTaskResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecutePlatformTaskResponseBody) GoString() string {
  return s.String()
}

func (s *ExecutePlatformTaskResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecutePlatformTaskResponseBody) GetVendorRequestId() *string  {
  return s.VendorRequestId
}

func (s *ExecutePlatformTaskResponseBody) GetVendorType() *string  {
  return s.VendorType
}

func (s *ExecutePlatformTaskResponseBody) SetRequestId(v string) *ExecutePlatformTaskResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecutePlatformTaskResponseBody) SetVendorRequestId(v string) *ExecutePlatformTaskResponseBody {
  s.VendorRequestId = &v
  return s
}

func (s *ExecutePlatformTaskResponseBody) SetVendorType(v string) *ExecutePlatformTaskResponseBody {
  s.VendorType = &v
  return s
}

func (s *ExecutePlatformTaskResponseBody) Validate() error {
  return dara.Validate(s)
}

