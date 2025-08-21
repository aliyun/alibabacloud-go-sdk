// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebIpSetSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyWebIpSetSwitchRequest
	GetConfig() *string
	SetDomain(v string) *ModifyWebIpSetSwitchRequest
	GetDomain() *string
	SetResourceGroupId(v string) *ModifyWebIpSetSwitchRequest
	GetResourceGroupId() *string
}

type ModifyWebIpSetSwitchRequest struct {
	// The details of the Blacklist/Whitelist (Domain Names) feature. This parameter is a JSON string. The value consists of the following fields:
	//
	// **bwlist_enable**: the status of the Blacklist/Whitelist (Domain Names) feature. This field is required and must be of the integer type. Valid values:
	//
	// 	- 0: turned off
	//
	// 	- 1: turned on
	//
	// This parameter is required.
	//
	// example:
	//
	// {"BwlistEnable":1}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
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

func (s ModifyWebIpSetSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebIpSetSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebIpSetSwitchRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyWebIpSetSwitchRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyWebIpSetSwitchRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyWebIpSetSwitchRequest) SetConfig(v string) *ModifyWebIpSetSwitchRequest {
	s.Config = &v
	return s
}

func (s *ModifyWebIpSetSwitchRequest) SetDomain(v string) *ModifyWebIpSetSwitchRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebIpSetSwitchRequest) SetResourceGroupId(v string) *ModifyWebIpSetSwitchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebIpSetSwitchRequest) Validate() error {
	return dara.Validate(s)
}
