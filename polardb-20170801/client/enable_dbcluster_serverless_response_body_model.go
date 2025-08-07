// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBClusterServerlessResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetDBClusterId(v string) *EnableDBClusterServerlessResponseBody
  GetDBClusterId() *string 
  SetRequestId(v string) *EnableDBClusterServerlessResponseBody
  GetRequestId() *string 
}

type EnableDBClusterServerlessResponseBody struct {
  // The ID of the serverless cluster.
  // 
  // example:
  // 
  // pc-bp10gr51qasnl****
  DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 5E71541A-6007-4DCC-A38A-F872C31FEB45
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableDBClusterServerlessResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableDBClusterServerlessResponseBody) GoString() string {
  return s.String()
}

func (s *EnableDBClusterServerlessResponseBody) GetDBClusterId() *string  {
  return s.DBClusterId
}

func (s *EnableDBClusterServerlessResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableDBClusterServerlessResponseBody) SetDBClusterId(v string) *EnableDBClusterServerlessResponseBody {
  s.DBClusterId = &v
  return s
}

func (s *EnableDBClusterServerlessResponseBody) SetRequestId(v string) *EnableDBClusterServerlessResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableDBClusterServerlessResponseBody) Validate() error {
  return dara.Validate(s)
}

