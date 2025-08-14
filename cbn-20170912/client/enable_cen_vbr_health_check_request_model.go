// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableCenVbrHealthCheckRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCenId(v string) *EnableCenVbrHealthCheckRequest
  GetCenId() *string 
  SetDescription(v string) *EnableCenVbrHealthCheckRequest
  GetDescription() *string 
  SetHealthCheckInterval(v int32) *EnableCenVbrHealthCheckRequest
  GetHealthCheckInterval() *int32 
  SetHealthCheckOnly(v bool) *EnableCenVbrHealthCheckRequest
  GetHealthCheckOnly() *bool 
  SetHealthCheckSourceIp(v string) *EnableCenVbrHealthCheckRequest
  GetHealthCheckSourceIp() *string 
  SetHealthCheckTargetIp(v string) *EnableCenVbrHealthCheckRequest
  GetHealthCheckTargetIp() *string 
  SetHealthyThreshold(v int32) *EnableCenVbrHealthCheckRequest
  GetHealthyThreshold() *int32 
  SetOwnerAccount(v string) *EnableCenVbrHealthCheckRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EnableCenVbrHealthCheckRequest
  GetOwnerId() *int64 
  SetResourceOwnerAccount(v string) *EnableCenVbrHealthCheckRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EnableCenVbrHealthCheckRequest
  GetResourceOwnerId() *int64 
  SetVbrInstanceId(v string) *EnableCenVbrHealthCheckRequest
  GetVbrInstanceId() *string 
  SetVbrInstanceOwnerId(v int64) *EnableCenVbrHealthCheckRequest
  GetVbrInstanceOwnerId() *int64 
  SetVbrInstanceRegionId(v string) *EnableCenVbrHealthCheckRequest
  GetVbrInstanceRegionId() *string 
}

type EnableCenVbrHealthCheckRequest struct {
  // The ID of the Cloud Enterprise Network (CEN) instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cen-hahhfskfkseig****
  CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
  // The description of the health check.
  // 
  // The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
  // 
  // example:
  // 
  // testdesc
  Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
  // The time interval at which probe packets are sent during a health check. Unit: seconds. Default value: **2**. Valid values: **2 to 3**.
  // 
  // example:
  // 
  // 2
  HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
  // Specifies whether to enable probing during the health check. Valid values:
  // 
  // 	- **true**: yes
  // 
  //         If you enable probing, the system does not switch to another route if the detected route is not reachable. 
  // 
  //           Make sure that a redundant route is available. Otherwise, network disconnections may occur. 
  // 
  // 	- **false*	- (default): no
  // 
  //         Probing is disabled by default. If a redundant route is specified, the system switches to the redundant route if the detected route is not reachable.
  // 
  // example:
  // 
  // false
  HealthCheckOnly *bool `json:"HealthCheckOnly,omitempty" xml:"HealthCheckOnly,omitempty"`
  // The source IP address for the health check. You can set the source IP address in the following ways:
  // 
  // 	- **Automatic IP Address*	- (recommended): The system automatically assigns an IP address from the 100.96.0.0/16 CIDR block.
  // 
  // 	- **Custom IP Address**: You must specify an idle IP address from the 10.0.0.0/8, 192.168.0.0/16, or 172.16.0.0/12 CIDR block. The specified IP address cannot be the IP address of the VBR on the Alibaba Cloud side, the IP address of the VBR on the customer side, or other IP addresses with which the VBR communicates through the CEN instance.
  // 
  // example:
  // 
  // 192.XX.XX.1
  HealthCheckSourceIp *string `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
  // The destination IP address for the health check.
  // 
  // Set the destination IP address to the IP address of the VBR on the customer side.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 10.XX.XX.1
  HealthCheckTargetIp *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
  // The number of probe packets that are sent during a health check. Unit: packets. Valid values: **3 to 8**. Default value: **8**.
  // 
  // example:
  // 
  // 8
  HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The ID of the VBR.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // vbr-wz95o9aylj181n5mzk****
  VbrInstanceId *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
  // The ID of the Alibaba Cloud account to which the VBR belongs.
  // 
  // > This parameter is required if the VBR and the CEN instance belong to different Alibaba Cloud accounts.
  // 
  // example:
  // 
  // 1250123456123456
  VbrInstanceOwnerId *int64 `json:"VbrInstanceOwnerId,omitempty" xml:"VbrInstanceOwnerId,omitempty"`
  // The ID of the region where the VBR is deployed.
  // 
  // You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-shenzhen
  VbrInstanceRegionId *string `json:"VbrInstanceRegionId,omitempty" xml:"VbrInstanceRegionId,omitempty"`
}

func (s EnableCenVbrHealthCheckRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableCenVbrHealthCheckRequest) GoString() string {
  return s.String()
}

func (s *EnableCenVbrHealthCheckRequest) GetCenId() *string  {
  return s.CenId
}

func (s *EnableCenVbrHealthCheckRequest) GetDescription() *string  {
  return s.Description
}

func (s *EnableCenVbrHealthCheckRequest) GetHealthCheckInterval() *int32  {
  return s.HealthCheckInterval
}

func (s *EnableCenVbrHealthCheckRequest) GetHealthCheckOnly() *bool  {
  return s.HealthCheckOnly
}

func (s *EnableCenVbrHealthCheckRequest) GetHealthCheckSourceIp() *string  {
  return s.HealthCheckSourceIp
}

func (s *EnableCenVbrHealthCheckRequest) GetHealthCheckTargetIp() *string  {
  return s.HealthCheckTargetIp
}

func (s *EnableCenVbrHealthCheckRequest) GetHealthyThreshold() *int32  {
  return s.HealthyThreshold
}

func (s *EnableCenVbrHealthCheckRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EnableCenVbrHealthCheckRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EnableCenVbrHealthCheckRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EnableCenVbrHealthCheckRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnableCenVbrHealthCheckRequest) GetVbrInstanceId() *string  {
  return s.VbrInstanceId
}

func (s *EnableCenVbrHealthCheckRequest) GetVbrInstanceOwnerId() *int64  {
  return s.VbrInstanceOwnerId
}

func (s *EnableCenVbrHealthCheckRequest) GetVbrInstanceRegionId() *string  {
  return s.VbrInstanceRegionId
}

func (s *EnableCenVbrHealthCheckRequest) SetCenId(v string) *EnableCenVbrHealthCheckRequest {
  s.CenId = &v
  return s
}

func (s *EnableCenVbrHealthCheckRequest) SetDescription(v string) *EnableCenVbrHealthCheckRequest {
  s.Description = &v
  return s
}

func (s *EnableCenVbrHealthCheckRequest) SetHealthCheckInterval(v int32) *EnableCenVbrHealthCheckRequest {
  s.HealthCheckInterval = &v
  return s
}

func (s *EnableCenVbrHealthCheckRequest) SetHealthCheckOnly(v bool) *EnableCenVbrHealthCheckRequest {
  s.HealthCheckOnly = &v
  return s
}

func (s *EnableCenVbrHealthCheckRequest) SetHealthCheckSourceIp(v string) *EnableCenVbrHealthCheckRequest {
  s.HealthCheckSourceIp = &v
  return s
}

func (s *EnableCenVbrHealthCheckRequest) SetHealthCheckTargetIp(v string) *EnableCenVbrHealthCheckRequest {
  s.HealthCheckTargetIp = &v
  return s
}

func (s *EnableCenVbrHealthCheckRequest) SetHealthyThreshold(v int32) *EnableCenVbrHealthCheckRequest {
  s.HealthyThreshold = &v
  return s
}

func (s *EnableCenVbrHealthCheckRequest) SetOwnerAccount(v string) *EnableCenVbrHealthCheckRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EnableCenVbrHealthCheckRequest) SetOwnerId(v int64) *EnableCenVbrHealthCheckRequest {
  s.OwnerId = &v
  return s
}

func (s *EnableCenVbrHealthCheckRequest) SetResourceOwnerAccount(v string) *EnableCenVbrHealthCheckRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EnableCenVbrHealthCheckRequest) SetResourceOwnerId(v int64) *EnableCenVbrHealthCheckRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnableCenVbrHealthCheckRequest) SetVbrInstanceId(v string) *EnableCenVbrHealthCheckRequest {
  s.VbrInstanceId = &v
  return s
}

func (s *EnableCenVbrHealthCheckRequest) SetVbrInstanceOwnerId(v int64) *EnableCenVbrHealthCheckRequest {
  s.VbrInstanceOwnerId = &v
  return s
}

func (s *EnableCenVbrHealthCheckRequest) SetVbrInstanceRegionId(v string) *EnableCenVbrHealthCheckRequest {
  s.VbrInstanceRegionId = &v
  return s
}

func (s *EnableCenVbrHealthCheckRequest) Validate() error {
  return dara.Validate(s)
}

