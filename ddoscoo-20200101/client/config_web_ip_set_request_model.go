// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigWebIpSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlackList(v []*string) *ConfigWebIpSetRequest
	GetBlackList() []*string
	SetDomain(v string) *ConfigWebIpSetRequest
	GetDomain() *string
	SetResourceGroupId(v string) *ConfigWebIpSetRequest
	GetResourceGroupId() *string
	SetWhiteList(v []*string) *ConfigWebIpSetRequest
	GetWhiteList() []*string
}

type ConfigWebIpSetRequest struct {
	// The IP addresses and CIDR blocks in the blacklist. You can add up to 200 IP addresses or CIDR blocks to the blacklist.
	//
	// example:
	//
	// 1.1.1.1
	BlackList []*string `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
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
	// The IP addresses and CIDR blocks in the whitelist. You can add up to 200 IP addresses or CIDR blocks to the whitelist.
	//
	// example:
	//
	// 2.2.2.2/24
	WhiteList []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s ConfigWebIpSetRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigWebIpSetRequest) GoString() string {
	return s.String()
}

func (s *ConfigWebIpSetRequest) GetBlackList() []*string {
	return s.BlackList
}

func (s *ConfigWebIpSetRequest) GetDomain() *string {
	return s.Domain
}

func (s *ConfigWebIpSetRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigWebIpSetRequest) GetWhiteList() []*string {
	return s.WhiteList
}

func (s *ConfigWebIpSetRequest) SetBlackList(v []*string) *ConfigWebIpSetRequest {
	s.BlackList = v
	return s
}

func (s *ConfigWebIpSetRequest) SetDomain(v string) *ConfigWebIpSetRequest {
	s.Domain = &v
	return s
}

func (s *ConfigWebIpSetRequest) SetResourceGroupId(v string) *ConfigWebIpSetRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigWebIpSetRequest) SetWhiteList(v []*string) *ConfigWebIpSetRequest {
	s.WhiteList = v
	return s
}

func (s *ConfigWebIpSetRequest) Validate() error {
	return dara.Validate(s)
}
