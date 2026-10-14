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
	// The status of the Location Blacklist feature (for domain names). This parameter is a string that is converted from a JSON struct. The JSON struct contains the following parameter:
	//
	// - **RegionblockEnable**: An Integer value that is required. The status of the Location Blacklist feature (for domain names). Valid values:
	//
	//   - **1**: enables the feature.
	//
	//   - **0**: disables the feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"RegionblockEnable": 1}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The domain name for which you want to set the Location Blacklist feature.
	//
	// > You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names that are protected by Anti-DDoS Proxy.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Proxy instance belongs in Resource Management.
	//
	// If you do not set this parameter, the default resource group is used.
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
