// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndpointStatusDetail interface {
  dara.Model
  String() string
  GoString() string
  SetIpPortMapping(v map[string]*IpPort) *EndpointStatusDetail
  GetIpPortMapping() map[string]*IpPort 
}

type EndpointStatusDetail struct {
  // example:
  // 
  // { 	"slot-j6co2fcd": {   	"Ip": "10.0.0.2", 		"Port": 7001 	}, 	"slot-asdgys4d": {   	"Ip": "10.0.0.3", 		"Port": 7002 	} }
  IpPortMapping map[string]*IpPort `json:"IpPortMapping,omitempty" xml:"IpPortMapping,omitempty"`
}

func (s EndpointStatusDetail) String() string {
  return dara.Prettify(s)
}

func (s EndpointStatusDetail) GoString() string {
  return s.String()
}

func (s *EndpointStatusDetail) GetIpPortMapping() map[string]*IpPort  {
  return s.IpPortMapping
}

func (s *EndpointStatusDetail) SetIpPortMapping(v map[string]*IpPort) *EndpointStatusDetail {
  s.IpPortMapping = v
  return s
}

func (s *EndpointStatusDetail) Validate() error {
  return dara.Validate(s)
}

