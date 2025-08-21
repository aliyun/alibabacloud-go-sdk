// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebAIProtectSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyWebAIProtectSwitchRequest
	GetConfig() *string
	SetDomain(v string) *ModifyWebAIProtectSwitchRequest
	GetDomain() *string
	SetResourceGroupId(v string) *ModifyWebAIProtectSwitchRequest
	GetResourceGroupId() *string
}

type ModifyWebAIProtectSwitchRequest struct {
	// The details of the Intelligent Protection policy. This parameter is a JSON string. The string contains the following fields:
	//
	// 	- **AiRuleEnable**: the status of the Intelligent Protection policy. This field is required and must be of the integer type. Valid values:
	//
	//     	- **0**: disabled
	//
	//     	- **1**: enabled
	//
	// This parameter is required.
	//
	// example:
	//
	// {"AiRuleEnable": 1}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for a domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
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

func (s ModifyWebAIProtectSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebAIProtectSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebAIProtectSwitchRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyWebAIProtectSwitchRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyWebAIProtectSwitchRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyWebAIProtectSwitchRequest) SetConfig(v string) *ModifyWebAIProtectSwitchRequest {
	s.Config = &v
	return s
}

func (s *ModifyWebAIProtectSwitchRequest) SetDomain(v string) *ModifyWebAIProtectSwitchRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebAIProtectSwitchRequest) SetResourceGroupId(v string) *ModifyWebAIProtectSwitchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebAIProtectSwitchRequest) Validate() error {
	return dara.Validate(s)
}
