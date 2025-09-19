// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCustomInstanceBlockRecordRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBlockIp(v string) *EnableCustomInstanceBlockRecordRequest
  GetBlockIp() *string 
  SetBound(v string) *EnableCustomInstanceBlockRecordRequest
  GetBound() *string 
  SetResourceOwnerId(v int64) *EnableCustomInstanceBlockRecordRequest
  GetResourceOwnerId() *int64 
  SetUuid(v string) *EnableCustomInstanceBlockRecordRequest
  GetUuid() *string 
}

type EnableCustomInstanceBlockRecordRequest struct {
  // The IP address that you want to block.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 192.168.xx.xx
  BlockIp *string `json:"BlockIp,omitempty" xml:"BlockIp,omitempty"`
  // The traffic direction from the IP address that you want to block. Valid value:
  // 
  // 	- **in**
  // 
  // 	- **out**
  // 
  // example:
  // 
  // in
  Bound *string `json:"Bound,omitempty" xml:"Bound,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The UUID of the server.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 50d213b4-3a35-427a-b8a5-04b0c7e1****
  Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s EnableCustomInstanceBlockRecordRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableCustomInstanceBlockRecordRequest) GoString() string {
  return s.String()
}

func (s *EnableCustomInstanceBlockRecordRequest) GetBlockIp() *string  {
  return s.BlockIp
}

func (s *EnableCustomInstanceBlockRecordRequest) GetBound() *string  {
  return s.Bound
}

func (s *EnableCustomInstanceBlockRecordRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableCustomInstanceBlockRecordRequest) GetUuid() *string  {
  return s.Uuid
}

func (s *EnableCustomInstanceBlockRecordRequest) SetBlockIp(v string) *EnableCustomInstanceBlockRecordRequest {
  s.BlockIp = &v
  return s
}

func (s *EnableCustomInstanceBlockRecordRequest) SetBound(v string) *EnableCustomInstanceBlockRecordRequest {
  s.Bound = &v
  return s
}

func (s *EnableCustomInstanceBlockRecordRequest) SetResourceOwnerId(v int64) *EnableCustomInstanceBlockRecordRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableCustomInstanceBlockRecordRequest) SetUuid(v string) *EnableCustomInstanceBlockRecordRequest {
  s.Uuid = &v
  return s
}

func (s *EnableCustomInstanceBlockRecordRequest) Validate() error {
  return dara.Validate(s)
}

