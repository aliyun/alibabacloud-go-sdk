// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebCacheSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ModifyWebCacheSwitchRequest
	GetDomain() *string
	SetEnable(v int32) *ModifyWebCacheSwitchRequest
	GetEnable() *int32
	SetResourceGroupId(v string) *ModifyWebCacheSwitchRequest
	GetResourceGroupId() *string
}

type ModifyWebCacheSwitchRequest struct {
	// The domain name for which you want to configure the Static Page Caching policy.
	//
	// > You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all the domain names that are added to Anti-DDoS Pro or Anti-DDoS Premium.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to enable or disable the Static Page Caching policy for a website. Valid values:
	//
	// 	- **1**: enables the policy.
	//
	// 	- **0**: disables the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not configure this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyWebCacheSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebCacheSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheSwitchRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyWebCacheSwitchRequest) GetEnable() *int32 {
	return s.Enable
}

func (s *ModifyWebCacheSwitchRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyWebCacheSwitchRequest) SetDomain(v string) *ModifyWebCacheSwitchRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebCacheSwitchRequest) SetEnable(v int32) *ModifyWebCacheSwitchRequest {
	s.Enable = &v
	return s
}

func (s *ModifyWebCacheSwitchRequest) SetResourceGroupId(v string) *ModifyWebCacheSwitchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebCacheSwitchRequest) Validate() error {
	return dara.Validate(s)
}
