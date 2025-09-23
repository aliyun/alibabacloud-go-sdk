// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBClusterOrcaRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBClusterId(v string) *EnableDBClusterOrcaRequest
  GetDBClusterId() *string 
  SetOwnerAccount(v string) *EnableDBClusterOrcaRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnableDBClusterOrcaRequest
  GetOwnerId() *int64 
  SetResourceOwnerAccount(v string) *EnableDBClusterOrcaRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableDBClusterOrcaRequest
  GetResourceOwnerId() *int64 
}

type EnableDBClusterOrcaRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // pc-************
  DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s EnableDBClusterOrcaRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableDBClusterOrcaRequest) GoString() string {
  return s.String()
}

func (s *EnableDBClusterOrcaRequest) GetDBClusterId() *string  {
  return s.DBClusterId
}

func (s *EnableDBClusterOrcaRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnableDBClusterOrcaRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableDBClusterOrcaRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableDBClusterOrcaRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableDBClusterOrcaRequest) SetDBClusterId(v string) *EnableDBClusterOrcaRequest {
  s.DBClusterId = &v
  return s
}

func (s *EnableDBClusterOrcaRequest) SetOwnerAccount(v string) *EnableDBClusterOrcaRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnableDBClusterOrcaRequest) SetOwnerId(v int64) *EnableDBClusterOrcaRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableDBClusterOrcaRequest) SetResourceOwnerAccount(v string) *EnableDBClusterOrcaRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableDBClusterOrcaRequest) SetResourceOwnerId(v int64) *EnableDBClusterOrcaRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableDBClusterOrcaRequest) Validate() error {
  return dara.Validate(s)
}

