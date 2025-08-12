// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSiteMonitorsRequest interface {
  dara.Model
  String() string
  GoString() string
  SetRegionId(v string) *EnableSiteMonitorsRequest
  GetRegionId() *string 
  SetTaskIds(v string) *EnableSiteMonitorsRequest
  GetTaskIds() *string 
}

type EnableSiteMonitorsRequest struct {
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The ID of the site monitoring task. Separate multiple IDs with commas (,).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 49f7b317-7645-4cc9-94fd-ea42e522****,49f7b317-7645-4cc9-94fd-ea42e522****
  TaskIds *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s EnableSiteMonitorsRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSiteMonitorsRequest) GoString() string {
  return s.String()
}

func (s *EnableSiteMonitorsRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableSiteMonitorsRequest) GetTaskIds() *string  {
  return s.TaskIds
}

func (s *EnableSiteMonitorsRequest) SetRegionId(v string) *EnableSiteMonitorsRequest {
  s.RegionId = &v
  return s
}

func (s *EnableSiteMonitorsRequest) SetTaskIds(v string) *EnableSiteMonitorsRequest {
  s.TaskIds = &v
  return s
}

func (s *EnableSiteMonitorsRequest) Validate() error {
  return dara.Validate(s)
}

