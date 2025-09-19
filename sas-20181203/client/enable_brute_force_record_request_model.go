// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBruteForceRecordRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBlockIp(v string) *EnableBruteForceRecordRequest
  GetBlockIp() *string 
  SetBound(v string) *EnableBruteForceRecordRequest
  GetBound() *string 
  SetId(v int64) *EnableBruteForceRecordRequest
  GetId() *int64 
  SetPort(v string) *EnableBruteForceRecordRequest
  GetPort() *string 
  SetResourceOwnerId(v int64) *EnableBruteForceRecordRequest
  GetResourceOwnerId() *int64 
  SetUuid(v string) *EnableBruteForceRecordRequest
  GetUuid() *string 
}

type EnableBruteForceRecordRequest struct {
  // The IP address that is specified in the policy.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 61.155.XX.XX
  BlockIp *string `json:"BlockIp,omitempty" xml:"BlockIp,omitempty"`
  // The traffic direction that is specified in the policy. Valid values:
  // 
  // 	- **in**: inbound
  // 
  // 	- **out**: outbound
  // 
  // example:
  // 
  // in
  Bound *string `json:"Bound,omitempty" xml:"Bound,omitempty"`
  // The ID of the policy that you want to enable.
  // 
  // > You can call the [DescribeBruteForceRecords](~~DescribeBruteForceRecords~~) operation to query the IDs of policies.
  // 
  // example:
  // 
  // 116602XX
  Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
  // The port number.
  // 
  // example:
  // 
  // 22/22
  Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The UUID of the server.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 71046acb-8bff-4c3b-9163-24deb007****
  Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s EnableBruteForceRecordRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableBruteForceRecordRequest) GoString() string {
  return s.String()
}

func (s *EnableBruteForceRecordRequest) GetBlockIp() *string  {
  return s.BlockIp
}

func (s *EnableBruteForceRecordRequest) GetBound() *string  {
  return s.Bound
}

func (s *EnableBruteForceRecordRequest) GetId() *int64  {
  return s.Id
}

func (s *EnableBruteForceRecordRequest) GetPort() *string  {
  return s.Port
}

func (s *EnableBruteForceRecordRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableBruteForceRecordRequest) GetUuid() *string  {
  return s.Uuid
}

func (s *EnableBruteForceRecordRequest) SetBlockIp(v string) *EnableBruteForceRecordRequest {
  s.BlockIp = &v
  return s
}

func (s *EnableBruteForceRecordRequest) SetBound(v string) *EnableBruteForceRecordRequest {
  s.Bound = &v
  return s
}

func (s *EnableBruteForceRecordRequest) SetId(v int64) *EnableBruteForceRecordRequest {
  s.Id = &v
  return s
}

func (s *EnableBruteForceRecordRequest) SetPort(v string) *EnableBruteForceRecordRequest {
  s.Port = &v
  return s
}

func (s *EnableBruteForceRecordRequest) SetResourceOwnerId(v int64) *EnableBruteForceRecordRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableBruteForceRecordRequest) SetUuid(v string) *EnableBruteForceRecordRequest {
  s.Uuid = &v
  return s
}

func (s *EnableBruteForceRecordRequest) Validate() error {
  return dara.Validate(s)
}

