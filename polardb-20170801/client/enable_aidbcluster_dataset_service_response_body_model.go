// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAIDBClusterDatasetServiceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetDBClusterId(v string) *EnableAIDBClusterDatasetServiceResponseBody
  GetDBClusterId() *string 
  SetRequestId(v string) *EnableAIDBClusterDatasetServiceResponseBody
  GetRequestId() *string 
}

type EnableAIDBClusterDatasetServiceResponseBody struct {
  // The instance ID of the dataset management service.
  // 
  // example:
  // 
  // pcs-2ze26***
  DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // C890995A-CF06-4F4D-8DB8-DD26C2******
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAIDBClusterDatasetServiceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableAIDBClusterDatasetServiceResponseBody) GoString() string {
  return s.String()
}

func (s *EnableAIDBClusterDatasetServiceResponseBody) GetDBClusterId() *string  {
  return s.DBClusterId
}

func (s *EnableAIDBClusterDatasetServiceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableAIDBClusterDatasetServiceResponseBody) SetDBClusterId(v string) *EnableAIDBClusterDatasetServiceResponseBody {
  s.DBClusterId = &v
  return s
}

func (s *EnableAIDBClusterDatasetServiceResponseBody) SetRequestId(v string) *EnableAIDBClusterDatasetServiceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableAIDBClusterDatasetServiceResponseBody) Validate() error {
  return dara.Validate(s)
}

