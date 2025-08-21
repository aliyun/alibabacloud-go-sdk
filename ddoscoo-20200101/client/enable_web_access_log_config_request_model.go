// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWebAccessLogConfigRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDomain(v string) *EnableWebAccessLogConfigRequest
  GetDomain() *string 
  SetResourceGroupId(v string) *EnableWebAccessLogConfigRequest
  GetResourceGroupId() *string 
}

type EnableWebAccessLogConfigRequest struct {
  // The domain name of the website.
  // 
  // >  A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/474212.html) operation to query all domain names.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // www.aliyun.com
  Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
  // The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
  // 
  // example:
  // 
  // default
  ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s EnableWebAccessLogConfigRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableWebAccessLogConfigRequest) GoString() string {
  return s.String()
}

func (s *EnableWebAccessLogConfigRequest) GetDomain() *string  {
  return s.Domain
}

func (s *EnableWebAccessLogConfigRequest) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *EnableWebAccessLogConfigRequest) SetDomain(v string) *EnableWebAccessLogConfigRequest {
  s.Domain = &v
  return s
}

func (s *EnableWebAccessLogConfigRequest) SetResourceGroupId(v string) *EnableWebAccessLogConfigRequest {
  s.ResourceGroupId = &v
  return s
}

func (s *EnableWebAccessLogConfigRequest) Validate() error {
  return dara.Validate(s)
}

