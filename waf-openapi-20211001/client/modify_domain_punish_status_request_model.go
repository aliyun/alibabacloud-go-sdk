// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDomainPunishStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ModifyDomainPunishStatusRequest
	GetDomain() *string
	SetInstanceId(v string) *ModifyDomainPunishStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyDomainPunishStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyDomainPunishStatusRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyDomainPunishStatusRequest struct {
	// The domain name that is penalized for failing to obtain an ICP filing.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.xxxxaliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-zxu****0g02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: Outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s ModifyDomainPunishStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainPunishStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainPunishStatusRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyDomainPunishStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDomainPunishStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDomainPunishStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyDomainPunishStatusRequest) SetDomain(v string) *ModifyDomainPunishStatusRequest {
	s.Domain = &v
	return s
}

func (s *ModifyDomainPunishStatusRequest) SetInstanceId(v string) *ModifyDomainPunishStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDomainPunishStatusRequest) SetRegionId(v string) *ModifyDomainPunishStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDomainPunishStatusRequest) SetResourceManagerResourceGroupId(v string) *ModifyDomainPunishStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDomainPunishStatusRequest) Validate() error {
	return dara.Validate(s)
}
