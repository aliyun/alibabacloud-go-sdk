// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBClusterCTableRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBClusterId(v string) *EnableDBClusterCTableRequest
  GetDBClusterId() *string 
}

type EnableDBClusterCTableRequest struct {
  // The ID of the cluster.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // pc-xxx
  DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s EnableDBClusterCTableRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableDBClusterCTableRequest) GoString() string {
  return s.String()
}

func (s *EnableDBClusterCTableRequest) GetDBClusterId() *string  {
  return s.DBClusterId
}

func (s *EnableDBClusterCTableRequest) SetDBClusterId(v string) *EnableDBClusterCTableRequest {
  s.DBClusterId = &v
  return s
}

func (s *EnableDBClusterCTableRequest) Validate() error {
  return dara.Validate(s)
}

