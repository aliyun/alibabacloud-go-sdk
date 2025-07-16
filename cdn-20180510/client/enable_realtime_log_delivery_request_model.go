// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableRealtimeLogDeliveryRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDomain(v string) *EnableRealtimeLogDeliveryRequest
  GetDomain() *string 
  SetLogstore(v string) *EnableRealtimeLogDeliveryRequest
  GetLogstore() *string 
  SetProject(v string) *EnableRealtimeLogDeliveryRequest
  GetProject() *string 
  SetRegion(v string) *EnableRealtimeLogDeliveryRequest
  GetRegion() *string 
}

type EnableRealtimeLogDeliveryRequest struct {
  // The accelerated domain name for which you want to enable real-time log delivery. You can specify multiple domain names and separate them with commas (,).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // example.com
  Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
  // The name of the Logstore where log entries are stored.
  // 
  // example:
  // 
  // LogstoreName
  Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
  // The name of the Log Service project that is used for real-time log delivery.
  // 
  // example:
  // 
  // ProjectName
  Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
  // The ID of the region where the Log Service project is deployed.
  // 
  // example:
  // 
  // cn-shanghai
  Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s EnableRealtimeLogDeliveryRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableRealtimeLogDeliveryRequest) GoString() string {
  return s.String()
}

func (s *EnableRealtimeLogDeliveryRequest) GetDomain() *string  {
  return s.Domain
}

func (s *EnableRealtimeLogDeliveryRequest) GetLogstore() *string  {
  return s.Logstore
}

func (s *EnableRealtimeLogDeliveryRequest) GetProject() *string  {
  return s.Project
}

func (s *EnableRealtimeLogDeliveryRequest) GetRegion() *string  {
  return s.Region
}

func (s *EnableRealtimeLogDeliveryRequest) SetDomain(v string) *EnableRealtimeLogDeliveryRequest {
  s.Domain = &v
  return s
}

func (s *EnableRealtimeLogDeliveryRequest) SetLogstore(v string) *EnableRealtimeLogDeliveryRequest {
  s.Logstore = &v
  return s
}

func (s *EnableRealtimeLogDeliveryRequest) SetProject(v string) *EnableRealtimeLogDeliveryRequest {
  s.Project = &v
  return s
}

func (s *EnableRealtimeLogDeliveryRequest) SetRegion(v string) *EnableRealtimeLogDeliveryRequest {
  s.Region = &v
  return s
}

func (s *EnableRealtimeLogDeliveryRequest) Validate() error {
  return dara.Validate(s)
}

