// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCustomBlockRecordRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBlockIp(v string) *EnableCustomBlockRecordRequest
  GetBlockIp() *string 
  SetBound(v string) *EnableCustomBlockRecordRequest
  GetBound() *string 
  SetResourceOwnerId(v int64) *EnableCustomBlockRecordRequest
  GetResourceOwnerId() *int64 
}

type EnableCustomBlockRecordRequest struct {
  // The IP address that is specified in the policy.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 43.248.XX.XX
  BlockIp *string `json:"BlockIp,omitempty" xml:"BlockIp,omitempty"`
  // The traffic direction that is specified in the policy. Valid values:
  // 
  // 	- **in**: inbound
  // 
  // 	- **out**: outbound
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // in
  Bound *string `json:"Bound,omitempty" xml:"Bound,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s EnableCustomBlockRecordRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableCustomBlockRecordRequest) GoString() string {
  return s.String()
}

func (s *EnableCustomBlockRecordRequest) GetBlockIp() *string  {
  return s.BlockIp
}

func (s *EnableCustomBlockRecordRequest) GetBound() *string  {
  return s.Bound
}

func (s *EnableCustomBlockRecordRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableCustomBlockRecordRequest) SetBlockIp(v string) *EnableCustomBlockRecordRequest {
  s.BlockIp = &v
  return s
}

func (s *EnableCustomBlockRecordRequest) SetBound(v string) *EnableCustomBlockRecordRequest {
  s.Bound = &v
  return s
}

func (s *EnableCustomBlockRecordRequest) SetResourceOwnerId(v int64) *EnableCustomBlockRecordRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableCustomBlockRecordRequest) Validate() error {
  return dara.Validate(s)
}

