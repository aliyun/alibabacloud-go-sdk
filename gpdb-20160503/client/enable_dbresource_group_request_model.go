// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBResourceGroupRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBInstanceId(v string) *EnableDBResourceGroupRequest
  GetDBInstanceId() *string 
  SetOwnerId(v int64) *EnableDBResourceGroupRequest
  GetOwnerId() *int64 
}

type EnableDBResourceGroupRequest struct {
  // The instance ID.
  // 
  // >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // gp-xxxxxxxxx
  DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s EnableDBResourceGroupRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableDBResourceGroupRequest) GoString() string {
  return s.String()
}

func (s *EnableDBResourceGroupRequest) GetDBInstanceId() *string  {
  return s.DBInstanceId
}

func (s *EnableDBResourceGroupRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableDBResourceGroupRequest) SetDBInstanceId(v string) *EnableDBResourceGroupRequest {
  s.DBInstanceId = &v
  return s
}

func (s *EnableDBResourceGroupRequest) SetOwnerId(v int64) *EnableDBResourceGroupRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableDBResourceGroupRequest) Validate() error {
  return dara.Validate(s)
}

