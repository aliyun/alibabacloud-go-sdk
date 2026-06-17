// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBClusterDynamoDBRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBClusterId(v string) *EnableDBClusterDynamoDBRequest
  GetDBClusterId() *string 
}

type EnableDBClusterDynamoDBRequest struct {
  // The ID of the cluster.
  // 
  // > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to list all clusters in the destination region, including their IDs.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // pc-**************
  DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s EnableDBClusterDynamoDBRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableDBClusterDynamoDBRequest) GoString() string {
  return s.String()
}

func (s *EnableDBClusterDynamoDBRequest) GetDBClusterId() *string  {
  return s.DBClusterId
}

func (s *EnableDBClusterDynamoDBRequest) SetDBClusterId(v string) *EnableDBClusterDynamoDBRequest {
  s.DBClusterId = &v
  return s
}

func (s *EnableDBClusterDynamoDBRequest) Validate() error {
  return dara.Validate(s)
}

