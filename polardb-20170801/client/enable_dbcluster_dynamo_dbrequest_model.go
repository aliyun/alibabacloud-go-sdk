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
  // The cluster ID.
  // 
  // > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query information about all clusters in the target region, including cluster IDs. Usage notes: Engine constraints that explicitly list the DBType and DBVersion ranges supporting the DynamoDB compatibility feature (such as Oracle 14+).
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

