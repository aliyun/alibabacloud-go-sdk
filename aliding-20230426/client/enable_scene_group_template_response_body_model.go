// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSceneGroupTemplateResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableSceneGroupTemplateResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableSceneGroupTemplateResponseBody
  GetSuccess() *bool 
  SetVendorRequestId(v string) *EnableSceneGroupTemplateResponseBody
  GetVendorRequestId() *string 
  SetVendorType(v string) *EnableSceneGroupTemplateResponseBody
  GetVendorType() *string 
}

type EnableSceneGroupTemplateResponseBody struct {
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

func (s EnableSceneGroupTemplateResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableSceneGroupTemplateResponseBody) GoString() string {
  return s.String()
}

func (s *EnableSceneGroupTemplateResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSceneGroupTemplateResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableSceneGroupTemplateResponseBody) GetVendorRequestId() *string  {
  return s.VendorRequestId
}

func (s *EnableSceneGroupTemplateResponseBody) GetVendorType() *string  {
  return s.VendorType
}

func (s *EnableSceneGroupTemplateResponseBody) SetRequestId(v string) *EnableSceneGroupTemplateResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableSceneGroupTemplateResponseBody) SetSuccess(v bool) *EnableSceneGroupTemplateResponseBody {
  s.Success = &v
  return s
}

func (s *EnableSceneGroupTemplateResponseBody) SetVendorRequestId(v string) *EnableSceneGroupTemplateResponseBody {
  s.VendorRequestId = &v
  return s
}

func (s *EnableSceneGroupTemplateResponseBody) SetVendorType(v string) *EnableSceneGroupTemplateResponseBody {
  s.VendorType = &v
  return s
}

func (s *EnableSceneGroupTemplateResponseBody) Validate() error {
  return dara.Validate(s)
}

