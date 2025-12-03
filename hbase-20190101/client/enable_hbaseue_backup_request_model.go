// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHBaseueBackupRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *EnableHBaseueBackupRequest
  GetClientToken() *string 
  SetColdStorageSize(v int32) *EnableHBaseueBackupRequest
  GetColdStorageSize() *int32 
  SetHbaseueClusterId(v string) *EnableHBaseueBackupRequest
  GetHbaseueClusterId() *string 
  SetNodeCount(v int32) *EnableHBaseueBackupRequest
  GetNodeCount() *int32 
}

type EnableHBaseueBackupRequest struct {
  // example:
  // 
  // xxx
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // example:
  // 
  // 800
  ColdStorageSize *int32 `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // ld-m5eznlga4k5bcxxxx
  HbaseueClusterId *string `json:"HbaseueClusterId,omitempty" xml:"HbaseueClusterId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 2
  NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
}

func (s EnableHBaseueBackupRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableHBaseueBackupRequest) GoString() string {
  return s.String()
}

func (s *EnableHBaseueBackupRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *EnableHBaseueBackupRequest) GetColdStorageSize() *int32  {
  return s.ColdStorageSize
}

func (s *EnableHBaseueBackupRequest) GetHbaseueClusterId() *string  {
  return s.HbaseueClusterId
}

func (s *EnableHBaseueBackupRequest) GetNodeCount() *int32  {
  return s.NodeCount
}

func (s *EnableHBaseueBackupRequest) SetClientToken(v string) *EnableHBaseueBackupRequest {
  s.ClientToken = &v
  return s
}

func (s *EnableHBaseueBackupRequest) SetColdStorageSize(v int32) *EnableHBaseueBackupRequest {
  s.ColdStorageSize = &v
  return s
}

func (s *EnableHBaseueBackupRequest) SetHbaseueClusterId(v string) *EnableHBaseueBackupRequest {
  s.HbaseueClusterId = &v
  return s
}

func (s *EnableHBaseueBackupRequest) SetNodeCount(v int32) *EnableHBaseueBackupRequest {
  s.NodeCount = &v
  return s
}

func (s *EnableHBaseueBackupRequest) Validate() error {
  return dara.Validate(s)
}

