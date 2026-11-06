// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBClusterDocumentDBRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBClusterId(v string) *EnableDBClusterDocumentDBRequest
  GetDBClusterId() *string 
}

type EnableDBClusterDocumentDBRequest struct {
  // The ID of the PolarDB cluster.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // pc-******************
  DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s EnableDBClusterDocumentDBRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableDBClusterDocumentDBRequest) GoString() string {
  return s.String()
}

func (s *EnableDBClusterDocumentDBRequest) GetDBClusterId() *string  {
  return s.DBClusterId
}

func (s *EnableDBClusterDocumentDBRequest) SetDBClusterId(v string) *EnableDBClusterDocumentDBRequest {
  s.DBClusterId = &v
  return s
}

func (s *EnableDBClusterDocumentDBRequest) Validate() error {
  return dara.Validate(s)
}

