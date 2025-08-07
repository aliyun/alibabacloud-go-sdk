// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalSecurityIPGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalSecurityIPGroup(v []*ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) *ModifyGlobalSecurityIPGroupResponseBody
	GetGlobalSecurityIPGroup() []*ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup
	SetRequestId(v string) *ModifyGlobalSecurityIPGroupResponseBody
	GetRequestId() *string
}

type ModifyGlobalSecurityIPGroupResponseBody struct {
	// The details of the global IP whitelist template.
	GlobalSecurityIPGroup []*ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup `json:"GlobalSecurityIPGroup,omitempty" xml:"GlobalSecurityIPGroup,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGlobalSecurityIPGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupResponseBody) GetGlobalSecurityIPGroup() []*ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	return s.GlobalSecurityIPGroup
}

func (s *ModifyGlobalSecurityIPGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGlobalSecurityIPGroupResponseBody) SetGlobalSecurityIPGroup(v []*ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) *ModifyGlobalSecurityIPGroupResponseBody {
	s.GlobalSecurityIPGroup = v
	return s
}

func (s *ModifyGlobalSecurityIPGroupResponseBody) SetRequestId(v string) *ModifyGlobalSecurityIPGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup struct {
	// The clusters that are associated with the IP address whitelist template.
	DBInstances []*string `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
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

func (s ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetDBInstances() []*string {
	return s.DBInstances
}

func (s *ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetGIpList() *string {
	return s.GIpList
}

func (s *ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetDBInstances(v []*string) *ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.DBInstances = v
	return s
}

func (s *ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetGIpList(v string) *ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.GIpList = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetGlobalIgName(v string) *ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.GlobalIgName = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetGlobalSecurityGroupId(v string) *ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) SetRegionId(v string) *ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup {
	s.RegionId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupResponseBodyGlobalSecurityIPGroup) Validate() error {
	return dara.Validate(s)
}
