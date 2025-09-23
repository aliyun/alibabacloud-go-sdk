// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBClusterOrcaResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetDBClusterId(v string) *EnableDBClusterOrcaResponseBody
  GetDBClusterId() *string 
  SetRequestId(v string) *EnableDBClusterOrcaResponseBody
  GetRequestId() *string 
}

type EnableDBClusterOrcaResponseBody struct {
  // example:
  // 
  // pc-***************
  DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 24A1990B-4F6E-482B-B8CB-75C612******
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableDBClusterOrcaResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableDBClusterOrcaResponseBody) GoString() string {
  return s.String()
}

func (s *EnableDBClusterOrcaResponseBody) GetDBClusterId() *string  {
  return s.DBClusterId
}

func (s *EnableDBClusterOrcaResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableDBClusterOrcaResponseBody) SetDBClusterId(v string) *EnableDBClusterOrcaResponseBody {
  s.DBClusterId = &v
  return s
}

func (s *EnableDBClusterOrcaResponseBody) SetRequestId(v string) *EnableDBClusterOrcaResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableDBClusterOrcaResponseBody) Validate() error {
  return dara.Validate(s)
}

