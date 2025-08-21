// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableKibanaPvlNetworkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetEndpointName(v string) *EnableKibanaPvlNetworkRequest
  GetEndpointName() *string 
  SetSecurityGroups(v []*string) *EnableKibanaPvlNetworkRequest
  GetSecurityGroups() []*string 
  SetVSwitchIdsZone(v []*EnableKibanaPvlNetworkRequestVSwitchIdsZone) *EnableKibanaPvlNetworkRequest
  GetVSwitchIdsZone() []*EnableKibanaPvlNetworkRequestVSwitchIdsZone 
  SetVpcId(v string) *EnableKibanaPvlNetworkRequest
  GetVpcId() *string 
}

type EnableKibanaPvlNetworkRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // es-cn-27a3mul6l000xxx-kibana-endpoint
  EndpointName *string `json:"endpointName,omitempty" xml:"endpointName,omitempty"`
  // This parameter is required.
  SecurityGroups []*string `json:"securityGroups,omitempty" xml:"securityGroups,omitempty" type:"Repeated"`
  // This parameter is required.
  VSwitchIdsZone []*EnableKibanaPvlNetworkRequestVSwitchIdsZone `json:"vSwitchIdsZone,omitempty" xml:"vSwitchIdsZone,omitempty" type:"Repeated"`
  // This parameter is required.
  // 
  // example:
  // 
  // vpc-xxx
  VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s EnableKibanaPvlNetworkRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableKibanaPvlNetworkRequest) GoString() string {
  return s.String()
}

func (s *EnableKibanaPvlNetworkRequest) GetEndpointName() *string  {
  return s.EndpointName
}

func (s *EnableKibanaPvlNetworkRequest) GetSecurityGroups() []*string  {
  return s.SecurityGroups
}

func (s *EnableKibanaPvlNetworkRequest) GetVSwitchIdsZone() []*EnableKibanaPvlNetworkRequestVSwitchIdsZone  {
  return s.VSwitchIdsZone
}

func (s *EnableKibanaPvlNetworkRequest) GetVpcId() *string  {
  return s.VpcId
}

func (s *EnableKibanaPvlNetworkRequest) SetEndpointName(v string) *EnableKibanaPvlNetworkRequest {
  s.EndpointName = &v
  return s
}

func (s *EnableKibanaPvlNetworkRequest) SetSecurityGroups(v []*string) *EnableKibanaPvlNetworkRequest {
  s.SecurityGroups = v
  return s
}

func (s *EnableKibanaPvlNetworkRequest) SetVSwitchIdsZone(v []*EnableKibanaPvlNetworkRequestVSwitchIdsZone) *EnableKibanaPvlNetworkRequest {
  s.VSwitchIdsZone = v
  return s
}

func (s *EnableKibanaPvlNetworkRequest) SetVpcId(v string) *EnableKibanaPvlNetworkRequest {
  s.VpcId = &v
  return s
}

func (s *EnableKibanaPvlNetworkRequest) Validate() error {
  return dara.Validate(s)
}

type EnableKibanaPvlNetworkRequestVSwitchIdsZone struct {
  // This parameter is required.
  // 
  // example:
  // 
  // vsw-xxxx
  VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou-h
  ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s EnableKibanaPvlNetworkRequestVSwitchIdsZone) String() string {
  return dara.Prettify(s)
}

func (s EnableKibanaPvlNetworkRequestVSwitchIdsZone) GoString() string {
  return s.String()
}

func (s *EnableKibanaPvlNetworkRequestVSwitchIdsZone) GetVswitchId() *string  {
  return s.VswitchId
}

func (s *EnableKibanaPvlNetworkRequestVSwitchIdsZone) GetZoneId() *string  {
  return s.ZoneId
}

func (s *EnableKibanaPvlNetworkRequestVSwitchIdsZone) SetVswitchId(v string) *EnableKibanaPvlNetworkRequestVSwitchIdsZone {
  s.VswitchId = &v
  return s
}

func (s *EnableKibanaPvlNetworkRequestVSwitchIdsZone) SetZoneId(v string) *EnableKibanaPvlNetworkRequestVSwitchIdsZone {
  s.ZoneId = &v
  return s
}

func (s *EnableKibanaPvlNetworkRequestVSwitchIdsZone) Validate() error {
  return dara.Validate(s)
}

