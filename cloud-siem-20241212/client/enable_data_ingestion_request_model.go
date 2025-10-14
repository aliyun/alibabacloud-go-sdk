// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDataIngestionRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDataIngestionId(v string) *EnableDataIngestionRequest
  GetDataIngestionId() *string 
  SetLang(v string) *EnableDataIngestionRequest
  GetLang() *string 
  SetProductId(v string) *EnableDataIngestionRequest
  GetProductId() *string 
  SetRegionId(v string) *EnableDataIngestionRequest
  GetRegionId() *string 
  SetRoleFor(v int64) *EnableDataIngestionRequest
  GetRoleFor() *int64 
}

type EnableDataIngestionRequest struct {
  // example:
  // 
  // alibaba_cloud_sas_netstat_ingestion_173326*******。
  DataIngestionId *string `json:"DataIngestionId,omitempty" xml:"DataIngestionId,omitempty"`
  // example:
  // 
  // zh。
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
  // example:
  // 
  // cn-hangzhou。
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // example:
  // 
  // 173326*******。
  RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s EnableDataIngestionRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableDataIngestionRequest) GoString() string {
  return s.String()
}

func (s *EnableDataIngestionRequest) GetDataIngestionId() *string  {
  return s.DataIngestionId
}

func (s *EnableDataIngestionRequest) GetLang() *string  {
  return s.Lang
}

func (s *EnableDataIngestionRequest) GetProductId() *string  {
  return s.ProductId
}

func (s *EnableDataIngestionRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableDataIngestionRequest) GetRoleFor() *int64  {
  return s.RoleFor
}

func (s *EnableDataIngestionRequest) SetDataIngestionId(v string) *EnableDataIngestionRequest {
  s.DataIngestionId = &v
  return s
}

func (s *EnableDataIngestionRequest) SetLang(v string) *EnableDataIngestionRequest {
  s.Lang = &v
  return s
}

func (s *EnableDataIngestionRequest) SetProductId(v string) *EnableDataIngestionRequest {
  s.ProductId = &v
  return s
}

func (s *EnableDataIngestionRequest) SetRegionId(v string) *EnableDataIngestionRequest {
  s.RegionId = &v
  return s
}

func (s *EnableDataIngestionRequest) SetRoleFor(v int64) *EnableDataIngestionRequest {
  s.RoleFor = &v
  return s
}

func (s *EnableDataIngestionRequest) Validate() error {
  return dara.Validate(s)
}

