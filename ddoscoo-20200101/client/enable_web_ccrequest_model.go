// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWebCCRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDomain(v string) *EnableWebCCRequest
  GetDomain() *string 
  SetResourceGroupId(v string) *EnableWebCCRequest
  GetResourceGroupId() *string 
}

type EnableWebCCRequest struct {
  // The domain name of the website.
  // 
  // > A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
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

func (s EnableWebCCRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableWebCCRequest) GoString() string {
  return s.String()
}

func (s *EnableWebCCRequest) GetDomain() *string  {
  return s.Domain
}

func (s *EnableWebCCRequest) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *EnableWebCCRequest) SetDomain(v string) *EnableWebCCRequest {
  s.Domain = &v
  return s
}

func (s *EnableWebCCRequest) SetResourceGroupId(v string) *EnableWebCCRequest {
  s.ResourceGroupId = &v
  return s
}

func (s *EnableWebCCRequest) Validate() error {
  return dara.Validate(s)
}

