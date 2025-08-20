// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExistRunningSQLEngineRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBClusterId(v string) *ExistRunningSQLEngineRequest
  GetDBClusterId() *string 
  SetResourceGroupName(v string) *ExistRunningSQLEngineRequest
  GetResourceGroupName() *string 
}

type ExistRunningSQLEngineRequest struct {
  // The cluster ID.
  // 
  // >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // amv-bp1cit7z8j****
  DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
  // The name of the resource group.
  // 
  // >  You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/459446.html) operation to query the name of the resource group for a cluster.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // spark_test
  ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s ExistRunningSQLEngineRequest) String() string {
  return dara.Prettify(s)
}

func (s ExistRunningSQLEngineRequest) GoString() string {
  return s.String()
}

func (s *ExistRunningSQLEngineRequest) GetDBClusterId() *string  {
  return s.DBClusterId
}

func (s *ExistRunningSQLEngineRequest) GetResourceGroupName() *string  {
  return s.ResourceGroupName
}

func (s *ExistRunningSQLEngineRequest) SetDBClusterId(v string) *ExistRunningSQLEngineRequest {
  s.DBClusterId = &v
  return s
}

func (s *ExistRunningSQLEngineRequest) SetResourceGroupName(v string) *ExistRunningSQLEngineRequest {
  s.ResourceGroupName = &v
  return s
}

func (s *ExistRunningSQLEngineRequest) Validate() error {
  return dara.Validate(s)
}

