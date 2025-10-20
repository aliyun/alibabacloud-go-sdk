// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBrandRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBrandId(v string) *EnableBrandRequest
  GetBrandId() *string 
  SetInstanceId(v string) *EnableBrandRequest
  GetInstanceId() *string 
}

type EnableBrandRequest struct {
  // 品牌化Id
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // brand_xxxx
  BrandId *string `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
  // IDaaS EIAM实例的ID。
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // idaas_ue2jvisn35ea5lmthk267xxxxx
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableBrandRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableBrandRequest) GoString() string {
  return s.String()
}

func (s *EnableBrandRequest) GetBrandId() *string  {
  return s.BrandId
}

func (s *EnableBrandRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableBrandRequest) SetBrandId(v string) *EnableBrandRequest {
  s.BrandId = &v
  return s
}

func (s *EnableBrandRequest) SetInstanceId(v string) *EnableBrandRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableBrandRequest) Validate() error {
  return dara.Validate(s)
}

