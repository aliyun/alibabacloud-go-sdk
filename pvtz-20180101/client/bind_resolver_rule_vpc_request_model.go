// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindResolverRuleVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *BindResolverRuleVpcRequest
	GetLang() *string
	SetRuleId(v string) *BindResolverRuleVpcRequest
	GetRuleId() *string
	SetVpc(v []*BindResolverRuleVpcRequestVpc) *BindResolverRuleVpcRequest
	GetVpc() []*BindResolverRuleVpcRequestVpc
}

type BindResolverRuleVpcRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the forwarding rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// hr****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The VPCs that you want to associate with the forwarding rule.
	Vpc []*BindResolverRuleVpcRequestVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
}

func (s BindResolverRuleVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s BindResolverRuleVpcRequest) GoString() string {
	return s.String()
}

func (s *BindResolverRuleVpcRequest) GetLang() *string {
	return s.Lang
}

func (s *BindResolverRuleVpcRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *BindResolverRuleVpcRequest) GetVpc() []*BindResolverRuleVpcRequestVpc {
	return s.Vpc
}

func (s *BindResolverRuleVpcRequest) SetLang(v string) *BindResolverRuleVpcRequest {
	s.Lang = &v
	return s
}

func (s *BindResolverRuleVpcRequest) SetRuleId(v string) *BindResolverRuleVpcRequest {
	s.RuleId = &v
	return s
}

func (s *BindResolverRuleVpcRequest) SetVpc(v []*BindResolverRuleVpcRequestVpc) *BindResolverRuleVpcRequest {
	s.Vpc = v
	return s
}

func (s *BindResolverRuleVpcRequest) Validate() error {
	if s.Vpc != nil {
		for _, item := range s.Vpc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BindResolverRuleVpcRequestVpc struct {
	// The region ID of the outbound VPC.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-f8zvrvr1payllgz38****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC type. Valid values:
	//
	// 	- STANDARD: standard VPC
	//
	// 	- EDS: Elastic Desktop Service (EDS) workspace VPC
	//
	// example:
	//
	// STANDARD
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s BindResolverRuleVpcRequestVpc) String() string {
	return dara.Prettify(s)
}

func (s BindResolverRuleVpcRequestVpc) GoString() string {
	return s.String()
}

func (s *BindResolverRuleVpcRequestVpc) GetRegionId() *string {
	return s.RegionId
}

func (s *BindResolverRuleVpcRequestVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *BindResolverRuleVpcRequestVpc) GetVpcType() *string {
	return s.VpcType
}

func (s *BindResolverRuleVpcRequestVpc) SetRegionId(v string) *BindResolverRuleVpcRequestVpc {
	s.RegionId = &v
	return s
}

func (s *BindResolverRuleVpcRequestVpc) SetVpcId(v string) *BindResolverRuleVpcRequestVpc {
	s.VpcId = &v
	return s
}

func (s *BindResolverRuleVpcRequestVpc) SetVpcType(v string) *BindResolverRuleVpcRequestVpc {
	s.VpcType = &v
	return s
}

func (s *BindResolverRuleVpcRequestVpc) Validate() error {
	return dara.Validate(s)
}
