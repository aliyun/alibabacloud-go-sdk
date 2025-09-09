// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceSecurityGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizedSecurityGroups(v []*string) *ConfigInstanceSecurityGroupsRequest
	GetAuthorizedSecurityGroups() []*string
	SetInstanceId(v string) *ConfigInstanceSecurityGroupsRequest
	GetInstanceId() *string
	SetLang(v string) *ConfigInstanceSecurityGroupsRequest
	GetLang() *string
	SetRegionId(v string) *ConfigInstanceSecurityGroupsRequest
	GetRegionId() *string
}

type ConfigInstanceSecurityGroupsRequest struct {
	// An array that consists of the IDs of authorized security groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp14u00sh39jvw5****
	AuthorizedSecurityGroups []*string `json:"AuthorizedSecurityGroups,omitempty" xml:"AuthorizedSecurityGroups,omitempty" type:"Repeated"`
	// The ID of the bastion host.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-78v1gh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID of the bastion host.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfigInstanceSecurityGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *ConfigInstanceSecurityGroupsRequest) GetAuthorizedSecurityGroups() []*string {
	return s.AuthorizedSecurityGroups
}

func (s *ConfigInstanceSecurityGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigInstanceSecurityGroupsRequest) GetLang() *string {
	return s.Lang
}

func (s *ConfigInstanceSecurityGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigInstanceSecurityGroupsRequest) SetAuthorizedSecurityGroups(v []*string) *ConfigInstanceSecurityGroupsRequest {
	s.AuthorizedSecurityGroups = v
	return s
}

func (s *ConfigInstanceSecurityGroupsRequest) SetInstanceId(v string) *ConfigInstanceSecurityGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceSecurityGroupsRequest) SetLang(v string) *ConfigInstanceSecurityGroupsRequest {
	s.Lang = &v
	return s
}

func (s *ConfigInstanceSecurityGroupsRequest) SetRegionId(v string) *ConfigInstanceSecurityGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigInstanceSecurityGroupsRequest) Validate() error {
	return dara.Validate(s)
}
