// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAdviceServiceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBClusterId(v string) *EnableAdviceServiceRequest
  GetDBClusterId() *string 
  SetRegionId(v string) *EnableAdviceServiceRequest
  GetRegionId() *string 
}

type EnableAdviceServiceRequest struct {
  // The cluster ID.
  // 
  // >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // am-bp1q10xxzq2z4****
  DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
  // The region ID of the cluster.
  // 
  // >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-shanghai
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableAdviceServiceRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableAdviceServiceRequest) GoString() string {
  return s.String()
}

func (s *EnableAdviceServiceRequest) GetDBClusterId() *string  {
  return s.DBClusterId
}

func (s *EnableAdviceServiceRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableAdviceServiceRequest) SetDBClusterId(v string) *EnableAdviceServiceRequest {
  s.DBClusterId = &v
  return s
}

func (s *EnableAdviceServiceRequest) SetRegionId(v string) *EnableAdviceServiceRequest {
  s.RegionId = &v
  return s
}

func (s *EnableAdviceServiceRequest) Validate() error {
  return dara.Validate(s)
}

