// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebCCRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAct(v string) *CreateWebCCRuleRequest
	GetAct() *string
	SetCount(v int32) *CreateWebCCRuleRequest
	GetCount() *int32
	SetDomain(v string) *CreateWebCCRuleRequest
	GetDomain() *string
	SetInterval(v int32) *CreateWebCCRuleRequest
	GetInterval() *int32
	SetMode(v string) *CreateWebCCRuleRequest
	GetMode() *string
	SetName(v string) *CreateWebCCRuleRequest
	GetName() *string
	SetResourceGroupId(v string) *CreateWebCCRuleRequest
	GetResourceGroupId() *string
	SetTtl(v int32) *CreateWebCCRuleRequest
	GetTtl() *int32
	SetUri(v string) *CreateWebCCRuleRequest
	GetUri() *string
}

type CreateWebCCRuleRequest struct {
	// The action on the requests that trigger the custom frequency control rule. Valid values:
	//
	// 	- **close**: blocks the requests.
	//
	// 	- **captcha**: triggers Completely Automated Public Turing test to tell Computers and Humans Apart (CAPTCHA) verification for the requests.
	//
	// This parameter is required.
	//
	// example:
	//
	// close
	Act *string `json:"Act,omitempty" xml:"Act,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The domain name of the website.
	//
	// >  A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The matching mode. Valid values:
	//
	// 	- **prefix**: prefix match.
	//
	// 	- **match**: exact match.
	//
	// >  If the **URI*	- of the check path contains parameters, you must set this parameter to **prefix**.
	//
	// This parameter is required.
	//
	// example:
	//
	// prefix
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The name of the rule. The name can be up to 128 characters in length and contain letters, digits, and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// testrule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Proxy instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The blocking duration. Valid values: **60*	- to **86400**. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The check path.
	//
	// >  The URI cannot be modified. The domain name of the website, the check path, and the rule name uniquely identify a rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// /abc/a.php
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateWebCCRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWebCCRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateWebCCRuleRequest) GetAct() *string {
	return s.Act
}

func (s *CreateWebCCRuleRequest) GetCount() *int32 {
	return s.Count
}

func (s *CreateWebCCRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateWebCCRuleRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateWebCCRuleRequest) GetMode() *string {
	return s.Mode
}

func (s *CreateWebCCRuleRequest) GetName() *string {
	return s.Name
}

func (s *CreateWebCCRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateWebCCRuleRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *CreateWebCCRuleRequest) GetUri() *string {
	return s.Uri
}

func (s *CreateWebCCRuleRequest) SetAct(v string) *CreateWebCCRuleRequest {
	s.Act = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetCount(v int32) *CreateWebCCRuleRequest {
	s.Count = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetDomain(v string) *CreateWebCCRuleRequest {
	s.Domain = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetInterval(v int32) *CreateWebCCRuleRequest {
	s.Interval = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetMode(v string) *CreateWebCCRuleRequest {
	s.Mode = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetName(v string) *CreateWebCCRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetResourceGroupId(v string) *CreateWebCCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetTtl(v int32) *CreateWebCCRuleRequest {
	s.Ttl = &v
	return s
}

func (s *CreateWebCCRuleRequest) SetUri(v string) *CreateWebCCRuleRequest {
	s.Uri = &v
	return s
}

func (s *CreateWebCCRuleRequest) Validate() error {
	return dara.Validate(s)
}
