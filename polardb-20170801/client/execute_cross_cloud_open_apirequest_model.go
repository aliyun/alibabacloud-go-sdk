// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteCrossCloudOpenAPIRequest interface {
  dara.Model
  String() string
  GoString() string
  SetProxyInfo(v string) *ExecuteCrossCloudOpenAPIRequest
  GetProxyInfo() *string 
}

type ExecuteCrossCloudOpenAPIRequest struct {
  // example:
  // 
  // {"Action":"DescribeDBClusters"}
  ProxyInfo *string `json:"ProxyInfo,omitempty" xml:"ProxyInfo,omitempty"`
}

func (s ExecuteCrossCloudOpenAPIRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteCrossCloudOpenAPIRequest) GoString() string {
  return s.String()
}

func (s *ExecuteCrossCloudOpenAPIRequest) GetProxyInfo() *string  {
  return s.ProxyInfo
}

func (s *ExecuteCrossCloudOpenAPIRequest) SetProxyInfo(v string) *ExecuteCrossCloudOpenAPIRequest {
  s.ProxyInfo = &v
  return s
}

func (s *ExecuteCrossCloudOpenAPIRequest) Validate() error {
  return dara.Validate(s)
}

