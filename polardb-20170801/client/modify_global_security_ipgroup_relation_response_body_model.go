// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalSecurityIPGroupRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyGlobalSecurityIPGroupRelationResponseBody
	GetDBClusterId() *string
	SetGlobalSecurityIPGroupRel(v []*ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) *ModifyGlobalSecurityIPGroupRelationResponseBody
	GetGlobalSecurityIPGroupRel() []*ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel
	SetRequestId(v string) *ModifyGlobalSecurityIPGroupRelationResponseBody
	GetRequestId() *string
}

type ModifyGlobalSecurityIPGroupRelationResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Details of the global IP address whitelist template.
	GlobalSecurityIPGroupRel []*ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel `json:"GlobalSecurityIPGroupRel,omitempty" xml:"GlobalSecurityIPGroupRel,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGlobalSecurityIPGroupRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBody) GetGlobalSecurityIPGroupRel() []*ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel {
	return s.GlobalSecurityIPGroupRel
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBody) SetDBClusterId(v string) *ModifyGlobalSecurityIPGroupRelationResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBody) SetGlobalSecurityIPGroupRel(v []*ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) *ModifyGlobalSecurityIPGroupRelationResponseBody {
	s.GlobalSecurityIPGroupRel = v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBody) SetRequestId(v string) *ModifyGlobalSecurityIPGroupRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBody) Validate() error {
	if s.GlobalSecurityIPGroupRel != nil {
		for _, item := range s.GlobalSecurityIPGroupRel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel struct {
	// The IP addresses in the whitelist template.
	//
	// > Separate multiple IP addresses with a comma. A maximum of 1,000 IP addresses or CIDR blocks can be added to all IP address whitelists.
	//
	// example:
	//
	// 192.168.0.1
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// The name of the IP address whitelist template. The name must meet the following requirements:
	//
	// - Consists of lowercase letters, digits, and underscores (_).
	//
	// - Starts with a letter and ends with a letter or a digit.
	//
	// - Is 2 to 120 characters in length.
	//
	// example:
	//
	// test_123
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	// The ID of the IP address whitelist template.
	//
	// example:
	//
	// g-zsldxfiwjmti0kcm****
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) GetGIpList() *string {
	return s.GIpList
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) SetGIpList(v string) *ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel {
	s.GIpList = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) SetGlobalIgName(v string) *ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel {
	s.GlobalIgName = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) SetGlobalSecurityGroupId(v string) *ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) SetRegionId(v string) *ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel {
	s.RegionId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) Validate() error {
	return dara.Validate(s)
}
