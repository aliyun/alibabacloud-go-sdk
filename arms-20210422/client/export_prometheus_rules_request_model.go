// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportPrometheusRulesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClusterId(v string) *ExportPrometheusRulesRequest
  GetClusterId() *string 
  SetName(v string) *ExportPrometheusRulesRequest
  GetName() *string 
  SetNameSpace(v string) *ExportPrometheusRulesRequest
  GetNameSpace() *string 
  SetRegionId(v string) *ExportPrometheusRulesRequest
  GetRegionId() *string 
}

type ExportPrometheusRulesRequest struct {
  // This parameter is required.
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  // This parameter is required.
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // This parameter is required.
  NameSpace *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
  // This parameter is required.
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ExportPrometheusRulesRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportPrometheusRulesRequest) GoString() string {
  return s.String()
}

func (s *ExportPrometheusRulesRequest) GetClusterId() *string  {
  return s.ClusterId
}

func (s *ExportPrometheusRulesRequest) GetName() *string  {
  return s.Name
}

func (s *ExportPrometheusRulesRequest) GetNameSpace() *string  {
  return s.NameSpace
}

func (s *ExportPrometheusRulesRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExportPrometheusRulesRequest) SetClusterId(v string) *ExportPrometheusRulesRequest {
  s.ClusterId = &v
  return s
}

func (s *ExportPrometheusRulesRequest) SetName(v string) *ExportPrometheusRulesRequest {
  s.Name = &v
  return s
}

func (s *ExportPrometheusRulesRequest) SetNameSpace(v string) *ExportPrometheusRulesRequest {
  s.NameSpace = &v
  return s
}

func (s *ExportPrometheusRulesRequest) SetRegionId(v string) *ExportPrometheusRulesRequest {
  s.RegionId = &v
  return s
}

func (s *ExportPrometheusRulesRequest) Validate() error {
  return dara.Validate(s)
}

