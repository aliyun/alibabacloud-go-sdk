// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpandGroupCapacityResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExpandGroupCapacityResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExpandGroupCapacityResponseBody
  GetSuccess() *bool 
  SetVendorRequestId(v string) *ExpandGroupCapacityResponseBody
  GetVendorRequestId() *string 
  SetVendorType(v string) *ExpandGroupCapacityResponseBody
  GetVendorType() *string 
}

type ExpandGroupCapacityResponseBody struct {
  // example:
  // 
  // 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  // example:
  // 
  // 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
  VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
  // example:
  // 
  // dingtalk
  VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ExpandGroupCapacityResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExpandGroupCapacityResponseBody) GoString() string {
  return s.String()
}

func (s *ExpandGroupCapacityResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExpandGroupCapacityResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExpandGroupCapacityResponseBody) GetVendorRequestId() *string  {
  return s.VendorRequestId
}

func (s *ExpandGroupCapacityResponseBody) GetVendorType() *string  {
  return s.VendorType
}

func (s *ExpandGroupCapacityResponseBody) SetRequestId(v string) *ExpandGroupCapacityResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExpandGroupCapacityResponseBody) SetSuccess(v bool) *ExpandGroupCapacityResponseBody {
  s.Success = &v
  return s
}

func (s *ExpandGroupCapacityResponseBody) SetVendorRequestId(v string) *ExpandGroupCapacityResponseBody {
  s.VendorRequestId = &v
  return s
}

func (s *ExpandGroupCapacityResponseBody) SetVendorType(v string) *ExpandGroupCapacityResponseBody {
  s.VendorType = &v
  return s
}

func (s *ExpandGroupCapacityResponseBody) Validate() error {
  return dara.Validate(s)
}

