// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAlertTemplateRequest interface {
  dara.Model
  String() string
  GoString() string
  SetId(v int64) *EnableAlertTemplateRequest
  GetId() *int64 
  SetRegionId(v string) *EnableAlertTemplateRequest
  GetRegionId() *string 
}

type EnableAlertTemplateRequest struct {
  // This parameter is required.
  Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
  // This parameter is required.
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableAlertTemplateRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableAlertTemplateRequest) GoString() string {
  return s.String()
}

func (s *EnableAlertTemplateRequest) GetId() *int64  {
  return s.Id
}

func (s *EnableAlertTemplateRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableAlertTemplateRequest) SetId(v int64) *EnableAlertTemplateRequest {
  s.Id = &v
  return s
}

func (s *EnableAlertTemplateRequest) SetRegionId(v string) *EnableAlertTemplateRequest {
  s.RegionId = &v
  return s
}

func (s *EnableAlertTemplateRequest) Validate() error {
  return dara.Validate(s)
}

