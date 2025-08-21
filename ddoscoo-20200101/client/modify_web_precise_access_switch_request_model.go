// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebPreciseAccessSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyWebPreciseAccessSwitchRequest
	GetConfig() *string
	SetDomain(v string) *ModifyWebPreciseAccessSwitchRequest
	GetDomain() *string
	SetResourceGroupId(v string) *ModifyWebPreciseAccessSwitchRequest
	GetResourceGroupId() *string
}

type ModifyWebPreciseAccessSwitchRequest struct {
	// The configuration of the Accurate Access Control policy. This parameter is a JSON string. The string contains the following fields:
	//
	// 	- **PreciseRuleEnable**: the status of the Accurate Access Control policy. This field is required and must be of the INTEGER type. Valid values:
	//
	//     	- **0**: disables the policy.
	//
	//     	- **1**: enables the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"PreciseRuleEnable":0}
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

func (s ModifyWebPreciseAccessSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebPreciseAccessSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebPreciseAccessSwitchRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyWebPreciseAccessSwitchRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyWebPreciseAccessSwitchRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyWebPreciseAccessSwitchRequest) SetConfig(v string) *ModifyWebPreciseAccessSwitchRequest {
	s.Config = &v
	return s
}

func (s *ModifyWebPreciseAccessSwitchRequest) SetDomain(v string) *ModifyWebPreciseAccessSwitchRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebPreciseAccessSwitchRequest) SetResourceGroupId(v string) *ModifyWebPreciseAccessSwitchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebPreciseAccessSwitchRequest) Validate() error {
	return dara.Validate(s)
}
