// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebAreaBlockSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyWebAreaBlockSwitchRequest
	GetConfig() *string
	SetDomain(v string) *ModifyWebAreaBlockSwitchRequest
	GetDomain() *string
	SetResourceGroupId(v string) *ModifyWebAreaBlockSwitchRequest
	GetResourceGroupId() *string
}

type ModifyWebAreaBlockSwitchRequest struct {
	// Specifies whether to enable or disable the Location Blacklist (Domain Names) policy for a domain name. The value is a string that consists of a JSON struct. The JSON struct contains the following parameters:
	//
	// 	- **RegionblockEnable**: the status of the Location Blacklist (Domain Names) policy. This parameter is required and must be of the INTEGER type. Valid values:
	//
	//     	- **1**: enables the policy.
	//
	//     	- **0**: disables the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"RegionblockEnable": 1}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The domain name for which you want to enable or disable the Location Blacklist policy.
	//
	// > You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all the domain names that are added to Anti-DDoS Pro or Anti-DDoS Premium.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not configure this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyWebAreaBlockSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebAreaBlockSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebAreaBlockSwitchRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyWebAreaBlockSwitchRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyWebAreaBlockSwitchRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyWebAreaBlockSwitchRequest) SetConfig(v string) *ModifyWebAreaBlockSwitchRequest {
	s.Config = &v
	return s
}

func (s *ModifyWebAreaBlockSwitchRequest) SetDomain(v string) *ModifyWebAreaBlockSwitchRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebAreaBlockSwitchRequest) SetResourceGroupId(v string) *ModifyWebAreaBlockSwitchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebAreaBlockSwitchRequest) Validate() error {
	return dara.Validate(s)
}
