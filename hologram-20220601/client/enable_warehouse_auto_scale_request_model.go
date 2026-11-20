// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWarehouseAutoScaleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetMaxClusterCount(v string) *EnableWarehouseAutoScaleRequest
  GetMaxClusterCount() *string 
  SetName(v string) *EnableWarehouseAutoScaleRequest
  GetName() *string 
}

type EnableWarehouseAutoScaleRequest struct {
  // The maximum number of clusters. The value must be greater than the number of reserved clusters and less than or equal to 4.
  // 
  // example:
  // 
  // 3
  MaxClusterCount *string `json:"maxClusterCount,omitempty" xml:"maxClusterCount,omitempty"`
  // The name of the compute group.
  // 
  // example:
  // 
  // my_warehouse
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s EnableWarehouseAutoScaleRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableWarehouseAutoScaleRequest) GoString() string {
  return s.String()
}

func (s *EnableWarehouseAutoScaleRequest) GetMaxClusterCount() *string  {
  return s.MaxClusterCount
}

func (s *EnableWarehouseAutoScaleRequest) GetName() *string  {
  return s.Name
}

func (s *EnableWarehouseAutoScaleRequest) SetMaxClusterCount(v string) *EnableWarehouseAutoScaleRequest {
  s.MaxClusterCount = &v
  return s
}

func (s *EnableWarehouseAutoScaleRequest) SetName(v string) *EnableWarehouseAutoScaleRequest {
  s.Name = &v
  return s
}

func (s *EnableWarehouseAutoScaleRequest) Validate() error {
  return dara.Validate(s)
}

