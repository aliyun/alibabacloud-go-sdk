// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyIpWhitelistRequest
	GetClusterId() *string
	SetGroupName(v string) *ModifyIpWhitelistRequest
	GetGroupName() *string
	SetIpList(v string) *ModifyIpWhitelistRequest
	GetIpList() *string
	SetIpVersion(v string) *ModifyIpWhitelistRequest
	GetIpVersion() *string
}

type ModifyIpWhitelistRequest struct {
	// The ID of target instance. You can call [DescribeInstances](https://help.aliyun.com/document_detail/144595.html) to obtain target instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1uoihlf82e8****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the whitelist group for the target instance. You can invoke [DescribeIpWhitelist](https://help.aliyun.com/document_detail/144606.html) to obtain the whitelist group name.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_01
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The IP addresses in the whitelist group after modification. Separate multiple IP addresses with commas (,).
	//
	// example:
	//
	// 42.120.XX.XX
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The version of the IP address. Valid values:
	//
	// - **4**: IPv4.
	//
	// - **6**: IPv6.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
}

func (s ModifyIpWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpWhitelistRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyIpWhitelistRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyIpWhitelistRequest) GetIpList() *string {
	return s.IpList
}

func (s *ModifyIpWhitelistRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ModifyIpWhitelistRequest) SetClusterId(v string) *ModifyIpWhitelistRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyIpWhitelistRequest) SetGroupName(v string) *ModifyIpWhitelistRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyIpWhitelistRequest) SetIpList(v string) *ModifyIpWhitelistRequest {
	s.IpList = &v
	return s
}

func (s *ModifyIpWhitelistRequest) SetIpVersion(v string) *ModifyIpWhitelistRequest {
	s.IpVersion = &v
	return s
}

func (s *ModifyIpWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
