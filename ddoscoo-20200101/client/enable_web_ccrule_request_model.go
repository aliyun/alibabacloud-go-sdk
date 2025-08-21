// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWebCCRuleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDomain(v string) *EnableWebCCRuleRequest
  GetDomain() *string 
  SetResourceGroupId(v string) *EnableWebCCRuleRequest
  GetResourceGroupId() *string 
}

type EnableWebCCRuleRequest struct {
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

func (s EnableWebCCRuleRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableWebCCRuleRequest) GoString() string {
  return s.String()
}

func (s *EnableWebCCRuleRequest) GetDomain() *string  {
  return s.Domain
}

func (s *EnableWebCCRuleRequest) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *EnableWebCCRuleRequest) SetDomain(v string) *EnableWebCCRuleRequest {
  s.Domain = &v
  return s
}

func (s *EnableWebCCRuleRequest) SetResourceGroupId(v string) *EnableWebCCRuleRequest {
  s.ResourceGroupId = &v
  return s
}

func (s *EnableWebCCRuleRequest) Validate() error {
  return dara.Validate(s)
}

