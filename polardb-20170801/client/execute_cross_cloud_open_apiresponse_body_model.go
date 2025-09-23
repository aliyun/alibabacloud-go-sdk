// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteCrossCloudOpenAPIResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetProxyData(v string) *ExecuteCrossCloudOpenAPIResponseBody
  GetProxyData() *string 
  SetRequestId(v string) *ExecuteCrossCloudOpenAPIResponseBody
  GetRequestId() *string 
}

type ExecuteCrossCloudOpenAPIResponseBody struct {
  // example:
  // 
  // {"ProxyData": "{\\"DBCluster\\":[{\\"AliyunRegionId\\":\\"cn-beijing\\",\\"CloudProvider\\":\\"huawei\\",\\"CreateTime\\":\\"2024-11-25T14:49:10Z\\",\\"CrossCloudRegionId\\":\\"cn-east-3\\",\\"DBClusterDescription\\":\\"\\",\\"DBClusterId\\":\\"pc-2zej3qvf5fg******\\",\\"DBClusterStatus\\":\\"Creating\\",\\"DBType\\":\\"polardb_mysql\\",\\"DBVersion\\":\\"8.0\\",\\"ProjectId\\":\\"pj-bp1m8oh1k68******\\"},{\\"AliyunRegionId\\":\\"cn-beijing\\",\\"CloudProvider\\":\\"huawei\\",\\"CreateTime\\":\\"2024-11-25T14:59:10Z\\",\\"CrossCloudRegionId\\":\\"cn-east-3\\",\\"DBClusterDescription\\":\\"\\",\\"DBClusterId\\":\\"pc-2ze29994l17******\\",\\"DBClusterStatus\\":\\"Running\\",\\"DBType\\":\\"polardb_mysql\\",\\"DBVersion\\":\\"8.0\\",\\"ProjectId\\":\\"pj-bp1m8oh1k68******\\"}]}","RequestId": "E56531A4-E552-40BA-9C58-137B80******"}
  ProxyData *string `json:"ProxyData,omitempty" xml:"ProxyData,omitempty"`
  // example:
  // 
  // E56531A4-E552-40BA-9C58-137B80******
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteCrossCloudOpenAPIResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteCrossCloudOpenAPIResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteCrossCloudOpenAPIResponseBody) GetProxyData() *string  {
  return s.ProxyData
}

func (s *ExecuteCrossCloudOpenAPIResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteCrossCloudOpenAPIResponseBody) SetProxyData(v string) *ExecuteCrossCloudOpenAPIResponseBody {
  s.ProxyData = &v
  return s
}

func (s *ExecuteCrossCloudOpenAPIResponseBody) SetRequestId(v string) *ExecuteCrossCloudOpenAPIResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteCrossCloudOpenAPIResponseBody) Validate() error {
  return dara.Validate(s)
}

