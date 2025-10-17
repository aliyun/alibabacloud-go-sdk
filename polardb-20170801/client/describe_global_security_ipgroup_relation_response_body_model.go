// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalSecurityIPGroupRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeGlobalSecurityIPGroupRelationResponseBody
	GetDBClusterId() *string
	SetGlobalSecurityIPGroupRel(v []*DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) *DescribeGlobalSecurityIPGroupRelationResponseBody
	GetGlobalSecurityIPGroupRel() []*DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel
	SetRequestId(v string) *DescribeGlobalSecurityIPGroupRelationResponseBody
	GetRequestId() *string
}

type DescribeGlobalSecurityIPGroupRelationResponseBody struct {
	// The ID of cluster.
	//
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The details of the global IP whitelist template.
	GlobalSecurityIPGroupRel []*DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel `json:"GlobalSecurityIPGroupRel,omitempty" xml:"GlobalSecurityIPGroupRel,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGlobalSecurityIPGroupRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalSecurityIPGroupRelationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBody) GetGlobalSecurityIPGroupRel() []*DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel {
	return s.GlobalSecurityIPGroupRel
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBody) SetDBClusterId(v string) *DescribeGlobalSecurityIPGroupRelationResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBody) SetGlobalSecurityIPGroupRel(v []*DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) *DescribeGlobalSecurityIPGroupRelationResponseBody {
	s.GlobalSecurityIPGroupRel = v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBody) SetRequestId(v string) *DescribeGlobalSecurityIPGroupRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBody) Validate() error {
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

type DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel struct {
	// The IP address in the global IP whitelist template.
	//
	// >  Separate multiple IP addresses with commas (,). You can add up to 1,000 IP addresses or CIDR blocks to all IP whitelists.
	//
	// example:
	//
	// 192.168.0.1
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// The name of the global IP whitelist template. The name must meet the following requirements:
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
	// The ID of the global IP whitelist template.
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

func (s DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) GoString() string {
	return s.String()
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) GetGIpList() *string {
	return s.GIpList
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) SetGIpList(v string) *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel {
	s.GIpList = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) SetGlobalIgName(v string) *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel {
	s.GlobalIgName = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) SetGlobalSecurityGroupId(v string) *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) SetRegionId(v string) *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationResponseBodyGlobalSecurityIPGroupRel) Validate() error {
	return dara.Validate(s)
}
