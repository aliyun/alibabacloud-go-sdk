// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalSecurityIPGroupNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalSecurityIPGroup(v []*ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) *ModifyGlobalSecurityIPGroupNameResponseBody
	GetGlobalSecurityIPGroup() []*ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup
	SetRequestId(v string) *ModifyGlobalSecurityIPGroupNameResponseBody
	GetRequestId() *string
}

type ModifyGlobalSecurityIPGroupNameResponseBody struct {
	// The details of the global IP whitelist template.
	GlobalSecurityIPGroup []*ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup `json:"GlobalSecurityIPGroup,omitempty" xml:"GlobalSecurityIPGroup,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGlobalSecurityIPGroupNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBody) GetGlobalSecurityIPGroup() []*ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup {
	return s.GlobalSecurityIPGroup
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBody) SetGlobalSecurityIPGroup(v []*ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) *ModifyGlobalSecurityIPGroupNameResponseBody {
	s.GlobalSecurityIPGroup = v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBody) SetRequestId(v string) *ModifyGlobalSecurityIPGroupNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup struct {
	// The IP address in the whitelist template.
	//
	// >  Separate multiple IP addresses with commas (,). You can add up to 1,000 IP addresses or CIDR blocks to all IP whitelists.
	//
	// example:
	//
	// 192.168.0.1
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// The name of the IP whitelist template. The name must meet the following requirements:
	//
	// 	- The name can contain lowercase letters, digits, and underscores (_).
	//
	// 	- The name must start with a letter and end with a letter or a digit.
	//
	// 	- The name must be 2 to 120 characters in length.
	//
	// example:
	//
	// test_123
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	// The ID of the IP whitelist template.
	//
	// example:
	//
	// g-zsldxfiwjmti0kcm****
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) GetGIpList() *string {
	return s.GIpList
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) SetGIpList(v string) *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup {
	s.GIpList = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) SetGlobalIgName(v string) *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup {
	s.GlobalIgName = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) SetGlobalSecurityGroupId(v string) *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) SetRegionId(v string) *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup {
	s.RegionId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameResponseBodyGlobalSecurityIPGroup) Validate() error {
	return dara.Validate(s)
}
