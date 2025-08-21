// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebCCRuleV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DeleteWebCCRuleV2Request
	GetDomain() *string
	SetOwner(v string) *DeleteWebCCRuleV2Request
	GetOwner() *string
	SetRuleNames(v string) *DeleteWebCCRuleV2Request
	GetRuleNames() *string
}

type DeleteWebCCRuleV2Request struct {
	// The domain name of the website.
	//
	// >  A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The source of the rule. Valid values:
	//
	// 	- **manual*	- (default): manually created.
	//
	// 	- **clover**: automatically created. Specify this value when you want to delete intelligent protection rules.
	//
	// example:
	//
	// manual
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The names of the rules that you want to delete.
	//
	// example:
	//
	// [\\"trdsss\\"]
	RuleNames *string `json:"RuleNames,omitempty" xml:"RuleNames,omitempty"`
}

func (s DeleteWebCCRuleV2Request) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebCCRuleV2Request) GoString() string {
	return s.String()
}

func (s *DeleteWebCCRuleV2Request) GetDomain() *string {
	return s.Domain
}

func (s *DeleteWebCCRuleV2Request) GetOwner() *string {
	return s.Owner
}

func (s *DeleteWebCCRuleV2Request) GetRuleNames() *string {
	return s.RuleNames
}

func (s *DeleteWebCCRuleV2Request) SetDomain(v string) *DeleteWebCCRuleV2Request {
	s.Domain = &v
	return s
}

func (s *DeleteWebCCRuleV2Request) SetOwner(v string) *DeleteWebCCRuleV2Request {
	s.Owner = &v
	return s
}

func (s *DeleteWebCCRuleV2Request) SetRuleNames(v string) *DeleteWebCCRuleV2Request {
	s.RuleNames = &v
	return s
}

func (s *DeleteWebCCRuleV2Request) Validate() error {
	return dara.Validate(s)
}
