// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDBClusterServerlessRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBClusterId(v string) *EnableDBClusterServerlessRequest
  GetDBClusterId() *string 
  SetOwnerAccount(v string) *EnableDBClusterServerlessRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnableDBClusterServerlessRequest
  GetOwnerId() *int64 
  SetResourceOwnerAccount(v string) *EnableDBClusterServerlessRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableDBClusterServerlessRequest
  GetResourceOwnerId() *int64 
  SetScaleApRoNumMax(v string) *EnableDBClusterServerlessRequest
  GetScaleApRoNumMax() *string 
  SetScaleApRoNumMin(v string) *EnableDBClusterServerlessRequest
  GetScaleApRoNumMin() *string 
  SetScaleMax(v string) *EnableDBClusterServerlessRequest
  GetScaleMax() *string 
  SetScaleMin(v string) *EnableDBClusterServerlessRequest
  GetScaleMin() *string 
  SetScaleRoNumMax(v string) *EnableDBClusterServerlessRequest
  GetScaleRoNumMax() *string 
  SetScaleRoNumMin(v string) *EnableDBClusterServerlessRequest
  GetScaleRoNumMin() *string 
}

type EnableDBClusterServerlessRequest struct {
  // The cluster ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // pc-**************
  DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The maximum number of stable AP read-only nodes. Valid values: 0 to 7.
  // 
  // example:
  // 
  // 1
  ScaleApRoNumMax *string `json:"ScaleApRoNumMax,omitempty" xml:"ScaleApRoNumMax,omitempty"`
  // The minimum number of stable AP read-only nodes. Valid values: 0 to 7.
  // 
  // example:
  // 
  // 1
  ScaleApRoNumMin *string `json:"ScaleApRoNumMin,omitempty" xml:"ScaleApRoNumMin,omitempty"`
  // The maximum number of PCUs per node for scaling. Valid values: 1 to 8 PCUs.
  // 
  // example:
  // 
  // 2
  ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
  // The minimum number of PolarDB capacity units (PCUs) per node for scaling. Valid values: 1 to 8 PCUs.
  // 
  // example:
  // 
  // 1
  ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
  // The maximum number of read-only nodes for scaling. Valid values: 0 to 7.
  // 
  // example:
  // 
  // 2
  ScaleRoNumMax *string `json:"ScaleRoNumMax,omitempty" xml:"ScaleRoNumMax,omitempty"`
  // The minimum number of read-only nodes for scaling. Valid values: 0 to 7.
  // 
  // example:
  // 
  // 1
  ScaleRoNumMin *string `json:"ScaleRoNumMin,omitempty" xml:"ScaleRoNumMin,omitempty"`
}

func (s EnableDBClusterServerlessRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableDBClusterServerlessRequest) GoString() string {
  return s.String()
}

func (s *EnableDBClusterServerlessRequest) GetDBClusterId() *string  {
  return s.DBClusterId
}

func (s *EnableDBClusterServerlessRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnableDBClusterServerlessRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableDBClusterServerlessRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableDBClusterServerlessRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableDBClusterServerlessRequest) GetScaleApRoNumMax() *string  {
  return s.ScaleApRoNumMax
}

func (s *EnableDBClusterServerlessRequest) GetScaleApRoNumMin() *string  {
  return s.ScaleApRoNumMin
}

func (s *EnableDBClusterServerlessRequest) GetScaleMax() *string  {
  return s.ScaleMax
}

func (s *EnableDBClusterServerlessRequest) GetScaleMin() *string  {
  return s.ScaleMin
}

func (s *EnableDBClusterServerlessRequest) GetScaleRoNumMax() *string  {
  return s.ScaleRoNumMax
}

func (s *EnableDBClusterServerlessRequest) GetScaleRoNumMin() *string  {
  return s.ScaleRoNumMin
}

func (s *EnableDBClusterServerlessRequest) SetDBClusterId(v string) *EnableDBClusterServerlessRequest {
  s.DBClusterId = &v
  return s
}

func (s *EnableDBClusterServerlessRequest) SetOwnerAccount(v string) *EnableDBClusterServerlessRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnableDBClusterServerlessRequest) SetOwnerId(v int64) *EnableDBClusterServerlessRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableDBClusterServerlessRequest) SetResourceOwnerAccount(v string) *EnableDBClusterServerlessRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableDBClusterServerlessRequest) SetResourceOwnerId(v int64) *EnableDBClusterServerlessRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableDBClusterServerlessRequest) SetScaleApRoNumMax(v string) *EnableDBClusterServerlessRequest {
  s.ScaleApRoNumMax = &v
  return s
}

func (s *EnableDBClusterServerlessRequest) SetScaleApRoNumMin(v string) *EnableDBClusterServerlessRequest {
  s.ScaleApRoNumMin = &v
  return s
}

func (s *EnableDBClusterServerlessRequest) SetScaleMax(v string) *EnableDBClusterServerlessRequest {
  s.ScaleMax = &v
  return s
}

func (s *EnableDBClusterServerlessRequest) SetScaleMin(v string) *EnableDBClusterServerlessRequest {
  s.ScaleMin = &v
  return s
}

func (s *EnableDBClusterServerlessRequest) SetScaleRoNumMax(v string) *EnableDBClusterServerlessRequest {
  s.ScaleRoNumMax = &v
  return s
}

func (s *EnableDBClusterServerlessRequest) SetScaleRoNumMin(v string) *EnableDBClusterServerlessRequest {
  s.ScaleRoNumMin = &v
  return s
}

func (s *EnableDBClusterServerlessRequest) Validate() error {
  return dara.Validate(s)
}

